## 用法

来源：

- [PoWA Archivist 5.2.0 README](https://github.com/powa-team/powa-archivist/blob/REL_5_2_0/README.md)
- [PoWA Archivist 5.2.0 control file](https://github.com/powa-team/powa-archivist/blob/REL_5_2_0/powa.control)
- [PoWA Archivist 5.2.0 release](https://github.com/powa-team/powa-archivist/releases/tag/REL_5_2_0)
- [PoWA documentation](https://powa.readthedocs.io/en/latest/components/powa-archivist/index.html)

PoWA 是一个工作负载分析框架，用于收集性能统计数据并通过 Web UI 提供实时图表。核心扩展 (`powa-archivist`) 会从多个统计扩展中快照统计信息。

### 架构

PoWA 包含以下几个组件：

- **powa-archivist**：PostgreSQL 扩展，用于收集和存储统计信息
- **powa-web**：一个 Web UI，用于可视化性能数据
- **pg_stat_statements**：必需的扩展，提供查询统计信息
- **pg_qualstats**（可选）：提供谓词统计信息和索引建议
- **pg_stat_kcache**（可选）：提供操作系统级指标（CPU、I/O）
- **pg_wait_sampling**（可选）：提供等待事件采样

在仓库数据库中启用该扩展。其控制文件需要 PL/pgSQL，`pg_stat_statements` 和 `btree_gist`：

```sql
CREATE EXTENSION pg_stat_statements;
CREATE EXTENSION btree_gist;
CREATE EXTENSION powa CASCADE;
```

`pg_stat_statements` 本身必须在 PostgreSQL 启动前配置在 `shared_preload_libraries` 中。

### 快照

PoWA 定期从所有注册的扩展中快照统计信息：

```sql
-- Manual snapshot
SELECT powa_take_snapshot();

-- Check snapshot status
SELECT * FROM powa_snapshot_metas;
```

### 配置

```sql
-- Register stat extensions (done automatically on CREATE EXTENSION)
SELECT powa_register_server(hostname => 'localhost', port => 5432);

-- Configure snapshot interval and retention
ALTER EXTENSION powa UPDATE;
```

### 关键表和视图

```sql
-- View collected query statistics
SELECT * FROM powa_statements;

-- View snapshot history
SELECT * FROM powa_snapshot_metas;
```

### Web UI

PoWA 的 Web 界面（单独安装）提供以下功能：

- 实时查询性能仪表板
- 查询钻取，带有计划详细信息
- 基于 `pg_qualstats` 的索引建议
- 等待事件分析
- 系统资源使用情况图表

文档：[powa.readthedocs.io](https://powa.readthedocs.io/)

### 版本和部署注意事项

- PoWA Archivist 5.2.0 添加了对 PostgreSQL 19 的支持，包括新的 `pg_stat_recovery` 和 `pg_stat_lock` 统计视图的收集器。核心快照工作流程与 5.1 系列保持兼容。
- `powa` 是数据库扩展；`powa-web` 是一个单独的可视化服务，而 `powa-collector` 用于远程收集架构。仅安装扩展不会部署 UI。
- 保留策略和快照频率直接影响仓库的增长。监控 PoWA 仓库数据库并调整保留策略以适应所需的数据库数量、查询次数以及启用的可选模块的数量。
