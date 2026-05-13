package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func ReplaceNote() {
	fmt.Println("-----------ReplaceNote()---------------")
	fmt.Println("使用正则替换原字符串并返回副本")
	fmt.Println(tab, "ReplaceAll, ReplaceAllString	替换字节切片和字符串, 第二个参数可以使用 '$' 符号, 会按照 expand 方法的规则进行解析")
	fmt.Println(tab, "ReplaceAllFunc, ReplaceAllStringFunc	替换字节切片和字符串, 第二个参数为接收与第一个参数类型相同并返回相同类型的函数, 不能使用 '$' 符号")
	fmt.Println(tab, "ReplaceAllLiteral, ReplaceAllLiteralString	替换字节切片和字符串, 第二个参数如果使用 '$' 符号, 不会使用 expand 进行扩展")
	fmt.Println("---------------------------------")

	// func (re *Regexp) ReplaceAll(src, repl []byte) []byte
	fmt.Println("func (re *Regexp) ReplaceAll(src, repl []byte) []byte")
	fmt.Println(tab, "返回 src 的一个拷贝, 将 src 中所有的匹配结果都替换为 repl, 在替换时, repl中的'$'符号会按照Expand方法的规则进行解释和替换, 例如$1会被替换为第一个分组匹配结果")
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAll([]byte(`seafood fool fook`), []byte(`new`)) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).ReplaceAll([]byte(`seafood fool fook`), []byte(`new`))) // seanew new new
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAll([]byte(`seafood fool fook`), []byte(`$1`)) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).ReplaceAll([]byte(`seafood fool fook`), []byte(`$1`))) // sead l k
	fmt.Println("--------------")

	// func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
	fmt.Println("func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte")
	fmt.Println(tab, "返回src的一个拷贝, 将src中所有re的匹配结果(设为matched)都替换为repl(matched). repl返回的切片被直接使用, 会使用Expand进行扩展")
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAllFunc([]byte(`seafood fool fook`), bytes.ToUpper) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).ReplaceAllFunc([]byte(`seafood fool fook`), bytes.ToUpper)) // seaFOOD FOOL FOOK
	fmt.Println("--------------")

	// func (re *Regexp) ReplaceAllString(src, repl string) string
	fmt.Println("func (re *Regexp) ReplaceAllString(src, repl string) string")
	fmt.Println(tab, "返回src的一个拷贝, 将src中所有的匹配结果都替换为repl. 在替换时, repl 中的'$'符号会按照Expand方法的规则进行解释和替换, 例如$1会被替换为第一个分组匹配结果")
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAllString(`seafood fool fook`, `new`) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).ReplaceAllString(`seafood fool fook`, `new`)) // seanew new new
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAllString(`seafood fool fook`, `$1`) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).ReplaceAllString(`seafood fool fook`, `$1`)) // sead l k
	fmt.Println("--------------")

	// func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
	fmt.Println("func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string")
	fmt.Println(tab, "返回src的一个拷贝, 将src中所有re的匹配结果(设为matched)都替换为repl(matched). repl返回的字符串被直接使用, 不会使用Expand进行扩展")
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAllStringFunc(`seafood fool fook`, strings.ToUpper) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).ReplaceAllStringFunc(`seafood fool fook`, strings.ToUpper)) // seaFOOD FOOL FOOK
	fmt.Println("--------------")

	// func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
	fmt.Println("func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte")
	fmt.Println(tab, "返回src的一个拷贝, 将src中所有re的匹配结果都替换为repl. repl参数被直接使用, 不会使用Expand进行扩展")
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAllLiteral([]byte(`seafood fool fook`), []byte(`$1`)) ")
	fmt.Printf("结果为 %q\n", regexp.MustCompile(`foo(.?)`).ReplaceAllLiteral([]byte(`seafood fool fook`), []byte(`$1`))) // sea$1 $1 $1
	fmt.Println("--------------")

	// func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
	fmt.Println("func (re *Regexp) ReplaceAllLiteralString(src, repl string) string")
	fmt.Println(tab, "返回src的一个拷贝, 将src中所有re的匹配结果都替换为repl. repl参数被直接使用, 不会使用Expand进行扩展")
	fmt.Print("regexp.MustCompile(`foo(.?)`).ReplaceAllLiteralString(`seafood fool fook`, `$1`) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`foo(.?)`).ReplaceAllLiteralString(`seafood fool fook`, `$1`)) // sea$1 $1 $1
	fmt.Println("---------------------------------")
}
