## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/MasaoFujii/postgres_fdw_plus/blob/e04bf0a9948333e415e3a5ea68a0af861b19e918/README.md)
- [1.0 版本 SQL 对象](https://github.com/MasaoFujii/postgres_fdw_plus/blob/e04bf0a9948333e415e3a5ea68a0af861b19e918/postgres_fdw_plus--1.0.sql)
- [FDW 实现](https://github.com/MasaoFujii/postgres_fdw_plus/blob/e04bf0a9948333e415e3a5ea68a0af861b19e918/postgres_fdw_plus.c)

`postgres_fdw_plus` 派生自 PostgreSQL `postgres_fdw`，为外部事务增加可选两阶段提交。它创建的包装器名称仍是 `postgres_fdw`，因此核心 `postgres_fdw` 与本扩展不能在同一数据库共存。

```sql
CREATE EXTENSION postgres_fdw_plus;
CREATE SERVER remote_pg FOREIGN DATA WRAPPER postgres_fdw
  OPTIONS (host 'db.internal', dbname 'app');
SET postgres_fdw.two_phase_commit = on;
```

每台远端服务器必须支持预备事务，并为产生的预备状态负载配置容量。应监控所有参与者的 `pg_prepared_xacts`。扩展记录成功提交的本地事务，使解析函数能够决定提交还是回滚遗留远端事务。

多个关键 GUC 可由普通用户修改，包括故意只执行 prepare 的调试选项，以及可能让解析变得不确定的跟踪开关。应强制安全的角色/数据库默认值，并阻止不可信会话。只有证明跟踪完整时才能强制回滚。必须测试 prepare、本地提交、远端提交、恢复、故障转移、超时、用户映射与清理间的所有崩溃点；2PC 会无限保留锁与 WAL。只能用于上游明确测试的 PostgreSQL 16 及以后版本。
