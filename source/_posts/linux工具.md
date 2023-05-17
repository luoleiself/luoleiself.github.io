---
title: linux工具
date: 2022-05-31 14:40:43
categories:
  - [linux, shell]
tags:
  - linux
  - shell
---

### awk

awk 是一种可以对文本和数据进行处理的编程语言, 默认情况下, awk 文件的每一行都被视为一条记录, 然后 awk 记录进一步分解成一系列的字段

```bash
awk '{ sum += $1 }; END { print sum }' file
awk -F : '{ print $1 }' /etc/passwd
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
git branch -r | awk 'BEGIN {print "hello awk, I am coming\n"}END{print "hello awk, good bye\n"}{printf "NF--%s NR--%s FNR--%s $0--%s\n",NF,NR,FNR,$0;}'
```

### xargs

将参数列表转换成小块分段传递给其他命令, 以避免参数列表过长的问题, 可单独使用, 也可以使用管道符、重定位符等其他命令配合使用

```bash
xargs [OPTION]... COMMAND INITIAL-ARGS...
```

- -E EOFString 指定逻辑 EOF 字符串以替换缺省的下划线(\_), xargs 命令读取标准输入直到达到 EOF 或指定的字符串
- -I replaceString 插入标准输入的每一行作为 command 参数的自变量, 把它插入每个发生的 replaceString 的 Argument 中
- -n number 指定传递给执行命令的参数个数, 默认是所有
- -p 每执行一个 argument 时询问一次用户
- -a file 从文件中读入作为 stdin
- -s size 命令行的最大字符数，指的是 xargs 后面那个命令的最大命令行字符数
- -t 执行命令之前先打印执行命令

```bash
# -I 指定参数自变量
echo "file1 file2 file3" | xargs -t -I % sh -c 'touch %;ls -l %'
sh -c touch file1 file2 file3;ls -l file1 file2 file3
```

- -d 设置自定义分隔符

```bash
echo -n file1#file2#file3#file4 | xargs -d \# -t touch
touch file1 file2 file3 file4
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
  grep -i -e "foo\|bar" # 使用 -e 参数 需要将正则中的部分字符转义才能使用
  ```

- -E,\-\-extended-regexp 使用扩展正则表达式匹配(ERE)

  ```bash
  grep -i -E "foo|bar" # 此处不需要进行字符转义
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

### firewall-cmd 防火墙

- \-\-permanent # 永久修改
- \-\-reload # 重新加载防火墙配置

```bash
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

### systemctl

systemd 系统控制和服务管理工具的主命令, systemd 开启和监督整个系统是基于 Unit 的概念, Unit 是由一个与配置文件名同名的名字和类型组成

| Runlevel | Target Unit                         | Description                              |
| :------: | ----------------------------------- | ---------------------------------------- |
|    0     | runlevel0.target, poweroff.target   | Shut down and power off the system       |
|    1     | runlevel1.target, rescue.target     | Set up a rescue shell                    |
|    2     | runlevel2.target, multi-user.target | Set up a non-graphical multi-user system |
|    3     | runlevel3.target, multi-user.target | Set up a non-graphical multi-user system |
|    4     | runlevel4.target, multi-user.target | Set up a non-graphical multi-user system |
|    5     | runlevel5.target, graphical.target  | Set up a graphical multi-user system     |
|    6     | runlevel6.target, reboot.target     | Shut down and reboot the system          |

- start 启动服务
- stop 停止服务
- reload 重新加载配置文件不重启服务
- restart 重启服务
- enable 允许开机启动
- disable 取消开机启动
- status 查看服务的状态
- is-active 查看服务是否正在运行
- is-enabled 查看服务是否开机启动
- show 查看服务的详细信息

- default 进入系统默认模式
- rescue 进入系统救援模式
- emergency 进入系统应急模式

- halt 关闭系统
- poweroff 关闭系统
- reboot 重启系统

- daemon-reload 重新加载 systemd 系统管理配置项
- daemon-reexec 重新执行 systemd 系统管理器

#### Unit 类型

每个配置单元都有一个对应的配置文件

