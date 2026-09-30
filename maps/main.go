package main

import (
	"fmt"
	"slices"
	"strings"
)

func countDistinctWords(messages []string) int {
	distinct:=[]string{}
	for _, msg:= range messages{
		slice := strings.Split(strings.ToLower(msg), " ")

		if isMember := slices.Contains(distinct, slice[0]); !isMember{
		distinct = append(distinct,slice[0] )
		}

		if isMember := slices.Contains(distinct, slice[1]); !isMember{
		distinct = append(distinct,slice[1] )
		}

		continue;
	}
	return len(distinct)
}


func main() {
	ages := make(map[string]int)
	ages["kemi"] = 23
	ages["remi"] = 29

	fmt.Println(ages)
	fmt.Println(ages["kemi"])

	//using literals
	newAges:= map[string]int{
		"kemi": 45,
		"remi": 23,
		"bosun": 29,
	}

	fmt.Println(newAges)

	//to delete
	delete(newAges, "kemi")
	fmt.Println(newAges)

	//to check if a key exists
	elem, ok:= newAges["kemi"]
	fmt.Println(elem, ok)


	names := map[string]int{}
missingNames := []string{}

if _, ok := names["Denna"]; !ok {
    // if the key doesn't exist yet,
    // append the name to the missingNames slice
    missingNames = append(missingNames, "Denna")
}

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)
fmt.Println(count)
}