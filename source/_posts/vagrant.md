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

本例使用的 vagrant 版本为 2.2.10, 安装目录：<br/> D:\HashiCorp\Vagrant\embedded\gems\2.2.10\gems\childprocess-4.0.0\lib\childprocess\windows\process_builder.rb 第 44 行

修改 newstr.encode!('UTF-16LE') 为 <br/> newstr.encode!('UTF-16LE', invalid: :replace, undef: :replace, replace: '?') [参考连接](https://blog.csdn.net/qq_41606390/article/details/122854431)
