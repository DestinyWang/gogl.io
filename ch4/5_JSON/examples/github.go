package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount  int         `json:"total_count"`
	Items       []*Issue
}

type Issue struct {
	Number      int
	HTMLURL     string      `json:"html_url"`
	Title       string
	State       string
	User        *User
	CreatedAt   time.Time   `json:"created_at"`
	Body        string      // Markdown 格式
}

type User struct {
	Login       string
	HTMLURL     string      `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	s := strings.Join(terms, " ")   // 将参数列表用空格进行拼接 (repo:golang/go is:open json decoder)
	q := url.QueryEscape(s)             // 对 s 进行转码使之可以安全的用在 URL 查询中 (repo%3Agolang%2Fgo+is%3Aopen+json+decoder)
	resp, err := http.Get(IssuesURL + "?q=" + q)    // https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
	if err != nil {
		return nil, err
	}
	// 必须在所有可能的分支上关闭 resp.Body
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {  // json.NewDecoder 使用流式解码器, 可以利用它来依次从字节流中解码出多个 JSON 实体
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	// issues 编号 提出者 标题
	// #11046     kurin encoding/json: Decoder internally buffers full input
	// #22369     Splik encoding/json: add the full path to the field in Unmars
	// #28189     adnsv encoding/json: confusing errors when unmarshaling custo
	// #28143 Carpetsmo proposal: encoding/json: add "readonly" tag
	// #16212 josharian encoding/json: do all reflect work before decoding
	// #26946    deuill encoding/json: clarify what happens when unmarshaling i
	// #12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
	// #14750 cyberphon encoding/json: parser ignores the case of member names
	// #27179  lavalamp encoding/json: no way to preserve the order of map keys
	// #5901        rsc encoding/json: allow override type marshaling
	// #22752  buyology proposal: encoding/json: add access to the underlying d
	// #7872  extempora encoding/json: Encoder internally buffers full output
	// #20754       rsc encoding/xml: unmarshal only processes first XML elemen
	// #7213  davechene cmd/compile: escape analysis oddity
	// #17609 nathanjsw encoding/json: ambiguous fields are marshalled
	// #20528  jvshahid net/http: connection reuse does not work happily with n
	// #22816 ganelon13 encoding/json: include field name in unmarshal error me
	// #21092  trotha01 encoding/json: unmarshal into slice reuses element data
	// #28578     cznic cmd/vendor/golang.org/x/arch/arm/arm64asm: TestObjdumpA
	// #15808 randall77 cmd/compile: loads/constants not lifted out of loop
	// #20206 markdryan encoding/base64: encoding is slow
	// #25426 josharian cmd/compile: revisit statement boundaries CL peformance
	// #19348 davidlaza cmd/compile: enable mid-stack inlining
	// #19109  bradfitz proposal: cmd/go: make fuzzing a first class citizen, l
	// #17244       adg proposal: decide policy for sub-repositories
	
}
