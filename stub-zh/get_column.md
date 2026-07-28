## 用法

来源：

- [官方扩展控制文件 (get_column.control)](https://api.pgxn.org/src/get_column/get_column-1.0.0/get_column.control)
- [官方扩展 SQL (get_column--1.0.sql)](https://api.pgxn.org/src/get_column/get_column-1.0.0/get_column--1.0.sql)

`get_column` — 通过名称从记录中获取列值。当 SQL 需要这些特殊函数或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION get_column;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `get_column(record, text)` 是一个扩展函数，返回 `anyelement`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为不可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
