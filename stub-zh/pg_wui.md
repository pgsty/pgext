## 用法

来源：

- [官方上游 README](https://github.com/siose-innova/pg_wui/blob/8b5e9016999bfcfddbbefd8dd2f17dfe5a69e6c3/README.md)
- [官方扩展控制文件 (pg_wui.control)](https://github.com/siose-innova/pg_wui/blob/8b5e9016999bfcfddbbefd8dd2f17dfe5a69e6c3/src/pg_wui.control)

`pg_wui` — PostgreSQL 扩展，用于使用 SIOSE 数据库进行野地城市界面数据的分析。请在相应的空间数据或地理空间工作流中使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_wui;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `0.2`。
- 首先安装确认的扩展依赖项：`postgis`。
- 控制文件将该扩展标记为不可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源的一致性。
