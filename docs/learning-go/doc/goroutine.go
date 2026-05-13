package main

import (
	"fmt"
	"runtime"
)

func goroutine() {
	fmt.Println("--------goroutine()----------")
	fmt.Println("线程是一个执行上下文, 它包含 CPU 执行指令流所需的所有信息。 假设您正在阅读一本书, 现在想休息一下, 但您希望在回来的时候能够从刚才停止的位置继续阅读。")
	fmt.Println("实现这一目的的一种方法是记下页码、行号和字号。这样您阅读一本书的执行上下文就可以用这 3 个数字表示。 如果你有一个室友, 你们学的是同样的专业, 她可以在你不使用这本书的时候拿走它, 然后从她先前停下的地方继续阅读。")
	fmt.Println("然后你后续可以把它拿回来, 从你原来暂停的地方继续阅读。 线程以相同的方式工作。 CPU 给您一种错觉, 即它正在同时进行多个计算。它通过在每个计算上分配一些时间片段来做到这一点。")
	fmt.Println("它可以做到这一点, 是因为它对每个计算都有一个执行上下文。就像您可以与朋友分享同一本书一样, 许多任务可以共享一个 CPU。")
	fmt.Println("协程跟线程是有区别的, 线程由CPU调度是抢占式的, 协程由用户态调度是协作式的, 一个协程让出CPU后, 才执行下一个协程")
	fmt.Println("------------------")

	fmt.Println("上下文切换: 为了保存任务切换时刻的基本信息, 当 CPU 重新执行任务的时候可以加载上下文下信息, 从当时退出的位置、状态重新开始执行任务. ")
	fmt.Println("上下文信息包括虚拟内存、栈、全局变量等用户态的资源, 也包括内核堆栈、寄存器等内核态的资源")
	fmt.Println("------------------")

	fmt.Println("以通信的方式共享内存(CSP), \"以共享内存的方式通信\"")
	fmt.Println("CSP并发模型通过 goroutine 和 channel 实现")
	fmt.Println("创建 goroutine 的语法格式")
	fmt.Println(tab, "go 函数名(参数列表)")
	fmt.Println(tab, "go func(参数列表){ /*函数体*/ }(调用参数列表)")
	fmt.Println("------------------")

	fmt.Println("goroutine 并发控制的方式")
	fmt.Println(tab, "1. 使用 sync.WaitGroup, 根据 goroutine 通过 wg.Add() 来记录已经开始的 goroutine 数量, 通过 wg.Wait() 来等待执行任务的 goroutine 的 wg.Done(), 实现同步的工作")
	fmt.Println(tab, "2. 使用 for/select + stop channel, 通过向 stop channel 中传递 stop signal 实现 goroutine 的结束")
	fmt.Println(tab, "3. 使用 context, 可以控制具有复杂层级关系的 goroutine 任务, 此时使用前两种方式会比较复杂, 使用 context 会更优雅")
	fmt.Println("------------------")

	fmt.Println("sysmon: Go 程序启动时创建的一个不依赖于任何 P 的自运行的独立线程, 主要作用监控 G 的运行并标识超时运行的 G")
	fmt.Println("--------------------")

	fmt.Println("go 协程(goroutine)是 go 中最基本的执行单元")
	fmt.Println(tab, "go 运行时会自动将和阻塞协程在同一线程上面的排队等候的其他协程移动到其他不同的、可运行的线程上面, 避免线程阻塞")
	fmt.Println(tab, "go 运行时有一个复杂的调度器, 能管理所有 goroutine 并为其分配执行时间, 这个调度器在操作系统之上, 将操作系统的线程与语言运行时的逻辑处理器绑定, ")
	fmt.Println(tab, "并在逻辑处理器上运行 goroutine. 调度器在任何给定的时间, 都会全面控制哪个 goroutine 要在哪个逻辑处理器上运行.")
	fmt.Println(tab, tab, "M 内核态线程, 每个 M 都代表了 1 个内核线程, go 调度器和 OS 调度器通过 M 结合起来, OS 调度器负责把内核线程分配到 CPU 的核上执行")
	fmt.Println(tab, tab, tab, "数量不固定(会有的休眠的 M 或不绑定 P 的 M), 最大 10000 个,")
	fmt.Println(tab, tab, tab, "当 M 数量大于 P 数量时, 多出来的 M 只有阻塞排队(回收或睡眠)等待绑定 P 才能执行")
	fmt.Println(tab, tab, "P 逻辑处理器, 代表了 M 所需的上下文环境, 负责衔接 M 和 G 的调度上下文")
	fmt.Println(tab, tab, tab, "在程序启动时创建, 并保存在数组中, 默认 GOMAXPROCS (最大 256) 个, 启动时固定, 一般不修改")
	fmt.Println(tab, tab, tab, "go 1.5 以后 GOMAXPROCS 默认为 系统 cpu 核心数")
	fmt.Println(tab, tab, tab, "P 的数量决定了系统内最大可并行的 G 的数量")
	fmt.Println(tab, tab, "G goroutine: 本质上也是一种轻量级的线程, 每个 G 都代表 1 个 goroutine, 包含调用栈, 任务函数, 调度信息, 状态等信息, 初始默认大小为 2KB")
	fmt.Println(tab, tab, tab, "有独立的栈空间, 共享程序的堆空间, 由用户态线程控制")
	fmt.Println(tab, "开销小, ")
	fmt.Println(tab, "调度性能好")
	fmt.Println(tab, "Go调度本质是把大量的 goroutine 分配到少量线程上去执行, 并利用多核并行, 实现更强大的并发")
	fmt.Println("------------------")

	fmt.Println("runtime.Version()", runtime.Version())                        // 1.22.3
	fmt.Println("runtime.GOMAXPROCS(0) // < 1 不更改当前设置", runtime.GOMAXPROCS(0)) // 8
	fmt.Println("runtime.NumCPU()", runtime.NumCPU())                          // 8
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())              // 1
	fmt.Println("runtime.GOROOT()", runtime.GOROOT())                          // D:\1.22.3
	fmt.Println("----------")
	var mm = &runtime.MemStats{} // 记录内存申请和分配的统计信息
	runtime.ReadMemStats(mm)     // 将内存申请和分配的统计信息写入 mm
	fmt.Println("内存申请和分配的统计信息 \t", mm)
	fmt.Println("----------")
	fmt.Println("描述某个调用栈序列申请和释放的活动对象等信息")
	var mr = runtime.MemProfileRecord{}                 // 描述某个调用栈序列申请和释放的活动对象等信息
	fmt.Println("mr.InUseBytes()", mr.InUseBytes())     // 0
	fmt.Println("mr.InUseObjects()", mr.InUseObjects()) // 0
	fmt.Println("mr.Stack()", mr.Stack())               // []
	fmt.Println("----------")
	fmt.Println("返回当前内存 profile中的记录数")
	var mrs = []runtime.MemProfileRecord{mr}  // MemProfileRecord 类型切片
	var n, ok = runtime.MemProfile(mrs, true) //  返回当前内存 profile中的记录数
	fmt.Println(ok, n)                        // true, 0
	fmt.Println("--------------------")

	fmt.Println("--------scheduler()----------")
	fmt.Println("调度器: go 调度器的本质是把大量的 goroutine 分配到少量线程上取执行, 并利用多核并行, 实现更强大的并发")
	fmt.Println("goroutine的调度是指程序代码按照一定的算法在适当的时候挑选出合适的goroutine并放到CPU上去运行的过程, 这些负责对goroutine进行调度的程序代码称为goroutine调度器")
	fmt.Println(tab, "调度协程: 每个逻辑处理器 P 中都有一个特殊的协程 g0, 其主要作用是执行协程调度")
	fmt.Println(tab, tab, "普通的协程 g 无差别地用于执行用户代码, 当 g 发生系统调用时, M 内将重新执行调度协程从 g 切换到 g0, 每个工作线程内部都在完成 g0->g->g0->g->g0")
	fmt.Println("------------------")

	fmt.Println("Go 指令调度过程")
	fmt.Println("程序启动时, 先初始化系统 CPU 核数的 P, 当创建一个 goroutine 时, 优先尝试放在本地队列中, 如果本地队列满了, 则会把本地队列的前半部分和这个新的 goroutine 一起放到全局队列中")
	fmt.Println("如果没有可用的 P 时, 新的 goroutine 加入到全局队列中, 如果获取到空闲的 P, 那么尝试唤醒(创建)一个 M, 并和 P 绑定开始执行")
	fmt.Println("当 P 的本地队列中没有 goroutine 时, 则尝试从全局队列中窃取一部分(min(len(GQ)/GOMAXPROCS + 1, len(GQ/2)))放到本地队列中执行, 这个过程是加锁的.")
	fmt.Println("当全局队列中也没有 goroutine 时, 则尝试从其他 P 的队列中窃取一半的 goroutine 放到本地队列执行")
	fmt.Println("当 G 发生 系统调用 时, P 断开和当前 M 的绑定, 尝试唤醒(创建)一个 M 继续执行,")
	fmt.Println("当 G 系统调用 结束后, M 尝试获取原来绑定的 P 继续执行, 如果获取不成功查看是否有空闲的 P, 有则绑定 P 继续执行 G,")
	fmt.Println("否则将这个 G 放入到全局队列中, 自己先进入空转状态然后进入睡眠等待唤醒")
	fmt.Println("--------------------")

	fmt.Println("当 M0 执行 G 遇到 主动让出(runtime.Gosched()),被动等待(I/O,channel,定时器),系统调用 时, M0 尝试唤起一个 M(可能是新建也可能来自线程池) 并释放 P1 等待 阻塞 的返回值.")
	fmt.Println("如果当前没有空闲的 M 时, P1 将加入空闲队列中等待 M 获取, 如果有空闲的 M1 时, M1 绑定 P1, 继续执行 P1 队列中其他的 goroutine")
	fmt.Println("当 阻塞 结束后, M0 会根据 G 查询一下上下文 P1, 如果不成功, M0 将 G 放到一个全局的 goroutine 队列中, 自己将进入线程池睡眠等待唤起")
	fmt.Println("P 会周期性的检查本地和全局的 goroutine 队列, 避免全局队列中的 go 任务等待时间过长")
	fmt.Println("如果本地和全局的 goroutine 队列中没有可执行的 G 并且其他 P 中队列有多余的 G 时, P 会窃取一部分 G 过来执行")
	fmt.Println("--------------------")

	fmt.Println(`
	mstart, Go程序初始化时, 第一个 M 是由 mstart 创建, 新的物理线程创建时, 调用的函数也是 mstart。
	startm, 当有新的 G 创建或者有 G 从 waiting 进入 running 且还有空闲的 P, 此时会调用 startm, 获取一个 M 和空闲的 P 绑定执行 G。
	newm, 当调用 startm 时, 如果没有空闲的 M 则会通过 newm 创建 M。
	stopm, 在2种情况下会执行 stopm, 一是当 M 绑定的 P 无可运行的 G 且无法从其它 P 窃取可运行的 G 时, M 先进入 spinning 状态, 然后退出。二是当 M 和 G 进入系统调用后, 长时间未退出, P 被 retake 且 M 找不到空闲的 P 绑定, 此时 M 会调用 stopm。
	spinning状态, 在 findrunnable 函数中, 会短暂进入 spinning 状态, 如果找不到可运行的 G 则调用 stopm`)
	fmt.Println("--------------------")
}
