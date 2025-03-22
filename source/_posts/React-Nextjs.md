---
title: React-Nextjs.md
date: 2022-02-15 15:37:49
categories:
  - ES
  - React
tags:
  - js
  - jsx
  - React
---

RSC(React Server Component)
ISR(Incremental Static Regeneration)

### 文件规范

应用根目录或 src 目录下

- middleware.ts 在请求完成之前在服务器上运行代码, 根据传入的请求修改响应, 对于实现自定义服务器端逻辑非常有用, 配合 `matcher` 使用过滤指定范围的请求
  - Props
    - request
- instrumentation.ts 用于将可观察工具集成到应用程序中, 能够跟踪性能和行为, 并在生产中调试问题

路由文件

- layout.tsx 在多个页面之间共享布局的UI, 能够保持跨路由的状态、交互性, 不会重新渲染.
  - Props
    - children
    - params 动态路由参数, 一个 Promise, Next.js 14 之前是同步的
    - [slot...] 动态插槽
- template.tsx 类似于 layout.tsx 能够包含布局和页面, 当路由发生改变时会重置状态
  - Props
    - children

- route.ts 使用 Web Request 和 Response API 为给定的路由创建自定义请求处理程序, 和 page.tsx 不能同时存在
  - Props
    - request
    - context
- page.tsx 定义路由独有的页面UI
  - Props
    - params 动态路由参数, 一个 Promise, Next.js 14 之前是同步的
    - searchParams 当前 URL 的查询字符串参数, 一个 Promise, Next.js 14 之前是同步的

- loading.tsx 创建基于加载时的状态, 配合 `Suspense` 组件使用

- not-found.tsx 路由未匹配到时渲染的 UI, 默认自动匹配 app 目录下的 not-found.tsx, 嵌套路由下手动调用 notFound 函数渲染局部 not-found.tsx

  ```tsx
  /*
    app
      products
        comments
          page.tsx // 调用 notFound 函数渲染当前路由段下的 not-found.tsx
          not-found.tsx
      layout.tsx
      page.tsx
      not-found.tsx
  */
  ```
<!-- more -->

- error.tsx 允许处理运行时的错误并显示回退 UI
  - Props
    - error
    - reset
- global-error.tsx 处理在根 layout 或 template 抛出的错误, 必须使用 html 和 body 标签, 这个文件将替换根 layout 或 template

```tsx
// app/blog/error.tsx
// 处理在嵌套路由段下抛出的错误
export default function Error({ error }: {
  error: Error;
  reset: () => void
}) {
  return (
    <div className='container text-center text-red-500 text-3xl'>
      blog error
      <br />
      {error.message}
    </div>
  )
}

// global-error.tsx
// 处理在根 layout 或 template 抛出的错误
export default function GlobalError({
  error,
  reset
}: {
  error: Error & { digest?: string };
  reset: () => void
}) {
  return (
    <html>
      <body>
        <div className='text-red-500 text-2xl'>
          global-error
          <br />
          {error.message}<br/>
          {error.digest}<br/>
          <Button type="primary" onClick={() => reset()}>Try again</Button>
        </div>
      </body>
    </html>
  )
}
```

- default.tsx 用于在 Nextjs 在加载完整页面后无法恢复 `插槽` 的活动状态时使用. 刷新页面(硬导航)时, 为与当前 URL 未匹配的子页面渲染内容, 如果不存在则渲染 404
  - Props
    - params 动态路由参数, 一个 Promise, Next.js 14 之前是同步的

```tsx
// 插槽不是路由段, 不能够影响 url, 插槽和正常页面合并之后形成与路由相关的最终页面,
// 因此, 在相同的路由段上不能有单独的静态渲染和动态插槽, 如果有一个插槽是动态的, 则该路由段的所有插槽都必须是动态的

// Link 组件软导航访问 /dashboard 和 /dashboard/visitor 页面时, 页面内容显示正常
// 硬导航 /dashboard/visitor 时, 无法恢复 并行路由(插槽) 状态,
// 需要提供已使用未匹配的 动态插槽 的 default.tsx 渲染内容 
// 否则将渲染 404
/* 
  app
    dashboard
      setting       ✕
        page.tsx
      @team
        default.tsx ✓
        page.tsx
      @blor           // 动态插槽未使用不需要提供 default.tsx
        page.tsx
      @analytics
        page.tsx
        visitor
          page.tsx
      default.tsx   ✓
      layout.tsx
      page.tsx
*/
```

