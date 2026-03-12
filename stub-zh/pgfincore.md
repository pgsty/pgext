

## 用法

> [pgfincore: 检查和管理操作系统缓冲区缓存](https://github.com/klando/pgfincore)

pgfincore 提供使用 `mincore` 和 `POSIX_FADVISE` 检查和管理 PostgreSQL 关系的操作系统页缓存内容的函数。

### 检查缓存状态

```sql
SELECT * FROM pgfincore('pgbench_accounts');
```

返回每段信息：`relpath`、`segment`、`os_page_size`、`rel_os_pages`、`pages_mem`（OS 缓存中的页数）、`group_mem`、`os_pages_free`、`pages_dirty`、`group_dirty`。

使用 `pgfincore('relation', true)` 可包含用于快照/恢复的 `databit` varbit 映射。

### 系统信息

```sql
SELECT * FROM pgsysconf();          -- os_page_size, os_pages_free, os_total_pages
SELECT * FROM pgsysconf_pretty();   -- 相同内容，人类可读格式输出
```

### 预加载到 OS 缓存

```sql
SELECT * FROM pgfadvise_willneed('pgbench_accounts');
```

### 从 OS 缓存中驱逐

```sql
SELECT * FROM pgfadvise_dontneed('pgbench_accounts');
```

### 其他 POSIX_FADVISE 标志

```sql
SELECT * FROM pgfadvise_normal('relation');
SELECT * FROM pgfadvise_sequential('relation');
SELECT * FROM pgfadvise_random('relation');
```

### 快照和恢复缓存状态

```sql
-- 快照
CREATE TABLE pgfincore_snapshot AS
  SELECT 'pgbench_accounts'::text AS relname, *, now() AS date_snapshot
  FROM pgfincore('pgbench_accounts', true);

-- 恢复
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, true,
               (SELECT databit FROM pgfincore_snapshot
                WHERE relname = 'pgbench_accounts' AND segment = 0));
```

### 直接页缓存控制

```sql
-- 加载前 3 页，卸载后 3 页
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, true, B'111000');
-- 仅加载
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, true, false, B'111000');
-- 仅卸载
SELECT * FROM pgfadvise_loader('pgbench_accounts', 0, false, true, B'111000');
```
