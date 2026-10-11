## 用法

来源：

- [gpdb-doc/markdown/ref_guide/modules/gp_pitr.html.md](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpdb-doc/markdown/ref_guide/modules/gp_pitr.html.md)
- [gpcontrib/gp_pitr/gp_pitr--1.1.sql](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_pitr/gp_pitr--1.1.sql)
- [gpcontrib/gp_pitr/Makefile](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_pitr/Makefile)
- [LICENSE](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/LICENSE)
- [gpcontrib/gp_pitr/gp_pitr.control](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_pitr/gp_pitr.control)

`gp_pitr` 为 Greenplum 6 提供集群级恢复辅助接口。1.1 版包含汇总归档状态的视图，以及供恢复工具使用的函数。仅安装此模块不会自动配置备份或 WAL 归档。

### 核心用法

```sql
CREATE EXTENSION gp_pitr;
SELECT * FROM gp_stat_archiver ORDER BY gp_segment_id;
```

### 运行边界

`gp_stat_archiver` 汇总协调节点和数据节点的归档统计，通过 `gp_segment_id` 区分来源；协调节点使用 -1。安装脚本向所有用户授予此视图的读取权限。

`gp_create_restore_point(text)` 在集群各节点创建命名恢复点，返回节点标识及恢复 LSN。`gp_switch_wal()` 切换 WAL 文件，返回节点标识、LSN 和文件名。两个函数都在协调节点执行，安装脚本撤销了 `PUBLIC` 的执行权限；它们是恢复工具使用的特权接口。

由超级用户注册模块。C 库按需加载，无需预加载配置，扩展不支持迁移模式。1.1 版提供从 1.0 升级的脚本。这些接口依赖 Greenplum 内核设施，不代表兼容原生 PostgreSQL。完整的恢复流程仍需配套的集群备份和 WAL 归档。
