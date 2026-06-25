---
title: React-umi.md
date: 2022-02-15 15:37:49
categories:
  - ES
  - React
tags:
  - js
  - jsx
  - React
---

# Umijs

## 配置

- `.umirc.ts` 和 `config/config.ts` 作用相同，前者优先级更高

### 运行时

`app.ts` 为运行时配置

- getInitialState() 在运行时中导出此函数, 其返回值将称为全局初始化状态
- useModel('@@initialState') 获取全局初始化状态

```ts
// src/app.ts
import { fetchInitialData } from '@/services/initial';

export async function getInitialState() {
  const initialData = await fetchInitialData();
  return initialData;
}

// src/pages/index.tsx
import { useModel } from 'umi';

export default function Page() {
  const { initialState, loading, error, refresh, setInitialState } =
    useModel('@@initialState'); // 获取全局初始化状态
  return <>{initialState}</>;
};
```

- layout 修改内置布局的配置
- onRouteChange 初始加载和路由切换时做一些事情
- patchRoutes 路由列表和路由组件映射
- patchClientRoutes 修改被 react-router 渲染前的树状路由表
- qiankun 提供微前端的能力
- render 覆写 render
- rootContainer 修改交给 react-dom 渲染时的根组件

## 路由

- 配置式路由, 导出 routes 配置项定义路由结构
- 约定式路由, 在 `src/pages/` 目录下定义 目录和文件名.tsx 的形式自动解析为路由结构

### 动态路由

- 配置式路由只支持 `:variable` 和 `*` 两种通配符
- 约定式路由只支持 `$` 前缀和 `$` 不指定参数名的形式

```ts
+ pages/
  + foo/
    - $slug.tsx
  + $bar/
    - $.tsx
  - index.tsx

[
  { path: '/', component: '@/pages/index.tsx' },
  { path: '/foo/:slug', component: '@/pages/foo/$slug.tsx' },
  { path: '/:bar/*', component: '@/pages/$bar/$.tsx' },
];
```

### 路由数据预加载

Umi 会自动根据当前路由或准备跳转的路由，并行地发起他们的数据请求，因此当路由组件加载完成后，已经有马上可以使用的数据了.

- 页面组件中, 导出 `clientLoader` 函数
- 页面组件中使用 `useClientLoaderData` Hook

```ts
// pages/.../some_page.tsx

import { useClientLoaderData } from 'umi';

export default function SomePage() {
  const { data } = useClientLoaderData();
  return <div>{data}</div>;
}

export async function clientLoader() {
  const data = await fetch('/api/data');
  return data;
}
```

## 数据流

是一种基于 hooks 范式的轻量级数据管理方案，可以在 Umi 项目中管理全局的共享数据

在 `src/models`，`src/pages/**/models/` 和 `src/pages/**/models/model.ts` 文件引入 Model 文件

```ts
// src/models/userModel.ts
// src/components/Username/index.tsx
import { useModel } from 'umi';

export default function Page() {
  const { user, loading } = useModel('userModel');

  return (
    {loading ? <></>: <div>{user.username}</div>}
  );
}
```

## 微生成器

快速生成页面

```ts
// 命令
umi g[enerate]

// 生成页面
// src/pages/foo.tsx
// src/pages/foo.less
umi g page foo

// 以目录方式生成页面
// src/pages/bar/index.tsx
// src/pages/bar/index.less
umi g page bar --dir

// 嵌套页面
// Write: src/pages/far/far/away/kingdom.tsx
// Write: src/pages/far/far/away/kingdom.less
umi g page far/far/away/kingdom

// 批量生成页面
// Write: src/pages/page1.tsx
// Write: src/pages/page1.less
// Write: src/pages/page2.tsx
// Write: src/pages/page2.less
// Write: src/pages/a/nested/page3.tsx
// Write: src/pages/a/nested/page3.less
umi g page  page1  page2   a/nested/page3
```

快速生成组件

```ts
umi g component bar
// Write: src/components/Bar/index.ts
// Write: src/components/Bar/component.tsx
```

生成 RouteAPI

```ts
// Write: api/films.ts
umi g api films
```

生成 mock

```ts
// Write: mock/acl.ts
umi g mock acl
```

## API

- createSearchParams() 包装 new URLSearchParams(init) 的工具函数，支持使用数组和对象创建

```ts
import { createSearchParams } from 'umi';

// 键值对对象
createSearchParams({ foo: 'bar', qux: 'qoo'}).toString()
// foo=bar&qux=qoo

// 键值元组数组
createSearchParams([["foo", "1"], ["bar", "2"]]).toString()
// foo=1&bar=2
```
- generatePath() 使用给定的带参数的 path 和对应的 params 生成实际要访问的路由

```ts
import { generatePath } from 'umi';

generatePath("/users/:id", { id: "42" }); // "/users/42"
generatePath("/files/:type/*", {
  type: "img",
  "*": "cat.jpg",
}); // "/files/img/cat.jpg"
```

