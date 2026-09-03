## 用法

来源：

- [官方 README](https://github.com/hyperiondb/search/blob/6dc33192199ed18a1be9d269b88a9027c76f2eac/README.md)
- [扩展控制文件](https://github.com/hyperiondb/search/blob/6dc33192199ed18a1be9d269b88a9027c76f2eac/hsearch.control)
- [pgrx 清单](https://github.com/hyperiondb/search/blob/6dc33192199ed18a1be9d269b88a9027c76f2eac/Cargo.toml)

`hsearch` 为 PostgreSQL 18 提供基于 Tantivy 的 `bm25` 索引访问方法，同时把索引字节保存在由 generic WAL 保护的 PostgreSQL 页面中。

### 启用

安装面向 PostgreSQL 18 的 hsearch 0.4.1，将其加入 `shared_preload_libraries`，重启后再创建扩展：

```ini
shared_preload_libraries = 'hsearch'
```

```sql
CREATE EXTENSION hsearch;
```

扩展不可迁移且仅限超级用户安装。其 SQL 对象位于 `hyper` schema。

### 建立并查询索引

第一个索引列是由 `key_field` 指定的稳定键，其余索引表达式是可搜索文本字段。

```sql
CREATE TABLE items (
  id      varchar(24) PRIMARY KEY,
  name    text,
  summary text
);

CREATE INDEX items_bm25 ON items
USING bm25 (
  id,
  (name::hyper.ngram(2,5,'ascii_folding=true')),
  (summary::hyper.ngram(2,5,'ascii_folding=true'))
) WITH (key_field='id');

SELECT id, hyper.score(id) AS score
FROM items
WHERE name &&& 'postgres'
ORDER BY score DESC
LIMIT 20;
```

`&&&` 要求所有生成的 ngram 都匹配，并重新检查堆表可见性。当前事务插入的行在提交后才可搜索。

### 运维与限制

索引通过 WAL 支持崩溃恢复与物理复制，但仍是派生结构：发生损坏或运维策略要求时，应使用 `REINDEX` 或 `hyper.reindex_all()` 重建。`hsearch.max_matches` 限制每个扫描键的候选数量，默认 1,000；提高它会改善深层过滤召回率，但会增加内存与延迟。0.4.1 不支持 unlogged 表，也不支持 `(2,5,'ascii_folding=true')` 之外的 tokenizer 配置。

