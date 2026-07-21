## 用法

来源：

- [官方 v0.2.0 README](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/README.md)
- [v0.2.0 更新日志](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/CHANGELOG.md)
- [v0.2.0 SQL API](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/pg_fts--0.2.0.sql)
- [v0.2.0 控制文件](https://codeberg.org/gregburd/pg_fts/src/tag/v0.2.0/pg_fts.control)

`pg_fts` 通过专用的 `ftsdoc` 和 `ftsquery` 类型以及 `fts` 倒排索引访问方法提供 BM25/BM25F 全文排序。它支持布尔、短语、NEAR、前缀、模糊和正则表达式术语，同时在索引中保留语料库统计信息以进行相关性评分。版本 `0.2.0` 要求 PostgreSQL 17 或更高版本。

### 创建并查询索引

```sql
CREATE EXTENSION pg_fts;

CREATE TABLE docs (
    id bigint PRIMARY KEY,
    body text NOT NULL
);

CREATE INDEX docs_fts
ON docs USING fts (to_ftsdoc('english', body));
```

使用相同的文本搜索配置用于文档和普通查询术语：

```sql
WITH q AS (
    SELECT to_ftsquery('english', 'postgres & "query planner" & index*') AS query
)
SELECT d.id,
       fts_snippet(d.body, q.query) AS excerpt
FROM docs AS d
CROSS JOIN q
WHERE to_ftsdoc('english', d.body) @@@ q.query
ORDER BY to_ftsdoc('english', d.body) <=> q.query
LIMIT 10;
```

`@@@` 匹配，而 `<=>` 升序距离按降序相关性排序行，并可驱动索引顺序扫描以支持 top-k 查询。

### 查询语言与 API 索引

- `to_ftsdoc([regconfig,] text)` 和 `to_ftsquery([regconfig,] text)`: 分析文档并解析查询。
- `quick brown`, `quick & brown`, `quick | brown` 和 `!slow`: 显式/隐式 AND、OR 和 NOT。
- `"quick brown"`, `NEAR(...)`, `term*`, `term~2` 和 `/regular-expression/`: 短语、接近度、前缀、模糊和正则表达式术语。
- `fts_bm25`, `fts_bm25_opts` 和 `fts_bm25f`: 显式的 BM25 排序变体和多字段排序。
- `fts_index_stats(index)` 和 `fts_index_df(index, query)`: 索引维护的文档计数、平均长度、词汇表大小和词频。
- `fts_highlight` 和 `fts_snippet`: 展示匹配文本。
- `fts_search(index, query, k)` 和 `fts_count(index, query)`: 索引本地 top-k 和 MVCC 意识的计数操作。
- `tsquery_to_ftsquery(tsquery)`: 迁移辅助程序；它不会使 `pg_fts` 成为 `tsvector`/GIN 的透明替代品。

### 维护与版本注意事项

```sql
SELECT fts_merge('docs_fts');
SELECT fts_vacuum('docs_fts');
```

- 插入进入立即可匹配的待处理列表，但排名 `<=>` 和 `fts_search` 结果覆盖合并段。当新插入文档必须立即参与排序时，请运行 `fts_merge()`。
- `fts_vacuum()` 会压缩段并回收可重用的索引页面；普通 `VACUUM` 也会参与待处理列表和删除标记维护。
- 版本 `0.2.0` 将访问方法从 `bm25` 重命名为 `fts`。由 `0.1.0` 创建且使用 `USING bm25` 的索引必须重新创建。
- 如果库报告磁盘格式不匹配，请遵循其 `REINDEX` 提示，而不是尝试用不同版本的格式读取索引。
- 访问方法是非覆盖的，并且在此版本中不提供并行扫描。在逻辑复制订阅者上分别安装扩展和索引；索引本身不会进行逻辑复制。
