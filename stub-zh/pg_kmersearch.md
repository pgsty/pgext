## 用法

来源：

- [官方上游 README](https://github.com/astanabe/pg_kmersearch/blob/2ef5eadf2f01eb29c6ca96bc747be754788b1eb5/README.md)
- [官方扩展控制文件 (pg_kmersearch.control)](https://github.com/astanabe/pg_kmersearch/blob/2ef5eadf2f01eb29c6ca96bc747be754788b1eb5/pg_kmersearch.control)
- [官方扩展 SQL (pg_kmersearch--1.0.sql)](https://github.com/astanabe/pg_kmersearch/blob/2ef5eadf2f01eb29c6ca96bc747be754788b1eb5/pg_kmersearch--1.0.sql)

`pg_kmersearch` — 一个基于 k-mer 索引的 PostgreSQL 扩展，用于 DNA 序列相似性搜索。将其用于相应的向量、模型或检索工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_kmersearch;

-- Create a table with DNA sequences
CREATE TABLE sequences (id SERIAL, name TEXT, dna DNA4);

-- Create GIN index (default: int4 keys for k-mer size 16)
CREATE INDEX ON sequences USING gin(dna kmersearch_dna4_gin_ops_int4);

-- Insert sequences
INSERT INTO sequences (name, dna) VALUES ('seq1', 'ATCGATCG');

-- Search using k-mer similarity
SELECT * FROM sequences WHERE dna =% 'ATCGATCG';

-- Calculate similarity scores
SELECT name, kmersearch_matchscore(dna, 'ATCGATCG') AS score
FROM sequences ORDER BY score DESC;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `bit_length(DNA2)` 是一个扩展函数，返回 `integer`。
- `bit_length(DNA4)` 是一个扩展函数，返回 `integer`。
- `char_length(DNA2)` 是一个扩展函数，返回 `integer`。
- `char_length(DNA4)` 是一个扩展函数，返回 `integer`。
- `kmersearch_actual_min_score_cache_free()` 是一个扩展函数，返回 `integer`。
- `kmersearch_actual_min_score_cache_stats()` 是一个扩展函数，返回 `TABLE`。
- `kmersearch_consistent_int2(internal, int2, text, int4, internal, internal)` 是一个扩展函数，返回 `boolean`。
- `kmersearch_consistent_int4(internal, int2, text, int4, internal, internal)` 是一个扩展函数，返回 `boolean`。
- `kmersearch_consistent_int8(internal, int2, text, int4, internal, internal)` 是一个扩展函数，返回 `boolean`。
- `kmersearch_delete_tempfiles()` 是一个扩展函数，返回 `TABLE`。
- `kmersearch_dna2_cmp(DNA2, DNA2)` 是一个扩展函数，返回 `integer`。
- `kmersearch_dna2_eq(DNA2, DNA2)` 是一个扩展函数，返回 `boolean`。
- `kmersearch_dna2_ge(DNA2, DNA2)` 是一个扩展函数，返回 `boolean`。
- `kmersearch_dna2_gt(DNA2, DNA2)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
