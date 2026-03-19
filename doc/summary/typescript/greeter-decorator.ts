
// 装饰器
console.group(
  "装饰器, 是一种特殊类型的声明, 它能够被附加到类声明, 方法, 访问符, 属性或参数上. 装饰器使用 @expression 形式, expression 求值后必须为一个函数, 它会在运行时被调用, 被装饰的声明信息做为参数传入" +
  "\n为在类的声明及成员上通过元编程语法添加标注提供了一种方式. " +
  "\n不能用在函数上, 因为存在函数提升"
);
// 装饰器求值顺序
//  1. 参数装饰器, 然后依次是方法装饰器, 访问符装饰器, 或属性装饰器应用到每个实例成员
//  2. 参数装饰器, 然后依次是方法装饰器, 访问符装饰器, 或属性装饰器应用到每个静态成员
//  3. 参数装饰器应用到构造函数
//  4. 类装饰器应用到类
console.log(`
// 装饰器求值顺序
//  1. 参数装饰器, 然后依次是方法装饰器, 访问符装饰器, 或属性装饰器应用到每个实例成员
//  2. 参数装饰器, 然后依次是方法装饰器, 访问符装饰器, 或属性装饰器应用到每个静态成员
//  3. 参数装饰器应用到构造函数
//  4. 类装饰器应用到类`);
function f() {
  console.log("f(): evaluated...");
  return function (
    target: any,
    propertyKey: Object
  ) {
    console.log("f(): called...", target, propertyKey);
  };
}
function g() {
  console.log("g(): evaluated...");
  return function (
    target: any,
    propertyKey: Object
  ) {
    console.log("g(): called...", target, propertyKey);
  };
}

class H {
  @f()
  name: string

  @f()
  @g()
  say() { }
}

// function sealed(constructor: Function) {
//   Object.seal(constructor);
//   Object.seal(constructor.prototype);
//   console.log('sealed...', constructor);
//   return function (
//     target: any,
//     propertyKey: string,
//     descriptor: PropertyDecorator
//   ) {
//     console.log(target, propertyKey, descriptor);
//   };
// }
// function enumerable(value: boolean) {
//   return function (
//     target: any,
//     propertyKey: string,
//     descriptor: PropertyDecorator
//   ) {
//     console.log(target, propertyKey, descriptor);
//     descriptor.enumerable = value;
//   };
// }
// function readonly(target: any, propertyKey: string, descriptor) {
//   console.log('readonly...', target, propertyKey, descriptor);
//   return descriptor;
// }
// function replaceMethod(target: any, propertyKey: string, descriptor) {
//   console.log('replaceMethod', target, propertyKey, descriptor);
//   target[propertyKey] = function () {
//     return `Hi, ${this.greeting}`;
//   };
//   return function () {
//     `Hi, ${this.greeting}`;
//   };
// }
// @sealed
// class Greeters {
//   // @readonly
//   greeting: string = 'Greeter';

//   constructor(message: string) {
//     this.greeting = message;
//   }

//   say() {
//     return 'Hello, say!';
//   }

//   @enumerable(false)
//   @replaceMethod
//   greet() {
//     return 'Hello, ' + this.greeting;
//   }
// }
// var gts = new Greeters('world');
// console.log(gts.greet());
console.log("---------------------")

function ValidataProperty(validateFn: (value: any) => boolean) {
  return function (target: Object, propertyKey: any) {
    let value: any;

    const getter = function () {
      return value;
    }

    const setter = function (newVal: any) {
      if (!validateFn(newVal)) {
        throw new Error(`Invalid value for ${propertyKey}`);
      }
      value = newVal;
    }

    Object.defineProperty(target, propertyKey, {
      get: getter,
      set: setter,
      enumerable: true,
      configurable: true,
    })
  }
}
class User {
  @ValidataProperty(value => typeof value === 'string' && value.length > 0)
  name: string

  @ValidataProperty(value => typeof value === 'number' && value >= 0)
  age: number

  constructor(name: string, age: number) {
    this.name = name;
    this.age = age;
  }
}
var u = new User('zhang san', 18)
console.groupEnd();
