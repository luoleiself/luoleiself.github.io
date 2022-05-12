---
title: Vue3.0知识点
date: 2021-06-19 15:19:26
categories:
  - ES
  - VueJs
tags:
  - js
  - VueJs
  - Vue3.0
---

> Vue 3.0.11

## 应用配置

```javascript
// 每个 Vue 应用会暴露一个 config 对象, 在挂载应用之前, 可修改其属性
import { createApp } from 'vue';
const app = createApp({});
console.log(app.config);
```

- errorHandler 处理组件渲染过程中抛出的未捕获错误

  - 类型: Function
  - 默认: undefined

  ```javascript
  app.config.errorHandler = (err, vm, info) => {
    // 处理错误
  };
  ```

- warnHandler 处理 Vue 运行中的警告, 开发环境下有效

  - 类型: Function
  - 默认: undefined

  ```javascript
  app.config.warnHandler = (msg, vm, trace) => {
    // 处理警告
  };
  ```

- globalProperties 添加应用程序内任何组件实例都可访问的全局 property, 属性名冲突时,组件内 property 优先

  - 类型: [key: string]: any
  - 默认: undefined

  ```javascript
  // Vue 2.x
  Vue.prototype.$xhr = () => {};
  // Vue 3.x
  app.config.globalProperties.$xhr = () => {};
  ```

- optionMergeStrategies 合并策略选项分别接收在父实例和子实例上定义的该选项的值作为第一个和第二个参数, 引用上下文实例被作为第三个参数传入

  - 类型: { [key: string]: Function }
  - 默认: {}

  ```javascript
  const app = Vue.createApp({
    mounted() {
      console.log(this.$options.hello);
    },
  });
  app.config.optionMergeStrategies.hello = (parent, child, vm) => {
    return `Hello, ${child}`;
  };
  app.mixin({
    hello: 'Vue',
  });
  // 'Hello, Vue'
  ```

- performance 启用对组件初始化、编译、渲染和更新的性能追踪

  - 类型: boolean
  - 默认: false

  ```javascript
  app.config.performance = true;
  ```

- isCustomElement (3.1 废弃) 指定一个方法, 用来识别在 Vue 之外定义的自定义元素, 如果组件符合此条件, 则不需要本地或全局注册

  > 注意, 所有原生 HTML 和 SVG 标记不需要在此函数中匹配——Vue 解析器自动执行此检查

  - 类型: (tag: string) => boolean
  - 默认: undefined

  ```javascript
  app.config.isCustomElement = (tag) => tag.startsWith('icon-');
  ```

- compilerOptions (3.1 新增) 配置运行时编译器的选项, 设置在这个对象上的值将会被传入浏览器内的模板编译器，并影响配置过的应用内的每个组件

  ```javascript
  app.config.compilerOptions = {
    isCustomElement: (tag) => {},
    whitespace: 'condense',
    delimiters: ['{{', '}}']
    comments: false,
  };
  ```

  - isCustomElement 作用同上
  - whitespace 默认情况下，Vue 会移除/压缩模板元素之间的空格以产生更高效的编译结果

    1. 元素内的多个开头/结尾空格会被压缩成一个空格
    2. 元素之间的包括折行在内的多个空格会被移除
    3. 文本结点之间可被压缩的空格都会被压缩成为一个空格

    ```javascript
    app.config.compilerOptions.whitespace = 'condense';
    // 设置 preserve 可以禁用 b 和 c
    ```

  - delimiters 设置用在模板内的文本插值的边界符
  - comments 默认生产环境移除模板内的 HTML 注释

    ```javascript
    app.config.compilerOptions.comments = true; // 生产环境保留注释
    ```

<!-- more -->

## 应用 API

> 调用 createApp 返回一个应用实例，该实例提供了一个应用上下文, 应用实例挂载的整个组件树共享相同的上下文

- component 注册或检索全局组件

  - 参数
    - {string} name
    - {Function | Object} [definition]
  - 返回值
    - 传入 definition 参数，返回应用实例
    - 不传入 definition 参数，返回组件定义

  ```javascript
  import { createApp } from 'vue';
  const app = createApp({});
  // 注册一个名为 my-component 的组件
  app.component('my-component', {
    /* ... */
  });
  // 检索注册的组件(始终返回构造函数)
  const MyComponent = app.component('my-component');
  ```

- config 应用配置对象
- directive 注册或检索全局指令

  - 参数
    - {string} name
    - {Function | Object} [definition]
  - 返回值
    - 传入 definition 参数，返回应用实例
    - 不传入 definition 参数，返回指令定义

  ```javascript
  import { createApp, onCreated, onMounted } from 'vue';
  const app = createApp({});
  // 注册
  app.directive('my-directive', {
    // 指令是具有一组生命周期的钩子：
    // 在绑定元素的 attribute 或事件监听器被应用之前调用
    created(el, binding, vnode, prevNode) {},
    // 在绑定元素的父组件挂载之前调用
    beforeMount(el, binding, vnode, prevNode) {},
    // 绑定元素的父组件被挂载时调用
    mounted(el, binding, vnode, prevNode) {},
    // 在包含组件的 VNode 更新之前调用
    beforeUpdate(el, binding, vnode, prevNode) {},
    // 在包含组件的 VNode 及其子组件的 VNode 更新之后调用
    updated(el, binding, vnode, prevNode) {},
    // 在绑定元素的父组件卸载之前调用
    beforeUnmount(el, binding, vnode, prevNode) {},
    // 卸载绑定元素的父组件时调用
    unmounted(el, binding, vnode, prevNode) {},
  });
  // 注册 (功能指令)
  app.directive('my-directive', () => {
    // 这将被作为 `mounted` 和 `updated` 调用
  });
  // getter, 如果已注册，则返回指令定义
  const myDirective = app.directive('my-directive');
  ```

