## 用法

来源：

- [官方上游 README](https://github.com/chadcatlett/pg-jq/blob/99b40b8d9d4924bab3465f90c655457cb2d9c50c/README.md)
- [官方扩展控制文件 (pg_jq.control)](https://github.com/chadcatlett/pg-jq/blob/99b40b8d9d4924bab3465f90c655457cb2d9c50c/pg_jq.control)
- [官方实现源码](https://github.com/chadcatlett/pg-jq/blob/99b40b8d9d4924bab3465f90c655457cb2d9c50c/src/lib.rs)

`pg_jq` — 这是一个玩具 PostgreSQL 扩展，它将基本的 libjq 功能暴露给 PostgreSQL。使用它来完成相应的 SQL 或数据库实用工具工作流。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_jq;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `json_jq` 是一个扩展函数。
- `jsonb_jq` 是一个扩展函数。
- `what_is_something_carlson_likes()` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
