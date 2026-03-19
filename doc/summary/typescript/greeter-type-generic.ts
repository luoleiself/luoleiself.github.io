// 泛型
console.group("泛型");
// 泛型函数
console.groupCollapsed("泛型函数");
function identity<T>(arg: T): T {
  return arg;
}
// console.log(identity<number>(5)); // 写法 1
console.log(identity(5)); // 写法 2
function loggingIdentity<T>(arg: T[]): T[] {
  console.log(arg.length);
  return arg;
}
let myIdentity: { <U>(arg: U): U } = identity;
console.log(`
function identity<T>(arg: T): T {
  return arg;
}
// console.log(identity<number>(5)); // 写法 1
console.log(identity(5)); // 写法 2
function loggingIdentity<T>(arg: T[]): T[] {
  console.log(arg.length);
  return arg;
}
let myIdentity: { <U>(arg: U): U } = identity;`);
console.groupEnd();
// 泛型接口
console.groupCollapsed("泛型接口");
// interface GenericIdentityFn {
//   <T>(arg: T): T;
// }
// let myIdentity: GenericIdentityFn = identity;
// 把泛型参数作为整个接口的一个参数,这样就能清楚的知道使用的具体是哪个泛型类型, 接口里的其它成员也能知道这个参数的类型
// 泛型参数不再描述函数, 而是把非泛型函数签名作为泛型类型一部分, 当使用 GenericIdentityFn 的时候, 还需要传入一个类型参数指定泛型类型, 锁定了之后代码里面使用的类型
interface GenericIdentityFn<T> {
  (arg: T): T;
}
let myIdentity2: GenericIdentityFn<number> = identity;
console.log(`
// 把泛型参数作为整个接口的一个参数,这样就能清楚的知道使用的具体是哪个泛型类型, 接口里的其它成员也能知道这个参数的类型
// 泛型参数不再描述函数, 而是把非泛型函数签名作为泛型类型一部分, 当使用 GenericIdentityFn 的时候, 还需要传入一个类型参数指定泛型类型, 锁定了之后代码里面使用的类型
interface GenericIdentityFn<T> {
  (arg: T): T;
}
let myIdentity2: GenericIdentityFn<number> = identity;`);
console.groupEnd();
// 泛型类
console.groupCollapsed(
  "泛型类, 类包含静态部分和实例部分, 泛型类指的是实例部分的类型, 类的静态部分不能使用这个泛型类型"
);
class GenericCls<T> {
  zeroValue: T;
  add: (x: T, y: T) => T;
}
let g1: GenericCls<number> = new GenericCls<number>();
g1.zeroValue = 10;
g1.add = function (x, y) {
  return x + y + this.zeroValue;
};
console.log(g1.add(2, 3));
let g2: GenericCls<string> = new GenericCls<string>();
g2.zeroValue = "hello";
g2.add = function (x, y) {
  return this.zeroValue + " " + x + " " + y;
};
console.log(g2.add("world", "St."));
console.log(`
class GenericCls<T> {
  zeroValue: T;
  add: (x: T, y: T) => T;
}
let g1: GenericCls<number> = new GenericCls<number>();
g1.zeroValue = 10;
g1.add = function (x, y) {
  return x + y + this.zeroValue;
};
console.log(g1.add(2, 3));
let g2: GenericCls<string> = new GenericCls<string>();
g2.zeroValue = 'hello';
g2.add = function (x, y) {
  return this.zeroValue + ' ' + x + ' ' + y;
};
console.log(g2.add('world', 'St.'));`);
console.groupEnd();
// 泛型约束
console.groupCollapsed("泛型约束, 定义接口描述约束条件");
// 接口 Lengthwise 包含一个 length 属性,
interface Lengthwise {
  length: number;
}
function loggingIdentity2<T extends Lengthwise>(arg: T): T {
  console.log(arg.length);
  return arg;
}
// error TS2345: Argument of type 'number' is not assignable to parameter of type 'Lengthwise'.
console.log(loggingIdentity2(2));
console.log(loggingIdentity2({ length: 10, value: 3 }));
console.log(`
interface Lengthwise {
  length: number;
}
function loggingIdentity2<T extends Lengthwise>(arg: T): T {
  console.log(arg.length);
  return arg;
}
// error TS2345: Argument of type 'number' is not assignable to parameter of type 'Lengthwise'.
console.log(loggingIdentity2(2));
console.log(loggingIdentity2({ length: 10, value: 3 }));`);
console.groupEnd();
console.groupEnd();

