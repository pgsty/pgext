## 用法

来源：

- [官方上游 README](https://github.com/cloudivian-org/aetheldb/blob/012ee8637f0c9130e2033cbd6f59e480809e166d/compute/README.md)
- [官方扩展控制文件 (aethel_smgr.control)](https://github.com/cloudivian-org/aetheldb/blob/012ee8637f0c9130e2033cbd6f59e480809e166d/compute/extension/aethel_smgr/aethel_smgr.control)
- [官方扩展 SQL (aethel_smgr--1.0.sql)](https://github.com/cloudivian-org/aetheldb/blob/012ee8637f0c9130e2033cbd6f59e480809e166d/compute/extension/aethel_smgr/aethel_smgr--1.0.sql)

`aethel_smgr` — AethelDB - 一个解耦的、无服务器的 PostgreSQL 平台。当应用程序需要这种特定的数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION aethel_smgr;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `aethel_smgr_status()` 是一个扩展函数，返回 `text`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
