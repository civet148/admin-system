# 1. 代码目录结构 


```shell script
├── cmd      # 服务程序
├── pkg      # 业务逻辑代码
    ├── api           # API接口
    ├── cache         # 验证码缓存
    ├── config        # 配置相关
    ├── controllers   # HTTP接口逻辑实现
    ├── crypto        # 加解密包
    ├── dal           # 数据访问层（数据库相关代码逻辑）                 
      ├── core          # 数据库业务代码（入口）
      ├── dao           # 对应每张表的数据库操作代码
      ├── db2go         # 自动由数据表结构生成数据模型脚本(执行gen_models.bat按提示下载db2go程序）
      └── models        # 数据库表结构定义（由db2go工具自动生成）
    ├── middleware    # 中间件（主要是JWT相关）
    ├── privilege     # casbin权限管理
    ├── proto         # 管理系统前端和后端HTTP通信协议
    ├── routers       # HTTP路由
    ├── services      # 服务代码目录
    ├── static        # 静态页面测试目录（hello.html）
    ├── storage       # 本地存储
    ├── sessions      # 用户session管理
    ├── types         # 公共常量/结构定义
    └── utils         # 工具包
├── deploy    # 部署脚本和SQL文件
├── docs      # 开发文档目录
├── dockers   # docker镜像编译以及测试运行脚本
├── static    # 静态文件目录
```
# 2. 编译运行

## 2.1. 物理机编译
```shell script
make
```

## 2.2. docker镜像编译

```shell
make docker 
```

## 2.3. 本机运行

```shell script
# --debug       打开调试模式(默认日志级别：info)
# --static      指定前端静态页面目录(不指定则默认当前路径下static目录)
./admin-system run --debug --static /opt/frontend/dist 0.0.0.0:8088
```

## 2.4. 容器运行/停止

```shell
# 启动
./run.sh
# 停止
./stop.sh
```