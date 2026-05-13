package main

import (
	"fmt"
	"path/filepath"
)

func filePathNote() {
	fmt.Println("----------filePathNote()-------------")
	var path = "E:\\workspace_test\\learning-go\\doc\\testing\\testdata\\test.txt"
	var path2 = "/root/workspace/learning-go/testing/testdata/test.txt"

	fmt.Println(path, "\n", path2)
	fmt.Println("----------")

	s1, _ := filepath.Abs(path)
	s2, _ := filepath.Abs(path2)
	fmt.Println("on windows: Abs", s1) // E:\workspace_test\learning-go\doc\testing\testdata\test.txt
	fmt.Println("on unix: Abs", s2)    // E:\root\workspace\learning-go\testing\testdata\test.txt
	fmt.Println("----------")
	fmt.Println("on windows: Base", filepath.Base(path)) //  test.txt
	fmt.Println("on unix: Base", filepath.Base(path2))   //  test.txt
	fmt.Println("----------")
	fmt.Println("on windows: Ext", filepath.Ext(path)) //  .txt
	fmt.Println("on unix: Ext", filepath.Ext(path2))   //  .txt
	fmt.Println("----------")
	fmt.Println("on windows: isAbs", filepath.IsAbs(path)) // true
	fmt.Println("on unix: isAbs", filepath.IsAbs(path2))   // false
	fmt.Println("----------")
	// 1.20 add
	fmt.Println("on windows: IsLocal", filepath.IsLocal(path)) // false
	fmt.Println("on unix: IsLocal", filepath.IsLocal(path2))   // false
	fmt.Println("----------")
	fmt.Println("filepath.Join(\"a\", \"b\", \"c\") Join", filepath.Join("a", "b", "c"))   //  a\b\c
	fmt.Println("filepath.Join(\"a\", \"b/c\") Join", filepath.Join("a", "b/c"))           //  a\b\c
	fmt.Println("filepath.Join(\"a/b\", \"c\") Join", filepath.Join("a/b", "c"))           //  a\b\c
	fmt.Println("filepath.Join(\"a/b\", \"/c\") Join", filepath.Join("a/b", "/c"))         //  a\b\c
	fmt.Println("filepath.Join(\"a/b\", \"../xyz\") Join", filepath.Join("a/b", "../xyz")) //  a\xyz
	fmt.Println("----------")
	fmt.Println("on windows: SplitList", filepath.SplitList(path)) //  [E:\workspace_test\learning-go\doc\testing\testdata\test.txt]
	fmt.Println("on unix: SplitList", filepath.SplitList(path2))   // [/root/workspace/learning-go/testing/testdata/test.txt]
	fmt.Println("----------")
	dir1, file1 := filepath.Split(path)
	dir2, file2 := filepath.Split(path2)
	fmt.Println("on windows:Split dir", dir1, "\non windows:Split file", file1) // E:\workspace_test\learning-go\doc\testing\testdata\   test.txt
	fmt.Println("on unix:Split dir", dir2, "\non unix:Split file", file2)       //  /root/workspace/learning-go/testing/testdata/   test.txt
	fmt.Println("----------")
	fmt.Println("on windows: FromSlash", filepath.FromSlash(path)) //  E:\workspace_test\learning-go\doc\testing\testdata\test.txt
	fmt.Println("on unix: FromSlash", filepath.FromSlash(path2))   // \root\workspace\learning-go\testing\testdata\test.txt
	fmt.Println("----------")
	fmt.Println("on windows: ToSlash", filepath.ToSlash(path)) //  E:/workspace_test/learning-go/doc/testing/testdata/test.txt
	fmt.Println("on unix: ToSlash", filepath.ToSlash(path2))   // /root/workspace/learning-go/testing/testdata/test.txt
	fmt.Println("----------")
	fmt.Println("on windows: VolumeName", filepath.VolumeName(path)) // E:
	fmt.Println("on unix: VolumeName", filepath.VolumeName(path2))   //
	fmt.Println("----------")
	fmt.Println("on windows: Dir", filepath.Dir(path)) //  E:\workspace_test\learning-go\doc\testing\testdata
	fmt.Println("on unix: Dir", filepath.Dir(path2))   // \root\workspace\learning-go\testing\testdata
	fmt.Println("----------")
	fmt.Println("------------------------------------")
}
