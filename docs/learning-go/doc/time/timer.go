package main

import (
	"fmt"
	"time"
)

func TimerNote() {
	fmt.Println("--------------timer 一次性计时器---------------")
	fmt.Println("func After(d Duration) <-chan Time // After 会在另一线程经过时间段 d 后向返回值发送当时的时间. 等价于 NewTimer(d).C")
	fmt.Println("func AfterFunc(d Duration, f func()) *Timer // AfterFunc 另起一个go程等待时间段 d 过去, 然后调用f. 它返回一个Timer, 可以通过调用其 Stop 方法来取消等待和对 f 的调用")
	fmt.Println("func NewTimer(d Duration) *Timer // 创建一个Timer, 它会在最少过去时间段 d 后到期, 向其自身的 C 字段发送当时的时间")
	fmt.Println(tab, "func (t *Timer) Reset(d Duration) bool // Reset使t重新开始计时, (本方法返回后再)等待时间段d过去后到期.")
	fmt.Println(tab, tab, "如果调用时t还在等待中会返回真; 如果t已经到期或者被停止了会返回假; 如果 t 未过期, 必须先停止计时器")
	fmt.Println(tab, "func (t *Timer) Stop() bool // Stop 停止 Timer 的执行. 如果停止了t会返回真; 如果t已经被停止或者过期了会返回假. Stop不会关闭通道t.C, 以避免从该通道的读取不正确的成功")
	fmt.Println("-----------------------------")

	var t1 = time.NewTimer(time.Second * 2) // 创建一个 2 秒后触发的 timer
	<-t1.C
	fmt.Println(`方式1: time.NewTimer(d Duration) *Timer
var t1 = time.NewTimer(time.Second * 2) // 创建一个 2 秒后触发的 timer
<-t1.C`)
	fmt.Println("timer 结构体内部封装了一个只读 channel, 阻塞当前 goroutine 执行直到 timer 触发")
	fmt.Println("-----------------------------")

	fmt.Println("\n方式2: time.After(d Duration) <-chan Time")
	<-time.After(time.Second * 2) // 创建一个 2 秒后触发的timer
	fmt.Println(`<-time.After(time.Second * 2) // 创建一个 2 秒后触发的timer`)
	fmt.Println("-----------------------------")

	var ch = make(chan bool)
	var t3 = time.AfterFunc(time.Second*2, func() {
		fmt.Println("方式3: time.AfterFunc(d Duration, f func) *Timer")
		ch <- true
	})
	fmt.Println("t3=", t3) // &{<nil> {824633926912 967132597320400 0 0xd6d9c0 0xd82840 0 0 1}}
	select {
	case n := <-ch:
		fmt.Println("\n方式3: 2秒以后执行的 select case... 没有 default 语句时自动阻塞当前 goroutine", n)
	default:
		fmt.Println("select case 直接执行无阻塞, 有 default 语句时需要手动阻塞当前 goroutine  time.Sleep(time.Second * 3)")
		time.Sleep(time.Second * 3)
	}
	fmt.Println("-----------------")
	fmt.Println(`
var ch = make(chan bool)
var t3 = time.AfterFunc(time.Second*2, func() {
    fmt.Println("方式3: time.AfterFunc(d Duration, f func) *Timer")
    ch <- true
})
fmt.Println("t3=", t3) // &{<nil> {824633926912 967132597320400 0 0xd6d9c0 0xd82840 0 0 1}}
select {
case n := <-ch:
    fmt.Println("方式3: 2秒以后执行的 select case... 没有 default 语句时自动阻塞当前 goroutine", n)
default:
    fmt.Println("select case 直接执行无阻塞, 有 default 语句时需要手动阻塞当前 goroutine  time.Sleep(time.Second * 3)")
    time.Sleep(time.Second * 3)
}`)
	fmt.Println("-----------------------------")

	fmt.Println(`
var unix = time.Now().Unix()
var t4 = time.NewTimer(time.Second * 5)
if !t4.Stop() {
  tn4 := <-t4.C
  fmt.Println("tn4 := <-t4.C", tn4)
}
t4.Reset(time.Second * 10)
tn41 := <-t4.C
fmt.Printf("time.NewTimer().Reset() 重置计时器经过 %\d 秒后执行\n", time.Now().Unix()-unix)
fmt.Println("tn41 := <-tn41.C", tn41)`)
	fmt.Println("-----------------")
	var unix = time.Now().Unix()
	var t4 = time.NewTimer(time.Second * 5)
	if !t4.Stop() {
		tn4 := <-t4.C
		fmt.Println("tn4 := <-t4.C", tn4)
	}
	t4.Reset(time.Second * 10)
	tn41 := <-t4.C
	fmt.Printf("time.NewTimer().Reset() 重置一次性计时器经过 %d 秒后执行\n", time.Now().Unix()-unix)
	fmt.Println("tn41 := <-tn41.C", tn41)
	fmt.Println("-----------------------------")
}
