## 用法

来源：

- [Extension versions](https://docs.databricks.com/aws/en/oltp/projects/extensions)
- [Enable Lakebase Search](https://docs.databricks.com/aws/en/oltp/projects/lakebase-search)
- [Extension guide](https://docs.databricks.com/aws/en/oltp/projects/lakebase-vector)

`lakebase_vector` 在托管 Lakebase Postgres 中提供兼容 pgvector 的近似向量搜索索引方法 `lakebase_ann`。供应商为 PostgreSQL 16–18 列出的版本为 1.1.1；它不是可下载的软件包。

### 启用与查询

先在项目设置中启用 Lakebase Search。该操作不可撤销，会重启项目中的全部计算节点并断开客户端连接。随后由具备权限的角色在每个数据库中安装：

```sql
CREATE EXTENSION IF NOT EXISTS lakebase_vector CASCADE;
CREATE TABLE items (id bigint, embedding vector(3));
INSERT INTO items VALUES (1, '[0.1,0.2,0.3]');
CREATE INDEX items_ann ON items USING lakebase_ann (embedding vector_l2_ops);
SELECT * FROM items ORDER BY embedding <-> '[0.2,0.1,0.3]' LIMIT 5;
```

`CASCADE` 会安装依赖 `vector`。在数据库外生成嵌入，并保持向量维度一致。ANN 返回近似最近邻结果。

### 维护

`ALTER EXTENSION lakebase_vector UPDATE` 更新 SQL 对象。索引存储具有独立版本，现有索引需通过 `REINDEX INDEX CONCURRENTLY` 重建才能采用新格式。升级前检查实际计算节点的 `pg_available_extensions`；供应商可能要求重启计算节点才能提供新版本。上述文档没有确认实现语言、库加载或预加载设置以及公开许可证。
