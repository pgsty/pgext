## 用法

来源：

- [官方上游 README](https://github.com/tvondra/hashset/blob/3ef8368d43eb6508de4f2b0e0a5a1d5e497bd44b/README.md)
- [官方扩展控制文件 (hashset.control)](https://github.com/tvondra/hashset/blob/3ef8368d43eb6508de4f2b0e0a5a1d5e497bd44b/hashset.control)
- [官方扩展 SQL (hashset--0.0.1.sql)](https://github.com/tvondra/hashset/blob/3ef8368d43eb6508de4f2b0e0a5a1d5e497bd44b/hashset--0.0.1.sql)

`hashset` — 这个 PostgreSQL 扩展实现了 hashset，这是一种数据结构（类型），提供了一组唯一的整数项，并且具有快速查找的功能。当应用程序数据需要这种类型、域或其操作符时，请使用它。上游明确表示该项目尚未准备好用于生产环境。

### 核心工作流

```sql
CREATE EXTENSION hashset;

SELECT hashset_add(NULL, 1); -- {1}
SELECT hashset_add('{NULL}', 1); -- {1,NULL}
SELECT hashset_add('{1}', NULL); -- {1,NULL}
SELECT hashset_add('{1}', 1); -- {1}
SELECT hashset_add('{1}', 2); -- {1,2}
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hashset_add(int4hashset, int)` 是一个扩展函数，返回 `int4hashset`。
- `hashset_capacity(int4hashset)` 是一个扩展函数，返回 `bigint`。
- `hashset_cardinality(int4hashset)` 是一个扩展函数，返回 `bigint`。
- `hashset_cmp(int4hashset, int4hashset)` 是一个扩展函数，返回 `integer`。
- `hashset_collisions(int4hashset)` 是一个扩展函数，返回 `bigint`。
- `hashset_contains(int4hashset, int)` 是一个扩展函数，返回 `boolean`。
- `hashset_difference(int4hashset, int4hashset)` 是一个扩展函数，返回 `int4hashset`。
- `hashset_eq(int4hashset, int4hashset)` 是一个扩展函数，返回 `boolean`。
- `hashset_ge(int4hashset, int4hashset)` 是一个扩展函数，返回 `boolean`。
- `hashset_gt(int4hashset, int4hashset)` 是一个扩展函数，返回 `boolean`。
- `hashset_hash(int4hashset)` 是一个扩展函数，返回 `integer`。
- `hashset_intersection(int4hashset, int4hashset)` 是一个扩展函数，返回 `int4hashset`。
- `hashset_le(int4hashset, int4hashset)` 是一个扩展函数，返回 `boolean`。
- `hashset_lt(int4hashset, int4hashset)` 是一个扩展函数，返回 `boolean`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为可重定位。
- 上游明确表示该项目尚未准备好用于生产环境。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
