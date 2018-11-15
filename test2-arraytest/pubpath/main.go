//[@:/vol1/site-pacakges-test/tablib/packages/odf3/text.py @:/vol1/site-pacakges-test/tablib/packages/odf/text.py @:/vol1/site-pacakges-test/tablib/packages/odf/text.pyc @:/vol1/site-pacakges-test/tablib/packages/odf3/text.pyc @:/vol1/site-pacakges-test/google/protobuf/text_encoding.py @:/vol1/site-pacakges-test/google/protobuf/text_encoding.pyc @:/vol1/site-pacakges-test/google/protobuf/text_format.py @:/vol1/site-pacakges-test/google/protobuf/text_format.pyc]

package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
1 2 3 4
1 2 3 9
1 2 5 7
1 2 6 7
*/

func test() {
	arr := [][]int{{1, 2, 3, 4}, {1, 2, 3, 9}, {1, 2, 5, 7}, {1, 2, 6, 7}}
	//fmt.Println(arr)

	pubPath := 0
	for j := 0; j < len(arr[0]); j++ { //列数
		fmt.Println("j:", j)

		for i := 0; i < len(arr); i++ { //行数
			//
			if arr[0][j] != arr[i][j] {
				pubPath = j - 1
				i = len(arr)
				j = len(arr[0])
			}
			//fmt.Println(arr[i][j] != arr[0][j])

		}
	}
	fmt.Println(pubPath)

}

func main() {
	//test()
	//test2()
	fmt.Println(test3())

	str := "abcdef"
	fmt.Println(str[0:1])

	fmt.Println(maxDepth(5))
}

//计算最大公共路径
func test2() {
	fsSearchs := []string{"a/b/c/d/e.a", "a/b/c/d/f/fs.a", "a/b/c/d/l/jw.q", "a/b/c/d/we/wee.r", "a/b/c/d/g.w", "a/b/c/d/e/efs/d.sd"}
	var fileArr [][]string
	//relateNameArr := make([]string, counter)

	for i, _ := range fsSearchs {
		//relateNameArr[i] = e.RelateName
		pathlList := strings.Split(fsSearchs[i], "/") //e.RelateName
		s1 := make([]string, 0, len(pathlList))
		for j := 0; j < len(pathlList); j++ {
			s1 = append(s1, pathlList[j])
		}
		fileArr = append(fileArr, s1)

	}
	fmt.Println(fileArr)
	//calcute the public path
	pubPathNum := 0
	fmt.Println(len(fileArr))
	if len(fileArr) > 1 {
		for col := 0; col < len(fileArr[0]); col++ { //column num
			for row := 0; row < len(fileArr); row++ { //row num
				if fileArr[0][col] != fileArr[row][col] { //compare if equal to
					pubPathNum = col - 1  //when not equal,the public is the front column
					row = len(fileArr)    //end row for
					col = len(fileArr[0]) //end column for
				}
			}
		}
	} else {
		pubPathNum = len(fileArr[0]) - 2
	}
	fmt.Println(pubPathNum)

	//join the pubPath
	pubPath := ""
	for i := 0; i <= pubPathNum; i++ {
		pubPath = pubPath + fileArr[0][i] + "/"
	}
	fmt.Println("Input_Dir:", pubPath)
}

//计算最大公共路径
func test3() string {
	//fsSearchs := []string{"a/z/sd/sd", "a/e/c", "a/e/c/a/a/.q", "a/e/c/d/e.a", "a/e/c/d/f/fs.a", "a/e/c/d/l/jw.q", "a/e/c/d/we/wee.r", "a/e/c/d/g.w", "a/e/c/d/e/efs/d.sd", "a/c/z", "a/d"}
	fsSearchs := []string{"a/e/sd/sd", "a/e/c", "a/e/c/a/a/.q", "a/e/c/d/e.a", "a/e/c/d/f/fs.a", "a/e/c/d/l/jw.q", "a/e/c/d/we/wee.r", "a/e/c/d/g.w", "a/e/c/d/e/efs/d.sd", "a/e/z", "a/e"}

	sort.Sort(sort.StringSlice(fsSearchs))
	minLen := len(fsSearchs[0])

	for i := 0; i < len(fsSearchs); i++ {
		if minLen > len(fsSearchs[i]) {
			minLen = len(fsSearchs[i])
		}
	}

	//
	for i := 0; i < minLen; i++ {
		if fsSearchs[0][i] != fsSearchs[len(fsSearchs)-1][i] {
			return fsSearchs[0][:i]
		}
	}
	return fsSearchs[0][0:minLen]

}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

//quickSort(data, 0, n, maxDepth(n))
