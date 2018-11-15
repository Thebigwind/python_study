package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	/*
		f, err := os.Stat("/home")
		if err != nil {
			fmt.Println("err")
		}
		fmt.Println(f.Name)
	*/
	str := "abcde"
	fmt.Println("lenth:", strings.Count(str, "")-1)

	// open input file
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	// read the whole file at once
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// write the whole body at once
	err = ioutil.WriteFile("output.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}
