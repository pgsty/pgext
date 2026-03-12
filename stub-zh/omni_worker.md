


## 用法

> [omni_worker: 通用工作进程池](https://docs.omnigres.org/omni_worker/intro/)

`omni_worker` 扩展提供通用的 PostgreSQL 工作进程池，在独立的后端上下文中执行任意工作负载。

### SQL 处理器

内置的 SQL 处理器在后台工作进程中调度 SQL 语句执行。

**即发即忘：**

```sql
SELECT omni_worker.sql('INSERT INTO logs VALUES (now(), $1)');
```

调度后立即返回 `null`。

**带超时：**

```sql
SELECT omni_worker.sql('VACUUM ANALYZE my_table', wait_ms => 1000);
```

返回值：
- `true` -- 执行成功完成
- `false` -- 执行失败
- `null` -- 执行未在超时时间内完成

SQL 处理器默认启用，可通过从 `handlers` 表中移除其条目来禁用。
