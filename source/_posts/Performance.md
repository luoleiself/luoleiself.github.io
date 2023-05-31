---
title: Performance
date: 2023-05-31 14:38:57
categories:
  - WebAPI
tags:
  - API
---

## Performance

Performance 接口可以获取当前页面中与性能相关的信息, 它是高级能力 Timing API 的一部分, 同时也融合了 Timeline API, Navigation Timing API, User Timing API, Resource Timing API

该类型的对象可以通过调用只读属性 `window.performance` 获取

### 只读属性

- navigation: 获取在指定的时间段内发生的操作相关信息, 包含页面的加载刷新、发生了多少次重定向等
- timing: 表示包含延迟相关的性能信息

  ```javascript
  performance.timing;
  // {
  //   "connectStart": 1685528410085,
  //   "navigationStart": 1685528410083,
  //   "secureConnectionStart": 0,
  //   "fetchStart": 1685528410085,
  //   "domContentLoadedEventStart": 1685528410878,
  //   "responseStart": 1685528410185,
  //   "domInteractive": 1685528410417,
  //   "domainLookupEnd": 1685528410085,
  //   "responseEnd": 1685528410186,
  //   "redirectStart": 0,
  //   "requestStart": 1685528410087,
  //   "unloadEventEnd": 0,
  //   "unloadEventStart": 0,
  //   "domLoading": 1685528410189,
  //   "domComplete": 1685528410896,
  //   "domainLookupStart": 1685528410085,
  //   "loadEventStart": 1685528410897,
  //   "domContentLoadedEventEnd": 1685528410878,
  //   "loadEventEnd": 1685528410897,
  //   "redirectEnd": 0,
  //   "connectEnd": 1685528410085
  // }
  ```

- memory: `非标准属性`, 表示基本内存使用情况的对象
- timeOrigin: 表示性能测量开始时的时间的高精度时间戳

<!-- more -->

### 事件

- resourcetimingbufferfull: 当浏览器的资源时间性能缓冲区已满时触发

```javascript
performance.onresourcetimingbufferfull = function (evt) {
  console.log('warning: Resource Timing Buffer is FULL!');
  performance.setResourceTimingBufferSize(300);
};
```

### 方法

#### p9e.clearMarks()

