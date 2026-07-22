## 用法

来源：

- [目录源码版本的官方 README](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/README.md)
- [目录源码版本的扩展控制文件](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/greenplum_fdw.control)
- [1.0 版扩展 SQL](https://github.com/luxms/greenplum-fdw/blob/fc8f03883391b72940ff8ceb9dc0aa56644a3c67/greenplum_fdw--1.0.sql)
- [PostgreSQL postgres_fdw 文档](https://www.postgresql.org/docs/current/postgres-fdw.html)

greenplum_fdw 是 PostgreSQL 外部数据包装器的一个分支，通过 libpq 查询远端 Greenplum Database。目录记录的源码会把远端事务改为可串行化隔离级别，因为较早的 Greenplum 版本会拒绝相应上游包装器启动的可重复读事务。

### 核心流程

创建包装器，描述远端 Greenplum 端点，把本地角色映射到远端凭据，然后定义或导入外部表。

```sql
CREATE EXTENSION greenplum_fdw;

CREATE SERVER gp_sales
FOREIGN DATA WRAPPER greenplum_fdw
OPTIONS (host 'greenplum.example.net', port '5432', dbname 'warehouse');

CREATE USER MAPPING FOR app_user
SERVER gp_sales
OPTIONS (user 'reporter', password 'secret');

CREATE FOREIGN TABLE gp_orders (
    order_id bigint,
    ordered_at timestamp,
    total numeric
)
SERVER gp_sales
OPTIONS (schema_name 'public', table_name 'orders');

SELECT order_id, total
FROM gp_orders
WHERE ordered_at >= current_date - interval '7 days';
```

### 重要选项

libpq 接受的连接关键字属于外部服务器，凭据属于用户映射。表和列映射可以覆盖远端名称。该包装器源码还识别从其 postgres_fdw 代码基础继承的远端估算、启动成本、元组成本、抓取大小、可更新性和可下推扩展设置。

### 兼容性与注意事项

应使用与本地 PostgreSQL 主版本完全匹配的 greenplum_fdw 源码分支。上游 README 描述了按 PostgreSQL 主版本拆分的分支以及较早的 Greenplum 兼容问题；不要假定目录中的 1.0 源码拥有当前 postgres_fdw 的功能集合或版本范围。

该包装器的关键行为是改变远端事务隔离级别。需要针对目标 Greenplum 版本和实际工作负载验证，尤其是写入与多语句事务。远端权限、网络可达性、TLS 参数和凭据处理仍属于普通 libpq 与外部服务器的职责。扩展没有声明预加载或服务器重启要求。
