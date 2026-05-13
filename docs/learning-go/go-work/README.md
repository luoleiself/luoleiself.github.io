go.work

- 初始化工作空间 `go work init`
- 克隆第三方模块仓库到 go.work 所在目录下 `git clone https://go.googlesource.com/example`
- 添加模块路径到工作空间 `go work use <dirPath>`
- 向 第三方包 stringutil 添加自定义方法 ToUpper 并调用
