## 用法

来源：

- [官方上游 README](https://github.com/sweatybridge/pg_ssh/blob/94ef25fb4a4666fa1447a489ca37382d2d84106a/README.md)
- [官方扩展控制文件 (pg_ssh.control)](https://github.com/sweatybridge/pg_ssh/blob/94ef25fb4a4666fa1447a489ca37382d2d84106a/pg_ssh.control)
- [官方实现源代码](https://github.com/sweatybridge/pg_ssh/blob/94ef25fb4a4666fa1447a489ca37382d2d84106a/src/lib.rs)

`pg_ssh` — PostgreSQL **18** (13–17 也通过 pgXX 特性构建) Rust 工具链 (稳定版) [cargo-pgrx][cpgrx] **0.19.1** — 必须与 pgrx crate 的版本完全匹配 系统包 (Debian/Ubuntu):. 使用它来进行相应的 SQL 或数据库实用工具工作流。使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_ssh;

SELECT convert_from(stdout, 'UTF8') AS stdout, exit_code
  FROM ssh.exec('web-1', 'uname -a; uptime');
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `exec` 是一个扩展函数。
- `keygen` 是一个扩展函数。
- `session_close` 是一个扩展函数。
- `session_exec` 是一个扩展函数。
- `session_open` 是一个扩展函数。
- `sessions()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.3.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
