package main

import (
	"log"
	"strings"
)

// 题目14: 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀,则返回""
// 输入不存在公共前缀。所有输入只包含小写字母 a-z
//
// 输入: ["flower","flow","flight"]
// 输出: "fl"
// 输入: ["dog","racecar","car"]
// 输出: ""
func main() {
	log.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	log.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
