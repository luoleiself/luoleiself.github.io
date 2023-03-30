---
title: Vue3-migration
date: 2021-08-17 18:06:52
categories:
  - [ES, VueJs]
tags:
  - js
  - VueJs
  - Vue3
---

### 全局 API

#### 全局 API 应用实例

- Vue.config -> app.config
- Vue.config.productionTip -> **移除**
- Vue.config.ignoreElements -> app.config.compilerOptions.isCustomElement
- Vue.component -> app.component
- Vue.directive -> app.directive
- Vue.mixin -> app.mixin
- Vue.use -> app.use
- Vue.prototype -> app.config.globalProperties
- Vue.extend -> **移除**

#### 全局 API Treeshaking

- Vue.nextTick() -> nextTick()
- Vue.observable -> reactive()
- Vue.version -> version
- Vue.compile(仅完整构建版本)
- Vue.set(仅兼容构建版本)
- Vue.delete(仅兼容构建版本)

### 模板指令

#### v-model

参见 vue3.md 的 v-model 指令部分

#### key 使用改变

对于 `v-if`, `v-else-if`, `v-else` 的各分支项 `key` 将不再是必须的, Vue 会自动生成唯一的 `key`

#### v-if 和 v-for 优先级

> 两者作用于同一个元素上时，v-if 会拥有比 v-for 更高的优先级

#### v-bind 合并行为

`v-bind` 的绑定顺序会影响渲染结果

- vue 2.x 独立绑定的 attribute 会覆盖 v-bind 中的 attribute
- vue 3.x 根据声明顺序决定如何被合并

```html
<!-- 模板 -->
<div id="red" v-bind="{ id: 'blue' }"></div>
<!-- 结果 -->
<div id="red"></div>

<!-- 模板 -->
<div v-bind="{ id: 'blue' }" id="red"></div>
<!-- 结果 -->
<div id="red"></div>
```

#### v-on.native 移除

`v-on` 的 `.native` 修饰符被**移除**

- 确保所有组件都是用新增的 `emits` 选项记录其事件

### 组件

#### 函数式组件

函数组件将接收两个参数: `props` 和 `context`, context 是一个包含 `attrs`, `slots`, `emit` 属性的上下文对象

- `<template>` 中的 `functional` 属性被**移除**
- `{functional: true}` 选项从通过函数创建的组件中**移除**
- `listeners` 作为 `$attrs` 的一部分传递

```javascript
import { h } from 'vue';
const DynamicHeading = (props, { attrs, slots, emit, expose }) => {
  return h(`h${props.level}`, attrs, slots);
};
DynamicHeading.props = ['level'];
export default DynamicHeading;
```

#### 异步组件

- 新的 `defineAsyncComponent` 助手方法,用于显式地定义异步组件
- `component` 选项重命名为 `loader`
- `Loader` 函数本身不再接收 `resolve` 和 `reject` 参数，且必须返回一个 Promise

```javascript
// Vue 2.x
const asyncPage = () => import('./NextPage.vue');
const asyncPage = {
  component: () => import('./NextPage.vue'),
  delay: 200,
  timeout: 3000,
  error: ErrorComponent,
  loading: LoadingComponent,
};
// Vue 3.x
// 不带选项的异步组件
const asyncPage = defineAsyncComponent(() => import('./NextPage.vue'));
// 带选项的异步组件
const asyncPageWithOptions = defineAsyncComponent({
  loader: () => import('./NextPage.vue'),
  delay: 200,
  timeout: 3000,
  errorComponent: ErrorComponent,
  loadingComponent: LoadingComponent,
});
```

#### emits 选项

`新增`项用来定义组件可以向其父组件触发的事件

### 渲染函数

#### 概览

- `h` 函数现在为全局引入，不再作为 `render` 函数的参数隐式提供

```javascript
import { h } from 'vue';
export default {
  render() {
    h('div');
  },
};
```

- 更改渲染函数参数, 使其在有状态组件和函数组件的表现更加一致
- VNode 现有一个扁平的 prop 结构

```javascript
// 2.x
{
  staticClass: 'button',
  class: { 'is-outlined': isOutlined },
  staticStyle: { color: '#34495E' },
  style: { backgroundColor: buttonColor },
  attrs: { id: 'submit' },
  domProps: { innerHTML: '' },
  on: { click: submitForm },
  key: 'submit-button'
}
// 3.x 语法
{
  class: ['button', { 'is-outlined': isOutlined }],
  style: [{ color: '#34495E' }, { backgroundColor: buttonColor }],
  id: 'submit',
  innerHTML: '',
  onClick: submitForm,
  key: 'submit-button'
}
```

