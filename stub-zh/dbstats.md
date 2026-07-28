## 用法

来源：

- [官方上游 README](https://gitlab.com/ottohahn/dbstats/-/blob/master/README.md)
- [官方扩展控制文件](https://gitlab.com/ottohahn/dbstats/-/blob/master/dbstats.control)
- [官方项目页面](https://gitlab.com/ottohahn/dbstats)

`dbstats` — 一个全面的 PL/pgSQL PostgreSQL 扩展，用于快速数据计算、统计分析和数据质量评估，直接在数据库中进行。当 SQL 需要这些特殊函数或聚合时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION dbstats;
```

在目标数据库中安装扩展，如果有可用示例，请运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
