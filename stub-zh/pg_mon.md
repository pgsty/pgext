## 用法

来源：

- [上游 README](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/README.md)
- [扩展 SQL](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/pg_mon--1.0.sql)
- [扩展 control 文件](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/pg_mon.control)
- [上游 CI 矩阵](https://github.com/RafiaSabih/pg_mon/blob/e49626862d19300412ee4bdcec8d7bac65ebd4b4/.github/workflows/tests.yaml)

`pg_mon` 在共享内存中记录查询执行时间和行数直方图，以及扫描与连接信息。它必须预加载：

```conf
shared_preload_libraries = 'pg_mon'
```

重启 PostgreSQL 后创建扩展并查询其视图。如果普通用户不应清空监控状态，应限制重置函数：

```sql
CREATE EXTENSION pg_mon;
REVOKE EXECUTE ON FUNCTION pg_mon_reset() FROM PUBLIC;

SELECT queryid, total_time, actual_rows,
       seq_scans, index_scans, hash_join_count,
       hist_time_ubounds, hist_time_freq
FROM pg_mon
ORDER BY total_time DESC;
```

### 运维说明

版本 `1.0` 的上游 CI 矩阵覆盖 PostgreSQL 11 至 17。`pg_mon.plan_info_immediate` 会在规划后立即公开计划信息，但可能增加锁开销；`pg_mon.plan_info_disable` 会关闭计划详情。直方图边界固定，收集状态保存在内存中，并非持久历史。安装 SQL 向 `PUBLIC` 授予该视图的 `SELECT` 权限，因此共享系统中应审查可见性要求。
