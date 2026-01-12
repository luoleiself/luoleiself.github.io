---
title: JSUtils
date: 2022-02-19 17:54:09
categories:
  - tools
tags:
  - API
  - js
---

## 函数防抖

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

## 函数节流

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

## 复制剪贴板

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

## 正则篇

### 检查身份证号

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

### 检查手机号

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

### 检查邮箱

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

## 数字千分位分割

- 使用 Number 包装对象的原型方法 toLocaleString
  - locales 指定地区语言, 默认使用系统的区域设置
  - options 对返回结果进行定制

```javascript
console.log(Number(2000).toLocaleString());
// 2,000

console.log(Number(123456.789).toLocaleString("zh-Hans-CN-u-nu-hanidec"));
// 一二三,四五六.七八九

// 定制货币格式
console.log(Number(123456.789).toLocaleString("de-DE", { style: "currency", currency: "EUR" }));
// 123.456,79 €

// 限制有效数字位数
console.log( Number(123456.789).toLocaleString("de-DE", { 
    style: "currency", 
    currency: "EUR", 
    maximumSignificantDigits: 5 
  })
);
// 123.460 €
```

- 手动实现位数截取

```javascript
/**
 * @method numberFormat 数字千分位分割
 * @param {String|Number} value 数字
 * @param {String} delimiter 千分位分隔符, 默认为 ','
 * @returns {String}
 */
const numberFormat = function (value, delimiter = ',') {
  if (
    typeof value === 'undefined' ||
    (typeof value === 'object' && value === null)
  ) {
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
      ? num.slice(0, remainder) +
          delimiter +
          num.slice(remainder, len).match(/\d{3}/g).join(delimiter)
      : num.slice(0, len).match(/\d{3}/g).join(delimiter);
  } catch (error) {
    return value;
  }
};
```

## 根据图片尺寸计算宽高比例

```javascript
/**
 * @method calcImgRatio 根据图片像素计算图片宽高比例
 * @param {Number} width 宽度
 * @param {Number} height 高度
 * @returns {String} '16|9'
 */
const calcImgRatio = function (width, height) {
  if (
    isNaN(width) ||
    typeof width !== 'number' ||
    isNaN(height) ||
    typeof height !== 'number'
  ) {
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

## Base64 图片转 File 文件

```javascript
/**
 * @method  dataURL2File Base64 图片转 File 文件
 * @param {String} dataUrl 图片 Base64 字符串编码
 * @param {String} fileName File 文件对象的 name 属性, 包含文件扩展名, 非必传
 * @return {File}  返回 File 文件对象
 */
function dataURL2File(dataUrl, fileName) {
  if (!/data:image\/.*?;base64,/.test(dataUrl)) {
    throw new Error('dataUrl 参数格式错误...');
  }
  if (!fileName) {
    fileName =
      Math.random().toString(16).substring(2) +
      '.' +
      dataUrl.match(/:(.*?);/i)[1].split('/')[1];
  }

  var arr = dataUrl.split(','),
    mime = arr[0].match(/:(.*?);/)[1],
    bstr = atob(arr[1]),
    n = bstr.length,
    u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new File([u8arr], fileName, { type: mime });
}
```

## 图片对象或链接转 Base64

```javascript
/**
 * @method  file2DataURL 图片对象或图片链接转 Base64
 * @param {String|File} arg 图片对象或者图片链接
 * @return {Promise}  返回 Promise 的 then 方法接收图片 Base64 编码字符串
 */
