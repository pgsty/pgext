## 用法

来源：

- [官方 README](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/README)
- [扩展控制文件](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/setrank.control)
- [1.0 版扩展 SQL](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/setrank--1.0.sql)
- [IDF 表加载器与缓存](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/idf.c)

`setrank` 1.0 为 PostgreSQL 全文检索添加 TF-IDF、覆盖密度 TF-IDF 和评分函数。与内置排名不同，它从调用方维护的表中读取文档频率统计，使常见词位获得较低权重。

### 构建统计表

表必须恰好有两列：词位 `text`，随后是文档数 `int4` 或 `int8`。第一列为 NULL 的行保存文档总数。

```sql
CREATE EXTENSION setrank;

CREATE TABLE docs (
  id bigint GENERATED ALWAYS AS IDENTITY,
  body text,
  fts tsvector
);

INSERT INTO docs(body, fts)
VALUES ('x ray star', to_tsvector('english', 'x ray star'));

CREATE TABLE docs_stat AS
SELECT word AS value, ndoc::bigint AS ndoc
FROM ts_stat('SELECT fts FROM docs');

INSERT INTO docs_stat(value, ndoc)
VALUES (NULL, (SELECT count(*) FROM docs));
```

每行第二列都必须为正，词位必须唯一，而且总数不能小于任何单个文档频率。

### 文档排名

```sql
SELECT id,
       ts_rank_tfidf('docs_stat', fts, plainto_tsquery('english', 'x ray star')) AS rank
FROM docs
WHERE fts @@ plainto_tsquery('english', 'x ray star')
ORDER BY rank DESC;
```

`get_idf(table_name, lexeme)` 返回缓存的逆文档频率。三个排名函数族是 `ts_rank_tfidf`、`ts_rank_cd_tfidf` 和 `ts_score_tfidf`；每个函数族都提供带可选 `float4[]` 权重和归一化标志的重载。

### 缓存与维护注意事项

后端第一次使用时，C 代码会把整个统计表读入 `TopMemoryContext`。它不会失效或刷新这个后端级缓存，因此同一会话在表修改后仍看不到变化。长连接池可能持续返回过时排名。

这些函数声明为 `IMMUTABLE`，但结果实际依赖表内容。不要将其用于表达式索引、生成列或其他要求真正不可变性的场景。语料发生变化后，应重建统计表，并回收所有曾调用 `setrank` 的会话，再信任新评分。

加载器把整个 `table_name` 参数作为单一标识符引用，因此所审计实现预期简单的非模式限定表名，而不是 `schema.table`。应限制谁能用任意表名调用这些函数，并把统计表放在当前搜索路径中。项目已废弃；继续使用前必须在目标 PostgreSQL 版本上验证行为。
