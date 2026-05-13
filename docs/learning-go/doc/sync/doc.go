package main

import (
	"fmt"
)

func readme() {
	fmt.Println("-------------readme()-------------")
	fmt.Println("死锁: 指两个或者两个以上的进程(或线程)在执行过程中, 因争夺资源而造成的一种互相等待的现象, 若无外力作用, 它们都将无法推进下去")
	fmt.Println(tab, "发生条件")
	fmt.Println(tab, tab, "互斥条件: 线程对资源的访问时排他性的, 如果一个线程占用了某资源, 则其他线程必须处于等待状态, 直到该资源被释放")
	fmt.Println(tab, tab, "请求和保持条件: 线程 T1 至少已经保持了一个资源 R1 占用, 但又请求使用另一个资源 R2, 此时, 资源 R2 被其他线程 T2 占用,")
	fmt.Println(tab, tab, "  于是该线程 T1 也必须等待, 但又对自己保持的资源 R1 不释放")
	fmt.Println(tab, tab, "不剥夺条件: 线程已获得的资源, 在未使用完之前, 不能被其他线程剥夺, 只能在使用完以后由自己释放")
	fmt.Println(tab, tab, "环路等待条件: 在死锁发生时, 必然存在一个 \"进程-资源环形链\", p0 等待 p1 占用的资源, p1 等待 p2 占用的资源, p2 等待 p0 占用的资源")
	fmt.Println(tab, "解决办法")
	fmt.Println(tab, tab, "如果并发查询多个表, 约定访问顺序")
	fmt.Println(tab, tab, "在同一个事务中, 尽可能做到一次锁定获取所需要的资源")
	fmt.Println(tab, tab, "对于容易产生死锁的业务场景, 尝试升级锁颗粒度, 使用表级锁")
	fmt.Println(tab, tab, "采用分布式事务锁或者使用乐观锁")
	fmt.Println("--------------------------")
	fmt.Println("活锁: 另一种形式的活跃性问题, 该问题尽管不会阻塞线程, 但也不能继续执行, 因为线程将不断重复同样的操作, 而且总会失败. ")
	fmt.Println(tab, "活锁和死锁的区别在于, 处于活锁的实体是在不断的改变状态, 处于死锁的实体表现为等待, 活锁有可能自行解开, 死锁则不能")
	fmt.Println("--------------------------")
	fmt.Println("饥饿: 指一个可运行的进程尽管能继续执行, 但被调度器无限期地忽视, 而不能被调度执行的情况")

	fmt.Println("--------------------------")
	fmt.Println("sync 包总结:")
	fmt.Println("  sync 包提供了基本的同步基元, 如互斥锁. 除了 Once 和 WaitGroup 类型, 大部分都是适用于低水平程序线程, 高水平的同步使用channel通信更好一些")
	fmt.Println("-------------")

	fmt.Println("Once 是只执行一次动作的对象, 提供 Do 方法接收一个回调函数 ")
	fmt.Println(tab, "func (o *sync.Once) Do(f func) 当且仅当第一次被调用时才执行函数 f, 即使每次调用 Do 提供的f值不同, 内部使用 Mutex 互斥锁实现了同步")
	fmt.Println("-------------")

	fmt.Println("WaitGroup 用于等待一组线程的结束. 父线程调用 Add 方法来设定应等待的线程的数量. 每个被等待的线程在结束时应调用 Done 方法. 同时, 主线程里可以调用 Wait 方法阻塞至所有线程结束")
	fmt.Println(tab, "func (wg *sync.WaitGroup) Add(delta int) 接收 int 参数设定应等待的线程的数量, 方法应在创建新的线程或者其他应等待的事件之前调用")
	fmt.Println(tab, "func (wg *sync.WaitGroup) Done() 减少 WaitGroup 计数器的值, 应在线程的最后执行")
	fmt.Println(tab, tab, "wg.Done() 和 wg.Add(-1) 是完全等价的")
	fmt.Println(tab, "func (wg *sync.WaitGroup) Wait() \033[1;32m阻塞\033[0m直到 WaitGroup 计数器减为0")
	fmt.Println("-------------")

	fmt.Println("Mutex 互斥锁, 零值为解锁状态. Mutex 类型的锁和线程无关, 可以由不同的线程加锁和解锁")
	fmt.Println(tab, "func (m *sync.Mutex) Lock() 锁住 m, 如果 m 已经加锁, 则\033[1;32m阻塞\033[0m直到 m 解锁")
	fmt.Println(tab, "func (m *sync.Mutex) Unlock() 解锁 m, 如果 m 未加锁会导致运行时错误. 锁和线程无关, 可以由不同的线程加锁和解锁")
	fmt.Println("-------------")

	fmt.Println("RWMutex 读写锁(多读单写锁). 零值为解锁状态. 读写锁可以被同时多个读取者持有或唯一个写入者持有. RWMutex 类型的锁也和线程无关, 可以由不同的线程加读取锁/写入和解读取锁/写入锁")
	fmt.Println(tab, "func (rw *sync.RWMutex) Lock() 将 rw 锁定为写入状态, 禁止其他线程读取或者写入, \033[1;32m阻塞\033[0m直到所有的读写锁解锁")
	fmt.Println(tab, "func (rw *sync.RWMutex) Unlock() 解除 rw 的写入锁状态, 如果 m 未加写入锁会导致运行时错误")
	fmt.Println(tab, "func (rw *sync.RWMutex) RLock() 将 rw 锁定为读取状态, 禁止其他线程写入, 但不禁止读取")
	fmt.Println(tab, "func (rw *sync.RWMutex) RUnlock() 解除 rw 的读取锁状态, 如果 m 未加读取锁会导致运行时错误")
	fmt.Println(tab, "func (rw *sync.RWMutex) RLocker() Locker 返回一个互斥锁, 通过调用 rw.Rlock 和 rw.Runlock 实现了 Locker 接口")
	fmt.Println("-------------")

	fmt.Println("Pool: 是一个可以分别存取的临时对象的集合")
	fmt.Println(tab, "Pool 中保存的任何 item 都可能随时不做通告的的释放掉, 如果 Pool 持有该对象的唯一引用, 这个对象就可能被回收, 所以不适合用于存放诸如 socket 长连接或数据库连接的对象")
	fmt.Println(tab, "Pool 可以安全的被多个线程同时使用")
	fmt.Println(tab, "Pool 的目的是缓存申请但未使用的 item 用于之后的重用, 以减轻 GC 的压力, 即让创建高效而线程安全的空闲列表更容易, 但 Pool 不适用于所有空闲列表")
	fmt.Println(tab, "Pool 的合理用法是用于管理一组静静的被多个独立并发线程共享并可能重用的临时 item, Pool 提供了让多个线程分摊内存申请消耗的方法")
	fmt.Println(tab, "func (p *Pool) Get() interface{} 从池中选择任意一个 item, 删除其在池中的引用计数, 并提供给调用者")
	fmt.Println(tab, "func (p *Pool) Put(x interface{}) 将 x 放入池中")
	fmt.Println("-------------")

	fmt.Println("Cond: 实现了一个条件变量, 一个线程集合地, 供线程等待或宣布某事件的发生, 用来协调多个 goroutine 之间的同步,")
	fmt.Println("  当共享资源的状态发生变化的时候, 可以通过条件变量来通知所有等待的 goroutine 区重新获取共享资源")
	fmt.Println("  每个 Cond 实例都有一个相关的锁(一般由 Mutex 或 RWMutex类型的值), 必须在改变条件时或者调用 Wait 方法时保持锁定多个 goroutine 等待, 1 个 goroutine 通知事件发生")
	fmt.Println(tab, "func NewCond(l Locker) *Cond // 使用锁创建一个 *Cond")
	fmt.Println(tab, "func (c *Cond) Broadcast() // 唤醒所有等待 c 的 goroutine, 在调用本方法时, 建议(并非必须)保持 c.L 的锁定")
	fmt.Println(tab, "func (c *Cond) Signal() // 唤醒等待 c 的一个线程(如果存在), 在调用本方法时, 建议(并非必须)保持 c.L 的锁定")
	fmt.Println(tab, "func (c *Cond) Wait() // 自行解锁 c.L 并\033[1;32m阻塞\033[0m当前线程, 在之后线程恢复执行时, Wait 方法会在返回前锁定 c.L")
	fmt.Println(tab, "  和其他系统不同, Wait 除非被 Broadcast 或者 Signal 唤醒, 不会主动返回")
	fmt.Println(tab, "  调用 Wait 方法之前需要先获取锁, 如果不加锁, 有可能会出现竞态条件")
	fmt.Println("-------------")

	fmt.Println("--------------------------")
}
