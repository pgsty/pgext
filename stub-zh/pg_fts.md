## 用法

来源：

- [README.md](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/README.md)
- [pg_fts.control](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/pg_fts.control)
- [CHANGELOG.md](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/CHANGELOG.md)
- [pg_fts--1.8.6--1.9.0.sql](https://github.com/gburd/pg_fts/blob/d5c645f615b7a4f4e1a9374c30200a33858be401/pg_fts--1.8.6--1.9.0.sql)

`pg_fts` 1.9.0 通过专用 fts 倒排索引提供 BM25/BM25F 排序的全文检索。支持 PostgreSQL 17 与 18；PostgreSQL 19/master 的 CI 属于尽力支持。控制文件标记可信且可重定位，核心使用无需共享预加载。

### 检索与排序

```sql
CREATE EXTENSION pg_fts;
CREATE TABLE docs (id bigint, body text);
CREATE INDEX docs_fts ON docs USING fts (to_ftsdoc('english', body));
SELECT id FROM docs
 WHERE to_ftsdoc('english', body) @@@ to_ftsquery('english', 'quick fox')
 ORDER BY to_ftsdoc('english', body) <=> to_ftsquery('english', 'quick fox')
 LIMIT 10;
SELECT fts_merge('docs_fts');
SELECT fts_vacuum('docs_fts');
```

### 对象与查询

`ftsdoc` 与 `ftsquery` 表示分析后的文档和查询。`to_ftsdoc()` 与 `to_ftsquery()` 构造它们；`@@@` 匹配文档，`<=>` 按相关性距离排序。索引 KNN 排序扫描必须包含匹配谓词。查询支持布尔项、短语、前缀、模糊与正则匹配。多列文档支持 BM25F 字段权重。`fts_count()` 与普通计数查询可使用精确的索引计数；`fts_search()` 提供直接排序结果。正则与长模糊项加速需通过 `trigrams = on` 显式启用。

### 维护与权限

待合并的新插入文档立即可搜索，合并不是可见性的前提。`fts_merge()` 合并分段，`fts_vacuum()` 回收物理空间。二者写入 WAL，需要索引所有权且必须在主库运行。较大的待合并文档可能暂时占用大量空间，批量导入时应定期合并与回收。`fts_search()` 与 `fts_anomalous_docs()` 会暴露索引内容，默认撤销 PUBLIC 权限；仅在明确评估后扩大访问。普通表查询可见性与直接辅助函数访问具有不同的权限边界。

### 升级至 1.9.0

安装匹配文件后执行 `ALTER EXTENSION pg_fts UPDATE TO '1.9.0'`。本版修复删除后排序错误和倒排块边界遗漏 top-k 结果的问题，没有磁盘格式变化，也不要求 REINDEX。`pg_fts.doclen_cache_mb` 默认 64，设为 0 关闭缓存；`pg_fts.dense_score_min_df` 默认 32768，设为 0 关闭该评分路径。应计算每个后端的缓存内存开销。
