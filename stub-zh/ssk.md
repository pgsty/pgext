## 用法

来源：

- [官方上游 README](https://github.com/marthinl/ssk/blob/be8781c8a4d3cb943fde0e6f039207f087fd5c1e/README.md)
- [官方扩展控制文件 (ssk.control)](https://github.com/marthinl/ssk/blob/be8781c8a4d3cb943fde0e6f039207f087fd5c1e/ssk.control)
- [官方扩展 SQL (ssk--1.0.sql)](https://github.com/marthinl/ssk/blob/be8781c8a4d3cb943fde0e6f039207f087fd5c1e/sql/ssk--1.0.sql)

`ssk` — **注意：** PostgreSQL 作为 SSK 参考实现的宿主，展示了其在生产数据库环境中的通用适用性。尽管根植于 PostgreSQL，但 SSK 的概念可以扩展到任何关系型数据库。当应用程序需要这种特定的数据库功能时，请使用它。上游将其描述为一个概念验证。

### 核心工作流

```sql
CREATE EXTENSION ssk;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `cardinality(ssk)` 是一个扩展函数，返回 `bigint`。
- `length(ssk)` 是一个扩展函数，返回 `bigint`。
- `size(ssk)` 是一个扩展函数，返回 `bigint`。
- `ssk()` 是一个扩展函数，返回 `ssk`。
- `ssk(bigint)` 是一个扩展函数，返回 `ssk`。
- `ssk(bigint[])` 是一个扩展函数，返回 `ssk`。
- `ssk(integer)` 是一个扩展函数，返回 `ssk`。
- `ssk(integer[])` 是一个扩展函数，返回 `ssk`。
- `ssk_add(ssk, bigint)` 是一个扩展函数，返回 `ssk`。
- `ssk_add(ssk, integer)` 是一个扩展函数，返回 `ssk`。
- `ssk_add_comm(bigint, ssk)` 是一个扩展函数，返回 `ssk`。
- `ssk_agg_finalfunc(bigint)` 是一个扩展函数，返回 `ssk`。
- `ssk_agg_sfunc(bigint, bigint)` 是一个扩展函数，返回 `bigint`。
- `ssk_cmp(ssk, ssk)` 是一个扩展函数，返回 `int`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为可重定位。
- 上游将该项目描述为一个概念验证。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
