


## 用法

> [pg_hint_plan: 赋予 PostgreSQL 手动干预执行计划的能力](https://github.com/ossc-db/pg_hint_plan)

`pg_hint_plan` 通过 SQL 注释提示覆盖 PostgreSQL 查询规划器的决策，允许你强制指定扫描方法、连接策略和连接顺序。

### 提示语法

提示嵌入在以 `/*+` 开头、`*/` 结尾的 SQL 注释中：

```sql
/*+
  HashJoin(a b)
  SeqScan(a)
*/
SELECT *
FROM pgbench_branches b
JOIN pgbench_accounts a ON b.bid = a.bid
ORDER BY a.aid;
```

### 扫描方法提示

| 提示 | 描述 |
|------|------|
| `SeqScan(table)` | 强制顺序扫描 |
| `IndexScan(table [index...])` | 强制索引扫描（可选限定特定索引） |
| `IndexOnlyScan(table [index...])` | 强制仅索引扫描 |
| `BitmapScan(table [index...])` | 强制位图扫描 |
| `TidScan(table)` | 强制 TID 扫描 |
| `NoSeqScan(table)` | 禁止顺序扫描 |
| `NoIndexScan(table)` | 禁止索引扫描和仅索引扫描 |
| `NoIndexOnlyScan(table)` | 禁止仅索引扫描 |
| `NoBitmapScan(table)` | 禁止位图扫描 |
| `NoTidScan(table)` | 禁止 TID 扫描 |
| `IndexScanRegexp(table [regexp...])` | 使用正则表达式匹配索引名的强制索引扫描 |
| `DisableIndex(table index...)` | 在规划期间禁用特定索引 |

### 连接方法提示

| 提示 | 描述 |
|------|------|
| `NestLoop(t1 t2 [t3...])` | 强制嵌套循环连接 |
| `HashJoin(t1 t2 [t3...])` | 强制哈希连接 |
| `MergeJoin(t1 t2 [t3...])` | 强制归并连接 |
| `NoNestLoop(t1 t2 [t3...])` | 禁止嵌套循环连接 |
| `NoHashJoin(t1 t2 [t3...])` | 禁止哈希连接 |
| `NoMergeJoin(t1 t2 [t3...])` | 禁止归并连接 |
| `Memoize(t1 t2 [t3...])` | 允许缓存内层结果 |
| `NoMemoize(t1 t2 [t3...])` | 禁止缓存 |

### 连接顺序提示

```sql
-- 简单的左深连接顺序
/*+ Leading(a b c) */

-- 带嵌套的显式连接树
/*+ Leading((a (b c))) */
```

### 行数修正

```sql
/*+ Rows(a b #100) */    -- 设置为绝对数量
/*+ Rows(a b +100) */    -- 增加估计值
/*+ Rows(a b -100) */    -- 减少估计值
/*+ Rows(a b *2.0) */    -- 倍乘估计值
```

### 并行查询控制

```sql
/*+ Parallel(table 4 hard) */   -- 强制 4 个并行工作进程
/*+ Parallel(table 0 hard) */   -- 禁用并行
```

### GUC 参数覆盖

```sql
/*+ Set(random_page_cost 1.0) Set(seq_page_cost 1.0) */
SELECT * FROM my_table WHERE id = 42;
```

### GUC 配置

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `pg_hint_plan.enable_hint` | 全局启用或禁用提示 | on |
| `pg_hint_plan.enable_hint_table` | 启用基于查询的提示表 | off |
| `pg_hint_plan.debug_print` | 打印已应用提示的调试信息 | off |
| `pg_hint_plan.parse_messages` | 提示解析消息的日志级别 | INFO |
| `pg_hint_plan.message_level` | 调试消息的日志级别 | LOG |
