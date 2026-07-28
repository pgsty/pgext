## 用法

来源：

- [官方上游 README](https://github.com/andrew/omni_git/blob/9bc1d4e34f11d49a10071e4c7bafe15ef5ac9fcc/README.md)
- [官方扩展控制文件 (omni_git.control)](https://github.com/andrew/omni_git/blob/9bc1d4e34f11d49a10071e4c7bafe15ef5ac9fcc/omni_git.control)
- [官方实现源代码](https://github.com/andrew/omni_git/blob/9bc1d4e34f11d49a10071e4c7bafe15ef5ac9fcc/omni_git.c)

`omni_git` — 一个 PostgreSQL 扩展，用于将 git 仓库存储在数据库表中，并提供 git 智能 HTTP 协议服务，从而将 Postgres 转变为一个 git 远程。当需要移植或模拟相应的数据库 API 时，请使用此扩展。在安装之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION omni_git;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 请先安装确认的扩展依赖项：`gitgres`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
