package main

import (
	"fmt"
	"reflect"
	"time"
)

func reflectChan() {
	fmt.Println("--------------reflectChan--------------")
	fmt.Println("type ChanDir int 表示通道类型的方向")
	fmt.Println("const (")
	fmt.Println("  RecvDir ChanDir = 1 << iota         // <-chan")
	fmt.Println("  SendDir                             // chan<-")
	fmt.Println("  BothDir         = RecvDir | SendDir // chan")
	fmt.Println(")")
	fmt.Println("--------------")
	fmt.Println("type SelectDir int 描述一个 SelectCase 的通信方向")
	fmt.Println("const (")
	fmt.Println("  _          SelectDir = iota")
	fmt.Println("  SelectSend SelectDir        // case Chan <- send")
	fmt.Println("  SelectRecv                  // case <-Chan")
	fmt.Println("  SelectDefault               // default")
	fmt.Println(")")
	fmt.Println("type SelectCase struct {")
	fmt.Println("  Dir SelectDir // case 的方向")
	fmt.Println("  Chan Value    // 使用的通道(收/发)")
	fmt.Println("  Send Value    // 用于发送的值")
	fmt.Println("}")
	fmt.Println(tab, "如果 Dir 是 SelectDefault, 该条 case 代表 default case, Chan 和 Send 字段必须是 Value 零值")
	fmt.Println(tab, "如果 Dir 是 SelectSend, 该条 case 代表一个发送操作, Chan 字段底层必须是一个 chan 类型, Send 的底层必须是可以直接赋值给该 chan 类型成员类型的类型,")
	fmt.Println(tab, "  如果 Chan 是 Value 零值, 则不管 Send 字段是不是零值, 该条 case 都会被忽略")
	fmt.Println(tab, "如果 Dir 是 SelectRecv, 该条 case 代表一个接收操作, Chan 字段底层必须是一个 chan 类型, Send 必须是一个 Value 零值,")
	fmt.Println(tab, "  如果 Chan 是 Value 零值, 该条 case 会被忽略, 但 Send 字段仍需是 Value 零值, 当该条 case 被执行时, 接收到的值会被 Select 返回")
	fmt.Println("--------------")
	fmt.Println("func ChanOf(dir ChanDir, t Type) Type")
	fmt.Println(tab, "// 返回元素类型为 t、方向为 dir 的通道类型.运行时GC强制将通道的元素类型的大小限定为64kb. 如果 t 的尺寸大于或等于该限制,本函数将会panic")
	fmt.Println("func MakeChan(typ Type, buffer int) Value")
	fmt.Println(tab, "// 创建一个元素类型为 typ、有 buffer 个缓存的通道类型的 Value 值")
	fmt.Println("func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)")
	fmt.Println(tab, "// 执行 cases 切片描述的 select 操作, 类似于 go 的 select 语句, 它会阻塞直到至少一条 case 可以执行, 从可执行的 case 中伪随机的选择一条并执行,")
	fmt.Println(tab, "   返回选择执行的 case 的索引, 如果执行的是接收 case 时, 会返回接收到的值 recv, 以及布尔值 recvOK 表示是否对应于通道中某次发送的值(用于区分通道关闭时接收到的零值)")
	fmt.Println(tab, "   Select 最多支持 65536 个case")
	fmt.Println("--------------")
	fmt.Println("func (v Value) Recv() (x Value, ok bool)")
	fmt.Println(tab, "// 从 v 持有的通道接收并返回一个值(的Value封装). 如果v的Kind不是Chan会panic. 方法会阻塞直到获取到值.")
	fmt.Println(tab, "   如果返回值x对应于某个发送到v持有的通道的值, ok为true; 如果因为通道关闭而返回,x为Value零值而ok为false")
	fmt.Println("func (v Value) TryRecv() (x Value, ok bool)")
	fmt.Println(tab, "// 尝试从v持有的通道接收一个值, 但不会阻塞. 如果v的Kind不是Chan会panic.")
	fmt.Println(tab, "   如果方法成功接收到一个值,会返回该值(的Value封装)和true; 如果不能无阻塞的接收到值,返回Value零值和false;")
	fmt.Println(tab, "   如果因为通道关闭而返回,返回值x是持有通道元素类型的零值的Value和false")
	fmt.Println("func (v Value) Send(x Value)")
	fmt.Println(tab, "// 向v持有的通道发送x持有的值. 如果v的Kind不是Chan,或者x的持有值不能直接赋值给v持有通道的元素类型,会panic")
	fmt.Println("func (v Value) TrySend(x Value) bool")
	fmt.Println(tab, "// 尝试向v持有的通道发送x持有的值,但不会阻塞.如果v的Kind不是Chan会panic.如果成功发送会返回true,否则返回false. x的持有值必须可以直接赋值给v持有通道的元素类型")
	fmt.Println("func (v Value) Close() // 关闭 v 持有的通道, 如果 v 的 Kind 不是 Chan 会panic")
	fmt.Println("--------------")

	fmt.Printf("reflect.SendDir 的类型为 %T 值为 %v\n", reflect.SendDir, reflect.SendDir) // reflect.ChanDir chan<-
	fmt.Printf("reflect.RecvDir 的类型为 %T 值为 %v\n", reflect.RecvDir, reflect.RecvDir) // reflect.ChanDir <-chan
	fmt.Printf("reflect.BothDir 的类型为 %T 值为 %v\n", reflect.BothDir, reflect.BothDir) // reflect.ChanDir chan
	fmt.Println("--------------")
	fmt.Printf("reflect.SelectSend 的类型为 %T 值为 %v\n", reflect.SelectSend, reflect.SelectSend)          // reflect.SelectDir 1
	fmt.Printf("reflect.SelectRecv 的类型为 %T 值为 %v\n", reflect.SelectRecv, reflect.SelectRecv)          // reflect.SelectDir 2
	fmt.Printf("reflect.SelectDefault 的类型为 %T 值为 %v\n", reflect.SelectDefault, reflect.SelectDefault) // reflect.SelectDir 3
	fmt.Println("----------------------------")

	var chanT = reflect.ChanOf(reflect.BothDir, reflect.TypeOf(250))
	fmt.Println("var chanT = reflect.ChanOf(reflect.BothDir, reflect.TypeOf(250))")
	fmt.Printf("chanT 的类型为 %T 值为 %v\n", chanT, chanT) // *reflect.rtype chan int
	var nc1 = reflect.MakeChan(chanT, 0)
	fmt.Println("var nc1 = reflect.MakeChan(chanT, 0)")
	fmt.Printf("nc1 的类型为 %T 值为 %v\n", nc1, nc1) // reflect.Value 0xc00001c120
	var nc2 = reflect.MakeChan(chanT, 5)
	fmt.Println("var nc2 = reflect.MakeChan(chanT, 5)")
	fmt.Printf("nc2 的类型为 %T 值为 %v\n", nc2, nc2) // reflect.Value 0xc00001e120

	var c1 = reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(struct{}{}))
	fmt.Println("var c1 = reflect.ChanOf(reflect.RecvDir, reflect.TypeOf(struct{}{}))")
	fmt.Printf("c1 的类型为 %T 值为 %v\n", c1, c1) // *reflect.rtype <-chan struct {}
	fmt.Println("----------------------------")

	var chan1 = reflect.MakeChan(chanT, 2)
	fmt.Println("var chan2 = reflect.MakeChan(chanT, 2)")
	var chan2 = reflect.MakeChan(chanT, 2)
	fmt.Println("var chan1 = reflect.MakeChan(chanT, 2)")
	succeed := chan1.TrySend(reflect.ValueOf(120))
	fmt.Printf("succeed := chan1.TrySend(reflect.ValueOf(120)) 尝试非阻塞发送值 %v\n", succeed) // true
	time.Sleep(time.Second * 2)
	v, ok := chan1.TryRecv()
	fmt.Printf("v, ok := chan1.TryRecv() 尝试非阻塞读取值 %v 状态 %v\n", v, ok) // 120 true
	time.Sleep(time.Second * 2)
	fmt.Println("--------------")
	succeed = chan2.TrySend(reflect.ValueOf(110))
	fmt.Printf("succeed = chan2.TrySend(reflect.ValueOf(110)) 尝试非阻塞发送值 %v\n", succeed) // true
	time.Sleep(time.Second * 2)
	v, ok = chan2.TryRecv()
	fmt.Printf("v, ok = chan2.TryRecv() 尝试非阻塞读取值 %v 状态 %v\n", v, ok) // 110 true
	time.Sleep(time.Second * 2)
	fmt.Println("--------------")

	var sc = []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: chan1, Send: reflect.ValueOf(nil)},
		{Dir: reflect.SelectRecv, Chan: chan2, Send: reflect.ValueOf(nil)},
		{Dir: reflect.SelectDefault, Chan: reflect.ValueOf(nil), Send: reflect.ValueOf(nil)},
	}
	// []reflect.SelectCase
	// [{Dir:2 Chan:<chan int Value> Send:<invalid Value>} {Dir:2 Chan:<chan int Value> Send:<invalid Value>} {Dir:3 Chan:<invalid Value> Send:<invalid Value>}]
	fmt.Printf("sc 的类型为 %T 值为 %+v\n", sc, sc)
	fmt.Println("--------------")

	go func() {
		for i := 0; i < 20; i++ {
			if i%2 == 0 {
				chan2.Send(reflect.ValueOf(i))
				continue
			}
			chan1.Send(reflect.ValueOf(i))
		}
		chan1.Close()
		chan2.Close()
	}()

	// go func() {
	// 	for {
	// 		v, ok := chan1.Recv()
	// 		if !ok {
	// 			break
	// 		}
	// 		fmt.Printf("second goroutine received value is %v\n", v)
	// 	}
	// }()
	time.Sleep(time.Second * 2)
	for {
		chosen, recv, ok := reflect.Select(sc)
		if !ok {
			// chosen = 0, recv = <invalid reflect.Value>, ok = false
			fmt.Printf("chosen = %v, recv = %v, ok = %t\n", chosen, recv, ok)
			break
		}
		/*
			chosen = 0, recv = 1, ok = true
			chosen = 1, recv = 0, ok = true
			chosen = 0, recv = 3, ok = true
			chosen = 1, recv = 2, ok = true
			chosen = 0, recv = 5, ok = true
			chosen = 1, recv = 4, ok = true
			chosen = 0, recv = 7, ok = true
			chosen = 1, recv = 6, ok = true
			chosen = 1, recv = 8, ok = true
			chosen = 1, recv = 10, ok = true
			chosen = 1, recv = 12, ok = true
			chosen = 0, recv = 9, ok = true
			chosen = 0, recv = 11, ok = true
			chosen = 0, recv = 13, ok = true
			chosen = 1, recv = 14, ok = true
			chosen = 0, recv = 15, ok = true
			chosen = 1, recv = 16, ok = true
			chosen = 0, recv = 17, ok = true
			chosen = 0, recv = 19, ok = true
			chosen = 0, recv = 0, ok = false
		*/
		fmt.Printf("chosen = %v, recv = %v, ok = %t\n", chosen, recv, ok)
	}
	fmt.Println("----------------------------")
}
