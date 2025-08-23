---
title: React-Tailwind.md
date: 2022-02-15 15:37:49
categories:
  - ES
  - React
tags:
  - js
  - jsx
  - React
---

## [react-transition-group](https://reactcommunity.org/react-transition-group/)

### Transition

包含 4 个状态

- entering
- entered
- exiting
- exited

props

- nodeRef 执行动画的关联 DOM 节点, 早期的版本使用 findDOMNode(deprected) 查找 DOM 节点, 报错的替换方案
- in 切换 enter 和 exit 的状态
- appear 控制组件首次挂载时的默认行为
- enter 控制进入的动画
- exit 控制退出的动画
- timeout 动画的时长
- addEvenetListener 添加自定义的事件
- onEnter
- onEntering
- onEntered
- onExit
- onExiting
- onExited

```jsx
import {useState, useRef} from 'react';
import {Transition} from 'react-transition-group';

function App(){
  const [isProp, setInProp] = useState(false);
  const nodeRef = useRef(null);

  return (
    <>
      <Transition nodeRef={nodeRef} in={inProp} timeout={500}>
        <h2>react-transition-group Transition component</h2>
      </Transition>
      <button onClick={() => setInProp(true)}>Click to Enter</button>
    </>
  )
}
```
<!--more-->
### CSSTransition

继承 Transition 的所有 props, 使用 CSS 设置动画

- classNames 当组件在 appear, enter, exit 时应用于组件的动画名称前缀

classNames="my-node" 将会应用以下几种样式

- my-node-appear, my-node-appear-active, my-node-appear-done
- my-node-enter, my-node-enter-active, my-node-enter-done
- my-node-exit, my-node-exit-active, my-node-exit-done

```jsx
import {useState, useRef} from 'react';
import {CSSTransition} from 'react-transition-group';

function App(){
  const [inProp, setInProp] = useState(false);
  const nodeRef = useRef(null);

  return (
    <>
      <CSSTransition nodeRef={nodeRef} in={inProp} timeout={500} classNames="my-node">
        <div ref={nodeRef}>
          react-transition-group CSSTransition component
        </div>
      </CSSTransition>
      <button onClick={() => setInProp(true)}>Click to Enter</button>
    </>
  )  
}

<style>
  .my-node-enter,
  .my-node-appear {
    opacity: 0;
  }
  .my-node-enter-active,
  .my-node-appear-active {
    opacity: 1;
    transition: opacity 500ms;
  }
  .my-node-exit {
    opacity: 1;
  }
  .my-node-exit-active {
    opacity: 0;
    transition: opacity 500ms;
  }
</style>
```

### SwitchTransition

切换两个组件的动画, 需要使用 key 作为 Transition 或 CSSTransition 的 props

- mode 动画的模式 in-out|out-in, 默认 out-in

```jsx
import {useState, useRef} from 'react';
import {SwitchTransition, CSSTransition} from 'react-transition-group';

function App(){
  const [state, setState] = useState(false);
  const helloRef = useRef(null);
  const goodByeRef = useRef(null);
  const nodeRef = state ? goodByeRef : helloRef;

  return (
    <SwitchTransition mode="out-in">
      <CSSTransition
        key={state ? 'goodeBye, world' : 'hello world'}
        nodeRef={nodeRef}
        classNames="fade"
      >
        <button
          ref={nodeRef}
          onClick={() => setState(state => !state)}
        >
          {state ? 'goodBye, world' : 'hello world'}
        </button>
      </CSSTransition>
    </SwitchTransition>
  )
}
```

### TransitionGroup

管理多个动画列表

- component 指定渲染的元素, 默认 div

```jsx
import {useState} from 'react';
import {TransitionGroup, CSSTransition} from 'raect-transition-group';
function App(){
  // ...
  const [books, setBooks] = useState([]);
  return (
    <TransitionGroup component="ul">
      {books.map({id, name, author, date} => {
        return (
          <CSSTransition key={id} timeout={500} classNames="book">
            <li>{name}-{author}-{date}</li>
          </CSSTransition>
        )
      })}
    </TransitionGroup>
  )
}
```

