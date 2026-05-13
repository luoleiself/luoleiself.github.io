package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"
	"strings"
	"syscall"
	"time"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("os.Args 获取命令行参数")
	fmt.Println("os.StartProcess(name, argv, attr) 启动一个新进程")
	fmt.Println("os.GetPagesize() 获取操作系统的页面大小")
	// fmt.Println("os.Environ() 获取当前进程的环境变量", os.Environ())
	hostname, _ := os.Hostname()
	fmt.Println("os.Hostname", hostname)
	fmt.Println("os.Getuid()", os.Getuid())
	fmt.Println("os.Getgid()", os.Getgid())
	fmt.Println("os.Geteuid() 获取调用进程的用户的 ID", os.Geteuid())
	fmt.Println("os.Getegid()", os.Getegid())
	groups, _ := os.Getgroups()
	fmt.Println("os.Getgroups()", groups)
	fmt.Println("os.Getpid() 获取当前进程的进程 ID", os.Getpid())
	fmt.Println("os.Getppid()", os.Getppid())
	dir, _ := os.Getwd()
	fmt.Println("os.Getwd() 获取当前进程的工作目录", dir)
	fmt.Println("os.Getenv()", os.Getenv("PATH"))
	fmt.Println("------------------------------------")

	users, _ := user.Current()
	fmt.Printf("user.Current() %+v\n", users)
	fmt.Print("users.GroupIds()")
	fmt.Println(users.GroupIds())
	fmt.Println()

	u, _ := user.Lookup("user")
	fmt.Printf("user.Lookup(\"user\") %+v\n", u)
	uById, _ := user.LookupId("3000")
	fmt.Printf("user.LookupId(\"3000\") %+v\n", uById)
	fmt.Println()

	g, _ := user.LookupGroup("user")
	fmt.Printf("user.LookupGroup(\"user\") %+v\n", g)
	gById, _ := user.LookupGroupId("3000")
	fmt.Printf("user.LookupGroupId(\"3000\") %+v\n", gById)
	fmt.Println()
	fmt.Println("------------------------------------")

	fmt.Println("os.Chdir() 切换当前进程的工作目录")
	fmt.Println("func Open(name string) (*File, error)")
	fmt.Println(tab, "Open 方法以 os.O_RDONLY 只读模式打开文件, 如果不存在则报错 *PathError.")
	fmt.Println(tab, "只能用于读取文件内容, 不支持向文件写入内容")
	fmt.Println("func Create(name string) (*File, error)")
	fmt.Println(tab, "只能用于读写文件内容，不支持向文件追加内容")
	fmt.Println(tab, "Create 方法以 os.O_RDWR 可读可写模式创建或者截断文件, 如果文件存在则截断, 如果文件不存在则以 umask 创建文件.")
	fmt.Println("func OpenFile(name string, flag int, perm FileMode) (*File, error) // 通常使用 Open 或 Create 方法代替")
	fmt.Println("---------------------")
	fmt.Println(`
  // file, err := os.Open("test.txt")
  file, err := os.Create("test.txt")
  // file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
  if err != nil {
    fmt.Println("os.OpenFile() err =", err)
  } else {
    n, err := file.Write([]byte("The first line!\n"))
    if err != nil {
      fmt.Println("file.Write() err =", err)
    }
    fmt.Printf("file.Write() write %\d byte\\n", n)
  }`)
	fmt.Println("------------------------------------")

	fmt.Println("func Lstat(name string) (FileInfo, error) // 返回一个描述 name 指定的文件对象的 FileInfo, 如果是符号链接, 则表示符号链接的信息")
	fmt.Println("func Stat(name string) (FileInfo, error) // 返回一个描述 name 指定的文件对象的 FileInfo")
	fmt.Println("------------------------------------")

	filePathNote()
	fmt.Println("------------------------------------")

	log.Println("os/signal.Notify...")
	stop := make(chan os.Signal, 1)

	fmt.Println("os.Interrupt == syscall.SIGINT", os.Interrupt == syscall.SIGINT)
	fmt.Println("os.Kill == syscall.SIGKILL", os.Kill == syscall.SIGKILL)
	// signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)
	signal.Notify(stop) // all signals will be sent to the channel.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	select {
	case sig := <-stop:
		cancel()
		log.Printf("signal.Notify() %s\n", sig)
	case <-ctx.Done():
		log.Println("context.WithTimeout()", ctx.Err())
	}
	log.Println("signal.Notify() end")
	fmt.Println("------------------------------------")

	fmt.Println("syscall")
	syscallNote()
}
