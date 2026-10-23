## 用法

来源：

- [Extension control](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/pg_documentdb_gw_host.control)
- [Cargo version and supported majors](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/Cargo.toml)
- [Worker initialization](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/src/bgworker.rs)
- [Required settings](https://github.com/documentdb/documentdb/blob/v1.0-0/pg_documentdb_gw_host/src/gucs.rs)
- [Official README](https://github.com/documentdb/documentdb/blob/v1.0-0/README.md)

`pg_documentdb_gw_host` 通过 PostgreSQL 后台工作进程运行 DocumentDB 协议网关。在 DocumentDB 1.0-0 中，该扩展版本仍为 0.1.0。源码声明支持 PostgreSQL 15–18，使用 pgrx 0.16.0；须确认所用发行包确实包含这一可选模块。

### 启用工作进程

以超级用户安装匹配的 DocumentDB 依赖与 SQL 对象：

```sql
CREATE EXTENSION documentdb CASCADE;
CREATE EXTENSION pg_documentdb_gw_host;
```

只安装 SQL 对象不会启用持久运行的网关。将 `pg_documentdb_gw_host` 加入部署现有的 `shared_preload_libraries`，把 `documentdb_gateway.database` 设置为安装扩展的数据库，并让 `documentdb_gateway.setup_configuration_file` 指向数据库主机上的有效网关配置文件。这两个设置仅允许超级用户修改，作用域为 postmaster；应在重启 PostgreSQL 前配置。

### 运行边界

工作进程连接配置的数据库，使用指定的初始化与 TLS 配置启动网关。缺少设置或无效配置会阻止启动。数据库连接失败会使工作进程退出，而非保持空闲；注册的工作进程重启间隔为一秒，应留意重复的启动错误。

该模块不提供独立的集合 API。客户端应通过配置的网关端点使用 DocumentDB 协议；完成身份验证、TLS、网络访问与连接额度配置后再开放监听端口。升级时应保持基础扩展与网关库版本匹配。
