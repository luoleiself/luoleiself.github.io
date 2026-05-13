package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	_ "strconv"
	"strings"
)

func bufferNote() {
	fmt.Println("----------------bufferNote()----------------")
	fmt.Println("strings 有 Reader, Builder, Replacer, 没有 Writer.")
	fmt.Println(tab, "Reader 不包含 ReadLine, ReadString, ReadSlice 方法")
	fmt.Println("bytes 有 Reader, Buffer, 没有 Writer.")
	fmt.Println(tab, "Reader 不包含 ReadBytes, ReadLine, ReadString, ReadSlice 方法")
	fmt.Println("bufio 有 Reader, Writer, ReaderWriter, Scanner")
	fmt.Println(tab, "Reader 有 strings、bytes 包的 Reader 不包含的方法")
	fmt.Println(tab, " ReadLine 是低水平的方法, 应使用 ReadBytes('\\n') 或 ReadString('\\n') 代替")
	fmt.Println("--------------------")

	fmt.Println("ReadRune  读取一个 UTF-8 编码的字符, 并返回其 Unicode 编码和字节数")
	fmt.Println("--------------------------------------------")

	fmt.Println("bufio.NewReader(rd io.Reader) *Reader")
	fmt.Println("bufio.NewReaderSize(rd io.Reader, size int) *Reader")
	fmt.Println("bufio.NewWriter(w io.Writer) *Writer")
	fmt.Println("bufio.NewWriterSize(w io.Writer, size int) *Writer")
	fmt.Println("bufio.NewReadWriter(r *Reader, w *Writer) *ReadWriter")
	fmt.Println("bufio.NewScanner(r io.Reader) *Scanner")
	fmt.Println("--------------------")
	fmt.Println("bytes.NewReader(b []byte) *Reader")
	fmt.Println("bytes.NewBuffer(buf []byte) *Buffer")
	fmt.Println("bytes.NewBufferString(s string) *Buffer")
	fmt.Println("--------------------")
	fmt.Println("strings.NewReader(s string) *Reader")
	fmt.Println("strings.NewReplacer(oldnew ...string) *Replacer")
	fmt.Println("var s strings.Builder")
	fmt.Println("--------------------------------------------")

	fmt.Println("var btb bytes.Buffer // 使用 bytes.Buffer 结构体实例化")
	var btb bytes.Buffer // 使用 bytes.Buffer 结构体
	btb.WriteString("hello")
	btb.WriteString(" ")
	btb.WriteString("world")
	// [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Println("btb.Bytes() // 返回一个长度为 b.Len() 的切片, 其中包含缓冲区的未读部分, 切片仅在下一次缓冲区修改之前有效", btb.Bytes())
	fmt.Println("btb.Cap() // 返回缓冲区底层字节切片的容量, 即为缓冲区数据分配的总空间", btb.Cap()) // 64
	fmt.Println("btb.Len() // 返回缓冲区中未读部分的字节数", btb.Len())                // 11
	fmt.Println(`
  var btb bytes.Buffer // 使用 bytes.Buffer 结构体
  btb.WriteString("hello")
  btb.WriteString(" ")
  btb.WriteString("world")`)
	fmt.Println("btb.String()", btb.String()) // hello world
	fmt.Println("--------------------")

	fmt.Println("bytes.NewBufferString(\"hello world\") 或 bytes.NewBuffer([]byte(\"hello world\"))")
	var b = bytes.NewBufferString("hello world") // 等价于下一行
	// var b = bytes.NewBuffer([]byte("hello world"))
	fmt.Print("b.ReadByte() ")
	fmt.Println(b.ReadByte()) // 104 <nil>
	// b.UnreadByte() // 撤销最近一次的读取操作的最后一个字节
	fmt.Println("b.Bytes() // 返回缓冲中未读部分的字节切片", b.Bytes())            // [101 108 108 111 32 119 111 114 108 100]
	fmt.Println("b.Cap() // 返回缓冲区底层字节切片的容量, 即为缓冲区数据分配的总空间", b.Cap()) // 16
	fmt.Println("b.Len() // 返回缓冲中未读部分的字节长度", b.Len())                // 10
	// b.Grow(n int)																							// 必要时增加缓冲区的容量, 以保证 n 字节的剩余空间
	// b.Reset()
	// 等价于 b.Truncate(0)
	b.Truncate(7)                                                    // 丢弃缓冲中除前 n 字节数据外的其它数据
	fmt.Println("b.Truncate(7) // 丢弃缓冲中除前 n 字节数据外的其它数据", b.String()) // ello wo
	fmt.Println("b.Next(5) // 返回未读取部分前 n 字节数据的切片并移动读取位置", b.Next(5)) // [101 108 108 111 32]
	fmt.Println("返回未读取部分的字节数据的字符串形式", b.String())                    // wo
	fmt.Println("--------------------")

	bb := bytes.NewReader([]byte("中国上下五千年"))
	z, zLen, _ := bb.ReadRune()                                                                       // 读取一个 UTF-8 编码的字符, 并返回其 Unicode 编码和字节数
	fmt.Printf("bytes 包 ReadRune 方法读取 \"中国上下五千年\" 的结果为 %v unicode值为 %d 大小为 %d\n", string(z), z, zLen) // 中	20013	3
	fmt.Println("bb.Size() // 返回基础字节切片的原始长度", bb.Size())                                              // 21
	fmt.Println("bb.Len() // 返回切片的未读部分的字节数", bb.Len())                                                // 18
	z2, _ := bb.ReadByte()
	fmt.Printf("bytes 包 ReadByte 方法读取 \"中国上下五千年\" 的结果为 %v unicode值为 %d\n", string(z2), z2) // å 229
	h, hLen, _ := bytes.NewReader([]byte("hello")).ReadRune()
	fmt.Printf("bytes 包 ReadRune 方法读取 \"hello\" 的结果为 %v unicode值为 %+v 大小为 %d\n", string(h), h, hLen) // h 104 1
	fmt.Println("--------------------------------------------")

	// strings.Builder 使用了建造者模式
	// 通过使用可变长度的缓冲区来存储字符串, 并在构建过程中动态地增加其大小, 以适应不断增长的字符串
	fmt.Println("var sbd strings.Builder // 使用 strings.Builder 结构体实例化")
	var sbd strings.Builder // 使用 strings.Builder 结构体实例化
	sbd.WriteString("你")
	sbd.WriteString("好")
	fmt.Println("sbd.Cap()", sbd.Cap()) // 8
	fmt.Println("sbd.Len()", sbd.Len()) // 6
	sbd.WriteRune('世')
	fmt.Println("sbd.String()", sbd.String()) // 你好世
	fmt.Println("--------------------")

	str := "hello world"
	sr := strings.NewReader(str)
	res1, err := sr.ReadByte() // 读取一个字节
	fmt.Printf("sr = %+v\n", sr)
	fmt.Println("sr.Size() // 返回基础字符串的原始长度", sr.Size()) // 11
	fmt.Println("sr.Len() // 返回字符串中未读部分的字节数", sr.Len()) // 10
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("strings 包 ReadByte 方法 res1 值为 %+v unicode值为 %d\n", string(res1), res1) // h 104
	fmt.Println("--------------------")

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println("strings.NewReplacer(\"<\", \"&lt;\", \">\", \"&gt;\")", r.Replace("This is <b>HTML</b>!")) // This is &lt;b&gt;HTML&lt;/b&gt;!
	fmt.Println("--------------------------------------------")

	bfr := bufio.NewReader(strings.NewReader(str))
	res2, err := bfr.ReadString('\n')
	fmt.Println("bfr.Size() // 返回以字节为单位基础缓冲区的大小", bfr.Size())       // 4096
	fmt.Println("bfr.Buffered() // 返回当前缓冲区可读取的字节数", bfr.Buffered()) // 0
	if err != nil && err != io.EOF {
		fmt.Println("err", err)
	}
	fmt.Printf("bufio 包 ReadString 方法 res2 值为 %v\n", res2) // hello world
	fmt.Println("--------------------")

	bfw := bufio.NewWriter(os.Stdout)
	bfw.WriteString("hello world")
	fmt.Println("bfw := bufio.NewWriter(os.Stdout)")
	fmt.Println("bfw.Size() // 返回以字节为单位基础缓冲区的大小", bfw.Size())       // 4096
	fmt.Println("bfw.Buffered() // 返回已写入当前缓冲区的字节数", bfw.Buffered()) // 11
	fmt.Println("--------------------")

	var bfrw = bufio.NewReadWriter(bufio.NewReader(strings.NewReader("hello world\n")), bufio.NewWriter(os.Stdout))
	fmt.Println("var bfrw = bufio.NewReadWriter(bufio.NewReader(strings.NewReader(\"hello world\\n\")), bufio.NewWriter(os.Stdout))")
	res3, _ := bfrw.ReadString('\n')
	fmt.Println("res3, _ := bfrw.ReadString('\\n')", res3) // hello world
	len, _ := bfrw.WriteString(res3)
	fmt.Println("len, _ := bfrw.WriteString(res3)", len) // 12
	fmt.Println("--------------------")

	bfs := bufio.NewScanner(strings.NewReader("AaBbCcDdEeFfGgHh"))
	bfs.Split(bufio.ScanBytes)
	for bfs.Scan() {
		fmt.Printf("bfs.Bytes() %q\t bfs.Text() %v\t\n", bfs.Bytes(), bfs.Text())
	}
	// bfs.Bytes() "A"  bfs.Text() A
	// bfs.Bytes() "a"  bfs.Text() a
	// bfs.Bytes() "B"  bfs.Text() B
	// bfs.Bytes() "b"  bfs.Text() b
	// bfs.Bytes() "C"  bfs.Text() C
	// bfs.Bytes() "c"  bfs.Text() c
	// bfs.Bytes() "D"  bfs.Text() D
	// bfs.Bytes() "d"  bfs.Text() d
	// bfs.Bytes() "E"  bfs.Text() E
	// bfs.Bytes() "e"  bfs.Text() e
	// bfs.Bytes() "F"  bfs.Text() F
	// bfs.Bytes() "f"  bfs.Text() f
	// bfs.Bytes() "G"  bfs.Text() G
	// bfs.Bytes() "g"  bfs.Text() g
	// bfs.Bytes() "H"  bfs.Text() H
	// bfs.Bytes() "h"  bfs.Text() h
	fmt.Println("--------------------------------------------")
}
