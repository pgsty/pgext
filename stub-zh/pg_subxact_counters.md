## 用法

来源：

- [官方 C 子项目 README](https://github.com/bdrouvot/pg_subxact_counters/blob/main/c/README.md)
- [C 扩展 SQL](https://github.com/bdrouvot/pg_subxact_counters/blob/main/c/pg_subxact_counters--1.0.sql)
- [C 计数器实现](https://github.com/bdrouvot/pg_subxact_counters/blob/main/c/pg_subxact_counters.c)

`pg_subxact_counters` 1.0 是仅可预加载的 C 扩展，暴露四个实例级原子计数器来记录子事务活动。定期采样可测量速率并发现子事务缓存溢出；它不会保留历史，也无法定位负责的数据库、会话或语句。

### 预加载与设置

把库加入预加载列表，重启 PostgreSQL，再在监控数据库创建 SQL 对象：

```conf
shared_preload_libraries = 'pg_subxact_counters'
```

```sql
CREATE EXTENSION pg_subxact_counters;

SELECT * FROM pg_subxact_counters;
```

查询函数明确要求超级用户，因此视图也有相同要求。应通过严格受控的监控路径导出样本，而不是开放普通访问。

### 计数器索引

- `subxact_start` 统计已启动的子事务。
- `subxact_commit` 统计已提交的子事务。
- `subxact_abort` 统计已中止的子事务。
- `subxact_overflow` 统计顶层事务的缓存子事务 ID 发生溢出的转换次数。

溢出计数器并不表示超出缓存上限的子事务数量。

### 解读与限制

计数器跨所有数据库全局累计，从共享内存初始化持续到 postmaster 重启，没有 SQL 重置接口。启动数减去提交数和中止数，可以粗略提示当前仍打开的子事务；但各原子值分别读取且并发活动持续发生，因此一次查询不是事务一致快照。应监控增量与速率，并和工作负载遥测关联，同时预期重启会重置序列。本文只覆盖仓库中的 C 子项目；上游报告测试过 PostgreSQL 13 及以上版本。
