
// 工具类型
console.group(
  "工具类型, TypeScript 提供了几种工具类型来促进常见的类型转换, 这些工具在全局作用域内可用."
);
// Promise<ResponseUserType>;
console.groupCollapsed(
  "Awaited<Type> // 对 async 函数中的 await 或 Promise 中的 then() 方法等进行建模"
);
// Awaited<Type> // 对 async 函数中的 await 或 Promise 中的 then() 方法等进行建模
type A1 = Awaited<Promise<string>>;
type A2 = Awaited<Promise<Promise<number>>>;
type A3 = Awaited<boolean | Promise<number>>;
console.groupEnd();

console.groupCollapsed(
  "Partial<Type> // 构造一个将 Type 的所有属性设置为可选的类型"
);
// Partial<Type> // 构造一个将 Type 的所有属性设置为可选的类型
interface ToDo {
  title: string;
  description: string;
  visible: boolean;
}
type B1 = Partial<ToDo>;
// 深度 Partial 类型
type DeepPartial<T> = {
  [K in keyof T]?: T[K] extends object ? DeepPartial<T[K]> : T[K];
}
interface Config {
  server: {
    port: number
    host: string
    options: {
      timeout: number
      retries: number
    }
  }
  database: {
    url: string
    name: string
  }
}
type PartialConfig = DeepPartial<Config>
console.groupEnd();

console.groupCollapsed(
  "Required<Type>  // 构造一个将 Type 的所有属性设置为 required 的类型, 与 Partial 相反."
);
// Required<Type>  // 构造一个将 Type 的所有属性设置为 required 的类型, 与 Partial 相反.
type C1 = Required<B1>;
// 深度 Required 类型
type DeepRequired<T> = {
  [K in keyof T]-?: T[K] extends Object ? DeepRequired<T[K]> : T[K];
}
console.groupEnd();

console.groupCollapsed(
  "Readonly<Type> // 构造一个将 Type 的所有属性设置为 readonly 的类型, 构造类型的属性不能重新分配"
);
// Readonly<Type> // 构造一个将 Type 的所有属性设置为 readonly 的类型, 构造类型的属性不能重新分配
type D1 = Readonly<B1>;
type D2 = Readonly<C1>;
// 深度 Readonly 类型
type DeepReadonly<T> = {
  readonly [K in keyof T] : T[K] extends Object ? DeepReadonly<T[K]> : T[K];
}
console.groupEnd();

console.groupCollapsed(
  "Record<Keys, Type> // 构造一个对象类型, 其属性键为 Keys, 其属性值为 Type."
);
// Record<Keys, Type> // 构造一个对象类型, 其属性键为 Keys, 其属性值为 Type.
type Language = "zh" | "us" | "uk" | "jp";
type Info = { title: string; description: string; visible: boolean };
const i18n: Record<Language, Info> = {
  zh: { title: "中国", description: "中", visible: true },
  us: { title: "美国", description: "美", visible: true },
  uk: { title: "英国", description: "加", visible: true },
  jp: { title: "日本", description: "日", visible: true },
};
console.log(i18n.zh, i18n.us);
console.groupEnd();
console.groupCollapsed(
  "Pick<Type, Keys> // 通过从 Type 中选取一组属性 Keys 来构造一个类型"
);
// Pick<Type, Keys> // 通过从 Type 中选取一组属性 Keys 来构造一个类型
interface Options {
  methods: "GET" | "POST" | "PUT";
  url: string;
  data?: any;
}
type Web = Required<Pick<Options, "methods">>;
function get(options: Web) {
  console.log(options.methods);
}
// error TS2353: Object literal may only specify known properties, and 'url' does not exist in type 'Required<Pick<Options, "methods">>'.
console.log(get({ methods: "POST", url: "hello.com" }));
console.groupEnd();
console.groupCollapsed(
  "Omit<Type, Keys> // 通过从 Type 中选择所有属性然后删除 Keys 来构造一个类型, 与 Pick 相反"
);
// Omit<Type, Keys> // 通过从 Type 中选择所有属性然后删除 Keys 来构造一个类型, 与 Pick 相反
type Web2 = Required<Omit<Options, "methods" | "data">>;
console.groupEnd();

// Extract 和 Exclude 主要用于联合类型，对象属性处理应该用 keyof、Pick、Omit 等工具类型
console.groupCollapsed(
  "Exclude<UnionType, ExcludedMembers> // 通过从 UnionType 中排除所有分配给 ExcludedMembers 的联合成员来构造一个类型"
);
// Exclude<UnionType, ExcludedMembers> // 通过从 UnionType 中排除所有分配给 ExcludedMembers 的联合成员来构造一个类型
type Hobbies = "js" | "css" | "html" | "java" | "golang" | "rust" | "docker";
type E1 = Exclude<Hobbies, "js" | "css" | "html">;
type E2 = Exclude<Hobbies, "java" | "golang" | "rust" | "docker">;
type E3 = Exclude<Hobbies, "docker">;
console.groupEnd();
console.groupCollapsed(
  "Extract<Type, Union> // 通过从 Type 中提取所有可分配给 Union 的联合成员来构造一个类型"
);
// Extract<Type, Union> // 通过从 Type 中提取所有可分配给 Union 的联合成员来构造一个类型
type Ext1 = Extract<Hobbies, "js" | "css" | "html">;
type Ext2 = Extract<Hobbies, "js" | "golang" | "docker" | "python">;
console.log('--------------');
type Ext3 = Extract<keyof Options, 'url' | 'methods'>;
console.groupEnd();

