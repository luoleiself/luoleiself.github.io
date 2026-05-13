// 类型注解是一种轻量级的为函数或变量添加约束的方式

// 接口： 只在两个类型内部的结构兼容那么这两个类型就是兼容的, 实现接口时只要包含了接口要求的结构即可而不必明确地使用 implements 语句
function greeter(person: Gentalman) {
  return "Hello, " + person.firstName + " " + person.lastName;
}

// let user = "Jane User";
// document.body.innerHTML = greeter(user);
// let user = [0, 1, 2];
// document.body.innerHTML = greeter(user);

// interface
interface Gentalman {
  firstName: string
  lastName: string
}
let user1 = { firstName: 'Jane', lastName: 'User' };
document.body.innerHTML = greeter(user1);


// interface 更专注于定义对象和类的结构, 支持继承、合并
// type 可以定义类型别名、联合类型、交叉类型, 但不支持继承和自动合并
interface PersonInfoInter {
  name: string
  age: number
}
interface PersonInfoInter {
  speak(): void
}
let user3: PersonInfoInter = {
  name: "hello user3",
  age: 18,
  speak() {
    console.log('user3 speak...');
  }
}
console.log(user3)

// Class
// 构造函数的参数使用 public 等同于创建了同名的成员变量
class Student {
  fullName: string;
  constructor(
    public firstName: string,
    public middleInitial: string,
    public lastName: string
  ) {
    this.fullName = firstName + " " + middleInitial + " " + lastName;
  }
}
let user2 = new Student("Jane", "M.", "User");
document.body.innerHTML = greeter(user2);

// 基础类型
let isDone: boolean = false;
console.log(isDone);

let decLiteral: number = 0;
let hexLiteral: number = 0xf00d;
let binaryLiteral: number = 0b1000;
let octalLiteral: number = 0o744;
console.log(decLiteral, hexLiteral, binaryLiteral, octalLiteral);

let sentence: string = "Jim";
console.log(sentence);

// 声明字面量类型
let a: 'hello world';
a = 'hello gg';
a = 100;
let b: 200;
b = 100;
b = true;

console.groupCollapsed("array");
// 数组 []
let list1: number[] = [1, 2, 3, 4, 5];
// 数组泛型
let list2: Array<number> = [6, 7, 8, 9, 10];
console.log(list1, list2);
console.groupEnd();

// 元组 Tuple: 允许表示一个已知元素数量和类型的数组, 各元素的类型不必相同
console.groupCollapsed(
  "元组 Tuple: 允许表示一个已知元素数量和类型的数组, 各元素的类型不必相同"
);
let t1: [string, number] = ["hello Tom", 18];
console.log(t1[0].substring(1));
console.log(`let t1: [string, number] = ["hello Tom", 18];`);
// error TS2322: Type '"world"' is not assignable to type 'undefined'.
// error TS2493: Tuple type '[string, number]' of length '2' has no element at index '3'.
// t1[3] = 'world';
console.groupEnd();

