

## 用法

> [pg_math: 使用 GSL 的 PostgreSQL 统计分布函数](https://github.com/chanukyasds/pg_math)

基于 GNU 科学计算库（GSL）提供 31 种统计分布的 CDF（累积分布函数）和 RDF（随机分布函数）。

```sql
CREATE EXTENSION pg_math;
```

### 支持的分布

高斯分布、单位高斯分布、高斯尾分布、二元高斯分布、F 分布、指数分布、拉普拉斯分布、指数幂分布、柯西分布、瑞利分布、瑞利尾分布、朗道分布、伽马分布、均匀分布、对数正态分布、卡方分布、t 分布、贝塔分布、逻辑斯蒂分布、帕累托分布、韦伯分布、第一类 Gumbel 分布、第二类 Gumbel 分布、泊松分布、伯努利分布、二项分布、负二项分布、帕斯卡分布、几何分布、超几何分布、对数分布。

### 函数命名约定

- `rdf_<distribution>(...)` -- 随机分布函数（概率密度值）
- `cdf_<distribution>_p(...)` -- 累积分布 P 值
- `cdf_<distribution>_q(...)` -- 累积分布 Q 值（1-P）
- `cdf_<distribution>_pinv(...)` -- 逆 CDF P
- `cdf_<distribution>_qinv(...)` -- 逆 CDF Q

### 示例

```sql
-- 高斯分布
SELECT rdf_gaussian(1.5, 2.0);            -- x=1.5, sigma=2.0 处的概率密度
SELECT cdf_gaussian_p(1.5, 2.0);          -- CDF P 值

-- 单位高斯分布（标准正态）
SELECT cdf_unit_gaussian_p(1.96);          -- ~0.975

-- 卡方分布
SELECT cdf_chisq_p(3.84, 1.0);            -- ~0.95

-- t 分布
SELECT cdf_tdist_pinv(0.975, 10.0);       -- df=10 时 95% 置信区间的临界值

-- 泊松分布
SELECT rdf_poisson(5, 3.0);               -- lambda=3 时 P(X=5)

-- 贝塔分布
SELECT rdf_beta(0.5, 2.0, 5.0);           -- x=0.5, a=2, b=5 处的概率密度
```
