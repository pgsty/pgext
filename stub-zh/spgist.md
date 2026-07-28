## 用法

来源：

- [官方上游 README](https://github.com/duckbill/sp-gist/blob/fd8516ec426ae0c7f0ae508e326fae7dcf674d1d/README.md)
- [官方扩展控制文件 (spgist.control)](https://github.com/duckbill/sp-gist/blob/fd8516ec426ae0c7f0ae508e326fae7dcf674d1d/spgist.control)
- [官方扩展 SQL (spgist--1.0.sql)](https://github.com/duckbill/sp-gist/blob/fd8516ec426ae0c7f0ae508e326fae7dcf674d1d/spgist--1.0.sql)

`spgist` — 可用于测试 sp-gist 索引（postresql）。在应用程序需要此特定数据库功能时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION spgist;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `spg_quad_choose(internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_quad_config(internal)` 是一个扩展函数，返回 `internal`。
- `spg_quad_inner_consistent(internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_quad_leaf_consistent(internal, internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_quad_picksplit(internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_text_choose(internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_text_config(internal)` 是一个扩展函数，返回 `internal`。
- `spg_text_inner_consistent(internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_text_leaf_consistent(internal, internal, internal)` 是一个扩展函数，返回 `internal`。
- `spg_text_picksplit(internal, internal)` 是一个扩展函数，返回 `internal`。
- `spgbeginscan(internal)` 是一个扩展函数，返回 `internal`。
- `spgbuild(internal)` 是一个扩展函数，返回 `internal`。
- `spgbuildempty(internal)` 是一个扩展函数，返回 `internal`。
- `spgbulkdelete(internal)` 是一个扩展函数，返回 `internal`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
