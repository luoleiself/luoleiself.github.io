---
title: FileSystemAccessAPI
date: 2023-04-12 14:02:45
categories:
  - WebAPI
tags:
  - API
---

> 目前仅在 chrome 86 (edge 86, opera 72)及以上版本支持, safari 和 firefox 暂时不支持

允许 Web 应用程序从用户设备的本地文件系统中操作文件, 它为 Web 应用程序提供了更多的灵活性和功能, 使其更接近于本地应用程序的体验

[File System Access API](https://developer.mozilla.org/en-US/docs/Web/API/File_System_Access_API) 遵循同源策略, 只允许 Web 应用程序在具有相同源的文件系统上进行操作, 当使用改 API 时, 会提示用户授权应用程序访问文件系统

- 将文件从本地文件系统上传到 Web 应用程序
- 将 Web 应用程中的数据写入到本地文件系统中
- 在用户的本地文件系统中创建、重命名和删除文件
- 读取本地文件系统上的文件内容

### [FileSystemHandle](https://developer.mozilla.org/en-US/docs/Web/API/FileSystemHandle)接口

是 File System Access API 表示文件或目录条目的对象, 多个句柄可以代表同一个条目，通常情况下, 使用它的子接口 `FileSystemFileHandle` 和 `FileSystemDirectoryHandle`

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

<!-- more -->

#### FSH 实例属性

- kind 返回条目的类型, `file` 表示文件, `directory` 表示目录
- name 返回关联条目的名称

#### FSH 实例方法

- isSameEntry() 比较两者 handle 以查看相关条目(文件或目录)是否匹配
- queryPermission() 查询当前句柄的当前权限状态
- remove() 请求从底层文件系统中删除由句柄表示的条目
- requestPermission() 请求文件句柄的读取或读写权限

### [showOpenFilePicker](https://developer.mozilla.org/zh-CN/docs/Web/API/Window/showOpenFilePicker)

显示一个文件选择器, 允许用户选择一个或多个文件并返回这些文件的句柄, 返回一个 Promise 对象, 并兑现一个包含 `FileSystemFileHandle` 对象的 Array 数组

```javascript
(async function () {
  // 可选参数, 传入一个文件选择器筛选文件的选项对象
  // 返回一个已兑现的包含 `FileSystemFileHandle` 对象的 Array 数组
  const fileHandles = await window.showOpenFilePicker({
    /* ... */
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

#### SOFP 实例方法 <em id="showopenfilePickerfn"></em> <!-- markdownlint-disable-line-->

##### getFile()

返回一个已兑现的由句柄表示的条目在磁盘上的状态的对象

##### createSyncAccessHandle()

返回一个已兑现的可用于同步读取和写入文件的 `FileSystemSyncAccessHandle` 对象

##### createWritable()

返回一个已兑现的可用于写入文件的新创建的 `FileSystemWritableFileStream` 对象

### [showSaveFilePicker](https://developer.mozilla.org/en-US/docs/Web/API/Window/showSaveFilePicker)

显示一个允许用户保存文件的文件选择器, 通过选择现有文件或者输入新文件的名称, 返回一个已兑现的 `FileSystemFileHandle` 对象

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

#### SSFP 实例方法

实例方法同 [`showOpenFilePicker`](#showopenfilePickerfn) 方法返回的实例的方法

### [showDirectoryPicker](https://developer.mozilla.org/en-US/docs/Web/API/Window/showDirectoryPicker)

显示一个目录选择器, 允许用户选择一个目录, 返回一个已兑现的 `FileSystemDirectoryHandle` 对象

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

#### SDP 实例方法

##### entries()

返回给定对象自己的可枚举属性对的新**异步迭代[key, value]**器

##### getFileHandle()

返回指定名称的文件的已兑现的 `FileSystemFileHandle` 对象

##### getDirectoryHandle()

返回指定名称的目录的已兑现的 `FileSystemDirectoryHandle` 对象

##### resolve()

返回一个从父句柄到指定子项的目录名称的 Promise, 子项的名称作为最后一个数组项

##### removeEntry()

尝试异步删除指定名称的文件或目录

##### keys()

返回一个新的**异步迭代**器, 其中包含每个项目的键 `FileSystemDirectoryHandle`

##### values()

返回一个新的**异步迭代**器, 其中包含对象中每个索引的值 `FileSystemDirectoryHandle`