// 高级类型
console.group("高级类型");
// 交叉类型
console.group(
  "交叉类型(Intersection Types), 是将多个类型合并为一个类型, 它包含了所需的所有类型的特性"
);
// Person & Serializable & Loggable 同时是 Person 和 Serializable 和 Loggable. 意味着这个类型的对象同时拥有了这三种类型的成员
// declare function loadEnv(
//   mode: string,
//   envDir: string,
//   prefixes?: string | string[]
// ): Record<string, string>;
// function loadEnv(mode, envDir, prefixes = 'VITE_') {
//   const env = {};
//   /* ... */
//   return env;
// }
console.groupEnd();
// 联合类型
console.groupCollapsed(
  "联合类型(Union Types), 表示一个值可以是几种类型之一, 用 (|) 分隔每个类型, 如果一个值是联合类型, 只能访问此联合类型的所有类型里共有的成员"
);
// 联合类型适合于那些值可以为不同类型的情况
// string | number | boolean 表示一个值可以是 string, number 或 boolean
interface Bird {
  fly(): void;
  layEggs(): void;
}
interface Fish {
  swim(): void;
  layEggs(): void;
}
function getSmallPet(): Fish | Bird {
  return {
    fly() {
      console.log("bird fly...");
    },
    layEggs() {
      console.log("bird layEggs...");
    },
  };
}
let pet = getSmallPet();
pet.layEggs(); // Ok
// pet.swim(); // Error, 如果是 Bird 类型, 调用 swim 方法报错
// 使用类型断言
if ((<Fish>pet).swim) {
  (<Fish>pet).swim();
} else {
  (<Bird>pet).fly();
}
console.log(`
interface Bird {
  fly(): void;
  layEggs(): void;
}
interface Fish {
  swim(): void;
  layEggs(): void;
}
function getSmallPet(): Fish | Bird {
  return {
    fly() {
      console.log("bird fly...");
    },
    layEggs() {
      console.log("bird layEggs...");
    },
  };
}
let pet = getSmallPet();
pet.layEggs(); // Ok
// pet.swim(); // Error, 如果是 Bird 类型, 调用 swim 方法报错
// 使用类型断言
if ((<Fish>pet).swim) {
  (<Fish>pet).swim();
} else {
  (<Bird>pet).fly();
}`);
console.groupEnd();
// 类型保护
console.groupCollapsed(
  "类型保护就是一些表达式, 它们会在运行时检查以确保在某个作用域里的类型, 定义一个函数作为类型保护表达式, 该函数的返回值是一个类型谓词: parameterName is Type"
);
function isFish(pet: Fish | Bird): pet is Fish {
  return (<Fish>pet).swim != undefined;
}
if (isFish(pet)) {
  pet.swim();
} else {
  pet.fly();
}
console.log(`
function isFish(pet: Fish | Bird): pet is Fish {
  return (<Fish>pet).swim != undefined;
}
if (isFish(pet)) {
  pet.swim();
} else {
  pet.fly();
}`);
console.groupEnd();
// typeof 类型保护
console.groupCollapsed("typeof 类型保护");
function padLeft(value: string, padding: number | string) {
  if (typeof padding === "number") {
    return Array(padding + 1).join(" ") + value;
  }
  if (typeof padding === "string") {
    return padding + value;
  }
  throw new Error(`Expected string or number, got '${padding}'.`);
}
console.log(`
function padLeft(value: string, padding: number | string) {
  if (typeof padding === "number") {
    return Array(padding + 1).join(" ") + value;
  }
  if (typeof padding === "string") {
    return padding + value;
  }
  throw new Error(\`Expected string or number, got '\$\{padding\}'.\`);
}`);
console.groupEnd();
console.groupCollapsed("instanceof 类型保护, 通过构造函数来细化类型的一种方式");
interface Padder {
  getPaddingString(): string
}
class SpaceRepeatingPadder implements Padder {
  constructor(private numSpaces: number) { }
  getPaddingString() {
    return Array(this.numSpaces + 1).join(" ");
  }
}
class StringPadder implements Padder {
  constructor(private value: string) { }
  getPaddingString() {
    return this.value;
  }
}
function getRandomPadder() {
  return Math.random() < 0.5 ?
    new SpaceRepeatingPadder(4) :
    new StringPadder("  ");
}
// 类型为SpaceRepeatingPadder | StringPadder
let padder: Padder = getRandomPadder();

if (padder instanceof SpaceRepeatingPadder) {
  padder; // 类型细化为'SpaceRepeatingPadder'
}
if (padder instanceof StringPadder) {
  padder; // 类型细化为'StringPadder'
}
console.groupEnd();

// 类型断言
console.groupCollapsed(
  "类型断言, 当在 TypeScript 中使用 JSX 时, 只允许使用 as 语法"
);
// 尖括号语法
// let someValue: any = 'this is a string';
// let strLength: number = (<string>someValue).length;
// console.log(strLength);
// as 语法
let someValue: any = "this is a string";
let strLength: number = (someValue as string).length;
console.log(strLength);
console.groupEnd();

// import interface-class

