

## 用法

> [weighted_statistics: PostgreSQL 的加权统计函数](https://github.com/schmidni/pg_weighted_statistics)

高性能 C 扩展，提供加权统计函数，具有自动稀疏数据处理能力（当 `sum(weights) < 1.0` 时）。

```sql
CREATE EXTENSION weighted_statistics;
```

### 函数

| 函数 | 描述 |
|---|---|
| `weighted_mean(values[], weights[])` | 加权平均值 |
| `weighted_variance(values[], weights[], ddof)` | 加权方差（ddof：0=总体，1=样本） |
| `weighted_std(values[], weights[], ddof)` | 加权标准差 |
| `weighted_quantile(values[], weights[], quantiles[])` | 经验 CDF 分位数 |
| `wquantile(values[], weights[], quantiles[])` | 第 7 类（Hyndman-Fan）分位数 |
| `whdquantile(values[], weights[], quantiles[])` | Harrell-Davis 分位数 |
| `weighted_median(values[], weights[])` | 第 50 百分位数快捷方式（经验 CDF） |

### 示例

```sql
-- 加权平均值
SELECT weighted_mean(ARRAY[1.0, 2.0, 3.0], ARRAY[0.2, 0.3, 0.5]);
-- 2.3

-- 加权分位数
SELECT weighted_quantile(ARRAY[10.0, 20.0, 30.0], ARRAY[0.3, 0.4, 0.3], ARRAY[0.25, 0.5, 0.75]);
-- {15.0, 20.0, 25.0}

-- 稀疏数据（当 sum(weights) < 1.0 时自动添加隐式零值）
SELECT weighted_mean(ARRAY[10, 20], ARRAY[0.2, 0.3]);
-- 8.0（等价于 weighted_mean(ARRAY[10, 20, 0, 0], ARRAY[0.2, 0.3, 0.25, 0.25])）

-- 多种统计量
SELECT weighted_mean(vals, weights),
       weighted_std(vals, weights, 1),
       weighted_quantile(vals, weights, ARRAY[0.05, 0.95])
FROM (SELECT array_agg(val) AS vals, array_agg(weight) AS weights FROM data) t;
```