- mixin 将一个 mixin 应用到整个应用范围内

  - 参数
    - {Object} mixin
  - 返回值: 应用实例

  ```javascript
  import { createApp } from 'vue';
  const app = createApp({});
  app.mixin({});
  ```

- mount 挂载应用实例

  - 参数
    - {Element | string} rootContainer
    - {boolean} isHydrate
  - 返回值: 根组件实例

  ```javascript
  <body>
    <div id='my-app'></div>
  </body>;
  import { createApp } from 'vue';
  const app = createApp({});
  // 做一些必要的准备
  app.mount('#my-app');
  ```

- unmount 卸载应用实例的根组件

  ```javascript
  <body>
    <div id='my-app'></div>
  </body>;
  import { createApp } from 'vue';
  const app = createApp({});
  // 做一些必要的准备
  app.mount('#my-app');
  // 挂载5秒后，应用将被卸载
  setTimeout(() => app.unmount(), 5000);
  ```

- provide 设置一个可以被注入到应用范围内所有组件中的值

  - 参数
    - {string | Symbol} key
    - value
  - 返回值: 应用实例

  ```javascript
  import { createApp } from 'vue';
  const app = createApp({
    inject: ['user'],
    template: `
      <div>
        {{ user }}
      </div>
    `,
  });
  app.provide('user', 'administrator');
  ```

- use 安装 Vue.js 插件。如果插件是一个对象，它必须暴露一个 install 方法。如果它本身是一个函数，它将被视为安装方法

  - 参数
    - {Object | Function} plugin
    - ...options (可选)
  - 返回值: 应用实例

  ```javascript
  import { createApp } from 'vue';
  import MyPlugin from './plugins/MyPlugin';
  const app = createApp({});
  app.use(MyPlugin);
  app.mount('#app');
  ```

- version 以字符串形式提供已安装的 Vue 的版本号

## 全局 API

- createApp 返回一个提供应用上下文的应用实例,应用实例挂载的整个组件树共享同一个上下文

  - 参数
    - {Object} rootConf 根组件选项对象
    - {Object} propConf 根 prop

  ```javascript
  import { createApp } from 'vue';
  const app = createApp(
    {
      props: ['username'],
      data() {
        return {};
      },
      methods: {},
      computed: {},
    },
    { username: 'Hello World' }
  );
  ```

- h 返回一个"虚拟节点",通常缩写为 VNode：一个普通对象，其中包含向 Vue 描述它应在页面上渲染哪种节点的信息,包括所有子节点的描述

  - 参数
    - {String | Object | Function} type
    - {Object} props
    - {String | Array} children

  ```javascript
  render() {
    // 使用返回 null 的函数将渲染一个注释
    return Vue.h('div', {}, [
      Vue.h(null, {}, [' Annotation start ']),
      Vue.h('h2', {}, ['Hello World']),
      Vue.h('p', {}, 'This is a label p, it contain description...'),
      Vue.h(null, {}, [' Annotation end '])
    ])
  }
  ```

- defineComponent 创建一个合成类型的构造函数,用于手动渲染函数、TSX 和 IDE 工具支持

  1. 具有组件选项的对象

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

  2. 一个 setup 函数, 函数名称作为组件名称使用

     ```javascript
     import { createApp, defineComponent, ref } from 'vue';
     const HelloWorld = defineComponent((props, { attrs, emit, slots }) => {
       const count = ref(0);
       return { count };
     });
     const app = createApp(HelloWorld).mount('#app');
     ```

- defineAsyncComponent 创建一个只有在需要时才会加载的异步组件

  1. 接受一个返回 Promise 的工厂函数

     ```javascript
     // 全局注册
     import { defineAsyncComponent } from 'vue';
     const AsyncComp = defineAsyncComponent(() => import('./components/AsyncComponent.vue'));
     app.component('async-component', AsyncComp);
     // 局部注册
     import { createApp, defineAsyncComponent } from 'vue';
     createApp({
       // ...
       components: {
         AsyncComponent: defineAsyncComponent(() => import('./components/AsyncComponent.vue')),
       },
     });
     ```

  2. 接受一个对象

     ```javascript
      import { defineAsyncComponent } from 'vue'
      const AsyncComp = defineAsyncComponent({
        // 工厂函数
        loader: () => import('./Foo.vue')
        // 加载异步组件时要使用的组件
        loadingComponent: LoadingComponent,
        // 加载失败时要使用的组件
        errorComponent: ErrorComponent,
        // 在显示 loadingComponent 之前的延迟 | 默认值：200（单位 ms）
        delay: 200,
        // 如果提供了 timeout ，并且加载组件的时间超过了设定值，将显示错误组件
        // 默认值：Infinity（即永不超时，单位 ms）
        timeout: 3000,
        // 定义组件是否可挂起 | 默认值：true
        suspensible: false,
        /**
        *
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
      })
     ```

