## 用法

来源：

- [1.6.0 README](https://github.com/timescale/pg_textsearch/blob/v1.6.0/README.md)
- [1.6.0 control](https://github.com/timescale/pg_textsearch/blob/v1.6.0/pg_textsearch.control)
- [1.6.0 SQL](https://github.com/timescale/pg_textsearch/blob/v1.6.0/sql/pg_textsearch--1.6.0.sql)
- [1.6.0 upgrade guide](https://github.com/timescale/pg_textsearch/blob/v1.6.0/UPGRADING.md)
- [1.6.0 release](https://github.com/timescale/pg_textsearch/releases/tag/v1.6.0)

`pg_textsearch` 1.6.0 通过 `bm25` 访问方法与 `<@>` 运算符，在 PostgreSQL 17–19 中提供 BM25 排序的文本检索。

### 启用与查询

管理员须将库加入已有预加载列表并重启 PostgreSQL。安装扩展需要提升权限。

```ini
shared_preload_libraries = 'pg_textsearch'
```

```sql
CREATE EXTENSION pg_textsearch;
CREATE TABLE documents (id bigserial PRIMARY KEY, content text);
INSERT INTO documents(content) VALUES
    ('PostgreSQL is a database system'),
    ('BM25 ranks full text search results');
CREATE INDEX docs_idx ON documents USING bm25(content)
WITH (text_config = 'english');

SELECT id, content <@> 'database system' AS score
FROM documents
ORDER BY content <@> 'database system'
LIMIT 5;
```

分数是 BM25 值的相反数，越小越靠前。部分索引、歧义表达式、预备或动态查询文本及 PL/pgSQL 应明确指定索引。独立评分要求对索引表或列具有 SELECT 权限。

```sql
SELECT id FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

### 过滤与索引选项

`text_config` 必填；`k1` 默认为 1.2，`b` 默认为 0.75。支持不可变文本表达式、文本类数组、部分索引及分区表。查询须重复索引表达式，部分索引还须有匹配的过滤条件。中文文本可选用 `zhparser` 等解析器。

```sql
SET default_text_search_config = 'english';
SELECT id FROM documents
WHERE content @@ to_tsquery('english', 'database & !mysql')
ORDER BY content <@> 'database'
LIMIT 5;
```

布尔查询支持 AND/OR/NOT、短语、前缀与权重。短语和权重条件可能需要表行复查。文本搜索配置须与索引一致；使用缓存的布尔查询计划期间若配置发生变化，应释放并重新准备语句。

1.6.0 可在排序扫描内部执行符合条件的常量布尔谓词，前提是每个匹配结果必然包含某个排序词项。其他形式，包括查询词未知的通用计划，继续由 PostgreSQL 过滤和排序。后置过滤的排序扫描会扩大评分批次，直到获得足够的合格行或遍历完匹配项。应检查 EXPLAIN；较小的 LIMIT 不保证工作量小。

### 维护与使用边界

持久化内存表位于受 WAL 保护的索引页中。共享内存缓存可选；`pg_textsearch.memory_limit` 默认 2GB，是准入预算的近似值，并非严格上限。`pg_textsearch.default_limit` 默认 1000，在 SQL 没有限制条数时决定初始批次大小。刷写和串行 VACUUM 可能执行同步压实。

```sql
SELECT bm25_spill_index('docs_idx');
SELECT bm25_force_merge('docs_idx');
```

修改索引的维护函数要求索引所有权，不处理分区父索引，也不能在恢复期间执行。已经发布的物理合并不会被 ROLLBACK 撤销。可选的受管理后台压实需要上游规定的额外配置，内联／手动模式不依赖该配置。执行查询的热备库需要 `hot_standby_feedback = on`，以便在活动快照期间延迟物理页复用。

`bm25` 访问方法名与 `pg_search`、`vchord_bm25` 等其他提供者冲突。RLS 隐藏行，但 BM25 语料统计仍反映全部索引行。超级用户可将 `pg_textsearch.allow_rls` 设为 off，拒绝新建或重建与 RLS 组合的索引；该设置不会禁用已有索引。

### 升级到 1.6.0

安装匹配的库与 SQL 文件，重启 PostgreSQL，然后更新每个数据库：

```sql
ALTER EXTENSION pg_textsearch UPDATE TO '1.6.0';
```

1.6.0 保持 1.5.1 的索引格式，并增加感知分区的评分函数。旧版涉及 HOT 链的并行构建可能遗漏行，中断的 VACUUM 也可能留下错误计数或位图。受影响的索引必须重建；升级二进制与 SQL 无法修复其内容。

```sql
REINDEX INDEX docs_idx;
```

普通 REINDEX 会阻塞写入。旧索引还须检查所有跨越版本的升级说明：1.3 之前未刷写的数据需在旧二进制下处理；1.5 之前的索引可能在否定布尔查询中漏掉空文档或仅含停用词的文档；1.5 之前的服务器无法读取版本 9 的索引头。在新二进制修改索引头之前须升级备库，降级二进制不会撤销格式变化。
