
# mxshop电商系统
### 主要技术栈：Go、Grpc、Gin、Mysql、Redis、Elasticsearch、RocketMQ、Nacos、Consul、Jaeger、Sentinel
### Consul的安装(docker)
```shell
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp  consul consul agent  -dev -client=0.0.0.0
```
开机启动consul
```shell
docker container update --restart=always 容器名字
```
浏览器访问 127.0.0.1:8500

### Nacos的安装(docker)
```shell
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
```
访问：http://192.168.1.103:8848/nacos/index.html
 用户名/密码：nacos/nacos
 
 配置开机启动：
 ```shell
 docker container update --restart=always xxx
```

### Elasticsearch的安装(docker)
[Elasticsearch安装](https://learnku.com/articles/72845)

### 项目介绍
网站地址：http://iceymoss.toop:8089
* 基于JWT做访问鉴权token，Gin做路由分发、表单验证、解决跨域等。

* 登录/注册功能：采用sever和web双层架构、使用viper包做配置解析、web层基于Gin做路由转发、使用redis实现注册验证码缓存服务、使用base64生成验证码图片做登录验证、srv层使用MD5盐值加密保证密码注册者知道的唯一性。

* 商品服务功能：基于Elasticsearch实现商品搜索；完成如下接口：1.商品相关、2.商品品牌相关、3.商品分类类目相关、4.商品分类相关、5.商品主页轮播图相关。
* 图片文件使用aliyun对象存储，使用服务端签名直传文件。
* 库存服务：库存服务的核心在于保持数据的一致性，可用性，高性能，解决在分布式高并发场景下，如何保证数据一致性，库存服务引入了Redis锁和RocketMQ，来实现分布式高并发场景下的数据一致性，如何扣减库存，库存超时归还，重复归还商品问题以及接口需要幂等性。
* 订单服务：基于grpc实现订单相关服务及购物车相关服务等各类接口，使用本地mysql事务保证本地数据一致性，从使用rocketMQ从订单服务到查询商品服务(跨服务)，调用库存服务扣减库存(跨服务)的跨微服务调用，保证信息一致性。
* 用户接口服务： 为用户提供操作接口其中实现了简单的地址，留言, 收藏等。
* 基于Jaeger做微服务间链路追踪，使用Sentinel实现限流。
