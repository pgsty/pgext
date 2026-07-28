## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/mvtbl/mvtbl-0.0.2/Readme.md)
- [官方扩展控制文件 (mvtbl.control)](https://api.pgxn.org/src/mvtbl/mvtbl-0.0.2/mvtbl.control)
- [官方扩展 SQL (mvtbl--0.0.1.sql)](https://api.pgxn.org/src/mvtbl/mvtbl-0.0.2/mvtbl--0.0.1.sql)

`mvtbl` — 一个用于轻松移动表到表空间的 PostgreSQL 扩展。在管理或自动化上述数据库行为时使用它。请使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION mvtbl;

SELECT pg_size_pretty(mvtbl('test','mvtbl_test_tblspace'));
 pg_size_pretty
----------------
 123 MB
(1 row)

SELECT pg_size_pretty(mvtbl('public.test','pg_default'));
 pg_size_pretty
----------------
 123 MB
(1 row)
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `mvtbl(tbl text, tblspace text)` 是一个扩展函数并返回 `bigint`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.0.2`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
