---
title: Docker-Compose
date: 2022-04-30 11:11:58
categories:
  - [tools]
  - [linux, Docker]
tags:
  - Docker
---
## Docker Compose

- Docker Compose 以服务为单位, 将为每一个服务部署一个容器
- 默认以 `应用名称-服务名称-数字` 方式作为容器名称
- 默认以 `应用名称_数据卷名` 方式作为数据卷名称
- 默认以 `应用名称_网络名` 方式作为网络名称

Docker Compose 是定义和运行多容器 Docker 应用程序的工具, 运行部分命令时需要在 `compose.yaml/docker-compose.yaml` 文件所在目录中, 以 `应用名称-服务名称-数字` 编号为规则命名容器, 配置文件使用 yaml 语法, yaml 是一个可读性高，用来表达数据序列化的格式.

yaml 文件中不能使用 tab 缩进, 只能使用空格

```yaml
${VAR:-default}   如果 变量名 对应的环境变量未设置或为空, 则使用 默认值, 否则使用环境变量的值 
${VAR-default}    仅当 变量名 对应的环境变量 完全未定义 时才用默认值(为空时不会替换)
${VAR:?error}     严格模式, 如果 变量名 对应的环境变量 未设置/为空, 则报错并显示错误信息 error
${VAR:+alternate} 反向逻辑, 如果 变量名 已设置且非空, 则使用 alternate 值
```

```bash
# 启动指定服务, 不加参数则默认启动所有服务
docker compose -f -p -c --env-file .env.development up [service_name]

# 以下的命令不带服务名称则默认对所有服务执行相同操作
```

### 参数

- \-\-all-resources 引入所有的资源, 即使未被服务使用
- -f, \-\-file stringArray 指定配置文件

- -p, \-\-project-name string 指定项目名称
- \-\-project-directory string 指定项目工作目录

- -c, \-\-context 指定上下文环境名称
- \-\-env-file stringArray 指定环境变量配置文件
- \-\-parallel 设置并行
- \-\-compatibility 运行 compose 兼容模式

- \-\-profile 启用指定的服务, web 服务默认启动

```bash
docker compose --profile db up -d # 只启动 db, web 服务
```

### 命令

- attach 连接运行服务的标准输入输出
- build 构建服务
- commit 从服务容器创建一个新的镜像
  - -a, \-\-author 作者
  - -m, \-\-message string 提交信息
  - -c, \-\-change list 应用 Dockerfile 指令创建镜像
  - \-\-index int 指定如果有多个副本的服务的容器
  - -p, \-\-pause 提交过程中是否中断容器运行, 默认为 true
- config 解析验证 compose.yaml 配置文件
  - \-\-environment 打印环境变量的插值
  - \-\-format string 格式化输出, 值可选 yaml(default) | json
  - \-\-images 输出镜像名称

  - \-\-hash  输出服务的配置 hash, 一个一行
  - \-\-profiles  输出指定的服务名, 一个一行
  
  - -o, \-\-output string 保存到指定文件中, 默认是标准输出流
  - -q, \-\-quiet 仅验证配置项, 不输出任何信息
  - \-\-services 输出服务名称
  - \-\-volumes  输出数据卷名称
- cp 在容器和本地文件系统之间拷贝文件
- events 接收一个来自容器的真实的事件
- export 导出容器文件系统为归档文件
- images 列出创建容器的镜像

- start 启动服务
- restart 重启服务
- wait 阻塞直到所有的服务容器停止
- stop 停止服务
- kill 强制停止容器
- pause 暂停服务
- unpasue 取消暂停服务
- rm 移除已经停止的容器, 默认情况下附加到容器上的匿名数据卷不会被移除
  - -f, \-\-force 不询问确认操作直接移除
  - -s, \-\-stop 在移除之前停止容器
  - -v, \-\-volumes 移除所有附加到容器上的匿名数据卷
- down 停止并移除容器, 网络
  - \-\-rmi string 移除服务使用的镜像
  - -t, \-\-timeout int 延迟关机的时长秒
  - -v, \-\-volumes 移除在 compose.yaml 中 volumes 顶层指令中声明的具名和匿名数据卷