## CSSInJs

### [标签函数](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Template_literals#%E5%B8%A6%E6%A0%87%E7%AD%BE%E7%9A%84%E6%A8%A1%E6%9D%BF)

模板字符串的高级用法

标签函数第一个参数包含一个字符串数组, 其余的参数与表达式相关

```jsx
const person = "Mike";
const age = 28;

function myTag(strings, personExp, ageExp) {
  const str0 = strings[0]; // "That "
  const str1 = strings[1]; // " is a "
  const str2 = strings[2]; // "."

  const ageStr = ageExp > 99 ? "centenarian" : "youngster";

  // 我们甚至可以返回使用模板字面量构建的字符串
  return `${str0}${personExp}${str1}${ageStr}${str2}`;
}

const output = myTag`That ${person} is a ${age}.`;

console.log(output);
// That Mike is a youngster.

// 以下都是等价的写法
fn`some string here`;
fn(['some string here']);

const aVar = 'good';
fn`this is a ${aVar} day`;
fn(['this is a ', 'day'], aVar);
```

### [styled-components](https://styled-components.com/docs/basics)

影响打包性能, 导致 js 文件体积变大

动态生成唯一的 className, 对样式覆写不太友好

```jsx
import styled from 'styled-components';
```

- 基本使用, 使用标签函数创建样式化的组件

```jsx
const Title = styled.h1`
  font-size: 1.5em;
  text-align: center;
  color: #bf4f74;
`;
const Wrapper = styled.section`
  padding: 4em;
  background: yellowlight;
`;

function App(){
  return (
    <Wrapper>
      <Title>title</Title>
    </Wrapper>
  )
}
// 渲染结果
<section>
  <h1>title</h1>
</section>
```

- 样式扩展

```jsx
// 创建 Button 组件
const Button = styled.button`
  color: #bf4f74;
  font-size: 1em;
  margin: 1em;
  padding: 0.25em 1em;
  border: 2px solid #bf4f74;
  border-radius: 3px;
`;
// 基于 Button 组件扩展并覆盖样式
const TomatoButton = styled(Button)`
  color: tomato;
  border-color: tomato;
`;

function App(){
  return (
    <>
      <Button>Normal Button</Button>
      <TomatoButton>Tomato Button</TomatoButton>
    </>
  )
}
// 渲染结果
<button>Normal Button</button>
<button>Tomato Button</button>
```

- 样式化任何组件

```jsx
// react-router-dom's Link Component
const Link = ({className, children}) => (
  <a className={className}>
    {children}
  </a>
);
const StyledLink = styled(Link)`
  color: #bf4f74;
  font-weight: bold;
`;
function App(){
  return (
    <>
      <Link>UnStyled Link</Link>
      <StyledLink>Styled Link</StyledLink>
    </>
  )
}
```

- 支持伪类, 伪元素, 嵌套样式

#### as 更改渲染结果

as props 更改样式化组件呈现的标签或组件

- as 接收标签

```jsx
function App(){
  return (
    <>
      <Button>Normal Button</Button>
      <Button as="a" href="#">Link with Button styles</Button>
      <TomatoButton as="a" href="#">Link with Tomato Button styles</TomatoButton>
    </>
  )
}
// 渲染结果
<button>Normal Button</button>
<a href="#">Link with Button styles</a>
<a href="#">Link with Tomato Button styles</a>
```

- as 接收组件

```jsx
// 接收 props 并将 props.children 反转
const ReverseButton = props => <Button { ...props } children={ props.children.split('').reverse() } />

function App(){
  return (
    <Button>Normal Button</Button>
    <Button as={ ReverseButton }>Custom Button with Normal Button styles</Button>
  )
}
// 渲染结果
<button>Normal Button</button>
<button>selyts nottuB lamroN htiw nottuB motsuC</button>
```

#### 获取 props

添加 `$` 前缀将 props 转换为短暂的 props, 防止样式化组件使用的 props 被传入到底层的 React 组件或者 DOM 元素上

