


## 用法

> [pg_profile: PostgreSQL 历史性能分析工具](https://github.com/zubkov-andrei/pg_profile)

pg_profile 定期采集 PostgreSQL 统计信息快照，并生成详细的历史性能报告。它依赖 `pg_stat_statements`，可选使用 `pg_stat_kcache` 和 `pg_wait_sampling` 获取额外指标。

### 采集快照

需要定期采集快照（例如通过 cron）。每次快照会捕获当前的统计信息状态：

```sql
SELECT profile.take_sample();
```

### 生成报告

通过指定两个快照 ID 来生成区间性能分析报告：

```sql
-- 生成快照 1 到 2 之间的常规报告
SELECT profile.get_report(1, 2);

-- 生成比较两个区间的差异报告
SELECT profile.get_diffreport(1, 2, 3, 4);
```

### 管理服务器

pg_profile 可以从远程集群采集统计信息：

```sql
-- 定义远程服务器
SELECT profile.create_server('remote', 'host=remote_host dbname=postgres');

-- 列出已定义的服务器
SELECT * FROM profile.show_servers();

-- 启用/禁用服务器
SELECT profile.enable_server('remote');
SELECT profile.disable_server('remote');
```

### 基线

基线可以保护快照范围不被自动清理：

```sql
-- 创建基线，保留快照 10 到 20
SELECT profile.create_baseline('incident_2024', 10, 20);

-- 列出基线
SELECT * FROM profile.show_baselines();

-- 删除基线
SELECT profile.drop_baseline('incident_2024');
```

### 保留策略

控制快照的保留时长：

```sql
-- 将本地服务器的保留期设置为 7 天
SELECT profile.set_server_max_sample_age('local', 7);
```

### 快照信息

```sql
-- 查看可用快照
SELECT * FROM profile.show_samples();

-- 查看快照采集耗时（需要 pg_profile.track_sample_timings = on）
SELECT * FROM v_sample_timings;
```

### 推荐配置

```
track_activities = on
track_counts = on
track_io_timing = on
track_wal_io_timing = on      # PG 14+
track_functions = all
```
