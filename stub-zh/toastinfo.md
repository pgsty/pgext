


## 用法

来源：[upstream README](https://github.com/credativ/toastinfo)、[upstream tags](https://github.com/credativ/toastinfo/tags)。

`toastinfo` 暴露 PostgreSQL 存储变长值（`varlena`）的方式，包括行内、压缩和行外 TOAST 存储形式。

### 函数

`pg_toastinfo(anyelement)` 描述一个 datum 的存储形式：

```sql
CREATE EXTENSION toastinfo;

SELECT a, length(b), pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;
```

可能的存储形式包括：

- `null`，表示 NULL 值。
- `ordinary`，表示非 varlena 数据类型。
- `short inline varlena`，最多 126 字节，使用 1 字节头。
- `long inline varlena, uncompressed`，最多 1 GiB，使用 4 字节头。
- `long inline varlena, compressed (pglz|lz4)`。
- `toasted varlena, uncompressed`。
- `toasted varlena, compressed (pglz|lz4)`。

在 PostgreSQL 14+ 上，压缩 varlena 会显示压缩方法。

`pg_toastpointer(anyelement)` 返回行外 toasted 值在 TOAST 表中的 `chunk_id` OID；对于非 toasted 输入返回 NULL：

```sql
SELECT pg_toastpointer(large_column)
FROM my_table;
```

### 存储示例

```sql
CREATE TABLE t (a text, b text);

ALTER TABLE t ALTER COLUMN b SET STORAGE EXTERNAL;
INSERT INTO t VALUES ('external-10000', repeat('x', 10000));

ALTER TABLE t ALTER COLUMN b SET STORAGE EXTENDED;
INSERT INTO t VALUES ('extended-1000000', repeat('x', 1000000));

ALTER TABLE t ALTER COLUMN b SET COMPRESSION lz4;
INSERT INTO t VALUES ('extended-lz4', repeat('x', 1000000));

SELECT a, pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;
```

### 注意事项

- Pigsty 元数据记录 `toastinfo` 1.6 覆盖 PostgreSQL 14-18，并同时有 Pigsty RPM 和 PGDG DEB。
- 上游 GitHub README 记录的是相同 SQL 表面；但评审时公开 GitHub tags 只可见到 `v1.5`，在该仓库中没有找到上游 1.6 changelog。
- `pg_toastpointer` 只对行外 toasted 数据有意义；普通、行内或 NULL 值都会返回 NULL。
