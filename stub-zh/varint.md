## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/varint/varint-0.1.0/README)
- [官方扩展控制文件 (varint.control)](https://api.pgxn.org/src/varint/varint-0.1.0/varint.control)
- [官方扩展 SQL (varint.sql)](https://api.pgxn.org/src/varint/varint-0.1.0/sql/varint.sql)

`varint` — PostgreSQL-varint 介绍 PostgreSQL-varint 是一种数据类型，用于在 PostgreSQL 中以可变宽度编码整数以节省空间。当应用程序数据需要此类型、域或其操作符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION varint;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `int2(varint64)` 是一个扩展函数，返回 `int2`。
- `int2(varuint64)` 是一个扩展函数，返回 `int2`。
- `int4(varint64)` 是一个扩展函数，返回 `int4`。
- `int4(varuint64)` 是一个扩展函数，返回 `int4`。
- `int8(varint64)` 是一个扩展函数，返回 `int8`。
- `int8(varuint64)` 是一个扩展函数，返回 `int8`。
- `varint64(int2)` 是一个扩展函数，返回 `varint64`。
- `varint64(int4)` 是一个扩展函数，返回 `varint64`。
- `varint64(int8)` 是一个扩展函数，返回 `varint64`。
- `varint64_cmp(varint64, varint64)` 是一个扩展函数，返回 `int4`。
- `varint64_eq(varint64, varint64)` 是一个扩展函数。
- `varint64_ge(varint64, varint64)` 是一个扩展函数。
- `varint64_gt(varint64, varint64)` 是一个扩展函数。
- `varint64_in(cstring)` 是一个扩展函数，返回 `varint64`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
