## 用法

来源：

- [固定提交的上游概述](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/README.md)
- [固定提交的 PostgreSQL 17 快速入门](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/docs/quickstart.md)
- [固定提交的查询支持参考](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/docs/query-support.md)
- [固定提交的限制说明](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/docs/limitations.md)
- [固定提交的扩展 control 文件](https://github.com/darthunix/pg_fusion/blob/08ad73af18c909bc50b95114c3e4316ffe01a3ed/pg/extension/pg_fusion.control)

pg_fusion 24.0.0 版是实验性的分析执行扩展。PostgreSQL 仍负责解析、catalog、snapshot、MVCC visibility、heap scan、TOAST 处理和面向客户端的 tuple；符合条件的扫描行通过共享内存 Arrow page 送入一个 background worker，由 Apache DataFusion 执行选定的 operator。

### 预加载与首个查询

该库会安装 hook、预留共享内存并启动 worker，因此必须在 postmaster 启动时加载：

```conf
shared_preload_libraries = 'pg_fusion'
```

重启后创建扩展，并且只在用于测试的会话中启用其严格执行路径：

```sql
CREATE EXTENSION IF NOT EXISTS pg_fusion;

CREATE TABLE fusion_demo AS
SELECT i AS id, i % 10 AS group_id, i::double precision AS value
FROM generate_series(1, 1000000) AS i;

ANALYZE fusion_demo;
SET pg_fusion.enable = on;

EXPLAIN
SELECT count(*), avg(value)
FROM fusion_demo
WHERE group_id >= 0;

SELECT count(*), avg(value)
FROM fusion_demo
WHERE group_id >= 0;
```

符合条件的 plan 会包含 Custom Scan (PgFusionScan)。当前 quick start 面向 pgrx PostgreSQL 17 开发集群。安装只允许超级用户执行。

### 实验性边界

受支持范围只是分析型 SELECT 的保守子集。启用运行时后，不支持的形状会抛出 pg_fusion planning error，而不是静默回退。当前排除项包括非 SELECT 工作、modifying CTE、bound/prepared parameter、SPI 和 PL/pgSQL 内部执行、table function、仍然存在的 subquery，以及不受支持的 PostgreSQL 类型或语义。

tuple-to-Arrow 转换和传输存在真实开销。宽 projection、返回大量原始行、filter pushdown 较弱或计算很轻的查询可能比原生 PostgreSQL 更慢。一个 worker 和固定共享内存池是集群共享资源；memory limit、spill directory、page count、control slot 与 runtime-filter 大小都需要容量规划。worker spill 使用操作系统临时存储，不受 PostgreSQL 临时文件控制项管理。

应将 pg_fusion 视为工程评估软件。需要分别在关闭和开启运行时的情况下对比结果与 EXPLAIN ANALYZE，验证文档列出的类型和 collation 限制，并从受控分析工作负载开始，而不是直接接入生产流量。
