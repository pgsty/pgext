## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/pg_tiering.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/pg_tiering--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tiering/pg_tiering.c)

`pg_tiering` `1.0` 提供按规则迁移分区表空间。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_tiering;
SELECT * FROM pg_tiering_status();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

pg_tiering_rules 指定父表、年龄阈值和目标表空间。pg_tiering_check_now 可手动检查；周期任务需预加载 pg_tiering、重启并设置 pg_tiering.database 和 pg_tiering.check_interval。关系迁移使用 ALTER TABLE SET TABLESPACE，会消耗锁与 I/O。启用规则前需核对可用空间、所有权及分区年龄选择；这是本地表空间迁移，不是对象存储归档。

此模块对应上游旧名 `alohadb_tiering`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
