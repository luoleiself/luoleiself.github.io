---
title: systemd
date: 2022-05-31 14:40:43
categories:
  - [linux, shell]
tags:
  - linux
  - shell
---

## systemd

systemd(system daemon)是 linux 下的一种 init 软件, 提供更优秀的框架以表示系统服务间的依赖关系, 并依此实现系统初始化时服务的并行启动, 同时达到降低 shell 的系统开销的效果, 最终代替常用的 System V 与 BSD 风格 init 程序

- 采用 socket 激活式与总线激活式服务, 以提高相互依赖的各服务的并行运行性能
- 采用 cgroup 代替 PID 来追踪进程, 依此即使是两次 fork 之后生成的守护进程也不会脱离 systemd 的控制

### CGroup

> cgroup 是 linux 内核的一个功能, 用来限制、控制与分离一个进程组的资源(如 CPU、内存、磁盘输入输出等)

cgroup 是 linux 内核提供的一种机制, 这种机制可以根据需求把一系列系统任务及其子任务整合(或分隔)到按资源划分等级的不同组内, 从而为系统资源管理提供一个统一的框架. 简单说, cgroup 可以限制、记录任务组所使用的物理资源, 本质上来说, cgroup 是内核附加在程序上的一系列钩子(hook), 通过程序运行时对资源的调度触发相应的钩子以达到资源追踪和限制的目的.

#### 作用

- 资源限制: cgroup 可以对任务需要的资源总额进行限制, 例如设定任务运行使用的内存上限, 一旦超出就触发 OOM
- 优先级分配: 通过分配的 CPU 时间片数量和磁盘 IO 带宽, 实际上就等于控制了任务运行的优先级
- 资源统计: cgroup 可以统计系统的资源使用量, 例如 CPU 使用时长、内存使用量等
- 任务控制: cgroup 可以对任务执行挂起、恢复等操作

<!-- more -->

#### 概念

- Task 表示系统中的进程, cgroup 对 task 对其 cpu, mem 等资源进行限制
- Subsystem 子系统, 表示一个资源调用控制器, cgroup 支持 cpu,mem 等 subsystem
  - blkio 对块设备的 IO 进行限制
  - cpu 限制 CPU 时间片的额分配, 和 cpuacct 挂载在同一目录
  - cpuacct 生成 cgroup 种的任务占用 CPU 资源的报告, 与 CPU 挂载在同一目录
  - cpuset 给 cgroup 中的任务分配独立的 CPU(多处理器系统)和内存节点
  - devices 允许或禁止 cgroup 中的任务访问块设备
  - freezer 暂停/恢复 cgroup 中的任务
  - hugetlb 限制使用的内存页数量
  - memory 对 cgroup 中的任务的可用内存进行限制, 并自动生成资源占用报告
  - net_cls 使用等级识别符(classid)标记网络数据包, 让 linux 流量控制器(tc 指令)可以识别来自特定 cgroup 任务的数据包, 并进行网络限制
  - net_prio 允许基于 cgroup 设置网络流量的优先级
  - perf_event 允许使用 perf 工具来监控 cgroup
  - pids 限制任务的数量
- Controller Group 控制, 对一种或多种资源设置限制, 是 group 进行资源控制的基本单位, task 可以加入到某个控制组种或一个控制组迁移到另一个
- Hierarchy 由一系 Controller Group 组成的树结构

### 架构图

![systemd-1](/images/systemd-1.webp)

### hostnamectl

管理当前主机信息

- status 查看当前主机的设置

```shell
[root@centos7 workspace]# hostnamectl status
  Static hostname: centos7.localdomain
        Icon name: computer-vm
          Chassis: vm
      Machine ID: afcca427b44c4f139ef788ed3b33b7e1
          Boot ID: 312745f18eaa4b3eb809d0f361ad43bc
  Virtualization: kvm
Operating System: CentOS Linux 7 (Core)
      CPE OS Name: cpe:/o:centos:centos:7
          Kernel: Linux 3.10.0-1160.90.1.el7.x86_64
    Architecture: x86-64
```

- set-hostname NAME 设置系统主机名
- set-icon-name NAME 设置主机的图标名称
- set-chassis NAME 设置主机的基础架构名称
- set-deployment NAME 设置主机的部署环境
- set-location NAME 设置主机的位置

