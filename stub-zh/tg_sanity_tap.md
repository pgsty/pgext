## 用法

来源：

- [官方扩展控制文件（tg_sanity_tap.control）](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/tg_sanity_tap.control)
- [官方扩展 SQL（tg_sanity_tap.sql）](https://api.pgxn.org/src/tg_sanity/tg_sanity-0.1.0/sql/tg_sanity_tap.sql)

`tg_sanity_tap` — pgtap 测试函数用于 tg 稳定性触发器。当应用程序需要此特定数据库功能时使用它。在安装扩展及其依赖项并验证它们之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION tg_sanity_tap;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `tg_sanity_tap(trigger_table regclass , trigger_name text , timing text , events text , trigger_arguments text)` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `0.1.0`。
- 首先安装并验证确认的扩展依赖项：`tg_sanity`, `pgtap`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
