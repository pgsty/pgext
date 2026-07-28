## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/session_hash_tools/session_hash_tools-1.0.0/README)

`session_hash_tools` — 开启 $dbname createlang plperl $dbname psql -c "CREATE SCHEMA tools" $dbname psql -f $PGDATA/share/contrib/tools.sql $dbname。当 SQL 需要这些特殊函数或聚合时使用它。在安装并验证其扩展依赖项之前，请勿使用。

### 核心工作流

此组件在审查的源代码中没有确认的独立 `CREATE EXTENSION` 工作流。仅通过上游机制进行构建、加载或启用，并在隔离数据库中验证最终服务器行为。

### 要求与注意事项

- 该目录记录版本 `1.0.0`。
- 首先安装确认的扩展依赖项：`plperl`, `plperlu`。
- 在生产使用前，需确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码的一致性。
