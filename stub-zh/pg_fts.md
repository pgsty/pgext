## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_fts/pg_fts-1.12.0/README.md)
- [pg_fts.control](https://api.pgxn.org/src/pg_fts/pg_fts-1.12.0/pg_fts.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_fts/pg_fts-1.12.0/CHANGELOG.md)
- [pg_fts--1.11.0--1.12.0.sql](https://api.pgxn.org/src/pg_fts/pg_fts-1.12.0/pg_fts--1.11.0--1.12.0.sql)

- [README.md](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/README.md)
- [pg_fts.control](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/pg_fts.control)
- [CHANGELOG.md](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/CHANGELOG.md)
- [pg_fts--1.9.1--1.10.0.sql](https://github.com/gburd/pg_fts/blob/8bccf0ce75f04b3af559b34d635a5f71e0d2d04e/pg_fts--1.9.1--1.10.0.sql)

`pg_fts` 1.12.0 通过专用 fts 倒排索引提供 BM25/BM25F 排序的全文检索。支持 PostgreSQL 17 与 18；PostgreSQL 19/master 的 CI 属于尽力支持。控制文件标记可信且可重定位，核心使用无需共享预加载。

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

### 升级至 1.10.0

安装匹配文件后执行 `ALTER EXTENSION pg_fts UPDATE TO '1.10.0'`。1.10.0 默认在服务端共享各索引分段的文档长度数组，并提供 `fts_shared_doclen_stats()` 查看共享缓存。超级用户可用 `pg_fts.shared_doclen` 切换此路径；共享内存不可用时回退到后端私有副本。没有磁盘格式变化，也不要求 REINDEX。旧版删除后排序和倒排块边界修复仍然包含在本版中；规划内存时应区分共享缓存与回退副本。

### 1.11.0 版本

1.11.0 修复结果错误及并发插入、VACUUM、合并故障。排序遍历按块上界优先处理，`pg_fts.bestfirst` 默认启用。索引格式不变，SQL 升级不新增对象。仅为升级不必重建索引，但 REINDEX 可以修复旧缺陷已遗留或丢失的索引项。

### 1.12.0 版本

1.12.0 在各后端缓存分段词典索引，计入 `pg_fts.doclen_cache_mb` 预算；放不下的分段回退到页链查找。修改该预算现可生效，无需等待分段目录变化。本版不改变索引格式，不新增 SQL 对象，也不要求 REINDEX。缓存重建期间的取消或错误可能导致内存一直保留到该后端退出；这一已知限制不改变查询结果。
