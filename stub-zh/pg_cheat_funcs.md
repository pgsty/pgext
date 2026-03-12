

## 用法

> [pg_cheat_funcs: 提供作弊（但实用）的函数](https://github.com/MasaoFujii/pg_cheat_funcs)

`pg_cheat_funcs` 扩展提供了一系列用于调试、诊断和底层 PostgreSQL 操作的实用函数。许多函数仅限超级用户使用。

### 进程控制

```sql
SELECT pg_signal_process(12345, 'TERM');       -- 向 PG 进程发送信号
SELECT pg_get_priority(pg_backend_pid());      -- 获取调度优先级
SELECT pg_set_priority(pg_backend_pid(), 10);  -- 设置调度优先级 (-20..19)
SELECT pg_postmaster_pid();                    -- 获取 postmaster PID
SELECT pg_backend_start_time();                -- 服务器进程启动时间
```

### 内存上下文检查

```sql
-- 显示内存上下文统计（PG 9.6-13）
SELECT * FROM pg_stat_get_memory_context();
-- 列: name, parent, level, total_bytes, total_nblocks, free_bytes, free_chunks, used_bytes
```

### 预备语句检查

```sql
-- 显示预备语句的缓存计划信息
SELECT * FROM pg_cached_plan_source('my_stmt');
-- 列: generic_cost, total_custom_cost, num_custom_plans, force_generic, force_custom
```

### 事务与 WAL 函数

```sql
SELECT pg_xlogfile_name('0/1234568'::pg_lsn, false);  -- LSN 转 WAL 文件名
SELECT pg_wait_syncrep('0/1234568'::pg_lsn);           -- 等待同步复制
SELECT * FROM pg_stat_get_syncrep_waiters();            -- 列出同步复制等待者
SELECT pg_set_next_xid('100'::xid);                    -- 设置下一个事务 ID（危险操作）
SELECT * FROM pg_xid_assignment();                      -- XID 状态信息
```

### 检查点与恢复

```sql
SELECT pg_checkpoint(true, true, true);  -- fast, wait, force
SELECT pg_promote(true);                 -- 提升备库（PG <= 11）
SELECT * FROM pg_recovery_settings();    -- 显示 recovery.conf 参数
SELECT pg_show_primary_conninfo();       -- 显示 primary_conninfo
```

### 文件操作

```sql
SELECT * FROM pg_list_relation_filepath('my_table'::regclass);  -- 列出段文件
SELECT pg_file_write_binary('/tmp/test', '\x48656c6c6f'::bytea); -- 写入二进制文件
SELECT pg_file_fsync('/tmp/test');                                -- fsync 文件
```

### 文本与编码转换

```sql
SELECT to_octal(255);                   -- '377'
SELECT pg_text_to_hex('PostgreSQL');     -- '506f737467726553514c'
SELECT pg_hex_to_text('506f737467726553514c'); -- 'PostgreSQL'
SELECT pg_chr(9731);                     -- 雪人字符
```

### 压缩

```sql
SELECT pglz_compress('some text data');        -- PGLZ 将文本压缩为 bytea
SELECT pglz_decompress(compressed_data);       -- 解压回文本
SELECT pglz_compress_bytea(data);              -- 压缩 bytea
SELECT pglz_decompress_bytea(compressed_data); -- 解压为 bytea
```

### 咨询锁管理

```sql
SELECT pg_advisory_xact_unlock(12345);              -- 释放排他咨询锁
SELECT pg_advisory_xact_unlock_shared(12345);       -- 释放共享咨询锁
```

### GUC 参数

| 参数 | 默认值 | 描述 |
|-----------|---------|-------------|
| `pg_cheat_funcs.log_memory_context` | `off` | 查询执行后记录内存上下文统计 |
| `pg_cheat_funcs.hide_appname` | `false` | 隐藏客户端 application_name |
| `pg_cheat_funcs.log_session_start_options` | `off` | 记录连接启动选项 |
| `pg_cheat_funcs.scheduling_priority` | `0` | 进程调度优先级 (-20..19) |
| `pg_cheat_funcs.exit_on_segv` | `off` | 开启时，段错误仅终止当前会话 |
