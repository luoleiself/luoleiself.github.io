---
title: shell语法
date: 2022-04-28 19:55:46
categories:
  - [linux, shell]
tags:
  - linux
  - shell
---

### shell 重置密码

1. 开机按 `e` 键进入内核编辑
2. 光标移动到倒数第二段 `Linux16` 末尾添加 `init=/bin/sh`
3. `ctrl + x` 进行引导启动, 成功进入命令提示界面
4. 输入 `mount -o remount, rw /` 挂载根目录
5. 使用 `passwd` 命令修改指定用户密码
6. 直到提示 `passwd: all authentication tokens updated successfully.`
7. 输入 `touch /.autorelabel` 回车
8. 输入 `exec /sbin/init` 回车重启系统

### 命令

- export 导出全局变量
- declare 声明变量
- unset 删除变量
- local 声明局部变量,一般用于函数内部

<!-- more -->

### 运算符

#### [ ] | [[]]

运算表达式两边必须要有空格

#### \${}

用于变量替换, 便于理解

```bash
A=B
echo $AB
echo ${A}B
BB
```

#### \$[ ] | $(( )) | (()) 高版本支持

用于进行数学计算，对于浮点数是当作字符串处理

```bash
a=5;b=7;c=2
echo $((a+b*c))
19
echo $[a+b*c]
19
echo $(($a+$b*$c))
19
echo $[$a+$b*$c]
19
```

#### \$( )

命令替换

```bash
[root@localhost ~]# docker rm -f $(docker ps -aq) # 移除所有容器
```

#### 算术运算符

| 算术运算符 |                      说明                       |          举例           |
| :--------: | :---------------------------------------------: | :---------------------: |
|     +      |                      加法                       |     echo \$((a+b))      |
|     -      |                      减法                       |     echo \$((a-b))      |
|     \*     |                      乘法                       |     echo \$((a\*b))     |
|     /      |                      除法                       |     echo \$((a/b))      |
|     %      |                      取余                       |     echo \$((a%b))      |
|     =      |                      赋值                       |           a=b           |
|     ==     |   相等，判断两个数字是否相等，相等则返回 true   | echo $[ ${a} == \${b} ] |
|     !=     | 不相等，判断两个数字是否相等，不相等则返回 true | echo $[ ${a} != \${b} ] |

#### 关系运算符

| 在 [ ] 中使用的比较符 | 在 (( )) 和 [[]] 中使用的比较符 |                说明                |                           举例                            |
| :-------------------: | :-----------------------------: | :--------------------------------: | :-------------------------------------------------------: |
|          -eq          |               ==                |       equal 的缩写，表示相等       | [ ${a} -eq ${b} ] \| [[${a} == ${b}]] \| ((${a} == ${b})) |
|          -ne          |               !=                |    not equal 的缩写，表示不相等    | [ ${a} -ne ${b} ] \| [[${a} != ${b}]] \| ((${a} != ${b})) |
|          -gt          |                >                |   greater than 的缩写，表示大于    |  [ ${a} -gt ${b} ] \| [[${a} > ${b}]] \| ((${a} > ${b}))  |
|          -ge          |               >=                | greater equal 的缩写，表示大于等于 |  [ ${a} -ge ${b} ] \| [[${a} >= ${b}]]\|((${a} >= ${b}))  |
|          -lt          |                <                |    least than 的缩写，表示小于     |  [ ${a} -lt ${b} ] \| [[${a} < ${b}]] \| ((${a} < ${b}))  |
|          -le          |               <=                |  least equal 的缩写，表示小于等于  |  [ ${a} -le ${b} ] \| [[${a} <= ${b}]]\|((${a} <= ${b}))  |

```bash
if [ $string -ne 'abc' ]; then
  echo "Not equal"
else
  echo "Equal"
fi
if [ $a -ge $b ]; then
  echo "Greater equal"
else
  echo "Not greater equal"
fi
```

#### 逻辑运算符

| 在 [ ] 中使用的比较符 | 在 (( )) 和 [[]] 中使用的比较符 |                   说明                    |                 举例                  |
| :-------------------: | :-----------------------------: | :---------------------------------------: | :-----------------------------------: |
|         赋值          |              a=10               |                   b=25                    |                                       |
|          -a           |               &&                |  与运算，两个表达式都为 true,才返回 true  | [ $a -lt 20 -a $b -gt 20 ] 返回 true  |
|          -o           |              \|\|               | 或运算，有一个表达式都为 true,则返回 true | [ $a -lt 20 -o $b -gt 100 ] 返回 true |
|           !           |                !                |    非运算，表达式为 true,则返回 false     |         [ ! false ] 返回 true         |

