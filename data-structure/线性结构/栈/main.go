package main

import (
	"fmt"
	"strconv"
)

// 150. 逆波兰表达式求值
// 用切片模拟栈的实现
func evalRPN(tokens []string) int {
	record := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		// 如果遍历不是算数运算字符串，则将该值转换为int值，压入栈中
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			data, _ := strconv.Atoi(tokens[i])
			record = append(record, data)
		} else if tokens[i] == "+" { // 如果是算数运算，则计算，将计算的两栈元素pop掉
			tmp := record[len(record)-2] + record[len(record)-1]
			record = record[0 : len(record)-2]
			record = append(record, tmp)
		} else if tokens[i] == "-" {
			tmp := record[len(record)-2] - record[len(record)-1]
			record = record[0 : len(record)-2]
			record = append(record, tmp)
		} else if tokens[i] == "*" {
			tmp := record[len(record)-2] * record[len(record)-1]
			record = record[0 : len(record)-2]
			record = append(record, tmp)
		} else if tokens[i] == "/" {
			tmp := record[len(record)-2] / record[len(record)-1]
			record = record[0 : len(record)-2]
			record = append(record, tmp)
		}
	}
	// 返回最终的结果
	return record[0]
}

func main() {
	tokens := []string{"4", "13", "5", "/", "+"}
	data := evalRPN(tokens)
	fmt.Print(data)
}
