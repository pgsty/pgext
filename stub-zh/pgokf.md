## 用法

来源：

- [官方 README](https://github.com/LogicOcean/pgokf/blob/v0.4.0/README.md)
- [扩展控制文件](https://github.com/LogicOcean/pgokf/blob/v0.4.0/crates/extension/pgokf.control)
- [SQL API](https://github.com/LogicOcean/pgokf/blob/v0.4.0/docs/sql-api.md)
- [升级契约](https://github.com/LogicOcean/pgokf/blob/v0.4.0/docs/api-stability.md)
- [发布与升级说明](https://github.com/LogicOcean/pgokf/releases/tag/v0.4.0)
- [pgrx 清单](https://github.com/LogicOcean/pgokf/blob/v0.4.0/crates/extension/Cargo.toml)
- [0.5.0 release](https://github.com/LogicOcean/pgokf/releases/tag/v0.5.0)
- [0.5.0 changelog](https://github.com/LogicOcean/pgokf/blob/v0.5.0/CHANGELOG.md)
- [0.4 to 0.5 upgrade SQL](https://github.com/LogicOcean/pgokf/blob/v0.5.0/crates/extension/sql/pgokf--0.4.0--0.5.0.sql)
- [Confluence ingestion](https://github.com/LogicOcean/pgokf/blob/v0.5.0/docs/confluence-integration.md)

`pgokf` 把 Open Knowledge Format bundle 物化为事务化 PostgreSQL catalog，提供全文检索、链接图、provenance、可选语义检索与 tenant 隔离。

### 启用与角色

0.4.0 支持 PostgreSQL 15–18；PostgreSQL 19 RC 的测试属于独立的兼容性目标。安装准确大版本的构件，并以超级用户创建不可迁移扩展：

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

核心 lexical search 不依赖其他扩展。可选 `vector` 启用 semantic/hybrid function，`pg_textsearch` 或 `pg_search` 启用 BM25，`pg_cron` 启用定时刷新。缺失它们时会按文档降级或报错，不会阻止基础安装。

注册过程读取服务端文件，部分管理函数使用 `SECURITY DEFINER`；应限制 writer/admin membership 并审查路径约束。0.4.0 仍是 pre-1.0，切换 minor version 前必须阅读升级脚本与 changelog。配套 ingest/embed/MCP binary 在 PostgreSQL 外运行，不是核心扩展依赖。


### 目录代次、新鲜度与升级

目录代次、持久化变更事件队列和带类型的关系用于协调派生知识的更新。`concept_search_fresh` 提供新鲜度过滤；`mark_reconciling` 和 `mark_fresh` 为生产者提供带检查的状态转换。语义排序要求文件、模型和向量契约的来源匹配，且内容有效状态为新鲜。升级到 0.3 时，旧 bundle 会被标为过期，旧向量需要重新生成才能参与排序。

先备份数据库并安装 0.4.0 文件，再执行更新：

```sql
ALTER EXTENSION pgokf UPDATE TO '0.4.0';
SELECT extversion FROM pg_extension WHERE extname = 'pgokf';
```

所提供的升级链覆盖更早的 0.1、0.2、0.3 及开发版本。只有旧搜索函数重载的定义、授权和依赖与预期目录完全匹配时，才会修正被替代的重载；遇到意外结构会中止升级。不要强制删除依赖对象来绕过检查。已声明的旧版 Web 表列顺序会保留。SQL 目录与单独部署的 MCP/Web 配套程序应分别遵循各自升级契约。

### 0.5.0 配套工具与升级

0.5.0 不改变 0.4.0 的目录对象或数据，提供的升级脚本只推进扩展版本。安装匹配的服务器文件后，使用常规升级路径：

```sql
ALTER EXTENSION pgokf UPDATE TO '0.5.0';
```

独立的 `pgokf-confluence` 配套工具会预览完整的 Confluence Cloud 空间，只有使用 `--apply` 时才写入专用内容包。它记录生产者来源、使用持久化检查点，并要求通过 `--allow-deletions` 确认页面删除。它适合受众一致的内容空间，不会把 Confluence 页面限制复制成 PostgreSQL 权限。令牌应保存在私密文件或 `CONFLUENCE_API_TOKEN` 中。配套工具的 `--tls` / `OKF_PG_TLS` 现在至少强制 `sslmode=require`，服务器拒绝 TLS 时会断开，不再悄悄回退到明文；需要时应使用更严格的证书验证。这些配套工具变化不会让扩展后端直接执行 Confluence 采集或外部向量计算。
