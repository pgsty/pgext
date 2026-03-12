

## 用法

> [pg_readonly: 将集群数据库设置为只读](https://github.com/pierreforstmann/pg_readonly)

pg_readonly 在 SQL 层面将 PostgreSQL 集群中的所有数据库设置为只读模式。必须通过 `shared_preload_libraries` 加载。只读状态在共享内存中通过全局标志管理（重启后不保留）。

### 检查只读状态

```sql
SELECT get_cluster_readonly();
-- 返回 false（可读写）或 true（只读）
```

### 设置集群为只读

```sql
SELECT set_cluster_readonly();
```

在只读模式下，SELECT 语句被允许（除非调用了写入函数），但 DML（INSERT、UPDATE、DELETE）、DDL（包括 TRUNCATE）和 DCL（GRANT、REVOKE）被阻止：

```sql
SELECT * FROM t;              -- 允许
UPDATE t SET x = 33;          -- ERROR: pg_readonly: invalid statement because cluster is read-only
CREATE TABLE tmp(c text);     -- ERROR: pg_readonly: invalid statement because cluster is read-only
```

注意：`set_cluster_readonly()` 会终止所有打开的事务。

### 设置集群为可读写

```sql
SELECT unset_cluster_readonly();
```

注意：后台进程（checkpointer、bgwriter、walwriter、autovacuum）在只读模式下继续运行——限制仅在 SQL 语句层面。
