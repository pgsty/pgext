


## 用法

> [toastinfo: 检查 varlena 列的 TOAST 存储详情](https://github.com/credativ/toastinfo)

toastinfo 暴露变长（varlena）数据类型的内部存储形式，显示 PostgreSQL 如何存储每个数据。

### 函数

**`pg_toastinfo(anyelement)`** -- 描述数据的存储形式：

```sql
SELECT a, length(b), pg_column_size(b), pg_toastinfo(b), pg_toastpointer(b)
FROM t;

        a         | length  | pg_column_size |              pg_toastinfo              | pg_toastpointer
------------------+---------+----------------+----------------------------------------+-----------------
 null             |       * |              * | null                                   |               *
 default          |       7 |              8 | short inline varlena                   |               *
 external-200     |     200 |            204 | long inline varlena, uncompressed      |               *
 external-10000   |   10000 |          10000 | toasted varlena, uncompressed          |           16427
 extended-10000   |   10000 |            125 | long inline varlena, compressed (pglz) |               *
 extended-1000000 | 1000000 |          11452 | toasted varlena, compressed (pglz)     |           16429
 extended-1000000 | 1000000 |           3936 | toasted varlena, compressed (lz4)      |           16430
```

可能的存储形式：
- `null` -- NULL 值
- `ordinary` -- 非 varlena 数据类型
- `short inline varlena` -- 最多 126 字节（1 字节头）
- `long inline varlena, (un)compressed` -- 最大 1GiB（4 字节头）
- `toasted varlena, (un)compressed` -- 最大 1GiB，存储在 TOAST 表中
- 压缩的 varlena 在 PG14+ 上显示压缩方法（`pglz`、`lz4`）

**`pg_toastpointer(anyelement)`** -- 返回 TOAST 表中的 `chunk_id` OID，非 TOAST 数据返回 NULL：

```sql
SELECT pg_toastpointer(large_column) FROM my_table;
```
