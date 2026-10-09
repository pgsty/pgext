## 用法

来源：

- [README.md](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/README.md)
- [bm25_native--1.0.sql](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/bm25_native--1.0.sql)
- [Makefile](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/Makefile)
- [bm25_native.control](https://github.com/pgexperts/pgx-bm25/blob/e0e1820586a53b1683c616f16ee0e64efcebdfff/bm25_native.control)

`bm25_native` 提供原生全文检索索引和 BM25 排序，索引存储在 PostgreSQL 索引页中，通过常规 WAL 与清理机制维护。v1.0.0 发布安装的 SQL 扩展版本为 1.0，正式支持 PostgreSQL 17 和 18。

### 核心用法

```sql
CREATE EXTENSION bm25_native;
CREATE TABLE bm25_demo (id integer PRIMARY KEY, body text);
INSERT INTO bm25_demo VALUES (1, 'quick fox'), (2, 'quiet database');
CREATE INDEX ON bm25_demo USING bm25_native (body);
SELECT id, bm25_score(ctid) AS score
FROM bm25_demo
WHERE body @@@ 'quick fox'
ORDER BY body &@@ 'quick fox'
LIMIT 10;
```

### 运行边界

应将 `@@@` 匹配条件与 `&@@` 排序一起使用，才能进行索引 top-k 检索；仅排序不会得到有用的排序索引扫描。`bm25_score` 读取当前排序扫描的分数；涉及多个扫描或分区子表时，应使用限定查询的重载，必要时再限定表身份。JSONB 查询构造函数包括 `bm25_term`、`bm25_boolean`、`bm25_phrase`、`bm25_wildcard` 和 `bm25_boost`。多列索引匹配与标量过滤的语义可能不同，行安全策略也可能触发这种差异；不支持的字段限定或短语过滤会报错。安装不受信任，安装后不允许迁移模式，PostgreSQL 19 的 CI 仍属实验验证。升级索引前应遵循对应版本的磁盘格式兼容与维护说明。
