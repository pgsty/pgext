

## 用法

> [lower_quantile: PostgreSQL 的下分位数聚合函数](https://github.com/tvondra/lower_quantile)

实现了"下分位数"聚合，与 `percentile_disc` 略有不同，它返回在排序多重集中秩为 `floor(1 + q*(n-1))` 的值。

```sql
CREATE EXTENSION lower_quantile;
```

### 函数

| 函数 | 描述 |
|---|---|
| `lower_quantile(value, quantile float)` | 计算给定分位数值（0 到 1）的下分位数 |

### 示例

```sql
-- 计算下分位数中位数
SELECT lower_quantile(i, 0.5)
FROM generate_series(1, 1000) s(i);

-- 计算第 95 百分位下分位数
SELECT lower_quantile(i, 0.95)
FROM generate_series(1, 1000) s(i);
```

该定义被部分论文（例如 DDSketch 论文）用于表述精度保证。
