---
title: React.md
date: 2022-02-15 15:37:49
categories:
  - ES
  - React
tags:
  - js
  - jsx
  - React
---

## React

React 18.3.1

## 设计思想

- 只有当在组件树中 相同的 位置渲染 相同的 组件时, React 才会一直保留着组件的 state
- 在一般安全的情况下采用批处理方式处理 state 更新

### 变换

设计 React 的核心前提是认为 UI 只是把数据通过映射关系变换成另一种形式的数据。同样的输入必会有同样的输出。这恰好就是纯函数。

### 抽象

需要把 UI 抽象成多个隐藏内部细节，又可复用的函数。通过在一个函数中调用另一个函数来实现复杂的 UI

### 组合

将两个或者多个不同的抽象通过组合再次抽象成一个抽象

### 状态

使用不可变的数据模型, 把可以改变 state 的函数串联起来作为原点放置在顶层

状态不存在于组件内, 状态是由 React 保存的, React 通过组件在渲染树中的位置将它保存的每个状态与正确的组件关联起来

- 这个变量是否通过 props 从父组件中获取，如果是，则不是一个状态
- 这个变量是否在组件的整个生命周期中都保持不变，如果是，则不是一个状态
- 这个变量是否可以通过其他状态(state)或者属性(props)计算得到，如果是，则不是一个状态
- 这个变量是否在组件的 render 方法中使用，如果不在，则不是一个状态

<!--more-->

#### 不变性

- 不直接操作数据源, 副本替换, 可以保持以前版本的数据完好无损, 并在以后重用它们
- 不直接操作数据源, 副本替换, 使组件比较其数据是否已更改的成本非常低, 提高子组件的渲染性能

#### 构建 state 原则

- 合并关联的 state, 如果总是同时更新两个或更多的 state 变量时, 考虑将它们合并为一个单独的 state
- 避免互相矛盾的 state, 当 state 结构中存在多个相互矛盾不一致的 state 时, 应避免这种情况
- 避免冗余的 state, 如果渲染期间从组件的 props 或其它现有的 state 中计算出一些信息, 则不应该作为 state
- 避免重复的 state, 当同一数据在多个 state 之间或多个嵌套对象中重复时, 这很难保持它们同步, 应避免这种情况
- 避免深度嵌套的 state, 深度分层的 state 更新起来很不方面, 尽量以扁平化方式构建 state

- 尽量避免在 state 中镜像 props, 会与从父组件传递的属性失去同步(将不会更新), 只有想要忽略特定 props 属性的所有更新时, 使用镜像 props 才有意义, 约定 prop 名称以 `initial` 或 `default` 开头, 以说明该 prop 的新值将被忽略

#### 响应式值

props 和 state 并不是唯一的响应式值, 从它们计算出的值都是响应式的, 如果 props 或 state 发生变化, 组件将重新渲染, 从中计算出的值也会随之变化, 所以需要将响应式值包括在 Effect 的依赖项中

- 所有在组件内部直接声明的变量和函数

#### 对 state 进行保留和重置

- 将组件渲染在不同的位置, 可以重置 state
- 使用 key 重置 state, key 不是全局唯一的, 只是标识父组件内部的顺序

```jsx
import { useState } from 'react';

export default function Scoreboard() {
  const [isPlayerA, setIsPlayerA] = useState(true);
  /* 如果只用一个 Counter 组件通过 props 传递 person 属性的方式切换用户会出现 state 不会被重置*/
  return (
    <div>
      {/* 第 1 种方式: 在不同位置上, 渲染两个 Counter 组件, 在切换不同的用户时可以重置 state */}
      {isPlayerA && <Counter person="Taylor" />}
      {!isPlayerA && <Counter person="Sarah" /> }

      {/* 第 2 种方式: 在同一位置上, 使用 key 为 Counter 组件添加标识, 让 React 区分组件 */}
      {isPlayerA 
        ? (<Counter key="Taylor" person="Taylor" />) 
        : (<Counter key="Sarah" person="Sarah" />)}
      <button onClick={() => setIsPlayerA(!isPlayerA) }>下一位玩家！</button>
    </div>
  );
}
// Counter 组件
function Counter({ person }) {
  const [score, setScore] = useState(0);
  const [hover, setHover] = useState(false);

  let className = 'counter';
  if (hover) {
    className += ' hover';
  }

  return (
    <div 
      className={className}
      onPointerEnter={() => setHover(true)}
      onPointerLeave={() => setHover(false)}
    >
      <h1>{person} 的分数：{score}</h1>
      <button onClick={() => setScore(score + 1)}>加一</button>
    </div>
  );
}
```

## React.Component(deprected)

### 组件的生命周期

- constructor(props)
  - 初始化内部 state，只能在构造函数中直接为 state 赋值，其他方法中应使用 this.setState()
  - 为事件处理函数绑定实例
- static getDerivedStateFromProps(props, state)
  - 在 render() 方法执行前调用, 初始化挂载及后续更新都会被调用
  - 返回一个对象来更新 state, 如果返回 null, 则不更新任何内容
  - 无法访问组件实例
- render() class 组件中唯一必须实现的方法, 返回以下类型之一
  - React 元素
  - 数组或者 Fragments
  - Portals
  - 字符串或者数值类型
  - 布尔类型或者 null,什么都不渲染
- componentDidMount()
  - 组件挂载后立即调，DOM 节点的初始化应该放在这里

- static getDerivedStateFromProps(props, state)
- shouldComponentUpdate(nextProps, nextState)
  - 返回值影响组件是否会重新渲染
  - 默认返回值为 true
  - 首次渲染或者使用 forceUpdate() 时不会调用该方法
  - 返回 false 不会调用 render() 和 componentDidUpdate() 方法
- render()
- getSnapshotBeforeUpdate(prevProps, prevState)
  - 在最近一次渲染输出(提交到 DOM 节点)之前调用
  - 返回值作为 componentDidUpdate 方法的第三个参数 snapshot 传递,否则此参数为 undefined
- componentDidUpdate(prevProps, prevState, snapshot)

- componentWillUnmount() 组件卸载及销毁之前直接调用,此方法中不应该使用 setState() 方法,组件不会被重新渲染

## 深入 JSX

- 只能返回一个根元素
- 标签必须闭合
- 大部分属性使用驼峰方式
- 自定义的组件必须以大写字母开头

- React 将 boolean, null, undefined 视为空值, 不做任何渲染
- 直接在 JSX 中渲染对象 React 将报错 (not valid React element)

### JSX 中的 Props

- JavaScript 表达式作为 Props

- 字符串字面量

  ```jsx
  // 两个JSX表达式是等价的,
  <MyComponent message="hello world" />
  <MyComponent message={'hello world'} />
  // 字符串字面量赋值给 prop 时，它的值是未转义的
  <MyComponent message="&lt;3" />
  <MyComponent message={'<3'} />
  ```

- Props 默认值为 True

  ```jsx
  <MyTextBox autocomplete />
  <MyTextBox autocomplete={true} />
  ```

- 属性展开

```jsx
// 此方法容易将不必要的 props 传递给不相关的组件，
// 或者将无效的 HTML 属性传递给 DOM,谨慎使用该语法.
const Button = (props) => {
  const { kind, ...other } = props;
  const className = kind === 'primary' ? 'PrimaryButton' : 'SecondaryButton';
  return <button className={className} {...other} />;
};

const App = () => {
  return (
    <div>
      <Button kind='primary' onClick={() => console.log('clicked!')}>
        Hello World!
      </Button>
    </div>
  );
};
```

### JSX 中的 children

- 包含在开始和结束标签之间的 JSX 内容将作为特定属性 props.children 传递给组件.

```jsx
function MyComponent({title, children}){
  return (
    <div title={title}>
      {children}
    </div>
  )
}
function App(){
  return (
    <MyComponent title='my-component'>
      This is children
      <p>tag p</p>
    </MyComponent>
  )
}
```

- 批量子元素

```jsx
// 调用子元素回调 numTimes 次，来重复生成组件
function Repeat(props) {
  let items = [];
  for (let i = 0; i < props.numTimes; i++) {
    items.push(props.children(i));
  }
  return <div>{items}</div>;
}

function ListOfTenThings() {
  return <Repeat numTimes={10}>{(index) => <div key={index}>This is item {index} in the list</div>}</Repeat>;
}
```

## React Hook

一个特殊的函数, 只能在组件或自定义 Hook 的顶层调用

- React 16.8 新增
- 可以在函数组件内"钩入" React State 及生命周期等特性的函数
- 不能在 class 组件中使用

### Hook 规则

- 只能在函数最顶层调用 Hook, 不能在循环、条件判断或者嵌套函数中调用
- 只能在 React 的函数组件中调用 Hook, 不能在其他 JavaScript 函数中调用
- 对 Hook 的每个调用完全独立于对同一个 Hook 的其他调用

### React 内置 Hook

#### useId <em id="useId"></em> <!--markdownlint-disable-line-->

生成传递给无障碍属性的唯一 ID, 不能用来生成数据列表中的 key

- 能够确保与服务器端渲染一起工作

```jsx
const id = useId();
```

```jsx
import {useId} from 'react';

function PasswordFiled(){
  const passwordHintId = useId();

  return (
    <>
      <input type="password" aria-describedby={passwordHintId}/>
      <p id={passwordHintId}>密码至少包含 18 个字符</p>
    </>
  )
}
```

#### useState

- 用于保存渲染间的数据
- 更新变量并触发 React 再次渲染组件
- 调用 set 函数不会改变已经执行的代码中当前的 state

```jsx
const [state, setState] = useState(initialState);
```

```jsx
// useState 实现原理
let componentHooks = []; // 把 state 存储在外面
let currentHookIndex = 0;

function useState(initialValue) {
  let pair = componentHooks[currentHookIndex];
  if(pair){
    currentHookIndex++;
    return pair;
  }

  // 第一次渲染时
  pair = [initialValue, setState];
  function setState(nextState){
    pair[0] = nextState;
    updateDOM(); // 更新DOM
  }

  // 存储这个 pair 用于将来的渲染
  // 并且为下一次 hook 的调用做准备
  componentHooks[currentHookIndex] = pair;
  currentHookIndex++;
  return pair;
}
```

- 初始 state 参数只在第一次渲染时会被用到

```jsx
import {useState} from 'react';

const [count, setCount] = useState(0);
const [fruit, setFruit] = useState('banana');
const [todos, setTodos] = useState([{ text: 'learn React Hooks' }]);
setTodos([...todos, { text: 'Hello setTodos' }]); // 和 this.setState 的区别: 不会进行 state 合并
```

- 函数式更新, 新的 state 需要通过之前的 state 计算得出
- React 将更新函数放入队列中, 在下一次渲染期间, React 将通过队列将所有更新函数应用于先前的状态来计算下一个状态

```jsx
setCount((prevCount) => prevCount - 1);
<button onClick={setCount.bind(null, (prevCount) => prevCount + 1)}>+</button>;
```

- 和当前的 state 合并更新, 如果更新函数返回值和当前 state 完全相同，则重渲染会被跳过

```jsx
setState((prevState) => {
  return { ...prevState, ...updateValues };
});
```

- 惰性初始 state, 初始 state 需要通过复杂计算获得, 可以使用回调函数返回计算后的初始 state

```jsx
const [count, setCount] = useState(() => {
  // ... 初始 state 的复杂计算
  return initialState;
});
```

#### useReducer <em id="useReducer"></em> <!--markdownlint-disable-line-->

整合组件的状态更新逻辑

将组件拥有的许多状态更新逻辑的事件处理程序整合到一个 外部函数 中, 这个函数称为 reducer

- 每个 action 都描述了一个单一的用户交互, 即使它会引发数据的多个变化
- reducer 用于更新 state 的纯函数, 即当输入相同时, 输出也相同, 参数为 state 和 action
- init 用于计算初始值的函数, 如果存在, 使用 `init(initialArg)` 的执行结果作为初始值, 否则使用 initialArg

```jsx
const [state, dispatch] = useReducer(reducer, initialArg, init);
```

```jsx
// 实现原理
import {useState} from 'react';
function reducer(state, action){
  switch(action.type){
    case 'add':
      return {
        ...state
      };
    case 'delete':
      return {
        ...state
      };
    default:
      throw new Error('Unknow action: ' + action.tye);
  }
}
function useReducer(reducer, initialState){
  const [state, setState] = useState(initialState);

  function dispatch(action){
    const nextState = reducer(state, action);
    setState(nextState);
  }

  return [state, dispatch];
}
```

- 指定初始 state
  - (redux 的参数约定)将初始 state 作为第二个参数传入 useReducer,可传入 undefined.
  - 如果 Reducer Hook 的返回值和当前 state 相同, React 将跳过子组件的渲染及副作用的执行

```jsx
import {useReducer} from 'react';

// 计数器 Demo
const initialState = { count: 0 };
function reducer(state, action) {
  switch (action.type) {
    case 'increment':
      return { count: state.count + 1 };
    case 'decrement':
      return { count: state.count - 1 };
    default:
      throw new Error();
  }
}
function Counter() {
  const [state, dispatch] = useReducer(reducer, initialState);
  return (
    <>
      Count: {state.count}
      <button onClick={() => dispatch({ type: 'decrement' })}>-</button>
      <button onClick={() => dispatch({ type: 'increment' })}>+</button>
    </>
  );
}
```

- 惰性初始化
  - useReducer 的第三个参数, initialArg 将作为 init 函数的参数并返回新的 state

  ```jsx
  const initialState = { count: 0 };
  function init(initialState) {
    return { ...initialState };
  }
  function reducer(state, action) {
    return { ...state };
  }
  function Counter({ initialState }) {
    const [state, dispatch] = useReducer(reducer, initialState, init);
    return <>{...state}</>;
  }
  ```

#### useContext <em id="useContext"></em> <!-- markdownlint-disable-line -->

深度传递信息

Context 允许父组件向其下层无论多深的任何组件提供信息, 而无需通过 props 显式传递

