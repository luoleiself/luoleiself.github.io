---
title: web-component
date: 2024-02-21 15:16:43
categories:
  - WebAPI
tags:
  - API
---

## 自定义元素

==封装== ==重用==

创建自定义元素, 扩展浏览器中可用的元素集,

### 类型

- 自定义内置元素, 继承自标准的 HTML 元素, 例如 HTMLImageElement, HTMLParagraphElement
- 独立自定义元素, 继承自 HTML 元素基类 HTMLElement, 必须从头实现它们的行为

### 实现自定义元素

使用 ES6 中的类实现一个自定义元素, 该类可以扩展 HTMLElement 或者其它定制的接口

- 在构造函数中设置初始化状态和默认值, 注册事件监听器, 创建一个影子根(shadowRoot)
- 在构造函数中不能检查元素的属性或子元素, 不能添加新的属性或子元素

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
    attributeChangedCallback(name, oldValue, newValue) {}
  }
);
```

#### 响应属性变化

```javascript
window.customElements.define(
  'my-custom-element',
  class extends HTMLElement {
    static observedAttributes = ['size']; // 静态属性
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

使用页面中指定的 DOM 元素作为影子宿主, 调用宿主的 `attachShadow()` 方法创建影子 DOM

- Element.shadowRoot 通过影子宿主的 shadowRoot 属性访问影子 DOM 的内部
- Element.attachShadow()
  - mode, 指定影子 DOM 树的封装模式
    - open, 允许从外部访问影子 DOM 根节点, Element.shadowRoot 返回一个 ShadowRoot 对象
    - closed, 拒绝从外部访问关闭的 影子 DOM 根节点, 返回 null
  - delegatesFocus, 焦点委托, 当设置为 true 时, 指定减轻自定义元素的聚焦性能问题行为

```javascript
const host = document.querySelector('#host');
const shadow = host.attachShadow({ mode: 'open' });
const span = document.createElement('span');
span.textContent = "I'm in the shadow DOM";
shadow.appendChild(span);
```

### CSS 封装
