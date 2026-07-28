## 用法

来源：

- [官方上游 README](https://github.com/elemdiscovery/proxquery/blob/74a314268002bd8b28542f630b0097b24a3e4dd7/README.md)
- [官方扩展控制文件 (proxquery.control)](https://github.com/elemdiscovery/proxquery/blob/74a314268002bd8b28542f630b0097b24a3e4dd7/proxquery.control)
- [官方实现源代码](https://github.com/elemdiscovery/proxquery/blob/74a314268002bd8b28542f630b0097b24a3e4dd7/src/lib.rs)

`proxquery` — proxquery 是一个 PostgreSQL 扩展，它在 tsvector 上添加了 @~@ 运算符，并提供更灵活的词项接近度搜索语法。使用它来进行相应的文本搜索、解析或语言学工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION proxquery;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `proxquery_build_query` 是一个扩展函数。
- `proxquery_is_analyzer` 是一个扩展函数。
- `proxquery_recheck` 是一个扩展函数。
- `ts_prox_chain` 是一个扩展函数。
- `ts_prox_not_within` 是一个扩展函数。
- `ts_prox_positions` 是一个扩展函数。
- `ts_prox_positions_prefix` 是一个扩展函数。
- `ts_prox_pre` 是一个扩展函数。
- `ts_prox_query_exact_cfg_droppable` 是一个扩展函数。
- `ts_prox_query_exact_string` 是一个扩展函数。
- `ts_prox_query_native_string` 是一个扩展函数。
- `ts_prox_query_skeleton` 是一个扩展函数。
- `ts_prox_query_support` 是一个扩展函数。
- `ts_prox_recheck` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了版本 `0.5.2`。
- 控制文件将扩展标记为可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为可信。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
