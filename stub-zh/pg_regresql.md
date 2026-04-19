## 用法

来源：[extension README](https://github.com/boringSQL/regresql/blob/master/pg_ext/README.md)，[control file](https://github.com/boringSQL/regresql/blob/master/pg_ext/pg_regresql.control)，[portable stats article](https://boringsql.com/posts/portable-stats/)

`pg_regresql` 是 RegreSQL 的 PostgreSQL planner-hook 组件。它让规划器优先信任 `pg_class` 中的统计信息，而不是根据物理文件大小重新缩放估算值，因此适合在注入生产统计后做执行计划回归测试。

### 激活 hook

```sql
LOAD 'pg_regresql';

EXPLAIN SELECT ...;
```

如果希望整个测试实例都启用它，上游建议：

```conf
session_preload_libraries = 'pg_regresql'
```

### 它覆盖的估算

README 说明该 hook 会在 `estimate_rel_size()` 之后运行，并用 catalog 中的值替换规划器默认估算：

- `rel->pages` 来自 `pg_class.relpages`
- `rel->tuples` 来自 `pg_class.reltuples`
- `rel->allvisfrac` 来自 `pg_class.relallvisible / relpages`
- `IndexOptInfo->pages` 来自索引的 `pg_class.relpages`
- `IndexOptInfo->tuples` 来自索引的 `pg_class.reltuples`

### 典型工作流

```sql
SELECT pg_restore_relation_stats(
    'schemaname', 'public',
    'relname', 'test_orders',
    'relpages', 123513::integer,
    'reltuples', 50000000::real,
    'relallvisible', 123513::integer
);

LOAD 'pg_regresql';

EXPLAIN SELECT * FROM test_orders WHERE created_at > '2024-06-01';
```

这个流程主要用于在本地复现生产计划，或在 migration、升级和 plan-regression 测试中复用已恢复的生产统计。

### 说明

- 控制文件当前声明 `default_version = '2.0'`。
- 公开仓库里的 tag 仍主要是 `v2.0.0-alpha*`，因此打包目标 `2.0.0` 领先于 GitHub 上清晰标注的最终正式 tag。
- 上游文档说明该扩展支持 PostgreSQL 14+。
