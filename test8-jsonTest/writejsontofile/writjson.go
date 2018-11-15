package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	test()
	fmt.Println("-------------")

	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatalln(err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(out)
	out.WriteTo(os.Stdout)
}

func test() {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	result, err := json.MarshalIndent(roads, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	//write to file .
	//if not provide arg, default current dir.
	file, err := os.OpenFile("user.json", os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	content := []byte(string(result))
	newWriter := bufio.NewWriterSize(file, 1024)

	_, err = newWriter.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = newWriter.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write to user.json successful")
}
