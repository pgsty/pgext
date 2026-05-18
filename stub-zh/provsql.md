## 用法

来源：[README](https://github.com/PierreSenellart/provsql/blob/master/README.md), [v1.4.0 release](https://github.com/PierreSenellart/provsql/releases/tag/v1.4.0), [latest release](https://github.com/PierreSenellart/provsql/releases/tag/v1.6.0), [v1.4.0 control](https://github.com/PierreSenellart/provsql/blob/v1.4.0/provsql.common.control), [getting started](https://provsql.org/docs/user/getting-provsql.html), [configuration](https://provsql.org/docs/user/configuration.html), [semirings](https://provsql.org/docs/user/semirings.html), [v1.4.0 upgrade](https://github.com/PierreSenellart/provsql/blob/v1.4.0/sql/upgrades/provsql--1.3.1--1.4.0.sql)

`provsql` 为 PostgreSQL 添加 semiring provenance 和 uncertainty management。上游文档覆盖 provenance tracking、semiring evaluation、probabilities、Shapley 和 Banzhaf values、where-provenance、update provenance，以及 temporal features。

### 加载并跟踪 Provenance

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

`CASCADE` 形式会在需要时自动安装 `uuid-ossp`。getting-started guide 说明 preload 步骤是强制的，因为 ProvSQL 安装了 planner hook。

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

计算 influence scores：

```sql
SELECT shapley(provenance(), m.token)
FROM mytable, my_mapping AS m;

SELECT banzhaf(provenance(), m.token)
FROM mytable, my_mapping AS m;
```

文档还描述了 `shapley_all_vars` 和 `banzhaf_all_vars`，用于一次性为所有输入变量计算分数。

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

`provsql` 1.4.0 为 `sr_how`、`sr_which`、`sr_tropical`、`sr_viterbi`、`sr_lukasiewicz`、`sr_minmax` 和 `sr_maxmin` 增加 compiled wrappers。对于 PostgreSQL 14 及之后版本，它还增加了基于 multirange values 的 `sr_temporal`、`sr_interval_num` 和 `sr_interval_int`。

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

高级用户仍可定义 custom semirings，并使用 `provenance_evaluate` 或 `aggregation_evaluate` 进行求值；当所需代数已有对应 compiled semiring 时，上游推荐使用 compiled semirings。

### 额外模式与辅助函数

上游文档化的 session GUC 包括：

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;
```

`provsql.tool_search_path` 用于 `d4`、`c2d`、`dsharp`、`minic2d`、`weightmc` 和 `graph-easy` 等外部概率与可视化工具。`provsql.aggtoken_text_as_uuid` 会让 aggregate-token cells 显示为 provenance UUID；`agg_token_value_text(token)` 可恢复这些 aggregate tokens 的显示文本。

用户指南还单独记录了 where-provenance helpers、update provenance，以及 `get_valid_time`、`timetravel`、`timeslice`、`history`、`undo` 等 temporal helpers。版本 1.4.0 还增加了 circuit inspection helpers `circuit_subgraph(root, max_depth)` 和 `resolve_input(uuid)`，它们供 ProvSQL Studio 使用，也可用于浏览 circuit fragments。

### 说明

- `db/extension.csv` 中的包行列出版本 `1.4.0`、package `provsql`、依赖 `uuid-ossp`，并支持 PostgreSQL 14 到 18。
- 上游文档说明 ProvSQL 已在 PostgreSQL 10 到 18 上测试。即使上游 GitHub 现在已有 `v1.6.0`，Pigsty row 仍跟踪 `1.4.0`，因此对 Pigsty 构建应以 package metadata 作为安装版本。
- `provsql.update_provenance` 和 multirange semirings 要求 PostgreSQL 14 或之后版本。
