---
title: Docker-上
date: 2022-03-22 18:37:24
categories:
  - [tools]
  - [linux, Docker]
tags:
  - Docker
---

## 敲黑板

- Package docker-ce is not available, but is referred to by another package.

  如果提示未发现可用的 docker-ce 包时,检查系统镜像源是否正确(如果不能翻墙时, 使用国内的镜像源修改 /etc/apt/source.list)

- Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get <http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json>: dial unix /var/run/docker.sock: connect: permission denied

  执行 docker 相关命令时提示, 表示 docker 权限不足

  - 使用 sudo 命令运行 docker 命令

  - 将当前用户加入到 docker 组中

```bash
[root@localhost ~]# cat /etc/group | grep docker
docker:x:994:vagrant
[root@localhost ~]# cat /etc/gshadow | grep docker
docker:!::vagrant
```

```bash
groupadd docker # 添加 docker 用户组
gpasswd -a $USER docker # 添加登陆用户到 docker 用户组中
newgrp docker # 更新用户组
或者
usermod -aG docker $USER # 给用户添加一个新的附属组
newgrp docker # 重新登陆组

systemctl restart docker # 重启 docker 服务

systemctl enable docker # 设置 docker 守护进程开机启动
```

<!-- more -->

### 概念

#### 架构

![docker-8](/images/docker-8.webp)

#### 命令

![docker-1](/images/docker-1.jpg)

## 镜像

是一个特殊的文件系统，除了提供容器运行时所需的程序、库、资源、配置等文件外，还包含了一些为运行时准备的一些配置参数（如匿名卷、环境变量、用户等）

镜像由多个层组成, 每层叠加之后, 从外部看来就如一个独立的对象

### UnionFS 联合文件系统

### 镜像加速

`/etc/docker/daemon.json` 配置文件修改默认镜像源

```json
{
  "registry-mirrors": []
}
```

### 镜像操作

镜像的操作命令可直接用在 `docker` 或 `docker image` 命令后面

- images 查看本地镜像列表, 作用同 `image ls`
- history 查看镜像的历史信息
- rmi 删除本地镜像, 作用同 `image rm`

```bash
docker rmi -f $(docker images -aq)
docker image rm -f $(docker image ls -aq) # 功能同上
```

- search [OPTIONS] TERM 从镜像仓库查找镜像

- load|import 从归档文件导入镜像
- save [OPTIONS] IMAGE [IMAGE...] 保存多个镜像到归档文件

```bash
[root@localhost ~]# docker [image] save -o bak.tar centos01:v1 centos02:v1 # 归档一个或多个镜像文件
```

- commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]] 基于容器创建一个镜像
- build 从 [Dockerfile 构建镜像](#build-image)

- pull [OPTIONS] NAME[:TAG|@DIGEST] 从远程仓库拉取镜像
- push [OPTIONS] NAME[:TAG] 将镜像上传到镜像仓库

- tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG] 创建一个指向 源镜像 的 目标镜像标签

```bash
# 创建一个指向 源镜像 的 目标镜像标签
[root@localhost ~]# docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]
```

![docker-3](/images/docker-3.png)

### 从容器构建镜像

```bash
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]] # 从容器构建镜像, 提交到本地仓库
```

- -a, \-\-author 提交作者
- -m, \-\-message 提交信息
- -p, \-\-pause 提交过程中是否中断容器运行, 默认为 true

```bash
[root@localhost ~]# docker images  # 显示所有镜像
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
nginx        latest    12766a6745ee   2 days ago     142MB
centos       latest    5d0da3dc9764   6 months ago   231MB
[root@localhost ~]# docker run -d -p 6666:80 nginx # 后台运行nginx并配置端口映射
[root@localhost ~]# curl -I localhost:6666 # 测试服务
HTTP/1.1 200 OK
Server: nginx/1.21.6
Date: Fri, 01 Apr 2022 08:44:03 GMT
Content-Type: text/html
Content-Length: 615
Last-Modified: Tue, 25 Jan 2022 15:03:52 GMT
Connection: keep-alive
ETag: "61f01158-267"
Accept-Ranges: bytes

[root@localhost ~]# docker ps -a # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                                   NAMES
91b34829d2d8   nginx     "/docker-entrypoint.…"   38 seconds ago   Up 36 seconds   0.0.0.0:6666->80/tcp, :::6666->80/tcp   laughing_kowalevski
a441e0564165   centos    "/bin/bash"              2 days ago       Up 2 hours                                              vigorous_turing
# 提交指定容器的镜像并添加作者,提交信息,版本号等
[root@localhost ~]# docker commit -a 'l.l' -m 'nginx 01' 91b34829d2d8 nginx01:1.0
[root@localhost ~]# docker images # 查看所有镜像
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
nginx01      1.0       9e81333043cc   2 seconds ago   142MB
nginx        latest    12766a6745ee   2 days ago      142MB
centos       latest    5d0da3dc9764   6 months ago    231MB
```

