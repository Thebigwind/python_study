package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func main() {
	pattern := "a([abc])"
	arr := []string{"abcfile", "abdfile", "abefile", "agefile", "aboutfile", "ashatfile", "ashotfile"}

	reg := regexp.MustCompile(pattern)

	groupMap := make(map[string][]string)

	//calcute group num
	for _, e := range arr {
		//fmt.Println("e:", e)
		groupName := reg.FindString(e)
		if groupName == "" {
			groupName = "no_match"
		}
		fmt.Println("groupName:", groupName)
		//numG = append(numG, groupName)
		groupMap[groupName] = append(groupMap[groupName], e)

	}
	fmt.Println(groupMap)

	result, err := json.MarshalIndent(groupMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}