### localectl

管理本地化设置

- status 查看本地化配置项

```shell
[root@centos7 workspace]# localectl status
System Locale: LANG=zh_CN.utf8
    VC Keymap: us
  X11 Layout: n/a
```

- list-locales 查看系统中的本地化配置
- set-locale LOCALE 设置系统的本地化

### timedatectl

管理系统的日期时间设置

RTC(real-time clock) 指硬件时间(BIOS 时间), 专用于记录时间, 有电池供电, 不受服务器和操作系统的开启关闭影响

NTP(network time protocol)网络时间协议, 用来同步化计算机时间的一种协议, 提高精准度的时间校正

- status 显示日期时间设置

```shell
[root@centos7 workspace]# timedatectl status
      Local time: Thu 2023-05-18 14:49:36 CST
  Universal time: Thu 2023-05-18 06:49:36 UTC
        RTC time: Thu 2023-05-18 14:49:36
       Time zone: Asia/Shanghai (CST, +0800)
     NTP enabled: yes
NTP synchronized: yes
 RTC in local TZ: yes
      DST active: n/a

Warning: The system is configured to read the RTC time in the local time zone.
         This mode can not be fully supported. It will create various problems
         with time zone changes and daylight saving time adjustments. The RTC
         time is never updated, it relies on external facilities to maintain it.
         If at all possible, use RTC in UTC by calling
         'timedatectl set-local-rtc 0'.
```

- set-time TIME 设置系统时间
- set-timezone ZONE 设置系统时区
- list-timezones 显示系统支持的时区

```shell
[root@centos7 workspace]# timedatectl set-timezone Asia/Shanghai # 设置时区

[root@centos7 workspace]# timedatectl list-timezones
Africa/Abidjan
Africa/Accra
Africa/Addis_Ababa
Africa/Algiers
Africa/Asmara
Africa/Bamako
Asia/Shanghai
```

- set-local-rtc BOOL 设置本地时间
- set-ntp BOOL 开启 NTP 同步

```shell
[root@centos7 workspace]# timedatectl set-local-rtc 1 # 设置本地时间
[root@centos7 workspace]# timedatectl set-local-rtc 0 # 设置 UTC 时间

[root@centos7 workspace]# timedatectl set-ntp 1 # 开启 NTP 同步
```

### loginctl

常看当前登录的用户

#### Session

- list-sessions 显示会话列表

```shell
[root@centos7 workspace]# loginctl list-sessions
SESSION        UID USER             SEAT
      8       1000 vagrant
      9       1000 vagrant

2 sessions listed.
```

- session-show [ID...] 显示会话的状态

```shell
[root@centos7 workspace]# loginctl session-status 9
9 - vagrant (1000)
           Since: Thu 2023-05-18 10:13:35 CST; 4h 53min ago
          Leader: 3294 (sshd)
          Remote: 10.0.2.2
         Service: sshd; type tty; class user
           State: active
            Unit: session-9.scope
                  ├─3294 sshd: vagrant [priv]
                  ├─3297 sshd: vagrant@pts/1
                  ├─3298 -bash
                  ├─3531 su - root
                  ├─3535 -bash
                  ├─4673 loginctl session-status 9
                  └─4674 less

May 18 10:13:35 centos7.localdomain systemd[1]: Started Session 9 of user vagrant.
May 18 10:13:35 centos7.localdomain sshd[3294]: pam_unix(sshd:session): session opened for user vagrant by (uid=0)
May 18 10:13:43 centos7.localdomain su[3531]: (to root) vagrant on pts/1
May 18 10:13:43 centos7.localdomain su[3531]: pam_unix(su-l:session): session opened for user root by vagrant(uid=1000)
```

- show-session [ID...] 显示会话的属性

```shell
[root@centos7 workspace]# loginctl show-session 9
Id=9
User=1000
Name=vagrant
Timestamp=Thu 2023-05-18 10:13:35 CST
TimestampMonotonic=3227303782
VTNr=0
Remote=yes
RemoteHost=10.0.2.2
Service=sshd
Scope=session-9.scope
Leader=3294
Audit=9
Type=tty
Class=user
Active=yes
State=active
IdleHint=no
IdleSinceHint=0
IdleSinceHintMonotonic=0
LockedHint=no
```

