
// 命名空间和模块
console.group("命名空间和模块");
// 命名空间
console.groupCollapsed(
    "命名空间(内部模块), 是位于全局命名空间下的一个普通的带有名字的 JavaScript 对象, 用来组织代码, 以便于在记录它们类型的同时还不用担心与其它对象产生命名冲突. " +
    "\n命名空间内的变量，函数，类，类型别名和接口需要使用 export 导出才能在命名空间外部使用. " +
    '\n命名空间可以分割成多个文件, 因为不同文件之间存在依赖关系, 需要使用引用标签 /// <reference path="" /> 告诉编译器文件之间的关联, 只能放在引入文件的顶部, 如果出现在其它语句之后则作为普通的单行注释'
);
// 通过命令行参数 --outFile 将多个命名空间文件结合在一起, 或者在 html 文件中使用 script 标签分别引入
// Validation.ts
namespace Validation {
    export interface StringValidator {
        isAcceptable(s: string): boolean;
    }
}

// LettersOnlyValidator.ts
/// <reference path="Validation.ts" />
namespace Validation {
    const lettersRegexp = /^[A-Za-z]+$/;
    export class LettersOnlyValidator implements StringValidator {
        isAcceptable(s: string): boolean {
            return lettersRegexp.test(s);
        }
    }
}

// ZipCodeValidator.ts
/// <reference path="Validation.ts" />
namespace Validation {
    const numberRegexp = /^\d+$/;
    export class ZipCodeValidator implements StringValidator {
        isAcceptable(s: string): boolean {
            return s.length == 5 && numberRegexp.test(s);
        }
    }
}

// Test.ts
/// <reference path="Validation.ts" />
/// <reference path="LettersOnlyValidator.ts" />
/// <reference path="ZipCodeValidator.ts" />
// Some samples to try
let strings = ["Hello", "98052", "101"];
// Validators to use
let validators: { [s: string]: Validation.StringValidator } = {};
validators["ZIP code"] = new Validation.ZipCodeValidator();
validators["Letters only"] = new Validation.LettersOnlyValidator();
// Show whether each string passed each validator
for (let s of strings) {
    for (let name in validators) {
        console.log(
            `"${s}" - ${validators[name].isAcceptable(s) ? "matches" : "does not match"
            } ${name}`
        );
    }
}
console.log(`
// Validation.ts
namespace Validation {
  export interface StringValidator {
    isAcceptable(s: string): boolean;
  }
}

// LettersOnlyValidator.ts
/// <reference path="Validation.ts" />
namespace Validation {
  const lettersRegexp = /^[A-Za-z]+$/;
  export class LettersOnlyValidator implements StringValidator {
    isAcceptable(s: string): boolean {
      return lettersRegexp.test(s);
    }
  }
}

// ZipCodeValidator.ts
/// <reference path="Validation.ts" />
namespace Validation {
  const numberRegexp = /^\\d+$/;
  export class ZipCodeValidator implements StringValidator {
    isAcceptable(s: string): boolean {
      return s.length == 5 && numberRegexp.test(s);
    }
  }
}

// Test.ts
/// <reference path="Validation.ts" />
/// <reference path="LettersOnlyValidator.ts" />
/// <reference path="ZipCodeValidator.ts" />
// Some samples to try
let strings = ["Hello", "98052", "101"];
// Validators to use
let validators: { [s: string]: Validation.StringValidator } = {};
validators["ZIP code"] = new Validation.ZipCodeValidator();
validators["Letters only"] = new Validation.LettersOnlyValidator();
// Show whether each string passed each validator
for (let s of strings) {
  for (let name in validators) {
    console.log(
      \`"\$\{s\}" - \$\{
        validators[name].isAcceptable(s) ? "matches" : "does not match"
      \} \$\{name}\`
    );
  }
}`);
console.groupEnd();
// 命名空间别名
console.groupCollapsed(
    "命名空间别名, 使用 import q = x.y.z 给常用的对象起一个短的名字."
);
namespace Shape {
    export namespace Polygons {
        export interface Draw {
            draw(): void;
        }
        export class Triangle { }
        export class Square {
            draw() {
                console.log("square draw...");
            }
        }
        export class Rectangle { }
    }
}
import polygons = Shape.Polygons; // 命名空间别名
let sq: polygons.Draw = new polygons.Square(); // Same as new Shape.Polygons.Square();
console.log(sq);
sq.draw();
console.log(`
namespace Shape {
  export namespace Polygons {
    export interface Draw {
      draw(): void;
    }
    export class Triangle {}
    export class Square {
      draw() {
        console.log("square draw...");
      }
    }
    export class Rectangle {}
  }
}
import polygons = Shape.Polygons; // 命名空间别名
let sq: polygons.Draw = new polygons.Square(); // Same as new Shape.Polygons.Square();`);
console.groupEnd();
console.groupCollapsed(
    "TypeScript 不同的模块永远也不会在相同的作用域内使用相同的名字, 因为使用模块的人会为它们命名, 所以完全没有必要把导出的符号包裹在一个命名空间里."
);
console.groupEnd();
console.groupEnd();

