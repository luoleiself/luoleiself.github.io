---
title: EventSource
date: 2024-02-23 16:09:48
categories:
  - WebAPI
tags:
  - API
---

## EventSource

> 当不使用 HTTP/2 时, 服务器发送事件(SSE)受到打开连接数(6)的限制, 这个限制是针对浏览器的. 每个浏览器的所有标签页中对相同域名的连接数最多支持 6 个

> 当使用 HTTP/2 时, 最大并发 HTTP 流的数量是由服务器和客户端协商的(默认为 100)

Web 内容与服务器发送事件通信的接口, 通信方向是单向的, 数据消息只能从服务器发送到客户端. 如果接收消息中有一个 event 字段, 触发的事件与 event 字段的值相同, 如果不存在 event 字段, 则将触发通用的 message 事件

- url, 表示远程资源的位置
- configuration, 可选, 一个对象
  - withCredentials, 标识 CORS 是否包含凭据, 默认为 false

```javascript
const sse = new EventSource(url, configuration);
/**
 * 没有 event 字段时触发
 * event: message
 * data: user data
 * id: someid
 */
sse.onmessage = function (e) {
  console.log(e);
};

/*
 * 触发 notice 事件回调
 *
 * event: notice
 * data: useful data
 * id: someid
 */
sse.addEventListener('notice', (e) => {
  console.log(evt);
});
```

### 实例属性

- readyState, 标识连接状态的数字
  - CONNECTING（0）
  - OPEN（1）
  - CLOSED（2）
- url, 表示远程资源的位置的字符串
- withCredentials, 标识是否使用跨域资源共享(CORS)凭据来实例化

### 实例方法

- close(), 关闭链接, 并将 readyState 属性设置为 CLOSED

### 事件

- error, 连接失败时触发
- message, 接收到数据时触发
- open, 连接打开时触发

```javascript
const sse = new EventSource('/api/v1/sse');

/*
 * 这将仅监听类似下面的事件
 *
 * event: notice
 * data: useful data
 * id: someid
 */
sse.addEventListener('notice', (e) => {
  console.log(e.data);
});

/*
 * 同理，以下代码将监听具有字段 `event: update` 的事件
 */
sse.addEventListener('update', (e) => {
  console.log(e.data);
});

/*
 * “message”事件是一个特例，因为它可以捕获没有 event 字段的事件，
 * 以及具有特定类型 `event：message` 的事件。
 * 它不会触发任何其他类型的事件。
 */
sse.addEventListener('message', (e) => {
  console.log(e.data);
});
```
