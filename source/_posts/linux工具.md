---
title: linux工具
date: 2022-05-31 14:40:43
categories:
  - [linux, shell]
tags:
  - linux
  - shell
---

### 删除无用内核

```shell
# 查看系统内核版本
[root@localhost ~]# uname -r
# 移除系统无用内容	
[root@localhost ~]# yum remove $(rpm -qa | grep kernel | grep -v $(uname -r))
# 查看系统已安装的包
[root@localhost ~]# rpm -qa | grep kernel
```

### awk

awk 是一种可以对文本和数据进行处理的编程语言, 默认情况下, awk 文件的每一行都被视为一条记录, 然后 awk 记录进一步分解成一系列的字段

```shell
[root@localhost ~]# awk '{ sum += $1 }; END { print sum }' file
[root@localhost ~]# awk -F : '{ print $1 }' /etc/passwd
```

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

  - BEGIN { 这里面是行处理语句执行前的语句 }
  - END { 这里面是行处理语句执行完成后的语句 }
  - { 这里面是处理每一行时要执行的语句 }

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

```shell
[root@localhost ~]# git branch -r | awk 'BEGIN {print "hello awk, I am coming\n"}END{print "hello awk, good bye\n"}{printf "NF--%s NR--%s FNR--%s $0--%s\n",NF,NR,FNR,$0;}'
```

### xargs

将参数列表转换成小块分段传递给其他命令, 以避免参数列表过长的问题, 可单独使用, 也可以使用管道符、重定位符等其他命令配合使用

```shell
xargs [OPTION]... COMMAND INITIAL-ARGS...
```

- -E EOFString 指定逻辑 EOF 字符串以替换缺省的下划线(\_), xargs 命令读取标准输入直到达到 EOF 或指定的字符串
- -I replaceString 插入标准输入的每一行作为 command 参数的自变量, 把它插入每个发生的 replaceString 的 Argument 中
- -n number 指定传递给执行命令的参数个数, 默认是所有
- -p 每执行一个 argument 时询问一次用户
- -a file 从文件中读入作为 stdin
- -s size 命令行的最大字符数，指的是 xargs 后面那个命令的最大命令行字符数
- -t 执行命令之前先打印执行命令

```shell
# -I 指定参数自变量
[root@localhost ~]# echo "file1 file2 file3" | xargs -t -I % sh -c 'touch %;ls -l %'
[root@localhost ~]# sh -c touch file1 file2 file3;ls -l file1 file2 file3
```

- -d 设置自定义分隔符

```shell
[root@localhost ~]# echo -n file1#file2#file3#file4 | xargs -d \# -t touch
[root@localhost ~]# touch file1 file2 file3 file4
```

### grep

> BRE 定义了 4 组元字符 `[ ]` `.` `^` `$`
> ERE 增加了 3 组元字符 `{ }` `()` `|`

```shell
grep [OPTION]... PATTERN [FILE]...
```

- -V,\-\-version 显示版本信息
- -c,\-\-count 统计符合字符串条件的行数
- -i,\-\-ignore-case 忽略大小写
- -e,\-\-regexp=PATTERN 使用正则表达式匹配

  ```shell
  # 使用 -e 参数 需要将正则中的部分字符转义才能使用
  [root@localhost ~]# grep -i -e "foo\|bar" 
  ```

- -E,\-\-extended-regexp 使用扩展正则表达式匹配(ERE)

  ```shell
  [root@localhost ~]# grep -i -E "foo|bar" # 此处不需要进行字符转义
  ```

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

```shell
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

```shell
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

```shell
scp [options] [[user@]host1:]file1 ... [[user@]host2:]file2
```

- -C 传输过程中允许压缩文件
- -p 保留源文件的修改时间, 访问时间, 访问权限
- -q 不显示传输进度条
- -r 递归复制整个目录
- -v 详细方式显示输出
- -P 指定数据传输的端口号
- -l 限制传输带宽 KB/s