// 模块解析
console.group("模块解析, 指编译器在查找导入模块内容时所遵循的流程.");
console.groupCollapsed(
    "相对导入(/, ./, ../), 相对于导入它的文件, 并且不能解析为一个外部模块声明."
);
// import Entry from './components/Entry';
// import { DefaultHeaders } from '../constants/http';
// import "/mod";
console.groupEnd();
console.groupCollapsed(
    "非相对导入, 可以相对于 baseUrl 或通过路径映射来进行解析, 还可以被解析成 外部模块声明"
);
// import * as $ from "jQuery";
// import { Component } from '@angular/core';
console.groupEnd();
console.group("解析规则");
// classic 解析规则, 相对路径导入基于当前模块所在路径解析导入它的模块路径后的路径中查找.
//                  非相对导入从导入它的模块的所在目录逐级往上尝试查找.
// node 解析规则, 相对路径导入 1. 检查导入模块文件是否存在
//                           2. 检查导入模块是否是目录并且是否包含一个 package.json 文件, 并且 package.json 文件指定了一个 main 字段
//                           3. 检查导入模块是否是目录并且是否包含一个 index.js 文件
//               非相对导入 从导入它的模块的所在目录逐级往上尝试查找 node_modules 目录, 并按照相对路径导入的解析顺序尝试查找
console.log(`
// classic 解析规则, 相对路径导入基于当前模块所在路径解析导入它的模块路径后的路径中查找.
//                  非相对导入从导入它的模块的所在目录逐级往上尝试查找.
// node 解析规则, 相对路径导入 1. 检查导入模块文件是否存在
//                           2. 检查导入模块是否是目录并且是否包含一个 package.json 文件, 并且 package.json 文件指定了一个 main 字段
//                           3. 检查导入模块是否是目录并且是否包含一个 index.js 文件
//               非相对导入 从导入它的模块的所在目录逐级往上尝试查找 node_modules 目录, 并按照相对路径导入的解析顺序尝试查找`);
console.groupEnd();
console.groupEnd();

// 声明合并
console.group(
    "声明合并, 指编译器将针对同一个名字的两个独立声明合并为单一声明, 合并后的声明同时拥有原先两个声明的特性, 任何数量的声明都可被合并, 不仅局限于两个声明."
);
// TypeScript 中的声明会创建以下三种实体之一: 命名空间, 类型或值
// 合并接口
console.groupCollapsed(
    "合并接口, \n同名接口中同时声明了同名的非函数成员且它们类型不同, 则编译器会报错, \n同名函数成员都会被当成这个函数的一个重载, 后面的接口具有更高的优先级" +
    "\n当出现特殊的函数签名时, 如果签名里有一个参数的类型是单一的字符串字面量(不是字符串字面量的联合类型), 那么它将会被提升到重载列表的最顶端"
);
// 如果两个接口中同时声明了同名的非函数成员且它们的类型不同, 则编译器会报错.
interface Box {
    height: number;
    width: number;
}
interface Box {
    // error TS2717: Subsequent property declarations must have the same type.  Property 'height' must be of type 'number', but here has type 'string'.
    height: string;
    scale: number;
}
let box: Box = { height: 5, width: 6, scale: 10 };
console.log(box);
console.log(`
  // 如果两个接口中同时声明了同名的非函数成员且它们的类型不同, 则编译器会报错.
  interface Box {
    height: number;
    width: number;
  }
  interface Box {
    // error TS2717: Subsequent property declarations must have the same type.  Property 'height' must be of type 'number', but here has type 'string'.
    height: string;
    scale: number;
  }
  let box: Box = { height: 5, width: 6, scale: 10 };
  console.log(box);`);
// 对于同名函数成员, 每个同名的函数声明都会被当成这个函数的一个重载, 后面的接口具有更高的优先级
console.log(`
  // 对于同名函数成员, 每个同名的函数声明都会被当成这个函数的一个重载, 后面的接口具有更高的优先级
  class Dog extends Animal {
    constructor(name: string) {
      super(name);
    }
  }
  class Cat extends Animal {
    constructor(name: string) {
      super(name);
    }
  }
  interface Cloner {
    clone(a: Animal): Animal;
  }
  interface Cloner {
    clone(a: Animal): Rhino;
  }
  interface Cloner {
    clone(a: Animal): Dog;
    clone(a: Animal): Cat;
  }
  // 合并后
  // interface Cloner {
  //   clone(a: Animal): Dog;
  //   clone(a: Animal): Cat;
  //   clone(a: Animal): Rhino;
  //   clone(a: Animal): Animal;
  // }`);
