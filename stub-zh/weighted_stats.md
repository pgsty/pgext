## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/weighted_stats/weighted_stats-1.0.0/README.md)
- [官方扩展控制文件 (weighted_stats.control)](https://api.pgxn.org/src/weighted_stats/weighted_stats-1.0.0/weighted_stats.control)
- [官方扩展 SQL (weighted_stats.sql)](https://api.pgxn.org/src/weighted_stats/weighted_stats-1.0.0/sql/weighted_stats.sql)

`weighted_stats` — 加权聚合函数。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION weighted_stats;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `weighted_mean` 是由扩展公开的聚合函数。
- `weighted_stddev_samp` 是由扩展公开的聚合函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
