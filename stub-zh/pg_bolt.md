## 用法

来源：

- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/README.md)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/pg_bolt.control)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/sql/pg_bolt--0.1.0.sql)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/Cargo.toml)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/LICENSE)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/docs/SQL_API.md)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/docs/CONFIGURATION.md)
- [Official source](https://github.com/craton-co/pg-bolt/blob/1481756dba79240bdc5fd7811a3e2b2b0bfe07fc/docs/MVCC_AND_CACHE.md)

`pg_bolt` 0.1.0 在显式缓存的 Arrow 表上提供实验性 GPU 分析查询。源码通过 pgrx 0.19 声明 PostgreSQL 15–18 目标。上游 README 报告了主机及 CUDA 测试，但说明作者的 Windows 主机未执行真实 PostgreSQL pgrx 测试；兼容范围应理解为源码目标，不能视为生产环境已经验证。

### 加载与查询快照

默认后台工作进程模式需要预加载和服务器重启，请保留其他预加载项。以超级用户创建扩展；它建立不可重定位的 `bolt` 模式。GPU 执行需要 NVIDIA 驱动和 CUDA Toolkit 12 或更高版本。`cuda-stub` 用于没有 GPU 的测试。

对于已有且包含 region_id 列的表，可先禁用自动规划器卸载，仅使用显式查询：

```ini
shared_preload_libraries = 'pg_bolt'
pg_bolt.max_vram_mb = 4096
```

```sql
CREATE EXTENSION pg_bolt;
SET pg_bolt.enable_custom_scan = off;
SELECT * FROM bolt.load_table('public.sales', 'sales');
SELECT bolt.explain_tier('SELECT region_id, count(*) FROM sales GROUP BY region_id');
SELECT bolt.query_json('SELECT region_id, count(*) FROM sales GROUP BY region_id');
SELECT * FROM bolt.tables();
SELECT bolt.status();
```

`bolt.load_table` 替换指定缓存快照，`bolt.cache_table` 复用未标记脏状态的缓存，`bolt.sync_table` 刷新缓存。`bolt.evict_table` 删除缓存表；`bolt.reset_vram` 释放运行时及缓存状态。`bolt.query` 流式返回 JSONB 行，`bolt.query_json` 返回单个 JSONB 数组。`bolt.explain_sql` 展示完整 Bolt 计划。

### 配置与边界

`pg_bolt.enabled` 控制执行，`pg_bolt.enable_custom_scan` 启用规划器集成。`pg_bolt.min_offload_rows` 默认 50000。`pg_bolt.max_vram_mb` 默认为零，表示没有配置预算；正值预算使用 LRU 淘汰，并拒绝单表超过预算的情况。`pg_bolt.statement_timeout_ms` 提供协作式取消。`pg_bolt.use_bgworker` 默认开启，依赖预加载初始化的工作进程；单后端开发模式另有边界。

缓存的 Arrow 数据是快照。上游设计跟踪 DML、DDL 和事务失效，但仍将跨后端目录及快照同步列为后续工作。向应用开放前，应与普通 PostgreSQL 查询对照验证新鲜度、事务可见性、权限、RLS 和回退行为。安装 SQL 没有撤销 PUBLIC 执行权限，管理员应按角色需要审查并限制这些 SQL 入口。

扩展源码采用 Apache-2.0，而链接的 craton-bolt 依赖采用 AGPL-3.0-only；上游对组合二进制另有许可说明。CUDA-stub 测试和性能数字不能证明普遍加速或生产安全性。
