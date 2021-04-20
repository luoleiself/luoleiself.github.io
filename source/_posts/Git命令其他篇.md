---
title: Git命令其他篇
date: 2021-04-20 16:57:42
categories:
  - tools
tags:
  - git
---

### 其他篇

#### submodule 仓库

##### submodule init 初始化

```bash
  git submodule init [<path>] # 初始化指定目录为嵌套仓库
```

##### submodule deinit 删除

```bash
  git submodule deinit [-f] # 强制删除嵌套仓库
  git submodule deinit [--all] # 删除所有嵌套仓库
```

##### submodule set-branch 设置分支

```bash
  git submodule set-branch (-b) <branch> [<path>] # 设置嵌套仓库的默认远程关联分支
```

##### submodule set-url 设置地址

```bash
  git submodule set-url <path> <newurl> # 设置嵌套仓库新的地址,会自动同步新的地址配置项
```

##### submodule add 添加

```bash
  git submodule add <repository> [<path>] # 在当前指定目录下添加仓库
```

##### status|summary 查看

```bash
  git submodule status [--recursive] [<path>] # 递归查看嵌套仓库的状态
  git submodule summary [<path>] # 查看嵌套仓库的提交记录
```

##### submodule update 更新

```bash
  git submodule update [--recursive] [<path>] # 递归更新嵌套仓库信息
```

#### archive 归档

- list

  ```bash
    git archive [-l] # 显示支持的归档文件格式
  ```

- format

  ```bash
    git archive --format=<fmt> [<path>...] # 指定归档文件格式
  ```

- prefix

  ```bash
    git archive --prefix=<prefix> [<path>...] # 指定归档文件前缀
  ```

- output|o

  ```bash
    git archive [-o <file>] [<path>...] # 将归档文件写入到指定文件非输出流
  ```

- add-file

  ```bash
    git archive --add-file <file> [<path>...] # 添加未被追踪的文件到归档文件中
  ```

#### clean 清除

- n|dry-run 不移除任何东西,只显示会做什么
- f|force 强制删除
- i|interactive 交互式操作
- e 使用正则表达式匹配

  ```bash
    git clean [-n] [-f] [-i] [-e <pattern>] <path> # 从工作区移除未被追踪的文件
  ```

---
