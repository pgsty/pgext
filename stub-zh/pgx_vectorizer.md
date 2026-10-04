## 用法

来源：

- [FEP 18 quick-start guide](https://www.postgresql.fastware.com/knowledge-base/quick-start-guides)
- [FEP 18 knowledge data management manual](https://www.postgresql.fastware.com/hubfs/_Global/Manuals/FEP-v18forx86-KnowledgeDataManagementUserGuide.pdf)

`pgx_vectorizer` 3.0 随 Fujitsu Enterprise Postgres 18 提供，用于自动生成嵌入与语义检索，依赖 `ai`、`vector` 和 `plpython3u`。

### 核心用法

先配置厂商 Python 运行环境、工作进程容量及嵌入服务，将 `pgx_vectorizer` 加入 `shared_preload_libraries` 并重启，再在目标数据库中启用：

```sql
CREATE EXTENSION pgx_vectorizer CASCADE;
SELECT pgx_vectorizer.start_vectorize_scheduler();
```

以扩展所有者身份，通过 `pgx_vectorizer.set_worker_setting` 登记工作进程的登录角色，并配置连接权限与密码文件。准备好提供 `all-minilm` 模型且可访问的 Ollama 服务后，为含有主键及文本列的表定义向量化：

```sql
CREATE TABLE sample_table (id bigint PRIMARY KEY, contents text);
INSERT INTO sample_table VALUES (1, 'PostgreSQL supports streaming replication.');
SELECT pgx_vectorizer.pgx_create_vectorizer(
  'sample_table'::regclass,
  destination => 'sample_embeddings',
  embedding => ai.embedding_ollama('all-minilm', 384),
  chunking => ai.chunking_recursive_character_text_splitter('contents'),
  scheduling => pgx_vectorizer.schedule_vectorizer(interval '1 hour')
);
```

向量化完成后，查询生成的视图：

```sql
SELECT * FROM pgx_vectorizer.pgx_similarity_search(
  'sample_embeddings'::regclass, 'database replication', 5, '<=>');
```

### 运行边界

后台工作进程与前台检索需要分别配置嵌入服务访问。检索用户需要 `ai` 模式中函数的执行权限。应保护 `pgx_vectorizer.worker_setting_table` 中的凭据。重命名源表、修改主键或文本列定义后，需要重新定义向量化。相关维护期间应暂停调度。此功能需要 Fujitsu 运行环境。
