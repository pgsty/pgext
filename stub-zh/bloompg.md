## 用法

来源：

- [官方 README](https://pgxn.org/dist/bloompg/0.1.2/README.html)
- [扩展控制文件](https://api.pgxn.org/src/bloompg/bloompg-0.1.2/bloompg.control)
- [版本化安装 SQL](https://api.pgxn.org/src/bloompg/bloompg-0.1.2/sql/bloompg--0.1.2.sql)
- [配置参考](https://pgxn.org/dist/bloompg/0.1.2/docs/CONFIGURATION.html)
- [安全策略](https://pgxn.org/dist/bloompg/0.1.2/SECURITY.html)

`bloompg` 0.1.2 是一个面向 PostgreSQL 18 受控分析负载的规划器 hook 扩展。它在规划阶段执行选定的关系输入，通过严格且可哈希连接的等值连接传播精确位图或布隆过滤器，物化缩减后的输入，再把更小的连接问题交给 PostgreSQL 进行第二次规划。现有 SQL 无需改写。当选择性较高的表可以排除其他连接表中的大量行时，它最有价值。

### 核心流程

安装库后，把 `bloompg` 加入 `shared_preload_libraries` 并重启 PostgreSQL，使每个新后端都能安装它的规划器 hook：

```conf
shared_preload_libraries = 'bloompg'
```

重启后，在每个需要使用它的数据库中以超级用户创建扩展，然后照常执行分析 SQL：

```sql
CREATE EXTENSION bloompg;
SELECT bloompg_version();

SET bloompg.enable = on;
SET bloompg.transfer_progress_metric = 'ndv';
SET bloompg.transfer_workers = 8;

SELECT count(*)
FROM fact
JOIN dimension USING (dimension_id)
WHERE dimension.region = 'APAC';
```

`bloompg.enable` 默认为 `on`。推荐的 `bloompg.transfer_progress_metric` 是 `ndv`，在适用时利用精确的去重键基数避免冗余传递；`rows` 保留基于行基数和过滤器谱系的传播。`bloompg.transfer_workers` 限制一次可用传递扫描的并行工作进程数，零值则让这些扫描串行执行。

### 函数与设置

- `bloompg_version()` 以 `text` 返回已安装的扩展版本。
- `bloompg_last_trace()` 以 `text` 返回当前后端最近一次传递跟踪。
- `bloompg_last_profile()` 以 `jsonb` 返回当前后端最近一次结构化画像。
- `bloompg.materialization_memory` 限制单条查询在优化阶段物化和并行 DSM 副本使用的内存；根据检测内存计算出的默认值最高为 `2GB`。
- `bloompg.sample_mode` 在持久化的 `prepared` 样本与查询局部的 `instant` 样本之间选择。`bloompg.sample_cache_dir` 指定预备样本缓存目录，并且只能由超级用户设置。
- `bloompg.profile` 收集结构化诊断信息，`bloompg.profile_log` 控制是否把已完成的画像写入 PostgreSQL 日志。
- `bloompg.index_transfer` 是一条实验性的、显式启用的精确键 B-tree 物化路径，默认值为 `off`。

仅在待调查的会话中启用诊断，执行查询，然后查看当前后端的最近结果：

```sql
SET bloompg.profile = on;
SET bloompg.profile_log = off;

SELECT count(*)
FROM fact
JOIN dimension USING (dimension_id)
WHERE dimension.region = 'APAC';

SELECT jsonb_pretty(bloompg_last_profile());
SELECT bloompg_last_trace();
```

### 范围、回退与安全

0.1.2 仅支持 PostgreSQL 18，上游在 Linux 上使用 PostgreSQL 18.4 开发和测试。控制文件设置 `superuser = true` 和 `trusted = false`，所以只能由超级用户创建；该扩展可重定位。修改 `shared_preload_libraries` 需要重启服务器。

受支持的核心范围是使用严格、可哈希连接等值条件的只读查询。修改语句、行锁、行级安全、易变表达式和不支持的查询范围会绕过 BloomPG，并保留 PostgreSQL 的规划方式。预备样本只影响调度和估算；用于排除行的过滤器来自当前语句快照下的精确扫描。

BloomPG 在规划期间执行真实扫描和内存物化。其单条查询预算默认为检测到的主机或 cgroup 内存的八分之一，最高为 2 GB。若下一次分配无法满足预算，它会放弃优化工作并执行原始 PostgreSQL 计划，而不会溢写到磁盘。应按分析负载的峰值并发调整 `bloompg.materialization_memory`，在有代表性的故障下验证原生回退，并让 `bloompg.index_transfer` 保持禁用，除非已经评估过其实验路径。

上游将 BloomPG 定位为分析研究扩展，不建议未经审查便默认部署在多租户或安全敏感集群中。生产使用前应测试结果等价性、规划延迟、内存压力、并行工作进程需求、hook 交互与日志处理；可以用 `bloompg.enable` 作为会话级回退开关。
