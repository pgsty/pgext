## 用法

来源：

- [官方上游 README](https://github.com/hrbigelow/streamvis/blob/51018a5199d3a26196b96836c69ca8142c7338ba/db/README.md)
- [官方扩展控制文件 (streamvis_udfs.control)](https://github.com/hrbigelow/streamvis/blob/51018a5199d3a26196b96836c69ca8142c7338ba/db/udf/streamvis_udfs.control)
- [官方扩展 SQL (streamvis_udfs--1.0.sql)](https://github.com/hrbigelow/streamvis/blob/51018a5199d3a26196b96836c69ca8142c7338ba/db/udf/streamvis_udfs--1.0.sql)

`streamvis_udfs` — 一个用于自托管数据日志和可视化的日志客户端和服务端。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION streamvis_udfs;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `decode_float_enc(e enc_typ)` 是一个扩展函数，返回 `REAL[]`。
- `decode_int_enc(e enc_typ)` 是一个扩展函数，返回 `INT[]`。
- `decode_text_enc(e enc_typ)` 是一个扩展函数，返回 `TEXT[]`。
- `encode_bool_enc(p_vals BOOLEAN[])` 是一个扩展函数，返回 `enc_typ`。
- `encode_float_enc(p_vals FLOAT[])` 是一个扩展函数，返回 `enc_typ`。
- `encode_int_enc(p_vals INT[])` 是一个扩展函数，返回 `enc_typ`。
- `encode_text_enc(p_vals TEXT[])` 是一个扩展函数，返回 `enc_typ`。
- `window_avg_finalfunc(internal)` 是一个扩展函数，返回 `enc_typ`。
- `window_avg_sfunc(internal, int, int, enc_typ, enc_typ[])` 是一个扩展函数，返回 `internal`。
- `window_avg` 是由扩展公开的聚合函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
