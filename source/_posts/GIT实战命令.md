---
title: GIT实战命令
date: 2021-04-17 18:25:55
categories:
  - tools
tags:
  - git
---

[![概念图](../../images/git-flow-1.jpg)](https://git-scm.com/)

### 配置篇

#### 重置 Git 仓库用户凭据

场景: 使用 https 方式与远程仓库同步时, 不想每次都在提示框中输入用户名和密码确认, 此方式在用户家目录下创建(修改) .git-credentials 文件, 存储用户名和密码

```bash
  git config --system --unset credential.helper # 清除本地存储的用户名和密码凭据
  git config --global credential.helper store # 存储凭据, 在第一次 push 或者 pull 时提示输入
```

#### 怎么舒服怎么来