#### 本地复制到远程

```shell
# 拷贝文件, 可以使用原文件名也可以重新命名文件
[root@localhost ~]# scp -Cp /home/workspace/file1.txt root@192.168.1.3:/home/workspace/

# 拷贝目录
[root@localhost ~]# scp -rCp /home/workspace/ root@192.168.1.3:/home/workspace/
```

#### 远程复制到本地

```shell
# 拷贝文件, 可以使用原文件名也可以重新命名文件
[root@localhost ~]# scp -Cp root@192.168.1.3:/home/workspace/file1.txt /home/workspace/
# 拷贝目录
[root@localhost ~]# scp -rCp root@192.168.1.3:/home/workspace/ /home/workspace
```

### firewall-cmd 防火墙

- \-\-permanent # 永久修改
- \-\-reload # 重新加载防火墙配置

```shell
[root@centos7 ~]firewall-cmd --list-all # 显示所有信息
[root@centos7 ~]firewall-cmd --list-ports # 显示端口信息
[root@centos7 ~]firewall-cmd --remove-ports=<port>/<protocol> # 显示端口信息

[root@centos7 ~]firewall-cmd --add-port=<port>/<protocol> --permanent # 永久修改防火墙配置
```

### tar 归档

```shell
tar [OPTION...] [FILE]...
```

- \-c,\-\-create 创建新的 tar 文件
- \-x,\-\-extract 解开 tar 文件
- \-r,\-\-append 添加文件到已经压缩的文件
- \-t,\-\-list 列出压缩文件中的信息
- \-u,\-\-update 用已打包的文件的较新版本更新 tar 文件
- \-d,\-\-diff,\-\-compare 将文件系统里的文件和 tar 文件里的文件进行比较
- \-\-delete 删除 tar 文件里面的文件, 不能用于已保存在磁带上的 tar 文件
- \-v,\-\-verbose 列出每一步处理涉及的文件的信息
- \-k,\-\-keep-old-files 不覆盖文件系统上已有的文件
- \-f,\-\-file 指定要处理的文件名
- \-w 每一步都要求确认

- \-\-atime-preserve 不改变转储文件的存取时间
- \-m,\-\-modification-time 当从一个档案中恢复文件时, 不使用新的时间标签
- \-C,\-\-directory DIR 转到指定的目录

#### 压缩工具

- \-j,\-\-bzip2 调用 bzip2 执行压缩或解压缩
- \-J,\-\-xz,\-\-lzma 调用 XZ Utils 执行压缩或解压缩
- \-z,\-\-gzip,\-\-gunzip,\-\-ungzip 调用 gzip 执行压缩或解压缩
- \-Z,\-\-compress,\-\-uncompress 调用 compress 执行压缩或解压缩

#### 归档压缩

```shell
# 使用 gzip 压缩归档 workspace 目录
[root@centos7 ~]tar -czvf workspace.tar.gz ./workspace
```

#### 解压缩

```shell
# 使用 gzip 解压缩到当前目录下的 redis-stable 目录
[root@centos7 ~]tar -xzvf redis-stable.tar.gz redis-stable
```

### crontab 定时任务

- -u 指定用户
- -e 使用编辑器设置时程表
- -r 删除时程表
- -l 列出目前的时程表

#### 每分钟向指定文件追加写入一条数据

```shell
#!/bin/bash

# /root/workspace/crontab-out-format.sh

file=/root/workspace/crontab-out-format.txt

function dateFormat() {
  echo $(/bin/date +"hello crontab %Y-%m-%d %H:%M:%S")
}

if [ -e $file ]; then
  echo $(dateFormat) >> $file
else
  touch $file
  echo $(dateFormat) >> $file
fi

[root@centos7 workspace]# crontab -l  # 列出当前用户的所有任务
* * * * * /bin/bash /root/workspace/crontab-out-format.sh # 定时任务
```