- resolveComponent 允许按名称解析 component

  > resolveComponent 只能在 render 或 setup 函数中使用

  - 参数

    - {String} name 已加载的组件的名称

  - 返回值: 返回一个 Component, 如果找不到则返回接收的参数 name

  ```javascript
  const app = Vue.createApp({});
  app.component('MyComponent', {
    /* ... */
  })
  import { resolveComponent } from 'vue';
  render() {
    const MyComponent = resolveComponent('MyComponent');
  }
  ```

- resolveDynamicComponent 返回已解析的 Component 或新创建的 VNode，其中组件名称作为节点标签。如果找不到 Component,将发出警告

  > 允许使用与 &lt;component :is=""&gt; 相同的机制来解析一个 component

  > resolveDynamicComponent 只能在 render 或 setup 函数中使用

  - 参数
    - {String | Object} component 组件

  ```javascript
  import { resolveDynamicComponent } from 'vue';
  render () {
    const MyComponent = resolveDynamicComponent('MyComponent');
  }
  ```

- resolveDirective 允许按名称解析一个 directive

  > resolveDirective 只能在 render 或 setup 函数中使用

  - 参数
    - {String} name 已加载的指令名称
  - 返回值: 返回一个 Directive, 如果没有找到则返回 undefined

  ```javascript
  import { createApp, resolveDirective } from 'vue';
  const app = createApp({});
  app.directive('highlight', {});
  render () {
    const highlightDirective = resolveDirective('highlight');
  }
  ```

