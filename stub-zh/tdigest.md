## 用法

来源：

- [tdigest v1.4.4 README](https://github.com/tvondra/tdigest/blob/v1.4.4/README.md)
- [v1.4.4 release](https://github.com/tvondra/tdigest/releases/tag/v1.4.4)
- [Extension control file](https://github.com/tvondra/tdigest/blob/v1.4.4/tdigest.control)

`tdigest` 实现了近似可合并的 t-digest，用于在线排名统计，如分位数、百分位秩和截尾均值。它支持并行聚合，并可以存储预聚合的摘要以便后续汇总。

```sql
CREATE EXTENSION tdigest;
```

### 直接聚合函数

| 函数 | 描述 |
|---|---|
| `tdigest_percentile(value, compression, quantile)` | 估算单一百分位数 |
| `tdigest_percentile(value, compression, quantiles[])` | 估算多个百分位数 |
| `tdigest_percentile_of(value, compression, value)` | 估算值的百分位秩 |
| `tdigest_percentile_of(value, compression, values[])` | 估算多个值的百分位秩 |

### 预聚合函数

| 函数 | 描述 |
|---|---|
| `tdigest(value, compression)` | 从值构建 t-digest |
| `tdigest_percentile(digest, quantile)` | 从预构建的摘要估算百分位数 |
| `tdigest_percentile(digest, quantiles[])` | 从预构建的摘要估算多个百分位数 |

### 增量更新函数

| 函数 | 描述 |
|---|---|
| `tdigest_add(digest, value)` | 将单个值添加到现有摘要中 |
| `tdigest_add(digest, values[])` | 将数组中的值添加到现有摘要中 |
| `tdigest_union(digest, digest)` | 合并两个摘要 |

### 辅助函数

| 函数 | 描述 |
|---|---|
| `tdigest_count(digest)` | 返回摘要中的项数 |
| `tdigest_sum(digest, low, high)` | 在值范围内截尾求和 |
| `tdigest_avg(digest, low, high)` | 在值范围内截尾平均 |

### 参数

- `compression` -- 控制准确性（越高 = 越准确，摘要越大）。误差大致为 `1/compression`。

### 示例

```sql
-- Instead of: SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT tdigest_percentile(a, 100, 0.95) FROM t;

-- Multiple percentiles
SELECT tdigest_percentile(a, 100, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- Pre-aggregate for fast repeated queries
CREATE TABLE p AS SELECT a, b, tdigest(c, 100) AS d FROM t GROUP BY a, b;

-- Query pre-aggregated data (~1.5ms vs ~7s for exact)
SELECT a, tdigest_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```

### 注意事项

- 结果是估算的。在设置精度目标之前，请使用代表性的数据验证所选压缩参数与精确 `percentile_cont` 结果。
- 较高的压缩通常会提高尾部准确性，但会增加状态大小和 CPU 成本。
- 存储的摘要可以在不同组和时间窗口之间合并。版本 `1.4.4` 修复了使用不同参数创建的摘要结合的问题，因此在异构状态可能相遇时，请使用该补丁级别。
- 版本 `1.4.4` 还增强了文本输入解析和验证，并增加了 PostgreSQL 19 的构建/测试覆盖率；以前被较旧版本接受的格式错误的序列化摘要现在可能会被拒绝。
