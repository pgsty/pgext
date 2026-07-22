## 用法

来源：

- [官方 README](https://github.com/keithf4/pg_doc_store/blob/cd2c8c4bc7235bc6443ee0aa9dbec19f1f945333/README.md)
- [扩展控制文件](https://github.com/keithf4/pg_doc_store/blob/cd2c8c4bc7235bc6443ee0aa9dbec19f1f945333/pg_doc_store.control)
- [官方 PGXN 发行页](https://pgxn.org/dist/pg_doc_store/)

`pg_doc_store` 在具备 ACID 能力的 PostgreSQL 表之上，用 PL/pgSQL 实现了精简的文档存储 API。它为 `jsonb` 文档创建集合表、分配 UUID、维护全文检索向量，并提供保存、包含查询与文本搜索函数。

### 核心流程

该扩展依赖 `pgcrypto` 生成 UUID。集合名必须包含模式限定。

```sql
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_doc_store;

SELECT * FROM create_document('public.customers');

SELECT save_document(
    'public.customers',
    '{"name":"Ada","city":"London"}'::jsonb
);
```

必要时，`save_document` 会创建集合。文档含有 `id` 时会按该 UUID 更新或插入；没有该键时，函数会生成 ID，并把它加入返回的文档。

```sql
SELECT *
FROM find_document(
    'public.customers',
    '{"city":"London"}'::jsonb,
    'name',
    'ASC'
);

SELECT *
FROM search_document('public.customers', 'Ada');
```

### 用户接口

- `create_document(p_tablename text)`：创建集合并返回其模式名和表名。
- `save_document(p_tablename text, p_doc_string jsonb)`：插入或更新文档，并返回保存后的 JSON。
- `find_document(p_tablename text, p_criteria jsonb, p_orderbykey text DEFAULT 'id', p_orderby text DEFAULT 'ASC')`：返回包含指定条件的文档，并按文档键排序。
- `search_document(p_tablename text, p_query text)`：按全文相关度返回文档。

每个集合包含 `id uuid`、`body jsonb`、`search tsvector`、`created_at` 和 `updated_at`，并带有主键、用于 JSON 包含与全文检索的 GIN 索引，以及刷新搜索向量的触发器。

### 运维说明

上游要求 PostgreSQL 9.5 或更高版本以及 `pgcrypto`；控制文件声明了该依赖，并固定了扩展的模式位置。集合表是由 API 动态创建的普通数据库对象，因此备份、权限、模式迁移和保留策略都必须明确包含它们。调用集合的函数应只使用可信且包含模式限定的标识符。无需共享预加载。上游版本 `0.2.1` 是较早且精简的 PL/pgSQL 实现；应在目标 PostgreSQL 版本上测试 JSON 键排序、全文配置、并发更新插入与升级行为。
