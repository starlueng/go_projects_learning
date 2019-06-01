package main

import (
	"fmt"
	"strconv"
)

func main() {
	var d = findAllNum(100, 999)
	fmt.Println(d)
}

func findAllNum(startNum int, endNum int) []string {
	var re []string

	for index := startNum; index < endNum; index++ {
		indexStr := strconv.Itoa(index)

		var result = 0
		for i := 0; i < len(indexStr); i++ {
			var num = int(indexStr[i] - '0')
			result += num * num * num
		}

		if result == index {
			fmt.Printf("%d 是需要的数据", index)
			re = append(re, indexStr)
		}

	}
	return re
}
