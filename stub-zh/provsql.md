## 用法

来源：[README](https://github.com/PierreSenellart/provsql/blob/master/doc/provsql.md)，[getting started](https://provsql.org/docs/user/getting-provsql.html)，[user docs](https://provsql.org/docs/)，[SQL API index](https://provsql.org/docs/sql/)

`provsql` 为 PostgreSQL 增加 semiring provenance 与 uncertainty management。上游文档覆盖 provenance tracking、semiring evaluation、probability、Shapley/Banzhaf 值、where-provenance、update provenance 以及 temporal features。

### 加载扩展

```ini
shared_preload_libraries = 'provsql'
```

```sql
CREATE EXTENSION provsql CASCADE;
```

`CASCADE` 形式会在需要时自动安装 `uuid-ossp`。getting-started 文档明确说明，预加载这一步是必须的，因为 ProvSQL 安装了 planner hook。

### 为表启用 provenance

```sql
SELECT provsql.add_provenance('mytable');

SELECT name, provenance()
FROM mytable;

SELECT provsql.remove_provenance('mytable');
```

用户文档还描述了 provenance mapping：

```sql
SELECT create_provenance_mapping('my_mapping', 'mytable', 'column_name');
SELECT create_provenance_mapping_view('my_mapping_view', 'mytable', 'column_name');
```

### 概率与 semiring 工作流

为 tuple token 赋概率：

```sql
SELECT set_prob(provenance(), 0.8)
FROM mytable
WHERE id = 1;
```

在 semiring 上计算 provenance：

```sql
SELECT city,
       provenance_evaluate(
         provenance(),
         'personnel_level',
         'unclassified'::classification_level,
         'security_plus',
         'security_times'
       )
FROM (SELECT DISTINCT city FROM personnel) AS t;
```

计算影响度分数：

```sql
SELECT shapley(provenance(), m.token)
FROM my_mapping AS m;
```

文档还说明了 `shapley_all_vars`、`banzhaf` 和 `banzhaf_all_vars`。

### 额外模式

上游记录了以下 session GUC：

```sql
SET provsql.where_provenance = on;
SET provsql.update_provenance = on;
```

用户指南还单独介绍了 `get_valid_time`、`timetravel`、`timeslice`、`history` 和 `undo` 等 temporal helper。

### 说明

- 上游测试覆盖 PostgreSQL 10 到 18。
- Git tag 显示仓库当前打包发布版本为 `v1.2.3`。
