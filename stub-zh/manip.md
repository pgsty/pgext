## 用法

来源：

- [官方上游 README](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/README.md)
- [官方扩展控制文件 (manip.control)](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/manip/manip.control)
- [官方扩展 SQL (manip--0.1.sql)](https://github.com/mkindahl/pg_examples/blob/277a29ac1b31478911c75ddd886d4dea02730aa8/manip/manip--0.1.sql)

`manip` — 各种操作函数。当需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION manip;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `get_prepared_plan(stmt_name text)` 是一个扩展函数，返回 `text`。
- `scan_table` 是一个扩展过程。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
