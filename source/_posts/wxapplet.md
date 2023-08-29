---
title: 微信小程序
date: 2021-12-24 11:54:05
categories:
  - wx
tags:
  - wx
---

### 敲黑板

#### [头像昵称 API 调整](https://developers.weixin.qq.com/community/develop/doc/00022c683e8a80b29bed2142b56c01)

- 20221026 起, 小程序 wx.getUserProfile 接口将被收回, wx.getUserInfo 接口获取用户头像将统一返回默认灰色头像，昵称将统一返回 “微信用户”
- 2.21.2 基础库开始, 如果业务需获取用户头像昵称，可以使用 [头像昵称填写能力](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/userProfile.html)

#### 地理位置

##### [地理位置 API 调整](https://developers.weixin.qq.com/community/develop/doc/000a02f2c5026891650e7f40351c01?blockType=1)

- 20220628 起, 使用地理位置 API 需要在 app.json 配置项 requirePrivateInfos 提前声明

```json
{ "requirePrivateInfos": ["chooseAddress", "getFuzzyLocation"] }
```

- 部分 API 需要在管理后台申请接口权限

![wx-1](/images/wx-1.png)

##### [接口权限申请](https://developers.weixin.qq.com/community/develop/doc/000e8ccb5ac498318cbd26c495bc01?blockType=1)

- getFuzzyLocation
- getLocation
- onLocationChange
- chooseAddress
- choosePoi
- chooseLocation

- 20220418 起, 如果使用以上接口, 需要在小程序管理后台申请接口权限
- 地理位置新增接口和相关流程调整

##### [choosePoi|chooseLocation](https://developers.weixin.qq.com/community/develop/doc/0006e45df2cac030e6edf367c56001?blockType=1)

- 20220613 起, 使用该接口不在需要用户授权 scope.userLocation
- wx.choosePoi 回调信息中不再返回真实的经纬度信息, 全部返回(0,0)

##### [getFuzzyLocation](https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.chooseLocation.html)

- 2.25.0 开始, 新增获取模糊地理位置接口, 接口规则同 [chooseLocation](https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.chooseLocation.html)

#### [getPhoneNumber](https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/phonenumber/phonenumber.getPhoneNumber.html)

- e.detail.encryptedData
- e.detail.iv

- 20230826 开始收费, 分为 手机号快速验证组件 和 手机号实时验证组件

```html
<!-- getRealtimePhoneNumber 回调函数参数不再包含 encryptedData 和 iv, 仅可通过返回的 code 换取手机号-->
<button
  open-type="getRealtimePhoneNumber"
  bind:getrealtimephonenumber="getrealtimephonenumber"
></button>
```

#### [wx.openSetting](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/setting/wx.openSetting.html)

