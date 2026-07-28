## 用法

来源：

- [官方上游 README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_delta/README.md)
- [官方扩展控制文件 (pg_delta.control)](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_delta/pg_delta.control)
- [官方实现源代码](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_delta/src/lib.rs)

`pg_delta` — PostgreSQL 与 Delta Lake 的流式集成。在 PostgreSQL 和 Delta Lake 表之间双向传输数据。当需要从 PostgreSQL 移动、转换或集成相应数据时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_delta;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `create_table` 是一个扩展函数。
- `drop_export` 是一个扩展函数。
- `drop_table` 是一个扩展函数。
- `export` 是一个扩展函数。
- `export_table` 是一个扩展函数。
- `extension_docs()` 是一个扩展函数。
- `history` 是一个扩展函数。
- `info` 是一个扩展函数。
- `list_exports()` 是一个扩展函数。
- `list_tables()` 是一个扩展函数。
- `read` 是一个扩展函数。
- `refresh` 是一个扩展函数。
- `schema` 是一个扩展函数。
- `status()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录版本 `0.2.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
