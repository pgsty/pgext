## 用法

来源：

- [官方 README](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/README.md)
- [2.1 版本升级 SQL](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/htc--2.0--2.1.sql)
- [C 估计器接口](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/functions/htc_compute.sql)
- [回归测试工作流](https://gitlab.com/nfiesta/nfiesta_htc/-/blob/267e25465ee4ea073fdbfbdddf1bf8d7d9c6359c/sql/htc.sql)

`htc` 为连续总体抽样实现 Horvitz-Thompson 估计器。它面向专门的调查与森林清查计算，接收抽样簇、密度、权重、面积、包含概率和辅助变量数组，而不是普通的行聚合。

### 核心流程

创建 `htc` 前，先安装上游版本使用的过程语言和数值计算前置项：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION plpython3u;
CREATE EXTENSION htc;

SELECT htc_version(), htc_version_ok();
```

使用原始单阶段接口时，将一个层的数据聚合为对齐数组，并显式传入层级常量：

```sql
SELECT (htc_compute(
    ARRAY[1, 2, 3, 4],
    ARRAY[1.1, 2.2, 3.3, 4.4]::double precision[],
    ARRAY[0.75, 0.65, 0.55, 0.45]::double precision[],
    123.456,
    75.0,
    100
)).*;
```

应以上游回归测试工作流作为准备数组和解释返回值 `total`、`var` 的规范指南；估计结果的有效性取决于是否保持所需顺序、形状、单位和抽样设计假设。

### 对象索引

- `htc_compute(...)` 使用 C 返回单阶段 `total` 和 `var` 估计值。
- `htc_compute_sweight_ha(...)` 是每公顷抽样权重变体。
- `htc_gbeta(...)` 从 JSONB 矩阵构建广义回归系数。
- `htc_total2p(...)` 返回两阶段总量、总方差以及各阶段方差分量。
- `htc_variance_phase1(...)` 和 `htc_variance_phase2(...)` 是两阶段方差辅助函数。
- `htc_version()` 和 `htc_version_ok()` 检查已加载共享库是否与 SQL 发行版匹配。

### 运维说明

版本 `2.1` 可重定位，无预加载要求。较新的两阶段例程是导入 NumPy 和 SciPy 的 `plpython3u` 函数，因此 PostgreSQL 服务端 Python 必须能使用这些模块。`plpython3u` 是非受信过程语言，通常需要超级用户级控制才能安装；应限制可以替换或调用过程代码的角色。

JSONB 例程会校验数组数量、维度、空值以及不允许为零的概率，但这并不能验证统计设计。应把上游回归测试输入视为数据模式，在二进制或 SQL 升级后检查 `htc_version_ok()`，并在用于运维决策前独立验证估计结果。