- 接收一个 context 对象(React.createContext 的返回值)并返回该 context 的当前值
- 仍需要在上层组件树中使用 Provider 提供 context

使用 Context: 如果以下方法不适合再考虑使用

- 从传递 props 开始, 这样做能够清晰的展示组件的数据流
- 抽象组件并将 jsx 作为 children 传递, 将使用 props 传递的属性抽象为 children 组件传递

使用场景:

- 主题
- 当前账户
- 路由
- 状态管理

```jsx
const someValue = useContext(someContext);
```

```jsx
import {createContext, useContext, useState} from 'react';

const themes = {
  light: {
    height: '30px',
    color: '#0088ff',
    background: '#eeeeee',
  },
  dark: {
    height: '30px',
    color: '#ffffff',
    background: '#222222',
  },
};
const ThemeContext = createContext(themes);
function ContextDemo() {
  const [theme, setTheme] = useState('dark');
  return (
    <ThemeContext.Provider value={themes[theme]}>
      <h2>ContextDemo</h2>
      <p>current theme: {theme}</p>
      <ThemeButton onClick={setTheme.bind(null, (prevState) => (prevState === 'dark' ? 'light' : 'dark'))} />
    </ThemeContext.Provider>
  );
}

function ThemeButton({onClick}) {
  // 使用 useContext 代替 ThemeContext.consumer 消费组件
  const theme = useContext(ThemeContext);
  return (
    <button onClick={onClick} style={{ ...theme }}>
      useContext button
    </button>
  );
}
```

- 层级标题

```jsx
import {createContext, useContext} from 'react';

const LevelContext = createContext(0);

function Heading({children}){
  // 组件内部读取上层的 level
  const level = useContext(LevelContext);
  switch(level){
    case 1:
      return <h1>{children}</h1>;
    case 2:
      return <h2>{children}</h2>;
    case 3:
      return <h3>{children}</h3>;
    case 4:
      return <h4>{children}</h4>;
    case 5:
      return <h5>{children}</h5>;
    case 6:
      return <h6>{children}</h6>;
    default:
      throw new Error('unknown level: ' + level);
  }
}

function Section({children}){
  // 使用 useContext 读取上层的 level, 不需要在 props 接收 level
  const level = useContext(LevelContext);

  return (
    <section>
      <LevelContext.Provider value={level + 1}>
        {children}
      </LevelContext.Provider>
    </section>
  )
}

function App(){
  return (
    {/* 
      在 Section 内部使用 useContext 读取上层的 level,
      不需要在此处向 Section 内部传递 level 
    */}
    <Section> 
      <Heading>主标题</Heading>
      <Section>
        <Heading>副标题</Heading>
        <Heading>副标题</Heading>
        <Heading>副标题</Heading>
        <Section>
          <Heading>子标题</Heading>
          <Heading>子标题</Heading>
          <Heading>子标题</Heading>
          <Heading>子标题</Heading>
          <Section>
            <Heading>子子标题</Heading>
            <Heading>子子标题</Heading>
            <Heading>子子标题</Heading>
            <Heading>子子标题</Heading>
          </Section>
        </Section>
      </Section>
    </Section>
  )
}
```

- 嵌套 Context

```jsx
import {createContext, useContext, useReducer} from 'react';

const TasksContext = createContext(null);
const TasksDispatchContext = createContext(null);

function App(){
  // ...
  return (
    <TasksProvider>
      {/* ... */}
    </TasksProvider>
  )
}

function TasksProvider({children}){
  const [tasks, dispatch] = useReducer(tasksReducer, initialTasks);
  return (
    <TasksContext value={tasks}>
      <TasksDispatchContext value={dispatch}>
        {children}
      </TasksDispatchContext>
    </TasksContext>
  )
}
function tasksReducer(state, action){
  // ...
}
```

#### useEffect

Effect 允许指定由渲染本身, 而不是特定事件引起的副作用

- 异步执行，不会阻塞浏览器更新
- useEffect 每次在调用一个新的 effect 之前对前一个 effect 进行清理，防止内存泄漏或崩溃的问题
- 尽量避免使用 对象 和 函数 作为 Effect 依赖, 否则会使 Effect 更频繁地重新同步
- 如果 Effect 的不同部分因不同原因需要重新运行, 将其拆分为多个 Effect
- 默认情况下，它在第一次渲染之后和每次更新之后都会执行
- 第二个参数为依赖项数组, 控制 effect 的执行时机
  - 依赖项为空数组, 表示只会在组件挂载后执行
  - 没有依赖项数组时, 表示每次重新渲染后重新执行

不使用 Effect 的情况:

- 不必为了渲染而使用 Effect 来转换数据
- 不必使用 Effect 来处理用户事件

日志出现两次挂载的原因, React 在开发环境中会在初始化挂载组件后, 立即再挂载一次, 帮助查找问题所在

```jsx
useEffect(setup, dependencies?);
```

```jsx
useEffect(() => {
  // ...
  return () => {
    // 组件卸载时调用, 可选
    // ...
  };
  // effect 会比较数组中所有参数是否和前一次的参数全等，如果有一个不相等则执行 effect
  // 依赖项数组为空，只执行一次 effect
}, [count]); // 仅在 count 更改时更新

useEffect(()=>{
  // 没有依赖项数组：每次重新渲染后重新运行
});
```

##### 在 Effect 中请求数据

- Effect 不能在服务器上运行, 服务器渲染的 HTML 将只包含没有数据的 Loading 状态
- 在 Effect 中请求数据容易造成网络瀑布, 当请求一些数据再渲染子组件, 然后重复这样的过程来请求子组件的数据
- 在 Effect 中直接请求数据通常意味着不会预加载或缓存数据
- 不符合工效学, 在调用 fetch 时, 需要编写大量样板代码, 以避免像竞争条件这样的 bug

#### useEffectEvent

- 只在 Effect 内部调用
- 永远不要把它们传递给其他的组件或 Hook

从 Effect 中提取非响应式逻辑, 在 响应式逻辑(Effect) 中使用响应式值而不用担心引起周围代码因为变化而重新执行

- EffectEvent 是非响应式的并且必须从依赖项中删除

```jsx
const onSomething = useEffectEvent(callback);
```

```jsx
import {useContext, useEffect, useEffectEvent} from 'react';
// 聊天室应用
// https://zh-hans.react.dev/learn/separating-events-from-effects

function Page({url}){
  const {items} = useContext(ShoppingCartContext);
  const numberOfItems = items.length;

  // 非响应式逻辑, 内部读取的 numberOfItems 一直都是最新的
  // 但是 numberOfItems 自己变化不会引起重新渲染
  const onVisit = useEffectEvent(visitedUrl => {
    logVisit(visitedUrl, numberOfItems);
  });
  useEffect(() => {
    // Effect 内部依然是响应式的, url 变化会去调用 onVisit
    onVisit(url);

    // 日志 logVisit 想加入购物车数量时, 需要添加依赖项 numberOfItems
    // 如果依赖项加入 numberOfItems, 就改变了 logVisit 日志的作用
    // 使用 useEffectEvent 提取非响应式逻辑
    // logVisit(url, numberOfItems); // 缺少依赖项 numberOfItems
  },[url]);

 // ...
}
```

#### useLayoutEffect

> 只在客户端运行, 在服务器渲染期间不会运行

在浏览器重新绘制屏幕之前触发, 可能会影响性能, 尽可能使用 useEffect

- 它会在所有的 DOM 变更之后同步调用 effect, 可以使用它来读取 DOM 布局并同步触发重渲染
- 计算布局, 不希望用户看到某些元素在移动

```jsx
useLayoutEffect(setup, dependencies?);
```

```jsx
import {useRef, useState, useLayoutEffect} from 'react';
import {createPortal} from 'react-dom';

function Tooltip({children, targetRect}){
  const ref = useRef(null);
  const [tooltipHeight, setTooltipHeight] = useState(0);

  useLayoutEffect(() => {
    const {height} = ref.current.getBoundingClientRect();
    setTooltipHeight(height); // 重新渲染提示框的位置
  }, []);

  let tooltipX = 0;
  let tooltipY = 0;
  if(targetRect !== null){
    tooltipX = targetRect.left;
    tooltipY = targetRect.top - tooltipHeight;
    if(tooltipY < 0){
      tooltipY = targetRect.bottom; 
    }
  }

  return createPortal(
    <ToolTipContainer x={tooltipX} y={tooltipY} contentRef={ref}>
      {children}
    </ToolTipContainer>,
    document.body
  )
}
```

#### useInsertionEffect

> 只在客户端运行, 在服务器渲染期间不会运行

为 CSS-in-js 库的作者特意打造的, 除非正在使用 CSS-in-js 库并且需要注入样式, 否则应该使用 useEffect 或 useLayoutEffect

- 为布局副作用触发之前将元素插入到 DOM 中
- 不能在内部更新状态

CSS-in-js 常见的实现方法:

- 使用编译器静态提取到 CSS 文件
- 内联样式, 例如 `<div style={{opacity: 1}}>`
- 运行时注入 `<style>` 标签

```jsx
useInsertionEffect(setup, dependencies?);
```

```jsx
import {useInsertionEffect} from 'react';

let isInserted = new Set();
function useCSS(rule){
  useInsertionEffec(()=>{
    // 在此注入 <style> 标签
    if(!isInserted.has(rule)){
      isInserted.add(rule);
      document.head.appendChild(getStyleForRule(rule));
    }
  });
  return rule;
}
function Button(){
  const className = useCss('...');
  return <div className={className} />;
}
```

#### useActionState(experimental)

可以根据某个表单动作的结果更新 state 的 Hook

调用 useActionState 时在参数中传入现有的表单动作函数以及一个初始状态, 返回一个新的 action 函数和一个 form state 以供在 form 中使用, 这个新的 form state 也会作为参数传入提供的表单动作函数

- form state 是一个只在表单被提交触发 action 后才会被更新的值
- 如果 action 是一个 url, form 的行为就是普通表单提交
- 如果 action 是一个函数, form 的行为由这个函数控制, action 同时会重写 `<button>`、`<input type="submit"/>`、`<input type="image"/>` 的 formAction(表单提交的 url) 属性

- action 作为函数, 当表单被提交时触发
  - 第一个参数为上一次调用 action 函数的返回值, 第一次调用时传入的是 initialState
  - 余下的参数为普通表单动作接到的参数

```jsx
const [state, formAction] = useActionState(action, initialState, permalink?);
```

```jsx
import {useActionState} from 'react';
async function increment(prevState, formData){
  // await fetch();
  await new Promise((resolve) => {
    setTimeout(() => resolve(), 2000);
  });
  return prevState + 1;
}
function StatusForm(){
  const [state, formAction] = useActionState(increment, 0);
  return (
    // action 是 url 则直接提交表单, 如果是函数, 则控制表单的提交行为
    <form action={formAction}>
      {state}
      {/* form 的 action, enctype, method, target 属性可以被 button, type="submit", type="image" */}
      {/* 的 formaction, formenctype, formmethod, formtarget 属性重写*/}
      <button formAction={formAction}>+1</button>
      {/* button 和 input type="submit", input type="image" 是等价的 */}
      <input type="submit" formAction={formAction}/>
      <input type="image" formAction={formAction}/>
    </form>
  );
}
```

#### useCallback

> 允许在多次渲染中缓存函数的 Hook, 通常应用于性能优化

- 把内联回调函数及依赖项数组作为参数传入 useCallback，它将返回该回调函数的 memoized 版本，该回调函数仅在某个依赖项改变时才会更新
- 优化针对于子组件渲染
- 第二个参数为依赖项数组
  - 没有依赖项数组时, 每次都会返回一个新的函数
  - 依赖项为空数组时, 不需要依赖项

```jsx
const cachedFn = useCallback(fn, dependencies);
```

```jsx
const memoizedCallback = useCallback(() => {
  doSomething(a, b);
}, [a, b]);

// 每次都返回一个新函数: 没有依赖项数组
const cachedFn = useCallback(()=>{
  doSomething();
});
```

#### useMemo

缓存每次重新渲染都需要计算的结果

把“创建”函数和依赖项数组作为参数传入 useMemo，它仅会在某个依赖项改变时才重新计算 memoized 值

- 优化针对于当前组件高开销的计算
- 传入 useMemo 的函数会在渲染期间执行
- 第二个参数为依赖项数组
  - 没有依赖项数组时, 每次渲染都会计算新的值
  - 依赖项为空数组时, 不需要依赖项

```jsx
const cachedValue = useMemo(calculateValue, dependencies);
```

```jsx
import {useMemo} from 'react';
function TodoList({todos, tab}){
  const visibleTodos = useMemo(() => filterTodos(todos, tab), [todos, tab]);
  // ...
}
```

#### useRef <em id="useRef"></em> <!--markdownlint-disable-line-->

希望 React 记住某些信息, 但又不想让这些信息触发新的渲染时, 使用 ref

- ref 在重新渲染之间由 React 保留, 更改 ref 不会触发更新
- 通过 .current 访问该 ref 的值, 不要在渲染期间写入或读取 ref.current, 会破坏这些预期行为
- 引用 DOM 节点, 在 DOM 节点被移除时, React 将重置 ref 的值为 null

```jsx
const refContainer = useRef(initialValue);

const refDemo = useRef(0);
console.log(refDemo.current); // 0
```

```jsx
import {useRef} from 'react';

function TextInputWithFocusButton() {
  const inputEl = useRef(null);
  function handleBtnClick(){
    // `current` 指向已挂载到 DOM 上的文本输入元素
    inputEl.current.focus();
  }
  return (
    <>
      <input ref={inputEl} type='text' />
      <button onClick={handleBtnClick}>Focus the input</button>
    </>
  );
}
```

- 批量操作 DOM, 使用 ref 回调, 将函数传递给 ref 属性, 当需要设置 ref 时, React 将调用 ref 回调并传入 DOM 节点,
  - 并在需要清除它时传入 null
  - 或者手动返回一个回调函数, 在回调函数内管理 map

