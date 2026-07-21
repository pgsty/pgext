## 用法

来源：

- [ProvSQL 1.11.0 文档](https://github.com/PierreSenellart/provsql/blob/v1.11.0/doc/provsql.md)
- [ProvSQL 1.11.0 发行版](https://github.com/PierreSenellart/provsql/releases/tag/v1.11.0)
- [ProvSQL 1.10.0 发行版](https://github.com/PierreSenellart/provsql/releases/tag/v1.10.0)
- [ProvSQL 1.11.0 控制文件](https://github.com/PierreSenellart/provsql/blob/v1.11.0/provsql.common.control)
- [ProvSQL 用户文档](https://provsql.org/docs/user/introduction.html)

`provsql` 将半环证明和不确定性管理添加到 PostgreSQL 中。上游文档介绍了证明跟踪、半环评估、概率、Shapley 和 Banzhaf 值、来源证明、更新证明以及时间特征。

### 加载并追踪证明

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

`CASCADE` 形式会自动安装 `uuid-ossp`（如果需要）。入门指南指出，预加载步骤是必需的，因为 ProvSQL 安装了一个计划器挂钩。

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

用户文档还描述了证明映射：

```sql
SELECT create_provenance_mapping('my_mapping', 'mytable', 'column_name');
SELECT create_provenance_mapping_view('my_mapping_view', 'mytable', 'column_name');
```

### 概率与影响

为元组标记分配概率：

```sql
SELECT set_prob(provenance(), 0.8)
FROM mytable
WHERE id = 1;

SELECT name, probability_evaluate(provenance()) AS prob
FROM mytable;
```

计算影响得分：

```sql
SELECT shapley(provenance(), m.token)
FROM mytable, my_mapping AS m;

SELECT banzhaf(provenance(), m.token)
FROM mytable, my_mapping AS m;
```

文档中也介绍了 `shapley_all_vars` 和 `banzhaf_all_vars`，用于一次性计算所有输入变量的得分。

### 内置半环

内置半环函数使用证明标记和证明映射表：

```sql
SELECT name, sr_boolean(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_formula(provenance(), 'my_mapping')
FROM mytable;

SELECT name, sr_how(provenance(), 'my_mapping')
FROM mytable;
```

当前文档包括编译后的包装器：`sr_how`、`sr_which`、`sr_tropical`、`sr_viterbi`、`sr_lukasiewicz`、`sr_minmax` 和 `sr_maxmin`。对于 PostgreSQL 14 及以上版本，它们还包括对多范围值的 `sr_temporal`、`sr_interval_num` 和 `sr_interval_int`。

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

高级用户仍然可以定义自定义半环，并使用 `provenance_evaluate` 或 `aggregation_evaluate` 评估它们；上游建议当一个半环匹配所需的代数时，使用编译的半环。

### 额外模式和助手

上游文档中记录的会话 GUC 包括：

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.last_eval_method = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;
```

`provsql.tool_search_path` 用于外部概率和可视化工具，如 `d4`、`c2d`、`dsharp`、`minic2d`、`weightmc` 和 `graph-easy`。`provsql.last_eval_method` 存储最后选择的概率评估方法。`provsql.aggtoken_text_as_uuid` 使聚合标记单元格显示为其证明 UUID；`agg_token_value_text(token)` 可以恢复这些聚合标记的显示文本。

用户指南单独记录了来源证明助手、更新证明、时间助手，如 `get_valid_time`、`timetravel`、`timeslice`、`history` 和 `undo`，电路检查助手 `circuit_subgraph(root, max_depth)` 和 `resolve_input(uuid)` 以及用于准备助手搜索路径的 `setup_search_path()`。

### 当前概率和推理界面

从 1.9 到 1.11 的多次发布显著扩展了 SQL 覆盖范围和概率评估：

- `FROM` 之外的子查询，包括 `EXISTS`、`NOT EXISTS`、`IN`、`NOT IN`、`ANY`、`ALL`、行值 `IN` 子查询和 `ARRAY(SELECT ...)`；
- `LEFT`、`RIGHT` 和 `FULL` 外连接，以及修正的 `EXCEPT` 和 `EXCEPT ALL` 证明；
- 遵循 SQL 的聚合处理方法和精确 `HAVING` 聚合概率，适用于 `COUNT`、`SUM`、`MIN`、`MAX` 和 `AVG`；
- 通过方法目录和成本选择器进行概率方法的选择，包括 `karp-luby`、`stopping-rule`、`sieve`、`d-tree` 和 `probability_bounds`；
- 精确的有界树宽递归可达性、不安全-UCQ 联合宽度编译、安全 UCQ 的 Möbius 反演以及循环递归的吸收证明；
- 通过 `target | evidence` 操作符和整个元组的 `given()`/前缀形式进行条件事件和分布处理；
- 连续和离散 `random_variable` 家族，包括正态、伽玛、对数正态、贝塔、威布尔、帕累托、逆伽玛、逆高斯、逻辑斯谛、泊松、二项式、几何、超几何和负二项分布；
- 分层贝叶斯模型，其中分布参数本身是随机变量，在可用闭合形式时进行共轭后验更新；
- 随着源数据变化而保持正确的维护证明映射，以及对 `NULL` 和可为空的随机变量的 SQL 合规 `NOT IN` 行为。

例如，根据观察到的证据条件化一个连续值并读取后验期望：

```sql
WITH model AS (
  SELECT normal(20, 5) AS reading
)
SELECT expected(reading | (reading > 25))
FROM model;
```

`agg_token` 类型支持概率聚合表达式的算术、一元负号和比较。使用官方的概率和连续分布章节来选择精确、编译或基于采样的评估方法。

### 备注

- 1.11.0 控制文件设置了 `default_version = '1.11.0'`，要求 `uuid-ossp`，标记扩展为受信任的，并且不可重定位。
- 上游文档指出 ProvSQL 已在 PostgreSQL 10 至 18 版本上进行了测试。
- `provsql.update_provenance` 和多范围半环需要 PostgreSQL 14 或更高版本。
- 更新证明追踪仍处于实验阶段；启用之前请验证其存储和性能成本。

还应分别验证 `NULL` 与 `EXCEPT` 的 SQL 语义是否符合应用预期。
