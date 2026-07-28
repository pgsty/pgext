## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/pgmpc/pgmpc-0.1.0/README)
- [官方扩展控制文件 (pgmpc.control)](https://api.pgxn.org/src/pgmpc/pgmpc-0.1.0/pgmpc.control)
- [官方扩展 SQL (pgmpc--0.1.0.sql)](https://api.pgxn.org/src/pgmpc/pgmpc-0.1.0/pgmpc--0.1.0.sql)

`pgmpc` — pgmpc, mpd 客户端 for PostgreSQL ================================. 请使用它来执行相应的 SQL 或数据库实用程序工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgmpc;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mpd_add(path text)` 是一个扩展函数并返回 `void`。
- `mpd_clear()` 是一个扩展函数并返回 `void`。
- `mpd_consume()` 是一个扩展函数。
- `mpd_load(path text)` 是一个扩展函数并返回 `void`。
- `mpd_ls(path text, OUT song text)` 是一个扩展函数并返回 `SETOF`。
- `mpd_lsplaylists(OUT playlist text)` 是一个扩展函数并返回 `SETOF`。
- `mpd_next()` 是一个扩展函数并返回 `void`。
- `mpd_pause()` 是一个扩展函数。
- `mpd_play()` 是一个扩展函数并返回 `void`。
- `mpd_playlist(OUT song text)` 是一个扩展函数并返回 `SETOF`。
- `mpd_playlist(path text, OUT song text)` 是一个扩展函数并返回 `SETOF`。
- `mpd_prev()` 是一个扩展函数并返回 `void`。
- `mpd_random()` 是一个扩展函数。
- `mpd_repeat()` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
