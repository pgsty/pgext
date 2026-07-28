## 用法

来源：

- [官方上游 README](https://github.com/marcelomendoncasoares/pg_hlc/blob/dc7691ef484e6016f5de38eb9ee0ad35dbbdd39b/README.md)
- [官方扩展控制文件 (pg_hlc.control)](https://github.com/marcelomendoncasoares/pg_hlc/blob/dc7691ef484e6016f5de38eb9ee0ad35dbbdd39b/pg_hlc.control)
- [官方扩展 SQL (pg_hlc--0.1.0.sql)](https://github.com/marcelomendoncasoares/pg_hlc/blob/dc7691ef484e6016f5de38eb9ee0ad35dbbdd39b/pg_hlc--0.1.0.sql)

`pg_hlc` — 一个与 Dart CRDT 库完全兼容的 PostgreSQL 扩展 (pg_hlc)，提供 Hybrid Logical Clock (HLC) 功能。该扩展使用 pgrx 框架构建，并实现了与 Dart 参考实现完全相同的 HLC 算法和 API。当应用程序数据需要这种类型、领域或其操作符时，请使用此扩展。在目标 PostgreSQL 构建中测试上游版本链接中的固定修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_hlc;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `hlc_compare(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `integer`。
- `hlc_eq(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `boolean`。
- `hlc_from_date(text, text)` 是一个扩展函数，返回 `hlctimestamp`。
- `hlc_gt(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `boolean`。
- `hlc_gte(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `boolean`。
- `hlc_increment(text)` 是一个扩展函数，返回 `hlctimestamp`。
- `hlc_lt(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `boolean`。
- `hlc_lte(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `boolean`。
- `hlc_merge(text, hlctimestamp)` 是一个扩展函数，返回 `hlctimestamp`。
- `hlc_ne(hlctimestamp, hlctimestamp)` 是一个扩展函数，返回 `boolean`。
- `hlc_now(text)` 是一个扩展函数，返回 `hlctimestamp`。
- `hlc_parse(text)` 是一个扩展函数，返回 `hlctimestamp`。
- `hlc_reset(text)` 是一个扩展函数，返回 `boolean`。
- `hlc_to_string(hlctimestamp)` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 该扩展的版本记录在目录中，版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
