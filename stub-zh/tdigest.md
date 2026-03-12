

## 用法

> [tdigest: PostgreSQL 的 t-digest 百分位估算](https://github.com/tvondra/tdigest)

实现 t-digest 算法，用于在线累积基于排名的统计量，如分位数和截尾均值。比 `percentile_cont` 快得多，支持并行处理，并允许预聚合。

```sql
CREATE EXTENSION tdigest;
```

### 直接聚合函数

| 函数 | 说明 |
|---|---|
| `tdigest_percentile(value, compression, quantile)` | 估算单个百分位数 |
| `tdigest_percentile(value, compression, quantiles[])` | 估算多个百分位数 |
| `tdigest_percentile_of(value, compression, value)` | 估算某个值的百分位排名 |
| `tdigest_percentile_of(value, compression, values[])` | 估算多个值的百分位排名 |

### 预聚合函数

| 函数 | 说明 |
|---|---|
| `tdigest(value, compression)` | 从数据值构建 t-digest |
| `tdigest_percentile(digest, quantile)` | 从预构建的摘要中估算百分位数 |
| `tdigest_percentile(digest, quantiles[])` | 从预构建的摘要中估算多个百分位数 |

### 增量更新函数

| 函数 | 说明 |
|---|---|
| `tdigest_add(digest, value)` | 向现有摘要中添加单个值 |
| `tdigest_add(digest, values[])` | 向现有摘要中添加一组值 |
| `tdigest_union(digest, digest)` | 合并两个摘要 |

### 工具函数

| 函数 | 说明 |
|---|---|
| `tdigest_count(digest)` | 返回摘要中的元素数量 |
| `tdigest_sum(digest, low, high)` | 指定值范围内的截尾求和 |
| `tdigest_avg(digest, low, high)` | 指定值范围内的截尾均值 |

### 参数

- `compression` -- 控制精度（值越大越精确，摘要越大）。误差大约为 `1/compression`。

### 示例

```sql
-- 替代: SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT tdigest_percentile(a, 100, 0.95) FROM t;

-- 多个百分位数
SELECT tdigest_percentile(a, 100, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- 预聚合以实现快速重复查询
CREATE TABLE p AS SELECT a, b, tdigest(c, 100) AS d FROM t GROUP BY a, b;

-- 查询预聚合数据（约 1.5ms vs 精确计算约 7s）
SELECT a, tdigest_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```