```tsx
const Button = styled.button<{ $primary?: boolean }>`
  background: ${props => props.$primary ? '#bf4f74' : 'white'};
  color: ${props => props.$primary ? 'white' : '#bf4f74'};
  font-size: 1em;
  margin: 1em;
  padding: 0.25em 1em;
  border: 2px solid #bf4f74;
  border-radius: 3px;
`;

function App(){
  return (
    <>
      <Button>Normal Button</Button>
      <Button $primary>Primary Button</Button>
    </>
  )
}
// 渲染结果
<button>Normal</button>
<button>Primary</button> // 和 Normal button 样式相反
```

#### attrs 属性重写

attrs 修改样式化组件的属性, 如果想防止样式化组件使用的 props 被传入到底层 React 组件或者 DOM 元素上, 使用 `$` 前缀修饰 props 将其转换为短暂的 props

```tsx
styled.tagName.attrs({} | Function);

const Input = styled.input.attrs<{ $size?: string }>(props => ({
  type: "text",
  $size: props.$size || '1em'
}))`
  border: 2px solid #BF4F74;

  /* use the dynamically computed prop */
  margin: ${props => props.$size};
  padding: ${props => props.$size};
`;
const PassWordInput = styled(Input).attrs({type: 'password'})`
  border: 2px solid aqua;
`;

function App(){
  return (
    <>
      <Input placeholder="A bigger text input" $size="2em" />
      <PassWordInput placeholder="A normal password input" />
    </>
  )
}
// 渲染结果
<input type="text" placeholder="A bigger text input"/> // margin 和 padding 都为传入的 2em
<input type="password" placeholder="A normal password input"/> // margin 和 padding 都为默认的 1em
```

#### 其他组件和API

- StyleSheetManager 管理样式化组件的辅助组件
- ThemeProvider
  - theme
- ThemeConsumer
- withTheme 高阶组件, 传递的组件将接收到一个包含 theme prop 的 theme 对象

```jsx
import {withTheme} from 'styled-components';

function App(props){
  return (
    <button style={{color: props.theme.color}}>Button</button>
  )
}
export default withTheme(App);
```

- useTheme 获取 ThemeProvider 传递的 theme 的 Hook

```jsx
import {ThemeProvider, ThemeConsumer, useTheme} from 'styled-components';

function App(){
  return (
    <ThemeProvider theme={{color: 'mediumseagreen'}}>
      <Button/>
    </ThemeProvider>
  )
}
function Button(){
  const theme = useTheme();
  const style = {color: theme.color};

  return (
    <button style={style}>
      Button
    </button>
  )
}
// 等价于
function Button(){
  return (
    <ThemeConsumer>
      {theme => <button style={{color: theme.coloer}}>Button</button>}
    </ThemeConsumer>
  )
}
```

- isStyledComponent 判断是否是样式化组件
- createGlobalStyle 创建全局样式
- css 创建样式片段
- keyframes 创建动画的辅助函数

```jsx
import {keyframes} from 'styled-components';
const fadeIn = keyframes`
  0% {
    opacity:0
  }
  100% {
    opacity: 1;
  }
`;
const FadeInButton = styled.button`
  animation: 1s ${fadeIn} ease-out;
`;
```

## tailwindCSS

tailwindcss

@tailwindcss/postcss

postcss

v3 版本支持 `tailwind.config.js` 配置文件

应用

```html
<!-- bg-[#1da1f2] 使用任意值 -->
<div class="3xl:text-lg text-gray-900 bg-white font-display rounded-sm bg-[#1da1f2]">demo</div>

<!-- 使用任意值(arbitrary values) -->
<div class="bg-[#316ff6] lg:top-[117px]">使用任意值</div>

<div class="grid grid-cols-[24rem_2.5rem_minmax(0, 1fr)]">使用任意值</div>
<div class="max-h-[calc(100vh - (--spacing(6)))]">使用任意值</div>

<!-- css 属性 -->
<div class="[mask-type:luminance] hover:[mask-type:alpha]">css 属性</div>

<!-- 样式优先级: 编译后样式末尾添加 !important -->
<div class="bg-red-500!">样式优先级</div>
```

