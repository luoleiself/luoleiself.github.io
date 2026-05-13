package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	tab := strings.Repeat(" ", 4)
	fmt.Println("链表: 一种物理存储单元上非连续、非顺序的存储容器, 由多个节点组成, 节点通过一些变量记录彼此之间的关系, 链表由多种实现方法, 如单链表、双链表、循环链表等. 链表没有具体元素类型的限制")
	fmt.Println(tab, "不必预先知道数据大小, 可以充分利用计算机内存空间, 实现灵活的内存动态管理")
	fmt.Println(tab, "失去了数组随机读取的优点, 增加了节点的指针域, 空间开销较大")
	fmt.Println("单向链表")
	fmt.Println(tab, "首元节点: 链表中第一个存储数据的节点")
	fmt.Println(tab, "头节点: 在 首元节点 之前, 指针域指向 首元节点, 头节点的数据域可以存储链表的 长度 或者其他信息")
	fmt.Println(tab, "头指针: 指向链表中的第一个节点的指针, 如果 链表 中有 头节点, 则 头指针 指向 头节点, 如果没有 头节点, 则 头指针 指向 首元节点")
	fmt.Println("初始化: list.New() 或者 list.List")
	list1 := list.New()
	var list2 list.List
	list1.PushBack("hello list1") // 在最后一个位置插入元素
	list1.PushBack("hello gg")    // 在最后一个位置插入元素
	fmt.Println(`
list1 := list.New()
var list2 list.List
list1.PushBack("hello list1") // 在最后一个位置插入元素
list1.PushBack("hello gg")    // 在最后一个位置插入元素
    `)
	fmt.Print("遍历 list1", tab)
	IterateList(*list1) // 结果示例 list1 = ["hello list1", "hello gg"] list2 = []
	fmt.Println("------------------------------")

	fmt.Println("PushFrontList(other *List) 创建链表other的拷贝, 添加 other 链表元素到头部, PushBackList(other *List) 创建链表other的拷贝, 添加 other 链表元素到尾部")
	list2.PushFrontList(list1) // 将链表 list1 插入到 list2 的第一个位置
	fmt.Println(`
    list2.PushFrontList(list1) // 将链表 list1 插入到 list2 的第一个位置
    `)
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例list2 = ["hello list1", "hello gg"]
	fmt.Println("------------------------------")

	fmt.Println("PushFront(v interface{}) *Element 在链表的第一个位置插入新元素并返回该新元素, PushBack(v interface{}) *Element 在链表的最后一个位置插入新元素并返回该新元素")
	list2.PushFront(1)            // 在第一个位置插入元素并返回该元素
	list2.PushBack(10)            // 在最后一个位置插入元素并返回该元素
	list2.PushFront("list.go")    // 在第一个位置插入元素并返回该元素
	list2.PushBack(1)             // 在最后一个位置插入元素并返回该元素
	list2.PushBack("hello world") // 在最后一个位置插入元素并返回该元素
	fmt.Println(`
    list2.PushFront(1)            // 在第一个位置插入元素并返回该元素
    list2.PushBack(10)            // 在最后一个位置插入元素并返回该元素
    list2.PushFront("list.go")    // 在第一个位置插入元素并返回该元素
    list2.PushBack(1)             // 在最后一个位置插入元素并返回该元素
    list2.PushBack("hello world") // 在最后一个位置插入元素并返回该元素
    `)
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例 list2 = ["list.go", 1, "hello list1", "hello gg", 10, 1, "hello world"]
	fmt.Println("----------------")

	fmt.Println("Init() *List 清空链表, Len() int 返回链表中元素的个数, Front() *Element 返回链表第一个元素或nil, Back() *Element 返回链表最后一个元素或nil")
	fmt.Println("获取链表的长度 list2.Len() ", list2.Len())     // 7
	fmt.Println("获取第一个元素 list2.Front() ", list2.Front()) // "list.go"
	fmt.Println("获取最后一个元素 list2.Back() ", list2.Back())  // "hello world"
	fmt.Println("----------------")

	fmt.Println("InsertBefore(v interface{}, mark *Element) *Element 将一个值为v的新元素插入到mark前面, 并返回生成的新元素. 如果mark不是list的元素, list不会被修改")
	fmt.Println("InsertAfter(v interface{}, mark *Element) *Element 将一个值为v的新元素插入到mark后面, 并返回新生成的元素. 如果mark不是list的元素, list不会被修改")
	el1 := list2.Front()                          // 获取第一个元素
	list2.InsertBefore("list2.InsertBefore", el1) // 在 el1 前面插入元素
	fmt.Println(`
    el1 := list2.Front()                          // 获取第一个元素
    list2.InsertBefore("list2.InsertBefore", el1) // 在 el1 前面插入元素
    `)
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例 list2 = ["list2.InsertBefore", "list.go", 1, "hello list1", "hello gg", 10, 1, "hello world"]
	fmt.Println("----------------")

	fmt.Println("MoveToFront(e *Element) 将元素e移动到链表的第一个位置, 如果e不是list的元素, list不会被修改")
	fmt.Println("MoveToBack(e *Element) 将元素e移动到链表的最后一个位置, 如果e不是list的元素, list不会被修改")
	el2 := list2.Front()  // 获取第一个元素
	list2.MoveToBack(el2) // el2 移动到链表末尾
	fmt.Println(`
    el2 := list2.Front()  // 获取第一个元素
    list2.MoveToBack(el2) // el2 移动到链表末尾
    `)
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例 list2 = ["list.go", 1, "hello list1", "hello gg", 10, 1, "hello world", "list2.InsertBefore"]
	fmt.Println("----------------")

	fmt.Println("MoveBefore(e, mark *Element) 将元素e移动到mark的前面。如果e或mark不是list的元素, 或者e==mark, list不会被修改")
	fmt.Println("MoveAfter(e, mark *Element) 将元素e移动到mark的后面。如果e或mark不是list的元素, 或者e==mark, list不会被修改")
	el3 := list2.Back().Prev() // 获取最后一个元素的上一个元素
	el4 := list2.Front()       // 获取第一个元素
	list2.MoveAfter(el3, el4)  // el3 移动到 el4后面
	fmt.Println(`
    el3 := list2.Back().Prev() // 获取最后一个元素的上一个元素
    el4 := list2.Front()       // 获取第一个元素
    list2.MoveAfter(el3, el4)  // el3 移动到 el4后面
    `)
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例 list2 = ["list.go", "hello world", 1, "hello list1", "hello gg", 10, 1, "list2.InsertBefore"]
	fmt.Println("----------------")

	fmt.Println("Remove(e *Element) interface{} 删除链表中的元素e, 并返回e.Value")
	list2.Remove(list2.Back().Prev()) // 移除链表中的元素, 倒数第二个
	fmt.Println(`list2.Remove(list2.Back().Prev()) // 移除链表中的元素, 倒数第二个`)
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例 list2 = ["list.go", "hello world", 1, "hello list1", "hello gg", 10, "list2.InsertBefore"]
	fmt.Println("----------------")

	fmt.Println("func (e *Element) Next() *Element 返回链表的后一个元素或者nil, func (e *Element) Prev() *Element 返回链表的前一个元素或者nil")
	el5 := list2.Back()                                  // 获取最后一个元素
	fmt.Println("list2.Back() ", list2.Back())           // "list2.InsertBefore"
	fmt.Println("el5.Prev() ", el5.Prev())               // 1
	fmt.Println("el5.Prev().Next() ", el5.Prev().Next()) // "list2.InsertBefore"
	fmt.Print("遍历 list2", tab)
	IterateList(list2) // 结果示例 list2 = ["list.go", "hello world", 1, "hello list1", "hello gg", 10, "list2.InsertBefore"]
	fmt.Print("遍历 list1", tab)
	IterateList(*list1)
	fmt.Println("------------------------------")
	fmt.Println(`
    // 遍历 list
    func IterateList(list list.List) {
        for el := list.Front(); el != nil; el = el.Next() {
            fmt.Print(el.Value, "\t")
        }
        fmt.Println()
    }
    `)
}

// 遍历 list
func IterateList(list list.List) {
	for el := list.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value, "\t")
	}
	fmt.Println()
}
