---
title: EventSource
date: 2024-02-23 16:09:48
categories:
  - WebAPI
tags:
  - API
---

## EventSource

> 当不使用 HTTP/2 时, 服务器发送事件(SSE)受到打开连接数(6)的限制, 这个限制是针对浏览器的. 每个浏览器的所有标签页中对相同域名的连接数最多支持 6 个(chrome 超过 6 个自动断开连接, firefox 超过 9 个自动断开连接)

> 当使用 HTTP/2 时, 最大并发 HTTP 流的数量是由服务器和客户端协商的(默认为 100)

Server-Sent Events 服务器发送事件

Web 内容与服务器发送事件通信的接口, 通信方向是单向的, 数据消息只能从服务器发送到客户端. 如果接收消息中有一个 event 字段, 触发的事件与 event 字段的值相同, 如果不存在 event 字段, 则将触发通用的 message 事件

- url, 表示远程资源的位置
- configuration, 可选, 一个对象
  - withCredentials, 标识 CORS 是否包含凭据, 默认为 false

### 消息格式

数据以 `文本流(text/event-stream)` 的形式从服务器发送到客户端, 每一行数据都必须以 `换行符(\n)` 结束, 每条消息以两个 `换行符(\n\n)` 结束

- event 自定义事件类型(可选), 不使用默认触发 message
- id 消息 ID, 用于断线重连定位
- retry 重连间隔(毫秒)
- data 数据内容

```javascript
const sse = new EventSource(url, configuration);
/**
 * 没有 event 字段时触发
 * event: message
 * id: someid
 * data: user data
 */
sse.onmessage = function (e) {
  console.log(e);
};

/*
 * 触发 notice 事件回调
 *
 * event: notice
 * id: someid
 * data: useful data
 */
sse.addEventListener('notice', (e) => {
  console.log(evt);
});
```

<!-- more -->

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

连接服务器

```html
<!-- client -->
<p><button id="close">Close</button></p>
<ul id="list"></ul>
<script>
  const sse = new EventSource('/evt-source', { withCredentials: true });
  const list = document.getElementById('list');
  const liMessageColor = 'rgb(' + Math.floor(Math.random() * 256) + ',' + Math.floor(Math.random() * 256) + ',' + Math.floor(Math.random() * 256) + ')'
  const liNoticeColor = 'rgb(' + Math.floor(Math.random() * 256) + ',' + Math.floor(Math.random() * 256) + ',' + Math.floor(Math.random() * 256) + ')'
  const btn = document.getElementById('close')

  btn.onclick = function () {
    sse.close()
  }

  sse.addEventListener('connected', function (e) {
    console.log('connected', e)
  })

  /*
  * 这将仅监听类似下面的事件
  *
  * event: notice
  * id: someid
  * data: useful data
  */
  sse.addEventListener('notice', function (e) {
    console.log('notice', e)
    const data = JSON.parse(e.data)
    const li = document.createElement('li')
    li.style.color = liNoticeColor
    li.innerHTML = 'event: ' + e.type + ' count:' + data.count
    list.appendChild(li)
  })

  /*
  * “message”事件是一个特例，因为它可以捕获没有 event 字段的事件，
  * 以及具有特定类型 `event：message` 的事件。
  * 它不会触发任何其他类型的事件。
  */
  sse.addEventListener('message', function (e) {
    console.log('message', e)
    const data = JSON.parse(e.data)
    const li = document.createElement('li')
    li.style.color = liMessageColor
    li.innerHTML = 'event: ' + e.type + ' count:' + data.count
    list.appendChild(li)
  })

  sse.addEventListener('closed', function (e) {
    console.log('closed', e)
  })

  sse.addEventListener('open', function (e) {
    console.log('open', e)
  })

  sse.addEventListener('error', function (e) {
    console.log('error', e)
  })
</script>
```

服务器推送数据

```javascript
// server.js
import express from 'express';

const app = express()
let count = 1

app.use(express.static('./'))

app.get('/evt-source', (req, res) => {
  res.setHeader('Content-Type', 'text/event-stream');
  res.setHeader('Cache-Control', 'no-cache');
  res.setHeader('Connection', 'keep-alive');

  // 定期发送事件给客户端
  const sendEvent = (evtName, data) => {
    /**
     * event: notice
     * id: 1738749223324
     * retry: 500
     * data: {"message":"Message at 2025-02-05T09:53:43.324Z","count":1}
     */
    // 数据以文本流的形式从服务器发送到客户端, 每一行数据都必须以换行符(\n)结束, 
    // 重连时间间隔 500 毫秒
    // 并且每条消息以两个换行符(\n\n)结束
    res.write(`event: ${evtName}\nid: ${Date.now()}\nretry: 500\ndata: ${JSON.stringify(data)}\n\n`);
  };

  // 发送初始事件
  sendEvent('connected', { message: 'Connection established' });

  // 模拟每 5 秒发送一次消息
  const timer = setInterval(() => {
    sendEvent(Math.random() > 0.5 ? 'notice' : 'message', { message: `Message at ${new Date().toISOString()}`, count: count });
    count++;
  }, 5000);

  // 当客户端断开连接时清除定时器
  req.on('close', () => {
    clearInterval(timer);
    res.end()
  });
})

app.get('/eventSource.html', (req, res) => {
  res.sendFile('./eventSource.html')
})

app.listen(8080, () => {
  console.log('app listen localhost:8080')
})
```
