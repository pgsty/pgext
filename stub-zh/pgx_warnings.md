## 用法

来源：

- [官方上游 README](https://github.com/valehdba/pgx_warnings/blob/723431d3557d7b341ec663cf3aa16f23dc2f5973/README.md)
- [官方扩展控制文件 (pgx_warnings.control)](https://github.com/valehdba/pgx_warnings/blob/723431d3557d7b341ec663cf3aa16f23dc2f5973/pgx_warnings.control)
- [官方扩展 SQL (pgx_warnings--1.0.sql)](https://github.com/valehdba/pgx_warnings/blob/723431d3557d7b341ec663cf3aa16f23dc2f5973/pgx_warnings--1.0.sql)

`pgx_warnings` — 一个 PostgreSQL 扩展，能够实时捕获服务器日志中的所有 WARNING、ERROR、FATAL 和 PANIC 消息，并即时通知到一个 Telegram 频道。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上方链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgx_warnings;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgx_warnings_clear()` 是一个扩展函数，返回 `void`。
- `pgx_warnings_list(IN max_entries integer DEFAULT 100, OUT "timestamp" timestamptz, OUT level text, OUT database text, OUT message text, OUT pid integer, OUT sent boolean)` 是一个扩展函数，返回 `SETOF`。
- `pgx_warnings_stats(OUT current_entries integer, OUT buffer_size integer, OUT total_captured bigint, OUT total_sent bigint, OUT total_failed bigint, OUT enabled boolean)` 是一个扩展函数，返回 `record`。
- `pgx_warnings_test()` 是一个扩展函数，返回 `text`。
- `pgx_warnings` 是一个扩展定义视图。
- `pgx_warnings_info` 是一个扩展定义视图。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
