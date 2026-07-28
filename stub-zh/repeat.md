## 用法

来源：

- [官方上游 README](https://github.com/evancarroll/pg-srf-repeat-benchmark/blob/866ae8b1d657bb03cf71d7e6f52a5406afd64aca/README.md)
- [官方扩展控制文件 (repeat.control)](https://github.com/evancarroll/pg-srf-repeat-benchmark/blob/866ae8b1d657bb03cf71d7e6f52a5406afd64aca/repeat.control)
- [官方扩展 SQL (repeat--0.0.1.sql)](https://github.com/evancarroll/pg-srf-repeat-benchmark/blob/866ae8b1d657bb03cf71d7e6f52a5406afd64aca/repeat--0.0.1.sql)

`repeat` — PostgreSQL Set-Returning-Function (SRF) C-Extension Benchmarks ====. 使用此扩展时，当应用程序需要此特定数据库功能，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION repeat;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证安装版本和返回值。

### 重要对象

- `repeat_materialize(object int4, times int4)` 是一个扩展函数，返回 `TABLE`。
- `repeat_materialize_preferred(object int4, times int4)` 是一个扩展函数，返回 `TABLE`。
- `repeat_valuepercall(object int4, times int4)` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
