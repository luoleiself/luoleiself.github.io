---
title: Charles-手机代理-Https-证书管理
date: 2021-12-14 14:37:01
categories:
  - tools
tags:
  - charles
---

### 敲黑板

- 开启 charles 无法代理 https, 检查电脑和手机上的 charles 证书是否过期

### 电脑删除证书

1. 使用 win + r 键调起系统运行窗口, 在输入框中输入 mmc 命令后回车
   ![charles-11](/images/charles-11.jpg)
2. 在弹出的对话框选择 "文件" -> "添加或删除管理单元"
   ![charles-12](/images/charles-12.png)

<!-- more -->

3. 在弹框中左侧列表选择 "证书", 点击 "添加" 按钮, 在弹出的对话框选择 "我的用户账户" 点击 "完成" 后, 右侧列表会出现 证书 选项, 最后点击 "确定"
   ![charles-13](/images/charles-13.jpg)
   ![charles-14](/images/charles-14.jpg)
4. 在第 2 步的窗口中会出现证书菜单, 依次展开这些菜单, 中间窗口会出现所有已安装的证书
   ![charles-15](/images/charles-15.jpg)
5. 选择需要删除的证书 "右键单击", 选择 "删除", 在弹出的对话框中选择 "是", 完成证书删除
   ![charles-16](/images/charles-16.jpg)
6. 点击右上角的 "关闭" 按钮关闭控制台, 在弹出的对话框中选择 "否" 不存入控制台配置, 如果有需要也可以选择 "是" 存入控制台配置
   ![charles-17](/images/charles-17.jpg)

### 电脑安装 charles 证书

1. 打开 charles, 选择 Help -> SSL Proxing -> Install Charles Root Certificate
   ![charles-1](/images/charles-1.png)

2. 在 "证书" 对话框中点击 "安装证书"
   ![charles-2](/images/charles-2.jpg)
3. 在 "证书导入向导" 对话框中选择 "当前用户" 或者 "本地计算机" 均可, 点击 "下一步"
   ![charles-3](/images/charles-3.jpg)
4. 在 "证书导入向导" 对话框中选择 "将所有的证书都放入下列存储", 点击 "浏览"
   ![charles-4](/images/charles-4.jpg)
5. 在 "选择证书存储" 对话框中选择 "受信任的根证书颁发机构", 点击 "确认"
   ![charles-5](/images/charles-5.jpg)
6. 在弹出的对话框选择 "是" 将证书导入到浏览器的证书列表中
   ![charles-10](/images/charles-10.jpg)
7. 点击 "确定" 完成证书添加
   ![charles-6](/images/charles-6.jpg)
8. charles 部分版本提示不能抓 HTTPS 的, 在 SSL Proxing Settings 下 Add \*:\* 允许所有 host 和端口

### 手机添加 charles 证书

1. 手机绑定 charles 代理(电脑 IP 和端口):
   1. WLAN -> 修改已连接网络
   2. 代理改为 "手动"
   3. 服务器主机名添加电脑 IP
   4. 服务器端口添加 charles 监听端口
   5. 保存后连接 wifi
2. 安卓手机使用 Chrome, 苹果手机使用 safari
3. 浏览器地址中输入 chls.pro/ssl 下载手机证书, 如果手机绑定代理后无法联网可以先去掉代理
4. 安卓手机双击已下载证书, 在弹出的对话框填写以下信息并确定
   1. 输入 "证明名称(可任意起名)", 后期删除证书时可根据此名称查找
   2. 凭证用途 设置为 "VPN 和应用", 改为 WLAN 会导致无效
5. 苹果手机打开设置 -> 通用 -> 关于本机 -> 证书信任设置 -> 找到已下载的 charles 证书打开信任
6. 使用手机访问代理, 如果访问失败, 手机重新连接 wifi 重试
   ![charles-7](/images/charles-7.jpg)

### 电脑使用 charles 无法访问百度等其他页面

黑名单和白名单设置

1. 打开 charles, 选择 Tools -> Block List | Allow List( 部分版本为 White List)
   ![charles-8](/images/charles-8.png)
2. 此配置为黑名单和白名单设置
   ![charles-9](/images/charles-9.jpg)

### 手机微信扫码后显示白屏

- 手机未安装 charles 证书会显示白屏
- 其他部分浏览器会在 地址栏 中出现非安全链接的提示图标

### 手机卸载安装的 charles 证书

1. 打开手机 设置 -> 安全 -> 更多安全设置 -> 加密和凭据 -> 用户凭据
2. 根据安装手机证书时填写的证明名称找到指定项, 点击后选择 "删除"