Component hierarchy

```tsx
<Layout>
  <Template>
    <ErrorBoundary fallback={<Error />}>
      <Suspense fallback={<Loading />}>
        <ErrorBoundary fallback={<NotFound />}>
          <Page />
        </ErrorBoundary>
      </Suspense>
    </ErrorBoundary>
  </Template>
</Layout>
```

### 路由结构

- pages 以当前目录下的 文件名 创建路由段, 目录下的 index.tsx 创建页面
  - [fileName] 动态路由, 动态路由参数可以在 `layout.tsx`, `page.tsx`, `route.ts` 和 `generateMetadata` 中获取
    - [...fileName] 截获所有动态路由参数
    - [[...fileName]] 可选的截获所有动态路由参数, 同时会截获不带任何动态参数的路由
- app 以当前目录下的 目录名 创建路由段, 目录下的 page.tsx 创建页面
  - _folderName 私有目录, 当前目录及子目录被 `路由解析 忽略`, 将 \_ 转义后命名目录路由段可正常访问
  - (folderName) 路由分组, 目录名被 `路由解析 忽略`, 使用相同的布局
  
  - @folderName 并行路由, 被 `路由解析 忽略`. 同时或有条件地 在同一 layout.tsx 中渲染一个或多个页面.

    不能够影响 url, 插槽和正常页面合并之后形成与路由相关的最终页面.
  
    使用 `插槽` 渲染页面, 硬导航时无法恢复未匹配路由的插槽的活动状态时使用插槽的 default.tsx 渲染.

  ```tsx
  // 访问 / 同时渲染 app/pages.tsx, @team/page.tsx, @analytics/page.tsx
  export default function Layout({ children, team, analytics }: readonly<{
      children: React.ReactNode;
      team: React.ReactNode;
      analytics: React.ReactNode;
    }>) {
      return (
        <>
          {children}
          <div className="flex justify-center items-center">
            {team}
            {analytics}
          </div>
        </>
      )
  }
  ```

  - [folderName] 动态路由, 动态路由参数可以在 `layout.tsx`, `page.tsx`, `route.ts` 和 `generateMetadata` 中获取
    - [...folderName] 截获所有动态路由参数
    - [[...folerName]] 可选的截获所有动态路由参数, 同时会截获不带任何动态参数的路由

    ```tsx
    /*
      app
        photo
          [id]
            page.tsx
        doc
          [[...slug]]
            page.tsx
        page.tsx
    */
    ```

  - (..)folderName 拦截路由, 在另一个页面中使用布局渲染拦截当前路由
    - (.)folderName 匹配同一级的路由
    - (..)folderName 匹配上一级的路由
    - (..)(..)folderName 匹配上上一级的路由
    - (...)folderName 匹配根路由

  ```tsx
  // 在 app/page.tsx 软导航 /photo/110 将渲染拦截路由 @modal/(.)photo/[id]/page.tsx 下的内容
  // 硬导航 /photo/110 时渲染 app/photo/[id]/page.tsx
  /* 
    app
      @modal
        default.tsx // 返回 null 在未匹配到 插槽 时不渲染内容
        (.)photo
          [id]
            page.tsx
      photo
        [id]
          page.tsx
      layout.tsx  
      page.tsx
  */
  ```

### 路由段

直接在 layout, page, route handlers 中导出以下配置修改行为

