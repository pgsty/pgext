## 用法

来源：

- [官方 pg_prob README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/README.md)
- [扩展 control 文件](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/pg_prob.control)
- [上游端到端示例](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_prob/example.md)

`pg_prob` 版本 `0.2.0` 增加了用于 SQL 蒙特卡洛建模的 `pgprob.dist` 概率类型。分布可以构造、通过算术组合、跨行聚合、采样和汇总，无需把包含不确定性的业务数据移出 PostgreSQL。

### 核心流程

```sql
CREATE EXTENSION pg_prob;

WITH estimate AS (
    SELECT pgprob.pert(5, 10, 20)
         + pgprob.pert(10, 15, 30) AS total_days
)
SELECT
    pgprob.summarize(total_days, 10000, 42) AS summary,
    pgprob.prob_below(total_days, 40, 10000, 42) AS finish_by_day_40
FROM estimate;
```

构造函数包括 `literal`、`normal`、`uniform`、`triangular`、`beta`、`lognormal`、`pert`、`poisson` 和 `exponential`。采样与统计函数包括 `sample`、`samples`、`summarize`、`mean`、`variance`、`stddev`、`percentile`、`prob_below`、`prob_above` 和 `prob_between`。聚合包括 `sum`、`avg`、`min`、`max` 及分布拟合变体；相关采样和条件构造函数支持更高级的模型。

### 运维说明

蒙特卡洛结果是估计值：样本数量同时决定误差与执行成本，若需可复现运行必须显式提供种子。对 `dist` 的算术会传播建模的不确定性，但模型假设与相关关系仍由调用方负责。扩展固定安装在 `pgprob` 模式中、不可重定位，且不限定超级用户安装。将结果用于财务或运维决策前，应以观测数据验证参数，并在目标工作负载上对大样本量进行基准测试。
