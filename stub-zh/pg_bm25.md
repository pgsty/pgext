## 用法

来源：

- [Fork README](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/README.md)
- [Extension control file](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/pg_bm25.control)
- [Index option implementation](https://github.com/devflowinc/pg_bm25_fork/blob/be44998446d1095e67970d42141cb2874c945289/pg_bm25/src/index_access/options.rs)

pg_bm25 是一个早期的 ParadeDB/Tantivy 全文检索扩展，为 PostgreSQL 加入 BM25 索引访问方法和 Lucene 风格查询。本页描述目录所收录的历史分叉，其 API 与当前 ParadeDB 扩展有实质差异；不能用当前 pg_search 文档替代这个版本的说明。

### 核心流程

control 文件要求由超级用户安装，并把对象固定在 `paradedb` 模式中。需要在表的复合行上创建索引，并声明要索引的字段：

```sql
CREATE EXTENSION pg_bm25;

CREATE TABLE articles (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  title text,
  body text,
  rating integer
);

CREATE INDEX articles_bm25
ON articles
USING bm25 ((articles.*))
WITH (text_fields='{"title": {}, "body": {}}',
      numeric_fields='{"rating": {}}');

SELECT id, title, paradedb.rank_bm25(ctid)
FROM articles
WHERE articles @@@ 'body:postgres AND rating:[4 TO *]'
ORDER BY paradedb.rank_bm25(ctid) DESC;
```

### 重要对象

- `bm25` 是由 Tantivy 文件支持的索引访问方法。
- `@@@` 针对行所在表关联的 BM25 索引执行文本查询。
- `text_fields`、`numeric_fields`、`boolean_fields` 和 `json_fields` 是 JSON 索引选项。
- `paradedb.rank_bm25` 返回当前查询中元组的分数。
- `paradedb.minmax_bm25` 针对指定索引和查询归一化分数。
- `paradedb.schema_bm25` 报告 Tantivy 字段模式。
- `paradedb.aggregation` 执行 Tantivy 聚合请求并返回 JSON。

### 运维说明

该分叉的 README 声称支持 PostgreSQL 12 及以上，但其预构建包与测试范围窄得多。索引数据由扩展代码维护在普通堆存储之外，因此依赖它之前必须测试崩溃恢复、备份还原、复制、升级、VACUUM 行为和并发写入。这个仓库是使用旧 0.3.x API 的历史分叉；新部署应优先选择仍受维护的检索扩展。
