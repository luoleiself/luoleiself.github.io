---
title: popover
date: 2025-11-12 17:32:32
categories:
  - WebAPI
tags:
  - API
---

## popover

用来指定一个元素为弹出框元素, 弹出框元素通过 `display: none` 被隐藏，直到通过调用/控制元素(即带有 popovertarget 属性的 button 或者 `type="button"` 的 input) 或 `HTMLElement.showPopover()` 调用打开。

当打开时，弹出框元素将出现在所有其他元素之上，即在顶层上，并且不会受到父元素的 position 或 overflow 样式的影响。

### 属性

- popover  指定元素为弹框出框元素
  - auto, 默认值, 可以在弹出框之外的区域进行选择或(Esc 键), 以达到 轻触关闭 的目的
  - manual, 手动模式, 不能 轻触关闭 或自动关闭, 必须通过声明式的 button 或 js 控制
  - hint, 提示模式, 在弹出框显示时不会自动关闭自动型弹出框, 但会关闭其他提示型弹出框, 并且会对关闭请求做出响应
  
- popovertarget 指定控制弹出框元素的 id，在 button 或 `type="button"` 的 input 元素上使用

- popovertargetaction 指定弹出框元素的行为
  - toggle, 默认值, 切换弹出框
  - hide, 隐藏弹出框
  - show, 显示弹出框

```html
<button popovertarget="myPopover" popovertargetaction="show">
  Show Popover
</button>
<button popovertarget="myPopover" popovertargetaction="hide">
  Hide Popover
</button>

<div id="myPopover" popover>popover content</div>
```

手动控制弹出框

```html
<button popovertarget="my-popover" popovertargetaction="toggle">Toggle Popover</button>

<div id="my-popover" popover="manual">Popover content</div>
```

### 方法

- HTMLElement.showPopover() 手动打开弹出框
- HTMLElement.hidePopover() 手动关闭弹出框
- HTMLElement.togglePopover() 手动切换弹出框

<!-- more -->

```html
<div id="mypopover" popover>
  <h2>Help!</h2>

  <p>You can use the following commands to control the app</p>

  <ul>
    <li>Press <ins>C</ins> to order cheese</li>
    <li>Press <ins>T</ins> to order tofu</li>
    <li>Press <ins>B</ins> to order bacon</li>
  </ul>
  <hr />
  <ul>
    <li>Say "Service" to summon the robot waiter to take your order</li>
    <li>Say "Escape" to engage the ejector seat</li>
  </ul>
</div>

<script>
const popover = document.getElementById("mypopover");

document.addEventListener("keydown", (event) => {
  if (event.key === "h") {
    popover.showPopover();
  }
});
</script>
```

### 事件

- HTMLElement.beforetoggle  事件，在弹出框状态切换之前触发, 可以阻止弹出框的打开
- HTMLElement.toggle  事件，在弹出框状态切换之后触发
