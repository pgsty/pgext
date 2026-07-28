## 用法

来源：

- [官方上游 README](https://github.com/evancarroll/pg-generate-up-down-series/blob/f7e5a0f9f1083efcd12dc4bbb906382d5acffc1e/README.md)
- [官方扩展控制文件 (generate_up_down_series.control)](https://github.com/evancarroll/pg-generate-up-down-series/blob/f7e5a0f9f1083efcd12dc4bbb906382d5acffc1e/generate_up_down_series.control)
- [官方扩展 SQL (generate_up_down_series--0.0.1.sql)](https://github.com/evancarroll/pg-generate-up-down-series/blob/f7e5a0f9f1083efcd12dc4bbb906382d5acffc1e/generate_up_down_series--0.0.1.sql)

`generate_up_down_series` — 首先你需要编译。在 Debian 系统上，你需要安装 postgresql-server-dev-all 和 build-essentials。然后你可以使用以下命令进行安装。在 SQL 中需要这些特殊函数或聚合时使用它。使用上述链接的上游版本作为 API 边界，并在目标 PostgreSQL 版本中进行测试。

### 核心工作流

```sql
CREATE EXTENSION generate_up_down_series;

SELECT *
FROM generate_up_down_series_evan(n,m);
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小 SQL 代码，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `generate_up_down_series_evan(n int4, m int4)` 是一个扩展函数，返回 `TABLE`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