```jsx
import {useRef, useState} from 'react';
function setupCatList() {
  const catList = [];
  for (let i = 0; i < 10; i++) {
    catList.push("https://loremflickr.com/320/240/cat?lock=" + i);
  }
  return catList;
}
// 批量 ref, 使用 ref 回调
function App(){
  const itemsRef = useRef(null);
  const [catList, setCatList] = useState(setupCatList);

  function getMap(){
    if(!itemsRef.current){
      itemsRef.current = new Map();
    }
    return itemsRef.current;
  }
  function scrollToCat(cat){
    const map = getMap();
    const node = map.get(cat); // 获取 map 中存储的 DOM 节点

    node.scrollToView({
      behivor: 'smooth',
      block: 'nearest',
      inline: 'center'
    });
  }

  return (
    <>
      <button onClick={() => scrollToCat(catList[0])}>first</button>
      <button onClick={() => scrollToCat(catList[1])}>second</button>
      <button onClick={() => scrollToCat(catList[2])}>third</button>
      <ul>
        {catList.map(cat => (
          <li 
            key={cat} 
            ref={ node => {
              const map = getMap();
              if(node){
                map.set(cat, node);
              } else {
                // 方式1: 清除 DOM
                map.delete(cat);
              }
              // 方式2: 清除 DOM
              return () => {
                map.delete(cat);
              }
            }}
          ></li>
          <img src={cat} />
        ))}
      </ul>
    </>
  )
}
```

