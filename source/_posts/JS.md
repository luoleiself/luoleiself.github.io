---
title: JS
date: 2022-03-05 10:35:34
categories:
  - ES
tags:
  - ES6
  - js
---

> --no-sandbox --disable-web-security --user-data-dir=C:\chromedata

> fetch() 表单上传时, 不能设置 `Content-Type` 头, 否则会丢失文件边界

### 浏览器引擎

| 浏览器  |    渲染引擎    |                               js 引擎                                |
| :-----: | :------------: | :------------------------------------------------------------------: |
|   IE    |    Trident     |                 JScript(IE3.0-IE8.0) / Chakra(IE9~)                  |
| Chrome  | webkit / Blink |                                  V8                                  |
| Safari  |     webkit     |                         Nitro(SquirrelFish)                          |
| Firefox |     Gecko      | ~Monkey 系列(SpiderMonkey / TraceMonkey / JaegerMonkey / OdinMonkey) |
|  Opera  | WebKit / Blink |                               Carakan                                |

### 改变原数组的方法

- pop 从数组中删除最后一个元素,并返回该元素的值(数组为空时返回 undefined). 此方法更改数组的长度
- push 将一个或多个元素添加到数组的末尾, 并返回该数组的新长度
- shift 从数组中删除第一个元素,并返回该元素的值(数组为空则返回 undefined). 此方法更改数组的长度
- unshift 将一个或多个元素添加到数组的开头. 并返回该数组的新长度(该方法修改原有数组)
- reverse 将数组中元素的位置颠倒, 并返回该数组. 该方法会改变原数组
- sort 用[原地算法](https://en.wikipedia.org/wiki/In-place_algorithm)对数组的元素进行排序, 并返回数组
  - compareFunction 用来指定按某种顺序进行排列的函数, 省略则按照转换为的字符串的 unicode 位点进行排序
    - firstEl 第一个用于比较的元素
    - secondEl 第二个用于比较的元素
- splice 通过删除或替换现有元素或者原地添加新的元素来修改数组, 并以数组形式返回被修改的内容. 此方法会改变原数组
  - start 指定修改的开始位置
  - deleteCount 整数, 表示要移除的数组元素的个数, 如果为 0 或者负数, 则不移除元素
  - item1, item2 要添加进数组的元素,从 start 位置开始, 不指定则删除数组元素
- fill 用一个固定值填充一个数组中从起始索引到终止索引内的全部元素. 不包括终止索引. 返回修改后的数组
  - value 用来填充数组元素的值
  - start 起始索引, 默认值为 0
  - end 终止索引, 默认值为 this.length
- copyWithin 浅复制数组的一部分到同一数组中的另一个位置, 并返回它, 不会改变原数组的长度. 返回改变后的数组
  - target 整数, 复制序列到该位置, 如果是负数, target 将从末尾开始计算. 如果 target 大于等于 arr.length, 将会不发生拷贝
  - start 整数, 开始复制元素的起始位置, 如果是负数, start 将从末尾开始计算. 如果 start 被忽略, copyWithin 将会从 0 开始复制
  - end 整数, 开始复制元素的结束位置, copyWithin 将会拷贝到该位置, 但不包括 end 这个位置的元素. 如果是负数, end 将从末尾开始计算. 如果忽略则复制到数组结尾

<!-- more -->

```javascript
arr.pop();
arr.push(element1, ..., elementN);
arr.shift();
arr.unshift(element1, ..., elementN);
arr.reverse();
arr.sort([compareFunction]);
array.splice(start[, deleteCount[, item1[, item2[, ...]]]]);

arr.fill(value[, start[, end]]);
const array1 = [1, 2, 3, 4];
// fill with 0 from position 2 until position 4
console.log(array1.fill(0, 2, 4));
// expected output: [1, 2, 0, 0]
// fill with 5 from position 1
console.log(array1.fill(5, 1));
// expected output: [1, 5, 5, 5]
console.log(array1.fill(6));
// expected output: [6, 6, 6, 6]

arr.copyWithin(target[, start[, end]]);
const array1 = ['a', 'b', 'c', 'd', 'e'];
// copy to index 0 the element at index 3
console.log(array1.copyWithin(0, 3, 4));
// expected output: Array ["d", "b", "c", "d", "e"]
// copy to index 1 all elements from index 3 to the end
console.log(array1.copyWithin(1, 3));
// expected output: Array ["d", "d", "e", "d", "e"]
```

### 表单 accept 属性

表单 input type="file" 上传图片时，accept 属性以文件名结尾格式在部分手机上使用时会提示 '没有应用可执行此操作', 将文件名结尾的格式改为 MIME 类型的格式

```html
<input type="file" accept=".png,.jpg,.jpeg" />

<input type="file" accept="image/png,image/jpeg" />
```

### beforescriptexecute|afterscriptexecute <!-- markdownlint-disable-line -->

当 HTML 文档中的 script 标签内的代码执行`前|后`触发该事件, 如果这个 script 标签是用 `appendChild` 等方法动态插入的, 则不会触发该类事件

### IOS v-model 短信验证码自动填充两次

原因是ios系统bug，复制验证码会触发 UITextFieldTextDidChangeNotification 监听事件，导致验证码出现两次

- 如果 type 是 text 或者 password 时, 可以给 input 添加 maxlength 属性限制最大长度.
- 如果 type 是 number 时, 可以监听 @input 事件进行截字.

### [IE 10 开始不再支持条件注释引入资源](<https://learn.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/compatibility/hh801214(v=vs.85)?redirectedfrom=MSDN>)

```html
<!--[if lt IE 9]>
  <link rel="stylesheet" type="text/css" href="./main.css" />
  <script type="text/javascript" src="./main.js"></script>
<![endif]-->
```

### ArrayBuffer

ArrayBuffer 对象用来表示通用的、固定长度的原始二进制数据缓冲区, 可以理解为一个字节数组. 不能直接操作 ArrayBuffer, 需要通过 `类型化数组对象(TypedArray)`或 `DataView` 操作
ArrayBuffer 构造函数创建一个以字节为单位的固定长度的新 ArrayBuffer, 或者从现有的数据中获取数组缓冲区(例如: Base64 字符串或者 Blob 类文件对象 )

```javascript
var buffer = new ArrayBuffer(10); // 创建一个 10 字节的缓冲区
var i8a = new Int32Array(buffer); // 并使用 Int32Array 视图引用它
```

#### TypedArray

不能实例化
描述底层 `二进制数据缓冲区(ArrayBuffer)` 的类数组视图, 没有可用的 TypedArray 全局属性和 TypedArray 构造函数, 其为所有类型化数组的子类提供了实用方法的通用接口, 当创建 TypedArray 子类(例如 Int8Array) 的实例时, 在内存中会创建数组缓冲区, 如果将 ArrayBuffer 实例作为构造函数参数时, 则使用该 ArrayBuffer.

- Int8Array -128 到 127, 1 字节, 8 位有符号整型(补码)
- Uint8Array 0 到 255, 1 字节, 8 位无符号整型
- Uint8ClampedArray 0 到 255, 1 字节, 8 位无符号整型(一定在 0 - 255 之间)
- Int16Array -32768 到 32767, 2 字节, 16 位有符号整型(补码)
- Uint16Array 0 到 65535, 2 字节, 16 位无符号整型
- Int32Array -2147483648 到 2147483647, 4 字节, 32 位有符号整型(补码)
- Uint32Array 0 到 4294967295, 4 字节, 32 位无符号整型
- Float32Array -3.4E38 到 3.4E38 并且 1.2E-38 是最小的正数, 4 字节, 32 位 IEEE 浮点数(7 位有效数字，例如 1.234567)
- Float64Array -1.8E308 到 1.8E308 并且 5E-324 是最小的正数, 8 字节, 64 位 IEEE 浮点数(16 位有效数字，例如 1.23456789012345)
- BigInt64Array -263 到 263 - 1, 8 字节, 64 位有符号整型(补码)
- BigUint64Array 0 到 264 - 1, 8 字节, 64 位无符号整型

```javascript
var ia = new Int8Array(10);
ia[0] = 42;
```

#### DataView

DataView 是一个可以从二进制 `ArrayBuffer` 对象中读写多种数值类型的底层接口, 使用它时, 不需要考虑不同平台的字节序问题
DataView 构造函数可以传入一个已经存在的 `ArrayBuffer` 或 `SharedArrayBuffer` 作为数据源, 第二个参数可以指定 buffer 中的字节偏移, 第三个参数可以指定 DataView 对象的字节长度, 返回表示指定数据缓冲区的新 DataView 对象

```javascript
var buffer = new ArrayBuffer(16);
var view = new DataView(buffer);

view.setUint8(0, 42); // 设置指定偏移量的值
view.getUint8(0); // 获取指定偏移量的值

view.setInt32(1, 2147483647);
view.getInt32(1);
```

### [document.activeElement](https://developer.mozilla.org/zh-CN/docs/Web/API/Document/activeElement)

> 只读属性

用来返回当前在 DOM 或 shadow DOM 树中处于焦点状态 Element, 如 HTMLInputElement 或 HTMLTextAreaElement 元素中有文字被选中时, activeElement 属性就会返回该元素, 其它情况如 select 元素或者 input, textarea 元素

#### [Element.scrollIntoViewIfNeeded](https://developer.mozilla.org/zh-CN/docs/Web/API/Element/scrollIntoViewIfNeeded)

> 非标准化

用来将不在浏览器窗口的可见区域内的元素滚动到浏览器窗口的可见区域内, 如果该元素已经处在浏览器窗口的可见区域内, 则不会发生滚动, 此方法是 Element.scrollIntoView 方法的专有变体

- true, 默认值, 元素将在其所在滚动区的可视区域中居中对齐
- false, 元素将在其所在可视区域最近的边缘对齐(根据元素距离顶部边缘或者底部边缘的最小值对齐)

```javascript
var el = document.getElementById('child');
el.scrollIntoViewIfNeeded(true);
```

#### [Element.scrollIntoView](https://developer.mozilla.org/zh-CN/docs/Web/API/Element/scrollIntoView)

滚动元素的父容器, 使被调用 scrollIntoView 的元素对用户可见, 取决于其它元素的布局情况, 此元素可能不会完全滚动到顶端或底端

- alignToTop, 对齐方式, 可选
  - true, 元素的顶端将和其所在滚动区的可视区域的顶端对齐, 相应参数 `scrollIntoViewOptions: {block: "start", inline: "nearest"}`
  - false, 元素的底端将和其所在滚动区的可视区域的底端对齐, 相应参数 `scrollIntoViewOptions: {block: "end", inline: "nearest"}`
- scrollIntoViewOptions 配置项, 可选
  - behavior, 定义滚动是立即的还是平滑的动画
    - smooth, 滚动应该是平滑的动画
    - instant, 滚动应该通过一次跳跃立刻发生
    - auto, 滚动行为由 scroll-behavior 的计算值决定
  - block, 定义垂直方向的对齐方式, start, center, end, nearest, 默认为 start
  - inline, 定义水平方向的对齐, start, center, end, nearest, 默认为 nearest

```javascript
var el = document.getElementById('box');

el.scrollIntoView();
el.scrollIntoView(false);
el.scrollIntoView({ block: 'end' });
el.scrollIntoView({ behavior: 'smooth', block: 'end', inline: 'nearest' });
```

### eval

避免使用 eval, 可以使用 `Function('"use strict"; console.log("hello world")')()` 代替, `Function` 直接调用此构造函数可以动态创建函数

### [JSON.stringify](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify)

- 转换值如果有 toJSON 方法则直接使用该方法的返回值
- 非数组对象的属性不能保证以特定的顺序出现在序列化后的字符串中
- 布尔值、数字、字符串的包装对象在序列化过程中自动转换成对应的原始值
- undefined、任意的函数、symbol 值在序列化过程中会被忽略(出现在非数组对象的属性值中时, 出现在数组中时会被转换为 null), undefined、函数单独转换时被转换为 undefined
- 对包含循环引用的对象(对象之间相互引用)会抛出错误
- 所有以 symbol 为属性键的属性都会被完全忽略掉, 即使 replacer 参数中指定包含了它们
- Date 日期调用了 toJSON 方法将其转换为 string 字符串
- NaN 和 Infinity 格式的数值及 null 都会被当作 null
- 其它类型的对象, 包括 Map/WeakMap/Set/WeakSet, 仅会序列化可枚举的属性

### export default import

- export default 向外暴露的成员，可以使用任意变量来接收
- 在一个模块中, export default 只允许向外暴露一次
- 在一个模块中, 可以同时使用 export default 和 export 向外暴露成员
- 使用 export 向外暴露的成员, 只能使用 `{ }` 的形式来接收, 这种形式称为 按需导出
- export 可以向外暴露多个成员, 同时, 如果某些成员在 import 导入时不需要, 可以不在 `{ }` 中定义
- 使用 export 导出的成员, 必须严格按照导出时候的名称来使用 `{ }` 按需接收
- 如果想更改变量名称可以使用 as 定义别名

### [window.requestIdleCallback](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/requestIdleCallback)

插入一个函数, 这个函数将在浏览器空闲时期被调用. 使开发者能够在主事件循环上执行后台和低优先级工作, 而不会影响延迟关键事件

- callback 一个在事件循环空闲时即将被调用的函数的引用
- options 可选
  - timeout 指定回调函数未被调用的超时时间毫秒, 如果指定 timeout 回调函数超时后未被调用将被放入事件循环中排队, 即使有可能对性能产生负面影响

```javascript
window.requestIdleCallback(() => {
  console.log('requestIdleCallback')
}, {timeout: 10});

window.scheduler.postTack();
```

### [window.crypto](https://developer.mozilla.org/zh-CN/docs/Web/API/Crypto)

> 只读属性

返回当前窗口的作用域的 `Crypto` 对象, 此对象允许网页访问某些加密相关的服务

- Crypto.subtle 返回一个 [SubtleCrypto](https://developer.mozilla.org/zh-CN/docs/Web/API/SubtleCrypto) 对象, 用来访问公共的密码学原语, 例如哈希、签名、加密以及解密
- Crypto.getRandomValues(typedArray) 使用密码学安全的随机数填充传入的 TypedArray.
- Crypto.randomUUID() 返回一个随机生成的, 长度为 36 字符的 UUID

```javascript
var arr = new Uint32Array(10);
crypto.getRandomValues(arr); // 使用安全随机数填充 arrayBuffer.

crypto.randomUUID(); // 随机生成 UUID

crypto.subtle.encrypt(algorithm, key, data); // 加密数据
```
