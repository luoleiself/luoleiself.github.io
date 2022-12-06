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

### ArrayBuffer

ArrayBuffer 对象用来表示通用的、固定长度的原始二进制数据缓冲区, 可以理解为一个字节数组. 不能直接操作 ArrayBuffer, 需要通过 `类型化数组对象(TypedArray)`或 `DataView` 操作
ArrayBuffer 构造函数创建一个以字节为单位的固定长度的新 ArrayBuffer, 或者从现有的数据中获取数组缓冲区(例如: Base64 字符串或者 Blob 类文件对象 )

```javascript
var buffer = new ArrayBuffer(10); // 创建一个 10 字节的缓冲区
var i8a = new Int32Array(buffer); // 并使用 Int32Array 视图引用它
```

#### TypedArray

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
- ReadableStream.cancel() 取消读取流, 可以传入 reason 参数表示取消原因, 该参数将回传给调用方
- ReadableStream.getReader() 创建一个读取器并将流锁定其上, 一旦流被锁定, 其他读取器将不能读取它直到释放
- ReadableStream.pipeThrough() 提供当前流管道输出到一个 transform 流或 writable/readable 流对的链式方法
- ReadableStream.pipeTo() 将当前 `ReadableStream` 管道输出到给定的 `WritableStream`, 并返回一个 Promise
- ReadableStream.tee() 返回包含两个 `ReadableStream` 实例分支的数组, 每个元素接收了相同的传输数据

```javascript
var rs = new ReadableStream(
  {
    // 当每个对象被构造时立刻调用的方法, 可以返回一个 Promise
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
  }
  // {
  //   highWaterMark: 1, // 非负整数, 定义了在应用程序之前可以包含在内部队列中的块的总数
  //   size(chunk) {}, // 表示每个分块使用的大小(以字节为单位)
  // }
);
```

##### ReadableStreamDefaultReader

表示一个用于读取来自网络提供的流数据(例如 fetch 请求)的默认读取器
`ReadableStreamDefaultReader` 可以用于读取底层为任意类型源的 `ReadableStream`(与 ReadableStreamBYOBReader 不同, 后者仅可以与底层为字节源的可读流一起使用)
构造方法创建并返回一个 `ReadableStreamDefaultReader` 实例, 通常不需要手动创建, 可以使用 `ReadableStream.getReader()` 方法代替

- ReadableStreamDefaultReader.closed 返回一个 Promise, 在流关闭时兑现
- ReadableStreamDefaultReader.cancel() 返回一个 Promise, 当流被取消时兑现, 调用此方法取消流可传入 reason 参数表示取消原因
- ReadableStreamDefaultReader.read() 返回一个 Promise, 提供对流内部队列中下一个分块的访问权限
- ReadableStreamDefaultReader.releaseLock() 释放读取这个流的锁

##### ReadableStreamDefaultController

是一个控制器, 允许控制 `ReadableStream` 的状态和内部队列, 默认控制器用于不是字节流的流
无构造函数, `ReadableStreamDefaultController` 实例会在构造 `ReadableStream` 时被自动创建

- ReadableStreamDefaultController.desiredSize 只读属性, 返回填充满流的内部队列所需要的大小
- ReadableStreamDefaultController.close() 关闭关联的流
- ReadableStreamDefaultController.enqueue() 将给定的块加入关联的流
- ReadableStreamDefaultController.error() 导致未来任何与关联流的交互都会出错

#### 可写流

##### WritableStream

##### WritableStreamDefaultWriter

##### WritableStreamDefaultController

是一个控制器, 允许控制 `WritableStream` 状态的控制器, 当构造 `WritableStream` 时, 会为底层的 sink 提供一个相应的`WritableStreamDefaultController`实例进行操作
无构造函数, `WritableStreamDefaultController` 实例会在构造 `WritableStream` 时被自动创建

- WritableStreamDefaultController.error() 导致未来任何与关联流的交互都会出错

#### 转换流

##### TransformStream