- 不能访问其他组件的 DOM 节点, 借助 [forwardRef](#forwardRef) 函数

```jsx
import {useRef} from 'react';

function MyInput(props){
  return (<input {...props}/>);
}
function App(){
  const inputRef = useRef(null);

  function handleClick(){
    // 报错, 不能访问其他组件的 DOM 节点, 需要结合 forwardRef 使用
    inputRef.current.focus(); 
  }

  return (
    <>
      <MyInput ref={inputRef}/>
      <button onClick={handleClick}>Click</button>
    </>
  )
}
```

#### useImperativeHandle

> 自定义由 ref 暴露出来的句柄

- useImperativeHandle 应当和 [forwardRef](#forwardRef) 一起使用
- 使用 ref 时自定义暴露给父组件的实例值

- createHandle 函数不需要参数, 返回想要暴露的 ref 的句柄

```jsx
useImperativeHandle(ref, createHandle, dependencies?);
```

```jsx
import {useRef, forwardRef, useImperativeHandle} from 'react';

const MyInput = forwardRef((props, ref) => {
  const realInputRef = useRef(null);
  // 限制对外层暴露的功能
  useImperativeHandle(ref, () => ({
    // 只暴露 focus 和 scrollIntoView 方法
    focus(){
      realInputRef.current.focus();
    },
    scrollIntoView(){
      realInputRef.current.scrollIntoView();
    })
  })

  return <input ref={realInputRef} {...props}/>
});

function App(){
  const inputRef = useRef(null);
  function handleClick(){
    inputRef.current.focus();
  }
  return (
    <>
      <MyInput ref={inputRef}/>
      <button onClick={handleClick}>input Focus</button>
    </>
  );
}
```

#### useDebugValue

> 可用于在 React 开发者工具中为自定义 Hook 添加标签

- 第二个可选参数, 只有在Hook被检查时才会被调用，接收debug值作为参数，并返回一个格式化的显示值

```jsx
useDebugValue(value， format?);

// useDebugValue(date, date => date.toLocalDateString());
```

```jsx
import {useDebugValue, useSyncExternalStore} from 'react';

function useOnlineStatus(){
  const isOnline = useSyncExternalStore(subscribe, () => navigator.onLine, () => true);
  // 在开发者工具中为 StatusBar 组件添加标签
  useDebugValue(isOnline ? 'Online' : 'Offline');
  return isOnline;
}
function StatusBar(){
  const isOnline = useOnlineStatus();
  return <h1></h1>;
}
function App(){
  return <StatusBar />;
}
```

#### useDeferredValue

> 允许延迟更新 UI 的非关键部分, 以让其他部分先更新

- 当 useDeferredValue 接收到与之前不同的值时, 除了当前渲染, 它还会安排一个后台重新渲染, 这个后台是可被中断的, 如果 value 有新的更新, React 会从头开始重新启动后台渲染
- useDeferredValue 本身不能阻止额外的网络请求
- useDeferredValue 本身不会引起任何固定的延迟, 一旦 React 完成原始的重新渲染, 它会立即开始使用新的延迟值处理后台重新渲染, 由事件(例如 输入)引起的任何更新都会中断后台重新渲染, 并被优先处理
- 由 useDeferredValue 引起的后台重新渲染在提交到屏幕之前不会触发 Effect, 如果后台重新渲染被暂停, Effect 将在数据加载后和 UI 更新后运行

```jsx
const deferredValue = useDeferredValue(value, initialValue);
```

```jsx
import {useState, useDeferredValue, Suspense} from 'react';
function App(){
  const [query, setQuery] = useState('');
  const deferredQuery = useDeferredValue(query);
  const isStale = query !== deferredQuery;
  
  const sty = {
    opacity: isStale ? 0.5 : 1,
    transition: isStale ? 'opacity 0.2s 0.2s linear': 'opacity 0.2s 0.2s linear'
  }

  // 每次输入时, 旧的列表会略微变暗, 直到新的结果列表加载完毕
  // 或者使用 CSS 过渡来延迟变暗的过程
  return (
    <>
      <label>
        Search albums:
        <input value={query} onChange={(e) => setQuery(e.target.value)}/>
      </label>
      <Suspense>
        <div style={sty}>
          {/* 查询结果列表 */}
          <SearhReasults query={deferredQuery}/>
        </div>
      </Suspense>
    </>
  )
}
```

#### useSyncExternalStore

订阅外部 store

- subscribe 接收一个单独的 callback 参数并把它订阅到 store 上, 当 store 发生改变时调用被提供的 callback, 引起组件重新渲染
- getSnapshot 返回组件需要的 store 中的数据快照, store 不变的情况下, 重复调用必须返回同一个值, 否则, React 就会重新渲染组件
- getServerSnapshot 返回 store 中数据的初始快照, 只会在服务器渲染时, 以及在客户端进行 hydration 时被用到

```jsx
const snapshot = useSyncExternalStore(subscribe, getSnapshot, getServerSnapshot?);
```

- 监听 store

```jsx
import {useSyncExternalStore} from 'react';
let nextId = 0;
let todos = [{id: nextId++, text: 'Todo # 1'}];
let listeners = [];

const todosStore = {
  addTodo(){
    todos = [...todos, {id: nextId++, text: 'Todo # ' + nextId}];
    emitChange();
  }
  subscribe(listener){
    // 接收一个 callback, 并返回一个 cleanup 函数
    listeners = [...listeners, listener];
    return () => {
      listeners = listeners.filter(l => l !== listener);
    }
  }
  getSnapshot(){
    return todos;
  }
}
function emitChange(){
  for(let listener of listeners){
    listener();
  }
}
function TodosApp(){
  const todos = useSyncExternalStore(todosStore.subscribe, todos.getSnapshot);

  return (
    <>
      <button onClick={() => todosStore.addTodo()}>Add Todo</button>
      <hr/>
      <ul>
        {todos.map(todo => (
          <li key={todo.id}>{todo.text}</li>
        ))}
      </ul>
    </>
  )
}
```

- 订阅浏览器 API

```jsx
import {useSyncExternalStore} from 'react';

function subscribe(callback){
  window.addEventListener('online', callback);
  window.addEventListener('offline', callback);
  return () => {
    window.removeEventListener('online', callback);
    window.removeEventListener('offline', callback);
  }
}
function useOnlineStatus(){
  const isOnline = useSyncExternalStore(subscribe, () => navigator.onLine, () => true);
  return isOnline;
}
```

#### useOptimistic(experimental)

更乐观的更新用户界面, 这种技术有助于使应用程序在感觉上响应地更加快速

- 界面立即更新为预期的结果, 而不是等待服务器的响应来反映更改
- 允许在进行异步操作时显示不同 state, 接受 state 作为参数, 并返回该 state 的副本, 在异步操作(如网络请求)期间可以不同

- optimisticState 结果乐观状态, 除非有操作挂起, 否则它等于 state, 在这种情况下, 它等于 updateFn 返回的值
- addOptimistic 触发乐观更新时调用的 dispatch 函数, 接收一个任意类型的参数 optimisticValue, 并以 state 和 optimisticValue 作为参数调用 updateFn
- updateFn 接受当前的 state 和传递给 addOptimistic 的乐观值, 并返回结果乐观状态

```jsx
cnost [optimisticState, addOptimistic] = useOptimistic(state, updateFn);
```

```jsx
import {useOptimistic, useRef, useState} from 'react';

function Thread({messages, sendMessage}){
  const formRef = useRef(null);
  const [optimisticMessages, addOptimisticMessage] = useOptimistic(messages, (state, optimisticValue) => {
    // 使用乐观值
    // 合并返回新的 state
    return [...state, {text: optimisticValue, sending: true}];
  });

  async function formAction(formData){
    addOptimisticMessage(formData.get('message'));
    formRef.current.reset();
    // 异步操作
    await sendMessage(formData); 
  }

  return (
    <>
      {optimisticMessages.map((message, index) => (
        <div key={index}>
          {message.text}
          {message.sending && <small>发送中...</small>}
        </div>
      ))}
      <form action={formAction} ref={formRef}>
        <input type="text" name="message" placeholder="hello..."/>
        <button>发送</button>
      </form>
    </>
  )
}
function App(){
  const [messages, setMessages] = useState([
    {text: '你好, 在这!', sending: false, key: 1}
  ]);

  async function sendMessage(formData){
    // 模拟异步操作
    const sentMessage = await deliverMessage(formData.get('message'));
    // 更新状态
    setMessages(messages => [...messages, {text: sentMessage}]);
  }

  return <Thread messages={messages} sendMessage={sendMessage}/>
}
async function deliverMessage(message){
  await new Promise(resolve => setTimeout(resolve, 3000));
  return message;
}
```

#### useTransition

在不阻塞 UI 的情况下更新状态, 将某些状态更新标记为 transition

- transition 更新不能用于控制文本输入
- isPending 是否存在待处理的 transition
- startTransition 调用此函数将状态更新标记为 transition, 传递给此函数的函数必须是同步的, React 会立即执行此函数, 并将在其执行期间发生的所有状态更新标记为 transition, 如果在其执行期间, 尝试稍后执行状态更新, 这些状态更新不会被标记为 transition

```jsx
const [isPending, startTransition] = useTransition();
```

### 自定义 Hook

- Hook 的名称必须以 use 开头
- 自定义 Hook 共享的是状态逻辑, 而不是状态本身

```jsx
function useDemo(prop) {
  const [count, setCount] = useState(0);
  useEffect(() => {
    // 执行订阅
    return () => {
      // 执行取消订阅
    };
  });
  return count;
}
```

- 网络状态 Hook

```jsx
import {useEffect,useState} from 'react';

// 无法检测已离线的网络
// 如果在生成初始 HTML 的服务端直接使用它是无效的
function useOnlineStatus(){
  const [isOnline, setIsOnline] = useState(true);

  useEffect(()=>{
    function handleOnline(){
      setIsOnline(true);
    }
    function handleOffline(){
      setIsOnline(false);
    }
    window.addEventListener('online', handleOnline);
    window.addEventListener('offline', handleOffline);
    return () => {
      window.removeEventListener('online', handleOnline);
      window.removeEventListener('offline', handleOffline);
    }
  },[]);
  return isOnline;
}

// 改进版, 使用 useSyncExternalStore
function subscribe(callback){
  window.addEventListener('online', callback);
  window.addEventListener('offline', callback);
  return () => {
    window.removeEventListener('online', callback);
    window.removeEventListener('offline', callback);
  }
}
function useOnlineStatus(){
  return useSyncExternalStore(subscribe, () => navigator.onLine, () => true);
}
```

## React 组件

### React 内置组件

#### Fragment

通常使用 `<>...</>` 简写形式代替, 允许在不添加额外节点的情况下将子元素组合

- 如果传递一个 key 给 Fragment 时, 不能使用简写形式

```jsx
import {Fragment} from 'react';

function Post(){
  return (
    <>
      <PostTitle/>
      <PostBody />
    </>
  )
}
// 渲染 key
function Blog(){
  return posts.map(post => (
    <Fragment key={post.id}>
      <PostTitle title={post.title}/>
      <PostBody body={post.body}/>
    </Fragment>
  ));
}
```

#### Profiler

允许编程式测量 React 树的渲染性能, 可以嵌套测量应用的不同部分

- id 字符串, 用于标识正在测量的 UI 部分
- onRender 当包裹的组件树更新时, React 回调此函数, 并传入有关渲染内容和所花费时间的信息

```jsx
import {Profiler} from 'react';
function onRender(id, phase, actualDuration, baseDuration, startTime, commitTime){
  // id 字符串, 为 Profiler 树的属性
  // phase 标识书的渲染阶段, 取值: mount, update, nested-update
  // actualDuration 渲染 Profiler 树的毫秒数
  // baseDuration 估算在没有任何优化的情况下重新渲染整棵 Profiler 子树所需的毫秒数
  // startTime 开始渲染此次更新时的时间戳
  // commitTime React 提交此次更新时的时间戳
}
function App(){

  return (
    <Profiler id="app" onRender={onRender}>
      {/* ... */}
    </Profiler>
  )
}
```

#### StrictMode

为整个应用启动严格模式, 尽早发现组件中的常见错误

所有的检查仅在开发环境中进行, 不会影响生产构建

- 组件将重新渲染一次, 以查找由于非纯渲染而引起的错误
- 组件将重新运行 Effect 一次, 以查找由于缺少 Effect 清理而引起的错误
- 组件将被检查是否使用了已弃用的 API

```jsx
import {StrictMode} from 'react';
import {createRoot} from 'react-dom/client';

const root = createRoot(document.getElementById('root'));
root.render(
  <StrictMode>
    <App/>
  </StrictMode>
);
// 部分开启严格模式
function App(){
  return (
    <>
      <h1>Header</h1>
        <StrictMode>
          {/* ... */}
        </StrictMode>
      <footer></footer>
    </>
  )
}

```

#### Suspense

允许子组件完成加载前展示后备方案, Suspense 无法检测到 Effect 或事件处理程序中获取数据的情况

- children 真正的 UI 渲染内容, 如果 children 在渲染中被挂起, Suspense 将会渲染 fallback
- fallback 真正的 UI 未渲染完成时代替其渲染的备用 UI

激活 Suspense 组件的数据源

- 支持 Suspense 的框架如 Relay 和 Next.js
- 使用 lazy 懒加载组件代码
- 使用 use 读取 Promise 的值

```jsx
import {Suspense} from 'react';

function App(){
  return (
    <Suspense fallback={<Loading />}>
      {/* ... */}
    </Suspense>
  )
}
```

## React API

### act

测试助手, 用于在做出断言之前应用挂起的 React 更新, 通常用在测试库中

```jsx
import {act} from 'react';
await act(async actFn);
```

### cache(experimental)

允许缓存数据获取或计算的结果, 在任何组件之外调用创建带有缓存的函数版本

- 仅供与 服务器组件一起使用

```jsx
import {cache} from 'react';
const cachedFn = cache(fn);
```

### createContext

创建一个 Context 提供给子组件, 通常和 [useContext](#useContext) 配合使用

### forwardRef <em id="forwardRef"></em> <!--markdownlint-disable-line-->

允许组件使用 ref 将 DOM 节点暴露给父组件

- render 渲染函数, React 将使用 props 和 ref 调用此函数, 返回的 JSX 作为组件的输出

```jsx
const SomeComponent = forwardRef(render);
```

```jsx
import {forwardRef, useRef} from 'react';

const MyInput = forwardRef((props, ref) => {
  return <input ref={ref} {...props}/>;
});

function App(){
  const inputRef = useRef(null);
  return <MyInput ref={inputRef}/>;
}
```

### lazy

在组件第一次被渲染之前延迟加载组件的代码, 通过将懒加载组件或其任何父级包装到 Suspense 边界中实现

- load 该函数返回一个 Promise 或一个 thenable 对象, React 首次调用 load 后将等待解析, 然后将解析值的 .default 渲染为 React 组件

```jsx
const LazyComponent = lazy(load);
```

```jsx
import {lazy} from 'react';

const MarkdownPreview = lazy(() => import('./MarkdownPreview.js'));

function App(){
  return (
    <Suspense fallback={<Loading/>}>
      <MarkdownPreview/>
    </Suspene>
  )
}
```

### memo

允许组件在 props 没有改变的情况下跳过重新渲染

只有当组件经常使用完全相同的 props 进行渲染时, 并且其重新渲染逻辑是非常昂贵的, 使用 memo 优化才有意义

- 记忆化只与从父组件传递给组件的 props 有关, 即使组件已被记忆化, 当其使用的 context 发生变化时, 仍将重新渲染
- componennt 要进行缓存的组件, memo 不会修改该组件, 而是返回一个新的、记忆化的组件, 接受任何有效的 React 组件, 包含函数式组件和 [forwardRef](#forwardRef) 组件
- arePropsEqual 该函数接收两个参数, 组件的上一个 props 和新的 props 进行比较
  - 新的 props 和旧的 props 具有相同的输出时返回 true
  - 否则返回 false

```jsx
const MemoizedComponent = memo(component, arePropsEqual?);
```

```jsx
import {memo} from 'react';
const Greeting = memo(function Greeting({name}){
  return <h1>{name}</h1>;
})
```

### startTransition

在不阻塞 UI 的情况下更新 state

```jsx
import {startTransition} from 'react';
function TabContainer(){
  const [tab, setTab] = useState('about');
  function selectTab(tab){
    startTransition(()=>{
      selectTab(tab);
    })
  }
  // ...
}
```

### use(experimental)

读取类似于 Promise 或 context 的资源的值

- 可以在 循环 或 条件 语句中调用 use, 调用 use 的函数仍然必须是一个 组件 或 Hook

```jsx
const value = use(resource);
```

```jsx
import {use} from 'react';
function MessageComponent({messagePromise}){
  const message = use(messagePromise);
  const theme - use(ThemeContext);
  // ...
}
```

### experimental_taintObjectReference(experimental)

允许阻止特定对象实例被传递给客户端组件, 例如 user 对象

- message 对象被传递给客户端组件时显示的消息
- object 被污染的对象, React 会阻止直接将 函数 和 类 传递给客户端组件, 并把默认的错误消息替换在 message 中定义的内容

```jsx
experimental_taintObjectReference(message, object);
```

### taintUniqueValue(experimental)

阻止将唯一值传递给客户端组件, 例如 密码、密钥或令牌

- message  value 被传递给客户端组件时显示的消息
- lifetime 指定 value 应该被污染多长时间的任何对象, 只要此对象仍然存在, 将阻止把 value 发送给任何客户端组件
- value 具有高熵的字符串或字节的唯一序列

```jsx
taintUniqueValue(message, lifetime, value);
```

## ReactDOM Hook

### useFormStatus(experimental)

提供上一次表单提交状态信息, 仅获取父级 form 的状态信息

- pending 标识父级 form 是否正在等待提交, 如果调用 useFormStatus 的组件未嵌套在 form 中, 总是返回 false
- data 包含父级 form 正在提交的 formData 数据, 如果没有进行提交为 null
- method  标识父级 form 使用 GET 或 POST 进行提交, 默认使用 GET
- action 传递给父级 form 的 action 属性的函数引用, 如果没有父级 form 则为 null

```jsx
const {pending, data, method, action} = useFormStatus();
```

```jsx
import {useFormStatus} from 'react-dom';

function Submit(){
  const {pending, data, method} = useFormStatus();
  return (
    <>
      <button type="submit" disabled={pending}>
        {pending ? '提交...' : '提交'}
      </button>
      <p>{data ? `请求 ${data.get('username')}...` : ''}</p>
      <p>method 为 {method}</p>
    </>
  )
}
function Form({action}){
  return (
    <form action={action}>
      <input name="username" type="text"/>
      <Submit/>
    </form>
  )
}
function App(){
  return <Form action={submitForm}/>;
}
```

- 不会返回在同一组件中渲染的 form 的状态信息, 仅获取父级 form 的状态信息

```jsx
function Form(){
  // useFormStatus 不会跟踪此组件中渲染的表单, pending 永远不会为 true
  // 需要将 useFormStatus 改成在 Form 组件的子组件中调用
  const {pending} = useFormStatus();
  return <form action={submit}></form>;
}
```

## ReactDOM 组件

### 自定义 HTML 元素

渲染一个带连字符的标签 如 `<my-elment>`, React 会认为渲染自定义 HTML 元素

- 所有自定义元素的 props 都将被序列化为字符串, 并且总是使用属性(attribute)进行设置
- 自定义元素接受 `class` 而不是 `className`, 接受 `for` 而不是 `htmlFor`
- 如果使用 is 属性渲染一个内置的浏览器 HTML 元素, 将被视为自定义元素

### 属性差异

在 React 中，所有的 DOM 特性和属性（包括事件处理）都应该是小驼峰命名的方式。例如，与 HTML 中的 tabindex 属性对应的 React 的属性是 tabIndex。例外的情况是 `aria-*` 以及 `data-*` 属性，一律使用小写字母命名。比如, 你依然可以用 aria-label 作为 aria-label。

- dangerouslySetInnerHTML React 为浏览器 DOM 提供 innerHTML 的替换方案
  - 不能同时传递 children 和 dangerouslySetInnerHTML
  - 直接设置 HTML 存在风险
  - key 为 \_\_html

  ```jsx
  function createMarkup() {
    return { __html: 'First &middot; Second' };
  }

  function MyComponent() {
    return <div dangerouslySetInnerHTML={createMarkup()} />;
  }
  ```

- for JavaScript 中的保留字，React 中内置元素使用 htmlFor 代替

  ```jsx
  <label htmlFor="htmlFor">htmlFor</label>
  <input name="htmlFor" id="htmlFor" value="" />
  ```

- suppressContentEditableWarning 此属性禁用 当 DOM 元素拥有 contentEditable 属性时，React 发出警告
- suppressHydrationWarning 此属性禁用警告, 如果 React 服务器与客户端渲染不同的内容时发出警告

#### 表单

- React 不支持在 option 元素上传递 selected 属性

- checked 控制复选框或单选按钮是否被选中
- value 控制文本框、下拉框、文本域的输入文本

以下属性仅在 `非受控元素` 中有效

- defaultChecked 指定复选框或单选按钮的初始值
- defaultValue 指定文本框、下拉框、文本域的初始值

### 合成事件

事件处理程序接收到一个 React 合成的事件对象

`SyntheticEvent` 为 React 的事件包装器, 是所有事件的基类型

- 事件: on + 事件名称 + Capture(捕获阶段触发)
- 事件处理器: handle + 事件名称

- 阻止事件冒泡手动调用 `e.stopPropagation()`
- 阻止部分 html 元素的浏览器默认行为调用 `e.preventDefault()`

```jsx
function App(){
  function handleClick(){
    // ...
  }
  function handleClickCapture(){
    // ...
  }
  return (
    <div 
      onClick={handleClick}
      onClickCapture={handleClickCapture}
    >
      <button onClick={handleClick}></button>
    </div>
  )
}
```

#### 属性

除了以下标准属性, React 针对不同事件额外附加了其他 属性, 如 ClipboardEvent 事件附加了 clipboardData

- bubbles 布尔值, 返回是否会在 DOM 中冒泡传播
- cancelable 布尔值, 返回事件是否可以被取消
- currentTarget DOM 节点, 返回当前事件处理程序所附加到的节点在 React 树中的位置
- defaultPrevented 布尔值, 返回是否调用了 preventDefault
- eventPhase 数字, 返回事件当前所处的阶段
- isTrusted 布尔值, 返回事件是否由用户发起
- target DOM 节点, 返回事件发生的节点(可能是远程子节点)
- timestamp 数字, 返回事件发生的时间
- nativeEvent DOM event 对象, 浏览器的原生事件对象

#### 方法

- preventDefault 阻止事件的默认浏览器行为
- stopPropagation 阻止事件在 React 树中的传播

- isDefaultPrevented 布尔值, 返回是否调用了 preventDefault 方法
- isPropagationStopped 布尔值, 返回是否调用 stopPropagation 方法

- persist 不适用 ReactDOM, 在 React Native 中, 调用此函数以读取事件后的属性
- isPersistent 不适用 ReactDOM, 在 React Native 中, 返回是否调用了 persist

## ReactDOM API

### createPortal

允许将 JSX 作为 children 渲染到 DOM 的不同部分

- children React 可以渲染的任何内容
- domNode 某个已经存在的 DOM 节点
- key 用作 portal key 的独特字符串或数字

```jsx
const createPortalDOM = createPortal(children, domNode, key?);
```

```jsx
import {createPortal} from 'react-dom';

const createPortalDOM = createPortal(<p></p>, document.body);
function List(){
  return (
    <>
      <ul>
        <li></li>
      </ul>
      {createPortalDOM}
    </>
  )
}
```

### flushSync

允许强制 React 立即同步更新 DOM, 使用此方法可能严重影响应用程序的性能

```jsx
flushSync(callback);
```

```jsx
import {useRef, useState} from 'react';
import {flushSync} from 'react-dom';

function App(){
  const listRef = useRef(null);
  const [text, setText] = useState('');
  const [todos, setTodos] = useState([]); 

  function handleAdd(){
    const newTodo = {id: nextId++, text: text};
    // 强制 React 立即同步更新 DOM
    flushSync(()=>{
      setText('');
      setTodos([...todos, newTodo]);
    });
    listRef.current.lastChild.scrollIntoView();
  }

  return (/* ... */);
}
```

### findDOMNode(deprected)

React 18 开始, 使用 ref 代替

获取组件实例对用的浏览器 DOM 节点

```jsx
const domNode = findDOMNode(componentInstance);
```

### hydrate(deprected)

React 18 开始, 使用 [hydrateRoot](#hydrateRoot) 代替

允许 React 17 及以下版本中使用 `react-dom/server` 生成的 HTML 内容作为浏览器 DOM 节点, 并在其中显示 React 组件

- reactNode 用于渲染现有的 HTML, 通常是 JSX 片段, 并且使用像 `renderToString(<App/>)` 的 ReactDOM Server 方法渲染
- domNode 在服务器中被渲染为根节点的 DOM 元素
- callback React 将在组件 hydrate 后调用

```jsx
hydrate(reactNode, domNode, callback?);
```

### render(deprected)

React 18 开始, 使用 [createRoot](#createRoot) 代替

将一段 JSX 片段渲染到浏览器 DOM 容器节点中

```jsx
render(reactNode, domNode, callback?);
```

```jsx
import {render} from 'react-dom';
import {App} from './App.js';

render(<App/>, document.getElementById('root'));
```

### unmountComponentAtNode(deprected)

React 18 开始, 使用 root.unmount 代替

从 DOM 中移除一个已挂载的 React 组件

```jsx
unmountComponentAtNode(domNode);
```

### preconnect(experimental) <em id="preconnect"></em> <!--markdownlint-disable-line-->

提前连接到一个期望从中加载资源的服务器

- 对同一服务器进行多次调用 preconnect 具有与单次调用相同的结果
- 在浏览器中, 可以在任何情况下调用 preconnect
- 服务器端渲染时, 只有在渲染组件或在从渲染组件中发起的异步上下文中调用 preconnect 时才会生效, 任何其他调用都会被忽略

```jsx
preconnect(href);
```

```jsx
import {preconnect} from 'react-dom';
function AppRoot(){
  // 预连接到主机
  preconnect('https://example.com');
}
```

### prefetchDNS(experimental)

允许提前查找期望从中加载资源的服务器的 IP, 和 [preconnect](#preconnect) 类似

```jsx
prefetchDNS(href);
```

### preinit(experimental) <em id="preinit"></em> <!--markdownlint-disable-line-->

> React 框架通常会内置资源处理方案, 不需要手动调用此 API

预获取和评估样式表或外部脚本

- 对于具有相同的 href 的多个 preinit 调用具有与单个调用相同的结果
- 在浏览器中, 可以在任何情况下调用 preinit
- 在服务器端渲染时, 只有在渲染组件或在从渲染组件中发起的异步上下文中调用 preinit 时才会生效, 任何其他调用都会被忽略

- href 要下载并执行的资源的 url
- options
  - as 标识资源的类型, 可能的值 script, style
  - precedence 与样式表一起使用时必需, 指定样式表相对于其他样式表的插入位置, 可能的值 reset, low, meduim, high
  - crossOrigin 标识要使用的 CORS 策略, 可能的值为 anonymous, use-credentials
  - integrity 标识资源的加密哈希, 用于验证其真实性
  - nonce 标识使用严格 CSP(安全内容策略) 时允许资源的加密随机数
  - fetchPriority 建议获取资源的相对优先级, 可能的值为 auto(默认), high, low

```jsx
preinit(href, options);
```

### preinitModule(experimental) <em id="preinitModule"></em> <!--markdownlint-disable-line-->

> React 框架通常会内置资源处理方案, 不需要手动调用此 API

预获取和评估 ESM 模块

执行时机参考 [preinit](#preinit)

- href 要下载并执行的模块的 url
- options
  - as 只能取值 script
  - crossOrigin 同 [preinit](#preinit)
  - integrity 同 [preinit](#preinit)
  - nonce 同 [preinit](#preinit)

```jsx
preinitModule(href, options)
```

### preload(experimental)

> React 框架通常会内置资源处理方案, 不需要手动调用此 API

预获取期望使用的资源, 比如样式表、字体、外部脚本

执行时机和参数参考 [preinit](#preinit)

```jsx
preload(href, options);
```

### preloadModule(experimental)

> React 框架通常会内置资源处理方案, 不需要手动调用此 API

预获取期望使用的 ESM 模块

执行时机和参数参考 [preinitModule](#preinitModule)

```jsx
preloadModule(href, options);
```

## 客户端 API

`react-dom/client` API 允许在客户端(浏览器) 渲染 React 组件, 通常在应用程序项目顶层调用

### createRoot <em id="createRoot"></em> <!--markdownlint-disable-line-->

允许在浏览器的 DOM 节点中创建根节点以显示 React 组件

- domNode 某个已经存在的 DOM 节点
- options
  - onCaughtError(experimental) 当 React 捕获到错误边界时调用
  - onUncaughtError(experimental) 当错误边界抛出了一个无法捕获的错误时调用
  - onRecoverableError 当 React 自动从错误中恢复时调用
  - identifierPrefix 一个字符串, React 使用此字符串作为 [useId](#useId) 生成的 id 的前缀, 当在一个页面中使用多个根节点时可以避免冲突

返回值

- root.render(reactNode) 将一段 JSX 片段渲染为 React 组件并显示, 首次调用时 React 将清空根节点所有已经存在的 HTML
- root.unmount() 销毁 React 根节点中的一个已经渲染的树

```jsx
const root = createRoot(domNode, options?);
```

```jsx
import {useState} from 'react';
import {createPortal} from 'react-dom';
import {createRoot} from 'react-dom/client';

const root = createRoot(document.getElementById('root'));
root.render(reactNode);
```

### hydrateRoot <em id="hydrateRoot"></em> <!--markdownlint-disable-line-->

<!-- 允许在先前路由 `react-dom/server` 生成的浏览器 HTML DOM 节点中展示 React 组件 -->
用 hydrateRoot 函数将 React 连接到 React 在服务器端环境中渲染的 HTML 中

开发模式下, React 会在 hydrate 期间发出不匹配警告. 在不匹配的情况下, 不能保证内容差异会被修补, 出于性能原因, 这很重要, 因为在大多数应用程序中, 不匹配很少见, 因此验证所有标记将是非常昂贵而不可行的

- domNode 一个在服务器端渲染时呈现为根元素的 DOM 元素
- reactNode 用于渲染已存在 HTML 的 React 节点, 通常是 JSX 片段, 并且使用像 `renderToPipeableStream(<App/>)` 的方法渲染
- options 参数同 [createRoot](#createRoot) 的 options

返回值

- root.render(reactNode) 更新一个 hydrate 根组件中的 React 组件来渲染浏览器端 DOM 元素
- root.unmount() 销毁 React 根节点内的渲染树

```jsx
const root = hydrateRoot(domNode, reactNode, options?);
```

```jsx
import {useState, useEffect} from 'react';
import {hydreateRoot} from 'react-dom/client';

// eg1
// index.html
// 由 react-dom/server 生成的 HMTL 内容
<html>
  <body>
    <div id="root">
      <h1>Is Server</h1>
    </div>
  </body>
</html>

function App(){
  const [isClient, setIsClient] = useState(false);
  useEffect(()=>{
    setIsClient(true);
  },[]);
  return <h1>{isClient ? 'Is Client' : 'Is Server'}</h1>;
}

hydreateRoot(document.getElementById('root'), <App/>);

// eg2
// index.html
// 由 react-dom/server 生成的 HTML 内容
<html>
  <body>
    <div id="root">
      <h1>Hello world! <!-- -->0</h1>
      <input placeholder="Type something here" />
    </div>
  </body>
</html>

function App({counter}){
  return (
    <>
      <h1>Hello world! {counter}</h1>
      <input placeholder="Type something here" />
    </>
  )
}

const root = hydrateRoot(document.getElementById('root'), <App counter={0}/>);
let i = 0;
setInterval(() => {
  root.render(<App counter={i}/>);
  i++;
}, 1000);
```

## 服务端 API

`react-dom/server` API 允许在服务器端将 React 组件渲染为 HTML, 仅在服务器端应用程序顶层调用

### Node.js 流服务器 API

#### renderToNodeStream(deprected)

React 18 开始, 改用 [renderToPipeableStream](#renderToPipeableStream)

输出 HTML 字符串的 Node.js 只读流, 此方法会缓冲所有输出, 因此实际上它并没有提供任何流式传输的好处

在客户端使用 [hydrateRoot](#hydrateRoot) 使服务器生成的 HTML 具有交互功能

- reactNode 要渲染为 HTML 的 React 节点
- options
  - identifierPrefix 字符串前缀, 由 [useId](#useId) 生成的 id 使用

```jsx
const stream = renderToNodeStream(reactNode, options?);
```

#### renderToPipeableStream <em id="renderToPipeableStream"></em> <!--markdownlint-disable-line-->

将一个 React 组件树渲染为管道化(pipeable)的 Node.js 流

在客户端使用 [hydrateRoot](#hydrateRoot) 使服务器生成的 HTML 具有交互功能

- options
  - bootstrapScriptContent 指定被放入 script 标签中作为其内容的字符串
  - bootstrapScripts 一个字符串数组, 将被转化为 script 标签嵌入页面
  - bootstrapModules 和 bootstrapScripts 类似, 但是嵌入页面的是 `<script type="module">`
  - identifierPrefix 字符串前缀, 由 [useId](#useId) 生成的 id 使用
  - namespaceURI 一个字符串, 指定与流相关联的 命名控件 URI, 默认是常规的 HTML, 可以指定 SVG
  - nonce 标识使用严格 CSP(安全内容策略) 时允许资源的加密随机数
  - onAllReady 函数, 在所有渲染完成时触发, 包括 shell 和 所有额外的 content, 可以代替 onShellReady 用于爬虫和静态内容生成
  - onShellReady 函数, 在 shell 初始化渲染后立即调用
  - onShellError 函数, 在 shell 发生错误渲染时调用
  - onError 函数, 出现异常错误时触发
  - progressiveChunkSize 一个块中的字节数

- 返回值: 包含 pipe() 和 abort() 方法的对象

```jsx
const {pipe, abort} = renderToPipeableStream(reactNode, options?);
```

```jsx
import {renderToPipeableStream} from 'react-dom/server';

app.use('/', (request, response) => {
  const {pipe, abort} = renderTopipeableStream(<App />, {
    bootstrapScriptContent: `window.assetMap = function(){ alert('assetMap') }`,
    bootstrapScripts: ['./main.js'],
    onShellReady(){
      response.setHeader('Content-Type', 'text/html');
      pipe(response);
    }
  });
});
```

#### renderToStaticNodeStream

将 Node.js 只读流渲染为非交互式 React 树, 无法 hydrate 交互功能

输出 HTML 字符串的 Node.js 只读流, 此方法会缓冲所有输出, 因此实际上它并没有提供任何流式传输的好处

此方法输出的 HTML 不能被客户端 hydrate 转换成具有交互功能

- options
  - identifierPrefix 字符串前缀, 由 [useId](#useId) 生成的 id 使用

```jsx
const stream = renderToStaticNodeStream(reactNode, options?);
```

### Web 流服务器 API

具有 web 流的环境中可用, 包括 浏览器, Deno, 以及一些现代 Edge 运行时

- ReadableStream
- WritableStream
- TransformStream

#### renderToReadableStream

将 React 树渲染后发送至 web 流

参数配置参考 [renderToPipeableStream](#renderToPipeableStream) 的 options

在客户端使用 [hydrateRoot](#hydrateRoot) 使服务器生成的 HTML 具有交互功能

- 返回一个 Promise
  - 如果 shell 渲染成功, 则 Promise 将 resolve 为 web 可读流
  - 如果 shell 渲染失败, 则 Promise 将 reject

```jsx
const stream = renderToReadableStream(reactNode, options?);
```

```jsx
import {renderToReadableStream} from 'react-dom/server';

async function handler(request){
  const stream = await renderToReadableStream(<App />,{
    bootstrapScripts: ['./main.js'],
    identifierPrefix: 'w',
    onError(error){
      console.error(error);
      logServerCrashReport(error);
    }
  });
  return new Response(stream, {
    headers: {'Content-Type': 'text/html'}
  });
}
```

### 非流环境 API

#### renderToString

将 React 树渲染为一个 HTML 字符串, 不支持流式传输或等待数据

在客户端使用 [hydrateRoot](#hydrateRoot) 使服务器生成的 HTML 具有交互功能

- 不完全支持 Suspense, 如果某个组件触发 Suspense, API 不会等待其内容解析完成, 将找到最近的 Suspense 边界并在 HTML 中渲染其 fallback 属性

- options
  - identifierPrefix 字符串前缀, 由 [useId](#useId) 生成的 id 使用

- 返回一个 HTML 字符串

```jsx
const html = renderToString(reactNode, options?);
```

```jsx
import {renderToString} from 'react-dom/server';

app.use('/', (request, response) => {
  const html = renderToString(<App/>);
  response.send(html);
});
```

#### renderToStaticMarkup

将非交互的 React 组件渲染成 HTML 字符串, 无法 hydrate 交互功能

- renderToStaticMarkup 的输出无法进行二次渲染, 仅需要呈现纯静态内容时使用
- renderToStaticMarkup 对 Suspense 的支持有限, 如果一个组件触发了 Suspense, API 立即将后备方案作为 HTML 输出
- renderToStaticMarkup 在浏览器中可以使用, 不建议在客户端代码中使用它

- options
  - identifierPrefix 字符串前缀, 由 [useId](#useId) 生成的 id 使用

- 返回一个 HTML 字符串

```jsx
const html = renderToStaticMarkup(reactNode, options?);
```

```jsx
import {renderToStaticMarkup} from 'react-dom/server';

app.use('/', (request, response) => {
  const html = renderToStaticMarkup(<App/>);
  response.send(html);
});
```

## React Router

### 路由器

#### 不支持 data APIs

- \<BrowserRouter\>
- \<MemoryRouter\>
- \<HashRouter\>
- \<NativeRouter\> 用于 React Native
- \<StaticRouter\>

#### 支持 data APIs

使用此方式创建路由, 同时启用用于数据获取的 loader, actions, fetchers 等 API

- createBrowserRouter
- createMemoryRouter
- createHashRouter
- createStaticRouter

##### createBrowserRouter <em id="createBrowserRouter"></em> <!--markdownlint-disable-line-->

- basename 基础路径
- future 用于启用新版本语法的配置对象
- hydrationData 当使用服务器端渲染时允许从服务器端获取数据
- unstable_dataStrategy 低水平 API, 将会覆盖 React Router 内部的 loader, action 的执行
- unstable_patchRoutesOnMis
- window 用于区分环境, 对开发者工具或者测试来说非常有用

返回值

- router 路由信息

```jsx
const router = createBrowserRouter(routes, {
  basename: '/app',
  hydrationData: {
    root: {
      // ...
    }
  },
});
```

##### createStaticHandler

通常用于服务器端渲染的 数据获取和提交, 配合 `createStaticRouter` 使用

- routes 路由信息
- opts
  - basename
  - future 用于启用新版本语法的配置对象
  - mapRouteProperties

返回值

- staticHandler.dataRoutes 路由信息
- staticHandler.query() 执行当前请求的 action, loader 并返回 context 包含了渲染页面的所有数据
  - request 请求
  - opts
    - routeId 如果需要调用不同的路由的 action 或 loader, 传入指定的 routeId
    - requestContext 将请求上下文信息传入 action 或 loader

staticHandler.query() 返回值

- context 包含渲染页面信息的请求上下文

##### createStaticRouter

- routes 路由信息
- context 请求的上下文信息
- opts
  - future 用于启用新版本语法的配置对象

返回值

- router 路由信息

##### StaticRouterProvider

接收来自 `createStaticHandler` 的 context 和  `createStaticRouter` 的 router, 用于服务器端渲染

- router 通过 createStaticRouter 创建的路由
- context 接收来自 staticHandler.query() 返回的结果作为数据
- hydrate 是否禁用客户端自动数据连接
- nonce 标识使用严格 CSP(安全内容策略) 时允许资源的加密随机数

```jsx
import {StrictMode} from 'react';
import {createStaticHandler, createStaticRouter, StaticRouterProvider} from 'react-router-dom';
import {renderToString} from 'react-dom/server';

const routes = [];
export async function renderHtml(req){
  const {query, dataRoutes} = createStaticHandler(routes);
  const fetchRequest = createFetchRequest(req);
  const context = await query(fetchRequest);

  if(context instanceof Response){
    throw context;
  }

  const router = createStaticRouter(dataRoutes, context);
  return renderToString(
    <StrictMode>
      <StaticRouterProvider router={router} context={context} />
    </StrictMode>
  )
}
```

##### RouterProvider

路由根组件, 所有的路由对象或者 Data APIS 都通过此组件注入 React 应用程序

- router 路由信息
- fallbackElement 后备内容
- future 用于启用新版本语法的配置对象

```jsx
import {StrictMode} from 'react';
import {createRoot} from 'react-dom/client';
import {createBrowserRouer, createRoutesFromElements, RouterProvider, Route} from 'react-router-dom';

// 使用 JSX 元素创建路由
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route 
      path="/"
      element={<Root/>}
      errorElement={<ErrorPage/>}
      loader={rootLoader}
      action={rootAction}
    >
      <Route index element={<Dashboard/>}/>
      {/* ... */}
    </Route>
  )
);
// 使用对象形式创建路由
const router = createBrowserRouter([
  {
    path: '/',
    element: <Root/>,
    loader: rootLoader,
    action: rootAction,
    errorElement: <ErrorPage/>,
    children: [
      {index: true, element: <Dashboard/>}
    ]
  }
])
createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router} fallbackElement={<SpinnerOfDom/>}/>
  </StrictMode>
)
```

### Route <em id="Route"></em> <!--markdownlint-disable-line-->

React Router 创建路由的 [内置组件](#internal-component), data APIs 由类似 [createBrowserRouter](#createBrowserRouter) 创建的路由才有效

- index 标识当路由未匹配到时默认匹配
- path 路由
- caseSensitive  path 是否区分大小写
- element/Component 当路由匹配时渲染
- handle 当前路由的任意数据, 作用同 [useMatches](#useMatches)

- 使用对象方式创建

```jsx
import {createBrowserRouter} from 'react-router-dom';
const router = createBrowserRouter([
  {
    path: '/',
    element: <Root />,
    errorElement: <ErrorPage />,
    loader: async ({request, params}) => {
      return fetch();
    },
    action: async ({request}) => {
      return update(await request.formData());
    },
    children: []
  }
]);
```

- 使用 JSX 元素创建

```jsx
import {createBrowserRouter, createRoutesFromElements, Route} from 'react-router-dom';

const router = createBrowserRouter(createRoutesFromElements(
  <Route
    path="/"
    element={<Root/>}
    errorElement={<ErrorPage/>}
    lazy={() => import('./a')}
    loader={async ({request, params}) => {
      return fetch();
    }}
    action={async ({request}) => {
      return update(await request.formData());
    }}
  >
    <Route index path="" element={<DashBoard/>}/>
  </Route>
))
```

#### Route.action <em id="Route.action"></em> <!--markdownlint-disable-line-->

当 React Router 抽象了异步 UI 和重新验证的复杂性时, 为应用程序提供了一种使用简单的 HTML 和 HTTP 语句执行数据更改的方法

每当应用程序向路由发送 non-get(POST, PUT, PATCH, DELETE) 请求的时候都会被调用

动态路由参数分别传递给 [loader](#Route.loader), [useMatch](#useParams), [useParams](#useParams)

- request  request 请求实例
- params 动态路由参数

```jsx
import {createBrowserRouter, createRoutesFromElements, Route} from 'react-router-dom';
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      path="/projects/:id/delete"
      action={async ({request, params}) => {
        console.log(params.id);
        const formData = request.formData();
        return deleteProjectById(params.id);
      }}
    >
      {/* .... */}
    </Route>
  )
)
```

#### Route.loader <em id="Route.loader"></em> <!--markdownlint-disable-line-->

组件渲染之前调用定义的 loader 函数并将返回的数据传入 React 元素

动态路由参数分别传递给 [action](#Route.action), [useMatch](#useMatch), [useParams](#useParams)

- params 动态路由参数
- request request 请求实例
- hydrate 服务器端渲染时处理 hydrate 数据

```jsx
import {createBrowserRouter, createRoutesFromElements, Route, useLoaderData} from 'react-router-dom';
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      path="/projects/:id"
      element={<Projects/>}
      loader={async ({request, params}) => {
        console.log(params.id);
        const res = await fetch();
        if(res.status == 404) {
          throw new Response('Not Found', {status: 404});
        }
        return res.json();
      }}
    >
      {/* ... */}
    </Route>
  )
)
function Projects(){
  const projects = useLoaderData();

  return (
    projects
  )
}
```

#### Route.lazy

路由懒加载

```jsx
import {createBrowserRouter, createRoutesFromElements, Route} from 'react-router-dom';
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      path="/"
      element={<Layout/>}
    >
      <Route path="a" lazy={() => import('./a')} />
      <Route path="b" lazy={() => import('./b')} />
    </Route>
  )
)
```

#### Route.shouldRevalidate

如果定义了此函数, 将在路由的 loader 调用之前执行此函数验证新数据, 如果返回 false 则不在调用 loader 并且保持当前页面数据不变

#### Route.errorElement/errorBoundary

当组件的 loader, action 或者在渲染过程中抛出错误时代替 element 显示

```jsx
import {createBrowserRouter, createRoutesFromElements, Route} from 'react-router-dom';
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      errorElement={<ErrorElement/>}
      loader={async ({request, params}) => {
        const res = await fetch();
        if(res.status == 404){
          throw new Response('Not Found', {status: 404});
        }
        const json = res.json();
        return {json}
      }}
    >
      {/* ... */}
    </Route>
  )
)
```

#### Route.hydrateFallbackElement/hydrateFallback

初始化服务器端渲染的内容没有被 hyrate 的组件, 如果未使用类似 [createBrowserRouter](#createBrowserRouter) 创建的路由则无效, 通常 SSR 的应用不会使用此项

```jsx
import {createBrowserRouter} from 'react-router-dom';
const router = createBrowserRouter([
  {
    id: 'root',
    path: '/',
    loader: rootLoader,
    Component: Root,
    children:[
      {
        id: 'invoice',
        path: 'invoice/:id',
        loader: invoiceLoader,
        Component: Invoice,
        hydrateFallback: InvoiceFallback
      }
    ]
  }
],
{
  future:{
    v7_partialHydration: true,
  },
  hydrationData:{
    root:{
      // ...
    }
  }
})
```

### React Router 内置组件 <em id="internal-component"></em> <!--markdownlint-disable-line-->

#### Await <em id="Await"></em> <!--markdownlint-disable-line-->

用于呈现具有自动错误处理功能的延迟值

- children React 元素或者一个函数
- resolve 返回一个 Promise 当延迟值被 resolve 后渲染
- errorElement 当 resolve 被 reject 后渲染

```jsx
import {Await, useLoaderData, defer, Route} from 'react-router-dom';
<Route
  loader={async () => {
    let book = await getBook();
    let reviews = getviews();

    return defer({book, reviews});
  }}
  element={<Book/>}
>
  {/* ... */}
</Route>

function Book(){
  const {book, reviews} = useLoaderData();
  return (
    <div>
      <h1>title</h1>
      <Await resolve={reviews}>
        <Review/>
      </Await>
    </div>
  )
}
```

#### Form

围绕普通 HTML 表单的包装器, 模拟浏览器进行客户端路由和数据更改

- action
- method
- navigate 标识表单默认提交行为提交之后的动作是否跳转
- fetchKey
- replace 标识表单提交行为替换当前历史记录栈
- relative 标识表单提交的后的跳转路径
- reloadDocument 标识跳过 React Router 的表单提交行为并使用浏览器内置的表单默认行为
- state
- preventScrollReset 标识表单提交行为是否滚动页面位置

#### Link <em id="Link"></em> <!--markdownlint-disable-line-->

路由导航

- relative 相对路径, 默认为 Route 的相对层级
- preventScrollRest 标识是否滚动到页面顶部
- replace 标识是否替换当前历史记录栈
- state 任何状态
- reloadDocument

#### NavLink

特殊的 Link, 可以标识当前活动状态的导航

- className 通过函数自定义样式
- style 通过函数自定义样式
- children
- end 改变路由匹配逻辑, 当前路由是否以 to 结尾
- caseSensitive 是否区分大小写
- aria-current
- reloadDocument

```jsx
import {NavLink} from 'react-router-dom';
<NavLink
  to="/message"
  className={({isActive, isPending, isTransitioning}) => {
    return isPending ? 'pending' : isActive ? 'active' : ''
  }}
  style={({isActive, isPending, isTransitioning}) => {
    return {
      fontWeight: isActive ? 'bold' : '',
      color: isPending ? 'red' : 'black',
      viewTransitionName: isTransitioning ? 'slide' : ''
    }
  }}
/>
```

#### Navigate

当组件渲染后改变当前的路由, 通常用在 class 组件中, 建议使用 [useNavigate](#useNavigate) Hook

- to 跳转的目标路由
- replace 是否使用替换模式
- state 任何状态
- relative

#### Outlet

渲染嵌套子路由

```jsx
function DashBoard(){
  return (
    <div>
      <h1>DashBoard</h1>
      {/* ... */}
      <Outlet/>
    </div>
  )
}

function App(){
  return (
    <Routes>
      <Route path="/" element={<DashBoard/>}>
        <Route
          path="message"
          element={<DashBoardMessage/>}
        />
        <Route path="tasks" element={<DashBoardTasks/>}/>
      </Route>
    </Routes>
  )
}
```

#### [Route](#Route)

React Router [内置组件](#internal-component)

#### Routes

匹配组件内的 Route, 通常用于不使用 [createBrowserRouter](#createBrowserRouter) 创建的 Route

#### ScrollRestoration

在渲染完成之后模拟浏览器在位置更改时的滚动恢复, 以确保滚动位置恢复到正确的位置

### React Router 内置 Hook

#### useActionData

获取上一个导航操作结果的返回值, 如果没有提交操作则返回 undefined

#### useAsyncError

获取最近的 [Await](#Await) 组件被 rejection 的结果

#### useAsyncValue

获取最近的 [Await](#Await) 组件被 resolved 的结果

#### useBeforeUnload

当用户离开页面时 (window.onbeforeunload) 保存重要的数据

#### useBlocker

阻止用户离开当前页面, 并呈现自定义 UI 提示用户允许确认导航

- state 当前 blocker 的状态
  - unblocked 空闲没有阻止状态
  - blocked 阻止状态
  - proceeding 正在从阻断器中前进
- proceed() 允许跳转
- reset() 重置 blocker 状态并留在当前位置

```jsx
const blocker = useBlocker();
```

#### useFetcher

不想在更改 URL 的情况下调用 [loader](#Route.loader), [action](#Route.action)获取页面的数据并重新验证, 或者需要同时进行多个更新

与服务器的许多交互不是导航事件, useFetcher 允许将 UI 插入到操作或 [loader](#Route.loader) 中而不引起导航

- key 默认为 内置组件 生成唯一的 key

- fetcher.Form 像 Form [内置组件](#internal-component) 一样, 只是不会引起导航

- fetcher.state 标识当前 Fetcher 的状态
  - idle 空闲
  - submiting 由 fetcher 使用 post, put, patch, delete 提交正在调用路由操作
  - loading fetcher 正在调用 fetcher.load 或者在单独提交或调用用 `useRevalidator` 之后重新验证
- fetcher.data 获取从 [loader](#Route.loader) 或 [action](#Route.action) 加载的数据
- fetcher.formData 当使用 fetcher.Form 和 `fetcher.submit()` 时, formData 可用
- fetcher.json 当使用 `fetcher.submit(data, {formEnctype: 'application/json'})` 提交时可用
- fetcher.text 当使用 `fetcher.submit(data, {formEnctype: 'text/plain'})` 提交时可用
- fetcher.formAction 提交时的 form 的 url
- fetcher.formMethod 提交时的方法 get, post, put, patch, delete

- fetcher.load(href, options) 从 [loader](#Route.loader) 中获取数据
- fetcher.submit(data, options?) 包含了 [useSubmit](#useSubmit) 调用的实例, 接收和 [useSubmit](#useSubmit) 相同的参数

```jsx
import {useEffect} from 'react';
import {useFetcher} from 'react-router-dom';
function SomeCompoent(){
  const fetcher = useFetcher({key: 'new-key'});

  useEffect(() => {
    fetcher.submit(data, options);
    fetcher.load(href);
  },[fetcher]);

  // 渲染的表单不会引起导航 
  return （
    <fetcher.Form action="/fetcher-action" method='post'>
      <button type="submit" onclick={(e) => {
        if(fetcher.state === 'idle' && !fetcher.data){
          fetcher.submit(fetcher.formData?.get('username'), {formEnctype: 'application/json'});
        }
      }}>Submit</button>
      <p>fetcher.formAction {fetcher.formAction}</p>
      <p>fetcher.formMethod {fetcher.formMethod}</p>
      {fetcher.json ? (<p>{fetcher.json}</p>) : (<p>json: null</p>)}
      {fetcher.data ? (<div>{fetcher.data}</div>) : (<div>loading data...</div>)}
    </fetcher.Form>
  ）  
}
```

#### useFetchers

获取除了 load, submit, Form 属性的 fetcher 数组

#### useFormAction

用在 Form [内置组件](#internal-component) 内部自动解析当前路由的默认和相关操作

- 可以直接计算当前的 formAction
- 也可以用在 [useSubmit](#useSubmit) 或者 `fetcher.submit` 中

```jsx
import {useFormAction} from 'react-router-dom';

function DeleteButton(){
  const formAction = useFormAction('destroy');
  return (
    <button
      formAction={formAction}
      formMethod="post"
    >
      Delete
    </button>
  )
}
```

```jsx
const submit = useSubmit();
const formAction = useFormAction('delete');
submit(formData, {formAction});
```

#### useHref

#### useInRouterContext

返回组件是否在 Router 的上下文环境中渲染的

#### useLinkClickHandler

获取 Link 的 click 事件句柄

#### useLoaderData

获取路由 [loader](#Route.loader) 返回的数据, 当路由 loader 被调用之后, 数据将自动重新验证并从 loader 中返回最新结果

useLoaderData 不会启动获取, 只读取 React Router 内部管理的结果

```jsx
import {StrictMode} from 'react';
import {createRoot} from 'react-dom/client';
import {useLoaderData, createBrowserRouter, createRoutesFromElements, RouterProvider} from 'react-router-dom';

function Albums(){
  const albums = useLoaderData();
  // ...
  return <div>Albums</div>;
}
const router = createBrowserRouter(createRoutesFromElements(
  <Route
    path="/"
    element={<Albums />}
    loader={async ({request, params}) => {
      return fakeFecth();
    }}
  />
));
createRoot(document.getElementById('root')).render(
  <StrictMode>
    <RouterProvider router={router}/>
  </StrictMode>
)
```

#### useLocation

获取当前 location 的对象

- location.hash
- location.key
- location.pathname
- location.search
- location.state 通过 [\<Link state/\>](#Link) 或者 [navigate](#useNavigate) 创建的

#### useMatch <em id="useMatch"></em> <!--markdownlint-disable-line-->

返回给定路径相对于当前位置上匹配的数据

动态路由参数分别传递给 [loader](#Route.loader), [action](#Route.action), [useParams](#useParams)

```jsx
import {useMatch, useParams} from 'react-router-dom';

function Random(){
  const match = useMatch('/projects/:projectId/tasks/:taskId');
  const params = useParams();

  console.log(match.params.projectId);
  console.log(match.params.taskId);

  console.log(params.projectId);
  console.log(params.taskId);
}
```

#### useMatches <em id="useMatches"></em> <!--markdownlint-disable-line-->

获取当前页面匹配到的路由信息

#### useNavigate <em id="useNavigate"></em> <!--markdownlint-disable-line-->

返回一个 navigate 函数, 能够以编程式导航, 该函数接收两个参数

- to 跳转的目标路由
- options
  - replace
  - state
  - preventScrollReset
  - relative

```jsx
import {useNavigate} from 'react-router-dom';

function useLogoutTimer(){
  const userIsInactive = useFakeInactive();
  const navigate = useNavigate();
  useEffect(() => {
    if(userIsInactive){
      fake.logout();
      naviagte('/session-time-out', {state: {token: 'token'}});
    }
  },[userIsInactive]);
}
```

#### useNavigation

获取当前页面的所有导航信息

- navigation.state
- navigation.location
- navigation.formData
- navigation.json
- navigation.text
- navigation.formAction
- navigation.formMethod
- navigation.formEnctype

#### useNavigationType

返回当前页的导航类型

```jsx
type NavigationType = 'POP' | 'PUSH' | 'REPLACE';
```

#### useParams <em id="useParams"></em> <!--markdownlint-disable-line-->

返回当前 url 中被 Route 匹配到的动态路由参数对象

动态路由参数分别传递给 [loader](#Route.loader), [action](#Route.action), [useMatch](#useMatch)

```jsx
function Books(){
  const {id} = useParams();
}
<Route
  path="/books/:id"
  element={<Books/>}
/>
```

#### useResolvedPath

返回给定 to 相对于当前位置的 pathname

#### useRevalidator

返回一个验证器对象, 允许重新验证数据

- revalidator.state
- revalidator.revalidate()

#### useRouteError <em id="useRouteError"></em> <!--markdownlint-disable-line-->

用在 errorElement 内部, 捕获由 [action](#Route.action), [loader](#Route.loader), 或者渲染期间抛出的错误

```jsx
import {useRouteError, isRouteErrorResponse, Route, json} from 'react-router-dom';

function ErrorBoundary(){
  const error = useRouteError();

  if(isRouteErrorResponse(error)){
    return (
      <div>
        <h1>Oops!</h1>
        <h2>{error.status}</h2>
        <p>{error.statusText}</p>
        {error.data?.message && <p>{error.data.message}</p>}
      </div>
    )
  } else {
    return <div>Oops</div>;
  }
}

<Route
  errorElement={<ErrorBoundary/>}
  action={async () => {
    throw json(
      {message: 'email is required' },
      {status: 400}
    )
  }}
/>
```

#### useRouteLoaderData

路由树上任何位置的当前渲染路线上的数据都可用, 对于树深层需要来自更高层路由的数据的组件以及需要树深层的子路由的数据的父路由非常有用

```jsx
import {useRouteLoaderData} from 'react-router-dom';

function SomeComp(){
  const user = useRouteLoaderData('root');
  // ...
}
createBrowserRouter([
  {
    path: '/',
    loader: () => fetchUser(),
    element: <Root />
    id: 'root',
    children: [
      {
        path: 'jobs/:jobId',
        loader: loaderJob,
        element: <JobListing />
      }
    ]
  }
])
```

#### useRoutes

相当于 Routes [内置组件](#internal-component) 的函数版本

#### useSearchParams

读取或修改当前 URL 的参数部分

```jsx
import {useSearchParams} from 'react-router-dom';

function App(){
  const [searchParams, setSearchParams] = useSearchParams();

  function handleSumbit(e){
    e.preventDefault();
    // 序列化字段
    const params = serializeFormQuery(e.target);
    setSearchParams(params);
  }
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input type="text" name="username"/>
      </form>
    </div>
  )
}
```

#### useSubmit <em id="useSubmit"></em> <!--markdownlint-disable-line-->

Form 表单提交的命令版本

- submit(data, options?) 手动提交方法
  - options 支持 form 表单的大多数属性

```jsx
import {useSubmit, Form} from 'react-router-dom';
function SearchFiled(){
  let submit = useSubmit();

  // 每次表单改动时提交 
  return (
    <Form onChange={(e) => {
      submit(null, {method: 'post', action: '/change'});
    }}>
      <input type="text" name="search"/>
      <button type="submit">Search</button>
    </Form>
  )
}
```

### React Router API

#### json

格式化数据

```jsx
import {json} from 'react-router-dom';

const loader = async () => {
  const data = fetchData();
  return json(data);
}
```

#### redirect

路由重定向

```jsx
import {redirect} from 'react-router-dom';

const loader = async () => {
  const res = await fetch();
  if(res.status == 401){
    return redirect('/login')
  }
  return null
}
```

#### redirectDocument

触发一个文档级别的重定向, 而不是基于客户端导航, 通常用于从一个应用跳转到另一个应用

#### replace

导航跳转替换当前的历史记录栈

#### createRoutesFromElements

使用 JSX 元素创建路由, 简写形式是 `createRoutesFromChildren`

#### createSearchParams

`new URLSearchParams(init)` 的包装写法

#### defer

延迟 [loader](#Route.loader) 的返回值

```jsx
import {defer} from 'react-router-dom';

const loader = async () => {
  let res = await fetch();

  return defer({name: 'zhangsan', age: 18});
}
```

#### generatePath

根据动态路由参数生成 url

```jsx
import {generatePath} from 'react-router-dom';

generatePath('/users/:id/:name', {id: 42, name: 'zhangsan'}); // /users/42/zhangsan
```

#### isRouteErrorResponse

判断是否是由 [useRouteError](#useRouteError) 捕获的路由错误

#### matchPath

将路由路径模式与 URL 路径进行匹配并返回有关的匹配信息, 否则返回 null

#### matchRoutes

执行一个路由匹配算法从给定的 routes 集合中找到匹配的路由并返回

#### renderMatches

渲染 matchRoutes 匹配结果中的 React 元素

#### resolvePath

根据给定的 to 解析为具有绝对路径的真实 path 对象

## Redux

- dispatch 只能处理同步的 action

- createStore  创建一个 Redux 存储实例
- combineReducers 将多个 reducer 函数合并成为一个更大的 reducer
- applyMiddleware 将多个中间件组合成一个 store 增强器
- compose 将多个 store 增强器合并成一个单一的 store 增强器

### @reduxjs/toolkit(RTK)

消除手写 Redux 逻辑中的样板代码, 防止常见错误, 并提供简化标准的 Redux 任务的 API

- [@reduxjs/toolkit/query](#RTK-Query) 独立可选的入口, 允许定义端点(REST, GraphQL或任何异步函数)并生成 reducer 和中间件来完整管理数据获取, 加载状态更新和结果缓存, 还可以自动生成 React Hooks, 可用于组件获取数据

#### configureStore

通过单个函数调用设置一个完善的 Redux Store, 包括合并 reducer, 添加 thunk 中间件以及设置 Redux Devtools 集成, 与 createStore 相比更容易配置

- reducer
  - 如果是一个函数, configureStore 直接使用其作为根 reducer
  - 如果是一个 slice reducers 的对象, configureStore 将使用 combineReducers 合并此对象并自动创建根 reducer
- middleware 函数, 接收 `getDefaultMiddleware` 函数作为参数, 并返回一个中间件数组, 如果未提供, configureStore 将调用 `getDefaultMiddleware` 设置中间件数组
- devTools 是否设置 Redux Devtools, 默认 true
- preloadedState 初始化状态
- enhancers 增强器函数, 和 middleware 参数作用类似

```jsx
import {configureStore} from '@reduxjs/toolkit';

const store = configureStore({
  reducer: {
    // ...
  }
});
```

##### 特点

- slice reducers 自动传递给 combineReducers
- 自动添加了 `redux-thunk` 中间件
- 添加了 Devtools 中间件来捕获更多意外的变更
- 自动设置了 Redux Devtools Extension
- 中间件和 Devtools 增强器被组合在一起添加到了 store 中

#### createAction <em id="createAction"></em> <!--markdownlint-disable-line-->

用于创建 action 的辅助函数

- type 字符串, 标识 action
- prepareAction() 可选, 函数, 接收任意个参数作为 action 的 payload 的值

返回值: actionCreator

- actionCreator.match 函数可以区分 action 是否是同一类型, TypeScript 中可以识别 action 中 payload 的类型

```tsx
import {createAction} from '@reduxjs/toolkit';

const actionCreator = createAction(type, prepareAction?);

// action 类型常量
function increment(amount: number){
  return {
    type: 'INCREMENT',
    payload: amount
  }
}
const action = increment(3);
// {type: 'INCREMENT', payload: 3}

// action 创建函数
const increment = createAction('INCREMENT', (text: string, age: number) => {
  return {
    type: "INCREMENT",
    payload: {
      text: text,
      age: age,
      id: nanoid(),
      createAt: new Date()
    }
  }
});
const action = increment('hello createAction', 18);
// {type: "INCREMENT", payload: {text: "hello createAction", age: 18, id, createAt}}

const increment = createAction<number>('INCREMENT');
function someFn(action: Action){
  if(increment.match(action)){
    // action.payload can be used as `number` here
  }
}
```

#### createReducer <em id="createReducer"></em> <!--markdownlint-disable-line-->

一个简化创建 reducer 函数的工具, 内部使用 Immer 通过在 reducer 中编写可变代码, 大大简化了不可变的更新逻辑, 并支持将特定的操作类型直接映射到 case reducer 函数

- initialState 初始化状态, 可以是一个返回 state 的函数
- builderCallback 回调函数接收一个 `builder` 对象通过 addCase 方法添加 reducer <em id="builderCallback"></em> <!--markdownlint-disable-line-->
  - addCase() 接收两个参数, 调用必须在 `addMatcher` 和 `addDefaultCase` 之前
    - actionCreatorOrType 指定 action.type
    - reducer
  - addMatcher() 匹配传入的 action, 调用必须在 `addCase` 之后和 `addDefaultCase` 之前
    - matcher() 匹配函数, 匹配传入的所有可能的 action.type, 并按定义的顺序调用
    - reducer
  - addDefaultCase() 添加默认的 reducer
    - reducer

返回值: reducer 函数

- 包含 `getInitialState` 函数, 调用 `getInitialState` 返回初始状态, 通常用于测试或者配合 React [useReducer](#useReducer) Hook

```jsx
import {createReducer} from '@reduxjs/toolkit';

const reducer = createReducer(initialState, builderCallback);

// 普通 reducer
function coutenReducer(state = initialState, action){
  switch(aciton.type){
    case 'ADD':
      return {...state, value: action.payload}
    case 'DELETE':
      return {...state, value: action.payload}
    default:
      return {...state}
  }
}

// createReducer
const counterReducer = createReducer(initialState, builder => {
  builder.addCase('ADD', (state, action) => {
    state.value++; // immer 创建的 state 副本, 直接修改
  }).addCase('DELETE', (state, action) => {
    state.value--;
  }).addCase('ADD_BY_AMOUNT', (state, action) => {
    state.value += action.payload;
  }).addMatcher((action) => isMatchedAction(action.type), (state, action) => {
    // ...
  }).addDefaultCase((state, action) => {
    // ...
  });
})
```

#### createSlice <em id="createSlice"></em> <!--markdownlint-disable-line-->

支持使用 Immer 库编写 reducer, 接受一个初始状态, 对象或者 reducer 函数, 并自动创建一个与 reducer 和 状态对应的 [actionCreator](#createAction), 在内部调用 `createAction` 和 `createReducer`

- name  标识 state, 将作为生成的 [actionCreator](#createAction) 的前缀
- initialState 初始化状态
- reducers reducer 配置
  - 如果需要自定义 case Reducer, 每个 reducer 将是一个具有 prepare 函数 和 reducer 函数的对象
    - prepare()
    - reducer
  - 如果是一个函数, 将接收一个 create 对象, 具有三个方法
    - create.reducer 标准的 reducer
    - create.prepareReducer 自定义 actionCreator 的 payload
    - create.asyncThunk 创建异步的函数代替 actionCreator
      - pending
      - fulfilled
      - rejected
- extraReducers 函数, 处理自己创建的 [actionCreator](#createAction) 之外的情况, 如处理异步请求的状态, 同 [builderCallback](#builderCallback)
- reducerPath 标识 slice 的位置, 默认 name
- selectors 接收 state 作为第一个参数和剩余的参数并返回指定结果

返回值, 包含上面的部分属性

- reducer
- actions
- caseReducers
- getInitialState()
- selectSlice 关联自动创建的一个 selector
- getSelectors()
- injectInfo() 注入 slice

```jsx
import {createSlice, configureStore} from '@reduxjs/toolkit';

const counterSlice = createSlice({
  name: 'counter',
  initialState: {value: 0},
  // reducers 为一个对象
  reducers: {
    increment(state, action) {
      state.value++;
    },
    decrement(state, action){
      state.value--;
    },
    // case reducer, prepareAction
    incrementByAmout: {
      reducer(state, action){
        state.value += action.payload.value;
      },
      prepare(text: string){
        return {payload: {text: text, value: 100}}
      }
    }
  },
  // reducers 为一个函数, 接收一个 create 对象 
  // reducer, prepareReducer, asyncThunk
  reducers(create){
    return {
      increment: create.reducer(state, action) => {
        state.value++;
      },
      decrement: create.reducer(state, action) => {
        state.value--;
      },
      incrementByAmount: create.prepareReducer((text: string) => {
        return { payload: {text: text, value: 100}}
      }, (state, action) => {
        state.value += action.payload.value;
      }),
      fetchTodo: create.asyncThunk(async (id: string, thunkApi) => {
        const res = await fetch(thunkApi);
        return (await res.json()) as Item
      },{
        pending: state => {
          state.loading = true;
        },
        rejected: state =>{
          state.loading = false;
        },
        fulfilled: (state, action) => {
          state.loading = false;
          state.todos.push(action.payload);
        }
      })
    }
  },
  extraReducers(builder){
    builder.addCase('INCREMENT', (state, action) => {
      state.value++;
    })
  }
});
const store = configureStore({
  reducer: {
    counter: counterSlice.reducer
  }
})
store.dispatch(counterSlice.actions.increment());
sotre.dispatch(counterSlice.actions.decrement());
store.dispatch(counterSlice.actions.incrementByAmount({value: 10}));

store.dispatch({type: 'counter/increment'})
store.dispatch({type: 'counter/decrement'})
```

两种获取 selector 的方式

- selectors

```jsx
const counterSlice = createSlice({
  name: 'counter',
  initialState: { value: 0 } satisfies CounterState as CounterState,
  reducers: {
    // ...
  },
  selectors: {
    selectValue: (sliceState) => sliceState.value,
  },
});
console.log(counterSlice.selectSlice({ counter: { value: 2 } })) // { value: 2 }

const { selectValue } = counterSlice.selectors
console.log(selectValue({ counter: { value: 2 } })) // 2
```

- getSelectors()

```jsx
const { selectValue } = counterSlice.getSelectors(
  (rootState: RootState) => rootState.aCounter,
)
console.log(selectValue({ aCounter: { value: 2 } })) // 2

const {selectValue} = counterSlice.getSelectors();
console.log(selectValue({value: 2})) //  2
```

dispatch 提交

- dispatch 提交 action 时, 如果参数是一个 action 对象形式, 则会忽略 case reducer 中配置的 prepare 方法

```jsx
const counterSlice = createSlice({
  name: 'counter',
  initialState: {
    count: 0,
  },
  reducers: {
    incrementByAmount:{
      reducer(state, action){
        state.count += action.payload;
      },
      prepare(val){
        return {payload: val + 2};
      }
    }
  }
});
// action 对象方式提交会忽略 case redcuer 的 prepare 方法
dipatch({type: 'counter/incrementByAmount', payload: 1});
```

#### combineSlices

合并多个 slice 为一个 reducer, 并允许初始化后更多的 reducer 注入

返回值

- withLazyLoadedSlices() 向 state 添加声明的 slice
- inject(slice, options) 添加 slice
  - options.overrideExisting 布尔值, 标识是否替换已存在的 slice
- selector() 将 reducer 包装在代理中以确保在当前状态未定义的情况下都能恢复到其初始状态

```jsx
import {combineSlices} from '@reduxjs/toolkit';

const lazeSlice = createSlice({
  name: 'counter',
  initialState: {value: 0}
});

const rootReducer = combineSlices(staticSlice, userSlice);
const injectReducer = rootReducer.inject(lazySlice);
// OR
const injectSlice = lazySlice.injectInfo(rootReducer);

const selectCounterValue = (rootState) => rootState.counter?.value // number | undefined
const wrappedSelectCounterValue = injectReducer.selector((rootState) => rooState.counter.value);
console.log(
  selectCounterValue({}), // undefined
  selectCounterValue({counter: {value: 2}}), // 2
  wrappedSelectCounterValue({}), // 0
  wrappedSelectCounterValue({counter: {value: 2}}), // 2
)
```

#### createAsyncThunk

接收一个 [actionCreator](#createAction)和一个回调函数并返回一个 Promise, 同时会创建三个 actionCreator 分别对应 pending, fulfilled, rejected 的状态, 不会生成 reducer

- type actionCreator, 如 `users/requestStatus` 将被创建为
  - pending: `users/requestStatus/pending`
  - fulfilled: `users/requestStatus/fulfilled`
  - rejected: `users/requestStatus/rejected`
- payloadCreator 函数, 将返回一个 promise, 接收两个参数
  - arg 包含了 thunk actionCreator 被 dispatch 时传入的参数, `dispatch(fetchUsers({status: 'active'}))`
  - thunkApi 包含了 thunk 函数的所有参数
    - dispatch()  Redux 的 dispatch 方法
    - getState()  Redux 的 getState 方法
    - extra 传递给 thunk 中间件的参数
    - requestId 自动生成的标识当前请求的唯一 id
    - signal 信号, AbortController.signal
    - rejectWithValue(value, [meta]) 修改当前 promise 的状态为 rejected
    - fulfilledWithValue(value, [meta])  修改当前 promise 的状态为 fulfilled
- options
  - condition(arg,{getState, extra}): boolean | Promise\<boolean\> 用来跳过执行 payloadCreator 和 所有的 dispatch
  - dispatchConditionRejection 布尔值, 如果 condition() 返回 false 所有的 action 都不会 dispatch, 如果想要当 thunk 结束 action 的状态标记为 rejected, 则设置为 true
  - idGenerator(arg): string 默认的 requestId 由 nanoid() 生成, 自定义生成 id 逻辑
  - serializeError(error: unknown) => any 替换内部的 `miniSerializeError` 方法
  - getPendingMeta({arg, requestId}, {getState, extra}): any 创建对象和 `pendingAction.meta` 合并

返回值

- thunk 函数, 带有 3 个状态
  - pending
  - fulfilled
  - rejected

```jsx
import {createAsyncThunk, createSlice} from '@reduxjs/toolkit';

const promise = createAsyncThunk(type, payloadCreator, options?);

const fetchUserById = createAsyncThunk('users/fetchUserById', async(userId: number, {requestId, getState, rejectWithValue}) => {
  try{
    const response = await fetch(userId);
    return response.data;
  }catch(err){
    return rejectWithValue(err.response.data);
  }
},{
  condition(userId, {getState, extra}){
    const {users} = getState();
    const fetchStatus = users.requests[userId];
    if(fetchStatus === 'fulfilled' || fetchStatus === 'loading'){
      // Already fetched or in progress, don't need to re-fetch
      return false;
    }
  }
});
const usersSlice = createSlice({
  name: 'users',
  initialState: { },
  reducers:{},
  extraReducers(builder){
    builder.addCase(fetchUserById.pending, (state, action) => {
      state.status = 'loading';
    }).addCase(fetchUserById.fulfilled, (state, action) => {
      state.status = 'fulfilled';
      state.user = action.payload;
    }).addCase(fetchUserById.rejected, (state, action) => {
      state.status = 'rejected';
    });
  }
});

dispatch(fetchUserById(123));
```

#### createEntityAdapter

生成一组预构建的 reducer 和 seletors, 用于对包含特定类型数据对象实例的规范化状态结构执行 CRUD 操作, 这些 reducer 函数可以作为 case reducer 传递给 [createReducer](#createReducer) 和 [createSlice](#createSlice), 也可以作为 `createReducer` 和 `createSlice` 的辅助函数

- selectId 可选, 函数, 接收一个 entity 实例并返回一个唯一 id, 如果未提供则默认为 entity => entity.id
- sortComparer 可选, 函数, 接收两个 entity 实例, 返回一个标准的 `Array.sort()` 排序之后的结果 (1, 0, -1) 以指示它们的排序相对排序, 如果未提供将不会排序, 也不会保证排序

- addOne/addMany 向 state 添加 items
- setOne/setMany 添加新 items 或替换现有 items
- setAll 替换所有 items
- removeOne/removeMany 根据 ID 删除 items
- removeAll 移除所有 items
- updateOne/updateMany 通过提供部分值更新现有 items
- upsertOne/upsertMany 添加新 items 或更新现有 items

- getInitialState() 如果传入对象参数, 将被合并到 initialState 中并返回
- getSelectors() 生成一组标准的 selector 函数

```jsx
import {createSlice, createAsyncThunk, createEntityAdapter} from '@reduxjs/toolkit';

const todosAdapter = createEntityAdapter({
  selectId: todo => todo.id,
  sortComparer: (a, b) => a.id < b.id
});
const initialState = todosAdapter.getInitialState({loading: 'idle'});

// Thunk 函数
const fetchTodos = createAsyncThunk("todos/fetchTodos", async () => {
  const response = await client.get("/fakeApi/todos");
  return response.todos;
});
const saveNewTodo = createAsyncThunk("todos/saveNewTodo",
  async (text) => {
    const initialTodo = { text };
    const response = await client.post("/fakeApi/todos", { todo: initialTodo });
    return response.todo;
  }
);

const todosSlice = createSlice({
  name: 'todos',
  initialState,
  reducers: {
    todoDeleted: todosAdapter.removeOne, // 根据 id 删除 todo
    completeTodosCleard(state, action) {
      const completeIds = Object.values(state.entities)
        .filter(todo => todo.complete)
        .map(todo => todo.id);
      // 删除所有已完成的 todo
      todosAdapter.removeMany(state, completedIds);
    }
  },
  extraReducers(builder){
    builder.addCase(fetchTodos.pending, (state, action) => {
      state.status = 'loading';
    }).addCase(fetchTodos.fulfilled, (state, action) => {
      state.status = 'idle';
    }).addCase(saveNewTodo.fulfilled, todosAdapter.addOne)
  }
})
```

#### createSelector <em id="createSelector"></em> <!--markdownlint-disable-line-->

函数组件每次重新渲染都会重新执行 selector, createSelector 用于创建带有**记忆化**的 selector, 当给定的 inputSelector 没有发生变化时返回已缓存的 selector

- inputSelectors 创建记忆化 selector 的依赖, 可以是一个函数, 也可以是多个函数组成的数组, 返回值依次作为 resultFn 的参数传入
  - selectorFn 接收 state 作为第一个参数和剩余的参数并返回指定结果
- resultFn 在 inputSelectors 之后调用并依次接收来自 inputSelectors 函数的返回值作为参数并返回结果

返回值, 带有记忆化的函数

```jsx
import {createSelector} from 'reselect';
import {useSelector} from 'react-redux';

const selectTodos = state => state.todos;
const selectTodosStatus = (state, completed) => completed;
const memoizedSelectTodoCount = createSelector([selectTodos, selectTodosStatus], (todos, completed) => {
  return todos.filter(todo => todo.completed === completed).length;
});

function CompletedTodosCount({completed}){
  const matchingCount = useSelector((state) => memoizedSelectTodoCount(state, complete));
  return <div>{matchingCount}</div>
}
function App(){
  return (
    <>
      <span>Number of done todos</span>
      <CompletedTodosCount completed={true}/>
    </>
  )
}
```

#### @reduxjs/toolkit/query <em id="RTK-Query"></em> <!--markdownlint-disable-line-->

独立可选的入口, 允许定义端点(REST, GraphQL或任何异步函数)并生成 reducer 和中间件来完整管理数据获取, 加载状态更新和结果缓存, 还可以自动生成 React Hooks, 可用于组件获取数据

### react-redux

#### Provider

- store
- serverState
- context
- stabilityCheck
- children

```jsx
import {Provider} from 'react-redux';
import {createRoot} from 'react-dom/client';
createRoot(document.getElementById('root')).render(
  <Provider store={store}>
    {/*  */}
  </Provider>
)
```

#### shallowEqual

#### useSelector

使用 selector 函数从 Redux store 中提取数据用于当前组件

使用 [createSelector](#createSelector) 创建记忆化的 selector

- selector
- equalityFn

```jsx
import {useSelector, shallowEqual} from 'react-redux';

const selectedData = useSelector(selectorReturningObject, shallowEqual);
// OR
const selectedData = useSelector(selectorReturningObject, {equalityFn: shallowEqual});

function TodoListItem(props){
  const todo = useSelector(state => state.todos[props.id]);
  return <div>{todo.text}</div>
}
```

#### useDispatch

```jsx
import {useCallback, memo} from 'react';
import {useDispatch} from 'react-redux';

function CounterComponent(){
  const dispatch = useDispatch();
  const incrementCounter = useCallbac(() => {
    dispatch({type:'increment-counter'});
  },[dispatch]);
  return (
    <div>
      <span>CounterComponent</span>
      <MyIncrement onIncrement={incrementCounter}/>
    </div>
  )
}
const MyIncrement = memo(({onIncrement}) => {
  return (<button onClick={onIncrement}>increment counter</button>)
})
```

#### useStore

大多数情况使用 `useSelector`

```jsx
import {useStore} from 'react-redux';

function MyComponent(){
  const store = useStore();

  return <div>{store.getState().todos.length}</div>;
}
```