- service 代表一个后台服务进程, 例如 mysqld、nginx
- socket 此类配置单元封装系统和互联网中的一个套接字, 每个套接字配置单元都有一个相应的服务配置单元, 相应的服务在第一个连接进入套接字时就会自动启动(例如 nscd.socket 在有新连接后会启动 nscd.service)
- device 此类配置单元封装一个存在于 linux 设备树中的设备, 每个使用 udev 规则标记的设备都会在 systemd 中作为一个设备配置单元出现
- mount 此类配置单元封装文件系统结构层次中的一个挂载点, systemd 将对这个挂载点进行监控和管理, systemd 会将 /etc/fstab 中的条目都转换为挂载点, 并在开机时处理
- automount 此类配置单元封装文件系统结构层次中的一个自动挂载点, 每个自动挂载配置单元对应一个挂载配置单元
- swap 和挂载配置单元类似, 可以用交换配置单元来定义系统中的交换分区, 让这些交换分区在启动时被激活
- target 此类配置单元为其它配置单元进行逻辑分组, 它们本身没有任何行为, 只是引用其他配置单元, 这样就可以对配置单元做一个统计的控制
- timer 定时器配置单元用来定时触发用户定义的操作, 这类配置单元取代了 atd, crond 等传统的定时服务
- snapshot 与 target 配置单元类似, 快照是一组配置单元, 保存了系统当前的运行状态

#### 配置文件项

##### Unit

用来定义单元的元数据, 以及配置与其他 Unit 的关系

- Description 当前服务的简单描述
- Documentation 文档地址
- Requires 表示强依赖关系, 即某些服务停止运行或退出, 该服务也必须停止或退出
- Wants 表示弱依赖关系, 即某些服务停止运行或退出不会影响该服务继续运行
- After 表示在什么服务之后启动
- Before 表示在什么服务之前启动
- Conflicts 表示指定的 Unit 不能与当前 Unit 同时运行
- Condition 表示当前 Unit 运行必须满足的条件, 否则不会运行
- Assert 表示当前 Unit 运行必须满足的条件, 否则会报启动失败

##### Install

定义如何安装此配置文件

- Alias 为当前 Unit 定义一个用于启动的别名
- Also 当前 Unit 被激活时, 同时被激活的其他 Unit
- RequiredBy 当前 Unit 被允许运行需要的一系列依赖 Unit, RequiredBy 列表从 Require 获得依赖信息
- DefaultInstance 实例单元的限制, 这个选项指定如果 Unit 被允许运行时的默认实例
- WantedBy 表示该服务所在的 target, target 表示一组服务, 大多的服务都附在 multi-user.target 组, 这个组的所有服务都将开机启动

##### Service

配置 service, 只有 service 类型的 Unit 才有此项

- Type 定义启动时的进程行为
  - simple 默认值, 执行 ExecStart 指定的命令, 启动主进程
  - forking 以 fork 方式从父进程创建子进程, 此时父进程将会退出, 子进程成为主进程
  - oneshot 与 simple 类似, 但只执行一次, Systemd 等待此进程执行完后, 才启动其他服务
  - dbus 与 simple 类似, 但会等待 D-Bus 信号后启动
  - notify 与 simple 类似, 启动结束后会发出通知信号, Systemd 再启动其他服务
  - idle 与 simple 类似, 等待其他任务都执行完成, 才会启动该服务
- ExecStart 定义启动进程时执行的命令
- ExecReload 定义重启服务时执行的命令
- ExecStop 定义停止服务时执行的命令
- ExecStartPre 定义启动服务之前执行的命令
- ExecStartPost 定义启动服务之后执行的命令
- ExecStopPost 定义停止服务之后执行的命令
- KillMode 定义 Systemd 如何停止服务
  - control-group 默认值, 当前控制组内的所有子进程都会被杀掉
  - process 只杀主进程
  - mixed 主进程将受到 SIGTERM 信号, 子进程受到 SIGKILL 信号
  - none 不杀掉任何进程, 只执行服务的 stop 命令
