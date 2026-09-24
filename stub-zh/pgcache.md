## 用法

来源：

- [README](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/README.md)
- [Control](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgcache/pgcache.control)
- [SQL](https://github.com/MEGALab/pgeverything/blob/517ff96f6c398b38008ccdb4586d2ab691921893/extensions/pgcache/pgcache--0.1.0.sql)

`pgcache` 是 PGEverything 的 SQL 键值缓存，支持 JSON 值与可选过期时间。此 0.1.0 源码快照用于上游 PostgreSQL 16 容器，存储的是可丢弃缓存，不是持久应用状态。

### 核心用法

安装 SQL 和控制文件后，在目标数据库启用扩展。删除与清理辅助函数需要 PL/pgSQL，没有共享库或预加载步骤。

```sql
CREATE EXTENSION pgcache;
SELECT cache_set('user:1', '{"name":"alice"}'::jsonb, ttl => 60);
SELECT cache_get('user:1');
SELECT cache_incr('hits');
SELECT cache_del('user:1');
SELECT pgcache.purge_expired();
```

### 接口与存储

| 函数 | 行为 |
| --- | --- |
| `cache_set(text,jsonb,integer)` | 插入或覆盖值；TTL 单位为秒，NULL 表示不过期 |
| `cache_get(text)` | 返回 JSON；缺失或过期时返回 NULL |
| `cache_del(text)` | 删除键，并返回此前是否存在该行 |
| `cache_incr(text,bigint)` | 原子递增整数 JSON 标量；增量默认为 1，更新时清除 TTL |
| `pgcache.purge_expired()` | 删除已过期行并返回数量 |

数据存放在非日志表 `pgcache.store` 中。崩溃恢复会清空非日志数据，它也不属于普通物理复制的数据源。读取会隐藏过期行，但不会删除它们，需要时应单独安排清理。`pg_cron` 是可选的外部调度器，不是安装依赖。

### 权限与限制

这些函数使用调用者权限，调用者仍需相应模式与表权限。控制文件设置了 `superuser=false`，但没有将扩展标记为可信；安装时仍须具备所需的常规 DDL 权限。内部模式固定，扩展不可迁移模式。上游未提供明确许可证，也没有公布 PostgreSQL 16 部署之外的支持矩阵。