function file2DataURL(arg) {
  var isImgUrl = /((https?)?:)?\/\/.*?\.(jpg|jpeg|bmp|png|webp|gif)/i.test(arg);
  var isFile =
    arg != null && typeof arg == 'object' && /image\/\w+/.test(arg.type);

  if (!isFile && !isImgUrl) {
    throw new Error('arg 参数不是一个图片对象或者图片链接...');
  }

  if (isImgUrl) {
    var mime = 'image/' + arg.match(/\.(jpg|jpeg|bmp|png|webp|gif)/i)[1];
    var canvas = document.createElement('canvas');
    var img = document.createElement('img');
    img.src = arg.includes('http')
      ? arg
      : arg.includes('://')
      ? 'https' + arg
      : 'https:' + arg;
    img.crossOrigin = 'anonymous';

    return new Promise(function (resolve, reject) {
      img.onload = function () {
        canvas.width = img.width;
        canvas.height = img.height;
        var ctx = canvas.getContext('2d');
        ctx.drawImage(img, 0, 0, img.width, img.height);
        var dataURL = canvas.toDataURL(mime);
        resolve(dataURL);
      };
      img.onerror = function (err) {
        reject(err);
      };
    });
  }

  if (typeof FileReader == 'undefined') {
    throw new Error('浏览器不支持 FileReader API, 请先升级浏览器...');
  }
  var reader = new FileReader();
  var p = new Promise(function (resolve, reject) {
    reader.onload = function (e) {
      resolve(e.currentTarget.result);
    };
    reader.onerror = function (e) {
      reject(e);
    };
  });
  reader.readAsDataURL(arg);
  return p;
}
```

## DOM 拖动

```javascript
var $popDom = null; // 指定拖动的 DOM
var disX,
  disY,
  dragging = false;
