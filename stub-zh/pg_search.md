
> [!NOTE] 此扩展由 ParadeDB 团队开发，通过 PIGSTY 仓库分发

## 用法

> [README](https://github.com/paradedb/paradedb/tree/dev/pg_search) | [Quickstart](https://docs.paradedb.com/documentation/getting-started/quickstart) | [Install docs](https://docs.paradedb.com/documentation/getting-started/install)

`pg_search` 是 ParadeDB 为 PostgreSQL 提供的全文检索扩展。它基于 Tantivy，为堆表提供 BM25 索引和查询能力。当前上游文档说明，它支持 PostgreSQL 15 及以上版本。

## 入门

安装文档强调一个关键要求：必须把 `pg_search` 加入 `shared_preload_libraries`，这样后台工作进程才能处理索引写入。

```ini
shared_preload_libraries = 'pg_search'
```

随后启用扩展：

```sql
CREATE EXTENSION pg_search;
ALTER SYSTEM SET paradedb.pg_search_telemetry TO 'off';
```

## 创建 BM25 索引

Quickstart 展示了在堆表上创建 BM25 索引的方式，且需要唯一键字段：

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating, in_stock, created_at, metadata, weight_range)
WITH (key_field='id');
```

现有文档强调，`key_field` 必须唯一，BM25 索引是搜索查询的核心访问方法。

## 查询

`@@@` 操作符用于执行搜索查询：

```sql
SELECT description, rating, category
FROM mock_items
WHERE description @@@ 'keyboard' AND rating > 2
ORDER BY rating
LIMIT 5;
```

ParadeDB 还提供用于相关性评分和摘要高亮的辅助函数：

```sql
SELECT description, paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;

SELECT description, paradedb.snippet(description), paradedb.score(id)
FROM mock_items
WHERE description @@@ 'keyboard'
ORDER BY paradedb.score(id) DESC
LIMIT 5;
```

短语搜索支持用引号包裹表达式：

```sql
SELECT description
FROM mock_items
WHERE description @@@ '"metal keyboard"';
```

## 文本配置

Quickstart 还展示了如何为文本字段配置分词器，例如英文词干提取：

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, (description::pdb.simple('stemmer=english')), category)
WITH (key_field='id');
```

更深入的部署和运行说明请参考上游 ParadeDB 文档站，它是该项目的主文档入口。