- activate [ID] 激活会话
- lock-session [ID...] 屏幕锁定一个或多个会话
- unlock-session [ID...] 屏幕解锁一个或多个会话
- lock-sessions 屏幕锁定当前所有的会话
- unlock-sessions 屏幕解锁当前所有的会话
- terminate-session ID... 终止退出多个会话
- kill-session ID... 终止退出多个会话

#### User

- list-users 显示所有用户

```shell
[root@centos7 workspace]# loginctl list-users
  UID USER
1000 vagrant

1 users listed.
```

- user-status [USER...] 显示用户状态

```shell
[root@centos7 workspace]# loginctl user-status vagrant
vagrant (1000)
           Since: Thu 2023-05-18 10:13:33 CST; 5h 4min ago
           State: active
        Sessions: 46 *9
            Unit: user-1000.slice
                  ├─session-46.scope
                  │ ├─4951 sshd: vagrant [priv]
                  │ ├─4954 sshd: vagrant@pts/0
                  │ ├─4955 -bash
                  │ ├─5236 su - root
                  │ └─5240 -bash
                  └─session-9.scope
                    ├─3294 sshd: vagrant [priv]
                    ├─3297 sshd: vagrant@pts/1
                    ├─3298 -bash
                    ├─3531 su - root
                    ├─3535 -bash
                    ├─5452 loginctl user-status vagrant
                    └─5453 less

May 18 15:11:35 centos7.localdomain su[4995]: pam_unix(su-l:session): session opened for user root by vagrant(uid=1000)
May 18 15:12:49 centos7.localdomain su[4995]: pam_unix(su-l:session): session closed for user root
May 18 15:12:59 centos7.localdomain sudo[5228]:  vagrant : TTY=pts/0 ; PWD=/home/vagrant ; USER=root ; COMMAND=/bin/logi
May 18 15:12:59 centos7.localdomain sudo[5228]: pam_unix(sudo:session): session opened for user root by vagrant(uid=0)
May 18 15:12:59 centos7.localdomain sudo[5228]: pam_unix(sudo:session): session closed for user root
May 18 15:13:56 centos7.localdomain su[5232]: pam_unix(su-l:auth): authentication failure; logname=vagrant uid=1000 euid
May 18 15:13:56 centos7.localdomain su[5232]: pam_succeed_if(su-l:auth): requirement "uid >= 1000" not met by user "root
May 18 15:13:58 centos7.localdomain su[5232]: FAILED SU (to root) vagrant on pts/0
May 18 15:14:08 centos7.localdomain su[5236]: (to root) vagrant on pts/0
May 18 15:14:08 centos7.localdomain su[5236]: pam_unix(su-l:session): session opened for user root by vagrant(uid=1000)
```

- show-user [USER...] 显示用户的属性

```shell
[root@centos7 workspace]# loginctl show-user vagrant
UID=1000
GID=1000
Name=vagrant
Timestamp=Thu 2023-05-18 10:13:33 CST
TimestampMonotonic=3225226215
RuntimePath=/run/user/1000
Slice=user-1000.slice
Display=9
State=active
Sessions=46 9
IdleHint=no
IdleSinceHint=0
IdleSinceHintMonotonic=0
Linger=no
```

- enable-linger [USER...] 启用一个或多个用户的延迟状态
- disable-linger [USER...] 禁用一个或多个用户的延迟状态
- terminate-user USER... 终止退出一个或多个用户的会话
- kill-user USER... 终止退出用户的会话

#### Seat

- list-seats 列出本机上的所有可用席位

```shell
[root@centos7 ~]# loginctl list-seats
SEAT
seat0

1 seats listed.
```

- seat-status [NAME...] 显示可用席位的状态

