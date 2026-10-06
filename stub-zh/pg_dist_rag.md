## 用法

来源：

- [src/postgres/yb-extensions/pg_dist_rag/README.md](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/README.md)
- [src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1.sql](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1.sql)
- [src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1--0.0.2.sql](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag--0.0.1--0.0.2.sql)
- [src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag.control](https://github.com/yugabyte/yugabyte-db/blob/771f263d34aa95952ca311bb4c63e1d7fa8e2159/src/postgres/yb-extensions/pg_dist_rag/pg_dist_rag.control)

`pg_dist_rag` 0.0.2 管理来源、向量索引、文档处理队列及列嵌入注册。核对的发行源码属于 YugabyteDB；这些 SQL 目录本身不会运行嵌入服务。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_dist_rag;
SELECT dist_rag.create_source(r_source_uri := 's3://example-bucket/documents/');
SELECT dist_rag.init_vector_index(
    r_index_name := 'knowledge',
    r_embedding_model_params := '{"dimensions":1536}'::jsonb
);
SELECT * FROM dist_rag.work_queue;
```

### 运行边界

先安装 `vector`。对象位于 `dist_rag`，扩展不可迁移模式，也没有声明预加载要求。本条目不为该 YugabyteDB 组件推定原版 PostgreSQL 的主版本兼容矩阵。

`dist_rag.create_source` 将来源任务入队，`dist_rag.init_vector_index` 创建嵌入存储，`dist_rag.build_index` 将预处理任务入队。队列需要外部工作进程及配置好的模型提供方处理；确认索引可用前应检查来源、文档和流水线状态。

0.0.2 通过 `dist_rag.create_column_embedding_mapping` 增加列映射，并由 `dist_rag.init_column_embedding` 激活；进度及暂停、恢复接口管理注册状态。部分管理函数以定义者权限执行，应将执行权限和直接访问表的权限限制给合适的操作人员与工作进程，并保护存储的提供方配置和凭据。
