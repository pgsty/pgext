## 用法

来源：

- [官方上游 README](https://gitlab.com/byard1/pg_kalman/-/blob/main/README.md)
- [官方扩展控制文件](https://gitlab.com/byard1/pg_kalman/-/blob/main/pg_kalman.control)
- [官方项目页面](https://gitlab.com/byard1/pg_kalman)

`pg_kalman` — 一个简单的 Kalman 过滤器扩展用于 PostgreSQL。主要是为了好玩。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_kalman;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.1`。
- 控制文件标记该扩展为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
