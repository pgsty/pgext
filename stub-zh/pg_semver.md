## 用法

来源：

- [官方上游 README](https://github.com/eendroroy/pg_semver/blob/42cddda9d5f36161841cd1fba5cd9fdf9a704cdb/README.md)
- [官方扩展控制文件 (pg_semver.control)](https://github.com/eendroroy/pg_semver/blob/42cddda9d5f36161841cd1fba5cd9fdf9a704cdb/pg_semver.control)
- [官方扩展 SQL (pg_semver--0.0.1.sql)](https://github.com/eendroroy/pg_semver/blob/42cddda9d5f36161841cd1fba5cd9fdf9a704cdb/pg_semver--0.0.1.sql)

`pg_semver` — **Version** 数据类型（SEMVER）用于 postgresql。当应用程序数据需要此数据类型、域或其操作符时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_semver;

SELECT PG_SEMVER_CMP('1.0.0-alpha.1', '1.0.0-alpha.2');
 pg_semver_cmp
---------------
            -1
(1 row)

SELECT PG_SEMVER_CMP('0.0.1', '0.0.1');
 pg_semver_cmp
---------------
             0
(1 row)

SELECT PG_SEMVER_CMP('0.0.2', '0.0.1');
 pg_semver_cmp
---------------
             1
(1 row)
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hash_ver(semver)` 是一个扩展函数，返回 `int`。
- `pg_semver_bump(semver, int)` 是一个扩展函数，返回 `semver`。
- `pg_semver_car(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_cmp(semver, semver)` 是一个扩展函数，返回 `int`。
- `pg_semver_eq(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_ge(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_gt(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_in(cstring)` 是一个扩展函数，返回 `semver`。
- `pg_semver_le(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_lt(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_ncar(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_ne(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_nsat(semver, semver)` 是一个扩展函数，返回 `boolean`。
- `pg_semver_out(semver)` 是一个扩展函数，返回 `cstring`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