```bash
a=5
b=12
if [ $a -lt 50 -a  $b -gt 8 ]; then
  echo "And(-a) expr result is true"
else
  echo "And(-a) expr result is false"
fi
if [ $a -lt 50 -o $b -gt 12 ]; then
  echo "Or(-o) expr result is true"
else
  echo "Or(-o) expr result is false"
fi
```

#### 字符串运算符, [ ] 字符串运算不需要转义

| 字符串运算符 |                     说明                     |       举例       |
| :----------: | :------------------------------------------: | :--------------: |
|      =       |   检测两个字符串是否相等，相等则返回 true    | [ ${a} = ${b} ]  |
|      !=      |  检测两个字符串是否相等，不相等则返回 true   | [ ${a} != ${b} ] |
|      -z      |   检测字符串长度是否为 0，为 0 则返回 true   |   [ -z ${b} ]    |
|      -n      | 检测字符串长度是否不为 0，不为 0 则返回 true |   [ -n ${b} ]    |
|     str      | 检测字符串是否为 null，不为 null 则返回 true |     [ ${b} ]     |

```bash
a=hello
b=world
if [ a = b ]; then
  echo "string a equal string b"
else
  echo "string a not equal string b"
fi
```

#### 文件测试运算符

| 文件测试运算符 |                                  说明                                   |          举例          |
| :------------: | :---------------------------------------------------------------------: | :--------------------: |
|       -b       |              检测文件是否是块设备文件，如果是，则返回 true              | [ -b $file ] 返回 true |
|       -c       |             检测文件是否是字符设备文件，如果是，则返回 true             | [ -c $file ] 返回 true |
|       -d       |               检测文件是否是目录文件，如果是，则返回 true               | [ -d $file ] 返回 true |
|       -f       | 检测文件是否是普通文件（既不是目录也不是设备文件），如果是，则返回 true | [ -f $file ] 返回 true |
|       -g       |             检测文件是否设置了 SGID 位，如果是，则返回 true             | [ -g $file ] 返回 true |
|       -k       |       检测文件是否设置了粘着位（stucky Bit），如果是，则返回 true       | [ -k $file ] 返回 true |
|       -p       |                检测文件是否具名管道，如果是，则返回 true                | [ -p $file ] 返回 true |
|       -u       |             检测文件是否设置了 SUID 位，如果是，则返回 true             | [ -u $file ] 返回 true |
|       -r       |                  检测文件是否可读，如果是，则返回 true                  | [ -r $file ] 返回 true |
|       -w       |                  检测文件是否可写，如果是，则返回 true                  | [ -w $file ] 返回 true |
|       -x       |                 检测文件是否可执行，如果是，则返回 true                 | [ -x $file ] 返回 true |
|       -s       |   检测文件是否为不为空（文件大小是否不为 0），如果不为 0，则返回 true   | [ -s $file ] 返回 true |
|       -e       |            检测文件(包括目录)是否存在，如果存在，则返回 true            | [ -e $file ] 返回 true |
|       -a       |     检测文件(包括目录)是否存在(此命令已废弃)，如果存在，则返回 true     | [ -e $file ] 返回 true |

```bash
if [ -e .node ]; then
  echo 'this file is exists'
else
  echo 'this file not exists'
fi
if [ -s .zshrc ]; then
  echo 'file not empty'
else
  echo 'file is empty'
fi
```

#### 重定向 | 管道 | 流

##### 重定向

- `>` 输出重定向, 文件覆盖
- `>>` 输出重定向, 文件末尾追加
- `2>` 标准错误输出, 文件覆盖
- `2>>` 标准错误输出, 文件末尾追加
- `2>&1` 标准输出和标准错误输出
- `<` 输入重定向
- `<<` 输入重定向

##### 管道

- `|` 连接两个命令, 第一个命令的输出作为第二个命令的输入

### 流程

#### if

```shell
# 当 then 单独另写一行时, 分号不能省略
if [-e /root/workspace/test.txt ]; then
  printf "hello world %s %s\n" $(/bin/date +"%Y-%m-%d %H:%M:%S")
# 当 then 单独另写一行时, 分号不能省略
elif [ -s /root/workspace/test.txt ]; then
  printf "hello world\n"
else
  printf "hello gg\n"
fi
```

#### case

```shell
a=20
case $a in
  10)
    echo "a的值为 10"
  ;;
  20)
    echo "a的值为 20" # 输出 a 的值为 20
  ;;
  *)
    echo "a的值不是10也不是20"
  ;;
esac
```

