## 用法

来源：

- [官方 README](https://github.com/LogicOcean/pgokf/blob/v0.1.13/README.md)
- [扩展控制文件](https://github.com/LogicOcean/pgokf/blob/v0.1.13/crates/extension/pgokf.control)
- [pgrx 清单](https://github.com/LogicOcean/pgokf/blob/v0.1.13/crates/extension/Cargo.toml)

`pgokf` 把 Open Knowledge Format bundle 物化为事务化 PostgreSQL catalog，提供全文检索、链接图、provenance、可选语义检索与 tenant 隔离。

### 启用与角色

0.1.13 版本为 PostgreSQL 15–19 提供 pgrx feature。安装准确大版本的构件，并以超级用户创建不可迁移扩展：

```sql
CREATE EXTENSION pgokf;
GRANT pgokf_writer TO app_user;
```

无需预加载。扩展创建 `pgokf_reader`、`pgokf_writer` 与 `pgokf_admin` 角色；应授予满足需求的最低权限角色。

### 注册并搜索 Bundle

文件系统路径由 PostgreSQL 服务端解析，必须是服务端可读的绝对路径。

```sql
SELECT *
FROM pgokf.register_bundle('/srv/okf/runbooks');

SELECT concept_id, title, rank
FROM pgokf.concept_search(
  'postgres failover',
  concept_type => 'runbook'
);

SELECT *
FROM pgokf.concept_neighbors('runbooks/database-failover', 2);
```

Bundle 仍是可移植的 source of truth；使用 `catalog_stats`、`health`、`search_index_status` 与 sync log 运维物化 projection。

### 可选搜索后端与边界

核心 lexical search 不依赖其他扩展。可选 `vector` 启用 semantic/hybrid function，`pg_search` 启用 BM25，`pg_cron` 启用定时刷新。缺失它们时会按文档降级或报错，不会阻止基础安装。

注册过程读取服务端文件，部分管理函数使用 `SECURITY DEFINER`；应限制 writer/admin membership 并审查路径约束。0.1.x 仍是 pre-1.0，切换 minor version 前必须阅读升级脚本与 changelog。配套 ingest/embed/MCP binary 在 PostgreSQL 外运行，不是核心扩展依赖。

