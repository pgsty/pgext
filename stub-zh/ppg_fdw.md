## 用法

来源：

- [上游 README](https://github.com/scarbrofair/ppg_fdw/blob/7f8d675da440400ad73d6cfbc397cb1e57a4e4ee/README.md)
- [1.0 版安装 SQL](https://github.com/scarbrofair/ppg_fdw/blob/7f8d675da440400ad73d6cfbc397cb1e57a4e4ee/ppg_fdw--1.0.sql)
- [分布式规划器钩子](https://github.com/scarbrofair/ppg_fdw/blob/7f8d675da440400ad73d6cfbc397cb1e57a4e4ee/ppg_planner.c)

ppg_fdw 是一个用于在水平分片 PostgreSQL 服务器上执行并行查询的遗留原型。它将 postgres_fdw 分支与全局规划器钩子结合，并曾用 TPC-H 事实表分片和复制的维度表演示无共享布局。

### 最小 FDW 对象

该扩展会安装 ppg_fdw 包装器、处理函数和验证器。单服务器定义使用熟悉的外部数据包装器对象：

```sql
CREATE EXTENSION ppg_fdw;

CREATE SERVER shard_group
FOREIGN DATA WRAPPER ppg_fdw
OPTIONS (
    host '127.0.0.1',
    dbname 'analytics',
    port '5432'
);

CREATE USER MAPPING FOR CURRENT_USER
SERVER shard_group
OPTIONS (
    user 'remote_user',
    password 'replace-with-managed-secret'
);

CREATE FOREIGN TABLE remote_orders (
    order_id bigint,
    customer_id bigint,
    total numeric(15,2)
)
SERVER shard_group;

SELECT * FROM remote_orders LIMIT 10;
```

用户映射凭据应通过密密管理流程处理，并强制使用加密且经过认证的网络连接。上游多主机示例通过直接更新 pg_catalog，将服务器系统目录选项数组替换为自定义 host1 与 host2 条目；这种不受支持的变更在此有意省略。

### 注意事项

- 代码与安装说明来自 PostgreSQL 9.3 时代，并复制了 PostgreSQL 规划器和 FDW 内部实现。没有证据表明它与当前 PostgreSQL 版本兼容。
- 上游分布式配置需要直接更新系统目录，会绕过 DDL 验证、依赖跟踪与受支持的管理接口。不得在生产集群上使用该方法。
- 全局规划器钩子只处理一部分计划节点和查询形状。不受支持的计划可能报错、产生意外行为，或丧失预期并行性。
- 正确结果取决于应用的分片键、数据放置、重复维度表与跨分片事务假设。该扩展不提供完整的分布式数据库控制面。
- 应将 ppg_fdw 视为历史研究代码。真实部署应评估有维护的 FDW 或分布式 PostgreSQL 系统，并要求它们具备有文档的故障处理、安全、兼容性与升级路径。
