## 用法

来源：

- [官方安装指南](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/docs/src/getting-started/installation.md)
- [官方五分钟教程](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/docs/src/getting-started/hello-world.md)
- [官方扩展控制文件](https://github.com/trickle-labs/pg-ripple/blob/31d4a2c97484988958ba986e5c5302fa24f0abbe/pg_ripple.control)

`pg_ripple` 是面向 PostgreSQL 18 的 RDF 知识图谱引擎，支持 SPARQL 查询、SHACL 校验、Datalog 推理、联邦查询与图分析。图数据存储在 PostgreSQL 内；本次复核的版本为 `0.128.0`，安装需要超级用户。

### 核心流程

HTAP 合并工作进程和共享内存字典缓存要求预加载。添加库、重启 PostgreSQL、创建扩展，然后加载 RDF 并查询。

```sql
ALTER SYSTEM SET shared_preload_libraries = 'pg_ripple';
-- Restart PostgreSQL before continuing.

CREATE EXTENSION pg_ripple;

SELECT pg_ripple.register_prefix('ex', 'http://example.org/');
SELECT pg_ripple.register_prefix('foaf', 'http://xmlns.com/foaf/0.1/');

SELECT pg_ripple.load_turtle('
  @prefix ex: <http://example.org/> .
  @prefix foaf: <http://xmlns.com/foaf/0.1/> .
  ex:alice foaf:name "Alice" .
  ex:bob foaf:name "Bob" .
');

SELECT *
FROM pg_ripple.sparql('
  PREFIX foaf: <http://xmlns.com/foaf/0.1/>
  SELECT ?person ?name WHERE {
    ?person foaf:name ?name .
  }
');
```

`load_turtle` 返回加载的三元组数量。`sparql` 每行返回一个 JSONB 对象，以查询变量为键。

### 主要 SQL 接口

- `register_prefix` 管理短命名空间前缀。
- `insert_triple`、`delete_triple` 与 `triple_count` 提供直接的三元组操作和基本检查。
- `load_turtle` 加载 Turtle 文档；命名图 API 可分隔租户或数据来源。
- `sparql` 及其解释和游标变体从 SQL 执行图查询。
- SHACL、Datalog、推理、向量/混合搜索、订阅、导出和管理 API 分别作为独立功能组提供文档。

### 运维边界

这一源码版本面向 PostgreSQL `18`，不能假设兼容更早的大版本。扩展会创建内部模式和按谓词划分的存储，包括 delta、main 与 tombstone 关系。应使用公开函数，不要直接修改内部对象。基本调用可以按需加载动态库，但生产级 HTAP 与共享内存功能需要 `shared_preload_libraries` 和重启。

这个宽广且快速演进的 API 必须按版本使用。应固定扩展及所有 HTTP 或 relay 配套组件的版本，验证备份恢复与升级，并在公开端点或远端数据源前检查安全和对外联邦设置。
