package main

import (
	"bufio"
	"fmt"
	"strings"
)

func scannerNote() {
	fmt.Println("----------------scannerNote()----------------")
	fmt.Println("Scanner类型提供了方便的读取数据的接口")
	// func NewScanner(r io.Reader) *Scanner
	fmt.Println(tab, "NewScanner(r io.Reader) *Scanner // 创建并返回一个从 r 读取数据的 Scanner, 默认的分割函数是 ScanLines")
	// func (s *Scanner) Buffer(buf []byte, max int)
	fmt.Println(tab, "func (s *Scanner) Buffer(buf []byte, max int) // 设置缓冲区的大小")
	// func (s *Scanner) Bytes() []byte
	fmt.Println(tab, "func (s *Scanner) Bytes() []byte // 返回最近一次 Scan 调用生成的 token. 底层数组指向的数据可能会被下一次 Scan 的调用重写")
	// func (s *Scanner) Err() error
	fmt.Println(tab, "func (s *Scanner) Err() error // 返回 Scanner 遇到的第一个非 EOF 的错误")
	// func (s *Scanner) Scan() bool
	fmt.Println(tab, "func (s *Scanner) Scan() bool")
	fmt.Println(tab, tab, "获取当前位置的 token, 并让 Scanner 的扫描位置移动到下一个 token, 当扫描因为抵达输入流结尾或者遇到错误而停止时, 本方法会返回 false, 除非是 io.EOF, 此时Err会返回 nil")
	// func (s *Scanner) Split(split SplitFunc)
	fmt.Println(tab, "func (s *Scanner) Split(split SplitFunc) // 设置该 Scanner 的分割函数. 本方法必须在 Scan 之前调用, 默认为 ScanLines")
	// func (s *Scanner) Text() string
	fmt.Println(tab, "func (s *Scanner) Text() string // 返回最近一次 Scan 调用生成的 token, 会申请创建一个字符串保存 token 并返回该字符串")
	fmt.Println("-------------------------------------")

	// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
	fmt.Println("type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error) // 定义 SplitFunc 类型, 代表用于对输出作词法分析的分割函数")
	// func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
	fmt.Println(tab, "func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)")
	fmt.Println(tab, tab, "符合 SplitFunc 类型, 将每个字节作为一个 token 返回")
	// func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
	fmt.Println(tab, "func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)")
	fmt.Println(tab, tab, "符合 SplitFunc 类型, 将每一行文本去掉末尾的换行标记作为一个 token 返回, 返回的行可以是空字符串")
	// func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
	fmt.Println(tab, "func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)")
	fmt.Println(tab, tab, "符合 SplitFunc 类型, 将每个 utf-8 编码的 unicode 码值作为一个 token 返回")
	// func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
	fmt.Println(tab, "func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)")
	fmt.Println(tab, tab, "符合 SplitFunc 类型, 将空白(参见 unicode.IsSpace )分隔的片段(去掉前后空白后)作为一个 token 返回, 永远不会返回空字符串")

	fmt.Println("-------------------------------------")

	str1 := "AaBbCcDdEeFf"
	var s1 *bufio.Scanner
	fmt.Println(str1)
	fmt.Println("var s1 *bufio.Scanner")

	fmt.Println(`
s1 = bufio.NewScanner(strings.NewReader(str1))
s1.Split(bufio.ScanBytes) // 按字节方式分割
for s1.Scan() {
    fmt.Printf("s1.Bytes() \%\q\t s1.Text() \%\v\t\n", s1.Bytes(), s1.Text())
}
    `)
	s1 = bufio.NewScanner(strings.NewReader(str1))
	s1.Split(bufio.ScanBytes) // 按字节方式分割
	for s1.Scan() {
		fmt.Printf("s1.Bytes() %q\t s1.Text() %v\t\n", s1.Bytes(), s1.Text())
	}
	fmt.Println("----------------")

	fmt.Println(`
s1 = bufio.NewScanner(strings.NewReader(str1))
s1.Split(bufio.ScanRunes) // 按 utf-8 编码的 unicode 码值分割
for s1.Scan() {
    fmt.Printf("s1.Bytes() \%\q\t s1.Text() \%\v\t\n", s1.Bytes(), s1.Text())
}    
    `)
	s1 = bufio.NewScanner(strings.NewReader(str1))
	s1.Split(bufio.ScanRunes) // 按 utf-8 编码的 unicode 码值分割
	for s1.Scan() {
		fmt.Printf("s1.Bytes() %q\t s1.Text() %v\t\n", s1.Bytes(), s1.Text())
	}
	fmt.Println("----------------")

	fmt.Println(`
s1 = bufio.NewScanner(strings.NewReader(str1))
s1.Split(bufio.ScanLines) // 按每一行文本末尾的换行标记(去掉换行标记)分割
for s1.Scan() {
    fmt.Printf("s1.Bytes() \%\q\t s1.Text() \%\v\t\n", s1.Bytes(), s1.Text())
}
    `)
	s1 = bufio.NewScanner(strings.NewReader(str1))
	s1.Split(bufio.ScanLines) // 按每一行文本末尾的换行标记(去掉换行标记)分割
	for s1.Scan() {
		fmt.Printf("s1.Bytes() %q\t s1.Text() %v\t\n", s1.Bytes(), s1.Text())
	}
	fmt.Println("----------------")

	fmt.Println(`
s1 = bufio.NewScanner(strings.NewReader(str1))
s1.Split(bufio.ScanWords) // 按空白字符分割的片段(去掉前后空白)分割
for s1.Scan() {
    fmt.Printf("s1.Bytes() \%\q\t s1.Text() \%\v\t\n", s1.Bytes(), s1.Text())
}
    `)
	s1 = bufio.NewScanner(strings.NewReader(str1))
	s1.Split(bufio.ScanWords) // 按空白字符分割的片段(去掉前后空白)分割
	for s1.Scan() {
		fmt.Printf("s1.Bytes() %q\t s1.Text() %v\t\n", s1.Bytes(), s1.Text())
	}
	fmt.Println("-------------------------------------")

	str2 := "练习 golang 的 scanner 结构体的各种方法"
	var s2 *bufio.Scanner
	fmt.Println(str2)
	fmt.Println("var s2 *bufio.Scanner")

	fmt.Println(`
s2 = bufio.NewScanner(strings.NewReader(str2))
s2.Split(bufio.ScanRunes) // 按 utf-8 编码的 unicode 码值分割
for s2.Scan() {
    fmt.Printf("s2.Bytes() \%\v\t s2.Text() \%\v\t\n", string(s2.Bytes()), s2.Text())
}
    `)
	s2 = bufio.NewScanner(strings.NewReader(str2))
	s2.Split(bufio.ScanRunes) // 按 utf-8 编码的 unicode 码值分割
	for s2.Scan() {
		fmt.Printf("s2.Bytes() %v\t s2.Text() %v\t\n", string(s2.Bytes()), s2.Text())
	}
	fmt.Println("----------------")

	fmt.Println(`
s2 = bufio.NewScanner(strings.NewReader(str2))
s2.Split(bufio.ScanLines) // 按每一行文本末尾的换行标记(去掉换行标记)分割
for s2.Scan() {
    fmt.Printf("s2.Bytes() \%\v\t s2.Text() \%\v\t\n", string(s2.Bytes()), s2.Text())
}
    `)
	s2 = bufio.NewScanner(strings.NewReader(str2))
	s2.Split(bufio.ScanLines) // 按每一行文本末尾的换行标记(去掉换行标记)分割
	for s2.Scan() {
		fmt.Printf("s2.Bytes() %v\t s2.Text() %v\t\n", string(s2.Bytes()), s2.Text())
	}
	fmt.Println("----------------")

	fmt.Println(`
s2 = bufio.NewScanner(strings.NewReader(str2))
s2.Split(bufio.ScanWords) // 按空白字符分割的片段(去掉前后空白)分割
for s2.Scan() {
    fmt.Printf("s2.Bytes() \%\v\t s2.Text() \%\v\t\n", string(s2.Bytes()), s2.Text())
}
    `)
	s2 = bufio.NewScanner(strings.NewReader(str2))
	s2.Split(bufio.ScanWords) // 按空白字符分割的片段(去掉前后空白)分割
	for s2.Scan() {
		fmt.Printf("s2.Bytes() %v\t s2.Text() %v\t\n", string(s2.Bytes()), s2.Text())
	}
	fmt.Println("-------------------------------------")

	fmt.Println("func SplitN(s string, sep string, n int) []string")
	fmt.Println(tab, "以去掉 sep 的方式分割 s 直到结尾, 每个 sep 都会进行一次分割, 即使两个 sep 相邻, 也会进行两次分割, n 控制返回的最多的分割项")
	fmt.Println("strings.SplitN(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\",1) 结果为", strings.SplitN("ab-axb-axxb-axxxb-axxxxxb", "x", 1))   // [ab-axb-axxb-axxxb-axxxxxb]
	fmt.Println("strings.SplitN(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\",2) 结果为", strings.SplitN("ab-axb-axxb-axxxb-axxxxxb", "x", 2))   // [ab-a b-axxb-axxxb-axxxxxb]
	fmt.Println("strings.SplitN(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\",-1) 结果为", strings.SplitN("ab-axb-axxb-axxxb-axxxxxb", "x", -1)) // [ab-a b-a  b-a   b-a     b]
	fmt.Println("-------------------------------------")

	fmt.Println("func SplitAfter(s string, sep string) []sting")
	fmt.Println(tab, "从 s 中出现 sep 的后面(保留 sep)进行分割直到结尾, 每个 sep 都会进行一次分割, 即使两个 sep 相邻, 也会进行两次分割")
	fmt.Println("strings.SplitAfter(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\") 结果为", strings.SplitAfter("ab-axb-axxb-axxxb-axxxxxb", "x")) // [ab-ax b-ax x b-ax x x b-ax x x x x b]
	fmt.Println("-------------------------------------")

	fmt.Println("func SplitAfterN(s string, sep string, n int) []string")
	fmt.Println(tab, "从 s 中出现 sep 的后面(保留 sep)进行分割直到结尾, 每个 sep 都会进行一次分割, 即使两个 sep 相邻, 也会进行两次分割, n 控制返回的最多的分割项")
	fmt.Println("strings.SplitAfterN(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\",1) 结果为", strings.SplitAfterN("ab-axb-axxb-axxxb-axxxxxb", "x", 1))   // [ab-axb-axxb-axxxb-axxxxxb]
	fmt.Println("strings.SplitAfterN(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\",2) 结果为", strings.SplitAfterN("ab-axb-axxb-axxxb-axxxxxb", "x", 2))   // [ab-ax b-axxb-axxxb-axxxxxb]
	fmt.Println("strings.SplitAfterN(\"ab-axb-axxb-axxxb-axxxxxb\",\"x\",-1) 结果为", strings.SplitAfterN("ab-axb-axxb-axxxb-axxxxxb", "x", -1)) // [ab-ax b-ax x b-ax x x b-ax x x x x b]
	fmt.Println("-------------------------------------")
}
