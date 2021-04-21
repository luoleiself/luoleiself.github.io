---
title: Git命令操作篇
date: 2021-04-20 16:48:57
categories:
  - tools
tags:
  - git
---

### 操作篇

#### 初始化仓库

```bash
  git init [project-name]
```

#### 克隆仓库

```bash
  git clone [url]
```

#### 分支

##### 查看

- a|all 列出所有本地和关联远程分支
- r 列出关联远程分支
- v|verbose 列出分支并显示当前提交信息摘要

  ```bash
    git branch [-a] [-r] [-v]
  ```

##### 新建

```bash
  git branch <branch-name>
```

##### 移动|修改

- m 移动|修改分支, old-branch 无则为移动操作
- M 强制移动|修改分支即使新分支存在, old-branch 无则为移动操作

  ```bash
    git branch [-m] [-M] [<old-branch>] <new-branch>
  ```

##### 复制

- c 复制分支和分支提交历史
- C 强制复制分支和分支提交历史

  ```bash
    git branch [-c] [-M] [<old-branch>] <new-branch>
  ```

##### 删除

- d 删除本地分支,一般和 r 配合删除关联远程分支
- D 强制删除本地分支, 即使分支未被合并

  ```bash
    git branch [-d] [-D] <branch-name>
  ```

##### 切换

```bash
  git checkout [branch-name]
```

###### 切换并创建新分支

基于远程分支创建新分支,自动建立追踪关系

```bash
  git checkout -b <branch-name> [-t] [<remote-branch>]
```

##### 分支追踪关系

- t|no-track 建立|取消分支追踪关系
- u|unset-upstream 建立|取消分支追踪关系

###### 当前分支和远程分支建立追踪关系

```bash
  git branch -t <remote-branch>
```

###### 指定分支和远程分支建立追踪关系

```bash
  git branch -u <local-branch> <remote-branch>
```

##### 分支追踪关系,提交摘要

```bash
  git branch [-vv]
```

![upstream](../../images/git-branch-1.jpg)

##### 分支合并

###### merge

- fast-forward(ff): 快速合并, 不创建新的 commit, 原分支删除后提交记录消失, 默认方式
- no-ff: 不快速合并, 保留原有分支记录, 创建新的 commit
- squash: 合并一些不必要的 commit, 创建新的 commit
- stat: 合并结束后统计显示区别
- continue: 解决冲突后结束合并
- abort: 中断解决冲突结束合并
- quit: 放弃合并

```bash
  git merge [--no-ff] <branch-name>
```

![merge](../../images/git-branch-2.png)

###### rebase

- i|interactive: 交互式操作
- continue: 解决冲突后结束合并
- abort: 中断解决冲突结束合并
- quit: 放弃合并
- skip: 重启合并跳过当前的修改

```bash
  git rebase <branch-name>
```

![rebase](../../images/git-branch-3.png)

###### 选择合并

- 选择一个或者多个 commit, 合并进当前分支, 手动 commit

  ```bash
    git cherry-pick [--no-commit|-n] <commit-ish>
  ```

- 选择 commit 区间合并, 含尾不含头

  ```bash
    git cherry-pick [--ff] commit1...commitN
  ```

- 选择 commit 区间合并, 包含头和尾

  ```bash
    git cherry-pick [--ff] commit1^...commitN
  ```

#### 文件操作

##### 添加文件

- all|A|. 添加所有文件提交信息列表
- i|interactive 交互式操作
- n|dry-run 不执行任何操作, 只显示做什么

```bash
  git add [-A] [-i] [-n] [<file>...]
```

##### 撤销文件

- 撤销工作区文件的变更

  ```bash
    git checkout -- [.|<file>...]
  ```

- 恢复上一个 commit 的所有文件到工作区

  ```bash
    git checkout .
  ```

- 恢复指定 commit 的指定文件到工作区

  ```bash
    git checkout [commit] [file]
  ```

##### 提交

- a|all git add -A 的缩写
- m|message commit 的注释
- amend 改写上一次 commit 的注释

```bash
  git commit [-am] [<file>...]
```

##### 版本变更

###### reset

- mixed: 还原版本库和暂存区, 工作区保持不变, 默认方式
- soft: 还原版本库, 暂存区和工作保持不变
- hard: 还原版本库和暂存区和工作区
- keep: 还原版本库和暂存区, 并更新工作区中的 commit 和 HEAD 之间不同的文件, 如果不同的文件本地有更改则中止
- merge: 还原版本库和暂存区, 并更新工作区中的 commit 和 HEAD 之间不同的文件, 但保留暂存区和工作区中不同的文件(既有未添加的更改), 如果不同的文件有未暂存的更改则中止

  ```bash
    git reset 3f405f2
    git reset --soft HEAD^
    git reset --hard HEAD~3
  ```

![reset](../../images/git-branch-4.png)

###### revert

撤销一个或多个 commit, 并手动提交 commit

```bash
  git revert [--no-commit|-n] [<commit-ish>...]
```

##### 比较

---

#### 暂存

保存当前工作区的状态以备以后继续使用并恢复干净的工作区

- list: 显示暂存区暂存记录
- show: 显示暂存区记录和当前工作区的不同
- pop: 取出指定的 stash 还原到工作区中并从暂存区中移除
- apply: 取出指定的 stash 还原到工作区不会从暂存区移除
- clear: 清空暂存区
- drop: 从暂存区移除指定的 stash
- create: 创建一个 stash 对象并返回 commit
- store: 使用返回的 commit 生成 stash 记录

```bash
  git stash
  git stash list
  git stash show|pop|apply|drop [<stash@{0}>]
  git stash clear
  git stash create [<message>]
  git stash store [-m <message>] <commit>
```

![stash](../../images/git-stash-1.jpg)

---

#### 日志

---

#### 远程同步

---
