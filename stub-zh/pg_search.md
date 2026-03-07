
> [!NOTE] 此扩展由 ParadeDB 团队开发，通过 PIGSTY 仓库分发

## 用法

https://docs.paradedb.com/documentation/getting-started/quickstart

```sql
CREATE EXTENSION pg_search;

ALTER SYSTEM SET paradedb.pg_search_telemetry TO 'off';

-- 创建 BM25 测试表
CALL paradedb.create_bm25_test_table(
  schema_name => 'public',
  table_name => 'mock_items'
);

SELECT description, rating, category FROM mock_items LIMIT 3;

-- 创建 BM25 索引（key_field 必须具有 UNIQUE 约束，每张表只能创建一个 BM25 索引）
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating, in_stock, created_at, metadata, weight_range)
WITH (key_field='id');

-- 使用 @@@ 操作符进行全文检索
SELECT description, rating, category
FROM mock_items
WHERE description @@@ 'keyboard' AND rating > 2
ORDER BY rating
LIMIT 5;

-- BM25 相关性评分
SELECT description, paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;

-- 高亮匹配的关键词
SELECT description, paradedb.snippet(description), paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;

-- 精确短语搜索（在单引号内使用双引号）
SELECT description, rating, category
FROM mock_items
WHERE description @@@ '"metal keyboard"';

-- 配置文本字段的分词器（例如英文词干提取）
DROP INDEX search_idx;
CREATE INDEX search_idx ON mock_items
USING bm25 (id, (description::pdb.simple('stemmer=english')), category)
WITH (key_field='id');
```
