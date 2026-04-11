
## 用法

> 语法：
>
> ```sql
> SELECT pglock.lock('b3d8a762-3a0e-495b-b6a1-dc8609839f7b', 'users');
> SELECT pglock.unlock('b3d8a762-3a0e-495b-b6a1-dc8609839f7b', 'users');
> SELECT pglock.ttl();
> ```
>
> 来源：[README](https://github.com/fraruiz/pglock)

`pglock` 是一个在 PostgreSQL 内部实现的轻量级分布式锁服务。它把锁存储在表中，并支持基于 TTL 的过期清理。

## 基本函数

README 记录了四个核心函数：

- `pglock.lock(id, resource)`，用于获取锁
- `pglock.unlock(id, resource)`，用于释放锁
- `pglock.ttl()`，用于清理过期锁
- `pglock.set_serializable()`，用于切换到可串行化隔离级别

获取锁：

```sql
SELECT pglock.lock('worker-1', 'users');
```

释放锁：

```sql
SELECT pglock.unlock('worker-1', 'users');
```

## 隔离级别

上游文档建议在并发场景下使用可串行化隔离级别以保证正确性：

```sql
SELECT pglock.set_serializable();
```

或者：

```sql
BEGIN ISOLATION LEVEL SERIALIZABLE;
SELECT pglock.lock('my-id', 'my-resource');
SELECT pglock.unlock('my-id', 'my-resource');
COMMIT;
```

## TTL 过期

锁带有可配置 TTL，默认值为 5 分钟。`pglock.ttl()` 会释放 `updated_at` 超过 TTL 的记录：

```sql
SELECT pglock.ttl();
```

如果安装了 `pg_cron`，README 说明可以每分钟运行一次 `pglock.ttl()`。

## 模式

锁表是 `pglock.locks`，字段包括：

- `id`
- `resource`
- `locked`
- `ttl`
- `created_at`
- `updated_at`

主键是 `(id, resource)`。
