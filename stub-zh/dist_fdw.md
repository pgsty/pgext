## 用法

来源：

- [Installation SQL](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/bulkload/dist_fdw--1.0.sql)
- [Extension control file](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/bulkload/dist_fdw.control)
- [Implementation (dist_fdw.cpp)](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/bulkload/dist_fdw.cpp)

`dist_fdw` 是 openGauss 用于外部文件及分布式存储的批量导入适配器。源码和 SQL 依赖 openGauss 内核设施，并非可直接用于原生 PostgreSQL 的扩展。

### 启用与对象

应使用自带该模块的兼容 openGauss 安装。版本化 SQL 注册 `pg_catalog.dist_fdw_handler()`、`pg_catalog.dist_fdw_validator(text[], oid)` 和 `dist_fdw` 外部数据包装器。配置导入前，先检查当前实例是否已安装该包装器。

```sql
SELECT extname, extversion FROM pg_extension WHERE extname = 'dist_fdw';
SELECT fdwname FROM pg_foreign_data_wrapper WHERE fdwname = 'dist_fdw';
```

### 导入流程与限制

应依据实际 openGauss 部署支持的存储协议、位置、格式和错误处理选项创建外部服务器与外部表，再从外部表加载到指定本地目标。源码具有本地文件、远端和对象存储等不同路径，行为取决于部署，不能跨内核版本套用通用文件或对象存储示例。文件路径由数据库主机解析，导入导出权限与错误日志保留策略须由管理员审查。扩展/control 版本为 `1.0`，不代表外围内核版本。
