# chat

## 数据库采用 postgresql
重启 pq 服务
`sudo service postgresql restart`
进入命令行
`psql -U chat`
编辑配置
`vim /etc/postgresql/11/main/postgresql.conf`
列出当前数据库所有表
`\dt`
列表所有数据库
`\l`
切换数据库
`\c 数据库名`
查看数据表结构
`\d 表名`

服务端启动 pgadmin web ui 容器
`docker run --env PGADMIN_DEFAULT_EMAIL=test@qq.com --env PGADMIN_DEFAULT_PASSWORD=test --env PGADMIN_LISTEN_PORT=8081 dpage/pgadmin4 -p 8081:8081`

[官方中文文档](http://www.postgres.cn/docs/10/admin.html)
[配置问题](https://stackoverflow.com/questions/18664074/getting-error-peer-authentication-failed-for-user-postgres-when-trying-to-ge)
[初始化问题](https://blog.csdn.net/zhangzeyuaaa/article/details/77941039)

[gorm 模型定义 -- tag](https://juejin.im/post/5ce2a5f3e51d455d86719f77)

## note
由于 user chan 采用 map 保存,并且键为 user id.所以需要保证每个用户只存在一次登录,此处应该在 jwt 中处理

## 1579070079
[比较两个 slice 是否相等 -- 遍历性能优于反射](https://www.jianshu.com/p/80f5f5173fca)
[比较 slice 前进行边界限制能提升性能](https://go101.org/article/bounds-check-elimination.html)

## 1580274197
需要开始接触单元测试了,不能一直采用 http 测试的方法!!!

## deploy
### 交叉编译
`GOOS=linux GOARCH=amd64 go build hello.go`

`https://www.jianshu.com/p/4b345a9e768e`

`$ docker logs -f -t CONTAINER_ID`

`https://www.jianshu.com/p/1eb1d1d3f25e`