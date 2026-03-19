### [node.green](http://node.green/)

### nvm 配置 node 全局模块

msi 在安装步骤中修改此设置, 同时会自动创建 `NVM_HOME` 和 `NVM_SYMLINK` 环境变量并在 Path 中导出

- nvm 安装目录下修改 settings.txt (安装程序自动修改以下路径)
  - root: D:\nvm # nvm 的安装路径和多版本 node 的存放路径
  - path: D:\nvm\nodejs # 当前版本 node 的工作路径
  - node_mirror: 可选项, node 指定版本的源, 默认为 <https://nodejs.org/dist/>
  - npm_mirror: 可选项, npm 指定版本的源, 默认为 <https://github.com/npm/cli/archive/>

#### nvm 安装目录下新建 node_cache 和 node_global 目录

#### 用户家目录下修改 .npmrc 文件或使用 npm 命令设置 prefix 和 cache

- 修改配置文件 .npmrc

```txt
cache=D:\nvm\node_cache
prefix=D:\nvm\node_global
```

- 使用 npm 命令修改

```bash
# 修改
npm config set cache "D:\nvm\node_cache" # 修改缓存目录
npm config set prefix "D:\nvm\node_global" # 修改全局模块存储目录

# 查询
npm config get cache
npm config get prefix
npm config ls -l # 查看所有的配置项
```

#### 创建 NODE_PATH 环境变量

NODE_PATH (可为任意名字) 值为 node_global 目录绝对路径, 并在 path 中导出

#### 安装全局模块并切换 node 版本验证

```bash
npm instal -g yarn # 安装全局 yarn
npm ls -g --depth 0 # 查看全局所有模块
yarn -v # 查看 yarn 版本

nvm use 20.16.0 # 切换 node 版本

npm ls -g --depth 0 # 查看全局所有模块
yarn -v # 查看 yarn 版本
```

### npm命令

- npm install modulesName@version  //安装版本

- npm install modulesName -S/--save //生产阶段依赖;
- npm install modulesName -D/--save-dev   //开发阶段依赖;
- npm install modulesName -O/--save-optional  //可选阶段的依赖;
- npm install modulesName -E/--save-exact     //安装指定版本;
- npm uninstall modulesName       //卸载指定模块
- npm outdated                    //检测模块是否过时
- npm ls -g --depth 0          //查看安装的模块
- npm init        //初始化项目，创建package.json文件，保存安装包的信息
- npm info package // 查看该模块的基本信息
- npm help        //查看某条命令的详细帮助
- npm home package // 浏览器打开该模块的首页
- npm root  -g  //查看系统包的安装路径
- npm config  list  -l  //查看config配置列表详情
- npm start       //启动模块
- npm cache      //管理模块的缓存
- npm stop        //停止模块
- npm restart    //重新启动模块
- npm test       //测试模块
- npm version    //查看模块版本
- npm view       //查看模块的注册信息
- npm publish     //发布模块
- npm config set prefix "globalDirPath"
- npm config set cache "cacheDirPath"

### [README_E.md](https://github.com/luoleiself/summary/blob/master/NodeJs/README_E.md)
