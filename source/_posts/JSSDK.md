---
title: JSSDK
date: 2022-05-06 12:19:31
categories:
  - wx
tags:
  - wx
  - jweixin
  - JSSDK
---

### 敲黑板

#### [onMenuShareAppMessage](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html) (微信提示: 即将废弃)

- jweixin-1.4.0.js

##### 企微(4.0.6) 中使用此 API 的位置

- 转发
- 转发给客户
- 群发到客户群
- 发表到客户的朋友圈
- 分享到同事吧
- 分享到微信
- 分享到微信朋友圈使用此 onMenuShareTimeline, 但无法设置分享信息

<!-- more -->

#### [updateAppMessageShareData](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html)

- jweixin-1.4.0.js

- 企微(4.0.6) 使用此 API 报错: Uncaught TypeError: wx.updateAppMessageShareData is not a function

```javascript
wx.updateAppMessageShareData({
  title: '', //分享信息的标题
  desc: '', //分享信息的描述
  link: '', //分享信息的链接
  imgUrl: '', //分享信息的图片
  success: function () {},
});
```

#### [updateTimelineShareData](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html)

- jweixin-1.4.0.js

- 企微(4.0.6) 使用此 API 报错: Uncaught TypeError: wx.updateTimelineShareData is not a function

```javascript
wx.updateTimelineShareData({
  title: '', //分享信息的标题
  desc: '', //分享信息的描述
  link: '', //分享信息的链接
  imgUrl: '', //分享信息的图片
  success: function () {},
});
```

#### [wx.invoke](https://developer.work.weixin.qq.com/document/path/90490)

- jweixin-1.4.0.js

```javascript
wx.invoke(
  'shareAppMessage',
  {
    title: '', //分享信息的标题
    desc: '', //分享信息的描述
    link: '', //分享信息的链接
    imgUrl: '', //分享信息的图片
  },
  function (res) {
    if (res.err_msg == 'shareAppMessage:ok') {
    }
  }
);
```

- shareAppMessage 自定义转发到会话
- shareWechatMessage 自定义转发到微信
- 企业微信 2.4.5 及以后版本支持,微信客户端不支持(微信开发者工具也不支持)
- 使用此 API 时, 在 `企微(4.0.6.6516)` 桌面版上打开页面时会自动弹出 转发到会话 弹框, 注释此方法后则没有自动弹框, 企微手机端不受此影响
