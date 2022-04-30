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

查找系统进程中已存在的 nginx 进程号, 使用 `kill -9 $PID` 关闭进程后重启

## 变量

- $uri 当前请求的 URI
- $arg_name 请求中的的参数名, 即"?"后面的 arg_name=arg_value 形式的 arg_name
- $args 请求中的参数值
- $is_args 如果请求中有参数, 值为 "?", 否则为空字符串
- $query_string 同$args

## 指令

### rewrite 重写请求中的 URI

至少有两个参数, 第一个参数匹配 URI 的正则表达式, 第二个参数是替换匹配的 URI, 第三个参数是标志位, 可以停止处理进一步的 rewrite 指令或发送重定向状态码(301|302)

- last 停止执行当前 server 上下文中的指令, 会继续搜索新 URI 匹配的位置
- break 停止执行当前 server 上下文中的指令, 取消搜索新 URI 匹配的位置, 不执行新位置中的 rewrite 指令

```conf
rewrite ^(/download/.*)/media/(\w+)\.?.*$ $1/mp3/$2.mp3 last;
rewrite ^/users/(.*)$ /show?user=$1 break;
```

### try_files 尝试检查文件

```conf
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

<!-- more -->

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
