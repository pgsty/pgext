## 用法

来源：

- [官方 README](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/README.md)
- [扩展控制文件](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/pg_statkit.control)
- [pgrx 清单](https://github.com/kirdmi/pg_statkit/blob/v1.1.0/Cargo.toml)

`pg_statkit` 在 PostgreSQL 数组上提供描述统计与临床效应量函数，无需把源数据导出到独立分析进程。

### 启用

1.1.0 版本通过 pgrx 0.18.1 构建并测试 PostgreSQL 13–17。安装匹配的原生构件后，以超级用户创建不可迁移、非 trusted 的扩展：

```sql
CREATE EXTENSION pg_statkit;
```

它不要求预加载或重启。

### 描述统计

函数接收 `double precision[]`；应显式转换数组字面量，因为小数字面量默认会形成 `numeric[]`。

```sql
SELECT statkit_mean('{1,2,3,4,5}'::float8[]);
SELECT statkit_median('{1,2,3,4,5}'::float8[]);
SELECT statkit_iqr('{1,2,3,4,5}'::float8[]);
SELECT statkit_percentile('{1,2,3,4,5}'::float8[], 40);

SELECT patient_id,
       statkit_median(array_agg(value::float8)) AS median_value
FROM measurements
GROUP BY patient_id;
```

函数面还包括 variance、population/sample standard deviation、MAD、coefficient of variation 与 standard error。

### 效应量与解释

`statkit_risk_difference`、`statkit_risk_ratio` 与 `statkit_odds_ratio` 用四个 `bigint` count 表示 2×2 表。`statkit_cohens_d`、`statkit_hedges_g` 与 `statkit_glass_delta` 接收两组 sample。

```sql
SELECT statkit_odds_ratio(a => 8, b => 2, c => 2, d => 8);

SELECT statkit_cohens_d(
  array_agg(value::float8) FILTER (WHERE arm = 'treatment'),
  array_agg(value::float8) FILTER (WHERE arm = 'control')
)
FROM trial;
```

若任一单元格为零，odds-ratio function 会应用 0.5 修正。未定义输入返回 `NULL`。1.1.0 不提供 confidence interval、hypothesis test 或 assumption check；研究结论需要这些能力时，应使用完整统计软件包。

