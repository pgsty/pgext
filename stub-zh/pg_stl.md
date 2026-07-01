


## 用法

- 来源：[pg_ts_analysis README](https://github.com/nadyaloseva/pg_ts_analysis)，[SQL definitions](https://github.com/nadyaloseva/pg_ts_analysis/blob/main/pg_stl--1.0.sql)，[control file](https://github.com/nadyaloseva/pg_ts_analysis/blob/main/pg_stl.control)

`pg_stl` 为 PostgreSQL 提供时间序列分析函数：自相关、偏自相关、STL 分解和 Holt-Winters 预测。上游 README 和 SQL 定义面向 PostgreSQL 16+。

### 自相关

`acf_array(data double precision[], lags integer)` 返回 lag `1..lags` 的自相关值：

```sql
CREATE EXTENSION pg_stl;

SELECT acf_array(
  array_agg(revenue ORDER BY date)::double precision[],
  28
)
FROM daily_sales;
```

README 描述了如何把 `7`、`14`、`21` 等 lag 上的峰值作为周周期性的信号。当序列太短、`lags < 1` 或 `lags >= n` 时，该函数返回 `NULL`。

### 偏自相关

`pacf_array(data double precision[], lags integer)` 使用 Durbin-Levinson 递推返回偏自相关值：

```sql
WITH series AS (
  SELECT array_agg(value ORDER BY ts)::double precision[] AS values
  FROM measurements
)
SELECT
  unnest(acf_array(values, 20)) AS acf,
  unnest(pacf_array(values, 20)) AS pacf
FROM series;
```

当你想在扣除较短 lag 的影响后观察某个 lag 的直接关系时，可以使用 PACF。

### STL 分解

`stl_decompose` 会把序列分解为 trend、seasonal 和 residual 三个数组：

```sql
WITH data AS (
  SELECT array_agg(revenue ORDER BY month)::double precision[] AS values
  FROM monthly_revenue
),
decomposed AS (
  SELECT (stl_decompose(values, 12)).*
  FROM data
)
SELECT
  unnest(trend) AS trend,
  unnest(seasonal) AS seasonal,
  unnest(residual) AS residual
FROM decomposed;
```

SQL 定义中的函数签名为：

```sql
stl_decompose(
  y double precision[],
  period integer,
  seasonal integer DEFAULT 7,
  robust boolean DEFAULT true,
  trend integer DEFAULT 0,
  low_pass integer DEFAULT 0,
  inner_iter integer DEFAULT 2,
  outer_iter integer DEFAULT 0
) RETURNS stl_result
```

只需要单个分量时，可以使用便捷函数：

```sql
SELECT stl_trend(values, 12) FROM series;
SELECT stl_seasonal(values, 12) FROM series;
SELECT stl_residual(values, 12) FROM series;
```

### 有序聚合辅助函数

SQL 文件还定义了 `stl_collect_ordered(tbl regclass, val text, ord text)`，用于按顺序把某一列收集为 `double precision[]`：

```sql
SELECT stl_decompose(
  stl_collect_ordered('monthly_revenue'::regclass, 'revenue', 'month'),
  12
);
```

### Holt-Winters 预测

`holt_winters_predict(seasonal_type text, period_length int, start_data_array real[])` 会向前预测一个季节周期。`seasonal_type` 为 `'mult'` 时表示乘法季节性，为 `'add'` 时表示加法季节性：

```sql
SELECT *
FROM holt_winters_predict(
  'mult',
  4,
  (SELECT array_agg(revenue ORDER BY date)::real[] FROM sales)
);
```

SQL 实现会自动选择平滑系数：先进行 500 次随机初始化，再以 `0.001` 步长细化搜索，以最小化平方误差。辅助函数 `holt_winters_mse(...)` 作为预测器使用的误差计算例程一并提供。

### 注意事项

- `stl_decompose` 需要一个不含 `NULL` 的 `double precision[]`。
- README 说明序列长度至少应为 `2 * period`。
- `seasonal` 必须是不小于 `3` 的奇数。
- Holt-Winters 需要 `real[]` 输入，并且只支持 `'mult'` 和 `'add'` 两种季节类型。
