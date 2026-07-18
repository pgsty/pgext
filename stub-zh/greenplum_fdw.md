## 用法

来源：

- [上游 README](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/README.md)
- [扩展 SQL](https://github.com/luxms/greenplum-fdw/blob/f643f9193fa43c672b2f8fec1fb9e1c90900d1be/greenplum_fdw--1.0.sql)
- [PostgreSQL postgres_fdw 文档](https://www.postgresql.org/docs/current/postgres-fdw.html)

`greenplum_fdw` 是 PostgreSQL `postgres_fdw` 面向 Greenplum 连接的分支。它将远端事务隔离级别从 `REPEATABLE READ` 改为 `SERIALIZABLE`，以解决文档所述 Greenplum 5 拒绝前一种隔离级别的问题。

安装服务器端动态库后，按常规 `postgres_fdw` 对象模型配置：

```sql
CREATE EXTENSION greenplum_fdw;

CREATE SERVER gp_remote
  FOREIGN DATA WRAPPER greenplum_fdw
  OPTIONS (host 'greenplum.example.net', port '5432', dbname 'analytics');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER gp_remote
  OPTIONS (user 'gp_reader', password 'change-me');

CREATE FOREIGN TABLE gp_sales (
  sale_id bigint,
  amount numeric
) SERVER gp_remote OPTIONS (schema_name 'public', table_name 'sales');
```

### 兼容性

上游版本 `1.0` 表示已发布分支面向 PostgreSQL 12 和 13，其余用法沿用 `postgres_fdw` 文档。该仓库没有发布当前 Greenplum/PostgreSQL 兼容矩阵，因此用于更新版本前必须进行构建与事务测试。应保护映射凭据，并仅向预期角色授予外部服务器访问权。
