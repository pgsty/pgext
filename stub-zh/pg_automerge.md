## 用法

来源：

- [README.md](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/README.md)
- [sql/snapshots/pg_automerge--0.1.0.sql](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/sql/snapshots/pg_automerge--0.1.0.sql)
- [docs/src/pages/operations/limitations.mdx](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/docs/src/pages/operations/limitations.mdx)
- [docs/src/pages/reference/configuration.mdx](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/docs/src/pages/reference/configuration.mdx)
- [pg_automerge.control](https://github.com/TheKnarf/pg-automerge/blob/77ce017eeb3c33cb47e011cabac99a6976967cf3/pg_automerge.control)

`pg_automerge` 将 Automerge 文档保存为 PostgreSQL 18 原生类型。应用负责生成变更与管理 actor ID，扩展负责存储、合并和查询，不编辑文档或自行生成变更。

### 核心用法

```sql
CREATE EXTENSION pg_automerge;
CREATE TABLE docs (id uuid PRIMARY KEY, doc automerge NOT NULL);
PREPARE store_document(uuid, bytea) AS
INSERT INTO docs (id, doc) VALUES ($1, $2)
ON CONFLICT (id) DO UPDATE SET doc = merge(docs.doc, EXCLUDED.doc);
SELECT doc->>'title' FROM docs WHERE doc @> '{"status":"open"}';
```

### 运行边界

输入必须是有效的 Automerge 二进制存档，而不是 JSON 文本。隐式 `jsonb` 转换提供当前文档，`merge()` 与 `merge_agg()` 合并历史。增量变更需要已有文档，因此首次插入必须使用完整存档。每个写入者应使用独立 actor ID，识别历史时比较 heads 而不是序列化字节。`pg_automerge.max_load_memory` 限制单文档的估算内存，但分配不受 PostgreSQL 内存核算约束，也不限制并发会话总量。原生加载或合并无法及时响应中断。安装要求超级用户，应信任写入者并约束输入资源成本。
