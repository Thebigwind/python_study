package main

import (
	"fmt"
	"regexp"
)

type Group struct {
	Files []string
}

func main() {
	pattern := "ab"
	arr := []string{"abcfile", "abdfile", "abefile", "agefile", "aboutfile", "ashatfile", "ashotfile"}

	reg := regexp.MustCompile(pattern)
	numGroup := make([]Group, 0, len(arr))

	//calcute group num
	for _, e := range arr {
		groupName := reg.FindAllString(e, -1)
		fmt.Println(groupName)
		group := Group{
			Files: []string{},
		}

		bools := arrContain(numGroup, group)
		if bools == false {
			numGroup = append(numGroup, group)
		}
	}

}

func arrContain(arr []Group, value Group) bool {
	if len(arr) == 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if len(arr[i].Files) == 0 && len(value.Files) == 0 {
			return true
		}
		for j := 0; j < len(arr[j].Files); j++ {
			if arr[i].Files[j] != value.Files[j] {
				return false
			}
		}

	}
	return true
}