- 注册组件

  - Vue3.x 中 VNode 是上下文无关的, 无法使用字符串 ID 隐式查找已注册的组件, 需要借助 `resolveComponent` 方法

```javascript
// Vue 2.x
Vue.component('my-component', {
  data() {
    return {
      count: 0,
    };
  },
  template: `<button @click="count++">
    Clicked {{ count }} times.
  </button>`,
});
export default {
  render(h) {
    return h('my-component');
  },
};
// Vue 3.x
import { h, resolveComponent } from 'vue';
// resolveComponent,resolveDynamicComponent,resolveDirective,withDirectives
// 全局API只能在 render 或 setup 函数中使用
export default {
  setup() {
    const MyComponent = resolveComponent('my-component');
    return () => h(MyComponent);
  },
};
```





























### v-for 中的 Ref 数组

- Vue 3.x 中 不再在 $ref 中自动创建数组填充相应的 $refs property

### 片段

- 组件可以包含多个根节点, 需要显示定义 attribute 的位置

  ```html
  <template>
    <header>...</header>
    <main v-bind="$attrs">...</main>
    <footer>...</footer>
  </template>
  ```

### 内联模板 Attribute 移除

```html
<my-component inline-template>
  <div>
    <p>它们被编译为组件自己的模板</p>
    <p>不是父级所包含的内容。</p>
  </div>
</my-component>
```

### 按键修饰符

- 不再支持使用数字(即键码)作为 v-on 修饰符
- 不再支持 config.keyCodes

### propsData 移除

- 使用 createApp 第二个参数, 向根组件传入 prop

### 在 prop 的默认函数中不能再访问 this

- 组件接收到的原始 prop 作为参数传递给默认函数
- inject API 在默认函数中使用

  ```javascript
  import { inject } from 'vue';
  export default {
    props: {
      theme: {
        default(props) {
          // `props` 是传递给组件的原始值。
          // 在任何类型/默认强制转换之前
          // 也可以使用 `inject` 来访问注入的 property
          return inject('theme', 'default-theme');
        },
      },
    },
  };
  ```

### 插槽统一

- this.$slots 插槽作为函数公开
- 移除 this.$scopedSlots

  ```javascript
  // 2.x 语法
  this.$scopedSlots.header;
  // 3.x 语法
  this.$slots.header();
  ```

- 见指令 v-slot
- 解构插槽 prop

  ```html
  <!-- 解构并且重命名 -->
  <todo-list v-slot="{ item: todo }">
    <i class="fas fa-check"></i>
    <span class="green">{{ todo }}</span>
  </todo-list>
  ```

- 动态插槽名

  ```html
  <base-layout>
    <template v-slot:[dynamicSlotName]> ... </template>
  </base-layout>
  ```

### 过渡的 class 名更改

- 过渡类名 v-enter 修改为 v-enter-from
- 过渡类名 v-leave 修改为 v-leave-from

### Transition Group 根元素

> 不再默认渲染根元素, 仍然可以用 tag prop 创建根元素

### VNode 生命周期事件

- 2.x

```html
<child-component @hook:updated="onUpdated"></child-component>
```

- 3.x
  - 绝大多数情况下只需要修改前缀. 生命周期钩子 beforeDestroy 和 destroyed 已经分别被重命名为 beforeUnmount 和 unmounted,所以相应的事件名也需要更新

```html
<child-component @vnode-updated="onUpdated"></child-component>
<!-- 或者使用驼峰命名法 -->
<child-component @vnodeUpdated="onUpdated"></child-component>
```

### Vue 3.x 迭代 新增或者废弃的语法和 API

### 源码优化

#### 代码管理方式 monorepo(monolithic repository)

- vue 2.x 源码托管在 src 目录下, 然后根据功能拆分出不同的目录 compiler, core, platforms, server, sfc, shared

- vue 3.0 通过 monorepo 方式管理代码结构,按功能拆分到不同的 package 中, 每个 package 有各自的 API, 类型定义, 测试

#### 使用 TypeScript 规范代码类型

#### 性能优化

- 源码体积优化: 移除不常用的 Feature, 更友好的 tree-shaking 减少打包体积
- 数据劫持优化: Vue 区别于 React 的一大特色是数据是响应式的,
- Vue 3.0 使用 Proxy 代理(嵌套对象在访问时递归响应式而非初始化递归, 提升性能)

#### 编译优化

- 静态模板优化：vue 3.0 编译阶段对静态模板的分析,区分动态内容和静态内容提升更新性能
- Slot, 事件监听函数缓存优化

#### Composition API 属于 API 的增强, 非开发规范

