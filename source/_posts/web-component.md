---
title: web-component
date: 2024-02-21 15:16:43
categories:
  - WebAPI
tags:
  - API
  - web-component
---

## Web Component

Web Component 是一套不同的技术, 允许创建可重用的定制元素(它们的功能封装在代码之外)并且在 web 应用中使用它们

- Custom element(自定义元素): 一组 JavaScript API, 允许自定义元素及其行为, 然后在用户界面中按照需要使用它们
- Shadow DOM(影子 DOM): 一组 JavaScript API, 用于将封装的 "影子"DOM 树附加到指定元素(与页面 DOM 分开呈现)并控制其关联的功能. 通过这种方式, 可以保持自定义元素的功能私有, 这样它们就可以被脚本化和样式化,而不用担心与文档的其他部分发生冲突
- HTML template(HTML 模板): `<template>` 和 `<slot>` 元素可以编写不在呈现页面中显示的标记模板. 然后它们可以作为自定义元素结构的基础被多次重用

## 自定义元素

==封装== ==重用==

创建自定义元素, 扩展浏览器中可用的元素集

### 类型

- 自定义内置元素, 继承自标准的 HTML 元素, 例如 HTMLImageElement, HTMLParagraphElement
- 独立自定义元素, 继承自 HTML 元素基类 HTMLElement, 必须从头实现它们的行为

### 实现自定义元素

使用 ES6 中的类实现一个自定义元素, 该类可以扩展 HTMLElement 或者其它定制的接口

- 在构造函数中设置初始化状态和默认值, 注册事件监听器, 创建一个影子根(shadowRoot)
- 在构造函数中不能检查元素的属性或子元素, 不能添加新的属性或子元素

<!-- more -->

```javascript
class WordCount extends HTMLParagraphElement {
  constructor() {
    super();
  }
  /* 自定义元素功能 */
}

class PopupInfo extends HTMLElement {
  constructor() {
    // 必须首先调用 super 方法
    super();
  }
  connectedCallback() {
    // 创建影子根
    const shadow = this.attachShadow({ mode: 'open' });

    // 创建几个 span
    const wrapper = document.createElement('span');
    wrapper.setAttribute('class', 'wrapper');

    const icon = document.createElement('span');
    icon.setAttribute('class', 'icon');
    icon.setAttribute('tabindex', 0);

    const info = document.createElement('span');
    info.setAttribute('class', 'info');

    // 获取属性内容然后将其放入 info 这个 span 内
    const text = this.getAttribute('data-text');
    info.textContent = text;

    // 插入图标
    let imgUrl;
    if (this.hasAttribute('img')) {
      imgUrl = this.getAttribute('img');
    } else {
      imgUrl = 'img/default.png';
    }

    const img = document.createElement('img');
    img.src = imgUrl;
    icon.appendChild(img);

    // 创建一些 CSS 应用于影子 DOM
    const style = document.createElement('style');
    console.log(style.isConnected);

    style.textContent = `
      .wrapper {
        position: relative;
      }

      .info {
        font-size: 0.8rem;
        width: 200px;
        display: inline-block;
        border: 1px solid black;
        padding: 10px;
        background: white;
        border-radius: 10px;
        opacity: 0;
        transition: 0.6s all;
        position: absolute;
        bottom: 20px;
        left: 10px;
        z-index: 3;
      }

      img {
        width: 1.2rem;
      }

      .icon:hover + .info, .icon:focus + .info {
        opacity: 1;
      }
    `;

    // 将创建好的元素附加到影子 DOM 上
    shadow.appendChild(style);
    console.log(style.isConnected);
    shadow.appendChild(wrapper);
    wrapper.appendChild(icon);
    wrapper.appendChild(info);
  }
}
```

#### 自定义元素生命周期

`attributeChangedCallback` 生命周期函数 监听所有的属性的变化, 为了提高性能, 使用 静态属性 `observedAttributes` 声明需要监听变化的属性

- connectedCallback() 每当元素添加到文档中时调用, 对自定义元素的操作在此钩子中实现
- disconnectedCallback() 每当元素从文档中移除时调用
- adoptedCallback() 每当元素被移动到新文档中时调用
- attributeChangedCallback() 在属性更改、添加、移除或替换时调用, 接收三个参数,
  - name, 发生变化的属性名称
  - oldValue, 属性的旧值
  - newValue, 属性的新值

```javascript
window.customElements.define(
  'my-custom-element',
  class extends HTMLElement {

    // 静态属性声明需要监听变化的属性
    static observedAttributes = ['foo', 'bar'];
    
    constructor() {
      super();
    }
    connectedCallback() {
      console.log('自定义元素添加到页面');
    }
    disconnectedCallback() {
      console.log('自定义元素从页面中移除');
    }
    adoptedCallback() {
      console.log('自定义元素移动到新页面');
    }
    // 回调函数将在 observedAttributes 声明的属性发生变化时调用
    attributeChangedCallback(name, oldValue, newValue) {}
  }
);
```

#### 响应属性变化

```javascript
window.customElements.define(
  'my-custom-element',
  class extends HTMLElement {
    
    // 静态属性声明需要监听变化的属性
    static observedAttributes = ['size'];
    
    constructor() {
      super();
    }
    attributeChangedCallback(name, oldValue, newValue) {
      // size 在被更改时触发
      console.log(`属性 ${name} 已由 ${oldValue} 变更为 ${newValue}`);
    }
  }
);
```

```html
<!--  size 在被更改时触发 -->
<my-custom-element size="100"></my-custom-element>
```

### 自定义元素注册

注册自定义元素使用 `window.customElements.define()` 方法, 接收三个参数

