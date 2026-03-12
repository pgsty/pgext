


## 用法

> [pg_track_optimizer: 通过基数估计误差检测次优查询计划](https://github.com/danolivo/pg_track_optimizer)

pg_track_optimizer 通过比较规划器预测与实际执行结果，自动检测具有较差基数估计的查询。它使用对数刻度计算多种误差指标。

### 启用追踪

```sql
-- 生产环境仅追踪有问题的查询
SET pg_track_optimizer.mode = 'normal';

-- 调试时追踪所有查询
SET pg_track_optimizer.mode = 'forced';

-- 当误差超过阈值时记录 EXPLAIN
SET pg_track_optimizer.log_min_error = 2.0;
```

### 查看追踪的查询

```sql
SELECT queryid, query,
       avg_avg, avg_min, avg_max,
       rms_avg, rms_min, rms_max,
       time_avg, blks_avg, nexecs
FROM pg_track_optimizer
ORDER BY avg_avg DESC
LIMIT 10;

-- 直接使用 RStats 类型
SELECT queryid, query,
       wca_error -> 'mean' AS avg_wca_error,
       blks_accessed -> 'mean' AS avg_blocks
FROM pg_track_optimizer()
WHERE blks_accessed -> 'mean' > 1000
ORDER BY wca_error -> 'mean' DESC;
```

### 误差指标

| 指标 | 描述 |
|--------|-------------|
| `avg_error` | 计划节点对数刻度误差的简单平均值 |
| `rms_error` | 均方根误差，强调大误差 |
| `twa_error` | 时间加权平均，突出慢节点 |
| `wca_error` | 代价加权平均，突出高代价节点 |
| `f_join_filter` | JOIN 过滤开销因子 |
| `f_scan_filter` | 扫描过滤开销因子 |

### 管理统计信息

```sql
-- 将统计信息保存到磁盘
SELECT pg_track_optimizer_flush();

-- 清除所有追踪的统计信息
SELECT pg_track_optimizer_reset();

-- 检查扩展状态
SELECT * FROM pg_track_optimizer_status;
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_track_optimizer.mode` | `disabled` | `disabled`、`normal`、`forced` |
| `pg_track_optimizer.log_min_error` | (无) | 触发 EXPLAIN 日志的误差阈值 |
| `pg_track_optimizer.hash_mem` | (默认) | 共享内存限制（KB） |
| `pg_track_optimizer.auto_flush` | `on` | 后端关闭时自动保存统计 |
