## 用法

来源：

- [官方上游 README](https://github.com/dalibo/opm/blob/6b757bfe413cb1a21d736cdc648971b4a2416213/README)
- [官方扩展控制文件 (pr_grapher.control)](https://github.com/dalibo/opm/blob/6b757bfe413cb1a21d736cdc648971b4a2416213/processes/pr_grapher/pr_grapher.control)
- [官方扩展 SQL (pr_grapher--1.0--1.1.sql)](https://github.com/dalibo/opm/blob/6b757bfe413cb1a21d736cdc648971b4a2416213/processes/pr_grapher/pr_grapher--1.0--1.1.sql)

`pr_grapher` — Grapher 进程用于 OPM。在收集或解释相应的 PostgreSQL 统计信息时使用它。在将其集成到应用程序 SQL 中之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pr_grapher;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pr_grapher.delete_graph(p_id bigint)` 是一个扩展函数，返回 `boolean`。
- `pr_grapher.get_categories()` 是一个扩展函数，返回 `TABLE`。
- `pr_grapher.js_time(timestamptz)` 是一个扩展函数，返回 `bigint`。
- `pr_grapher.js_timetz(timestamptz)` 是一个扩展函数，返回 `bigint`。
- `pr_grapher.list_graph()` 是一个扩展函数，返回 `TABLE`。
- `pr_grapher.categories` 是一个由扩展安装或管理的表。
- `pr_grapher.graph_categories` 是一个由扩展安装或管理的表。
- `pr_grapher.graphs` 是一个由扩展安装或管理的表。
- `pr_grapher.nested_categories` 是一个由扩展安装或管理的表。
- `pr_grapher.series` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.1`。
- 先安装并验证确认的扩展依赖项：`opm_core`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
