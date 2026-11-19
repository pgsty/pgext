## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/pg_learned_stats.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/pg_learned_stats--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_learned_stats/learned_stats.c)

`pg_learned_stats` `1.0` 提供供规划器使用的实验性学习式基数估计。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_learned_stats;
SELECT * FROM pg_learned_stats_status();
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

使用函数前需将 pg_learned_stats 加入 shared_preload_libraries 并重启。pg_learned_stats.enabled 控制钩子，pg_learned_stats.train_interval 控制训练间隔，pg_learned_stats_train 与 pg_learned_stats_reset 管理学习状态。模块使用 AlohaDB 规划器集成及共享内存。应以代表性负载演练计划变化，学习式估计不保证生成更好的执行计划。

此模块对应上游旧名 `alohadb_learned_stats`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
