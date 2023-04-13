---
title: Stream
date: 2023-04-12 15:52:05
categories:
  - WebAPI
tags:
  - API
---

## 可读流

### ReadableStream

创建并从给定的 Handler 返回一个可读流对象

#### RS 构造方法

创建并从给定的处理程序返回一个可读的流对象

##### 可读流配置项

`underlyingSource` 可选的定义可读流的行为方式的配置项

- start(controller) 由开发人员定义, 当流对象被创建时立刻调用, 执行其他任何必须的设置流功能, 如果过程是异步的则返回一个 Promise
  - controller 根据 `type='bytes'` 属性传递的 `ReadableStreamDefaultController` 或 `ReadableByteStreamController` 控制器实例
- pull(controller) 由开发人员定义, 当流的内部队列不满时, 会重复调用这个方法, 直到队列补满
- cancel(reason) 由开发人员定义, 当该流被取消时调用
- type 控制正在处理的可读类型的流, 默认 `ReadableStreamDefaultController`, `type='bytes'` 表示 `ReadableByteStreamController`
- autoAllocateChunkSize 开启流自动分配缓冲区, 使用正整数设置

<!-- more -->

##### 可读流队列策略配置项

`queuingStrategy` 可选的为流定义排队策略的配置项

- highWaterMark 非负整数, 定义应用在背压之前可以包含在内部队列中的块的总数
- size(chunk) 表示每个分块使用的大小(以字节为单位)

```javascript
const queuingStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
const rs = new ReadableStream(
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

#### RS 实例属性

- locked 只读属性, 返回该可读流是否被锁定到一个 reader

#### RS 实例方法

##### RS.cancel(reason)

取消读取流, 可以传入 reason 参数表示取消原因, 该参数将回传给调用方

##### RS.getReader(mode)

将流锁定到该实例, 一旦流被锁定, 其他 reader 不能读取直到释放

- 默认返回一个 `ReadableStreamDefaultReader` 实例
- `{ mode: 'byob' }` 返回一个 `ReadableStreamBYOBReader` 实例

```javascript
// 获取 ReadableStreamDefaultReader 实例
const reader = rs.getReader();

// 获取 ReadableStreamBYOBReader 实例
const reader = rs.getReader({ mode: 'byob' });
```

##### RS.pipeThrough(transformStream, options)

提供当前流管道输出到一个 transform 流或 writable/readable 流对的链式方法

```javascript
fetch('png-logo.png')
  .then((response) => response.body)
  .then((body) => body.pipeThrough(new PNGTransformStream()))
  .then((rs) => rs.pipeTo(new FinalDestinationStream()));
```

##### RS.pipeTo(destination, options)

将当前 `ReadableStream` 管道输出到给定的 `WritableStream`, 并返回一个 Promise

##### RS.tee()

返回包含两个 `ReadableStream` 实例分支的数组, 每个元素接收了相同的传输数据

### 默认读取流

> 通常不需要手动创建, 可以使用 `ReadableStream.getReader()` 方法获取

`ReadableStreamDefaultReader` 表示一个用于读取来自网络提供的流数据(例如 fetch 请求)的**默认读取器**

`ReadableStreamDefaultReader` 可以用于读取底层为任意类型源的 `ReadableStream`(与 ReadableStreamBYOBReader 不同, 后者仅可以与底层为字节源的可读流一起使用)

#### RSDR 构造方法

构造方法创建并返回一个 `ReadableStreamDefaultReader` 实例, 通常不需要手动创建

```javascript
// 参数 stream 为将被读取的 ReadableStream;
const rsdr = new ReadableStreamDefaultReader(stream);
```

#### RSDR 实例属性

- closed 只读属性, 返回一个 Promise, 在流关闭时兑现

#### RSDR 实例方法

##### RSDR.cancel(reason)

返回一个 Promise, 当流被取消时兑现, 调用此方法取消流可传入 reason 参数表示取消原因

##### RSDR.read()

返回一个 Promise, 提供对流内部队列中下一个分块的访问权限 {value: theChunk, done: false} 表示可用, {value: undefined, done: true} 表示流已关闭

```javascript
const rsdr = rs.getReader();
rsdr.read().then(function ({ done, value }) {
  if (done) {
    console.log('rsdr complete...');
    // 取消流
    rsdr.cancel();
    return;
  }
  console.log(value);
});

