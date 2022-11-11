---
title: Redis-Set
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

### Set 命令

Set 是无序不重复的集合

- SCARD key 获取 set 成员的数量, set 为空或者不存在返回 0
- SADD key member [member ...] 向 set 中添加多个成员并返回添加成功的数量, 0 表示有重复成员
