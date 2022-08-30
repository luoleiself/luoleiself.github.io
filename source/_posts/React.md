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

## React.Component

## React.PureComponent

## Hooks
