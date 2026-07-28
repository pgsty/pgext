## 用法

来源：

- [官方上游 README](https://github.com/rajeshkumarblr/pg_bloom/blob/00dc31b46450d2032cc22069e4c8b83f20285f73/README.md)
- [官方扩展控制文件 (pg_bloom.control)](https://github.com/rajeshkumarblr/pg_bloom/blob/00dc31b46450d2032cc22069e4c8b83f20285f73/pg_bloom.control)
- [官方扩展 SQL (pg_bloom--0.0.1.sql)](https://github.com/rajeshkumarblr/pg_bloom/blob/00dc31b46450d2032cc22069e4c8b83f20285f73/pg_bloom--0.0.1.sql)

`pg_bloom` — 一个实现布隆过滤器数据结构的 PostgreSQL 扩展。当应用程序数据需要这种类型、域或其操作符时，请使用此扩展。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_bloom;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `bloom_contains(bloom, text)` 是一个扩展函数，返回 `boolean`。
- `bloom_in(cstring)` 是一个扩展函数，返回 `bloom`。
- `bloom_out(bloom)` 是一个扩展函数，返回 `cstring`。
- `bloom` 是一个扩展定义的类型。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
