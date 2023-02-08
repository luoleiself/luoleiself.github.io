---
title: JS小结
date: 2022-03-05 10:35:34
categories:
  - ES
tags:
  - ES6
  - js
---

### 改变原数组的方法

- pop 从数组中删除最后一个元素,并返回该元素的值(数组为空时返回 undefined). 此方法更改数组的长度

  ```javascript
  arr.pop();
  ```

- push 将一个或多个元素添加到数组的末尾, 并返回该数组的新长度

  ```javascript
  arr.push(element1, ..., elementN);
  ```

- shift 从数组中删除第一个元素,并返回该元素的值(数组为空则返回 undefined). 此方法更改数组的长度

  ```javascript
  arr.shift();
  ```

- unshift 将一个或多个元素添加到数组的开头. 并返回该数组的新长度(该方法修改原有数组)

  ```javascript
  arr.unshift(element1, ..., elementN);
  ```

  <!-- more -->

- reverse 将数组中元素的位置颠倒, 并返回该数组. 该方法会改变原数组

  ```javascript
  arr.reverse();
  ```

- sort 用[原地算法](https://en.wikipedia.org/wiki/In-place_algorithm)对数组的元素进行排序, 并返回数组

  - compareFunction 用来指定按某种顺序进行排列的函数, 省略则按照转换为的字符串的 unicode 位点进行排序
    - firstEl 第一个用于比较的元素
    - secondEl 第二个用于比较的元素

  ```javascript
  arr.sort([compareFunction]);
  ```

- splice 通过删除或替换现有元素或者原地添加新的元素来修改数组, 并以数组形式返回被修改的内容. 此方法会改变原数组

  - start 指定修改的开始位置
  - deleteCount 整数, 表示要移除的数组元素的个数, 如果为 0 或者负数, 则不移除元素
  - item1, item2 要添加进数组的元素,从 start 位置开始, 不指定则删除数组元素

  ```javascript
  array.splice(start[, deleteCount[, item1[, item2[, ...]]]]);
  ```

- fill 用一个固定值填充一个数组中从起始索引到终止索引内的全部元素. 不包括终止索引. 返回修改后的数组

  - value 用来填充数组元素的值
  - start 起始索引, 默认值为 0
  - end 终止索引, 默认值为 this.length

  ```javascript
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
  ```

- copyWithin 浅复制数组的一部分到同一数组中的另一个位置, 并返回它, 不会改变原数组的长度. 返回改变后的数组

  - target 整数, 复制序列到该位置, 如果是负数, target 将从末尾开始计算. 如果 target 大于等于 arr.length, 将会不发生拷贝
  - start 整数, 开始复制元素的起始位置, 如果是负数, start 将从末尾开始计算. 如果 start 被忽略, copyWithin 将会从 0 开始复制
  - end 整数, 开始复制元素的结束位置, copyWithin 将会拷贝到该位置, 但不包括 end 这个位置的元素. 如果是负数, end 将从末尾开始计算. 如果忽略则复制到数组结尾

  ```javascript
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

### [IE 10 开始不再支持条件注释引入资源](<https://learn.microsoft.com/en-us/previous-versions/windows/internet-explorer/ie-developer/compatibility/hh801214(v=vs.85)?redirectedfrom=MSDN>)

```html
<!--[if lt IE 9]>
  <link rel="stylesheet" type="text/css" href="./main.css" />
  <script type="text/javascript" src="./main.js"></script>
<![endif]-->
```

### eval

避免使用 eval, 可以使用 `Function('"use strict"; console.log("hello world")')()` 代替, `Function` 直接调用此构造函数可以动态创建函数

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

### Stream

Streams API 允许 JavaScript 以编程方式访问从网络接收的数据流, 并且允许根据需要处理它们

#### 可读流

##### ReadableStream

创建并从给定的 Handler 返回一个可读流对象

- ReadableStream.locked 只读属性, 返回该可读流是否被锁定到一个 reader
- ReadableStream.cancel(reason) 取消读取流, 可以传入 reason 参数表示取消原因, 该参数将回传给调用方
- ReadableStream.getReader(mode) 返回一个 `ReadableStreamDefaultReader` 实例并将流锁定到该实例, 一旦流被锁定, 其他 reader 不能读取直到释放
- ReadableStream.pipeThrough(transformStream, options) 提供当前流管道输出到一个 transform 流或 writable/readable 流对的链式方法
- ReadableStream.pipeTo(destination, options) 将当前 `ReadableStream` 管道输出到给定的 `WritableStream`, 并返回一个 Promise
- ReadableStream.tee() 返回包含两个 `ReadableStream` 实例分支的数组, 每个元素接收了相同的传输数据

```javascript
const queuingStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
var rs = new ReadableStream(
  {
    // 当创建实例时执行, 用于设置流功能, 可以返回一个 Promise
    start(controller) {
      interval = setInterval(() => {
        let string = randomChars();

        // Add the string to the stream
        controller.enqueue(string);

        // show it on the screen
        let listItem = document.createElement('li');
        listItem.textContent = string;
        list1.appendChild(listItem);
      }, 1000);

      button.addEventListener('click', function () {
        clearInterval(interval);
        fetchStream();
        controller.close();
      });
    },
    // // 当流的内部队列不满时, 会重复调用这个方法, 直到队列补满, 如果返回一个 promise, 此方法将不会再被调用
    // pull(controller) {
    //   // We don't really need a pull in this example
    // },
    // // 应用程序调用此方法将该流取消, 可以返回一个 Promise
    // cancel() {
    //   // This is called if the reader cancels,
    //   // so we should stop generating strings
    //   clearInterval(interval);
    // },
    // type: '', // 表示该流的类型
    // autoAllocateChunkSize: '', // 开启流自动分配缓冲区
  },
  queuingStrategy
);
```

##### ReadableStreamDefaultReader

表示一个用于读取来自网络提供的流数据(例如 fetch 请求)的默认读取器
`ReadableStreamDefaultReader` 可以用于读取底层为任意类型源的 `ReadableStream`(与 ReadableStreamBYOBReader 不同, 后者仅可以与底层为字节源的可读流一起使用)
构造方法创建并返回一个 `ReadableStreamDefaultReader` 实例, 通常不需要手动创建, 可以使用 `ReadableStream.getReader()` 方法代替

- ReadableStreamDefaultReader.closed 只读属性, 返回一个 Promise, 在流关闭时兑现
- ReadableStreamDefaultReader.cancel(reason) 返回一个 Promise, 当流被取消时兑现, 调用此方法取消流可传入 reason 参数表示取消原因
- ReadableStreamDefaultReader.read() 返回一个 Promise, 提供对流内部队列中下一个分块的访问权限 {value: theChunk, done: false} 表示可用, {value: undefined, done: true} 表示流已关闭
- ReadableStreamDefaultReader.releaseLock() 释放读取这个流的锁

##### ReadableStreamDefaultController

是一个控制器, 允许控制 `ReadableStream` 的状态和内部队列, 默认控制器用于不是字节流的流
无构造函数, `ReadableStreamDefaultController` 实例会在构造 `ReadableStream` 时被自动创建

- ReadableStreamDefaultController.desiredSize 只读属性, 返回填充满流的内部队列所需要的大小
- ReadableStreamDefaultController.close() 关闭关联的流
- ReadableStreamDefaultController.enqueue(chunk) 将给定的块加入关联的流
- ReadableStreamDefaultController.error(message) 导致未来任何与关联流的交互都会出错

#### 可写流

##### WritableStream

将流数据写入目的地(称为 sink) 提供了一个标准的抽象

- WritableStream.locked 只读属性, 返回可写流是否已锁定一个 writer
- WritableStream.abort(reason) 中止流, 表示不再向流中写入数据(立刻返回一个错误状态) 并丢球所有已入队的数据
- WritableStream.getWriter() 返回一个新的 `WritableStreamDefaultWriter` 实例并且将流锁定到该实例, 一旦流被锁定, 其他 writer 不能写入直到释放

```javascript
const decoder = new TextDecoder('utf-8');
const queuingStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
let result = '';
var writer = new WritableStream(
  {
    // 当创建实例时执行, 用于设置流功能, 可以返回一个 Promise
    start(controller) {},
    // 当一个新的 chunk 准备好写入 sink 时调用此方法, 可以返回一个 Promise 来表示写入操作的成功或者失败
    write(chunk, controller) {
      return new Promise((resolve, reject) => {
        var buffer = new ArrayBuffer(1);
        var view = new Uint8Array(buffer);
        view[0] = chunk;
        var decoded = decoder.decode(view, { stream: true });
        var listItem = document.createElement('li');
        listItem.textContent = 'Chunk decoded: ' + decoded;
        list.appendChild(listItem);
        result += decoded;
        resolve();
      });
    },
    // 当应用程序完成所有 chunk 的写入时执行此方法, 可以返回一个 Promise 表示操作成功或失败
    close(controller) {
      var listItem = document.createElement('li');
      listItem.textContent = '[MESSAGE RECEIVED] ' + result;
      list.appendChild(listItem);
    },
    // 立即关闭流并且丢弃所有入队数据时执行此方法, 可以返回一个 Promise 表示操作成功或失败
    abort(reason) {
      console.log('Sink error:', err);
    },
  },
  queuingStrategy
);
```

##### WritableStreamDefaultWriter

由 `WritableStream.getWriter()` 返回的对象, 并且一旦创建就会将 writer 锁定到 WritableStream, 确保没有其他流可以写入底层 sink
构造方法创建并返回一个 `WritableStreamDefaultWriter` 实例, 通常不需要手动创建, 可以使用 `WritableStream.getWriter()` 方法代替

- WritableStreamDefaultWriter.closed 只读属性, 返回一个 Promise, 在流关闭时兑现
- WritableStreamDefaultWriter.desiredSize 只读属性, 返回填充满流的内部队列所需要的大小
- WritableStreamDefaultWriter.ready 只读属性, 返回一个 Promise, 当流填充内部队列的所需大小从非正数变为正数时兑现
- WritableStreamDefaultWriter.abort(reason) 中止流, 返回一个 Promise, 表示生产者不能再向流写入数据(会立刻返回一个错误状态), 并丢弃所有已入队了数据
- WritableStreamDefaultWriter.close() 关闭关联的可写流, 返回一个 Promise, 如果所有剩余的分块在关闭之前成功写入, 则使用 undefined 兑现, 如果遇到问题则拒绝并返回相关错误
- WritableStreamDefaultWriter.releaseLock() 释放可写流的锁
- WritableStreamDefaultWriter.write(chunk) 将传递的数据写入 WritableStream 和它的底层 sink, 返回一个 Promise, 成功写入后使用 undefined 兑现, 如果遇到问题则拒绝并返回相关错误

##### WritableStreamDefaultController

是一个控制器, 允许控制 `WritableStream` 状态的控制器, 当构造 `WritableStream` 时, 会为底层的 sink 提供一个相应的`WritableStreamDefaultController`实例进行操作
无构造函数, `WritableStreamDefaultController` 实例会在构造 `WritableStream` 时被自动创建

- WritableStreamDefaultController.error(message) 导致未来任何与关联流的交互都会出错

#### 转换流

##### TransformStream

TransformStream 接口表示链式管道传输转换流概念的具体实现, 可以传给 `ReadableStream.pipeThrough()` 方法将流数据从一种格式转换成另一种, 例如, 可以用于解码(编码)视频帧, 解压数据或者将流从 XML 转换成 JSON

- TransformStream.readable 只读属性, 转换流的 readable 端
- TransformStream.writable 只读属性, 转换流的 writable 端

```javascript
var writableStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
var readableStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
var ts = new TransformStream({
  // 当创建实例时执行, 通常用于使用 enqueue 对分块进行排队
  start(controller) {},
  // 当一个写入可写端的分块准备好转换时调用，并且执行转换流的工作, 如果不提供此方法则使用恒等变换并且分块将在没有更改的情况下排队
  transform(chunk, controller) {},
  // 当所有写入可写端的分块成功转换后被调用, 并且可写端将会关闭
  flush(controller) {},
  writableStrategy,
  readableStrategy,
});
```

##### TransformStreamDefaultController

提供操作关联的 `ReadableStream` 和 `WritableStream` 的方法
无构造函数, `TransformStreamDefaultController` 实例会在构造 `TransformStream` 时被自动创建, 通过 `TransformStream` 的回调函数获取

- TransformStreamDefaultController.desiredSize 只读属性, 返回填充满内部队列的可读端所需要的大小
- TransformStreamDefaultController.enqueue(chunk) 将给定的 chunk 加入流的可读端
- TransformStreamDefaultController.error(reason) 导致转换流的可读端和可写端都会出错
- TransformStreamDefaultController.terminate() 关闭流的可读端并且流的可写端出错

#### TextEncoder

接受码位流作为输入并提供 UTF-8 字节流作为输出

- TextEncoder.prototype.encoding 只读属性, 总是返回 utf-8
- TextEncoder.encode(string) 接受一个字符串输入并返回一个 UTF-8 编码的文本的 `Uint8Array`
- TextEncoder.encodeInto(string, Uint8Array) 接受一个字符串和一个目标(Uint8Array 用于存放 UTF-8 编码的文本), 并且返回一个只是编码进度的对象, 此方法性能会比 encode 好一些

```javascript
var te = new TextEncoder();
var u8arr = te.encode('hello world'); // Uint8Array [104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100]
```

#### TextDecoder

接收一个文本编码, 将字节流作为输入并提供码位流作为输出

- TextDecoder.prototype.encoding 只读属性, 表示将使用的编码格式
- TextDecoder.prototype.fatal 只读属性, 表示错误模式是否致命
- TextDecoder.prototype.ignoreBOM 只读属性, 表示是否忽略字节顺序标记(BOM)
- TextDecoder.prototype.decode(buffer, {stream: false}) 返回一个使用指定编码格式解码的字符串
  - buffer 一个 ArrayBuffer, TypedArray 或包含要解码的编码文本的 DataView 对象
  - stream 默认 false 不使用分块方式, true 表示以分块方式处理数据,后续调用 decode 将跟随附加数据

```javascript
// 第一个参数表示编码, 默认 utf-8, 第二个参数的 fatal 属性表示在解码无效数据时是否必须抛出 TypeError, 默认 false
var td = new TextDecoder('utf-8', { fatal: true });
td.decode(u8arr); // hello world
```

#### TextEncoderStream

将一个字符串流转换为 UTF-8 编码的字节, 与 `TextEncoder` 的流形式等价

- TextEncoderStream.encoding 只读属性, 总是返回 utf-8
- TextEncoderStream.readable 只读属性, 返回此对象控制的 `ReadableStream` 实例
- TextEncoderStream.writable 只读属性, 返回此对象控制的 `WritableStream` 实例

```javascript
var tes = new TextEncoderStream();
console.log(tes.readable);
```

#### TextDecoderStream

将二进制编码的文本流转换字符串流, 与 `TextDecoder` 的流形式等价

- TextDecoderStream.encoding 只读属性, 表示将使用的编码格式
- TextDecoderStream.fatal 只读属性, 表示错误是否致命
- TextDecoderStream.ignoreBOM 只读属性, 表示是否忽略字节顺序标记
- TextDecoderStream.readable 只读属性, 返回此对象控制的 `ReadableStream` 实例
- TextDecoderStream.writable 只读属性, 返回此对象控制的 `WritableStream` 实例

```javascript
var tds = new TextDecoderStream('utf-8', { fatal: false });
console.log(tds.writable);
```
