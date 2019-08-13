# smartapi
  个人文件存储服务器

# 项目结构：
```

├─doc 文档
│  ├─readme.md 公共参数说明
│  ├─user.md 用户接口文档
├── etc
│   ├── config.json 配置文件
│   └── database
│       └── init_db.sql 数据库初始化sql脚本
├── internal
│   ├── common 公共文件目录
│   │   └── config.go 
│   ├── log 日志库
│   │   ├── log.go
│   │   ├── signal.go
│   │   └── signal_windows.go
│   ├── service 接口服务目录
│   │   ├── apis
│   │   │   ├── auth.go 接口鉴权中间件
│   │   │   ├── error_code.go 错误码定义
│   │   │   ├── p.go 请求参数，返回参数定义
│   │   │   ├── router.go api路由
│   │   │   └── user_handler.go 用户的handler
│   │   └── models
│   │       ├── init.go 数据库连接初始化
│   │       └── user_model.go 表格定义
│   └── tools 工具库
│       ├── util.go
│       └── util_test.go
├── main.go 
├── Makefile 
└── README.md
```

# ToDo
> 文件系统注册、登陆
> 文件上传下载
> 文件共享
> 文件资源分类
> 评论、统计