$popDom.on('mousedown', function (evt) {
  evt.stopPropagation();
  var popWidth = $popDom[0].clientWidth;
  var popHeight = $popDom[0].clientHeight;
  disX = evt.pageX - Number.parseInt($popDom.css('left'), 10);
  disY = evt.pageY - Number.parseInt($popDom.css('top'), 10);
  dragging = true;

  $(document)
    .on('mousemove.move-dom', function (evt) {
      evt.stopPropagation();
      evt.preventDefault();
      if (!dragging) {
        return;
      }
      var clientWidth = document.documentElement.clientWidth;
      var clientHeight = document.documentElement.clientHeight;

      var videoPopTop = Number.parseInt(evt.pageY - disY, 10);
      var videoPopLeft = Number.parseInt(evt.pageX - disX, 10);

      if (videoPopTop < 0) {
        videoPopTop = 0;
      } else if (clientHeight - videoPopTop < popHeight) {
        videoPopTop = clientHeight - popHeight;
      }

      if (videoPopLeft < 0) {
        videoPopLeft = 0;
      } else if (clientWidth - videoPopLeft < popWidth) {
        videoPopLeft = clientWidth - popWidth;
      }
      $popDom.css({ top: videoPopTop, left: videoPopLeft });
    })
    .on('mouseup.move-dom', function (evt) {
      evt.stopPropagation();
      dragging = false;
      $(document).off('mousemove.move-dom');
      $(document).off('mouseup.move-dom');
    });
});
```

## 日期倒计时

```javascript
(function () {
  if (
    (Object.hasOwn && Object.hasOwn(Date.prototype, 'leftDown')) ||
    Date.prototype.hasOwnProperty('leftDown')
  ) {
    throw new Error('倒计时方法名重复了...想别的招儿吧!');
  }

  function gtNine(val) {
    return val > 9 ? `${val}` : `0${val}`;
  }

  function hasOwnProp(obj, prop) {
    return obj !== null && typeof obj === 'object' && obj.hasOwnProperty(prop);
  }
  /**
   * @method leftDown 倒计时方法
   * @description Date原型对象的方法, 所有的日期时间对象都可调用
   * @param {function} callback 时间计算过程返回的差值对象
   *    @param {days: '00', hours: '00', minutes: '00', seconds: '00'}
   * @property paused 只读属性, 是否暂停倒计时
   * @property resumed 只读属性, 是否恢复倒计时
   * @property canceled 只读属性, 是否取消倒计时
   * @property kept 只读属性, 是否跟随上一次暂停时间继续计算
   * @property left 只读属性, 获取倒计时当前计算后的差值
   * @returns Object
   *    @javascript { cancel: cancel, pause: pause, resume: resume }
   *    @method cancel 取消倒计时
   *    @method pause 暂停倒计时
   *    @method resume 恢复倒计时,  keep 参数控制是否跟随上一次暂停时间继续计算, 默认为 true
   */
  Date.prototype.leftDown = function (callback) {
    if (
      hasOwnProp(this, 'paused') &&
      hasOwnProp(this, 'resumed') &&
      hasOwnProp(this, 'canceled') &&
      hasOwnProp(this, 'kept')
    ) {
      console.warn(
        'leftDown 方法不能在一个实例上重复调用, 可以使用返回的 pause, cancel, resume 方法控制当前倒计时, 或者创建新的 Date 实例调用此方法'
      );
      return;
    }

    var _timer = null;
    var _referrerTime = Date.now();
    var _paused = false; // 是否暂停
    var _canceled = false; // 是否取消
    var _resumed = false; // 是否恢复
    var _keep = true; // 是否跟随上一次暂停时间继续计算
    var obj = { days: '00', hours: '00', minutes: '00', seconds: '00' };
    var evt = new Event('leftDown', { bubbles: true, cancelable: true });
    var evtPause = new Event('leftDown:pause', { bubbles: true, cancelable: true });
    var evtResume = new Event('leftDown:resume', { bubbles: true, cancelable: true });
    var evtCancel = new Event('leftDown:cancel', { bubbles: true, cancelable: true });
    evt.leftDown = obj;

    var getDiff = function () {
      _referrerTime = _keep ? _referrerTime + 1000 : Date.now();
      return this.getTime() - _referrerTime;
    }.bind(this);

    var calcLeftTimeFn = function () {
      var leftTime = getDiff();
      if (leftTime < 0) {
        clearInterval(_timer);
        _timer = null;
        _canceled = true;
        evtCancel.leftDown = evt.leftDown = obj;
        window && window.dispatchEvent(evt);
        window && window.dispatchEvent(evtCancel);
        return typeof callback === 'function' && callback(obj);
      }
      var d = Math.floor(leftTime / 1000 / 60 / 60 / 24);
      var h = Math.floor((leftTime / 1000 / 60 / 60) % 24);
      var m = Math.floor((leftTime / 1000 / 60) % 60);
      var s = Math.floor((leftTime / 1000) % 60);

      obj['days'] = gtNine(d);
      obj['hours'] = gtNine(h);
      obj['minutes'] = gtNine(m);
      obj['seconds'] = gtNine(s);
      evt.leftDown = obj;
      window && window.dispatchEvent(evt);
      typeof callback === 'function' && callback(obj);
    };

    if (getDiff() <= 0) {
      _canceled = true;
      evtCancel.leftDown = evt.leftDown = obj;
      window && window.dispatchEvent(evt);
      window && window.dispatchEvent(evtCancel);
      return typeof callback === 'function' && callback(obj);
    }

    _timer = setInterval(calcLeftTimeFn.bind(this), 1000);

    Object.defineProperties(this, {
      paused: {
        get() { return _paused },
        set(newVal) {
          console.log('paused is the read-only property');
          return 'paused is the read-only property';
        },
      },
      resumed: {
        get() { return _resumed },
        set(newVal) {
          console.log('resumed is the read-only property');
          return 'resumed is the read-only property';
        },
      },
      canceled: {
        get() { return _canceled },
        set(newVal) {
          console.log('canceled is the read-only property');
          return 'canceled is the read-only property';
        },
      },
      kept: {
        get() { return _keep },
        set(newVal) {
          console.log('kept is the read-only property');
          return 'kept is the read-only property';
        },
      },
      left: {
        get() { return obj },
        set() {
          console.log('left is the read-only property');
          return 'left is the read-only property';
        },
      },
    });
    // 取消倒计时
    var cancel = function () {
      if(_canceled) return;
      _canceled = true;
      if (_timer) {
        clearInterval(_timer);
        _timer = null;
      }
      evtCancel.leftDown = obj;
      window && window.dispatchEvent(evtCancel);
      typeof callback === 'function' && callback(obj);
    }.bind(this);
    // 暂停倒计时
    var pause = function () {
      if(_canceled || _paused) return;
      _paused = true;
      _resumed = !_paused;
      if (_timer) {
        clearInterval(_timer);
        _timer = null;
      }
      evtPause.leftDown = obj;
      window && window.dispatchEvent(evtPause);
      typeof callback === 'function' && callback(obj);
    }.bind(this);
    // 是否跟随上一次暂停时间继续计算, 默认为 true, 传入 false 则按照当前时间重新计算
    var resume = function (keep = true) {
      if(_canceled || _resumed) return;
      if(_timer){
        clearInterval(_timer);
        _timer = null;
      }
      _resumed = true;
      _paused = !_resumed;
      _keep = keep === false ? false : true;
      calcLeftTimeFn();
      _timer = setInterval(calcLeftTimeFn.bind(this), 1000);
      evtResume.leftDown = obj;
      window && window.dispatchEvent(evtResume);
      typeof callback === 'function' && callback(obj);
    }.bind(this);

    return { cancel: cancel, pause: pause, resume: resume };
  };
})();