- withDirectives 允许将指令应用于 VNode。返回一个包含应用指令的 VNode

  > withDirectives 只能在 render 或 setup 函数中使用

  - 参数

    - {Element} vnode 使用 h() 创建的虚拟节点
    - {Array} directives 指令数组

      - 每个指令都是一个数组, 最多可以定义 4 个索引

        - [directive] 指令本身
        - [directive, value] 上述内容,指令的值
        - [directive, value, arg] 上述内容,一个 String 参数,eg: v-on:click 中的 click
        - [directive, value, arg, modifiers] 上述内容,定义任意修饰符的 key:value 键值对

        ```javascript
        const MyDirective = resolveDirective('MyDirective');
        const nodeWithDirectives = withDirectives(h('div'), [[MyDirective, 100, 'click', { prevent: true }]]);
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

- createRenderer 自定义渲染器可以传入特定于平台的类型

  - 参数
    - HostNode 宿主环境中的节点
    - HostElement 宿主环境中的元素

  ```javascript
  import { createRenderer } from 'vue';
  const { render, createApp } = createRenderer<Node, Element>({
    patchProp,
    ...nodeOps
  })
  ```

- nextTick 将回调推迟到下一个 DOM 更新周期之后执行

  ```javascript
  import { createApp, nextTick } from 'vue';
  const app = createApp({
    setup() {
      const message = ref('Hello!');
      const changeMessage = async (newMessage) => {
        message.value = newMessage;
        await nextTick();
        console.log('Now DOM is updated');
      };
    },
  });
  ```

- mergeProps 将包含 VNode prop 的多个对象合并为一个单独的对象, 返回一个新创建的对象, 而作为参数传递的对象则不会被修改.

  ```javascript
  import { h, mergeProps } from 'vue';
  export default {
    inheritAttrs: false,
    render() {
      const props = mergeProps(
        {
          // 该 class 将与 $attrs 中的其他 class 合并。
          class: 'active',
        },
        this.$attrs
      );
      return h('div', props);
    },
  };
  ```

- useCssModule 允许在 setup 的单文件组件函数中访问 CSS 模块

  > useCssModule 只能在 render 或 setup 函数中使用

  - 参数
    - {String} name CSS 模块的名称, 默认为 '$style'

  ```javascript
  <script>
  import { h, useCssModule } from 'vue'
  export default {
    setup () {
      const style = useCssModule()
      return () => h('div', {
        class: style.success
      }, 'Task complete!')
    }
  }
  </script>
  <style module>
  .success {
    color: #090;
  }
  </style>
  ```

- version 以字符串形式提供已安装的 Vue 的版本号

## 选项

### Data

#### data

- vm.a 等价于 vm.$data.a
- 以 \_ 或 $ 开头的 property 不会被组件实例代理, vm.$data.\_property

#### props 用于接收来自父组件的数据

#### computed 计算属性

- 计算属性的结果会被缓存,依赖的响应式 property 变化才会重新计算

#### methods

#### watch

#### emits 定义组件触发的自定义事件

> emits 选项中列出的事件不会从组件的根元素继承，也将从 $attrs property 中移除

- Array&lt;string&gt; | Object

  ```javascript
  import { createApp } from 'vue';
  const app = createApp({});
  // 数组语法
  app.component('todo-item', {
    emits: ['check'],
    created() {
      this.$emit('check');
    },
  });
  // 对象语法
  app.component('reply-form', {
    emits: {
      // 没有验证函数
      click: null,
      // 带有验证函数
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
      this.$emit('submit', { email: '', password: '' });
    },
  });
  ```

### DOM

#### template

> 如果 Vue 选项中包含渲染函数，该模板将被忽略

#### render

> render 函数的优先级高于从挂载元素 template 选项或内置 DOM 提取出的 HTML 模板编译渲染函数

### 生命周期钩子

> 所有的生命周期钩子自动绑定 this 上下文到实例中, 不能使用箭头函数来定义一个生命周期方法

#### beforeCreate

#### created

#### beforeMount

> 该钩子在服务器端渲染期间不被调用

#### mounted

> 该钩子在服务器端渲染期间不被调用

#### beforeUpdate

> 该钩子在服务器端渲染期间不被调用，因为只有初次渲染会在服务端进行

#### updated

> 该钩子在服务器端渲染期间不被调用

#### activated 被 keep-alive 缓存的组件激活时调用

> 该钩子在服务器端渲染期间不被调用

#### deactivated 被 keep-alive 缓存的组件停用时调用

> 该钩子在服务器端渲染期间不被调用

#### beforeUnmount

> 该钩子在服务器端渲染期间不被调用

#### unmounted

> 该钩子在服务器端渲染期间不被调用

#### errorCaptured 当捕获一个来自子孙组件的错误时被调用

- 参数
  - err: Error
  - instance: Component
  - info: string
- 返回值: 可以返回 false 阻止错误继续向上传播

#### renderTracked 跟踪虚拟 DOM 重新渲染时调用

- 参数
  - e: DebuggerEvent

#### renderTriggered 当虚拟 DOM 重新渲染被触发时调用

- 参数
  - e: DebuggerEvent

### 选项/资源

#### directives 包含组件实例可用指令的哈希表

#### components 包含组件实例可用组件的哈希表

### 组合

#### mixins

- Mixin 钩子按照传入顺序依次调用，并在调用组件自身的钩子之前被调用

#### extends 允许声明扩展另一个组件

```javascript
const CompA = { ... }
// 在没有调用 `Vue.extend` 时候继承 CompA
const CompB = {
  extends: CompA,
  ...
}
```

#### provide / inject

> 允许一个祖先组件向其所有子孙后代注入一个依赖,不论组件层次深度

> provide 和 inject 绑定并不是响应式的

```javascript
// 父级组件 provide  'foo'
const Provider = {
  provide: {
    foo: 'bar',
  },
  // 或者是返回一个对象的函数
  // provide() {
  //   return {
  //     foo: 'bar',
  //   };
  // },
};
// 子组件 inject  'foo'
const Child = {
  inject: ['foo'],
  // 或者是一个对象
  // inject: {
  //   foo: {
  //     default: 'foo',
  //   },
  // },
  created() {
    console.log(this.foo); // => "bar"
  },
  // ...
};
```

#### setup 组合式 API 的入口点

- 在创建组件实例时, 在初始 prop 解析之后立即调用 setup
- 在生命周期方面, 在 beforeCreate 钩子之前调用
- 参数
  - props 接收父组件传入的属性, 不要结构 props 对象,会失去响应式
  - ctx 上下文对象, 包含了 attrs, slots, emit
- 返回值

  - Object 对象的属性合并到组件实例的上下文
  - h 返回一个渲染函数, 该函数可以直接使用在同一作用域中声明的响应式状态

    ```javascript
    import {h, ref, reactive} from 'vue';
    setup() {
      const count = ref(0)
      const object = reactive({ foo: 'bar' })
      return () => h('div', [count.value, object.foo])
    }
    ```

### 杂项

#### name

#### inheritAttrs

- 控制是否在子组件的根元素上显示不被认作 props 的 attributes 绑定

- 默认值: true

#### delimiters (3.1 废弃)

- 默认值: ['{{', '}}'] 模板文本插值的分隔符

#### compilerOptions (3.1 新增) 配置运行时编译器的选项

配置同应用 API compilerOptions

## 实例 Property

- $data
- $props
- $el
- $options 当前组件实例的初始化选项
- $parent 父组件实例
- $root 根组件实例
- $slots 用来访问被插槽分发的内容
- $refs
- $attrs

## 实例方法

- $watch
- $emit
- $forceUpdate 迫使组件实例重新渲染,仅仅影响实例本身和插入插槽内容的子组件
- $nextTick

## 指令

- v-text
- v-html
- v-show
- v-else-if
- v-for
- v-on
- v-bind
- v-model
- v-slot 只能使用在 template 标签上, 当被提供的内容只有默认插槽时, 组件的标签才可以当作插槽的模板使用(v-slot 写在标签上)
- v-pre 跳过这个元素和它的子元素的编译过程
- v-cloak 这个指令保持在元素上直到关联组件实例结束编译
- v-once
- v-is (3.1 废弃) 对于某些 HTML 元素只能出现在固定位置的解析规则的重定义

  > 注意：本节仅影响直接在页面的 HTML 中写入 Vue 模板的情况

  ```html
  <!-- vue 2.x -->
  <!-- 这样做是有必要的，因为 `<my-component-row>` 放在一个 -->
  <!-- `<table>` 内可能无效且被放置到外面 -->
  <table>
    <tr is="my-component-row"></tr>
  </table>
  <!-- vue 3.x -->
  <table>
    <tr v-is="my-component-row"></tr>
    <!-- 不正确，不会渲染任何内容 -->
  </table>
  <table>
    <tr v-is="'my-component-row'"></tr>
    <!-- 正确 -->
  </table>
  ```

## 特殊指令

- key
- ref

  ```javascript
    <template>
      <div ref="root">This is a root element</div>
      <div v-for="(item, i) in list" :ref="el => { if (el) divs[i] = el }">
        {{ item }}
      </div
    </template>
    <script>
      import { ref, onMounted } from 'vue';

      export default {
        setup() {
          const root = ref(null);
          const divs = ref([]);
          // 确保在每次更新之前重置ref
          onBeforeUpdate(() => {
            divs.value = [];
          });

          onMounted(() => {
            // DOM元素将在初始渲染后分配给ref
            console.log(root.value); // <div>这是根元素</div>
          });
          return { root, divs }
        }
      }
    </script>
  ```

- is

## 内置组件

- component
  - is 渲染一个元组件为动态组件
- transition

  - props
    - name
    - appear
    - persisted
    - css
    - type
    - mode
    - duration
    - enter/leave-from-class
    - appear-class
    - enter/leave/appear-to-class
    - enter/leave/appear-active-class
  - events
    - before-enter/leave
    - enter/leave/appear
    - after-enter/leave/appear
    - enter/leave/appear-cancelled

  ```javascript
  <!-- 单个元素 -->
  <transition>
    <div v-if="ok">toggled content</div>
  </transition>

  <!-- 动态组件 -->
  <transition name="fade" mode="out-in" appear>
    <component :is="view"></component>
  </transition>

  <!-- 事件钩子 -->
  <div id="transition-demo">
    <transition @after-enter="transitionComplete">
      <div v-show="ok">toggled content</div>
    </transition>
  </div>
  ```

- transition-group
  - props
    - tag
    - move-class
  - events 和 transition 相同
- keep-alive

  - props
    - {String|RegExp|Array} include 哪些组件实例可以被缓存
    - {String|RegExp|Array} exclude 哪些组件实例不被缓存
    - {String|Number} max 最多可以缓存多少组件实例

  ```javascript
  <!-- 逗号分隔字符串 -->
  <keep-alive include="a,b">
    <component :is="view"></component>
  </keep-alive>

  <!-- regex (使用 `v-bind`) -->
  <keep-alive :include="/a|b/">
    <component :is="view"></component>
  </keep-alive>

  <!-- Array (使用 `v-bind`) -->
  <keep-alive :include="['a', 'b']">
    <component :is="view"></component>
  </keep-alive>
  ```

- slot

  - props
    - {String} name 插槽命名

- teleport 移动实际 DOM 节点(非销毁重建),并保持任何组件实例的活动状态

  - props
    - {String} to 必须是有效的查询选择器或者 HTMLElement(如果在浏览器环境中时), 指定将在其中移动 &lt;teleport&gt; 内容的目标元素
    - {Boolean} disabled 可选属性可用于禁用 &lt;teleport&gt; 的功能

  ```javascript
  <!-- 正确 -->
  <teleport to="#some-id" />
  <teleport to=".some-class" />
  <teleport to="[data-teleport]" />

  <!-- 错误 -->
  <teleport to="h1" />
  <teleport to="some-string" />
  ```

## 响应性 API

### 响应性基础 API

- reactive 返回对象的响应式副本

  > 当将 ref 分配给 reactive property 时，ref 将被自动解构

  ```javascript
  const count = ref(1);
  const obj = reactive({});
  obj.count = count;
  console.log(obj.count); // 1
  console.log(obj.count === count.value); // true
  // 它会更新 `obj.value`
  count.value++;
  console.log(count.value); // 2
  console.log(obj.count); // 2
  // 它也会更新 `count` ref
  obj.count++;
  console.log(obj.count); // 3
  console.log(count.value); // 3
  ```

- readonly 接受一个对象 (响应式或纯对象) 或 ref 并返回原始对象的只读代理, 任何被访问的嵌套 property 也是只读的

  ```javascript
  const original = reactive({ count: 0 });
  const copy = readonly(original);
  watchEffect(() => {
    console.log(copy.count); // 用于响应性追踪
  });
  original.count++; // 变更 original 会触发依赖于副本的侦听器
  copy.count++; // 警告! // 变更副本将失败并导致警告
  ```

- isProxy 检查对象是否是由 reactive 或 readonly 创建的 proxy
- isReactive 检查对象是否是由 reactive 创建的响应式代理, 如果代理是由 readonly 创建的并包含了由 reactive 创建的另一个代理, 同样返回 true
- isReadonly 检查对象是否是由 readonly 创建的只读代理
- toRaw 返回 reactive 或 readonly 代理的原始对象

  > 可用于临时读取数据而无需承担代理访问/跟踪的开销，也可用于写入数据而避免触发更改。不建议保留对原始对象的持久引用

  ```javascript
  const foo = {};
  const reactiveFoo = reactive(foo);
  console.log(toRaw(reactiveFoo) === foo); // true
  ```

- markRaw 标记一个对象, 使其永远不会转换为 proxy. 返回对象本身

  - 有些值不应该是响应式的，例如复杂的第三方类实例或 Vue 组件对象。
  - 当渲染具有不可变数据源的大列表时，跳过 proxy 转换可以提高性能

  ```javascript
  const foo = markRaw({});
  console.log(isReactive(reactive(foo))); // false

  // 嵌套在其他响应式对象中时也可以使用
  const bar = reactive({ foo });
  console.log(isReactive(bar.foo)); // false
  ```

- shallowReactive 创建一个响应式代理,它跟踪其自身 property 的响应性,但不执行嵌套对象的深层响应式转换 (暴露原始值)

  ```javascript
  const state = shallowReactive({
    foo: 1,
    nested: {
      bar: 2,
    },
  });
  state.foo++; // 改变 state 本身的性质是响应式的
  // ...但是不转换嵌套对象
  isReactive(state.nested); // false
  state.nested.bar++; // 非响应式
  ```

- shallowReadonly 创建一个 proxy, 使其自身的 property 为只读, 但不执行嵌套对象的深度只读转换 (暴露原始值)

  ```javascript
  const state = shallowReadonly({
    foo: 1,
    nested: {
      bar: 2,
    },
  });
  state.foo++; // 改变 state 本身的 property 将失败
  // ...但适用于嵌套对象
  isReadonly(state.nested); // false
  state.nested.bar++; // 适用
  ```

### Refs

- ref 接受一个内部值并返回一个响应式且可变的 ref 对象. ref 对象具有指向内部值的单个属性 .value

  ```javascript
  const count = ref(0);
  console.log(count.value); // 0

  count.value++;
  console.log(count.value); // 1
  ```

- unref 如果参数是一个 ref, 则返回内部值, 否则返回参数本身
  > val = isRef(val) ? val.value : val 的语法糖函数
- toRef 为源响应式对象上的某个 property 新创建一个 ref, 然后 ref 可以被传递,它会保持对其源 property 的响应式连接

  ```javascript
  const state = reactive({
    foo: 1,
    bar: 2,
  });
  const fooRef = toRef(state, 'foo');
  fooRef.value++;
  console.log(state.foo); // 2

  state.foo++;
  console.log(fooRef.value); // 3
  ```

- toRefs 将响应式对象转换为普通对象, 其中结果对象的每个 property 都是指向原始对象相应 property 的 ref

  > 方便消费组件可以在不丢失响应性的情况下对返回的对象进行分解/扩散

  ```javascript
  const state = reactive({
    foo: 1,
    bar: 2,
  });
  const stateAsRefs = toRefs(state);
  /*
  stateAsRefs 的类型:
  {
    foo: Ref<number>,
    bar: Ref<number>
  }
  */

  // ref 和原始 property 已经“链接”起来了
  state.foo++;
  console.log(stateAsRefs.foo.value); // 2

  stateAsRefs.foo.value++;
  console.log(state.foo); // 3
  ```

- isRef 检查值是否为一个 ref 对象
- customRef 创建一个自定义的 ref, 并对其依赖项跟踪和更新触发进行显式控制

  > 它需要一个工厂函数，该函数接收 track 和 trigger 函数作为参数，并且应该返回一个带有 get 和 set 的对象

  ```javascript
  <input v-model='text' />;
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
  export default {
    setup() {
      return {
        text: useDebouncedRef('hello'),
      };
    },
  };
  ```

- shallowRef 创建一个跟踪自身 .value 变化的 ref, 但不会使其值也变成响应式的

  ```javascript
  const foo = shallowRef({});
  // 改变 ref 的值是响应式的
  foo.value = {};
  // 但是这个值不会被转换。
  isReactive(foo.value); // false
  ```

- triggerRef 手动执行与 shallowRef 关联的任何副作用

### Computed | watch

- computed 接受一个 getter 函数,并为从 getter 返回的值返回一个不变的响应式 ref 对象

  ```javascript
  // 1.
  const count = ref(1);
  const plusOne = computed(() => count.value + 1);
  console.log(plusOne.value); // 2
  plusOne.value++; // 错误
  // 2.
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

