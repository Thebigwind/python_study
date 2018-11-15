package main

import (
	"fmt"
	"sort"
)

func main() {
	//定义一个map并初始化，根据key获取value
	colors := map[string]string{
		"bird":  "blue",
		"snake": "green",
		"cat":   "black",
	}
	// Get color of snake.
	c := colors["snake"]

	// Display string.
	fmt.Println(c)

	//通过==进行map赋值
	colors["rabbit"] = "white"
	fmt.Println(colors)

	names := map[int]string{}
	// Add three pairs to the map in separate statements.
	names[990] = "file.txt"
	names[1009] = "data.xls"
	names[1209] = "image.jpg"

	// There are three pairs in the map.
	fmt.Println(len(names))

	//通过delete删除
	stus := map[int]string{}
	stus[0] = "zhangsan"
	stus[1] = "lisi"
	stus[2] = "wang5"
	stus[3] = "ma6"
	fmt.Println(stus)

	delete(stus, 2)
	fmt.Println(stus)

	//map的range遍历
	animals := map[string]string{}
	animals["cat"] = "Mitty"
	animals["dog"] = "hashi"
	//look over the map
	for key, value := range animals {
		fmt.Println(key, "=", value)
	}

	//获取map中所有的key
	sizes := map[string]int{
		"XL": 20,
		"L":  10,
		"M":  5,
	}

	// Loop over map and append keys to empty slice.
	keys := []string{}
	for key, _ := range sizes {
		keys = append(keys, key)
	}

	// This is a slice of the keys.
	fmt.Println(keys)

	//通过make进行构建
	lookup := make(map[string]int, 200)

	// Use the new map.
	lookup["cat"] = 10
	result := lookup["cat"]
	fmt.Println(result)

	//map作为函数参数
	// This map has two string keys.
	colors2 := map[string]int{
		"blue":  10,
		"green": 20,
	}
	// Pass map to func.
	PrintGreen(colors2)

	//按key进行排序
	// To create a map as input
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	// To store the keys in slice in sorted order
	var keys2 []int
	for k := range m {
		fmt.Println(k) //每次都随机输出key的值
		keys2 = append(keys2, k)
	}
	sort.Ints(keys2)

	// To perform the opertion you want
	for _, k := range keys2 {
		fmt.Println("Key:", k, "Value:", m[k])
	}

	//panicTest
	panicTest()

}

func PrintGreen(colors map[string]int) {
	// Handle map argument.
	fmt.Println(colors["green"])
}

func panicTest() { //向一个值为nil的字典上添加键值对
	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n",
		key, elem, ok)

	fmt.Printf("The length of nil map: %d\n",
		len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n",
		key)
	delete(m, key)

	elem = 2
	fmt.Println("Add a key-element pair to a nil map...")
	m["two"] = elem // 这里会引发panic。
}
