## 用法

来源：

- [官方上游 README](https://github.com/hyz1840/pg_jiebaparser/blob/a244c1332f02b1383a822c0c952d55b6949da329/README.md)
- [官方扩展控制文件 (jiebaparser.control)](https://github.com/hyz1840/pg_jiebaparser/blob/a244c1332f02b1383a822c0c952d55b6949da329/jiebaparser.control)
- [官方扩展 SQL (jiebaparser.sql)](https://github.com/hyz1840/pg_jiebaparser/blob/a244c1332f02b1383a822c0c952d55b6949da329/jiebaparser.sql)

`jiebaparser` — 用于中文全文搜索的 Postgresql 扩展（jiaba 引擎），使用共享内存。请在相应的文本搜索、解析或语言工作流中使用它。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION jiebaparser;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `jbprs_end(internal)` 是一个扩展函数，返回 `void`。
- `jbprs_getlexeme(internal, internal, internal)` 是一个扩展函数，返回 `internal`。
- `jbprs_lextype(internal)` 是一个扩展函数，返回 `internal`。
- `jbprs_start(internal, int4)` 是一个扩展函数，返回 `internal`。
- `jbprs_start_q(internal, int4)` 是一个扩展函数，返回 `internal`。
- `jiebaparser_reset()` 是一个扩展函数，返回 `void`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以与固定源进行比对。
