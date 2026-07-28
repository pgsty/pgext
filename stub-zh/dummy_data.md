## 用法

来源：

- [官方上游 README](https://gitlab.com/3manuek/dummy_fdw/-/blob/master/README.md)
- [官方扩展控制文件](https://gitlab.com/3manuek/dummy_fdw/-/blob/master/dummy_data.control)
- [官方项目页面](https://gitlab.com/3manuek/dummy_fdw)

`dummy_data` — 一个适用于 PostgreSQL 9.3+ 的可读且支持空值的外部数据封装器。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION dummy_data;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
