## 用法

来源：

- [官方扩展控制文件（range_tools.control）](https://api.pgxn.org/src/range_tools/range_tools-0.1.2/range_tools.control)
- [官方扩展 SQL（range_tools.sql）](https://api.pgxn.org/src/range_tools/range_tools-0.1.2/sql/range_tools.sql)

`range_tools` — 用于范围类型工具。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION range_tools;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `range_type` 是一个扩展定义视图。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.2`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
