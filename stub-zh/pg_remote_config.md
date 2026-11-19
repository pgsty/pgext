## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/pg_remote_config.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/pg_remote_config--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_remote_config/pg_remote_config.c)

`pg_remote_config` `1.0` 提供远程配置状态与管理表。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_remote_config;
SELECT * FROM remote_config_status();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

远程模式需预加载 pg_remote_config 并重启。通过管理员控制的配置设置 pg_remote_config.mode、pg_remote_config.host、pg_remote_config.port 和私密的 pg_remote_config.password。这是服务器控制面集成，远端来源可改变服务器配置。状态 SQL 本身不会启用远程管理。应严格限制管理表、令牌和网络访问。

此模块对应上游旧名 `alohadb_remote_config`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
