## 用法

来源：

- [Installation SQL](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/mot/fdw_adapter/mot_fdw--1.0.sql)
- [Extension control file](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/mot/fdw_adapter/mot_fdw.control)
- [Build configuration](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/mot/fdw_adapter/Makefile)
- [MOT table workflow in openGauss 3.0](https://docs.opengauss.org/en/docs/3.0.0/docs/Developerguide/creating-dropping-an-mot-table.html)

`mot_fdw` 将 openGauss SQL 引擎连接到其内存优化表引擎。它要求启用 MOT 支持的 openGauss 构建，不是原生 PostgreSQL 的内存表实现。

### 启用与对象

应使用实际 openGauss 版本文档中的 MOT 服务器和外部表流程。安装 SQL 注册 `mot_fdw_handler()`、`mot_fdw_validator(text[], oid)` 和 `mot_fdw` 包装器。首先核实内核是否已安装扩展和包装器。

```sql
SELECT extname, extversion FROM pg_extension WHERE extname = 'mot_fdw';
SELECT srvname FROM pg_foreign_server
WHERE srvfdw = (SELECT oid FROM pg_foreign_data_wrapper WHERE fdwname = 'mot_fdw');
```

### 运行边界

MOT 表的内存、索引、事务集成、检查点与恢复由 openGauss MOT 引擎负责，不由普通网络 FDW 的远端服务器选项决定。承载需持久保存的数据前，应规划内存并遵循内核的备份恢复流程。control 版本 `1.0` 不代表 PostgreSQL 大版本，也不代表现代可移植二进制包。其 SQL 使用内核专有语言语法，不应将这些文件安装到原生 PostgreSQL。

### MOT 表用法

在已经启用 MOT 且提供 `mot_server` 的部署中，官方 3.0 文档使用以下表语法；列类型与容量限制仍以所用内核为准：

```sql
CREATE FOREIGN TABLE mot_example (x integer) SERVER mot_server;
INSERT INTO mot_example VALUES (1);
SELECT * FROM mot_example;
```
