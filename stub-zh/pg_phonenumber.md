## 用法

来源：

- [官方上游 README](https://github.com/greenbea/pg_phonenumber/blob/bb0ed97ddb393445b6bed288fb99f7c113118ddc/README.md)
- [官方扩展控制文件 (pg_phonenumber.control)](https://github.com/greenbea/pg_phonenumber/blob/bb0ed97ddb393445b6bed288fb99f7c113118ddc/pg_phonenumber.control)
- [官方扩展 SQL (pg_phonenumber--1.0.sql)](https://github.com/greenbea/pg_phonenumber/blob/bb0ed97ddb393445b6bed288fb99f7c113118ddc/pg_phonenumber--1.0.sql)

`pg_phonenumber` — C++ libphonenumber 支持的电话类型、相等运算符以及支持区域辅助函数。当应用程序数据需要此类型、域或其运算符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_phonenumber;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `get_supported_calling_codes()` 是一个扩展函数，返回 `INT[]`。
- `get_supported_regions()` 是一个扩展函数，返回 `TEXT[]`。
- `phone_eq(phone, phone)` 是一个扩展函数。
- `phone_in(cstring, oid, integer)` 是一个扩展函数，返回 `phone`。
- `phone_ne(phone, phone)` 是一个扩展函数。
- `phone_out(phone)` 是一个扩展函数，返回 `cstring`。
- `phone` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码进行比对。
