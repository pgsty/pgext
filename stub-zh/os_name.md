## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/os_name/os_name-0.0.3/README.md)
- [官方扩展控制文件 (os_name.control)](https://api.pgxn.org/src/os_name/os_name-0.0.3/os_name.control)
- [官方扩展 SQL (os_name--0.0.2.sql)](https://api.pgxn.org/src/os_name/os_name-0.0.3/os_name--0.0.2.sql)

`os_name` — 可枚举的移动操作系统类型，存储在一个字节的固定长度类型中。当应用程序数据需要此类型、域或其操作符时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION os_name;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，验证安装的版本和返回值，然后再将其集成到应用程序 SQL 中。

### 重要对象

- `hash_os_name(os_name)` 是一个扩展函数，返回 `integer`。
- `os_name_cmp(os_name, os_name)` 是一个扩展函数，返回 `integer`。
- `os_name_eq(os_name, os_name)` 是一个扩展函数，返回 `boolean`。
- `os_name_ge(os_name, os_name)` 是一个扩展函数，返回 `boolean`。
- `os_name_gt(os_name, os_name)` 是一个扩展函数，返回 `boolean`。
- `os_name_in(cstring)` 是一个扩展函数，返回 `os_name`。
- `os_name_le(os_name, os_name)` 是一个扩展函数，返回 `boolean`。
- `os_name_lt(os_name, os_name)` 是一个扩展函数，返回 `boolean`。
- `os_name_ne(os_name, os_name)` 是一个扩展函数，返回 `boolean`。
- `os_name_out(os_name)` 是一个扩展函数，返回 `cstring`。
- `os_name_recv(internal)` 是一个扩展函数，返回 `os_name`。
- `os_name_send(os_name)` 是一个扩展函数，返回 `bytea`。
- `os_name` 是一个扩展定义的类型。
- `btree_os_name_ops` 是一个扩展定义的操作符类。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.3`。
- 控制文件将扩展标记为可重定位。
- 2026-07-28 审查期间前 GitHub 仓库 URL 返回 404；请将上述固定在 PGXN 分发视为可用源边界。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