// 类型推论
console.groupCollapsed("类型推论：最佳通用类型, 上下文类型");
console.groupEnd();

// 类型兼容性
console.group("类型兼容性");
console.groupCollapsed(
  "ts 的类型兼容性是基于结构子类型, 结构类型是一种只使用其成员来描述类型的方式, 与名义类型形成对比, (基于名义类型的类型系统中, 数据类型的兼容性或等价性是通过明确的声明和/或类型的名称来决定的)"
);
console.groupEnd();
console.groupEnd();

// 类型别名
console.groupCollapsed(
  "类型别名, 给类型起一个新名字, 可以作用于原始值, 联合类型, 元组以及其他自定义类型," +
  "\n类型别名不能出现在声明右侧的任何位置," +
  "\n类型别名不能被 extends 和 implements," +
  "\n无法重新打开类型以添加新属性, 接口始终可扩展."
);
type Easing = "ease-in" | "ease-out" | "ease-in-out";
type Name = string;
type NameResolver = () => string;
type NameOrResolver = Name | NameResolver;
function getName(n: NameOrResolver): Name {
  if (typeof n === "string") {
    return n;
  } else {
    return n();
  }
}
type Container<T> = { value: T }; // 泛型类型别名
type Tree<T> = {
  value: T;
  left: Tree<T>;
  right: Tree<T>;
};
console.log(`
type Easing = "ease-in" | "ease-out" | "ease-in-out";
type Name = string;
type NameResolver = () => string;
type NameOrResolver = Name | NameResolver;
function getName(n: NameOrResolver): Name {
  if (typeof n === "string") {
    return n;
  } else {
    return n();
  }
}
type Container<T> = { value: T }; // 泛型类型别名
type Tree<T> = {
  value: T;
  left: Tree<T>;
  right: Tree<T>;
};`);
console.groupEnd();
console.groupEnd();

// 类型操作
console.group("类型操作");
// typeof
console.groupCollapsed(
  `typeof 类型运算符, 在 JavaScript 中, 用于在运行时获取值的类型
    在 TypeScript 中, 可以作为类型查询操作符在类型上下文中使用, 用来获取变量表达式的类型`);
let s1 = "hello";
let n: typeof s1 = "world";
console.log("typeof n", typeof n);

const user = {
  name: "Alice",
  age: 30
};
type UserType = typeof user;
// 等价于: { name: string; age: number; }
console.groupEnd();

// keyof
console.groupCollapsed(
  `keyof 类型运算符, 采用对象类型并生成其键的字符串或数字字面联合(已知 T 的公共属性名的联合)
    如果 type 包含了 string 或 number 的下标的签名, keyof 将都返回它们, 因为 js 对象总是会自动转换它们`
);
type Pointer = { x: number; y: number };
type Po = keyof Pointer;
const poDemo: Po = 'x'
console.log(`
type Pointer = { x: number; y: number };
type Po = keyof Pointer;`);
// 字符串类型的 key 将自动返回 string 和 number 的联合类型
type Arrayish = { [k: number]: boolean }
type A = keyof Arrayish; // type A = number;
type Mapish = { [k: string]: unknown };
type M = keyof Mapish; // type M = string | number;
console.groupEnd();

// 索引访问类型 indexed access types
console.groupCollapsed("索引访问类型");
type PErson = { age: number; name: string; alive: boolean };
type Age = PErson["age"];
type PE1 = PErson["age" | "alive"];
let page: Age = 110;
console.log(`
type PErson = { age: number; name: string; alive: boolean };
type Age = PErson["age"];
type PE1 = PErson["age" | "alive"];
let page: Age = 110;`);
console.log(page);

interface CustomeMap<T> {
  [key: string]: T;
}
let keys: keyof CustomeMap<number>; // string
let value: CustomeMap<number>['foo']; // number
// 使用 number 获取数组中一个元素的类型
const MyArray = [
  { name: 'Alice', age: 15 },
  { name: 'Bob', age: 23 },
  { name: 'Eve', age: 38 }
]
type MA = typeof MyArray[number];
type Age1 = typeof MyArray[number]['age'];
console.groupEnd();

// 条件类型 conditional types
console.groupCollapsed("条件类型");
interface Animal {
  live(): void;
}
interface Dog extends Animal {
  woof(): void;
}
type Example1 = Dog extends Animal ? number : string;
type Example2 = RegExp extends Animal ? number : string;
console.log(`
interface Animal {
  live(): void;
}
interface Dog extends Animal {
  woof(): void;
}
type Example1 = Dog extends Animal ? number : string;
type Example2 = RegExp extends Animal ? number : string;`);
console.log('------------')
// 条件复杂类型
type isArray<T> = T extends any[] ? true : false;
type isString<T> = T extends string ? true : false;

type CheckArray = isArray<number[]>;
type CheckStringArray = isArray<string>;
type CheckString = isString<"hello">;
type CheckNumber = isString<123>;

