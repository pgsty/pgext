## 用法

来源：

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)
- [Enable Lakebase Search](https://docs.databricks.com/aws/en/oltp/projects/lakebase-search)
- [Extension guide](https://docs.databricks.com/aws/en/oltp/projects/lakebase-text)

`lakebase_text` 通过 `lakebase_bm25` 对 PostgreSQL 标准全文检索向量提供 BM25 排序。托管 Lakebase 目录为 PostgreSQL 16–18 列出的 SQL 版本为 0.1.3。

### 启用与搜索

在项目设置中启用 Lakebase Search；该操作不可撤销，会重启项目的全部计算节点并断开客户端连接。使用具备数据库权限的角色安装，写入文本后再建立索引：

```sql
CREATE EXTENSION IF NOT EXISTS lakebase_text;
CREATE TABLE documents (
  id bigint, passage text,
  terms tsvector GENERATED ALWAYS AS (to_tsvector('english', passage)) STORED
);
INSERT INTO documents (id, passage) VALUES (1, 'Postgres search');
CREATE INDEX documents_bm25 ON documents USING lakebase_bm25 (terms);
SELECT id FROM documents
ORDER BY terms <@> to_bm25query(to_tsvector('english', 'search'), 'documents_bm25')
LIMIT 5;
```

分数为负，按升序排列时最相关的结果在前。索引扫描可能省略零分行。文档与查询应使用一致的全文检索配置。

### 维护

在已有数据的表上建立索引，保证自动清理有效运行；批量写入后可用 `VACUUM` 刷新统计。`ALTER EXTENSION lakebase_text UPDATE` 更新 SQL 对象；索引格式升级另需执行 `REINDEX INDEX CONCURRENTLY`。更新前检查供应商提供的版本。这里没有公开源码、许可证、实现语言和预加载设置的说明，本文仅描述供应商托管功能。
