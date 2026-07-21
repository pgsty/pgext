## 用法

来源：

- [FbSQL 0.1.0 README](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/README.md)
- [FbSQL 0.1.0 changes](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/Changes)
- [Extension control file](https://github.com/dsc-chiba-u/FbSQL/blob/v0.1.0/fbsql.control)
- [PGXN release](https://pgxn.org/dist/fbsql/0.1.0/)

`fbsql` 是一个统计建模 DSL 证明概念，它将关系型 SQL 查询作为输入并返回行数据，同时使用 R 公式语法描述模型。0.1.0 版本通过 PL/R 实现广义线性模型的拟合，并通过纯 PL/pgSQL 进行预测。

### 前提条件

FbSQL 在 PostgreSQL 16 上开发和测试，并需要 PL/R 8.4.0 或更高版本以及 R。`plr` 是一个不受信任的语言，因此超级用户必须安装依赖项和扩展。

```sql
CREATE EXTENSION fbsql CASCADE;
SELECT fbsql.version();
```

仅授予常规用户所需的函数访问权限和源数据权限。

### 核心工作流

拟合二元流失模型并保留返回的关系：

```sql
CREATE TEMP TABLE churn_model AS
SELECT *
FROM fbsql.fit_glm(
  relation => $$
    SELECT churn_flag, age, gender
    FROM customer
    WHERE created_at >= DATE '2025-01-01'
      AND created_at <  DATE '2026-01-01'
  $$,
  formula => 'churn_flag ~ age + gender',
  family => 'binomial'
);
```

预测接受新行的查询以及保存模型的查询。由于它返回 `SETOF record`，因此在调用时提供输出列：

```sql
SELECT customer_id, churn_flag_predicted
FROM fbsql.predict_glm(
  relation => $$SELECT customer_id, age, gender FROM customer_2026$$,
  model    => $$SELECT * FROM churn_model$$
) AS p(
  customer_id bigint,
  age integer,
  gender text,
  churn_flag_predicted double precision
);
```

### 重要对象

- `fbsql.fit_glm(relation, formula, family)` 返回每个模型项的一行、重复拟合统计信息以及包含预测所需信息的 `metadata jsonb`。
- `fbsql.predict_glm(relation, model, on_new_levels)` 将 `<response>_predicted` 添加到输入行。默认情况下，`on_new_levels` 为 `error` 或者设置为 `na` 以在未见过的因子水平上生成空预测。
- `fbsql.version()` 报告扩展版本。

### 支持的功能和注意事项

0.1.0 版支持带有恒等连接的高斯模型以及带有对数几率连接的二项式模型，使用数值和因子预测变量。拟合应用完整案例分析并报告已用和丢弃的行数；预测在预测变量为空时返回 `NULL`。预测使用存储的系数和元数据，并不会在运行时调用 R。

交互、自定义对比度、偏移量、权重、预测区间、其他家庭和连接以及分布式拟合均不支持。`relation` 和 `model` 参数包含 SQL 代码：从可信的 SQL 构造它们，而不是未经清理的用户输入，并审查执行角色的权限。
