---
title: Redis-PubSub
date: 2022-11-09 18:03:19
categories:
  - [server, Redis]
tags:
  - Redis
---

## 发布订阅

Redis 发布/订阅(pub/sub)是一种消息通信模式: 发送者(pub)发送消息, 订阅者(sub)接收消息
它采用事件作为基本的通信机制，提供大规模系统所要求的松散耦合的交互模式: 订阅者(如客户端)以事件订阅的方式表达出它有兴趣接收的一个事件或一类事件;发布者(如服务器)可将订阅者感兴趣的事件随时通知相关订阅者
订阅者对一个或多个频道感兴趣,只需接收感兴趣的消息,不需要知道什么样的发布者发布的. 这种发布者和订阅者的解耦合可以带来更大的扩展性和更加动态的网络拓扑

- 发布者: 无需独占链接, 可以在 publish 发布消息的同时, 使用同一个链接进行其他操作
- 订阅者: 需要独占链接, 在 subscribe 期间, 以阻塞的方式等待消息

### 发布消息

- PUBLISH channel message 给指定的频道发送消息并返回接收到消息的订阅者数量, 0 表示没有订阅者
- SPUBLISH shardchannel message 给指定的碎片频道发送消息并返回接收到消息的订阅者数量, 0 表示没有订阅者, 7.0.0 支持

### 普通订阅

- SUBSCRIBE channel [channel ...] 订阅指定频道立即进入阻塞状态等待接收消息
- UNSUBSCRIBE [channel [channel ...]] 根据给定频道取消客户端订阅, 如果未指定则取消所有频道订阅

<!-- more -->

```bash
# 1
127.0.0.1:6379> SUBSCRIBE first second
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "first"
3) (integer) 1
1) "subscribe"
2) "second"
3) (integer) 2
# 2
127.0.0.1:6379> SUBSCRIBE first third
Reading messages... (press Ctrl-C to quit)
1) "subscribe"
2) "first"
3) (integer) 1
1) "subscribe"
2) "third"
3) (integer) 2
# 3
127.0.0.1:6379> PUBSUB CHANNELS
1) "third"
2) "first"
3) "second"

# 3
127.0.0.1:6379> PUBLISH first 'hello first'
(integer) 2
# 1
127.0.0.1:6379> SUBSCRIBE first second
...
1) "message"
2) "first"
3) "hello first"
# 2
127.0.0.1:6379> SUBSCRIBE first third
...
1) "message"
2) "first"
3) "hello first"

# 3
127.0.0.1:6379> PUBLISH second 'hello second'
(integer) 1
# 1
127.0.0.1:6379> SUBSCRIBE first second
...
1) "message"
2) "second"
3) "hello second"

# 3
127.0.0.1:6379> PUBLISH third 'hello third'
(integer) 1
# 2
127.0.0.1:6379> SUBSCRIBE first third
...
1) "message"
2) "third"
3) "hello third"
```

### 模式订阅

- PSUBSCRIBE pattern [pattern ...] 根据给定模式订阅频道立即进入阻塞状态等待接收消息
  - pattern 可以使用正则表达式匹配多个频道
- PUNSUBSCRIBE [pattern [pattern ...]] 根据给定模式取消客户端订阅, 如果未指定则取消所有模式订阅

```bash
# 1
127.0.0.1:6379> PSUBSCRIBE __key*__:*
Reading messages... (press Ctrl-C to quit)
1) "psubscribe"
2) "__key*__:*"
3) (integer) 1
# 2
127.0.0.1:6379> PSUBSCRIBE __key*__:*
Reading messages... (press Ctrl-C to quit)
1) "psubscribe"
2) "__key*__:*"
3) (integer) 1
# 3
127.0.0.1:6379> PUBSUB NUMPAT
(integer) 1

# 3
127.0.0.1:6379> PUBLISH __key@__:foo 'hello key at foo'
(integer) 2
# 1
127.0.0.1:6379> PSUBSCRIBE __key*__:*
...
1) "pmessage"
2) "__key*__:*"
3) "__key@__:foo"
4) "hello key at foo"
# 2
127.0.0.1:6379> PSUBSCRIBE __key*__:*
...
1) "pmessage"
2) "__key*__:*"
3) "__key@__:foo"
4) "hello key at foo"

# 3
127.0.0.1:6379> PUBLISH __key@__:bar 'hello key at bar'
(integer) 2
# 1
127.0.0.1:6379> PSUBSCRIBE __key*__:*
...
1) "pmessage"
2) "__key*__:*"
3) "__key@__:bar"
4) "hello key at bar"
# 2
127.0.0.1:6379> PSUBSCRIBE __key*__:*
...
1) "pmessage"
2) "__key*__:*"
3) "__key@__:bar"
4) "hello key at bar"
```

### 碎片频道订阅

- SSUBSCRIBE shardchannel [shardchannel ...] 订阅指定的碎片频道, 7.0.0 支持
- SUNSUBSCRIBE [shardchannel [shardchannel ...]] 根据给定碎片频道取消客户端订阅, 如果未指定则取消所有碎片频道订阅, 7.0.0 支持

### 统计订阅信息

```bash
127.0.0.1:6379> PUBSUB HELP
 1) PUBSUB <subcommand> [<arg> [value] [opt] ...]. Subcommands are:
 2) CHANNELS [<pattern>]
 3)     Return the currently active channels matching a <pattern> (default: '*').
 4) NUMPAT
 5)     Return number of subscriptions to patterns.
 6) NUMSUB [<channel> ...]
 7)     Return the number of subscribers for the specified channels, excluding
 8)     pattern subscriptions(default: no channels).
 9) SHARDCHANNELS [<pattern>]
10)     Return the currently active shard level channels matching a <pattern> (default: '*').
11) SHARDNUMSUB [<shardchannel> ...]
12)     Return the number of subscribers for the specified shard level channel(s)
13) HELP
14)     Prints this help.
```

- PUBSUB CHANNELS [pattern] 返回当前活跃频道列表(不包含使用模式订阅的频道)
- PUBSUB NUMSUB [channel [channel ...]] 返回订阅者的数量(不包含使用模式订阅的频道)
  - 如果不指定 channel 将返回 (empty array)

```bash
127.0.0.1:6379> PUBSUB CHANNELS
1) "conn"
2) "__sentinel__:hello"

127.0.0.1:6379> PUBSUB NUMSUB hello conn
1) "hello"
2) (integer) 1
3) "conn"
4) (integer) 1
```

- PUBSUB NUMPAT 返回订阅者通过模式订阅的频道的数量

```bash
127.0.0.1:6379> PUBSUB NUMPAT
(integer) 0
127.0.0.1:6379> PUBSUB NUMPAT
(integer) 1
```

- PUBSUB SHARDCHANNELS [pattern] 返回当前活动的碎片频道, 未找到返回 empty array, 7.0.0 支持
- PUBSUB SHARDNUMSUB [shardchannel [shardchannel ...]] 返回指定的碎片频道的订阅者数量, 未找到返回 empty arryay, 7.0.0 支持

```bash
127.0.0.1:6379> PUBSUB SHARDNUMSUB conn
1) "conn"
2) (integer) 0
```
