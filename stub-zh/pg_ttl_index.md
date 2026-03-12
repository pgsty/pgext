


## 用法

> [pg_ttl_index: 基于 TTL 索引的自动数据过期](https://github.com/ibrahimkarimeddin/postgres-extensions-pg_ttl)

`pg_ttl_index` 通过将 TTL（生存时间）与时间戳列关联来实现自动数据过期。后台工作进程会定期删除时间戳超过配置过期间隔的行。

### 快速开始

```sql
-- 启动后台工作进程
SELECT ttl_start_worker();

-- 创建带时间戳列的表
CREATE TABLE user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER,
    session_data JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 1 小时（3600 秒）后过期
SELECT ttl_create_index('user_sessions', 'created_at', 3600);
```

### 函数

| 函数 | 描述 |
|------|------|
| `ttl_start_worker()` | 启动自动清理的后台工作进程 |
| `ttl_worker_status()` | 检查工作进程是否正在运行 |
| `ttl_runner()` | 手动触发清理 |
| `ttl_create_index(table, column, expire_seconds [, batch_size])` | 配置 TTL 过期策略 |
| `ttl_drop_index(table, column)` | 移除 TTL 配置 |
| `ttl_summary()` | 列出所有 TTL 索引及统计信息 |

### 示例

24 小时过期的会话管理：

```sql
SELECT ttl_create_index('sessions', 'created_at', 86400, 5000);
```

7 天日志保留：

```sql
SELECT ttl_create_index('app_logs', 'logged_at', 604800);
```

自定义过期列的缓存条目（0 表示该列本身存储绝对过期时间）：

```sql
SELECT ttl_create_index('cache_entries', 'expires_at', 0);
```

### 监控

```sql
SELECT * FROM ttl_summary();
```

暂停特定表的清理：

```sql
UPDATE ttl_index_table SET active = false WHERE table_name = 'user_sessions';
```

### 配置

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `pg_ttl_index.naptime` | 清理间隔（秒） | 60 |
| `pg_ttl_index.enabled` | 全局启用/禁用工作进程 | on |

```sql
ALTER SYSTEM SET pg_ttl_index.naptime = 30;
SELECT pg_reload_conf();
```
