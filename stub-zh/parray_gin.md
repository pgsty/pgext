## 用法

- 来源：[README](https://github.com/theirix/parray_gin/blob/master/README.md)，[reference doc](https://github.com/theirix/parray_gin/blob/master/doc/parray_gin.md)

`parray_gin` 为 `text[]` 增加一个 GIN operator class，并提供严格匹配与部分匹配操作符。上游把它描述为基于 `pg_trgm` trigram 分解的数组索引。

### 创建扩展与索引

```sql
CREATE EXTENSION parray_gin;

CREATE TABLE test_table (
  id  bigserial,
  val text[]
);

CREATE INDEX test_tags_idx
ON test_table
USING gin (val parray_gin_ops);
```

### 可索引操作符

reference doc 说明 `parray_gin_ops` 支持下列操作符：

- `@>`：严格包含。
- `<@`：严格被包含。
- `@@>`：部分包含，数组元素可以使用 `LIKE` 模式。
- `<@@`：部分被包含。

示例：

```sql
SELECT * FROM test_table WHERE val @> ARRAY['must','contain'];
SELECT * FROM test_table WHERE val @@> ARRAY['what%like%'];
SELECT * FROM test_table WHERE val <@ ARRAY['galaxy','ago','vader'];
SELECT * FROM test_table WHERE val <@@ ARRAY['%ar%','vader'];
```

### 匹配行为

严格匹配要求数组元素完全相等。部分匹配允许 `'foo%'` 或 `'%oo%'` 这类模式。由于 trigram 索引可能返回误命中，文档说明索引查找之后还会做 recheck。

### 注意事项

README 表示支持范围一直到 PostgreSQL 18，而 reference doc 仍写成 9.1-14。两份文档对操作符和 opclass 行为的描述是一致的，但版本说明在上游尚未完全同步。