- name, 自定义元素的名称, 使用 `kebab-case` 格式
- constructor, 自定义元素的构造函数
- options, 可选, 一个对象,
  - extends, 指定自定义元素要扩展的内置元素名称

```javascript
window.customElements.define(
  'word-count',
  class extends HTMLParagraphElement {
    constructor() {
      super();
    }
    connectedCallback() {
      console.log('自定义元素添加到页面');
    }
  },
  { extends: 'p' }
);

window.customElements.define('popup-info', PopupInfo);
```

### 使用自定义元素

- 自定义内置元素

使用内置元素时, 将自定义元素的名称作为 is 属性的值

```html
<p is="word-count"></p>
```

- 独立自定义元素

```html
<popup-info>
  <!-- 元素内容 -->
</popup-info>
```

## 影子 DOM

将一个 DOM 树附加到一个元素上, 并且使该树的的内部与页面中运行的 Javascript 和 CSS 相互隔离

- 影子宿主(Shadow host), 影子 DOM 附加到的常规 DOM 节点
- 影子树(Shadow tree), 影子 DOM 内部的 DOM 树
- 影子边界(Shadow boundary), 影子 DOM 终止, 常规 DOM 开始的地方
- 影子根(Shadow root), 影子树的根节点

### 创建影子 DOM

使用页面中指定的 DOM 元素作为影子宿主, 调用宿主的 `Element.attachShadow()` 方法创建影子 DOM

- Element.shadowRoot 通过影子宿主的 shadowRoot 属性访问影子 DOM 的内部
- Element.attachShadow() 创建影子 DOM
  - mode, 指定影子 DOM 树的封装模式
    - open, 允许从外部访问影子 DOM 根节点, Element.shadowRoot 返回一个 ShadowRoot 对象
    - closed, 拒绝从外部访问关闭的 影子 DOM 根节点, Element.shadowRoot 返回 null
  - delegatesFocus, 焦点委托, 当设置为 true 时, 指定减轻自定义元素的聚焦性能问题行为

```html
<div id="host"></div>

<script>
  const host = document.querySelector('#host');
  const shadow = host.attachShadow({ mode: 'open' });
  const span = document.createElement('span');
  span.textContent = "I'm in the shadow DOM";
  shadow.appendChild(span);
</script>
```

### CSS 封装

#### 编程式

> 创建单一样式表并将其与多个 DOM 树共享

通过构建一个 `CSSStyleSheet` 对象并将其附加到影子根

- replace() 和 replaceSync() 替换当前样式表的内容, 只能用在通过 CSSStyleSheet 构造函数创建的 styleSheet 对象上
  - replace() 方法异步的设置其内容, 返回一个 Promise
  - replaceSync() 方法同步的设置其内容

```javascript
// 创建一个空的 CSSStyleSheet 对象
const sheet = new CSSStyleSheet();
// 使用 replace 或 replaceSync 方法设置其内容
sheet.replaceSync('span { color: red; border: 2px dotted black;}');

const host = document.querySelector('#host');
// 创建影子 DOM
const shadow = host.attachShadow({ mode: 'open' });
// 将 styleSheet 添加到影子根的 adoptedStyleSheets 属性中
shadow.adoptedStyleSheets = [...shadow.adoptedStyleSheets, sheet];

const span = document.createElement('span');
span.textContent = "I'm in the shadow DOM";
shadow.appendChild(span);
```

- deleteRule(index) 从样式表中删除指定的样式规则, index 为样式规则的索引
- insertRule(rule [, index]) 向当前样式表指定位置插入样式规则, index 默认为 0, 返回值为新插入的规则在样式表中的索引

```javascript
sheet.insertRule("#blanc {color: white}", 0);

document.adoptedStyleSheets = [...document.adoptedStyleSheets, sheet];
```

#### 声明式

> 不需要在不同组件之间共享样式表

通过在 `template` 元素的声明中添加一个 `<style>` 元素

```html
<template id="my-custom-element">
  <style>
    span {
      color: red;
      border: 2px solid blue;
    }
  </style>
  <span>shadow DOM</span>
</template>

<div id="host"></div>

<script>
  const host = document.querySelector('#host');
  const shadow = host.attachShadow({ mode: 'open' });
  const template = document.getElementById('my-custom-element');

  shadow.appendChild(template.content);
</script>
```

## template 和 slots

### template

`template` 元素中的内容不会在 DOM 中呈现, 但仍可用 javascript 去引用它

```html
<template id="my-paragraph">
  <style>
    p {
      color: red;
      font-size: 18px;
    }
  </style>
  <p>This is my paragraph</p>
</template>

<script>
  window.customElements.define(
    'my-paragraph',
    class extends HTMLElement {
      constructor() {
        super();
        let template = document.getElementById('my-paragraph').content;

        const shadowRoot = this.attachShadow({ mode: 'open' });
        // Node.cloneNode() 返回调用该方法的节点的一个副本
        shadowRoot.appendChild(template.cloneNode(true));
      }
    }
  );
</script>
```

### slots

> 使用 slots 增加 template 元素的灵活性

- 标记中包含未定义相关的插槽内容或者浏览器不支持 slot 属性时显示默认内容

```html
<template id="my-paragraph">
  <style>
    p {
      color: red;
      font-size: 18px;
    }
  </style>
  <p>
    <slot name="my-text">default text</slot>
  </p>
</template>

<my-paragraph>
  <span slot="my-text">from slot</span>
</my-paragraph>

<script>
  window.customElements.define(
    'my-paragraph',
    class extends HTMLElement {
      constructor() {
        super();
        const template = document.getElementById('my-paragraph').content;
        const shadowRoot = this.attachShadow({ mode: 'open' }).appendChild(
          template.cloneNode(true)
        );
      }
    }
  );
</script>
```
