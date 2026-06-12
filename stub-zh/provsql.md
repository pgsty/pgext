## 用法

来源：[README](https://github.com/PierreSenellart/provsql/blob/v1.9.0/doc/provsql.md)、[v1.9.0 release](https://github.com/PierreSenellart/provsql/releases/tag/v1.9.0)、[v1.9.0 control](https://github.com/PierreSenellart/provsql/blob/v1.9.0/provsql.common.control)、[getting started](https://provsql.org/docs/user/getting-provsql.html)、[configuration](https://provsql.org/docs/user/configuration.html)、[semirings](https://provsql.org/docs/user/semirings.html)

`provsql` 为 PostgreSQL 增加 semiring provenance 与不确定性管理能力。上游文档覆盖 provenance tracking、semiring evaluation、probabilities、Shapley and Banzhaf values、where-provenance、update provenance 和 temporal features。

### 加载并跟踪 Provenance

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

如果需要，`CASCADE` 形式会自动安装 `uuid-ossp`。getting-started guide 说明 preload 步骤是必需的，因为 ProvSQL 会安装 planner hook。

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

用户文档还描述了 provenance mappings：

```sql
SELECT create_provenance_mapping('my_mapping', 'mytable', 'column_name');
SELECT create_provenance_mapping_view('my_mapping_view', 'mytable', 'column_name');
```

### 概率与影响力

为 tuple tokens 分配概率：

```sql
SELECT set_prob(provenance(), 0.8)
FROM mytable
WHERE id = 1;

SELECT name, probability_evaluate(provenance()) AS prob
FROM mytable;
```

计算影响力分数：

```sql
SELECT shapley(provenance(), m.token)
FROM mytable, my_mapping AS m;

SELECT banzhaf(provenance(), m.token)
FROM mytable, my_mapping AS m;
```

文档还描述了 `shapley_all_vars` 和 `banzhaf_all_vars`，用于一次性计算所有输入变量的分数。

### 内置 Semirings

内置 semiring 函数使用 provenance token 和 provenance mapping table：

```sql
SELECT name, sr_boolean(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_formula(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_how(provenance(), 'my_mapping')
FROM mytable;
```

当前文档包含 `sr_how`、`sr_which`、`sr_tropical`、`sr_viterbi`、`sr_lukasiewicz`、`sr_minmax` 和 `sr_maxmin` 的 compiled wrappers。对于 PostgreSQL 14 及之后版本，还包含基于 multirange values 的 `sr_temporal`、`sr_interval_num` 和 `sr_interval_int`。

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

高级用户仍可以定义 custom semirings，并通过 `provenance_evaluate` 或 `aggregation_evaluate` 求值；如果已有 compiled semiring 符合所需代数，上游建议优先使用它。

### 额外模式与辅助函数

上游文档记录的 session GUC 包括：

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.last_eval_method = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;
```

`provsql.tool_search_path` 用于 `d4`、`c2d`、`dsharp`、`minic2d`、`weightmc` 和 `graph-easy` 等外部概率与可视化工具。`provsql.last_eval_method` 会保存上一次选用的概率求值方法。`provsql.aggtoken_text_as_uuid` 会让 aggregate-token 单元格显示为其 provenance UUID；`agg_token_value_text(token)` 可恢复这些 aggregate tokens 的显示文本。

用户指南另行记录了 where-provenance helpers、update provenance、`get_valid_time`、`timetravel`、`timeslice`、`history` 和 `undo` 等 temporal helpers，`circuit_subgraph(root, max_depth)` 与 `resolve_input(uuid)` 这类 circuit-inspection helpers，以及用于准备 helper search path 的 `setup_search_path()`。

### v1.9.0 查询与概率说明

Release `1.9.0` 显著扩展了 provenance-aware queries 的 SQL 覆盖范围：

- `FROM` 之外的 subqueries，包括 `EXISTS`、`NOT EXISTS`、`IN`、`NOT IN`、`ANY`、`ALL`、row-valued `IN`、scalar subqueries 和 `ARRAY(SELECT ...)`；
- `LEFT`、`RIGHT` 和 `FULL` outer joins，并修正了 `EXCEPT` 与 `EXCEPT ALL` 的 provenance；
- aggregates 的 SQL-faithful `NULL` handling，以及 `COUNT`、`SUM`、`MIN`、`MAX` 和 `AVG` 的精确 `HAVING` aggregate probabilities；
- 通过 method catalog 和 cost chooser 选择 probability method，支持 `karp-luby`、`stopping-rule`、`sieve`、`d-tree` 和 `probability_bounds`；
- 幂等的 `add_provenance` 和 `create_provenance_mapping` 调用。

该 release 移除了旧的 `probability_benchmark` helper。`agg_token` 现在为 aggregate-token expressions 提供原生 arithmetic、unary minus 和 comparison 支持。

### 说明

- `db/extension.csv` 中的包行列出 version `1.9.0`、package `provsql`、dependency `uuid-ossp`，并标注 PostgreSQL 14 到 18 支持。
- v1.9.0 control file 设置 `default_version = '1.9.0'`，要求 `uuid-ossp`，将扩展标记为 trusted，且不可 relocatable。
- 上游文档说明 ProvSQL 已在 PostgreSQL 10 到 18 上测试；Pigsty package matrix 为 PostgreSQL 14-18。
- `provsql.update_provenance` 和 multirange semirings 要求 PostgreSQL 14 或更新版本。