- logs 查看服务输出日志
  - \-f, \-\-follow  监听日志输出
  - \-\-index int  指定哪个容器执行命令, 如果服务有多个副本时
  - \-n, \-\-tail  输出容器日志的最后几行
  - \-\-since  显示指定时间开始的日志
  - \-\-until  显示截止到指定时间的日志
- stats 查看容器资源的使用情况

- ls 列出正在运行的 compose 项目
- port 查看公共端口绑定信息
  - \-\-index int 指定如果有多个副本的服务的容器
  - \-\-protocol string  指定协议, tcp(default) | udp
- ps 查看所有容器
  - -a, \-\-all 列出所有容器
  - \-f, \-\-filter \<FILTER_TYPE\>=\<VALUE\> 根据指定条件过滤服务容器
  - \-\-format string 使用自定义模板格式化输出, table(default) | table TEMPLATE | json | TEMPLATE
  - \-\-services 显示服务名称
  - \-\-status stringArray 通过状态过滤服务, [paused | restarting | removing | running | dead | created | exited]
- pull 拉取服务镜像
- push 推送服务镜像

- top 显示运行的进程信息
- exec 在运行的容器中执行命令
  - -d, \-\-detach  后台运行命令
  - -e, \-\-env stringArray 设置环境变量
  - \-\-index int  指定哪个容器执行命令, 如果服务有多个副本时
  - -T, \-\-no-TTY  不分配伪 TTY, 默认每次执行命令时都分配 TTY
  - -w, \-\-workdir string 设置本次命令的工作目录
- run 在服务上运行一次性命令
  - \-\-build 启动容器之前构建镜像
  - -d, \-\-detach
  - -i, \-\-interactive 交互式运行
  - -e, \-\-env stringArray 设置环境变量
  - -l, \-\-label stringArray 添加或覆盖 label
  - -P, \-\-service-ports 运行命令所有服务的端口都能映射到宿主机
  - \-\-name  给容器定义名字
  - \-\-pull  运行之前拉取镜像
  - -p, \-\-publish stringArray 发布容器的端口到宿主机
  - \-\-rm 当容器退出时自动移除
  - -v, \-\-volume stringArray 挂载数据卷
  - -w, \-\-workdir string 设置容器内的工作目录

- scale 调整服务
  - \-\-no-deps 不启动关联的服务

- version 查看版本信息
- watch 监听文件系统更新时服务容器重构/重启的构建上下文

- create 为服务创建容器

  - \-\-build 启动容器之前构建镜像
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-force-recreate 即使配置项或镜像没有改变也要重新创建容器
  - \-\-no-recreate 如果容器存在则不创建新的容器
  - \-\-pull 创建之前拉取镜像
  - \-\-scale scale 调整服务实例数量, 并覆盖 `compose.yaml` 配置文件中的 scale 配置

- up 创建服务并启动容器

  - -d, \-\-detach 后台运行容器
  - \-\-attach 连接服务的输出
  - \-\-no-attach 不连接服务的输出
  - \-\-build 启动容器之前构建镜像
  - \-\-no-build 不构建镜像即使镜像不存在
  - \-\-no-start 创建服务之后不启动它
  - \-\-no-deps 不启动关联的服务
  - \-\-pull 启动之前拉取镜像
  - \-\-scale scale 调整服务实例数量, 覆盖 `compose.yaml` 配置文件中的 scale 配置
  - \-\-no-log-prefix 打印日志时不适用前缀
  - \-\-no-recreate 如果容器存在则不创建新的容器
  - -y 非交互式运行命令, 所有的提示都回答 yes

  - -t, \-\-timeout int 延迟关闭容器

```bash
[root@localhost ~]# docker compose up service_id # 启动指定服务

# 调整指定服务实例数量, 先去掉 compose.yaml/docker-compose.yaml 配置文件 service 指定的端口, 在单机中会出现端口占用问题
[root@localhost ~]# docker compose up --scale web=5 -d
```