- Helmet 用于在页面中动态配置 head 中的标签，例如 title

```ts
import { Helmet } from 'umi';

export default function Page() {
  return (
    <Helmet>
      <title>Hello World</title>
    </Helmet>
  );
}
```

- matchPath() 将给定的路径以及一个已知的路由格式进行匹配，并且返回匹配结果
- matchRoutes() 将给定的路径以及多个可能的路由选择进行匹配，并且返回匹配结果

```ts
import { matchPath, matchRoutes } from 'umi';
const match = matchPath(
  { path: "/users/:id" },
  "/users/123",
);
// {
//   "params": { "id": "123" },
//   "pathname": "/users/123",
//   "pathnameBase": "/users/123",
//   "pattern": { "path": "/users/:id" }
// }

const match = matchRoutes(
  [
    {
      path: "/users/:id",
    },
    {
      path: "/users/:id/posts/:postId",
    },
  ],
  "/users/123/posts/456",
);
// [
//  {
//    "params": {
//      "id": "123",
//       "postId": "456"
//     },
//     "pathname": "/users/123/posts/456",
//     "pathnameBase": "/users/123/posts/456",
//     "route": {
//       "path": "/users/:id/posts/:postId"
//     }
//   }
// ]
```

- terminal() 在开发阶段在浏览器向 node 终端输出日志的工具

```ts
import {terminal} from 'umi';
// 下面三条命令会在 umi 启动终端上打出用不同颜色代表的日志
terminal.log('i am log level');
terminal.warn('i am warn level');
terminal.error('i am error level');
```

### Hooks

- useAppData() 返回全局的应用数据
- useLocation() 返回当前 location 对象
- useMatch() 返回传入 path 的匹配信息；如果匹配失败将返回 null
- useNavigate() 返回一个可以控制跳转的函数；比如可以用在提交完表单后跳转到其他页面

- useOutlet() 当前匹配的子路由元素， \<Outlet\> 内部使用此 hook
- useOutletContext() 返回 Outlet 组件上挂载的 context

```ts
import { useMatch, useOutlet } from 'umi';

// when url = '/events/12'
const match = useMatch('/events/:eventId');
console.log(match?.pathname, match?.params.eventId);
// '/events/12 12'

const Layout = ()=>{
  const outlet = useOutlet()

  return <div className="fancyLayout">
    {outlet}
  </div>
}
```

- useParams() 返回动态路由的匹配参数键值对对象
- useSearchParams() 读取和修改当前 URL 的 query string, 其返回包含两个值的数组，当前 URL 的 search 参数和用于更新 search 参数的函数
- useResolvedPath() 根据当前路径将目标地址解析出完整的路由信息

```ts
import { useSearchParams, useResolvedPath } from 'umi';

function App() {
  let [searchParams, setSearchParams] = useSearchParams();
  function handleSubmit(event) {
    event.preventDefault();
    setSearchParams(serializeFormQuery(event.target));
  }
  return <form onSubmit={handleSubmit}>{/* ... */}</form>;
}

const path = useResolvedPath('docs')
/* path
{ pathname: '/a/new/page/docs', search: '', hash: '' }
*/
```

- useRouteData() 返回当前匹配路由的数据
- useRoutes() 渲染路由的钩子函数，传入路由配置和可选参数 location, 即可得到渲染结果；如果没有匹配的路由，结果为 null

```ts
import * as React from "react";
import { useRoutes } from "umi";

function App() {
  let element = useRoutes([
    {
      path: "/",
      element: <Dashboard />,
      children: [
        {
          path: "messages",
          element: <DashboardMessages />,
        },
        { path: "tasks", element: <DashboardTasks /> },
      ],
    },
    { path: "team", element: <AboutPage /> },
  ]);

  return element;
}
```

- useRouteProps() 读取当前路由在路由配置里的 props 属性，用此 hook 来获取路由配置中的额外信息

```ts
// .umirc.ts
routes: [
  {
    path: '/',
    custom_key: '1',
  }
]
import { useRouteProps } from 'umi'

export default function Page() {
  const routeProps = useRouteProps()

  // use `routeProps.custom_key`
} 
```

- useSelectedRoutes() 用于读取当前路径命中的所有路由信息, 例如在 layout 布局中可以获取到当前命中的所有子路由信息，同时可以获取到在 routes 配置中的参数
  
```ts
// layouts/index.tsx
import { useSelectedRoutes } from 'umi'

export default function Layout() {
  const routes = useSelectedRoutes()
  const lastRoute = routes.at(-1)

  if (lastRoute?.pathname === '/some/path') {
    return <div>1 : <Outlet /></div>
  }

  if (lastRoute?.extraProp) {
    return <div>2 : <Outlet /></div>
  }

  return <Outlet />
}
```
