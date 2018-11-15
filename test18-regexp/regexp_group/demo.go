package main

import (
	//"fmt"
	"regexp"
)

func main() {
	pattern := ""
	arr := []string{"abcfile", "abdfile", "abefile", "agefile", "aboutfile", "ashatfile", "ashotfile"}

	reg := regexp.MustCompile(pattern)

	//var str2arr [][]string
	//numG := make([]string, 0, len(arr))
	var numG []string
	//calcute group num
	for _, e := range arr {
		//groupName := reg.FindAllString(e, -1)
		groupName := reg.FindString(e)

		bools := arrContain(numG, groupName)
		if bools == false {
			numG = append(numG, groupName)
		}

	}

}

func arrContain(arr []string, value string) bool {
	if len(arr) == 0 {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return true
		}
	}
	return false
}