### 配置文件

- 使用副本不能指定容器名称, Compose 自动使用 应用名称-服务名称-数字 形式命名容器

```yaml
# compose.yaml/docker-compose.yaml
version: 3.9   # 版本, obsolete(已过时)
name: myapp # 定义默认的项目名称, 将以环境变量 ${COMPOSE_PROJECT_NAME} 的方式公开
services:
  web:   # 服务名称
    annotations: # 容器声明, 可以是 arr 或 map
      com.example.foo: bar
      # - com.example.foo=bar
    attach: false # 设置为 false 时不会主动收集服务日志, 默认为 false, v2.20.0 以上支持
    build:
      context: './web'  # 指定构建 web 服务的镜像的上下文环境目录
      dockerfile: Dockerfile  # 指定构建镜像的配置文件名称
    command: ['bundle', 'exec', 'thin', '-p', '3000']  # 覆盖镜像配置文件(Dockerfile)中的CMD指令
    ports: # 端口映射
      - '5000:5000'
      - '0.0.0.0:80:80/tcp' # 指定 ip 地址和协议, 或修改 /etc/docker/daemon.json 配置项"ipv6":false
      - target: 80
        host_ip: 127.0.0.1
        published: 8080
        protocol: tcp
        mode: host
    privileged: true   # 配置容器目录权限
    read_only: true    # 开启容器文件系统只读模式
    restart: always    # 定义容器重启模式 "no" | always | on-failure | unless-stopped
    container_name: my-web  # 容器名称, 使用副本不能指定容器名称, Compose 自动使用 应用名称-服务名称-数字 形式命名容器
    env_file: .env # 环境变量配置文件
    environment:  # 设置容器内环境变量
      RACK_ENV: development
      SHOW: 'true'
      USER_INPUT:
    entrypoint:   # 覆盖镜像配置文件(Dockerfile)中的 ENTRYPOINT 指令
      - php
      - -d
      - vendor/bin/phpunit
    external_links:   # 将服务容器连接到 compose 应用管理以外的服务, 作用同 links
      - redis
      - database:mysql
      - database:postgresql
    extra_hosts:  # 添加主机 ip 映射关系到容器网络接口配置中(/etc/hosts)
      - 'somehost:162.242.195.82'
      - 'otherhost:50.31.209.229'
    volumes: # 挂载数据卷
      - type: volume
        source: db-data
        target: /data
        read_only: true
        volume:
          nocopy: true
          subpath: sub
        tempfs:
          size: 1024
          mode: 755
      - type: bind
        source: /home/workspace
        target: /home/workVolume
      - /home/workspace:/var/workspace  # 定义指定路径数据卷
    tmpfs:
      - /run  # 挂载容器内临时文件系统
      - /tmp
    volumes_from:     # 挂载共享数据卷
      - service_name
      - service_name:ro
      - container:container_name
      - container:container_name:rw
    network_mode: "host|none|service:[service name]" # 设置服务容器的网络模式
    networks:  # 自定义网络模式
      - my-web-network
    platform: linux/amd64 # 设置服务容器运行的目标平台
    # 当前服务启动的依赖优先于当前服务启动
    # 当前服务关闭优先于当前服务的依赖关闭
    depends_on:  # 服务启动依赖, 
      db:
        condition: service_healthy
        restart: true
      redis:
        condition: service_healthy
      # - db
      # - redis
    deploy:  # 部署
      # 外部客户端连接服务的方式
      # vip(Virtual IP) 为服务分配虚拟 IP, 客户端使用虚拟 IP 连接
      # dnsrr 平台配置 dns 条目, 使用服务名称查询 IP 地址列表连接
      endpoint_mode: vip
      labels:
        com.example.description: 'This label will appear on the web server' # 服务元数据
      mode: replicated # 服务运行模式, global | replicaated(default) | replicated-job | global-job
      replicas: 6 # 实例
      restart_policy: # 服务重启策略, 如果缺失, compose 会使用服务 restart 项
        condition: on-failure
        delay: 5s
        max_attempts: 3
        window: 120s
      rollback_config: # 服务回滚设置
      update_config: # 服务升级设置
        parallelism: 2
        delay: 10s
        order: stop-first
    dns:
      - 8.8.8.8  # 自定义网络的 DNS 服务器
    extends:
      file: common.yml   # 当前配置中扩展另一个服务
    labels:    # 添加容器元数据
      - 'com.example.description=Accounting webapp'
    develop: # 定义开发模式容器同步
      devices:
      dns:
      dns_opt:
      dns_search:
    healthcheck: # 服务健康检查
      test: ["CMD", "curl", "-f", "http://localhost"]
      interval: 1m30s
      timeout: 10s
      retries: 3
      start_period: 40s
      start_interval: 5s

  redis: # 服务名称 
    # 当同时使用 image 和 build 指令构建时, 将按照 pull_policy 的定义进行构建, 
    # 如果未定义 pull_policy 时, Compose 会尝试先拉取镜像, 如果在镜像仓库或缓存中找不到镜像时从源构建
    pull_policy: always   # Compose 总是从镜像仓库拉取镜像
                 never    # Compose 不会从镜像仓库拉取镜像而是依赖缓存中的镜像, 如果缓存中不存在则报告错误
                 missing  # 默认选项, Compose 仅当缓存中镜像不可用时从镜像仓库拉取
                 build    # Compose 构建镜像如果镜像已经存在则重新构建
    image: redis
    build:
      context: redis
      dockerfile: /redis.Dockerfile
    volumes:   # 挂载数据卷
      - /home/workspace   # 定义匿名数据卷
      - db-data:/var/lib/redis  # 挂载公共数据卷 db-data
    networks:   # 自定义网络模式
      - my-web-network
    links:    # 定义网络连接另一个服务的容器
      - db:mysql   # 可以直接使用 服务名, 或者使用 服务名:别名 方式
    scale: 6  # 设置容器数量, 如果 scale 和 deploy.replicas 同时存在则必须保持一致
    profiles: # 指定启动时的 profile
      - 'db'

  db:
    image: mysql
    ports:
      - '3306:3306'
    environment:
      MYSQL_ROOT_PASSWORD: 123456
      MYSQL_DATABASE: test
      MYSQL_USER: test
      MYSQL_PASSWORD: test123
    volumes:
      - db-data:/var/lib/mysql   # 挂载公共数据卷 db-data
    networks:
      - my-web-network
    profiles: # 指定启动时的 profile
      - 'debug'

# 跨服务共享数据卷定义到顶层指令 volumes
volumes:
  db-data:   # 声明卷名, compose 自动创建该卷名并会添加项目名前缀
  data:
    name: 'my-app-data'

networks:
  front-tier:
  back-tier:
  backend:
    driver: custom-driver
  my-web-network:   # 声明自定义网络模式, compose 自动创建该网络并会添加项目名前缀
    driver: bridge
    name: myapp-network # 定义网络名称, 在 docker network 列表中显示
    attachable: true  # 允许独立的容器连接到此网络
    enable_ipv6: true
    external: true  # 指定此网络的生命周期在应用程序的生命周期之外进行维护, Compose 不会尝试创建这些网络, 如果不存在则返回错误
configs: # 允许服务调整其行为而无须重新构建 docker 镜像
  http_config:
    file: ./httpd.conf
secrets: # 针对敏感数据的配置
  server-certificate:
    file: ./server.cert
  token:
    environment: 'OAUTH_TOKEN'
```

