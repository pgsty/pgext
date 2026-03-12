

## 用法

> [pg_dirtyread: 读取已删除但未被清理的死元组](https://github.com/df7cb/pg_dirtyread)

`pg_dirtyread` 允许读取尚未被 VACUUM 清理的死元组（已删除/已更新的行）。该函数返回 RECORD 类型，因此需要描述模式的表别名。

### 基本用法

```sql
SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

### 示例

```sql
CREATE TABLE foo (bar bigint, baz text);
ALTER TABLE foo SET (autovacuum_enabled = false, toast.autovacuum_enabled = false);

INSERT INTO foo VALUES (1, 'Test'), (2, 'New Test');
DELETE FROM foo WHERE bar = 1;

SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
 bar |   baz
-----+----------
   1 | Test
   2 | New Test
```

### 已删除的列

使用 `dropped_N`（第 N 列，从 1 开始）访问已删除列的内容，前提是表未被重写（例如通过 `VACUUM FULL` 或 `CLUSTER`）：

```sql
ALTER TABLE ab DROP COLUMN b;
DELETE FROM ab;
SELECT * FROM pg_dirtyread('ab') ab(a text, dropped_2 text);
```

### 系统列

在别名中包含系统列即可获取它们。特殊的 `dead` 布尔列标识死元组：

```sql
SELECT * FROM pg_dirtyread('foo')
    AS t(tableoid oid, ctid tid, xmin xid, xmax xid, cmin cid, cmax cid, dead boolean,
         bar bigint, baz text);
```

`dead` 列在恢复期间（例如备用服务器上）不可用。
