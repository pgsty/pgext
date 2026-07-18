


## 用法

来源：[README](https://github.com/PierreSenellart/provsql/blob/v1.9.0/doc/provsql.md)、[v1.9.0 发行版](https://github.com/PierreSenellart/provsql/releases/tag/v1.9.0)、[v1.9.0 控制文件](https://github.com/PierreSenellart/provsql/blob/v1.9.0/provsql.common.control)、[入门指南](https://provsql.org/docs/user/getting-provsql.html)、[配置](https://provsql.org/docs/user/configuration.html)、[半环](https://provsql.org/docs/user/semirings.html)

`provsql` 为 PostgreSQL 增加半环溯源与不确定性管理能力。上游文档覆盖溯源跟踪、半环求值、概率、Shapley 值与 Banzhaf 值、位置溯源、更新溯源和时态功能。

### 加载并跟踪溯源信息

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

如果需要，`CASCADE` 形式会自动安装 `uuid-ossp`。入门指南说明预加载步骤是必需的，因为 ProvSQL 会安装规划器钩子。

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

用户文档还描述了溯源映射：

```sql
SELECT create_provenance_mapping('my_mapping', 'mytable', 'column_name');
SELECT create_provenance_mapping_view('my_mapping_view', 'mytable', 'column_name');
```

### 概率与影响力

为元组令牌分配概率：

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

### 内置半环

内置半环函数使用溯源令牌和溯源映射表：

```sql
SELECT name, sr_boolean(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_formula(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_how(provenance(), 'my_mapping')
FROM mytable;
```

当前文档包含 `sr_how`、`sr_which`、`sr_tropical`、`sr_viterbi`、`sr_lukasiewicz`、`sr_minmax` 和 `sr_maxmin` 的编译型封装。对于 PostgreSQL 14 及之后版本，还包含基于多范围值的 `sr_temporal`、`sr_interval_num` 和 `sr_interval_int`。

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

高级用户仍可以定义自定义半环，并通过 `provenance_evaluate` 或 `aggregation_evaluate` 求值；如果已有编译型半环符合所需代数，上游建议优先使用它。

### 额外模式与辅助函数

上游文档记录的会话 GUC 包括：

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.last_eval_method = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;
```

`provsql.tool_search_path` 用于 `d4`、`c2d`、`dsharp`、`minic2d`、`weightmc` 和 `graph-easy` 等外部概率与可视化工具。`provsql.last_eval_method` 会保存上一次选用的概率求值方法。`provsql.aggtoken_text_as_uuid` 会让聚合令牌单元格显示为其溯源 UUID；`agg_token_value_text(token)` 可恢复这些聚合令牌的显示文本。

用户指南另行记录了位置溯源辅助函数、更新溯源、`get_valid_time`、`timetravel`、`timeslice`、`history` 和 `undo` 等时态辅助函数，`circuit_subgraph(root, max_depth)` 与 `resolve_input(uuid)` 这类电路检查辅助函数，以及用于准备辅助函数搜索路径的 `setup_search_path()`。

### v1.9.0 查询与概率说明

`1.9.0` 发行版显著扩展了溯源感知查询的 SQL 覆盖范围：

- `FROM` 之外的子查询，包括 `EXISTS`、`NOT EXISTS`、`IN`、`NOT IN`、`ANY`、`ALL`、行值 `IN`、标量子查询和 `ARRAY(SELECT ...)`；
- `LEFT`、`RIGHT` 和 `FULL` 外连接，并修正了 `EXCEPT` 与 `EXCEPT ALL` 的溯源；
- 聚合中忠实于 SQL 语义的 `NULL` 处理，以及 `COUNT`、`SUM`、`MIN`、`MAX` 和 `AVG` 的精确 `HAVING` 聚合概率；
- 通过方法目录和成本选择器选择概率方法，支持 `karp-luby`、`stopping-rule`、`sieve`、`d-tree` 和 `probability_bounds`；
- 幂等的 `add_provenance` 和 `create_provenance_mapping` 调用。

该发行版移除了旧的 `probability_benchmark` 辅助函数。`agg_token` 现在为聚合令牌表达式提供原生算术、一元负号和比较支持。

### 说明

- `db/extension.csv` 中的软件包行列出版本 `1.9.0`、软件包 `provsql`、依赖 `uuid-ossp`，并标注支持 PostgreSQL 14 到 18。
- v1.9.0 控制文件设置 `default_version = '1.9.0'`，要求 `uuid-ossp`，将扩展标记为可信，且不可迁移。
- 上游文档说明 ProvSQL 已在 PostgreSQL 10 到 18 上测试；Pigsty 软件包矩阵覆盖 PostgreSQL 14-18。
- `provsql.update_provenance` 和多范围半环要求 PostgreSQL 14 或更新版本。
