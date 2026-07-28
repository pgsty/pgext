## 用法

来源：

- [官方上游 README](https://github.com/lowdown-labs/pg_fela/blob/247a5038f9b88d783524a83a69afaffcf08fdaa7/README.md)
- [官方扩展控制文件 (pg_fela.control)](https://github.com/lowdown-labs/pg_fela/blob/247a5038f9b88d783524a83a69afaffcf08fdaa7/pg_fela.control)
- [官方实现源代码](https://github.com/lowdown-labs/pg_fela/blob/247a5038f9b88d783524a83a69afaffcf08fdaa7/src/lib.rs)

`pg_fela` — 一个完整的 Rust pgrx PostgreSQL 扩展，可以在数据库内部运行冻结的表型基础模型 (FelaTab)：分类、填充缺失值、聚类、评分异常、排名特征重要性，并通过单个 SELECT 语句解释预测，无需训练步骤，没有任何内容离开 Postgres。将其用于相应的向量、模型或检索工作流。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_fela;

-- AutoML in a SELECT: learns from the labeled rows, predicts the ones where target is NULL
SELECT * FROM fela_automl('my_table', 'target_column');

-- Same, plus a per row trust/OOD score so a confident prediction on unfamiliar data gets flagged
SELECT * FROM fela_predict_trust('my_table', 'target_column');

-- Why did row 42 get this prediction? Top contributing features, signed toward/away from it
SELECT * FROM fela_explain_row('my_table', 'target_column', 42);

-- Implicit AutoML: builds my_table_ml, joining prediction/confidence/trust/ood/cluster back onto the base table
SELECT fela_create_view('my_table', 'target_column');
SELECT * FROM my_table_ml WHERE ood;             -- rows unlike anything the model learned from
SELECT * FROM my_table_ml ORDER BY confidence;   -- triage the least sure predictions first
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `fela_anomaly` 是一个扩展函数。
- `fela_argmax` 是一个扩展函数。
- `fela_automl` 是一个扩展函数。
- `fela_caps()` 是一个扩展函数。
- `fela_classify` 是一个扩展函数。
- `fela_classify_gated` 是一个扩展函数。
- `fela_cluster` 是一个扩展函数。
- `fela_cluster_ex` 是一个扩展函数。
- `fela_confidence` 是一个扩展函数。
- `fela_conformal_regress` 是一个扩展函数。
- `fela_conformal_threshold` 是一个扩展函数。
- `fela_create_view` 是一个扩展函数。
- `fela_detect_task` 是一个扩展函数。
- `fela_explain` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `1.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