// 枚举 enum
console.group(
  "枚举 enum, 常量枚举(const enum)只能使用常量枚举表达式不允许包含计算成员, 在编译阶段会被删除, 在使用的地方用内联的方式替换."
);
// 数字枚举
console.groupCollapsed("数字枚举, 默认从 0 开始自动累加, 也可以手动指定");
enum Direction {
  Up,
  Down,
  Left,
  Right,
}
console.log(Direction.Up); // 0
enum Resp {
  No = 1,
  Yes,
  Ha = 10,
  He,
  Xi = 5,
  Xa,
}
console.log(
  "No %d, Yes %d, Ha %d, He %d, Xi %d, Xa %d",
  Resp.No,
  Resp.Yes,
  Resp.Ha,
  Resp.He,
  Resp.Xi,
  Resp.Xa
); // No 1, Yes 2, Ha 10, He 11, Xi 5, Xa 6
enum Access {
  None,
  Read = 1 << 1,
  Write = 1 << 2,
  ReadWrite = 1 << 3,
}
console.log(
  "None %d, Read %d, Write %d, ReadWrite %d",
  Access.None,
  Access.Read,
  Access.Write,
  Access.ReadWrite
); // greeter.js:92 None 0, Read 2, Write 4, ReadWrite 8
console.log(`
enum Direction {
  Up,
  Down,
  Left,
  Right,
}
enum Resp {
  No = 1,
  Yes,
  Ha = 10,
  He,
  Xi = 5,
  Xa,
}
enum Access {
  None,
  Read = 1 << 1,
  Write = 1 << 2,
  ReadWrite = 1 << 3,
}`);
console.groupEnd();
// 字符串枚举
console.groupCollapsed("字符串枚举, 没有自增长的行为, 可以很好的序列化");
enum Position {
  Top = "TOP",
  Left = "LEFT",
  Bottom = "BOTTOM",
  Right = "RIGHT",
}
console.log(
  "Top %s, Right %s, Bottom %s, Left %s",
  Position.Top,
  Position.Right,
  Position.Bottom,
  Position.Left
); // Top TOP, Right RIGHT, Bottom BOTTOM, Left LEFT
console.log(`
enum Position {
  Top = 'TOP',
  Left = 'LEFT',
  Bottom = 'BOTTOM',
  Right = 'RIGHT',
}`);
console.groupEnd();
// 常量枚举
console.log('常量枚举在编译阶段会被删除, 在使用的地方使用内联的方式替换');
const enum Directions {
  Up,
  Right,
  Down,
  Left
}
console.log(Directions.Up);
console.log(`
const enum Directions {
  Up,
  Right,
  Down,
  Left
}
console.log(Directions.Up);
console.log('0' /* Directions.Up */);
`)

// 异构枚举
console.groupCollapsed("异构枚举, 可以混合字符串和数字成员");
console.groupEnd();
console.groupEnd();

// 常量断言
console.groupCollapsed("常量断言");
console.log('Typescript 3.4 引入的常量断言(const assertion), 主要作用是告诉 Typescript 编译器将值视为不可变的字面量类型, 而不是更宽泛的类型.');
console.log(`
1. 创建字面量类型: 将值推断为具体的字面量类型
2. 防止类型扩宽: 阻止 Typescript 自动扩宽类型
3. 创建只读类型: 使对象和数组变为只读
4. 减少运行时错误: 在编译时捕获更多错误
5. 更好的智能提示: IDE 可以提供更准确的代码补全
`);
console.groupCollapsed('普通声明')
// 普通声明 - 类型会被扩宽
const color = "red";        // 类型: string
const number = 42;          // 类型: number
const flag = true;          // 类型: boolean

// 使用 as const - 字面量类型
const color2 = "red" as const;    // 类型: "red"
const number2 = 42 as const;      // 类型: 42
const flag2 = true as const;      // 类型: true
console.groupEnd();

console.groupCollapsed('数组声明')
// 普通数组声明
const fruits = ["apple", "banana", "orange"];
// 类型: string[]

// 使用 as const
const fruits2 = ["apple", "banana", "orange"] as const;
// 类型: readonly ["apple", "banana", "orange"]

// 尝试修改会报错
// fruits2.push("grape");  // Error: Property 'push' does not exist
// fruits2[0] = "pear";    // Error: Cannot assign to read-only property
console.groupEnd();

console.groupCollapsed('对象声明');
// 普通对象声明
const config = {
  apiUrl: "https://api.example.com",
  timeout: 5000,
  retries: 3
};
// 类型: { apiUrl: string; timeout: number; retries: number; }

// 使用 as const
const config2 = {
  apiUrl: "https://api.example.com",
  timeout: 5000,
  retries: 3
} as const;
// 类型: {
//   readonly apiUrl: "https://api.example.com";
//   readonly timeout: 5000;
//   readonly retries: 3;
// }
console.groupEnd();
// 替代枚举
// 传统枚举
enum Status {
  PENDING = "pending",
  SUCCESS = "success",
  ERROR = "error"
}
// 使用 as const 替代
const STATUS = {
  PENDING: "pending",
  SUCCESS: "success",
  ERROR: "error"
} as const;
type StatusType = typeof STATUS[keyof typeof STATUS];
// 类型: "pending" | "success" | "error"
console.groupEnd();

