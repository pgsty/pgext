## 用法

来源：

- [pgslot.control](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/pgslot.control)
- [README.md](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/README.md)
- [sql/pgslot--0.5.sql](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/sql/pgslot--0.5.sql)
- [scripts/roles.sql](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/scripts/roles.sql)
- [src/README.md](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/src/README.md)
- [Makefile](https://github.com/bugraaktug/pgslot/blob/48ea289780f2d1f163c1c617697c18cc396e3d5e/Makefile)

`pgslot` 记录复制槽历史，并计算 WAL 保留量、消费速率及下游流水线健康状态。SQL 扩展不会推进或删除复制槽。

### 核心用法

```sql
CREATE EXTENSION pgslot;
SELECT pgslot.collect();
SELECT * FROM pgslot.slot_health;
SELECT * FROM pgslot.wal_summary;
SELECT pgslot.prune(168);
```

### 运行边界

由超级用户安装。固定的 `pgslot` 模式提供 `slot_health`、`wal_summary` 和 `slot_pipeline`，`report_metric` 接收适配器指标。按上游脚本授予角色权限：`pgslot_monitor` 读取视图，`pgslot_collector` 执行采集与清理，`pgslot_adapter` 上报指标。原始历史表和特权函数已撤销 PUBLIC 权限。可在外部定时调用 `collect()` 与 `prune(integer)`；可选的 C 工作进程则要求预加载 `pgslot`、重启，并设置 `pgslot.database` 及属于采集组的登录角色 `pgslot.role`。该工作进程仅覆盖一个数据库，上游报告在 PostgreSQL 15 上测试。CLI/TUI 是可选配套程序。
