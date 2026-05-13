
// import decorator

// 三斜线指令
console.group(
  "三斜线指令, 只能放在引入文件的顶部, 如果出现在其它语句之后则作为普通的单行注释." +
  '\n/// <reference path="" />' +
  '\n/// <reference type="node"/> 引入声明文件, 表明这个文件使用了 @types/node/index.d.ts 里面声明的名字, 并且这个包需要在编译阶段与声明文件一起被包含进来, 仅当需要写一个 d.ts 文件时才使用这个指令' +
  '\n/// <reference no-default-lib="true" /> 把一个文件标记成默认库, 并告诉编译器在编译过程中不要包含这个默认库' +
  '\n/// <amd-module name="namedModule"/> 指定生成的 AMD 模块的名字, 默认都是匿名的'
);
console.groupEnd();

// 声明文件
console.group("声明文件, 帮助识别库以哪种方式被使用");
console.groupCollapsed("示例");
// 全局库
// 模块化库
// declare namespace Cats {
//   export interface KittingSettings {}
// }
console.log(`
// 全局变量---声明
declare var foo: number;
declare const FOO: string;
declare let bar: boolean;
// 全局函数---声明
declare function greet(message: string): void;
// 带属性的对象
let result = myLib.makeGreeting('Hello, world');
console.log('The computed greeting is:' + result);
let count = myLib.numberOfGreetings;
// 带属性的对象---声明
declare namespace myLib {
  function makeGreeting(message: string): string;
  let numberOfGreetings: number;
}`);
console.log(`
// 函数重载
let x: Widget = getWidget(43);
let arr: Widget[] = getWidget('all of them');
// 函数重载---声明
declare function getWidget(n: number): Widget;
declare function getWidget(s: string): Array<Widget>;
// 可重用类型(接口)
greet({ greeting: 'hello world', duration: 4000 });
// 可重用类型(接口)---声明
interface GreetingSettings {
  greeting: string;
  duration?: number;
  color?: string;
}
declare function greet(setting: GreetingSettings): void;`);
console.log(`
// 可重用类型(类型别名)
function getGreeting() {
  return 'Howdy';
}
class MyGreeter extends Greeter {}
greet('hello');
greet(getGreeting);
greet(new MyGreeter('world'));
// 可重用类型(类型别名)---声明
type GreetingLike = string | (() => string) | MyGreeter;
declare function greet(g: GreetingLike): void;`);
// 组织类型---使用命名空间组织类型
// 类
// const dg = new DIYGreeter('hello, world');
// dg.greeting = 'howDy';
// dg.showGreeting();
// class SpecialGreeter extends Greeter {
//   constructor() {
//     super('Very special greetings...');
//   }
// }
// 类---声明
// declare class DIYGreeter {
//   constructor(message: string);
//   greeting: string;
//   showGreeting(): void;
// }
console.groupEnd();
// 规范
console.groupCollapsed("规范");
console.log(
  "不要使用基本数据类型的包装类型" +
  "\n不用定义一个从来没使用过其类型参数的泛型类型"
);
console.log(
  "不要为返回值被忽略的回调函数设置一个 any 类型的返回值类型. 通常使用 void"
);
console.log("重载");
console.log(
  "不要因为回调函数参数个数不同而写不同的重载, 应该只使用最大参数个数写一个重载" +
  "\n不要把一般的重载放在精确的重载前面. TypeScript 匹配到第一个符合条件时不再往下查找" +
  "\n不要仅在末尾参数不同时写不同的重载, 优先使用可选参数" +
  "\n不要为仅在某个位置上的参数类型不同的情况下定义重载, 优先使用联合类型"
);

/* Wrong */
interface Moment {
  utcOffset(): number;
  utcOffset(b: number): Moment;
  utcOffset(b: string): Moment;
}
/*  OK */
interface Moment {
  utcOffset(): number;
  utcOffset(b: number | string): Moment;
}
console.groupEnd();
// 原理
console.groupCollapsed(
  "原理" +
  "\n类型别名声明, 接口声明, 类声明, 枚举声明, 指向某个类型的 import 声明, 这几种声明方式都会创建一个新的类型名称" +
  "\nlet, const, var 声明, 包含值 的 namespace 或 module 声明, enum 声明, class 声明, 指向值的 import 声明, function 声明, 这几种声明方式都会创建一个值"
);
// 类型别名声明, 接口声明, 类声明, 枚举声明, 指向某个类型的 import 声明, 这几种声明方式都会创建一个新的类型名称
// let, const, var 声明, 包含值 的 namespace 或 module 声明, enum 声明, class 声明, 指向值的 import 声明, function 声明, 这几种声明方式都会创建一个值

// export var SomeVar: { a: SomeType };
// export interface SomeType {
//   count: number;
// }
// import * as foo from './foo';
// let x: foo.SomeType = foo.SomeVal.a;
// console.log(x.count);
// 使用 Bar 作为类型和值,
// export var Bar: { a: Bar };
// export interface Bar {
//   count: number;
// }
//
// import { Bar } from './foo';
// let x: Bar = Bar.a;
// console.log(x.count);

// class D {} 和 interface D {} 同样可以作为 D 类型的属性, 不能使用接口类型往类型别名里添加成员
class Foo {
  x: number;
}
interface Foo {
  y: number;
}
// let a: Foo = ...;
// console.log(a.x + a.y);
class D { }
namespace D {
  export let x: number;
}
let y = D.x;
console.groupEnd();
console.groupEnd();
