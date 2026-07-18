## 用法

来源：

- [已复核 commit 的 Quria README](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/README.md)
- [已复核 commit 的 quria.control](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/quria.control)
- [Cargo 软件包元数据](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/Cargo.toml)
- [上游演示 SQL](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/demo.sql)
- [Bootstrap 与预加载行为](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/sql/_bootstrap.sql)
- [公开函数与钩子](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/lib.rs)
- [全文操作符实现](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/fulltext/operators.rs)
- [磁盘倒排索引实现](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/fulltext/index/inverted/inverted.rs)
- [已复核 commit 的向量模块](https://github.com/iantbutler01/Quria/blob/347740b52d85f5f64aa513464afc18b3c8776d0b/src/vector/hnsw/mod.rs)

`quria` 是一个版本 `0.0.0` 的 pgrx BM25 风格全文搜索原型。它定义 `quria.fulltext` 和 `quria.query` 类型、`quria_fts` 索引访问方法、`~>` 匹配操作符以及 `quria.ft_score`。

### 原型全文搜索

以超级用户安装后，应重新连接，使 bootstrap 创建的数据库级会话预加载设置生效：

```sql
CREATE EXTENSION quria;

CREATE TABLE docs (body quria.fulltext);
INSERT INTO docs VALUES ('tepid text turtle'), ('fast green turtle');
CREATE INDEX docs_body_quria ON docs USING quria_fts (body);

SELECT body,
       quria.ft_score(ctid, '{"query":"turtle","column":"body"}'::quria.query) AS score
FROM docs
WHERE body ~> '{"query":"turtle","column":"body"}'::quria.query
ORDER BY score DESC;
```

### 注意事项

- README 宣称支持 HNSW 向量搜索，但已复核向量模块没有实现。只有实验性全文路径包含实质代码。
- Bootstrap 会修改整个数据库的 `session_preload_libraries`，并把模式 `quria` 的全部权限授予 `PUBLIC`。在任何共享测试系统中，都应先审查并收紧这两项变更。
- 索引数据被序列化到 PostgreSQL 操作系统用户本地的 `.quria/fulltext` 目录，位于普通 PostgreSQL 关系存储和 WAL 之外。标准备份、复制、崩溃恢复和事务保证不会自动覆盖这些文件。
- 源码标为 WIP，没有 release tag，使用旧版内部索引 API，并包含许多 panic 路径。Cargo feature 覆盖 PostgreSQL 11 至 15，默认版本为 13；这不构成生产兼容性声明。
