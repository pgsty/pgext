## 用法

来源：[README](https://github.com/sraoss/pg_ivm/blob/master/README.md), [release 1.14](https://github.com/sraoss/pg_ivm/releases/tag/v1.14)

`pg_ivm` 为 PostgreSQL 物化视图提供 immediate Incremental View Maintenance。它不会全量重算整个视图，而是通过 `AFTER` triggers 应用 delta，并将元数据保存在 `pgivm` schema 中。

```sql
CREATE EXTENSION pg_ivm;
```

### 所需配置

上游指出应预加载 `pg_ivm`，这样 IMMV 才能被正确维护：

```conf
shared_preload_libraries = 'pg_ivm'
session_preload_libraries = 'pg_ivm'
```

当前 README 说明该扩展兼容 PostgreSQL 13 到 18；最新 GitHub release 为 `1.14`，发布日期是 2026-03-31。

### 主要函数

- `pgivm.create_immv(name, query)` 创建可增量维护物化视图（IMMV）、其维护触发器，以及在可能时创建唯一索引。
- `pgivm.refresh_immv(name, with_data)` 全量刷新 IMMV，并可禁用或重新启用维护。
- `pgivm.get_immv_def(regclass)` 重建已存储的 `SELECT` 定义。
- `pgivm.pg_ivm_immv` 保存 IMMV 元数据，包括 `immvrelid`、`viewdef`、`ispopulated` 与 `lastivmupdate`。

### 常见模式

创建 IMMV：

```sql
SELECT pgivm.create_immv(
  'immv_agg',
  'SELECT bid, count(*), sum(abalance), avg(abalance)
   FROM pgbench_accounts JOIN pgbench_branches USING(bid)
   GROUP BY bid'
);
```

在基表变更后查询已维护的结果：

```sql
UPDATE pgbench_accounts SET abalance = abalance + 1000 WHERE aid = 4112345;
SELECT * FROM immv_agg WHERE bid = 42;
```

检查或刷新 IMMV：

```sql
SELECT immvrelid AS immv, pgivm.get_immv_def(immvrelid)
FROM pgivm.pg_ivm_immv;

SELECT pgivm.refresh_immv('immv_agg', true);
```

批量操作前暂停维护，之后再重建：

```sql
SELECT pgivm.refresh_immv('myview', false);
-- bulk changes
SELECT pgivm.refresh_immv('myview', true);
```

### 注意事项

- 上游只支持受限的视图定义子集：join、`DISTINCT`、简单子查询/CTE，以及内建聚合 `count`、`sum`、`avg`、`min`、`max`。
- 不支持的构造包括 `HAVING`、window functions、`ORDER BY`、`LIMIT/OFFSET`、`UNION`/`INTERSECT`/`EXCEPT`、`DISTINCT ON` 与用户自定义聚合。
- 高效维护依赖合适的唯一索引；只有在定义允许时，`create_immv` 才会自动创建该索引。
