---
title: Tapable总结
date: 2021-07-04 21:08:54
categories:
  - webpack
tags:
  - js
  - webpack
  - Tapable
---

[webpack](https://webpack.docschina.org/) 是一个用于现代 JavaScript 应用程序的静态模块打包工具, webpack 处理应用程序时会在内部构建一个依赖图, 此依赖图对应映射到项目所需的每个模块，并生成一个或多个 bundle

[plugin](https://webpack.docschina.org/api/plugins/) 是 webpack 生态的关键部分, 采用基于事件流的机制, 将各个插件串联起来完成相关的功能, [compiler](https://webpack.docschina.org/api/compiler-hooks/) 模块是 [webpack](https://webpack.docschina.org/) 的主要引擎, 它扩展(extend)自 [Tapable](https://github.com/webpack/tapable#tapable) 类, 用来注册和调用插件.

> [Tapable](https://github.com/webpack/tapable#tapable) 是一个用于事件发布订阅执行的插件架构, 类似于 Node.js 的 EventEmitter 库.

![tapable-1](/images/tapable-1.jpg)

```javascript
const {
  SyncHook,
  SyncBailHook,
  SyncWaterfallHook,
  SyncLoopHook,
  AsyncParallelHook,
  AsyncParallelBailHook,
  AsyncSeriesHook,
  AsyncSeriesBailHook,
  AsyncSeriesWaterfallHook,
  AsyncSeriesLoopHook,
  HookMap,
  MultiHook,
} = require('tapable');
```

### 钩子分类

#### 执行方式

##### Basic Hook (基础)

钩子调用所在行中调用的每个钩子函数

##### WaterFall (瀑布)

与基础钩子不同, 它将一个返回值从每个函数传递到下一个函数

##### Bail (保证)

钩子函数执行中, 只要其中有一个钩子返回 非 undefined 时, 则剩余的钩子函数不再执行

##### Loop (循环)

循环执行钩子, 当循环钩子函数返回 非 undefined 时, 则从第一个钩子重新启动, 直到所有的钩子返回 undefined 时结束

<!-- more -->

#### 定义类型

##### Sync 同步

同步钩子只能使用 tap 注册钩子, 使用 call | callAsync 调用钩子

```javascript
myHooks.tap('x', () => {});
myHooks.call('x');
myHooks.callAsync('x', () => {});
```

##### AsyncSeries 异步串行

异步串行钩子可以使用 tap | tapSync | tapPromise 注册钩子, 使用 call | callAsync | promise 调用钩子

```javascript
myHooks.tapAsync('x', (callback) => {
  callback();
});
myHooks.callAsync('x', () => {});
```

##### AsyncParallel 异步并行

异步并行钩子可以使用 tap | tapSync | tapPromise 注册钩子, 使用 call | callAsync | promise 调用钩子

```javascript
myHooks.tapPromise('x', () => {
  return new Promise((resolve, reject) => {
    resolve();
  });
});
myHooks.promise('x').then((res) => {});
```

### 使用

#### 同步钩子

##### SyncHook

钩子函数依次全部执行，如果有 hook 回调，则最后执行

```javascript
const { SyncHook } = require('tapable');

const syncHook = new SyncHook(['name']);
// 钩子函数依次全部执行，如果有 hook 回调，则最后执行
syncHook.tap('x', (name) => {
  console.log('x done ', name);
});
syncHook.tap('y', (name) => {
  console.log('y done ', name);
});
syncHook.call('call');
syncHook.callAsync('callAsync', () => {
  console.log('syncHook.callAsync');
});
/*
$ node SyncHook.js
x done  call
y done  call
x done  callAsync
y done  callAsync
syncHook.callAsync
*/
// 模拟 SyncHook 类
class MySyncHook {
  constructor(args) {
    this.args = args;
    this.tasks = [];
  }
  tap(name, task) {
    this.tasks.push(task);
  }
  call(...args) {
    // 也可在参数不足时抛出异常
    if (args.length < this.args.length) throw new Error('参数不足');

    // 传入参数严格对应创建实例传入数组中的规定的参数，执行时多余的参数为 undefined
    args = args.slice(0, this.args.length);

    // 依次执行事件处理函数
    this.tasks.forEach((task) => task(...args));
  }
}
```

##### SyncBailHook

钩子函数依次执行, 如果某个钩子函数的返回值为 非 undefined，则后面的钩子不再执行，如果有 hook 回调，则最后执行

```javascript
const { SyncBailHook } = require('tapable');

const syncBailHook = new SyncBailHook(['name']);
// 钩子函数依次执行, 如果某个钩子函数的返回值为 非 undefined，则后面的钩子不再执行，如果有 hook 回调，则最后执行
syncBailHook.tap('x', (name) => {
  console.log('x done ', name);
  return 'undefined'; // 返回值为非 undefined 不再执行其他 hooks
});
syncBailHook.tap('y', (name) => {
  console.log('y done ', name);
});
syncBailHook.call('call');
syncBailHook.callAsync('callAsync', () => {
  console.log('syncBailHook.callAsync');
});
/*
$ node SyncBailHook.js
x done  call
x done  callAsync
syncBailHook.callAsync
*/
// 模拟 SyncBailHook 类
class MySyncBailHook {
  constructor(args) {
    this.args = args;
    this.tasks = [];
  }
  tap(name, task) {
    this.tasks.push(task);
  }
  call(...args) {
    // 传入参数严格对应创建实例传入数组中的规定的参数，执行时多余的参数为 undefined
    args = args.slice(0, this.args.length);

    // 依次执行事件处理函数，如果返回值不为空，则停止向下执行
    let i = 0,
      ret;
    do {
      ret = this.tasks[i++](...args);
    } while (!ret);
  }
}
```

##### SyncWaterfallHook

钩子函数依次全部执行，上一个钩子函数的返回值作为下一个钩子函数的参数，如果有 hook 回调，则最后执行

```javascript
const { SyncWaterfallHook } = require('tapable');

const syncWaterfallHook = new SyncWaterfallHook(['name']);
// 钩子函数依次全部执行，上一个钩子函数的返回值作为下一个钩子函数的参数，如果有 hook 回调，则最后执行
syncWaterfallHook.tap('x', (name) => {
  console.log('x done ', name);
  return `${name} from x...`;
});
syncWaterfallHook.tap('y', (name) => {
  console.log('y done ', name);
});
syncWaterfallHook.call('call');
syncWaterfallHook.callAsync('callAsync', () => {
  console.log('syncWaterfallHook.callAsync');
});
/*
$ node SyncWaterfallHook.js
x done  call
y done  call from x...
x done  callAsync
y done  callAsync from x...
syncWaterfallHook.callAsync
*/
// 模拟 SyncWaterfallHook 类
class MySyncWaterfallHook {
  constructor(args) {
    this.args = args;
    this.tasks = [];
  }
  tap(name, task) {
    this.tasks.push(task);
  }
  call(...args) {
    // 传入参数严格对应创建实例传入数组中的规定的参数，执行时多余的参数为 undefined
    args = args.slice(0, this.args.length);

    // 依次执行事件处理函数，事件处理函数的返回值作为下一个事件处理函数的参数
    let [first, ...others] = this.tasks;
    return reduce((ret, task) => task(ret), first(...args));
  }
}
```

##### SyncLoopHook

钩子函数依次全部执行，当循环钩子中回调函数返回非 undefined 时，钩子将从第一个重新启动，直到所有的钩子返回非 undefined 时结束 如果有 hook 回调，则最后执行

```javascript
const { SyncLoopHook } = require('tapable');

const syncLoopHook = new SyncLoopHook(['name']);
// 钩子函数依次全部执行，当循环钩子中回调函数返回非 undefined 时，钩子将从第一个重新启动，直到所有的钩子返回非 undefined 时结束 如果有 hook 回调，则最后执行
syncLoopHook.tap('x', (name) => {
  let flag = Math.floor(Math.random() * 10);
  if (flag > 5) {
    console.log('x done ', name, flag);
    return undefined;
  } else {
    console.log('x loop ', name, flag);
    return 'undefined';
  }
});
syncLoopHook.tap('y', (name) => {
  let flag = Math.floor(Math.random() * 15);
  if (flag > 10) {
    console.log('y done ', name, flag);
    return undefined;
  } else {
    console.log('y loop ', name, flag);
    return 'undefined';
  }
});
syncLoopHook.tap('z', (name) => {
  console.log('z done ', name);
  return undefined;
});
syncLoopHook.call('call');
syncLoopHook.callAsync('callAsync', () => {
  console.log('syncLoopHook.callAsync');
});
/*
$ node SyncLoopHook.js
x done  call 8
y loop  call 5
x done  call 7
y done  call 14
z done  call
x loop  callAsync 1
x loop  callAsync 3
x loop  callAsync 2
x loop  callAsync 1
x loop  callAsync 1
x loop  callAsync 4
x done  callAsync 8
y loop  callAsync 2
x loop  callAsync 2
x done  callAsync 8
y done  callAsync 13
z done  callAsync
syncLoopHook.callAsync
*/
// 模拟 SyncLoopHook 类
class MySyncLoopHook {
  constructor(args) {
    this.args = args;
    this.tasks = [];
  }
  tap(name, task) {
    this.tasks.push(task);
  }
  call(...args) {
    // 传入参数严格对应创建实例传入数组中的规定的参数，执行时多余的参数为 undefined
    args = args.slice(0, this.args.length);

    // 依次执行事件处理函数，如果返回值为 true，则继续执行当前事件处理函数
    // 直到返回 undefined，则继续向下执行其他事件处理函数
    this.tasks.forEach((task) => {
      let ret;
      do {
        ret = this.task(...args);
      } while (ret === true || !(ret === undefined));
    });
  }
}
```

#### 异步并行

##### AsyncParallelHook

钩子函数异步并行全部执行，所有钩子回调执行完后，hook 回调执行

```javascript
const { AsyncParallelHook } = require('tapable');

const asyncParallelHook = new AsyncParallelHook(['name']);
// 钩子函数异步并行全部执行，所有钩子回调执行完后，hook 回调执行
asyncParallelHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 5s ', name);
    callback();
  }, 5000);
});
asyncParallelHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 2s ', name);
    callback();
  }, 2000);
});
asyncParallelHook.tapAsync('z', (name, callback) => {
  console.log('z done ', name);
  setTimeout(() => {
    console.log('z done setTimeout 3s ', name);
    callback();
  }, 3000);
});
asyncParallelHook.tapPromise('w', (name) => {
  return new Promise((resolve, reject) => {
    console.log('w done ', name);
    setTimeout(() => {
      console.log('w done setTimeout 8s', name);
      resolve('hello world');
    }, 8000);
  });
});
asyncParallelHook.callAsync('callAsync', () => {
  console.log('asyncParallelHook.callAsync');
});
// 等价
// asyncParallelHook.promise('callAsync').then((res) => {
//   console.log('asyncParallelHook.callAsync');
// });
/*
$ node AsyncParallelHook.js
x done  callAsync
y done  callAsync
z done  callAsync
w done  callAsync
y done setTimeout 2s  callAsync
z done setTimeout 3s  callAsync
x done setTimeout 5s  callAsync
w done setTimeout 8s callAsync
asyncParallelHook.callAsync
*/
// 模拟 AsyncParallelHook 类：tapAsync/callAsync
class MyAsyncParallelHook {
  constructor(args) {
    this.args = args;
    this.tasks = [];
  }
  tabAsync(name, task) {
    this.tasks.push(task);
  }
  callAsync(...args) {
    // 先取出最后传入的回调函数
    let finalCallback = args.pop();

    // 传入参数严格对应创建实例传入数组中的规定的参数，执行时多余的参数为 undefined
    args = args.slice(0, this.args.length);

    // 定义一个 i 变量和 done 函数，每次执行检测 i 值和队列长度，决定是否执行 callAsync 的回调函数
    let i = 0;
    let done = () => {
      if (++i === this.tasks.length) {
        finalCallback();
      }
    };

    // 依次执行事件处理函数
    this.tasks.forEach((task) => task(...args, done));
  }
}
```

##### AsyncParallelBailHook

钩子函数全部异步并行执行, 只要有一个钩子返回了非 undefined 值时， hook 回调会立即执行

```javascript
const { AsyncParallelBailHook } = require('tapable');

const asyncParallelBailHook = new AsyncParallelBailHook(['name']);
// 钩子函数全部异步并行执行, 只要有一个钩子返回了非 undefined 值时， hook 回调会立即执行
asyncParallelBailHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 1s');
    callback();
  }, 1000);
});
asyncParallelBailHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 2s');
    callback('undefined');
  }, 2000);
});
asyncParallelBailHook.tapAsync('z', (name, callback) => {
  console.log('z done ', name);
  setTimeout(() => {
    console.log('z done setTimeout 3s');
    callback();
  }, 3000);
});
asyncParallelBailHook.callAsync('callAsync', () => {
  console.log('asyncParallelBailHook.callAsync');
});
/*
$ node AsyncParallelBailHook.js
x done  callAsync
y done  callAsync
z done  callAsync
x done setTimeout 1s
y done setTimeout 2s
asyncParallelBailHook.callAsync
z done setTimeout 3s
*/
```

#### 异步串行

##### AsyncSeriesHook

钩子函数全部异步串行执行，执行顺序按照注册顺序执行，上一个钩子执行结束后下一个执行开始, hook 回调最后执行

```javascript
const { AsyncSeriesHook } = require('tapable');

const asyncSeriesHook = new AsyncSeriesHook(['name']);
// 钩子函数全部异步串行执行，执行顺序按照注册顺序执行，上一个钩子执行结束后下一个执行开始, hook 回调最后执行
asyncSeriesHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 2s');
    callback();
  }, 2000);
});
asyncSeriesHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 1s');
    callback();
  }, 1000);
});
asyncSeriesHook.tapAsync('z', (name, callback) => {
  console.log('z done', name);
  setTimeout(() => {
    console.log('z done setTimeout 3s');
    callback();
  }, 3000);
});
asyncSeriesHook.callAsync('callAsync', () => {
  console.log('asyncSeriesHook.callAsync');
});
/*
$ node AsyncSeriesHook.js
x done  callAsync
x done setTimeout 2s
y done  callAsync
y done setTimeout 1s
z done callAsync
z done setTimeout 3s
asyncSeriesHook.callAsync
*/
```

##### AsyncSeriesBailHook

钩子函数全部异步串行执行，只要有一个钩子返回了 非 undefined 值时，hook 回调立即执行，其他钩子有可能不再执行

```javascript
const { AsyncSeriesBailHook } = require('tapable');

const asyncSeriesBailHook = new AsyncSeriesBailHook(['name']);
// 钩子函数全部异步串行执行，只要有一个钩子返回了 非 undefined 值时，hook 回调立即执行，其他钩子有可能不再执行
asyncSeriesBailHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 2s');
    callback();
  }, 2000);
});
asyncSeriesBailHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 1s');
    callback('undefined');
  }, 1000);
});
asyncSeriesBailHook.tapAsync('z', (name, callback) => {
  console.log('z done ', name);
  setTimeout(() => {
    console.log('z done setTimeout 3s');
    callback();
  }, 3000);
});
asyncSeriesBailHook.callAsync('callAsync', () => {
  console.log('asyncSeriesBailHook.callAsync');
});
/*
$ node AsyncSeriesBailHook.js
x done  callAsync
x done setTimeout 2s
y done  callAsync
y done setTimeout 1s
asyncSeriesBailHook.callAsync
*/
```

##### AsyncSeriesWaterfallHook

钩子函数全部异步串行执行, 上一个钩子的返回的结果作为下一个钩子的参数，hook 回调在所有钩子回调返回后执行

```javascript
const { AsyncSeriesWaterfallHook } = require('tapable');

const asyncSeriesWaterfallHook = new AsyncSeriesWaterfallHook(['name']);
// 钩子函数全部异步串行执行, 上一个钩子的返回的结果作为下一个钩子的参数，hook 回调在所有钩子回调返回后执行
asyncSeriesWaterfallHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 1s');
    callback(null, ' from x... '); // 不会阻止 y 的执行
  }, 1000);
});

asyncSeriesWaterfallHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 2s');
    callback(null, ' from y... '); // 不会阻止 z 的执行
  }, 2000);
});
asyncSeriesWaterfallHook.tapAsync('z', (name, callback) => {
  console.log('z done ', name);
  callback(' from z...');
});
asyncSeriesWaterfallHook.callAsync('callAsync', (...args) => {
  console.log('asyncSeriesWaterfallHook.callAsync', args);
});
/*
$ node AsyncSeriesWaterfallHook.js
x done  callAsync
x done setTimeout 1s
y done   from x...
y done setTimeout 2s
z done   from y...
asyncSeriesWaterfallHook.callAsync [ ' from z...' ]
*/
```

##### AsyncSeriesLoopHook

钩子函数全部异步串行执行，当循环钩子中回调函数返回非 undefined 时，钩子将从第一个重新启动，直到所有的钩子返回非 undefined 时结束，hook 回调最后执行

```javascript
const { AsyncSeriesLoopHook } = require('tapable');

const asyncSeriesLoopHook = new AsyncSeriesLoopHook(['name']);
// 钩子函数全部异步串行执行，当循环钩子中回调函数返回非 undefined 时，钩子将从第一个重新启动，直到所有的钩子返回非 undefined 时结束，hook 回调最后执行
asyncSeriesLoopHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 1s');
    let flag = Math.floor(Math.random() * 10);
    if (flag > 5) {
      console.log('x done ', name, flag);
      callback(null, undefined);
    } else {
      console.log('x loop ', name, flag);
      callback(null, 'undefined');
    }
  }, 1000);
});
asyncSeriesLoopHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 2s');
    let flag = Math.floor(Math.random() * 15);
    if (flag > 10) {
      console.log('y done ', name, flag);
      callback(null, undefined);
    } else {
      console.log('y loop ', name, flag);
      callback(null, 'undefined');
    }
  }, 2000);
});
asyncSeriesLoopHook.tapAsync('z', (name, callback) => {
  console.log('z done ', name);
  setTimeout(() => {
    console.log('z done setTimeout 3s');
    callback();
  }, 3000);
});

asyncSeriesLoopHook.callAsync('callAsync', (...args) => {
  console.log('asyncSeriesLoopHook.callAsync', args);
});
/*
$ node AsyncSeriesLoopHook.js
x done  callAsync
x done setTimeout 1s
x loop  callAsync 2
x done  callAsync
x done setTimeout 1s
x done  callAsync 6
y done  callAsync
y done setTimeout 2s
y loop  callAsync 0
x done  callAsync
x done setTimeout 1s
x loop  callAsync 3
x done  callAsync
x done setTimeout 1s
x done  callAsync 8
y done  callAsync
y done setTimeout 2s
y done  callAsync 13
z done  callAsync
z done setTimeout 3s
asyncSeriesLoopHook.callAsync []
*/
```

### 拦截器

> 拦截器使用 intercept 方法注册, 在钩子注册, 调用过程中触发

- call: (...args) => void
  向拦截器添加调用方法将在调用钩子时触发, 可以访问钩子参数
- tap: (tap: Tap) => void
  向拦截器添加点击方法将在执行钩子时触发, 可以访问 tap 对象, 无法更改对象
- loop: (...args) => void
  向拦截器添加循环方法将在每个循环钩子触发时触发
- register: (tap: Tap)=> Tap | undefined
  向拦截器添加一个注册器方法将在每次注册钩子时触发, 并可以修改 tap 对象
- Context
  插件和拦截器可以访问的上下文对象, 该对象可以将任意值传递给后续的插件和拦截器

```javascript
myHooks.intercept({
  context: true,
  tap: (context, tapInfo) => {
    console.log(context, tapInfo);
  },
});
```

```javascript
// 示例:
const asyncParallelHook = new AsyncParallelHook(['name']);

asyncParallelHook.intercept({
  call: (...args) => {
    console.log('asyncParallelHook intercept call ', args);
  },
  tap: (...args) => {
    console.log('asyncParallelHook intercept tap ', args);
  },
});
asyncParallelHook.tapAsync('x', (name, callback) => {
  console.log('x done ', name);
  setTimeout(() => {
    console.log('x done setTimeout 2s');
    callback();
  }, 2000);
});
asyncParallelHook.tapAsync('y', (name, callback) => {
  console.log('y done ', name);
  setTimeout(() => {
    console.log('y done setTimeout 5s');
    callback();
  }, 5000);
});
asyncParallelHook.callAsync('callAsync', () => {
  console.log('asyncParallelHook.callAsync');
});
/*
// 输出:
asyncParallelHook intercept call  [ 'callAsync' ]
asyncParallelHook intercept tap  [ { type: 'async', fn: [Function], name: 'x' } ]
x done  callAsync
asyncParallelHook intercept tap  [ { type: 'async', fn: [Function], name: 'y' } ]
y done  callAsync
x done setTimeout 2s
y done setTimeout 5s
asyncParallelHook.callAsync
*/
```

### 辅助函数

#### HookMap

将 hook 分组，方便 hook 组批量调用

```javascript
// 示例:
const { HookMap, SyncHook } = require('tapable');

const hookMap = new HookMap((key) => new SyncHook(['name']));
// 将 hook 分组，方便 hook 组批量调用
hookMap.for('key1').tap('x', (name) => {
  console.log('key1-x ', name);
});
hookMap.for('key1').tap('y', (name) => {
  console.log('key1-y ', name);
});
hookMap.for('key2').tap('z', (name) => {
  console.log('key2-z ', name);
});
hookMap.get('key1').call('call');
/*
// 输出:
key1-x  call
key1-y  call
*/
```

#### MultiHook

向 hook 批量注册钩子函数

```javascript
// 示例:
const { MultiHook, SyncHook, SyncBailHook } = require('tapable');

const multiHook = new MultiHook([new SyncHook(['name']), new SyncBailHook(['name'])]);
// 向 hook 批量注册钩子函数
multiHook.tap('plugin', (name) => {
  console.log('multiHook plugin ');
});
Array.prototype.forEach.call(multiHook.hooks, (hooks) => {
  hooks.callAsync('multiHook.hooks.call ', () => {
    console.log('hooks.callAsync ');
  });
});
/*
// 输出:
multiHook plugin
hooks.callAsync
multiHook plugin
hooks.callAsync
*/
```
