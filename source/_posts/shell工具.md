---
title: shell工具
date: 2022-05-31 14:40:43
categories:
  - [linux, shell]
tags:
  - linux
  - shell
---

### awk

#### 内建变量

- OFMT 输出数字格式(默认%.6g)
- CONVFMT 数字转换格式(默认值为%.6g)ENVIRON 环境变量关联数组
- ENVIRON 数组, 存储当前系统的环境变量, ENVIRON['PATH']
- ERRNO 最后一个系统错误的描述
- FILENAME 当前正在处理的文件名
- FIELDWIDTHS 以空格分隔的字段宽度
- PROCINFO 数组, 存储当前系统运行时信息, PROCINFO['uid']
- IGNORECASE 如果为真, 则进行忽略大小写的匹配
- 字段分隔符
  - FS 字段的分隔符, 默认为空格, 可以使用 -F 参数设置
  - OFS 输出的字段分隔符, 默认为空格
- 记录分隔符

  - RS 输入记录的分隔符, 默认为换行符
  - ORS 输出的记录分隔符, 默认为换行符

<!-- more -->

- 记录统计
  - $0 完整的输入记录
  - $1-$n 表示当前行第几个字段, $(NF - n) 动态计算第几个字段
  - NF 当前记录中的字段个数
  - NR 已经读出的记录数(默认从 1 开始)
  - FNR 当前处理的记录号
- 参数
  - ARGC 命令行参数个数(不包括 awk 的选项和 awk 的程序内容)
  - ARGV 命令行参数序列数组,下标从 0 开始
  - ARGIND 命令行中当前文件的位置(从 0 开始算)
- 流程控制

  - BEGIN { 这里面放的是执行前的语句 }
  - END { 这里面放的是处理完所有的行后要执行的语句 }
  - { 这里面放的是处理每一行时要执行的语句 }

- 内置函数
  - gsub(r, s) 在整个 $0 中用 s 替换 r
  - gsub(r, s, t) 在整个 t 中用 s 替换 r
  - index(s, t) 返回 s 中字符串 t 的第一位置
  - length(s) 返回 s 的长度
  - match(s, r) 测试 s 是否包含匹配 r 的字符串
  - split(s, a, fs) 在 fs 上将 s 分成序列 a
  - sprint(fmt, exp) 返回经 fmt 格式化后的 exp
  - sub(r, s) 用 $0 中最左边最长的子串代替 s
  - substr(r, p) 返回字符串 s 中从 p 开始的后缀部分
  - substr(s, p, n) 返回字符串 s 中从 p 开始长度为 n 的后缀部分

### xargs

- -t 执行命令之前先打印执行命令
- -n 指定传递给执行命令的参数个数, 默认是所有
- -I 同时运行多个命令, 并替换所有匹配的项为传递给 xargs 的参数
- -p 每执行一个 argument 时询问一次用户
- -a file 从文件中读入作为 stdin
- -s num 命令行的最大字符数，指的是 xargs 后面那个命令的最大命令行字符数

```bash
echo "file1 file2 file3"| xargs -t -I % sh -c 'touch %;ls -l %'
sh -c touch file1 file2 file3;ls -l file1 file2 file3
```

- -d 设置自定义分隔符

```bash
echo -n file1#file2#file3#file4|xargs -d \# -t touch
touch file1 file2 file3 file4
```

### grep

- -V,\-\-version 显示版本信息
- -c,\-\-count 统计符合字符串条件的行数
- -i,\-\-ignore-case 忽略大小写
- -e,\-\-regexp=PATTERN 使用正则表达式匹配
- -E,\-\-extended-regexp 使用扩展正则表达式匹配(ERE)
- -G,\-\-basic-regexp 使用基础正则表达式(BRE)
- -A{NUM},\-\-after-context=NUM 查找某些字符的内容, 并向下延伸 `NUM` 行
- -B{NUM},\-\-before-context=NUM 查找某些字符的内容, 并向上延伸 `NUM` 行
- -C{NUM},\-\-context=NUM 查找某些字符的内容, 并向上和向下各延伸 `NUM` 行
- -f File,\-\-file=File 从文件中提取模板
- -h,\-\-no-filename 当搜索多个文件时, 不显示匹配文件名前缀
- -o,\-\-only-matching 只显示正则表达式匹配的部分
- -q,\-\-quiet 取消显示,只返回退出状态, `0` 表示找到了匹配的行
- -l,\-\-files-with-matches 打印匹配模板的文件清单
- -L,\-\-files-without-match 打印不匹配模板的文件清单
- -n,\-\-line-number 在匹配的行前面打印行号
- -v,\-\-invert-match 显示不包括文本的所有信息
- -R,-r,\-\-recursive 递归的读取目录下的所有文件,包括子目录

