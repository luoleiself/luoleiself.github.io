---
title: MutationObserver
date: 2021-07-26 17:05:57
categories:
  - WebAPI
tags:
  - API
---

## MutationObserver

MutationObserver 接口提供监视对 DOM 树所做更改的能力, 用于替代 Mutation Events 的新 API, 与 Events 不同的是: 事件是同步触发, 即 DOM 发生变动会立刻触发相应事件, MutationObserver 则是**异步触发**, DOM 发生变动以后并不会马上触发, 而是要等到当前所有 DOM 操作都结束后才触发, 所有监听操作以及相应的处理都是在其他任务执行完成之后异步执行的, 并且在 DOM 更改触发之后,将更改记录存储在数组之中, 统一进行回调通知

### 构造函数

创建并返回一个新的 `MutationObserver` 实例, 会在指定的 DOM 发生变化时被调用

- 参数 callback

  当被指定的节点或子树以及配置项有 Dom 变化时会被调用, 回调函数有两个参数:

  - MutationRecord 描述所有被触发改动的记录对象数组
  - MutationObserver 调用该函数的 MutationObserver 对象

```javascript
// 创建一个观察器并传入回调函数
const observer = new MutationObserver(function (MutationRecord, observer) {
  console.log(MutationRecord, observer);
  // [{
  //   addedNodes: NodeList []
  //   attributeName: ""
  //   attributeNamespace: null
  //   nextSibling: null
  //   oldValue: ""
  //   previousSibling: null
  //   removedNodes: NodeList []
  //   target:
  //   type: "attributes"
  // }]
});
// 指定观察变动的 DOM 节点和配置项
observer.observe(document.querySelector('#someElement'), {
  subtree: true,
  childList: true,
  attributes: true,
});
// 停止观察器
observer.disconnect();
```

- 返回值 MutationObserver 对象

<!-- more -->

### 实例方法

#### observer.disconnect()

阻止 MutationObserver 实例继续接收的通知, 直到再次调用其 observe 方法, 该观察者对象包含的回调函数都不会再被调用

- 参数 无
- 返回值 undefined

```javascript
observer.disconnect();
```

#### observer.observe()

配置 MutationObserver 在 DOM 更改匹配给定选项时, 通过其回调函数开始接收通知

- 参数
  - target DOM 树种一个要观察变化的 DOM Node
  - options 可选, 一个可选的 `MutationObserverInit` 对象, 此对象的配置描述了 DOM 的哪些变化应该提供给当前观察者的 callback
- 返回值 undefined

```javascript
observer.observe(Element, { subtree: true, childList: true, attributes: true });
```

#### observer.takeRecords()

> 使用场景是在断开观察者之前立即获取所有未处理的更改记录, 以便在停止观察者时可以处理任何未处理的更改

返回所有匹配 DOM 更改的**挂起的**状态队列并清除队列, 使变更队列保持为空

- 参数 无
- 返回值 `MutationRecord` 对象列表, 每个对象都描述了应用于 DOM 树某部分的一次改动

```javascript
const mutations = observer.takeRecords();
```

##### MutationRecord

每个 `MutationRecord` 代表一个独立的 DOM 变化在每次随 DOM 变化时作为 MutationObserver 回调函数的参数传入

- type String,
  - 如果是属性变化, 则返回 "attributes"
  - 如果是 characterData 节点变化, 则返回 "characterData"
  - 如果是子节点树 childList 变化, 则返回 "childList"
- target Node,
  - 根据 type 类型返回变化所影响的节点
- addedNodes NodeList, 返回被添加的节点,如果无则为空的 NodeList
- removedNodes NodeList, 返回被移除的节点, 如果无则为空的 NodeList
- previousSibling Node, 返回被添加或移除的节点之前的兄弟节点, 或者 null
- nextSibling Node, 返回被添加或移除的节点之后的兄弟节点,或者 null
- attributeName String, 返回被修改的属性的属性名, 或者 null
- attributeNamespace String, 返回被修改属性的命名空间, 或者 null
- oldValue String, 如果使当前属性有效,则需要在 MutationObserverInit 参数中配置 attributeOldValue 或者 characterDataOldValue 为 true
  - 如果属性 attributes 变化, 返回变化之前的属性值
  - 如果 characterData 变化, 返回变化之前的数据
  - 如果子节点树 childList 变化, 返回 null

```javascript
const observer = new MutationObserver((mutationRecord, observer) => {
  console.log(mutationRecord, observer);
});
observer.observe(document.querySelector('#root'), {
  attributes: true,
  attributeOldValue: true,
  attributeFilter: ['title'],
  characterData: true,
  characterDataOldValue: true,
});
// [{
//   addedNodes: NodeList []
//   attributeName: "title"
//   attributeNamespace: null
//   nextSibling: null
//   oldValue: "hello new root 1"
//   previousSibling: null
//   removedNodes: NodeList []
//   target: "div#root.helloRoothelloRoot1.helloRoot1"
//   type: "attributes"
// }]
```

### MutationObserverInit

描述了 MutationObserver 的配置, 主要被用作 observe 方法的参数模型

#### 参数属性

- attributeFilter 可选, 要监视的特定属性名称的数组, 未包含此属性则对所有属性的更改都会触发变动通知, 无默认值

- attributeOldValue 可选, 当监视节点的属性改动时, 将此属性设为 true 将记录任何有改动的属性的上一个值, 无默认值

- attributes 三选一,设为 true 以观察受监视元素的属性值变更, 默认 false

- characterData 三选一, 设为 true 以监视指定目标节点或子节点树中节点所包含的字符数据的变化, 无默认值

- characterDataOldValue 设为 true 以在文本在受监视节点上发生更改时记录节点文本的先前值, 无默认值

- childList 三选一, 设为 true 以监视目标节点(如果 subtree 为 true, 则包含子孙节点)添加或删除新的子节点, 默认 false

- subtree 设为 true 以将监视范围扩展至目标节点整个节点树中的所有节点, 默认 false

#### 注意

当调用 observe 方法时, childList, attributes, characterData 三个属性之中, 至少有一个必须为 true, 否则抛出 TypeError 异常
