## 用法

来源：

- [官方上游 README](https://github.com/zvdy/pgao/blob/45f09972a1d6e551125d95279c0fd863a5533aa3/extension/README.md)
- [官方扩展控制文件 (pgao.control)](https://github.com/zvdy/pgao/blob/45f09972a1d6e551125d95279c0fd863a5533aa3/extension/pgao.control)
- [官方扩展 SQL (pgao--0.1.0.sql)](https://github.com/zvdy/pgao/blob/45f09972a1d6e551125d95279c0fd863a5533aa3/extension/pgao--0.1.0.sql)

`pgao` — 所有函数均为 SQL 仅用，STABLE/IMMUTABLE，并不需要超级用户权限。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgao;
SELECT * FROM pgao.health();
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgao.health()` 是一个扩展函数，返回 `TABLE`。
- `pgao.replication_lag_ms()` 是一个扩展函数，返回 `bigint`。
- `pgao.table_bloat()` 是一个扩展函数，返回 `TABLE`。
- `pgao.version()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
