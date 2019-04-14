# Beego Example 
![Travis](https://img.shields.io/badge/-Beijing--研发-green.svg?logo=Docker&style=popout-square)
<h2>Go语言项目</h2>

使用beego框架开发的一个例子。

## 如何安装

```
$ git clone http://116.196.78.233:10080/wushaoqiang/beego.git
```
## 如何运行
### 环境依赖
- Mysql
- Redis
- Kafka

### app.conf
```
[mysql]
db_alias = "default"
db_name = "beego"
db_user = "root"
db_pwd = "123456"
db_host = "localhost"
db_port = 3306
db_charset = "utf8"

[cache]
#redis
redis_host = "127.0.0.1:6379"
redis_password = "659309"

[logs]
#添加输出引擎
level = 7

[kafka]
servers = "localhost:9092"
```