// 释放读取流的锁
rsdr.releaseLock();
```

##### RSDR.releaseLock()

释放读取这个流的锁

### 字节读取流

> 通常不需要手动创建，使用 `ReadableStream.getReader({ mode: 'byob' })` 方法获取

`ReadableStreamBYOBReader` 表示一个支持从底层字节源进行零拷贝读取的读取器, 用于从底层源进行高效复制, 其中数据以匿名字节序列的形式传递

#### RSBYOBR 构造方法

构造方法创建并返回一个 `ReadableStreamBYOBReader` 实例, 通常不需要手动创建

```javascript
// 参数 stream 为将被读取的 ReadableStream
const rsbyobr = new ReadableStreamBYOBReader(stream);
```

#### RSBYOBR 实例属性

- closed 只读属性, 返回一个 Promise, 在流关闭时兑现

#### RSBYOBR 实例方法

##### RSBYOBR.cancel(reason)

返回一个 Promise, 当流被取消时兑现, 调用此方法取消流可传入 reason 参数表示取消原因

##### RSBYOBR.read(view)

传递一个必须写入数据的视图, 并返回一个 Promise 解析为流中的下一个块的视图或者表示流已关闭或出错的指示的视图

```javascript
const rsbyobr = rs.getReader({ mode: 'byob' });
let buffer = new Uint8Array(new ArrayBuffer(4000));
rsbyobr.read(buffer).then(function ({ done, value }) {
  if (done) {
    console.log('rsbyobr complete...');
    // 取消流
    rsbyobr.cancel();
    return;
  }
  console.log(value);
});

