package main

import "fmt"

func getNext(next []int, str string) {
	// 定义俩指针，初始化
	j := 0
	next[0] = j
	for i := 1; i < len(str); i++ {
		// 循环判断前后缀不相等
		for j > 0 && str[j] != str[i] {
			j = next[j-1]
		}

		// 判断前后缀相等
		if str[j] == str[i] {
			j++
		}
		next[i] = j
	}
}

func main() {
	// str := "abcabcabcabc"
	// str2 := "asdfasdfasdf"
	str3 := "aabaabaaf"
	next := make([]int, len(str3))
	getNext(next, str3)
	fmt.Println(next)
}