- watchEffect 在响应式地跟踪其依赖项时立即运行一个函数, 并在更改依赖项时重新运行它

  - 停止侦听器

    ```javascript
    const stop = watchEffect(() => {});
    stop();
    ```

  - 清除副作用时回调 onInvalidate

    - 副作用即将重新执行时
    - 侦听器被停止(setup 或 lifeCycle Hooks 中使用过, 则在组件卸载时)

    ```javascript
    watchEffect((onInvalidate) => {
      const token = performAsyncOperation(id.value);
      onInvalidate(() => {
        // id has changed or watcher is stopped.
        // invalidate previously pending async operation
        token.cancel();
      });
    });
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
    watchEffect(() => {}, { onTrack(e) {}, onTrigger(e) {} });
    ```

  ```javascript
  const count = ref(0);
  watchEffect(() => console.log(count.value));
  // -> logs 0
  setTimeout(() => {
    count.value++;
    // -> logs 1
  }, 100);
  ```

- watch 使用方式和 this.$watch 和 watch 选项完全等效

  - 与 watchEffect 的区别
    - 惰性地执行副作用
    - 更具体地说明应触发侦听器重新运行的状态
    - 访问被侦听状态的先前值和当前值
  - 侦听单一源

    ```javascript
    // 侦听一个 getter
    const state = reactive({ count: 0 });
    watch(
      () => state.count,
      (count, prevCount) => {
        /* ... */
      }
    );
    // 直接侦听一个 ref
    const count = ref(0);
    watch(count, (count, prevCount) => {
      /* ... */
    });
    ```

  - 侦听多个源

    ```javascript
    watch([fooRef, barRef], ([foo, bar], [prevFoo, prevBar]) => {
      /* ... */
    });
    ```

