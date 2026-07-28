## 用法

来源：

- [官方上游 README](https://github.com/yuuch/pg_yaap/blob/41eda4a6020c9e0c0185586ffdd5ed4c55e94585/README.md)
- [官方扩展控制文件 (pg_yaap.control)](https://github.com/yuuch/pg_yaap/blob/41eda4a6020c9e0c0185586ffdd5ed4c55e94585/pg_yaap.control)
- [官方扩展 SQL (pg_yaap--1.0.sql)](https://github.com/yuuch/pg_yaap/blob/41eda4a6020c9e0c0185586ffdd5ed4c55e94585/pg_yaap--1.0.sql)

`pg_yaap` — pg_yaap 是一个 PostgreSQL 扩展，它用 YAAP 自有的优化器和 C++ 列式执行引擎替换了支持的分析查询的规划器和执行路径。在相应的分析或存储工作流中使用它。在目标 PostgreSQL 构建上使用链接的上游固定版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_yaap;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
