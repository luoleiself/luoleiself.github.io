---
title: TreeWalker
date: 2022-03-05 10:35:34
categories:
  - WebAPI
  - ES
tags:
  - API
  - ES6
  - js
---

## [TreeWalker](https://developer.mozilla.org/zh-CN/docs/Web/API/TreeWalker)

表示文档子树中的节点和它们的位置

- root 表示对象的根节点
- whatShow, 默认值: 0xFFFFFFFF, 表示位掩码的 unsigned long, 由 [NodeFilter](https://dom.spec.whatwg.org/#interface-nodefilter) 的常用属性组合而成, 此参数便于筛选出特定类型的节点
  - 0xFFFFFFFF, NodeFilter.SHOW_ALL 显示所有节点
  - 0x2       , NodeFilter.SHOW_ATTIBUTE 显示 Attr 节点
  - 0x80      , NodeFilter.SHOW_COMMENT 显示 Comment 节点
  - 0x1       , NodeFilter.SHOW_ELEMENT 显示 Element 节点
  - 0x4       , NodeFilter.SHOW_TEXT    显示 Text 节点
- filter, 回调函数或包含 acceptNode() 方法的对象, 其返回值为 NodeFilter.FILTER_ACCEPT, NodeFilter.FILTER_REJECT 或 NodeFilter.FILTER_SKIP
  - NodeFilter.FILTER_ACCEPT 包含此节点
  - NodeFilter.FILTER_REJECT 不包含以此节点为根的子树中的任意节点
  - NodeFilter.FILTER_SKIP   不包含此节点

```javascript
// 返回新创建的 TreeWalker 对象 
const tw = document.createTreeWalker(root, whatToShow, filter);

const treeWalker = document.createTreeWalker(document.body, NodeFilter.SHOW_ElEMENT, (node) => 
  node.classList.contains('no-escape') 
  ? NodeFilter.FILTER_REJECT
  : node.closest('.escape')
    ? NodeFilter.FILTER_ACCEPT
    : NodeFilter.FILTER_SKIP
);
while (treeWalker.nextNode()) {
  for (const node of treeWalker.currentNode.childNodes) {
    if (node.nodeType === Node.TEXT_NODE && /\S/.test(node.data)) {
      // 排除仅含空白符的文本节点
      node.data = encodeURI(node.data.replace(/\s+/g, " "));
    }
  }
}
```

<!--more-->

### treeWalker

- root, 只读属性, 表示新建 TreeWalker 时所声明的根节点
- whatToShow, 只读属性, 返回一个 unsigned long 类型的常量位掩码, 表示需要筛选的 Node 类型
- filter, 只读属性, 返回一个实现 NodeFilter 接口对象, 用来挑选相关的节点
- currentNode, 返回 TreeWalker 当前指向的 Node

- parentNode(), 移动当前 Node 到文档顺序中的第一个可见的祖先节点并返回该节点, 如果没有则返回 null 同时不会发生移动
- firstChild(), 移动当前 Node 到当前节点的第一个可见子节点并返回该节点, 如果没有则返回 null 同时不会发生移动
- lastChild(), 移动当前 Node 到当前节点的最后一个可见子节点并返回该节点, 如果没有则返回 null 同时不会发生移动
- previousSibling(), 移动当前 Node 到当前节点的前一个兄弟节点并返回该节点, 如果没有则返回 null 同时不会发生移动
- nextSibling(), 移动当前 Node 到当前节点的后一个兄弟节点并返回该节点, 如果没有则返回 null 同时不会发生移动
- previousNode(), 移动当前 Node 到文档顺序中前一个节点并返回该节点, 如果没有则返回 null 同时不会发生移动
- nextNode(), 移动当前 Node 到文档顺序中下一个节点并返回该节点, 如果没有则返回 null 同时不会发生移动

```html
<template>
  <!DOCTYPE html>
  <html lang="en">
    <head><title>Demo</title>
    <body>
      <div id="container"></div>
    </body>
  </html> 
</template>
<script>
  const w1 = document.createTreeWalker(document.body, NodeFilter.SHOW_ALL);
  const node = w1.firstChild(); // nodeName: #Text

  const w2 = document.createTreeWalker(document.body, NodeFilter.SHOW_ELEMENT);
  const el = w2.firstChild(); // nodeName: DIV  
</script>
```
