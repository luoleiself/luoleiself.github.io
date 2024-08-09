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
const para = document.querySelector("p");

const ifr = document.querySelector("iframe");

ifr.addEventListener("load", iframeLoaded, false);

function iframeLoaded() {
  ifr.otherWindow.postMessage("来自主页的问候！", "*", [channel.port2]);
}

channel.port1.onmessage = function(e) {
  para.innerHTML = e.data;
};
```
