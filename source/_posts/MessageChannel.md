---
title: MessageChannel
date: 2024-08-09 15:20:07
categories:
  - WebAPI
tags:
  - API
---

## MessageChannel

允许创建一个新的消息通道, 并通过它的两个 MessagePort 属性发送数据

- port1 只读属性
- port2 只读属性

```javascript
const channel = new MessageChannel();
const output = document.querySelector(".output");
const iframe = document.querySelector("iframe");

// 处理 port1 收到的消息
channel.port1.onmessage = function(e) {
  output.innerHTML = e.data;
};

iframe.addEventListener("load", iframeLoaded, false);
function iframeLoaded() {
  // 把 port2 传给 iframe
  iframe.contentWindow.postMessage('Hello from the main page!', '*', [channel.port2]);
}

// 在 iframe 中...
window.addEventListener('message', (evt) => {
  const messagePort = evt.ports?.[0];
  messagePort.postMessage('Hello from the iframe!');
})
```
