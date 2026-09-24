


## 用法

来源：

- [pgGraph v1.2.1 README](https://github.com/Evokoa/pgGraph/blob/v1.2.1/README.md)
- [v1.2.1 发行说明](https://github.com/Evokoa/pgGraph/releases/tag/v1.2.1)
- [SQL API 参考](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/api-reference.mdx)
- [Schema 注册](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/schema-registration.mdx)
- [管理与安全](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/administration-and-security.mdx)
- [故障排查](https://github.com/Evokoa/pgGraph/blob/v1.2.1/docs/user_guide/troubleshooting.mdx)
- [扩展控制文件](https://github.com/Evokoa/pgGraph/blob/v1.2.1/graph/graph.control)
- [v1.2.0 到 v1.2.1 升级 SQL](https://github.com/Evokoa/pgGraph/blob/v1.2.1/graph/sql/graph--1.2.0--1.2.1.sql)

`pggraph` 是包名与 PGXN 发行名，但安装到 PostgreSQL 中的扩展名是 `graph`。它从普通 PostgreSQL 表构建派生图产物，并以源表作为事实来源，通过 `graph` schema 提供图搜索、遍历、最短路径、GQL 风格读取，以及部分映射式写入。

版本 1.2.1 支持 PostgreSQL 14-18、命名图、按图隔离的授权与配额、持久同步、有界遍历与分析、维护任务，以及选定的 GQL 读写 profile。它还通过有边界的开放词汇类型字典，移除了历史上 254 个关系标签的上限。它不声明支持完整 ISO GQL、完整 openCypher 或公开 SQL/PGQ `GRAPH_TABLE` surface。标准 PostgreSQL SQLSTATE 会与稳定的 `PGxxx` detail 配对，供应用诊断。

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

除非使用显式的 `*_to_graph` / `*_from_graph` 辅助函数，注册操作会作用于当前选中的图。节点标识符必须匹配主键，或匹配唯一的 `NOT NULL` 索引。`columns` 控制可搜索和可被 GQL 读取的属性；遍历过滤下推需要通过 `graph.add_filter_column()` 单独注册。边表和 junction table 形式的关系也受支持，`label_column` 可在文档规定的公开上限内提供动态边标签。

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

### 关系类型与注册恢复

版本 1.2.0 允许单个图包含最多 1,000,000 种关系类型。每个 UTF-8 标签上限为 1,024 字节，字典累计上限为 256 MiB，关系类型过滤数组在分配前受 4,096 项和 4 MiB 限制。`graph.status()` 只会按稳定 ID 顺序预览前 64 个已提交类型；要分页读取完整有效字典，请使用：

```sql
SELECT type_id, label
FROM graph.edge_types(after_type_id := 0, max_rows := 1000);
```

通过触发器同步提交的动态标签会由 `graph.apply_sync()` 加入字典，无需全量重建。不存在的关系类型会返回无匹配，而端点映射存在歧义时会失败关闭。

无参形式会删除选中图的派生引擎与产物，但保留注册。只在表重建或逻辑恢复导致 relation identity 过期时，才使用 boolean 重载：

```sql
SELECT graph.reset();

-- This also clears table, edge, and filter registrations for the selected graph.
SELECT graph.reset(true);
-- Reapply reviewed graph.add_table(...), graph.add_edge(...), and
-- graph.add_filter_column(...) calls before rebuilding.
SELECT * FROM graph.build();
```

两种形式都不会修改 PostgreSQL 源表或其他命名图。`graph.reset(true)` 会清空选中图的注册目录，因此使用前必须保留已审核的注册 SQL。

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

`graph.gql()` 为每条 SQL 行返回一个 `jsonb` 对象。节点标签映射到已注册表名，关系类型映射到已注册边标签。受支持的可变 GQL profile 包含已注册关系创建：映射式写入仍通过 PostgreSQL 源表 DML 执行，源表依然是权威数据。未支持的 openCypher 或 SQL/PGQ 形状会返回显式能力错误，而不是部分执行。

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

### 升级到 1.2.1

安装匹配的 1.2.1 软件包前先备份 PostgreSQL，再更新各数据库，并重建每个已注册图。命名图也必须逐一选中后重建：

```sql
ALTER EXTENSION graph UPDATE TO '1.2.1';
SELECT * FROM graph.build();
SELECT extversion FROM pg_extension WHERE extname = 'graph';
SELECT * FROM graph.status();
```

此次重建是必需步骤，用于建立数据库级文件根目录、目录来源校验和事务化代际状态；旧产物不会自动接管，源表和注册信息仍是权威来源。逻辑恢复后也需重建。更新保留既有 SQL 对象标识、所有者与显式授权。不支持就地替换二进制降级；回滚应恢复升级前备份与匹配的旧软件包，再重建派生图状态。

新增的 `graph.sync_retention()` 是所选图的管理员诊断接口，报告同步日志保留与清理阻碍，但不会执行清理或获取写锁。版本 1.2.1 还修复了同步回放、事务回滚、数据库间产物隔离与部分 GQL 查询结果问题。

### 注意事项

- 源表仍是事实来源。图产物、projection 文件、同步状态和运行时引擎都是派生状态，可从源表重建。
- 注册信息变化后需要运行 `graph.build()` 或图级构建辅助函数；依赖增量 projection 时，应使用 sync/maintenance API。
- `graph._graphs`、授权、配额、任务、同步日志、projection 元数据等内部表是实现细节，应用代码应使用公开 SQL 函数。
- 版本 1.2.0 源码构建使用 Rust 1.96 与 `cargo-pgrx` 0.19.1。上游支持 PostgreSQL 14 到 18，默认 release gate 目标是 PostgreSQL 17。
