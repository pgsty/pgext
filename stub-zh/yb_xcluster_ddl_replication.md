## 用法

来源：

- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.control](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.control)
- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/README.md](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/README.md)
- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication--1.0.sql](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication--1.0.sql)
- [src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.c](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/src/postgres/yb-extensions/yb_xcluster_ddl_replication/yb_xcluster_ddl_replication.c)
- [LICENSE.md](https://github.com/yugabyte/yugabyte-db/blob/ad24ce143e5b9e28997082eecaac353374dc23a4/LICENSE.md)

`yb_xcluster_ddl_replication` 是 YugabyteDB 专用扩展，用于 xCluster 自动 DDL 复制。它记录源端 DDL 与目标端回放状态，实际回放由外部 xCluster 处理器执行。

### 核心用法

先通过 YugabyteDB 管理流程启用自动 xCluster 复制。在所引源码版本中，YugabyteDB 已将该库列入默认 `shared_preload_libraries`，xCluster 配置流程会在两端参与复制的数据库中创建扩展。随后可查询数据库的复制角色：

```sql
SELECT yb_xcluster_ddl_replication.get_replication_role();
```

### 运行边界

固定模式内包含 `ddl_queue`、`replicated_ddls` 及事件触发器处理函数。只有 `get_replication_role()` 授予 PUBLIC 执行权限，内部表与回放设置不属于应用 API。目标端用户 DDL 受到限制。安装脚本包含建表等复杂 DDL 的扩展，必须在数据库加入自动复制前安装。

所引源码修订的 control 版本为 1.0；它不能直接用于原生 PostgreSQL，也不是独立复制服务。
