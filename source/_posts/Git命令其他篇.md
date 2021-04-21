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
  git submodule init [<path>]
```

##### submodule deinit 删除

```bash
  git submodule deinit [-f]
  git submodule deinit [--all]
```

##### submodule set-branch 设置分支

设置嵌套仓库的默认远程关联分支

```bash
  git submodule set-branch (-b) <branch> [<path>]
```

##### submodule set-url 设置地址

设置嵌套仓库新的地址,会自动同步新的地址配置项

```bash
  git submodule set-url <path> <newurl>
```

##### submodule add 添加

在当前指定目录下添加仓库

```bash
  git submodule add <repository> [<path>]
```

##### status|summary 查看

- 递归查看嵌套仓库的状态

  ```bash
    git submodule status [--recursive] [<path>]
  ```

- 查看嵌套仓库的提交记录

  ```bash
    git submodule summary [<path>]
  ```

##### submodule update 更新

递归更新嵌套仓库信息

```bash
  git submodule update [--recursive] [<path>]
```

#### archive 归档

- list: 显示支持的归档文件格式

  ```bash
    git archive [-l]
  ```

- format: 指定归档文件格式

  ```bash
    git archive --format=<fmt> [<path>...]
  ```

- prefix: 指定归档文件前缀

  ```bash
    git archive --prefix=<prefix> [<path>...]
  ```

- output|o: 将归档文件写入到指定文件非输出流

  ```bash
    git archive [-o <file>] [<path>...]
  ```

- add-file: 添加未被追踪的文件到归档文件中

  ```bash
    git archive --add-file <file> [<path>...]
  ```

#### clean 清除

- n|dry-run: 不移除任何东西,只显示会做什么
- f|force: 强制删除
- i|interactive: 交互式操作
- e: 使用正则表达式匹配

##### 从工作区移除未被追踪的文件

```bash
  git clean [-n] [-f] [-i] [-e <pattern>] <path>
```

---
