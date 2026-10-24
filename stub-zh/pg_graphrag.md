## 用法

来源：

- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/README.md)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/pg_graphrag.control)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/Cargo.toml)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/src/sql/schema.sql)
- [Official source](https://github.com/Riku-KANO/pg-graph/blob/16986a73d3b45f02ff1744efc0efcc688517da9f/examples/quickstart.sql)

`pg_graphrag` 0.1.0 存储带来源证据、时态关系和权限感知检索的知识图谱。SQL 模式使用 `uuidv7` 与 `WITHOUT OVERLAPS` 等 PostgreSQL 18 功能；即使 Cargo 清单包含较早版本的构建开关，也需要 PostgreSQL 18。该预览版仓库尚未提供可确认的许可证。

### 基本流程

由超级用户安装，依赖 `btree_gist`，对象位于 `graphrag` 模式。数据库不调用 LLM，抽取结果和社区摘要由应用提供。下面的轻量流程不需要 LLM。

```sql
CREATE EXTENSION pg_graphrag CASCADE;
SELECT graphrag.create_graph('demo');
SELECT graphrag.add_document('demo', content => 'Alice works at Example Labs.',
                             external_id => 'doc1', title => 'Team');
SELECT graphrag.lazy_index('demo', id)
FROM graphrag.document WHERE external_id = 'doc1';
SELECT graphrag.build_context('demo', query => 'Alice', max_tokens => 600);
```

`ingest_extraction` 为来源片段存储结构化抽取结果。`local_search` 结合图、稠密向量和词法检索，`relational_paths` 返回带证据的解释路径。`build_communities`、`pending_report` 和 `set_report` 支持应用生成摘要。时态检索可指定 `as_of`；通过受支持的 API 删除文档时，会清理其派生证据。

### 权限

文档和片段的 ACL 使用 PostgreSQL 角色名，空 ACL 表示公开。调用 `graphrag.enable_row_security()` 保护直接表访问，再按需使用 `graphrag.grant_reader` 或 `graphrag.grant_writer` 授权。检索应以实际调用者角色执行，例如使用事务内的角色切换，不应以绕过行级安全的所有者或超级用户执行。扩展根据可见证据过滤派生实体和报告；不依附证据的人工维护属性可能公开可见。

### 运行边界

不需要预加载或重启。稠密向量检索会扫描候选，不依赖 `pgvector`，不能据此宣称具备向量索引检索。LLM 输出仍是应用提供的不可信数据，摘要也可能过期。部署前应使用有代表性的语料验证访问规则、更新与删除流程。
