---
title: JSUtils
date: 2022-02-19 17:54:09
categories:
  - tools
tags:
  - API
  - js
---

#### 函数防抖

```javascript
/**
 * @method debounce 函数防抖
 * @param {Function} fn 执行方法
 * @param {Number} delay 延迟时间 默认 300 毫秒
 * @returns {Function}
 */
const debounce = function (fn, delay) {
  if (typeof fn !== 'function') throw new Error('fn is not Function');
  delay = delay >= 0 ? delay : 300;

  let timer;
  return function () {
    let context = this;
    let args = arguments;

    if (timer) clearTimeout(timer);

    timer = setTimeout(() => {
      fn.apply(context, args);
    }, delay);
  };
};
```

#### 函数节流

```javascript
/**
 * @method throttle 函数节流
 * @param {Function} fn 执行方法
 * @param {Number} delay 延迟时间 默认 300 毫秒
 * @returns {Function}
 */
const throttle = function (fn, delay) {
  if (typeof fn !== 'function') throw new Error('fn is not Function');
  delay = delay >= 0 ? delay : 300;

  var previous = 0;
  return function () {
    var _this = this;
    var args = arguments;
    var now = new Date();
    if (now - previous > delay) {
      fn.apply(_this, args);
      previous = now;
    }
  };
};
```

<!-- more -->

#### 复制剪贴板

```javascript
/**
 * @method copyText 剪贴板方法
 * @param {String} text 需要剪贴的内容
 * @returns {Boolean}
 */
const copyText = function (text) {
  // 数字没有 .length 不能执行selectText 需要转化成字符串
  const textString = text.toString();
  let input = document.querySelector('#copy-input');
  if (!input) {
    input = document.createElement('input');
    input.id = 'copy-input';
    input.readOnly = 'readOnly'; // 防止ios聚焦触发键盘事件
    input.style.position = 'absolute';
    input.style.left = '-10000px';
    input.style.zIndex = '-1000';
    document.body.appendChild(input);
  }

  input.value = textString;
  // ios必须先选中文字且不支持 input.select();
  selectText(input, 0, textString.length);
  console.log(document.execCommand('copy'), 'execCommand');
  if (document.execCommand('copy')) {
    document.execCommand('copy');
  }
  input.blur();

  // input自带的select()方法在苹果端无法进行选择，所以需要自己去写一个类似的方法
  // 选择文本。createTextRange(setSelectionRange)是input方法
  function selectText(textbox, startIndex, stopIndex) {
    if (textbox.createTextRange) {
      //ie
      const range = textbox.createTextRange();
      range.collapse(true);
      range.moveStart('character', startIndex); //起始光标
      range.moveEnd('character', stopIndex - startIndex); //结束光标
      range.select(); //不兼容苹果
    } else {
      //firefox/chrome
      textbox.setSelectionRange(startIndex, stopIndex);
      textbox.focus();
    }
    return true;
  }
};
```

#### 检查身份证号

```javascript
/**
 * @method checkCardNo 剪贴板方法
 * @param {String} value 身份证号码
 * @returns {Boolean}
 */
const checkCardNo = function (value) {
  let reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/;
  return reg.test(value);
};
```

#### 检查手机号

```javascript
/**
 * @method checkTel 检查手机号
 * @param {String|Number} value 手机号
 * @returns {Boolean}
 */
const checkTel = function (value) {
  return /^1[3,4,5,6,7,8,9][0-9]{9}$/.test(value.toString());
};
```

#### 检查邮箱

```javascript
/**
 * @method checkEmail 检查邮箱格式
 * @param {String} value 邮箱
 * @returns {Boolean}
 */
const checkEmail = function (value) {
  return /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test(value);
};
```

#### 数字千分位分割

```javascript
/**
 * @method numberFormat 数字千分位分割
 * @param {String|Number} value 数字
 * @param {String} delimiter 千分位分隔符, 默认为 ','
 * @returns {String}
 */
const numberFormat = function (value, delimiter = ',') {
  if (typeof value === 'undefined' || (typeof value === 'object' && value === null)) {
    return value;
  }

  try {
    let num = value.toString();
    let len = num.length;
    if (len <= 3) {
      return num;
    }
    // 判断数字的长度是否为 3 的倍数
    let remainder = len % 3;
    return remainder > 0
      ? num.slice(0, remainder) + delimiter + num.slice(remainder, len).match(/\d{3}/g).join(delimiter)
      : num.slice(0, len).match(/\d{3}/g).join(delimiter);
  } catch (error) {
    return value;
  }
};
```

#### 根据图片像素计算图片宽高比例

```javascript
/**
 * @method calcImgRatio 根据图片像素计算图片宽高比例
 * @param {Number} width 宽度
 * @param {Number} height 高度
 * @returns {String} '16|9'
 */
const calcImgRatio = function (width, height) {
  if (isNaN(width) || typeof width !== 'number' || isNaN(height) || typeof height !== 'number') {
    throw new TypeError('Params are not Number type....');
  }

  let num = getMaxCommonDivisor(width, height);

  return `${width / num}|${height / num}`;

  // 获取两个数的最大公约数
  function getMaxCommonDivisor(a, b) {
    if (a % b) {
      return getMaxCommonDivisor(b, a % b);
    }
    return b;
  }
};
```