```tsx
// 阻止页面预渲染, 如果使用 cookie, headers, searchParams 等函数页面自动被视为动态渲染
export const dynamic: string = 'force-dynamic';  // auto | force-dynamic | error | force-static

// layout 和 page 启用部分渲染
export const experimental_ppr: boolean = true;

// 控制访问非 generateStaticParams 生成的动态段时会发生什么
export const dyanmicParams: boolean = true;

// 设置 layout 和 page 的验证时间间隔(秒)
export const revalidate: boolean | number = false; // false | 0 | number

// 高级设置, 如果需要重置默认行为时使用
// export const fetchCache: string = 'auto'; // auto | default-cache | only-cache | force-cache | force-no-store | default-no-store | only-no-store

// 设置运行时
export const runtime: string = 'nodejs';  //nodejs | edge

// 设置首选区域
export const preferredRegion: string = 'auto'; // auto | global | home | string | string[]

// 限制服务器端逻辑的执行时长, next.js 默认不限制
// export const maxDuration: number = 0;
```

### 缓存

#### Request Memoization

是 React 的一个特性, next.js 扩展了 fetch API, 自动缓存相同的请求, 在 react 组件树中为相同的数据多次调用 fetch 函数将只执行一次

渲染路由时, 第一次调用特定请求时结果不在内存中而是缓存 `MISS`, 函数将被执行获取外部数据后存储到内存中,
在同一渲染过程中, 请求的后续函数调用将是缓存 `HIT`, 数据在不执行函数的情况下从内存中返回,
一旦路由被渲染并且渲染过程完成时, 内存将会被 重置, 所有请求记忆都会被清除

- 仅适用于 fetch 请求中的 GET 方法, 其他请求方法不会被记忆
- 仅适用于 React 组件树, 例如 `generateMetadata`、`generateStaticParams`、Layout、Page 和其他服务器组件中, route handlers 不适用因为不属于 React 组件树

不推荐退出请求记忆

#### Data Cache 数据缓存

next.js 有一个内置的数据缓存, 可以在传入的服务器请求和部署中持久保持数据获取的结果

- 使用 fetch('', {cache: 'force-cache'}) 强制使用缓存
- 使用 fetch('', {cache: 'force-cache', next: { revalidate: 3600 }}) 设置 next.js 验证数据的时间间隔(秒)
- cache mode, default | no-store | reload | no-cache | force-cache
  - force-cache, 自己先在缓存中查找资源, 如果有不管是否过期直接返回
  - default, 自己先在缓存中查找资源, 然后验证资源是否过期, 如果过期再询问服务器资源是否过期
  - no-cache, 自己先在缓存中查找资源, 然后再询问服务器资源是否过期
  - reload, 不查看缓存, 直接从服务器获取资源, 然后使用下载的资源更新缓存
  - no-store, 不查看缓存, 直接从服务器获取资源, 并且不会更新缓存资源

退出数据缓存

- 使用 fetch 不指定 cache 参数或者指定 {cache: 'no-store'}

#### Full Route Cache 完整路由缓存

next.js 在构建时自动渲染和缓存路由, 而不是在服务器上为每个请求渲染从而加快页面加载速度

- 使用流式 `服务器端组件载荷`(特殊的数据格式, 使用 React Server Component Payload 和 Client Component 渲染 HTML) 返回响应而无需等待所有渲染完成
- 默认缓存路由的渲染结果

退出完整路由缓存

- 使用 dynamic APIs
- 导出 dynamic = 'force-dynamic' 或者 revalidate = 0
- 退出 Data Cache, 如果路由有一个未缓存的获取请求, 这将该路由退出完整路由缓存为每个请求获取特定数据, 其他未退出数据缓存的获取请求仍将缓存在数据缓存中
  这允许缓存和未缓存数据的混合.
  
#### Router Cache 路由缓存

next.js 有一个客户端的路由缓存, 用于存储路由段的 RSC(React Server Component) 载荷, 按 layout、加载状态和 page 划分

当用户在路由之间导航时, next.js 会缓存访问过的路由段, 并预取用户可能导航到的路由, 导航之间不会重新加载整个页面, 并保留 React 状态和浏览器状态

- 布局被缓存并在导航时重用(部分渲染)
- 加载状态被缓存并在导航中重用, 以实现即时导航
- 默认页面不会被缓存, 但在浏览器向前和向后导航期间会被重用

### 元数据

元数据支持两种方式配置, `配置文件` 和 `动态生成`

favicon

- 使用图片文件放在 app 目录下
- 使用代码生成图标文件

