## 用法

来源：

- [官方上游 README](https://github.com/baofuhann/postgres/blob/f0cf77cfdf5c548de599237c47b7f63e72217024/contrib/README)
- [官方扩展控制文件 (alex_gist_int8.control)](https://github.com/baofuhann/postgres/blob/f0cf77cfdf5c548de599237c47b7f63e72217024/contrib/alex_gist_int8/alex_gist_int8.control)
- [官方扩展 SQL (alex_gist_int8--1.0.sql)](https://github.com/baofuhann/postgres/blob/f0cf77cfdf5c548de599237c47b7f63e72217024/contrib/alex_gist_int8/alex_gist_int8--1.0.sql)

`alex_gist_int8` — GiST 索引支持常见 PostgreSQL 数据类型。当应用程序需要此特定数据库功能时使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION alex_gist_int8;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `gbt_decompress(internal)` 是一个扩展函数，返回 `internal`。
- `gbt_int8_compress(internal)` 是一个扩展函数，返回 `internal`。
- `gbt_int8_consistent(internal,int8,int2,oid,internal)` 是一个扩展函数。
- `gbt_int8_distance(internal,int8,int2,oid,internal)` 是一个扩展函数，返回 `float8`。
- `gbt_int8_fetch(internal)` 是一个扩展函数，返回 `internal`。
- `gbt_int8_penalty(internal,internal,internal)` 是一个扩展函数，返回 `internal`。
- `gbt_int8_picksplit(internal, internal)` 是一个扩展函数，返回 `internal`。
- `gbt_int8_same(gbtreekey16, gbtreekey16, internal)` 是一个扩展函数，返回 `internal`。
- `gbt_int8_train(oid)` 是一个扩展函数。
- `gbt_int8_union(internal, internal)` 是一个扩展函数，返回 `gbtreekey16`。
- `gbtreekey16_in(cstring)` 是一个扩展函数，返回 `gbtreekey16`。
- `gbtreekey16_out(gbtreekey16)` 是一个扩展函数，返回 `cstring`。
- `gbtreekey8_in(cstring)` 是一个扩展函数，返回 `gbtreekey8`。
- `gbtreekey8_out(gbtreekey8)` 是一个扩展函数，返回 `cstring`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
