## 用法

来源：

- [官方来源](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/README.md)
- [官方来源](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/pg_governor.control)
- [官方来源](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/sql/pg_governor--0.1.sql)
- [官方来源](https://github.com/peterrobert401/pg_workload/blob/c006fa19836bdf0d2994895a9245a7f0fcfd8e7b/LICENSE)

`pg_governor` 0.1 为直接分配到角色组的用户限制顶层查询并发与连接数，支持 PostgreSQL 15–17。准入优先级控制排队顺序，不分配 CPU 或 I/O 资源。

### 配置工作负载组

预加载库并重启，在 `pg_governor.database` 指定的数据库中创建扩展；示例假定角色 `alice` 已存在。

```ini
shared_preload_libraries = 'pg_governor'
pg_governor.database = 'postgres'
```

```sql
CREATE EXTENSION pg_governor;
SELECT governor.create_group('reporting', 2, 5000, 65536, 20, 3);
SELECT governor.assign_role('alice', 'reporting');
SELECT * FROM governor.stats;
SELECT * FROM governor.instance_status();
```

### 限制与管理

示例允许两个活动查询、5000 毫秒等待、65536 KB 的 `work_mem`、20 个连接，优先级为 3。优先级越大越先执行，同级按 FIFO 排队。活动查询和连接上限为零表示无限制，等待值为零表示不超时。`pg_governor.max_active_total` 与 `pg_governor.max_connections_total` 为受管理角色增加实例级上限。

`governor.alter_group`、`governor.assign_role`、`governor.unassign_role`、`governor.drop_group` 与 `governor.reload` 用于变更配置。管理函数已撤销 `PUBLIC` 执行权限，统计信息可公开读取。后台进程定期保存累计计数，崩溃可能丢失最近一个保存周期的数据。

### 运行边界

`governor.reload()` 非事务性地修改共享状态。回滚配置变更后应再次调用，使共享内存与表一致。仅管理直接分配的角色，组关系不会经角色成员身份继承。嵌套查询及游标打开期间的查询不独立计数。共享队列最多容纳 1024 个等待查询。移除扩展前先删除组：`DROP EXTENSION` 不会清空共享状态。这不是 CPU、内存或 I/O 隔离机制。
