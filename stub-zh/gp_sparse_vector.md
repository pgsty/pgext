## 用法

来源：

- [官方上游 README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_sparse_vector/README)
- [官方扩展控制文件 (gp_sparse_vector.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_sparse_vector/gp_sparse_vector.control)
- [官方扩展 SQL (gp_sparse_vector--1.0.0--1.0.1.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_sparse_vector/gp_sparse_vector--1.0.0--1.0.1.sql)

`gp_sparse_vector` — 例如，假设我们有一个存储在 Postgres 中的 "float8[]" 数组：'{0, 33,...40,000 零..., 12, 22 }'::float8[]。使用它作为相应的向量、模型或检索工作流的一部分。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION gp_sparse_vector;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `dimension(svec)` 是一个扩展函数，返回 `integer`。
- `dmax(float8,float8)` 是一个扩展函数，返回 `float8`。
- `dmin(float8,float8)` 是一个扩展函数，返回 `float8`。
- `dot(float8[],float8[])` 是一个扩展函数，返回 `float8`。
- `dot(float8[],svec)` 是一个扩展函数，返回 `float8`。
- `dot(svec,float8[])` 是一个扩展函数，返回 `float8`。
- `dot(svec,svec)` 是一个扩展函数，返回 `float8`。
- `float8arr_cast_float4(float4)` 是一个扩展函数，返回 `float8[]`。
- `float8arr_cast_float8(float8)` 是一个扩展函数，返回 `float8[]`。
- `float8arr_cast_int2(int2)` 是一个扩展函数，返回 `float8[]`。
- `float8arr_cast_int4(int4)` 是一个扩展函数，返回 `float8[]`。
- `float8arr_cast_int8(bigint)` 是一个扩展函数，返回 `float8[]`。
- `float8arr_cast_numeric(numeric)` 是一个扩展函数，返回 `float8[]`。
- `float8arr_div_float8arr(float8[],float8[])` 是一个扩展函数，返回 `svec`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0.1`。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