console.groupCollapsed(
  "NonNullable<Type> // 通过从 Type 中排除 null 和 undefined 来构造一个类型"
);
// NonNullable<Type> // 通过从 Type 中排除 null 和 undefined 来构造一个类型
type N1 = NonNullable<string | number | undefined>;
type N2 = NonNullable<boolean | null | undefined>;
console.groupEnd();
console.groupCollapsed(
  "Parameters<Type> // 从函数类型 Type 的参数中使用的类型构造元组类型"
);
// Parameters<Type> // 从函数类型 Type 的参数中使用的类型构造元组类型
type P1 = Parameters<() => void>;
type P2 = Parameters<(s: string, n: number) => void>;
type P3 = Parameters<<T>(arg: T) => T>;
console.groupEnd();
console.groupCollapsed(
  "ConstructorParameters<Type> // 从构造函数类型的类型构造元组或数组类型, 生成一个包含所有参数类型的元组类型(如果 Type 不是函数类型则生成类型 never)"
);
// ConstructorParameters<Type> // 从构造函数类型的类型构造元组或数组类型, 生成一个包含所有参数类型的元组类型(如果 Type 不是函数类型则生成类型 never)
type Q1 = ConstructorParameters<FunctionConstructor>;
class Q {
  constructor(a: number, b: string) {}
}
type Q2 = ConstructorParameters<typeof Q>;
console.groupEnd();
console.groupCollapsed(
  "ReturnType<Type> // 构造一个由函数 Type 的返回类型组成的类型, 对于重载函数, 将是最后一个签名的返回类型"
);
// ReturnType<Type> // 构造一个由函数 Type 的返回类型组成的类型, 对于重载函数, 将是最后一个签名的返回类型
type R1 = ReturnType<() => void>;
type R2 = ReturnType<() => string>;
type R3 = ReturnType<<T>() => T>;
console.groupEnd();
console.groupCollapsed(
  "InstanceType<Type> // 构造一个由 Type 中的构造函数的实例类型组成的类型"
);
// InstanceType<Type> // 构造一个由 Type 中的构造函数的实例类型组成的类型
class I {
  x = 0;
  y = 10;
}
type I1 = InstanceType<typeof I>;
type I2 = InstanceType<any>;
type I3 = InstanceType<never>;
console.groupEnd();
console.groupCollapsed(
  "NoInfer<Type> // 阻止对所包含类型的推断, 除了阻止推断之外, NoInfer<Type> 与 Type 相同"
);
// NoInfer<Type> // 阻止对所包含类型的推断, 除了阻止推断之外, NoInfer<Type> 与 Type 相同
function createStreetLight<C extends string>(
  colors: C[],
  defaultColor?: NoInfer<C>
) {
  // ...
}
createStreetLight(["red", "yellow", "green"], "red"); // OK
// error TS2345: Argument of type '"blue"' is not assignable to parameter of type '"red" | "green" | "yellow"'.
createStreetLight(["red", "yellow", "green"], "blue"); // Error
console.groupEnd();
console.groupCollapsed(
  "ThisParameterType<Type> // 提取函数类型的 this 参数的类型, 如果函数类型没有 this 参数则提取 unknown."
);
// ThisParameterType<Type> // 提取函数类型的 this 参数的类型, 如果函数类型没有 this 参数则提取 unknown.
function toHex(this: Number) {
  return this.toString(16);
}
function numberToString(n: ThisParameterType<typeof toHex>) {
  return toHex.apply(n);
}
console.log(numberToString(2));
console.groupEnd();
console.groupCollapsed(
  "OmitThisParameter<Type> // 从 Type 中删除 this 参数, 如果 Type 没有显式声明的 this 参数, 则结果只是 Type. 否则将从 Type 创建一个没有 this 参数的新函数类型. 泛型被删除, 只有最后一个重载签名被传播到新的函数类型中."
);
// OmitThisParameter<Type> // 从 Type 中删除 this 参数, 如果 Type 没有显式声明的 this 参数, 则结果只是 Type. 否则将从 Type 创建一个没有 this 参数的新函数类型. 泛型被删除, 只有最后一个重载签名被传播到新的函数类型中.
const fiveToHex: OmitThisParameter<typeof toHex> = toHex.bind(5);
console.log(fiveToHex());
console.groupEnd();
console.groupCollapsed(
  "ThisType<Type> // 此工具不返回转换后的类型, 相反, 它用作上下文 this 类型的标记, 必须启用 noImplicitThis 标志才能使用此工具"
);
// ThisType<Type> // 此工具不返回转换后的类型, 相反, 它用作上下文 this 类型的标记, 必须启用 noImplicitThis 标志才能使用此工具
console.groupEnd();
console.groupCollapsed(
  "Uppercase<StringType> // 将字符串中的每个字符转换为大写版本"
);
// Uppercase<StringType> // 将字符串中的每个字符转换为大写版本
type U1 = Uppercase<"hello">;
type ASCIICacheKey<Str extends string> = `ID-${Uppercase<Str>}`;
type MainID = ASCIICacheKey<"my_app">;
console.groupEnd();
console.groupCollapsed(
  "Lowercase<StringType> // 将字符串中的每个字符转换为小写版本"
);
// Lowercase<StringType> // 将字符串中的每个字符转换为小写版本
type L1 = Lowercase<"HELLO">;
console.groupEnd();
console.groupCollapsed(
  "Capitalize<StringType> // 将字符串中的第一个字符转换为等效的大写字母"
);
// Capitalize<StringType> // 将字符串中的第一个字符转换为等效的大写字母
type Cap1 = Capitalize<"hello, world">;
console.groupEnd();
console.groupCollapsed(
  "Uncapitalize<StringType> // 将字符串中的第一个字符转换为等效的小写字母"
);
// Uncapitalize<StringType> // 将字符串中的第一个字符转换为等效的小写字母
type Unc1 = Uncapitalize<"HELLO, WORLD">;
console.groupEnd();
console.groupEnd();