### 从 Dockerfile 构建镜像 <em id='build-image'></em> <!--markdownlint-disable-line-->

`.` 上下文路径

- -f 指定配置文件, 默认 `${PWD}/Dockerfile`
- -t 指定新创建的镜像名称和标签
- \-\-no-cache 构建镜像时不使用缓存
- \-\-compress 使用 gzip 压缩构建上下文环境
- \-\-label 设置镜像元数据
- \-\-network 设置构建过程中 `RUN` 指令的网络模式
- \-\-build-arg 通过命令行设置构建镜像过程中的参数, 可以覆盖 `ARG` 设置的参数

```bash
docker build -f /path/to/Dockerfile -t name:tag .
```

## 容器

容器独立运行的一个或一组应用，是镜像运行时的实体, 它可以被启动、开始、停止、删除. 每个容器之间都是相互隔离的, 保证安全的平台

- 应用容器化 将应用整合到容器中并且运行起来的这个过程

```bash
[root@localhost ~]# docker ps -a # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED        STATUS                   PORTS     NAMES
a441e0564165   centos    "/bin/bash"   28 hours ago   Exited (0) 3 hours ago             vigorous_turing
```

### 容器操作

容器的操作命令可直接用在 `docker` 或 `docker container` 命令后面

- create 创建新容器
- run [创建并启动一个容器](#run-container)
- start|stop|kill 启动|停止容器
- pause 暂停容器
- restart 重启容器
- rename 重命名容器
- events 获取 docker 服务器的实时事件
- diff 显示容器文件系统的前后变化
- ps 查看容器列表, 作用同 `container ls`
- inspect 查看容器详细信息, 作用同 `container inspect`
- logs 输出容器运行日志

  - -f, \-\-follow 实时输出容器运行日志
  - -n, \-\-tail 查看指定行数
  - -t, \-\-timestamps 输出日志添加时间戳

- port 查看容器映射端口
- wait 阻塞一个或多个容器直到停止运行, 并打印容器的退出码

```bash
[root@localhost ~]# docker wait centos01 # 阻塞一个或多个容器直到停止运行, 并打印容器的退出码
0
```

- export [OPTIONS] CONTAINER 导出容器的文件系统到归档文件

```bash
[root@localhost ~]# docker [image] export -o bak.tar centos01  # 导出容器的文件系统到归档文件
```

- rm 删除容器

```bash
[root@localhost ~]# docker rm -f $(docker ps -aq) # 删除所有容器
```

- stats 统计容器使用情况

```bash
[root@localhost ~]# docker stats
CONTAINER ID   NAME              CPU %     MEM USAGE / LIMIT     MEM %     NET I/O      BLOCK I/O     PIDS
a441e0564165   vigorous_turing   0.00%     1.336MiB / 481.6MiB   0.28%     1.6kB / 0B   6.79MB / 0B   1
```

### 运行 <em id='run-container'></em> <!--markdownlint-disable-line-->

如果本地不存在镜像时则先从远程拉取镜像(docker pull 镜像名)

```bash
docker run --name 'helloWorld' -it 镜像名 在启动的容器里执行的命令
```

- \-\-name 容器名称
- -d, \-\-detach 后台方式运行
- -a, \-\-attach 附加到标准输出中
- -i, \-\-interactive 即使没有附加也保持 STDIN 打开, 如果需要执行命令则需要开启这个选项
- -t, \-\-tty 分配一个伪终端进行执行, 一个连接用户的终端与容器 stdin 和 stdout 的桥梁
- -P, \-\-publish-all 将宿主机的随机端口映射到容器使用的端口上
- -p, \-\-publish 指定容器的端口

  - -p 宿主机 IP:宿主机端口:容器端口/协议
  - -p 宿主机端口:容器端口/协议
  - -p 容器端口/协议

- -e, \-\-env 设置容器运行的环境变量
- \-\-env-file 设置容器运行的环境变量文件
- -w, \-\-workdir 设置容器内部的工作目录
- -h, \-\-hostname 设置容器的主机名
- -v, \-\-volumes 设置容器数据卷映射
- \-\-mount 挂载文件系统
- \-\-tmpfs 挂载临时文件系统
- \-\-volumes-from 指定继承的数据卷容器

- -l, \-\-label 设置容器的元数据
- \-\-ip 设置容器 IP 地址
- \-\-ip6 IPv6 地址
- \-\-link 连接到另一个容器
- \-\-network 连接到指定的网络, 默认为 docker0

- \-\-privileged 授予此容器扩展权限
- \-\-entrypoint 覆盖镜像文件中默认的 ENTRYPOINT
- \-\-read-only 只读模式挂在容器文件系统
- \-\-restart string 当容器退出后的重启策略
  - no, 默认值
  - on-failure[:max-retries]
  - always
  - unless-stopped

- \-\-add-host list 添加主机 ip 映射
- \-\-dns list 设置 dns 服务
- \-\-cpus 指定容器运行时的 cpu 数量
- -m, \-\-memory 指定容器运行时的内存大小

- \-\-rm 测试时临时运行容器关闭后自动删除容器

```bash
docker run -d --rm -p 6666:80 --name nginx01 nginx
```

```bash
[root@localhost ~]# docker images # 查看本地所有镜像
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
nginx        latest    12766a6745ee   44 hours ago   142MB
centos       latest    5d0da3dc9764   6 months ago   231MB
[root@localhost ~]# docker ps -a  # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED        STATUS          PORTS     NAMES
a441e0564165   centos    "/bin/bash"   29 hours ago   Up 27 minutes             vigorous_turing
[root@localhost ~]# docker run -d --rm -p 6666:80 nginx # 后台运行容器同时配置端口映射
[root@localhost ~]# curl -I localhost:6666 # 测试 nginx 服务
HTTP/1.1 200 OK
Server: nginx/1.21.6
Date: Thu, 31 Mar 2022 11:52:34 GMT
Content-Type: text/html
Content-Length: 615
Last-Modified: Tue, 25 Jan 2022 15:03:52 GMT
Connection: keep-alive
ETag: "61f01158-267"
Accept-Ranges: bytes

[root@localhost ~]# docker ps -a  # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS          PORTS                                   NAMES
2ccf4c0204ab   nginx     "/docker-entrypoint.…"   11 seconds ago   Up 10 seconds   0.0.0.0:6666->80/tcp, :::6666->80/tcp   thirsty_ishizaka
a441e0564165   centos    "/bin/bash"              29 hours ago     Up 28 minutes                                           vigorous_turing
[root@localhost ~]# docker stop 2ccf4c0204ab  # 停止指定容器
2ccf4c0204ab
[root@localhost ~]# docker ps -a  # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED        STATUS          PORTS     NAMES
a441e0564165   centos    "/bin/bash"   29 hours ago   Up 28 minutes             vigorous_turing
```

- 示例：最简单的 nginx 服务集群

```bash
[root@localhost ~]# docker images # 查看所有镜像
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
nginx        latest    12766a6745ee   3 days ago     142MB
centos       latest    5d0da3dc9764   6 months ago   231MB
[root@localhost ~]# docker run -d --name 'nginx01' -p 6666:80 nginx # 后台启动 nginx 服务配置端口映射
[root@localhost ~]# docker run -d --name 'nginx02' -p 7777:80 nginx # 后台启动 nginx 服务配置端口映射
[root@localhost ~]# docker ps -a  # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND                  CREATED          STATUS                    PORTS                                   NAMES
335120e70611   nginx     "/docker-entrypoint.…"   6 seconds ago    Up 5 seconds              0.0.0.0:7777->80/tcp, :::7777->80/tcp   nginx02
771fd45df0ef   nginx     "/docker-entrypoint.…"   47 seconds ago   Up 46 seconds             0.0.0.0:6666->80/tcp, :::6666->80/tcp   nginx01
a441e0564165   centos    "/bin/bash"              3 days ago       Exited (0) 22 hours ago                                           vigorous_turing
[root@localhost ~]# curl -I localhost:6666  # 测试 nginx 服务
HTTP/1.1 200 OK
Server: nginx/1.21.6
Date: Sat, 02 Apr 2022 10:42:29 GMT
Content-Type: text/html
Content-Length: 615
Last-Modified: Tue, 25 Jan 2022 15:03:52 GMT
Connection: keep-alive
ETag: "61f01158-267"
Accept-Ranges: bytes

[root@localhost ~]# curl -I localhost:7777  # 测试 nginx 服务
HTTP/1.1 200 OK
Server: nginx/1.21.6
Date: Sat, 02 Apr 2022 10:42:37 GMT
Content-Type: text/html
Content-Length: 615
Last-Modified: Tue, 25 Jan 2022 15:03:52 GMT
Connection: keep-alive
ETag: "61f01158-267"
Accept-Ranges: bytes
```

### 进入容器

#### exec 和 attach

- exec
  - 进入容器打开一个新的终端
  - 退出容器, 容器正常运行
- attach
  - 进入容器打开正在运行的终端
  - 退出容器, 容器自动停止

```bash
[root@localhost ~]# docker ps -a # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED        STATUS          PORTS     NAMES
a441e0564165   centos    "/bin/bash"   25 hours ago   Up 31 minutes             vigorous_turing
[root@localhost ~]# docker exec -it a441e0564165 /bin/bash # 交互方式进入容器
[root@a441e0564165 /]# exit # 退出
exit
[root@localhost ~]# docker ps -a # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED        STATUS          PORTS     NAMES
a441e0564165   centos    "/bin/bash"   25 hours ago   Up 32 minutes             vigorous_turing

[root@localhost ~]# docker attach a441e0564165 # 进入容器
[root@a441e0564165 /]# exit # 退出
exit
[root@localhost ~]# docker ps -a # 查看所有容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED        STATUS                     PORTS     NAMES
a441e0564165   centos    "/bin/bash"   25 hours ago   Exited (0) 2 seconds ago             vigorous_turing
```

- -e 运行容器时指定环境变量
- -w 容器工作目录

```bash
[root@localhost ~]# docker exec -it -e PATH=/usr/local/v12.22.1/bin:$PATH centos01 node -v
v12.22.1
[root@localhost ~]# docker exec -it -w /usr/local centos01  pwd
/usr/local
```

### 文件拷贝 cp <em id="docker-cp"></em> <!--markdownlint-disable-line-->

- -a, \-\-archive 复制文档的所有信息

```bash
docker cp [宿主机路径] [容器标识]:[容器内路径]  # 拷贝宿主机文件到 容器内 目录

docker cp [容器标识]:[容器内路径] [宿主机路径]  # 拷贝容器内文件 到 宿主机目录
```

#### 拷贝宿主机到容器内

```bash
# 拷贝宿主机文件到 a441e0564165 /user/local 下
[root@localhost ~]# docker cp -a .nvm/versions/node/v12.22.1/ a441e0564165:/usr/local/v12.22.1/
[root@localhost ~]# docker exec -it a441e0564165 ls /usr/local
bin  etc  games  include  lib  lib64  libexec  sbin  share  src  v12.22.1
```

#### 拷贝容器内到宿主机

```bash
[root@localhost ~]# ls /home/
ubuntu  vagrant

[root@localhost ~]# docker exec -it a441e0564165 /bin/bash # 进入 a441e0564165
[root@a441e0564165 /]# ls -al /home
total 8
drwxr-xr-x 2 root root 4096 Nov  3  2020 .
drwxr-xr-x 1 root root 4096 Apr 10 08:06 ..
[root@a441e0564165 /]# touch hello docker > /home/hello.txt # 创建文件

[root@localhost ~]# docker cp -a a441e0564165:/home/hello.txt /home/vagrant
[root@localhost ~]# ls
centos01  description-pak  hello.txt
```

## 数据卷

卷就是目录或者文件,存在于一个或多个容器中,由 docker 挂载到容器中,卷的设计目的就是数据的持久化,完全独立于容器的生存周期, 因此 Docker 不会在容器删除时删除其挂载的数据

- create 创建数据卷
- inspect 显示数据卷的详细信息
- ls 显示所有数据卷
- rm 删除数据据按
- prune 移除所有未使用的本地数据卷

### 特点

- 卷中的更改可以直接生效
- 数据卷可在容器之间共享数据
- 数据卷中的更改不会包含在镜像的更新中
- 数据卷的生命周期一直持续到没有容器使用它为止

### 数据卷类型

默认挂载数据卷的权限为 `RW`

可以在挂载数据卷时指定数据卷的权限 `-v [source/path]:[destination/path]:[rw]`

- \-\-mount 不指定 type 选项默认为 volume
- -v 不能建立 tmpfs mounts

|       对比项       | --volume 或 -v | --mount type=bind |
| :----------------: | :------------: | :---------------: |
| 若是主机路径不存在 |    自动建立    |     命令报错      |

|    对比项    |     bind mount     |           volume           |
| :----------: | :----------------: | :------------------------: |
| Source 位置  |      用户指定      |  /var/lib/docker/volumes/  |
| Source 为空  |   覆盖 dest 为空   |       保留 dest 内容       |
| Source 非空  |   覆盖 dest 内容   |       覆盖 dest 内容       |
| Source 种类  |     文件或目录     |         只能是目录         |
|   可移植性   |  一般（自行维护）  |     强（docker 托管）      |
| 宿主直接访问 | 容易（仅需 chown） | 受限（需登陆 root 用户）\* |

#### volume mounts

指定 docker 挂载区域, Docker 管理宿主机文件系统的一部分(/var/lib/docker/volumes)

```bash
docker volume create [OPTIONS] [VOLUME] # 创建数据卷
docker run --mount type=volume,source=${PWD}/${CONTAINER_NAME}/app,destination=/app centos01 /bin/bash
docker run -v ${PWD}/${CONTAINER_NAME}/app:/app # 作用同上一行
```

#### bind mounts

是宿主机任意文件系统, 可以存储在宿主机系统的任意位置

```bash
docker run --mount type=bind,source=${PWD}/${CONTAINER_NAME}/app,destination=/app centos01 /bin/bash
# 如果挂载到容器中的非空目录, 则会隐藏容器中非空目录中的文件
docker run -v ${PWD}/${CONTAINER_NAME}/app:/app # 作用同上一行
```

#### tmpfs mounts

临时挂载到宿主机系统的内存中, 不会写入宿主机的文件系统

```bash
docker run --mount type=tmpfs,tmpfs-size=512M,destination=/path/in/container
```

![docker-5](/images/docker-5.jpg)

### 挂载数据卷 <em id="guazaishujujuan"></em> <!--markdownlint-disable-line-->

#### -v 挂载方式

##### -v 容器内路径 匿名挂载 <em id="nimingguazai"></em> <!--markdownlint-disable-line-->

```bash
[root@localhost ~]# docker run -tid --name centos01 -v /centosVolume centos /bin/bash
[root@localhost ~]# docker container inspect centos01
"Mounts": [
  {
    "Type": "volume",
    "Name": "2d25068e7e073ed051b075c3d454de4cd8db89871e6e19b38aa2c44b34cee647",
    "Source": "/var/lib/docker/volumes/2d25068e7e073ed051b075c3d454de4cd8db89871e6e19b38aa2c44b34cee647/_data",
    "Destination": "/centosVolume",
    "Driver": "local",
    "Mode": "",
    "RW": true,
    "Propagation": ""
  }
]
```

##### -v 卷名:容器内路径 具名挂载

- 如果 `不需要` 对容器内的数据卷挂载点进行 `写操作` 时, 使用具名挂载方式备份容器内数据到宿主机中

```bash
[root@localhost ~]# docker run -tid --name centos02 -v summary:/myVolume centos /bin/bash
[root@localhost ~]# docker container inspect centos02
"Mounts": [
  {
    "Type": "volume",
    "Name": "summary",
    "Source": "/var/lib/docker/volumes/summary/_data",
    "Destination": "/myVolume",
    "Driver": "local",
    "Mode": "z",
    "RW": true,
    "Propagation": ""
  }
]
```

##### -v /宿主机路径:容器内路径 指定路径挂载

- 如果 `需要` 对容器内的数据卷挂载点进行 `写操作` 时, 使用指定路径挂载方式, 此方式会将宿主机中的数据卷挂载点数据覆盖容器内指定路径

```bash
[root@localhost ~]# docker run -tid --name centos03 -v ${PWD}/react-app/:/containerVolume centos /bin/bash
[root@localhost ~]# docker container inspect centos03
"Mounts": [
  {
    "Type": "bind",
    "Source": "/home/workspace/react-app",
    "Destination": "/containerVolume",
    "Mode": "",
    "RW": true,
    "Propagation": "rprivate"
  }
]
```

#### \-\-mount 挂载方式

##### \-\-mount 具名挂载

```bash
[root@localhost ~]# docker run -tid --name centos04 --mount type=volume,source=applet_ui,destination=/appletVolume centos /bin/bash
[root@localhost ~]# docker container inspect centos04
"Mounts": [
  {
    "Type": "volume",
    "Name": "applet_ui",
    "Source": "/var/lib/docker/volumes/applet_ui/_data",
    "Destination": "/appletVolume",
    "Driver": "local",
    "Mode": "z",
    "RW": true,
    "Propagation": ""
  }
]
```

##### \-\-mount 指定路径挂载

```bash
[root@localhost ~]# docker run -tid --name centos05 --mount type=bind,source=${PWD}/applet_uni,destination=/uniVolume centos /bin/bash
[root@localhost ~]# docker container inspect centos05
"Mounts": [
  {
    "Type": "bind",
    "Source": "/home/workspace/applet_uni",
    "Destination": "/uniVolume",
    "Mode": "",
    "RW": true,
    "Propagation": "rprivate"
  }
]
```

### 共享数据卷

- \-\-volumes-from 创建数据卷容器共享数据

```bash
# 创建容器并挂载数据卷
[root@localhost ~]# docker run -tid --name centos01 -v /home/vagrant/centos01:/home/vagrant/centos01 centos /bin/bash
# 容器数据卷共享数据
[root@localhost ~]# docker run -tid --name centos02 --volumes-from centos01 centos /bin/bash
[root@localhost ~]# docker ps -a  # 查看容器信息
CONTAINER ID   IMAGE     COMMAND       CREATED          STATUS          PORTS     NAMES
7e2097a1476d   centos    "/bin/bash"   4 seconds ago    Up 2 seconds              centos02
ac564d66b6d4   centos    "/bin/bash"   16 minutes ago   Up 16 minutes             centos01
# 宿主机中创建文件
[root@localhost ~]# touch /home/vagrant/centos01/hello.txt
[root@localhost ~]# docker exec -it centos02 ls /home/vagrant/centos01 # 查看 centos02 文件状态
hello.txt
[root@localhost ~]# docker rm -f centos01 # 删除 centos01
[root@localhost ~]# docker ps -a
CONTAINER ID   IMAGE     COMMAND       CREATED              STATUS              PORTS     NAMES
7e2097a1476d   centos    "/bin/bash"   About a minute ago   Up About a minute             centos02
[root@localhost ~]# ls /home/vagrant/centos01/ # 查看宿主机文件状态
hello.txt
```

## Dockerfile

- .dockerignore 构建上下文忽略文件

Dockerfile 是用来构建 Docker 镜像的一个指令脚本, 脚本中的每条指令执行一次都会在镜像上新建一层

Docker 的多阶段构建特性允许在同一个 Dockerfile 中使用多个 FROM 指令来定义多个构建阶段.
每个阶段都可以从不同的基础镜像开始, 并且可以选择性地将文件从一个阶段复制到另一个阶段. 有助于创建更精简的最终镜像
docker build 时会为每个阶段构建一个临时的中间镜像, 但最终只会输出最后一个阶段的镜像

### 指令

- FROM 构建镜像时的基础镜像层, 可以出现多次以创建多个镜像或者将当前构建作为另一个构建的依赖
  - \-\-platform 指定镜像的平台, 用来处理那些支持多平台的镜像, linux/amd64、linux/arm64、windows/amd64
  - AS \<name\> 为当前的构建阶段命名
- MAINTAINER(deprecated) 维护者信息, 使用 `LABEL` 指令代替
- EXPOSE 对外暴露端口, 仅仅是声明容器使用的端口, 并不会自动在宿主机进行端口映射
  - 端口声明可以帮助镜像使用者理解这个镜像服务的端口使用, 以方便配置映射
  - 使用 -p 参数会忽略此项声明的端口
  - 使用 -P 参数会自动将宿主机上的随机端口映射到此项声明的端口

- ADD 复制指令, 增强版的 `COPY` 指令, 支持文件解压和远程 URL 资源
- COPY 复制指令, 从上下文目录中复制文件或者目录到容器里指定路径
- RUN 构建镜像时执行的命令 可以存在多条指令
  - RUN [OPTIONS] \<command\> ...
  - RUN [OPTIONS] ['\<command\>', ...]
- WORKDIR 为`RUN`, `CMD`, `ENTRYPOINT`, `COPY`, `ADD` 指定工作目录

- ARG 构建参数, 作用与 ENV 一致, ARG 中的环境变量仅在 `Dockerfile` 内有效
- ENV 设置持久化环境变量, 如果只想在构建构建阶段有效使用 `ARG` 指令

```yaml
# Dockerfile
# ARG BUILDPLATFORM linux/arm64
# FROM --platform=$BUILDPLATFORM nginx
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# EXPOSE 80/tcp
# EXPOSE 80/udp
# EXPOSE [8080, 8081]

# RUN yum -y install vim
# # 合并多个相同作用的指令以减少新建镜像层数
# RUN yum -y install wget \
#   && wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
#   && tar -xvf redis.tar.gz
# RUN --mount # 构建阶段挂载文件系统

# COPY [--chown=<user>:<group>] 可选参数，用户改变复制到容器内文件的拥有者和属组
# COPY ["src", "dest"]
# COPY --link

ARG VERSION1 1
ARG VERSION=1

ENV NAME1 hello
ENV NAME2=hello

FROM alpine:latest
WORKDIR /root
# 镜像构建完成之后将自动删除临时生成的镜像
COPY --from=builder /app/main . # 从某个镜像或构建阶段或命名上下文中复制文件
# ADD --keep-git-dir=true # 保持 .git 目录, 默认会被排除
# ADD --checksum=<hash>  # 验证远程资源
CMD ["./main"]
```

- CMD 容器运行时执行的命令, 如果存在多个 `CMD` 指令, 仅最后一个生效
- ENTRYPOINT 容器运行时执行的命令, 参数不会被 `docker run` 的命令行参数覆盖, 如果存在多个 `ENTRYPOINT` 指令，仅最后一个生效

```yaml
ENTRYPOINT '<exec_cmd>' '<param1>'
ENTRYPOINT ["<executeable>","<param1>","<param2>",...]

CMD '<exec_cmd>' '<param1>'
CMD ["<可执行文件或命令>","<param1>","<param2>",...]
```

- VOLUME 定义匿名数据卷, 在启动容器时会自动挂载到 /var/lib/docker/volumes/
  - 不支持 具名挂载 和 指定路径挂载, 只能使用 [匿名挂载](#nimingguazai) 方式
  - compose.yaml 配置项支持任意挂载方式 [挂载数据卷](#guazaishujujuan)

```yaml
# 出于可移植和分享的考虑, 不支持 具名挂载 和 指定路径挂载 在 Dockerfile 中配置
# 只能使用 匿名挂载 配置方式
# 因为宿主机目录是依赖于特定宿主机的, 并不能够保证在所有的宿主机上都存在这样的目录
VOLUME ["<路径1>", "<路径2>"...]
VOLUME <路径> <路径>
```

- USER 指定执行后续命令的用户和用户组, 用户名和用户组必须提前存在
- LABEL 给镜像添加元数据
- SHELL 允许重写默认的 shell
- STOPSIGNAL 设置当容器退出时系统调用的指令
- ONBUILD 当前镜像作为其他镜像的基础镜像构建时触发

```yaml
USER <用户名>[:<用户组>]

ONBUILD ADD . /app/src
```

- HEALTHCHECK 指定监控 docker 容器服务的运行状态的方式

#### CMD 和 ENTRYPOINT

##### 区别

- CMD 情况下, run 后面的参数将作为整体替换 `CMD` 配置项中的命令

```yaml
# Dockerfile
FROM ubuntu
#...
CMD ['/bin/bash', 'cat', '/etc/hosts']
```

```bash
# /bin/bash ls -al 命令会整体替换 Dockerfile 中的 CMD 指令
docker run -it ubuntu /bin/bash ls -al

# 实际命令
/bin/bash ls -al
```

- ENTRYPOINT 情况下, `docker run` 后面的参数将作为 ENTRYPOINT 配置项指令的一部分

```yaml
# Dockerfile
FROM ubuntu
#...
ENTRYPOINT ['/bin/bash', 'ls', '-l']
```

```bash
# run 后面的参数作为 ENTRYPOINT 配置项命令参数的一部分
docker run -it ubuntu -a

# 实际命令
/bin/bash ls -al
```

- CMD 和 ENTRYPOINT 同时存在时, CMD 作为 ENTRYPOINT 配置项命令的一部分, 与书写位置没有关系

```yaml
# Dockerfile
FROM ubuntu
#...
ENTRYPOINT ['nginx', '-c']
CMD ['/etc/nginx/nginx.conf']
```

|     是否传参     | Dockerfile 配置项指令          | 传参运行                             |
| :--------------: | ------------------------------ | ------------------------------------ |
|   Docker 命令    | docker run nginx               | docker run nginx /etc/nginx/new.conf |
| 衍生出的实际命令 | nginx -c /etc/nginx/nginx.conf | nginx -c /etc/new.conf               |

##### 使用场景

- Dockerfile 应至少指定一个 CMD 或 ENTRYPOINT 指令
- ENTRYPOINT 应在将容器作为可执行文件时定义
- CMD 应该用作为 ENTRYPOINT 命令定义默认参数或在容器中执行临时命令的一种方式
- CMD 使用替代参数运行容器时将被覆盖

|                         | No ENTRYPOINT         | ENTRYPOINT entry p1_entry | ENTRYPOINT ['entry','p1_entry']      |
| ----------------------- | --------------------- | ------------------------- | ------------------------------------ |
| No CMD                  | error, not allowed    | /bin/sh -c entry p1_entry | entry p1_entry                       |
| CMD ['cmd','p1_cmd']    | cmd p1_cmd            | /bin/sh -c entry p1_entry | entry p1_entry cmd p1_cmd            |
| CMD ['p1_cmd','p2_cmd'] | p1_cmd p2_cmd         | /bin/sh -c entry p1_entry | entry p1_entry p1_cmd p2_cmd         |
| CMD cmd p1_cmd          | /bin/sh -c cmd p1_cmd | /bin/sh -c entry p1_entry | entry p1_entry /bin/sh -c cmd p1_cmd |

### 镜像优化

- 合并多个相同作用的指令以减少新建镜像层数
- 使用较小的基础镜像
- 使用 .dockerignore 文件, 从构建上下文中忽略指定的文件
- 在 RUN 之后放置 COPY 指令, 有助于 docker 能够更好的使用缓存功能
- 在安装后删除软件包
- 使用 Docker 镜像缩容工具

### 应用

- Dockerfile 配置文件

```yaml
# Dockerfile
FROM centos:7 # 基础镜像

LABEL authors=luoleiself<luoleiself@163.com> # 镜像元数据

WORKDIR /usr/local/   # 设置工作目录

# ADD hello.txt $PWD  # 复制文件, 增强版的 COPY 指令
COPY hello.txt $PWD   # 复制文件

VOLUME ["/var/log"] # 挂载匿名数据卷

# ARG NAME hello  # 构建参数, 仅在 Dockerfile 内有效
# ENV NAME world  # 持久化环境变量

# RUN npm install

CMD ["cat", "/usr/local/hello.txt"]   # 容器运行时执行的命令
# ENTRYPOINT ["cat", "/usr/local/hello.txt"] # 容器运行执行的命令
```

#### 构建镜像

```bash
# 构建镜像并指定镜像名称和版本号, 最后的 . 很重要,表示在当前目录下进行构建, 可以使用
[root@localhost workspace]# docker build -t hello:v1.0 .
Sending build context to Docker daemon  3.072kB
Step 1/6 : FROM centos:7
 ---> eeb6ee3f44bd
Step 2/6 : LABEL maintainer=luoleiself<luoleiself@163.com>
 ---> Running in 42aee4c83238
Removing intermediate container 42aee4c83238
 ---> c0a8427ea1ca
Step 3/6 : WORKDIR /usr/local/
 ---> Running in 0b1c0da5bfdb                   # 启动临时容器
Removing intermediate container 0b1c0da5bfdb    # 修改元数据内部使用 `docker commit` 提交镜像后, 立即删除临时容器
 ---> 63cd399d1589
Step 4/6 : COPY hello.txt $PWD
 ---> 2e6fe81c553d
Step 5/6 : VOLUME ["/var/log"]                  # 挂载匿名数据卷
 ---> Running in 35e67913cfb2
Removing intermediate container 35e67913cfb2
 ---> 60419e54ac19
Step 6/6 : CMD ["cat", "/usr/local/hello.txt"]
 ---> Running in 1785744ac7aa
Removing intermediate container 1785744ac7aa
 ---> 4cd89659cce0
Successfully built 4cd89659cce0
Successfully tagged hello:v1.0

[root@localhost workspace]# docker run --name hello-v1 hello:v1.0 # 基于镜像创建容器
hello docker
```

#### 查看镜像信息

##### 镜像配置信息

```bash
[root@localhost workspace]# docker image inspect hello:v1.0
...
"Config": {
  "Env": [
    "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
  ],
  "Cmd": [
    "cat",
    "/usr/local/hello.txt"
  ],
  "Image": "sha256:60419e54ac191ef9170c15d90ab3e0a040ae75ab0c15b320918977336710f17e",
  "Volumes": {
    "/var/log": {}
  },
  "WorkingDir": "/usr/local/",
  "Entrypoint": null,
}
...
```

##### 镜像历史信息

```bash
[root@localhost workspace]# docker image history hello:v1.0
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
4cd89659cce0   25 minutes ago   /bin/sh -c #(nop)  CMD ["cat" "/usr/local/he…   0B
60419e54ac19   25 minutes ago   /bin/sh -c #(nop)  VOLUME [/var/log]            0B
2e6fe81c553d   25 minutes ago   /bin/sh -c #(nop) COPY file:f5a7cb03621adea9…   13B
63cd399d1589   25 minutes ago   /bin/sh -c #(nop) WORKDIR /usr/local/           0B
c0a8427ea1ca   25 minutes ago   /bin/sh -c #(nop)  LABEL maintainer=luoleise…   0B
eeb6ee3f44bd   7 months ago     /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B
<missing>      7 months ago     /bin/sh -c #(nop)  LABEL org.label-schema.sc…   0B
<missing>      7 months ago     /bin/sh -c #(nop) ADD file:b3ebbe8bd304723d4…   204MB
```

#### 查看容器状态

```bash
[root@localhost workspace]# docker container inspect hello-v1
...
"Mounts": [
  {
    "Type": "volume",
    "Name": "08785005cb382a42f41b8f7230564cf468d38be82ce74c94fb3a277ef5a64f66",
    "Source": "/var/lib/docker/volumes/08785005cb382a42f41b8f7230564cf468d38be82ce74c94fb3a277ef5a64f66/_data",
    "Destination": "/var/log",
    "Driver": "local",
    "Mode": "",
    "RW": true,
    "Propagation": ""
  }
]
...
```
