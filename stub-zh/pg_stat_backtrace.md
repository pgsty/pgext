## 用法

> 来源：[pg_stat_backtrace upstream README](https://github.com/Nickyoung0/pg_stat_backtrace)、[upstream changelog](https://github.com/Nickyoung0/pg_stat_backtrace/blob/main/CHANGELOG.md)、[local metadata](../db/extension.csv)、本地源码归档 `pg_stat_backtrace-1.0.0.tar.gz`。

`pg_stat_backtrace` 用于捕获或记录同一 Linux 主机上 PostgreSQL 后端进程或辅助进程的 C 层调用栈。它使用 `ptrace(PTRACE_SEIZE)` 和 `libunwind`；不需要 `shared_preload_libraries`，也不会向目标进程发送 `SIGSTOP`。

```sql
CREATE EXTENSION pg_stat_backtrace;
```

### 捕获调用栈

先从 PostgreSQL 视图中找到目标进程，再调用 `pg_get_backtrace(pid)`：

```sql
SELECT pid, backend_type, state, wait_event, query
FROM pg_stat_activity
WHERE backend_type = 'autovacuum worker';

SELECT pg_get_backtrace(123456);
```

返回文本采用类似 `pstack(1)` 的格式：

```text
#0  0x00007f5e6c1a4d9e in __epoll_wait+0x4e
#1  0x000055f1a8c2a213 in WaitEventSetWaitBlock+0x83
#2  0x000055f1a8c2a04e in WaitEventSetWait+0xfe
```

### 写入服务器日志

如果希望结果进入常规 PostgreSQL 日志管道，可以使用 `pg_log_backtrace(pid)`：

```sql
SELECT pg_log_backtrace(123456);

SELECT pid, pg_log_backtrace(pid)
FROM pg_stat_activity
WHERE backend_type = 'walsender';
```

该函数成功时返回 `true`。

### 权限

默认情况下，两个函数都会从 `PUBLIC` 撤销执行权限。只应把访问权授予可信的监控角色：

```sql
GRANT EXECUTE ON FUNCTION pg_get_backtrace(int) TO observability;
GRANT EXECUTE ON FUNCTION pg_log_backtrace(int) TO observability;
```

C 代码仍会执行目标检查：

- 超级用户可以定位该实例中的任意 PostgreSQL 进程，包括 `walwriter`、`checkpointer`、`walsender`、autovacuum worker、startup、archiver 等辅助进程。
- 非超级用户只能定位由其成员角色拥有的普通后端进程。
- 辅助进程没有角色所有权，因此非超级用户总会被拒绝。
- 非超级用户不能定位超级用户拥有的后端，即使存在角色成员关系也不行。

### 输入与错误行为

两个函数都是 `VOLATILE STRICT PARALLEL RESTRICTED`。

```sql
SELECT pg_get_backtrace(NULL::int);  -- NULL, no ptrace attempt
SELECT pg_log_backtrace(NULL::int);  -- NULL, no ptrace attempt
SELECT pg_get_backtrace(0);          -- WARNING, then NULL
SELECT pg_log_backtrace(-1);         -- WARNING, then false
```

自我定位会被拒绝，因为 Linux 进程不能 ptrace 自己：

```sql
SELECT pg_get_backtrace(pg_backend_pid());
```

### 运行注意事项

- Pigsty 将 `pg_stat_backtrace` 1.0.0 打包给 PostgreSQL 14-18。上游 1.0.0 同时声明兼容 PostgreSQL 19。
- 该扩展仅支持 Linux，构建和运行时都依赖 `libunwind` / `libunwind-ptrace`。
- 在启用了 Yama ptrace 限制的主机上，后端之间互相捕获调用栈可能需要设置 `kernel.yama.ptrace_scope = 0`。
- 展开调用栈期间，目标进程会被短暂暂停。不要在繁忙主库上对 `walwriter`、`checkpointer` 或同步复制 `walsender` 等关键进程进行紧密循环调用。
- Linux 对每个目标进程只允许一个 tracer。多个会话同时定位同一个 PID 时可能因 `EPERM` 失败；等待正在进行的调用结束后重试。
