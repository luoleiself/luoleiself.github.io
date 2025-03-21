---
title: React-Router.md
date: 2022-02-15 15:37:49
categories:
  - ES
  - React
tags:
  - js
  - jsx
  - React
---

## React Router

[React Router v7](https://reactrouter.com/home) 示例

工作模式

- Declareative 声明式, 使用路由组件匹配 url, 导航路由
- Data 在 React rendering 外配置路由, 支持 data APIs
- Framework 包含 Data mode 模式的全栈框架

```tsx
// npx create-react-router@latest my-react-router-app
// Declareative
createRoot(root).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={ <Root /> } />
    </Routes>
  </BrowserRouter>
);

// Data
const router = createBrowserRouter([
  {path: '/', Component: Root, loader: rootLoader}
]);
createRoot(root).render(<RouterProvider router={router}/>);

// Framework
// routes.ts
export default [
  index('./home.tsx'), // 默认渲染组件
  route('about', './about.tsx'),
  route('dashboard', './dashboard.tsx', [ // 嵌套路由, 添加路由段
    index('./home.tsx'),
    route('settings', './setting.tsx')
  ]),
  layout('./auth/layout.tsx', [ // 使用 布局, 但不会添加路由段
    route('login', './auth/login.tsx'),
    route('register', './auth/register.tsx')
  ]),
  ...prefix("projects", [ // 路由前缀
    index("./projects/home.tsx"),
    layout("./projects/project-layout.tsx", [
      route(":pid", "./projects/project.tsx"),
      route(":pid/edit", "./projects/edit-project.tsx"),
    ]),
  ])
]
```

- 自动将 loaderData, actionData, params, matches 作为 props 传递给组件

```tsx
export async function loader(){} // 路由组件渲染之前执行
export async function clientLoader({params}){} // 客户端获取数据
export async function action(){}
export async function clientAction(){}
export function ErrorBoundary(){} // 路由发生错误时渲染
export function headers(){} // 定义响应头
export function links(){} // 定义页面 head 中的 <link> 元素信息
export function meta(){} // 定义页面 head 中的 meta 信息
// 允许在 useMatches 中向 路由 匹配 中添加任意内容, 像面包屑
export const handle = {} 
// 是否允许弹出路由重新验证不影响其数据的操作, 默认所有路由在操作后都会重新验证
export function shouldRevalidate(){} 
```

以下部分为 React Router v6

<!--more-->

### 路由器

创建路由方式

- 使用 createBrowserRouter 和 RouterProvider
  - 对象形式
  - JSX 元素(createRoutesFromElements)
- 使用 Routes, Route, BrowserRouter|HashRouter 内置组件
- 使用 useRoutes hook 和 BrowserRouter|HashRouter 内置组件

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
const routes = [];
const router = createBrowserRouter(routes, {
  basename: '/app',
  hydrationData: {
    root: {
      // ...
    }
  },
});
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

// const router = createBrowserRouter();

const root = createRoot(document.getElementById('root'))
root.render(
  <StrictMode>
    <RouterProvider router={router} fallbackElement={<SpinnerOfDom/>}/>
  </StrictMode>
);
```

- 使用对象形式创建路由

```jsx
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
```

- 使用 JSX 元素创建路由

```jsx
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
"server.jsx"
import {StrictMode} from 'react';
import {createStaticHandler, createStaticRouter, StaticRouterProvider} from 'react-router-dom/server';
import {renderToString} from 'react-dom/server';

// routes

let handler = createStaticHandler(routes);

app.get('*', async (req, res) => {
  let fetchRequest = createRequest(req, res);
  let context = await handler.query(fetchRequest);

  let router = createStaticRouter(handler.dataRoutes, context);
  let html = renderToString(
    <StrictMode>
      <StaticRouterProvider router={router} context={context} />
    </StrictMode>
  );

  res.send("<!DOCTYPE html>" + html);
});
const listener = app.listen(3000, () => {
  let {port} =  listener.address();
  console.log(`listening on port ${port}`);
});

"client.jsx"
import {StrictMode} from 'react';
import {createBrowserRouter, RouterProvider} from 'react-router-dom';
import {hydrateRoot} from 'react-dom/client';

// routes
let router = createBrowserRouter(routes);
const root = hydrateRoot(
  document.getElementById('root'),
  <StrictMode>
    <RouterProvider router={router}/>
  </StrictMode>
);
```

### Route <em id="Route"></em> <!--markdownlint-disable-line-->

React Router 创建路由的 [内置组件](#internal-component), data APIs 由类似 [createBrowserRouter](#createBrowserRouter) 创建的路由才有效

- index 标识当路由未匹配到时默认匹配
- path 路由
- caseSensitive  path 是否区分大小写
- handle 当前路由的任意数据, 作用同 [useMatches](#useMatches)
- element/component 当路由匹配时渲染, 使用 element 意味着不需要再额外的使用 passProps 风格的方式传递 props

```jsx
// 需要使用其他方式传递 props
<Route path=":userId" component={Profile} passProps={{animate: true}} />
// 或者使用 renderProps 传递 props
// 或者使用 HOC 传递 props
<Route path=":userId" render={(routeProps) => (<Profile routeProps={routeProps} animate={true} />)} />
<Route path=":userId" children={({match}) => (
  match ? <Profile match={match} animate={true} /> : <NotFound /> 
)} />

// 使用 element 传递 props
<Route path=":userId" element={<Profile animate={true} />} />
```

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

每当应用程序向路由发送 non-get(POST, PUT, PATCH, DELETE) 提交时, 都将调用此 action

动态路由参数分别传递给 [loader](#Route.loader), [useMatch](#useParams), [useParams](#useParams)

- request  request 请求实例
- params 动态路由参数

```jsx
import {createBrowserRouter, createRoutesFromElements, Route} from 'react-router-dom';
const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      path="/projects/:id/edit"
      action={async ({request, params}) => {
        console.log(params.id);
        const formData = request.formData();
        return editProjectById(params.id);
      }}
    >
      {/* .... */}
    </Route>
  )
);
```

- 以下几种方式都将调用 Route 的 action

```jsx
import {useFetcher, useSubmit} from 'react-router-dom';

