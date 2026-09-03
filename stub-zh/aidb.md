## 用法

来源：

- [AIDB 官方文档](https://www.enterprisedb.com/docs/aidb/7/)
- [官方配置指南](https://www.enterprisedb.com/docs/aidb/7/install/configure/)
- [官方 pipeline 指南](https://www.enterprisedb.com/docs/aidb/7/data-pipelines/)

`aidb` 是 EDB 的 in-database AI pipeline 扩展，用于对 PostgreSQL 数据执行解析、切块、embedding、index、retrieval、reranking 与 generation。

### 启用与访问

安装 AIDB 7，预加载、重启，并通过 `CASCADE` 创建扩展及其所需 `vector` 依赖：

```ini
shared_preload_libraries = 'aidb'
```

```sql
CREATE EXTENSION aidb CASCADE;
GRANT aidb_users TO app_user;
```

`aidb_users` 授予 AIDB routine 访问权，但不授予应用表权限；source/target table privilege 必须单独授予。PGD 用户创建后还必须调用 `aidb.bdr_setup()`。

### 创建并查询 Pipeline

```sql
SELECT aidb.create_pipeline(
  name               => 'support_kb',
  source             => 'support_articles',
  source_key_column  => 'id',
  source_data_column => 'body',
  step_1             => 'KnowledgeBase',
  step_1_options     => aidb.knowledge_base_config(
    model       => 'bge-small-en-v1.5-f16',
    data_format => 'Text'
  ),
  auto_processing    => 'Disabled'
);

SELECT aidb.run_pipeline('support_kb');
SELECT * FROM aidb.retrieve_text('support_kb', 'connection exhaustion', 3);
```

### 资源与安全边界

模型可能在首次使用时下载，推理会消耗 backend/background-worker CPU、内存与存储。应限制 pipeline role、source-table privilege、model provenance、network egress、concurrency 与 threading。外部文件或 object storage 还依赖 EDB `pgfs`；由于该 provider identity 与同名但无关的开源扩展冲突，必须验证安装软件包，不能只依赖名称。
