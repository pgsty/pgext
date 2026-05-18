## 用法

来源：[README v1.2.0](https://github.com/timescale/pg_textsearch/blob/v1.2.0/README.md), [v1.2.0 release notes](https://github.com/timescale/pg_textsearch/releases/tag/v1.2.0)

`pg_textsearch` 为 PostgreSQL 提供 BM25-ranked full-text search，包含 `bm25` access method 和 `<@>` scoring operator。上游将 `v1.2.0` 标记为 production ready。

`v1.2.0` 支持 PostgreSQL 17 和 18。Linux 与 macOS 上两个 PostgreSQL 版本都有预构建 release assets。执行 `CREATE EXTENSION` 前，必须通过 `shared_preload_libraries` 加载该扩展。

### 启用扩展

```ini
shared_preload_libraries = 'pg_textsearch'  # add to any existing list
```

```sql
CREATE EXTENSION pg_textsearch;
```

运行扩展升级前，请先安装新 binary 并重启 PostgreSQL：

```sql
ALTER EXTENSION pg_textsearch UPDATE;
```

`v1.2.0` release 增加了 physical replication 支持，并修复了 update-heavy 工作负载下的正确性问题。依赖新版本前，请安装匹配的 binary 并运行 SQL extension upgrade。

### 构建和查询 BM25 索引

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

`<@>` 返回负的 BM25 score，因为 PostgreSQL operator index scans 是升序；值越低匹配越好。使用 `ORDER BY ... LIMIT` 进行快速 top-k 搜索。

显式引用索引时，使用 `to_bm25query()`：

```sql
SELECT *
FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

主要文档化 SQL 表面包括：

- `text <@> 'query'`：在 planner 检测到的 index context 中对文本评分。
- `text <@> bm25query`：使用显式 `bm25query` 评分。
- `to_bm25query(text)`：用于 planner 选择索引上下文的 `ORDER BY`。
- `to_bm25query(text, text)`：查询文本加索引名。
- `bm25query = bm25query`：等值检查。

### 索引选项与数据形状

```sql
CREATE INDEX ON documents USING bm25(content)
WITH (text_config = 'english', k1 = 1.5, b = 0.8);
```

索引选项包括 `text_config`（必填）、`k1`（默认 `1.2`）和 `b`（默认 `0.75`）。`english`、`simple`、`french`、`german` 等 text search configurations 使用 PostgreSQL text search configuration 名称。

扩展原生支持 `text[]`、`varchar[]` 和 `bpchar[]` 列作为数组输入；数组元素会在 tokenization 前拼接。

```sql
CREATE TABLE posts (id serial PRIMARY KEY, tags text[]);
CREATE INDEX posts_tags_idx ON posts USING bm25(tags)
WITH (text_config = 'english');

SELECT *
FROM posts
ORDER BY tags <@> 'database'
LIMIT 10;
```

Expression indexes 支持 immutable text expressions，包括 JSONB 提取、文本转换和多列拼接：

```sql
CREATE INDEX events_msg_idx ON events USING bm25 ((data->>'message'))
WITH (text_config = 'english');

SELECT *
FROM events
ORDER BY (data->>'message') <@> to_bm25query('network error', 'events_msg_idx')
LIMIT 10;
```

Partial indexes 会将搜索限定到行子集。查询时请显式指定索引名：

```sql
CREATE INDEX docs_en_idx ON docs USING bm25(content)
WITH (text_config = 'english')
WHERE lang = 'en';

SELECT *
FROM docs
WHERE lang = 'en'
ORDER BY content <@> to_bm25query('databases', 'docs_en_idx')
LIMIT 10;
```

### 运维与 GUC

```sql
SELECT bm25_force_merge('docs_idx');
SELECT * FROM bm25_memory_usage();
```

`bm25_force_merge(index_name)` 会将所有 segments 合并成一个，适合批量加载后使用，不适合稳定写入流量中频繁使用。`bm25_memory_usage()` 报告 memtables 的 shared memory 使用情况。

v1.2.0 中文档化的 `pg_textsearch` GUC 包括：

- `pg_textsearch.default_limit`
- `pg_textsearch.compress_segments`
- `pg_textsearch.segments_per_level`
- `pg_textsearch.memory_limit`
- `pg_textsearch.bulk_load_threshold`
- `pg_textsearch.memtable_spill_threshold`（deprecated；新部署请使用 `memory_limit`）

`pg_textsearch.memory_limit` 默认 `2GB`，限制 memtables 使用的 dynamic shared memory。README 还将 `bm25_spill_index(index_name)`、`bm25_dump_index(index_name)` 和 `bm25_summarize_index(index_name)` 记录为开发或诊断辅助函数。

### 注意事项

- `pg_textsearch` 要求 `shared_preload_libraries = 'pg_textsearch'`，并且必须在 `CREATE EXTENSION` 前重启 PostgreSQL。
- `bm25` access method 名称会与 `pg_search` 和 `vchord_bm25` 冲突；避免在同一个数据库中安装这些 BM25 access-method 扩展。
- 在 PL/pgSQL 和 stored procedures 内，隐式 `text <@> 'query'` 形式不会使用 planner hooks；上游说明应在那里使用带显式索引名的 `to_bm25query()`。
- phrase queries 不是原生能力，因为索引存储 term frequencies 而不是 term positions；需要类似 phrase 匹配时，可使用 BM25 ranking 加 post-filter。
- Partial indexes 要求使用带索引名的 `to_bm25query()`，因为隐式查询形式会跳过它们。
- 分区表上的 BM25 indexes 使用 partition-local statistics，因此跨分区 score 可能不能直接比较。
- 超过 PostgreSQL `tsvector` word length limit 的词会在 tokenization 中被忽略。
- `pg_textsearch` 使用固定 LWLock tranche IDs 1001-1008；另一个扩展若使用相同固定 ID，会导致 wait-event 名称不准确。
