---
title: BroadcastChannel
date: 2024-03-05 15:20:07
categories:
  - WebAPI
tags:
  - API
---

## BroadcastChannel

该接口代理了一个命名频道, 可以让同源下的任意浏览上下文订阅它, 它允许同源的不同浏览器窗口(非 chrome 和 qq 浏览器这种方式), tab 页, frame 或者 iframe 下的不同文档之间相互通信,
通过触发一个 message 事件, 消息可以广播到所有监听了该频道的 BroadCastChannel 对象.

- 构造函数

  - channelName, 表示通道名称的字符串, 对于相同的来源下的所有浏览上下文, 一个名称只对应一个通道

```javascript
const bc = new BroadcastChannel(channelName);
```

<!-- more -->

- 实例属性

  - name, 返回频道名称

- 实例方法

  - postMessage(), 向所有监听了相同频道的 bc 对象发送一条消息, 消息内容可以是任意类型的数据
  - close(), 关闭当前频道对象不再接收新的消息, 并允许被垃圾回收

- 事件

  - message, 当频道接收到一条消息时触发
  - messageerror, 当频道接收到一条无法反序列化的消息时触发

```javascript
const bc = new BroadcastChannel(channelName);
bc.onmessage = function(evt){
  console.log(evt)l;
}

bc.postMessage('hello BroadcastChannel');
```

![BroadcastChannel-1](/images/BroadcastChannel-1.jpg)

## localStorage 同源窗口通信

当 localStorage 被修改时会触发 storage 事件

```javascript
// 设置监听
window.addEventListener('storage', function (evt) {
  console.log('evt', evt);
});

// 移除监听
window.removeEventListener('storage', function () {
  console.log('evt', evt);
});
```

![localStorage-1](/images/localStorage-1.jpg)