// 演示用效果方法
window.randRGBA = function (a) {
  let r = Math.floor(Math.random() * 256);
  let g = Math.floor(Math.random() * 256);
  let b = Math.floor(Math.random() * 256);
  a =
    a >= 1
      ? 1
      : Math.random()
          .toString()
          .match(/\d\.\d{1}/)[0];
  return `rgba(${r},${g},${b},${a})`;
};

var date = new Date(Date.now() + 1000000);
// 方式一：回调函数的参数
var { pause, resume, cancel } = date.leftDown(function (obj) {
  console.log(obj);
  document.body.style.color = randRGBA(1);
});
// // 方式二：事件监听回调函数参数的 leftDown 字段获取
// window.addEventListener('leftDown', function (evt) {
//   console.log(evt)
// });
// window.addEventListener('leftDown:pause', function (evt) {
//   console.log(evt)
// });
// window.addEventListener('leftDown:resume', function (evt) {
//   console.log(evt)
// });
// window.addEventListener('leftDown:cancel', function (evt) {
//   console.log(evt)
// });
```

## JS 批量样式

```javascript
// Dom 添加边框
[].forEach.call($$('*'), (dom) => {
  dom.style.outline =
    '1px solid #' + (~~(Math.random() * (1 << 24))).toString(16);
});

// 随机生成颜色
() =>
  '#' +
  Math.floor(Math.random() * 0xffffff)
    .toString(16)
    .padEnd(6, '0')
    .toUpperCase();