- 2.3.0 开始, [打开小程序设置页](https://developers.weixin.qq.com/community/develop/doc/000cea2305cc5047af5733de751008)

改为由用户行为触发, 直接调用此 API 无效, 需要在页面中使用 button 组件

```html
<!-- 方法一 -->
<button open-type="openSetting" bind:opensetting="callback">打开设置页</button>
<!-- 方法二 -->
<button bind:tap="openSetting">打开设置页</button>
openSetting(){ wx.openSetting() }
```

#### [wx.getSetting](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/setting/wx.getSetting.html)

- 2.10.1 开始, withSubscriptions 参数控制是否同时获取用户订阅消息的订阅状态, 默认 false
- [SubscriptionsSetting](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/setting/SubscriptionsSetting.html) 订阅消息设置
  - mainSwitch 布尔值, 标识订阅消息的总开关
  - itemSettings 一个对象, 包含每一项订阅消息的订阅状态

```javascript
wx.getSetting({ withSubscriptions: true })
  .then((res) => {
    console.log(res.authSetting); // 用户授权结果
    console.log(ers.subscriptionsSetting); // 当参数 withSubscriptions 为 true 时返回此结果
    /*{
      mainSwitch: true, // 订阅消息总开关
      // itemSettings 只返回用户勾选过订阅面板中的 "总是保持以上选择, 不再询问" 的订阅消息
      itemSettings: {
        // 每一项开关
        // 每一项已授权键为消息模板的 id, 值为 accept, reject, ban, filter 其中一种
      } 
    }*/
    console.log(res.miniprogramAuthSetting); // 在插件种调用时, 返回宿主小程序的用户授权结果
  })
  .catch((err) => {
    console.log(err);
  });
```

#### [wx.requestSubscribeMessage](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeMessage.html) 订阅消息

调起客户端小程序订阅消息界面, 返回用户订阅消息的操作结果, 如果用户勾选了订阅消息界面的 **总是保持以上选择，不再询问** 选项时, 消息模板会被记录在用户的小程序设置页, 并且在每次调用此 API 时不再弹出订阅消息界面(只返回订阅消息结果), 通过 `wx.getSetting` API 可以获取用户订阅消息的订阅状态

- 一次性模板 id 和永久模板 id 不能同时混用
- IOS 7.0.6/Android 7.0.7 之后支持多个同类型消息模板
- 2.8.2 开始, 用户发生点击行为或者发起支付回调后, 才可以调起订阅消息界面
  - 可以在事件处理函数中使用同步调用方式
  - 不能在异步回调中调用, 否则报错(requestSubscribeMessage:fail can only be invoked by user TAP gesture.)
- 2.10.0 开始, 支持订阅语音消息提醒

```javascript
wx.requestSubscribeMessage({ tmplIds: [] /*消息模板 id*/ })
  .then((res) => {
    // API 调用成功之后返回的参数
    {
      [TEMPLATE_ID]: 'accept', // reject, ban, filter
      errMsg: 'requestSubscribeMessage:ok'
    }
  })
  .catch((err) => {
    // API 调用失败返回的参数(部分状态码)
    {
      errCode: 20004,
      errMsg: "requestSubscribeMessage:fail:The main switch is switched off"
    }
  });
```

##### **总是保持以上选择，不再询问**

- 已勾选(默认), 每次调用此 API 不再弹出订阅消息界面, 只返回记录在用户的小程序设置页中的订阅消息结果
- 未勾选, 每次调用此 API 都会弹出订阅消息界面

##### 订阅消息的总开关

订阅消息的总开关 **关闭** 时, 调用此 API 会报一个小程序错误, 可以使用 App.onError 或者 wx.onError 捕获错误

![wx-2](/images/wx-2.jpg)

#### 组件[behaviors](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/behaviors.html)

- 自定义组件混入 behaviors, 在 attached 钩子函数中调用混入的方法获取不到, 需要在 ready 中调用

#### ios 微信小程序 input 输入框使用 bind:focus 时, 在 input 获取焦点后会出现输入法键盘,然后点击其他下拉框或自定义组件会出现 input 仍然可以编辑状态

- 修改 input 的 focus 属性

#### 小程序多 appId 平台提审和接口域名切换

使用以下 API 获取小程序 appId, 根据 appid 推送不同平台和切换不同接口域名

- [wx.getLaunchOptionsSync](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/life-cycle/wx.getLaunchOptionsSync.html) 获取小程序启动时的参数, 与 App.onLaunch 的回调参数一致
  - 2.1.2 支持
- [wx.getAccountInfoSync](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/account-info/wx.getAccountInfoSync.html) 获取当前帐号信息, 线上小程序版本号仅支持在正式版小程序中获取, 开发版和体验版中无法获取
  - 2.2.2 支持

#### 单页应用添加微信分享打包后分享时好好坏

- [企微](https://developer.work.weixin.qq.com/document/path/90542) 微信 JSSDK 签名的 url 和当前显示的 url 不一致导致自定义分享信息失败,在路由钩子函数或者监听路由变化时重新请求接口签名
- [微信](https://developers.weixin.qq.com/community/develop/doc/000088945c4320dae71d677d15b400?jumpto=reply&parent_commentid=000062cf8fcaf80be71d58bd15b4&commentid=000a68a75ac9188cf81d2be8353c) 微信中打开普通 URL 链接分享仍为普通 URL 链接, 卡片形式的链接不受此影响

#### 小程序内嵌 H5 页区分运行环境

小程序内嵌 H5 页区分运行环境是小程序还是 H5 执行不同逻辑调用 wx.miniProgram.getEnv 依赖微信 JSSDK 1.4 以上版本, 微信 JSSDK 1.0 报错

```javascript
var ua = navigator.userAgent.toLowerCase();
var result = ua.match(/MicroMessenger/i);
if (result && result[0] == `micromessenger`) {
  //ios的ua中无miniProgram，但都有MicroMessenger（表示是微信浏览器）
  wx.miniProgram.getEnv(function (res) {
    if (res.miniprogram) {
      // 小程序内
    } else {
      // 非小程序内
    }
  });
} else {
  // 非小程序内
}
```

<!-- more -->

### 逻辑层

#### [注册页面](https://developers.weixin.qq.com/miniprogram/dev/framework/app-service/page.html)

- options 2.10.1, 页面的组件选项, 同 [Component 构造器](#zidingyizujian)中的 options
- behaviors 2.9.2, 使多个页面可以共享相同的数据和方法

```javascript
// my-behavior.js
module.exports = Behavior({
  data: {},
  methods: {
    sharedMethod: function () {
      this.data.sharedText === 'This is a piece of data shared between pages.';
    },
  },
});
// page-a.js
const myBehaviors = require('./my-behavior.js');
Page({
  behaviors: [myBehaviors],
  onLoad() {
    this.data.sharedText === 'This is a piece of data shared between pages.';
  },
});
```

#### [API](https://developers.weixin.qq.com/miniprogram/dev/framework/app-service/api.html#API)

- 2.10.2 支持 callback 和 promise 两种调用方式. 接口参数 Object 对象不包含 success/fail/complete 时默认返回 promise, 否则按回调方式执行,无返回值

### 视图层

#### [WXML](https://developers.weixin.qq.com/miniprogram/dev/framework/view/wxml/)

- dataset 自定义数据以 `data-` 开头, 多个单词使用连字符 `-` 连接, 此写法会被自动转换成驼峰写法, 大写字符会自动转换成小写字符

```javascript
data-element-type /* 最终会呈现为 */ event.currentTarget.dataset.elementType
data-elementType /* 最终会呈现为 */ event.currentTarget.dataset.elementtype

// vue自定义事件绑定使用 kebab-case 写法
v-bind:my-event
```

- 2.9.3 支持简易双向绑定 `model:property="{{myProperty}}"`

#### [事件](https://developers.weixin.qq.com/miniprogram/dev/framework/view/wxml/event.html)

- 1.5.0 支持事件捕获阶段, 可使用 `capture-bind`, `capture-catch` 关键字

```html
<!-- 执行顺序: handleTap2、handleTap4、handleTap3、handleTap1 -->
<!-- 如果将第一个 capture-bind 改为 capture-catch，将只触发 handleTap2  -->
<view
  id="outer"
  bind:touchstart="handleTap1"
  capture-bind:touchstart="handleTap2"
>
  outer view
  <view
    id="inner"
    bind:touchstart="handleTap3"
    capture-bind:touchstart="handleTap4"
  >
    inner view
  </view>
</view>
```

- 2.4.4 支持 [WXS 函数响应事件](https://developers.weixin.qq.com/miniprogram/dev/framework/view/interactive-animation.html)
- 2.8.1 支持所有组件 `bind:event_name` 事件绑定方式
- 2.8.2 支持 `mut-bind` 事件绑定, 所有 `mut-bind` 之间是"互斥"的, 只会有其中一个绑定函数被触发, 同时,不会影响到 `bind` 和 `catch` 的绑定效果

```html
<!-- 点击 inter view 会依次调用 handleTap3、handleTap2 -->
<!-- 点击 middle view 会依次调用 handleTap2、handleTap1 -->
<view id="outer" mut-bind:tap="handleTap1">
  outer view
  <view id="middle" bind:tap="handleTap2">
    middle view
    <view id="inner" mut-bind:tap="handleTap3"> inner view </view>
  </view>
</view>
```

#### [响应显示区域](https://developers.weixin.qq.com/miniprogram/dev/framework/view/resizable.html)

- 2.4.0 支持屏幕旋转, 可配置全局或者页面的屏幕旋转

```json
{ "pageOrientation": "auto" }
```

- 2.5.0 支持 `pageOrientation` 设置为 `landscape` 表示为固定横屏显示
- 2.3.0 支持 iPad 小程序屏幕旋转, `仅支持`在 app.json 中配置

```json
{ "resizable": true }
```

- 2.4.0 支持页面 `resize` 生命周期监听页面显示区域变化
- windows wx 3.3 支持小程序`分栏模式`, 在 app.json 中配置

```json
{ "resizable": true, "frameset": true }
```

#### [动画](https://developers.weixin.qq.com/miniprogram/dev/framework/view/animation.html)

- 2.9.0 支持[关键帧动画](https://developers.weixin.qq.com/miniprogram/dev/framework/view/animation.html)代替旧的 `wx.createAnimation`

```javascript
// selector String, required, 选择器
// keyframes Array, required, 关键帧信息
// duration Number, required, 动画持续时长(单位毫秒)
// callback function, not required, 动画完成后的回调
this.animate(selector, keyframes, duration, callback);
// 清楚动画执行完毕后的属性
this.clearAnimation(selector, options, callback);
```

#### [初始渲染缓存](https://developers.weixin.qq.com/miniprogram/dev/framework/view/initial-rendering-cache.html)

- 2.11.1 支持

  - 支持的内置组件
    - view
    - text
    - button
    - image
    - scroll-view
    - rich-text
  - 启用初始化渲染缓存指定页面,在 `页面.json` 配置文件中配置, 不包含任何 `setData` 的结果

  ```json
  { "initialRenderingCache": "static" }
  ```

  - 缓存所有页面, 在 `app.json` 中配置

  ```json
  {
    "window": {
      "initialRenderingCache": "static"
    }
  }
  ```

- 初始渲染缓存页面中添加动态内容, `页面.json` 配置文件

  - 此方法调用不能早于 `Page` 的 `onReady` 或者 `Component` 的 `ready` 生命周期, 否则影响性能
  - 参数传入 `null` 可禁用初始渲染缓存

```json
{ "initialRenderingCache": "dynamic" }
```

```javascript
<view wx:if='{{loading}}'>{{ loadingHint }}</view>;
// 页面中调用此方法启用
Page({
  data: {
    loading: true,
  },
  onReady: function () {
    // 这一部分数据将被应用于界面上，相当于在初始 data 基础上额外进行一次 setData
    this.setInitialRenderingCache({ loadingHint: '正在加载' });
  },
});
```

### 运行

#### [运行机制](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/operating-mechanism.html)

- 前台进入后台 `5秒` 后进入挂起状态, 小程序代码停止运行
- 挂起状态维持 `30分钟` 后会被销毁, 如果小程序占用系统资源过高,可能会被系统销毁或者微信客户端主动回收

#### 重新启动策略

- 2.8.0 支持如果冷启动时不带`path`参数, 默认进入小程序的首页, 在 `页面.json` 或者 `app.json` 修改小程序冷启动时的默认行为

  - 如果小程序退出时间过久 `1天`, 下次冷启动时不遵循此规则

```json
/* homePageAndLatestPage */
//如果从这个页面退出小程序，下次冷启动后立刻加载这个页面，页面的参数保持不变（不可用于 tab 页）
{ "restartStrategy": "homePage" }
```

#### 退出状态

小程序可能被销毁之前, 页面回调函数 `onSaveExitState` 会被调用, 如果需要保留页面页面中的状态, 可以在这个回调函数中保存一些数据, 下次启动时可以通过 `exitState` 获得已保存得数据, 如果小程序退出时间过久`1天`会丢弃保存的数据

- 2.7.4 支持

```json
{ "restartStrategy": "homePageAndLatestPage" }
```

```javascript
Page({
  onLoad: function () {
    // 尝试获得上一次退出前 onSaveExitState 保存的数据
    var prevExitState = this.exitState;
    if (prevExitState !== undefined) {
      // 如果是根据 restartStrategy 配置进行的冷启动，就可以获取到
      prevExitState.myDataField === 'myData';
    }
  },
  onSaveExitState: function () {
    var exitState = { myDataField: 'myData' }; // 需要保存的数据
    return {
      data: exitState,
      // 超时时刻, 默认:当前时间 + 1天
      expireTimeStamp: Date.now() + 24 * 60 * 60 * 1000,
    };
  },
});
```

### 自定义组件 <em id="zidingyizujian"></em> <!-- markdownlint-disable-line -->

- 1.6.3 支持

#### [构造器](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/component.html)

- data
- methods
- externalClasses 组件接收的外部样式类, 一个字符串数组
- properties
  - 配置项中使用驼峰写法(propertyName), 在 xml 文件中使用连字符写法(property-name)
- behaviors
  - 自定义组件引入的 behaviors 定义的方法在 attached 钩子函数中获取不到，需要在 ready 中调用
- [relations](#relations) 组件关系定义
- [observers](#observers) 数据监听器
- options
  - multipleSlots 是否在组件定义时的选项中启用多 slot 支持
  - styleIsolation 组件样式隔离选项, 取值 'isolated', 'apply-shared', 'shared'
    - 'isolated' 启用样式隔离, 默认值
    - 'apply-shared' 页面 wxss 样式将影响到自定义组件, 但反之不会影响
    - 'shared' 页面样式和自定义组件样式相互影响, 插件中不可用, 使用 Component 构造器创建页面时, 默认此项
    - 'page-isolated' 此页面禁用 app.wxss, 同时, 页面的 wxss 不会影响到其它自定义组件
    - 'page-apply-shared' 此页面禁用 app.wxss, 同时, 页面 wxss 不会影响到其他自定义组件, 但设为 shared 的自定义组件会影响到页面
    - 'page-shared' 此页面禁用 app.wxss, 同时, 页面 wxss 样式会影响到其他设为 apply-shared 或 shared 的自定义组件, 也会受到设为 shared 的自定义组件的影响
  - addGlobalClass: true // 等价于设置 `styleIsolation: apply-shared`
  - [pureDataPattern](#pureDataPattern) 指定符合规则的字段为纯数据字段, 仅作为当前组件内部使用不参与页面渲染, 2.10.1 开始可以在 json 文件中配置
  - virtualHost: true 虚拟化组件节点
- 生命周期

  - lifetimes 2.2.3 支持, 优先级最高会覆盖外部定义的生命周期钩子
    - created 在组件实例刚刚被创建时执行，注意此时不能调用 setData
    - attached 在组件实例进入页面节点树时执行, 大部分初始化工作可以在此方法内执行
    - ready 在组件布局完成后执行
    - moved 在组件实例被移动到节点树另一个位置时执行
    - detached 在组件实例被从页面节点树移除时执行
    - error 当组件方法抛出错误时执行, 基础库 2.4.1 支持
  - pageLifetimes 组件所在页面的生命周期
    - show 组件所在的页面被展示时执行
    - hide 组件所在的页面被隐藏时执行
    - resize 组件所在的页面尺寸变化时执行
    - routeDone 组件所在页面路由动画完成时执行, 基础库 2.31.2 支持

- [抽象节点](#generic-node)

- this.selectComponent(selector) 父组件中获取子组件实例对象
- this.triggerEvent(evtName, evtDetail, evtOptions) 组件触发事件

#### [引用](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/)

- 在 `页面.json` 文件中配置

```json
{
  "usingComponents": {
    "component-tag-name": "path/to/the/custom/component"
  }
}
```

- 页面

```html
<view>
  <!-- 以下是对一个自定义组件的引用 -->
  <component-tag-name inner-text="Some text"></component-tag-name>
</view>
```

#### [模板和样式](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/wxml-wxss.html)

- slot

  - 开启多个 slot

  ```javascript
  Component({
    options: {
      multipleSlots: true, // 在组件定义时的选项中启用多slot支持
    },
    properties: {},
    methods: {},
  });
  ```

  ```html
  <!-- 组件模板 -->
  <view class="wrapper">
    <slot name="before"></slot>
    <view>这里是组件的内部节点</view>
    <slot name="after"></slot>
  </view>
  <!-- 引用组件的页面模板 -->
  <view>
    <component-tag-name>
      <view slot="before">这里是插入到组件slot-before中的内容</view>
      <view slot="after">这里是插入到组件slot-after中的内容</view>
    </component-tag-name>
  </view>
  ```

- 组件样式隔离, 避免组件内部使用的`标签下选择器样式`污染页面

  - 2.6.5 支持 `js` 中配置

  ```javascript
  Component({
    options: { styleIsolation: 'isolated' },
  });
  ```

  - 2.10.1 支持在 `页面.json` 或者 `自定义组件.json` 中配置

  ```json
  { "styleIsolation": "isolated" }
  ```

- 外部样式应用

  - 1.9.90 支持
  - 2.7.1 支持多个 `class`

  ```javascript
  /* 组件 custom-component.js */
  Component({
    externalClasses: ['my-class'],
  });
  ```

  ```html
  <!-- 组件 custom-component.wxml -->
  <custom-component class="my-class"
    >这段文本的颜色由组件外的 class 决定</custom-component
  >
  ```

#### [behaviors](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/behaviors.html)

- 内置 behaviors

  - wx://form-field
  - wx://form-field-group 2.10.2 支持
  - wx://form-field-button 2.10.3 支持
  - wx://component-export 2.2.3 支持

- 同名字段覆盖和组合规则
  - 同名属性和方法
    - 组件本身的属性和方法覆盖 behavior 中的
    - behavior 中定义靠后的覆盖靠前的
    - 嵌套引用 behavior 的情况, 父 behavior 覆盖子 behavior 中
  - 同名数据字段
    - 对象类型进行对象合并
    - 其余情况进行数据覆盖: 组件 > 父 behavior > 子 behavior 、 靠后的 behavior > 靠前的 behavior
  - 生命周期钩子不会覆盖,逐个调用
    - behavior 优先于组件执行
    - 子 behavior 优先于 父 behavior 执行
    - 靠前的 behavior 优先于 靠后的 behavior 执行
    - 多次引用 behavior, 生命周期只会被执行一次
- 2.15.0 [behaviors 中声明的生命周期钩子会被 Page 和 Component 构造器中声明的同名钩子覆盖执行, 和上面文档中描述的不一样](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/behaviors.html)

#### [组件之间关系](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/relations.html) <em id="relations"></em> <!-- markdownlint-disable-line -->

- type 目标组件的相对关系
  - parent
  - child
  - ancestor
  - descendant
- linked 关系生命周期函数, 当关系被建立在页面节点树中时触发, 触发时机在组件 attached 之后
- linkChanged 关系生命周期函数, 当关系在页面节点树中发生改变时触发, 触发时机在组件 moved 之后
- unlinked 关系生命周期函数, 当关系脱离页面节点树时触发, 触发时机在组件 detached 之后
- target 关联的目标节点所应具有的 behavior，所有拥有这一 behavior 的组件节点都会被关联

```html
<custom-ul>
  <custom-li> item 1 </custom-li>
  <custom-li> item 2 </custom-li>
</custom-ul>
```

```js
// path/to/custom-ul.js
Component({
  relations: {
    './custom-li': {
      type: 'child', // 关联的目标节点应为子节点
      linked: function (target) {
        // 每次有custom-li被插入时执行，target是该节点实例对象，触发在该节点attached生命周期之后
      },
      linkChanged: function (target) {
        // 每次有custom-li被移动后执行，target是该节点实例对象，触发在该节点moved生命周期之后
      },
      unlinked: function (target) {
        // 每次有custom-li被移除时执行，target是该节点实例对象，触发在该节点detached生命周期之后
      },
    },
  },
  methods: {
    _getAllLi: function () {
      // 使用getRelationNodes可以获得nodes数组，包含所有已关联的custom-li，且是有序的
      var nodes = this.getRelationNodes('path/to/custom-li');
    },
  },
  ready: function () {
    this._getAllLi();
  },
});
// path/to/custom-li.js
Component({
  relations: {
    './custom-ul': {
      type: 'parent', // 关联的目标节点应为父节点
      linked: function (target) {
        // 每次被插入到custom-ul时执行，target是custom-ul节点实例对象，触发在attached生命周期之后
      },
      linkChanged: function (target) {
        // 每次被移动后执行，target是custom-ul节点实例对象，触发在moved生命周期之后
      },
      unlinked: function (target) {
        // 每次被移除时执行，target是custom-ul节点实例对象，触发在detached生命周期之后
      },
    },
  },
});
```

#### [数据监听器](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/observer.html) <em id="observers"></em> <!-- markdownlint-disable-line-->

用于监听和响应任何属性和数据字段的变化,作用类似于计算属性

- 2.6.1 支持

```javascript
// this.data.sum 永远是 this.data.numberA 与 this.data.numberB 的和
Component({
  attached: function () {
    this.setData({ numberA: 1, numberB: 2 });
  },
  observers: {
    'numberA, numberB': function (numberA, numberB) {
      // 在 numberA 或者 numberB 被设置时，执行这个函数
      this.setData({ sum: numberA + numberB });
    },
  },
});
```

#### [纯数据字段](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/pure-data.html) <em id="pureDataPattern"></em> <!-- markdownlint-disable-line -->

- 纯数据字段是一些不用于界面渲染的 data 字段(包括 setData 设置的字段), 既不会展示在界面上,也不会传递给其他组件，可以用于提升页面更新性能

- 属性中的纯数据字段的属性不会触发 observer, 需要使用 observers([数据监听器](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/observer.html)) 监听
- 2.8.2 支持
- 2.10.1 支持在 json 配置文件中配置 pureDataPattern 项

```javascript
Component({
  options: {
    pureDataPattern: /^_/, // 指定所有 _ 开头的数据字段为纯数据字段
  },
  data: {
    a: true, // 普通数据字段
    _b: true, // 纯数据字段
  },
});
```

#### [抽象节点](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/generics.html) <em id="generic-node"></em> <!-- markdownlint-disable-line -->

> 1.9.6 支持

有时, 自定义组件模板中的一些节点，其对应的自定义组件不是由自定义组件本身确定的，而是自定义组件的调用者确定的。这时可以把这个节点声明为**抽象节点**

- 抽象节点 generic 引用的 `generic:xxx="yyy"` 中的值 yyy 只能是静态值, 不能包含数据绑定, 抽象节点特性不适用动态节点绑定的场景
- componentGenerics 抽象节点配置项

index

```html
<!-- index.wxml -->
<!-- 使用抽象节点组件时必须指定使用的具体组件 -->
<view>
  <!-- 自定义单选框 -->
  <generic-node generic:selectable="custom-radio" />
  <!-- 自定义复选框 -->
  <generic-node generic:selectable="custom-checkbox" />
</view>
```

```json
/* index.json */
{
  "usingComponents": {
    "generic-node": "/pages/generic-node",
    "custom-radio": "/pages/custom-radio",
    "custom-checkbox": "/pages/custom-checkbox"
  }
}
```

generic-node

```html
<!-- generic-node.wxml -->
<view>
  <view>generic-node header</view>
  <selectable selected="{{selected}}" disabled="{{false}}"></selectable>
  <view>generic-node footer</view>
</view>
```

```json
/* generic-node.json */
{
  "component": true,
  "usingComponents": {},
  "componentGenerics": {
    "selectable": true /* 在 componentGenerics 中声明抽象节点*/,
    "selectable": {
      "default": "path/to/default/component" /* 为抽象节点指定默认组件 */
    }
  }
}
```

#### [自定义组件扩展](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/extend.html)

- 2.2.3 支持, 详细用法看文档

#### [开发第三方自定义组件](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/trdparty.html)

- 2.2.1 支持

#### [单元测试](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/unit-test.html)

- 2.2.1 支持

#### [获取更新性能统计信息](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/update-perf-stat.html)

- 2.12.0 支持

#### [占位组件](https://developers.weixin.qq.com/miniprogram/dev/framework/custom-component/placeholder.html)

在使用 [分包异步化](https://developers.weixin.qq.com/miniprogram/dev/framework/subpackages/async.html) 或 用时注入 等特性时, 不可用的自定义组件使用 占位组件 临时替代渲染

- 2.17.3 支持

```json
{
  "componentPlaceholder": {
    "comp-a": "view",
    "comp-b": "comp-c"
  }
}
```

### 插件

插件是对一组 js 接口、[自定义组件](#zidingyizujian)或页面的封装, 用于嵌入到小程序中使用. 插件不能独立运行, 必须嵌入到其它小程序中才能使用
第三方小程序在使用插件时, 无法看到插件的代码, 因此, 插件适合用来封装自己的功能或服务, 提供给第三方小程序进行展示和使用.

插件开发者可以像开发小程序一样编写一个插件并上传代码, 在插件发布之后, 其它小程序方可调用, 小程序平台托管插件代码, 其它小程序调用时上传的插件代码会随小程序一起下载运行

插件拥有更强的独立性, 拥有独立的 API 接口、域名列表等, 但同时会受到一些限制, 如 一些 API 无法调用或功能受限

#### 开发插件

插件跳转链接: `plugin-private://PLUGIN_APPID/PATH/TO/PAGE`

- [requireMiniProgram](https://developers.weixin.qq.com/miniprogram/dev/reference/api/requireMiniProgram.html) 在插件中获取由使用者小程序导出的内容, 基础库 2.11.1 支持

```javascript
// 使用者小程序
module.exports = {
  greeting() {
    return 'Greetings from Weixin MiniProgram!';
  },
};

// 插件
const miniProgramExports = requireMiniProgram();
miniProgramExports.greeting(); // "Greetings from Weixin MiniProgram!"
```

##### 插件目录结构

```conf
plugin
  |- components
    |- hello-component.js
    |- hello-component.json
    |- hello-component.wxml
    |- hello-component.wxss
  |- pages
    |- hello-page.js
    |- hello-page.json
    |- hello-page.wxml
    |- hello-page.wxss
  |- index.js
  |- plugin.json
```

##### 插件配置文件

> 向使用者小程序开放的所有自定义组件、页面和 js 接口都必须在插件配置文件 `plugin.json` 中列出

```json
/* plugin.json */
{
  "publicComponents": {
    "hello-component": "components/hello-component"
  },
  "pages": {
    "hello-page": "pages/hello-page"
  },
  "main": "index.js"
}
```

##### 引用小程序的自定义组件

如果需要在页面或自定义组件中将一部分区域交给使用者小程序渲染, 但在插件中不能直接指定使用者小程序的自定义组件路径, 因此无法直接通过 `usingComponents` 得方式来引用, 需要使用 [抽象节点](#generic-node)

插件自定义组件 plugin-view

```json
/* plugin/components/plugin-view.json */
{
  "componentGenerics": {
    "mp-view": true
  }
}
```

```html
<!-- plugin/components/plugin-view.wxml -->
<view>小程序插件组件</view>
<mp-view />
```

小程序中引用 plugin-view

```html
<!-- miniprogram/page/index.wxml -->
<plugin-view generic:mp-view="comp-from-miniprogram" />
```

- 如果是插件页本身就是一个页面顶层组件, 小程序不会引用它, 无法通过 `generic:xxx=""` 的方式来指定抽象节点实现, 因此, 基础库 2.12.2 支持, 小程序可以在插件的配置里为插件页指定抽象节点实现

例如插件页面名为 plugin-index

```json
/* app.json */
{
  /* ... */
  "plugins": {
    "plugin-name": {
      "provider": "wxAPPID",
      "version": "1.0.0",
      /* 2.12.2 为插件页指定抽象节点实现 */
      "genericsImplementation": {
        "plugin-index": {
          "mp-view": "components/comp-from-miniprogram"
        }
      }
    }
  }
  /* ... */
}
```

##### 其它注意事项

- 插件可以预览和上传, 但没有体验版
- 插件可以同时有多个线上版本, 由使用插件的使用者小程序决定具体使用的版本号
- 手机预览和提审插件时, 会使用一个特殊的小程序来套用项目中 miniprogram 文件夹下的小程序, 从而预览插件

插件之间互相调用

- 插件不能直接引用其它插件, 必须在使用者小程序的配置中声明引用之后互相调用
- 对于 js 接口, 可以使用 `requirePlugin`, 但不能直接在文件开头使用, 因为被依赖的插件可能还没有被初始化

#### 使用插件

使用插件之前需要先在小程序管理后台添加插件, 如果插件无需申请, 添加后可直接使用, 否则需要申请并等待插件开发者通过后方可使用

- 主包引用插件

```json
/* app.json */
{
  /* ... */
  "plugins": {
    "plugin-name": {
      "version": "1.0.0",
      "provider": "wxAPPID",
      /* 2.11.1 使用者小程序通过 export 向插件导出内容 */
      /* 插件内使用 requireMiniProgram 全局函数获取使用者小程序导出的内容 */
      "export": "index.js"
    }
  }
  /* ... */
}
```

```javascript
// index.js
module.exports = { whoami: 'Wechat MiniProgram' };

// plugin.js
requireMiniProgram().whoami; // 'Wechat MiniProgram'
```

- 分包引用插件
  - 默认情况下, 仅能在分包内使用当前分包引用的插件, 除非通过 分包异步化 进行异步的跨分包引用
  - 同一个插件不能被多个分包同时引用
  - 如果基础库 < 2.9.0, 不能从分包外的页面直接跳入到分包内的插件页面, 需要先跳入分包内的非插件页面、再跳入同一分包内的插件页面

```json
/* app.json */
{
  /* ... */
  "subPackages": [
    {
      "root": "packageA",
      "name": "" /* 分包别名, 预下载时可用 */,
      "pages": ["pages/page-a", "pages/page-b"],
      "plugins": {
        "plugin-name": {
          "version": "1.0.0",
          "provider": "wxAPPID"
        }
      }
    }
  ]
  /* ... */
}
```

##### 跳转插件页面

`plugin://PLUGIN_NAME/PLUGIN_PAGE`

```html
<navigator url="plugin://PLUGIN_NAME/PLUGIN_PAGE">To plugin page!</navigator>
```

#### 插件使用组件的限制

不能在插件页面中使用

- 开放能力(open-type)为以下之一的 button
  - contact(打开客服会话)
  - getPhoneNumber(获取用户手机号)
  - getUserInfo(获取用户信息)
- open-data
- web-view

以下组件在插件中使用需要基础库版本支持

- navigator 需要基础库 2.1.0
- live-player 和 live-pusher 需要基础库 2.3.0

#### 插件功能页

> 2.1.0 支持, 使用插件功能页之前, 先激活功能页特性, 配置对应的功能页函数, 再使用 `functional-page-navigator` 组件跳转到插件功能页

插件功能页是插件所有者小程序中的一个特殊页面, 使用插件功能页可以完成某些接口或 API 在插件中的调用限制

- 插件所有者小程序配置激活插件功能页

```json
/* app.json */
{
  /* ... */
  "functionalPages": {
    /* independent: true 表示插件功能页的代码独立于其它代码,*/
    /* 这意味着插件功能页可以被独立下载、加载, 具有更好的性能表现 */
    /* 但同时使插件功能页目录 functional-page/ 不能 require 这个目录以外的文件(反之亦然) */
    "independent": true
  },
  "functionalPages": true /* 旧式写法 */
  /* ... */
}
```

- 跳转到插件功能页
  - version 跳转到的小程序版本, 线上版本必须是 release
  - name 要跳转的功能页
  - args 功能页参数, 格式与具体功能页相关
  - bind:success 功能页返回且操作成功时触发, detail 格式与具体功能页相关
  - bind:fail 功能页返回且操作失败时触发, detail 格式与具体功能页相关
  - bind:cancel 因用户操作从功能页返回时触发

不能使用 wx.navigateTo 进行跳转, 需要使用 `functional-page-navigator` 组件跳转

```html
<functional-page-navigator
  name="loginAndGetUserInfo"
  args=""
  version=""
  bind:success=""
  ><button>登录到插件</button>
</functional-page-navigator>
```

##### 用户信息功能页

用于帮助插件获取用户信息, 相当于 wx.login 和 wx.getUserInfo 的功能

##### 支付功能页

支付功能页用于帮助完成支付, 相当于 wx.requestPayment 的功能

插件使用支付功能时, 需要在管理后台进行额外的权限申请, 主体为个人小程序在使用插件时, 无论是否通过申请都无法正常使用插件里的支付功能

- 2.22.1 开始, 插件内可以直接调用 wx.requestPluginPayment 实现跳转支付, 通过 `functional-page-navigator` 跳转将被废弃
- 满足以下条件时, 调用 wx.requestPluginPayment 或点击 navigator 将直接拉起支付收银台, 跳过支付功能页
  - 小程序与插件绑定在同一个 open 平台账号上
  - 小程序与插件均为 open 账号的同主体/关联主体时

```javascript
// functional-page/request-payment.js
exports.beforeRequestPayment = function (paymentArgs, callback) {
  // paymentArgs 通过 functional-page-navigator 的 arg 参数中 paymentArgs 字段传递到此功能页的自定义数据
  // callback 回调函数, 调用该函数后, 小程序将发起支付, 类似于 wx.requestPayment
};
```

###### 配置功能页函数

支付功能页需要在插件中提供一个函数来响应插件中的支付调用, 即在插件中跳转到支付功能页或调用 wx.requestPluginPayment 时, 这个函数就会在合适的时机被调用来帮助完成支付, 如果不提供功能页函数, 功能页将通过 fail 事件返回错误

支付功能页函数应以导出函数的方式提供在插件所有者小程序的根目录下的 functional-page/request-payment.js 文件中名为 beforeRequestPayment

##### 收获地址功能页

- 2.16.1 开始, 插件内可直接使用 wx.chooseAddress 实现对应的功能, 点击 `functional-page-navigator` 将不再进入功能页

##### 发票功能页

- 2.16.1 开始, 插件内可直接使用 wx.chooseInvoice 实现对应的功能, 点击 `functional-page-navigator` 将不再进入功能页

##### 发票抬头功能页

- 2.16.1 开始, 插件内可直接使用 wx.chooseInvoiceTitle 实现对应的功能, 点击 `functional-page-navigator` 将不再进入功能页

### 基础能力

### 连接硬件能力

### 开放能力