#### volumes

- 使用路径方式挂载数据卷

- 使用卷名方式挂载数据卷, 需要在 `一级配置项` 中声明, compose 会自动创建以项目名为前缀的卷名, 如果不需要卷名前缀, 则使用 `external: true` 指定卷名, 但是需要手动创建该卷名

### 多实例 Web 应用

```yaml
# compose.yaml
name: myapp
service:
  nginx:
    image: nginx:latest
    port:
      - '80:80'
    volumes:
      - ./nginx.conf:/etc/nginx/nginx.conf  # 自定义 nginx 配置文件
    depends_on:
      - web
    network:
      - webnet
  web:
    build:
      context: web
      dockerfile: ./my-web-app.Dockerfile
    deploy:
      replicas: 3 # 启动 3 个 web 实例
    environment:
      - ENV=production
    network:
      - webnet
networks:
  webnet:
    driver: bridge
```

nginx

```conf
# nginx.conf
http {
  upstream web_backend {
    server web:80;
    server web:80;
    server web:80;
  }
  server {
    listen 80;
    location / {
      proxy_pass http://web_backend;

      proxy_set_header Host $host;
      proxy_set_header X-Reap-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
  }
}
```

web 应用

```python
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello, World!'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
```

```Dockerfile
# my-web-app.Dockerfile
# 使用官方 Python 镜像作为基础镜像
FROM python:3.9-slim
# 设置工作目录
WORKDIR /app

# 复制依赖文件并安装依赖
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# 复制应用代码
COPY . .
EXPOSE 80
CMD ["python", "app.py"]
```

