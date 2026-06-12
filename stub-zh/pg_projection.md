## 用法

- 来源：[pg_projection README](https://github.com/suissa/pg_projection)，[SQL definitions](https://github.com/suissa/pg_projection/blob/main/pg_projection--1.0.sql)，[control file](https://github.com/suissa/pg_projection/blob/main/pg_projection.control)

`pg_projection` 为 PostgreSQL `jsonb` 提供类似 MongoDB 的读取投影能力。1.0 SQL 文件定义了两个函数：`pg_project(jsonb, jsonb)` 用于单个 JSON 文档，`pg_project_set(text, jsonb)` 用于把查询结果转换为 JSON 数组后再做投影。

### 投影单个 JSONB 值

投影值使用数字标志：`1` 表示包含字段，`0` 表示排除字段。

```sql
CREATE EXTENSION pg_projection;

SELECT pg_project(
  '{"_id": 7, "name": "Ada", "email": "ada@example.test", "secret": "x"}'::jsonb,
  '{"name": 1, "email": 1}'::jsonb
);
-- {"_id": 7, "name": "Ada", "email": "ada@example.test"}
```

在包含模式下，如果 `_id` 存在，默认会被保留。调用方只想返回选中字段时，需要显式排除它：

```sql
SELECT pg_project(
  '{"_id": 7, "name": "Ada", "email": "ada@example.test"}'::jsonb,
  '{"_id": 0, "name": 1}'::jsonb
);
-- {"name": "Ada"}
```

### 排除字段

当投影使用 `0` 时，函数会从原始文档出发，移除匹配的顶层 key：

```sql
SELECT pg_project(
  '{"name": "Ada", "internal_id": "a-1", "secret_key": "k"}'::jsonb,
  '{"internal_id": 0, "secret_key": 0}'::jsonb
);
-- {"name": "Ada"}
```

### 投影查询结果

`pg_project_set(query_text, projection_json)` 会执行传入的 SQL 文本，用 `to_jsonb(t)` 转换每一行，应用 `pg_project`，并返回 JSON 数组：

```sql
SELECT pg_project_set(
  'SELECT id, username, password_hash FROM users WHERE active',
  '{"password_hash": 0}'::jsonb
);
```

由于 `query_text` 是动态 SQL，只应传入由你控制的应用代码或迁移代码组装出的可信查询字符串。不要把不可信的用户输入拼接进这个参数。

### 注意事项

- SQL 实现只投影顶层 key；它没有实现 MongoDB 的嵌套路径投影。
- 投影值会在内部转换为整数，因此应使用数字 `0` 和 `1` 标志。
- `pg_project(jsonb, jsonb)` 声明为 `IMMUTABLE STRICT`；`pg_project_set(text, jsonb)` 声明为 `STABLE`。
