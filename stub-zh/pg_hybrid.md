## 用法

来源：

- [官方上游 README](https://github.com/daviszhen/pg_hybrid/blob/7e5d30c151d263f5101c3f61cef969a67646dadc/README.md)
- [官方扩展控制文件 (pg_hybrid.control)](https://github.com/daviszhen/pg_hybrid/blob/7e5d30c151d263f5101c3f61cef969a67646dadc/pg_hybrid.control)
- [官方扩展 SQL (pg_hybrid--1.0.sql)](https://github.com/daviszhen/pg_hybrid/blob/7e5d30c151d263f5101c3f61cef969a67646dadc/pg_hybrid--1.0.sql)

`pg_hybrid` — PostgreSQL 16 扩展，提供 IVFFlat 索引访问方法以支持向量相似性搜索。使用此扩展进行相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_hybrid;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `array_to_hvector(double precision[], integer, boolean)` 是一个扩展函数，返回 `hvector`。
- `array_to_hvector(integer[], integer, boolean)` 是一个扩展函数，返回 `hvector`。
- `array_to_hvector(numeric[], integer, boolean)` 是一个扩展函数，返回 `hvector`。
- `array_to_hvector(real[], integer, boolean)` 是一个扩展函数，返回 `hvector`。
- `hvector(hvector, integer, boolean)` 是一个扩展函数，返回 `hvector`。
- `hvector_accum(double precision[], hvector)` 是一个扩展函数，返回 `double`。
- `hvector_add(hvector, hvector)` 是一个扩展函数，返回 `hvector`。
- `hvector_avg(double precision[])` 是一个扩展函数，返回 `hvector`。
- `hvector_binary_quantize(hvector)` 是一个扩展函数，返回 `bit`。
- `hvector_cmp(hvector, hvector)` 是一个扩展函数，返回 `int4`。
- `hvector_combine(double precision[], double precision[])` 是一个扩展函数，返回 `double`。
- `hvector_concat(hvector, hvector)` 是一个扩展函数，返回 `hvector`。
- `hvector_cosine_distance(hvector, hvector)` 是一个扩展函数，返回 `float8`。
- `hvector_dims(hvector)` 是一个扩展函数，返回 `integer`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
