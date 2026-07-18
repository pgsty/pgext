## 用法

来源：

- [上游 README](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/README.md)
- [扩展 control 文件](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/redis_link.control)
- [SQL 安装脚本](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/redis_link--1.0.sql)
- [C 实现](https://github.com/fkfk000/redis_link/blob/5a92f6ff642ff86c1b6b84c8b8e343fdd3d9322d/redis_link.c)

`redis_link` `1.0` 版是上游在 PostgreSQL 16 测试的 Redis 外部数据包装器。它可以扫描 Redis string、hash 和 set 数据，处理 Redis Cluster 重定向，并支持 `INSERT`；没有实现 `UPDATE` 或 `DELETE` 回调。

### 示例

```sql
CREATE EXTENSION redis_link;
CREATE SERVER redis_server FOREIGN DATA WRAPPER redis_link
  OPTIONS (host_port '127.0.0.1:6379');
CREATE FOREIGN TABLE redis_items (name text, content text, expire integer)
  SERVER redis_server
  OPTIONS (key 'name', value 'content', prefix 'demo:', table_type 'string', ttl 'expire');
SELECT * FROM redis_items;
```

验证器接受 `password` 选项，但复核的 C 实现从未读取它，因此不能假定支持 Redis 认证。实现通过拼接标识符和行值构造 Redis 命令，再交给 hiredis 命令格式化；空格、百分号或不可信数据可能改变解析或触发内存安全缺陷。只能连接严格隔离、内容可信的 Redis 命名空间，不要向应用角色开放写入或选项创建权限。
