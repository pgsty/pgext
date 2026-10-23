## 用法

来源：

- [README.md](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/README.md)
- [trino_fdw--1.0.sql](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/trino_fdw--1.0.sql)
- [Makefile](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/Makefile)
- [trino_fdw.control](https://github.com/cmsedore/pgtrino/blob/acf9f98087b14ed7c93a380b761190d2705554d1/trino_fdw.control)

`trino_fdw` 使用 libcurl 将 Trino 表映射为 PostgreSQL 外部表。上游说明支持 PostgreSQL 14–17。先由超级用户安装扩展，再配置外部服务器与调用者映射。无需共享预加载。

### 核心用法

```sql
CREATE EXTENSION trino_fdw;
CREATE SERVER trino FOREIGN DATA WRAPPER trino_fdw
  OPTIONS (host 'trino.example.com', port '8443', ssl 'true', catalog 'hive');
CREATE USER MAPPING FOR CURRENT_USER SERVER trino OPTIONS (user 'analytics');
CREATE SCHEMA sales;
IMPORT FOREIGN SCHEMA sales FROM SERVER trino INTO sales;
EXPLAIN SELECT * FROM sales.orders WHERE order_id = 42;
```

### 运行边界

服务器选项包括 host、port、ssl、ssl_verify、catalog、schema、spooling 和 batch_size。传输远程凭证时应使用 HTTPS 并校验证书。映射选项 user 与 password 指定 Trino 身份，PostgreSQL 外部表权限和 Trino 授权共同生效。投影、符合条件的过滤、连接、聚合与排序可下推；EXPLAIN 显示远程 SQL，ANALYZE 收集供本地规划器使用的统计信息。`trino_query` 返回记录，`trino_execute` 返回更新计数，两者都要求外部服务器的 USAGE 权限。UPDATE/DELETE 要求整个语句可下推且连接器支持；不支持 RETURNING、带连接的修改或行触发器。每条远程语句或每批 INSERT 都独立提交，PostgreSQL 回滚无法撤销。启用 spooling 时，PostgreSQL 服务器还须能够访问 Trino 使用的对象存储；该路径不可达时应关闭 spooling。
