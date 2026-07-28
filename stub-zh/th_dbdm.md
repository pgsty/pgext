## 用法

来源：

- [官方上游 README](https://github.com/xingjianwei/pg_analytics/blob/924deb4044b4e1c40766cd342ec990bf47066702/README.md)
- [官方扩展控制文件 (th_dbdm.control)](https://github.com/xingjianwei/pg_analytics/blob/924deb4044b4e1c40766cd342ec990bf47066702/th_dbdm.control)
- [官方实现源代码](https://github.com/xingjianwei/pg_analytics/blob/924deb4044b4e1c40766cd342ec990bf47066702/src/lib.rs)

`th_dbdm` — pg_analytics（以前名为 pg_lakehouse）将 DuckDB 内置于 Postgres 中。安装 pg_analytics 后，Postgres 可以查询像 AWS S3 这样的外部对象存储以及像 Iceberg 或 Delta Lake 这样的表格式。查询会被推送到 DuckDB，这是一个高性能的分析查询引擎。使用它来进行相应的分析或存储工作流。经过审核的上游材料已将此功能标记为弃用。

### 核心工作流

```sql
CREATE EXTENSION th_dbdm;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 要求与注意事项

- 该目录记录了版本信息 `1.3.3`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 上游材料包含显式的弃用边界。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
