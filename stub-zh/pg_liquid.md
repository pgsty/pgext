


## 用法

来源：[文档站点](https://michael-golfi.github.io/pg_liquid/)，[README](https://github.com/michael-golfi/pg_liquid/blob/main/README.md)，[v0.1.7 发行版](https://github.com/michael-golfi/pg_liquid/releases/tag/v0.1.7)，[SQL 安装脚本](https://github.com/michael-golfi/pg_liquid/blob/main/sql/pg_liquid--0.1.7.sql)

`pg_liquid` 将 Liquid 风格的图与复合查询引入 PostgreSQL。它在 `liquid` 模式中存储图状态，并提供普通查询、绑定主体的查询与最小权限读取的 SQL 入口。

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

- `liquid.query(program text)`：执行 Liquid 事实、规则和一个最终查询。
- `liquid.query_as(principal text, program text)`：以主体绑定方式执行。
- `liquid.read_as(principal text, program text)`：最小权限读取封装，面向应用侧读取。

### 语言与建模特性

- 以 `.` 结尾的事实和规则
- 每个程序只能有一个最终 `?` 查询
- 通过 `Edge(...)` 表示图边
- 形如 `Type@(role=value, ...)` 的类型化复合项
- 查询局部递归规则
- 内建策略复合项，例如 `CompoundReadByRole` 与 `liquid/acts_for`

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

- `liquid.create_row_normalizer(...)`：将关系行投影为 Liquid 复合项。
- `liquid.rebuild_row_normalizer(...)`：在表结构变化后重新生成绑定。
- `liquid.drop_row_normalizer(...)`：删除规范化器，并可选清理已生成绑定。

### 注意事项

- 上游在 PostgreSQL 14 到 18 上验证该扩展。
- 随附 SQL 会从 `PUBLIC` 撤销 `query_as` 与 `read_as`；应按需显式授权。
- 当客户端不应声明新事实时，`read_as(...)` 是更安全的应用入口。
