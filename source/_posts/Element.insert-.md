---
title: Element.insert*
date: 2023-07-21 12:02:45
categories:
  - WebAPI
tags:
  - API
---

### Element.insertAdjacentElement(position, element)

将一个给定的元素节点插入到相对于被调用元素的给定的一个位置, 执行成功返回插入的元素, 失败则返回 null

- position 表示相对于该元素的位置
  - 'beforebegin' 在该元素本身的前面
  - 'afterbegin' 只在该元素中, 在该元素中的第一个子元素前面
  - 'beforeend' 只在该元素中, 在该元素中的最后一个子元素的后面
  - 'afterend' 在该元素本身的后面
- element 要插入的元素

### Element.insertAdjacentHTML(position, text)

将指定的文本解析为 Element 元素, 并将结果插入到 DOM 树中的指定位置. 它不会重新解析它正在使用的元素, 因此不会破坏元素内的现有元素, 这避免了额外的序列化步骤, 使其比直接使用 innerHTML 操作更快

- position 表示插入内容相对于元素的位置
  - 'beforebegin' 在该元素本身的前面
  - 'afterbegin' 只在该元素中, 在该元素中的第一个子元素前面
  - 'beforeend' 只在该元素中, 在该元素中的最后一个子元素的后面
  - 'afterend' 在该元素本身的后面
- text 被解析为 HTML 或 XML 元素并插入到 DOM 树中的 DOMString

### Element.insertAdjacentText(position, element)

将一个给定的文本节点插入到相对于被调用元素的给定的一个位置

- position 表示相对于该元素的位置
  - 'beforebegin' 在该元素本身的前面
  - 'afterbegin' 只在该元素中, 在该元素中的第一个子元素前面
  - 'beforeend' 只在该元素中, 在该元素中的最后一个子元素的后面
  - 'afterend' 在该元素本身的后面
- element 被插入的文本节点的 DOMString

<!-- more -->