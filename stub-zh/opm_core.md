## 用法

来源：

- [官方上游 README](https://github.com/opmdg/opm-core/blob/ae89f025407ab144e1e30abd7d6580f258945d61/README.md)
- [官方扩展控制文件 (opm_core.control)](https://github.com/opmdg/opm-core/blob/ae89f025407ab144e1e30abd7d6580f258945d61/pg/opm_core.control)
- [官方扩展 SQL (opm_core--2.5--2.6.sql)](https://github.com/opmdg/opm-core/blob/ae89f025407ab144e1e30abd7d6580f258945d61/pg/opm_core--2.5--2.6.sql)

`opm_core` — 中央模块，用于收集或解释相应的 PostgreSQL 统计信息。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION opm_core;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `public.clone_graph(p_id_graph bigint)` 是一个扩展函数，返回 `bigint`。
- `public.drop_account(IN p_account text)` 是一个扩展函数，返回 `TABLE`。
- `public.get_graph(p_id_graph bigint)` 是一个扩展函数，返回 `TABLE`。
- `public.get_sampled_metric_data(p_id_metric bigint, p_timet_begin timestamp with time zone, p_timet_end timestamp with time zone, p_sample_num integer)` 是一个扩展函数，返回 `TABLE`。
- `public.get_server(IN p_id bigint)` 是一个扩展函数，返回 `TABLE`。
- `public.get_service(IN p_id bigint)` 是一个扩展函数，返回 `TABLE`。
- `public.grant_appli(IN p_role name)` 是一个扩展函数，返回 `TABLE`。
- `public.grant_dispatcher(IN p_whname name, IN p_rolname name)` 是一个扩展函数，返回 `TABLE`。
- `public.js_time(timestamptz)` 是一个扩展函数，返回 `bigint`。
- `public.js_timetz(timestamptz)` 是一个扩展函数，返回 `bigint`。
- `public.list_accounts()` 是一个扩展函数，返回 `TABLE`。
- `public.list_graphs()` 是一个扩展函数，返回 `TABLE`。
- `public.list_graphs(p_id_server bigint)` 是一个扩展函数，返回 `TABLE`。
- `public.list_graphs_templates(IN p_id bigint)` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `2.6`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
