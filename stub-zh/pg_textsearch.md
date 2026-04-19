## 用法

来源：[README](https://github.com/timescale/pg_textsearch/blob/main/README.md)，[release notes](https://github.com/timescale/pg_textsearch/releases)，[Timescale changelog](https://github.com/timescale/docs/blob/latest/about/changelog.md)

`pg_textsearch` 为 PostgreSQL 提供基于 BM25 的排序全文检索，核心接口是 `bm25` access method 和 `<@>` 评分运算符。上游已将 `v1.0.0` 标记为 production ready。

### 启用扩展

```ini
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
```

### 构建 BM25 索引并查询

```sql
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);

CREATE INDEX docs_idx
ON documents USING bm25(content)
WITH (text_config = 'english');

SELECT *
FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

README 说明 `<@>` 返回的是负 BM25 分数，因此值越低表示匹配越好。

### 显式查询与索引选项

```sql
SELECT *
FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;

CREATE INDEX ON documents USING bm25(content)
WITH (text_config = 'english', k1 = 1.5, b = 0.8);
```

README 还记录了 expression index、partial index 和 multilingual partial index。

### 常用函数与 GUC

```sql
SELECT bm25_force_merge('docs_idx');
```

当前上游文档列出的 GUC 包括：

- `pg_textsearch.default_limit`
- `pg_textsearch.segments_per_level`
- `pg_textsearch.memory_limit`
- `pg_textsearch.log_scores`

### 注意事项

- 上游当前仅支持 PostgreSQL 17 和 18。
- 在 PL/pgSQL 和 stored procedure 中，隐式 `text <@> 'query'` 形式不会使用 planner hook；上游要求改用带显式索引名的 `to_bm25query()`。
- `v1.0.0` 引入了 production-ready 状态、`bm25_force_merge()` 以及较新的 GUC，这些变化都已在 README 和官方 changelog 中记录。
