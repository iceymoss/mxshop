
# mxshop电商系统
### 主要技术栈：go、grpc、gin、mysql、redis
### 功能介绍
* 登录/注册功能：采用sever和web双层架构、使用viper包做配置解析、使用redis实现注册验证码缓存服务、使用base生成验证码图片、使用MD5盐值加密保证密码只有注册者知道

* 商品服务功能：1.商品相关、2.商品品牌相关、3.商品分类类目相关、4.商品分类相关、5.商品主页轮播图相关