```shell
[root@centos7 ~]# loginctl seat-status seat0
seat0
  Devices:
          ├─/sys/devices/LNXSYSTM:00/LNXPWRBN:00/input/input0
          │ input:input0 "Power Button"
          ├─/sys/devices/LNXSYSTM:00/LNXSLPBN:00/input/input1
          │ input:input1 "Sleep Button"
          ├─/sys/devices/LNXSYSTM:00/device:00/PNP0A03:00/LNXVIDEO:00/input/input4
          │ input:input4 "Video Bus"
          ├─/sys/devices/pci0000:00/0000:00:02.0/drm/card0
          │ drm:card0
          ├─/sys/devices/pci0000:00/0000:00:02.0/graphics/fb0
          │ [MASTER] graphics:fb0 "vboxdrmfb"
          ├─/sys/devices/pci0000:00/0000:00:04.0/input/input6
          │ input:input6 "VirtualBox mouse integration"
          ├─/sys/devices/platform/i8042/serio0/input/input2
          │ input:input2 "AT Translated Set 2 keyboard"
          ├─/sys/devices/platform/i8042/serio1/input/input3
          │ input:input3 "ImExPS/2 Generic Explorer Mouse"
          └─/sys/devices/platform/pcspkr/input/input5
            input:input5 "PC Speaker"
```

- show-seat [NAME...] 显示席位的属性

```shell
[root@centos7 ~]# loginctl show-seat seat0
Id=seat0
CanMultiSession=yes
CanTTY=yes
CanGraphical=yes
Sessions=
IdleHint=yes
IdleSinceHint=0
IdleSinceHintMonotonic=0
```

- attach NAME DEVICE... 将指定的设备连接到指定的席位上
- flush-devices 删除所有之前用 attach 命令连接的设备(同时也删除了所有之前用 attach 命令创建的席位)
- terminate-seat NAME... 结束指定席位上的所有会话, 将杀死指定席位上的所有会话进程, 同时释放与之关联的所有资源

### journalctl

管理系统运行日志

- \-S,\-\-since=DATE 显示指定日期开始之后的日志
- \-U,\-\-until=DATE 显示指定日期开始之前的日志
- \-b,\-\-boot[=ID] 显示指定 boot 的日志
- \-\-list-boots 显示所有的 boot

```shell
[root@centos7 ~]# journalctl --list-boots
0 312745f18eaa4b3eb809d0f361ad43bc 四 2023-05-18 09:19:48 CST—四 2023-05-18 16:01:01 CST
```

- \-k,\-\-dmesg 显示本次启动时的日志
- \-u,\-\-unit=UNIT 显示指定 Unit 的日志

```shell
[root@centos7 ~]# journalctl --since "2 hours ago" -u redis.service
-- Logs begin at 四 2023-05-18 16:09:37 CST, end at 四 2023-05-18 16:22:59 CST. --
5月 18 16:20:54 centos7 systemd[1]: Stopping redis-server...
5月 18 16:20:54 centos7 systemd[1]: Stopped redis-server.
5月 18 16:22:59 centos7 systemd[1]: Starting redis-server...
5月 18 16:22:59 centos7 systemd[1]: Started redis-server.
```

- \-e,\-\-pager-end 立刻跳到页面结尾
- \-f,\-\-follow 阻塞 journalctl 进程监听系统运行日志
- \-n,\-\-lines[=INTEGER] 控制显示日志的行数
- \-r,\-\-reverse 倒叙输出日志
- \-0,\-\-output=STRING 设置输出日志的格式
  - short
  - short-iso
  - short-precise
  - short-monotonic
  - verbose
  - export
  - json
  - json-pretty
  - json-sse
  - cat
- \-\-verify 校验日志的一致性
- \-\-header 显示 journal 的头部信息

```shell
[root@centos7 ~]# journalctl --header
File Path: /run/log/journal/afcca427b44c4f139ef788ed3b33b7e1/system.journal
File ID: 498fe2e44a684bf7a8353f3e9d09b4f4
Machine ID: afcca427b44c4f139ef788ed3b33b7e1
Boot ID: 33002b9c227345c98ada60a3f1a386c5
Sequential Number ID: 498fe2e44a684bf7a8353f3e9d09b4f4
State: ONLINE
Compatible Flags:
Incompatible Flags: COMPRESSED-XZ
Header size: 240
Arena size: 8388368
Data Hash Table Size: 20913
Field Hash Table Size: 333
Rotate Suggested: no
Head Sequential Number: 1
Tail Sequential Number: 977
Head Realtime Timestamp: 四 2023-05-18 16:09:37 CST
Tail Realtime Timestamp: 四 2023-05-18 16:30:01 CST
Tail Monotonic Timestamp: 20min 23.207s
Objects: 4090
Entry Objects: 977
Data Objects: 2241
Data Hash Table Fill: 10.7%
Field Objects: 43
Field Hash Table Fill: 12.9%
Tag Objects: 0
Entry Array Objects: 827
Disk usage: 8.0M
```

