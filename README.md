# Gin Example 
![Travis](https://img.shields.io/badge/-Beijing--研发-green.svg?logo=Docker&style=popout-square)

<h2>Go语言项目</h2>

使用Gin框架开发的一个例子，使用GORM，INI，Swagger，Redis，Kafka，ElaticSearch等库。

<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="159" hegiht="159" align="center" />


## 如何安装

```
$ git clone http://116.196.78.233:10080/wushaoqiang/Gin_demo.git
```
## 如何运行
### 环境依赖
- Mysql
- Redis
- Kafka

### app.ini
```
[server]
App = gin
Port = 8090
RunMode = debug
ReadTimeOut = 60
WriteTimeOut = 60
JwtSecret = love

[mysql]
user = root
password = 123456
host = 127.0.0.1:3306
db = bloc
```

### swagger

```
http://localhost:port/swagger/index.html
```

![avatar](/static/this.jpg)