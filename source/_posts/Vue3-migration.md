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

<!-- more -->

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

- 新增 `defineAsyncComponent` 助手方法, 用于显式地定义异步组件
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
- VNode Prop 格式化

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

  - Vue3.x `setup()` 中 VNode 是上下文无关的, 无法使用字符串 ID 隐式查找已注册的组件, 需要借助 `resolveComponent` 方法

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

#### 插槽统一

- `this.$slots` 插槽作为函数公开
- `this.$scopedSlots` **移除**

#### $listeners **移除**

> `$listeners` 对象在 Vue3 中被移除, 事件监听器现在是 `$attrs` 的一部分

#### $attrs 包含 class 和 style

`$attrs` 现在包含了所有传递给组件的 attribute, 包含 `class` 和 `style`

- Vue 2.x 当 `inheritAttrs: false` 时, `class` 和 `style` 不是 `$attrs` 的一部分, 仍然会被应用到组件的根元素中

### 自定义元素

#### 与自定义元素的互操作性

> 特殊的 `is` 属性的使用被严格限制在保留的 `<component>` 标签中
> 使用 `vue:` 前缀解决 DOM 内模板解析问题

```html
<!-- Vue 2.x -->
<table>
  <tr is="blog-post-row"></tr>
</table>

<!-- Vue 3.x -->
<table>
  <tr is="vue:blog-post-row"></tr>
</table>
```

### 移除的 APIs

#### 按键修饰符

- 不再支持使用数字(按键)作为 `v-on` 修饰符
- 不再支持 `config.keyCodes`

```html
<!-- Vue 2.x -->
<!-- 键码 -->
<input v-on:keyup.13="submit" />
<!-- 别名 -->
<input v-on:keyup.enter="submit" />

<!-- Vue 3.x -->
<input v-pn:keyup.page-down="nextPage" />
<!-- 同时匹配 q 和 Q -->
<input v-on:keypress.q="quit" />
```

#### 事件 API

`$on`, `$off`, `$once` 实例方法已被 **移除**, 组件实例不再实现事件触发接口

#### 过滤器

`$filter` 过滤被**移除**

#### 内联模板 Attribute

对 **内联模板特性** 的支持被**移除**

- Vue 2.x 使用 `inline-template` 属性将其内容作为模板使用, 而不是作为分发内容

```html
<my-component inline-template>
  <div>
    <p>它们被编译为组件自己的模板</p>
    <p>不是父级所包含的内容。</p>
  </div>
</my-component>
```

#### $children

`$children` 实例属性已被 **移除**

#### propsData

- Vue 2.x `propsData` 选项用于在创建 Vue 实例的过程中传入的 prop, 现在已被 **移除**
- Vue 3.x 使用 `createApp` 的第二个参数传入 prop

```javascript
// Vue 2.x
const Comp = Vue.extend({
  props: ['name'],
  template: '<div>{{name}}</div>',
});
new Comp({
  propsData: {
    name: 'hello world',
  },
});

// Vue 3.x
import { createApp } from 'vue';
const app = createApp(
  {
    props: ['name'],
    template: '<div>{{name}}</div>',
  },
  { name: 'hello world' }
);
```

### 其他变化

#### 片段

- 组件可以包含多个根节点, 需要显示定义 attribute 的位置

```html
<template>
  <header>...</header>
  <main v-bind="$attrs">...</main>
  <footer>...</footer>
</template>
```

#### attribute 强制行为

> 底层的内部 API 更改, 绝大多数开发人员不会受到影响

#### 自定义指令

- 指令的钩子函数已经被重命名, 以更好地与组件的生命周期保持一致
- `expression` 字符串不再作为 `binding` 对象的一部分被传入

##### 钩子函数

|     Vue 2.x      |    Vue 3.x    |
| :--------------: | :-----------: |
|                  |    created    |
|       bind       |  beforeMount  |
|     inserted     |    mounted    |
|                  | beforeUpdate  |
|      update      |               |
| componentUpdated |    updated    |
|                  | beforeUnmount |
|      unbind      |   unmounted   |

#### Data 选项

- 组件选项 `data` 的声明不再接收纯 JavaScript `Object`, 而是接收一个 `function`
- 当合并来自 `mixin` 或 `extend` 的多个 `data` 返回值时, 合并操作现在是浅层次的而非深层次的(只合并根级属性)

#### mount API

> 挂载元素时, 被渲染的应用作为子元素插入, 不再替换要挂载的目标元素

- Vue 2.x 当挂载一个具有 `template` 的应用时, 被渲染的内容会替换要挂载的目标元素
- Vue 3.x 被挂载的应用会作为子元素插入, 从而替换目标元素的 `innerHTML`

```html
<!-- Vue 2.x 替换目标元素 -->
<template>
  <div id="rendered">hello world!</div>
</template>
<script>
  const app = new Vue({
    data() {
      return {
        message: 'hello world!',
      };
    },
    template: '<div id="rendered">{{message}}</div>',
  });
  app.$mount('#app');
</script>
```

```html
<!-- Vue 3.x 作为子元素插入到目标元素 -->
<template>
  <div id="app">
    <div id="rendered">hello world!</div>
  </div>
</template>
<script>
  import { createApp } from 'vue';
  const app = createApp({
    data() {
      return { message: 'hello world!' };
    },
    template: '<div id="rendered">{{message}}</div>',
  });
  app.mount('#app');
</script>
```

#### 在 prop 的默认函数中不能再访问 this

- 组件接收到的原始 prop 作为参数传递给默认函数
- `inject` API 在默认函数中使用

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

#### Transition class

过渡类名 `v-enter` 修改为 `v-enter-from`, 过渡类名 `v-leave` 修改为 `v-leave-from`

- v-enter-from
- v-enter-active
- v-enter-to
- v-leave-from
- v-leave-active
- v-leave-to

#### Transition 作为根节点

当使用 `<Transition>` 作为根节点的组件从外部被切换时将不再触发过渡效果

#### TransitionGroup 根元素

`<TransitionGroup>` 不再默认渲染根元素, 但仍然可以用 `tag` 属性创建根元素

#### VNode 生命周期事件

- Vue 2.x 监听组件生命周期的事件名以 `hook:` 前缀开头, 并跟随相应的生命周期钩子的名字

```html
<template>
  <child-component @hook:updated="onUpdated" />
</template>
```

- Vue 3.x 前缀改为 `vue:` 开头, 这些事件也可用于 HTML 元素, 和在组件上的用法一样
  - 绝大多数情况下只需要修改前缀. 生命周期钩子 `beforeDestroy` 和 `destroyed` 已经分别被重命名为 `beforeUnmount` 和 `unmounted`, 所以相应的事件名也需要更新

```html
<template>
  <child-component @vue:mounted="onMounted" />
  <child-component @vue:before-update="onBeforeUpdate" />
  <!-- 等同于 -->
  <child-component @vue:beforeUpdate="onBeforeUpdate" />
</template>
```

#### 侦听数组

当侦听一个数组时, 只有当数组被替换时才会触发回调, 如果需要在数组被改变时触发回调, 必须指定 `deep` 选项