```tsx
// app/icon.tsx
import {ImageResponse} from 'next/og';

export const size = {width: 32, height: 32};
export const contentType = 'image/png';
export default function Icon(){
  return new ImageResponse(
    (
      // ImageResponse JSX element
      <div
        style={{
          fontSize: 24,
          background: 'black',
          width: '100%',
          height: '100%',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          color: 'white',
        }}
      >
        A
      </div>
    ),
    // ImageResponse options
    {
      // For convenience, we can re-use the exported icons size metadata
      // config to also set the ImageResponse's width and height.
      ...size,
    }
  )
}
```

metadata <em id ="metadata"></em> <!--markdownlint-disable-line--> 

不能在相同的路由段中同时使用 静态配置 和 动态生成 两种方式, 从 layout.tsx 或 page.tsx 中导出

导出 metadata 时不能使用 `'use client'` 指令标识组件为客户端组件

- 静态配置
- 使用 `generateMetadata` 函数动态生成

```tsx
// layout.tsx, page.tsx
import type {Metadata, ResolvingMetadata} from "next";
import Counter from './counter'; // import client component with 'use client' directive.

// static metatdata
export const metadata: Metadata = {
  title: "",
  description: "",
  generator: 'Next.js',
  applicationName: 'Next.js',
  referrer: 'origin-when-cross-origin',
  keywords: ['Next.js', 'React', 'JavaScript'],
  authors: [{ name: 'Seb' }, { name: 'Josh', url: 'https://nextjs.org' }],
  creator: 'Jiachi Liu',
  publisher: 'Sebastian Markbåge',
  formatDetection: {
    email: false,
    address: false,
    telephone: false,
  },
  openGraph: {},
  robots: {}
  icons: {},
  colorScheme: '',
  manifest: '',
  twitter: {},
  viewport: {},
  alternates: {},
  assets: [],
  category: '',
  bookmarks: [],
}
// dynamic metadata
type Props = {
  params: Promise<{id: string}>,
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>
}
export async function generateMetadata({params, searchParams}: Props, parent: ResolvingMetadata): Promise<Metadata> {
  const {id} = params;
  const res = await fetch();
  return {title: "", description: ""};
}
export default function Page({params, searchParams}: Props){
  return <Counter />
}

// couter.tsx 使用 client component hooks.
'use client';
import React, { useState } from 'react';

export default function Counter() {
  const [count, setCount] = useState(0);
  return (
    <>
      <p>count is {count}</p>
      <button onClick={() => setCount(count + 1)}>Click</button>
    </>
  )
}
```

viewport 页面视窗 <em id="viewport"></em> <!--markdownlint-disable-line-->

```tsx
// layout.tsx, page.tsx
// static viewport
import type {Viewport} from 'next';
export const viewport: Viewport = {
  themeColor: 'black'
}

// dynamic viewport
export function generateViewport(){
  return {}
}
export default function Page(){
  // ...
}
```

manifest

```tsx
// app/manifest.json
// static manifest
{
  "name": "My Next.js Application",
  "short_name": "Next.js App",
  "description": "An application built with Next.js",
  "start_url": "/"
  // ...
}

// app/manifest.ts
// dynamic manifest
import type { MetadataRoute } from 'next'
export default function manifest(): MetadataRoute.Manifest {
  return {
    name: 'Next.js App',
    short_name: 'Next.js App',
    description: 'Next.js App',
    start_url: '/',
    display: 'standalone',
    background_color: '#fff',
    theme_color: '#fff',
    icons: [
      {
        src: '/favicon.ico',
        sizes: 'any',
        type: 'image/x-icon',
      },
    ],
  }
}
```

robots

```tsx
// app/robots.txt
// static robots
/* */

// app/robots.ts
// dynamic robots
import type { MetadataRoute } from 'next'
 
export default function robots(): MetadataRoute.Robots {
  return {
    rules: {
      userAgent: '*',
      allow: '/',
      disallow: '/private/',
    },
    sitemap: 'https://acme.com/sitemap.xml',
  }
}
```

sitemap <em id="sitemap"></em> <!--markdownlint-disable-line-->

