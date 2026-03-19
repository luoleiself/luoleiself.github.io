// interface
console.group("接口, 声明对值所具有的结构进行类型检查时的约束");
interface labeledValue {
  label: string;
}
function printLabel(labelObj: labeledValue) {
  console.log(labelObj.label);
}
let labelParam = { size: 10, label: "Size 10 object." };
printLabel(labelParam);
// 可选属性, 可以对可能存在的属性进行预定义, 可以捕获引用了不存在的属性时的错误
// 额外属性检查, 绕过额外属性检查的最简方式是类型断言，最佳方式是字符串索引签名 interface C {name: string; age: number; [propName: string]: string}
console.groupCollapsed(
  "可选属性, 可以对可能存在的属性进行预定义, 可以捕获引用了不存在的属性时的错误"
);
console.log(`interface SquareConfig {
  color?: string;
  width?: number;
}`);
interface SquareConfig {
  color?: string;
  width?: number;
}
function createSquare(config: SquareConfig): { color: string; area: number } {
  let newSquare = { color: "white", area: 100 };
  if (config.color) {
    newSquare.color = config.color;
  }
  if (config.width) {
    newSquare.area = config.width * config.width;
  }
  return newSquare;
}
console.log(createSquare({ color: "red" }));
console.log(createSquare({ color: "blue", width: 5 }));
// error TS2561: Object literal may only specify known properties, but 'colour' does not exist in type 'SquareConfig'. Did you mean to write 'color'?
// console.log(createSquare({ colour: 'green', width: 5 }));
console.log(createSquare({ colour: "green", width: 7 } as SquareConfig));
// 可索引的类型
interface StringArray {
  [index: number]: string;
}
let myArray: StringArray = ["Bob", "Fred"];
let myStr: string = myArray[1];
console.groupEnd();
// 只读属性, 只能在对象刚刚创建时修改其值
console.groupCollapsed("只读属性, 只能在对象刚刚创建时修改其值");
console.log(`interface Point {
  readonly x: number;
  readonly y: number;
}`);
interface Point {
  readonly x: number;
  readonly y: number;
}
let p1: Point = { x: 10, y: 20 };
// p1.x = 50; // error TS2540: Cannot assign to 'x' because it is a read-only property.
console.error(
  "p1.x = 50; // error TS2540: Cannot assign to 'x' because it is a read-only property."
);
console.groupEnd();
let ro: ReadonlyArray<number> = [1, 2, 3, 4, 5];
// ro.push(6); // error TS2339: Property 'push' does not exist on type 'readonly number[]'.
// 函数类型, 包含只有参数列表和返回值类型的函数定义
console.groupCollapsed("函数类型");
interface SearchFunc {
  (source: string, subString: string): boolean;
}
let mySearch: SearchFunc = function (src: string, sub: string) {
  let result = src.search(sub);
  return result > -1;
};
console.groupEnd();
// 类实现接口, 接口描述了类的公共部分, 类型检查器只对其实例部分进行类型检查, 构造函数属于类的一部分(静态部分)
interface ClockInterface {
  currentTime: Date;
  setTime(d: Date): void;
}
class Clock implements ClockInterface {
  currentTime: Date;
  setTime(d: Date) {
    this.currentTime = d;
  }
  constructor(h: number, m: number) { }
}
let d = new Date();
let c: ClockInterface = new Clock(d.getHours(), d.getMinutes());
c.setTime(new Date())
console.log(d);
// 继承接口,
console.groupCollapsed("继承接口,");
console.log(`
interface Shape {
  color: string;
}
interface PenStroke {
  penWidth: number;
}
interface Square extends Shape, PenStroke {
  sideLength: number;
}
let square = <Square>{}; // 类型断言
square.color = 'green';
square.sideLength = 10;
square.penWidth = 25;
`);
interface Shape {
  color: string;
}
interface PenStroke {
  penWidth: number;
}
interface Square extends Shape, PenStroke {
  sideLength: number;
}
let square = <Square>{};
square.color = "green";
square.sideLength = 10;
square.penWidth = 25;
console.log(square);
console.groupEnd();
console.groupCollapsed("混合类型");
console.log(`
interface Counter {
  (start: number): string;
  interval: number;
  reset(): void;
}

function getCounter(): Counter {
  let counter = <Counter>function (start: number) { };
  counter.interval = 123;
  counter.reset = function () { };
  return counter;
}

let c = getCounter();
c(10);
c.reset();
c.interval = 5.0;`);
console.groupEnd();
// 接口继承类, 只会继承类的成员但不包括实现, 如果一个接口继承了一个拥有私有或受保护的成员的类时, 这个接口类型只能被这个类或其子类所实现(implements)
console.groupCollapsed(
  "接口继承类, 只会继承类的成员但不包括实现, 如果一个接口继承了一个拥有私有或受保护的成员的类时, 这个接口类型只能被这个类或其子类所实现(implements)"
);
class Control {
  private state: any;
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(private name: string) { }
  print() {
    console.log(this.name + " toString func...");
  }
}
// SelectableControl 包含了 Control 的所有成员, 包括私有成员 state
// 因为 state 是私有成员, 所以只有 Control 或者 Control 的子类才能实现 SelectableControl 接口
interface SelectableControl extends Control {
  select(): void;
}
class Button extends Control implements SelectableControl {
  constructor(name: string) {
    super(name);
  }
  select() { }
}
class TextBox extends Control {
  constructor(name: string) {
    super(name);
  }
  select() { }
}
// Images 没有继承 Control 类, 不能实现 SelectableControl 接口
// error TS2420: Class 'Images' incorrectly implements interface 'SelectableControl'. Property 'state' is missing in type 'Images' but required in type 'SelectableControl'.
class Images implements SelectableControl {
  print(): void {
    throw new Error("Method not implemented.");
  }
  select() { }
}
let b: Button = new Button("Button");
let t: TextBox = new TextBox("TextBox");
console.log(b, t);
b.print();
t.print();
console.log(`
class Control {
  private state: any;
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(private name: string) {}
  print() {
    console.log(this.name + ' toString func...');
  }
}
// SelectableControl 包含了 Control 的所有成员, 包括私有成员 state
// 因为 state 是私有成员, 所以只有 Control 或者 Control 的子类才能实现 SelectableControl 接口
interface SelectableControl extends Control {
  select(): void;
}
class Button extends Control implements SelectableControl {
  constructor(name: string) {
    super(name);
  }
  select() {}
}
class TextBox extends Control {
  constructor(name: string) {
    super(name);
  }
  select() {}
}
// Images 没有继承 Control 类, 不能实现 SelectableControl 接口
// error TS2420: Class 'Images' incorrectly implements interface 'SelectableControl'. Property 'state' is missing in type 'Images' but required in type 'SelectableControl'.
class Images implements SelectableControl {
  print(): void {
    throw new Error('Method not implemented.');
  }
  select() {}
}
let b: Button = new Button('Button');
let t: TextBox = new TextBox('TextBox');
console.log(b, t);
b.print();
t.print();`);
console.groupEnd();
console.groupEnd();

