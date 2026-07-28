## 用法

来源：

- [官方上游 README](https://github.com/magnusp/pg_hello/blob/83ab1b18a51ff0d0d42d3639caf79b5885bb7040/README.md)
- [官方扩展控制文件 (pg_hello.control)](https://github.com/magnusp/pg_hello/blob/83ab1b18a51ff0d0d42d3639caf79b5885bb7040/pg_hello.control)
- [官方扩展 SQL (pg_hello--1.0.sql)](https://github.com/magnusp/pg_hello/blob/83ab1b18a51ff0d0d42d3639caf79b5885bb7040/pg_hello--1.0.sql)

`pg_hello` — pg_hello，一个非常基础的 Postgres 扩展 =========================================。当应用程序需要这种特定的数据库功能时使用它。经过审核的上游材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION pg_hello;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `hello(TEXT)` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 经过审核的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 上游材料包含显式的弃用边界。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
