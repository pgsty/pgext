## 用法

来源：

- [官方文档](https://github.com/wahicona/pg_sag_rag/blob/9bd58193d84cff7762794b75c56f40b4cbdc387a/README.md)
- [扩展控制文件](https://github.com/wahicona/pg_sag_rag/blob/9bd58193d84cff7762794b75c56f40b4cbdc387a/pg_sag_rag.control)
- [官方仓库](https://github.com/wahicona/pg_sag_rag)

`pg_sag_rag` 纯 SQL 实现的多跳事件/实体检索、查询路由与库内 RAG 评估。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_sag_rag`：

```sql
CREATE EXTENSION pg_sag_rag CASCADE;
```

经审查的控制文件或官方流程要求 `vector`, `pg_trgm`。只有服务器已安装这些扩展文件时，`CASCADE` 才会成功。

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT sag_rag.add_evaluation_set('my-rag-eval');
SELECT sag_rag.add_evaluation_question(1, 'How much is AGI Bar foam?', '[0.86,0.11,0.10]'::vector);
SELECT sag_rag.link_evaluation_answer_event(1, 2);

SELECT sag_rag.run_evaluation_hybrid(1, p_top_k => 1);
SELECT sag_rag.run_evaluation_multihop(1, p_seed_k => 1, p_top_k => 10);
SELECT sag_rag.run_evaluation_auto(1);

SELECT run_id, strategy, parameters
FROM sag_rag.evaluation_run
ORDER BY run_id;

SELECT * FROM sag_rag.recall_at_k(1, 1);
SELECT * FROM sag_rag.recall_at_k(2, 10);
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `sag_rag.event` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `sag_rag.entity` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `sag_rag.create_hnsw_indexes` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `sag_rag.document` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `sag_rag.event_entity` | TABLE | 扩展自有表；备份与升级时需计入其中数据。 |
| `sag_rag.recall_at_k` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `sag_rag.run_evaluation_auto` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `sag_rag.search_events_auto` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 14, 15, 16, 17；不要推断未列出的主版本。
