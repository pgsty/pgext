


## 用法

> [pg_wait_sampling: 基于采样的等待事件统计](https://github.com/postgrespro/pg_wait_sampling)

pg_wait_sampling 通过定期采样每个后端正在等待的事件来收集等待事件统计。它提供近期历史环形缓冲区和累积统计两种视图。

### 视图

**当前等待事件：**

```sql
SELECT * FROM pg_wait_sampling_current;
```

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `pid` | int4 | 进程 ID |
| `event_type` | text | 等待事件类型 |
| `event` | text | 等待事件名称 |
| `queryid` | int8 | 查询 ID |

**等待事件历史**（近期采样的环形缓冲区）：

```sql
SELECT * FROM pg_wait_sampling_history;
```

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `pid` | int4 | 进程 ID |
| `ts` | timestamptz | 采样时间戳 |
| `event_type` | text | 等待事件类型 |
| `event` | text | 等待事件名称 |
| `queryid` | int8 | 查询 ID |

**等待事件剖析**（累积计数）：

```sql
SELECT * FROM pg_wait_sampling_profile;
```

| 列名 | 类型 | 描述 |
|--------|------|-------------|
| `pid` | int4 | 进程 ID |
| `event_type` | text | 等待事件类型 |
| `event` | text | 等待事件名称 |
| `queryid` | int8 | 查询 ID |
| `count` | text | 采样计数 |

**获取特定进程的当前等待事件：**

```sql
SELECT * FROM pg_wait_sampling_get_current(12345);
```

### 重置剖析

```sql
SELECT pg_wait_sampling_reset_profile();
```

### 配置

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_wait_sampling.history_size` | 5000 | 环形缓冲区大小 |
| `pg_wait_sampling.history_period` | 10 | 历史采样周期（毫秒） |
| `pg_wait_sampling.profile_period` | 10 | 剖析采样周期（毫秒） |
| `pg_wait_sampling.profile_pid` | `true` | 按进程收集剖析 |
| `pg_wait_sampling.profile_queries` | `top` | 按查询剖析：`none`、`top`、`all` |
| `pg_wait_sampling.sample_cpu` | `true` | 采样 CPU 上的后端（事件列为 NULL） |

### 示例：热点等待事件

```sql
SELECT event_type, event, count
FROM pg_wait_sampling_profile
ORDER BY count DESC
LIMIT 10;
```
