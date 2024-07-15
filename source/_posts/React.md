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

## State

- 4 种条件判定是否可作为 state
  - 这个变量是否通过 props 从父组件中获取，如果是，则不是一个状态
  - 这个变量是否在组件的整个生命周期中都保持不变，如果是，则不是一个状态
  - 这个变量是否可以通过其他状态(state)或者属性(props)计算得到，如果是，则不是一个状态
  - 这个变量是否在组件的 render 方法中使用，如果不在，则不是一个状态

## 组件类

### 类组件

#### React.Component

#### React.PureComponent

对 props 和 state 进行浅比较, 跳过不必要的更新, 提高组件性能

### 函数组件

## Hooks

> 以 use 开头的函数被称为 Hook, Hook 比普通函数更为严格, 只能在组件的**顶层**调用 Hook.

### 状态值

- useState
- useReducer

### 生命周期

- useEffect
- useLayoutEffect

### 状态共享

- useContext

### 性能优化

- useMemo
- useCallback

### 属性

- useRef
- useId
