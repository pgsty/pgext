## 用法

来源：

- [官方上游 README](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/scheduling/README.md)
- [官方扩展控制文件 (scheduling.control)](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/scheduling/scheduling.control)
- [官方扩展 SQL (scheduling--0.1.0.sql)](https://github.com/constructive-io/pgpm-demo/blob/7793d0dca37c96d92e0cd4c16e25c8140142039e/packages/scheduling/sql/scheduling--0.1.0.sql)

`scheduling` — **🛠 构建于 Constructive 团队 — Postgres 工具化模块的创造者，致力于构建安全、可组合的后端。如果您喜欢我们的工作，请在 GitHub 上贡献。** 请使用此扩展进行相应的调度、时间序列或时间工作流。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION scheduling;
```

在目标数据库中安装扩展，并在可用时运行最小的上游示例，然后验证已安装的版本和返回值，再将其集成到应用程序 SQL 中。

### 重要对象

- `scheduling.appointments` 是由扩展安装或管理的表。
- `scheduling.encounters` 是由扩展安装或管理的表。
- `scheduling` 是由扩展创建的模式。

### 要求与注意事项

- 控制文件声明默认版本为 `0.1.0`。
- 首先安装已确认的扩展依赖项：`plpgsql`，`patients`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
