package main

import (
	"fmt"
	"strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	f("To Upper: %s\n", strings.ToUpper("Hello THERE"))
	f("To Lower: %s\n", strings.ToLower("Hello THERE")) // hello there

	// 忽略大小写比较两个字符串
	f("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAlis")) // true
	f("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAli"))  // false

	// 子字符串出现的索引
	f("Index: %v\n", strings.Index("Mihalis", "ha")) // Index: 2
	f("Index: %v\n", strings.Index("Mihalis", "Ha")) // Index: -1

	// 子字符串出现次数
	f("Count: %v\n", strings.Count("Mihalis", "i"))
	f("Count: %v\n", strings.Count("Mihalis", "I"))

	// 重复
	f("Repeat: %s\n", strings.Repeat("ab", 5)) // ababababab

	f("TrimSpace: %s\n", strings.TrimSpace(" \tThis is a line. \n"))            // This is a line.
	f("TrimLeft: %s", strings.TrimLeft(" \tThis is a\t line. \n", "\n\t "))     // This is a	 line.
	f("TrimRight: %s\n", strings.TrimRight(" \tThis is a\t line. \n", "\n\t ")) //  	This is a	 line.

	f("Prefix: %v\n", strings.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", strings.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", strings.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", strings.HasSuffix("Mihalis", "IS"))

	// 提取非空白字符，返回一个字符串切片
	t := strings.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = strings.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	// 分割字符串
	f("%s\n", strings.Split("abcd efg", ""))            // [a b c d   e f g]
	f("%s\n", strings.Replace("abcd efg", "", "_", -1)) // _a_b_c_d_ _e_f_g_
	f("%s\n", strings.Replace("abcd efg", "", "_", 4))  // _a_b_c_d efg
	f("%s\n", strings.Replace("abcd efg", "", "_", 2))  // _a_bcd efg

	// 保留分割符
	f("SplitAfter: %s\n", strings.SplitAfter("123++432++", "++")) // [123++ 432++ ]

	// 自定义函数
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", strings.TrimFunc("123 abc ABC \t .", trimFunction)) // abc ABC
}
