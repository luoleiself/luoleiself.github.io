---
title: BABELRC
date: 2021-12-09 16:56:30
categories:
  - ES
tags:
  - ES6
  - Babel
---

## @babel/cli

### install

```bash
npm i -D @babel/cli @babel/core @babel/preset-env
```

### 运行

#### 参数

- \-\-watch | \-w 监听文件改变自动编译
- \-\-out-file | \-o 输出指定文件名
- \-\-out-dir | \-d 编译整个目录
  - 编译目录下所有文件输出合并为一个文件

```bash
npx babel index.js -w # 编译并监听 index.js

npx babel index.js -o index.min.js # 编译 index.js 文件输出到 index.min.js

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

- \-\-presets 使用预设

  ```bash
  npx babel index.js -o index.min.js --presets=@babel/preset-env,@babel/preset-react # 使用预设
  ```

- \-\-plugins 使用插件

  ```bash
  # 使用插件
  npx babel index.js -o index.min.js --plugins=@babel/proposal-class-properties,@babel/plugin-transform-runtime
  ```

- 通过管道输入文件

  ```bash
  npx babel -o bundle.min.js < index.js # 读取 index.js 的内容通过管道流编译输出到 bundle.min.js
  ```

- \-\-no-babelrc 忽略 .babelrc 配置文件
- \-\-config-file 自定义配置文件路径
- \-\-copy-files 复制文件
- \-\-ignore 忽略某些文件

  ```bash
  npx babel index.js --no-babelrc # 忽略项目中的 .babelrc 配置文件

  npx babel index.js --config-file /path/to/.babelrc.json # 自定义配置文件路径

  npx babel src -d dist --copy-files "libs/**/*.js" # 复制不需要编译的文件

  npx babel src -d dist --ignore "src/**/*.test.js","lib/**/*.*" # 忽略编译文件
  ```

## @babel/core

### Presets

- @babel/preset-env
- @babel/preset-typescript
- @babel/preset-react
- @babel/preset-flow

### 插件

```json
{
  // 开启默认预设
  "presets": [
    [
      "@babel/preset-env",
      {
        "modules": false, // 关闭 esm 转化，统一交由 rollup 处理，防止冲突
        "targets": "> 0.25%, not dead"
      }
    ]
  ],
  "plugins": [
    "@babel/plugin-external-helpers",
    [
      // 开启 babel 各依赖联动，由此插件负责自动导入 helper 辅助函数，从而形成沙箱 polyfill
      "@babel/plugin-transform-runtime",
      {
        "corejs": { "version": 3, "proposals": true },
        "helper": true,
        "regenerator": true
        // "useESModules": true // 关闭 esm 转化，交由 rollup 处理，同上防止冲突 7.13.0 开启废弃
      }
    ]
  ]
}
```

### 集成

- @babel/cli babel 命令行工具
- @babel/polyfill 7.4.0 开始被废弃, JS 标准新增的原生对象和 API 的 shim, 实现上仅仅是 core-js 和 regenerator-runtime 两个包的封装
- @babel/preset-xxx transform 阶段使用到的一系列的 plugin
- @babel/plugin-xxx babel 转译过程中使用到的插件

  - @babel/plugin-transform-runtime 包含 babel 重建模块化运行时助手的插件
    - 为所有辅助函数创建 `@babel/runtime` 模块的引用, 避免编译输出中的重复引用
    - 为代码创建一个沙盒环境, 避免直接的引入垫片而引起的全局环境污染

- @babel/register 通过 require 钩子方式使用 babel 自动转译引用的 js 代码文件
- @babel/standalone 提供一个 js 编译环境,不建议在生产环境中使用此工具

### 工具

- @babel/parser babel 的 js 解析函数
- @babel/core 包含整个 babel 工作的核心模块, 提供了 babel 的转译 API
- @babel/generator 根据 AST 生成代码
- @babel/code-frame 用于生成错误信息, 打印出错误点源代码帧以及指出出错位置
- @babel/runtime 包含 babel 模块化运行时助手的库, 功能类似 babel-polyfill
  - 为每个 js 文件添加运行时辅助函数
  - 可能会在输出文件中注入一些跨文件相同的可能被重复使用的代码
- @babel/template 辅助函数, 用于从字符串形式的代码来构建 AST 树节点
- @babel/traverse 用于对 AST 的遍历, 只要提供功能给 plugins 使用
- @babel/types 此模块包含手动构建 AST 和检查 AST 节点类型的方法
- @babel/helpers 一系列预制的 babel-template 函数, 用于提供给一些 plugins 使用
