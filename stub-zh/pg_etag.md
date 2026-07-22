## 用法

来源：

- [官方 README](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/README.md)
- [扩展 SQL](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/pg_etag--1.0.sql)
- [BLAKE2 实现](https://github.com/ergo70/pg_etag/blob/444c21d9df4475179d9ee40ec07fcf51a48e5501/pg_etag.c)

`pg_etag` 为 `text`、`bytea`、有序结果流和 PostgreSQL 大对象生成十六进制 BLAKE2b 摘要。只有在定义了稳定序列化和行顺序后，才能把这些摘要用作应用缓存验证器。

### 单值与行

```sql
CREATE EXTENSION pg_etag;

SELECT etag('canonical payload');
SELECT etag(convert_to('canonical payload', 'UTF8'));

SELECT id,
       etag(concat_ws(E'\x1f', id::text, updated_at::text, payload)) AS row_etag
FROM documents;

SELECT blake2('short digest', 32);
```

`etag(text)` 和 `etag(bytea)` 返回 64 字节 BLAKE2b 摘要，以 128 个小写十六进制字符表示。`blake2(text, integer)` 和 `blake2(bytea, integer)` 允许调用方选择 1 到 64 字节的摘要长度；实现会把区间外的值钳制到边界。

### 结果集 ETag

```sql
SELECT etag_agg(row_etag ORDER BY id) AS result_etag
FROM (
    SELECT id,
           etag(concat_ws(E'\x1f', id::text, updated_at::text, payload)) AS row_etag
    FROM documents
    WHERE tenant_id = 42
) AS ordered_rows;
```

`etag_agg(text)` 和 `etag_agg(bytea)` 按聚合输入顺序散列字节，并忽略空输入。必须在聚合内部写出确定性的 `ORDER BY`。状态函数拼接输入字节时不会加入分隔符，因此要在被散列的表示中编码字段和行边界；否则不同分段可能产生相同字节流。

### 大对象与注意事项

重载 `public.etag(oid)` 按页码顺序散列 `pg_largeobject.data` 页。调用者必须有权访问大对象；虽然控制文件声明扩展可迁移，该函数仍硬编码安装到 `public` 模式。

ETag 的稳定性取决于序列化。时间戳、数值、JSON、排序规则或编码的文本格式可能随设置和版本变化。应散列规范化字节，包含决定表示的所有字段，并注意 BLAKE2 的抗碰撞性并不提供身份认证或授权。该扩展依赖 `libb2`。
