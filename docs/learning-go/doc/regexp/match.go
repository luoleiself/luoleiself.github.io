package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func MatchNote() {
	fmt.Println("-----------MatchNote()---------------")
	fmt.Println("检查文本是否包含指定内容")
	fmt.Println(tab, "QuoteMeta, Match, MatchString, MatchReader")
	fmt.Println("---------------------------------")

	fmt.Println("func QuoteMeta(s string) string // 返回将 s 中所有正则表达式元字符都进行转义后字符串. 该字符串可以用在正则表达式中匹配字面值")
	// func QuoteMeta(s string) string
	res := regexp.QuoteMeta("[hello]\\s.")
	fmt.Println(`regexp.QuoteMeta("[hello]\\s.") 的结果为`, res) // \[hello\]\\s\.
	fmt.Println("-----------------")

	fmt.Println(`func Match(pattern string, b []byte) (matched bool, err error) // 检查 b 中是否存在匹配 pattern 的子序列`)
	// func Match(pattern string, b []byte) (matched bool, err error)
	res1, _ := regexp.Match("[hello]\\s.", []byte("hello world"))
	fmt.Println(`regexp.Match("[hello]\\s.", []byte("hello world")) 的结果为`, res1) // true
	fmt.Println("-----------------")

	fmt.Println(`func MatchString(pattern string, s string) (matched bool, err error) // 检查 s 中是否存在匹配 pattern 的子序列, 用法和 Match 类似`)
	// func MatchString(pattern string, s string) (matched bool, err error)
	res2, _ := regexp.MatchString("[hello]\\s.", "hello world")
	fmt.Println(`regexp.MatchString("[hello]\\s.", "hello world") 的结果为`, res2) // true
	fmt.Println("-----------------")

	fmt.Println(`func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) // 检查 r 中是否存在 匹配 pattern 的子序列, 用法和 Match 类似`)
	// func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
	res3, _ := regexp.MatchReader("[hello]\\s.", bytes.NewReader([]byte("hello world")))
	fmt.Println(`regexp.MatchReader("[hello]\\s.", bytes.NewReader([]byte("hello world"))) 的结果为`, res3) // true
	fmt.Println("---------------------------------")
}
