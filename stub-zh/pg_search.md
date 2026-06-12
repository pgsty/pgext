## 用法

来源：[ParadeDB extension install docs](https://docs.paradedb.com/deploy/self-hosted/extension)、[create-index docs](https://docs.paradedb.com/documentation/indexing/create-index.md)、[match docs](https://docs.paradedb.com/documentation/full-text/match.md)、[score docs](https://docs.paradedb.com/documentation/sorting/score.md)、[highlight docs](https://docs.paradedb.com/documentation/full-text/highlight.md)、[v0.24.0 release](https://github.com/paradedb/paradedb/releases/tag/v0.24.0)、[pg_search README](https://github.com/paradedb/paradedb/blob/v0.24.0/pg_search/README.md)

`pg_search` 是 ParadeDB 为 PostgreSQL 提供的 BM25 搜索扩展。上游 README 说明支持从 PostgreSQL 15 开始；Pigsty 为 PostgreSQL 15-18 打包 `0.24.0`，并使用 `cargo-pgrx` 0.18.1 构建。

### 启用并创建扩展

```conf
shared_preload_libraries = 'pg_search'
```

```sql
CREATE EXTENSION pg_search;
```

自托管扩展文档要求在 `CREATE EXTENSION` 之前先设置 `shared_preload_libraries = 'pg_search'`。

### 创建 BM25 索引

当前示例使用 `bm25` access method，并指定唯一键字段：

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating)
WITH (key_field = 'id');
```

每张表只支持一个 BM25 索引。`key_field` 是必填项，必须唯一，且必须是第一个被索引的列；文本类型 key fields 必须不分词。

### 查询操作符与辅助函数

当前文档使用以下查询操作符：

- `|||`：析取匹配，等价于 `term1 OR term2`。
- `&&&`：合取匹配，等价于 `term1 AND term2`。

示例：

```sql
SELECT description, rating
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY rating
LIMIT 5;

SELECT description, pdb.score(id) AS score
FROM mock_items
WHERE description &&& 'running shoes'
ORDER BY score DESC
LIMIT 5;

SELECT description, pdb.snippet(description) AS snippet, pdb.score(id) AS score
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY score DESC
LIMIT 5;
```

常用结果辅助函数包括 `pdb.score(id)`、`pdb.snippet(field)`、`pdb.snippets(field)` 和 `pdb.snippet_positions(field)`。高亮相对昂贵，且不支持 fuzzy search queries。

### 说明

- 旧 quickstart URL 已移除；当前 `|||`、`&&&`、scoring 和 highlighting 语法应以上方版本化文档页面为准。
- Release `0.24.0` 要求 preload `pg_search`，将 pgrx 升级到 0.18.1，并记录了 crash-recovery、`ltree` 与 inline-tokenizer 相关工作；上面的基础 BM25 查询示例没有变化。
- Pigsty metadata 说明 `bm25` access method 与 `pg_textsearch`、`vchord_bm25` 冲突；未测试目标组合前，不要在同一集群中 preload 竞争性的 BM25 access-method 扩展。
