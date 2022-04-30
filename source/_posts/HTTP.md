---
title: HTTP
date: 2021-07-26 19:12:29
categories:
  - network
tags:
  - http
---

#### 状态码

- 100 Continue 请求继续
- 101 Switching Protocols 协议切换，通过 upgrade 消息头切换协议
- 102 Processing 继续执行

- 200 OK 请求成功
- 204 No Content 没有响应实体
- 206 Partial Content 服务器处理部分 GET 请求

- 301 Moved Permanently 资源永久移动
- 302 Move Temporarily 资源临时移动
- 303 See Other 参见其他
- 304 Not Modified 资源没有改动
- 305 Use Proxy 使用代理
- 307 Temporary Redirect 资源临时从不同的 URI 响应

- 400 Bad Request 错误的请求
- 401 Unauthorized 需要验证
- 403 Forbidden 服务器拒绝执行
- 404 Not Found 资源未找到
- 405 Method Not Allowed 请求方法不能用于相应资源
- 408 Request Timeout 请求超时
- 413 Request Entity Too Large 请求体超长
- 415 Unsupported Media Type 不支持的媒体类型

- 500 Internal Server Error 服务器错误
- 502 Bad Gateway 网关错误
- 504 Gateway Timeout 网关超时
- 505 HTTP Version Not Supported 服务器不支持的使用的 HTTP 版本

#### URL 编码常用

- %21 !
- %22 "
- %23 #
- %24 $
- %25 %
- %26 &
- %27 '
- %28 (
- %29 )
- %2F /
- %30-9 0..9
- %3A :
- %3B ;
- %3C <
- %3D =
- %3E >
- %3F ?
- %40 @
- %41 A
