## 用法

来源：

- [官方 README](https://github.com/fugu13/pgdmn/blob/b68ce1244c5b555a1c4405b669eb5ab5aa11c77f/README.md)
- [扩展控制文件](https://github.com/fugu13/pgdmn/blob/b68ce1244c5b555a1c4405b669eb5ab5aa11c77f/pgdmn.control)
- [pgrx 清单](https://github.com/fugu13/pgdmn/blob/b68ce1244c5b555a1c4405b669eb5ab5aa11c77f/Cargo.toml)

`pgdmn` 在 PostgreSQL 内计算 Decision Model and Notation (DMN) 模型与 Friendly Enough Expression Language (FEEL) 表达式。

### 启用

0.1.0 版本为 PostgreSQL 14–18 提供 pgrx feature。针对准确大版本构建并安装后，以超级用户创建不可迁移扩展：

```sql
CREATE EXTENSION pgdmn;
```

它不要求预加载。扩展内嵌 vendored dsntk 0.3 引擎；升级应视为应用规则引擎变更，并在发布前测试已保存模型。

### 计算 FEEL

`feel_eval` 返回 JSONB，并接受可选 JSONB context。若结果不能表示为所请求的 PostgreSQL 类型，typed variant 会报错。

```sql
SELECT feel_eval('1 + 2');
SELECT feel_eval('x * 2', '{"x":21}'::jsonb);
SELECT feel_eval_numeric('x * 2', '{"x":21}'::jsonb);
```

Typed function 包括 `feel_eval_text`、`feel_eval_numeric`、`feel_eval_bool`、`feel_eval_date`、`feel_eval_timestamp` 与 `feel_eval_interval`。

### 加载并计算 DMN

`dmn_load` 把 DMN XML 解析为 `dmnmodel`。`dmn_eval` 使用 JSONB 输入计算具名 decision、business knowledge model 或 decision service。

```sql
WITH model AS (
  SELECT dmn_load($dmn$<definitions>...</definitions>$dmn$) AS value
)
SELECT dmn_eval(
  value,
  'Eligibility',
  '{"Age":30,"Income":75000}'::jsonb
)
FROM model;
```

使用 `dmn_invocables`、`dmn_info`、`dmn_name`、`dmn_namespace` 与 `dmn_xml` 检查模型。计算前应校验不可信 XML 并限制其大小；模型解析与规则执行会占用调用会话所在后端的 CPU 和内存。

