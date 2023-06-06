---
title: vagrant
date: 2022-04-29 17:19:28
categories:
  - tools
tags:
  - vagrant
---

## 敲黑板

### vagrant 命令行报错 Encoding::UndefinedConversionError

使用 vagrant 命令时提示 process_builder.rb:44:in `encode!': "\\xE5" to UTF-8 in conversion from ASCII-8BIT to UTF-8 to UTF-16LE (Encoding::UndefinedConversionError)

本例使用的 vagrant 版本为 2.2.10, 安装目录：\<br/\> D:\HashiCorp\Vagrant\embedded\gems\2.2.10\gems\childprocess-4.0.0\lib\childprocess\windows\process_builder.rb 第 44 行

修改 newstr.encode!('UTF-16LE') 为 \<br/\> newstr.encode!('UTF-16LE', invalid: :replace, undef: :replace, replace: '?') [参考连接](https://blog.csdn.net/qq_41606390/article/details/122854431)

### vagrant 自制 box 启动时 Authentication failure

- vagrant 2.3.6
- virtualBox 7.0.8
- win 11

> Authentication failure 是 ssh 登录证书错误, 但虚拟机已经启动完成

- `.ssh/authorized_keys` 或 `.vagrant/machines/default/virtualbox/private_key` 有变化(内容改动/文件存在)

```shell
...
==> default: Forwarding ports...
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Authentication failure. Retrying...
    default: Warning: Authentication failure. Retrying...
    default: Warning: Authentication failure. Retrying...
    default: Warning: Authentication failure. Retrying...
...
```

<!-- more -->

因为 vagrant 官方提供了一对默认的 keypair,
公钥预先存放在 box 的 vagrant 家目录 `.ssh/authorized_keys` 文件中,
私钥存放在宿主机 vagrant 安装目录的 .vagrant.d 下的 `insecure_private_key` 文件中.

当 vagrant up 启动虚拟机时, 第一次登录虚拟机使用的是官方提供的 keypair, 然后 vagrant 在宿主机生成一对新的 keypair,
将公钥更新虚拟机中 vagrant 家目录 `.ssh/authorized_keys` 文件内容,
将私钥存放到初始化虚拟机目录的 `.vagrant/machines/default/virtualbox/private_key` 文件中.

> 为了安全考虑才会有替换公钥的过程, 否则任何人使用官方公钥都可以登录

解决办法:

- 通过用户密码登录虚拟机
- 使用命令 `wget` 将官方不安全的公钥更新指定文件 `/home/vagrant/.ssh/authorized_keys`
  - \-O file, 指定文件名
  - `https://raw.githubusercontent.com/hashicorp/vagrant/master/keys/vagrant.pub`
- 使用命令 `chmod` 修改 `authorized_keys` 文件仅属主可读写
- 退出并关闭虚拟机
- 使用命令 `vagrant package` 重新打包
- 使用命令 `vagrant box add` 导入打包的 box 文件
- 使用导入的 box 创建并启动虚拟机
