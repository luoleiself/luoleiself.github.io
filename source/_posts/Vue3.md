---
title: Vue3
date: 2021-06-19 15:19:26
categories:
  - [ES, VueJs]
tags:
  - js
  - VueJs
  - Vue3
---

## 全局 API

### 应用实例 API

> 调用 createApp 返回一个应用实例，该实例提供了一个应用上下文, 应用实例挂载的整个组件树共享相同的上下文

#### createApp()

创建一个应用实例

- 参数
  - {Object} rootComponent 根组件选项
  - {Object} rootProps 传递给根组件的 props

```javascript
import { createApp } from 'vue';
const app = createApp(
  {
    /* root component options */
  },
  {
    /* root component props */
  }
);
```

#### createSSRApp()

以 `SSR` 模式创建一个应用实例, 用法和 `createApp()` 相同

#### app.mount()

将应用实例挂载到一个容器元素中

- 参数可以是一个实际的 DOM 元素或一个 CSS 选择器, 返回根组件实例
- 如果该组件有 `template` 模板或定义了 `render` 函数, 则替换容器内所有现存的 DOM 节点, 否则使用容器元素的 innerHTML 作为模板

```javascript
import { createApp } from 'vue';
import App from './App.vue';

const app = createApp({
  /* */
});

// app.mount('#app');
app.mount(App);
```

#### app.unmount()

卸载一个已挂载的应用实例, 同时触发该应用组件树内所有组件的卸载生命周期钩子

#### app.provide()

提供一个值, 可以在应用中的所有后代组件中注入使用

- 参数
  - key, 注入的 key
  - value, 注入的 key 对应的值, 返回应用实例本身

```javascript
import { createApp } from 'vue';
const app = createApp({
  inject: ['name'],
  template: '<span>{{name}}</span>',
});
app.provide('name', 'hello world');
```

#### app.component()

注册或查找全局组件, 根据参数个数区分

```javascript
import { createApp } from 'vue';
const app = createApp(/* */);
app.component('my-component', {
  /*组件配置项*/
});

app.component('my-component'); // 查找已注册的组件
```

<!-- more -->

#### app.directive() <em id="directive"></em> <!-- markdownlint-disable-line -->

注册或查找全局指令, 根据参数个数区分

> 自定义指令主要是为了**重用**涉及普通元素的底层 DOM 访问的逻辑
> 只有当所需功能只能通过直接的 DOM 操作来实现时, 才应该使用自定义指令

一个自定义指令由一个包含类似组件生命周期钩子的对象来定义

- 当在组件上使用自定义指令时, 它会始终应用于组件的根节点
- 如果组件存在多个根节点时, 指令将会被忽略并且抛出一个警告

