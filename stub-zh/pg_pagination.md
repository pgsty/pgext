## 用法

来源：

- [官方上游 README](https://github.com/arunsahoo-xt/pg_extensions/blob/63fe9163fe6864fbfd020c1e72507613780120c7/pg_pagination/README.md)
- [官方扩展控制文件 (pg_pagination.control)](https://github.com/arunsahoo-xt/pg_extensions/blob/63fe9163fe6864fbfd020c1e72507613780120c7/pg_pagination/pg_pagination.control)
- [官方实现源代码](https://github.com/arunsahoo-xt/pg_extensions/blob/63fe9163fe6864fbfd020c1e72507613780120c7/pg_pagination/src/lib.rs)

`pg_pagination` — 一个 **高性能的 PostgreSQL 扩展**，使用 pgrx 编写在 Rust 中。当 SQL 需要这些特殊功能或聚合时，请使用此扩展。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_pagination;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `paginate_by_cursor` 是一个扩展函数。
- `paginate_by_offset` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
