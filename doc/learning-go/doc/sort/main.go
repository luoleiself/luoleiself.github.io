package main

import (
	"fmt"
	"sort"
	"strings"
)

var tab = strings.Repeat(" ", 2)

func main() {
	fmt.Println("排序")
	fmt.Println(tab, "IntSlice, Float64Slice, StringSlice 默认提供三个类型已实现 Interface 接口(Len,Less,Swap方法)")
	fmt.Println(tab, "func Sort(data Interface) // 排序 data, 参数类型必须是实现了 Interface 接口(Len,Less,Swap方法)的类型, 如果不是, 需要先进行类型转换")
	fmt.Println(tab, "func Stable(data Interface) // 排序 data, 并保证排序的稳定性, 相等元素的相对次序不变,参数类型必须是实现了 Interface 接口(Len,Less,Swap方法)的类型, 如果不是, 需要先进行类型转换")
	fmt.Println(tab, "func IsSorted(data Interface) bool // 判断 data 是否已排序, 参数类型必须是实现了 Interface 接口(Len,Less,Swap方法)的类型, 如果不是, 需要先进行类型转换")
	fmt.Println(tab, "func Search(n int,f func(int) bool) int // 采用二分法搜索找到[0, n)区间内最小的满足 f(i)==true 的值 i, 如果未找到则返回该值插入的索引")
	fmt.Println(tab, "---------整型排序-------")
	fmt.Println(tab, "func Ints(x []int) // 升序排序")
	fmt.Println(tab, "func IntsAreSorted(x []int) bool")
	fmt.Println(tab, "func SearchInts(a []int, x int) int // 查询给定值在切片中出现的下标, 如果不存在则返回插入位置的下标")
	fmt.Println(tab, "type IntSlice []int // 声明本地类型 IntSlice")
	fmt.Println(tab, tab, "func (p IntSlice) Search(x int) int // IntSlice 类型的方法, 普通的 int 切片使用此方法需要先将类型转为 IntSlice 类型, 见下面的 先类型转换")
	fmt.Println(tab, tab, "func (x IntSlice) Sort() // IntSlice 类型的方法, 普通的 int 切片使用此方法需要先将类型转为 IntSlice 类型, 见下面的 先类型转换")
	fmt.Println(tab, "---------浮点数排序-------")
	fmt.Println(tab, "func Float64s(x []float64) // 升序排序")
	fmt.Println(tab, "func Float64sAreSorted(x []float64) bool")
	fmt.Println(tab, "func SearchFloat64s(a []float64, x float64) int // 查询给定值在切片中出现的下标, 如果不存在则返回插入位置的下标")
	fmt.Println(tab, "type Float64Slice []float64 // 声明本地类型 Float64Slice")
	fmt.Println(tab, tab, "func (p Float64Slice) Search(x float64) int // Float64Slice 类型的方法, 普通的 float64 切片使用此方法需要先将类型转为 Float64Slice 类型, 见下面的 先类型转换")
	fmt.Println(tab, tab, "func (x Float64Slice) Sort() // Float64Slice 类型的方法, 普通的 float64 切片使用此方法需要先将类型转为 Float64Slice 类型, 见下面的 先类型转换")
	fmt.Println(tab, "---------字符串排序-------")
	fmt.Println(tab, "func Strings(x []string) // 升序排序")
	fmt.Println(tab, "func StringsAreSorted(x []string) bool")
	fmt.Println(tab, "func SearchStrings(a []string, x string) int // 查询给定值在切片中出现的下标, 如果不存在则返回插入位置的下标")
	fmt.Println(tab, "type StringSlice []string // 声明本地类型 StringSlice")
	fmt.Println(tab, tab, "func (p StringSlice) Search(x string) int // StringSlice 类型的方法, 普通的 string 切片使用此方法需要先将类型转为 StringSlice 类型, 见下面的 先类型转换")
	fmt.Println(tab, tab, "func (x StringSlice) Sort() // StringSlice 类型的方法, 普通的 string 切片使用此方法需要先将类型转为 StringSlice 类型, 见下面的 先类型转换")
	fmt.Println(tab, "---------其他类型切片排序-------")
	fmt.Println(tab, "func Slice(x interface{}, less func(i, j int) bool) // 根据给定的函数排序 x, 如果 x 不是 slice 则报错")
	fmt.Println(tab, "func SliceIsSorted(x interface{}, less func(i, j int) bool) bool")
	fmt.Println(tab, "func SliceStable(x interface{}, less func(i, j int) bool)")
	fmt.Println(tab, "-------------")
	fmt.Println(tab, "func Reverse(data Interface) Interface // 返回一个实现了 Less 方法的结构体, 参数类型同顶部声明")
	fmt.Println(tab, `type reverse struct {
        // This embedded Interface permits Reverse to use the methods of
        // another Interface implementation.
        Interface
    }

    // Less returns the opposite of the embedded implementation's Less method.
    func (r reverse) Less(i, j int) bool {
        return r.Interface.Less(j, i)
    }

    // Reverse returns the reverse order for data.
    func Reverse(data Interface) Interface {
        return &reverse{data}
    }
	`)
	fmt.Println("---------------------------------")

	fmt.Println("下面3个方法使用时需要将参数类型转换为实现了 Interface 接口的类型(IntSlice, Float64Slice, StringSlice)")
	fmt.Println("sort.IsSorted(data Interface), sort.Sort(data Interface), sort.Stable(data Interface)，sort.Reverse(data Interface) Interface")
	fmt.Println("-------------")

	var is = []int{19, 1, 32, 28, 6, 3, 16, 34, 0, 13, 11, 25, 44}
	fmt.Printf("is 的类型为 %T 值为 %+v\n", is, is)                                                                                              // []int [19 1 32 28 6 3 16 34 0 13 11 25 44]
	fmt.Println("是否排序 sort.IntsAreSorted(is) 或 sort.IsSorted(sort.IntSlice(is)", sort.IntsAreSorted(is), sort.IsSorted(sort.IntSlice(is))) // false false
	// sort.Ints(is)
	sort.IntSlice(is).Sort()
	fmt.Println("sort.Ints(is) 或 sort.IntSlice(is).Sort() 排序后", "is 结果为 ", is)                                                              // [0 1 3 6 11 13 16 19 25 28 32 34 44]
	fmt.Println("是否排序 sort.IntsAreSorted(is) 或 sort.IsSorted(sort.IntSlice(is))", sort.IntsAreSorted(is), sort.IsSorted(sort.IntSlice(is))) // true true
	fmt.Println("切片必须按升序排序, 查询给定值在切片中出现的下标, 如果不存在则返回插入位置的下标")
	fmt.Println(tab, "sort.SearchInts(is, 2)", sort.SearchInts(is, 2))                    // 2
	fmt.Println(tab, "sort.IntSlice(is).Search(5) // 先类型转换", sort.IntSlice(is).Search(5)) // 3
	sort.Sort(sort.Reverse(sort.IntSlice(is)))
	fmt.Println("逆序 sort.Sort(sort.Reverse(sort.IntSlice(is)))", is) // [44 34 32 28 25 19 16 13 11 6 3 1 0]
	fmt.Println("-------------")

	var f64s = []float64{9.87, 3.1415, 13.33, 6.59, 2.84, 11.11, 4.87}
	fmt.Printf("f64s 的类型为 %T 值为 %+v\n", f64s, f64s)                                                                                                                 // []float64 [9.87 3.1415 13.33 6.59 2.84 11.11 4.87]
	fmt.Println("是否排序 sort.Float64sAreSorted(f64s) 或 sort.IsSorted(sort.Float64Slice(f64s))", sort.Float64sAreSorted(f64s), sort.IsSorted(sort.Float64Slice(f64s))) // false false
	// sort.Float64s(f64s)
	sort.Float64Slice(f64s).Sort()
	fmt.Println("sort.Float64s(f64s) 或 sort.Float64Slice(f64s).Sort() 排序后", "f64s 结果为 ", f64s)                                                                      //  [2.84 3.1415 4.87 6.59 9.87 11.11 13.33]
	fmt.Println("是否排序 sort.Float64sAreSorted(f64s) 或 sort.IsSorted(sort.Float64Slice(f64s))", sort.Float64sAreSorted(f64s), sort.IsSorted(sort.Float64Slice(f64s))) // true true
	fmt.Println("切片必须按升序排序, 查询给定值在切片中出现的下标, 如果不存在则返回插入位置的下标")
	fmt.Println(tab, "sort.SearchFloat64s(f64s, 3.1415) ", sort.SearchFloat64s(f64s, 3.1415))                     // 1
	fmt.Println(tab, "sort.Float64Slice(f64s).Search(55.1415) // 先类型转换", sort.Float64Slice(f64s).Search(55.1415)) // 7
	sort.Sort(sort.Reverse(sort.Float64Slice(f64s)))
	fmt.Println("逆序 sort.Sort(sort.Reverse(sort.Float64Slice(f64s)))", f64s) // [13.33 11.11 9.87 6.59 4.87 3.1415 2.84]
	fmt.Println("-------------")

	var ss = []string{"hello", "world", "gg", "Jack", "Tom", "jerry", "JACK", "Alen", "Bob", "TOM", "bob"}
	fmt.Printf("ss 的类型为 %T 值为 %+q\n", ss, ss)                                                                                                           // []string ["hello" "world" "gg" "Jack" "Tom" "jerry" "JACK" "Alen" "Bob" "TOM" "bob"]
	fmt.Println("是否排序 sort.StringsAreSorted(ss) 或 sort.IsSorted(sort.StringSlice(ss))", sort.StringsAreSorted(ss), sort.IsSorted(sort.StringSlice(ss))) // false false
	// sort.Strings(ss)
	sort.StringSlice(ss).Sort()
	fmt.Println("sort.Strings(ss) 或 sort.StringSlice(ss).Sort() 排序后", "ss 结果为 ", ss)                                                                    // [Alen Bob JACK Jack TOM Tom bob gg hello jerry world]
	fmt.Println("是否排序 sort.StringsAreSorted(ss) 或 sort.IsSorted(sort.StringSlice(ss))", sort.StringsAreSorted(ss), sort.IsSorted(sort.StringSlice(ss))) // true true
	fmt.Println("切片必须按升序排序, 查询给定值在切片中出现的下标, 如果不存在则返回插入位置的下标")
	fmt.Println(tab, "sort.SearchStrings(ss, \"GG\")", sort.SearchStrings(ss, "GG"))                      // 2
	fmt.Println(tab, "sort.StringSlice(ss).Search(\"YMD\") // 先类型转换", sort.StringSlice(ss).Search("YMD")) // 6
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Println("逆序 sort.Sort(sort.Reverse(sort.StringSlice(ss)))", ss) // [world jerry hello gg bob Tom TOM Jack JACK Bob Alen]
	fmt.Println("---------------------------------")

	sort.Stable(sort.IntSlice(is))
	fmt.Println("sort.Stable(sort.IntSlice(is)) 排序后", "is 结果为", is)
	sort.Stable(sort.Float64Slice(f64s))
	fmt.Println("sort.Stable(sort.Float64Slice(f64s)) 排序后", "f64s 结果为", f64s)
	sort.Stable(sort.StringSlice(ss))
	fmt.Println("sort.Stable(sort.StringSlice(ss)) 排序后", "ss 结果为", ss)
	fmt.Println("---------------------------------")

	stu := []struct {
		Name string
		Age  uint8
	}{
		{Name: "Gopher", Age: 18},
		{Name: "Alice", Age: 25},
		{Name: "Vera", Age: 7},
		{Name: "Gopher", Age: 6},
		{Name: "Bob", Age: 75},
	}
	fmt.Println(`
stu := []struct {
    Name string
    Age  uint8
}{
    {Name: "Gopher", Age: 18},
    {Name: "Alice", Age: 25},
    {Name: "Vera", Age: 7},
    {Name: "Gopher", Age: 6},
    {Name: "Bob", Age: 75},
}
	`)
	fmt.Printf("stu 的类型为 %T 值为 %+v\n", stu, stu)                                                                                                                                             // []struct { Name string; Age uint8 }  [{Name:Gopher Age:18} {Name:Alice Age:25} {Name:Vera Age:7} {Name:Gopher Age:6} {Name:Bob Age:75}]
	fmt.Println("是否排序 sort.SliceIsSorted(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name })", sort.SliceIsSorted(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name })) // false
	sort.Slice(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name })
	fmt.Println("sort.Slice(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name }) 排序后", "stu 结果为 ", stu) // [{Alice 25} {Bob 75} {Gopher 18} {Gopher 6} {Vera 7}]
	fmt.Println("如果希望相等的元素保持其原始顺序, 需使用 sort.SliceStable(x interface{}, func(i, j int) bool)")
	sort.SliceStable(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name })
	fmt.Println("sort.SliceStable(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name }) 排序后", "stu 结果为 ", stu)                                                                      // [{Alice 25} {Bob 75} {Gopher 18} {Gopher 6} {Vera 7}]
	fmt.Println("是否排序 sort.SliceIsSorted(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name })", sort.SliceIsSorted(stu, func(i, j int) bool { return stu[i].Name < stu[j].Name })) // true
	fmt.Println("---------------------------------")

	fmt.Println(`
a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
var x = 15
var i = sort.Search(len(a), func(i int) bool { return a[i] >= x })
fmt.Println("采用二分法搜索找到[0, n)区间内最小的满足 f(i) == true 的值 i, 如果未找到则返回该值插入的索引")
fmt.Println("x = 15, i = ", i)
x = 7
i = sort.Search(len(a), func(i int) bool { return a[i] >= x })
fmt.Println("x = 7, i = ", i)
	`)
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	var x = 15
	var i = sort.Search(len(a), func(i int) bool { return a[i] >= x })
	fmt.Println("采用二分法搜索找到[0, n)区间内最小的满足 f(i) == true 的值 i, 如果未找到则返回该值插入的索引")
	fmt.Println("x = 15, i = ", i)
	x = 7
	i = sort.Search(len(a), func(i int) bool { return a[i] >= x })
	fmt.Println("x = 7, i = ", i)
	fmt.Println("---------------------------------")

	CustomSort()
}
