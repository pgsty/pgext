## 用法

- 来源：[README](https://github.com/fraruiz/pglock/blob/master/README.md)

`pglock` 是一个在 PostgreSQL 内实现的轻量分布式锁服务。它把锁状态保存在 `pglock.locks` 中，并通过基于 TTL 的清理来回收陈旧行。

### 创建扩展

```sql
CREATE EXTENSION pglock;
```

上游 README 标明要求 PostgreSQL 9.1+ 和 `plpgsql`。

### 获取与释放锁

```sql
SELECT pglock.lock('worker-1', 'users');
SELECT pglock.unlock('worker-1', 'users');
```

`pglock.lock(id, resource)` 在成功拿到锁时返回 `true`，如果锁已被其他进程持有则返回 `false`。文档说明 `pglock.unlock(id, resource)` 是幂等的。

### 隔离级别与过期

README 建议在并发场景下使用 serializable isolation 以保证正确性：

```sql
SELECT pglock.set_serializable();
```

```sql
BEGIN ISOLATION LEVEL SERIALIZABLE;
SELECT pglock.lock('my-id', 'my-resource');
SELECT pglock.unlock('my-id', 'my-resource');
COMMIT;
```

陈旧锁通过下面的方式过期清理：

```sql
SELECT pglock.ttl();
```

文档给出的默认 TTL 是 5 分钟。如果环境里有 `pg_cron`，README 说明可以每分钟调度一次 `pglock.ttl()`。

### 锁表

上游模式中的锁表是 `pglock.locks`，包含 `id`、`resource`、`locked`、`ttl`、`created_at` 和 `updated_at` 列。主键为 `(id, resource)`。
