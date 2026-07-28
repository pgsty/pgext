## 用法

来源：

- [官方上游 README](https://github.com/ggnmstr/extrema/blob/75fd0c4e50957cc59428e2dbb6933c91fbb82e6c/README.md)
- [官方扩展控制文件 (extrema.control)](https://github.com/ggnmstr/extrema/blob/75fd0c4e50957cc59428e2dbb6933c91fbb82e6c/extrema.control)
- [官方扩展 SQL (extrema--1.0.sql)](https://github.com/ggnmstr/extrema/blob/75fd0c4e50957cc59428e2dbb6933c91fbb82e6c/extrema--1.0.sql)

`extrema` — 此扩展允许用户通过将其他扩展添加到相应的 cgroups 来限制其资源使用（目前仅限 CPU、RAM、VmSwap 和 cpuset）。这些限制可以通过 PostgreSQL 的 GUC 机制轻松配置。在管理或自动化上述数据库行为时使用此扩展。请使用链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION extrema;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `ema_lib_info()` 是一个扩展函数，返回 `SETOF`。
- `ema_reload()` 是一个扩展函数，返回 `VOID`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