// 释放读取流的锁
rsbyobr.releaseLock();
```

##### RSBYOBR.releaseLock()

释放读取这个流的锁

### 默认读取流控制器

> **无构造函数**, `ReadableStreamDefaultController` 实例会在构造 `ReadableStream` 时被自动创建

`ReadableStreamDefaultController` 接口是一个控制器, 允许控制 `ReadableStream` 的状态和内部队列, 默认控制器用于不是字节流的流

#### RSDC 实例属性

- desiredSize 只读属性, 返回填充满流的内部队列所需要的大小

#### RSDC 实例方法

##### RSDC.close()

关闭关联的流

##### RSDC.enqueue(chunk)

将给定的块加入关联的流

```javascript
// 按下按钮时, 关闭流并运行另一个函数从流中读取数据
let interval;
const rs = new ReadableStream({
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

    button.addEventListener('click', () => {
      clearInterval(interval);
      fetchStream();
      controller.close();
    });
  },
  pull(controller) {},
  cancel() {
    clearInterval(interval);
  },
});
```

##### RSDC.error(message)

导致未来任何与关联流的交互都会出错

### 字节读取流控制器

> **无构造函数**, 当构造 `ReadableStream` 时, `type="bytes"` 将自动创建 `ReadableByteStreamController` 实例

`ReadableByteStreamController` 接口是一个控制器, 允许控制具有底层字节源的状态和内部队列, 并在流的内部队列为空时实现从底层源到消费者的高效零拷贝数据传输

#### RBSC 实例属性

- byobRequest 返回当前的 `BYOB` 拉取请求, 如果没有未处理的请求则返回 null
- desiredSize 返回填充流内部队列所需的大小

#### RBSC 实例方法

##### RBSC.close()

关闭关联的流

##### RBSC.enqueue(chunk)

将给定的块加入关联的流

##### RBSC.error(message)

导致未来任何与关联流的交互都会出错

## 可写流

### WritableStream

将流数据写入目的地(称为 sink) 提供了一个标准的抽象

#### WS 构造方法

创建一个新的 `WritableStream` 实例

##### 可写流配置项

`underlyingSource` 可选的定义可读流的行为方式的配置项

- start(controller) 由开发人员定义, 当流对象被创建时立刻调用, 执行其他任何必须的设置流功能, 如果过程是异步的则返回一个 Promise
  - controller 传递的 `WritableStreamDefaultController` 控制器实例
- write(chunk, controller) 由开发人员定义, 当新数据块准备好写入底层接收器时将被调用
  - chunk 指定的数据块
- close(controller) 由开发人员定义, 如果应用发出信号表明它已完成块写入流, 则将调用此方法
- abort(reason) 由开发人员定义, 如果应用发出信号表示它希望突然关闭流并将其置于错误状态, 则将调用此方法

##### 可写流队列策略配置项

`queuingStrategy` 可选的为流定义排队策略的配置项

- highWaterMark 非负整数, 定义应用在背压之前可以包含在内部队列中的块的总数
- size(chunk) 表示每个分块使用的大小(以字节为单位)

```javascript
const queuingStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
const ws = new WritableStream({
  start(controller) {},
  write(chunk, controller) {},
  close(controller) {},
  abort(reason) {
    console.log('error ', reason);
  },
  queuingStrategy,
});
```

#### WS 实例属性

- locked 只读属性, 返回可写流是否已锁定一个 writer

#### WS 实例方法

##### WS.close()

关闭关联的流

##### WS.abort(reason)

中止流, 表示不再向流中写入数据(立刻返回一个错误状态) 并丢球所有已入队的数据

##### WS.getWriter()

返回一个新的 `WritableStreamDefaultWriter` 实例并且将流锁定到该实例, 一旦流被锁定, 其他 writer 不能写入直到释放

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

### 默认写入流

> 通常不需要手动创建, 可以使用 `WritableStream.getWriter()` 方法获取

`WritableStreamDefaultWriter` 一旦创建就会锁定 Writer, 以确保没有其他流可以写入底层接收器

#### WSDW 构造方法

构造方法创建并返回一个 `WritableStreamDefaultWriter` 实例, 通常不需要手动创建

```javascript
// 参数 stream 为将被写入的 WritableStream
const wsdw = new WritableStreamDefaultWriter(stream);
```

#### WSDW 实例属性

- closed 只读属性, 返回一个 Promise, 在流关闭时兑现
- desiredSize 只读属性, 返回填充满流的内部队列所需要的大小
- ready 只读属性, 返回一个 Promise, 当流填充内部队列的所需大小从非正数变为正数时兑现

#### WSDW 实例方法

##### WSDW.abort(reason)

中止流, 返回一个 Promise, 表示生产者不能再向流写入数据(会立刻返回一个错误状态), 并丢弃所有已入队了数据

##### WSDW.close()

关闭关联的可写流, 返回一个 Promise, 如果所有剩余的分块在关闭之前成功写入, 则使用 undefined 兑现, 如果遇到问题则拒绝并返回相关错误

##### WSDW.releaseLock()

释放可写流的锁

##### WSDW.write(chunk)

将传递的数据写入 WritableStream 和它的底层 sink, 返回一个 Promise, 成功写入后使用 undefined 兑现, 如果遇到问题则拒绝并返回相关错误

```javascript
const wsdw = ws.getWriter();
const encoder = new TextEncoder();
const encoded = encoder.encode('hello world', { stream: true });
encoded.forEach((chunk) => {
  wsdw.ready
    .then(() => wsdw.write(chunk))
    .then(() => {
      console.log('chunk written is to sink.');
    })
    .catch((err) => {
      console.log('chunk error ', err);
    });
});
// Call ready again to ensure that all chunks are written
// before closing the writer.
wsdw.ready
  .then(() => {
    wsdw.close();
  })
  .then(() => {
    console.log('All chunks written');
  })
  .catch((err) => {
    console.log('Stream error ', err);
  });