## 组合式 API

- setup 组件选项，在创建组件之前执行

  - 参数
    - {Data} props
    - {SetupContext} context
      - attrs
      - slots
      - emit

  ```javascript
  import { defineComponent } from 'vue';
  const MyComponent = defineComponent({
    setup(props, { attrs, slots, emit }) {
      return {};
    },
  });
  ```

  - attrs 和 slots 是非响应式的, 如果需要根据 attrs 或者 slots 更改应用的的副作用, 需要在 onUpdated 钩子中执行此操作

- 生命周期钩子

  ```javascript
  import { defineComponent, onBeforeMount, onMounted, onBeforeUnmount, onUnmounted } from 'vue';
  const MyComponent = defineComponent({
    setup(props, { attrs, slots, emit }) {
      onBeforeMount(() => {});
      onMounted(() => {});
      onBeforeUnmount(() => {});
      onUnmounted(() => {});
      return {};
    },
  });
  ```

- Provide / Inject
- getCurrentInstance
  - 支持访问内部组件实例，用于高阶用法或库的开发
  - 只能在 setup 或生命周期钩子中调用

## Migration

### v-for 中的 Ref 数组

- Vue 3.x 中 不再在 $ref 中自动创建数组填充相应的 $refs property

### 异步组件

- 新的 defineAsyncComponent 助手方法，用于显式地定义异步组件
- component 选项重命名为 loader
- Loader 函数本身不再接收 resolve 和 reject 参数，且必须返回一个 Promise

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

