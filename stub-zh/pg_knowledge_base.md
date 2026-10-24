## 用法

来源：

- [Official source](https://github.com/igor-suhorukov/pg_knowledge_base/blob/a25f8dbec782699981e57b864afe33ab074d56a5/README.md)
- [Official source](https://github.com/igor-suhorukov/pg_knowledge_base/blob/a25f8dbec782699981e57b864afe33ab074d56a5/pg_knowledge_base.control)
- [Official source](https://github.com/igor-suhorukov/pg_knowledge_base/blob/a25f8dbec782699981e57b864afe33ab074d56a5/sql/pg_knowledge_base--0.1.sql)

`pg_knowledge_base` 0.1 将事务性 RDF 表、Oxigraph SPARQL 引擎和关系表到 RDF 的映射结合起来。该预览版本面向原生 PostgreSQL 19，上游在 19rc1 上测试。SQL 函数不需要预加载；可选网络端点有额外要求。

### 基本流程

以超级用户安装。对象位于固定的 `kb` 模式。

```sql
CREATE EXTENSION pg_knowledge_base;
SELECT kb.load('<urn:alice> <urn:knows> <urn:bob> .');
SELECT * FROM kb.sparql('SELECT ?s ?o WHERE { ?s <urn:knows> ?o }')
  AS t(s text, o text);
SELECT kb.sparql_json('SELECT ?s WHERE { ?s <urn:knows> <urn:bob> }');
```

`kb.load` 导入 RDF 文本；`kb.load_file` 在通过权限检查后读取服务端文件。`kb.sparql` 返回记录，调用时须显式指定 SQL 列定义；`kb.sparql_json` 和 `kb.sparql_text` 返回序列化结果。`kb.sparql_update` 在调用者事务中修改 RDF 存储。回滚、WAL、备份和复制使用 PostgreSQL 的表机制。

### 映射与网络访问

`kb.add_mapping` 注册 R2RML 映射，`kb.direct_mapping` 创建直接映射。查询以调用角色的权限与行级安全规则读取映射表。SPARQL 更新修改 RDF 存储，不会写入被映射的关系表。不支持远程 `SERVICE` 求值。

HTTP 端点默认关闭。启用时需要将 `pg_knowledge_base` 加入 `shared_preload_libraries`，设置 `kb.endpoint_listen_addresses`，然后重启 PostgreSQL，同时保留已有预加载项。默认端口为 7878；各数据库的查询、更新与图存储路径通过 PostgreSQL 角色和 `pg_hba.conf` 认证。传输凭据时使用 TLS，并核对连接与工作进程限制。可选追踪功能也需要预加载及文档指定的追踪扩展。

### 运行边界

应明确授予 `kb` 模式与底层对象的访问权限。批量加载会锁定四元组表并重建索引，应考虑并发操作受阻的影响。Rust 引擎有独立的内存上限，不受 PostgreSQL `work_mem` 代管。在开放任意查询前，应核对支持的 RDF/结果格式和资源设置。
