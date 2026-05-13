package main

import (
	"fmt"
)

func cmdInternalPkg() {
	fmt.Println("---------------cmdInternalPkg()----------------")
	fmt.Println("内部包(internal package): 只能被它们所在的源子树的中包导入")
	fmt.Println(tab, "导入包的路径为包所在的目录到模块根目录的\033[1;32m绝对路径\033[0m")
	fmt.Println("将包放在一个名为 internal 的目录或一个名为 internal 的目录的子目录中,")
	fmt.Println("当 go 命令看到在其包的导入路径中包含 internal 时, 会验证执行导入的包是否在以父目录为根的 internal 目录中")
	fmt.Println(tab, "内部包不能被它们所在的源子树之外的包导入")
	fmt.Println(tab, "父级包只能访问内部包中可导出的内容")
	fmt.Println(`
  internal-demo
    |-- go.mod
    |-- main.go	×
    |-- hello
      |-- hello.go	×
    |-- world
      |-- world.go	√
      |-- internal
        |-- internala
          |-- internala.go	√
        |-- internalb
          |-- internalb.go	√
      |-- worldinner
        |-- worldinner.go	√`)
	fmt.Println("internal 目录下的包只能被 world包, worldinner包, internala包, internalb包等导入")
	fmt.Println("在 world 包之外的包中导入时, 编译错误")
	fmt.Println(`
$ go run .
package internaldemo
        hello.go:7:2: use of internal package internaldemo/world/internal/internala not allowed
package internaldemo
        hello.go:8:2: use of internal package internaldemo/world/internal/internalb not allowed
$ go run .
package internaldemo
       imports internaldemo/hello
       hello\hello.go:7:2: use of internal package internaldemo/world/internal/internala not allowed
				`)
	fmt.Println("-------------------------------")
}
