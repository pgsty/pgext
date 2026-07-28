## 用法

来源：

- [官方上游 README](https://github.com/saltypatron/hartonomous-001/blob/1ea12b4dfaa0a24b46ffa8c41290d9e1767fec4b/README.md)
- [官方扩展控制文件 (hartonomous.control)](https://github.com/saltypatron/hartonomous-001/blob/1ea12b4dfaa0a24b46ffa8c41290d9e1767fec4b/ext/hartonomous_pg/hartonomous.control)
- [官方扩展 SQL (hartonomous--1.0.sql)](https://github.com/saltypatron/hartonomous-001/blob/1ea12b4dfaa0a24b46ffa8c41290d9e1767fec4b/ext/hartonomous_pg/sql/hartonomous--1.0.sql)

`hartonomous` — Hartonomous substrate — 架构、类型、表、BLAKE3、S^3 几何、遍历、UCD/UCA 原子。当应用程序需要此特定数据库功能时使用它。上游明确表示该项目尚未准备好生产使用。

### 核心工作流

```sql
CREATE EXTENSION hartonomous;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `antipode(point4d)` 是一个扩展函数，返回 `point4d`。
- `array_to_linestring4d(double precision[])` 是一个扩展函数，返回 `linestring4d`。
- `array_to_point4d(double precision[])` 是一个扩展函数，返回 `point4d`。
- `bbox(linestring4d)` 是一个扩展函数，返回 `box4d`。
- `bbox(point4d)` 是一个扩展函数，返回 `box4d`。
- `bbox_4d_combine(box4d, box4d)` 是一个扩展函数，返回 `box4d`。
- `bbox_4d_sfunc(box4d, point4d)` 是一个扩展函数，返回 `box4d`。
- `bbox_expand(box4d, point4d)` 是一个扩展函数，返回 `box4d`。
- `bbox_union(box4d, box4d)` 是一个扩展函数，返回 `box4d`。
- `blake3_hash(bytea)` 是一个扩展函数，返回 `bytea`。
- `blake3_hash_text(text)` 是一个扩展函数，返回 `bytea`。
- `box4d_contained_by_box(box4d, box4d)` 是一个扩展函数，返回 `boolean`。
- `box4d_contains_box(box4d, box4d)` 是一个扩展函数，返回 `boolean`。
- `box4d_contains_point(box4d, point4d)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 请首先安装确认的扩展依赖项：`postgis`、`btree_gist`、`pg_trgm`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 上游明确表示该项目尚未准备好生产使用。
- 上游材料包含显式的弃用边界。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
