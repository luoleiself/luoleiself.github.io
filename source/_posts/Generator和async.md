---
title: Generator和async
date: 2021-07-23 18:03:53
categories:
  - ES
tags:
  - ES6
---

### Generator

> Generator 函数是 ES6 提供的一种异步编程解决方案, 语法行为和传统普通函数完全不同, Generator 函数是一个状态机, 封装了多个内部状态, 执行函数会返回一个迭代器对象, 返回的迭代器对象可以依次遍历 Generator 函数内部的每一个状态.

- function 关键字和函数名之间有一个星号, 星号写在哪个位置都可以
- 函数体内部使用 yield 表达式定义不同的状态

```javascript
function* helloWorld() {
  yield 'hello';
  yield 'world';
}

var gen = helloWorld();
gen.next();
// { value: 'hello', done: false }
gen.next();
// { value: 'world', done: false }
gen.next();
// { value: undefined, done: true }
```

调用 Generator 函数, 返回一个迭代器对象, 代表 Generator 函数的内部指针, 每次调用迭代器对象的 next 方法, 就会返回一个包含 value 和 done 属性的对象, value 属性表示当前的内部状态(yield 表达式)的值, done 属性表示遍历是否结束的 boolean 值

#### yield 表达式

> yield 表达式只能用在 Generator 函数里面, 用在其他地方都会报错

1. 遇到 yield 表达式, 就暂停执行后面的操作, 并将 yield 后面紧跟的表达式的值作为返回对象的 value 属性值
2. 下一次调用 next 方法时, 再继续往下执行, 直到遇到下一个 yield 表达式
3. 如果没有再遇到新的 yield 表达式, 就一直运行到函数结束, 直到 return 语句为止, 并将 return 语句后面的表达式的值作为返回的对象的 value 属性值
4. 如果该函数没有 return 语句, 则返回的对象的 value 属性值为 undefined

<!-- more -->

#### 与 Iterator 接口的关系

> 任意数据结构的 Symbol.iterator 方法,等于该数据结构的迭代器生成函数,调用该方法会返回一个该数据结构的迭代器对象

```javascript
var myIterator = {
  // [Symbol.iterator]: function* () {
  // 简洁写法
  *[Symbol.iterator]() {
    yield 1;
    yield 2;
    yield 3;
  },
};
[...myIterator]; // [1, 2, 3]
```

#### next 方法的参数

> next 方法可以带一个参数,该参数被作为上一个 yield 表达式的结果, 第一次使用 next 方法时, 传递参数会被解释器忽略

```javascript
function* foo(x) {
  var y = 2 * (yield x + 1); // 第一个 yield 表达式
  var z = yield y / 3; // 第二个 yield 表达式
  return x + y + z;
}

var a = foo(5);
a.next(); // Object{value:6, done:false}
a.next(); // Object{value:NaN, done:false}
a.next(); // Object{value:NaN, done:true}

var b = foo(5);
b.next(); // { value:6, done:false }
b.next(12); // { value:8, done:false } // 12 作为第一个 yield 表达式的结果
b.next(13); // { value:42, done:true } // 13 作为第二个 yield 表达式的结果
```

#### for...of 循环

> for...of 循环可以自动循环 Generator 函数运行时生成的 Iterator 对象, 此时不再需要调用 next 方法, next 返回的对象的 done 属性为 true 时则自动中止循环且不包含该返回对象

```javascript
function* foo() {
  yield 1;
  yield 2;
  yield 3;
  yield 4;
  return 5;
}

for (let v of foo()) {
  console.log(v);
}
// 1 2 3 4
```

#### throw

迭代器对象的 throw 方法可以在函数体外抛出错误异常, 在 Generator 函数体内捕获

```javascript
function* foo() {
  try {
    yield;
  } catch (err) {
    console.log('内部捕获', err);
  }
}
var gen = foo();
gen.next();
try {
  gen.throw('err1');
  gen.throw('err2');
} catch (err) {
  console.log('外部捕获', err);
}
// 内部捕获 err1
// 外部捕获 err2
```

#### return

迭代器对象的 return 方法返回给定的值并且中止遍历 Generator 函数, 如果 Generator 函数内部有 try...finally, 并且正在执行 try 内部代码, 则立刻进入 finally 执行直到函数结束

```javascript
function* numbers() {
  yield 1;
  try {
    yield 2;
    yield 3;
  } finally {
    yield 4;
    yield 5;
  }
  yield 6;
}
var g = numbers();
g.next(); // { value: 1, done: false }
g.next(); // { value: 2, done: false }
g.return(7); // { value: 4, done: false }
g.next(); // { value: 5, done: false }
g.next(); // { value: 7, done: true }
```

#### next throw return 的共同点

> 此三个方法本质上是一件事, 共同的作用是让 Generator 函数恢复执行并且使用不同的语句替换 yield 表达式

- next 将 yield 表达式替换成一个值
- throw 将 yield 表达式替换成一个 throw 语句
- return 将 yield 表达式替换成一个 return 语句

#### yield\* 表达式

用于在一个 Generator 函数内执行另外一个 Generator 函数

```javascript
function* bar() {
  yield 'x';
  yield* foo();
  yield 'y';
}

// 等同于
function* bar() {
  yield 'x';
  yield 'a';
  yield 'b';
  yield 'y';
}

// 等同于
function* bar() {
  yield 'x';
  for (let v of foo()) {
    yield v;
  }
  yield 'y';
}

for (let v of bar()) {
  console.log(v);
}
// "x"
// "a"
// "b"
// "y"
```

### async

> ES8 标准, Generator 函数的语法糖

#### 与 Generator 函数的区别

##### 内置执行器

Generator 函数的自动执行需要依赖执行器 co 模块, async 函数自带执行器只要调用即可

##### 更好的语义

async 表示函数内有异步操作, await 表示后面的表达式需要等待结果

##### 更广的适用性

co 模块规定 yield 命令后面只能是 thunk 函数或者 Promise 对象, async 函数的 await 命令后面可以是 Promise 对象和原始类型的值(数值、字符串和布尔值,此时会自动转成 Resolve 的 Promise 对象)

##### 返回值是 Promise

async 的返回值是 Promise, 可以使用 then 方法指定下一步的操作

```javascript
async function(){
  await '1';
  await '2';
}
```

#### async 的实现原理

将 Generator 函数和自动执行器包装在一个函数内

#### await

##### Promise 对象

返回该对象的结果

##### 定义了 then 方法的对象

await 将当作 Promise 对象执行

##### reject 状态

如果 Promise 对象变成 reject 状态,则整个 async 函数都会中断执行
