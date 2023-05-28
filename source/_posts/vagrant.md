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

### 修改默认虚拟电脑存储位置启动虚拟机报错

- windows 11
- vagrant 2.3.6
- VirtualBox 7.0.8

win 11 下, 使用 VirtualBox 的默认虚拟电脑存储位置(在C盘)时, vagrant 可以正常启动虚拟机,

如果修改 VirtualBox 的默认虚拟电脑存储位置为 D:\VirtualBox\vms 时, vagrant 启动虚拟机会提示以下错误

重新安装 VirtualBox 6.1 版本可以正常操作

Bringing machine 'default' up with 'virtualbox' provider...

==> default: Importing base box 'hashicorp/bionic64'...

There was an error while executing `VBoxManage`, a CLI used by Vagrant

for controlling VirtualBox. The command and stderr is shown below.


Command: ["import", "\\\\?\\C:\\Users\\Administrator\\.vagrant.d\\boxes\\hashicorp-VAGRANTSLASH-bionic64\\1.0.282\\virtualbox\\box.ovf", "--vsys", "0", "--vmname", "ubuntu-18.04-amd64_1685247283628_30042", "--vsys", "0", "--unit", "11", "--disk", "D:/VirtualBox/vms/ubuntu-18.04-amd64_1685247283628_30042/ubuntu-18.04-amd64-disk001.vmdk"]


Stderr: 0%...10%...20%...30%...40%...50%...60%...70%...80%...90%...100%

Interpreting \\?\C:\Users\Administrator\.vagrant.d\boxes\hashicorp-VAGRANTSLASH-bionic64\1.0.282\virtualbox\box.ovf...

OK.

0%...E_INVALIDARG

VBoxManage.exe: error: Appliance import failed

VBoxManage.exe: error: Code E_INVALIDARG (0x80070057) (extended info not available)

VBoxManage.exe: error: Context: "enum RTEXITCODE __cdecl handleImportAppliance(struct HandlerArg *)" at line 1416 of file VBoxManageAppliance.cpp
