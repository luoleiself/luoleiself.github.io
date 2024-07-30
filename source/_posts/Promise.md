---
title: Promise
date: 2021-06-26 16:10:04
categories:
  - ES
tags:
  - js
  - Promise
  - ES6
---

### 概念

> 是异步编程的一种解决方案, 解决 js 异步回调地狱的问题

- 状态唯一

  - pending: 初始状态, 既没有被兑现, 也没有被拒绝
  - fulfilled: 操作成功完成
  - rejected: 操作失败

- 状态不受外界影响

### 简单用法

```javascript
let p = new Promise((resolve, reject) => {
  resolve(200);
});
p.then((res) => {
  console.log(res);
});
```

<!-- more -->

### 手动实现

#### 基础代码

```javascript
// 状态常量
const PENDING = 'pending';
const FULFILLED = 'fulfilled';
const REJECTED = 'rejected';
// 创建一个 MyPromise 类
class MyPromise {
  // 储存状态的变量，初始值是 pending
  status = PENDING;
  // 成功之后的值
  value = null;
  // 失败之后的原因
  reason = null;

  // 存储成功回调函数
  onFulfilledCallbacks = [];
  // 存储失败回调函数
  onRejectedCallbacks = [];

  constructor(executor) {
    let resolve = (value) => {
      if (this.status === PENDING) {
        // 状态修改为成功
        this.status = FULFILLED;
        // 保存成功之后的值
        this.value = value;
        while (this.onFulfilledCallbacks.length) {
          // shift操作改变原数组, 循环从队列取出第一个方法执行
          this.onFulfilledCallbacks.shift()(value);
        }
      }
    };
    let reject = (reason) => {
      if (this.status === PENDING) {
        // 状态成功为失败
        this.status = REJECTED;
        // 保存失败后的原因
        this.reason = reason;
        while (this.onRejectedCallbacks.length) {
          // shift操作改变原数组, 循环从队列取出第一个方法执行
          this.onRejectedCallbacks.shift()(reason);
        }
      }
    };

    try {
      executor(resolve, reject);
    } catch (error) {
      reject(error); // 失败执行 reject 方法
    }
  }

  then(onFulfilled, onRejected) {
    const realOnFulfilled =
      typeof onFulfilled === 'function' ? onFulfilled : (value) => value;
    const realOnRejected =
      typeof onRejected === 'function' ? onRejected : (reason) => throw Error(reason);
    if (this.status === FULFILLED) {
      realOnFulfilled(this.value);
    }
    if (this.status === REJECTED) {
      realOnRejected(this.reason);
    }
    if (this.status === PENDING) {
      // 等待状态将成功和失败回调存储起来
      this.onFulfilledCallbacks.push(() => realOnFulfilled(this.value));
      this.onRejectedCallbacks.push(() => realOnRejected(this.reason));
    }
  }

  catch(onRejected) {
    const realOnRejected =
      typeof onRejected === 'function' ? onRejected : (reason) => throw Error(reason);
    if (this.status === REJECTED) {
      realOnRejected(this.reason);
    }
    if (this.status === PENDING) {
      this.onRejectedCallbacks.push(() => realOnRejected(this.reason));
    }
  }
}
```

```javascript
// 使用
var p = new MyPromise((resolve, reject) => {
  setTimeout(() => resolve('gg'), 5000);
}).then(
  (res) => console.log('res', res),
  (err) => console.log('err', err)
);
// 5秒后输出: 'res' gg
```

#### 静态方法

```javascript
// resolve 静态方法
static resolve (parameter) {
  // 如果传入 MyPromise 就直接返回
  if (parameter instanceof MyPromise) {
    return parameter;
  }

  // TODO 如果是一个 thenable 对象, 则通过传入一对解决函数作为参数调用该 thenable 对象的 then 方法后
  // 得到的状态作为返回的 Promise 对象的状态
  if(Object.prototype.toString.call(parameter) === '[object Object]' && typeof parameter.then === 'function') {
    return new MyPromise(parameter.then);
  }

  // 转成常规方式
  return new MyPromise(resolve => {
    resolve(parameter);
  });
}

// reject 静态方法, 不管是否是 Promise 都将返回新的 Promise
static reject (reason) {
  return new MyPromise((resolve, reject) => {
    reject(reason);
  });
}
```

