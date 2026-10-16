## 用法

来源：

- [src/pg_fdw/README.md](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/README.md)
- [src/pg_fdw/springtail_fdw--1.0.sql](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/springtail_fdw--1.0.sql)
- [src/pg_fdw/CMakeLists.txt](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/CMakeLists.txt)
- [LICENSE.md](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/LICENSE.md)
- [src/pg_fdw/springtail_fdw.control](https://github.com/Springtail-inc/springtail/blob/a6755b462920480787e0b80673861a6d7fd73d21/src/pg_fdw/springtail_fdw.control)

`springtail_fdw` 1.0 是 Springtail 面向 SQL 的存储适配器，需要已有的 Springtail 部署；仅安装扩展不会建立读取扩容集群。

### 核心用法

```sql
CREATE EXTENSION springtail_fdw;
CREATE SERVER springtail_fdw_server FOREIGN DATA WRAPPER springtail_fdw;
```

### 运行边界

使用前，将 `springtail_fdw.config_file_path` 指向部署中的 system.json。上游安装指南使用 PostgreSQL 16，注册服务器 `springtail_fdw_server`，并通过 IMPORT FOREIGN SCHEMA 导入 Springtail 命名空间。XID 管理守护进程必须运行，且 XID 不低于表的 XID。完整部署还有其他守护进程、Redis 和复制配置要求。默认仅超级用户可安装。源码采用 Elastic License 2.0，包含托管服务限制。
