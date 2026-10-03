## 用法

来源：

- [README 0.1.0](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/README.md)
- [Control](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/snout_graphql.control)
- [SQL 0.1.0](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/sql/snout_graphql--0.1.0.sql)
- [Cargo](https://github.com/snoutdata/snout-graphql/blob/4f549d1738072780d9cb0df7fe9d848e6666bbab/Cargo.toml)

`snout_graphql` 0.1.0 根据可见的表、视图和函数，在 PostgreSQL 内执行 GraphQL 请求。这个 Rust 扩展面向 PostgreSQL 17。安装需要超级用户，会在固定的 `graphql` schema 中创建对象，并注册两个事件触发器。

### 基本用法

```sql
CREATE EXTENSION snout_graphql;
COMMENT ON SCHEMA public IS '@graphql({"inflect_names": true})';
SELECT graphql.resolve($$ { __typename } $$);
```

`graphql.resolve(query, variables, "operationName", extensions)` 返回 JSONB 响应，变量值与请求文档分别传入。请求按调用者的表权限和行级安全策略执行；调用者还需要访问扩展 schema 与对象的权限。发生错误时，响应包含错误且数据为空，并回滚该请求已执行的写入。若要提供 HTTP 端点，需要另设服务，以请求对应的数据库角色调用该函数。

### 对象与配置

- `graphql._internal_resolve`：直接抛出错误的内部入口。
- `graphql.comment_directive`：读取对象注释中的 JSON 指令。
- `graphql.get_schema_version`、`graphql.increment_schema_version`、`graphql.seq_schema_version`：维护反射 schema 缓存的失效状态。
- `graphql_watch_ddl`、`graphql_watch_drop`：跟踪相关 DDL 和对象删除的事件触发器。
- `graphql.exception`：报告修改行数超过允许范围的错误。
- 注释指令包括 `inflect_names`、`max_rows`、`introspection`、`totalCount`、`aggregate`、`primary_key_columns` 和 `foreign_keys`。内省默认关闭，集合分页默认每页 30 行。扩展没有 GUC 参数或后台工作进程。

### 使用边界

入口函数和事件触发器与 `pg_graphql` 使用的对象重叠；不要将两者安装到同一个 schema，也不要假定存在自动迁移路径。当前修订中缺少上游 README 链接的差异说明文档，因此这里不保证与其他实现等价。

每个连接会缓存反射 schema、执行计划和内省结果。对象注释会影响对外 API，表授权和行级安全仍由应用负责。GraphQL 文本在 PostgreSQL 后端进程内解析；收录这个源码预览版本不代表通过生产环境认证。
