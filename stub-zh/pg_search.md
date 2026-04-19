## 用法

- 来源：[ParadeDB extension install docs](https://docs.paradedb.com/deploy/self-hosted/extension)，[quickstart](https://docs.paradedb.com/documentation/getting-started/quickstart)，[v0.23.0 release](https://github.com/paradedb/paradedb/releases/tag/v0.23.0)，[pg_search README](https://github.com/paradedb/paradedb/blob/dev/pg_search/README.md)

`pg_search` 是 ParadeDB 提供的、基于 BM25 的 PostgreSQL 搜索扩展。上游 README 说明支持从 PostgreSQL 15 开始，而 v0.23.0 的自托管安装文档仍要求在 `CREATE EXTENSION` 之前先 preload 该库。

### 启用并创建扩展

```conf
shared_preload_libraries = 'pg_search'
```

```sql
CREATE EXTENSION pg_search;
```

v0.23.0 的自托管扩展文档说明提供了面向 Postgres 15+ 的预编译二进制包。

### 创建 BM25 索引

quickstart 示例使用 `bm25` access method，并要求指定唯一键字段：

```sql
CREATE INDEX search_idx ON mock_items
USING bm25 (id, description, category, rating)
WITH (key_field = 'id');
```

v0.23.0 release 还提到，现在可以按字段调节 BM25 的 `k1` 和 `b` 参数。

### 查询操作符与辅助函数

当前 quickstart 使用以下查询操作符：

- `|||`：析取匹配，等价于 `term1 OR term2`。
- `&&&`：合取匹配，等价于 `term1 AND term2`。

示例：

```sql
SELECT description, rating
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY rating
LIMIT 5;
```

```sql
SELECT description, pdb.score(id)
FROM mock_items
WHERE description &&& 'running shoes'
ORDER BY score DESC
LIMIT 5;
```

```sql
SELECT description, pdb.snippet(description), pdb.score(id)
FROM mock_items
WHERE description ||| 'running shoes'
ORDER BY score DESC
LIMIT 5;
```

### 说明

开发分支 README 已把安装和用法细节指向官方文档站，而不是在 README 中内联维护 SQL 细节。因此，对当前 `pg_search` 语法而言，quickstart 才是最权威的用法来源；它反映的是 0.20 之后的 API，而不是一些次级资料里仍能看到的旧 `@@@` 示例。
