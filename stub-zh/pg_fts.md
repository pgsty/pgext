## 用法

来源：

- [1.14.0 upgrade](https://api.pgxn.org/src/pg_fts/pg_fts-1.14.0/pg_fts--1.13.0--1.14.0.sql)
- [README.md](https://api.pgxn.org/src/pg_fts/pg_fts-1.14.0/README.md)
- [pg_fts.control](https://api.pgxn.org/src/pg_fts/pg_fts-1.14.0/pg_fts.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_fts/pg_fts-1.14.0/CHANGELOG.md)
- [pg_fts--1.14.0.sql](https://api.pgxn.org/src/pg_fts/pg_fts-1.14.0/pg_fts--1.14.0.sql)

`pg_fts` 1.14.0 通过专用 fts 倒排索引提供 BM25/BM25F 排序的全文检索。支持 PostgreSQL 17 与 18；PostgreSQL 19/master 的 CI 属于尽力支持。控制文件标记可信且可重定位，核心使用无需共享预加载。

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

### 复合邻近查询与 1.13.0 修复

查询新增无序邻近 `w/N`、有序邻近 `p/N`、精确间隔 `<N>`、相邻 `<->` 和同位置 `<0>`。两侧可以是词项、短语、OR 分组或其他邻近分组；不接受直接作为操作数的 AND/NOT。引号短语保留被移除停用词留下的间隔。带排序的前缀、模糊和正则查询现会包含全部匹配行，待合并文档也参与排序；仅通过扩展项匹配的文档仍可能得零分，因为评分依据字面查询词项。

本版修复畸形二进制查询导致的崩溃、非 ASCII 模糊匹配错误，以及并发合并期间位置扫描的重复释放。安装匹配文件后执行 `ALTER EXTENSION pg_fts UPDATE TO '1.13.0'`；索引格式不变，无需例行 REINDEX。

### 1.14.0 的清理与写入负载修复

1.14.0 用 `maintenance_work_mem` 限制倒排记录清理批次，并在并行 VACUUM 中跳过索引尾部截断。暂未释放的空间仍可复用，之后可由非并行 VACUUM 或 `fts_vacuum()` 回收。新增的 `fts_matchsel` 估算器改善排序检索的执行计划选择。`pg_fts.pending_flush_threshold` 默认在待合并文档达到 20000 时触发刷新，设为 0 则禁用；插入路径只尝试获取维护锁，不等待锁释放。

安装匹配文件后执行 `ALTER EXTENSION pg_fts UPDATE TO '1.14.0'`，更新匹配操作符的估算函数。索引格式不变，无需例行 REINDEX。
