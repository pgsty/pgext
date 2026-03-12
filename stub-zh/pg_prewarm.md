

## 用法

> [pg_prewarm: 预热关系数据](https://www.postgresql.org/docs/current/pgprewarm.html)

`pg_prewarm` 扩展提供将关系数据加载到操作系统缓冲区缓存或 PostgreSQL 缓冲区缓存的函数，减少后续查询的 I/O 延迟。

### 预热关系

```sql
-- 将整个表加载到 PostgreSQL 缓冲区缓存（默认模式）
SELECT pg_prewarm('my_table');

-- 使用特定模式加载
SELECT pg_prewarm('my_table', 'prefetch');  -- 异步 OS 预取
SELECT pg_prewarm('my_table', 'read');      -- 同步读入 OS 缓存
SELECT pg_prewarm('my_table', 'buffer');    -- 加载到 PG 缓冲区缓存

-- 加载特定块范围
SELECT pg_prewarm('my_table', 'buffer', 'main', 0, 999);

-- 预热索引
SELECT pg_prewarm('my_table_pkey');
```

### 函数签名

```sql
pg_prewarm(regclass,
           mode text DEFAULT 'buffer',
           fork text DEFAULT 'main',
           first_block int8 DEFAULT NULL,
           last_block int8 DEFAULT NULL
) RETURNS int8
```

返回预热的块数。

| 参数 | 描述 |
|-----------|-------------|
| `mode` | `prefetch`（异步 OS）、`read`（同步 OS）或 `buffer`（PG 缓存） |
| `fork` | 要预热的关系分支（例如 `main`、`fsm`、`vm`） |
| `first_block` | 起始块号（默认：0） |
| `last_block` | 结束块号（默认：关系的最后一个块） |

### 自动预热

通过 `shared_preload_libraries` 加载时，自动预热会定期将共享缓冲区缓存中的缓冲区列表保存到磁盘，并在重启时恢复。

```sql
-- 手动启动自动预热工作进程
SELECT autoprewarm_start_worker();

-- 强制立即转储缓冲区状态
SELECT autoprewarm_dump_now();  -- 返回写入的记录数
```

### GUC 参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_prewarm.autoprewarm` | `true` | 启用自动预热后台工作进程 |
| `pg_prewarm.autoprewarm_interval` | `300s` | `autoprewarm.blocks` 文件更新间隔（0 = 仅在关闭时转储） |

缓冲区状态保存到数据目录中的 `autoprewarm.blocks` 文件。重启后，两个后台工作进程会重新加载保存的缓冲区。
