

## 用法

> [quantile: PostgreSQL 的分位数聚合函数](https://github.com/tvondra/quantile)

提供计算分位数的聚合函数，支持 `int`、`bigint`、`double precision` 和 `numeric` 类型的重载。

```sql
CREATE EXTENSION quantile;
```

### 函数

| 函数 | 描述 |
|---|---|
| `quantile(value, quantile float)` | 计算单个分位数（0 到 1） |
| `quantile(value, quantiles float[])` | 一次计算多个分位数，返回数组 |

### 示例

```sql
-- 计算中位数（0.5 分位数）
SELECT quantile(i, 0.5) FROM generate_series(1, 1000) s(i);
-- 500

-- 计算第 95 百分位数
SELECT quantile(i, 0.95) FROM generate_series(1, 1000) s(i);

-- 一次计算所有四分位数（比分别调用更高效）
SELECT quantile(i, ARRAY[0.25, 0.5, 0.75])
FROM generate_series(1, 1000) s(i);
-- {250, 500, 750}
```

注意：自 PostgreSQL 9.4 起，内置的 `percentile_cont` 和 `percentile_disc` 函数已经可用。建议优先使用内置函数，仅在该扩展对您的工作负载有明显性能优势时才使用它。
