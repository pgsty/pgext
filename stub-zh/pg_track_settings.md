


## 用法

> [pg_track_settings: 追踪 PostgreSQL 配置变更](https://github.com/rjuju/pg_track_settings)

pg_track_settings 记录 PostgreSQL 配置随时间的变更，追踪全局设置（`pg_settings`）和按角色/数据库的覆盖设置（`pg_db_role_setting`）。

### 采集快照

定期调用（通过 cron 或 PoWA）以捕获当前设置：

```sql
SELECT pg_track_settings_snapshot();
```

### 查看某一时间点的设置

```sql
-- 指定时间的所有设置
SELECT * FROM pg_track_settings('2024-01-15 10:00:00');

-- 指定时间的所有覆盖（按角色/数据库）设置
SELECT * FROM pg_track_db_role_settings('2024-01-15 10:00:00');
```

### 比较两个时间点的设置

```sql
-- 查看最近一小时内的变更
SELECT * FROM pg_track_settings_diff(now() - interval '1 hour', now());

-- 比较覆盖设置
SELECT * FROM pg_track_db_role_settings_diff(now() - interval '1 hour', now());
```

### 查看变更历史

```sql
-- 特定设置的历史
SELECT * FROM pg_track_settings_log('work_mem');

-- 覆盖设置的历史
SELECT * FROM pg_track_db_role_settings_log('work_mem');

-- PostgreSQL 重启历史
SELECT * FROM pg_reboot;
```

### 重置历史

```sql
SELECT pg_track_settings_reset();
```

### 函数摘要

| 函数 | 描述 |
|----------|-------------|
| `pg_track_settings_snapshot()` | 捕获当前设置 |
| `pg_track_settings(timestamptz)` | 指定时间的所有设置 |
| `pg_track_settings_diff(timestamptz, timestamptz)` | 两个时间点之间变更的设置 |
| `pg_track_settings_log(text)` | 特定设置的历史 |
| `pg_track_db_role_settings(timestamptz)` | 指定时间的覆盖设置 |
| `pg_track_db_role_settings_diff(timestamptz, timestamptz)` | 覆盖设置的变更 |
| `pg_track_db_role_settings_log(text)` | 特定覆盖设置的历史 |
| `pg_track_settings_reset()` | 清除所有历史 |
