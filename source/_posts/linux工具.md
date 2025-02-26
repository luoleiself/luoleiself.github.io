---
title: linux工具
date: 2022-05-31 14:40:43
categories:
  - [linux, shell]
tags:
  - linux
  - shell
---

### 命令行输出内容变身

格式: \033\[显示方式;前景色;背景色 m ...... \033\[0m

- \033[ 固定格式
- \033[0m 非必需, 如果省略表示后面输出内容的样式都会应用当前设置的样式

#### 属性集

| 前景色 | 背景色 | 色值   |
| :------: | :------: | ------ |
| 30     | 40     | 黑色   |
| 31     | 41     | 红色   |
| 32     | 42     | 绿色   |
| 33     | 43     | 黄色   |
| 34     | 44     | 蓝色   |
| 35     | 45     | 紫红色 |
| 36     | 46     | 青蓝色 |
| 37     | 47     | 白色   |

#### 显示方式

| 显示方式 | 表现行为 |
| :------: | -------- |
|    0     | 默认     |
|    1     | 高亮     |
|    4     | 下划线   |
|    5     | 闪烁     |
|    7     | 反色     |
|    8     | 不可见   |

```bash
# 输出字体为绿色的 你好hello world
[root@localhost ~]# printf '\033[1;32m你好hello world\033[0m\n'
# 输出字体为绿色并带有下划线的 你好hello world
[root@localhost ~]# printf '\033[4;32m你好hello world\033[0m\n'
# 输出背景色为绿色的 你好hello world
[root@localhost ~]# printf '\033[7;32m你好hello world\033[0m\n'
# 输出内容不可见
[root@localhost ~]# printf '\033[8;32m你好hello world\033[0m\n'
```

<!-- more -->
### date

```bash
$ date
Thu Oct 10 15:38:10     2024
$ date "+%F %X"
2024-10-10 15:39:33
```

- \-s, \-\-set 设置系统日期时间
- \-u, \-\-utc, \-\-universal 显示 utc 时间

```bash
$ date -u
Thu Oct 10 08:15:04 UTC 2024
```

- [+FORMAT] 格式化输出日期时间

- %a 星期的缩写
- %A 星期的完整格式
- %b 月份名称的缩写格式
- %h 月份名称的缩写格式, 同 %b
- %B 月份名称的完整格式

```bash
$ date +%a
Thu
$ date +%A
Thursday

$ date +%b
Oct
$ date +%h
Oct
$ date +%B
October
```

#### Y-m-d

- %Y 年份的完整格式
- %y 年份的最后两位数字, 00..99
- %m 月份
- %d 月份中的天
- %e 月份中的天, 使用空格补齐, 类似 %_d

```bash
$ date +%Y
2024
$ date +%y
24

$ date +%m
10

$ date +%d
10
```

- %g 年份的最后两位数字, ISO 格式
- %G 年份的完整格式, ISO 格式

- %D 日期格式, %m/%d/%y
- %F 日期格式, %+4Y-%m-%d
- %x 本地日期表示格式, 类似 %D

```bash
$ date +%g
24
$ date +%G
2024

$ date +%D
10/10/24
$ date +%F
2024-10-10
$ date +%x
10/10/24
```

- %j 当天在年中的天数
- %q 年份中的哪个季度, 1..4

```bash
$ date +%j
284
$ date +%q
4
```

#### H:M:S

- %H 小时的24时格式
- %I 小时的12时格式
- %k 小时的24时格式, 使用空格补齐, 类似 %_H
- %l 小时的12时格式, 使用空格补齐, 类似 %_I
- %M 分, 00..59
- %S 秒, 00..60

```bash
$ date +%H
15
$ date +%I
03
$ date +%k
16
$ date +%l
 4
$ date +%M
39
$ date +%S
08
```

- %r 本地时间的12时格式, 类似 03:43:18 PM
- %R 24时的小时和分钟, 类似 %H:%M
- %T 时间, 类似 %H:%M:%S
- %X 本地时间的表示格式, 类似 %T

```bash
$ date +%r
03:43:18 PM

$ date +%R
15:44

$ date +%T
15:48:22
$ date +%X
16:06:53
```

- %s 秒数, 1970年至今的秒数
- %N 纳秒, 000000000..999999999

```bash
$ date +%s
1728546311

$ date +%N
322722000
```

#### week

- %t a tab
- %u 一周中的第几天, 1为周一, 1..7
- %w 一周中的第几天, 0为周日, 0..6
- %U 一年中的第几周, 周日为第一天, 00..53
- %W 一年中的第几周, 周一为第一天, 00..53
- %V ISO 一年中的第几周, 周一为第一天, 01..53

```bash
$ date +%u
4
$ date +%w
4
$ date +%U
40
$ date +%W
41

$ date +%V
41
```

- %z 当前时区, 类似 +0800
- %:z 当前时区, 类似 +08:00
- %::z 当前时区, 类似 +08:00:00
- %Z 当前时区的缩写格式, 类似 EDT

```bash
$ date +%z
+0800
$ date +%:z
+08:00
$ date +%::z
+08:00:00
```

### ssh 操作

- ssh-keyscan 收集公钥中的主机地址
- ssh-copy-id 将本地的公钥文件复制到远程主机对应账户下的 authorized_keys 文件中
- ssh-keygen 生成非对称密钥对

```bash
[vagrant@centos8s ~]# ssh-keygen -t <密钥类型> -f <output_keyfile> -C <comment>
```

- ssh-agent ssh 认证代理, 通常和 ssh-add 配合使用管理本地密钥
  - \-s 在标准输出上初始化 bourne shell
  - \-k 关闭当前的 agent(代理), 通过环境变量 SSH_AGNENT_PID
- ssh-add 管理本地密钥
  - \-l 列出所有的密钥摘要信息
  - \-L 列出所有的公钥信息
  - \-d \<input_keyfile\> 移除指定的密钥
  - \-D 移除所有的密钥
  - \-x 锁定 agent(代理)
  - \-X 解锁 agent(代理)

```bash
# 使用 ssh -T 测试连通性
[vagrant@centos8s ~]$ ssh -T git@github.com
git@github.com: Permission denied (publickey).
# 或者使用 -i 每次都指定密钥
[vagrant@centos8s ~]$ ssh -i ~/.ssh/github_25519 -T git@github.com
Hi ......! You\'ve successfully authenticated, but GitHub does not provide shell access.

# 如果出现提示
# Could not open a connection to your authentication agent.
# 初始化一个 bash
[vagrant@centos8s ~]$ eval `ssh-agent -s`

# 使用 ssh-add 将密钥添加到 ssh 认证代理
[vagrant@centos8s ~]$ ssh-add ~/.ssh/github_25519
Identity added: /home/vagrant/.ssh/github_25519 (......@163.com)
[vagrant@centos8s ~]$ ssh-add -l
256 SHA256:CfRvLFZMgJ/p7r5ywt8BSQ2T1qEYtjCGDvVDVeYOtmY ......@163.com (ED25519)
# 不需要使用 -i 指定密钥
[vagrant@centos8s ~]$ ssh -T git@github.com
Hi ......! You\'ve successfully authenticated, but GitHub does not provide shell access.

# 移除指定的密钥
[vagrant@centos8s ~]$ ssh-add -d ~/.ssh/github_25519
Identity removed: /home/vagrant/.ssh/github_25519 (......@163.com)
# 移除所有的密钥
[vagrant@centos8s ~]$ ssh-add -D
All identities removed.
# 列出所有的密钥摘要信息
[vagrant@centos8s ~]$ ssh-add -l
The agent has no identities.
```

### awk

awk 是一种可以对文本和输入数据进行处理的编程语言, 默认情况下, awk 对文件的每一行都视为一条记录, 然后每一条记录被进一步分解成一系列的字段

- \-f scriptFile 从脚本文件中读取 awk 命令
- \-v var=value 赋值一个用户定义变量, 将外部变量传递给 awk

```bash
[root@localhost ~]# awk -v name="hello world" -F : '{print $1}' /etc/passwd
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

```bash
[root@localhost ~]# git branch -r | awk 'BEGIN {print "hello awk, I am coming\n"}END{print "hello awk, good bye\n"}{printf "NF--%s NR--%s FNR--%s $0--%s\n",NF,NR,FNR,$0;}'
```

##### 字段分隔符

- \-F 字段分隔符
- FS 字段的分隔符, 默认为空格, 可以使用 -F 参数设置
- OFS 输出的字段分隔符, 默认为空格

```bash
[root@localhost ~]# awk '{FS=":";print $1}' /etc/passwd
[root@localhost ~]# awk -F : '{print $1}' /etc/passwd
```

##### 记录分隔符

- RS 输入记录的分隔符, 默认为换行符
- ORS 输出的记录分隔符, 默认为换行符

##### 记录统计

- $0 完整的输入记录
- $1-$n 表示当前行第几个字段, $(NF - n) 动态计算第几个字段
- NF 当前记录中的字段个数
- NR 已经读出的记录数(默认从 1 开始)
- FNR 当前处理的记录号

```bash
[root@localhost ~]# echo 2 | awk -v sum=2 '{printf "FS = %s\nOFS = %s\nRS = %s\nORS = %s\nNF = %d\nNR = %d\nFNR = %d\n$0 = %s\n$1 = %s\n$2 = %s", FS, OFS, RS, ORS, NF, NR, FNR, $0, $1, $2}'
FS =
OFS =
RS =

ORS =

NF = 1
NR = 1
FNR = 1
$0 = 2
$1 = 2
$2 =
```

##### 参数

- ARGC 命令行参数个数(不包括 awk 的选项和 awk 的程序内容)
- ARGV 命令行参数序列数组,下标从 0 开始
- ARGIND 命令行中当前文件的位置(从 0 开始算)

##### 流程控制

- BEGIN { 这里面是行处理语句执行前的语句 }
- END { 这里面是行处理语句执行完成后的语句 }
- { 这里面是处理每一行时要执行的语句 }

```bash
[root@localhost ~]# awk '{ sum += $1 }END{ print sum }' file
[root@localhost ~]# echo 2 | awk -v sum=2 'BEGIN{print "hello awk"}END{print "good bye awk"}{printf "sum = %.2f\n", sum += $1}'
hello awk
sum = 4.00
good bye awk


[root@localhost ~]# echo 1.1 | awk -v sum=2 'BEGIN{print "hello awk"}END{print "good bye awk"}{printf "sum = %.2f\n", sum += $1}'
hello awk
sum = 3.10
good bye awk
```

- if 语句

```bash
[root@localhost ~]# echo -e "1 a\n2 b\n3 c" | awk '{if ($1 > 1) print $0}'
2 b
3 c
```

- while 和 do...while

```bash
[root@localhost ~]# echo 3 | awk '{i=$1; while (i <= 5) {print i; i++}}'
3
4
5
```

- for 语句

```bash
[root@localhost ~]# echo 3 | awk '{for(i=1; i <= 2;i++) {print i}}'
1
2
```

- switch 语句

```bash
[root@localhost ~]# echo 3 | awk '{switch ($1) {case 1: print "一";break; case 2: print "二";break; case 3: print "三";break; default: print "其它";}}'
三
```

##### 内置函数

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

将参数列表转换成小块分段传递给其他命令, 以避免参数列表过长的问题, 可单独使用, 也可以使用管道符、重定位符等其他命令配合使用

```bash
xargs [OPTION]... COMMAND INITIAL-ARGS...
```

- -E EOFString 指定逻辑 EOF 字符串以替换缺省的下划线(\_), xargs 命令读取标准输入直到达到 EOF 或指定的字符串
- -I replaceString 插入标准输入的每一行作为 command 参数的自变量, 把它插入每个发生的 replaceString 的 Argument 中
- -n number 指定传递给执行命令的参数个数, 默认是所有
- -p 每执行一个 argument 时询问一次用户
- -P 指定并行执行的进程数
- -a file 从文件中读入作为 stdin
- -s size 命令行的最大字符数，指的是 xargs 后面那个命令的最大命令行字符数
- -t 执行命令之前先打印执行命令

```bash
# -I 指定参数自变量
[root@localhost ~]# echo "file1 file2 file3" | xargs -t -I % sh -c 'touch %;ls -l %'

[root@localhost ~]# sh -c "touch file{1..3}";ls -l
```

- -d 设置自定义分隔符

```bash
[root@localhost ~]# echo -n file1#file2#file3#file4 | xargs -d \# -t touch

[root@localhost ~]# touch file{1..4}
```

### grep

> BRE 定义了 4 组元字符 `[ ]` `.` `^` `$`
> ERE 增加了 3 组元字符 `{ }` `()` `|`

```bash
grep [OPTION]... PATTERN [FILE]...
```

- -V,\-\-version 显示版本信息
- -c,\-\-count 统计符合字符串条件的行数
- -i,\-\-ignore-case 忽略大小写
- -e,\-\-regexp=PATTERN 使用正则表达式匹配

  ```bash
  # 使用 -e 参数 需要将正则中的部分字符转义才能使用
  [root@localhost ~]# grep -i -e "foo\|bar"
  ```

- -E,\-\-extended-regexp 使用扩展正则表达式匹配(ERE)

  ```bash
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
- -x,\-\-line-regexp 仅匹配整行
- -w,\-\-word-regexp 仅匹配整个单词
- -v,\-\-invert-match 显示不包括文本的所有信息
- -R,-r,\-\-recursive 递归的读取目录下的所有文件,包括子目录

```bash
grep -A 3 'pattern' file.txt # 向下搜索并显示匹配内容的后 3 行
grep -B 3 'pattern' file.txt # 向上搜索并显示匹配内容的前 3 行
grep -C 3 'pattern' file.txt 
```

#### 批量删除本地关联的 git 远程分支

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

```bash
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

```bash
# 拷贝文件, 可以使用原文件名也可以重新命名文件
[root@localhost ~]# scp -Cp /home/workspace/file1.txt root@192.168.1.3:/home/workspace/

# 拷贝目录
[root@localhost ~]# scp -rCp /home/workspace/ root@192.168.1.3:/home/workspace/
```

#### 远程复制到本地

```bash
# 拷贝文件, 可以使用原文件名也可以重新命名文件
[root@localhost ~]# scp -Cp root@192.168.1.3:/home/workspace/file1.txt /home/workspace/
# 拷贝目录
[root@localhost ~]# scp -rCp root@192.168.1.3:/home/workspace/ /home/workspace
```

### crontab 定时任务

- -u 指定用户
- -e 使用编辑器设置时程表
- -r 删除时程表
- -l 列出目前的时程表

|参数|取值|
|--|:--:|
|Minutes|0-59|
|Hours|0-23|
|Day of month|1-31|
|Month|1-12 or JAN-DEC|
|Day of week|0-6 or SUN-SAT|

#### 特殊字符

- \* 表示所有可能的值, 表示在所有时间点都执行任务
- \, 用于分隔多个值, 表示多个时间点执行任务
- \- 用于指定iyge范围内的连续值, 表示一个范围内的时间点执行任务
- \/ 用于指定一个步长, 表示每隔一定时间执行任务

```bash
*/5 * * * * # 表示每隔5分钟执行任务
30 0 1 * * # 表示每月第一天凌晨0点30分执行任务
10 1 1 1,4,7,10 * # 表示1,4,7,10月的第一天凌晨1点10分执行任务
0 1-3 * * 1 # 表示每周一周四的凌晨1点到3点执行任务
```

#### 每分钟向指定文件追加写入一条数据

```bash
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

### tar 归档

```bash
tar [OPTION...] [FILE]...
```

- \-u,\-\-update 用已打包的文件的较新版本更新 tar 文件
- \-d,\-\-diff,\-\-compare 将文件系统里的文件和 tar 文件里的文件进行比较
- \-\-delete 删除 tar 文件里面的文件, 不能用于已保存在磁带上的 tar 文件
- \-v,\-\-verbose 列出每一步处理涉及的文件的信息
- \-k,\-\-keep-old-files 不覆盖文件系统上已有的文件
- \-f,\-\-file 指定要处理的文件名
- \-w 每一步都要求确认

- \-\-atime-preserve 不改变转储文件的存取时间
- \-m,\-\-modification-time 当从一个档案中恢复文件时, 不使用新的时间标签

#### 压缩工具

- \-j,\-\-bzip2 调用 bzip2 执行压缩或解压缩
- \-J,\-\-xz,\-\-lzma 调用 XZ Utils 执行压缩或解压缩
- \-z,\-\-gzip,\-\-gunzip,\-\-ungzip 调用 gzip 执行压缩或解压缩
- \-Z,\-\-compress,\-\-uncompress 调用 compress 执行压缩或解压缩

#### 归档压缩

- \-c,\-\-create 创建新的 tar 文件

```bash
# 使用 gzip 压缩归档 workspace 目录
[root@centos7 ~]tar -czvf workspace.tar.gz ./workspace
```

#### 解压缩

- \-x,\-\-extract 解开 tar 文件

```bash
# 使用 gzip 解压缩到当前目录下的 redis-stable 目录
[root@centos7 ~]tar -xzvf redis-stable.tar.gz redis-stable
```

#### 列出文件

- \-t,\-\-list 列出压缩文件中的信息

```bash
[root@centos7 ~]tar -ztf archive.tar.gz # 输出压缩包的文件列表
```

#### 添加文件

- \-r,\-\-append 向压缩包中添加文件, 如果压缩包不存在则新建压缩包

```bash
# 向 archive.tar 压缩包中添加文件 file1.txt file2.txt
[root@centos7 ~]tar -rf archive.tar file1.txt file2.txt
```

#### 提取文件

- \-\-files-from=FILE 提取从 FILE 文件中列出的文件列表
- \-\-wildcards 使用通配符匹配文件名
- \-\-strip-components=NUMBER 去掉指定数量的目录层级
- \-C,\-\-directory DIR 转到指定的目录

```bash
# 从 redis-stable.tar.gz 包 提取 files.txt 中列出的文件列表 到当前目录下
[root@centos7 ~]tar -zxf redis-stable.tar.gz --files-from=files.txt -C ./

# 从 redis-stable.tar.gz 包 提取 redis.conf 文件 到当前目录下, 去掉 1 层目录
[root@centos7 ~]tar -zxf redis-stable.tar.gz --wildcards '*/redis.conf' --strip-components=1 -C ./

# 从 redis-stable.tar.gz 包 提取 sentinel.conf 文件 到当前目录下, 去掉 1 层目录
[root@centos7 ~]tar -zxf redis-stable.tar.gz --wildcards '*/sentinel.conf' --strip-components=1 -C ./ 
```

### curl

- \-\-help [all] For all options use the manual
- \-s,\-\-silent Silent mode
- \-v,\-\-verbose Make the operation more talkative
- \-#,\-\-progress-bar Display transfer progress as a bar
- \-L,\-\-location Follow redirects
- \-X,\-\-request \<method\> Specify request method to use
- \-H,\-\-header \<header\/@file\> Pass custom headers to server

```bash
curl --help all

curl --location --request POST 'https://developer.toutiao.com/api/apps/v2/jscode2session' \
--header 'Access-Token: clt.01*********3d3d' \
--header 'Content-Type: application/json' \
--data '{
  "name": "zhangsan",
  "age": 18
}'
```

- \-b,\-\-cookie \<data|filename\> Send cookies from string/load from file

```bash
curl --location --request POST 'https://developer.toutiao.com/api/apps/v2/jscode2session' \
--cookie 'name=zhangsan;age=18' \
-v

curl --location --request POST 'https://developer.toutiao.com/api/apps/v2/jscode2session' \
--header 'Cookie: name=zhangsan;age=18' \
-v
```

- \-d
  - \-\-data \<data\> HTTP POST data
  - \-\-data-ascii \<data\> HTTP POST ASCII data
  - \-\-data-binary \<data\> HTTP POST binary data
  - \-\-data-raw \<data\> HTTP POST data, '@' allowed
  - \-\-data-urlencode \<data\> HTTP POST data URL encoded

```bash
curl --location --request POST 'https://developer.toutiao.com/api/apps/v2/jscode2session' \
--header 'Access-Token: clt.01*********3d3d' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'client-key=tt10abc********' \
--data-urlencode 'client_secret=7820************' \
--data-urlencode 'code=ffab5ec*********' \
--data-urlencode 'grant_type=authorization_code'

curl --location --request POST 'https://developer.toutiao.com/api/apps/v2/jscode2session' \
--header 'Access-Token: clt.01*********3d3d' \
--header 'Content-Type: application/json' \
--data-raw '{
  "app_id": "tt*******",
  "app_name": "douyin",
  "path":"xxxx",
  “query”: "{xxx:xxxx}"，
  “expire_time”: 16444464021,
}'
```

- \-F

  - \-\-form \<name-content\> Specify multipart MIME data
  - \-\-form-escape Escape form fields using backslash
  - \-\-form-string \<name=string\> Specify multipart MIME data

```bash
curl --location --request POST 'https://developer.toutiao.com/api/apps/v2/jscode2session' \
--header 'Access-Token: clt.01*********3d3d' \
--header 'Content-Type: multipart/form-data' \
--header 'Accept: */*' \
--form 'appid=tt1yyyy' \
--form 'material_type="1000"' \
--form 'material_file=@"/Users/xxx.jpg"'
```

- \-o,\-\-output \<file\> Write to file instead of stdout
- \-\-output-dir \<dir\> Directory to save files in

### 批量重命名文件

```bash
#!/usr/bin/env bash
cd $1

count=1
for file in *.png; do
  if [ -e "$file" ]; then
    new_name="${count}.png"
    # 遍历重命名文件是否存在
    while [ -e "$new_name" ]; do
      ((count++))
      new_name="${count}.png"
    done
    # 重命名文件
    mv "$file" "$new_name"
    # 输出重命名信息
    echo "重命名 '$file' 为 '$new_name'"
    # 计数器
    ((count++))
  fi
done
if [ $count -eq 1 ]; then
  echo "没有找到 png 文件"
fi
```

### firewall-cmd

- \-\-check-config 检查永久配置是否有错误

- \-\-state 查看防火墙状态
- \-\-reload 重新加载配置
- \-\-permanent 配置永久有效

- \-\-get-zones 查看可用的区域
- \-\-set-default-zone 设置系统的默认区域
- \-\-get-active-zones 查看所有活动区域以及包含的接口
- \-\-zone=\<zone\> 指定区域
- \-\-info-zone=\<zone\> 查看指定区域的详细信息

- \-\-list-all 查看当前活动的防火墙规则
- \-\-list-service 查看允许的服务
- \-\-list-ports 查看允许的端口

- \-\-add-interface 添加接口
- \-\-add-service 添加服务, 允许指定的服务通过防火墙
- \-\-add-port  添加端口, 允许指定的接口通过防火墙
- \-\-remove-service 移除服务
- \-\-remove-port 移除端口
- \-\-set-log-denied=[all|off] 启用/禁用防火墙的日志记录功能

```bash
firewall-cmd --check-config # 检查永久配置是否有错误

firewall-cmd --get-zones # 查看可用的区域
firewall-cmd --zone=public --list-all # 列出指定区域的所有规则
firewall-cmd --zone-public --list-service # 列出指定区域的服务

firewall-cmd --zone=public --add-service=http --permanent # 允许 http 服务
firewall-cmd --zone=public --add-service=http --permanent # 允许 https 服务
firewall-cmd --zone=public --add-port=22 --permanent # 允许 22 端口
firewall-cmd --reload # 重新加载配置

firewall-cmd --zone=public --remove-port=22 --permanent # 移除 22 端口
firewall-cmd --zone-public --remove-service=http --permanent # 移除 http 服务
```
