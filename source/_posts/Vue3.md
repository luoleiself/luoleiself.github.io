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

- createApp() 创建一个应用实例

  - 参数
    - {Object} rootComponent 根组件选项
    - {Object} rootProps 传递给根组件的 props

  ```javascript
  import { createApp } from 'vue';
  const app = createApp({
    /* root component options */
  });
  ```

- createSSRApp() 以 `SSR` 模式创建一个应用实例, 用法和 `createApp()` 相同

- app.mount() 将应用实例挂载到一个容器元素中
  - 参数可以是一个实际的 DOM 元素或一个 CSS 选择器, 返回根组件实例
  - 如果该组件有 `template` 模板或定义了 `render` 函数, 则替换容器内所有现存的 DOM 节点, 否则使用容器元素的 innerHTML 作为模板
- app.unmount() 卸载一个已挂载的应用实例, 同时触发该应用组件树内所有组件的卸载生命周期钩子
- app.provide() 提供一个值, 可以在应用中的所有后代组件中注入使用

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

- app.component() 注册或查找全局组件, 根据参数个数区分

  ```javascript
  import { createApp } from 'vue';
  const app = createApp(/* */);
  app.component('my-component', {
    /*组件配置项*/
  });

  app.component('my-component'); // 查找已注册的组件
  ```

