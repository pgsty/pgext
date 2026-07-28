## 用法

来源：

- [官方上游 README](https://github.com/michelp/pgs2/blob/2a193a0f76578eb78eada9041a5da8d85e449f3b/README.md)
- [官方扩展控制文件 (pgs2.control)](https://github.com/michelp/pgs2/blob/2a193a0f76578eb78eada9041a5da8d85e449f3b/pgs2.control)
- [官方扩展 SQL (pgs2--0.0.1.sql)](https://github.com/michelp/pgs2/blob/2a193a0f76578eb78eada9041a5da8d85e449f3b/pgs2--0.0.1.sql)

`pgs2` — Postgres 扩展用于 S2 球面几何。使用它来进行相应的空间数据或地理空间工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pgs2;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `btS2Cellsortsupport(internal)` 是一个扩展函数，返回 `void`。
- `S2Cap(center S2Point, radium float8 = 0.0)` 是一个扩展函数，返回 `S2Cap`。
- `S2Cap_eq(A S2Cap, B S2Cap)` 是一个扩展函数。
- `S2Cap_in(cstring)` 是一个扩展函数，返回 `S2Cap`。
- `S2Cap_out(S2Cap)` 是一个扩展函数，返回 `cstring`。
- `S2Cell_as_S2LatLng(c S2Cell)` 是一个扩展函数，返回 `S2LatLng`。
- `S2Cell_as_S2Point(c S2Cell)` 是一个扩展函数，返回 `S2Point`。
- `S2Cell_cmp(A S2Cell, B S2Cell)` 是一个扩展函数，返回 `int`。
- `S2Cell_distance(A S2Cell, B S2Cell)` 是一个扩展函数，返回 `float8`。
- `S2Cell_eq(A S2Cell, B S2Cell)` 是一个扩展函数。
- `S2Cell_ge(A S2Cell, B S2Cell)` 是一个扩展函数。
- `S2Cell_gt(A S2Cell, B S2Cell)` 是一个扩展函数。
- `S2Cell_in(cstring)` 是一个扩展函数，返回 `S2Cell`。
- `S2Cell_le(A S2Cell, B S2Cell)` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
