package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	// Year 和 Color 后面的字符串字面量是成员的标签
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color, omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
}

func main() {
	data, err := json.Marshal(movies)   // Marshal 生成了一个字节 slice, 其中包含一个不带有任何多余空白字符的很长的字符串
	data1, _ := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)    // [{"Title":"Casablanca","released":1942,"color":false,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]}]
	fmt.Printf("%s\n", data1)
	//  [
	//      {
	//  	    "Title": "Casablanca",
	//  	    "released": 1942,
	//  	    "color": false,
	//  	    "Actors": [
	//  	        "Humphrey Bogart",
	//  		    "Ingrid Bergman"
	//  	    ]
	//      },
	//      {
	//  	    "Title": "Cool Hand Luke",
	//  	    "released": 1967,
	//  	    "color": true,
	//  	    "Actors": [
	//  	        "Paul Newman"
	//  	    ]
	//      }
	//  ]
}
