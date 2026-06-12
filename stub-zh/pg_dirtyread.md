## 用法

来源：[upstream README](https://github.com/df7cb/pg_dirtyread)、[Debian changelog](https://github.com/df7cb/pg_dirtyread/blob/master/debian/changelog)、[tags](https://github.com/df7cb/pg_dirtyread/tags)。

`pg_dirtyread` 可以读取尚未被 vacuum 清理掉的死亡或已更新堆表行。该函数返回 `record`，因此每次调用都需要带一个表别名，用来声明希望返回的列。

### 基本用法

```sql
CREATE EXTENSION pg_dirtyread;

SELECT *
FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

列按名称匹配，因此别名可以省略部分列，也可以使用不同的列顺序。

### 示例

```sql
CREATE TABLE foo (bar bigint, baz text);
ALTER TABLE foo SET (autovacuum_enabled = false, toast.autovacuum_enabled = false);

INSERT INTO foo VALUES (1, 'Test'), (2, 'New Test');
DELETE FROM foo WHERE bar = 1;

SELECT * FROM pg_dirtyread('foo') AS t(bar bigint, baz text);
```

被删除的行在 vacuum 移除之前仍可能可见。

### 已删除的列

只要表还没有被 `VACUUM FULL` 或 `CLUSTER` 等操作重写，就可以取回已删除列的内容。使用 `dropped_N`，其中 `N` 是原始的 1 基列序号：

```sql
CREATE TABLE ab(a text, b text);
INSERT INTO ab VALUES ('Hello', 'World');
ALTER TABLE ab DROP COLUMN b;
DELETE FROM ab;

SELECT *
FROM pg_dirtyread('ab') AS ab(a text, dropped_2 text);
```

由于 PostgreSQL 会移除已删除列的原始类型元数据，这里只能执行有限的类型检查。

### 系统列

在别名中包含系统列即可取回它们。特殊的 `dead boolean` 列会报告确定已经死亡的行：

```sql
SELECT *
FROM pg_dirtyread('foo') AS t(
  tableoid oid,
  ctid tid,
  xmin xid,
  xmax xid,
  cmin cid,
  cmax cid,
  dead boolean,
  bar bigint,
  baz text
);
```

`dead` 列在恢复期间不可用，包括在备用服务器上。`oid` 系统列只在 PostgreSQL 11 及更早版本中可用。

### 注意事项

- Pigsty 将 `pg_dirtyread` 2.8 作为 PostgreSQL 14-18 的 RPM 打包；DEB 可用性来自 PGDG 的 `postgresql-$v-dirtyread`。
- 上游 2.8 是面向 PostgreSQL 19 默认 TOAST 压缩改为 `lz4` 的兼容性发布；没有记录新的面向用户 SQL 函数。
- `pg_dirtyread` 用于取证和恢复类检查。它绕过常规 MVCC 可见性预期，不应被用于应用读取。
