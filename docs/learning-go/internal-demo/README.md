# 内部包(internal package)

将包放在一个名为 internal 的目录或一个名为 internal 的目录的子目录中, 当 go 命令看到在其包的导入路径中包含 internal 时, 会验证执行导入的包是否在以父目录为根的 internal 目录中

- 内部包不能被它们所在的源子树之外的包导入
- 父级包只能访问内部包中可导出的内容
