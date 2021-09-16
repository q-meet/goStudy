package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
假设我们有一个数字到字母表的映射
1 -> ['a','b','c']
2 -> ['d','e']
3 -> ['f','g','h']
实现一个函数 对于给定的一串数字， 例如 “1”, "233" 返回一个所有可能的组合字符串列表
*/
var dict = map[int][]string{
	1: {"a", "b", "c"},
	2: {"d", "e"},
	3: {"f", "g"},
}

// getDictByIndex 通过单索引获取列表
func getDictByIndex(index int) []string {
	return dict[index]
}

//getDictByIds 通过ids获取内容
func getDictByIds(ids string)(res []string) {
	if ids == "" {
		return
	}
	for _, v := range ids {
		v1 :=  string(v)
		v2, _ := strconv.Atoi(v1)
		res = append(res, strings.Join(getDictByIndex(v2), ","))
	}
	return
}

// strToStr 字符串切换顺序
func strToStr(str string) {
	if str == "" {
		return
	}
	res := []byte(str)
	resLen := len(res)
	//任意组合的结果
	var ret []string
	for i := 0; i < resLen; i++ {

		for j := 1; j < resLen; j++ {
			//二位数的结合
			ret = append(ret, string(res[i]) + string(res[j]))
		}
	}
	fmt.Println(ret)
	return

}


func main() {
	list := "123"
	strToStr(list)
  	println( strings.Join(getDictByIds(list), ","))
}
