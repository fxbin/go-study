package main

import "fmt"

// 寻找最长不含有重复字符的子串
// 对于一个字母X
// lastOccurred[x] 不存在, 或者 < start ——> 无需操作
// lastOccurred[x] >= start ——> 更新start
// 更新lastOccurred[x], 更新maxLength
func lengthOfNonRepeatingSubStr(s string) int {

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("aaa"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcc"))
	fmt.Println(lengthOfNonRepeatingSubStr("你好你你你"))
}
