---
title: FileSystemAPI
date: 2023-04-12 14:02:45
categories:
  - WebAPI
tags:
  - API
---

> 目前仅在 chrome 86 (edge 86, opera 72)及以上版本支持, safari 和 firefox 暂时不支持

允许 Web 应用程序从用户设备的本地文件系统中操作文件, 它为 Web 应用程序提供了更多的灵活性和功能, 使其更接近于本地应用程序的体验

[File System API](https://developer.mozilla.org/zh-CN/docs/Web/API/File_System_API) 遵循同源策略, 只允许 Web 应用程序在具有相同源的文件系统上进行操作, 当使用该 API 时, 会提示用户授权应用程序访问文件系统

- 将文件从本地文件系统上传到 Web 应用程序
- 将 Web 应用程中的数据写入到本地文件系统中
- 在用户的本地文件系统中创建、重命名和删除文件
- 读取本地文件系统上的文件内容

```javascript
/**
 * 继承
 * FileSystemHandle
 *  <- FileSystemFileHandle
 *  <- FileSystemDirectoryHandle
 * FileSystemSyncAccessHandle
 * WritableStream
 *  <- FileSystemWritableFileStream
 */
```

### [FileSystemHandle](https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle)接口

是 File System API 表示文件或目录条目的对象, 多个句柄可以代表同一个条目，通常情况下, 使用它的子接口 `FileSystemFileHandle` 和 `FileSystemDirectoryHandle`

- FileSystemFileHandle
- [FileSystemDirectoryHandle](#FileSystemDirectoryHandle)

<!-- more -->

#### FSH 实例属性

- kind 返回条目的类型, `file` 表示文件, `directory` 表示目录
- name 返回关联条目的名称

#### FSH 实例方法

- isSameEntry() 比较两者 handle 以查看相关条目(文件或目录)是否匹配
- queryPermission() 查询当前句柄的当前权限状态
- remove() 请求从底层文件系统中删除由句柄表示的条目
- requestPermission() 请求文件句柄的读取或读写权限

#### FileSystemFileHandle <em id="FileSystemFileHandle"></em> <!--markdownlint-disable-line-->

表示一个指向文件系统条目的句柄

- [window.showOpenFilePicker](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/showOpenFilePicker) 显示一个文件选择器, 允许用户选择一个或多个文件并返回这些文件的句柄, 返回一个 Promise 对象, 并兑现一个包含 `FileSystemFileHandle` 对象的 Array 数组 <em id="showOpenFilePicker"></em> <!--markdownlint-disable-line-->
  - options 可选
    - excludeAcceptAllOption 默认 false, 选择器应包含一个不应用任何文件类型过滤器的选项, 设为 true 表示该选项不可用
    - id 指定 id, 浏览器可以为不同的 id 记住不同的目录
    - multiple 默认 false, 是否支持选择多个文件
    - startIn 一个 FileSystemHandle 对象或一个已知的(desktop, documents, downloads, music, pcitures, videos)目录, 以指定打开选择器的起始目录
    - types 允许选择文件的对象数组
      - description 文件类型的可选描述
      - accept 键为 MIME 类型, 值为文件扩展名的数组

```javascript
(async function () {
  // 可选参数, 传入一个文件选择器筛选文件的选项对象
  // 返回一个已兑现的包含 `FileSystemFileHandle` 对象的 Array 数组
  // FileSystemHandle <- FileSystemFileHandle
  const fileHandles = await window.showOpenFilePicker({
    excludeAcceptAllOption: true,
    multiple: false,
    types: [
      {
        description: 'Images', 
        'image/**': ['.png', '.jpg', '.gif']
      }
    ]
  });

  // 获取返回结果中第一个文件对象
  const file = await fileHandles[0].getFile();

  // 获取返回结果中第一个可用于同步读取和写入文件的对象,
  // 但只能用在专用的 web workers 中
  const syncAccessFile = await fileHandles[0].createSyncAccessHandle();

  // 获取返回结果中第一个可用于写入文件的新创建的对象
  // WritableStream <- FileSystemWritableFileStream
  // 继承 writableStream 接口的 `FileSystemWritableFileStream` 实例
  // 在流关闭之前, 通常将对流做的任何修改写入临时文件来实现,
  // 并且仅在可写文件流关闭时才将文件句柄表示的文件替换为临时文件
  const writer = await fileHandles[0].createWritable();
})();
```

- [window.showSaveFilePicker](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/showSaveFilePicker) 显示一个允许用户保存文件的文件选择器, 通过选择现有文件或者输入新文件的名称, 返回一个已兑现的 `FileSystemFileHandle` 对象

  - options 参数同 [window.showOpenFilePicker](#showOpenFilePicker)

```javascript
(async function () {
  // 可选参数, 传入一个文件选择器筛选文件的选项对象
  // 返回一个已兑现的 `FileSystemFileHandle` 对象
  const fileHandle = await window.showSaveFilePicker({
    /* ... */
  });

  // 创建一个可写流
  // 继承 writableStream 接口的 `FileSystemWritableFileStream` 实例
  const writable = await fileHandle.createWritable();

  // 写入数据
  await writable.write('Hello world!');

  // 关闭流
  await writable.close();
})();
```

##### getFile()

返回一个已兑现的由句柄表示的条目在磁盘上的状态的对象

##### createSyncAccessHandle()

返回一个已兑现的可用于同步读取和写入文件的 [FileSystemSyncAccessHandle](#FileSystemSyncAccessHandle) 对象

```javascript
(async function(){
  // ...
  // 创建同步读写文件句柄
  const accessHandle = await fileHandles[0].createSyncAccessHandle();

  // 读取文件内容
  const fileSize = accessHandle.getSize();
  const buffer = new DataView(new ArrayBuffer(fileSize));
  const readBuffer = accessHandle.read(buffer, {at: 0});

  // 写入文件
  const encoder = new TextEncoder();
  const encoderMessage = encoder.encode('hello world');
  const writeBuffer = accessHandle.write(encoderMessage,{at: readBuffer});

  // 持久化
  accessHandle.flush();
  // 关闭
  accessHandle.close();
})();
```

##### createWritable()

返回一个已兑现的可用于写入文件的新创建的 [FileSystemWritableFileStream](#FileSystemWritableFileStream) 对象

```javascript
(async function(){
  // ...
  // 创建写句柄
  const writeHandle = await fileHandles[0].createWritable();
  // 写入文件
  await writeHandle.write("hello world");
  // 关闭
  await writeHandle.close();
})();
```

#### FileSystemDirectoryHandle <em id="FileSystemDirectoryHandle"></em> <!--markdownlint-disable-line-->

表示一个指向文件系统目录的句柄

- [window.showDirectoryPicker](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/showDirectoryPicker) 显示一个目录选择器, 允许用户选择一个目录, 返回一个已兑现的 `FileSystemDirectoryHandle` 对象
  - options 可选
    - id 指定 id, 浏览器可以为不同的 id 记住不同的目录
    - mode 默认 read, 标识当前的句柄模式, 支持 readwrite
    - startIn 一个 FileSystemHandle 对象或一个已知的(desktop, documents, downloads, music, pictures, videos)目录, 用于指定选择器的起始目录

```javascript
(async function () {
  // 可选参数, 传入一个目录选择器筛选目录的选项对象
  // 返回一个已兑现的 `FileSystemDirectoryHandle` 对象
  // FileSystemHandle <- FileSystemDirectoryHandle
  const dirHandle = await window.showDirectoryPicker({
    /* ... */
  });

  // 遍历目录
  for await (const [name, handle] of dirHandle.entries()) {
    if (handle.key == 'file') {
      const fileHandle = await dirHandle.getFileHandle(name);
      console.log(fileHandle);
    } else {
      const dh = dirHandle.getDirectoryHandle(name);
      console.log(dh);
    }
  }
})();
```

- [navigator.storage.getDirectory()](https://developer.mozilla.org/zh-CN/docs/Web/API/StorageManager/getDirectory) 获取 FileSystemDirectoryHandle 对象的引用, 允许访问存储在 源私有文件系统(OPFS) 中的目录及目录的内容
- [源私有文件系统 OPFS](https://developer.mozilla.org/zh-CN/docs/Web/API/File_System_API/Origin_private_file_system) 作为 File System API 的一部分提供了一个存储端点, 它是页面所属的源专用的, 并且不像常规文件系统那样对用户可见, 它提供了一种特殊类型文件的访问能力, 这种文件经过高度性能优化, 并提供对其内容的原地写入访问特性

```javascript
(async function(){
  // 获取文件句柄
  const root = await navigator.storage.getDirectory();
  const fileHandle = await root.getFileHandle('temp.txt', {create: true});
  // 获取同步访问句柄
  const accessHandle = await fileHandle.createSyncAccessHandle();
})();
```

##### getDirectoryHandle()

返回一个 Promise, 兑现一个调用此方法的目录句柄内指定名称的子目录的 [FileSystemDirectoryHandle](#FileSystemDirectoryHandle)

- name 子目录的名称
- options
  - create 默认 false, 为 true 时表示没有找到对应的目录将会创建一个指定名称的目录并返回

```javascript
(async function(){
  // 不存在则新建目录并返回
  const subDirHandle = await dirHandle.getDirectoryHandle('temp', {create: true});
})();
```

##### getFileHandle()

返回一个 Promise, 兑现一个调用此方法的目录句柄内指定名称的文件的 [FileSystemFileHandle](#FileSystemFileHandle)

- name 文件名称
- options
  - create 默认 false, 为 true 时表示没有找到对应的文件将会创建一个指定名称的文件并返回

```javascript
(async function(){
  // 不存在则新建文件并返回
  const fileHandle = await dirHandle.getFileHandle('temp.txt', {create: true});
})();
```

##### removeEntry()

尝试异步删除指定名称的文件或目录

- name 移除条目的名称
- options
  - recursive 默认 false, 为 true 时表示条目将会被递归删除

```javascript
(async function(){
  // 允许递归删除条目
  await dirHandle.removeEntry('subDir', {recursive: true});
})();
```

##### resolve()

返回一个 Promise, 兑现一个包含从父目录前往指定子条目中间的目录的名称的数组, 数组的最后一项是子条目的名称

- possibleDescendant 要返回其相对路径的 FileSystemHandle

```javascript
(async function(){
  // 检查文件句柄是否存在于目录句柄的目录中
  const relativePaths = await dirHandle.resolve(filehandles[0]);
  if(relativePaths === null){
    // 不在目录句柄中
  } else {
    // relativePaths 是一个包含名称的数组, 指示相对路径
  }
})();
```

##### entries()

返回给定对象自己的可枚举属性对的新**异步迭代[key, value]**器

```javascript
(async function(){
  const dirHandle = await window.showDirectoryPicker();

  for await(const [key, value] of dirHandle.entries()){
    console.log(key, value);
  }
})();
```

##### keys()

返回一个新的**异步迭代**器, 其中包含每个项目的键 `FileSystemDirectoryHandle`

##### values()

返回一个新的**异步迭代**器, 其中包含每个项目的值 `FileSystemDirectoryHandle`

#### FileSystemSyncAccessHandle <em id="FileSystemSyncAccessHandle"></em> <!--markdownlint-disable-line-->

表示一个指向文件系统条目的同步句柄

##### close()

关闭一个打开的同步文件句柄, 禁止之后对其的任何操作并且释放之前加在与文件句柄相关联的文件上的独占锁

##### flush()

将通过 write 方法对句柄相关联的文件所做的所有更改持久化到磁盘上

##### getSize()

返回与句柄相关联文件的字节大小

##### read()

将与句柄相关联文件的内容读取到指定的缓冲区中, 可选择在给定的偏移处开始读取

- buffer 指定的缓冲区
- options
  - at 开始读取的字节偏移量数字

```javascript
(async function(){
  const [fileHandle] = await window.showOpenFilePicker();
  const accessHandle = await fileHandle.createSyncAccessHandle();

  // 获取文件字节大小
  const fileSize = accessHandle.getSize();
  // 创建缓冲区
  const buffer = new DataView(new ArrayBuffer(fileSize));
  // 读取文件到缓冲区
  const readBuffer = accessHandle.read(buffer, {at: 0});
})();
```

##### truncate()

将与句柄相关联文件的大小调整为指定的字节数

- newSize 调整的字节数

```javascript
// 将文件裁剪至 0 字节
accessHandle.truncate(0);
```

##### write()

将指定缓冲区中的内容写入到与句柄相关联的文件, 可选择在指定的偏移处开始写入

- buffer 指定的缓冲区
- options
  - at 开始写入的字节偏移量数字

```javascript
(async function(){
  // ...

  // 写入文件
  const encoder = new TextEncoder();
  const encoderMessage = encoder.encode('hello world');
  const writeBuffer = accessHandle.write(encoderMessage, {at: readBuffer});
  accessHandle.flush();
  accessHandle.close();
})();
```

#### FileSystemWritableFileStream <em id="FileSystemWritableFileStream"></em> <!--markdownlint-disable-line-->

表示一个操作磁盘上单个文件的 WritableStream 对象, 通过 FileSystemFileHandle.createWritable() 方法访问

#### write(data)

向调用此方法的文件写入内容, 写入到文件当前指针偏移处

#### seek()

更新文件当前指针偏移到指定位置(字节)

- position 从文件开头起的字节位置

#### truncate(size)

将与流相关的文件调整为指定的字节大小

- size 调整到的字节数
