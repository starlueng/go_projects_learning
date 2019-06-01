package main

import (
	"fmt"
	"strconv"
)

func main() {
	// d := findAllNum(100, 999)
	// fmt.Println(d)
	c := sushu(101, 200)
	fmt.Printf("101～200的素数共有%d 个 具体参数为%v", len(c), c)
}

/*水仙数 153 = 1*1*1 + 5 *5 *5  +3 *3 *3 */
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
			re = append(re, indexStr)
		}

	}
	return re
}

/*判断101～200有多少素数*/

func sushu(start int, end int) []int {

	var resultnums []int
	for index := start; index < end+1; index++ {
		testnum := 2
		issushu := true

		for {
			if testnum < index {
				if index%testnum == 0 {
					issushu = false
					break
				} else {
					testnum++
				}
			} else {
				break
			}

		}
		if issushu {

			resultnums = append(resultnums, index)
		}
	}
	return resultnums
}
