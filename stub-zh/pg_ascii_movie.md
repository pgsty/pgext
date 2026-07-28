## 用法

来源：

- [官方上游 README](https://github.com/yamatattsu/pg_ascii_movie/blob/0bcc5f8836afcbde84139dcbbc0054a5b24acbd6/README.md)
- [官方扩展控制文件 (pg_ascii_movie.control)](https://github.com/yamatattsu/pg_ascii_movie/blob/0bcc5f8836afcbde84139dcbbc0054a5b24acbd6/pg_ascii_movie.control)
- [官方扩展 SQL (pg_ascii_movie--1.0.sql)](https://github.com/yamatattsu/pg_ascii_movie/blob/0bcc5f8836afcbde84139dcbbc0054a5b24acbd6/pg_ascii_movie--1.0.sql)

`pg_ascii_movie` — pg_ascii_movie 是一个 PostgreSQL 扩展，用于观看 ASCII 动画。当应用程序需要此特定数据库功能时，请使用此扩展。在安装此扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_ascii_movie;
```

在目标数据库中安装扩展，当有可用示例时，运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `get_frame(in pos integer)` 是一个扩展函数，返回 `text`。
- `get_wait_time(in pos integer)` 是一个扩展函数，返回 `double`。
- `play_sw1` 是一个扩展过程。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 请先安装并验证确认的依赖项：`file_fdw`。
- 控制文件将此扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
