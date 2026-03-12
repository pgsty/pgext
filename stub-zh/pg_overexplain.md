


## 用法

> [pg_overexplain: 带有内部规划器详情的扩展 EXPLAIN](https://www.postgresql.org/docs/current/pgoverexplain.html)

pg_overexplain 使用额外的调试选项扩展 `EXPLAIN` 命令，以显示内部规划器数据结构。主要用于规划器调试和开发。

### 加载

```sql
LOAD 'pg_overexplain';
-- 或在 postgresql.conf 中：
-- session_preload_libraries = 'pg_overexplain'
```

### EXPLAIN (DEBUG)

显示内部计划树信息：

```sql
EXPLAIN (DEBUG) SELECT * FROM my_table WHERE id = 1;
```

显示每个节点的字段：
- **Disabled Nodes** -- 禁用节点的原始计数
- **Parallel Safe** -- 节点是否可以出现在 Gather 下
- **Plan Node ID** -- 用于并行查询协调的内部 ID
- **extParam / allParam** -- 影响节点的参数

显示每个查询的字段：
- **Command Type** -- 查询类型（select、update 等）
- **Flags** -- hasReturning、hasModifyingCTE、canSetTag、transientPlan 等
- **Subplans Needing Rewind** -- 需要回退的子计划 ID
- **Relation OIDs** -- 计划依赖的 OID
- **Parse Location** -- 查询字符串中的位置

### EXPLAIN (RANGE_TABLE)

显示查询范围表条目的信息：

```sql
EXPLAIN (RANGE_TABLE) SELECT * FROM t1 JOIN t2 ON t1.id = t2.id;
```

显示范围表索引引用（`Scan RTI`、`Nominal RTI`、`Append RTIs` 等），并转储每个范围表条目及其类型（relation、subquery、join、cte 等）和条目特定字段。

### 注意事项

- 输出反映内部规划器数据结构，可能需要阅读源代码才能解释
- 输出格式可能因 PostgreSQL 版本而异
- 不建议在生产环境中使用
