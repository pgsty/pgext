## 用法

来源：

- [官方上游 README](https://gitlab.com/daamien/pgrx-tuto/-/blob/main/README.md)
- [官方扩展控制文件](https://gitlab.com/daamien/pgrx-tuto/-/blob/main/world/world.control)
- [官方项目页面](https://gitlab.com/daamien/pgrx-tuto)

`world` — 该仓库包含一系列使用 PGRX 开发的 PostgreSQL 扩展。当 SQL 需要这些特殊函数或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION world;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