// 当出现特殊的函数签名时, 如果签名里有一个参数的类型是单一的字符串字面量(不是字符串字面量的联合类型), 那么它将会被提升到重载列表的最顶端
console.log(`
  // 当出现特殊的函数签名时, 如果签名里有一个参数的类型是单一的字符串字面量(不是字符串字面量的联合类型), 那么它将会被提升到重载列表的最顶端
  interface Document {
    createElement(tagName: any): Element;
  }
  interface Document {
    createElement(tagName: 'div'): HTMLDivElement;
    createElement(tagName: 'span'): HTMLSpanElement;
  }
  interface Document {
    createElement(tagName: string): HTMLElement;
    createElement(tagName: 'canvas'): HTMLCanvasElement;
  }
  // 合并后
  // interface Document {
  //   createElement(tagName: 'canvas'): HTMLCanvasElement;
  //   createElement(tagName: 'div'): HTMLDivElement;
  //   createElement(tagName: 'span'): HTMLSpanElement;
  //   createElement(tagName: string): HTMLElement;
  //   createElement(tagName: any): Element;
  // }`);
console.groupEnd();
// 合并命名空间
console.groupCollapsed(
    "合并命名空间, \n同名命名空间合并其成员, 模块导出的同名接口进行合并, 构成单一命名空间内含合并后的接口" +
    "\n命名空间里值的合并, 如果当前已经存在给定名字的命名空间, 那么后来的命名空间的导出成员会被加到已经存在的那个模块里" +
    "\n非导出成员, 仅在其原有的(合并前)命名空间内可见, 从其它命名空间合并进来的成员无法访问非导出成员"
);
// 同名命名空间合并其成员, 模块导出的同名接口进行合并, 构成单一命名空间内含合并后的接口
console.log(`
  // 同名命名空间合并其成员, 模块导出的同名接口进行合并, 构成单一命名空间内含合并后的接口
  namespace Animal {
    export class Zebra {}
  }
  namespace Animal {
    export interface Legged {
      numberOfLegs: number;
    }
    export class Dog {}
  }
  // 合并后
  // namespace Animal {
  //   export interface Legged {
  //     numberOfLegs: number;
  //   }
  
  //   export class Zebra {}
  //   export class Dog {}
  // }`);
// 命名空间里值的合并, 如果当前已经存在给定名字的命名空间, 那么后来的命名空间的导出成员会被加到已经存在的那个模块里
// 非导出成员, 仅在其原有的(合并前)命名空间内可见, 从其它命名空间合并进来的成员无法访问非导出成员
console.log(`
  // 命名空间里值的合并, 如果当前已经存在给定名字的命名空间, 那么后来的命名空间的导出成员会被加到已经存在的那个模块里
  // 非导出成员, 仅在其原有的(合并前)命名空间内可见, 从其它命名空间合并进来的成员无法访问非导出成员
  namespace Animal {
    let haveMuscles = true;
  
    export function animalsHaveMuscles() {
      return haveMuscles;
    }
  }
  namespace Animal {
    export function doAnimalHaveMuscles() {
      // error TS2304: Cannot find name 'haveMuscles'.
      return haveMuscles; // 不能访问未导出的成员
    }
  }`);
console.groupEnd();
// 命名空间与类和函数和枚举类型合并
console.groupCollapsed(
    "命名空间与类和函数和枚举类型合并, 只要命名空间的定义符合将要合并类型的定义, 合并结果包含两者的声明类型."
);
console.log(`
  // 合并命名空间和类
  // class Album {
  //   label: Album.AlbumLabel;
  // }
  // namespace Album {
  //   export class AlbumLabel {}
  // }
  // 合并命名空间和函数
  // function buildLabel(name: string): string {
  //   return buildLabel.prefix + name + buildLabel.suffix;
  // }
  // namespace buildLabel {
  //   export let suffix = ';';
  //   export let prefix = 'Hello, ';
  // }
  // 合并命名空间和枚举
  // enum Color {
  //   red = 1,
  //   green = 2,
  //   blue = 4,
  // }
  // namespace Color {
  //   export function mixColor(colorName: string) {
  //     if (colorName == 'yellow') {
  //       return Color.red + Color.green;
  //     } else if (colorName == 'white') {
  //       return Color.red + Color.green + Color.blue;
  //     } else if (colorName == 'magenta') {
  //       return Color.red + Color.blue;
  //     } else if (colorName == 'cyan') {
  //       return Color.green + Color.blue;
  //     }
  //   }
  // }`);
console.groupEnd();
console.groupEnd();