```

### 默认写入流控制器

> **无构造函数**, `WritableStreamDefaultController` 实例会在构造 `WritableStream` 时被自动创建

`WritableStreamDefaultController` 接口是一个控制器, 允许控制 `WritableStream` 状态的控制器, 当构造 `WritableStream` 时, 会为底层的 sink 提供一个相应的`WritableStreamDefaultController`实例进行操作

#### WSDC 实例属性

- signal 返回与 `AbortSignal` 关联的控制器

#### WSDC 实例方法

##### WSDC.error(message) 导致未来任何与关联流的交互都会出错

```javascript
const ws = new WritableStream({
  start(controller) {
    controller.error('Stream is broken');
  },
  write(chunk, controller) {},
  close(controller) {},
  abort(err) {},
});
```

## 转换流

### TransformStream

`TransformStream` 接口表示链式管道传输转换流概念的具体实现, 可以传给 `ReadableStream.pipeThrough()` 方法将流数据从一种格式转换成另一种, 例如, 可以用于解码(编码)视频帧, 解压数据或者将流从 XML 转换成 JSON

#### TS 构造方法

创建并返回一个转换流对象, 可以选择为流指定一个转换对象和排队策略

##### 转换流配置项

transformer 可选的表示转换流的对象, 如果未提供, 则生成的流将是一个恒等交换流, 它将所有写入可写端的分块转发到可读端, 不会有任何该表

- start(controller) 当 `TransformStream` 创建时被调用, 通常用于使用 `TransformStreamDefaultController.enqueue()` 对分块进行排队
- transform(chunk, controller) 当一个写入可写端的分块准备好转换时调用, 并执行转换流的工作, 如果没有提供则使用恒等交换
- flush(controller) 当所有写入可写端的分块成功转换后被调用, 并且可写端将会关闭

##### 写入流队列策略配置项

`writableStrategy` 可选的定义写入流队列策略的配置项

- highWaterMark 非负整数, 定义应用在背压之前可以包含在内部队列中的块的总数
- size(chunk) 表示每个分块使用的大小(以字节为单位)

##### 读取流队列策略配置项

`readableStrategy` 可选的定义读取流队列策略的配置项

- highWaterMark 非负整数, 定义应用在背压之前可以包含在内部队列中的块的总数
- size(chunk) 表示每个分块使用的大小(以字节为单位)

```javascript
const writableStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
const readableStrategy = new CountQueuingStrategy({ highWaterMark: 1 });
const ts = new TransformStream({
  // 当创建实例时执行, 通常用于使用 enqueue 对分块进行排队
  start(controller) {},
  // 当一个写入可写端的分块准备好转换时调用，并且执行转换流的工作
  // 如果不提供此方法则使用恒等变换并且分块将在没有更改的情况下排队
  transform(chunk, controller) {},
  // 当所有写入可写端的分块成功转换后被调用, 并且可写端将会关闭
  flush(controller) {},
  writableStrategy,
  readableStrategy,
});
```

#### TS 实例属性

- readable 只读属性, 转换流的 readable 端
- writable 只读属性, 转换流的 writable 端

```javascript
const writableStrategy = new ByteLengthQueuingStrategy({
  highWaterMark: 1024 * 1024,
});
const el = document.body;
const ts = new TransformStream(
  {
    transform(chunk, controller) {
      controller.enqueue(chunk.toUpperCase());
    },
  },
  writableStrategy
);
const ws = new WritableStream({
  write(chunk, controller) {
    el.append(chunk);
  },
});

