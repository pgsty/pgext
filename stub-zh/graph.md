


## 用法

来源：

- [pgGraph v0.1.8 README](https://github.com/Evokoa/pgGraph/blob/v0.1.8/README.md)
- [v0.1.8 release notes](https://github.com/Evokoa/pgGraph/blob/v0.1.8/docs/release-notes.mdx)
- [SQL API Reference](https://github.com/Evokoa/pgGraph/blob/v0.1.8/docs/user_guide/api-reference.mdx)
- [Schema Registration](https://github.com/Evokoa/pgGraph/blob/v0.1.8/docs/user_guide/schema-registration.mdx)
- [Administration and Security](https://github.com/Evokoa/pgGraph/blob/v0.1.8/docs/user_guide/administration-and-security.mdx)

`pggraph` 是包名与 PGXN 发行名，但安装到 PostgreSQL 中的扩展名是 `graph`。它从普通 PostgreSQL 表构建派生图产物，并以源表作为事实来源，通过 `graph` schema 提供图搜索、遍历、最短路径、GQL 风格读取，以及部分映射式写入。

v0.1.8 增加了命名图管理、按图隔离的 catalog、图级授权与配额、托管维护任务、GQL 关系创建，以及更明确的 openCypher / SQL/PGQ 预览能力边界。上游仍把 pgGraph 标记为早期 alpha；应先在可丢弃或开发数据库中试用，并始终把图产物视为可从源表重建的派生状态。

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
SELECT * FROM graph.build();

SELECT node_count, edge_count, edge_types
FROM graph.status();
```

`graph.auto_discover('public')` 会扫描该 schema 中的主键与外键，注册发现到的源表和边，并为 `graph.build()` 准备图。生产 schema 建议显式注册，确保标签、搜索列、过滤列、权重、租户行为和图身份都经过设计。

### 手工注册与命名图

```sql
SELECT graph.create_graph('customer_360', namespace := 'analytics');
SELECT graph.set_current_graph('customer_360', namespace := 'analytics');

SELECT graph.add_table(
  table_name := 'public.people'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name'],
  tenant_column := NULL
);

SELECT graph.add_table_to_graph(
  graph_name := 'customer_360',
  table_name := 'public.companies'::regclass,
  id_column  := 'id',
  columns    := ARRAY['name'],
  graph_namespace := 'analytics'
);

SELECT graph.add_edge_to_graph(
  graph_name := 'customer_360',
  from_table := 'public.people'::regclass,
  from_column := 'company_id',
  to_table := 'public.companies'::regclass,
  to_column := 'id',
  label := 'works_at',
  bidirectional := true,
  graph_namespace := 'analytics'
);

SELECT * FROM graph.build_graph('customer_360', graph_namespace := 'analytics');
```

除非使用显式的 `*_to_graph` / `*_from_graph` 辅助函数，注册操作会作用于当前选中的图。节点标识符必须匹配主键，或匹配唯一的 `NOT NULL` 索引。`columns` 控制可搜索和可被 GQL 读取的属性；遍历过滤下推需要通过 `graph.add_filter_column()` 单独注册。边表和 junction table 形式的关系也受支持，`label_column` 可提供动态边标签，但 v0.1.8 对用户可见标签数量有上限。

### 搜索、遍历与路径

```sql
SELECT node_table_name, node_id, node
FROM graph.search(
  property_key   := 'name',
  property_value := 'Alice',
  table_filter   := 'public.people'::regclass,
  mode           := 'exact',
  hydrate        := true
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

`hydrate := false` 返回紧凑的图坐标。启用 hydration 后，源表行的可见性仍由 PostgreSQL ACL 与 RLS 控制。过期坐标会失败关闭，而不会伪造行。

### GQL 查询与关系写入

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

`graph.gql()` 为每条 SQL 行返回一个 `jsonb` 对象。节点标签映射到已注册表名，关系类型映射到已注册边标签。v0.1.8 扩展了可变 GQL 写入面，支持创建已注册关系：映射式写入仍通过 PostgreSQL 源表 DML 执行，源表依然是权威数据。未支持的 openCypher 或 SQL/PGQ 形状会返回更清晰的能力边界错误，而不是部分执行。

### 管理与运维

```sql
GRANT USAGE, CREATE ON SCHEMA graph TO graph_admin;

SELECT * FROM graph.grant_graph('customer_360', 'app_reader', 'read', namespace := 'analytics');
SELECT * FROM graph.set_graph_quota('owner', 'max_named_graphs', 25, scope_key := 'app_owner');
SELECT * FROM graph.select_graph('customer_360', namespace := 'analytics');
SELECT * FROM graph.add_sync_policy('customer_360', schedule_interval_secs := 300, graph_namespace := 'analytics');
SELECT * FROM graph.run_due_jobs();
SELECT * FROM graph.projection_status();
```

图管理覆盖 catalog 变更、构建、同步回放、维护、配额、运行时图加载和全局分析。命名图权限包括 `read`、`write`、`build`、`admin`，但图级 `read` 本身不够：hydrated 读取仍需要源表 `SELECT` 权限。选中图的 tenant 也会默认约束遍历、搜索、GQL 与 Cypher 调用，除非显式传入匹配的 tenant。

### 注意事项

- 源表仍是事实来源。图产物、projection 文件、同步状态和运行时引擎都是派生状态，可从源表重建。
- 注册信息变化后需要运行 `graph.build()` 或图级构建辅助函数；依赖增量 projection 时，应使用 sync/maintenance API。
- `graph._graphs`、授权、配额、任务、同步日志、projection 元数据等内部表是实现细节，应用代码应使用公开 SQL 函数。
- v0.1.8 将源码构建基线提升到 Rust 1.96 和 `cargo-pgrx` 0.19.1。上游仍支持 PostgreSQL 14 到 18，默认 release gate 目标是 PostgreSQL 17。
