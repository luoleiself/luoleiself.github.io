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

#### 批量删除本地关联的 git 远程分支

- awk 和 xargs 命令结合使用

  - xargs -t 执行命令之前先打印执行命令
  - xargs -n 指定传递给执行命令的参数个数, 默认是所有
  - xargs -I 同时运行多个命令, 并替换所有匹配的项为传递给 xargs 的参数
  - xargs -p 每执行一个 argument 时询问一次用户
  - xargs -a file 从文件中读入作为 stdin
  - xargs -s num 命令行的最大字符数，指的是 xargs 后面那个命令的最大命令行字符数

    ```shell
    echo "file1 file2 file3"| xargs -t -I % sh -c 'touch %;ls -l %'
    sh -c touch file1 file2 file3;ls -l file1 file2 file3
    ```

  - xargs -d 设置自定义分隔符

    ```shell
    echo -n file1#file2#file3#file4|xargs -d \# -t touch
    touch file1 file2 file3 file4
    ```

- 不带主机名的过滤

```shell
# 格式化输出所有本地关联的分支名
$ git branch -r | \
  grep 'origin/feature*' | \
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
  grep 'origin/revert*' | \
  awk '{printf "%4s\n",$0;}' | \
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
