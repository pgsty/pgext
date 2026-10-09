## 用法

来源：

- [packages/function-resolution/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/README.md)
- [packages/function-resolution/sql/pgpm-function-resolution--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/sql/pgpm-function-resolution--0.47.0.sql)
- [packages/function-resolution/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/Makefile)
- [packages/function-resolution/pgpm-function-resolution.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/function-resolution/pgpm-function-resolution.control)

`pgpm-function-resolution` 沿 PGPM 作用域链选择函数定义，并可将任务加入队列。`function_resolution.resolve` 返回匹配的定义；`function_resolution.enqueue` 解析路由后调用应用任务队列。

### 核心用法

```sql
CREATE EXTENSION "pgpm-function-resolution";
SELECT function_definition_id, resolved_scope
FROM function_resolution.resolve(
  app_scope.platform_database_id(), 'platform', NULL, 'my:task', false
);
```

### 运行边界

需要已填充的 metaschema 与函数目录、JWT 执行上下文以及非 NULL 的作用域。目录中没有匹配定义，与目录本身无法完成查找，是不同情况。此版本还提供 `function_resolution.resolve_invocation`、`function_resolution.resolve_bucket`、`function_resolution.resolve_api` 及能力解析。入队只写入队列表，不会运行工作进程。多个辅助函数采用 `SECURITY DEFINER`，部署时必须同时控制应用角色与执行权限。 依赖为 `plpgsql`, `errors`, `pgpm-verify`, `metaschema-schema`, `metaschema-modules`, `pgpm-app-scope`, `pgpm-database-jobs`, `pgpm-jwt-claims`。 安装脚本会向预先存在的角色 `administrator`, `authenticated`, `anonymous` 授权。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
