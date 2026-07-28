## 用法

来源：

- [官方上游 README](https://github.com/adjust/pg_intmap/blob/5ddc90c908a79c2f97878008d8f98886ace4b3af/README.md)
- [官方扩展控制文件 (pg_intmap.control)](https://github.com/adjust/pg_intmap/blob/5ddc90c908a79c2f97878008d8f98886ace4b3af/pg_intmap.control)
- [官方扩展 SQL (pg_intmap--0.1.sql)](https://github.com/adjust/pg_intmap/blob/5ddc90c908a79c2f97878008d8f98886ace4b3af/pg_intmap--0.1.sql)

`pg_intmap` — 压缩整数到整数映射。当应用程序数据需要此类型、域或其操作符时使用。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION pg_intmap;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `intarr_get_val(intarr, int4)` 是一个扩展函数，返回 `int8`。
- `intarr_in(cstring)` 是一个扩展函数，返回 `intarr`。
- `intarr_out(intarr)` 是一个扩展函数，返回 `cstring`。
- `intmap(int8[], int8[])` 是一个扩展函数，返回 `intmap`。
- `intmap_get_val(intmap, int8)` 是一个扩展函数，返回 `int8`。
- `intmap_in(cstring)` 是一个扩展函数，返回 `intmap`。
- `intmap_meta(intmap)` 是一个扩展函数，返回 `cstring`。
- `intmap_out(intmap)` 是一个扩展函数，返回 `cstring`。
- `intarr` 是一个扩展定义的类型。
- `intmap` 是一个扩展定义的类型。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
