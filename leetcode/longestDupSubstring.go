package main

import (
	"fmt"
)

func main() {
	// 给定 abcdefabcd，期望输出 abcd
	fmt.Println(longestDupSubstring("abcdefabcd"))
	fmt.Println(longestDupSubstring("aaaaaaaaaa"))
}

// longestDupSubstring 最长重复子串
func longestDupSubstring(s string) string {
	// 检查长度为 l 的重复子串
	check := func(l int) string {
		vis := make(map[string]bool) // 记录子串是否重复
		for i := 1; i+l-1 <= len(s); i++ {
			j := i + l - 1
			// 当前子串
			t := s[i-1 : j]
			// 如果子串已经存在，则返回
			if vis[t] {
				return t
			}
			// 记录当前子串
			vis[t] = true
		}
		return "" // 没有找到重复子串
	}

	// 二分搜索最长重复子串的长度
	left, right := 0, len(s)
	ans := ""
	for left < right {
		mid := (left + right + 1) / 2 // 计算中间值
		t := check(mid)               // 检查长度为 mid 的重复子串
		if len(t) > 0 {               // 如果找到，更新左边界和结果
			left = mid
			ans = t
		} else { // 否则更新右边界
			right = mid - 1
		}
	}
	return ans // 返回最长重复子串
}
