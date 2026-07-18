## 用法

来源：

- [上游 README](https://github.com/bitner/pg_table_range/blob/bd20fd89667c50ddb476c19434d7ae65390654ef/README.md)
- [版本与 PostgreSQL 特性矩阵](https://github.com/bitner/pg_table_range/blob/bd20fd89667c50ddb476c19434d7ae65390654ef/Cargo.toml)
- [规划器钩子与 GUC 注册](https://github.com/bitner/pg_table_range/blob/bd20fd89667c50ddb476c19434d7ae65390654ef/src/lib.rs)

pg_table_range 是面向 PostgreSQL 16–18 的实验性扩展，用于在规划期按非分区键列剪枝。其 table_range 索引访问方法保存每个分区的实际标量最小值与最大值、范围外延，或 PostGIS 几何外延。只有当摘要能够证明分区不可能匹配时，规划器钩子才会移除它。

### 基本用法

在分区表上构建摘要索引，并用代表性谓词对比执行计划：

```sql
CREATE EXTENSION pg_table_range;

CREATE INDEX events_tr
ON events USING table_range (val, created_at);

EXPLAIN (COSTS OFF)
SELECT *
FROM events
WHERE val >= 250;

SET table_range.enable_pruning = off;
EXPLAIN (COSTS OFF)
SELECT *
FROM events
WHERE val >= 250;
RESET table_range.enable_pruning;
```

该索引永远不会被选中来扫描行；它拥有供规划器使用的紧凑摘要。插入会自动拓宽摘要。删除仍然安全，但可能让摘要比实际存活数据更宽，直到 VACUUM 或 REINDEX 将其收紧。

受支持的剪枝谓词包括标量比较、BETWEEN、空值测试、常量 IN 或 ANY 列表，以及范围与 PostGIS 几何类型的重叠。不受支持的表达式会保守地保留分区。

### 注意事项

- 上游明确将项目标记为实验性且高度由 AI 驱动。采用前应在仿真生产数据上验证正确性、崩溃恢复、升级和规划器行为。
- 规划成本随分区数量线性增长。原生分区键剪枝便宜得多；只有当选择性非键谓词的执行收益足以抵消规划工作时，才应使用 pg_table_range。
- 索引创建与非键规划都会锁定所有分区。当分区数达到数千时，默认 max_locks_per_transaction 可能导致共享内存不足错误；增大它会消耗共享内存，且需要重启。
- 删除可能降低选择性，直到维护重建或收紧摘要。失效的 table_range 索引会默默禁用剪枝。
- 控制文件要求由超级用户安装。应限制表与索引 DDL 权限，并针对确切的 PostgreSQL 大版本测试 Rust/pgrx 构建。
