package main

import (
	"fmt"
	"regexp"
)

func CompileNote() {
	fmt.Println("-----------CompileNote()---------------")
	fmt.Println("Compile 和 CompilePOSIX, MustCompile 和 MustCompilePOSIX 的区别")
	fmt.Println(tab, "解析并返回一个正则表达式")
	fmt.Println(tab, "Compile 和 CompilePOSIX 返回第二个值表示解析错误")
	fmt.Println(tab, "MustCompile 和 MustCompilePOSIX 解析失败直接 panic")
	fmt.Println(tab, "Compile 和 MustCompile 在匹配文本时, 该正则表达式会尽可能早的开始匹配, 并且在匹配过程中选择回溯搜索到的第一个匹配结果. 这种模式被称为 \"leftmost-first\", Perl、Python和其他实现都采用了这种模式")
	fmt.Println(tab, "CompilePOSIX 和 MustCompilePOSIX 在匹配文本时, 该正则表达式会尽可能早的开始匹配,  并且在匹配过程中选择搜索到的最长的匹配结果. 这种模式被称为 \"leftmost-longest\"")
	fmt.Println("---------------------------------")

	// Compile(expr string) (*Regexp, error) // 解析并返回一个正则表达式, 在匹配文本时, 该正则表达式会尽可能早的开始匹配, 并且在匹配过程中选择回溯搜索到的第一个匹配结果. 这种模式被称为"leftmost-first", Perl、Python和其他实现都采用了这种模式, 但本包的实现没有回溯的损耗
	fmt.Println("func Compile(expr string) (*Regexp, error)")
	fmt.Println(tab, "解析并返回一个正则表达式, 在匹配文本时, 该正则表达式会尽可能早的开始匹配, 并且在匹配过程中选择回溯搜索到的第一个匹配结果. 这种模式被称为\"leftmost-first\", Perl、Python和其他实现都采用了这种模式, 但本包的实现没有回溯的损耗")
	res4, err := regexp.Compile("foo|foobar")
	if err != nil {
		fmt.Println(`regexp.Compile("foo|foobar") 结果报错`, err)
	} else {
		fmt.Println(`regexp.Compile("foo|foobar") 结果为`, res4)                                                // foo|foobar
		fmt.Println("res4.String() 返回用于编译成正则表达式的字符串", res4.String())                                         // foo|foobar
		fmt.Println("res4.FindString(\"foobarbaz\") 结果为", res4.FindString("foobarbaz"), "返回所有匹配项的最左侧的一个匹配项") // foo
		fmt.Printf("res4.Find([]byte(\"foobarbaz\")) 结果为 %q\n", res4.Find([]byte("foobarbaz")))              // foo
	}
	fmt.Println("-----------------")

	// CompilePOSIX(expr string) (*Regexp, error) // 作用类似 Compile 但会将语法约束到POSIX ERE（egrep）语法, 并将匹配模式设置为leftmost-longest. 在匹配文本时, 该正则表达式会尽可能早的开始匹配,  并且在匹配过程中选择搜索到的最长的匹配结果. 这种模式被称为"leftmost-longest"
	fmt.Println("func CompilePOSIX(expr string) (*Regexp, error)")
	fmt.Println(tab, "作用类似 Compile 但会将语法约束到POSIX ERE(egrep)语法, 并将匹配模式设置为leftmost-longest. 在匹配文本时, 该正则表达式会尽可能早的开始匹配,  并且在匹配过程中选择搜索到的最长的匹配结果. 这种模式被称为\"leftmost-longest\"")
	res5, err := regexp.CompilePOSIX("foo|foobar")
	if err != nil {
		fmt.Println(`regexp.CompilePOSIX("foo|foobar") 结果报错`, err)
	} else {
		fmt.Println(`regexp.CompilePOSIX(foo|foobar") 结果为`, res5)                                            // foo|foobar
		fmt.Println("res5.String() 返回用于编译成正则表达式的字符串", res5.String())                                         // foo|foobar
		fmt.Println("res5.FindString(\"foobarbaz\") 结果为", res5.FindString("foobarbaz"), "返回所有匹配项的最左侧的一个匹配项") // foobar
		fmt.Printf("res5.Find([]byte(\"foobarbaz\")) 结果为 %q\n", res5.Find([]byte("foobarbaz")))              // foobar
	}
	fmt.Println("-----------------")

	// MustCompile(str string) *Regexp // 作用类似 Compile, 解析失败时会触发 panic, 主要用于全局正则表达式变量的安全初始化
	fmt.Println(`func MustCompile(str string) *Regexp`)
	fmt.Println(tab, "作用类似 Compile, 解析失败时会触发 panic, 主要用于全局正则表达式变量的安全初始化")
	res6 := regexp.MustCompile("[hello]\\s.{4}")
	if res6 != nil {
		fmt.Println(`regexp.MustCompile("[hello]\\s.{4}") 结果为`, res6)                                          // [hello]\s.{4}
		fmt.Println("res6.String() 返回用于编译成正则表达式的字符串", res6.String())                                           // [hello]\s.{4}
		fmt.Println(`res6.FindString("hello world") 结果为`, res6.FindString("hello world"), "返回所有匹配项的最左侧的一个匹配项") // o worl
	}
	fmt.Println("-----------------")

	// MustCompilePOSIX(str string) *Regexp // 作用类似 CompilePOSIX, 解析失败时会触发 panic, 主要用于全局正则表达式变量的安全初始化
	fmt.Println(`func MustCompilePOSIX(str string) *Regexp`)
	fmt.Println(tab, "作用类似 CompilePOSIX, 解析失败时会触发 panic, 主要用于全局正则表达式变量的安全初始化")
	res7 := regexp.MustCompilePOSIX("[he][:space:].{3}")
	if res7 != nil {
		fmt.Println(`regexp.MustCompilePOSIX("[he][:space:].{3}") 结果为`, res7)                                  // [he][:space:].{3}
		fmt.Println("res7.String() 返回用于编译成正则表达式的字符串", res7.String())                                           // [he][:space:].{3}
		fmt.Println(`res7.FindString("hello world") 结果为`, res7.FindString("hello world"), "返回所有匹配项的最左侧的一个匹配项") // hello
	}
	fmt.Println("---------------------------------")
}