### 批量删除本地关联的 git 远程分支

- awk 和 xargs 命令结合使用

- 不带远程名称的过滤

```bash
# 格式化输出所有本地关联的分支名
$ git branch -r | \
  grep -i 'origin/feature*' | \
  awk '{FS="/"; \
    printf "FS %4s OFS %4s NF %4s NR %4s $0 %4s $1 %4s\n",FS,OFS,NF,NR,$0,$1; \
  }'
# 或
$ git branch -r | \
  awk 'BEGIN{IGNORECASE=1} \
    /origin\/feature/{ \
      FS="/"; \
      printf "FS %4s OFS %4s NF %4s NR %4s $0 %4s $1 %4s\n",FS,OFS,NF,NR,$0,$1; \
    }'
FS    / OFS      NF    1 NR    1 $0   origin/feature-BUSINESS-11269 $1 origin/feature-BUSINESS-11269
FS    / OFS      NF    2 NR    2 $0   origin/feature-BUSINESS-11374 $1   origin
FS    / OFS      NF    2 NR    3 $0   origin/feature-mall-20220406 $1   origin
FS    / OFS      NF    2 NR    4 $0   origin/feature-mall-entry $1   origin
FS    / OFS      NF    2 NR    5 $0   origin/feature-privacy-20220329 $1   origin
FS    / OFS      NF    2 NR    6 $0   origin/feature-service_tpl-20220418 $1   origin
FS    / OFS      NF    2 NR    7 $0   origin/feature_BUSINESS-0707 $1   origin
...

# 执行删除分支操作
# BEGIN{printf "\nawk begin filtering\n\n";}END{printf "\ndelete successfully\n";}
$ git branch -r | \
  grep -i 'origin/revert*' | \
  awk '{printf "%4s\n",$0;}' | \  # awk '{printf $0 "\n";}'
  xargs -t -I {} git branch -dr {}
Deleted remote-tracking branch origin/revert-0946083d (was 795de8b941).
Deleted remote-tracking branch origin/revert-4835d8ea (was eac88d3e28).
## 或
$ git branch -r | \
  awk 'BEGIN{IGNORECASE=1}/origin\/revert/{printf "%4s\n",$0}' | \
  xargs -t -I {} git branch -dr {}
```

- 包含主机名的过滤

```bash
# IGNORECASE=1 开启忽略大小写
# $0~/origin\/feature/ 判断分支名是否包含 origin/feature
# 使用内置函数 gsub 全局替换 remotes 为空
$ git branch -a | \
  awk 'BEGIN{IGNORECASE=1}{ \
    if($0~/origin\/feature/){ \
      gsub(/remotes\//, " ", $0); \
      printf "%4s\n",$0; \
    } \
  }' | \
  xargs -t -I {} git branch -dr {}
```

### scp 主机之间复制文件

- -C 传输过程中允许压缩文件
- -p 保留源文件的修改时间, 访问时间, 访问权限
- -q 不显示传输进度条
- -r 递归复制整个目录
- -v 详细方式显示输出
- -P 指定数据传输的端口号
- -l 限制传输带宽 KB/s

#### 本地复制到远程

```bash
scp local_file remote_username@remote_ip:remote_folder
# 拷贝文件, 可以使用原文件名也可以重新命名文件
scp -Cp /home/workspace/file1.txt root@192.168.1.3:/home/workspace/

# 拷贝目录
scp -rCp /home/workspace/ root@192.168.1.3:/home/workspace/
```

#### 远程复制到本地

```bash
# 拷贝文件, 可以使用原文件名也可以重新命名文件
scp -Cp root@192.168.1.3:/home/workspace/file1.txt /home/workspace/
# 拷贝目录
scp -rCp root@192.168.1.3:/home/workspace/ /home/workspace
```
