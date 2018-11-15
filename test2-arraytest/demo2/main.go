package main

import (
	"fmt"
)

const OUTER_COUNT = 4
const INNER_COUNT = 8

func main() {
	var s [][]string
	for i := 0; i < OUTER_COUNT; i++ {
		sl := make([]string, 0, INNER_COUNT)
		for j := 0; j < INNER_COUNT; j++ {
			sl = append(sl, "a")
		}
		s = append(s, sl)
	}
	fmt.Println(s)

	arr := []string{"abcfile", "abdfile", "abefile", "agefile", "aboutfile", "ashatfile", "ashotfile"}
	fmt.Println(arr)
	fmt.Println(splitSlice(arr))
}

//切片分组
func splitSlice(arr []string) [][]string {
	returnData := make([][]string, 0)
	i := 0
	var j int
	for {
		if i >= len(arr) {
			break
		}

		for j = i + 1; j < len(arr); j++ {
		}

		returnData = append(returnData, arr[i:j])
		i = j
	}
	return returnData
}

/*
//切片分组
func splitSlice(arr []string, n int, front bool) [][]string {
	returnData := make([][]string, 0)
	i := 0
	var j int
	for {
		if i >= len(arr) {
			break
		}
		if front { //前n个字符分组
			for j = i + 1; j < len(arr) && arr[i][0:n-1] == arr[j][0:n-1]; j++ {
			}

		} else { //后n个字符分组
			for j = i + 1; j < len(arr) && arr[i][len(arr[i])-n:] == arr[j][len(arr[j])-n:]; j++ {
			}
		}

		returnData = append(returnData, arr[i:j])
		i = j
	}
	return returnData
}
*/