- 钩子参数
  - el 指令绑定的元素, 可用于直接操作 DOM
  - binding 一个对象
    - value 传递给指令的值, 例如 `v-my-directive="1 + 1"` 的值为 2
    - oldValue 之前的值, 仅在 [`beforeUpdate`](#onBeforeUpdate) 和 [`updated`](#onUpdated) 中可用
    - arg 传递给指令的参数, 例如 `v-my-directive:name` 的参数为 `name`
    - modifiers 一个包含修饰符的对象, 例如 `v-my-directive.foo.bar` 的修饰符对象为 `{foo: true, bar: true}`
    - instance 使用该指令的组件实例
    - dir 指令的定义对象
  - vnode 代表绑定元素的底层 VNode
  - prevVnode 之前的渲染中代表指令所绑定元素的 VNode, 仅在 [`beforeUpdate`](#onBeforeUpdate) 和 [`updated`](#onUpdated) 中可用

```html
<template>
  <MyComponent v-my-directive:name.foo.bar="test" />
</template>
<script>
  import { createApp } from 'vue';
  const app = createApp(/* */);
  app.directive('my-directive', {
    /**
     * 自定义指令钩子:
     * created 在绑定元素的 attribute 或事件监听器被应用之前调用
     * beforeMount 在绑定元素的父组件挂载之前调用
     * mounted 绑定元素的父组件被挂载时调用
     * beforeUpdate 在包含组件的 VNode 更新之前调用
     * updated 在包含组件的 VNode 及其子组件的 VNode 更新之后调用
     * beforeUnmount 在绑定元素的父组件卸载之前调用
     * unmounted 卸载绑定元素的父组件时调用
     */
  });

  // 简化形式: 仅需要在 `mounted` 和 `updated` 上实现相同的行为
  app.directive('my-directive', (el, binding) => {
    /* 在 mounted 和 updated 时都调用 */
  });
</script>
```

#### app.use()

安装一个 **插件**, 插件可以是一个包含 `install()` 方法的对象或者是一个安装函数本身

- 参数
  - 第一个参数为插件本身
  - 可选, 第二个参数作为插件选项将会传递给插件的 `install()`方法

```javascript
import { createApp } from 'vue';
const app = createApp(/* */);
// 包含 install 方法的对象
const myPlugins = {
  install(app, options) {
    /* 配置此应用 */
  },
};
app.use(myPlugins, {
  name: 'hello world',
  /* 传递给 install 方法的可选的选项 */
});

// 安装函数
app.use((app, options) => {
  /* 配置此应用 */
});
```

#### app.mixin()

应用一个全局的 mixin, 作用于应用中的每个组件实例 (不推荐使用), 在 Vue 3 中为了向后兼容

#### app.runWithContext()

> Vue 3.3 支持

使用当前应用作为注入上下文执行回调函数, 在回调同步调用期间, 即使没有当前活动的组件实例, inject() 调用也可以从当前应用提供的值中查找注入

```javascript
import { inject } from 'vue';

app.provide('id', 1);

const injected = app.runWithContext(() => {
  return inject('id');
});
console.log(injected); // 1
```

#### app.version

提供当前应用所使用的 Vue 版本号, 插件中可根据此执行不同的逻辑

#### app.config

应用实例暴露出的一个 `config` 对象, 其中包含了对此应用实例的配置

##### app.config.errorHandler

用于为应用实例内抛出的未捕获错误指定一个全局处理函数

##### app.config.warnHandler

用于为 Vue 的运行时警告指定一个自定义处理函数

##### app.config.performance

设置为 `true` 可在浏览器工具的 **性能/时间线** 页启用对组件初始化、编译、渲染和修改的性能表现追踪

##### app.config.compilerOptions

配置 **运行时编译器** 的选项

- app.config.compilerOptions.isCustomElement 用于指定一个检查方法来识别原生自定义元素

- app.config.compilerOptions.whitespace 用于调整模板中空格的处理行为 `condense(default) | preserve`

- app.config.compilerOptions.delimiters 用于调整模板内文本插值的分隔符, 默认 ['{{', '}}']

- app.config.compilerOptions.comments 用于调整是否移除模板中的 HTML 注释

##### app.config.globalProperties

用于注册能够被应用实例内所有组件实例访问到的全局属性的对象, 对 Vue 2 中 `Vue.prototype` 使用方式的一种替代

##### app.config.optionMergeStrategies

用于定义自定义组件选项的合并策略的对象

```javascript
import { createApp } from 'vue';
const app = createApp(/* */);
app.config.errorHandler = (err, instance, info){/* */}

app.config.globalProperties.name = "hello world"
app.config.globalProperties.$xhr = () => {};

app.config.compilerOptions.isCustomElement = (tag){
  return  tag.startsWith('icon-');
}
```

### 通用

#### version

暴露当前所使用的 Vue 的版本号

#### nextTick() <em id="nextTick"></em> <!-- markdownlint-disable-line -->

> Dom 更新不是同步的, Vue 会在 `next tick` 更新周期中缓冲所有状态的修改, 以确保进行了多次状态修改, 每个组件都只会被更新一次

等待下一次 DOM 更新刷新的工具方法, 可以在状态改变后立即使用以等待 DOM 更新完成

- 传递一个回调函数作为参数
- 或者 await 返回的 Promise

```javascript
import { version, nextTick } from 'vue';

console.log(version); // 打印当前使用的 Vue 版本
async function increment() {
  console.log('DOM 还未更新');
  await nextTick();
  console.log('DOM 已更新');
}
```

#### defineComponent() <em id="defineComponent"></em> <!-- markdownlint-disable-line -->

创建一个合成类型的构造函数, 用于手动渲染函数、TSX 和 IDE 工具支持

- 参数为组件选项对象

```javascript
import { createApp, defineComponent } from 'vue';

const MyComponent = defineComponent({
  data() {
    return { count: 1 };
  },
  methods: {
    increment() {
      this.count++;
    },
  },
});
const app = createApp(MyComponent).mount('#app');
```

- Vue 3.3 支持, 备用签名, 旨在与[组合式 API](#compositionapi)和[渲染函数](#renderingfunc)或 JSX 一起使用
- 参数为 setup 函数, 函数名称作为组件名称使用

```javascript
import { createApp, defineComponent, ref, h } from 'vue';

const HelloWorld = defineComponent(
  (props, ctx) => {
    const count = ref(0);

    // 渲染函数或 JSX
    return () => h('div', count.value);
  },
  {
    /* 其他选项, 例如声明 props 和 emits */
    props: {},
  }
);
const app = createApp(HelloWorld).mount('#app');
```

#### defineAsyncComponent()

创建一个只有在需要时才会加载的异步组件

- 参数为配置加载行为的选项对象

```javascript
import { defineAsyncComponent } from 'vue';

const AsyncComp = defineAsyncComponent({
  loader: () => import('./Foo.vue') // 工厂函数
  loadingComponent: LoadingComponent,  // 加载异步组件时要使用的组件
  errorComponent: ErrorComponent, // 加载失败时要使用的组件
  delay: 200, // 在显示 loadingComponent 之前的延迟 | 默认值：200（单位 ms）
  // 如果提供了 timeout ，并且加载组件的时间超过了设定值，将显示错误组件
  // 默认值：Infinity（即永不超时，单位 ms）
  timeout: 3000,
  suspensible: false, // 定义组件是否可挂起 | 默认值：true
  /**
   * @param {*} error 错误信息对象
    * @param {*} retry 一个函数，用于指示当 promise 加载器 reject 时，加载器是否应该重试
    * @param {*} fail  一个函数，指示加载程序结束退出
    * @param {*} attempts 允许的最大重试次数
    */
  onError(error, retry, fail, attempts) {
    if (error.message.match(/fetch/) && attempts <= 3) {
      // 请求发生错误时重试，最多可尝试 3 次
      retry()
    } else {
      // 注意，retry/fail 就像 promise 的 resolve/reject 一样：
      // 必须调用其中一个才能继续错误处理。
      fail()
    }
  }
});
```

- 参数为异步加载函数

```javascript
import { defineAsyncComponent } from 'vue';

// 全局注册异步组件
const AsyncComp = defineAsyncComponent(() =>
  import('./components/AsyncComponent.vue');
);
app.component('async-component', AsyncComp);

// 局部注册异步组件
import { createApp, defineAsyncComponent } from 'vue';
createApp({
  // ...
  components: {
    AsyncComponent: defineAsyncComponent(() =>
      import('./components/AsyncComponent.vue');
    ),
  },
});
```

#### [defineCustomElement()](#defineComponent)

和 `defineComponent()` 接收的参数相同, 不同的是返回一个原生 **自定义元素** 类的构造器

```javascript
import { defineCustomElement } from 'vue';
// <my-element></my-element>
const MyElement = defineCustomElement({});
customElements.define('my-element', MyElement); // 注册该自定义元素
document.body.append(new MyElement(/* 初始化 prop */));
```

## 组合式 API <em id="compositionapi"></em> <!-- markdownlint-disable-line -->

组合式 API 的核心思想是直接在函数作用域内定义响应式状态变量, 并将从多个函数中得到的状态组合起来处理复杂问题, 这种形式更加自由灵活和高效.

组合式 API(composition API) 是一系列 API 的集合, 能够通过函数而不是声明选项的方式书写 Vue 组件来实现更加简洁高效的逻辑复用([选项式 API](#optionalapi) 中主要的逻辑复用机制是 mixins), 涵盖以下方面的 API

- 响应式 API: 例如 [ref()](#ref) 和 [reactive()](#reactive), 可以直接创建响应式状态、计算属性和侦听器
- 生命周期钩子: 例如 [onMounted()](#onMounted) 和 [onUnmounted()](#onUnmounted), 可以在组件各个生命周期阶段添加逻辑
- 依赖注入: 例如 provide() 和 inject(), 可以在使用响应式 API 时, 利用 Vue 的依赖注入系统

### setup()

> 对于结合单文件组件使用的组合式 API 推荐使用 `<script setup>` 语法

`setup()` 是在组件中使用组合式 API 的入口, 通常在两个情况下使用

- 需要在**非单文件组件**中使用组合式 API 时
- 需要在基于[选项式 API](#optionalapi) 的组件中集成基于组合式 API 的代码时

#### 基本使用

- 在创建组件实例时, 在初始 prop 解析之后立即调用 setup
- 在生命周期方面, 在 `beforeCreate` 钩子之前调用

- getCurrentInstance

  - 支持访问内部组件实例，用于高阶用法或库的开发
  - 只能在 setup 或生命周期钩子中调用

- `setup()` 应该同步地返回一个对象, 唯一可以使用 `async setup()` 的情况是该组件是 [\<Suspense\>](#suspense) 组件地后裔

#### 访问 Props

- `setup()` 的第一个参数, 是响应式地并且会在传入新的 props 时同步更新
- 不能直接对 props 进行解构操作, 会丢失响应性, 可以通过 `toRefs()` 和 `toRef()` 工具函数辅助完成

#### 上下文

`setup()` 的第二个参数为一个**上下文**对象, 暴露了其他一些在 `setup()` 中可能会用到的值, 该上下文对象是非响应式的, 可以安全地解构, attrs 和 slots 都是非响应式的, 如果需要根据 attrs 或 slots 的改变执行副作用, 需要在 onBeforeUpdate 钩子中执行相关逻辑

- attrs 透传 Attributes, 等价于 $attrs, 未被声明为 `props` 或 `emits` 的属性或者 `v-on` 事件监听器都将作为 `attrs` 的一部分
  - 当一个组件以单个元素为根作渲染时, 透传的 attribute 会自动被添加到根元素上, 如果有多个根节点时没有自动 attribute 透传行为
- slots [插槽](#v-slot), 等价于 $slots
- emit 触发事件, 等价于 $emit
- expose 用于显示的限制该组件暴露出的属性, 父组件将仅能访问 expose 函数暴露出的内容

```javascript
import { ref, createApp } from 'vue';

const app = createApp({
  setup(props, { emit, expose }) {
    const publicCount = ref(0);
    expose({ count: publicCount });

    emit('my-event', { name: 'hello world' });

    return { publicCount };
  },
});
```

##### 使用 [slot](#v-slot) 渲染

```javascript
import { createApp, defineComponent, h } from 'vue';

const HelloWorld = defineComponent((props, { slots }) => {
  // 使用 `?.` 可选链运算符判断插槽函数不存在则使用默认值渲染
  return () => [
    h(
      'p',
      slots?.default?.() || 'rendered content from self by default slot...'
    ),
    h('p', slots?.header?.() || 'rendered content from self by header slot...'),
  ];
});

const app = createApp({
  setup(props, { slots }) {
    return () =>
      h(
        HelloWorld,
        {
          // VNode 生命周期事件琢磨中...
          'onVue:before-mount': () => {
            console.log('child component hooks triggered...');
          },
        },
        // 传递单个默认插槽函数
        // () => 'rendered content from father component by default slot...',
        // 传递具名插槽函数, 使用插槽函数对象形式传递
        {
          default: () =>
            'rendered content from father component by default slot...',
          header: () =>
            h(
              'span',
              'rendered content from father component by header slot...'
            ),
        }
      );
  },
});
app.component('hello-world', HelloWorld);
app.mount('#app');
```

##### 使用 [slot](#v-slot)scope 渲染

```javascript
import { createApp, defineComponent, h } from 'vue';

const HelloWorld = defineComponent((props, { attrs, slots }) => {
  const message = 'from hello world component';
  const age = attrs.age > 0 ? attrs.age : 18;
  console.log(props); // {}
  console.log(slots); // {default: renderFnWithContext()}
  console.log(attrs); // {name: 'from createApp', age: -1}

  return () => h('p', slots.default({ message: message, age: age }));
});

const app = createApp({
  setup(props, ctx) {
    return () =>
      h(
        HelloWorld,
        { name: 'from createApp', age: -1 },
        // 传递单个默认插槽函数
        (slotScope) =>
          slotScope.message + ' - ' + slotScope.age + ' - others from createApp'
      );
  },
});
app.component('hello-world', HelloWorld);
app.mount('#app');
```

#### 返回[渲染函数](#renderingfunc)

> 返回[**渲染函数**](#renderingfunc)将会阻止返回其他东西, 对于父组件通过模板引用组件暴露的属性使用 `expose()` 方法解决

- 返回一个 [**渲染函数**](#renderingfunc), 此时在渲染函数中可以直接使用在同一作用域下声明的响应式状态

```javascript
import { h, ref, reactive } from 'vue';
const app = createApp({
  setup(props, { expose }) {
    const count = ref(0);
    const object = reactive({ foo: 'bar' });
    const increment = () => ++count.value;

    expose({ increment }); // 组件暴露 increment 方法

    return () => h('div', [count.value, object.foo]);
  },
});
```

### 响应式: 核心

- 响应式状态默认是深层次的, 即对深层次的响应式状态的更改也能被检测到
- 只有代理对象是响应式的, 更改原始对象不会触发更新, 使用响应式系统时仅使用声明对象的代理版本

#### ref() <em id="ref"></em> <!-- markdownlint-disable-line -->

> 当在模板中使用了一个 ref, 然后改变这个 ref 的值, Vue 会自动检测到这个变化并相应地更新 DOM, 这个过程通过一个基于依赖追踪的响应式系统实现的, 当一个组件首次渲染时, Vue 会追踪在渲染过程中使用的每一个 ref, 当一个 ref 被修改时, 它会触发追踪它的组件的一次重新渲染, 在标准的 javascript 中无法检测普通变量的修改, 可以通过 getter 和 setter 方法来拦截对象属性的 get 和 set 操作.
> .value 属性给予了 Vue 一个机会来检测 ref 何时被访问或修改, 在其内部, Vue 在它的 getter 中执行追踪, 在它的 setter 中执行触发, 概念上 ref 可以看作是一个这样的对象.

接受一个内部值, 返回一个响应式可更改的 ref 对象, 此对象只有一个指向其内部值的属性 `.value`

- 将一个对象赋值给 ref, 那么这个对象将通过 `reactive()` 转为具有深层次响应式的对象, 如果对象中包含了嵌套的 ref, 它们将被深层地解包

```javascript
import { ref } from 'vue';
const obj = ref({
  nested: { count: 0 },
  arr: ['foo', 'bar'],
});
obj.value.nested.count++;
obj.value.arr.push('baz');
```

- 一个包含对象类型值的 ref 可以响应式的替换整个对象

```javascript
const count = ref(0);
console.log(count.value); // 0
count.value++;
console.log(count.value); // 1

const objRef = ref({ count: 0 });
objRef.value = { count: 1 }; // 响应式替换
```

- ref 被传递给函数或是从**一般对象**上被解构时, 不会丢失响应性

```javascript
const obj = {
  foo: ref(0),
  bar: ref(1),
};
// 该函数接收一个 ref 需要通过 .value 取值, 但会保持响应性
callSomeFn(obj.foo);
// 解构 ref 仍然是响应性的
const { foo, bar } = obj;
```

- 当 ref 在模板中作为**顶层属性**被访问时, 它们会被自动解包, 不需要使用 `.value`

```html
<script setup>
  import { ref } from 'vue';
  // ref 非渲染上下文顶层属性
  const obj = { foo: ref(1) };

  // ref 为渲染上下文顶层属性
  const count = ref(0);

  function increment() {
    count.value++;
  }
</script>
<template>
  <!-- 不需要使用 .value 访问 -->
  <button @click="increment">{{count}}</button>
  <!-- 需要使用 .value 访问 -->
  <p>{{obj.foo.value}}</p>
</template>
```

#### computed() <em id="computed"></em> <!-- markdownlint-disable-line -->

返回一个只读的响应式 ref 对象, 该 ref 通过 .value 暴露 getter 函数的返回值

- 接受一个 getter 函数
- 接受一个带有 get 和 set 函数的对象

```javascript
// 接受一个 getter 函数
const count = ref(1);
const plusOne = computed(() => count.value + 1);
console.log(plusOne.value); // 2
plusOne.value++; // 错误

// 接受一个带有 get 和 set 函数的对象
const count = ref(1);
const plusOne = computed({
  get: () => count.value + 1,
  set: (val) => {
    count.value = val - 1;
  },
});
plusOne.value = 1;
console.log(count.value); // 0
```

#### reactive() <em id="reactive"></em> <!-- markdownlint-disable-line -->

返回一个对象的响应式代理

- 对同一个**原始对象**调用 `reactive()` 总是返回同样的**代理对象**
- 对一个已存在的**代理对象**调用 `reactive()` 总是返回其本身

**局限性** <!--markdownlint-disable-line-->

- 有限的值类型: 只能用于对象类型(对象、数组、Map、Set 这样的集合类型), 不能持有 `string`, `number`, `boolean` 这样的原始类型
- 不能替换整个对象: 因为 Vue 的响应式系统是通过属性访问进行追踪的, 因此需要始终保持对响应式对象的**相同引用**,
- 对解构操作不友好: 将响应式对象的属性赋值或解构至本地变量时、或是将该属性传入一个函数时将失去响应性

==建议使用 `ref()` 作为生命响应式状态的主要 API==

- 将一个 ref 作为响应式对象的属性被访问或修改时自动解包

```javascript
const count = ref(0);
// ref 会解包
const state = reactive({
  count,
});
console.log(state.count); // 0
console.log(state.count === count.value); // true

// 自动更新 `state.value`
count.value++;
console.log(count.value, state.count); // 1 1

// 自动更新 `count` ref
state.count++;
console.log(count.value, state.count); // 2 2
```

- 将一个新的 ref 赋值给一个关联了已有 ref 的属性, 那么旧的 ref 会被替换

```javascript
const otherCount = ref(4);
state.count = otherCount;
console.log(count.value, state.count); // 2 4
```

- 只有当嵌套在一个深层响应式对象内时, 才会发生 ref 解包, 当其作为 [**浅层响应式对象**](#shallowReactive) 的属性被访问时不会解包
- 当 ref 作为响应式数组或原生集合类型(如 Map)中的元素被访问时, 不会被解包

```javascript
// 浅层响应式对象, ref 作为其属性被访问时不会被解包
const sr = shallowReactive({
  arr: ref([1, 2, 3]),
});

// 原生集合中包含 ref 元素时, ref 不会解包
const books = reactive([ref('Vue 3.0')]);
console.log(books[0].value); // 需要使用 .value

const map = reactive(new Map([['count', ref(0)]]));
console.log(map.get('count').value); // 需要使用 .value
```

#### readonly() <em id="readonly"></em> <!-- markdownlint-disable-line -->

- 接受一个对象(响应式或普通)或一个 ref, 返回原值的只读代理
- 任何被访问的嵌套属性也是只读的, 它的 ref 解包行为与 reactive() 相同, 但解包得到的值是只读的

```javascript
import { reactive, readonly, watchEffect } from 'vue';

const original = reactive({ count: 0 });
const copy = readonly(original);
watchEffect(() => {
  console.log(copy.count); // 用于响应性追踪
});
// 变更 original 会触发依赖于副本的侦听器
original.count++;
// 警告! // 变更副本将失败并导致警告
copy.count++;
```

#### watchEffect() <em id="watchEffect"></em> <!-- markdownlint-disable-line -->

立即执行一个函数, 同时响应式地追踪其依赖, 并在依赖更新时重新执行函数

- 第一个参数是要运行的副作用函数
- 第二个参数是可选项, 用来调整副作用的刷新时机或调试副作用的依赖
- 返回值是一个用来停止该副作用的函数

```javascript
const count = ref(0);
watchEffect(() => {
  console.log(count.value); // 输出 0
});
count.value++; // 输出 1
```

- 副作用清除 onCleanup

  - 副作用即将重新执行时
  - 侦听器被停止(setup 或 lifeCycle Hooks 中使用过, 则在组件卸载时)

```javascript
watchEffect(async (onCleanup) => {
  const { response, cancel } = doAsyncWork(id.value);
  // cancel 会在 id 更改时调用
  // 取消之前未完成的请求
  onCleanup(cancel);
  data.value = await response;
});
```

- 停止侦听器

```javascript
const stop = watchEffect(() => {});
// 不再需要此侦听器时
stop();
```

- 副作用刷新时机, 会在组件更新之前执行副作用

  - 如果需要在组件更新后重新运行侦听器副作用
  - flush
    - pre: '默认值', 指定的回调应该在渲染前被调用
    - post: 将回调推迟到渲染之后调用
    - sync: '始终同步触发', 低效

```javascript
// 在组件更新后触发，这样你就可以访问更新的 DOM。
// 注意：这也将推迟副作用的初始运行，直到组件的首次渲染完成。
watchEffect(() => {}, {
  flush: 'post',
});
```

- 侦听器调试, 只能用于开发模式下

  - onTrack 响应式 property 和 ref 作为依赖项被追踪时被调用
  - onTrigger 依赖项变更导致副作用被触发时被调用

```javascript
watchEffect(() => {}, {
  flush: 'post',
  onTrack(e) {},
  onTrigger(e) {},
});
```

#### [watchPostEffect()](#watchEffect)

`watchEffect()` 使用 flush: 'post' 选项时的别名

#### [watchSyncEffect()](#watchEffect)

`watchEffect()` 使用 flush: 'sync' 选项时的别名

#### watch() <em id="watch"></em> <!-- markdownlint-disable-line -->

- 侦听一个或多个响应式数据源, 并在数据源变化时调用所给的回调函数, 使用方式和 this.$watch 和 watch 选项完全等效
- 默认是浅层侦听, 仅在侦听的属性被赋新值时才触发回调, 而嵌套属性的变化不会触发, 如果需要侦听嵌套属性, 使用 `deep: true` 选项
- 默认是懒侦听的, 仅在侦听源发生变化时才触发回调, 如果需要在创建侦听器时立即执行一遍回调, 使用 `immediate: true` 选项

##### 参数

- 第一个参数是侦听器的源, 支持包含返回值的函数、ref、响应式对象、或者以上类型的值组成的数组
- 第二个参数是侦听源发生变化时调用的函数, 函数接收三个参数: 新值、旧值，及一个用于注册副作用清理的回调函数
- 第三个参数是一个配置项对象

  - immediate 在侦听器创建时立即触发回调, 第一次调用时旧值为 `undefined`
  - deep 如果源是对象, 强制深度遍历, 以便在深层级变更时触发回调
  - flush 调整回调函数的刷新时机, 见 [watchEffect()](#watchEffect)
  - onTrack/onTrigger 调试侦听器的依赖, 见 [watchEffect()](#watchEffect)
  - once 回调函数只会执行一次, 侦听器将在回调函数首次运行后自动停止, 3.4 支持

```javascript
import { reactive, ref, watch } from 'vue';

// 侦听一个 getter 函数
const state = reactive({ count: 0 });
watch(
  () => state.count,
  (newVal, oldValue) => {
    /* */
  },
  // 当侦听 getter 函数, 回调函数只在此函数的返回值变化时才会触发, 监听深层级变更时触发需要设置 {deep: true}
  // 当前侦听一个响应式对象, 默认自动开启深层级模式
  {
    deep: true,
    once: true, // 侦听器只会执行一次后自动停止
  }
);

// 侦听一个 ref
const count = ref(0);
watch(count, (count, prevCount) => {
  /* */
});

// 侦听多个源
const stop = watch([fooRef, barRef], ([foo, bar], [prevFoo, prevBar]) => {
  /* */
});
// 停止侦听器
stop();

// 副作用清理
watch(id, async (newValue, oldValue, onCleanup) => {
  const { response, cancel } = doAsyncWork(newValue);
  // 当 id 变化时, cancel 被调用
  // 取消之前的未完成的请求
  onCleanup(cancel);
  data.value = await response;
});
```

##### 与 [watchEffect()](#watchEffect) 的区别

- 惰性执行副作用
- 更具体地说明应触发侦听器重新运行的状态
- 访问被侦听状态的先前值和当前值
- 侦听多个源

### 响应式: 工具

#### isRef()

检查某个值是否是 ref

#### unref()

如果参数是 ref, 则返回 ref 指向的内部值, 否则返回参数本身. 是 `isRef(val) ? val.value : val` 的一个语法糖

#### toRef()

- Vue 3.3 支持将值、refs 或 getters 规范化为 refs

基于响应式对象上的一个属性, 新创建一个对应的 ref, 此 ref 与其源属性保持同步, 改变源属性的值将更新 ref 的值, 反之亦然

```javascript
import { reactive, toRef } from 'vue';

const state = reactive({ foo: 1, bar: 2 });
const fooRef = toRef(state, 'foo');
// 更改 ref 会更新源属性
fooRef.value++;
console.log(state.foo); // 2
// 更改源属性会更新 ref
state.foo++;
console.log(fooRef.value); // 3
```

#### toValue()

> Vue 3.3 支持

将值、refs 或 getters 规范化为值, 与 unref() 类似, 不同的是此函数也会规范化 getter 函数, 如果参数是一个 getter, 它将会被调用并且返回它的返回值

```javascript
import { toValue, ref } from 'vue';
toValue(1); // 1
toValue(ref(1)); // 1
toValue(() => 1); // 1
```

#### toRefs()

> 方便消费组件可以在不丢失响应性的情况下对返回的对象进行分解/扩散

将一个响应式对象转换为一个普通对象, 这个普通对象的每个属性都指向源对象相应属性的 ref, 每个单独的 ref 都是使用 `toRef()` 创建的

- toRefs 在调用时只为源对象上的可以枚举的属性创建 ref, 如果为可能还不存在的属性创建 ref 时, 使用 toRef

```javascript
import { reactive, toRefs } from 'vue';

const state = reactive({ foo: 1, bar: 2 });
const stateAsRefs = toRefs(state);
/*
stateAsRefs 的类型:
{
  foo: Ref<number>,
  bar: Ref<number>
}
*/

// ref 和原始 property 已经 "链接" 起来了
state.foo++;
console.log(stateAsRefs.foo.value); // 2

stateAsRefs.foo.value++;
console.log(state.foo); // 3
```

#### isProxy()

检查一个对象是否由 `reactive()`, `readonly()`, `shallowReactive()`, `shallowReadonly()` 创建的代理

#### isReactive()

检查一个对象是否由 `reactive()`, `shallowReactive()` 创建的代理

#### isReadonly()

检查对象是否是由 `readonly()`, `shallowReadonly()` 创建的只读代理, 只读对象的属性可以更改, 但不能通过传入的对象直接赋值

### 响应式: 进阶

#### shallowRef()

[ref()](#ref) 的浅层作用形式

- 浅层 ref 的内部值将会原样存储和暴露, 并且不会被深层递归地转为响应式
- 只有对 `.value` 的访问是响应式的

```javascript
import { shallowRef } from 'vue';

const state = shallowRef({ count: 1 });
// 不会触发更改
state.value.count = 2;

// 会触发更新
state.value = { count: 2 };
```

#### triggerRef()

强制触发依赖一个 `浅层 ref` 的副作用, 通常对浅引用的内部值进行深度变更后使用

```javascript
import { shallowRef, watchEffect, triggerRef } from 'vue';

const shallow = shallowRef({ name: 'hello world' });

// 立刻执行一次副作用, 输出 hello world
watchEffect(() => {
  console.log(shallow.value.name);
});

// 更改不会触发副作用, ref 是浅层的
shallow.value.name = 'hello gg';

// 手动触发浅层 ref 的副作用, 输出 hello gg
triggerRef(shallow);
```

#### customRef()

> 它需要一个工厂函数，该函数接收 track 和 trigger 函数作为参数，并且应该返回一个带有 get 和 set 的对象

创建一个自定义的 ref, 显式声明对其依赖追踪和更新触发的控制方式

```html
<template> <input v-model="text" /> </template>

<script setup>
  const text = useDebouncedRef('hello');
</script>

<script>
  import { customRef } from 'vue';

  // 创建一个防抖 ref, 只在最后一次 set 调用后的一段固定间隔后再调用
  function useDebouncedRef(value, delay = 200) {
    let timeout;
    return customRef((track, trigger) => {
      return {
        get() {
          track();
          return value;
        },
        set(newValue) {
          clearTimeout(timeout);
          timeout = setTimeout(() => {
            value = newValue;
            trigger();
          }, delay);
        },
      };
    });
  }
</script>
```

#### shallowReactive() <em id="shallowReactive"></em> <!-- markdownlint-disable-line -->

[reactive()](#reactive) 的浅层作用形式

- 没有深层级的转换, 浅层响应式对象里只有根级别的属性是响应式的
- 属性的值会被原样存储和暴露, 值为 ref 的属性不会自动解包

```javascript
import { shallowReactive, isReactive } from 'vue';

const state = shallowReactive({ foo: 1, nested: { bar: 2 } });
// 更改状态自身的属性是响应式的
state.foo++;
// 嵌套对象不会被转为响应式
isReactive(state.nested); // false
// 不是响应式的
state.nested.bar++;
```

#### shallowReadonly()

[readonly()](#readonly) 的浅层作用形式

- 没有深层级的转换, 只有根层级的属性变为了只读
- 属性的值会被原样存储和暴露, 值为 ref 的属性不会自动解包

```javascript
import { shallowReadonly, isReadonly } from 'vue';

const state = shallowReadonly({ foo: 1, nested: { bar: 2 } });
// 更改状态自身的属性会失败
state.foo++;
// 可以更改嵌套的属性
isReadonly(state.nested); // false
// 允许更改嵌套的属性
state.nested.bar++;
```

#### toRaw()

> 可用于临时读取数据而无需承担代理访问/跟踪开销，也可用于写入数据而避免触发更改. 不建议保留对原始对象的持久引用

返回由 `reactive()`, `readonly()`, `shallowReactive()`, `shallowReadonly()` 创建的的代理对应的原始对象

```javascript
const foo = {};
const reactiveFoo = reactive(foo);
console.log(toRaw(reactiveFoo) === foo); // true
```

#### markRaw()

将一个对象标记为不可转为代理并返回该对象本身

- 有些值不应该是响应式的，例如复杂的第三方类实例或 Vue 组件对象
- 当渲染具有不可变数据源的大列表时，跳过代理转换可以提高性能

```javascript
import { markRaw, reactive, isReactive } from 'vue';

const foo = markRaw({});
console.log(isReactive(reactive(foo))); // false

// 嵌套在其他响应式对象中时也可以使用
const bar = reactive({ foo });
console.log(isReactive(bar.foo)); // false
```

#### effectScope() <em id='effectscope'></em> <!-- markdownlint-disable-line -->

创建一个 effect 作用域, 可以捕获其中所创建的响应式副作用(计算属性和侦听器), 这样捕获到的副作用可以一起处理

```javascript
import { effectScope, watch, watchEffect } from 'vue';

// 创建 effect 作用域
const scope = effectScope();

scope.run(() => {
  const doubled = computed(() => counter.value * 2);

  watch(doubled, (newValue, oldValue) =>
    console.log(doubled.value, newValue, oldValue)
  );

  watchEffect(() => console.log('Count: ', doubled.value));
});

// 停止当前作用域内的所有 effect
scope.stop();
```

#### [getCurrentScope()](#effectscope)

如果存在则返回当前活跃的 effect 作用域

#### [onScopeDispose()](#effectscope)

> 此方法可以作为可复用的组合式函数中 `onUnmounted` 的替代品, 它并不与组件耦合, 因为每个 Vue 组件的 setup 函数也是在一个 effect 作用域中调用的

在当前活跃的 effect 作用域上注册一个处理回调函数, 当相关的 effect 作用域停止时会调用注册的回调函数, 这个方法可以作为可复用的组合式函数中的 onUnmounted 的替代

```javascript
import { onScopeDispose } from 'vue';

onScopeDispose(() => {
  console.log('活跃的 effect 作用域被停止...');
});
```

### 生命周期钩子

> 所有生命周期钩子函数必须在组件的 `setup()` 阶段**同步调用**

#### VNode 生命周期事件

VNode 生命周期事件前缀从 `hook:` 更改为 `vue:`, 这些事件也可用于 HTML 元素, 和在组件上的用法一样

- `vue:` 前缀为固定格式, 生命周期事件名可以使用 `kebab-case` 或者 `camelCase` 格式

```html
<template>
  <!-- Vue 2.x -->
  <child-component @hook:mounted="onMounted"></child-component>

  <!-- Vue 3.x -->
  <child-component @vue:mounted="onMounted"></child-component>
  <child-component @vue:before-update="onBeforeUpdate"></child-component>
  <!-- 等同于 -->
  <child-component @vue:beforeUpdate="onBeforeUpdate"></child-component>
</template>
```

```javascript
import { h, createApp, defineComponent } from 'vue';

const HelloWorld = defineComponent((props, ctx) => {
  return () => h('p', 'hello world component');
});
const app = createApp({
  data() {
    return {};
  },
  template: `<h1>This is template option.</h1>
    <hello-world @vue:before-mount="helloWorldBeforeMount"></hello-world>`,
  methods: {
    helloWorldBeforeMount() {
      console.log('child component hooks before-mount triggered...');
    },
  },
});
app.component('hello-world', HelloWorld);
app.mount('#app');
```

#### onBeforeMount() <em id="onBeforeMount"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件 **挂载之前** 调用, 组件已经完成其响应式状态的设置, 但还没有创建 DOM 节点

#### onMounted() <em id="onMounted"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件 **挂载完成** 之后执行

- 其所有同步子组件都已经被挂载(不包含 **异步组件** 或 [\<Suspense\>](#suspense) 树内的组件)
- 其自身的 DOM 树已经创建完成并插入了父容器中, 仅当根容器存在于文档中

```html
<template>
  <div ref="el"></div>
</template>
<script setup>
  import { ref, onMounted } from 'vue';
  const el = ref();

  onMounted(() => {
    console.log(el.value); // <div>
  });
</script>
```

#### onBeforeUpdate() <em id="onBeforeUpdate"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件因为响应式状态变更而更新其 DOM 树之前调用

#### onUpdated() <em id="onUpdated"></em> <!-- markdownlint-disable-line -->

> 父组件的更新钩子在其子组件的更新钩子之后调用, 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件因为响应式状态变更而更新其 DOM 树之后调用

- 如果需要在某个特定的状态更改后访问更新后的 DOM, 使用 [`nextTick()`](#nextTick)

```html
<template>
  <button id="count" @click="count++">{{count}}</button>
</template>
<script setup>
  import { ref, onUpdated } from 'vue';

  const count = ref(0);

  onUpdated(() => {
    console.log(document.getElementById('count').textContent);
  });
</script>
```

#### onBeforeUnmount() <em id="onBeforeUnmount"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件实例被 **卸载之前** 调用, 此时组件实例还保有全部的功能

#### onUnmounted() <em id="onUnmounted"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件实例被 **卸载之后** 调用

- 其所有子组件都已经被卸载
- 所有相关的响应式作用(**渲染作用** 以及 `setup()` 时创建的计算属性和侦听器)都已经停止

```html
<script setup>
  import { onMounted, onUnmounted } from 'vue';

  let intervalId;
  onMounted(() => {
    intervalId = setInterval(() => {
      /* */
    });
  });

  onUnmounted(() => clearInterval(intervalId));
</script>
```

#### onActivated() <em id="onActivated"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数, 如果组件实例是 `<KeepAlive>` 缓存树的一部分, 当组件被插入到 DOM 中时调用

#### onDeactivated() <em id="onDeactivated"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数, 如果组件实例是 `<KeepAlive>` 缓存树的一部分, 当组件从 DOM 中移除时调用

#### onErrorCaptured() <em id="onErrorCaptured"></em> <!-- markdownlint-disable-line -->

注册一个回调函数在捕获了后代组件传递的错误时调用

##### 错误来源

- 组件渲染
- 事件处理器
- 生命周期钩子
- `setup()` 函数
- 侦听器
- 自定义指令钩子
- 过渡钩子

函数参数包含三个, 通过返回 `false` 阻止错误继续向上传递

- 错误对象
- 触发该错误的组件实例
- 说明错误来源类型的的信息字符串

##### 错误传递规则

- 默认情况下, 所有的错误都会被发送到应用级 `app.config.errorHandler`(前提已经定义), 这样这些错误都能在一个统一的地方报告给分析服务
- 如果组件的继承链或组件链上存在多个 `errorCaptured` 钩子, 对于同一个错误, 这些钩子会被按从底到上的顺序一一调用, 这个过程称为 **向上传递**, 类似于原生 DOM 事件的冒泡机制
- 如果 `errorCaptured` 钩子本身抛出了一个错误, 那么这个错误和原来捕获到的错误都将被发送到 `app.config.errorHandler`
- `errorCaptured` 钩子可以通过返回 `false` 来阻止错误继续向上传递

#### onRenderTracked() <em id="onRenderTracked"></em> <!-- markdownlint-disable-line -->

> 仅在开发模式下可用, 且在服务器端渲染期间不会被调用

注册一个回调函数在组件渲染过程中追踪到响应式依赖时调用

#### onRenderTriggered() <em id="onRenderTriggered"></em> <!-- markdownlint-disable-line -->

> 仅在开发模式下可用, 且在服务器端渲染期间不会被调用

注册一个回调函数在组件的响应式依赖的变更触发了组件渲染时调用

#### onServerPrefetch() <em id="onServerPrefetch"></em> <!-- markdownlint-disable-line -->

注册一个异步函数, 在组件实例在服务器上被渲染之前调用, 可用于执行一些仅存在于服务端的数据抓取过程

- 如果返回一个 Promise, 服务端渲染会在渲染该组件前等待该 Promise 完成

```html
<script setup>
  import { ref, onServerPrefetch, onMounted } from 'vue';

  const data = ref(null);

  onServerPrefetch(async () => {
    // 在服务器上预加载数据
    data.value = await fetchOnServer();
  });

  onMounted(() => {
    if(!data.value) {
      // 在客户端组件挂载时再进行数据加载
      data.value = await fetchOnClient();
    }
  });
</script>
```

### 依赖注入

#### provide() <em id="provide"></em> <!-- markdownlint-disable-line -->

> `provide()` 必须在组件的 `setup()` 阶段同步调用

允许组件向其所有后代组件注入一个依赖, 不论组件层次深度

```html
<script setup>
  import { ref, provide } from 'vue';

  provide('name', 'hello world');

  // 或者是返回一个对象的函数
  provide(() => {
    return { foo: 'foo' };
  });

  const count = ref(0);
  provide('count', count);
</script>
```

#### inject() <em id="inject"></em> <!-- markdownlint-disable-line -->

> `inject()` 必须在组件的 `setup()` 阶段同步调用

注入一个由祖先组件或整个应用(通过 `app.provide()` ) 提供的值

- 第一个参数为注入的 key, 通过遍历父组件链匹配 key 来确定最近的组件所提供的值, 否则将返回 undefined
- 第二个参数可选, 即在没有匹配到 key 时使用的默认值,

  - 如果为一个工厂函数, 则用来返回某些创建复杂的值
  - 如果默认值本身是一个函数, 则需要将 true 作为第三个参数传入, 表明这个函数就是默认值而不是工厂函数

```html
<script setup>
  import { inject } from 'vue';

  // 注入值的默认方式
  const count = inject('count');

  // 注入一个值, 如果为空则使用提供的默认值
  const foo = inject('foo', 'default value');

  // 注入一个值, 如果为空则使用提供的工厂函数
  const bar = inject('bar', () => new Map());

  // 注入一个值, 表明提供的默认值是一个函数
  const fn = inject('fn', () => {}, true);
</script>
```

## 选项式 API <em id="optionalapi"></em> <!-- markdownlint-disable-line -->

选项式 API 以 `组件实例` 的概念为中心(this), 将响应性相关的细节抽象出来, 并强制按照选项来组织代码, 从而对初学者而言更为友好

### 状态选项

#### data

> 以 \_ 和 $ 开头的属性不会被组件实例代理, 因为它们可能和 Vue 的内置属性, API 方法冲突

用于声明组件初始响应式状态的函数

#### props <em id="props"></em> <!-- markdownlint-disable-line -->

用于声明组件的 props

- 使用字符串数组的简易形式
- 使用对象的完整形式, 可以对单个 prop 进行更详细的配置

  - type 定义 prop 的类型, 可以为原生构造函数之一
  - default 为该 prop 指定一个当其没有被传入值或值为 undefined 时的默认值,
    对象或数组的默认值必须从一个工厂函数返回, 工厂函数也接收原始 prop 对象作为参数
  - required 定义该 prop 是否必需传入
  - validator 将 prop 值作为唯一参数传入的自定义验证函数

```javascript
export default {
  data() {
    return { name: 'hello world' };
  },
  // 简易形式
  props: ['name', 'age'],
  // 对象形式
  props: {
    name: String, // 类型检查
    age: {
      type: Number,
      default: 18,
      required: true,
      validator: (value) => {
        return value > 0;
      },
    },
    hobbies: {
      type: Array,
      default: (prop) => ['reading'],
    },
  },
};
```

#### computed

用于声明在组件实例上暴露的计算属性

- 包含一个只有 getter 函数的方法, 方法名为计算属性的名称
- 包含一个具有 get 和 set 函数的对象

```javascript
export default {
  data() {
    return { age: 18 };
  },
  computed: {
    // 只读计算属性
    name() {
      return 'hello world';
    },
    // 可读可写计算属性
    agePlus: {
      get() {
        return this.age;
      },
      set(val) {
        this.age = this.age + val;
      },
    },
  },
};
```

#### methods

> 在声明方法时避免使用箭头函数, 因为它们不能通过 this 访问组件实例

用于声明要混入到组件实例中的方法

#### [watch](#watch)

用于声明在数据更改时调用的侦听回调

- 普通形式
- 对象形式

```javascript
export default {
  data() {
    return { age: 18 };
  },
  watch: {
    // 侦听根级属性
    age(val, oldVal) {
      console.log(val, oldVal);
    },
    // 字符串方法名称
    b: 'otherMethod',
    // 深度侦听属性
    c: {
      handler(val, oldVal) {},
      deep: true,
      flush: 'post',
      onTrack(e) {},
    },
    // 侦听单个嵌套属性
    'c.d': function (val, oldVal) {},
    // 该回调函数在侦听开始之后立即调用
    e: {
      handler(val, oldVal) {},
      immediate: true,
    },
    // 回调数组, 将会被逐一调用
    f: [
      'handler',
      function handle2(val, oldVal) {
        console.log('handle2 triggered');
      },
      {
        handler: function handle3(val, oldVal) {
          console.log('handle3 triggered');
        },
        /* ... */
      },
    ],
  },
};
```

#### emits <em id="emits"></em> <!-- markdownlint-disable-line -->

用于声明由组件触发的自定义事件

- 使用字符串数组的简易形式
- 使用对象的完整形式

```javascript
export default {
  data() {
    return { name: 'hello world' };
  },
  // 简易形式
  emits: ['check'],
  // 对象形式
  emits: {
    // 没有验证函数
    click: null,
    // 具有验证函数
    submit: (payload) => {
      if (payload.email && payload.password) {
        return true;
      } else {
        console.warn(`Invalid submit event payload!`);
        return false;
      }
    },
  },
  mounted() {
    this.$emit('check');
  },
};
```

#### expose

> 保持私有的内部状态或方法, 以避免紧耦合

用于声明当组件实例被父组件通过模板引用访问时暴露的公共属性, 当使用 expose 时, 只有显式列出的属性将在组件实例上暴露

```javascript
export default {
  data() {
    return { name: 'hello world' };
  },
  expose: ['publicProp', 'publicMethod'],
};
```

### 渲染选项

#### template

> 通过 `template` 选项提供的模板将会在运行时即时编译

用于声明组件的字符串模板

- 如果字符串以 `#` 开头, 它将被用作 `querySelector` 的选择器, 并使用所选中元素的 `innerHTML` 作为模板字符串
- 如果存在 `render` 选项, 则 `template` 选项将被忽略
- 如果应用的根组件不含任何 `template` 或 `render` 选项, Vue 将尝试使用所挂载元素的 `innerHTML` 来作为模板

#### render

> `render` 具有比 `template` 更高的优先级

用于编程式地创建组件虚拟 DOM 树的函数

#### compilerOptions

用于配置组件模板的运行时编译选项, 支持与应用级 `app.config.compilerOptions` 相同的选项, 并针对当前组件有更高的优先级

```javascript
import { h } from 'vue';
export default {
  data() {
    return { name: 'hello world' };
  },
  template: '#tpl',
  render() {
    return h('div', {}, [h('p', 'This is tag p'), h('p', this.name)]);
  },
  compilerOptions: {
    delimiters: ['{{', '}}'],
    comments: true,
    isCustomElement(tag) {
      return tag.startsWith('icon-');
    },
  },
};
```

### 生命周期选项

> [组合式 API](#compositionapi) 中的 `setup()` 钩子函数会在所有 [选项式 API](#optionalapi) 钩子之前调用

#### beforeCreate

在组件实例初始化完成之后立即调用

#### created

在组件实例处理完所有与状态相关的选项后调用, 此时挂载阶段还未开始, `$el` 属性仍不可用

#### [beforeMount](#onBeforeMount)

> 钩子函数在服务器端渲染期间不会被调用

在组件被 **挂载之前** 调用, 组件已经完成了其响应式状态的设置, 但还没有创建 DOM 节点, 将首次执行 DOM 渲染过程

#### [mounted](#onMounted)

> 钩子函数在服务器端渲染期间不会被调用

在组件被 **挂载之后** 调用

#### [beforeUpdate](#onBeforeUpdate)

> 钩子函数在服务器端渲染期间不会被调用

在组件因为响应式状态变更而更新其 DOM 树之前调用

#### [updated](#onUpdated)

> 钩子函数在服务器端渲染期间不会被调用

在组件因为响应式状态变更而更新其 DOM 树之后调用

#### [beforeUnmount](#onBeforeUnmount)

> 钩子函数在服务器端渲染期间不会被调用

在组件实例被 **卸载之前** 调用

#### [unmounted](#onUnmounted)

> 钩子函数在服务器端渲染期间不会被调用

在组件实例被 **卸载之后** 调用

#### [activated](#onActivated)

> 钩子函数在服务器端渲染期间不会被调用

如果组件实例是 `<KeepAlive>` 缓存树的一部分, 当组件被插入到 DOM 中时调用

#### [deactivated](#onDeactivated)

> 钩子函数在服务器端渲染期间不会被调用

如果组件实例是 `<KeepAlive>` 缓存树的一部分, 当组件从 DOM 中移除时调用

#### [errorCaptured](#onErrorCaptured)

在捕获了后代组件传递的错误时调用

#### [renderTracked](#onRenderTracked)

> 仅在开发模式下可用, 且在服务器端渲染期间不会被调用

在一个响应式依赖被组件的渲染作用追踪后调用

#### [renderTriggered](#onRenderTriggered)

> 仅在开发模式下可用, 且在服务器端渲染期间不会被调用

在一个响应式依赖被组件触发了重新渲染之后调用

#### [serverPrefetch](#onServerPrefetch)

在组件实例在服务器上被渲染之前要完成的异步函数

### 组合选项

#### [provide](#provide)

用于提供可以被后代组件注入的值

#### [inject](#inject)

用于声明要通过从上层提供方匹配并注入当前组件的属性

- 一个字符串数组
- 一个对象
  - 匹配可注入的 key(String 或者 Symbol)
  - 一个对象
    - from 属性表示匹配可用的注入的来源
    - default 属性用作候补值, 和 props 的默认值类似

```javascript
export default {
  data() {
    return { name: 'hello world' };
  },
  // 字符串数组
  inject: ['foo'],
  // 对象形式指定默认值
  inject: {
    foo: { default: 'foo' },
  },
  // 对象形式指定来源
  inject: {
    foo: {
      from: 'bar',
      default: 'foo',
    },
  },
  // 对象形式使用工厂函数
  inject: {
    foo: {
      from: 'bar',
      default: () => [1, 2, 3],
    },
  },
};
```

#### mixins

> Mixin 钩子的调用顺序与提供它们的选项顺序相同, 且会在组件自身的钩子前调用

一个包含组件选项对象的数组, 这些选项都将被混入到当前组件的实例中(不推荐使用)

#### extends

> `extends` 和 `mixin` 实现上几乎相同, 但是表达的目标不同, `mixins` 选项基本用于组合功能, `extends` 一般更关注继承关系, 为[选项式 API](#optionalapi)设计的

要继承的 **基类** 组件, 同 `mixins` 一样, 所有选项都将使用相关的策略进行合并, 不会处理 setup() 钩子的合并

### 其他杂项

#### name

> 使用 name 选项可以覆盖推导出的名称, 或是在没有推导出名字是显式提供一个

用于显式声明组件展示时的名称

- 在组件自己的模板中递归引用自己时
- 在 Vue 开发者工具中的组件树显示时
- 在组件抛出的警告追踪栈信息中显示时

##### 场景 1

使用单文件组件时, 组件会根据其文件名推导出其名称, 例如 `MyComponent.vue` 的文件会推导出显式名称为 `MyComponent`

##### 场景 2

当使用 `app.component` 注册全局组件时, 这个全局 ID 会自动设置为其名称

#### inheritAttrs <em id="inheritAttrs"></em> <!-- markdownlint-disable-line -->

> 默认情况下, 父组件传递的没有被子组件解析为 `props` 的 `attributes` 绑定会被透传

用于控制是否启用默认的组件 `attribute` 透传行为, 默认为 true

- 使用 [\<script setup\>](#script-setup) 的[组合式 API](#compositionapi) 中声明这个选项时, 需要一个额外的 `<script>` 块
- Vue 3.3 支持, 使用 [defineOptions](#defineOptions) 声明

```html
<!-- 单独 script 块声明 inheritAttrs 选项 -->
<script>
  export default {
    inheritAttrs: false,
  };
</script>
<!-- SFC 组合式 API -->
<script setup>
  const msg = 'hello world';

  const props = defineProps(['name', 'age']);
  const emit = defineEmits(['input']);

  // 暴露外部可访问的公共属性
  defineExpose({ msg });

  // vue 3.3 使用宏函数定义其它杂项
  defineOptions({ name: 'Comp-A', inheritAttrs: false });
</script>
```

#### components

用于注册对当前组件实例可用的组件的配置对象

```javascript
import { h } from 'vue';
export default {
  data() {
    return { name: 'hello world' };
  },
  components: {
    Foo,
    'my-component': {
      setup() {
        return () => h('h1', 'Register local components...');
      },
    },
  },
};
```

#### [directives](#directive)

用于注册对当前组件实例可用的指令的配置对象

```javascript
export default {
  data() {
    return { name: 'hello world' };
  },
  directives: {
    // 在模板中启用 v-focus 指令
    focus: {
      /* 自定义指令的钩子函数配置 */
    },
  },
};
```

### 组件实例

除了 $data 下的嵌套属性外, 其它的属性都是只读的

#### $data

从 `data` 选项函数返回的对象, 会被组件赋为响应式, 组件实例将会代理其数据对象的属性访问

#### [$props](#props)

表示组件当前已解析的 `props` 对象

#### $el

该组件实例管理的 DOM 根节点, $el 直到组件 **挂载完成** 之前都是 undefined

#### $options

已解析的用于实例化当前组件的组件选项

- 全局 mixin
- 组件 `extends` 的基组件
- 组件级 mixin

#### $parent

当前组件可能存在的父组件实例, 如果当前组件是顶层组件, 则为 null

#### $root

当前组件树的根组件实例, 如果当前组件实例没有父组件, 则为本身

#### [$slots](#v-slot)

表示父组件传入 **插槽** 的对象

#### $refs

包含 DOM 元素和组件实例的对象, 通过 **模板引用** 注册

#### $attrs

包含了组件所有透传 `attributes` 的对象

#### [$watch()](#watch)

用于命令式地创建侦听器的 API

#### $emit()

在当前组件触发一个自定义事件, 任何额外的参数都将传递给事件监听器的回调函数

#### $forceUpdate()

强制当前组件重新渲染, 仅仅影响实例本身和插入插槽内容的子组件

#### [$nextTick()](#nextTick)

> 和全局的 `nextTick` 的区别是传递给 `this.$nextTick()` 的回调函数会带上绑定当前组件实例上下文的 `this`

绑定在实例上的 `nextTick()` 函数

```javascript
export default {
  data() {
    return { name: 'hello world' };
  },
  updated() {
    this.$nextTick(() => {
      console.log(this.name);
    });
  },
};
```

## 内置内容

### 内置指令

#### v-text

更新元素的文本内容

#### v-html

更新元素的 `innerHTML`

#### v-show

基于表达式值的真假来改变元素的可见性, 通过设置内联样式的 `display` CSS 属性来工作

#### v-if

基于表达式值的真假来条件性地渲染元素或者模板片段, 同时使用 `v-if` 和 `v-for` 时, 前者的优先级更高

#### v-else

表示 `v-if` 或 `v-if` / `v-else-if` 链式调用的块

#### v-for

基于原始数据多次渲染元素或模板块

#### v-on

给元素绑定事件监听器, 缩写 `@`

- 和原生 DOM 事件不一样, 组件触发的事件**没有冒泡机制**, 只能监听直接子组件触发的事件, 平级组件或嵌套组件间通信, 应使用一个外部事件总线或全局状态管理方案

##### 事件处理器 <em id="handler"></em> <!--markdownlint-disable-line-->

- 内联事件处理器: 事件被触发时执行的内联 javascript 语句(与 onClick 类似)

  - 需要处理原生 DOM 事件时, 手动传入一个特殊的 $event 变量, 或者使用内联箭头函数

    ```html
    <button @click="log('click event', $event)">click</button>

    <button @click="(event) => log('click event', event)">click</button>
    ```

- 方法事件处理器: 一个指向组件上定义的方法的属性名或是路径
  - 自动接收原生 DOM 事件并立即触发执行

```html
<!-- 内联事件处理器  -->
<button @click="count++">Add</button>
<p>{{count}}</p>

<!-- 方法事件处理器 -->
<button @click="greet">Greet</button>
<!-- 
const name = ref('Hello Greet');

function greet(evt){
  console.log(name.value);
}
-->
```

##### 事件修饰符

- .stop 调用 `event.stopPropagation()`
- .prevent 调用 `event.preventDefault()`
- .capture 在捕获模式添加事件监听器
- .self 只有事件从元素本身发出才触发处理函数
- .{keyAlias} 只有在某些按键下触发处理函数
- .once 最多触发一次处理函数
- .passive 通过 `{passive: true}` 附加一个 DOM 事件

##### 按键修饰符

- .enter
- .tab
- .delete 捕获 delete 和 backspace 两个按键
- .esc
- .space
- .up
- .down
- .left 只在鼠标左键事件触发处理函数
- .right 只在鼠标右键事件触发处理函数
- .middle 只在鼠标中键事件触发处理函数

##### 系统按键修饰符

- .ctrl
- .alt
- .shift
- .meta
- .exact 允许控制触发一个事件所需的确定组合的系统按键修饰符

```html
<!-- 当按下 ctrl 时, 即使同时按下 alt 或者 shift 也会触发 -->
<button @click.ctrl="onClick">click</button>

<!-- 仅当按下 ctrl 且未按任何其它键时才会触发 -->
<button @click.ctrl.exact="onClick">click</button>

<!-- 仅当没有按下任何系统按键时触发 -->
<button @click.exact="onClick">click</button>
```

#### v-bind <em id="v-bind"></em> <!-- markdownlint-disable-line -->

> `v-bind` 的绑定顺序会影响渲染结果

动态的绑定一个或多个 attribute, 也可以是组件的 prop, 缩写 `:` 或 `.`(当使用 `.prop` 修饰符)

##### 绑定修饰符

- .camel 将 `kebab-case` 命名的属性转变为 `camelCase` 命名
- .prop 强制绑定为 DOM property, 3.2 支持
- .attr 强制绑定为 DOM attribute, 3.2 支持

```html
<svg :view-box.camel="viewBox"></svg>
<!-- 使用 .prop 修饰符，会从组件选项 props 中移除, 
  以 .[attr] 形式出现在组件 attrs 参数中并且**不会显示**在 DOM 上 -->
<!-- 使用 .attr 修饰符，会从组件选项 props 中移除, 
  以 ^[attr] 形式出现在组件 attrs 参数中并且**会显示**在 DOM 上 -->
<div :someProperty.prop="someObject"></div>
<!-- 等价于 -->
<div .someProperty="someObject"></div>
```

##### 同名缩写

> Vue 3.4 支持

```html
<!-- 缩写形式的动态 attribute, 扩展为 :src="src" -->
<img :src />
```

#### v-model <em id="v-model"></em> <!-- markdownlint-disable-line -->

在表单输入元素或组件上创建双向绑定

- \<input\>
- \<select\>
- \<textarea\>
- components

##### 修饰符

- .lazy 监听 change 事件而不是 input 事件
- .number 将输入的合法字符换转为数字
- .trim 移除输入内容两端空格

##### 版本迭代

- [`v-bind`](#v-bind) 的 `.sync` 修饰符和组件的 model 选项被**移除**, 使用 `v-model` 和参数代替
- 同一组件上可以使用多个 `v-model` 进行双向绑定
- 可自定义 `v-model` 修饰符
- 自定义组件时 `v-model` 的 `prop` 和 `event` 默认名称已更改

  - prop: `value` -> `modelValue`
  - event: `input` -> `update:modelValue`

##### migration

- 所有子组件 `.sync` 修饰符的部分替换为 `v-model`
- 未带参数的 `v-model`, 修改子组件的 prop -> `modelValue`, event -> `update:modelValue`

```html
<template>
  <!-- 带参数 -->
  <my-component :title.sync="pageTitle" />
  <!-- 替换为 -->
  <my-component :title="pageTitle" @update:title="pageTile = $event" />
</template>
<script setup>
  const props = defineProps({ title: String });
  const emit = defineEmits(['update:title']);

  const changePageTitle = function (title) {
    emit('update:title', title);
  };
</script>

<template>
  <!-- 未带参数替换为 v-model -->
  <ChildComponent v-model="pageTitle" />
</template>
<script setup>
  // 以前是`value: String`
  const props = defineProps({ modelValue: String });
  const emit = defineEmits(['update:modelValue']);

  const changePageTitle = function (title) {
    emit('update:modelValue', title); // 以前是 `this.$emit('input', title)`
  };
</script>
```

##### Vue 2.0

- `v-model` 只能使用 `value` 作为 prop, 并监听子组件抛出的 `input` 事件, 如果使用其他 prop, 必须使用 `v-bind.sync` 同步

```html
<template>
  <my-component :value="pageTitle" @input="pageTitle = $event" />
  <!-- 简写方式 -->
  <my-component v-model="pageTitle" />
</template>
<script>
  export default {
    props: ['value'],
    created() {
      this.$emit('input', 'hello value');
    },
  };
</script>
```

##### Vue 2.2

- 增加组件选项 `model`, 允许自定义 `v-model` 的 prop 和 event, 只能在组件上使用一个 model

```html
<template>
  <my-component :value="pageTitle" @change="pageTitle = $event" />
  <!-- 简写方式 -->
  <my-component v-model="pageTitle" />
</template>
<script>
  export default {
    model: {
      prop: 'title',
      event: 'change',
    },
    props: {
      // 这将允许 `value` 属性用于其他用途
      value: String,
      // 使用 `title` 代替 `value` 作为 model 的 prop
      title: {
        type: String,
        default: 'Default title',
      },
    },
  };
</script>
```

##### Vue 2.3

- 增加 `.sync` 修饰符

```html
<template>
  <my-component :title="pageTitle" @update:title="pageTitle = $event" />
  <!-- 简写方式 -->
  <my-component :title.sync="pageTitle" />
</template>
```

##### Vue 3.x

- `v-model` 默认传递 `modelValue` prop, 并接收子组件抛出的 `update:modelValue` 事件

```html
<template>
  <!-- 单个 v-model 绑定 -->
  <my-component
    :modelValue="pageTitle"
    @update:modelValue="pageTitle = $event"
  />
  <!-- 简写方式 -->
  <my-component v-model="pageTitle" />
</template>
<script setup>
  const props = defineProps(['modelValue']);
  const emit = defineEmits(['update:modelValue']);

  // 触发事件
  emit('update:modelValue', 'hello modelValue');
</script>
```

- 多个 `v-model` 绑定

```html
<template>
  <!-- 多个 v-model 绑定 -->
  <my-component
    :title="pageTitle"
    @update:title="pageTitle = $event"
    :content="pageContent"
    @update:content="pageContent = $event"
  />
  <!-- 简写方式 -->
  <my-component v-model:title="pageTitle" v-model:content="pageContent" />
</template>
<script setup>
  const props = defineProps({ title: String, content: String });
  const emit = defineEmits(['update:title', 'update:content']);

  emit('update:title', 'hello title');
  emit('update:content', 'hello content');
</script>
```

- 处理 `v-model` 修饰符

  - 不带参数: 生成的 prop 名称为 `modelModifiers` 的对象, 包含传入的修饰符
  - 带参数: 生成的 prop 名称为 `arg + 'Modifiers'`

```javascript
import { createApp, ref, h, defineComponent } from 'vue';

// v-model 不带参数
// <with-out-args v-model.capitalize="myText" />
const WithoutArgs = defineComponent({
  props: ['modelValue', 'modelModifiers'],
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    console.log(props);
    const emitValue = function (e) {
      let value = e.target.value;
      if (props.modelModifiers.capitalize) {
        value = value.charAt(0).toUpperCase() + value.slice(1);
      }
      emit('update:modelValue', value);
    };
    return () =>
      h('p', [
        'without args ',
        h('input', {
          type: 'text',
          placeholder: 'write something...',
          value: props.modelValue,
          onInput: emitValue,
        }),
      ]);
  },
});

// v-model 带参数
// <with-args v-model:description.uppercase="myText" />
const WithArgs = defineComponent({
  props: ['description', 'descriptionModifiers'],
  emits: ['update:description'],
  setup(props, { emit }) {
    console.log(props);
    const emitValue = function (e) {
      let value = e.target.value;
      if (props.descriptionModifiers.uppercase) {
        value = value.toUpperCase();
      }
      emit('update:description', value);
    };

    return () =>
      h('p', [
        'with args ',
        h('input', {
          type: 'text',
          placeholder: 'write something...',
          value: props.description,
          onInput: emitValue,
        }),
      ]);
  },
});

const app = createApp({
  setup(props, ctx) {
    const v1 = ref('hello without args');
    const v2 = ref('hello with args');

    return () => [
      // 渲染 v-model 不带参数的组件
      h(WithoutArgs, {
        modelValue: v1.value,
        modelModifiers: { capitalize: true },
        'onUpdate:modelValue': (value) => (v1.value = value),
      }),
      // 渲染 v-model 带参数的组件
      h(WithArgs, {
        description: v2.value,
        descriptionModifiers: { uppercase: true },
        'onUpdate:description': (value) => (v2.value = value),
      }),
    ];
  },
});
app.mount('#app');
```

#### v-slot <em id="v-slot"></em> <!-- markdownlint-disable-line -->

> 如果混用了 **具名插槽** 和 **默认插槽**, 则需要为 **默认插槽** 使用显式的 `<template>` 标签, 直接为组件添加 `v-slot` 指令将导致编译错误

用于声明具名插槽或是期望接收 props 的作用域插槽, 缩写 `#`

##### 限制使用

能够合法在函数参数位置使用的 js 表达式, 支持解构语法. 绑定的值是可选的(只有在给作用域插槽传递 props 才需要)

- \<template\>
- components(用于带有 prop 的单个默认插槽)

##### 具名插槽

组件中包含多个插槽出口, `<slot>` 内置元素的特殊属性 `name` 用来给每个插槽分配唯一的 ID 以确定每一处要渲染的内容

```html
<!-- BaseComponent 组件模板 -->
<div class="container">
  <header>
    <slot name="header"></slot>
  </header>
  <main>
    <slot></slot>
  </main>
  <footer>
    <slot name="footer"></slot>
  </footer>
</div>

<!-- 父组件使用 -->
<BaseComponent>
  <template #header>
    <!-- header 插槽的内容 -->
  </template>

  <!-- 所有位于顶级的非 template 节点都被隐式地当作默认插槽的内容 -->
  <p>This is the first tag p</p>
  <p>This is the second tag p</p>

  <template #footer>
    <!-- footer 插槽的内容 -->
  </template>
</BaseComponent>
```

##### 动态插槽

```html
<!-- 动态插槽名，支持 `#SlotName` 缩写 -->
<base-layout>
  <template v-slot:[dynamicSlotName]></template>
  <!-- 等同于 -->
  <template #[dynamicSlotName]></template>
</base-layout>
```

##### 作用域插槽

- 默认插槽接收 props, 通过子组件标签上的 `v-slot` 指令接收一个插槽 props 对象

```html
<slot text="hello text" message="hello message"></slot>

<!-- 单个默认作用域插槽, 直接使用子组件标签 -->
<MyComponent v-slot="slotScope">
  {{slotScope.text}} - {{slotScope.message}}
</MyComponent>
```

- 具名插槽 props 可以作为 `v-slot` 指令的值被访问到 `v-slot:name="slotScope"`

```html
<!-- 具名作用域插槽 -->
<MyComponent>
  <template #header="headerScope"> {{ headerScope }} </template>

  <!-- 使用显式的默认插槽 -->
  <template #default="{ message }">
    <p>{{ message }}</p>
  </template>

  <template #footer>
    <p>Here's some contact info</p>
  </template>
</MyComponent>
```

#### v-pre

跳过该元素及其所有子元素的编译

#### v-once

仅渲染元素和组件一次, 并跳过之后的更新

#### v-memo

> Vue 3.2 支持

缓存一个模板的子树, 根据传入的依赖值数组的比较结果控制子树的更新

- 如果依赖值为空数组, 功能等同于 `v-once`
- 结合 `v-for` 使用, 必须确保和 `v-for` 用在同一个元素上, 否则无效

```html
<!-- 普通用法 -->
<div v-memo="[valueA, valueB]"></div>
<!-- v-for | v-memo -->
<div v-for="item in list" :key="item.id" v-memo="[item.id === selected]">
  <p>ID: {{ item.id }} - selected: {{ item.id === selected }}</p>
  <p>more child nodes</p>
</div>
```

#### v-cloak

> 该指令只在没有构建步骤的环境下需要使用

用于隐藏尚未完成编译的 DOM 模板

### 内置组件 <em id="builtincomponent"></em> <!-- markdownlint-disable-line -->

> 内置组件无需注册便可以直接在模板中使用，同时也支持 `tree-shake`; 仅在使用时才会包含在构建中
> 在 [**渲染函数**](#renderingfunc) 中使用内置组件时, 需要显式引入

```javascript
import { h, KeepAlive, Transition } from 'vue';

export default {
  setup(props, ctx) {
    return () => h(Transition, { mode: 'out-in' } /* ... */);
  },
};
```

#### \<Transition\>

为单个元素或组件提供动画过渡效果

##### \<Transition\> props

- name
- css
- type
- duration
- mode
- appear
- enterFromClass
- enterActiveClass
- enterToClass
- appearFromClass
- appearActiveClass
- appearToClass
- leaveFromClass
- leaveActiveClass
- leaveToClass

##### \<Transition\> 事件

- @before-enter
- @before-leave
- @enter
- @leave
- @appear
- @after-enter
- @after-leave
- @after-appear
- @enter-cancelled
- @leave-cancelled(v-show only)
- @appear-cancelled

```html
<!-- 单个元素 -->
<Transition>
  <div v-if="ok">toggled content</div>
</Transition>

<!-- 动态组件 -->
<Transition name="fade" mode="out-in" appear>
  <component :is="view"></component>
</Transition>

<!-- 事件钩子 -->
<div id="transition-demo">
  <Transition @after-enter="transitionComplete">
    <div v-show="ok">toggled content</div>
  </Transition>
</div>
```

#### \<TransitionGroup\>

为列表中的多个元素或组件提供过渡效果

##### \<TransitionGroup\> props

- tag 如果未定义, 则渲染为片段(fragment)
- moveClass 用于自定义过渡期间被应用的 CSS class, 使用 `kebab-case` 格式

##### \<TransitionGroup\> 事件

`<TransitionGroup>` 抛出与 `<Transition>` 相同的事件

#### \<KeepAlive\>

缓存包裹在其中的动态切换组件

##### \<KeepAlive\> props

- include 哪些组件实例可以被缓存
- exclude 哪些组件实例不被缓存
- max 最多可以缓存多少组件实例

```html
<!-- 逗号分隔字符串 -->
<KeepAlive include="a,b">
  <component :is="view"></component>
</KeepAlive>

<!-- regex (使用 `v-bind`) -->
<KeepAlive :include="/a|b/">
  <component :is="view"></component>
</KeepAlive>

<!-- Array (使用 `v-bind`) -->
<KeepAlive :include="['a', 'b']">
  <component :is="view"></component>
</KeepAlive>
```

#### \<Teleport\>

移动实际 DOM 节点(非销毁重建),并保持任何组件实例的活动状态

##### \<Teleport\> props

- to 必填项, 指定目标容器, 可以是选择器或实际元素
- disabled 值为 true 时, 内容将保留在其原始位置不做移动, 值可动态修改

```html
<!-- 正确 -->
<Teleport to="#some-id" />
<Teleport to=".some-class" />
<Teleport to="[data-teleport]" />
<!-- 错误 -->
<Teleport to="h1" />
<Teleport to="some-string" />

<button @click="open = true">Open Modal</button>
<Teleport to="body">
  <div v-if="open" class="modal">
    <p>Hello from the modal!</p>
    <button @click="open = false">Close</button>
  </div>
</Teleport>
```

#### \<Suspense\> <em id="suspense"></em> <!-- markdownlint-disable-line -->

用于协调对组件树中嵌套的异步依赖的处理

##### \<Suspense\> props

- timeout 渲染新内容耗时超时时间

##### \<Suspense\> 事件

- @pending 在 suspense 进入挂起状态时触发
- @resolve 在 default 插槽完成获取新内容时触发
- @fallback 在 fallback 插槽的内容显示时触发

##### \<Suspense\> 插槽

- #default
- #fallback

```html
<Suspense>
  <!-- 具有深层异步依赖的组件 -->
  <Dashboard />
  <!-- 在 #fallback 插槽中显示 “正在加载中” -->
  <template #fallback> Loading... </template>
</Suspense>
```

### 内置特殊元素

> `<component>`, `<slot>`, `<template>` 具有类似组件的特性, 也是模板语法的一部分. 但它们并非真正的组件
> 同时在模板编译期间会被编译掉. 因此, 它们通常在模板中使用小写字母

#### \<component\>

用于渲染动态组件或元素的 `元组件`

##### \<component\> props

- is 要渲染的实际组件由 `is` prop 决定
  - 如果是字符串时, 可以是 HTML 标签名或者组件的注册名
  - 或者是直接绑定到组件的定义

#### \<slot\>

表示模板中的插槽内容出口

##### \<slot\> props

- name 指定插槽名, 缺少时将会渲染默认插槽

#### \<template\>

当使用内置指令而不在 DOM 中渲染元素时, `<template>` 标签可以作为占位符使用

### 内置特殊 Attributes

#### key

主要作为 Vue 的虚拟 DOM 算法提示, 在比较新旧节点列表时用于识别 vnode

#### ref

用于注册元素或子组件的 `模板引用`

```html
<template>
  <div ref="root">This is a root element</div>
  <div v-for="item in list" :ref="itemRefs">{{ item }}</div>
</template>
<script setup>
  import { ref, onBeforeUpdate, onUpdated, onMounted } from 'vue';

  const root = ref(null);
  const itemRefs = ref([]);
  // 确保在每次更新之前重置 ref
  onBeforeUpdate(() => {
    itemRefs.value = [];
  });

  onUpdated(() => {
    console.log(itemRefs.value);
  });

  onMounted(() => {
    // DOM元素将在初始渲染后分配给ref
    console.log(root.value); // <div>这是根元素</div>
  });
</script>
```

#### is

- 用于动态绑定组件

```html
<script setup>
  import Foo from './Foo.vue';
  import Bar from './Bar.vue';
</script>
<template>
  <component :is="Foo" />
  <!-- 三目运算中的组件使用 -->
  <component :is="someCondition ? Foo : Bar" />
</template>
```

- 用于原生元素时, 将被作为 `Customized built-in element`, 如果需要用 Vue 组件替换原生元素, 需要加上 `vue:` 前缀

```html
<template>
  <table>
    <tr is="vue:my-row-component"></tr>
  </table>
</template>
```

## 单文件组件

### SFC 语法定义

- `<template>` 每个 `*.vue` 文件最多可以包含一个顶层 `<template>` 块, 包含的内容将被提取传递给 `@vue/compiler-dom` 编译生成为 [**渲染函数**](#renderingfunc)
- `<script>` 每个 `*.vue` 文件最多可以包含一个 `<script>` 块(使用 `<script setup>` 除外), 默认导出是 Vue 的组件选项对象
- `<script setup>` 每个 `*.vue` 文件最多可以包含一个 `<script setup>` 块, 此脚本块将被预处理为组件的 `setup()` 函数
- `<style>` 每个 `*.vue` 文件可以包含多个 `<style>` 块

#### src 导入

可以将单文件组件拆分成多个文件中, 使用 src 导入外部文件

```html
<template src="./template.html"></template>
<script src="./script.js"></script>
<style src="./style.css"></style>
```

### \<script setup\>

<em id="script-setup"></em> <!--markdownlint-disable-line-->

> `<script setup>` 是在单文件组件(SFC) 中使用 [组合式 API](#compositionapi) 的编译时语法糖
> `<script setup>` 中的代码会在每次组件实例被创建的时候执行

- 更少的样板内容, 更简洁的代码
- 能够使用纯 TypeScript 声明 props 和 自定义事件
- 更好的运行时性能(其模板会被编译成同一作用域内的渲染函数, 避免了渲染上下文代理对象)
- 更好的 IDE 类型推导性能(减少了语言服务器从代码中抽取类型的工作)

#### 顶层绑定

任何在 `<script setup>` 声明的 **顶层的绑定**(包括变量, 函数声明, 以及 import 导入的内容)都能在模板中直接使用

#### 响应式

```html
<script setup>
  import { ref, reactive } from 'vue';
  const count = ref(0);
  const user = reactive({ name: 'hello world', age: 18 });
</script>
<template>
  <button @click="count++">{{ count }}</button>
</template>
```

#### 使用组件

单文件组件模板中使用组件可以使用 `kebab-case` 或者 `PascalCase` 两种格式, 推荐使用后者

```html
<script setup>
  import { MyComponent } from './MyComponent.vue';
</script>
<template>
  <MyComponent />
</template>
```

##### 动态组件

```html
<script setup>
  import Foo from './Foo.vue';
  import Bar from './Bar.vue';
</script>
<template>
  <component :is="Foo" />
  <!-- 三目运算中的组件使用 -->
  <component :is="someCondition ? Foo : Bar" />
</template>
```

##### 递归组件

一个单文件组件可以通过它的文件名被其自己所引用, 为防止具名的导入和组件自身推导的名字冲突可以使用别名的方式

```javascript
import { FooBar as FooBarChild } from './components';
```

##### 命名空间组件

可以使用带 `.` 的组件标签来引用嵌套在对象属性中的组件

```html
<script setup>
  import * as Form from './form-components';
</script>
<template>
  <Form.Input>
    <Form.Label>label</Form.Label>
  </Form.Input>
</template>
```

#### 使用[自定义指令](#directive)

本地声明自定义指令在 `<script setup>` 中不需要显式注册, 但必须遵循 `vNameOfDirective` 的命名规范

```html
<script setup>
  const vMyDirective = {
    mounted(el, binding, vnode prevVnode){},
    updated(el, binding, vnode prevVnode){},
    unmounted(el, binding, vnode prevVnode){}
  }
</script>
<template>
  <h1 v-my-directive>This is a heading.</h1>
</template>
```

#### defineProps|defineEmits

- `defineProps` 和 `defineEmits` 都是只能在 `<script setup>` 中使用的 **编译器宏**, 不需要导入直接使用, 且会随着 `<script setup>` 的处理过程一同被编译掉
- `defineProps` 接收和 [props](#props) 选项相同的值, `defineEmits` 接收和 [emits](#emits) 选项相同的值
- `defineProps` 和 `defineEmits` 在选项传入后会提供恰当的类型推导
- 传入 `defineProps` 和 `defineEmits` 的选项会从 setup 中提升到模块的作用域, 因此, 传入的选项不能引用在 setup 作用域中声明的局部变量

##### 使用类型声明时的默认 props 值

defineProps 不能给使用类型声明的 props 提供默认值, 使用 `withDefaults` **编译器宏** 解决

```javascript
import { withDefaults } from 'vue';

export interface Props {
  msg?: string
  labels?: string[]
}

const props = withDefaults(defineProps<Props>(), {
  msg: 'hello withDefaults',
  labels: () => ['one', 'two']
});
```

##### props|emit 类型声明

`props` 和 `emits` 都可以通过给 `defineProps` 和 `defineEmits` 传递纯类型参数的方式声明

```html
<script setup>
  const props = defineProps({ title: string, age: number });
  const emit = defineEmits({
    // 没有验证函数
    click: null,
    // 有验证函数
    submit: (payload) => {
      if (payload.email && payload.password) {
        return true;
      } else {
        console.warn(`Invalid submit event payload!`);
        return false;
      }
    },
  });
</script>
```

#### defineExpose()

使用 `<script setup>` 的组件是 **默认关闭** 的, 不会暴露任何在 `<script setup>` 中声明的绑定. 使用 defineExpose 显式指定组件中要暴露出去的属性

```html
<script setup>
  import { ref } from 'vue';

  const a = 1;
  const b = ref(0);
  defineExpose({ a, b });
</script>
```

#### defineOptions() <em id="defineOptions"></em> <!--markdownlint-disable-line-->

> Vue 3.3 支持

在 `<script setup>` 中使用 [选项式 API](#optionalapi) 的**宏**, 无法访问 `<script setup>` 中不是字面常数的局部变量

```html
<script setup>
  import { useSlots } from 'vue';

  defineOptions({
    name: 'Foo',
    inheritAttrs: false,
  });

  const slots = useSlots();
</script>

<!-- Compiled Code -->

<script>
  export default {
    name: 'Foo',
    inheritAttrs: false,
  };
</script>
<script setup>
  import { useSlots } from 'vue';

  const slots = useSlots();
</script>
```

#### defineSlots()

> Vue 3.3 支持

只接受类型参数, 没有运行时参数. 用于为 IDE 提供插槽名称 和 props 类型检查的类型提示

```html
<script setup lang="ts">
  const slots = defineSlots<{
    default(props: { msg: string }): any;
  }>();
</script>
```

#### definePropsRefs()

> Vue 3.3 支持

在 `<script setup>` 中定义一个不会丢失响应性的解构 props 的响应式对象

```html
<script setup>
  // won't lose reactivity with destructuring
  const {foo, bar} = definePropsRefs<{
    foo: string,
    bar: number
  }>();

  // Ref<string>
  console.log(foo.value, bar.value);
</script>
```

#### defineRender()

> Vue 3.3 支持 <em id="vue3.3"></em> <!-- markdownlint-disable-line -->

在 `<script setup>` 中定义一个渲染函数

```html
<script setup>
  // JSX passed directly
  defineRender(
    <div>
      <span>hello world</span>
    </div>
  );

  // Or using render function
  defineRender(() => {
    return (
      <div>
        <span>hello world</span>
      </div>
    );
  });
</script>
```

#### defineModel()

> Vue 3.4 支持

```html
<script setup>
  // 声明 modelValue prop, 由父组件通过 v-model 使用
  const model = defineModel();
  // 或者声明带选项的 modelValue prop
  const model = defineModal({ type: String });
  // 在被修改时, 触发 update:modelValue 事件
  model.value = 'hello world';

  // 声明 count prop, 由父组件通过 v-model:count 使用
  const count = defineModel('count');
  // 或者声明带选项的 count prop
  const count = defineModel('count', { type: Number, default: 0 });
  // 在被修改时, 触发 update:count 事件
  count.value++;
</script>
```

##### 修饰符和转换器

```html
<script setup>
  const [modelValue, modelModifiers] = defineModel();
  // 对应 v-model.trim
  if (modelModifiers.trim) {
    // ...
  }

  // 通过修饰符使用 set 和 get 转换器对其值进行转换
  const [modelValue, modelModifiers] = defineModel({
    // get() 省略了，因为这里不需要它
    set(value) {
      // 如果使用了 .trim 修饰符，则返回裁剪过后的值
      if (modelModifiers.trim) {
        return value.trim();
      }
      // 否则，原样返回
      return value;
    },
  });
</script>
```

#### useSlots()|useAttrs()

- 在 SFC 中使用的辅助函数获取 slots 和 attrs
- 需要手动从 vue 中导入
- 返回的值和 `setupContext.slots` 和 `setupContext.attrs` 是等价的

```html
<script setup>
  import { useSlots, useAttrs } from 'vue';

  const slots = useSlots();
  const attrs = useAttrs();
</script>
```

#### 与普通 script 一起用

- 声明无法在 `<script setup>` 中声明的选项, 例如 [inheritAttrs](#inheritAttrs) 或插件的自定义选项(Vue 3.3 使用 [defineOptions](#defineOptions) 替代)
- 声明模块的具名导出(named exports)
- 运行只需要在模块作用域执行一次的副作用, 或是创建单例对象

#### 顶层 await

> `<script setup>` 中可以使用顶层 await, 结果代码会被编译成 `async setup()` > `async setup()` 必须与 [\<Suspense\>](#suspense) [**内置组件**](#builtincomponent)组合使用

```html
<script setup>
  const post = await fetch('/api/post/1').then(res => res.json())
</script>
```

#### 限制

由于模块执行语义的差异, `<script setup>` 中的代码依赖单文件组件的上下文, 当将其移动到外部的 `.js` 或 `.ts` 文件中时, 对于开发者和工具来说都会感到混乱. 因此, `<script setup>` 不能和 `src` 属性一起使用

### CSS 功能

#### 组件作用域

> 使用 `scoped` 后, 父组件的样式不会透传到子组件中，不过, 子组件的 **根节点** 会同时被父组件的作用域样式和子组件的作用域样式影响, 这样设计是为了让父组件可以从布局的角度出发, 调整其子组件根元素的样式

##### 深度选择器

```html
<style scoped>
  .a :deep(.b) {
    /* */
  }
</style>
```

##### 插槽选择器

默认情况下, 作用域样式不会影响到 `<slot />` 渲染出来的内容, 使用 `:slotted` 伪类以明确地将插槽内容作为选择器的目标

```html
<style scoped>
  :slotted(div) {
    color: red;
  }
</style>
```

##### 全局选择器

```html
<style scoped>
  :global(.red) {
    color: red;
  }
</style>
```

#### css Modules

一个 `<style module>` 标签会被编译成 `CSS Modules` 并且将生成的 CSS class 作为 `$style` 对象暴露给组件

```html
<template>
  <p :class="$style.red">This is should be red.</p>
</template>
<style module>
  .red {
    color: red;
  }
</style>
```

##### 自定义注入名称

module 属性可以接受一个值作为自定义注入名称代替 `$style`

```html
<template>
  <p :class="classes.red">This is should be red.</p>
</template>
<style module="classes">
  .red {
    color: red;
  }
</style>
```

##### 与[组合式 API](#compositionapi)一起使用

- 使用 `useCssModule` API 在 `setup()` 和 `<script setup>` 中访问注入的 class
- 使用 **自定义注入名称** 的 `<style module>`, `useCssModule` 接收一个匹配的 `module` attribute 值作为第一个参数

```html
<script setup>
  import { h, useCssModule } from 'vue';
  // 默认情况下, 返回 <style module> 的 class
  const style = useCssModule();
  // 自定义注入名称, 返回 <style module="classes"> 的 class
  const classes = useCssModule('classes');
</script>
```

#### CSS 中的 v-bind()

单文件组件的 `<style>` 标签支持使用 `v-bind` CSS 函数将 CSS 的值链接到动态的组件状态

- [选项式 API](#optionalapi)使用

```html
<template>
  <div class="text">hello world</div>
</template>
<script>
  export default {
    data() {
      return { color: 'red' };
    },
  };
</script>
<style scoped>
  .text {
    color: v-bind(color);
  }
</style>
```

- [组合式 API](#compositionapi) 使用

```html
<template>
  <p>hello world</p>
</template>
<script setup>
  const theme = {
    color: 'red',
  };
</script>
<style scope>
  p {
    color: v-bind('theme.color');
  }
</style>
```

## 进阶 API

### 渲染函数 <em id="renderingfunc"></em> <!-- markdownlint-disable-line-->

- 如果组件定义了 setup 并且返回值是一个函数, 则其返回值作为该组件的渲染函数
- 如果组件定义了 render, 则将其作为渲染函数
- 如果组件定义了 template, 则将其作为模板进行编译成可执行的渲染函数
- 如果以上条件都不满足, 则使用容器的 innerHTML 作为模板

#### h()

> 当创建组件的 vnode 时, 子节点必须以 **[插槽](#v-slot)函数** 的形式传递, 如果组件只有默认插槽, 可以使用单个 **[插槽](#v-slot)函数** 传递, 否则, 必须以 **[插槽](#v-slot)函数** 的对象形式传递

创建虚拟 DOM 节点(vnode)

- 第一个参数是一个字符串(用于原生元素)或者一个 Vue 组件定义
- 第二个参数是要传入的 prop, 如果 **[插槽](#v-slot)函数** 不是对象形式时, 可以省略此参数
- 第三个参数是子节点

```html
<script setup>
  import { h } from 'vue';

  h(
    'div',
    {
      class: 'bar',
      style: { color: 'red' },
      innerHtml: 'hello',
      // 事件监听以 onXxx 的形式
      onClick: () => {},
    },
    ['hello world', h('span', 'gg')]
  );
</script>
```

- 创建组件

```html
<script setup>
  import { h } from 'vue';
  import Foo from './Foo.vue';

  // 传递 prop
  h(Foo, {
    // 等价于 some-prop="hello world"
    someProp: 'hello world',
    // 等价于 @update="() => {}"
    onUpdate: () => {},
  });

  // 传递单个默认插槽函数, 可以省略 prop 参数
  h(Foo, () => 'default slot');

  // 传递具名插槽函数
  // 需要使用 null 来避免插槽对象被当作 prop
  h(MyComponent, null, {
    default: () => 'default slot',
    header: () => h('div', 'hello div'),
    footer: () => [h('span', 'one'), h('span', 'two')],
  });
</script>
```

#### mergeProps()

> `class`, `style` 将被合并成一个对象, `onXxx` 将被合并成一个数组

合并多个 props 对象, 用于处理含有特定的 props 参数的情况

```html
<script setup>
  import { h, mergeProps } from 'vue';

  const one = {
    class: 'foo',
    onClick: handlerA,
  };

  const two = {
    class: { bar: true },
    onClick: handlerB,
  };

  const merged = mergeProps(one, two);
  /**
   * {
   *  class: 'foo bar',
   *  onClick: [handlerA, handlerB]
   * }
   */
</script>
```

#### cloneVNode()

克隆一个 vnode, 可在原有的基础上添加一些额外的 prop

```html
<script setup>
  import { h, cloneVNode } from 'vue';

  const original = h('div');
  const cloned = cloneVNode(original, { id: 'foo' });
</script>
```

#### isVNode()

判断一个值是否为 vnode 类型

#### resolveComponent()

> `resolveComponent()` 只能在 [**渲染函数**](#renderingfunc) 或 `setup()` 中使用
> 如果可以直接引入组件就不需要使用此方法

按名称手动解析已注册的组件, 未找到则抛出一个运行时警告并返回组件名字符串

```html
<script setup>
  import { h, resolveComponent } from 'vue';

  const ButtonComponent = resolveComponent('ButtonComponent');
  return () => h(ButtonComponent);
</script>
```

#### [resolveDirective()](#directive)

> `resolveDirective()` 只能在 [**渲染函数**](#renderingfunc) 或 `setup()` 中使用
> 如果可以直接引入组件就不需要使用此方法

按名称手动解析已注册的指令, 未找到则抛出一个运行时警告并返回 undefined

```html
<script setup>
  import { resolveDirective } from 'vue';

  const myDirective = resolveDirective('myDirective');
</script>
```

#### [withDirectives()](#directive)

用于给 vnode 增加自定义指令

- 第一个参数为要添加指令的 vnode
- 第二个参数为自定义指令数组, 每个自定义指令表示为 `[Directive, value, argument, modifiers]` 形式的数组
  - [directive] 指令本身
  - [directive, value] 上述内容, 指令的值
  - [directive, value, arg] 上述内容, 一个 String 参数,eg: v-on:click 中的 click
  - [directive, value, arg, modifiers] 上述内容, 定义任意修饰符的 key:value 键值对

```html
<script setup>
  import { h, withDirectives } from 'vue';

  const pin = {
    created() {
      console.log('pin directive created...');
    },
    beforeMount() {
      console.log('pin directive beforeMount...');
    },
    mounted(el, binding, vnode, prevVnode) {
      console.log(binding.value, binding.arg, binding.modifiers);
      el.style.fontSize = '20px';
      el.style.color = '#08f';
    },
  };

  return () =>
    // <div v-pin:top.animate="200"></div>
    withDirectives(h('div', 'hello withDirectives'), [
      [pin, 200, 'top', { animate: true }],
    ]);
</script>
```

```javascript
import { withDirectives, resolveDirective } from 'vue';
const foo = resolveDirective('foo');
const bar = resolveDirective('bar');

return withDirectives(h('div'), [
  [foo, this.x],
  [bar, this.y],
]);
```

#### withModifiers()

用于向事件处理函数添加内置 `v-on` 修饰符

```html
<script setup>
  import { h, withModifiers } from 'vue';

  const clk = function (e) {
    console.log(e);
  };

  return () =>
    h(
      'button',
      {
        // 等价于 v-on.stop.prevent
        onClick: withModifiers(
          (e) => {
            console.log(e);
          },
          ['stop', 'prevent']
        ),
      },
      'Click Me'
    );
</script>
```

#### 综合使用

```javascript
import {
  createApp,
  h,
  ref,
  defineComponent,
  resolveComponent,
  resolveDirective,
  withDirectives,
  withModifiers,
} from 'vue';

const HelloWorld = defineComponent({
  props: ['name', 'age'],
  // 定义局部指令
  directives: {
    foo: {
      created(el, binding, vnode, prevVnode) {
        console.log('directives foo hooks created trigger... ', binding);
      },
      /* beforeMount, mounted, beforeUpdate, updated, beforeUnmount */
      unmounted(el, binding, vnode, prevVnode) {
        console.log('directives foo hooks unmounted trigger... ', binding);
      },
    },
  },
  setup(props, { slots }) {
    const message = ref('from hello world component');
    const fooV = ref(250);
    const show = ref(true);

    // 定义点击切换自定义指令值的方法
    const changeFooDirective = function () {
      fooV.value = Math.ceil(Math.random() * 100);
      show.value = fooV.value % 4 === 0 ? true : false;
    };

    // 解析一个已注册的指令, 未找到则抛出运行时警告并返回 undefined
    const foo = resolveDirective('foo');

    // Must use `.value` to read or write the value wrapped by `ref()`
    return () =>
      h('div', [
        h(
          'p',
          slots?.default?.({
            message: message.value,
            age: props.age > 0 ? props.age : 18,
          })
        ),
        h('p', slots?.header?.()),
        show.value
          ? withDirectives(
              h(
                'p',
                withDirectives(
                  h('span', ['v-foo:bar2.uppercase=' + fooV.value]),
                  [[foo, fooV.value, 'bar2', { uppercase: true }]]
                )
              ),
              [[foo, fooV.value, 'bar1', { uppercase: true }]]
            )
          : '',
        h(
          'p',
          h(
            'button',
            { onClick: changeFooDirective },
            'Click Me - ' + props.age
          )
        ),
      ]);
  },
});

const app = createApp({
  setup(props, { slots }) {
    const gg = {
      created(el, binding, vnode, prevVnode) {
        console.log('withDirectives gg hooks created trigger... ', binding);
      },
      /* beforeMount, mounted, beforeUpdate, updated, beforeUnmount */
      unmounted(el, binding, vnode, prevVnode) {
        console.log('withDirectives gg hooks unmounted trigger... ', binding);
      },
    };

    // 解析一个已注册的组件，未找到则抛出运行时警告并返回组件名字符串
    const HelloWorld = resolveComponent('hello-world');

    return () => [
      withDirectives(
        h(
          HelloWorld,
          { name: 'from createApp', age: -1 },
          {
            default: (slotScope) =>
              slotScope.message +
              ' - ' +
              slotScope.age +
              ' - others from createApp',
            header: () => 'from createApp by header slot...',
          }
        ),
        [[gg, 1000, 'bottom', { animate: true }]]
      ),
      h(
        'button',
        { onClick: withModifiers((e) => console.log(e), ['stop', 'prevent']) },
        'Click Me withModifiers'
      ),
    ];
  },
});
app.component('hello-world', HelloWorld);
app.mount('#app');
```

##### [自定义指令](#directive)挂载

![vue3-directives-hooks-1](/images/vue3-directives-hooks-1.jpg)

##### [自定义指令](#directive)更新

![vue3-directives-hooks-2](/images/vue3-directives-hooks-2.jpg)

##### [自定义指令](#directive)卸载

![vue3-directives-hooks-3](/images/vue3-directives-hooks-3.jpg)

### 服务端渲染

#### renderToString()

> 导出自 `vue/server-renderer`

- 传入第二个可选的上下文对象用来在渲染过程中记录额外的数据

```javascript
import { createSSRApp } from 'vue';
import { renderToString } from 'vue/server-renderer';

const app = createSSRApp({
  data() {
    return { name: 'hello world' };
  },
  template: `<div>{{name}}</div>`,
});

const ctx = {};
(async () => {
  const html = await renderToString(app, ctx);
  console.log(html);
})();
```

#### renderToNodeStream()

> 导出自 `vue/server-renderer`

将输入渲染为一个 `Node.js Readable Stream` 实例

```javascript
import { renderToNodeStream } from 'vue/server-renderer';

// 在 Node.js http 处理函数中使用
const stream = renderToNodeStream(app);
stream.pip(res);
```

#### pipeToNodeWritable()

> 导出自 `vue/server-renderer`

将输入渲染并 pipe 到一个 `Node.js Writable Stream` 实例

```javascript
import { pipeToNodeWritable } from 'vue/server-renderer';

// 在 Node.js http 处理函数中使用
pipeToNodeWritable(app, {}, res);
```

#### renderToWebStream()

> 导出自 `vue/server-renderer`

将输入渲染为一个 `Web ReadableStream` 实例

```javascript
import { renderToWebStream } from 'vue/server-renderer';

(async () => {
  const res = new Response(renderToWebStream(app));
  console.log(res); // Response
  // text() 返回 UTF-8 编码的包含 USVString 对象的 Promise 对象
  console.log(await res.text());
})();
```

#### pipeToWebWritable()

> 导出自 `vue/server-renderer`

将输入渲染并 pipe 到一个 `Web WritableStream` 实例

```javascript
import { pipeToWebWritable } from 'vue/server-renderer';

(async () => {
  // 创建一个转换流
  const tfs = new TransformStream();
  // 将内容渲染并 pipe 到转换流的可写流
  pipeToWebWritable(app, {}, tfs.writable);
  // 根据转换流的可读流创建响应对象
  const res = new Response(tfs.readable);
  console.log('pipeToWebWritable ', res);
  // text() 返回 UTF-8 编码的包含 USVString 对象的 Promise 对象
  console.log('pipeToWebWritable ', await res.text());
})();
```

#### renderToSimpleStream()

> 导出自 `vue/server-renderer`

通过一个简单的接口, 将输入以 `stream` 模式进行渲染

```javascript
import { renderToSimpleStream } from 'vue/server-renderer';

let res = '';

renderToSimpleStream(
  app,
  {},
  {
    push(chunk) {
      if (chunk == null) {
        console.log('renderToSimpleStream render complete: ', res);
      } else {
        res += chunk;
      }
    },
    destroy(err) {
      console.log(err);
    },
  }
);
```

#### useSSRContext()

运行时 API, 用于获取传递给 `renderToString` 或者其他服务端渲染 API 的上下文对象

```html
<script setup>
  import { useSSRContext } from 'vue';

  // 确保只在服务端渲染时调用
  if(import.meta.env.SSR){
    const ctx = useSSRContext();
    // 为上下文对象添加属性
  }
</script>
```

### 工具类型

#### PropType\<T\>

用于在用运行时 props 声明时给一个 prop 标注更复杂的类型定义

```javascript
  import type { PropType } from 'vue';

  interface Book{
    title: string,
    author: string,
    year: number
  }

  export default {
    props: {
      book: {
        // 提供一个比 `Object` 更具体的类型
        type: Object as PropType<Book>,
        required: true,
      }
    }
  }
```

#### ComponentCustomProperties

用于增强组件实例类型以支持自定义全局属性

```javascript
  import axios from 'axios';

  declare module 'vue'{
    interface ComponentCustomProperties {
      $http: typeof axios
      $translate: (key: string) => string
    }
  }
```

#### ComponentCustomOptions

用来扩展组件选项类型以支持自定义选项

```javascript
import {Route} from 'vue-router';

declare module 'vue' {
  interface ComponentCustomOptions {
    beforeRouteEnter?(to: any, from: any, next: () => void): void
  }
}
```

#### ComponentCustomProps

用于扩展全局可用的 TSX props, 以便在 TSX 元素上使用没有在组件选项上定义过的 props

```javascript
declare module 'vue'{
  interface ComponentCustomProps {
    hello?: string
  }
}
export {}
```

#### CSSProperties

用于扩展在样式属性绑定上允许的值的类型

```javascript
declare module 'vue' {
  interface CSSProperties {
    [key: `--${string}`]: string
  }
}
```

### 自定义渲染

#### createRenderer()

创建一个自定义渲染器, 可以在非 DOM 环境中使用 Vue 核心运行时的特性

```javascript
import { createRenderer } from '@vue/runtime-core';

const { render, createApp } = createRenderer({
  patchProp,
  insert,
  remove,
  createElement,
  // ...
});

// render 是底层 API
// createApp 返回一个应用实例
export { render, createApp };

// 重新导出 Vue 的核心 API
export * from '@vue/runtime-core';
```

## Router

- useLink
- useRoute 返回当前的路由地址, 相当于模板中使用 $route
- useRouter 返回路由器实例, 相当于模板中使用 $router
- 路由组件传参, 使用 props 将路由和组件解耦

  - 布尔模式: `{path: '/users/:id', component: User, props: true}`
  - 命名视图: `{path: '/users/:id', components: {default: User, sidebar: SideBar}, props: {default: true, sidebar: true}}`
  - 对象模式: `{path: '/users/profile', component: User, props: {newsLetterPopup: false}}`
  - 函数模式: `{path: '/search', component: Search, props: (route) => ({query: route.query})}`

```javascript
// 组件 User 和 路由强耦合
const User = { template: '<div>User {{$route.params.id}}</div>' };
const routes = [{ path: '/users/:id', component: User }];

// 将 props 设置 true 时, route.params 将被设置为组件的 props
const User = { props: ['id'], template: '<div>User {{id}}</div>' };
const routes = [{ path: '/users/:id', component: User, props: true }];
```

- beforeEnter 路由独享的守卫, **只在进入路由时触发**, 不会在 params, query, hash 改变时触发

```javascript
const routes = [
  {
    path: '/users/:id',
    component: User,
    beforeEnter: (to, from) => {
      // 从 /users/2 到 /users/3，/users/2#info 到 /users/2#projects 不会触发
      // reject navigation
      return false;
    },
  },
];
```

### 编程式导航

- 当同时提供了 path, params 参数时, params 会被忽略

```javascript
import { useRouter } from 'vue-router';
const router = useRouter();

// params 不能和 path 同时使用
router.push({ path: '/user', params: { username: 'zhangsan' } });

// 可以使用 name 和 params 组合
router.push({ name: 'user', params: { username: 'zhangsan' } });
```

- 替换当前位置

```javascript
router.push({ path: '/home', replace: true });
// 等价于
router.replace({ path: '/home' });
```

## Store

store 是一个用 [reactive](#reactive) 包装的对象, 不需要使用 .value 访问, 使用解构的方式将会丢失响应性

- defineStore 创建 store, 可使用 Option 对象 或 Setup 函数
- storeToRefs() 从 store 中提取属性时保持其响应性, 并且跳过所有的 action 或非响应式(不是 ref 或 reactive)的属性

```javascript
import { createApp, ref } from 'vue';
import { createPinia, defineStore, storeToRefs } from 'pinia';
// 创建 pinia 并挂载到 vue 实例
const pinia = createPinia();
const app = createApp({});
app.use(pinia);
app.mount('#app');
// 定义 store
const useCounterStore = defineStore('counter', () => {
  const count = ref(1);
  const doubleCount = count.value * 2;
  const increment = function () {
    count.value++;
  };
  return { count };
});

const counter = useCounterStore();
// count, doubleCount 是响应式的 ref
// 同时通过插件添加的属性也会被提取为 ref
// 并且会跳过所有的 action 或非响应式(不是 ref 或 reactive) 的属性
const { count, doubleCount } = storeToRefs(counter);
// 作为 action 的 increment 可以直接解构
const { increment } = counter;
```

- store.$dispose() 停止 store 的相关作用域, 并从 store 注册表中删除它. 插件可以覆盖此方法来清理已添加的任何副作用函数
- store.$reset() 重置 state 为初始值
  - 使用 选项式 API 创建的 store 调用此方法, 使用 setup 创建的 store 调用此方法报错 `Error: ... is built using the setup syntax and does not implement $reset()`
- store.$patch() 批量修改 state, 可接收一个对象或者一个函数, 如果参数为函数, 函数接收一个参数 state 表示当前 store
- store.$subscribe() 订阅 state, 侦听 state 及其变化在 patch 后只触发一次,
  默认情况下, state 订阅器被绑定在使用的组件上, 当组件卸载时, 它们将被自动移除, 如果想在组件卸载时仍保留它们, 传入第二个参数 `{ detached: true }`, 将订阅器从组件中分离
- store.$onAction() 订阅 action, 传递给它的回调函数会在 action 本身之前执行,
  默认情况下, action 订阅器被绑定在使用的组件上, 当组件卸载时, 它们将被自动移除, 如果想在组件卸载时仍保留它们, 传入第二个参数 true, 将订阅器从组件中分离

```javascript
import { useCounterStore } from './useCounterStore.js'; // counter
const counter = useCounterStore();
const unsubscribe = counter.$onAction(
  ({ name, store, args, after, onError }) => {
    /* ... */
    // after 函数将在 action 成功并完全运行后触发, 等待着任何返回的 promise
    after((result) => {});
    // onError 函数将在 action 抛出或返回一个拒绝的 promise 时触发
    onError((error) => {});
  },
  true // 组件卸载时订阅器仍会被保留
);
// 取消订阅
unsubscribe();
```

- setMapStoreSuffix() 修改 pinia 为每个 store 的 id 后面加上后缀, 默认 'Store', 修改后会影响调用辅助函数 mapStores 后的 store 的访问方式
- mapStores() 将整个 store 映射为组件的计算属性

```javascript
import { mapStores } from 'pinia';
import { useCounterStore } from './useCounterStore.js'; // counter
import { useUserStore } from './useUserStore.js'; // user
export default {
  setup() {},
  computed: {
    // 需要使用 id+'Store' 的形式访问每个 store
    ...mapStores(useCounterStore, useUserStore),
  },
  methods: {
    m1() {
      console.log(this.counterStore.count);
      this.counterStore.increment();
    },
  },
};
```

- mapState() 辅助函数, 将 state 属性映射为只读的计算属性
- mapWritableState() 辅助函数, 将 state 映射为可修改的属性

```javascript
import { mapState } from 'pinia';
import { useCounterStore } from './useCounterStore.js'; // counter
export default {
  setup() {},
  computed: {
    ...mapState(useCounterStore, ['count']),
    ...mapState(useCounterStore, {
      myCount: 'count',
      double: (store) => store.count * 2,
    }),
  },
};
```

- mapActions() 辅助函数, 将 action 属性映射为组件的方法

```javascript
import { mapActions } from 'pinia';
import { useCounterStore } from './useCounterStore.js'; // counter
export default {
  setup() {},
  methods: {
    // 将 increment 方法注册为组件的 myIncrement 方法
    ...mapActions(useCounterStore, {
      myIncrement: 'increment',
    }),
  },
};
```
