---
title: 即将废弃或已经废弃的API
date: 2021-07-05 12:38:23
categories:
  - WebAPI
tags:
  - API
---

- [String.prototype.substr()](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/substr)

  > 并非 JavaScript 核心语言的一部分,未来将可能被移除掉, 建议使用 substring() 代替

- [arguments.callee](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/arguments/callee)

  > ES5 严格模式下禁止使用此方法，建议使用方法名调用自己

- [Function.arguments](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Function/arguments) 已废弃

  > 是一个类数组对象,表示传入函数的实参. 在函数内部直接使用 [arguments](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Functions/arguments) 对象访问

- [regexObj.compile](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/RegExp/compile) 已废弃

  > 用于编译正则表达式, 和 RegExp 构造函数的作用基本一样

- [Date.prototype.getYear](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Date/getYear) 已废弃

  > 使用 getFullYear 代替
