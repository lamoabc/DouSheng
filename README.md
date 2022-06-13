# 抖声


-----

## 介绍
字节跳动青训营项目

## [文档说明](https://github.com/jhy625/DouSheng/tree/master/Doc)

抖声APP,服务端配置以及功能说明

## [代码](https://github.com/jhy625/DouSheng/tree/master/Code)

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

### 下载此项目并运行的方式

1. 使用官方提供的APP,配置ip为47.95.117.216即可使用

   **注意:如果使用官方提供的APP,并且下载手机上配上面提供ip使用,投稿的发布按钮会没有反应,这个问题是客户端的错误,抓包后会发现,会报408错误,这个错误的主要原因是客户端发送数据时,有字段的长度没有指定,服务端校验会发生错误,此问题不确定是否跟手机型号有关,在最后的项目不同版本介绍中有解决办法**

2. 下载项目至本地,将其中连接远程数据库的配置以及OSS的配置更改为自己的账号也可以使用

### 目前实现接口

[接口功能详细说明](https://github.com/jhy625/DouSheng/blob/master/Doc/%E6%8E%A5%E5%8F%A3%E5%8A%9F%E8%83%BD%E8%AF%A6%E7%BB%86%E8%AF%B4%E6%98%8E.md)

#### 基础接口

1. 注册接口
2. 登录接口
3. 视频流接口
4. 视频投稿接口
5. 发布列表接口
6. 用户信息接口

### 扩展接口

1. 赞操作
2. 点赞操作
3. 评论操作
4. 视频评论列表
5. 关系操作
6. 用户关注列表
7. 用户粉丝列表



**[项目不同版本介绍](https://github.com/jhy625/DouSheng/blob/master/Doc/%E9%A1%B9%E7%9B%AE%E4%B8%8D%E5%90%8C%E7%89%88%E6%9C%AC%E4%BB%8B%E7%BB%8D.md)**
