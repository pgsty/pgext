## 用法

来源：

- [docs/factors.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/factors.md)
- [docs/volatility.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/volatility.md)
- [docs/covariance.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/covariance.md)
- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/1973ed494fa0c1c59d7a145b342144d20010832d/docs/risk.md)

`pgquant` 计算简单/对数收益率、累积与年化收益率，以及历史、Gaussian 和 Student-t VaR/ES。0.2.0 还提供波动率、协方差和因子构建；动量与投资组合优化仍属于计划功能。

### 核心用法

```sql
CREATE EXTENSION pgquant;
SELECT pgquant_cumulative_simple_return(ARRAY[0.02,-0.01,0.03]::float8[]);
SELECT pgquant_var_historical(ARRAY[-0.03,0.01,-0.02,0.02]::float8[], 0.95);
SELECT pgquant_es_t(0.001, 0.015, 5.0, 0.95);
```

### 运行边界

上游文档列出 PostgreSQL 14–17。控制文件允许非超级用户安装，但未将扩展标记为可信；没有预加载要求。基于查询的函数要求按顺序返回标的、日期和双精度价格三列，每个标的的首条记录不产生收益率。查询文本应来自可信来源。置信水平、分布假设和输入校验会影响结果；Student-t ES 要求自由度大于一。

### 波动率与协方差

`pgquant_rolling_vol` 使用有序的收益率查询，`pgquant_ewma_vol` 返回数组，`pgquant_ewma_lambda_mle` 估计衰减参数。`pgquant_garch11_normal` 与 `pgquant_garch11_t` 拟合 GARCH 参数，至少需要十个观测值。`pgquant_sample_cov` 和 `pgquant_shrinkage_cov_lw` 接收包含标的、日期、双精度收益率三列的查询，基于共同日期的完整数据返回标的对协方差。输入顺序、缺失数据和分布假设都会影响估计。

```sql
SELECT pgquant_ewma_vol(ARRAY[0.02,-0.01,0.03,-0.02]::float8[], 0.94);
```

### 因子构建

`pgquant_portfolio_sort(query, num_buckets)` 要求依次返回 symbol TEXT、date DATE 和 characteristic_value DOUBLE PRECISION 三列，输出 symbol、date 和 bucket，bucket 1 包含最小值。`pgquant_construct_smb(query)` 与 `pgquant_construct_hml(query)` 接收 symbol、date、return DOUBLE PRECISION、size_bucket INT、btm_bucket INT 和 weight DOUBLE PRECISION。规模分桶使用 1/2，账面市值比分桶使用 1/2/3；结果为 date 及 smb_return 或 hml_return。须准备类型正确、时间对齐的面板，且只传入可信 SQL 文本。

```sql
SELECT * FROM pgquant_portfolio_sort(
  $$ SELECT * FROM (VALUES
     ('A'::text, DATE '2026-01-01', 100.0::float8),
     ('B'::text, DATE '2026-01-01', 200.0::float8)
  ) AS panel(symbol, date, characteristic_value) $$, 2);
```
