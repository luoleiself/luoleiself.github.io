---
title: nginx.md
date: 2022-04-15 11:15:45
categories:
  - server
tags:
  - nginx
---

## 敲黑板

### 启动 nginx 失败

命令行提示错误 98: address already in use

查找系统进程中已存在的 nginx 进程号, 使用 `kill -9 $PID` 关闭进程后重启 nginx 服务

## 变量

- $nginx_version nginx 版本

- $connection_requests TCP 链接当前的请求数量

- $proxy_protocol_addr 获取代理访问服务器的客户端地址，如果是直接访问，该值为空字符串
<!-- more -->

### $request

- $content_length 请求头字段 `Content-Length`
- $content_type 请求头字段 `Content-Type`

- $document_uri 同 $uri
- $document_root 当前请求的文档根目录或别名
- $realpath_root 当前请求的文档根目录或别名的真实路径，会将所有符号连接转换为真实路径

- $remote_addr 客户端的 IP 地址
- $remote_port 客户端的端口号
- $remote_user 用于 HTTP 基础认证服务的用户名

- $host 请求信息中的 `Host`, 如果请求中没有 Host 行，则等于设置的服务器名;
- $hostname 主机名

- $http_user_agent 客户端 agent 信息
- $http_cookie 客户端所有的 cookie 信息
- $cookie_NAME 获取指定 cookie, 后面的 `NAME` 为 cookie 的 key

- $http_referer 引用地址

- $http_via 最后一个访问服务器的 Ip 地址

- $http_x_forwarded_for 相当于网络访问路径

#### $request_uri

- $request 代表客户端的请求地址
- $request_uri 包含请求参数的原始 URI, 不包含主机名
- $request_body 客户端的请求主体
- $request_method HTTP 请求方法，一般为 `GET` 或 `POST`
- $request_body_file 将客户端请求主体保存在临时文件中. 文件处理结束后, 此文件需删除
- $request_filename 当前连接请求的文件路径, 由 root 或 alias 指令与 URI 请求生成
- $request_length 请求的长度 (包括请求的地址, http 请求头和请求主体)
- $request_time 处理客户端请求使用的时间; 从读取客户端的第一个字节开始计时

- $uri 当前请求的 URI

- $arg_name 请求中的的参数名, 即"?"后面的 arg_name=arg_value 形式的 arg_name
- $args 请求中的参数值
- $is_args 如果请求中有参数, 值为 "?", 否则为空字符串
- $query_string 同$args

### $server

- $time_local 服务器时间(LOG Format 格式)

- $sent_http_NAME 可以设置任意 http 响应头字段, 变量名中的后半部分 NAME 能够替换成任意响应头字段, 连字符用下划线代替

  - $sent_http_content_type 'text/html'
  - $sent_http_content_length 1024

#### $server_uri

- $scheme HTTP 方法(如 http, https)
- $server_protocol 请求使用的协议, 通常是 HTTP/1.0 或 HTTP/1.1
- $server_addr 服务器地址
- $server_name 服务器名称
- $server_port 服务器端口号

- $status HTTP 响应代码

## 指令

### add_header 添加响应头字段

```nginx
add_header name value [always]; # 基础语法
```

#### 作用场景

- http
- server
- location
- if in location

#### 如果当前层添加了 `add_header`, 则不能从上层继承

#### 仅当状态码为 `ngx_http_headers_module` 模块列出时, `add_header` 添加的标头字段有效

200、201 (1.3.10)、204、206、301、302、303、304、307 (1.1.16、1.0.13) 或 308

#### 如果设置一个非有效状态码, 则会忽略 `add_header` 添加的标头字段

#### always 忽略状态码强制添加标头字段

### rewrite 重写请求中的 URI

至少有两个参数, 第一个参数匹配 URI 的正则表达式, 第二个参数是替换匹配的 URI, 第三个参数是标志位, 可以停止处理进一步的 rewrite 指令或发送重定向状态码(301|302)

- last 停止执行当前 server 上下文中的指令, 会继续搜索新 URI 匹配的位置
- break 停止执行当前 server 上下文中的指令, 取消搜索新 URI 匹配的位置, 不执行新位置中的 rewrite 指令

```nginx
rewrite ^(/download/.*)/media/(\w+)\.?.*$ $1/mp3/$2.mp3 last;
rewrite ^/users/(.*)$ /show?user=$1 break;
```

### try_files 尝试检查文件

```nginx
# 如果源文件不存在则内部重定向最后一个参数指定的 URI, 返回 /www/data/images/default.gif
location /images/ {
  root /www/data;

  try_files $uri /images/default.gif;
}

# 如果文件或者目录不存在则返回404
location / {
  try_files $uri $uri/ $uri.html =404;
}

# 如果文件或者目录不存在, 则请求重定向到指定位置然后传递给代理服务器
location / {
  try_files $uri $uri/ @backend;
}
location @backend {
  proxy_pass http://backend.example.com;
}
```

### [location 路径匹配](http://nginx.org/en/docs/http/ngx_http_core_module.html#location)

#### 匹配规则

| 符号 | 说明                                                                 |
| :--: | -------------------------------------------------------------------- |
|  ~   | 正则匹配，区分大小写                                                 |
| ~\*  | 正则匹配，不区分大小写                                               |
|  ^~  | 普通字符匹配，如果该选项匹配，则，只匹配改选项，不再向下匹配其他选项 |
|  =   | 普通字符匹配，精确匹配                                               |
|  @   | 定义一个命名的 location，用于内部定向，例如 error_page，try_files    |

#### 匹配优先级( 跟 location 的书写顺序关系不大 )

1. 精确匹配：

   = 前缀的指令严格匹配这个查询。

   如果找到，停止搜索。

2. 普通字符匹配：

   所有剩下的常规字符串，最长的匹配。

   如果这个匹配使用^〜前缀，搜索停止。

3. 正则匹配：

   正则表达式，在配置文件中定义的顺序，匹配到一个结果，搜索停止；

4. 默认匹配：

   如果第 3 条规则产生匹配的话，结果被使用。

   否则，如同从第 2 条规则被使用。
