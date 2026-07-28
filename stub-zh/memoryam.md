## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/memoryam/memoryam-0.0.1/README.md)
- [官方扩展控制文件 (memoryam.control)](https://api.pgxn.org/src/memoryam/memoryam-0.0.1/memoryam.control)
- [官方扩展 SQL (memoryam--0.0.1.sql)](https://api.pgxn.org/src/memoryam/memoryam-0.0.1/memoryam--0.0.1.sql)

`memoryam` — MemoryAM 是一个用 C++ 实现的 PostgreSQL 存储方法的内存临时表实现。其使命是简单地实现一个 TableAM 存储系统。因此，我们将在内存中存储所有更改，并仅允许单个连接访问。当应用程序需要此特定数据库功能时，请使用它。上游明确表示该项目尚未准备好生产使用。

### 核心工作流

```sql
CREATE EXTENSION memoryam;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `memoryam_relation_details(IN regclass, OUT row_number bigint, OUT xmin integer, OUT xmax integer, OUT is_deleted bool)` 是一个扩展函数，返回 `SETOF`。
- `memoryam_storage_details(OUT table_name text, OUT row_count bigint, OUT deleted_count bigint, OUT transaction_count bigint)` 是一个扩展函数，返回 `SETOF`。
- `memoryam_tableam_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `memoryam` 是一个扩展定义的访问方法。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为不可重定位。
- 上游明确表示该项目尚未准备好生产使用。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