```tsx
// app/sitemap.xml
// static sitemap
/* 
  <xml>
    <property></property>
  </xml>
*/

// app/sitemap.ts
// dynamic sitemap
import type { MetadataRoute } from 'next'
export function generateSitemaps(){
  return {}
}
export default function sitemap(): MetadataRoute.Sitemap {
 // ...
}
```

### 函数

- headers 一个 async 函数, 在服务器组件内读取请求头信息
- cookies 一个 async 函数, 在服务器组件内读取请求中的 cookies

```tsx
// page.tsx
import {cookies, headers} from 'next/headers';
export default async function Page(){
  const headersList = await headers();
  const ua = headersList.get('user-agent');

  const cookieStore = await cookies();
  const theme = cookieStore.get('theme');

  return '...';
}
```

- NextRequest 扩展了 Web Request API
- NextResponse 扩展了 Web Response API
- notFound 调用方法将抛出 `NEXT_NOT_FOUND` 错误, 渲染 not-found.tsx 内容
- permanentRedirect 永久重定向, 返回 308(HTTP), 如果资源不存在可以使用 notFound 函数代替
- redirect 重定向
  - path
  - type, replace(default) | push
- revalidatePath 按需清理特定路径的缓存数据
  - path
  - type, page | layout
- revalidateTag 按需清理特定缓存标记的缓存数据
  - tag

- after 注册在响应结束之后执行的任务, 通常记录日志和数据分析

```tsx
// layout.tsx
import {after} from 'next/server';
export default function Layout({children}){
  after(() => {
    // layout 渲染完成发送给请求后执行
    log();
  })
  return (
    <div>
      Hello World
      {children}
    </div>
  )
}
```

- connection 标记渲染内容等待用户的请求传入

当不使用 dynamic APIs 时希望在运行时动态渲染而不是在构建时静态渲染, 通常用在访问有意更改渲染结果的外部信息时

```tsx
// page.tsx
import {connection} from 'next/server';
export default async function Page(){
  await connection(); // 等待请求传入
  // Everything below will be excluded from prerendering
  const rand = Math.rand();
  return <span>{rand}</span>
}
```

- drafMode 启用或禁用草稿模式(draftMode), async 函数
  - isEnabled, 标识 draftMode 是否启用
  - enable(), 启用 draftMode
  - disable(), 禁用 draftMode

草稿模式允许在 next.js 应用程序中预览无头 CMS 中的草稿内容而无需重建整个网站, 对于在构建时静态渲染的内容允许切换到动态渲染并查看更改非常有用

```tsx
// page.tsx
import {draftMode} from 'next/server';
async function getData(){
  const {isEnabled} = await draftMode();
  const url = isEnabled ? 'https://draft.example.com' : 'https://product.example.com';

  const res = await fetch(url);
  return res.json();
}
export default async function Page(){
  const {title, desc} = await getData();
  return (
    <main>
      <h1>{title}</h1>
      <p>{desc}</p>
    </main>
  )
}

// app/api/draft/route.ts
import {draftMode, NextRequest} from 'next/server';
import {redirect} from 'next/navigation';
export async function GET(request: NextRequest) {
  // Parse query string parameters
  const { searchParams } = new URL(request.nextUrl)
  const secret = searchParams.get('secret')
  const slug = searchParams.get('slug')
 
  // Check the secret and next parameters
  // This secret should only be known to this Route Handler and the CMS
  if (secret !== 'MY_SECRET_TOKEN' || !slug) {
    return new Response('Invalid token', { status: 401 })
  }
 
  // Fetch the headless CMS to check if the provided `slug` exists
  // getPostBySlug would implement the required fetching logic to the headless CMS
  const post = await getPostBySlug(slug)
 
  // If the slug doesn't exist prevent draft mode from being enabled
  if (!post) {
    return new Response('Invalid slug', { status: 401 })
  }
 
  // Enable Draft Mode by setting the cookie
  const draft = await draftMode()
  draft.enable()
 
  // Redirect to the path from the fetched post
  // We don't redirect to searchParams.slug as that might lead to open redirect vulnerabilities
  redirect(post.slug)
}
```

