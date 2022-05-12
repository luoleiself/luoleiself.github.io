---
title: JS小结
date: 2022-03-05 10:35:34
categories:
  - ES
tags:
  - ES6
  - js
---

### 改变原数组的方法

- pop 从数组中删除最后一个元素,并返回该元素的值(数组为空时返回 undefined). 此方法更改数组的长度

  ```javascript
  arr.pop();
  ```

- push 将一个或多个元素添加到数组的末尾, 并返回该数组的新长度

  ```javascript
  arr.push(element1, ..., elementN);
  ```

- shift 从数组中删除第一个元素,并返回该元素的值(数组为空则返回 undefined). 此方法更改数组的长度

  ```javascript
  arr.shift();
  ```

- unshift 将一个或多个元素添加到数组的开头. 并返回该数组的新长度(该方法修改原有数组)

  ```javascript
  arr.unshift(element1, ..., elementN);
  ```

  <!-- more -->

- reverse 将数组中元素的位置颠倒, 并返回该数组. 该方法会改变原数组

  ```javascript
  arr.reverse();
  ```

- sort 用[原地算法](https://en.wikipedia.org/wiki/In-place_algorithm)对数组的元素进行排序, 并返回数组

  - compareFunction 用来指定按某种顺序进行排列的函数, 省略则按照转换为的字符串的 unicode 位点进行排序
    - firstEl 第一个用于比较的元素
    - secondEl 第二个用于比较的元素

  ```javascript
  arr.sort([compareFunction]);
  ```

- splice 通过删除或替换现有元素或者原地添加新的元素来修改数组, 并以数组形式返回被修改的内容. 此方法会改变原数组

  - start 指定修改的开始位置
  - deleteCount 整数, 表示要移除的数组元素的个数, 如果为 0 或者负数, 则不移除元素
  - item1, item2 要添加进数组的元素,从 start 位置开始, 不指定则删除数组元素

  ```javascript
  array.splice(start[, deleteCount[, item1[, item2[, ...]]]]);
  ```

- fill 用一个固定值填充一个数组中从起始索引到终止索引内的全部元素. 不包括终止索引. 返回修改后的数组

  - value 用来填充数组元素的值
  - start 起始索引, 默认值为 0
  - end 终止索引, 默认值为 this.length

  ```javascript
  arr.fill(value[, start[, end]]);

  const array1 = [1, 2, 3, 4];
  // fill with 0 from position 2 until position 4
  console.log(array1.fill(0, 2, 4));
  // expected output: [1, 2, 0, 0]

  // fill with 5 from position 1
  console.log(array1.fill(5, 1));
  // expected output: [1, 5, 5, 5]

  console.log(array1.fill(6));
  // expected output: [6, 6, 6, 6]
  ```

- copyWithin 浅复制数组的一部分到同一数组中的另一个位置, 并返回它, 不会改变原数组的长度. 返回改变后的数组

  - target 整数, 复制序列到该位置, 如果是负数, target 将从末尾开始计算. 如果 target 大于等于 arr.length, 将会不发生拷贝
  - start 整数, 开始复制元素的起始位置, 如果是负数, start 将从末尾开始计算. 如果 start 被忽略, copyWithin 将会从 0 开始复制
  - end 整数, 开始复制元素的结束位置, copyWithin 将会拷贝到该位置, 但不包括 end 这个位置的元素. 如果是负数, end 将从末尾开始计算. 如果忽略则复制到数组结尾

  ```javascript
  arr.copyWithin(target[, start[, end]]);

  const array1 = ['a', 'b', 'c', 'd', 'e'];
  // copy to index 0 the element at index 3
  console.log(array1.copyWithin(0, 3, 4));
  // expected output: Array ["d", "b", "c", "d", "e"]

  // copy to index 1 all elements from index 3 to the end
  console.log(array1.copyWithin(1, 3));
  // expected output: Array ["d", "d", "e", "d", "e"]
  ```

### 表单 accept 属性

表单 input type="file" 上传图片时，accept 属性以文件名结尾格式在部分手机上使用时会提示 '没有应用可执行此操作', 将文件名结尾的格式改为 MIME 类型的格式

```html
<input type="file" accept=".png,.jpg,.jpeg" />

<input type="file" accept="image/png,image/jpeg" />
```

### 检测浏览器版本号

```javascript
class VersionDiff {
  /**
   * @property {private} #_comparator 关系符加版本号
   */
  static #_comparator = 'gt:0.0.0';
  /**
   * @method #greaterThan
   * @param {Array} arr 3位数版本号
   * @param {Array} arr1 3位数版本号,默认为关系符加版本号解析得到
   * @returns Boolean
   */
  static #greaterThan(arr, arr1) {
    return arr[0] > arr1[0] || (arr[0] == arr1[0] && arr[1] > arr1[1]) || (arr[0] == arr1[0] && arr[1] == arr1[1] && arr[2] > arr1[2])
      ? true
      : false;
  }
  /**
   * @method #lessThan
   * @param {Array} arr 3位数版本号
   * @param {Array} arr1 3位数版本号,默认为关系符加版本号解析得到
   * @returns Boolean
   */
  static #lessThan(arr, arr1) {
    return arr[0] < arr1[0] || (arr[0] == arr1[0] && arr[1] < arr1[1]) || (arr[0] == arr1[0] && arr[1] == arr1[1] && arr[2] < arr1[2])
      ? true
      : false;
  }
  /**
   * @method #equalTo
   * @param {Array} arr 3位数版本号
   * @param {Array} arr1 3位数版本号,默认为关系符加版本号解析得到
   * @returns Boolean
   */
  static #equalTo(arr, arr1) {
    return arr[0] == arr1[0] && arr[1] == arr1[1] && arr[2] == arr1[2] ? true : false;
  }
  /**
   * @method #notEqual
   * @param {Array} arr 3位数版本号
   * @param {Array} arr1 3位数版本号,默认为关系符加版本号解析得到
   * @returns Boolean
   */
  static #notEqual(arr, arr1) {
    return arr[0] != arr1[0] || (arr[0] == arr1[0] && arr[1] != arr1[1]) || (arr[0] == arr1[0] && arr[1] == arr1[1] && arr[2] != arr1[2])
      ? true
      : false;
  }
  /**
   * @method #versionDiff
   * @param {String} version 版本号
   * @param {RegExp} reg 匹配正则
   * @param {String} comparator 关系符加版本号
   * @returns Object
   */
  static #versionDiff(version, reg, comparator) {
    if (comparator == null || comparator == undefined) {
      comparator = VersionDiff.#_comparator;
    }
    if (typeof comparator != 'string') {
      return new Error('TypeError: params comparator must be string...');
    }
    let regExp = reg ? reg : /(wxwork)\/([^\s]*)/i;
    let _version = version ? version.match(regExp) : navigator.userAgent.match(regExp);

    if (!Array.isArray(_version)) {
      return { version: 0, [comparator]: 'unknow' };
    }

    let [operator, baseVersion] = comparator.split(':');
    let result = { version: _version[2], [_version[1].toString()]: true };
    let _ver = _version[2].split('.');
    let _baseVer = baseVersion.split('.');

    if (_ver.length != _baseVer.length) {
      if (_ver.length < _baseVer.length) {
        for (let i = 0, len = _baseVer.length - _ver.length; i < len; i++) {
          _ver.push(0);
        }
      } else {
        for (let i = 0, len = _ver.length - _baseVer.length; i < len; i++) {
          _baseVer.push(0);
        }
      }
    }

    switch (operator.toLowerCase()) {
      case 'gt':
        return Object.assign({}, result, { [comparator]: VersionDiff.#greaterThan(_ver, _baseVer) });
      case 'ge':
        return Object.assign({}, result, {
          [comparator]: VersionDiff.#equalTo(_ver, _baseVer) ? true : VersionDiff.#greaterThan(_ver, _baseVer),
        });
      case 'lt':
        return Object.assign({}, result, { [comparator]: VersionDiff.#lessThan(_ver, _baseVer) });
      case 'le':
        return Object.assign({}, result, {
          [comparator]: VersionDiff.#equalTo(_ver, _baseVer) ? true : VersionDiff.#lessThan(_ver, _baseVer),
        });
      case 'eq':
        return Object.assign({}, result, { [comparator]: VersionDiff.#equalTo(_ver, _baseVer) });
      case 'ne':
        return Object.assign({}, result, { [comparator]: VersionDiff.#notEqual(_ver, _baseVer) });
      default:
        return Object.assign({}, result, { [comparator]: 'unknow' });
    }
  }
  /**
   * @method test 静态测试方法
   * @param {String} browser 浏览器型号
   * @param {String} comparator 关系符加版本号
   * @returns Object
   */
  static test(browser = 'safari', comparator) {
    if (arguments.length == 0) {
      browser = `safari`;
      comparator = VersionDiff.#_comparator;
    } else if (arguments.length == 1) {
      comparator = browser;
      browser = `safari`;
    }
    return VersionDiff.#versionDiff(
      `${browser}/${Math.floor(Math.random() * 10)}.${Math.floor(Math.random() * 10)}.${Math.floor(Math.random() * 10)}`,
      new RegExp('(' + browser + ')/([^s]*)', 'i'),
      comparator
    );
  }
  /**
   * @method test 静态比较方法
   * @param {String} comparator 关系符加版本号
   * @returns Object
   */
  static diff(browser = 'safari', comparator) {
    if (arguments.length == 0) {
      browser = `safari`;
      comparator = VersionDiff.#_comparator;
    } else if (arguments.length == 1) {
      comparator = browser;
      browser = `safari`;
    }
    return VersionDiff.#versionDiff(null, new RegExp('(' + browser + ')/([^s]*)', 'i'), comparator);
  }
}

// 测试方式：使用 类名.test 调用
console.log(`VersionDiff.test('chrome', 'eq:5.3.2') `, VersionDiff.test('chrome', 'eq:5.3.2'));
console.log(`VersionDiff.test('chrome', 'ne:5.3.2') `, VersionDiff.test('chrome', 'ne:5.3.2'));
console.log(`VersionDiff.test('chrome', 'gt:5.3.2') `, VersionDiff.test('chrome', 'gt:5.3.2'));
console.log(`VersionDiff.test('chrome', 'le:5.3.2') `, VersionDiff.test('chrome', 'le:5.3.2'));
// 使用方式: 使用 类名.diff 调用方式
console.log(`VersionDiff.diff('gt:3.2.0') `, VersionDiff.diff('gt:3.2.0'));
```
