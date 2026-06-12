来源：[pgGraph v0.1.7 README](https://github.com/Evokoa/pgGraph/blob/v0.1.7/README.md)、[Quickstart](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/quickstart.mdx)、[SQL API Reference](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/user_guide/api-reference.mdx)、[Schema Registration](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/user_guide/schema-registration.mdx)、[Configuration](https://github.com/Evokoa/pgGraph/blob/v0.1.7/docs/user_guide/configuration.mdx)。

## 用法

`pggraph` 是包名与 PGXN 发行名，但安装到 PostgreSQL 中的扩展名是 `graph`。该扩展会从普通 PostgreSQL 表构建派生图索引，并以这些表作为事实来源，通过 `graph` schema 提供图搜索、遍历、GQL 风格模式读取和路径函数。

上游项目将 v0.1.7 标记为早期 alpha。建议先在可丢弃或开发数据库中使用，并从源表重建图，而不是把生成的图产物当作权威数据。

### 基本图构建

```sql
CREATE EXTENSION IF NOT EXISTS graph;
SELECT graph.reset();

CREATE TABLE companies (
  id   text PRIMARY KEY,
  name text NOT NULL
);

CREATE TABLE people (
  id         text PRIMARY KEY,
  name       text NOT NULL,
  company_id text REFERENCES companies(id)
);

INSERT INTO companies VALUES
  ('c1', 'Acme Bank'),
  ('c2', 'Northwind Trading');

INSERT INTO people VALUES
  ('p1', 'Alice', 'c1'),
  ('p2', 'Bob', 'c1'),
  ('p3', 'Carol', 'c2');

SELECT * FROM graph.auto_discover('public');

SELECT node_count, edge_count, edge_types
FROM graph.status();
```

`graph.auto_discover('public')` 会扫描该 schema 中的主键和外键，注册发现到的源表与边，然后构建图。对于生产 schema，建议使用显式注册，确保标签、搜索列、权重和租户行为都经过有意设计。

### 手工注册

```sql
SELECT graph.reset();

SELECT graph.add_table(
  table_name := 'public.people'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name']
);

SELECT graph.add_table(
  table_name := 'public.companies'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name']
);

SELECT graph.add_edge(
  from_table    := 'public.people'::regclass,
  from_column   := 'company_id',
  to_table      := 'public.companies'::regclass,
  to_column     := 'id',
  label         := 'works_at',
  bidirectional := true
);

SELECT * FROM graph.build();
```

节点标识符必须匹配主键，或匹配唯一的 `NOT NULL` 索引。`columns` 控制可用于搜索和 GQL 的源表属性。遍历过滤下推需要通过单独的 `graph.add_filter_column()` 注册。

### 搜索、遍历与路径

```sql
SELECT node_table_name, node_id, node
FROM graph.search(
  property_key  := 'name',
  property_value := 'Alice',
  table_filter  := 'public.people'::regclass,
  mode          := 'exact',
  hydrate       := true
);

SELECT depth, node_table_name, node_id, edge_path
FROM graph.traverse(
  'public.people'::regclass,
  'p1',
  2,
  hydrate := false
);

SELECT step, node_table_name, node_id, edge_label
FROM graph.shortest_path(
  'public.people'::regclass,
  'p1',
  'public.companies'::regclass,
  'c1',
  hydrate := false
);
```

`hydrate := false` 返回紧凑的图坐标。启用 hydration 后，源行可见性仍由 PostgreSQL ACL 和 RLS 控制，过期坐标会失败关闭，而不会伪造行。

### GQL 查询

```sql
SELECT row
FROM graph.gql(
  'MATCH (p:people)-[:works_at]->(c:companies)
   WHERE p.name = $name
   RETURN p.id AS person_id, c.name AS company
   ORDER BY company',
  params  := '{"name":"Alice"}'::jsonb,
  hydrate := true
);
```

`graph.gql()` 为每条 SQL 行返回一个 `jsonb` 对象。节点标签映射到已注册的表名，关系类型映射到已注册的边标签。支持的 GQL/openCypher 子集覆盖常见读取、有界路径、部分聚合，以及在启用 mutable overlay 时的窄范围映射写入。

### 运维注意事项

- 注册信息变更后，或源表发生所选同步模式未覆盖的变化后，需要用 `graph.build()` 重建。
- 动态边标签使用紧凑 ID；v0.1.7 最多支持 254 个面向用户的边标签。
- 加权最短路径需要数值型 `weight_column`；缺失或 NULL 的权重默认为 `1`。
- 重要 GUC 包括 `graph.max_nodes`、`graph.max_frontier`、`graph.memory_limit_mb`、`graph.query_freshness`、`graph.default_projection_mode` 和 `graph.mutable_enabled`。
- 映射式 GQL 写入需要 `graph.default_projection_mode = 'mutable_overlay'` 且 `graph.mutable_enabled = on`。
