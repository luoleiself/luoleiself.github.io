---
title: CSS3
date: 2024-07-26 22:24:03
categories: CSS
tags: CSS
---

### accent-color

为某些元素所生成的用户界面控件设置了**强调色**

- auto 表示用户代理所选颜色, 应匹配平台的强调色(若有)
- \<color\> 指定用作强调色的颜色

```html
<input type="checkbox" checked/>
<input type="radio" />
<style>
  input {
    accent-color: #74992e;
  }
</style>
```
<!-- more -->

### color-scheme

允许元素指示它可以舒适地呈现哪些颜色方案, 操作系统颜色方案的常见选择为"亮色"和"暗色", 或"日间模式"和"夜间模式". 当选择其中一种配色方案时, 操作系统会对用户界面进行调整, 包括表单控件、滚动条和 CSS 系统颜色的使用值

- normal 表示元素未指定任何配色方案, 因此应使用浏览器的默认配色方案呈现
- light 表示可以使用操作系统**亮色**配色方案渲染元素
- dark 表示可以使用操作系统**深色**配色方案渲染元素
- only 禁止用户代理覆盖元素的配色方案

为整个页面选择配色方案在 `:root` 元素上指定 color-scheme, 设置之后不会根据系统的主题色自动调整

```html
<!-- 为元素指定配色方案 -->
<textarea style="color-scheme:dark;"></textarea>
<input type="text" style="color-scheme:only light;"/>

<style>
  /* 为整个页面指定暗色 */
  :root {
    color-scheme: dark;
  }
</style>
```

### prefers-color-scheme

CSS 媒体特性用于检测用户是否有将系统的主题色设置为**亮色**或**深色**, 可以根据系统的主题色自动调整, 前提需要定义好配色方案

- no-preference 表示系统未得知用户在配色方案的选项, 在布尔值上下文中执行结果为 false, 此结果无法通过媒体特性获取系统是否支持设置主题色
- light 表示用户已告知系统选择使用亮色主题的界面
- dark 表示用户已告知系统选择使用深色主题的界面

```css
/* 设备最小高度为 680px, 或者为纵向模式的屏幕设备 */
@media(min-height: 680px), screen and (orientation: portrait) {
  /* ... */
}
.light {
  background: #eee;
  color: black;
}
.dark {
  background: #333;
  color: white;
}
/* 媒体查询配色方案 */
@media(prefers-color-scheme: light){
  .light.light-scheme {
    background: white;
    color: #555;
  }
  .dark.light-scheme {
    background: #eee;
    color: black;
  }
}
@media(prefers-color-scheme: dark){
  .light.dark-scheme {
    background: #333;
    color: white;
  }
  .dark.dark-scheme {
    background: black;
    color: #ddd;
  }
}
```

```html
<div class="light">Day (initial)</div>
<div class="light light-scheme">Day (changes in light scheme)</div>
<div class="light dark-scheme">Day (changes in dark scheme)</div>
<br />

<div class="dark">Night (initial)</div>
<div class="dark light-scheme">Night (changes in light scheme)</div>
<div class="dark dark-scheme">Night (changes in dark scheme)</div>
```

### pointer-events

指定在什么情况下(如果有)某个特定的图形元素可以成为鼠标事件的 target

- none 指定元素永远不会成为鼠标事件的 target, 当其后代元素的 pointer-events 属性指定其他值时, 鼠标事件可以指向后代元素
      例如: 页面上显示自定义水印效果不影响页面交互, 使用容器元素包含水印覆盖整个 body 将优先级设置为最高并设置此属性使容器元素不被鼠标事件选中

### 布局

#### place-content

简写属性

- align-content 设置 flex 布局的纵轴和 grid 布局的主轴在内容项之间和周围分配空间
- justify-content 设置 flex 布局和 grid 布局主轴上的对齐方式

#### place-items 简写属性

简写属性

- align-items 设置 flex 布局的交叉轴和 grid 布局的块向轴上的子元素的对齐方式
- justify-items 设置所有盒中的项目沿适当轴线对齐到每个盒子

#### 单个盒子属性

- align-self  设置 flex 布局和 grid 布局中的子元素的行为, 覆盖 align-items 的值
- justify-self  设置单个盒子在其布局容器适当轴中的对齐方式
