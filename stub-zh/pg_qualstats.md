


## 用法

> [pg_qualstats: PostgreSQL 谓词统计收集器](https://github.com/powa-team/pg_qualstats)

pg_qualstats 收集 `WHERE` 子句和 `JOIN` 条件中谓词的统计信息。它追踪哪些列被最频繁查询以及哪些列被一起查询，从而支持索引建议。

### 查看谓词统计

```sql
-- 当前数据库的原始谓词统计
SELECT * FROM pg_qualstats;

-- 人类可读的聚合形式
SELECT * FROM pg_qualstats_pretty;

-- 按属性聚合的统计
SELECT * FROM pg_qualstats_all;

-- 按查询聚合的谓词
SELECT * FROM pg_qualstats_by_query;
```

### 索引顾问

基于收集的谓词统计生成索引建议：

```sql
-- 建议索引（过滤行数 >1000 且选择率 >30% 的谓词）
SELECT v FROM json_array_elements(
    pg_qualstats_index_advisor(min_filter => 50)->'indexes') v;

-- 显示无法优化的谓词
SELECT v FROM json_array_elements(
    pg_qualstats_index_advisor(min_filter => 50)->'unoptimised') v;
```

### 实用函数

```sql
-- 获取 queryid 对应的存储查询文本
SELECT pg_qualstats_example_query(queryid);

-- 获取所有存储的查询文本
SELECT * FROM pg_qualstats_example_queries();

-- 重置所有统计信息
SELECT pg_qualstats_reset();
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_qualstats.enabled` | `true` | 启用/禁用收集 |
| `pg_qualstats.track_constants` | `true` | 追踪各个常量值 |
| `pg_qualstats.max` | 1000 | 最大追踪谓词数和查询文本数 |
| `pg_qualstats.resolve_oids` | `false` | 在查询时解析 OID（占用更多空间） |
| `pg_qualstats.track_pg_catalog` | `false` | 追踪 pg_catalog 对象上的谓词 |
| `pg_qualstats.sample_rate` | -1 | 查询采样比例（-1 = 自动：1/max_connections） |
