## 用法

来源：

- [ProvSQL 1.12.0 文档](https://github.com/PierreSenellart/provsql/blob/v1.12.0/doc/provsql.md)
- [ProvSQL 1.12.0 发行版](https://github.com/PierreSenellart/provsql/releases/tag/v1.12.0)
- [ProvSQL 1.12.0 变更日志](https://github.com/PierreSenellart/provsql/blob/v1.12.0/CHANGELOG.md)
- [ProvSQL 1.12.0 控制文件](https://github.com/PierreSenellart/provsql/blob/v1.12.0/provsql.common.control)
- [ProvSQL 用户文档](https://provsql.org/docs/user/introduction.html)

`provsql` 为 PostgreSQL 添加半环溯源与不确定性管理功能。上游文档涵盖溯源追踪、半环求值、概率、Shapley 和 Banzhaf 值、位置溯源、更新溯源以及时态功能。

### 加载并追踪溯源

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

如果需要，`CASCADE` 形式会自动安装 `uuid-ossp`。入门指南指出，预加载步骤是强制要求，因为 ProvSQL 会安装规划器钩子。

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

用户文档还介绍了溯源映射：

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

计算影响力得分：

```sql
SELECT shapley(provenance(), m.token)
FROM mytable, my_mapping AS m;

SELECT banzhaf(provenance(), m.token)
FROM mytable, my_mapping AS m;
```

文档还介绍了 `shapley_all_vars` 和 `banzhaf_all_vars`，用于一次计算所有输入变量的得分。

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

当前文档包含 `sr_how`、`sr_which`、`sr_tropical`、`sr_viterbi`、`sr_lukasiewicz`、`sr_minmax` 和 `sr_maxmin` 的编译式包装器。对于 PostgreSQL 14 及更高版本，还包含在多范围值上运行的 `sr_temporal`、`sr_interval_num` 和 `sr_interval_int`。

```sql
SELECT city,
       sr_minmax(provenance(), 'personnel_level',
                 'unclassified'::classification_level) AS clearance
FROM (SELECT DISTINCT city FROM personnel) AS t;

SELECT entity_id, sr_temporal(provenance(), 'validity_mapping')
FROM mytable;
```

高级用户仍可定义自定义半环，并通过 `provenance_evaluate` 或 `aggregation_evaluate` 对其求值；如果已有编译半环与所需代数相符，上游建议优先使用该编译半环。

### 额外模式与辅助工具

上游记录的会话 GUC 包括：

```sql
SET provsql.active = on;
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
SET provsql.tool_search_path = '/opt/d4:/home/postgres/bin';
SET provsql.aggtoken_text_as_uuid = on;

-- After probability_evaluate(...), inspect the route that actually ran:
SHOW provsql.last_eval_method;
```

`provsql.tool_search_path` 用于 `d4`、`c2d`、`dsharp`、`minic2d`、`weightmc` 和 `graph-easy` 等外部概率及可视化工具。`provsql.last_eval_method` 保存最近一次选用的概率求值方法。`provsql.aggtoken_text_as_uuid` 会让聚合令牌单元格显示为对应的溯源 UUID；`agg_token_value_text(token)` 可以恢复这些聚合令牌的显示文本。

用户指南还分别介绍了位置溯源辅助工具、更新溯源、`get_valid_time`、`timetravel`、`timeslice`、`history` 和 `undo` 等时态辅助工具、电路检查工具 `circuit_subgraph(root, max_depth)` 和 `resolve_input(uuid)`，以及用于准备辅助搜索路径的 `setup_search_path()`。

### 当前概率与推理接口

近期发行版实质性扩展了 SQL 覆盖范围和概率求值能力：

- `FROM` 之外的子查询，包括 `EXISTS`、`NOT EXISTS`、`IN`、`NOT IN`、`ANY`、`ALL`、行值 `IN`、标量子查询以及 `ARRAY(SELECT ...)`；
- `LEFT`、`RIGHT` 和 `FULL` 外连接，以及修正后的 `EXCEPT` 和 `EXCEPT ALL` 溯源；
- 聚合中遵循 SQL 语义的 `NULL` 处理，以及 `COUNT`、`SUM`、`MIN`、`MAX` 和 `AVG` 的精确 `HAVING` 聚合概率；
- 通过方法目录和成本选择器选取概率方法，支持 `karp-luby`、`stopping-rule`、`sieve`、`d-tree` 和 `probability_bounds`；
- 精确的有界树宽递归可达性、不安全 UCQ 的联合宽度编译、安全 UCQ 的 Möbius 反演，以及循环递归的吸收式溯源；
- 通过 `target | evidence` 操作符和整元组 `given()`/前缀形式实现条件事件与分布；
- 连续和离散的 `random_variable` 家族，包括正态、伽马、对数正态、贝塔、威布尔、帕累托、逆伽马、逆高斯、逻辑斯蒂、泊松、二项、几何、超几何和负二项分布；
- 分布参数本身也是随机变量的分层贝叶斯模型，并在存在闭式解时执行共轭后验更新；
- 随源数据变化仍保持正确的维护式溯源映射，以及 `NOT IN`、`EXCEPT` 和可空随机变量中符合 SQL 规范的 `NULL` 行为。

例如，可以基于已观察证据对连续值进行条件化，并读取后验期望：

```sql
WITH model AS (
  SELECT normal(20, 5) AS reading
)
SELECT expected(reading | (reading > 25))
FROM model;
```

`agg_token` 类型支持概率聚合表达式所需的算术运算、一元负号和比较操作。请参考官方概率与连续分布章节，在精确求值、编译式求值和基于采样的求值方法之间进行选择。

### 注意事项

- 版本 1.11.1 修正了聚合比较的溯源，并使空的 `sum`、`min`、`max` 和 product 分组返回 SQL `NULL`。它还改变了为 `count(*)` 生成的内容寻址令牌；如果旧的空输入行为很重要，请重新运行曾将这些令牌物化的查询。
- 版本 1.12.0 将 Möbius 求值扩展至自连接，允许 `sr_formula` 使用可选映射呈现所有门类型，并通过 `provsql.last_eval_method` 报告 `sq-rewrite`、`bounded-jw` 和 `reachability`。它还修复了 DML 重写场景，并使重复调用 `remove_provenance()` 变得安全。
- `ALTER EXTENSION provsql UPDATE` 会安装 SQL 变更。由于 ProvSQL 会按会话缓存函数 OID，请重新连接在更新期间一直保持打开的会话。
- 1.12.0 控制文件设置 `default_version = '1.12.0'`，要求 `uuid-ossp`，将扩展标记为受信任且不可重定位。
- 上游文档称 ProvSQL 已在 PostgreSQL 10 到 18 上通过测试。
- `provsql.update_provenance` 和多范围半环要求 PostgreSQL 14 或更高版本。
- 更新溯源追踪仍处于实验阶段；在大范围启用前，请验证其存储与性能开销。
