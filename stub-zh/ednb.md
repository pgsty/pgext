## 用法

来源：

- [官方上游 README](https://github.com/scpatt/pg_edn/blob/92613c2719e660be1d21b867d2c0dd9d089d0e4f/README.md)
- [官方扩展控制文件 (ednb.control)](https://github.com/scpatt/pg_edn/blob/92613c2719e660be1d21b867d2c0dd9d089d0e4f/ednb/ednb.control)
- [官方扩展 SQL (ednb--0.0.1.sql)](https://github.com/scpatt/pg_edn/blob/92613c2719e660be1d21b867d2c0dd9d089d0e4f/ednb/ednb--0.0.1.sql)

`ednb` — Postgres 扩展，用于支持 EDN 作为自定义数据类型。当应用程序需要此类型、域或其操作符时，请使用此扩展。上游将其描述为一个正在进行中的项目。

### 核心工作流

```sql
CREATE EXTENSION ednb;
```

在目标数据库中安装扩展，运行可用的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `ednb_in(cstring)` 是一个扩展函数，返回 `ednb`。
- `ednb_out(ednb)` 是一个扩展函数，返回 `cstring`。
- `ednb` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为可重定位。
- 上游将该项目描述为一个正在进行中的项目。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
