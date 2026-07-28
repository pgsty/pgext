## 用法

来源：

- [官方上游 README](https://github.com/greengagedb/pxf/blob/4cf00f3908fd63481666646fcc3bf9a5c3df68e6/fdw/README.md)
- [官方扩展控制文件 (pxf_fdw.control)](https://github.com/greengagedb/pxf/blob/4cf00f3908fd63481666646fcc3bf9a5c3df68e6/fdw/pxf_fdw.control)
- [官方扩展 SQL (pxf_fdw--1.0.sql)](https://github.com/greengagedb/pxf/blob/4cf00f3908fd63481666646fcc3bf9a5c3df68e6/fdw/pxf_fdw--1.0.sql)

`pxf_fdw` — 此 Greengage 扩展实现了一个 PXF 外部数据封装器（FDW）。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION pxf_fdw;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pxf_fdw_handler()` 是一个扩展函数，返回 `fdw_handler`。
- `pxf_fdw_validator(text[], oid)` 是一个扩展函数，返回 `void`。
- `adl_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `file_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `gs_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `hbase_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `hdfs_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `hive_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `jdbc_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `s3_pxf_fdw` 是一个由扩展定义的外部数据封装器。
- `wasbs_pxf_fdw` 是一个由扩展定义的外部数据封装器。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源中的信息一致。
