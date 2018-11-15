package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Lable struct {
	Files []string
}

func main() {
	pattern := "a([abc])"
	arr := []string{"abcfile", "abdfile", "abefile", "agefile", "aboutfile", "ashatfile", "ashotfile"}

	reg := regexp.MustCompile(pattern)

	//groupMap := make(map[string][]map[string][]string)
	groupMap := make(map[string][]string)
	labelMap := make(map[string]Lable)

	//calcute group num
	for _, e := range arr {
		groupName := reg.FindString(e)
		if groupName == "" {
			groupName = "no_match"
		}
		fmt.Println("groupName:", groupName)

		//groupMap[groupName]["Files"] = e
		groupMap[groupName] = append(groupMap[groupName], e)

	}
	fmt.Println(groupMap)

	//遍历map，
	for k, _ := range groupMap {
		label := Lable{
			Files: groupMap[k],
		}
		labelMap[k] = label

	}

	result, err := json.MarshalIndent(labelMap, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(result))
}