// 类, 构造函数使用 protected 修饰符后, 这个类不能在包含它的类外被实例化, 但是能被继承.
console.group("类");
console.groupCollapsed(
  "public  默认修饰符, 在任何位置都可以访问该修饰符修饰的成员"
);
class Greeter {
  constructor(public greeting: string) { }
  public greet() {
    return "Hello, " + this.greeting;
  }
}
let gt = new Greeter("world");
console.log(gt);
console.log(gt.greet());
console.groupEnd();
// private
console.groupCollapsed("private 不能在该修饰符的类的外部访问");
class Animal {
  constructor(private name: string) { }
}
class Rhino extends Animal {
  constructor(name: string) {
    super(name);
  }
  public say() {
    // error TS2341: Property 'name' is private and only accessible within class 'Animal'.
    console.log(this.name);
  }
}
console.log(`
class Animal {
  constructor(private name: string) { }
}
class Rhino extends Animal {
  constructor(name: string) {
    super(name);
  }
  public say() {
    // error TS2341: Property 'name' is private and only accessible within class 'Animal'.
    console.log(this.name);
  }
}`);
console.groupEnd();
// protected
console.groupCollapsed("protected 只能在该修饰符的类或子类中访问");
class Person {
  protected name: string;
  constructor(name: string) {
    this.name = name;
  }
}
class Employee extends Person {
  private department: string;
  constructor(name: string, department: string) {
    super(name);
    this.department = department;
  }
  public getElevatorPitch() {
    console.log("Employee public func...");
    return `Hello, my name is ${this.name} and I work in ${this.department}`;
  }
}
let howard = new Employee("Howard", "Sales");
// TS2445: Property 'name' is protected and only accessible within class 'Person' and its subclasses.
console.log(howard.name);
console.log(howard.getElevatorPitch());
console.log(`
class Person {
  protected name: string;
  constructor(name: string) {
    this.name = name;
  }
}
class Employee extends Person {
  private department: string;
  constructor(name: string, department: string) {
    super(name);
    this.department = department;
  }
  public getElevatorPitch() {
    console.log('Employee public func...');
    return \`Hello, my name is \$\{this.name\} and I work in \$\{this.department\}\`;
  }
}
let howard = new Employee('Howard', 'Sales');
// TS2445: Property 'name' is protected and only accessible within class 'Person' and its subclasses.
console.log(howard.name);
console.log(howard.getElevatorPitch());`);
console.groupEnd();
// readonly
console.groupCollapsed("readonly 只读属性必须在声明时或构造函数里被初始化");
class Octopus {
  readonly numberOfLegs: number = 10;
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(readonly name: string) { }
}
let dad = new Octopus("Man with the 8 strong legs.");
// error TS2540: Cannot assign to 'name' because it is a read-only property.
dad.name = "Man with the 3-piece suit.";
console.log(dad);
console.log(`
class Octopus {
  readonly numberOfLegs: number = 10;
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(readonly name: string) {}
}
let dad = new Octopus('Man with the 8 strong legs.');
// error TS2540: Cannot assign to 'name' because it is a read-only property.
dad.name = 'Man with the 3-piece suit.';`);
console.groupEnd();
// static
console.groupCollapsed("static 只存在于类本身上而不是类的实例上.");
class Grid {
  static origin = { x: 0, y: 0 };
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(public scale: number) { }
  calculateDistanceFormOrigin(point: { x: number; y: number }) {
    let xDist = point.x - Grid.origin.x;
    let yDist = point.y - Grid.origin.y;
    return Math.sqrt(xDist * xDist + yDist * yDist) / this.scale;
  }
}
let grid1 = new Grid(1.0);
let grid2 = new Grid(5.0);
console.log(grid1.calculateDistanceFormOrigin({ x: 10, y: 10 }));
console.log(grid2.calculateDistanceFormOrigin({ x: 10, y: 10 }));
console.log(`
class Grid {
  static origin = { x: 0, y: 0 };
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(public scale: number) {}
  calculateDistanceFormOrigin(point: { x: number; y: number }) {
    let xDist = point.x - Grid.origin.x;
    let yDist = point.y - Grid.origin.y;
    return Math.sqrt(xDist * xDist + yDist * yDist) / this.scale;
  }
}
let grid1 = new Grid(1.0);
let grid2 = new Grid(5.0);
console.log(grid1.calculateDistanceFormOrigin({ x: 10, y: 10 }));
console.log(grid2.calculateDistanceFormOrigin({ x: 10, y: 10 }));`);
console.groupEnd();
// 存取器
console.groupCollapsed("存取器");
let passCode = "secret passCode";
class St {
  private _fullName: string;
  get fullName(): string {
    return this._fullName;
  }
  set fullName(name: string) {
    if (passCode && passCode == "secret passcode") {
      this._fullName = name;
    } else {
      console.log("Error: Unauthorized update of st.");
    }
  }
}
let s = new St();
s.fullName = "Bob Smith";
if (s.fullName) {
  alert(s.fullName);
}
console.groupEnd();
// abstract
console.groupCollapsed(
  "abstract 抽象类可以包含成员的实现细节, 不能被直接实例化, 一般作为基类使用, 抽象类中的抽象方法不包含具体的实现并且必须在派生类中实现"
);
abstract class Department {
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(public name: string) { }
  // 普通方法
  printName(): void {
    console.log("Department name: " + this.name);
  }
  // 抽象方法
  abstract printMeeting(): void;
}
class AccountingDepartment extends Department {
  constructor() {
    super("Accounting and Auditing");
  }
  printMeeting(): void {
    console.log("The Accounting Department meets each Monday at 10am.");
  }
  generateReports(): void {
    console.log("Generating accounting reports...");
  }
}
let de: Department;
// error TS2511: Cannot create an instance of an abstract class.
de = new Department();
de = new AccountingDepartment();
console.log(de);
de.printName();
de.printMeeting();
// error TS2339: Property 'generateReports' does not exist on type 'Department'.
de.generateReports();
console.log(`
abstract class Department {
  // 构造函数的参数使用 public, private, protected, readonly 修饰符等于同时创建了同名的成员变量
  constructor(public name: string) {}
  // 普通方法
  printName(): void {
    console.log('Department name: ' + this.name);
  }
  // 抽象方法
  abstract printMeeting(): void;
}
class AccountingDepartment extends Department {
  constructor() {
    super('Accounting and Auditing');
  }
  printMeeting(): void {
    console.log('The Accounting Department meets each Monday at 10am.');
  }
  generateReports(): void {
    console.log('Generating accounting reports...');
  }
}
let de: Department;
// error TS2511: Cannot create an instance of an abstract class.
de = new Department();
de = new AccountingDepartment();
console.log(de);
de.printName();
de.printMeeting();
// error TS2339: Property 'generateReports' does not exist on type 'Department'.
de.generateReports();`);
console.groupEnd();
console.groupEnd();
