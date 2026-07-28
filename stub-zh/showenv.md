## 用法

来源：

- [官方上游 README](https://github.com/mkindahl/pg_showenv/blob/bad19503239223fad3ce77dc8133178bff91a663/README.md)
- [官方扩展控制文件 (showenv.control)](https://github.com/mkindahl/pg_showenv/blob/bad19503239223fad3ce77dc8133178bff91a663/showenv.control)
- [官方扩展 SQL (showenv--1.0.sql)](https://github.com/mkindahl/pg_showenv/blob/bad19503239223fad3ce77dc8133178bff91a663/showenv--1.0.sql)

`showenv` — 一个 PostgreSQL 扩展，用于以 SQL 结果集的形式显示服务器进程的环境变量。在进行数据库管理或自动化上述描述的行为时，请使用此扩展。审核过的上游项目已归档或不再维护。

### 核心工作流

```sql
CREATE EXTENSION showenv;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `environment_variables()` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 扩展控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为可信。
- 上游材料表明该项目已被放弃或不再维护。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
