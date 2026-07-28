## 用法

来源：

- [官方扩展控制文件（lambda.control）](https://api.pgxn.org/src/pg_lambda/pg_lambda-1.0.5/lambda.control)
- [官方扩展 SQL（lambda.sql）](https://api.pgxn.org/src/pg_lambda/pg_lambda-1.0.5/sql/lambda.sql)

`lambda` — Lambda 函数用于 Postgres。当数据库代码必须在该过程语言中运行或与其进行交互时，请使用此扩展。在目标 PostgreSQL 构建上使用链接的上游修订版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION lambda;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `lambda(code text , VARIADIC inputs anyarray)` 是一个扩展函数并返回 `anyelement`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0.5`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