```

## 修改 history

```javascript
// 修改历史记录阻止后退功能
$(document).ready(function () {
  if (window.history && window.history.pushState) {
    $(window).on('popstate', function () {
      window.history.pushState('forward', null, '');
      window.history.forward(1);
    });
  }
  window.history.pushState('forward', null, ''); //在IE中必须得有这两行
  window.history.forward(1);
});
```

## 存储单位转换

```javascript
function formatSizeUnits(kb) {
  let units = ['KB', 'MB', 'GB', 'TB', 'PB'];
  let unitIndex = 0;

  while (kb >= 1024 && unitIndex < units.length - 1) {
    kb /= 1024;
    unitIndex++;
  }

  return `${kb.toFixed(2)} ${units[unitIndex]}`;
}
```

## requestAnimationFrame 数字累加

```javascript
function roll(o, s , e){
  if (!e) {
    const r = {total: 3e3, start: performance.now() };
    const l = function (e) {
      var t, n, i, e = e - r.start, a = (t = e,
        n = 0,
        a = 1,
        i = r.total,
        -a * ((t = t / i - 1) * t * t * t - 1) + n);
      
      s.innerHTML = Math.round(a * o).toLocaleString();
      e < r.total && window.requestAnimationFrame(l);
      // console.log("t=", t, "n=", n, "i=", i, "e=", e, "a=", a, "o=", o, "s=", s);
    };
    window.requestAnimationFrame(l)
  }
}
```

## JS 创建 Element

```javascript
/*
继承关系
EventTarget
  <- Node
    <- Document
    <- DocumentFragment
    <- Element
      <- HTMLElement
      <- SVGElement
    <- [Attr](https://developer.mozilla.org/zh-CN/docs/Web/API/Attr)
    <- CharacterData
      <- Comment 
*/
/* is.js */
export const array = Array.isArray;
export function undefinedOrNull(v) {
  return v === undefined || v === null;
}
export function primitive(v) {
  return (
    typeof v === 'string' ||
    typeof v === 'number' ||
    typeof v === 'boolean' ||
    v instanceof String ||
    v instanceof Number ||
    v instanceof Boolean
  );
}
export function string(v) {
  return Object.prototype.toString.call(v) === '[object String]';
}
export function object(v) {
  return v !== null && Object.prototype.toString.call(v) === '[object Object]';
}
export function fn(v) {
  return Object.prototype.toString.call(v) === '[object Function]';
}
export function promise(v) {
  return Object.prototype.toString.call(v) === '[object Promise]';
}
export function asyncFn(v) {
  return Object.prototype.toString.call(v) === '[object AsyncFunction]';
}
export function generatorFn(v) {
  return Object.prototype.toString.call(v) === '[object GeneratorFunction]';
}

import * as is from './is.js';

export function h(tag, props, child) {
  var children,
    attrs,
    el =
      is.undefinedOrNull(tag) || (is.string(tag) && tag.trim() === '')
        ? document.createDocumentFragment()
        : document.createElement(tag);

  if (child !== undefined && child !== null) {
    if (is.object(props)) {
      attrs = props;
    }
  } else if (props !== undefined && props !== null) {
    if (is.object(props)) {
      attrs = props;
    } else {
      child = props;
    }
  }

  if (is.primitive(child)) {
    children = [document.createTextNode(child)];
  } else if (child instanceof HTMLElement) {
    children = [child];
  } else if (is.array(child)) {
    for (var i = 0; i < child.length; i++) {
      if (is.primitive(child[i])) {
        child[i] = h(undefined, child[i]);
      }
    }
    children = child;
  }

  if (children !== undefined) {
    for (var i = 0; i < children.length; i++) {
      el.appendChild(children[i]);
    }
  }

  if (is.object(attrs)) {
    setElAttrs(el, attrs);
  }

  return el;
}
function setElAttrs(el, props) {
  if (
    !(el instanceof HTMLElement) ||
    !(el instanceof Node) ||
    !is.object(props)
  ) {
    return false;
  }
  for (var key in props) {
    if (key === 'cls' || key === 'class') {
      if (!is.string(props[key]) && !is.array(props[key])) {
        continue;
      }
      el.setAttribute(
        'class',
        is.array(props[key]) ? [...new Set(props[key])].join(' ') : props[key]
      );
      continue;
    }
    if (key === 'style') {
      var sty = '';
      if (is.string(props[key])) {
        sty = props[key];
      } else if (is.object(props[key])) {
        for (var k in props[key]) {
          sty += k + ':' + props[key][k] + ';';
        }
      }
      el.setAttribute(key, sty);
      continue;
    }
    if (key === 'data' || key.includes('data-')) {
      if (is.primitive(props[key])) {
        el.setAttribute(key, props[key]);
      } else if (is.object(props[key])) {
        for (var k in props[key]) {
          el.setAttribute(key + '-' + k, props[key][k]);
        }
      }
      continue;
    }
    if (key === 'on') {
      if (!is.object(props[key])) {
        continue;
      }
      for (var k in props[key]) {
        el.addEventListener(k, props[key][k]);
      }
      continue;
    }
    el.setAttribute(key, props[key]);
  }
}
```
