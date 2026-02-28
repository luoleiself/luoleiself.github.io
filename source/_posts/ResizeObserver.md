---
title: ResizeObserver
date: 2021-07-26 17:05:57
categories:
  - WebAPI
tags:
  - API
---

## ResizeObserver

ResizeObserver 接口监视 Element 内容盒或边框盒或者 SVGElement 边界尺寸的变化

### 构造函数

ResizeObserver 构造函数创建一个新的 ResizeObserver 对象，它可以用于监听 Element 内容盒或边框盒或者 SVGElement 边界尺寸的大小

- 参数 callback

  每当观测的元素调整大小时，调用该函数。该函数接收两个参数

  - entries 一个 `ResizeObserverEntry` 对象数组，可以用于获取每个元素改变后的新尺寸
  - observer  对 ResizeObserver 自身的引用

```tsx
const observer = new ResizeObserver(function (entries, observer) {
  console.log(entries, observer);
  // [{
  //   ResizeObserverEntry.borderBoxSize
  //   ResizeObserverEntry.contentBoxSize
  //   ResizeObserverEntry.devicePixelContentBoxSize
  //   ResizeObserverEntry.contentRect
  //   ResizeObserverEntry.target
  // }]
});
// 开始监视指定 Element
observer.observer(document.querySelector('#someElement'));
// 停止所有监视
observer.disconnect();
```

### 实例方法

#### observer.disconnect()

法取消所有的对 Element 或 SVGElement 目标的监听

#### observer.observe()

监听指定的 Element 或 SVGElement

- 参数
  - target 监听的目标
  - options 配置参数
    - box 设置监听的盒模型, 值为 content-box(默认), border-box, device-pixel-content-box

#### observer.unobserve()

结束对指定的 Element 或 SVGElement 的监听

- 参数
  - target 取消监听的目标

### ResizeObserverEntry

传递给 ResizeObserver() 构造函数中的回调函数参数的对象

#### 参数属性

- borderBoxSize 只读, 一个对象，当运行回调时，该对象包含着正在观察元素的新边框盒的大小。
- contentBoxSize 只读, 一个对象，当运行回调时，该对象包含着正在观察元素的新内容盒的大小。
- devicePixelContentBoxSize 只读, 一个对象，当运行回调时，该对象包含着正在观察元素的新内容盒的大小（以设备像素为单位）。
- contentRect 只读, 一个对象，当运行回调时，该对象包含着正在观察元素新大小的 DOMRectReadOnly 对象。请注意，这比以上两个属性有着更好的支持，但是它是 Resize Observer API 早期实现遗留下来的，出于对浏览器的兼容性原因，仍然被保留在规范中，并且在未来的版本中可能被弃用。
- target 只读, 对正在观察 Element 或 SVGElement 的引用。
