## 用法

来源：

- [官方上游 README](https://github.com/ffeller/mastermobilitydb/blob/60004d6b8f5ac491dcbad17f3a7f5bf0aab69e00/readme.txt)
- [官方扩展控制文件 (mastermobilitydb.control)](https://github.com/ffeller/mastermobilitydb/blob/60004d6b8f5ac491dcbad17f3a7f5bf0aab69e00/mastermobilitydb.control)
- [官方扩展 SQL (mastermobilitydb--1.0.0.sql)](https://github.com/ffeller/mastermobilitydb/blob/60004d6b8f5ac491dcbad17f3a7f5bf0aab69e00/mastermobilitydb--1.0.0.sql)

`mastermobilitydb` — 用于管理移动对象轨迹的 MobilityDB 架构、类型和实用工具。适用于相应的空间数据或地理空间工作流。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION mastermobilitydb;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `ASPECT_ATTRIBUTE_COUNT()` 是一个扩展函数，返回 `INTEGER`。
- `ASPECT_ATTRIBUTE_FIND_ALL()` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_ATTRIBUTE_FIND_BY_ID(IN P_ASPECT_ID INTEGER, IN P_ATTRIBUTE_ID INTEGER)` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_COUNT()` 是一个扩展函数，返回 `INTEGER`。
- `ASPECT_FIND_ALL()` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_FIND_BY_ID(IN P_ASPECT_ID INTEGER)` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_FIND_BY_NAME(IN P_DESCRIPTION VARCHAR(50))` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_TYPE_COUNT()` 是一个扩展函数，返回 `INTEGER`。
- `ASPECT_TYPE_FIND_ALL()` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_TYPE_FIND_BY_ID(IN P_ASPECT_TYPE_ID INTEGER)` 是一个扩展函数，返回 `SETOF`。
- `ASPECT_TYPE_FIND_BY_NAME(IN P_DESCRIPTION VARCHAR(50))` 是一个扩展函数，返回 `SETOF`。
- `ATTRIBUTE_COUNT()` 是一个扩展函数，返回 `INTEGER`。
- `ATTRIBUTE_FIND_ALL()` 是一个扩展函数，返回 `SETOF`。
- `ATTRIBUTE_FIND_BY_ID(IN P_ATTRIBUTE_ID INTEGER)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0.0`。
- 先安装并验证确认的扩展依赖项：`file_fdw`, `postgis`, `mobilitydb`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
