package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main() {

}

func execTest(s string) (error, string) {
	cmd := exec.Command("/bin/bash", "-c", s)

	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	return err, strings.Trim(out.String(), "\n")
}
