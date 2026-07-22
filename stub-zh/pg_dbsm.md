## 用法

来源：

- [项目 README](https://github.com/japinli/pg_dbsm/blob/4d8f0928f3cc4923d4ba6b6d1c358a8fdba7d00e/README.md)
- [扩展 control 文件](https://github.com/japinli/pg_dbsm/blob/4d8f0928f3cc4923d4ba6b6d1c358a8fdba7d00e/pg_dbsm.control)
- [后台工作进程实现](https://github.com/japinli/pg_dbsm/blob/4d8f0928f3cc4923d4ba6b6d1c358a8fdba7d00e/pg_dbsm.c)

`pg_dbsm` 0.0.1 是一个后台工作进程扩展，用于定期记录集群中每个数据库的大小。工作进程会连接到一个指定数据库，在其中创建未限定 schema 的 `dbsm` 表，并持续写入数据库名、采集时间、总大小和大小增量。

### 配置工作进程

```conf
shared_preload_libraries = 'pg_dbsm'
dbsm.database = 'postgres'
dbsm.naptime = 86400
```

把 `pg_dbsm` 加入 `shared_preload_libraries` 后必须重启 PostgreSQL。`dbsm.database` 指定存储数据库，默认是 `postgres`；`dbsm.naptime` 是以秒计的采集间隔，默认值为 86400。随后可在该数据库中查询工作进程创建的表：

```sql
CREATE EXTENSION pg_dbsm;

SELECT datname, created, datsize, incsize
FROM dbsm
ORDER BY created DESC;
```

### 对象归属与故障行为

0.0.1 版安装 SQL 实际为空：`CREATE EXTENSION` 既不会注册后台工作进程，也不会创建或接管 `dbsm` 表。真正注册工作进程的是预加载库，表和索引由工作进程自行创建。因此，删除扩展不会停止仍在预加载的库，也不会清理这些对象。

每个数据库的第一条记录中，`incsize` 是空值而不是零：源码使用永远不可能为真的 `m.datsize = NULL` 判断，随后又减去不存在的历史值。编制增长报表时必须处理这一情况。

源码将工作进程配置为故障后不自动重启。应监控其进程和 PostgreSQL 日志、保护集中存储表，并制定保留策略，因为该表会持续增长。项目没有发布现代兼容矩阵；在生产使用前，应针对实际 PostgreSQL 大版本验证这份较早的 C 后台工作进程代码。
