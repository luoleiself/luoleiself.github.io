package main

import (
	"fmt"
)

func readme() {
	fmt.Println("在被测试包相同的包下面, 创建包含 TestXxx 函数的以 _test.go 为文件名结尾的文件")
	fmt.Println(tab, "测试用例文件不会参与正常源码的编译, 不会被包含到可执行文件中")
	fmt.Println(tab, "测试用例的文件名必须以 _test.go 结尾")
	fmt.Println(tab, "测试函数的名称要以 TestXxx 或 BenchmarkXxx 或 FuzzXxx 开头, 后面可以跟任意字母组成的字符串, 采用驼峰写法,")
	fmt.Println(tab, " 一个测试用例文件中可以包含多个测试函数")
	fmt.Println("基准测试: 用于测量和评估软件性能指标的方法")
	fmt.Println("包内测试和包外测试")
	fmt.Println(tab, "包内测试: 测试代码放在与被测试包同名的包中")
	fmt.Println(tab, " 包内测试本质上是一种白盒测试方法, 测试代码可以访问该包下的所有符号, 可以很容易地达到较高的测试覆盖率")
	fmt.Println(tab, "缺点")
	fmt.Println(tab, tab, "测试代码自身需要经常性的维护, 测试代码的测试数据构造和测试逻辑和被测试包的特定数据结构是紧耦合的")
	fmt.Println(tab, tab, "包循环引用, Go 编译器不允许")
	fmt.Println(tab, "包外测试: 仅针对导出 API 的测试, 包外测试本质上是一种黑盒测试方法, 例如 strings 包的测试")
	fmt.Println(tab, " 包名遵循 packageName_test 规则")
	fmt.Println(tab, "缺点")
	fmt.Println(tab, tab, "仅能访问被测试的包导出的函数、方法、类型和变量")
	fmt.Println(tab, tab, "容易出现对被测试包的测试覆盖不足的情况")
	fmt.Println(tab, "安插后门: 在被测试包同级目录下, 创建 export_test.go 文件(导出包内部访问的数据), 同时创建 被测试包名_test.go 文件, 在此文件中导入被测试的包")
	fmt.Println(tab, " 该文件既不会被包含在正式产品代码中, 又不包含任何测试代码, 而仅用于将被测试包的内部符号在测试阶段暴露给包外测试代码")
	fmt.Println(tab, " 例如 fmt 包下的 export_test.go")
	fmt.Println(`
  // fmt/print.go
  package fmt
  import (
    "internal/fmtsort"
    "io"
    "os"
    "reflect"
    "strconv"
    "sync"
    "unicode/utf8"
  )

  // fmt/export_test.go
  package fmt
  var IsSpace = isSpace
  var Parsenum = parsenum

  // fmt/fmt_test.go
  package fmt_test

  import (
    . "fmt"
    "internal/race"
    "io"
    "math"
    "reflect"
    "runtime"
    "strings"
    "testing"
    "time"
    "unicode"
  )`)
	fmt.Println("----------------------------------")

	fmt.Println("测试固件")
	fmt.Println(tab, "指一个人造的、确定性的环境, 一个测试用例或一个测试套件(下的一组测试用例)在这个环境中进行测试, 其测试结果是可重复的(多次测试运行的结果是相同的),")
	fmt.Println(tab, " 通常使用 setUp 和 tearDown 来代表测试固件的 创建/设置 和 拆除/销毁 的动作.")
	fmt.Println(tab, "常见场景")
	fmt.Println(tab, tab, "将一组已知的特定数据加载到数据库中, 测试结束后清除这些数据")
	fmt.Println(tab, tab, "复制一组特定的已知文件, 测试测试结束后清除这些文件")
	fmt.Println(tab, tab, "创建伪对象(fake object)或模拟对象(mock object), 并为这些对象设定测试时所需的特定数据和期望结果")
	fmt.Println("静态测试固件")
	fmt.Println(" Go 工具链将会忽略名为 testdata 的目录, 可以在 testdata 目录下存放和管理测试代码依赖的数据文件,")
	fmt.Println("  而 go test 命令在执行时会将被测试包源码所在目录设置为其工作目录, 这样如果要使用 testdata 目录下的某个数据文件,")
	fmt.Println("  就不需要再处理各种路径问题, 直接在测试代码中定位到充当测试固件的数据文件")
	fmt.Println("----------------------------------")

	fmt.Println("go 1.14 之前测试固件的 setUp 和 tearDown 的实现")
	fmt.Println("go 1.14 testing 包新增了 Cleanup 方法支持测试固件的销毁操作")
	fmt.Println("   t.Cleanup(f func()) 注册测试(或子测试)及其所有子测试完成时要调用的函数, 清理函数将按 LIFO 的顺序调用")
	fmt.Println(`
  func setUp() func() {
    ...
    return func(){
      ...
    }
  }

  func TestXxx(t *testing.T){
    defer setUp()() // go 1.14 之前
    t.Cleanup(setUp()) // go 1.14 之后
    ...
  }`)
	fmt.Println("go 1.4 添加 M 通过 TestMain 函数去执行测试")
	fmt.Println(`
  func pkgSetUp() func() {
    ...
    return func(){
      ...
    }
  }

  func TestMain(m *testing.M) {
    defer pkgSetUp("TestMain")()
    m.Run()
  }`)
	fmt.Println("----------------------------------")

	fmt.Println("Fail: 失败, 继续当前函数")
	fmt.Println("FailNow: 失败, 终止当前函数")
	fmt.Println("SkipNow: 跳过, 终止当前函数")
	fmt.Println("Log: 输出信息, 仅失败或 -v 时有效")
	fmt.Println("Logf: Log + format")
	fmt.Println("Error: Fail + Log")
	fmt.Println("Errorf: Fail + Logf")
	fmt.Println("Fatal: FailNow + Log")
	fmt.Println("Fatalf: FailNow + Logf")
	fmt.Println("Skip: SkipNow + Log")
	fmt.Println("Skipf: SkipNow + Logf")
	fmt.Println("----------------------------------")
}