const fetcher = useFetcher();
const submit = useSubmit();

// 以下几种方式都将调用 Route 的 action
<Form method="post" action="/projects"/>;
<fetcher.Form method="put" action="/projects/123/edit" />;
submit(data, {method: 'post', action: '/projects'});
fetcher.submit(data, {method: 'put', action: '/projects/123/edit'})
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
      errorElement={<ErrorElement />}
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

- to
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

匹配组件内的 Route, 用于不使用 [createBrowserRouter](#createBrowserRouter) 创建 Route 的情况

也可以使用 [useRoutes](#useRoutes) Hook 创建路由

```jsx
import {Routes, Route, BrowserRouter} from 'react-router-dom';

function App(){
  return (
    <>
      <header>header</header>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={ <DashBoard/> }>
            <Route path="message" element={ <DashBoardMessage /> } />
            <Route path="tasks" element={ <DashBoardTasks /> } />
          </Route>
          <Route path="team" element={ <Team /> } />
        </Routes>
      </BrowserRouter>
      <footer>footer</footer>
    </>
  )
}
```

#### ScrollRestoration

在渲染完成之后模拟浏览器在位置更改时的滚动恢复, 以确保滚动位置恢复到正确的位置

#### Links

> React Router v7 支持

渲染所有的 link 标签

#### Meta

> React Router v7 支持

渲染所有的 meta 标签

#### PrefetchPageLinks

> React Router v7 支持

为即将导航的另一个页面的资源添加预取标签 \<link rel='prefetch|preload|modulepreload' \>

```tsx
import { PrefetchPageLinks } from 'react-router';

<PrefetchPageLinks page='/absolute/path'/>;
```

### React Router 内置 Hook

#### useActionData

获取上一个导航操作结果的返回值, 如果没有提交操作则返回 undefined

通常用于表单验证错误, 如果表单不正确可以返回错误并让用户重试

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
      navigate('/session-time-out', {state: {token: 'token'}});
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

#### useRoutes <em id="useRoutes"></em> <!--markdownlint-disable-line-->

相当于 Routes [内置组件](#internal-component) 的函数版本

```jsx
import {useRoutes, BrowserRouter} from 'react-router-dom';

function App(){
  const element = useRoutes([
    { path: '/', element: <DashBoard />, children: [
      { path: 'message', element: <DashBoardMessage /> },
      { path: 'tasks', element: <DashBoardTasks /> }
    ]},
    { path: 'team', element: <Team /> }
  ]);
  return (
    <>
      <header></header>
      <BrowserRouter>
        {element}
      </BrowserRouter>
      <footer></footer>
    </>
  )
}
```

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

### React Router Utils

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

#### createPath

> React Router v7 支持

根据传入的 路径、参数、哈希创建 URL

#### parsePath

> React Router v7 支持

解析一个 URL, 作用 和 createPath 相反

#### data

> React Router v7 支持

创建一个包含 status, headers 的 Response