### attribute 强制行为

- Vue 3.x 中, 移除 attribute 使用 null 或者 undefined
- 非布尔 attribute, 如果 attribute 为 false, 将停止删除它们, 相反强制它们为 'false'

### $attrs 包括 class 和 style

- $attrs 中的 attribute 不再自动添加到根元素中,而是由开发者决定

### $children 移除

### 自定义指令

- 特定 is prop 用法仅限于保留的 component 标签
- 新增 v-is 指令来处理原生 HTML 解析限制

```javascript
// Vue 2.x
Vue.directive('my-directive', {
  bind() {}, // - 指令绑定到元素后发生。只发生一次。
  inserted() {}, // - 元素插入父 DOM 后发生。
  update() {}, // - 当元素更新，但子元素尚未更新时，将调用此钩子。
  componentUpdated() {}, // - 一旦组件和子级被更新，就会调用这个钩子。
  unbind() {}, // - 一旦指令被移除，就会调用这个钩子。也只调用一次
});
// Vue 3.x
const app = createApp({});
app.directive('my-directive', {
  created() {},
  beforeMount() {},
  mounted() {},
  beforeUpdate() {},
  updated() {},
  beforeUnmount() {},
  unmounted() {},
});
```

### Data 选项

> 标准化只接受返回的 Object 的 function

### emits

### 事件 API

> $on，$off 和 $once 实例方法已被移除，应用实例不再实现事件触发接口

### 过滤器 移除

### 片段

- 组件可以包含多个根节点, 需要显示定义 attribute 的位置

  ```javascript
  <template>
    <header>...</header>
    <main v-bind='$attrs'>...</main>
    <footer>...</footer>
  </template>
  ```

### 函数式组件

- 单文件组件的 functional attribute 被移除
- { functional: true} 选项通过函数创建组件不推荐使用,和状态组件的性能差别不大

  ```javascript
  import { h } from 'vue';
  const DynamicHeading = (props, { attrs, slots, emit }) => {
    return h(`h${props.level}`, attrs, slots);
  };
  DynamicHeading.props = ['level'];
  export default DynamicHeading;
  ```

### 全局 API 变更

- 测试期间,全局配置很容易意外地污染其他测试用例, Vue.use | Vue.mixin 无恢复效果的方法
- 全局配置使同一页面的多个 app 之间共享同一个 Vue 副本非常困难

  ```javascript
  import { createApp } from 'vue';
  import Foo from './Foo.vue';
  import Bar from './Bar.vue';

  const createMyApp = (options) => {
    const app = createApp(options);
    app.directive('focus' /* ... */);

    return app;
  };

  createMyApp(Foo).mount('#foo');
  createMyApp(Bar).mount('#bar');
  ```

### 全局 API Treeshaking

- 全局 API 现在只能作为 ES 模块构建的命名导出进行访问, Vue 应用程序将从未使用的 api 从最终捆绑包中消除,从而获得最佳的文件大小

  ```javascript
  <transition>
    <div v-show='ok'>hello</div>
  </transition>;
  import { h, Transition, withDirectives, vShow } from 'vue';
  export function render() {
    return h(Transition, [withDirectives(h('div', 'hello'), [[vShow, this.ok]])]);
  }
  ```

### 内联模板 Attribute 移除

```javascript
<my-component inline-template>
  <div>
    <p>它们被编译为组件自己的模板</p>
    <p>不是父级所包含的内容。</p>
  </div>
</my-component>
```

### key attribute

- 条件判断分支项不再必须 key

### 按键修饰符

- 不再支持使用数字(即键码)作为 v-on 修饰符
- 不再支持 config.keyCodes

### $listeners 移除

- 事件监听器作为 $attrs 的一部分

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

### 渲染函数 API

