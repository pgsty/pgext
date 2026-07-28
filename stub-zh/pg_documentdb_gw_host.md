## 用法

来源：

- [官方上游 README](https://github.com/documentdb/documentdb/blob/fd46318bf292780238eac4ef2f9e0011f7234539/README.md)
- [官方扩展控制文件 (pg_documentdb_gw_host.control)](https://github.com/documentdb/documentdb/blob/fd46318bf292780238eac4ef2f9e0011f7234539/pg_documentdb_gw_host/pg_documentdb_gw_host.control)
- [官方实现源代码](https://github.com/documentdb/documentdb/blob/fd46318bf292780238eac4ef2f9e0011f7234539/pg_documentdb_gw_host/src/lib.rs)

`pg_documentdb_gw_host` — pg documentdb gw host：由 pgrx 创建。在移植或模拟相应的数据库 API 时使用。必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pg_documentdb_gw_host;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了版本信息 `0.1.0`。
- 请首先安装确认的扩展依赖项：`documentdb`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
