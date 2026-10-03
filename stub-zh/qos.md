## 用法

来源：

- [README.md](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/README.md)
- [qos.control](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos.control)
- [qos--1.0--1.1.sql](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos--1.0--1.1.sql)
- [qos--1.1.sql](https://github.com/appstonia/pg_qos/blob/fd3462b7fa81f8bb8aed2113a1f75bcf0e5dfe00/qos--1.1.sql)

`qos` 1.1（发行包 1.1.0）在 PostgreSQL 15+ 上按角色与数据库限制资源。先将 `qos` 加入 `shared_preload_libraries` 并重启，再由管理员创建 SQL 对象。CPU 亲和性限制要求 Linux。

### 配置限制

```sql
CREATE EXTENSION qos;
ALTER ROLE app_user SET qos.work_mem_limit = '32MB';
ALTER ROLE app_user SET qos.max_concurrent_select = '100';
ALTER ROLE app_user SET qos.max_select_rate = '10/500ms';
SELECT * FROM qos_stat_rate;
```

### 限制语义

`qos.work_mem_limit` 限制有效工作内存；`qos.cpu_core_limit` 在 Linux 上控制 CPU 亲和性，在其他平台限制并行工作进程。`qos.max_concurrent_tx`、`qos.max_concurrent_select`、`qos.max_concurrent_update`、`qos.max_concurrent_delete` 与 `qos.max_concurrent_insert` 限制并发操作。

`qos.max_tx_rate`、`qos.max_select_rate`、`qos.max_update_rate`、`qos.max_delete_rate` 与 `qos.max_insert_rate` 使用 100/1s 这样的次数/窗口组合，默认 -1 禁用各项速率限制。窗口为 100 毫秒至一天。速率或并发限制触发 SQLSTATE 54000，客户端应参考重试提示。适用的角色与数据库配置取最严格值，速率组合按规范化速率比较。

### 观测与升级

`qos_stat_rate` 暴露实时窗口，其他 `qos_stat` 视图提供活动与计数器。`qos_prometheus_metrics()` 输出 Prometheus 文本格式，计数器在服务重启后重置。1.1 用这些视图替换旧的占位函数 `qos_get_stats()`。

升级会改变共享内存布局，因此必须替换库并重启 PostgreSQL，再在各数据库执行 `ALTER EXTENSION qos UPDATE TO '1.1'`。新速率限制在配置前保持禁用。这些控制不能替代应用准入限制或操作系统隔离。