- h 全局导入, 不再作为参数传递给渲染函数

  ```javascript
  // Vue 2.x
  export default {
    render(h) {
      return h();
    },
  };
  // Vue 3.x
  import { h } from 'vue';
  export default {
    render() {
      return h();
    },
  };
  ```

- 渲染函数参数更改以在有状态组件和函数组件之间更加一致
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

- 注册组件, 需要借助 resolveComponent API

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

  ```javascript
  // 解构并且重命名
  <todo-list v-slot='{ item: todo }'>
    <i class='fas fa-check'></i>
    <span class='green'>{{ todo }}</span>
  </todo-list>
  ```

- 动态插槽名

  ```javascript
  <base-layout>
    <template v-slot:[dynamicSlotName]>
      ...
    </template>
  </base-layout>
  ```

### 过渡的 class 名更改

- 过渡类名 v-enter 修改为 v-enter-from
- 过渡类名 v-leave 修改为 v-leave-from

### Transition Group 根元素

> 不再默认渲染根元素, 仍然可以用 tag prop 创建根元素

### v-on.native 修饰符 移除

新增 emits 选项允许子组件定义真正会被触发的事件

### v-model

- v-bind 的 .sync 修饰符和组件的 model 选项被移除, 使用 v-model 代替
- 同一组件上可以使用多个 v-model 进行双向绑定
- 可自定义 v-model 修饰符
- 自定义组件时 v-model prop 和事件默认名称已更改

  - prop: value -> modelValue
  - event: input -> update:modelValue

  - Vue 2.0 v-model 只能使用 value 作为 prop, 并接收抛出的 input 事件, 如果使用其他 prop, 必须使用 v-bind.sync 同步

    ```javascript
    <my-component :value="pageTitle" @input="pageTitle = $event" />
    // 简写方式
    <my-component v-model="pageTitle" />
    ```

  - Vue 2.2 增加组件选项 model, 允许自定义 v-model 的 prop 和事件,只能在组件上使用一个 model

    ```javascript
    <my-component :value="pageTitle" @input="pageTitle = $event" />
    // 简写方式
    <my-component v-model="pageTitle" />
    export default {
      model: {
        prop: 'title',
        event: 'change'
      },
      props: {
        // 这将允许 `value` 属性用于其他用途
        value: String,
        // 使用 `title` 代替 `value` 作为 model 的 prop
        title: {
          type: String,
          default: 'Default title'
        }
      }
    }
    ```

  - .sync 修饰符 2.3.0 新增

    ```javascript
    <my-component :title="pageTitle" @update:title="pageTitle = $event" />
    // 简写方式
    <my-component :title.sync="pageTitle" />
    ```

  - Vue 3.x v-model 传递 modelValue prop 并接收抛出的 update:modelValue 事件

    ```javascript
    <my-component :modelValue="pageTitle" @update:modelValue="pageTitle = $event"/>
    // 简写方式
    <my-component v-model="pageTitle" />

    // 修改名称，使用多个 v-model
    <my-component :title="pageTitle" @update:title="pageTitle = $event" :content="pageContent" @update:content="pageContent = $event"/>
    // 简写方式
    <my-component v-model:title="pageTitle" v-model:content="pageContent"/>
    ```

  - migration

    - 所有子组件 .sync 修饰符的替换为 v-model
    - 未带参数的 v-model, 修改子组件的 prop 和 event 命令为 modelValue 和 update:modelValue

    ```javascript
    <my-component :title.sync="pageTitle" />
    // 替换为
    <my-component v-model:title="pageTitle" />

    // 未带参数的 v-model
    <ChildComponent v-model="pageTitle" />
    export default {
      props: {
        modelValue: String // 以前是`value：String`
      },
      emits: ['update:modelValue'],
      methods: {
        changePageTitle(title) {
          this.$emit('update:modelValue', title) // 以前是 `this.$emit('input', title)`
        }
      }
    }
    ```

- 处理 v-model 修饰符

  - 不带参数: Vue 3.x 通过 modelModifiers 提供给 prop
  - 带参数: 生成的 prop 名称将为 arg + 'Modifiers'

  ```javascript
  // 不带参数的 v-model
  <my-component v-model.capitalize="myText" />
  app.component('my-component', {
    props: {
      modelValue: String,
      modelModifiers: {
        default: () => ({})
      }
    },
    emits: ['update:modelValue'],
    template: `
      <input type="text"
        :value="modelValue"
        @input="$emit('update:modelValue', $event.target.value)">
    `,
    created() {
      console.log(this.modelModifiers) // { capitalize: true }
    }
  });
  // 带参 v-model
  <my-component v-model:description.capitalize="myText" />
  app.component('my-component', {
    props: ['description', 'descriptionModifiers'],
    emits: ['update:description'],
    template: `
      <input type="text"
        :value="description"
        @input="$emit('update:description', $event.target.value)">
    `,
    created() {
      console.log(this.descriptionModifiers) // { capitalize: true }
    }
  })
  ```

### v-if 与 v-for 的优先级对比

> 两者作用于同一个元素上时，v-if 会拥有比 v-for 更高的优先级

### v-bind 合并行为

> 声明绑定的顺序决定了合并顺序

### Watch on Arrays

> 当侦听一个数组时，只有当数组被替换时才会触发回调。如果你需要在数组改变时触发回调，必须指定 deep 选项
