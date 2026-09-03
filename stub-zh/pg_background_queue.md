## 用法

来源：

- [官方文档](https://github.com/hank-cp/pg_background_queue/blob/8e6f1c5429df50d5475a0d3943753e5adc125be9/README.md)
- [扩展控制文件](https://github.com/hank-cp/pg_background_queue/blob/8e6f1c5429df50d5475a0d3943753e5adc125be9/pg_background_queue.control)
- [官方仓库](https://github.com/hank-cp/pg_background_queue)

`pg_background_queue` 支持主题、并发限制与重试的后台 SQL 任务队列。

### 启用

将 `pg_background_queue` 合并到现有预加载列表，重启 PostgreSQL，然后在每个需要其 SQL 对象的数据库中创建 `pg_background_queue`：

```ini
shared_preload_libraries = 'pg_background_queue'
```

```sql
CREATE EXTENSION pg_background_queue;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Check active worker count
SELECT pg_background_queue_active_workers_count();

-- Ensure workers are running (❗️️SUGGESTING: trigger by pg_cron)
SELECT pg_background_queue_ensure_workers();

-- Calibrate worker count (❗️️SUGGESTING: trigger by pg_cron)
SELECT pg_background_queue_calibrate_workers_count();
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `pg_background_tasks` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `pg_background_enqueue` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_background_queue_ensure_workers` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_background_queue_calibrate_workers_count` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_background_queue_active_workers_count` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `pg_background_queue_task_state` | TYPE | 扩展创建的用户数据类型。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
- 预加载 `pg_background_queue` 会改变集群启动状态；配置与重启应和 `CREATE EXTENSION` 分开实施。
