
// 函数
console.group("函数");
let add: (x: number, y: number) => number = function (
  x: number,
  y: number
): number {
  return x + y;
};
// 上下文归类
let add2: (x: number, y: number) => number = function (x, y): number {
  return x + y;
};
console.log(add(1, 2));
// 可选参数和默认参数
console.groupCollapsed(
  "可选参数和默认参数, 可选参数必须放在必须参数的后面. 默认参数不必须放在必须参数的后面,如果出现在必须参数前面则需要明确传入 undefined 值来获得默认值"
);
function buildName(
  fistName: string,
  lastName?: string,
  label: string = "Smith"
): string {
  return fistName + " " + lastName + " , " + label;
}
console.log(buildName("Bob"));
console.log(buildName("Bob", "Adams"));
console.log(buildName("Bob", "Adams", "Sr."));
console.groupEnd();
// 剩余参数
console.groupCollapsed("剩余参数, 必须放在参数列表末尾");
function build(
  firstName: string,
  lastName: string,
  // ...resetOfName: Array<string>, // 泛型
  ...restOfName: string[]
): string {
  return firstName + " " + lastName + " " + restOfName.join(" ");
}
console.log(build("Bob", "Adams", "Cat", "dat", "eric", "fedora"));
console.log(`
function build(
  firstName: string,
  lastName: string,
  // ...resetOfName: Array<string>, // 泛型
  ...restOfName: string[]
): string {
  return firstName + ' ' + lastName + ' ' + restOfName.join(' ');
}
console.log(build('Bob', 'Adams', 'Cat', 'dat', 'eric', 'fedora'));`);
console.groupEnd();
console.groupCollapsed(
  "重载, js 不支持重载, ts 方法为同一个函数提供多个函数类型定义进行函数重载. 定义重载时, 把最精确的定义放在最前面, 编译器查找重载列表时, 如果匹配成功则直接使用. function pickCard(x): any 不是重载列表的一部分"
);
let suits = ["hearts", "spades", "clubs", "diamonds"];
// function pickCard(x: Array<{ suit: string; card: number }>): number; // 泛型
function pickCard(x: { suit: string; card: number }[]): number;
function pickCard(x: number): { suit: string; card: number };
function pickCard(x): any {
  if (typeof x == "object") {
    let pickedCard = Math.floor(Math.random() * x.length);
    return pickedCard;
  } else if (typeof x == "number") {
    let pickedSuit = Math.floor(x / 13);
    return { suit: suits[pickedSuit], card: x % 13 };
  }
}
let myDeck = [
  { suit: "diamonds", card: 2 },
  { suit: "spades", card: 10 },
  { suit: "hearts", card: 4 },
];
let pickedCard1 = myDeck[pickCard(myDeck)];
console.log("card: " + pickedCard1.card + " of " + pickedCard1.suit);

let pickedCard2 = pickCard(15);
console.log("card: " + pickedCard2.card + " of " + pickedCard2.suit);
console.log(`
let suits = ['hearts', 'spades', 'clubs', 'diamonds'];
// function pickCard(x: Array<{ suit: string; card: number }>): number; // 泛型
function pickCard(x: { suit: string; card: number }[]): number;
function pickCard(x: number): { suit: string; card: number };
function pickCard(x): any {
  if (typeof x == 'object') {
    let pickedCard = Math.floor(Math.random() * x.length);
    return pickedCard;
  } else if (typeof x == 'number') {
    let pickedSuit = Math.floor(x / 13);
    return { suit: suits[pickedSuit], card: x % 13 };
  }
}
let myDeck = [
  { suit: 'diamonds', card: 2 },
  { suit: 'spades', card: 10 },
  { suit: 'hearts', card: 4 },
];
let pickedCard1 = myDeck[pickCard(myDeck)];
console.log('card: ' + pickedCard1.card + ' of ' + pickedCard1.suit);

let pickedCard2 = pickCard(15);
console.log('card: ' + pickedCard2.card + ' of ' + pickedCard2.suit);`);
console.groupEnd();
console.groupEnd();