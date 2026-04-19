## 用法

来源：[Docs site](https://michael-golfi.github.io/pg_liquid/)，[README](https://github.com/michael-golfi/pg_liquid/blob/main/README.md)，[Release v0.1.7](https://github.com/michael-golfi/pg_liquid/releases/tag/v0.1.7)，[SQL install script](https://github.com/michael-golfi/pg_liquid/blob/main/sql/pg_liquid--0.1.7.sql)

`pg_liquid` 将 Liquid 风格的 graph 与 compound query 引入 PostgreSQL。它在 `liquid` schema 中存储 graph state，并提供普通 query、principal-bound query 与 least-privilege read 的 SQL 入口。

### 核心查询接口

```sql
CREATE EXTENSION pg_liquid;

SELECT *
FROM liquid.query($$
  Edge("a","b").
  Edge("b","c").
  Path(X,Y) :- Edge(X,Y).
  Path(X,Y) :- Edge(X,Z), Path(Z,Y).
  Path("a",Y)?
$$) AS t(y text);
```

- `liquid.query(program text)`：执行 Liquid facts、rules 和一个最终 query。
- `liquid.query_as(principal text, program text)`：以 principal 绑定方式执行。
- `liquid.read_as(principal text, program text)`：最小权限读取封装，面向应用侧读取。

### 语言与建模特性

- 以 `.` 结尾的 facts 和 rules
- 每个 program 只能有一个最终 `?` query
- 通过 `Edge(...)` 表示 graph edges
- 形如 `Type@(role=value, ...)` 的 typed compounds
- query-local recursive rules
- 内建策略 compound，例如 `CompoundReadByRole` 与 `liquid/acts_for`

### 行规范化器

```sql
SELECT liquid.create_row_normalizer(
  'account_profiles'::regclass,
  'account_profile_normalizer',
  'AccountProfile',
  '{"account_id":"id","display_name":"display_name","tier":"tier"}'::jsonb,
  true
);
```

- `liquid.create_row_normalizer(...)`：将关系行投影为 Liquid compounds。
- `liquid.rebuild_row_normalizer(...)`：在表结构变化后重新生成绑定。
- `liquid.drop_row_normalizer(...)`：删除规范化器，并可选清理已生成绑定。

### 注意事项

- 上游在 PostgreSQL 14 到 18 上验证该扩展。
- 随附 SQL 会从 `PUBLIC` 撤销 `query_as` 与 `read_as`；应按需显式授权。
- 当客户端不应声明新 facts 时，`read_as(...)` 是更安全的应用入口。