### docker compose 数据库

```yaml
name: myapp # ${COMPOSE_PROJECT_NAME} 以环境变量形式访问项目名称
services:
  redis:
    image: redis:7
#   使用副本不能指定容器名称, Compose 自动使用 应用名称-服务名称-数字 形式命名容器
#   container_name: redis-container
    ports:
      - '6379:6379'
    command: ['redis-server', '--appendonly yes', '--logfile /data/redis.log']
    volumes:
      - /var/lib/redis:/data
      - /var/lib/redis/redis.conf:/usr/local/etc/redis/redis.conf
    networks:
      - my-app-network
#   deploy:
#     replicas: 3
#       labels:
#         com.myapp.redis.description: 'This label will appear on the redis server'
  mysql:
    image: mysql:latest
    container_name: mysql-container
    ports:
      - '3306:3306'
    environment:
      MYSQL_ROOT_PASSWORD: 123456
      MYSQL_DATABASE: test
      MYSQL_USER: test
      MYSQL_PASSWORD: test123
    volumes:
      - /var/lib/mysql:/var/lib/mysql
    networks:
      - my-app-network
  mongodb:
#   mongodb 6.0 以上 docker 镜像不再包含 mongo shell 工具. 只包含 mongod 数据库服务器
#   手动下载 mongosh 工具, 或者使用 mongodb 6.0 之前的版本
    image: mongo:latest
    container_name: mongodb-container
    ports:
      - '27017:27017'
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: 123456
      MONGODB_BIND_IP: 0.0.0.0
    command: ['mongod', '--logpath', '/var/log/mongodb/mongod.log']
    volumes:
      - /var/lib/mongodb:/data/db
      - /var/lib/mongodb/logs:/var/log/mongodb
    networks:
      - my-app-network
#   挂载临时文件系统添加初始化脚本安装 mongosh
    tmpfs:
      - /tmp
#   首次创建容器需要先执行 entrypoint 命令安装 mongosh, 然后再注释 entrypoint
#   entrypoint: ['/bin/sh', '-c', 'apt-get update && apt-get install -y wget && wget -qO- https://downloads.mongodb.com/compass/mongosh-1.5.4-linux-x64.tgz | tar -xz -C /usr/local/bin --strip-components 1 mongosh-1.5.4-linux-x64/bin/mongosh && exec docker-entrypoint.sh $$MONGO_INITDB_ROOT_USERNAME $$MONGO_INITDB_ROOT_PASSWORD']
# sqllite:
#   image: sqllite3:latest
#   command: ['sqllite3', '/data/database.db'] # 启动 sqllite 并指向数据库文件
#   volumes:
#     - /var/lib/sqllite:/data
#   networks:
#     - my-app-network
volumes:
  db-data:
    labels:
      - 'com.myapp.volumes.description=share data vaolume'
networks:
  my-app-network:
    driver: bridge
    name: myapp-network # 定义网络名称, 在 docker network 列表中显示
    attachable: true  # 允许独立的容器连接到此网络
    enable_ipv6: true
```
