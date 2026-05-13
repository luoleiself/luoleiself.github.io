package externaltesting

import (
	"fmt"
	"time"
)

type Month uint

const (
	january Month = iota + 1
	february
	march
	april
	may
	june
	july
	august
	september
	october
	november
	december
)

func init() {
	fmt.Println("此处位于 externaltesting 包的 init 方法中, 仅包内可访问的变量")
	fmt.Printf(" january %d february %d march %d april %d may %d june %d july %d august %d september %d october %d november %d december %d \n", january, february, march, april, may, june, july, august, september, october, november, december)
	fmt.Println("----------------")
	fmt.Println("externaltesting_test 包是对 externaltesting 包的包外测试")
	fmt.Println(" state.go 文件属于 externaltesting 包, 不能包含 main 方法, 使用 init 方法作为程序初始化")
	fmt.Println(" export_test.go 固定命名, 属于 externaltesting 包, 仅用于将被测试包的内部访问符号暴露给包外测试代码使用")
	fmt.Println(" externaltesting_test.go 文件属于 externaltesting_test 包, 对 externaltesting 包的包外测试代码")
	fmt.Println("----------------------------------")
	time.Sleep(5 * time.Second)
}
