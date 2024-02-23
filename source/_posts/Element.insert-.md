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

```html
<!-- beforebegin -->
<p>
  <!-- afterbegin -->
  foo
  <!-- beforeend -->
</p>
<!-- afterend -->

<script>
  // 只有当节点处于 DOM 树中且有一个父节点时 beforebegin 和 afterend 才会生效
  beforeBtn.addEventListener('click', () => {
    const tempDiv = document.createElement('div');
    tempDiv.style.backgroundColor = randomColor();
    if (activeElem) {
      activeElem.insertAdjacentElement('beforebegin', tempDiv);
    }
    setListener(tempDiv);
  });

  afterBtn.addEventListener('click', () => {
    const tempDiv = document.createElement('div');
    tempDiv.style.backgroundColor = randomColor();
    if (activeElem) {
      activeElem.insertAdjacentElement('afterend', tempDiv);
    }
    setListener(tempDiv);
  });
</script>
```

### Element.insertAdjacentHTML(position, text)

将指定的文本解析为 Element 元素, 并将结果插入到 DOM 树中的指定位置. 它不会重新解析它正在使用的元素, 因此不会破坏元素内的现有元素, 这避免了额外的序列化步骤, 使其比直接使用 innerHTML 操作更快

- position 表示插入内容相对于元素的位置
  - 'beforebegin' 在该元素本身的前面
  - 'afterbegin' 只在该元素中, 在该元素中的第一个子元素前面
  - 'beforeend' 只在该元素中, 在该元素中的最后一个子元素的后面
  - 'afterend' 在该元素本身的后面
- text 被解析为 HTML 或 XML 元素并插入到 DOM 树中的 DOMString

```html
<script>
  // 只有当节点处于 DOM 树中且有一个父节点时 beforebegin 和 afterend 才会生效
  // 原为 <div id="one">one</div>
  var d1 = document.getElementById('one');
  d1.insertAdjacentHTML('afterend', '<div id="two">two</div>');

  // 此时，新结构变成：
  // <div id="one">one</div><div id="two">two</div>
</script>
```

### Element.insertAdjacentText(position, element)

将一个给定的文本节点插入到相对于被调用元素的给定的一个位置

- position 表示相对于该元素的位置
  - 'beforebegin' 在该元素本身的前面
  - 'afterbegin' 只在该元素中, 在该元素中的第一个子元素前面
  - 'beforeend' 只在该元素中, 在该元素中的最后一个子元素的后面
  - 'afterend' 在该元素本身的后面
- element 被插入的文本节点的 DOMString

```html
<script>
  // 只有当节点处于 DOM 树中且有一个父节点时 beforebegin 和 afterend 才会生效
  beforeBtn.addEventListener('click', function () {
    para.insertAdjacentText('afterbegin', textInput.value);
  });

  afterBtn.addEventListener('click', function () {
    para.insertAdjacentText('beforeend', textInput.value);
  });
</script>
```

<!-- more -->

### Element.replaceChildren()

将多个 Node 对象替换为该节点的后代集合

### Element.replaceWith()

将多个 Node 对象或字符串对象替换该节点父节点下的子节点

```html
<script>
  const div = document.createElement('div');
  const p = document.createElement('p');
  div.appendChild(p);
  const span = document.createElement('span');

  p.replaceWith(span);

  console.log(div.outerHTML);
  // "<div><span></span></div>"
</script>
```
