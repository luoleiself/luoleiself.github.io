# 运算符

!! 将任意类型强制转换为 布尔 类型，与 Boolean() 的效果相同, 但更简洁

```ts
// 第一个 ! 将值转换为相反的布尔值, 第二个 ! 再将结果反转回来得到原始值的真实布尔表示

// !!undefined  => false
// !!null       => false
// !!0,-0       => false
// !!NaN        => false
// !!''         => false
// !!false      => false
// !!其他所有值  => true

const value = 'some string';
console.log(!!value); // true
```

! 尾非空断言符, 用于告诉 TS 编译器 "我确定这个值不是 null 或 undefined, 请跳过空值检查"

- 强制类型系统忽略可能的 null/undefined
- 避免编译器报错
- 不产生任何运行时效果(编译后会直接移除)

```ts
// 函数返回值的非空断言
function maybeString(): string | undefined {
  return Math.random() > 0.5 ? 'hello' : undefined;
}
const str = maybeString()!; // 断言返回值非空
console.log(str.toUpperCase()); // 运行时可能报错

// 数组查找结果的非空断言
const arr: (string | null)[] = ['a', null, 'b'];
const item = arr.find(x => x !== null)!; // 断言找到的元素非空

// React 的 Ref 操作
const inputRef = useRef<HTMLInputElement>(null);
useEffect(() => {
  inputRef.current!.focus(); // 断言选然后 ref 已绑定
},[])
```

!. 非空断言操作符, 是一个组合操作符, 用于告诉 TS 编译器 "我确定这个值不是 null 或 undefined, 请允许我访问它的属性"

- ?. 是安全的，遇到 null/undefined 停止并返回 undefined
- !. 不安全的, 告诉编译器此处不会为 null/undefined, 会禁用 TS 的空检查可能会导致运行时报错, 优先使用 可选链 或 空值检查

```ts
interface User {
  name: string
  address?: {
    street: string
  }
}
const user: User = {name: 'Alice'};
// 正常情况下会报错, 因为 address 可能是 undefined
const street = user.address.street; // 错误: Object is possibly 'undefined'

// 使用 !. 断言
const street = user.address!.street; // 通过编译, 可能在运行时报错

// 使用类型保护
function isNotNull<T>(value: T | null): value is T {
  return value !== null;
}
```
