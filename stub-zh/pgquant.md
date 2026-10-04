## 用法

来源：

- [docs/volatility.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/volatility.md)
- [docs/covariance.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/covariance.md)
- [pgquant.control](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/pgquant.control)
- [README.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/README.md)
- [Cargo.toml](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/Cargo.toml)
- [src/lib.rs](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/src/lib.rs)
- [docs/returns.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/returns.md)
- [docs/risk.md](https://github.com/pranshu05/pgQuant/blob/3c42fc3fc0f7630944b05c3a31a377835c2e1c6e/docs/risk.md)

`pgquant` 可计算简单／对数收益率、累计与年化收益率，以及历史、正态和 Student-t 分布下的 VaR/ES。0.1.5 新增波动率与协方差模型，因子构建和投资组合优化仍在计划中。

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