// infer 关键字用于在条件类型的 extends 子句中声明一个待推断的类型变量，它只能在条件类型的右侧使用
// 使用 infer 关键字从 true 分支中与之进行比较的类型中进行推断的方法
type ArrayElement<T> = T extends (infer U)[] ? U : T; // 如果 T 是数组类型, 则推断出元素类型 U, 否则返回 T 本身
type UnpackPromise<T> = T extends Promise<infer U> ? U : never;
type NumberArray = ArrayElement<number[]>; // number
type PromiseString = UnpackPromise<Promise<string>>; // string

// 内置 ReturnType 的简化实现
type ReturnType2<T extends (...args: any) => any> = T extends (...args: any) => infer R ? R : any;

type ReturnExample = ReturnType2<(a: number, b: string) => boolean>;
console.groupEnd();

// 映射类型 mapped types
console.groupCollapsed(
  "映射类型: " +
  "\n 映射修饰符: 使用 readonly 或者 ? 分别影响可变性和可选型" +
  "\n 添加前缀 + 或者 - 来添加或移除这些修饰符, 不使用默认为 +"
);
type OnlyBoolsAndHorses = {
  [key: string]: boolean | number;
}
type MaybeUser = {
  id: string;
  name: string;
  age: number;
};
type User1 = {
  [P in keyof MaybeUser]?: MaybeUser[P];
};
type Concrete<T> = {
  [Property in keyof T]-?: T[Property];
};
type User2 = Concrete<User1>;
type User3 = {
  +readonly [P in keyof MaybeUser]+?: MaybeUser[P];
};
type Decrete<T> = {
  -readonly [Property in keyof T]-?: T[Property];
}
type User4 = Decrete<User3>;
console.log(`
type MaybeUser = {
  id: string;
  name: string;
  age: number;
};
type User1 = {
  [P in keyof MaybeUser]?: MaybeUser[P];
};
type Concrete<T> = {
  [Property in keyof T]-?: T[Property];
};
type User2 = Concrete<User1>;
type User3 = {
  +readonly [P in keyof MaybeUser]+?: MaybeUser[P];
};
type Decrete<T> = {
 -readonly [Property in keyof T]-?: T[Property];
}
type User4 = Decrete<User3>;`);
console.log("--------------")

console.log('映射类型与键值映射')
type Getter<T> = {
  [K in keyof T as `get${Capitalize<string & K>}`]: () => T[K]
}
type MapPerson = {
  name: string
  age: number
}
type PersonGetter = Getter<MapPerson> // { getName: () => string; getAge: () => number }

type FilteredKeys<T, U> = {
  [K in keyof T as T[K] extends U ? K : never]: T[K]
}
interface MapMixed {
  name: string
  counl: number
  isActive: boolean
  data: Object
}
type StringKeys = FilteredKeys<MapMixed, string> // { name: string }
console.log(`
type Getter<T> = {
  [K in keyof T as \`get\${Capitalize<string & K>}\`]: () => T[K]
}
type MapPerson = {
  name: string
  age: number
}
type PersonGetter = Getter<MapPerson> // { getName: () => string; getAge: () => number }

type FilteredKeys<T, U> = {
  [K in keyof T as T[K] extends U ? K : never]: T[K]
}
interface MapMixed {
  name: string
  counl: number
  isActive: boolean
  data: Object
}
type StringKeys = FilteredKeys<MapMixed, string> // { name: string }
`)
console.groupEnd();

// 模板字面量类型 template literal types
console.groupCollapsed("模板字面类型");
type World = "world";
type Greeting = `hello ${World}`;
type EmailLocaleIDs = "welcome_email" | "email_heading";
type FooterLocaleIDs = "footer_title" | "footer_sendoff";
type AllLocaleIDs = `${EmailLocaleIDs | FooterLocaleIDs}_id`;

var w1: World = "hello"
var greeting1: Greeting = "hello world"
var e1: EmailLocaleIDs = "email_heading"

// 动态模板字面量类型
type PropEventType<T extends string> = `on${Capitalize<T>}`
type ButtonEvents = PropEventType<'click' | 'hover' | 'focus'> // onClick | onHover | onFocus

console.log(`
type World = "world";
type Greeting = \`hello \$\{World\}\`;
type EmailLocaleIDs = "welcome_email" | "email_heading";
type FooterLocaleIDs = "footer_title" | "footer_sendoff";
type AllLocaleIDs = \`\${EmailLocaleIDs | FooterLocaleIDs\}_id\`;
// 动态模板字面量类型
type PropEventType<T extends string> = \`on\${Capitalize<T>}\`
type ButtonEvents = PropEventType<'click' | 'hover' | 'focus'> // onClick | onHover | onFocus
`);
console.groupEnd();
console.groupEnd();
