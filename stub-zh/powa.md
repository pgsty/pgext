


## 用法

> [powa: PostgreSQL 工作负载分析器](https://github.com/powa-team/powa)

PoWA 是一个工作负载分析框架，收集性能统计信息并通过 Web UI 提供实时图表。核心扩展（`powa-archivist`）从多个统计扩展中采集快照。

### 架构

PoWA 由几个组件组成：

- **powa-archivist**：收集和存储统计信息的 PostgreSQL 扩展
- **powa-web**：用于可视化性能数据的 Web UI
- **pg_stat_statements**：必需，提供查询统计
- **pg_qualstats**：可选，提供谓词统计和索引建议
- **pg_stat_kcache**：可选，提供操作系统级指标（CPU、I/O）
- **pg_wait_sampling**：可选，提供等待事件采样

### 采集快照

PoWA 定期从所有已注册的扩展中采集统计快照：

```sql
-- 手动采集快照
SELECT powa_take_snapshot();

-- 检查快照状态
SELECT * FROM powa_snapshot_metas;
```

### 配置

```sql
-- 注册统计扩展（在 CREATE EXTENSION 时自动完成）
SELECT powa_register_server(hostname => 'localhost', port => 5432);

-- 配置快照间隔和保留策略
ALTER EXTENSION powa UPDATE;
```

### 关键表和视图

```sql
-- 查看收集的查询统计
SELECT * FROM powa_statements;

-- 查看快照历史
SELECT * FROM powa_snapshot_metas;
```

### Web UI

PoWA Web 界面（单独安装）提供：

- 实时查询性能仪表板
- 按查询下钻及计划详情
- 基于 `pg_qualstats` 的索引建议
- 等待事件分析
- 系统资源使用图表

文档：[powa.readthedocs.io](https://powa.readthedocs.io/)
