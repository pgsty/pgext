## 用法

来源：

- [官方文档](https://github.com/hank-cp/pg_vector_embedding/blob/c53bf7f93ad9f4cc27fa657b40749488a1000eae/README.md)
- [扩展控制文件](https://github.com/hank-cp/pg_vector_embedding/blob/c53bf7f93ad9f4cc27fa657b40749488a1000eae/pg_vector_embedding.control)
- [官方仓库](https://github.com/hank-cp/pg_vector_embedding)

`pg_vector_embedding` 通过外部 HTTP 服务与后台队列触发生成向量嵌入。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_vector_embedding`：

```sql
CREATE EXTENSION pg_vector_embedding CASCADE;
```

经审查的控制文件或官方流程要求 `http`, `pg_background_queue`, `vector`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- 1. Setup
CREATE EXTENSION pg_vector_embedding CASCADE;

ALTER DATABASE mydb SET pg_vector_embedding.embedding_url = 'https://api.example.com/v1/embeddings';
ALTER DATABASE mydb SET pg_vector_embedding.embedding_api_key = 'sk-xxxxx';

\c  -- Reconnect to apply settings

-- 2. Create and register table
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    embedding VECTOR(1024)
);

SELECT ve_enable('public', 'articles', ARRAY['title', 'content'], 'embedding');

-- 3. Insert data (embeddings computed automatically in background)
INSERT INTO articles (title, content) VALUES
    ('PostgreSQL Extensions', 'Learn how to build powerful PostgreSQL extensions'),
    ('Vector Search', 'Implementing semantic search with pgvector');

-- 4. Wait for background processing (or check if embeddings are ready)
SELECT COUNT(*) FROM articles WHERE embedding IS NOT NULL;

-- 5. Perform similarity search
WITH search_query AS (
    SELECT ve_compute_embedding(
        '{"title":"PostgreSQL","content":"tutorial"}'::text
    ) AS query_embedding
)
SELECT id, title, embedding <-> query_embedding AS distance
FROM articles, search_query
WHERE embedding IS NOT NULL
ORDER BY distance
LIMIT 5;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `ve_compute_embedding` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ve_enable` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ve_trigger` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ve_compact_row_data` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `ve_disable` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
