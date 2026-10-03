## 用法

来源：

- [Official README](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/README.md)
- [Extension control file](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/redis_cluster_fdw.control)
- [Installation SQL](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/redis_cluster_fdw--1.0.sql)
- [Implementation (redis_cluster_fdw.c)](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/redis_cluster_fdw.c)

`redis_cluster_fdw` 是面向 Redis 集群的独立 Redis FDW 变体。包装器名称和选项以版本化 SQL 与 C 校验器为准，README 的部分示例仍沿用旧的非集群名称。

### 基本用法

安装扩展和 hiredis-cluster 库后，配置集群节点与逐用户密码。该集群校验器不接受数据库编号选项。

```sql
CREATE EXTENSION redis_cluster_fdw;
CREATE SERVER redis_cluster FOREIGN DATA WRAPPER redis_cluster_fdw
  OPTIONS (nodes '127.0.0.1:6379');
CREATE USER MAPPING FOR CURRENT_USER SERVER redis_cluster
  OPTIONS (password 'replace-with-cluster-password');
CREATE FOREIGN TABLE redis_values (key text, value text)
  SERVER redis_cluster;
SELECT * FROM redis_values;
```

### 选项与限制

`nodes` 为逗号分隔的服务器列表，`password` 放在用户映射中。表选项包括 `tabletype`、`tablekeyprefix`、`tablekeyset` 和 `singleton_key`。单键列表、集合及有序集合分值需要使用相应列布局。读写直接作用于 Redis，不能假设 PostgreSQL 回滚会撤销外部写入。不支持模式导入、截断 API 或通用表达式下推。上游声称支持 PostgreSQL 13+，但未提供完整的大版本测试矩阵；应使用匹配的构建和 UTF-8 数据库编码。
