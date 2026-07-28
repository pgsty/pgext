## 用法

来源：

- [官方扩展控制文件 (pg_sqlite_fs.control)](https://github.com/silverdaz/pg_sqlite_fs/blob/87cbad18f8d31f47287e92991a21a96feb9c49c1/pg_sqlite_fs.control)
- [官方扩展 SQL (pg_sqlite_fs.sql)](https://github.com/silverdaz/pg_sqlite_fs/blob/87cbad18f8d31f47287e92991a21a96feb9c49c1/pg_sqlite_fs.sql)

`pg_sqlite_fs` — SQLite 文件系统创建。使用它进行相应的 SQL 或数据库实用程序工作流。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_sqlite_fs;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `delete_attribute(filename text, inode bigint, name text)` 是一个扩展函数并返回 `boolean`。
- `delete_attributes(filename text, inode bigint)` 是一个扩展函数并返回 `boolean`。
- `delete_entry(text, bigint)` 是一个扩展函数并返回 `SETOF bigint`。
- `insert_attribute(filename text, inode bigint, name text, value text)` 是一个扩展函数并返回 `boolean`。
- `insert_entries(text, text)` 是一个扩展函数并返回 `boolean`。
- `insert_entry(text, bigint, text, bigint, ctime bigint DEFAULT 0, mtime bigint DEFAULT 0, nlink bigint DEFAULT 1, size bigint DEFAULT 0, is_dir boolean DEFAULT TRUE)` 是一个扩展函数并返回 `void`。
- `insert_file(filename text, inode bigint, mountpoint text, relative_path text, header bytea, payload_size bigint, prepend bytea, append bytea)` 是一个扩展函数并返回 `boolean`。
- `insert_files(text, text)` 是一个扩展函数并返回 `boolean`。
- `make(filepath text, umask integer default 0o007)` 是一个扩展函数并返回 `boolean`。
- `remove(text)` 是一个扩展函数并返回 `boolean`。
- `truncate_attributes(text)` 是一个扩展函数并返回 `boolean`。
- `truncate_entries(text)` 是一个扩展函数并返回 `boolean`。
- `truncate_files(text)` 是一个扩展函数并返回 `boolean`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
