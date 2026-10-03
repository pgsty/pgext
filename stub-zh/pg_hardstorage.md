## 用法

来源：

- [Official pg_hardstorage.control](https://github.com/cybertec-postgresql/pg_hardstorage/blob/b47541b7e1cea69ce6ec63b26e154eb25fc4ca91/ext/pg_hardstorage_extension/pg_hardstorage.control)
- [Official pg_hardstorage--1.0.sql](https://github.com/cybertec-postgresql/pg_hardstorage/blob/b47541b7e1cea69ce6ec63b26e154eb25fc4ca91/ext/pg_hardstorage_extension/pg_hardstorage--1.0.sql)
- [Official pg_hardstorage_db_install-extension.md](https://github.com/cybertec-postgresql/pg_hardstorage/blob/b47541b7e1cea69ce6ec63b26e154eb25fc4ca91/docs/reference/cli/pg_hardstorage_db_install-extension.md)

`pg_hardstorage` SQL 扩展 1.0 通过表、视图和写入函数暴露备份状态。它是独立 Go 备份命令行工具的数据库集成界面；创建扩展不会启动备份或 WAL 流。

### 启用与查看

```sql
CREATE EXTENSION pg_hardstorage;
SELECT * FROM pg_hardstorage.backups;
SELECT * FROM pg_hardstorage.health;
SELECT * FROM pg_hardstorage.rpo;
```

### 写入与读取权限

扩展文件安装后，由超级用户在固定模式中创建扩展。安装还会在集群中创建尚不存在的 `pg_hardstorage_writer` NOLOGIN 角色。读取方获得模式使用权与三个公开视图的 SELECT 权限。`pg_hardstorage.upsert_backup()`、`pg_hardstorage.upsert_health()` 与 `pg_hardstorage.upsert_rpo()` 是固定搜索路径的 SECURITY DEFINER 函数；执行函数和直接写入状态表的权限仅授予写入角色。仅应让可信的状态发布者成为该角色的成员。

### 数据填充与维护

此源码版本的内置备份流程尚未调用这些写入函数，因此视图在操作人员发布状态前保持为空。命令行工具另提供 `pg_hardstorage db install-extension`，带有试运行和打印 SQL 选项；该路径安装其内嵌 SQL 界面，不应与已注册扩展的升级混淆。SQL 扩展在安装时使用 PL/pgSQL，不需要 PostgreSQL 共享库或预加载，也未单独给出 PostgreSQL 大版本支持矩阵。命令行工具的发行版本与 SQL 版本 1.0 不同。
