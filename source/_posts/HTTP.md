---
title: HTTP
date: 2021-07-26 19:12:29
categories:
  - network
tags:
  - http
---

### 状态码

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

### URL 编码常用

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

### [HSTS](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Strict-Transport-Security)

HTTP Strict-Transport-Security, HTTP 严格传输安全 是一种安全策略机制, 旨在强制客户端仅通过 HTTPS 协议与服务器进行通信, 并且在未来的一段时间内, 所有对该站点的请求都必须使用 HTTPS.

通过在 HTTP 响应头中包含一个特殊的指令: `Strict-Transport-Security` 来实现.

浏览器接收到这个响应头, 会将该站点标记为 HSTS 站点, 并在指定的时间内自动将所有的 HTTP 请求转换为 HTTPS 请求。即使是手动输入的 http url 或者点击的一个 http 链接也会被自动升级为 HTTPS 请求.

部分主流浏览器维护了一个预加载的 HSTS 站点列表, 这些站点默认被强制使用 HTTPS, 无需等待首次访问后的响应头设置.

- Chrome 的 HSTS 预加载列表: <https://www.chromium.org/hsts>
- Firefox 的 HSTS 预加载列表: [nsSTSPreloadList.inc](https://hg.mozilla.org/mozilla-central/raw-file/tip/security/manager/ssl/nsSTSPreloadList.inc)

```http
Strict-Transport-Security: max-age=<expire-time>; includeSubDomains; preload
```

- max-age, 浏览器记录的只能使用 HTTPS 访问站点的最大时间(单位秒)
- includeSubDomains, 可选, 表示该策略适用于主域名及其所有子域名
- preload, 可选, 表示该站点希望被列入浏览器的 HSTS 预加载列表

```bash
$ curl -I https://www.douyin.com
HTTP/1.1 404 Not Found
Server: Tengine
Content-Type: text/plain; charset=utf-8
Connection: keep-alive
Date: Wed, 11 Dec 2024 09:11:02 GMT
cache-control: no-store
Vary: Accept-Encoding, Accept-Encoding
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block
X-Download-Options: noopen
X-Content-Type-Options: nosniff
content-security-policy: upgrade-insecure-requests ;frame-ancestors https://pc.xgo.bytedance.net https://tcs.bytedance.net https://*.douyin.com https://aidp.bytedance.com https://aidp.bytedance.net https://www.aidp-cqc.com;script-src 'report-sample' 'strict-dynamic' 'nonce-4QE2wZ2rc-ztbxrA3IWHy' 'wasm-unsafe-eval' 'unsafe-eval' 'self' *.bytedance.com *.bytedance.net *.pstatp.com *.bytednsdoc.com *.bytescm.com *.douyin.com *.bytegoofy.com *.snssdk.com *.byted-static.com *.huoshanstatic.com *.douyinstatic.com *.ibytedapm.com *.zijieapi.com *.bytetos.com *.yhgfb-cn-static.com *.byteimg.com;report-uri https://i.snssdk.com/log/sentry/v2/api/slardar/main/?ev_type=csp&bid=douyin_web;report-to main-endpoint
Reporting-Endpoints: main-endpoint="https://mon.zijieapi.com/monitor_browser/collect/batch/security/?bid=douyin_web", default="https://mon.zijieapi.com/monitor_browser/collect/batch/security/?bid=douyin_web"
X-Tt-Logid: 20241211171102BFBC2AC0A24F4C05B24E
X-Agw-Info: N1yG692xcEzUcPRkqur4q1w8MlUT0hGFq0kNWMMV07yBP0rl8Ro_MQ8j_XpZ93UiGF_yYjjOLU0-LJkSol7AQbN5g-cwiWw79Mh3Nn_6upuBUoABSTZXiw5xsf5QXhFwSLwtY4xBUmCBdBQIjt7NUQ==
Server-Timing: inner; dur=276,tt_agw; dur=264
Set-Cookie: ttwid=1%7Ct5T1yUApy2aoqf8tihUlKDs27fKNe5akRSMvf5zx5nk%7C1733908262%7C209a31899a105d3c878e20feea33411d21a34e256860ec7250eec309baab71e4; Domain=.douyin.com; Path=/; Expires=Sat, 06 Dec 2025 09:11:02 GMT; HttpOnly
Set-Cookie: UIFID_TEMP=92559570181449f7274d01beb9fdcc50ede300fee49b9150f96a9b11ec20d282f33b1ce26fee42ba9461631433a50fa5108eb52097c37c530c81c602f7edb85602265a215753110a5ddc46f4e991edf4; path=/; expires=Sat, 25 Apr 2026 09:11:02 GMT; domain=douyin.com; samesite=none; secure
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
Access-Control-Allow-Credentials: true
x-tt-trace-host: 011ff03de9066d57c184f3d11219de2a9150664283b7d9b81d900f399ddf465413a8c909a66ed35f5f7f27db6607d80ca8d1103636a7e5a90d5b92105a23ba0862ae5cdac345035fa6422cb605c9b9dddb594f35c29470ce5c159141f9bdd32f9a
x-tt-trace-tag: id=03;cdn-cache=miss;type=dyn
server-timing: cdn-cache;desc=MISS,edge;dur=0,origin;dur=303
x-alicdn-da-ups-status: endOs,0,404
Via: live5.cn2073[303,0]
Timing-Allow-Origin: *
EagleId: 3db6831717339082625713776e
```

### [CSP](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CSP)

内容安全策略 是一个额外的安全层, 用于检测并消弱某些特定类型的攻击, 包括 跨站脚本(XSS) 和数据注入攻击等. 它通过定义哪些资源可以被加载和执行来限制网页的攻击面.

CSP 通过 HTTP 响应头或 HTML \<meta\> 标签传递给浏览器, 浏览器接收到这些策略后根据规则阻止不符合条件的资源加载或执行.

```http
Content-Security-Policy: default-src 'self'; img-src *; 
  media-src media1.com media2.com; 
  script-src userscripts.example.com;
  report-uri /_/csp_reports
```

```html
<meta http-equiv="Content-Security-Policy" content="default-src 'self'; img-src https://*; child-src 'none';" />
```

- default-src, 默认策略, 适用于所有未明确指定的资源类型, 如果指定了其他指令将覆盖此值
- script-src, 允许加载的 js 源, 可以指定域名、协议或路径, 或者使用关键字 `self`,`unsafe-inline`,`unsafe-eval`
- style-src, 允许加载的 CSS 样式表源
- img-src, 允许加载的图像源
- connect-src, 允许通过脚本接口加载的 URL
- font-src, 允许加载的字体源
- object-src, 允许加载的插件对象源, 如 flash
- frame-src, 允许嵌入的框架源
- media-src, 允许加载的音频和视频源
- child-src, (使用 frame-src 和 worker-src 代替)允许加载的子资源, 如 iframe 或 worker
- form-action, 允许提交表单的目的URL
- base-uri, 允许使用的 \<base\> 标签URI
- sandbox, 启用类似沙盒的限制
- prefetch-src, 允许预加载或预渲染的资源源
- report-uri, 指定违反 CSP 时发送报告的 URL
- report-to, 指定一个或多个组名, 用于接收 CSP 违规报告, 需要配合 Report-To HTTP 头使用

CSP 仅报告模式, 对任何违规行为将会报告一个指定的 URI 地址, 但不具有强制性

```http
Content-Security-Policy-Report-Only: default-src 'self'; 
  img-src *; 
  media-src media1.com media2.com; 
  script-src userscripts.example.com; 
  report-uri /_/csp-reports
```
