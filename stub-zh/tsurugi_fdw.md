## 用法

来源：

- [官方上游 README](https://github.com/project-tsurugi/tsurugi_fdw/blob/e777ab56b5cbff4df43f608299cd73739e92e7aa/README.md)
- [官方扩展控制文件 (tsurugi_fdw.control)](https://github.com/project-tsurugi/tsurugi_fdw/blob/e777ab56b5cbff4df43f608299cd73739e92e7aa/tsurugi_fdw.control)
- [官方扩展 SQL (tsurugi_fdw--1.4.0--1.5.0.sql)](https://github.com/project-tsurugi/tsurugi_fdw/blob/e777ab56b5cbff4df43f608299cd73739e92e7aa/tsurugi_fdw--1.4.0--1.5.0.sql)

`tsurugi_fdw` — tsurugi_fdw 是一个 PostgreSQL 扩展，提供了访问 Tsurugi 的外部数据封装。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION tsurugi_fdw;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `tg_execute_ddl(text DEFAULT null, text DEFAULT null)` 是一个扩展函数，返回 `TEXT`。
- `tg_set_exclusive_read_areas(variadic text[])` 是一个扩展函数，返回 `cstring`。
- `tg_set_inclusive_read_areas(variadic text[])` 是一个扩展函数，返回 `cstring`。
- `tg_set_transaction(text)` 是一个扩展函数，返回 `cstring`。
- `tg_set_transaction(text, text)` 是一个扩展函数，返回 `cstring`。
- `tg_set_transaction(text, text, text)` 是一个扩展函数，返回 `cstring`。
- `tg_set_write_preserve(variadic text[])` 是一个扩展函数，返回 `cstring`。
- `tg_show_tables(text DEFAULT null, text DEFAULT null, text DEFAULT 'detail', boolean DEFAULT true)` 是一个扩展函数，返回 `JSON`。
- `tg_show_transaction()` 是一个扩展函数，返回 `cstring`。
- `tg_verify_tables(text DEFAULT null, text DEFAULT null, text DEFAULT null, text DEFAULT 'summary', boolean DEFAULT true)` 是一个扩展函数，返回 `JSON`。
- `tsurugi_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `tsurugi_fdw_validator(options text[], catalog oid)` 是一个扩展函数，返回 `void`。
- `tsurugi_fdw` 是一个扩展定义的外部数据封装。

### 要求与注意事项

- 控制文件声明默认版本为 `1.5.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