#### for

```shell
# 当 do 单独另写一行时, 分号不能省略
for i in {1..10}; do
  # 依次输出 for do 1 到 10
  echo "for do " ${i}
done
```

#### while

```shell
j=1
# 当 do 单独另写一行时, 分号不能省略
while [ $j -lt 10 ]; do
  # 依次输出 while do 1 到 9
  echo "while do " ${j}
  j=$[j+1]
done
```

#### until

```shell
k=1
# 当 do 单独另写一行时, 分号不能省略
until [ $k -gt 10 ]; do
  # 依次输出 until do 1 到 10
  echo "until do " ${k}
  k=$[k+1]
done
```

### 数组

#### 一维数组

- bash 只支持一维数组
- 初始化时不需要定义数组的大小, 定义时用小括号将用空格分隔的元素包含起来
- 数组元素的下标由 0 开始

```shell
arr=(1 2 3 'a')
echo "arr[0]=" ${arr[0]}
echo "arr[1]=" ${arr[1]}
echo "arr[2]=" ${arr[2]}
echo "arr[3]=" ${arr[3]}
echo "arr[*]=" ${arr[*]}  # 输出所有元素
echo "arr[@]=" ${arr[@]}  # 输出所有元素
echo "arr的键为 " ${!arr[*]}
# arr的键为  0 1 2 3
echo "arr的长度为 " ${#arr[*]}
# 获取数组的长度与获取字符串长度的方法相同

str=helloworld
echo "str 的长度为 " ${#str}  # str 的长度为 10
```

#### 关联数组

- 关联数组, 使用 `declare -A` 声明

```shell
declare -A site=(['google']='www.google.com' ['baidu']='www.baidu.com' ['taobao']='www.taobao.com')
echo "site['google']=" ${site['google']}
# site['google']= www.google.com
echo "site['baidu']=" ${site['baidu']}
# site['baidu']= www.baidu.com
echo "site['taobao']=" ${site['taobao']}
# site['taobao']= www.taobao.com
echo "site的所有元素是 " ${site[*]}
# site的所有元素是  www.google.com www.taobao.com www.baidu.com
echo "site的所有元素是 " ${site[@]}
# site的所有元素是  www.google.com www.taobao.com www.baidu.com
echo "site的键为 " ${!site[*]}
# site的键为  google taobao baidu
echo "site的长度为 " ${#site[*]}
# site的长度为  3
```

#### 数组操作

- \*|@ 获取数组的所有元素, `${arr[*]}`
- ! 获取数组的所有键, `${!arr[*]}`
- \# 获取数组的长度, `${#arr[*]}`

### 函数

- function 是关键字, 用来定义函数
- name 是函数名
- statements 函数体中的执行语句
- return value 函数的返回值, 一般表示函数的返回状态, 0 表示成功, 其他值表示失败, 只能是 0 - 255 之间的数字

```shell
# 标准语法
[function] function_name [()] {
  statements
  [return value]
}
```

- 定义函数时, 关键字 function 和 () 可以二选一
- 定义函数时, 不需要提前指定参数, 在函数体内使用参数时, 使用 `特殊变量` 获取
- 函数调用时, 函数名后面不需要带小括号, 如果有参数时，多个参数之间用空格分隔
- 获取函数的返回值
  - 一种是借助全局变量, 将得到的结果赋值给全局变量
  - 一种是在函数内使用 `echo`, `printf` 命令将结果输出, 在函数外部使用 `$( )` 或者 `\`\`` 捕获结果

```shell
#!/bin/bash
function dateFormat() {
  echo $(/bin/date +"hello crontab %Y-%m-%d %H:%M:%S")
}

echo $(dateFormat) # 获取函数的结果
```

#### 特殊变量

|   变量   |                                                含义                                                 |
| :------: | :-------------------------------------------------------------------------------------------------: |
|    $0    |                                          当前脚本的文件名                                           |
| $n(n>=1) |                        传递给脚本或函数的参数, n 是一个数字, 表示第几个参数                         |
|    $#    |                                     传递给脚本或函数的参数个数                                      |
|   $\*    |                     传递给脚本或函数的所有参数, 当前脚本的所有参数作为一个参数                      |
|    $@    | 传递给脚本或函数的所有参数,当前脚本的所有参数分别作为一个参数, 当被双引号包含时, $@ 和 $\* 有所不同 |
|    $?    |                                  上个命令的退出状态,或函数的返回值                                  |
|    $$    |                   当前 Shell 进程 ID, 对于 Shell 脚本, 就是这些脚本所在的进程 ID                    |