- Composition API 非 Vue 3.0 的开发规范, 可以根据需求选择使用 Composition API 还是 Options API

<!-- more -->

### vue 3.1

#### 新增

- compilerOptions 配置运行时编译器的选项

  - isCustomElement 指定一个方法来识别 Vue 以外 (例如通过 Web Components API) 定义的自定义元素
  - whitespace 编译模板时对模板元素之间的空格的处理方式, 默认值: condense
    - condense
    - preserve
  - delimiters 设置模板内的文本插值的边界符
  - comments 是否在生产环境保留注释

  ```javascript
  // 应用中使用
  app.config.compilerOptions = {
    isCustomElement: (tag) => tag.startWith('icon'),
    whitespace: 'condense', // 默认值
    delimiters: ['{{', '}}'], // 默认值
    comments: false, // 默认值
  };
  // 组件中使用
  createApp({
    // ...
    compilerOptions: {},
  });
  ```

- is 特殊属性

  - 动态组件中正常使用

    ```html
    <component :is="currentView"></component>
    ```

  - 原生元素上使用 is 属性, vue: 前缀声明的元素会被 Vue 组件进行渲染替换

    ```html
    <table>
      <tr is="vue:my-row-component"></tr>
    </table>
    ```

#### 废弃

- isCustomElement 使用 compilerOptions.isCustomElement 代替
- delimiters 使用 compilerOptions.delimiters 代替
- v-is 使用 is 代替, 参考上方新增中的语法

### vue 3.2

#### 新增

- defineCustomElement 接受和 defineComponent 相同的参数, 但是返回一个原生的自定义元素, 该元素可以用于任意框架或不基于框架使用

  ```javascript
  import { defineCustomElement } from 'vue';
  // <my-element></my-element>
  const MyElement = defineCustomElement({});
  customElements.define('my-element', MyElement); // 注册该自定义元素
  document.body.append(new MyElement(/* 初始化 prop */));
  ```

- v-bind 修饰符

  - .camel 将 kebab-case attribute 名转换为 camelCase, 3.2 之前新增

    ```html
    <svg :view-box.camel="viewBox"></svg>
    ```

  - .prop 将一个绑定强制设置为一个 DOM Property
  - .attr 将一个绑定强制设置为一个 DOM Attribute

  ```html
  <!-- 使用 .prop 修饰符，会从组件选项 props 中移除, 
    以 .[attr] 形式出现在组件 attrs 参数中并且不会显示在 DOM 上 -->
  <!-- 使用 .attr 修饰符，会从组件选项 props 中移除, 
    以 ^[attr] 形式出现在组件 attrs 参数中并且会显示在 DOM 上 -->
  <div :someProperty.prop="someObject"></div>
  <!-- // 相当于 -->
  <div .someProperty="someObject"></div>
  ```

- SFC 单文件组件

  - 每次组件实例被创建的时候都会执行
  - 任何声明的顶层的绑定(包括变量，函数声明，以及 import 引入的内容)都可以在模板中直接使用

  - defineProps 和 defineEmits

    - 只能在 SFC 中使用, 不需要导入
    - defineProps 接收 props 相同的参数
    - defineEmits 接收 emits 相同的参数
    - 传入的选项不能引用在 setup 范围中声明的局部变量, 会引起编译错误

  - defineExpose 明确向外暴露出去的属性

    - 不需要导入直接使用

  - useSlots 和 useAttrs

    - 在 SFC 中使用的辅助函数获取 slots 和 attrs
    - 需要手动从 vue 中导入
    - 返回的值和 setupContext.slots 和 setupContext.attrs 是等价的

  - 顶层 await, 结果代码会被编译成 `async setup()`

    ```html
    <script setup>
      // setup 会被编译成 async setup()
      const post = await fetch('/api/example');
    </script>
    ```

  ```html
  <script setup>
    import { capitalize } from './helpers';
    import { ref } from 'vue';
    import MyComponent from './MyComponent.vue'; // 使用组件
    console.log('hello script setup');
    const count = ref(0);
    // 变量
    const msg = 'Hello!';
    // 函数
    function log() {
      console.log(msg);
    }

    // defineProps 和 defineEmits
    const props = defineProps(['foo']);

    const emit = defineEmits(['change', 'delete']);

    // defineExpose
    defineExpose({ msg });

    // useSlots 和 useAttrs
    import { useSlots, useAttrs } from 'vue';
    const slots = useSlots();
    const attrs = useAttrs();
  </script>
  <template>
    <div @click="log">{{ msg }}</div>
    <div>{{capitalize('hello')}}</div>
    <button @click="count++">{{count}}</button>
    <MyComponent />
  </template>
  ```

#### 废弃
