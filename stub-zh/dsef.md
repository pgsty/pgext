## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/dsef/dsef-2024.4.9/README.md)
- [官方扩展控制文件 (dsef.control)](https://api.pgxn.org/src/dsef/dsef-2024.4.9/dsef.control)
- [官方扩展 SQL (dsef--unpackaged--2024.4.9.sql)](https://api.pgxn.org/src/dsef/dsef-2024.4.9/sql/dsef--unpackaged--2024.4.9.sql)

`dsef` — *详细的 SQL 报告用于第三方帮助和支持*。在收集或解释相应的 PostgreSQL 统计信息时使用它。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION dsef;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `ds_capture()` 是一个扩展函数，返回 `int`。
- `ds_insert(p_run int)` 是一个扩展函数，返回 `int`。
- `ds_report(IN global_scope boolean DEFAULT TRUE,IN all_rows boolean DEFAULT FALSE)` 是一个扩展函数，返回 `TABLE`。
- `ds_report_diff(IN global_scope boolean DEFAULT TRUE,IN all_rows boolean DEFAULT FALSE,IN cnt_diff_pct_threshold numeric DEFAULT 0)` 是一个扩展函数，返回 `TABLE`。
- `ds_set(p_setting text)` 是一个扩展函数，返回 `int`。
- `ds_start()` 是一个扩展函数，返回 `int`。
- `ds_version()` 是一个扩展函数，返回 `text`。
- `explain_analyze_full(p_sql text,p_format text DEFAULT 'TEXT',p_verbose boolean DEFAULT false)` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `2024.4.9`。
- 先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
