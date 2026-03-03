package main

import (
	"fmt"
	"regexp"
)

func Readme() {
	fmt.Println("-----------Readme()---------------")
	fmt.Println("POSIX 正则表达式包含 BRE(Basic regular Expressions) 和 ERE(Extended Regexp Expressions)")
	fmt.Println("POSIX 字符组为了兼容各地区语言, 例如 è 字符无法使用 \\w 或者 a-z 正常匹配只能使用下面的字符组匹配")
	fmt.Println(tab, tab, "[:alnum:] => \\w", "字母数字字符")
	fmt.Println(tab, tab, "[:alpha:] => \\a", "字母字符")
	fmt.Println(tab, tab, "[:cntrl:] =>    ", "控制字符")
	fmt.Println(tab, tab, "[:digit:] => \\d", "数字字符")
	fmt.Println(tab, tab, "[:graph:] =>    ", "非空白字符(非空格, 控制字符等)")
	fmt.Println(tab, tab, "[:lower:] =>    ", "小写字母")
	fmt.Println(tab, tab, "[:upper:] =>    ", "大写字母")
	fmt.Println(tab, tab, "[:print:] => \\p", "与[:graph:]相似, 但是包含空格字符")
	fmt.Println(tab, tab, "[:punct:] =>    ", "标点字符")
	fmt.Println(tab, tab, "[:space:] => \\s", "所有空白字符(换行符, 空格, 制表符)")
	fmt.Println(tab, tab, "[:xdigit:] => \\x", "允许十六进制的数字(0-9a-fA-F)")
	fmt.Println("")
	fmt.Println(tab, "分组")
	fmt.Println(tab, tab, "(re)           编号的捕获分组, 例如 (\\d{3})\\.(\\d{3}) 通过 NumSubexp 获取正则表达式分组的数量", regexp.MustCompile(`(\d{3})\.(\d{3})`).NumSubexp()) // 2
	fmt.Println(tab, tab, "(?P<name>re)   命名并编号的捕获分组, 例如 (?P<first>\\d{3})\\.(?P<second>\\d{3}) 通过 SubexpNames 获取正则表达式子表达式的名称 [\"first\" \"second\"]")
	fmt.Println(tab, tab, "(?:re)         不捕获的分组, 例如 (?:\\d{3})\\.(\\d{3}) 通过 NumSubexp 获取正则表达式分组的数量会忽略 ?: 的分组", regexp.MustCompile(`(?:\d{3})\.(\d{3})`).NumSubexp()) // 1
	fmt.Println(tab, tab, "(?flags)       设置当前所在分组的标志, 不捕获也不匹配, 例如 (?i)(\\w{3})\\.(\\w{3}) 忽略大小写, 配合标志位使用")
	fmt.Println(tab, tab, "(?flags:re)    设置re段的标志, 不捕获的分组, 例如 (?i:^hello).*go  查找以 hello 开头(忽略大小写), 以 go 结尾的字符串, 配合标志位使用")
	fmt.Println(tab, "标志位")
	fmt.Println(tab, tab, "I              大小写敏感(默认关闭)")
	fmt.Println(tab, tab, "m              ^ 和 $ 在匹配文本开始和结尾之外, 还可以匹配行首和行尾(默认开启)")
	fmt.Println(tab, tab, "s              让 . 可以匹配 \\n (默认关闭)")
	fmt.Println(tab, tab, "U              非贪婪的: 交换 x* 和 x*?、x+ 和 x+?……的含义(默认关闭)")
	fmt.Println(tab, "边界匹配")
	fmt.Println(tab, tab, "^              匹配文本开始, 标志m为真时, 还匹配行首")
	fmt.Println(tab, tab, "$              匹配文本结尾, 标志m为真时, 还匹配行尾")
	fmt.Println(tab, tab, "\\A             匹配文本开始")
	fmt.Println(tab, tab, "\\b             单词边界(一边字符属于\\w, 另一边为文首、文尾、行首、行尾或属于\\W)")
	fmt.Println(tab, tab, "\\B             非单词边界")
	fmt.Println(tab, tab, "\\z             匹配文本结尾")
	fmt.Println(tab, "转义序列")
	fmt.Println(tab, tab, "\\a             响铃符(\\007)")
	fmt.Println(tab, tab, "\\f             换纸符(\\014)")
	fmt.Println(tab, tab, "\\t             水平制表符(\\011)")
	fmt.Println(tab, tab, "\\n             换行符(\\012)")
	fmt.Println(tab, tab, "\\r             回车符(\\015)")
	fmt.Println(tab, tab, "\\v             垂直制表符(\\013)")
	fmt.Println(tab, tab, "\\123           八进制表示的字符码(最多三个数字)")
	fmt.Println(tab, tab, "\\x7F           十六进制表示的字符码(必须两个数字)")
	fmt.Println(tab, tab, "\\x{10FFFF}     十六进制表示的字符码")
	fmt.Println(tab, tab, "\\*             字面值'*'")
	fmt.Println(tab, tab, "\\Q...\\E        反斜线后面的字符的字面值")
	fmt.Println("")

	res8, err := regexp.Compile("[\\w]")
	matchByte := []byte("ÀàÉèÇçÔôÛûŸÿ")
	if err != nil {
		fmt.Println("regexp.Compile 报错了", err)
	} else {
		fmt.Printf("字符 %q \n", matchByte)
		fmt.Println("res8, err := regexp.Compile(\"[\\\\w]\")\nres8.Match 正则匹配结果为 ", res8.Match(matchByte)) // false
	}
	fmt.Println("---------------------------------")

	// func (re *Regexp) Longest()
	fmt.Println("regexp.MustCompile(`foo.?`).Longest() // 正则表达式在之后的搜索中都采用 \"leftmost-longest\" 模式")
	fmt.Println("---------------------------------")

	// func (re *Regexp) NumSubexp() int // 返回该正则表达式中捕获分组的数量
	fmt.Println("func (re *Regexp) NumSubexp() int // 返回该正则表达式中捕获分组的数量")
	fmt.Print("regexp.MustCompile(`a(x*)b-c(de)f`).NumSubexp() ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b-c(de)f`).NumSubexp()) // 2
	fmt.Println("---------------------------------")

	// func (re *Regexp) LiteralPrefix() (prefix string, complete bool)
	fmt.Println("func (re *Regexp) LiteralPrefix() (prefix string, complete bool)")
	fmt.Println(tab, "返回一个字符串字面值 prefix, 任何匹配本正则表达式的字符串都会以 prefix 起始. 如果该字符串字面值包含整个正则表达式, 返回值 complete 会设为真")
	lp1, c1 := regexp.MustCompile(`https?`).LiteralPrefix()
	fmt.Printf("regexp.MustCompile(`https?`).LiteralPrefix() 结果为 %v, 是否包含整个正则表达式 %t\n", lp1, c1) // http, false
	lp2, c2 := regexp.MustCompile(`https`).LiteralPrefix()
	fmt.Printf("regexp.MustCompile(`https`).LiteralPrefix() 结果为 %v, 是否包含整个正则表达式 %t\n", lp2, c2) // https, true
	fmt.Println("---------------------------------")

	// func (re *Regexp) Split(s string, n int) []string
	fmt.Println("func (re *Regexp) Split(s string, n int) []string // 返回正则表达式分割 s 后的多个字符串, 结果中不包含正则匹配项")
	fmt.Println(tab, "如果 n >= 0, 方法最多返回 n 个匹配项/子匹配项; 否则, 它将返回所有匹配项")
	fmt.Print("regexp.MustCompile(`a(x*)b`).Split(`-1axxxb2-1axxb2-1axb2-1ab2-`, 1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).Split(`-1axxxb2-1axxb2-1axb2-1ab2-`, 1)) // [-1axxxb2-1axxb2-1axb2-1ab2-]
	fmt.Print("regexp.MustCompile(`a(x*)b`).Split(`-1axxxb2-1axxb2-1axb2-1ab2-`, 2) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).Split(`-1axxxb2-1axxb2-1axb2-1ab2-`, 2)) // [-1 2-1axxb2-1axb2-1ab2-]
	fmt.Print("regexp.MustCompile(`a(x*)b`).Split(`-1axxxb2-1axxb2-1axb2-1ab2-`, -1) ")
	fmt.Printf("结果为 %v\n", regexp.MustCompile(`a(x*)b`).Split(`-1axxxb2-1axxb2-1axb2-1ab2-`, -1)) // [-1 2-1 2-1 2-1 2-]
	fmt.Println("---------------------------------")

	// func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
	// func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
	fmt.Println("func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte")
	fmt.Println("func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte")
	fmt.Println(tab, "Expand 用从 src 中提取的由 match 返回的匹配项替换模板 template 中的变量, 并将结果附加到 dst 并返回结果")
	content := `
        # comment line
        option1: value1
        option2: value2

        # another comment line
        option3: value3
	`
	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	result := []byte{}
	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		// Template to convert "key: value" to "key=value" by
		// referencing the values captured by the regex pattern.
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.ExpandString(result, "$key=$value\n", content, submatches)
	}
	fmt.Println(string(result))
	fmt.Println("---------------------------------")

	// func (re *Regexp) SubexpIndex(name string) int
	// func (re *Regexp) SubexpNames() []string
	fmt.Println("func (re *Regexp) SubexpIndex(name string) int // 返回具有给定名称的第一个子表达式的索引, 如果没有具有该名称的子表达式, 则返回 -1")
	re := regexp.MustCompile(`he(?P<first>ll)o\s wo(?P<second>rl)d-f(?P<third>oo)`)
	fmt.Println(tab, "re := regexp.MustCompile(`he(?P<first>ll)o\\s wo(?P<second>rl)d-f(?P<third>oo)`)")
	fmt.Println(tab, "re.SubexpIndex(`first`)", re.SubexpIndex(`first`)) // 1
	fmt.Println("func (re *Regexp) SubexpNames() []string // 返回正则表达式中子表达式的名称")
	fmt.Println(tab, "re.SubexpNames() ", re.SubexpNames()) // [ first second third]
	fmt.Println("-----------")
	fmt.Println("不捕获分组: regexp.MustCompile(`(?:\\d{4})-(\\d{2})-(\\d{2})`).NumSubexp() ", regexp.MustCompile(`(?:\d{4})-(\d{2})-(\d{2})`).NumSubexp()) // 2
	fmt.Println("regexp.MustCompile(`(\\d{4})-(\\d{2})-(\\d{2})`).SubexpNames() ", regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`).SubexpNames())        // []
	fmt.Println("---------------------------------")

}
