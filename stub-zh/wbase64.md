## 用法

来源：

- [官方上游 README](https://github.com/wuriyanto48/pgwbase64/blob/963c280ca190b23f555e199aacd8feb5a222dde8/README.md)
- [官方扩展控制文件 (wbase64.control)](https://github.com/wuriyanto48/pgwbase64/blob/963c280ca190b23f555e199aacd8feb5a222dde8/wbase64.control)
- [官方扩展 SQL (wbase64--0.0.1.sql)](https://github.com/wuriyanto48/pgwbase64/blob/963c280ca190b23f555e199aacd8feb5a222dde8/wbase64--0.0.1.sql)

`wbase64` — Base64 PostgreSQL 扩展。用于相应的 SQL 或数据库实用程序工作流。使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION wbase64;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgw_b64_decode(text)` 是一个扩展函数，返回 `text`。
- `pgw_b64_encode(text)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
