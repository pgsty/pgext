## 用法

来源：

- [官方上游 README](https://github.com/solidcoredata/pgwire4/blob/13f90441a022cce962f6e6bf1b710703c21af19a/README.md)
- [官方扩展控制文件 (pgwire4.control)](https://github.com/solidcoredata/pgwire4/blob/13f90441a022cce962f6e6bf1b710703c21af19a/ext/pgwire4.control)
- [官方扩展 SQL (pgwire4--1.0.sql)](https://github.com/solidcoredata/pgwire4/blob/13f90441a022cce962f6e6bf1b710703c21af19a/ext/sql/pgwire4--1.0.sql)

`pgwire4` — PostgreSQL 扩展，提供新的网络协议。当应用程序需要此特定数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgwire4;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgwire4.bulk_int8(stream_name text)` 是一个扩展函数，返回 `SETOF`。
- `pgwire4.bulk_scan(stream_name text)` 是一个扩展函数，返回 `SETOF`。
- `pgwire4.stats(OUT accepted bigint, OUT dispatch_failures bigint, OUT workers integer, OUT databases integer, OUT listener_pid integer)` 是一个扩展函数，返回 `record`。
- `pgwire4.status(OUT slot integer, OUT pid integer, OUT database text, OUT state text, OUT sessions bigint, OUT queries bigint, OUT rows_streamed bigint, OUT cancels bigint, OUT errors bigint, OUT cache_hits bigint, OUT cache_misses bigint, OUT session_user_name text, OUT sess…)` 是一个扩展函数，返回 `SETOF`。
- `pgwire4.version()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码中的信息一致。