### systemd-analyze

查看当前系统的启动耗时

- time 有无此命令都可以

```shell
[root@centos7 workspace]# systemd-analyze time
Startup finished in 492ms (kernel) + 3.958s (initrd) + 38.325s (userspace) = 42.776s
```

- blame 查看每个服务的启动耗时
- critical-chain 显示瀑布状的启动过程流

### systemd-cgls

递归显示 cgroup 内容

```shell
[root@centos7 ~]# systemd-cgls
├─1 /usr/lib/systemd/systemd --switched-root --system --deserialize 22
├─user.slice
│ └─user-1000.slice
│   ├─session-51.scope
│   │ ├─6018 sshd: vagrant [priv]
│   │ ├─6021 sshd: vagrant@pts/1
│   │ ├─6022 -bash
│   │ ├─6255 su - root
│   │ └─6259 -bash
│   └─session-50.scope
│     ├─5989 sshd: vagrant [priv]
│     ├─5992 sshd: vagrant@pts/0
│     ├─5993 -bash
│     ├─6047 su - root
│     ├─6051 -bash
│     ├─6557 systemd-cgls
│     └─6558 less
└─system.slice
  ├─vboxadd-service.service
  │ └─3167 /usr/sbin/VBoxService --pidfile /var/run/vboxadd-service.sh
  ├─docker.service
  │ └─1182 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
  ├─postfix.service
  │ ├─1123 /usr/libexec/postfix/master -w
  │ ├─1126 qmgr -l -t unix -u
  │ └─4421 pickup -l -t unix -u
  ├─redis.service
  │ └─1057 /usr/local/bin/redis-server 127.0.0.1:6379
  ├─sshd.service
  │ └─1028 /usr/sbin/sshd -D -u0
  ...
```

### systemd-cgtop

显示 cgroup 的资源使用情况, 类似与 top 命令

```shell
Path                                                                             Tasks   %CPU   Memory  Input/s Output/s
/                                                                                   77    2.3   437.4M        -        -
/user.slice                                                                         11    1.8    35.1M        -        -
/system.slice                                                                        -    0.2   373.4M        -        -
/system.slice/redis.service                                                          1    0.2     9.2M        -        -
/system.slice/containerd.service                                                     1    0.0    58.5M        -        -
/system.slice/rsyslog.service                                                        1    0.0     2.5M        -        -
/system.slice/tuned.service                                                          1    0.0    13.2M        -        -
/system.slice/vboxadd-service.service                                                1    0.0     1.4M        -        -
/system.slice/haveged.service                                                        1    0.0     5.8M        -        -
/system.slice/NetworkManager.service                                                 2      -    11.9M        -        -
/system.slice/auditd.service                                                         1      -     2.9M        -        -
/system.slice/chronyd.service                                                        1      -     1.1M        -        -
/system.slice/crond.service                                                          1      -   756.0K        -        -
/system.slice/dbus.service                                                           1      -     1.7M        -        -
/system.slice/docker.service                                                         1      -   123.8M        -        -
/system.slice/firewalld.service                                                      1      -    34.0M        -        -
/system.slice/sshd.service                                                           1      -     5.1M        -        -
/system.slice/sys-kernel-debug.mount                                                 -      -   256.0K        -        -
/system.slice/system-getty.slice                                                     1      -   188.0K        -        -
/system.slice/system-getty.slice/getty@tty1.service                                  1      -        -        -        -
/system.slice/system-lvm2\x2dpvscan.slice                                            -      -     4.0K        -        -
/system.slice/systemd-journald.service                                               1      -     1.2M        -        -
/system.slice/systemd-logind.service                                                 1      -   960.0K        -        -
/system.slice/systemd-udevd.service                                                  1      -    20.1M        -        -
/system.slice/vagrant_data.mount                                                     -      -    16.0K        -        -
/user.slice/user-1000.slice/session-50.scope                                         6      -        -        -        -
/user.slice/user-1000.slice/session-51.scope                                         5      -        -        -        -
```

### systemd-nspawn

生成一个用于调试、测试和构建的最小命名空间容器

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
