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

[官方中文文档](http://www.postgres.cn/docs/10/admin.html)
[配置问题](https://stackoverflow.com/questions/18664074/getting-error-peer-authentication-failed-for-user-postgres-when-trying-to-ge)
[初始化问题](https://blog.csdn.net/zhangzeyuaaa/article/details/77941039)