fetch('./lorem-ipsum.txt').then((response) =>
  response.body.pipeThrough(ts).pipeTo(ws)
);
```

### 默认转换流控制器

> **无构造函数**, `TransformStreamDefaultController` 实例会在构造 `TransformStream` 时被自动创建

`TransformStreamDefaultController` 接口提供操作关联的 `ReadableStream` 和 `WritableStream` 的方法

#### TSDC 实例属性

- desiredSize 只读属性, 返回填充满内部队列的可读端所需要的大小

#### TSDC 实例方法

##### TSDC.enqueue(chunk)

将给定的 chunk 加入流的可读端

##### TSDC.error(reason)

导致转换流的可读端和可写端都会出错

##### TSDC.terminate()

关闭流的可读端并且流的可写端出错

```javascript
const ts = new TransformStream({
  transform(chunk, controller) {
    controller.enqueue(new TextEncoder().encode(chunk));
  },
  flush(controller) {
    controller.terminate();
  },
});
```

## TextEncoder

接受码位流作为输入并提供 UTF-8 字节流作为输出

### TE 构造函数

创建并返回一个新的 `TextEncoder` 实例, 该实例将生成具有 `UTF-8` 编码的字节流

```javascript
const te = new TextEncoder();
```

### TE 实例属性

- encoding 只读属性, 总是返回 utf-8

### TE 实例方法

#### TE.encode(string)

接受一个字符串输入并返回一个 `UTF-8` 编码的文本的 `Uint8Array`

```javascript
const te = new TextEncoder();
const u8arr = te.encode('hello world'); // Uint8Array [104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100]
```

#### TE.encodeInto(string, Uint8Array)

接受一个字符串和一个目标(Uint8Array 用于存放 UTF-8 编码的文本), 并且返回一个只是编码进度的对象, 此方法性能会比 encode() 更高

```javascript
const te = new TextEncoder();
const uint8Arr = new Unit8Array(8);
te.encodeInto('hello world', uint8Arr);
console.log(uint8Arr.join());
```

## TextDecoder

`TextDecoder` 接口表示一个文本解码器, 一个解码器只支持一种特定文本编码, 解码器将字节流作为输入并提供码位流作为输出

### TD 构造函数

根据参数指定的编码创建并返回一个新的 `TextDecoder` 实例

- utfLabel 可选的一个字符串, 默认 `utf-8`, 可以为任意有效的编码
- options 可选的配置项
  - fatal 布尔值, 表示在解码无效数据时, `decode()` 方法是否必须抛出 `TypeError`, 默认 false

```javascript
const td1 = new TextDecoder('iso-8859-2');
// Allow TypeError exception to be thrown
const td2 = new TextDecoder('csiso2022kr', { fatal: true });
```

### TD 实例属性

- encoding 只读属性, 表示将使用的编码格式
- fatal 只读属性, 表示错误模式是否致命
- ignoreBOM 只读属性, 表示是否忽略字节顺序标记(BOM)

### TD 实例方法

#### TD.decode(buffer, {stream: false})

返回一个使用指定编码格式解码的字符串

- buffer 可选的一个 ArrayBuffer, TypedArray 或包含要解码的编码文本的 DataView 对象
- options 可选的配置项
  - stream 默认 false 不使用分块方式, true 表示以分块方式处理数据,后续调用 decode 将跟随附加数据

```javascript
const te = new TextEncoder();
const array = te.encode('hello world');
// Unit8Array(6) [228,189,160,229,165,189]
document.getElementById('encode-value').textContent = array;

const td = new TextDecoder();
const str = td.decode(array);
// 你好
document.getElementById('decode-value').textContent = str;
```

## TextEncoderStream

将一个字符串流转换为 `UTF-8` 编码的字节, 与 `TextEncoder` 的流形式等价

### TES 构造函数

创建并返回一个新的 `TextEncoderStream` 实例, 该对象使用 `UTF-8` 编码将字符串流转换为字节

```javascript
const tes = new TextEncoderStream();
```

### TES 实例属性

- encoding 只读属性, 总是返回 utf-8
- readable 只读属性, 返回此对象控制的 `ReadableStream` 实例
- writable 只读属性, 返回此对象控制的 `WritableStream` 实例

```javascript
var tes = new TextEncoderStream();
console.log(tes.readable);
```

## TextDecoderStream

将二进制编码的文本流转换字符串流, 与 `TextDecoder` 的流形式等价

### TDS 构造函数

- label 可选的一个字符串, 默认 `utf-8`, 可以为任意有效的编码
- options 可选的配置项
  - fatal 布尔值, 表示在错误的模式, 如果为 true, decoder 则在遇到错误时抛出一个 DOMException, 默认为 false

```javascript
const response = await fetch('https://example.com');
const stream = response.body.pipeThrough(new TextDecoderStream());
```

### TDS 实例属性

- encoding 只读属性, 表示将使用的编码格式
- fatal 只读属性, 表示错误是否致命
- ignoreBOM 只读属性, 表示是否忽略字节顺序标记
- readable 只读属性, 返回此对象控制的 `ReadableStream` 实例
- writable 只读属性, 返回此对象控制的 `WritableStream` 实例

```javascript
var tds = new TextDecoderStream('utf-8', { fatal: true });
console.log(tds.writable);
```