#### 链式调用

```javascript
then(onFulfilled, onRejected) {
  const realOnFulfilled = typeof onFulfilled === 'function' ? onFulfilled : value => value;
  const realOnRejected = typeof onRejected === 'function' ? onRejected : reason => throw Error(reason);

  // 为了链式调用这里直接创建一个 MyPromise，并在后面 return 出去
  const promise2 = new MyPromise((resolve, reject) => {
    const fulfilledMicrotask = () =>  {
      // 创建一个微任务等待 promise2 完成初始化
      queueMicrotask(() => {
        try {
          // 获取成功回调函数的执行结果
          const x = realOnFulfilled(this.value);
          // 传入 resolvePromise 集中处理
          resolvePromise(promise2, x, resolve, reject);
        } catch (error) {
          reject(error)
        }
      })
    }

    const rejectedMicrotask = () => {
      // 创建一个微任务等待 promise2 完成初始化
      queueMicrotask(() => {
        try {
          // 调用失败回调，并且把原因返回
          const x = realOnRejected(this.reason);
          // 传入 resolvePromise 集中处理
          resolvePromise(promise2, x, resolve, reject);
        } catch (error) {
          reject(error)
        }
      })
    }
    // 判断状态
    if (this.status === FULFILLED) {
      fulfilledMicrotask()
    } else if (this.status === REJECTED) {
      rejectedMicrotask()
    } else if (this.status === PENDING) {
      // 等待
      // 因为不知道后面状态的变化情况，所以将成功回调和失败回调存储起来
      // 等到执行成功失败函数的时候再传递
      this.onFulfilledCallbacks.push(fulfilledMicrotask);
      this.onRejectedCallbacks.push(rejectedMicrotask);
    }
  })

  return promise2;
}
// 处理 MyPromise 实例之间的关系
function resolvePromise(promise2, x, resolve, reject) {
  // 如果相等了，说明return的是自己，抛出类型错误并返回
  if (promise2 === x) {
    return reject(new TypeError('Chaining cycle detected for promise #<Promise>'))
  }
  // 判断x是不是 MyPromise 实例对象
  if(x instanceof MyPromise) {
    // 执行 x，调用 then 方法，目的是将其状态变为 fulfilled 或者 rejected
    // x.then(value => resolve(value), reason => reject(reason))
    // 简化之后
    x.then(resolve, reject)
  } else{
    // 普通值
    resolve(x)
  }
}
```

#### 完整代码

