---
title: ArrayBuffer
date: 2022-03-05 10:35:34
categories:
  - ES
tags:
  - ES6
  - js
  - Buffer
---

## ArrayBuffer

ArrayBuffer 对象用来表示通用的、固定长度的原始二进制数据缓冲区, 可以理解为一个字节数组. 不能直接操作 ArrayBuffer, 需要通过 `类型化数组对象(TypedArray)`或 `DataView` 操作
ArrayBuffer 构造函数创建一个以字节为单位的固定长度的新 ArrayBuffer, 或者从现有的数据中获取数组缓冲区(例如: Base64 字符串或者 Blob 类文件对象 )

```javascript
var buffer = new ArrayBuffer(10); // 创建一个 10 字节的缓冲区
var i8a = new Int32Array(buffer); // 并使用 Int32Array 视图引用它
```

### TypedArray

不能实例化

描述底层 `二进制数据缓冲区(ArrayBuffer)` 的类数组视图, TypedArray 没有可用的全局属性和构造函数, 其为所有类型化数组的子类提供了实用方法的通用接口, 当创建 TypedArray 子类(例如 Int8Array) 的实例时, 在内存中会创建数组缓冲区, 如果将 ArrayBuffer 实例作为构造函数参数时, 则使用该 ArrayBuffer.

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
var ia = new Int8Array(10); // Int8Array(length) 创建一个长度为 10 的 Int8Array
ia[0] = 42;
```

#### 参数

实例化子类时可传入以下参数

- typedArray, 描述底层二进制数据缓冲区的类数组视图
- object, 描述底层二进制数据缓冲区的类数组视图
- length, 描述底层二进制数据缓冲区的字节长度
- buffer,byteOffset,length 创建一个指定字节长度的 ArrayBuffer

```typescript
var ia = new Int8Array(10); //  创建一个 10 个元素的 Int8Array
var ia = new Int8Array([1, 2, 3]); // 创建一个 3 个元素的 Int8Array
var ia = new Int8Array(new ArrayBuffer(10), 1, 8); // 创建一个 8 个元素的 Int8Array
```

<!--more-->

#### 静态属性

- TypedArray.BYTES_PER_ELEMENT  代表了强类型数组中每个元素所占用的字节数
- TypedArray[Symbol.species]   返回用于构造类型化数组方法返回值的构造函数

```typescript
Int8Array.BYTES_PER_ELEMENT; // 1
Uint8Array.BYTES_PER_ELEMENT; // 1
Uint8ClampedArray.BYTES_PER_ELEMENT; // 1
Int16Array.BYTES_PER_ELEMENT; // 2
Uint16Array.BYTES_PER_ELEMENT; // 2
Int32Array.BYTES_PER_ELEMENT; // 4
Uint32Array.BYTES_PER_ELEMENT; // 4
Float32Array.BYTES_PER_ELEMENT; // 4
Float64Array.BYTES_PER_ELEMENT; // 8

Int8Array[Symbol.species]; // function Int8Array()
Uint8Array[Symbol.species]; // function Uint8Array()
Float32Array[Symbol.species]; // function Float32Array()
```

#### 静态方法

- TypedArray.from() 从类数组对象或可迭代对象创建一个新的类型化数组, 此方法与 Array.from() 几乎相同
- TypedArray.of() 创建一个具有可变数量参数的新类型数组, 此方法与 Array.of() 很相似

```typescript
const u16 = Uint16Array.from([1, 2, 3]); // Uint16Array [1, 2, 3]
const u16f = Uint16Array.of(1, 2, 3); // Uint16Array [1, 2, 3]

const u8 = Uint8Array.from([1, 2, 3]); // Uint8Array [1, 2, 3]
const u8f = Uint8Array.of(1, 2, 3); // Uint8Array [1, 2, 3]
```

#### 实例属性

- TypedArray.prototype.buffer 获取缓冲区对象
- TypedArray.prototype.byteLength 获取缓冲区字节长度
- TypedArray.prototype.byteOffset 获取缓冲区字节偏移
- TypedArray.prototype.length 获取数组长度

```typescript
var buffer = new ArrayBuffer(8);
var uint16 = new Uint16Array(buffer);
uint16.buffer; // ArrayBuffer { byteLength: 8 }

var buffer = new ArrayBuffer(8);
var uint8 = new Uint8Array(buffer, 1, 5); // Uint8Array(buffer, byteOffset, length); 创建一个长度为 5 的 Uint8Array
uint8.byteLength; // 5 (在 Uint8Array 构造时指定)
uint8.byteOffset; // 1 (在 Uint8Array 构造时指定)
uint8.length; // 5 (在 Uint8Array 构造时指定)
```

#### 实例方法

为 TypedArray 实例提供和 Array 类相同的方法

### DataView

DataView 是一个可以从二进制 `ArrayBuffer` 对象中读写多种数值类型的底层接口, 使用它时, 不需要考虑不同平台的字节序问题

DataView 构造函数可以传入一个已经存在的 `ArrayBuffer` 或 `SharedArrayBuffer` 作为数据源, 第二个参数可以指定 buffer 中的字节偏移, 第三个参数可以指定 DataView 对象的字节长度, 返回表示指定数据缓冲区的新 DataView 对象

#### 构造函数参数

- buffer, 描述底层二进制数据缓冲区的类数组视图
- byteOffset, 描述底层二进制数据缓冲区的字节偏移
- byteLength, 描述底层二进制数据缓冲区的字节长度

```typescript
new DataView(buffer)
new DataView(buffer, byteOffset)
new DataView(buffer, byteOffset, byteLength)
```

#### dv 实例属性

- DataView.prototype.buffer 获取缓冲区对象
- DataView.prototype.byteLength 获取缓冲区字节长度
- DataView.prototype.byteOffset 获取缓冲区字节偏移

#### dv 实例方法

- DataView.prototype.get*(byteOffset) 获取指定字节偏移量的值
- DataView.prototype.set*(byteOffset, value) 设置指定字节偏移量的值

```javascript
var buffer = new ArrayBuffer(16);
var view = new DataView(buffer);

view.setUint8(0, 42); // 设置指定偏移量的值
view.getUint8(0); // 获取指定偏移量的值

view.setInt32(1, 2147483647);
view.getInt32(1);
```
