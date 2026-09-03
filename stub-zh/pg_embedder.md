## 用法

来源：

- [官方文档](https://github.com/riclab/pg_embedder/blob/230cfc5f1785975d5e2fe78c2f2646be953c649a/README.md)
- [扩展控制文件](https://github.com/riclab/pg_embedder/blob/230cfc5f1785975d5e2fe78c2f2646be953c649a/pg_embedder.control)
- [构建清单](https://github.com/riclab/pg_embedder/blob/230cfc5f1785975d5e2fe78c2f2646be953c649a/Cargo.toml)

`pg_embedder` 使用 Rust、Candle 与打包模型在数据库后端内执行文本嵌入推理。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_embedder`：

```sql
CREATE EXTENSION pg_embedder;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- index side
UPDATE docs SET embedding = embed_encode(content);

-- search side, correct for every model
SELECT id, content, cosine_similarity(embedding, embed_query('how does vector search work')) AS score
FROM docs
ORDER BY score DESC
LIMIT 5;
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `embed_encode` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `embed_text` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `docs` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `cosine_similarity` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `embed_init` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `embed_model` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `embed_query` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `embed_info` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 13, 14, 15, 16, 17, 18；不要推断未列出的主版本。
