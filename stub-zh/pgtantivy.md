## 用法

来源：

- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/README.md)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/Cargo.toml)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/pgtantivy.control)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/src/lib.rs)
- [Official source](https://github.com/alitrack/pgtantivy/blob/v0.5.0/src/am.rs)

`pgtantivy` 0.5.0 在 PostgreSQL 17–19 中提供 Tantivy BM25 检索，支持中文 jieba 和英文分词。它提供显式快照索引函数、PostgreSQL 索引访问方法，以及通过 pgvector 实现的可选混合检索。安装需要超级用户，扩展不受信任且不可重定位。

### 快照检索

```sql
CREATE EXTENSION pgtantivy;
CREATE TABLE docs (doc_id bigint PRIMARY KEY, body text NOT NULL);
INSERT INTO docs VALUES (1, 'PostgreSQL 数据库全文检索');
SELECT pgtantivy_create(
    'docs_idx',
    (SELECT json_agg(json_build_object('id', doc_id, 'text', body))::text FROM docs),
    'chinese');
SELECT id, score FROM pgtantivy_search('docs_idx', '数据库 检索', 10);
```

查询词之间使用 AND 语义。新后端使用命名索引前，应调用 `pgtantivy_load('docs_idx')`。`pgtantivy_add` 按 ID 插入或替换 JSON 文档，`pgtantivy_remove` 删除指定 ID，`pgtantivy_snippet` 生成高亮片段。

### 存储与同步

索引文件位于可配置的 `pgtantivy.index_dir`，默认使用 PostgreSQL 数据目录内的子目录；已加载索引的注册表属于各自后端。手动构建的索引是快照，普通表写入需要显式同步，或使用文档中的 `pgtantivy_sync_trigger`。触发器暂存操作，在 PRE_COMMIT 阶段批量写入；回滚会清空待处理操作。不能据此推定这些独立 Tantivy 文件具有 PostgreSQL WAL/MVCC 或崩溃原子性保证。

`pgtantivy_drift_check` 将记录的 LSN 与全实例 WAL 进度比较。即使只有无关表变化，也可能报告陈旧信号。`pgtantivy_touch` 只推进锚点，不会同步表内容。

### 其他接口与边界

`pgtantivy` 索引访问方法和 `@@@` 运算符支持位图索引扫描。可选的 `pgtantivy_hybrid` 将 BM25 与 pgvector 结果结合；核心扩展没有将 pgvector 声明为强制依赖。

本地与 S3 索引包支持导出、导入及版本选择。FNV 校验和可检测意外损坏，但不能认证不可信内容。SQL 函数可接收服务器路径及对象存储凭据，应将执行和文件访问权限限制给可信角色，避免将凭据写入会被记录的 SQL。侧存索引应与源数据一起备份并验证恢复，或准备重建方案。
