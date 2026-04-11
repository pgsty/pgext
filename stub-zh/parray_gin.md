
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION parray_gin;
> CREATE INDEX test_tags_idx ON test_table USING gin (val parray_gin_ops);
> SELECT * FROM test_table WHERE val @> ARRAY['must','contain'];
> SELECT * FROM test_table WHERE val @@> ARRAY['what%like%'];
> ```
>
> 来源：[README](https://github.com/theirix/parray_gin)，[参考文档](https://github.com/theirix/parray_gin/blob/master/doc/parray_gin.md)

`parray_gin` 为 `text[]` 数组提供 GIN 索引和操作符支持，既支持严格匹配，也支持部分匹配。上游文档将其描述为基于 `pg_trgm` 三元组实现的数组索引方案。

## 数组索引

该扩展提供 `parray_gin_ops` 操作符类，可用于 `text[]` 上的 GIN 索引：

```sql
CREATE TABLE test_table(id bigserial, val text[]);
CREATE INDEX test_tags_idx ON test_table USING gin (val parray_gin_ops);
```

参考文档指出，被索引的值和查询都会拆分为 trigrams。由于 GIN 可能返回误命中，操作符匹配后还会进行复核。

## 操作符

### 严格匹配

`@> (text[], text[]) -> bool`

当左侧数组包含右侧数组中的所有元素时返回 `true`。

```sql
SELECT * FROM test_table WHERE val @> ARRAY['far'];
```

`<@ (text[], text[]) -> bool`

当左侧数组被右侧数组包含时返回 `true`。

```sql
SELECT * FROM test_table WHERE val <@ ARRAY['galaxy','ago','vader'];
```

### 部分匹配

`@@> (text[], text[]) -> bool`

当左侧数组以部分匹配方式包含右侧所有项时返回 `true`，例如 `'foobar' ~~ 'foo%'` 或 `'foobar' ~~ '%oo%'`。

```sql
SELECT * FROM test_table WHERE val @@> ARRAY['%ar%'];
```

`<@@ (text[], text[]) -> bool`

当左侧数组被右侧所有模式部分匹配包含时返回 `true`。

```sql
SELECT * FROM test_table WHERE val <@@ ARRAY['%ar%','vader'];
```

## 说明

上游文档指出，GIN 可用于 `@>`, `<@`, `@@>` 和 `<@@`。文档还提到，该扩展可用于从 JSON 文本字段中提取出的 JSON 数组，相关场景中曾与 `json_accessors` 扩展配合使用。
