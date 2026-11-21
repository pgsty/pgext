## 用法

来源：

- [README.en.md](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/README.en.md)
- [pgwarc_lance.control](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/pgwarc_lance.control)
- [Cargo.toml](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/Cargo.toml)
- [lib.rs](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/src/lib.rs)
- [pgwarc_lance--0.1.26--0.1.27.sql](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/sql/pgwarc_lance--0.1.26--0.1.27.sql)
- [SECURITY.md](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/SECURITY.md)
- [sql-cli-contract.md](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/docs/sql-cli-contract.md)
- [LICENSE](https://github.com/bjm1244/pgwarc_lance/blob/v0.1.27/LICENSE)


`pgwarc_lance` 0.1.27 面向 WARC/Web 归档，将基于表的 BM25、本地 Lance 向量检索和倒数排名融合结合起来。它是研究原型。本次发行版验证范围为 Linux x86-64 与 PostgreSQL 16.15，其他构建选项不代表已支持相应发布目标。

### 索引与检索文本

以超级用户安装。扩展不可迁移模式，不要求预加载。新安装会撤销所有函数的 PUBLIC 执行权限；只向可信应用角色授予必需的函数和表权限。

```sql
CREATE EXTENSION pgwarc_lance;
SELECT bm25_index_document(1, 'PostgreSQL archive search');
SELECT bm25_index_document(2, 'Database backup and recovery');
SELECT * FROM bm25_search('archive', 10);
```

BM25 数据保存在 `pgwarc_lance` 模式中。`tokenize_korean` 使用 CJK 双字切分与 ASCII 分词。`bm25_doc_count` 和 `bm25_avg_doc_len` 提供语料统计。`hybrid_search` 合并文本与向量排名；`hybrid_warc_search_filtered` 对两侧候选集应用相同的资格过滤条件。

### Lance 与发件箱

`lance_create_table`、`lance_insert_many` 和 `lance_vector_search` 管理并查询本地数据集。管理员应将 `pgwarc_lance.allowed_uri_prefix` 设置为已有的本地数据目录，文件访问函数只开放给可信角色。不支持远程或对象存储 Lance，也不支持多个独立写入者。

直接 Lance 写入会立即改变外部文件，PostgreSQL 回滚不会撤销它。需要协调摄取时，`lance_outbox_enqueue_revision` 在调用方 SQL 事务中保存向量载荷和文档修订号。提交后，外部工作进程以独立自动提交调用执行 `lance_outbox_process_next`。这提供最终一致性，不提供 PostgreSQL 与 Lance 之间的原子提交。

`lance_outbox_status` 显示待处理与失败工作。依赖向量写后读行为前应确认任务已完成。`lance_document_delete` 核对期望修订号，并写入墓碑和有序删除事件；还要求破坏性操作权限及对应管理员设置。不要直接修改内部修订或队列表。

### 备份与升级

Lance 文件不受 PostgreSQL WAL 和 `pg_dump` 保护。应停止导入器和工作进程、核对队列，并保存相互匹配的 SQL 数据和 Lance 目录。恢复两部分后才能继续工作。单独回退 Lance 版本不会清除 SQL 墓碑，也不会恢复逻辑文档。

提供的 0.1.26 至 0.1.27 迁移仅推进版本，不修改 SQL 对象或授权：

```sql
ALTER EXTENSION pgwarc_lance UPDATE TO '0.1.27';
```

升级应使用匹配文件和可用的升级链。上游尚未完成独立相关性评估、持续负载容量验证及生产支持，使用前应验证这些边界。
