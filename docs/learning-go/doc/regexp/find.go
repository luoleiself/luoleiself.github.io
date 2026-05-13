package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func FindNote() {
	fmt.Println("-----------FindNote()---------------")
	fmt.Println("所有的字符都被视为utf-8编码的码值")
	fmt.Println("Regexp 类型提供了多达16个方法, 用于匹配正则表达式并获取匹配的结果, 它们的名字满足正则表达式 Find(All)?(String)?(Submatch)?(Index)?")
	fmt.Println(tab, "如果 'All' 出现了, 该方法会返回输入中所有互不重叠的匹配结果, 如果一个匹配结果的前后(没有间隔字符)存在长度为0的成功匹配, 该空匹配会被忽略. 包含 'All' 的方法要求一个额外的整数参数n, 如果 n >= 0, 方法会返回最多前 n 个匹配项, 否则, 将返回所有的匹配项.")
	fmt.Println(tab, "如果 'String' 出现了, 匹配对象为字符串, 否则应为 []byte 类型, 返回值和匹配对象的类型都是对应的.")
	fmt.Println(tab, "如果 'Submatch' 出现了, 返回值是表示正则表达式中成功的组匹配(子匹配/次级匹配)的切片, 组匹配是正则表达式内部的括号包围的次级表达式(也被称为'捕获分组'),从左到右按左括号的顺序编号. 索引为 0 的组匹配为完整表达式的匹配项, 1 为第一个分组的匹配项, 依次类推.")
	fmt.Println(tab, "如果 'Index' 出现了, 匹配/分组匹配会用输入流的字节索引对表示 result[2*n:2*n+1] 表示第 n 个分组匹配的匹配结果,如果没有 'Index', 匹配结果表示为匹配到的文本, 如果索引为负数, 表示分组匹配没有匹配到输入流中的文本. ")
	fmt.Println(tab, "")
	fmt.Println("---------------------------------")

	// func (re *Regexp) Find(b []byte) []byte
	fmt.Println("func (re *Regexp) Find(b []byte) []byte // 返回所有匹配项最左侧的一个匹配项")
	fmt.Print("regexp.MustCompile(`foo.?`).Find([]byte(`seafood fool fook`)) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo.?`).Find([]byte(`seafood fool fook`))) // "food"
	fmt.Println("-----------------")

	// func (re *Regexp) FindIndex(b []byte) (loc []int) // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindIndex(b []byte) (loc []int) // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)")
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindIndex([]byte(`seafood fool fook`)) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).FindIndex([]byte(`seafood fool fook`))) // [3 7]
	fmt.Println("-----------------")

	// func (re *Regexp) FindSubmatch(b []byte) [][]byte	// 返回所有匹配项最左侧的一个匹配项和子匹配项
	fmt.Println("func (re *Regexp) FindSubmatch(b []byte) [][]byte // 返回所有匹配项最左侧的一个匹配项和子匹配项")
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindSubmatch([]byte(`seafood fool fook`)) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).FindSubmatch([]byte(`seafood fool fook`))) // ["food" "d"]
	fmt.Println("-----------------")

	// func (re *Regexp) FindSubmatchIndex(b []byte) []int	// 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindSubmatchIndex(b []byte) []int	// 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)")
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindSubmatchIndex([]byte(`seafood fool fook`)) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).FindSubmatchIndex([]byte(`seafood fool fook`))) // [3 7 6 7]
	fmt.Println("-----------------")

	// func (re *Regexp) FindString(s string) string
	fmt.Println("func (re *Regexp) FindString(s string) string // 返回所有匹配项中最左侧的一个匹配项")
	fmt.Print("regexp.MustCompile(`foo.?`).FindString(`seafood fool fook`) 结果为 ")
	fmt.Printf("%q\n", regexp.MustCompile(`foo.?`).FindString(`seafood fool fook`)) // "food"
	fmt.Println("-----------------")

	// func (re *Regexp) FindStringIndex(s string) (loc []int)	// 返回所有匹配项最左侧的一个匹配的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindStringIndex(s string) (loc []int) // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)")
	fmt.Print("regexp.MustCompile(`foo.?`).FindStringIndex(`seafood fool fook`) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.?`).FindStringIndex(`seafood fool fook`)) // [3 7]
	fmt.Println("-----------------")

	// func (re *Regexp) FindStringSubmatch(s string) []string // 返回所有匹配项最左侧的一个匹配项和子匹配项
	fmt.Println("func (re *Regexp) FindStringSubmatch(s string) []string // 返回所有匹配项最左侧的一个匹配项和子匹配项")
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindStringSubmatch(`seafood fool fook`) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).FindStringSubmatch(`seafood fool fook`)) // ["food" "d"]
	fmt.Println("-----------------")

	// func (re *Regexp) FindStringSubmatchIndex(s string) []int // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindStringSubmatchIndex(s string) []int // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)")
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindStringSubmatchIndex(`seafood fool fook`) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).FindStringSubmatchIndex(`seafood fool fook`)) // [3 7 6 7]
	fmt.Println("-----------------")

	// func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int) // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int) // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)")
	fmt.Print("regexp.MustCompile(`foo.?`).FindReaderIndex(bytes.NewReader([]byte(`seafood fool fook`))) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.?`).FindReaderIndex(bytes.NewReader([]byte(`seafood fool fook`)))) // [3 7]
	fmt.Println("-----------------")

	// func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int	// 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int // 返回所有匹配项最左侧的一个匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)")
	fmt.Print("regexp.MustCompile(`foo.?`).FindReaderSubmatchIndex(bytes.NewReader([]byte(`seafood fool fook`))) ")
	fmt.Printf("结果为 %v // 正则没有分组,所以只有匹配项的结果,没有子匹配项的结果\n", regexp.MustCompile(`foo.?`).FindReaderSubmatchIndex(bytes.NewReader([]byte(`seafood fool fook`)))) // [3 7]
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindReaderSubmatchIndex(bytes.NewReader([]byte(`seafood fool fook`))) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).FindReaderSubmatchIndex(bytes.NewReader([]byte(`seafood fool fook`)))) // [3 7 6 7]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAll(b []byte, n int) [][]byte
	fmt.Println("func (re *Regexp) FindAll(b []byte, n int) [][]byte  // 返回所有匹配项")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`foo.?`).FindAll([]byte(`seafood fool fook`), 1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo.?`).FindAll([]byte(`seafood fool fook`), 1)) // ["food"]
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindAll([]byte(`seafood fool fook`), -1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).FindAll([]byte(`seafood fool fook`), -1)) // ["food" "fool" "fook"]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
	fmt.Println("func (re *Regexp) FindAllIndex(b []byte, n int) [][]int  // 返回所有匹配项的起止位置(含头不含尾)")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`foo.`).FindAllIndex([]byte(`seafood fool fook`), 1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.`).FindAllIndex([]byte(`seafood fool fook`), 1)) // [[3 7]]
	fmt.Print("regexp.MustCompile(`foo(.)`).FindAllIndex([]byte(`seafood fool fook`), 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.)`).FindAllIndex([]byte(`seafood fool fook`), 2)) // [[3 7] [8 12]]
	fmt.Print("regexp.MustCompile(`foo.`).FindAllIndex([]byte(`seafood fool fook`), -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.`).FindAllIndex([]byte(`seafood fool fook`), -1)) // [[3 7] [8 12] [13 17]]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
	fmt.Println("func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte // 返回所有匹配项和子匹配项")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindAllSubmatch([]byte(`seafood fool fook`), 1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).FindAllSubmatch([]byte(`seafood fool fook`), 1)) // [["food" "d"]]
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindAllSubmatch([]byte(`seafood fool fook`), 2) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).FindAllSubmatch([]byte(`seafood fool fook`), 2)) // [["food" "d"] ["fool" "l"]]
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindAllSubmatch([]byte(`seafood fool fook`), -1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).FindAllSubmatch([]byte(`seafood fool fook`), -1)) // [["food" "d"] ["fool" "l"] ["fook" "k"]]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int //  返回所有匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)
	fmt.Println("func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int // 返回所有匹配项的起止位置(含头不含尾)和子匹配项的起止位置(含头不含尾)")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`foo.?`).FindAllSubmatchIndex([]byte(`seafood fool fook`), 1) ")
	fmt.Printf("结果为 %v // 正则没有分组,所以只有匹配项的结果,没有子匹配项的结果\n", regexp.MustCompile(`foo.?`).FindAllSubmatchIndex([]byte(`seafood fool fook`), 1)) // [[3 7]]
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindAllSubmatchIndex([]byte(`seafood fool fook`), 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).FindAllSubmatchIndex([]byte(`seafood fool fook`), 2)) // [[3 7 6 7] [8 12 11 12]]
	fmt.Print("regexp.MustCompile(`foo(.?)`).FindAllSubmatchIndex([]byte(`seafood fool fook`), -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).FindAllSubmatchIndex([]byte(`seafood fool fook`), -1)) // [[3 7 6 7] [8 12 11 12] [13 17 16 17]]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllString(s string, n int) []string
	fmt.Println("func (re *Regexp) FindAllString(s string, n int) []string // 返回所有匹配项")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`foo.`).FindAllString(`seafood fool fook`, 1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo.`).FindAllString(`seafood fool fook`, 1)) // ["food"]
	fmt.Print("regexp.MustCompile(`foo.`).FindAllString(`seafood fool fook`, 2) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo.`).FindAllString(`seafood fool fook`, 2)) // ["food" "fool"]
	fmt.Print("regexp.MustCompile(`foo.`).FindAllString(`seafood fool fook`, -1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo.`).FindAllString(`seafood fool fook`, -1)) // ["food" "fool" "fook"]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
	fmt.Println("func (re *Regexp) FindAllStringIndex(s string, n int) [][]int // 返回所有匹配项的起止位置(含头不含尾)")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`foo.`).FindAllStringIndex(`seafood fool fook`, 1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.`).FindAllStringIndex(`seafood fool fook`, 1)) // [[3 7]]
	fmt.Print("regexp.MustCompile(`foo.`).FindAllStringIndex(`seafood fool fook`, 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.`).FindAllStringIndex(`seafood fool fook`, 2)) // [[3 7] [8 12]]
	fmt.Print("regexp.MustCompile(`foo.`).FindAllStringIndex(`seafood fool fook`, -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo.`).FindAllStringIndex(`seafood fool fook`, -1)) // [[3 7] [8 12] [13 17]]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
	fmt.Println("func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string // 返回所有匹配项和子匹配项")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-ab-`, 1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-ab-`, 1)) // [["ab" ""]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-ab-`, 2) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-ab-`, 2)) // [["ab" ""]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-ab-`, -1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-ab-`, -1)) // [["ab" ""]]
	fmt.Println("------")
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxb-`, 1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxb-`, 1)) // [["axxb" "xx"]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxb-`, 2) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxb-`, 2)) // [["axxb" "xx"]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxb-`, -1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxb-`, -1)) // [["axxb" "xx"]]
	fmt.Println("------")
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxxb-axxb-axb-ab-`, 1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxxb-axxb-axb-ab-`, 1)) // [["axxxb" "xxx"]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxxb-axxb-axb-ab-`, 2) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxxb-axxb-axb-ab-`, 2)) // [["axxxb" "xxx"] ["axxb" "xx"]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxxb-axxb-axb-ab-`, -1) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatch(`-axxxb-axxb-axb-ab-`, -1)) // [["axxxb" "xxx"] ["axxb" "xx"] ["axb" "x"] ["ab" ""]]
	fmt.Println("-----------------")

	// func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
	fmt.Println("func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int // 返回所有匹配项和子匹配项的起止位置(含头不含尾)")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-ab-`, 1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-ab-`, 1)) // [[1 3 2 2]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-ab-`, 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-ab-`, 2)) // [[1 3 2 2]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-ab-`, -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-ab-`, -1)) //[[1 3 2 2]]
	fmt.Println("------")
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxb-`, 1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxb-`, 1)) // [[1 5 2 4]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxb-`, 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxb-`, 2)) //[[1 5 2 4]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxb-`, -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxb-`, -1)) //[[1 5 2 4]]
	fmt.Println("------")
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxxb-axxb-axb-ab-`, 1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxxb-axxb-axb-ab-`, 1)) // [[1 6 2 5]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxxb-axxb-axb-ab-`, 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxxb-axxb-axb-ab-`, 2)) // [[1 6 2 5] [7 11 8 10]]
	fmt.Print("regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxxb-axxb-axb-ab-`, -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).FindAllStringSubmatchIndex(`-axxxb-axxb-axb-ab-`, -1)) // [[1 6 2 5] [7 11 8 10] [12 15 13 14] [16 18 17 17]]
	fmt.Println("---------------------------------")

}