// any 任意类型, 类型检查器跳过检查
console.groupCollapsed("any 任意类型, 类型检查器跳过检查");
let notSure: any = 4;
console.log(notSure);
notSure = "maybe a string instead;";
console.log(notSure);
notSure = true;
console.log(notSure);

let l3: any[] = [1, "free", true];
console.log("l3", l3);
let l4: Array<any> = [2, "freeze", false];
console.log("l4", l4);
console.groupEnd();

// void 没有任何类型, 某种程度上与 any 类型相反, 只能赋值 undefined 和 null
console.groupCollapsed("void 没有任何类型, 只能赋值 undefined 和 null");
function warnUser(): void {
  console.log("This is my warning message.");
}
warnUser();
let unusable1: void = undefined;
// let unusable2: void = null; // error TS2322: Type 'null' is not assignable to type 'void'.
console.log(unusable1 /* unusable2 */);
console.groupEnd();

// Undefined 和 Null, 严格模式下, null 和 undefined 只能赋值给 void 和它们本身
console.groupCollapsed(
  "Undefined 和 Null, 严格模式下, null 和 undefined 只能赋值给 void 和它们本身"
);
let und: undefined = undefined;
let nul: null = null;
// let und1: number = undefined; // error TS2322: Type 'undefined' is not assignable to type 'number'.
// let nul1: string = null; // error TS2322: Type 'null' is not assignable to type 'string'.
console.log(und, nul /* und1 ,*/ /* nul1 */);
console.groupEnd();

// never 永不存在的值的类型
console.groupCollapsed("never 永不存在的值的类型");
console.log(
  "总是会抛出异常或者根本就不会有返回值的函数表达式或箭头函数表达式的返回值类型"
);
// let n1: number = never; // error TS2693: 'never' only refers to a type, but is being used as a value here.
function error(message: string): never {
  throw new Error(message);
}
function fail() {
  return error("Something failed.");
}
try {
  fail();
} catch (error) {
  console.error(error);
}
console.groupEnd();

// Object 非原始类型, 除 number, string, boolean, symbol, null, undefined 之外的类型
console.groupCollapsed(
  "Object 非原始类型, 除 number, string, boolean, symbol, null, undefined 之外的类型"
);

console.groupEnd();

// Symbol
console.groupCollapsed("Symbol, 不能使用 new 调用构造函数");
// 静态方法
console.log("https://www.tslang.cn/docs/handbook/symbols.html");
// Symbol.hasInstance; // 用于判断某对象是否为某构造器的实例
// Symbol.isConcatSpreadable; // 配置某对象作为 'Array.prototype.concat()' 方法的参数时是否展开其数组元素
// Symbol.iterator; // 为对象定义一个默认的迭代器, 可以被 'for ... of' 循环使用
// Symbol.asyncIterator; // 为对象定义一个默认的异步迭代器, 可以被 'for await ... of' 循环使用

class Array1 {
  static [Symbol.hasInstance](ins) {
    return Array.isArray(ins);
  }
}
console.log([] instanceof Array1); // true

var numeric = [1, 2, 3, 4];
numeric[Symbol.isConcatSpreadable] = false;
console.log([5, 6, 7, 8].concat(numeric)); // [5,6,7,8,[1,2,3,4]]

const iterable1 = {
  *[Symbol.iterator]() {
    yield 1;
    yield 2;
    yield 3;
  },
};
console.log([...iterable1]); // [1, 2, 3]
console.log(`
class Array1 {
  static [Symbol.hasInstance](ins) {
    return Array.isArray(ins);
  }
}
console.log([] instanceof Array1); // true

var numeric = [1, 2, 3, 4];
numeric[Symbol.isConcatSpreadable] = false;
console.log([5, 6, 7, 8].concat(numeric)); // [5,6,7,8,[1,2,3,4]]

const iterable1 = {
  *[Symbol.iterator]() {
    yield 1;
    yield 2;
    yield 3;
  },
};
console.log([...iterable1]); // [1, 2, 3]`);
console.groupEnd();
