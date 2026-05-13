package main

import (
	"io/fs"
	"os"
	"testing"
)

func TestFS(t *testing.T) {
	t.Run("io/fs.FS", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")

		// 打开一个文件
		f, err := fsys.Open("testdata/Book1.xlsx")
		if err != nil {
			t.Fatalf("fsys.Open() file Error: %s\n", err.Error())
		}
		defer f.Close()

		// 读取文件
		data := make([]byte, 100)
		n, err := f.Read(data)
		if err != nil {
			t.Fatalf("f.Read() error: %s\n", err.Error())
		}
		t.Logf("f.Read(): %d\n%s\n", n, string(data[:n]))
	})
}

func TestDirEntry(t *testing.T) {
	t.Run("io/fs.DirEntry", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")

		// 读取当前目录
		entries, err := fs.ReadDir(fsys, ".")
		if err != nil {
			t.Fatalf("fs.ReadDir() Error: %s\n", err.Error())
		}
		for _, entry := range entries {
			t.Log("entry.Name()", entry.Name())
			t.Log("entry.IsDir()", entry.IsDir())
			t.Log("entry.Type()", entry.Type())
			info, err := entry.Info()
			if err == nil {
				t.Log("info.Name()", info.Name())
				t.Log("info.Size()", info.Size())
				t.Log("info.Mode()", info.Mode())
				t.Log("info.ModTime()", info.ModTime())
				t.Log("info.IsDir()", info.IsDir())
			}
			t.Log("---------")
		}
	})
}

func TestReadFileFS(t *testing.T) {
	t.Run("io/fs.ReadFileFS", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")

		// 第一种方式, 类型断言
		if readFileFS, ok := fsys.(fs.ReadFileFS); ok {
			f, err := readFileFS.ReadFile("main.go")
			if err != nil {
				t.Fatalf("readFileFS.ReadFile() error %s\n", err.Error())
			}
			t.Logf("readFileFS.ReadFile() len %d\n", len(f))
		}
		// 第二种方式, fs.ReadFile() 读取文件
		f, err := fs.ReadFile(fsys, "main.go")
		if err != nil {
			t.Fatalf("fs.ReadFile() error %s\n", err.Error())
		}
		t.Logf("fs.ReadFile() len %d\n", len(f))
	})
}

func TestReadDirFS(t *testing.T) {
	t.Run("io/fs.ReadDirFS", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")

		// 类型断言
		if readDirFS, ok := fsys.(fs.ReadDirFS); ok {
			entries, err := readDirFS.ReadDir(".")
			if err != nil {
				t.Fatalf("readDirFS.ReadDir() error: %s\n", err.Error())
			}
			t.Log("Directory content:")
			for _, entry := range entries {
				t.Log("entry.Name()", entry.Name())
			}
		}
	})
}

func TestStatFS(t *testing.T) {
	t.Run("io/fs.StatFS", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")

		// 第一种方式, 类型断言
		// if statFS, ok := fsys.(fs.StatFS); ok {
		// 	info, err := statFS.Stat("testdata/Book1.xlsx")
		// }
		// 第二种方式, fs.Stat() 获取文件信息
		info, err := fs.Stat(fsys, "testdata/Book1.xlsx")
		if err != nil {
			t.Fatalf("fs.Stat() error %s\n", err.Error())
		}
		t.Log("info.Name()", info.Name())
		t.Log("info.Size()", info.Size())
		t.Log("info.Mode()", info.Mode())
		t.Log("info.ModTime()", info.ModTime())
		t.Log("info.IsDir()", info.IsDir())
	})
}

func TestSubFS(t *testing.T) {
	t.Run("io/fs.SubFS", func(t *testing.T) {
		fsys := os.DirFS(".")

		// 使用 fs.Sub() 获取一个表示子目录的文件系统
		subFS, err := fs.Sub(fsys, "testdata")
		if err != nil {
			t.Fatalf("fs.Sub() error %s\n", err.Error())
		}
		// 第一种方式, 类型断言
		// if s, ok := subFS.(fs.SubFS); ok {
		// 	sub, err := s.Sub("testdata")
		// }
		// 第二种方式, fs.ReadDir 读取目录信息
		entries, err := fs.ReadDir(subFS, ".")
		if err != nil {
			t.Fatalf("fs.ReadDir() error %s\n", err.Error())
		}
		for _, entry := range entries {
			t.Log("entry.Name()", entry.Name())
		}
	})
}

func TestGlobFS(t *testing.T) {
	t.Run("io/fs.GlobFS", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")
		// 第一种方式, 类型断言
		// if globFS, ok := fsys.(fs.GlobFS); ok {
		// 	matches, err := globFS.Glob("*.go")
		// }
		// 第二种方式, fs.Glob() 通配符匹配文件
		matches, err := fs.Glob(fsys, "*.go")
		if err != nil {
			t.Fatalf("fs.Glob() error %s\n", err.Error())
		}
		t.Log("fs.Glob()", matches)
	})
}

func TestWalkDir(t *testing.T) {
	t.Run("io/fs.WalkDir", func(t *testing.T) {
		// 创建一个表示当前目录的文件系统
		fsys := os.DirFS(".")

		fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				t.Fatalf("fs.WalkDir() error %s\n", err.Error())
			}
			t.Logf("path %v dirEntry %v\n", path, d)
			return nil
		})
	})
}
