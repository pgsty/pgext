## 用法

来源：

- [Version 1.4.0 README](https://github.com/timescale/pg_textsearch/blob/v1.4.0/README.md)
- [Control file](https://github.com/timescale/pg_textsearch/blob/v1.4.0/pg_textsearch.control)
- [Versioned SQL](https://github.com/timescale/pg_textsearch/blob/v1.4.0/sql/pg_textsearch--1.4.0.sql)
- [Version 1.4.0 release](https://github.com/timescale/pg_textsearch/releases/tag/v1.4.0)
- [Memtable architecture](https://github.com/timescale/pg_textsearch/blob/v1.4.0/docs/memtable_v2.md)

`pg_textsearch` 通过 `bm25` 访问方法和 `<@>` 运算符提供 BM25 全文检索。上游 1.4.0 支持 PostgreSQL 17 和 18，需要预加载并重启。使用本版本前应准备匹配的上游构建；目录中的 Pigsty 软件包基线仍为 1.2.0。

### 构建与查询

```conf
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

分数是 BM25 值的相反数，因此数值越低越靠前。使用 `ORDER BY` 与 `LIMIT` 进行 top-k 检索。独立评分、部分索引或 PL/pgSQL 中应明确指定索引：

```sql
SELECT id FROM documents
ORDER BY content <@> to_bm25query('database system', 'docs_idx')
LIMIT 5;
```

### 索引与查询选项

`text_config` 是必填项，指定 PostgreSQL 文本搜索配置。`k1` 默认是 1.2，`b` 默认是 0.75。`bm25query` 类型和 `to_bm25query(text, text)` 显式携带查询及索引上下文。独立评分需要目标表或索引列的 SELECT 权限。

扩展支持 `text[]`、`varchar[]` 和 `bpchar[]`、不可变文本表达式索引、部分索引及分区表。查询中须使用与索引匹配的表达式。部分索引需要匹配的过滤条件及明确的索引名。1.4.0 改进了带过滤条件的 top-k 执行，但严格的后置过滤仍可能导致结果数少于请求数。

中文分词可配置 `zhparser` 等解析器，再使用相应的文本搜索配置；这是该工作流的可选依赖。对于缺少空白分词边界的超长文本，上游建议由应用确定分块边界，并使用文本数组保存。

### 维护与配置

从 1.3.0 开始，内存表结构存放于索引页中，使用 PostgreSQL 标准 WAL 回放。旧版共享内存中的内存表及其内存上限配置不再适用于本版本。自动压实在刷写过程中同步执行，重写入负载可能受到压实延迟影响。

```sql
SELECT bm25_spill_index('docs_idx');
SELECT bm25_force_merge('docs_idx');
```

强制合并适合在批量加载后执行，持续写入时应谨慎使用。VACUUM 也会刷写待处理的内存表页。相关配置包括：

| 配置 | 默认值 | 用途 |
| --- | --- | --- |
| `pg_textsearch.default_limit` | 1000 | 查询无条数限制时的评分上限 |
| `pg_textsearch.compress_segments` | on | 倒排块压缩 |
| `pg_textsearch.segments_per_level` | 8 | 压实阈值 |
| `pg_textsearch.bulk_load_threshold` | 100000 | 每事务触发刷写的词项数量 |
| `pg_textsearch.memtable_pages_threshold` | 64 | 触发刷写的链页数量 |

并行构建要求 `maintenance_work_mem` 至少为 64 MB，且有可用的并行维护进程。升级时先准备匹配的二进制文件，重启 PostgreSQL，再按发布说明执行扩展更新。

```sql
ALTER EXTENSION pg_textsearch UPDATE;
```

### 使用边界

`bm25` 访问方法名与 `pg_search`、`vchord_bm25` 冲突，不应在同一数据库安装相互冲突的提供者。由于不保存词项位置，扩展不原生支持短语匹配。分数采用各分区自身的统计信息，跨分区可能不可直接比较。超长词项受 PostgreSQL 文本搜索限制影响。固定的 LWLock tranche ID 也可能与其他扩展冲突，导致等待事件名称不准确。
