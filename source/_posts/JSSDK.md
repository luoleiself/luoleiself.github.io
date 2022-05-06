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

- 使用版本 https://res.wx.qq.com/open/js/jweixin-1.4.0.js

##### 企微(4.0.6) 中使用此 API 的位置

- 转发
- 转发给客户
- 群发到客户群
- 发表到客户的朋友圈
- 分享到同事吧
- 分享到微信
- 分享到微信朋友圈使用此 onMenuShareTimeline, 但无法设置分享信息

#### [updateAppMessageShareData](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html)

- 使用版本 https://res.wx.qq.com/open/js/jweixin-1.4.0.js

- 企微(4.0.6) 使用此 API 报错: Uncaught TypeError: wx.updateAppMessageShareData is not a function

#### [updateTimelineShareData](https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html)

- 使用版本 https://res.wx.qq.com/open/js/jweixin-1.4.0.js

- 企微(4.0.6) 使用此 API 报错: Uncaught TypeError: wx.updateTimelineShareData is not a function
