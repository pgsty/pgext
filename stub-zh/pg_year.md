## 用法

来源：

- [官方上游 README](https://github.com/pierreforstmann/pg_year/blob/8e18cd6bd7872807b6e13e22741b23c6f60a5b54/README.md)
- [官方扩展控制文件 (pg_year.control)](https://github.com/pierreforstmann/pg_year/blob/8e18cd6bd7872807b6e13e22741b23c6f60a5b54/pg_year.control)
- [官方扩展 SQL (pg_year--0.0.1.sql)](https://github.com/pierreforstmann/pg_year/blob/8e18cd6bd7872807b6e13e22741b23c6f60a5b54/pg_year--0.0.1.sql)

`pg_year` — 该扩展已在 PostgreSQL 12、13、14、15、16、17 和 18 版本中验证通过。当应用程序需要此类型、域或其操作符时，请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_year;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hash_year(year)` 是一个扩展函数，返回 `integer`。
- `year_add(year, int)` 是一个扩展函数，返回 `year`。
- `year_cmp(year, year)` 是一个扩展函数，返回 `integer`。
- `year_eq(year, year)` 是一个扩展函数，返回 `boolean`。
- `year_ge(year, year)` 是一个扩展函数，返回 `boolean`。
- `year_gt(year, year)` 是一个扩展函数，返回 `boolean`。
- `year_in(cstring)` 是一个扩展函数，返回 `year`。
- `year_le(year, year)` 是一个扩展函数，返回 `boolean`。
- `year_lt(year, year)` 是一个扩展函数，返回 `boolean`。
- `year_minus(year, int)` 是一个扩展函数，返回 `year`。
- `year_ne(year, year)` 是一个扩展函数，返回 `boolean`。
- `year_out(year)` 是一个扩展函数，返回 `cstring`。
- `year` 是一个扩展定义的类型。
- `btree_year_ops` 是一个扩展定义的操作符类。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
