## 用法

来源：

- [官方 pg_bleve README](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/README.md)
- [1.0.0 版扩展 SQL](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/pg_bleve--1.0.0.sql)
- [Go 实现](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/main.go)
- [PostgreSQL 索引访问方法包装器](https://github.com/MagicFun1241/pg_bleve/blob/2757cbd1e3db950a723b785611a89d42b4f8c548/pg_wrapper.c)

`pg_bleve` 是一个面向 PostgreSQL 17 的早期实验，通过 cgo 调用 Go Bleve 搜索库。其实际可用接口是一组手动索引管理函数，数据文件位于 `/tmp/pg_bleve_indexes`；声明的 PostgreSQL `bleve` 索引访问方法只是占位实现，不会维护可搜索的 SQL 索引。

### 手动索引流程

```sql
CREATE EXTENSION pg_bleve;

SELECT bleve_create_index('products');

SELECT bleve_index_document(
    'products',
    'sku-42',
    '{"title":"Wireless headphones","category":"electronics"}'
);

SELECT bleve_search('products', 'category:electronics');
SELECT 'products' @@@ 'title:wireless';
```

`bleve_create_index(index_name, config_json)` 创建磁盘上的 Bleve 索引，但此版本会忽略 `config_json`。`bleve_index_document(index_name, doc_id, document_json)` 插入或替换一份 JSON 文档。`bleve_search(index_name, query_string)` 以 JSON 文本返回搜索结果，并硬编码最多返回十条命中。

`@@@` 运算符执行直接查询；`|||` 将空白分隔的词组合为析取条件，`&&&` 将其组合为合取条件。`bleve_populate_index(index_name, table_name)` 会序列化表行，但最多处理 1,000 行；`bleve_populate_index_batch(index_name, table_name, batch_size)` 是实验性的分批版本。使用任一辅助函数后都应核对实际数量与结果。

### 关键限制

不要依赖 `CREATE INDEX ... USING bleve`。此修订版的构建回调只创建空的外部索引而不索引堆元组，插入回调不执行索引写入，扫描不返回元组，VACUUM 回调也不做维护。因此 PostgreSQL 的更新和删除不会自动同步。

Bleve 文件位于 PostgreSQL 事务、MVCC、WAL、备份、恢复和权限控制之外。`/tmp` 可能被清理，索引名直接参与文件系统路径，同一主机不同数据库或集群之间也可能重名。若干具有副作用或读取文件系统的函数还被声明为 `IMMUTABLE`，与实际行为不符。此扩展只能视为一次性实验代码；函数调用成功不代表数据库级持久性或一致性。
