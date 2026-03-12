

## 用法

> [ddsketch: PostgreSQL 的 DDSketch 分位数估计](https://github.com/tvondra/ddsketch)

实现了 DDSketch，一种完全可合并的分位数草图，提供相对误差保证。比 `percentile_cont` 快得多，并支持并行处理。

```sql
CREATE EXTENSION ddsketch;
```

### 直接聚合函数

| 函数 | 描述 |
|---|---|
| `ddsketch_percentile(value, alpha, nbuckets, quantile)` | 估计单个百分位数 |
| `ddsketch_percentile(value, alpha, nbuckets, quantiles[])` | 估计多个百分位数 |
| `ddsketch_percentile_of(value, alpha, nbuckets, value)` | 估计某个值的百分位排名 |
| `ddsketch_percentile_of(value, alpha, nbuckets, values[])` | 估计多个值的百分位排名 |

### 预聚合函数

| 函数 | 描述 |
|---|---|
| `ddsketch(value, alpha, nbuckets)` | 从值构建 ddsketch |
| `ddsketch_percentile(sketch, quantile)` | 从预构建的草图估计百分位数 |
| `ddsketch_percentile(sketch, quantiles[])` | 从预构建的草图估计多个百分位数 |

### 工具函数

| 函数 | 描述 |
|---|---|
| `ddsketch_count(sketch)` | 返回草图中的项目数量 |
| `ddsketch_sum(sketch, low, high)` | 指定值范围内的截断求和 |
| `ddsketch_avg(sketch, low, high)` | 指定值范围内的截断平均值 |

### 参数

- `alpha` -- 控制精度和草图大小（越小越精确，但体积越大）
- `nbuckets` -- 最大桶数（每个 8 字节）

### 示例

```sql
-- 替代：SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT ddsketch_percentile(a, 0.05, 1024, 0.95) FROM t;

-- 一次估计多个百分位数
SELECT ddsketch_percentile(a, 0.05, 1024, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- 预聚合以加速重复查询
CREATE TABLE p AS SELECT a, b, ddsketch(c, 0.05, 1024) AS d FROM t GROUP BY a, b;

-- 查询预聚合数据（约 1.5ms vs 精确计算约 7s）
SELECT a, ddsketch_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```
