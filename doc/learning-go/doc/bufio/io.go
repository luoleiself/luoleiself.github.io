package main

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	_ "io/ioutil"
	"os"
)

type FileMode uint32

const (
	ModeDir FileMode = 1 << (32 - 1 - iota)
	ModeAppend
	ModeExclusive
	ModeTemporary
	ModeSymlinks
)

func (m FileMode) String() string {
	const str = "dalTLDpSugct?"
	var buf [32]byte // Mode is uint32.
	w := 0
	for i, c := range str {
		if m&(1<<uint(32-1-i)) != 0 {
			buf[w] = byte(c)
			w++
		}
	}
	if w == 0 {
		buf[w] = '-'
		w++
	}
	const rwx = "rwxrwxrwx"
	for i, c := range rwx {
		if m&(1<<uint(9-1-i)) != 0 {
			buf[w] = byte(c)
		} else {
			buf[w] = '-'
		}
		w++
	}
	return string(buf[:w])
}

func ioNote() {
	fmt.Println("----------------ioNote()----------------")
	fmt.Println("io/fs", "定义了文件系统的抽象接口, 提供了一种统一访问不同类型的文件系统, 包括本地文件系统、内存文件系统、zip文件等")
	fmt.Println("  抽象文件系统, io/fs 包定义了一组接口, 用于描述文件系统的基本操作, 如打开文件、读取目录等, 通过这些接口可以编写与具体文件系统无关的代码")
	fmt.Println("  统一访问方式, 无论底层文件系统是什么类型, 只要实现了 io/fs 包定义的接口, 就可以使用相同的代码进行访问")
	fmt.Println("  提高代码可测试性, 通过使用 io/fs 包, 可以方便地 mock 文件系统, 从而提高代码的可测试性")
	fmt.Println("fs.ModeDir", fs.ModeDir)
	fmt.Println("fs.ModeAppend", fs.ModeAppend)
	fmt.Println("fs.ModeExclusive", fs.ModeExclusive)
	fmt.Println("fs.ModeTemporary", fs.ModeTemporary)
	fmt.Println("fs.ModeSymlink", fs.ModeSymlink)
	fmt.Println("fs.ModeDevice", fs.ModeDevice)
	fmt.Println("------------------------------")
	var b bytes.Buffer
	r := io.TeeReader(os.Stdin, &b)
	io.ReadAll(r)
	fmt.Println(b.String())

	fmt.Println(ModeDir, ModeAppend, ModeExclusive, ModeTemporary, ModeSymlinks)
	fmt.Println("------------------------------")

	fmt.Println("1<<32 - 1", 1<<32-1)
	fmt.Println(2147483648, 1<<32-1, "\n2147483648 & 1<<32-1", 2147483648&1<<32-1)
	fmt.Println("--------------")
	fmt.Println("io/fs 包定义 FileMode 类型部分源码")
	var str = "dalTLDpSugct?"
	var buf [32]byte
	var w = 0
	for i, c := range str {
		fmt.Println(i, c, 1<<(32-1-i), ModeDir&(1<<(32-1-i)))
		if ModeDir&(1<<(32-1-i)) != 0 {
			buf[w] = byte(c)
			w++
		}
	}
	if w == 0 {
		buf[w] = '-'
		w++
	}
	fmt.Printf("buf 类型为 %T 值为 %v\n", buf, buf)
	fmt.Println("------------------------------")

	fmt.Println("func WriteFile(filename string, data []byte, perm fs.FileMode) error")
	fmt.Println(tab, "As of Go 1.16, this function simply calls os.WriteFile.")
	fmt.Println(`
err := ioutil.WriteFile("test.txt", []byte("hello ioutil.WriteFile"), 0664)
// err := ioutil.WriteFile("test.txt", []byte("hello ioutil.WriteFile"), fs.ModePerm)
if err != nil {
  fmt.Println("ioutil.WriteFile() write err =", err)
} else {
  fmt.Println("ioutil.WriteFile() write file successful!")
}`)
	fmt.Println("------------------------------")

	fmt.Println(`
// Go 1.16 以前 io/ioutil 包的 WriteFile 方法
// WriteFile writes data to a file named by filename.
// If the file does not exist, WriteFile creates it with permissions perm
// (before umask); otherwise WriteFile truncates it before writing, without changing permissions.
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// Go 1.16 以后 io/ioutil 包的 WriteFile 方法
// As of Go 1.16, this function simply calls os.WriteFile.
func WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(filename, data, perm)
}`)
	fmt.Println("------------------------------")

	fmt.Println("io/fs 包部分源码")
	fmt.Println(`
package fs

import "io"

// An FS provides access to a hierarchical file system.
//
// The FS interface is the minimum implementation required of the file system.
// A file system may implement additional interfaces,
// such as ReadFileFS, to provide additional or optimized functionality.
type FS interface {
  // Open opens the named file.
  //
  // When Open returns an error, it should be of type *PathError
  // with the Op field set to "open", the Path field set to name,
  // and the Err field describing the problem.
  //
  // Open should reject attempts to open names that do not satisfy
  // ValidPath(name), returning a *PathError with Err set to
  // ErrInvalid or ErrNotExist.
  Open(name string) (File, error)
}

// ReadFileFS is the interface implemented by a file system
// that provides an optimized implementation of ReadFile.
type ReadFileFS interface {
	FS

	// ReadFile reads the named file and returns its contents.
	// A successful call returns a nil error, not io.EOF.
	// (Because ReadFile reads the whole file, the expected EOF
	// from the final Read is not treated as an error to be reported.)
	//
	// The caller is permitted to modify the returned byte slice.
	// This method should return a copy of the underlying data.
	ReadFile(name string) ([]byte, error)
}
// ReadFile reads the named file from the file system fs and returns its contents.
// A successful call returns a nil error, not io.EOF.
// (Because ReadFile reads the whole file, the expected EOF
// from the final Read is not treated as an error to be reported.)
//
// If fs implements ReadFileFS, ReadFile calls fs.ReadFile.
// Otherwise ReadFile calls fs.Open and uses Read and Close
// on the returned file.
func ReadFile(fsys FS, name string) ([]byte, error) {
	// 如果实现了 ReadFileFS 接口, ReadFile 方法将会调用 fs.ReadFile, 否则 ReadFile 将调用 fs.Open 并使用 Read 和 Close 方法在返回的文件上
	// 此处使用类型断言
	if fsys, ok := fsys.(ReadFileFS); ok { 
		return fsys.ReadFile(name)
	}

	file, err := fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var size int
	if info, err := file.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}

	data := make([]byte, 0, size+1)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := file.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}

// ReadDirFS is the interface implemented by a file system
// that provides an optimized implementation of ReadDir.
type ReadDirFS interface {
	FS

	// ReadDir reads the named directory
	// and returns a list of directory entries sorted by filename.
	ReadDir(name string) ([]DirEntry, error)
}

// ReadDir reads the named directory
// and returns a list of directory entries sorted by filename.
//
// If fs implements ReadDirFS, ReadDir calls fs.ReadDir.
// Otherwise ReadDir calls fs.Open and uses ReadDir and Close
// on the returned file.
func ReadDir(fsys FS, name string) ([]DirEntry, error) {
	// 如果实现了 ReadDirFS 接口, ReadDir 方法将会调用 fs.ReadDir, 否则 ReadDir 将调用 fs.Open 并使用 ReadDir 和 Close 方法在返回的文件上
	// 此处使用类型断言
	if fsys, ok := fsys.(ReadDirFS); ok {
		return fsys.ReadDir(name)
	}

	file, err := fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dir, ok := file.(ReadDirFile)
	if !ok {
		return nil, &PathError{Op: "readdir", Path: name, Err: errors.New("not implemented")}
	}

	list, err := dir.ReadDir(-1)
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, err
}

// dirInfo is a DirEntry based on a FileInfo.
type dirInfo struct {
	fileInfo FileInfo
}

func (di dirInfo) IsDir() bool {
	return di.fileInfo.IsDir()
}

func (di dirInfo) Type() FileMode {
	return di.fileInfo.Mode().Type()
}

func (di dirInfo) Info() (FileInfo, error) {
	return di.fileInfo, nil
}

func (di dirInfo) Name() string {
	return di.fileInfo.Name()
}

// FileInfoToDirEntry returns a DirEntry that returns information from info.
// If info is nil, FileInfoToDirEntry returns nil.
func FileInfoToDirEntry(info FileInfo) DirEntry {
	if info == nil {
		return nil
	}
	return dirInfo{fileInfo: info}
}`)
	fmt.Println("------------------------------")
}
