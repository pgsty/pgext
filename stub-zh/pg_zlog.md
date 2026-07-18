## 用法

来源：

- [上游架构与配置文档](https://github.com/cruzdb/pg_zlog/blob/5d428db977b0eec10c07f155a30db54b2c1af51c/README.md)
- [扩展控制文件](https://github.com/cruzdb/pg_zlog/blob/5d428db977b0eec10c07f155a30db54b2c1af51c/pg_zlog.control)

`pg_zlog` 是通过 Ceph 上的 ZLog 共享日志复制指定表的研究型扩展。节点访问已注册表前会推进日志，元数据则记录预期的 Ceph 集群与存储池标识。

```sql
CREATE EXTENSION pg_zlog;
SELECT pgzlog_add_cluster('ceph', '/etc/ceph/ceph.conf');
SELECT pgzlog_add_pool('ceph', 'pg_pool');
SELECT pgzlog_add_log('pg_pool', 'pg_log', 'seq-host', 5678);
SELECT pgzlog_replicate_table('pg_log', 'coordinates');
```

该项目已废弃，并依赖过时的 ZLog/Ceph 集成，不能替代受支持的 PostgreSQL 物理或逻辑复制。它只能用于隔离的历史测试环境；评估任何结果前，都要核实扩展预加载要求、Ceph 凭据、集群身份、重放顺序、DDL 行为、故障恢复和数据一致性。
