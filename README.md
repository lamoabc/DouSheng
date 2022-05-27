# 抖声


-----

## 介绍
字节跳动青训营项目

## [文档说明](https://gitee.com/set-sail0/DouSheng/tree/develop/Doc)

抖声APP,服务端配置以及功能说明

## [代码](https://gitee.com/set-sail0/DouSheng/tree/master/Code)

抖声App,不同版本服务端代码

## 功能以及环境简要说明

### 项目正常运行需要安装依赖

```golang
// gin
go get -u github.com/gin-gonic/gin

// gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

// oss
go get github.com/aliyun/aliyun-oss-go-sdk/oss

// JWT
go get -u github.com/dgrijalva/jwt-go
```

### 目前实现接口

[接口功能详细说明]()

#### 基础接口

1. 注册接口
2. 登录接口
3. 视频流接口
4. 视频投稿接口
5. 发布列表接口
6. 用户信息接口
