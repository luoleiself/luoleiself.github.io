---
title: IntersectionObserver
date: 2023-04-20 19:04:10
categories:
  - WebAPI
tags:
  - API
---

## IntersectionObserver

> 之前通常要使用事件监听, 并且需要频繁调用 Element.getBoundingClientRect() 方法以获取相关元素的边界信息, 并使用事件通知的方式同步进行计算, 可能造成性能问题

`IntersectionObserver` 接口提供了一种**异步观察**目标元素与祖先元素或顶级文档视口(viewport)交叉状态的方法
当一个 `IntersectionObserver` 对象被创建时, 其被配置为监听根中一段给定比例的可见区域, 一旦 `IntersectionObserver` 被创建, 则无法更改其配置, 所以一个给定的观察者对象只能用来监听可见区域的特定变化值, 但是, 可以在同一个观察者对象中配置监听多个目标元素

- 交叉状态触发是成对出现的, 就像 `mouseenter` 和 `mouseleave` 事件一样, 有进入状态在未来某一时刻肯定会有移出状态
- 在非交叉区域内的操作不会触发, 和 `scroll` 事件的触发是有区别的

### 应用范围

- 图片懒加载, 当图片滚动到可见区域时才进行加载
- 内容无限滚动, 用户滚动到接近内容底部时直接加载更多而无需用户操作翻页
- 检测广告的曝光情况, 为了计算广告收益, 需要直到广告元素的曝光情况
- 在到达可见区域时执行任务或播放动画

### 触发时机

通常只需要关注文档最接近的可滚动祖先元素的交集更改, 如果元素不是可滚动元素的后代, 则默认为设备视窗
如果要观察相对于根元素的交集, 则指定根元素为 null

- 每当目标元素与设备视窗或者其他指定元素发生交集的时候执行, 设备视窗或者其他元素被称为根元素或根
- Observer 第一次监听目标元素时

### 构造函数

创建并返回一个新的 `IntersectionObserver` 实例, 当其监听到目标元素的可见部分(的比例)超过了一个或多个**阈值**(threshold)时, 执行指定的回调函数
如果指定 `rootMargin` 则会检查其是否符合语法规定, 检查阈值以确保全部在 `0.0` 到 `1.0` 之间, 并且阈值列表会按升序排列, 如果阈值列表为空, 则默认为一个 [0.0] 的数组

```javascript
const ob = new IntersectionObserver(
  (entries, observer) => {
    // 如果 intersectionRatio 为 0, 则目标在视野外
    if (entries[0].intersectionRatio <= 0) {
      return;
    }
    console.log(entries, observer);
  },
  {
    // 使用文档视口作为 root
    root: null,
    // 添加到根的边界盒中的一组偏移量用于计算交叉值
    rootMargin: '0px 0px 0px 0px',
    // 指定元素在根的可见程度每多 25% 执行一次回调函数
    threshold: [0, 0.25, 0.5, 0.75, 1],
  }
);
```

#### 参数 callback

当元素的可见比例超过指定阈值后, 会调用此回调函数

- entries 一个 `IntersectionObserverEntry` 对象的数组, 每个被触发的阈值都稍微与指定阈值有偏差
- observer 被调用的 `IntersectionObserver` 实例

#### 可选 options

用来配置 observer 实例的对象, 如果未指定, observer 实例默认使用文档视口作为 root, 并且没有 `rootMargin`, 阈值为 0%

##### root

监听元素的祖先元素对象, 其边界盒将被视作视口, 如果未指定或者为 null, 则默认为浏览器视窗

##### rootMargin

一个计算交叉值时添加至根的边界盒中的一组偏移量, 类型为字符串, 可以有效的缩小或扩大根的判定范围从而满足计算需要, 默认为 `0px 0px 0px 0px`

##### threshold

规定了一个监听目标与边界盒交叉区域的比例值, 可以为一个具体的数值或是一组 `0.0` 到 `1.0` 之间的数组
如果指定元素在根的可见程度每多 25% 执行一次回调, 则指定为数组 `[0, 0.25, 0.5, 0.75, 1]`
如果指定元素在根的可见程度超过 50% 执行一次回调, 则指定为 `0.5`

- 若指定为 0.0, 监听元素与根即使有 1 像素交叉, 此元素也会被视为可见, 默认为 0.0
- 若指定为 1.0, 监听元素都在可见范围内时才被视为可见

### 实例属性

#### ob.root

只读属性, 获取当前的 `IntersectionObserver` 实例的根元素

#### ob.rootMargin

只读属性, 获取当前 `IntersectionObserver` 实例的根元素的边界盒的偏移量

#### ob.thresholds

只读属性, 获取当前 `IntersectionObserver` 实例的根元素与监听目标的交叉区域的比例值

### 实例方法

#### ob.disconnect()

停止 `IntersectionObserver` 实例对所有目标元素的监听行为

#### ob.observe(targetElement)

开始监听指定元素

- targetElement 可见性区域被监控的元素, 此元素必须是根元素的后代元素

```javascript
const ob = new IntersectionObserver(
  (entries, observer) => {
    entries.forEach((entry) => {
      if (entry.intersectionRatio > 0) {
        entry.target.classList.add('active');
      } else {
        entry.target.classList.remove('active');
      }
    });
  },
  {
    root: null,
    rootMargin: '0px 0px 0px 0px',
    threshold: [0, 0.5, 1],
  }
);
// 监听多个目标元素
document.querySelector('.section-content').forEach((el) => {
  ob.observe(el);
});
```

#### ob.takeRecords()

> 使用构造函数回调监听则不需要调用此方法, 调用此方法会清除挂起的交叉区域状态列表并不会执行构造函数中的回调

返回所有观察目标的**挂起的**交叉区域状态的 `IntersectionObserverEntry` 对象的数组

#### ob.unobserve(target)

取消对指定目标的监听行为

- target 要取消监听的目标, 如果未提供参数, 则不做任何操作

### IntersectionObserverEntry

接口描述了目标元素与其根元素容器在某一特定过渡时刻的交叉状态, `IntersectionObserverEntry` 的实例作为 entries 参数被传递到 IntersectionObserver 实例的回调函数中.
此外, 这些对象只能通过调用 `observer.takeRecords()` 来获取

#### 属性

- boundingClientRect 只读属性, 返回包含目标元素的边界信息, 计算方式与 `Element.getBoundingClientRect()` 相同
- intersectionRatio 只读属性, 返回 `intersectionRect` 与 `boundingClientRect` 的比例值
- intersectionRect 只读属性, 返回一个 `DOMRectReadOnly` 用来描述根与目标元素的相交区域
- isIntersecting 只读属性, 返回一个布尔值, 如果目标元素与交叉区域观察者对象的根相交则返回 true, 否则返回 false
- rootBounds 只读属性, 返回一个 `DOMRectReadOnly` 用来描述交叉区域观察者中的根
- target 只读属性, 与根出现相交区域改变的元素
- time 只读属性, 返回一个记录从 IntersectionObserver 的时间原点到交叉被触发的时间的时间戳

![IntersectionObserver-1](/images/IntersectionObserver-1.jpg)
![IntersectionObserver-2](/images/IntersectionObserver-2.jpg)
