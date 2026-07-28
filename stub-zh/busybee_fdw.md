## 用法

来源：

- [官方上游 README](https://gitlab.com/byard1/busybee_fdw/-/blob/master/README.md)
- [官方扩展控制文件](https://gitlab.com/byard1/busybee_fdw/-/blob/master/busybee_fdw.control)
- [官方项目页面](https://gitlab.com/byard1/busybee_fdw)

`busybee_fdw` — 一个 MQTT 外部数据封装器用于 PostgreSQL。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION busybee_fdw;

CREATE SERVER busybee_server FOREIGN DATA WRAPPER busybee_fdw;
```

在目标数据库中安装扩展，如果有可用示例，请运行上游最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
