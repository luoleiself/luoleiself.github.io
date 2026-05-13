package main

import (
	"fmt"
	"syscall"
)

func syscallNote() {
	fmt.Println("syscall.SIGHUP", syscall.SIGHUP, "挂起(Hangup)信号, 通常在终端断开连接时发送给控制进程")
	fmt.Println("syscall.SIGINT", syscall.SIGINT, "中断信号, 通常由用户通过 Ctrl + C 发送, 终止当前正在运行的进程")
	fmt.Println("syscall.SIGQUIT", syscall.SIGQUIT, "退出信号, 通常由用户通过 Ctrl + \\ 发送, 终止程序并生成核心转储文件, 用于调试")
	fmt.Println("syscall.SIGILL", syscall.SIGILL, "非法指令信号, 当程序尝试执行非法指令时发送, 通知进程发生了非法操作导致程序崩溃")
	fmt.Println("syscall.SIGTRAP", syscall.SIGTRAP, "陷阱信号, 通常由调试器设置的断点触发, 用于调试目的, 允许调试器捕获断点事件")
	fmt.Println("syscall.SIGABRT", syscall.SIGABRT, "异常终止信号, 通常由 abort() 函数调用产生, 通知进程进行异常终止, 并生成核心转储文件")
	fmt.Println("syscall.SIGKILL", syscall.SIGKILL, "强制终止信号, 无法被捕获或忽略, 立即终止进程, 通常用于紧急情况下的进程终止")
	fmt.Println("syscall.SIGPIPE", syscall.SIGPIPE, "管道错误信号, 当试图写入一个已关闭的管道或套接字时发送, 通知进程管道或套接字操作失败")
	fmt.Println("syscall.SIGALRM", syscall.SIGALRM, "定时器信号, 通常由 alarm() 函数触发, 用于定时任务或超时处理")
	fmt.Println("syscall.SIGTERM", syscall.SIGTERM, "终止信号, 通常由 kill 命令发送, 请求进程正常终止, 允许进程清理资源")
}
