## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/pg_restapi.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/pg_restapi--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_restapi/restapi_http.c)

`pg_restapi` `1.0` 提供实验性的 SQL 与后台 HTTP CRUD 处理器。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_restapi;
SELECT * FROM pg_restapi_endpoints();
SELECT * FROM pg_restapi_status();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

SQL 处理器与网络工作进程分别启用。工作进程要求将 pg_restapi 加入 shared_preload_libraries 并重启，database、schema、port 与 default_role 控制暴露范围。源码接受 X-PG-Role 请求头并切换 SQL 角色，工作进程连接时也未指定独立数据库角色。这不构成安全的公网认证边界。认证及权限处理经独立审核前，应隔离监听端口，仅使用可丢弃的测试数据。

此模块对应上游旧名 `alohadb_restapi`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
