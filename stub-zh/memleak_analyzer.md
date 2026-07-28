## 用法

来源：

- [官方上游 README](https://github.com/vostralis/pg-memleak-analyzer/blob/f629e05e3fe7f0f67b26ab35974a2d476cc4c631/README.md)
- [官方扩展控制文件 (memleak_analyzer.control)](https://github.com/vostralis/pg-memleak-analyzer/blob/f629e05e3fe7f0f67b26ab35974a2d476cc4c631/memleak_analyzer.control)
- [官方扩展 SQL (memleak_analyzer--1.0.sql)](https://github.com/vostralis/pg-memleak-analyzer/blob/f629e05e3fe7f0f67b26ab35974a2d476cc4c631/memleak_analyzer--1.0.sql)

`memleak_analyzer` — pg-memleak-analyzer 是一个诊断工具和 PostgreSQL 扩展，用于分析和定位 PostgreSQL 后端会话和后台工作进程中的逻辑内存泄漏。在收集或解释相应的 PostgreSQL 统计信息时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION memleak_analyzer;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `memleak_analyzer.analyze_bgw(target_pid INTEGER, observation_interval INTEGER)` 是一个扩展函数，返回 `TABLE`。
- `memleak_analyzer.analyze_query(query text)` 是一个扩展函数，返回 `TABLE`。
- `memleak_analyzer.get_bgw_snapshot(target_pid INTEGER)` 是一个扩展函数，返回 `TABLE`。
- `memleak_analyzer` 是由扩展创建的一个模式。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
