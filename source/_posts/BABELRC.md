---
title: BABELRC
date: 2021-12-09 16:56:30
categories:
  - ES
tags:
  - ES6
  - Babel
---

# @babel/cli

## install

```bash
npm i -D @babel/cli @babel/core
```

## 运行

### 配置项

- 编译文件

  ```bash
  npx babel index.js # 编译 index.js
  ```

- \-\-watch | \-w 监听文件改变自动编译

  ```bash
  npx babel index.js -w # 编译并监听 index.js
  ```

- \-\-out-file | \-o 输出指定文件名

  ```bash
  npx babel index.js -o index.min.js # 编译 index.js 文件输出到 index.min.js
  ```

- \-\-out-dir | \-d 编译整个目录

  - 编译目录下所有文件输出合并为一个文件

    ```bash
    npx babel src -o index.min.js # 编译 src 目录下所有文件输出到 index.min.js
    ```

  ```bash
  npx babel src -d dist # 编译 src 目录下文件输出到 dist 下
  ```

<!-- more -->

- \-\-source-maps | \-s 编译输出源码映射表

  - true 输出 .map 文件
  - inline 文件末尾追加 data: URL 映射关系表
  - both 以上两种方式都用

  ```bash
  npx babel index.js -w -s # 输出 .map 文件
  npx babel index.js -w -s inline # 文件末尾追加
  ```

- \-\-ignore 忽略某些文件

  ```bash
  npx babel src -d dist --ignore "src/**/*.test.js","lib/**/*.*" # 忽略编译文件
  ```

- \-\-copy-files 复制文件

  ```bash
  npx babel src -d dist --copy-files "libs/**/*.js" # 复制不需要编译的文件
  ```

- 通过管道输入文件

  ```bash
  npx babel -o bundle.min.js < index.js # 读取 index.js 的内容通过管道流编译输出到 bundle.min.js
  ```

- \-\-presets 使用预设

  ```bash
  npx babel index.js -o index.min.js --presets=@babel/preset-env,@babel/preset-react # 使用预设
  ```

- \-\-plugins 使用插件

  ```bash
  # 使用插件
  npx babel index.js -o index.min.js --plugins=@babel/proposal-class-properties,@babel/plugin-transform-runtime
  ```

- \-\-no-babelrc 忽略 .babelrc 配置文件

  ```bash
  npx babel index.js --no-babelrc # 忽略项目中的 .babelrc 配置文件
  ```

- \-\-config-file 自定义配置文件路径

  ```bash
  npx babel index.js --config-file /path/to/.babelrc.json # 自定义配置文件路径
  ```

# @babel/core

## Presets

- @babel/preset-env
- @babel/preset-react
  - @babel/plugin-syntax-jsx
  - @babel/plugin-transform-react-jsx
  - @babel/plugin-transform-react-display-name
- @babel/preset-typescript
  - @babel/plugin-transform-typescript
- @babel/preset-flow
  - @babel/plugin-transform-flow-strip-types
- TC39 规则
  - Stage 0 - Strawman: just an idea, possible Babel plugin.
  - Stage 1 - Proposal: this is worth working on.
  - Stage 2 - Draft: initial spec.
  - Stage 3 - Candidate: complete spec and initial browser implementations.
  - Stage 4 - Finished: will be added to the next yearly release.

```json
{ "presets": ["babel-preset-myPreset", "@babel/preset-env"] }
```

a

## 插件

```json
{ "plugins": ["@babel/plugin-transform-runtime"] }
```

## 集成

- @babel/cli
- @babel/plugin-transform-runtime
- @babel/register 通过 require 钩子方式使用 babel

  ```javascript
  require('@babel/register')({
    presets: ['@babel/preset-env'],
    plugins: ['@babel/plugin-transform-runtime'],
    extensions: ['.es6', '.es', '.jsx', '.js', '.mjs'],
    cache: true,
  });
  ```

- @babel/standalone 提供一个 js 编译环境,不建议在生产环境中使用此工具

## 工具

- @babel/generator AST 转换
- @babel/runtime 包含 babel 模块化运行时帮助程序的库
