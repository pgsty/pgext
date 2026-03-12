

## 用法

> [btree_gin: B 树等价的 GIN 操作符类](https://www.postgresql.org/docs/current/btree-gin.html)

为通常仅支持 B 树索引的数据类型提供 GIN 索引操作符类。适用于将 GIN 可索引列和 B 树可索引列组合的多列 GIN 索引。

```sql
CREATE EXTENSION btree_gin;
```

### 支持的数据类型

`int2`、`int4`、`int8`、`float4`、`float8`、`numeric`、`timestamp with time zone`、`timestamp without time zone`、`time with time zone`、`time without time zone`、`date`、`interval`、`oid`、`money`、`char`、`varchar`、`text`、`bytea`、`macaddr`、`macaddr8`、`inet`、`cidr`、`uuid`、`bit`、`varbit`、`bool`、`name`、`bpchar` 以及所有 `enum` 类型。

### 示例

```sql
-- 整数列上的 GIN 索引
CREATE INDEX idx ON test USING GIN (a);
SELECT * FROM test WHERE a < 10;

-- 将全文搜索与标量过滤结合的多列 GIN 索引
CREATE INDEX idx ON articles USING GIN (body_tsvector, category);
SELECT * FROM articles
WHERE body_tsvector @@ to_tsquery('PostgreSQL')
  AND category = 'tech';
```

注意：btree_gin 在单列查询时不会优于标准 B 树索引。其主要优势在于将标量列与 GIN 原生列（如 tsvector 或数组）组合到单个多列索引中。
