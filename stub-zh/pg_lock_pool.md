## 用法

来源：

- [官方 README](https://github.com/adjust/pg_lock_pool/blob/4d8a59ad82b6170b4fe827e0fa5ce4c8bc4b4695/Readme.md)
- [扩展控制文件](https://github.com/adjust/pg_lock_pool/blob/4d8a59ad82b6170b4fe827e0fa5ce4c8bc4b4695/pg_lock_pool.control)
- [0.0.1 版扩展 SQL](https://github.com/adjust/pg_lock_pool/blob/4d8a59ad82b6170b4fe827e0fa5ce4c8bc4b4695/pg_lock_pool--0.0.1.sql)

`pg_lock_pool` 0.0.1 提供 PL/pgSQL 辅助函数，在由双键 PostgreSQL 咨询锁表示的池中等待任意空闲槽位。它无需创建锁表，就能把协作任务的并发持有者限制为 `poolsize` 个。

### 会话级锁池

```sql
CREATE EXTENSION pg_lock_pool;

SELECT get_lock_pool(999, 3, 30) AS slot;
SELECT pg_advisory_unlock(999, 1);
```

`get_lock_pool(key, poolsize, timeout)` 使用 `pg_try_advisory_lock(key, slot)` 探测 `1..poolsize` 的槽位。成功时返回正槽位号；调用方必须把相同的键和返回槽位传给 `pg_advisory_unlock`。否则会话级锁会一直保留到数据库会话结束，事务回滚也不会释放它。

### 事务级锁池

```sql
BEGIN;
SELECT get_xact_lock_pool(999, 3, 30) AS slot;
COMMIT;
```

`get_xact_lock_pool` 使用 `pg_try_advisory_xact_lock`，并在事务结束时自动释放选中的槽位。两个函数的 `timeout` 均默认为 10 秒；循环到达超时时仍未取得槽位则返回 `-1`。

### 运维语义

每个失败轮次会尝试所有槽位、检查整数重试计数器，然后休眠一秒。因此超时是近似重试时长，不是精确墙钟截止时间，而且每个等待者都会占用一个 PostgreSQL 后端。

咨询锁是协作机制：无关 SQL 可以忽略它，所有参与者必须使用相同的 `key`、`poolsize` 以及会话级/事务级约定。应用代码应校验 `poolsize` 和 `timeout` 为正值。对于会话锁，务必保存返回槽位并在错误路径释放；否则连接池可能使锁远超逻辑请求的生命周期。