```javascript
// 状态常量
const PENDING = 'pending';
const FULFILLED = 'fulfilled';
const REJECTED = 'rejected';
// 创建一个 MyPromise 类
class MyPromise {
  // 储存状态的变量，初始值是 pending
  status = PENDING;
  // 成功之后的值
  value = null;
  // 失败之后的原因
  reason = null;

  // 存储成功回调函数
  onFulfilledCallbacks = [];
  // 存储失败回调函数
  onRejectedCallbacks = [];

  constructor(executor) {
    let resolve = (value) => {
      if (this.status === PENDING) {
        // 状态修改为成功
        this.status = FULFILLED;
        // 保存成功之后的值
        this.value = value;
        while (this.onFulfilledCallbacks.length) {
          // shift操作改变原数组, 循环从队列取出第一个方法执行
          this.onFulfilledCallbacks.shift()(value);
        }
      }
    };
    let reject = (reason) => {
      if (this.status === PENDING) {
        // 状态成功为失败
        this.status = REJECTED;
        // 保存失败后的原因
        this.reason = reason;
        while (this.onRejectedCallbacks.length) {
          // shift操作改变原数组, 循环从队列取出第一个方法执行
          this.onRejectedCallbacks.shift()(reason);
        }
      }
    };

    try {
      executor(resolve, reject);
    } catch (error) {
      reject(error); // 失败执行 reject 方法
    }
  }
  // resolve 静态方法
  static resolve(parameter) {
    // 如果传入 MyPromise 就直接返回
    if (parameter instanceof MyPromise) {
      return parameter;
    }

    // 转成常规方式
    return new MyPromise((resolve) => {
      resolve(parameter);
    });
  }

  // reject 静态方法
  static reject(reason) {
    return new MyPromise((resolve, reject) => {
      reject(reason);
    });
  }

  then(onFulfilled, onRejected) {
    const realOnFulfilled =
      typeof onFulfilled === 'function' ? onFulfilled : (value) => value;
    const realOnRejected =
      typeof onRejected === 'function' ? onRejected : (reason) => throw Error(reason);

    // 为了链式调用这里直接创建一个 MyPromise，并在后面 return 出去
    const promise2 = new MyPromise((resolve, reject) => {
      const fulfilledMicrotask = () => {
        // 创建一个微任务等待 promise2 完成初始化
        queueMicrotask(() => {
          try {
            // 获取成功回调函数的执行结果
            const x = realOnFulfilled(this.value);
            // 传入 resolvePromise 集中处理
            resolvePromise(promise2, x, resolve, reject);
          } catch (error) {
            reject(error);
          }
        });
      };

      const rejectedMicrotask = () => {
        // 创建一个微任务等待 promise2 完成初始化
        queueMicrotask(() => {
          try {
            // 调用失败回调，并且把原因返回
            const x = realOnRejected(this.reason);
            // 传入 resolvePromise 集中处理
            resolvePromise(promise2, x, resolve, reject);
          } catch (error) {
            reject(error);
          }
        });
      };
      // 判断状态
      if (this.status === FULFILLED) {
        fulfilledMicrotask();
      } else if (this.status === REJECTED) {
        rejectedMicrotask();
      } else if (this.status === PENDING) {
        // 等待
        // 因为不知道后面状态的变化情况，所以将成功回调和失败回调存储起来
        // 等到执行成功失败函数的时候再传递
        this.onFulfilledCallbacks.push(fulfilledMicrotask);
        this.onRejectedCallbacks.push(rejectedMicrotask);
      }
    });

    return promise2;
  }

  catch(onRejected) {
    const realOnRejected =
      typeof onRejected === 'function' ? onRejected : (reason) => throw Error(reason);
    if (this.status === REJECTED) {
      realOnRejected(this.reason);
    }
    if (this.status === PENDING) {
      this.onRejectedCallbacks.push(() => realOnRejected(this.reason));
    }
  }
}

// 处理 MyPromise 实例之间的关系
function resolvePromise(promise2, x, resolve, reject) {
  // 如果相等了，说明return的是自己，抛出类型错误并返回
  if (promise2 === x) {
    return reject(
      new TypeError('Chaining cycle detected for promise #<Promise>')
    );
  }
  // 判断x是不是 MyPromise 实例对象
  if (x instanceof MyPromise) {
    // 执行 x，调用 then 方法，目的是将其状态变为 fulfilled 或者 rejected
    // x.then(value => resolve(value), reason => reject(reason))
    // 简化之后
    x.then(resolve, reject);
  } else {
    // 普通值
    resolve(x);
  }
}
```

### 官方 API

- Promise.resolve() 将给定的值转换为一个 Promise
  - 如果该值本身是一个 Promise, 将该 Promise 返回
  - 如果该值是一个 thenable 对象, 将调用其 then() 方法及其两个回调函数
  - 否则, 返回的 Promise 将以该值兑现
- Promise.reject() 返回一个已拒绝的 Promise 对象, 拒绝原因为给定的参数
- Promise.all() 在**所有**传入的 Promise 都被兑现时兑现, 在**任意一个** Promise 被拒绝时拒绝
- Promise.allSettled() 在**所有**的 Promise 都被敲定时兑现
- Promise.any() 在**任意一个** Promise 被兑现时兑现, 仅在**所有**的 Promise 都被拒绝时才会拒绝
- Promise.race() 在**任意一个** Promise 敲定时敲定, 即在**任意一个** Promise 被兑现时兑现, 在**任意一个** Promise 被拒绝时拒绝
