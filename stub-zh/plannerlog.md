## 用法

来源：

- [官方上游 README](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/README.md)
- [官方扩展控制文件 (plannerlog.control)](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/rust/pgx/plannerlog/plannerlog.control)
- [官方实现源代码](https://github.com/unknowntpo/playground-2022/blob/c2e4935c7a575bec01d6d5301d027f6adf801a62/rust/pgx/plannerlog/src/lib.rs)

`plannerlog` — pgrx 规划器挂钩示例，用于记录每个计划的查询。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION plannerlog;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `hello_hello_world()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