- fetch 扩展了 Web fetch API
- generateImageMetadata 生成一个或多个不同版本的图片元数据, 希望避免硬编码元数据时例如 Icon
  - params, 一个 Promise, Next.js 14 之前是同步的
  - 返回值
    - id, string,required
    - alt, string
    - size, {width: number, height: number}
    - contentType, string

```tsx
import {ImageResponse} from 'next/og';
export function generateImageMetadata(){
  return [
    {id: 'small', contentType: 'image/png', size: {width: 40, height: 40}},
    {id: 'medium', contentType: 'image/png', size: {width: 72, height: 72}}
  ]
}
export default function Icon({id}: {id: string}){
  return new ImageResponse((
    <div
      style={{
        width: '100%',
        height: '100%',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        fontSize: 88,
        background: '#000',
        color: '#fafafa',
      }}>
      Icon {id}
    </div>
  ))
}
```

- [generateMetadata](#metadata) 生成页面元数据
- [generateSitemaps](#sitemap) 生成应用站点地图
- generateStaticParams 合并动态路由段和静态路由, 在构建时生成路由而不是在请求时按需生成
  - Props
    - params 动态路由参数, 一个 Promise, Next.js 14 之前是同步的

```tsx
export async function generateStaticParams() {
  return [
    { category: 'a', product: '1' },
    { category: 'b', product: '2' },
    { category: 'c', product: '3' },
  ];
}
// Three versions of this page will be statically generated
// using the `params` returned by `generateStaticParams`
// - /products/a/1
// - /products/b/2
// - /products/c/3
export default async function Page({params}: {
  params: Promise<{category: string, product: string}>
}){
  const {categor, product} = await params;
  // ...
}
```

- [generateViewport](#viewport) 生成页面的视窗配置
- ImageResponse 图片构造函数, 生成动态图片 `import { ImageResponse } from 'next/og'`;

#### hook

Client Component Hook

- useParams 获取动态路由参数
- usePathname 获取当前 url 的路径
- useReportWebVitals 获取网站性能指标
- useRouter 编程式改变路由

```tsx
'use client';
export default function Page(){
  const router = useRouter();
  
  return (<div>
    Hello World!
    <button onClick={() => router.push('/login')}>login</button>
  </div>);
}
```

- useSearchParams 获取当前 url 查询参数

- useSelectedLayoutSegment 获取当前 layout 下面一层的活动路由段, 通常用于在父布局中改变子段的状态
- useSelectedLayoutSegments 获取当前 layout 下的活动路由段, 通常用于在父布局中改变子段的状态

```tsx
// app/blog/blog-nav-link.tsx
'use client'
 
import Link from 'next/link'
import { useSelectedLayoutSegment } from 'next/navigation'
 
// This *client* component will be imported into a blog layout
export default function BlogNavLink({
  slug,
  children,
}: {
  slug: string
  children: React.ReactNode
}) {
  // Navigating to `/blog/hello-world` will return 'hello-world'
  // for the selected layout segment
  const segment = useSelectedLayoutSegment()
  const isActive = slug === segment
 
  return (
    <Link
      href={`/blog/${slug}`}
      // Change style depending on whether the link is active
      style={{ fontWeight: isActive ? 'bold' : 'normal' }}
    >
      {children}
    </Link>
  )
}

// app/blog/layout.tsx
// Import the Client Component into a parent Layout (Server Component)
import { BlogNavLink } from './blog-nav-link'
import getFeaturedPosts from './get-featured-posts'
 
export default async function Layout({
  children,
}: {
  children: React.ReactNode
}) {
  const featuredPosts = await getFeaturedPosts()
  return (
    <div>
      {featuredPosts.map((post) => (
        <div key={post.id}>
          <BlogNavLink slug={post.slug}>{post.title}</BlogNavLink>
        </div>
      ))}
      <div>{children}</div>
    </div>
  )
}
```

- userAgent 获取 request 请求中的 user-agent

```tsx
import { NextRequest, NextResponse, userAgent } from "next/server";
export function middleware(request: NextRequest){
  const ua = userAgent(request);
  console.log('ua', ua)

  const requestHeaders = new Headers(request.headers);
  requestHeaders.set('x-hello-form-middleware', 'hello');

   return NextResponse.next({
    request: {
      headers: requestHeaders
    }
  })
}
```
