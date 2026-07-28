## 用法

来源：

- [官方上游 README](https://github.com/polardb/polardb-for-postgresql/blob/b5f78e9e5af18ea3ffb54b8d638208ed7648edca/README.md)
- [官方扩展控制文件 (polar_vfs.control)](https://github.com/polardb/polardb-for-postgresql/blob/b5f78e9e5af18ea3ffb54b8d638208ed7648edca/src/polar_vfs/polar_vfs.control)
- [官方扩展 SQL (polar_vfs--1.0.sql)](https://github.com/polardb/polardb-for-postgresql/blob/b5f78e9e5af18ea3ffb54b8d638208ed7648edca/src/polar_vfs/polar_vfs--1.0.sql)

`polar_vfs` — 一个基于 PostgreSQL 开发的云原生数据库，由阿里云打造。当应用程序需要此特定数据库功能时，请使用它。使用上述链接的上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION polar_vfs;
```

在目标数据库中安装扩展，在可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `polar_libpfs_version()` 是一个扩展函数，返回 `text`。
- `polar_vfs_disk_expansion(text)` 是一个扩展函数。
- `polar_vfs_mem_status()` 是一个扩展函数，返回 `setof`。
- `mm_type` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件将扩展标记为可信。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
