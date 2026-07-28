## 用法

来源：

- [官方上游 README](https://github.com/cjavdev/believe-sql/blob/84bda689e85dd82e9d0ed044db139d86a703d519/README.md)
- [官方扩展控制文件 (believe.control)](https://github.com/cjavdev/believe-sql/blob/84bda689e85dd82e9d0ed044db139d86a703d519/believe.control)

`believe` — > [!NOTE] > > Believe API PostgreSQL 扩展目前处于 **实验性** 阶段，我们期待您能够尝试使用它！请在相应的 SQL 或数据库实用工具工作流中使用它。上游将此功能描述为实验性。

### 核心工作流

```sql
CREATE EXTENSION believe;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.11.0`。
- 控制文件标记该扩展为可重定位。
- 上游将项目的一部分或全部标记为实验性。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源。
