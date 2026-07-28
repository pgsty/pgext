## 用法

来源：

- [pgContext 0.2.0 pgvector 共存指南](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_coexist.md)
- [pgContext 0.2.0 pgvector 迁移指南](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/pgvector_migration.md)
- [pgcontext_pgvector 控制文件](https://github.com/evokoa/pgcontext/blob/v0.2.0/pgcontext_pgvector.control)
- [pgcontext_pgvector 扩展 SQL](https://github.com/evokoa/pgcontext/blob/v0.2.0/sql/pgcontext_pgvector--0.2.0.sql)
- [pgContext 0.2.0 发行说明](https://github.com/evokoa/pgcontext/blob/v0.2.0/docs/user_guide/release_notes.md)

`pgcontext_pgvector` 是可选的 pgContext companion 桥，用于在 pgvector 扩展拥有的列上提供 pgContext HNSW 索引。它不会合并两套类型系统，也不会复制应用数据；它增加经过认证的类型转换、支持函数和操作符类，而精确距离语义仍绑定到 pgvector 操作符。

### 认证组合与安装

0.2.0 只有在数据库使用 PostgreSQL 17、pgContext 0.2.0，并把 pgvector 0.8.x 安装在 `public` 时才会通过检查。应显式安装前置扩展与桥：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgcontext;
CREATE EXTENSION pgcontext_pgvector;
```

两个前置扩展也可以按相反顺序安装，但 `pgcontext_pgvector` 必须在二者之后安装。安装需要超级用户权限。

### 为已有 pgvector 列创建索引

```sql
CREATE INDEX items_embedding_pgc
    ON items USING pgcontext_hnsw
       (embedding pgcontext.vector_hnsw_pgvector_cosine_ops);

SELECT id
FROM items
ORDER BY embedding <=> $1::public.vector
LIMIT 10;
```

已有 pgvector 写法的 SQL 可以使用 pgContext 访问方法。ANN 候选会重新解析到可见 heap 行，并通过 pgvector 操作符精确重排，从而保留其 `double precision` 距离结果语义。

### 重要对象

- `pgcontext.vector_hnsw_pgvector_l2_ops`、`pgcontext.vector_hnsw_pgvector_ip_ops`、`pgcontext.vector_hnsw_pgvector_cosine_ops` 与 `pgcontext.vector_hnsw_pgvector_l1_ops` 用于已有 `public.vector` 列。
- `pgcontext.sparsevec_hnsw_pgvector_cosine_ops` 用于认证范围内的 `public.sparsevec` 列，但受文档规定的 16,000 维和页面包络限制。
- `pgcontext.migration_report()` 即使没有安装桥，也能盘点 pgvector 列、依赖、HNSW 与 IVFFlat。
- 所有权转换函数提供经过审查的 fast 或 restricted-online 工作流；IVFFlat 会重建为 HNSW，而不是就地转换。

### 依赖与移除边界

主 `pgcontext` 扩展仍独立于 pgvector。桥索引依赖 `pgcontext_pgvector`，而桥依赖两个父扩展，因此在这些索引被删除或转换之前，PostgreSQL 会阻止以 `RESTRICT` 移除它们。

不要把 `DROP EXTENSION vector CASCADE` 当作迁移方式。应先盘点数组、视图、函数、预备会话、表达式索引及其他应用依赖。该桥不会提供 pgvector 的全部辅助函数、IVFFlat、迭代扫描 GUC、并行构建、子向量或进度报告行为。

无需预加载或重启。该桥是带权限的兼容面，并不承诺未来任意 pgContext、pgvector、PostgreSQL 大版本或磁盘索引组合都兼容；任一组件变化后都要重新执行认证预检与重建验证。
