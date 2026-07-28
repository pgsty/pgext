## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/notify_now/notify_now-1.0.0/README.md)
- [官方扩展控制文件 (notify_now.control)](https://api.pgxn.org/src/notify_now/notify_now-1.0.0/notify_now.control)
- [官方扩展 SQL (notify_now--1.0.sql)](https://api.pgxn.org/src/notify_now/notify_now-1.0.0/notify_now--1.0.sql)

`notify_now` — 这个简单的扩展允许你通过内置的 PostgreSQL NOTIFY API 从单个查询返回多个响应。没有额外的依赖项。当应用程序需要这种特定的数据库功能时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION notify_now;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `notify_now(text, text)` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
