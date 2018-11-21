package main

import "fmt"

func main() {
	ages1 := map[string]int {
		"destiny": 24,
		"camery": 25,
		"freedom": 26,
		"justice": 27,
	}
	
	ages2 := make(map[string]int)
	ages2["destiny"] = 24
	ages2["camery"] = 25
	ages2["freedom"] = 26
	ages2["justice"] = 27
	
	fmt.Println("ages1")
	for k, v := range ages1 {
		fmt.Printf("%s\t%d\n", k, v)
	}
	// ages1
	// justice	27
	// destiny	24
	// camery	25
	// freedom	26
	
	fmt.Println("ages2")
	for k, v := range ages2 {
		fmt.Printf("%s\t%d\n", k, v)
	}
	//ages2
	//camery	25
	//freedom	26
	//justice	27
	//destiny	24
}
