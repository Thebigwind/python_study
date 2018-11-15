package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{"Tom", 20}
	p2 := Person{"Lily", 21}
	p3 := Person{"Linda", 23}
	p4 := Person{"Jass", 25}
	p5 := Person{"Tonny", 20}
	p6 := Person{"Pite", 25}
	p7 := Person{"Paul", 21}
	p8 := Person{"Kriss", 27}
	p9 := Person{"Jake", 23}
	p10 := Person{"Rose", 20}

	personList := []Person{}
	personList = append(personList, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10)
	fmt.Println("分组前的数据:", personList)
	fmt.Println("分组后的数据:", splitSlice(personList))
}

//按某个字段排序
type sortByAge []Person

func (s sortByAge) Len() int           { return len(s) }
func (s sortByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByAge) Less(i, j int) bool { return s[i].Age < s[j].Age }

//切片分组
func splitSlice(list []Person) [][]Person {
	sort.Sort(sortByAge(list))
	returnData := make([][]Person, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		for j = i + 1; j < len(list) && list[i].Age == list[j].Age; j++ {
		}

		returnData = append(returnData, list[i:j])
		i = j
	}
	return returnData
}