- Restart 定义 Systemd 重启服务的方式
  - no 默认值, 退出后不会重启
  - on-success 只有正常退出时(退出状态码为 0), 才会重启
  - on-failure 非正常退出时(退出状态码非 0), 包括信号被终止和超时才会重启
  - on-abnormal 只有被信号终止和超时才会重启
  - on-abort 只有在收到没有捕捉到的信号终止时才会重启
  - on-watchdog 超时退出才会重启
  - always 不管什么原因总是重启
- RestartSec 定义 Systemd 重启服务之前等待的秒数
- TimeoutSec 定义 Systemd 停止服务之前等待的秒数
- user 定义服务的用户名
- WorkingDirectory 定义服务的安装目录
- EnvironmentFile 定义环境变量配置文件
- PrivateTmp 定义是否分配独立空间

##### Timer

- OnBootSec 当开机多久后才执行当前 Unit
- OnUnitActiveSec 这个 timer 配置文件所管理的那个 Unit 在最后一次启动后, 相隔多久再执行一次
- OnUnitInactiveSec 这个 timer 配置文件所管理的那个 Unit 在最后一次停止后, 相隔多久再执行一次
- OnCalendar 使用实际时间(非循环时间)的方式来启动服务
- OnActiveSec 当 timers.target 启动多久后才执行当前 Unit
- OnStartupSec 当 Systemd 第一次启动后多久才执行当前 Unit

#### 配置 service

- Redis.service

```conf
[Unit]
Description=redis-server
After=network.target

[Service]
Type=forking
ExecStart=/usr/local/bin/redis-server /root/workspace/redis6379.conf
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

#### 配置 target

- multi-user.target

```conf
#  This file is part of systemd.
#
#  systemd is free software; you can redistribute it and/or modify it
#  under the terms of the GNU Lesser General Public License as published by
#  the Free Software Foundation; either version 2.1 of the License, or
#  (at your option) any later version.

[Unit]
Description=Multi-User System
Documentation=man:systemd.special(7)
Requires=basic.target
Conflicts=rescue.service rescue.target
After=basic.target rescue.service rescue.target
AllowIsolate=yes
```

#### 配置 timer

- systemd-tmpfiles-clean.timer

```conf
#  This file is part of systemd.
#
#  systemd is free software; you can redistribute it and/or modify it
#  under the terms of the GNU Lesser General Public License as published by
#  the Free Software Foundation; either version 2.1 of the License, or
#  (at your option) any later version.

[Unit]
Description=Daily Cleanup of Temporary Directories
Documentation=man:tmpfiles.d(5) man:systemd-tmpfiles(8)

[Timer]
OnBootSec=15min # 当开机多久后才执行当前 Unit
OnUnitActiveSec=1d  # 这个 timer 配置文件所管理的那个 Unit 在最后一次启动后, 相隔多久再执行一次
#OnUnitInactiveSec # 这个 timer 配置文件所管理的那个 Unit 在最后一次停止后, 相隔多久再执行一次
#OnCalendar # 使用实际时间(非循环时间)的方式来启动服务
#OnActiveSec  # 当 timers.target 启动多久后才执行当前 Unit
#OnStartupSec # 当 Systemd 第一次启动后多久才执行当前 Unit
```

#### 配置 mount

- tmp.mount

```conf
#  This file is part of systemd.
#
#  systemd is free software; you can redistribute it and/or modify it
#  under the terms of the GNU Lesser General Public License as published by
#  the Free Software Foundation; either version 2.1 of the License, or
#  (at your option) any later version.

[Unit]
Description=Temporary Directory
Documentation=man:hier(7)
Documentation=http://www.freedesktop.org/wiki/Software/systemd/APIFileSystems
ConditionPathIsSymbolicLink=!/tmp
DefaultDependencies=no
Conflicts=umount.target
Before=local-fs.target umount.target

[Mount]
What=tmpfs
Where=/tmp
Type=tmpfs
Options=mode=1777,strictatime

# Make 'systemctl enable tmp.mount' work:
[Install]
WantedBy=local-fs.target
```
