---
title: vue3.x-changelog
date: 2021-08-17 18:06:52
categories:
  - ES
tags:
  - js
  - VueJs
  - Vue3.0
---

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

- v-memo 记忆一个模板的子树,元素和组件上都可使用,接受一个固定长度的数组作为依赖值进行记忆比对,如果长度相同则跳过子树的更新

  - 如果依赖值为空数组,功能等同于 v-once
  - 结合 v-for 使用, 必须确保和 v-for 用在同一个元素上, 否则无效

  ```html
  <!-- 普通用法 -->
  <div v-memo="[valueA, valueB]"></div>
  <!-- v-for | v-memo -->
  <div v-for="item in list" :key="item.id" v-memo="[item.id === selected]">
    <p>ID: {{ item.id }} - selected: {{ item.id === selected }}</p>
    <p>more child nodes</p>
  </div>
  ```

- watchPostEffect watchEffect 的别名, 带有 flush: 'post' 选项
- watchSyncEffect watchEffect 的别名, 带有 flush: 'sync' 选项
- Effect 作用域 API, 主要用来开发库
  - effectScope 创建一个 effect 作用域对象, 以捕获在其内部创建的响应式 effect, 使得这些 effect 可以一起被处理
  - getCurrentScope 如果有, 则返回当前活跃的 effect 作用域
  - onScopeDispose 在当前活跃的 effect 作用域上注册一个处理回调, 该回调会在相关的 effect 作用域结束之后被调用
- SFC 单文件组件中使用组合式 API

  - 每次组件实例被创建的时候都会执行
  - 任何声明的顶层的绑定(包括变量，函数声明，以及 import 引入的内容)都可以在模板中直接使用

  - defineProps 和 defineEmits

    - 只能在 SFC 中使用, 不需要导入
    - defineProps 接收 props 相同的参数
    - defineEmits 接收 emits 相同的参数
    - 传入的选项不能引用在 setup 范围中声明的局部变量, 会引起编译错误

  - defineExpose 明确向外暴露出去的属性

  - useSlots 和 useAttrs

    - 在 SFC 中使用的辅助函数获取 slots 和 attrs
    - 需要手动从 vue 中导入
    - 返回的值和 setupContext.slots 和 setupContext.attrs 是等价的

  - 顶层 await

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