### 指令

- @tailwind 插入样式到 base, components, utilities, variants, v4 不再支持该指令
  - base 用于重置规则或应用于html元素的默认样式
  - components 基于类的样式, 能够使用实用程序覆盖这些样式
  - utilities 实用程序用于小型、单一用途的类, 这些类应始终优先于任何其他样式

- @import CSS 指令, 导入 tailwindcss
- @theme 添加自定义主题变量或重置默认主题

```css
/* 导入 tailwindcss */
/* 使用前缀, 如果存在类名冲突的情况时给 tailwind 生成的 class 和 css 变量添加前缀 */
@import 'tailwindcss' prefix(tw);
/* 给 tailwindcss 所有样式末尾添加 !important */
@import 'tailwindcss' important;

/* 自定义主题变量 */
/* 命名空间颜色变量前缀: --color-*  */
/* 命令空间字体变量前缀: --font-*  */
/* 命令空间间距变量前缀: --spacing-*  */
/* 命令空间圆角变量前缀: --radius-*  */
/* ... */
@theme {
  --color-dblue-300: #0088ff;
  --spacing-50: '50rem';
  --tab-size-github: 8;
}
```

- @source 排除不被 tailwind 自动识别的资源

- @apply 将已有的实用程序应用到自定义样式中
- @layer 指定自定义样式应该放在哪个层中
- @utility 添加实用程序
  - \-\-value(--theme-key-*) 解析一组主题 keys 的应用程序的值
- @variant 在样式中应用 tailwind 的变体
- @custom-variant 添加自定义的变体

```css
@import 'tailwindcss';

/* 添加默认样式 */
@layer base {
  h1 {
    font-size: var(--text-2xl);
  }
}

/* 添加基于类的样式 */
@layer components {
 .card {
    background-color: var(--color-white);
    border-radius: var(--rounded-lg);
    padding: var(--spacing-6);
    box-shadow: var(--shadow-xl);
    @apply rounded-2xl;
  }
}

/* 定义实用程序 */
@utility content-auto {
  content-visibility: auto;
  &::-webkit-scrollbar {
    display: none;
  }
}
@utility tab-* {
  tab-size: --value(integer);
}

/* 添加应用程序到自定义样式中 */
.select2-dropdown {
  @apply rounded-b-lg shadow-md;
}

/* 添加变体 */
.my-element {
  background: white;
  @variant dark {
    background: black;
  }
}
```

```html
<!-- use @layer components difine classes -->
<div class="card">use @layer components difine classes</div>

<!-- use utility -->
<div class="content-auto hover:content-auto">Hello World!</div>

<!-- use `--value(--theme-key-*)` function resolve utility value -->
<!-- --value(type) validate the bare value -->
<div class="tab-2 tab-4">Function utility</div>
```

- @reference 直接在 css 模块, style 标签内, 其他组件内使用 @apply, @variant, 自定义的主题中定义的样式避免重复的 css 输出

```html
<style>
  /* 使用自定义主题 */
  @reference '../../app.css';
  /* 使用默认主题 */
  @reference 'tailwindcss'; 

  h1 {
    @apply text-2xl font-bold text-red-500;
  }
</style>
```

- @config 加载旧版本的 js 配置文件, v4 兼容旧版本
- @plugin 加载旧版本的 js 配置文件中的插件, v4 兼容旧版本

#### 函数

- \-\-alpha() 调整颜色的透明度
- \-\-spacing() 生成主题的间距值

```css
.my-element {
  color: --alpha(var(--color-lime-300) / 50%);
  margin: --spacing(4);
  margin: calc(var(--spacing) * 4);
}
```

- theme() 在编译时允许使用 tailwind 配置的值, 函数在构建时执行， v4 兼容旧版本
- screen() 使用 tailwind 配置的值创建媒体查询

```css
.my-element {
  margin: theme(spacing.12);
}
```

## classnames