- app.directive() 注册或查找全局指令, 根据参数个数区分 <em id="directive"></em> <!-- markdownlint-disable-line -->

  - 钩子函数参数
    - el 指令绑定的元素, 可用于直接操作 DOM
    - binding 一个对象
      - value 传递给指令的值, 例如 `v-my-directive="1 + 1"` 的值为 2
      - oldValue 之前的值, 仅在 [`beforeUpdate`](#onBeforeUpdate) 和 [`updated`](#onUpdated) 中可用
      - arg 传递给指令的参数, 例如 `v-my-directive:foo` 的参数为 `foo`
      - modifiers 一个包含修饰符的对象, 例如 `v-my-directive.foo.bar` 的修饰符对象为 `{foo: true, bar: true}`
      - instance 使用该指令的组件实例
      - dir 指令的定义对象
    - vnode 代表绑定元素的底层 VNode
    - prevVnode 之前的渲染中代表指令所绑定元素的 VNode, 仅在 [`beforeUpdate`](#onBeforeUpdate) 和 [`updated`](#onUpdated) 中可用

  ```html
  <template>
    <MyComponent v-my-directive="test" />
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

  <!-- more -->

- app.use() 安装一个 **插件**, 插件可以是一个包含 `install()` 方法的对象或者是一个安装函数本身

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

- app.mixin() 应用一个全局的 mixin, 作用于应用中的每个组件实例 (不推荐使用)
- app.version 提供当前应用所使用的 Vue 版本号, 插件中可根据此执行不同的逻辑
- app.config 应用实例暴露出的一个 `config` 对象, 其中包含了对此应用实例的配置

  - app.config.errorHandler 用于为应用实例内抛出的未捕获错误指定一个全局处理函数
  - app.config.warnHandler 用于为 Vue 的运行时警告指定一个自定义处理函数
  - app.config.performance 设置为 `true` 可在浏览器工具的 **性能/时间线** 页启用对组件初始化、编译、渲染和修改的性能表现追踪
  - app.config.compilerOptions 配置 **运行时编译器** 的选项
    - app.config.compilerOptions.isCustomElement 用于指定一个检查方法来识别原生自定义元素
    - app.config.compilerOptions.whitespace 用于调整模板中空格的处理行为 `condense(default) | preserve`
    - app.config.compilerOptions.delimiters 用于调整模板内文本插值的分隔符, 默认 ['{{', '}}']
    - app.config.compilerOptions.comments 用于调整是否移除模板中的 HTML 注释
  - app.config.globalProperties 用于注册能够被应用实例内所有组件实例访问到的全局属性的对象
  - app.config.optionMergeStrategies 用于定义自定义组件选项的合并策略的对象

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

- version 暴露当前所使用的 Vue 的版本号
- nextTick() 等待下一次 DOM 更新刷新的工具方法, 可以在状态改变后立即使用以等待 DOM 更新完成 <em id="nextTick"></em> <!-- markdownlint-disable-line -->

  - 传递一个回调函数作为参数
  - 或者 await 返回的 Promise

```javascript
import { Version, nextTick } from 'vue';

console.log(Version); // 打印当前使用的 Vue 版本
async function increment() {
  console.log('DOM 还未更新');
  await nextTick();
  console.log('DOM 已更新');
}
```

- defineComponent() 创建一个合成类型的构造函数, 用于手动渲染函数、TSX 和 IDE 工具支持 <em id="defineComponent"></em> <!-- markdownlint-disable-line -->

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

  - 参数为 setup 函数, 函数名称作为组件名称使用

    ```javascript
    import { createApp, defineComponent, ref } from 'vue';

    const HelloWorld = defineComponent((props, ctx) => {
      const count = ref(0);
      return { count };
    });
    const app = createApp(HelloWorld).mount('#app');
    ```

- defineAsyncComponent() 创建一个只有在需要时才会加载的异步组件

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

- [defineCustomElement()](#defineComponent) 和 `defineComponent()` 接收的参数相同, 不同的是返回一个原生 **自定义元素** 类的构造器

## 组合式 API <em id="combinedapi"></em> <!-- markdownlint-disable-line -->

### setup()

> 对于结合单文件组件使用的组合式 API 推荐使用 `<script setup>` 语法

#### 基本使用

- 需要在非单文件组件中使用组合式 API 时
- 需要在基于选项式 API 的组件集成基于组合式 API 的代码时

- 在创建组件实例时, 在初始 prop 解析之后立即调用 setup
- 在生命周期方面, 在 beforeCreate 钩子之前调用

- getCurrentInstance
  - 支持访问内部组件实例，用于高阶用法或库的开发
  - 只能在 setup 或生命周期钩子中调用

#### 访问 Props

- setup 函数地第一个参数, 是响应式地并且会在传入新的 props 时同步更新
- 不能直接对 props 进行解构操作, 会丢失响应性, 可以通过 `toRefs()` 和 `toRef()` 工具函数辅助完成

#### 上下文

setup 函数的第二个参数, 暴露了其他一些在 setup 中可能会用到的值, 该上下文对象是非响应式的, 可以安全地解构, attrs 和 slots 是非响应式的, 如果需要根据 attrs 或者 slots 的改变执行副作用, 需要在 onBeforeUpdate 钩子中执行相关逻辑

- attrs 透传 Attributes, 等价于 $attrs
- slots 插槽, 等价于 $slots
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

#### 返回渲染函数

- setup 应该同步地返回一个对象, 唯一可以使用 `async setup()` 的情况是该组件是 [&lt;Suspense&gt;](#suspense) 组件地后裔
- 也可以返回一个 **渲染函数**, 此时在渲染函数中可以直接使用在同一作用域下声明的响应式状态

```javascript
import { h, ref, reactive } from 'vue';
const app = createApp({
  setup() {
    const count = ref(0);
    const object = reactive({ foo: 'bar' });
    return () => h('div', [count.value, object.foo]);
  },
});
```

### 响应式: 核心

#### ref() <em id="ref"></em> <!-- markdownlint-disable-line -->

接受一个内部值, 返回一个响应式可更改的 ref 对象, 此对象只有一个指向其内部值的属性 `.value`

- 如果将一个对象赋值给 ref, 那么这个对象将通过 `reactive()` 转为具有深层次响应式的对象, 如果对象中包含了嵌套的 ref, 它们将被深层地解包

```javascript
const count = ref(0);
console.log(count.value); // 0

count.value++;
console.log(count.value); // 1
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

- 仅对对象类型有效(对象、数组、Map、Set 这样的集合类型), 而对 `string`, `number`, `boolean` 这样的原始类型无效
- 因为 Vue 的响应式系统是通过属性访问进行追踪的, 因此需要始终保持对响应式对象的相同引用, 将响应式对象的属性赋值或解构至本地变量时、或是将该属性传入一个函数时将失去响应性
- 对同一个原始对象调用 `reactive()` 总是返回同样的代理对象, 对一个已存在的代理对象调用 `reactive()` 总是返回其本身

- 将一个 ref 赋值给一个 reactive 属性时, 该 ref 会自动解包

```javascript
const count = ref(1);
const obj = reactive({});
obj.count = count;
console.log(obj.count); // 1
// ref 会解包
console.log(obj.count === count.value); // true
// 自动更新 `obj.value`
count.value++;
console.log(count.value); // 2
console.log(obj.count); // 2
// 自动更新 `count` ref
obj.count++;
console.log(obj.count); // 3
console.log(count.value); // 3
```

- 当访问到某个响应式 **数组** 或 `Map` 这样的原生集合类型中的 ref 元素时, 不会执行 ref 的解包

```javascript
// 原生集合中包含 ref 元素时, ref 不会解包
const books = reactive([ref('Vue 3.0')]);
console.log(books[0].value); // 需要使用 .value

const map = reactive(new Map([['count', ref(0)]]));
console.log(map.get('count').value); // 需要使用 .value
```

#### readonly() <em id="readonly"></em> <!-- markdownlint-disable-line -->

接受一个对象(响应式或普通)或一个 ref, 返回原值的只读代理, 任何被访问的嵌套属性也是只读的, 它的 ref 解包行为与 reactive() 相同, 但解包得到的值是只读的

```javascript
import { reactive, readonly, watchEffect } from 'vue';
const original = reactive({ count: 0 });
const copy = readonly(original);
watchEffect(() => {
  console.log(copy.count); // 用于响应性追踪
});
original.count++; // 变更 original 会触发依赖于副本的侦听器
copy.count++; // 警告! // 变更副本将失败并导致警告
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
- 默认是懒侦听的, 仅在侦听源发生变化时才执行回调函数

##### 参数

- 第一个参数是侦听器的源, 支持包含返回值的函数、ref、响应式对象、或者以上类型的值组成的数组
- 第二个参数是侦听源发生变化时调用的函数, 函数接收三个参数: 新值、旧值，及一个用于注册副作用清理的回调函数
- 第三个参数是一个配置项对象

  - immediate 在侦听器创建时立即触发回调, 第一次调用时旧值为 `undefined`
  - deep 如果源是对象, 强制深度遍历, 以便在深层级变更时触发回调
  - flush 调整回调函数的刷新时机, 见 [watchEffect()](#watchEffect)
  - onTrack/onTrigger 调试侦听器的依赖, 见 [watchEffect()](#watchEffect)

```javascript
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

基于响应式对象上的一个属性, 新创建一个对应的 ref, 此 ref 与其源属性保持同步, 改变源属性的值将更新 ref 的值, 反之亦然

```javascript
const state = reactive({ foo: 1, bar: 2 });
const fooRef = toRef(state, 'foo');
// 更改 ref 会更新源属性
fooRef.value++;
console.log(state.foo); // 2
// 更改源属性会更新 ref
state.foo++;
console.log(fooRef.value); // 3
```

#### toRefs()

> 方便消费组件可以在不丢失响应性的情况下对返回的对象进行分解/扩散

将一个响应式对象转换为一个普通对象, 这个普通对象的每个属性都指向源对象相应属性的 ref, 每个单独的 ref 都是使用 `toRef()` 创建的

- toRefs 在调用时智慧为源对象上的可以枚举的属性创建 ref, 如果为可能还不存在的属性创建 ref 时, 使用 toRef

```javascript
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
const state = shallowRef({ count: 1 });
// 不会触发更改
state.value.count = 2;

// 会触发更新
state.value = { count: 2 };
```

#### triggerRef()

强制触发依赖一个 `浅层 ref` 的副作用, 通常对浅引用的内部值进行深度变更后使用

```javascript
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

#### shallowReactive()

[reactive()](#reactive) 的浅层作用形式

- 没有深层级的转换, 浅层响应式对象里只有根级别的属性是响应式的
- 属性的值会被原样存储和暴露, 值为 ref 的属性不会自动解包

```javascript
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
const foo = markRaw({});
console.log(isReactive(reactive(foo))); // false

// 嵌套在其他响应式对象中时也可以使用
const bar = reactive({ foo });
console.log(isReactive(bar.foo)); // false
```

#### effectScope() <em id='effect'></em> <!-- markdownlint-disable-line -->

创建一个 effect 作用域, 可以捕获其中所创建的响应式副作用(计算属性和侦听器), 这样捕获到的副作用可以一起处理

```javascript
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

#### [getCurrentScope()](#effect)

如果存在则返回当前活跃的 effect 作用域

#### onScopeDispose()

> 此方法可以作为可复用的组合式函数中 `onUnmounted` 的替代品, 它并不与组件耦合, 因为每个 Vue 组件的 setup 函数也是在一个 effect 作用域中调用的

在当前活跃的 effect 作用域上注册一个处理回调函数, 当相关的 effect 作用域停止时会调用注册的回调函数

### 生命周期钩子

> 所有生命周期钩子函数必须在组件的 `setup()` 阶段同步调用

#### onBeforeMount() <em id="onBeforeMount"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件 **挂载之前** 调用, 组件已经完成其响应式状态的设置, 但还没有创建 DOM 节点

#### onMounted() <em id="onMounted"></em> <!-- markdownlint-disable-line -->

> 钩子函数在服务器端渲染期间不会被调用

注册一个回调函数在组件 **挂载完成** 之后执行

- 其所有同步子组件都已经被挂载(不包含 **异步组件** 或 [&lt;Suspense&gt;](#suspense) 树内的组件)
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

- 第一个参数为注入的 key, 通过匹配最近的组件提供的值, 否则将返回 undefined
- 第二个参数可选, 即在没有匹配到 key 时使用的默认值,

  - 如果为一个工厂函数, 则用来返回某些创建复杂的值
  - 如果默认值本身是一个函数, 则需要将 false 作为第三个参数传入, 表明这个函数就是默认值而不是工厂函数

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
  const fn = inject('fn', () => {}, false);
</script>
```

## 选项式 API

### 状态选项

#### data

> 以 \_ 和 $ 开头的属性不会被组件实例代理, 因为它们可能和 Vue 的内置属性, API 方法冲突

用于声明组件初始响应式状态的函数

#### props <em id="props"></em> <!-- markdownlint-disable-line -->

用于声明组件的 props

- 使用字符串数组的简易形式
- 使用对象的完整形式, 可以对单个 prop 进行更详细的配置

  - type 定义 prop 的类型, 可以为原生构造函数之一
  - default 为该 prop 指定一个当其没有被传入值或值为 undefined 时的默认值
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

> 组合式 API 中的 `setup()` 钩子函数会在所有选项式 API 钩子之前调用

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

一个包含组件选项对象的数组, 这些选项都将被混入到当前组件的实例中

#### extends

> `extends` 和 `mixin` 实现上几乎相同, 但是表达的目标不同, `mixins` 选项基本用于组合功能, `extends` 一般更关注继承关系

要继承的 **基类** 组件, 同 `mixins` 一样, 所有选项都将使用相关的策略进行合并

### 其他杂项

#### name

用于显式声明组件展示时的名称, 使用 name 选项可以覆盖推导出的名称, 或是在没有推导出名字是显式提供一个

- 在组件自己的模板中递归引用自己时
- 在 Vue 开发者工具中的组件树显示时
- 在组件抛出的警告追踪栈信息中显示时

#### inheritAttrs <em id="inheritAttrs"></em> <!-- markdownlint-disable-line -->

> 默认情况下, 父组件传递的没有被子组件解析为 `props` 的 `attributes` 绑定会被透传

用于控制是否启用默认的组件 `attribute` 透传行为, 默认为 true

- 使用 [&lt;script setup&gt;](#scriptsetup) 的组合式 API 中声明这个选项时, 需要一个额外的 `<script>` 块

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

#### $slots

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

### 指令

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

##### 事件修饰符

- .stop 调用 `event.stopPropagation()`
- .prevent 调用 `event.preventDefault()`
- .capture 在捕获模式添加事件监听器
- .self 只有事件从元素本身发出才触发处理函数
- .{keyAlias} 只有在某些按键下触发处理函数
- .once 最多触发一次处理函数
- .left 只在鼠标左键事件触发处理函数
- .right 只在鼠标右键事件触发处理函数
- .middle 只在鼠标中键事件触发处理函数
- .passive 通过 `{passive: true}` 附加一个 DOM 事件

#### v-bind <em id="v-bind"></em> <!-- markdownlint-disable-line -->

动态的绑定一个或多个 attribute, 也可以是组件的 prop, 缩写 `:` 或 `.`(当使用 `.prop` 修饰符)

##### 绑定修饰符

- .camel 将短横线命名的 attribute 转变为驼峰式命名
- .prop 强制绑定为 DOM property
- .attr 强制绑定为 DOM attribute

```html
<template>
  <div :someProperty.prop="someObject"></div>
  <!-- 等价于 -->
  <div .someProperty="someObject"></div>
</template>
```

#### v-model

在表单输入元素或组件上创建双向绑定

##### 修饰符

- .lazy 监听 change 事件而不是 input 事件
- .number 将输入的合法字符换转为数字
- .trim 移除输入内容两端空格

##### 版本迭代

- [`v-bind`](#v-bind) 的 .sync 修饰符和组件的 model 选项被移除, 使用 v-model 代替
- 同一组件上可以使用多个 v-model 进行双向绑定
- 可自定义 v-model 修饰符
- 自定义组件时 `v-model` 的 `prop` 和 `event` 默认名称已更改

  - prop: `value` -> `modelValue`
  - event: `input` -> `update:modelValue`

###### migration

- 所有子组件 `.sync` 修饰符的替换为 `v-model`
- 未带参数的 `v-model`, 修改子组件的 prop -> `modelValue`, event -> `update:modelValue`

```html
<template>
  <my-component :title.sync="pageTitle" />
  <!-- 替换为 -->
  <my-component v-model:title="pageTitle" />

  <!-- 未带参数的 v-model -->
  <ChildComponent v-model="pageTitle" />
</template>
<script>
  export default {
    props: {
      modelValue: String, // 以前是`value：String`
    },
    emits: ['update:modelValue'],
    methods: {
      changePageTitle(title) {
        this.$emit('update:modelValue', title); // 以前是 `this.$emit('input', title)`
      },
    },
  };
</script>
```

##### V 2.0

- Vue 2.0 `v-model` 只能使用 `value` 作为 prop, 并监听抛出的 `input` 事件, 如果使用其他 prop, 必须使用 `v-bind.sync` 同步

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

##### V 2.2

- Vue 2.2 增加组件选项 `model`, 允许自定义 `v-model` 的 prop 和 event, 只能在组件上使用一个 model

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

##### V 2.3

- Vue 2.3 增加 `.sync` 修饰符

```html
<template>
  <my-component :title="pageTitle" @update:title="pageTitle = $event" />
  <!-- 简写方式 -->
  <my-component :title.sync="pageTitle" />
</template>
```

##### V 3.x

- Vue 3.x `v-model` 默认传递 `modelValue` prop, 并接收子组件抛出的 `update:modelValue` 事件

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

```html
<template>
  <!-- v-model 不带参数  -->
  <my-component v-model.capitalize="myText" />

  <!-- v-model 带参数  -->
  <my-component v-model:description.capitalize="myText" />
</template>
<script setup></script>
<script>
  // v-model 不带参数
  app.component('my-component', {
    props: {
      modelValue: String,
      modelModifiers: {
        default: () => ({}),
      },
    },
    emits: ['update:modelValue'],
    template: `<input type="text" :value="modelValue" @input="emitValue" />`,
    created() {
      console.log(this.modelModifiers); // { capitalize: true }
    },
    methods: {
      emitValue(e) {
        let value = e.target.value;
        if (this.props.modelModifiers.capitalize) {
          value = value.charAt(0).toUpperCase() + value.slice(1);
        }
        this.$emit('update:modelValue', $event.target.value);
      },
    },
  });

  // v-model 带参数
  app.component('my-component', {
    props: ['description', 'descriptionModifiers'],
    emits: ['update:description'],
    template: `
      <input type="text"
        :value="description"
        @input="$emit('update:description', $event.target.value)" />
    `,
    created() {
      console.log(this.descriptionModifiers); // { capitalize: true }
    },
  });
</script>
```

#### v-slot

用于声明具名插槽或是期望接收 props 的作用域插槽, 缩写 `#`

##### 限制使用

> 如果混用了 **具名插槽** 和 **默认插槽**, 则需要为 **默认插槽** 使用显式的 `<template>` 标签, 否则编译错误

- &lt;template&gt;
- components(用于带有 prop 的单个默认插槽)

```html
<template>
  <MyComponent>
    <!-- 使用显式的默认插槽 -->
    <template #default="{ message }">
      <p>{{ message }}</p>
    </template>

    <template #footer>
      <p>Here's some contact info</p>
    </template>
  </MyComponent>

  <!-- 单个默认作用域插槽, 直接使用子组件标签 -->
  <MyComponent v-slot="slotProps">
    {{slotProps.text}} - {{slotProps.message}}
  </MyComponent>
</template>
```

#### v-pre

跳过该元素及其所有子元素的编译

#### v-once

仅渲染元素和组件一次, 并跳过之后的更新

#### v-memo

缓存一个模板的子树, 根据传入的依赖值数组的比较结果控制子树的更新

#### v-cloak

> 该指令只在没有构建步骤的环境下需要使用

用于隐藏尚未完成编译的 DOM 模板

### 组件

> 内置组件无需注册便可以直接在模板中使用，同时也支持 `tree-shaking`; 仅在使用时才会包含在构建中
> 在 **渲染函数** 中使用它们时, 需要显式引入

```javascript
import { h, KeepAlive, Transition } from 'vue';

export default {
  setup(props, ctx) {
    return () => h(Transition, { mode: 'out-in' } /* ... */);
  },
};
```

#### &lt;Transition&gt;

为单个元素或组件提供动画过渡效果

##### Transition props

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

##### Transition 事件

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

#### &lt;TransitionGroup&gt;

为列表中的多个元素或组件提供过渡效果

##### TransitionGroup props

- tag 如果未定义, 则渲染为片段(fragment)
- moveClass 用于自定义过渡期间被应用的 CSS class, 使用 `kebab-case` 格式

##### TransitionGroup 事件

`<TransitionGroup>` 抛出与 `<Transition>` 相同的事件

#### &lt;KeepAlive&gt;

缓存包裹在其中的动态切换组件

##### KeepAlive props

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

#### &lt;Teleport&gt;

移动实际 DOM 节点(非销毁重建),并保持任何组件实例的活动状态

##### Teleport props

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

#### &lt;Suspense&gt; <em id="suspense"></em> <!-- markdownlint-disable-line -->

用于协调对组件树中嵌套的异步依赖的处理

##### Suspense props

- timeout 渲染新内容耗时超时时间

##### Suspense 事件

- @pending 在 suspense 进入挂起状态时触发
- @resolve 在 default 插槽完成获取新内容时触发
- @fallback 在 fallback 插槽的内容显示时触发

##### Suspense 插槽

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

### 特殊元素

> `<component>`, `<slot>`, `<template>` 具有类似组件的特性, 也是模板语法的一部分. 但它们并非真正的组件, 同时在模板编译期间会被编译掉. 因此, 它们通常在模板中使用小写字母

#### &lt;component&gt;

用于渲染动态组件或元素的 `元组件`

##### component props

- is 要渲染的实际组件由 `is` prop 决定
  - 如果是字符串时, 可以是 HTML 标签名或者组件的注册名
  - 或者是直接绑定到组件的定义

#### &lt;slot&gt;

表示模板中的插槽内容出口

##### slot props

- name 指定插槽名, 缺少时将会渲染默认插槽

#### &lt;template&gt;

当使用内置指令而不在 DOM 中渲染元素时, `<template>` 标签可以作为占位符使用

### 特殊 Attributes

#### key

主要作为 Vue 的虚拟 DOM 算法提示, 在比较新旧节点列表时用于识别 vnode

#### ref

用于注册元素或子组件的 `模板引用`

```html
<template>
  <div ref="root">This is a root element</div>
  <div v-for="item in list" :ref="itemRefs">
    {{ item }}
  </div
</template>
<script setup>
  import { ref, onBeforeUpdate, onUpdated, onMounted } from 'vue';

  const root = ref(null);
  const itemRefs = ref([]);
  // 确保在每次更新之前重置 ref
  onBeforeUpdate(() => {
    itemRefs.value = [];
  });

  onUpdated(()=>{
    console.log(itemRefs.value);
  });

  onMounted(() => {
    // DOM元素将在初始渲染后分配给ref
    console.log(root.value); // <div>这是根元素</div>
  });
</script>
```

#### is

用于动态绑定组件

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

- `<template>` 每个 `*.vue` 文件最多可以包含一个顶层 `<template>` 块, 包含的内容将被提取传递给 `@vue/compiler-dom` 编译生成为 **渲染函数**
- `<script>` 每个 `*.vue` 文件最多可以包含一个 `<script>` 块(使用 `<script setup>` 除外), 默认导出是 Vue 的组件选项对象
- `<script setup>` 每个 `*.vue` 文件最多可以包含一个 `<script setup>` 块, 此脚本块将被预处理为组件的 `setup()` 函数
- `<style>` 每个 `*.vue` 文件可以包含多个 `<style>` 块

#### src 导入

```html
<template src="./template.html"></template>
<script src="./script.js"></script>
<style src="./style.css"></style>
```

### &lt;script setup&gt; <em id="scriptsetup"></em> <!-- markdownlint-disable-line -->

> `<script setup>` 是在单文件组件(SFC) 中使用 [组合式 API](#combinedapi) 的编译时语法糖
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

##### 默认 props 值

> `defineProps` 没有可以给 props 提供默认值的方式, 使用 `withDefaults` **编译器宏** 解决

```html
<script setup>
  const props = withDefaults(defineProps(), {
    msg: 'hello world',
    labels: () => ['one', 'two'],
  });
</script>
```

#### defineExpose()

> 使用 `<script setup>` 的组件是 **默认关闭** 暴露任何在 `<script setup>` 中声明的绑定

```html
<script setup>
  import { ref } from 'vue';

  const a = 1;
  const b = ref(0);
  defineExpose({ a, b });
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

- 声明无法在 `<script setup>` 中声明的选项, 例如 [inheritAttrs](#inheritAttrs) 或插件的自定义选项
- 声明模块的具名导出(named exports)
- 运行只需要在模块作用域执行一次的副作用, 或是创建单例对象

#### 顶层 await

> `<script setup>` 中可以使用顶层 await, 结果代码会被编译成 `async setup()`
> `async setup()` 必须与 [&lt;Suspense&gt;](#suspense) 内置组件组合使用

```html
<script setup>
  const post = await fetch('/api/post/1').then(res => res.json())
</script>
```

























## 全局 API

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
        const nodeWithDirectives = withDirectives(h('div'), [
          [MyDirective, 100, 'click', { prevent: true }],
        ]);
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

  ```html
  <script>
    import { h, useCssModule } from 'vue';
    export default {
      setup() {
        const style = useCssModule();
        return () =>
          h(
            'div',
            {
              class: style.success,
            },
            'Task complete!'
          );
      },
    };
  </script>
  <style module>
    .success {
      color: #090;
    }
  </style>
  ```

- version 以字符串形式提供已安装的 Vue 的版本号

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

### 片段

- 组件可以包含多个根节点, 需要显示定义 attribute 的位置

  ```html
  <template>
    <header>...</header>
    <main v-bind="$attrs">...</main>
    <footer>...</footer>
  </template>
  ```

### 函数式组件

- 单文件组件的 functional attribute 被移除
- { functional: true} 选项通过函数创建组件不推荐使用,和状态组件的性能差别不大

  ```javascript
  import { h } from 'vue';
  const DynamicHeading = (props, { attrs, slots, emit, expose }) => {
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

  ```html
  <template>
    <transition> <div v-show="ok">hello</div> </transition>;
  </template>
  <script>
    import { h, Transition, withDirectives, vShow } from 'vue';
    export function render() {
      return h(Transition, [
        withDirectives(h('div', 'hello'), [[vShow, this.ok]]),
      ]);
    }
  </script>
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

- 如果组件中定义了 setup 配置项并且返回值是一个函数, 则其返回值作为该组件的渲染函数
- 如果组件中定义了 render 配置项, 则将其作为渲染函数
- 如果以上条件都不满足, 当前组件包含 template 配置项, 则将其作为模板进行编译成可执行的渲染函数
- 如果以上条件都不满足, 则使用容器的 innerHTML 作为模板
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

- cloneVNode 克隆一个 vnode
- isVNode 判断一个值是否为 vnode 类型
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

- withModifiers 用于向事件处理函数添加内置 `v-on` 修饰符

  ```javascript
  import { h, withModifiers } from 'vue';

  const vnode = h('button', {
    // 等价于 v-on.stop.prevent
    onClick: withModifiers(() => {
      // ...
    }, ['stop', 'prevent']),
  });
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

### v-on.native 修饰符 移除

新增 emits 选项允许子组件定义真正会被触发的事件

### v-if 与 v-for 的优先级对比

> 两者作用于同一个元素上时，v-if 会拥有比 v-for 更高的优先级

### v-bind 合并行为

> 声明绑定的顺序决定了合并顺序

```html
<!-- 模板 -->
<div id="red" v-bind="{ id: 'blue' }"></div>
<!-- 结果 -->
<div id="blue"></div>

<!-- 模板 -->
<div v-bind="{ id: 'blue' }" id="red"></div>
<!-- 结果 -->
<div id="red"></div>
```

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
