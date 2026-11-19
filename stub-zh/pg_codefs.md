## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/pg_codefs.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/pg_codefs--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_codefs/pg_codefs.c)

`pg_codefs` `1.0` 提供带版本历史的数据库文件工作区。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_codefs;
SELECT codefs_write('src/demo.txt', 'hello', 'main', 'initial');
SELECT codefs_read('src/demo.txt');
SELECT * FROM codefs_history('src/demo.txt');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

SQL 工作区由相对路径、版本、分支、锁和策略表组成。磁盘同步及策略工作进程要求在 shared_preload_libraries 中加入 pg_codefs 并重启；该可选模式还需配置 pg_codefs.database、pg_codefs.sync_interval 和 pg_codefs.policy_interval。策略变更由超级用户控制。磁盘同步属于外部副作用，需隔离文件系统，并通过授权保护源码与历史。

此模块对应上游旧名 `alohadb_codefs`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
