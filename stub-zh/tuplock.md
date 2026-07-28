## 用法

来源：

- [官方扩展控制文件 (tuplock.control)](https://api.pgxn.org/src/tuplock/tuplock-1.2.2/tuplock.control)
- [官方扩展 SQL (tuplock.sql)](https://api.pgxn.org/src/tuplock/tuplock-1.2.2/sql/tuplock.sql)

`tuplock` — 使用布尔属性锁定元组（行）。在管理或自动化上述数据库行为时使用此扩展。请使用链接中的上游最新版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION tuplock;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `tuplock()` 是一个扩展函数，返回 `TRIGGER`。
- `test` 是由扩展安装或管理的表。
- `test2` 是由扩展安装或管理的表。
- `test_tuplock` 是由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.2.2`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