将给定的 mark 从浏览器的性能输入缓冲区中移除, 如果未指定参数则所有 entryType 值为 mark 的 [PerformanceEntry](#PerformanceEntry) 将从缓冲区中移除

- 参数 name

```javascript
performance.clearMarks();
```

#### p9e.clearMeasures()

将给定的 measure 从浏览器的性能输入缓冲区中移除, 如果未指定参数则所有 entryType 值为 measure 的 [PerformanceEntry](#PerformanceEntry) 将从缓冲区移除

- 参数 name

```javascript
performance.clearMeasures();
```

#### p9e.clearResourceTimings()

将所有 entryType 值为 resource 的 [PerformanceEntry](#PerformanceEntry) 从缓冲区移除

```javascript
performance.clearResourceTimings();
```

#### p9e.getEntries()

基于给定的 filter 返回一个 [PerformanceEntry](#PerformanceEntry) 对象列表, 如果没有符合 filter 条件的返回空数组, 如果未指定参数则返回所有 [PerformanceEntry](#PerformanceEntry)

- 参数
  - options: 可选属性, 一个包含键值对的过滤配置项
    - name: `PerformanceEntry` 的名字
    - entryType: `PerformanceEntry` 的 entryType, 合法的类型可以从 `PerformanceEntry.entryType` 获取
    - initiatorType: 初始化资源的类型, 在 `PerformanceResourceTiming.initiatorType` 接口中定义

```javascript
// 获取所有的 PerformanceEntry
const entries = performance.getEntries();
// 获取指定的 PerformanceEntry
const entries = performance.getEntries({
  name: 'https://developer.mozilla.org/favicon-192x192.png',
  entryType: 'resource',
  initiatorType: 'other',
});
```

#### p9e.getEntriesByName()

基于给定的 name 和 entryType 返回一个 [PerformanceEntry](#PerformanceEntry) 对象列表, 如果未找到返回空数组

- 参数
  - name: `PerformanceEntry` 的名字
  - type: `PerformanceEntry` 的 entryType

```javascript
performance.mark('Begin');
// do something in 20 seconds
performance.mark('End');

const entries = performance.getEntriesByName('Begin', 'mark');
// [{
//   name: 'Begin',
//   entryType: 'mark',
//   startTime: 92068.30000019073,
//   duration: 0,
// }];
const entries = performance.getEntriesByName('End', 'mark');
// [{
//   name: 'End',
//   entryType: 'mark',
//   startTime: 112071.60000038147,
//   duration: 0,
// }];
```

#### p9e.getEntriesByType()

基于给定的 entryType 返回一个 [PerformanceEntry](#PerformanceEntry) 对象的列表, 未指定参数返回 TypeError, 未找到返回空数组

- 参数 entryType: `PerformanceEntry` 的 entryType

```javascript
const entries = performance.getEntriesByType('paint');
// [{
//     "name": "first-paint",
//     "entryType": "paint",
//     "startTime": 682,
//     "duration": 0
//   },
//   {
//     "name": "first-contentful-paint",
//     "entryType": "paint",
//     "startTime": 682,
//     "duration": 0
// }]
```

#### p9e.mark()

> 通过 `performance.getEntries*` 方法可以获取到

根据给定的 name, 在浏览器的性能输入缓冲区中创建一个相关的 timestamp, 如果指定的 name 已经存在于 `PerformanceTiming` 接口则抛出一个 SyntaxError

- 参数 name: 指定标记的名字

标记的 [PerformanceEntry](#PerformanceEntry) 默认包含的属性

- entryType: 默认为 mark
- name: 调用 mark 方法时指定的 name
- startTime: 调用 mark 方法时的时间戳
- duration: 默认为 0(标记没有持续时间)

```javascript
const dogMark = performance.mark('dog');
// {
//   "name": "dog",
//   "entryType": "mark",
//   "startTime": 671450.1999998093,
//   "duration": 0
// }
performance.getEntriesByName('dog', 'mark');
// [{
//   name: 'dog',
//   entryType: 'mark',
//   startTime: 671450.1999998093,
//   duration: 0,
// }];
```

#### p9e.measure()

> 通过 `performance.getEntries*` 方法可以获取到

在浏览器的指定 start mark 和 end mark 间的性能输入缓冲区中创建一个指定的 timestamp

- 参数
  - name: 指定测量的名字
  - startMark: 可选属性, 表示测量开始的标记名字
  - endMark: 可选属性, 表示测量结束的标记名字

测量的 [PerformanceEntry](#PerformanceEntry) 默认包含的属性

- entryType: 默认为 measure
- name: 调用 measure 方法时指定的 name
- startTime: 调用 measure 方法开始的时间戳
- duration: 测量的持续时间, 通常为结束时间戳减去开始时间戳的差值

```javascript
performance.mark('my-measure-start-mark');
setTimeout(function () {
  performance.mark('my-measure-end-mark');

  // 如果使用的 mark 不存在需要先创建, 否则报错
  performance.measure(
    'my-measure',
    'my-measure-start-mark',
    'my-measure-end-mark'
  );

  const entries = performance.getEntriesByName('my-measure');
  // [{
  //   name: 'my-measure',
  //   entryType: 'measure',
  //   startTime: 111720.0999994278,
  //   duration: 20006.300000190735,
  // }];

  // 清除缓冲区 entryType 为 mark 的标记
  performance.clearMarks();
  // 清除缓冲区 entryType 为 measure 的标记
  performance.clearMeasures();
}, 20000);
```

#### p9e.now()

返回一个表示从性能测量时刻开始经过的毫秒数, 这个时间戳不是高精度的

```javascript
performance.now();
// 24152.300000190735;
```

#### p9e.setResourceTimingBufferSize()

将浏览器的 resource timing 缓冲区的大小设置为 [PerformanceEntry](#PerformanceEntry) 对象的指定数量

#### p9e.toJSON()

返回 performance 对象的 JSON 对象

```javascript
performance.toJSON();
// {
//   "timeOrigin": 1685527847332.7,
//   "timing": {
//     "connectStart": 1685527848166,
//     "navigationStart": 1685527848164,
//     "secureConnectionStart": 0,
//     "fetchStart": 1685527848166,
//     "domContentLoadedEventStart": 1685527848860,
//     "responseStart": 1685527848276,
//     "domInteractive": 1685527848521,
//     "domainLookupEnd": 1685527848166,
//     "responseEnd": 1685527848277,
//     "redirectStart": 0,
//     "requestStart": 1685527848169,
//     "unloadEventEnd": 0,
//     "unloadEventStart": 0,
//     "domLoading": 1685527848281,
//     "domComplete": 1685527849039,
//     "domainLookupStart": 1685527848166,
//     "loadEventStart": 1685527849039,
//     "domContentLoadedEventEnd": 1685527848860,
//     "loadEventEnd": 1685527849039,
//     "redirectEnd": 0,
//     "connectEnd": 1685527848166
//   },
//   "navigation": {
//     "type": 2,
//     "redirectCount": 0
//   }
// }
```

## PerformanceEntry <em id="PerformanceEntry"></em> <!-- markdownlint-disable-line -->

PerformanceEntry 对象代表了 performance 时间列表中的单个 metric 数据, 每个 performance entry 都可以在应用运行过程中通过手动调用 `performance.mark` 或 `performance.measure` 方法创建, 同时, performance entry 在资源加载的时候也会被动生成(例如 img, css, script 等资源)

### 属性

- name: 只读属性, 表示 PerformanceEntry 的名字
- entryType: 只读属性, 表示上报的 performance metric 的 entryType 类型, 例如 'mark'
  - frame/navigation: 文档的地址
  - resource: 请求的资源的地址
  - mark: 当调用 `performance.mark()` 方法时作为 `PerformanceEntry` 的 name
  - measure: 当调用 `performance.measure()` 方法时作为 `PerformanceEntry` 的 name
  - paint: 其他例如 `first-paint` 或 `first-contentful-paint`
- startTime: 只读属性, 表示 metric 上报时的时间
- duration: 只读属性, 表示 DOMHighResTimeStamp 事件的耗时

### 实例方法

- toJSON() 返回 PerformanceEntry 对象的 JSON 格式数据

```javascript
performance.getEntries()[0].toJSON();
// {
//   "name": "https://developer.mozilla.org/zh-CN/docs/Web/API/PerformanceEntry",
//   "entryType": "navigation",
//   "startTime": 0,
//   "duration": 2278.7000007629395,
//   "initiatorType": "navigation",
//   "nextHopProtocol": "http/1.1",
//   "renderBlockingStatus": "blocking",
//   "workerStart": 0,
//   "redirectStart": 0,
//   "redirectEnd": 0,
//   "fetchStart": 6.100000381469727,
//   "domainLookupStart": 6.100000381469727,
//   "domainLookupEnd": 6.100000381469727,
//   "connectStart": 9.100000381469727,
//   "connectEnd": 831.1000003814697,
//   "secureConnectionStart": 722.1000003814697,
//   "requestStart": 831.3000001907349,
//   "responseStart": 1025.1000003814697,
//   "responseEnd": 1026.5,
//   "transferSize": 12336,
//   "encodedBodySize": 12036,
//   "decodedBodySize": 55753,
//   "serverTiming": [],
//   "unloadEventStart": 0,
//   "unloadEventEnd": 0,
//   "domInteractive": 1506.2000007629395,
//   "domContentLoadedEventStart": 1912.9000005722046,
//   "domContentLoadedEventEnd": 1913.5,
//   "domComplete": 2278.5,
//   "loadEventStart": 2278.7000007629395,
//   "loadEventEnd": 2278.7000007629395,
//   "type": "navigate",
//   "redirectCount": 0,
//   "activationStart": 0
// }
```

## PerformanceObserver

PerformanceObserver 用于检测性能度量事件, 在浏览器的性能时间轴记录新的 performance entry 的时候将会被通知

### 构造函数

使用给定的观察者生成一个新的 `PerformanceObserver` 实例, 当通过 `observe` 方法注册的条目类型的性能条目事件被记录下来时, 会调用该观察者回调

- 参数 callback
  观察者的性能事件被记录时将调用构造函数注册的回调, 回调函数有两个参数:

  - list: 描述性能观察条目列表
  - observer: 调用该函数的 PerformanceObserver 对象

```javascript
const observer = new PerformanceObserver(function (list, observer) {
  console.log(list, observer);
  console.log(list.getEntries());
  // console.log(list.getEntriesByName());
  // console.log(list.getEntriesByType());
  // [{
  //   name: 'https://pvx.xcar.com.cn/ckl.gif?1=1&b=chrome&pw=1068&ph=1057&pvh=1057&st=0&dm=https%3A%2F%2Faikahao.xcar.com.cn&lp=https%3A%2F%2Faikahao.xcar.com.cn%2Fmy%2Fregister.html&ref=https%3A%2F%2Faikahao.xcar.com.cn%2F&ds=1920x1200&uv=640afc3c3b6e1&pvx_uv=640afc3c76993&_t=1685520132758&_wd=false&_uid=16034823&tpinfo=%5B%7B%220%22%3A%22127167%22%7D%5D',
  //   entryType: 'resource',
  //   startTime: 10342.599999427795,
  //   duration: 24.40000057220459,
  //   initiatorType: 'img',
  //   nextHopProtocol: '',
  //   renderBlockingStatus: 'non-blocking',
  //   workerStart: 0,
  //   redirectStart: 0,
  //   redirectEnd: 0,
  //   fetchStart: 10342.599999427795,
  //   domainLookupStart: 0,
  //   domainLookupEnd: 0,
  //   connectStart: 0,
  //   connectEnd: 0,
  //   secureConnectionStart: 0,
  //   requestStart: 0,
  //   responseStart: 0,
  //   responseEnd: 10367,
  //   transferSize: 0,
  //   encodedBodySize: 0,
  //   decodedBodySize: 0,
  //   serverTiming: [],
  // }];
});
observer.observe({ entryTypes: ['mark', 'element', 'measure', 'resource'] });
```

- 返回值 PerformanceObserver 对象

### 静态属性

- supportedEntryTypes: 只读属性, 返回用户代理支持的 PerformanceObserver 接口的属性数组

```javascript
PerformanceObserver.supportedEntryTypes;
// [
//   'element',
//   'event',
//   'first-input',
//   'largest-contentful-paint',
//   'layout-shift',
//   'longtask',
//   'mark',
//   'measure',
//   'navigation',
//   'paint',
//   'resource',
// ];
```

### observer 方法

#### observer.disconnect()

用于阻止性能观察者接收任何性能条目事件

- 参数 无
- 返回值 undefined

#### observer.observe()

用于观察传入的参数中指定的性能条目类型的集合, 当记录一个指定类型的性能条目时, 性能检测对象的回调函数将会被调用

- 参数
  - options: 一个包含单个键值对的对象, 该键值对的键名为 entryTypes, 如果未传入 options 或 options 实参为空数组会抛出 TypeError
    - entryTypes: 一个放字符串的数组, 字符串的有效值取值在性能条目类型中, 如果取值无效则浏览器会自动忽略
- 返回值 undefined

```javascript
observer.observe({ entryTypes: ['element', 'mark', 'measure'] });
```

#### observer.takeRecords()

返回当前存储在性能观察器中的性能条目列表, 并将其清空

- 参数 无
- 返回值 [PerformanceEntry](#PerformanceEntry) 对象列表

```javascript
const records = observer.takeRecords();
```
