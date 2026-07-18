## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/README.md)
- [0.0.1 版本 SQL API](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router--0.0.1.sql)
- [规划器钩子实现](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/src/planner/planner.c)
- [扩展 control 文件](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router.control)

`router` 是用于 OLTP 分片的实验性请求路由器。其 C 库安装规划器行为，并提供修改分片、分片键、键范围和本地表元数据的函数。

```sql
CREATE EXTENSION router;
SELECT mdbr_add_shard('shard1', '127.0.0.1', '5432', 'sharddb', 'router_user', 'replace-me');
SELECT create_sharding_key('customer_key', 'customer_id');
SELECT assign_key_range_2_shard('shard1', 'customer_key', 1, 100000);
```

SQL 文件中的加载提示把项目称为 `mdb_router`，但实际 control 与 SQL 文件名使用 `router`。上游只提供由 make 驱动的演示集群，没有受支持的安装与运维流程；不能从上述示例推断生产语义。

README 明确列出内存管理不良、缺少元数据一致性控制，以及重启后共享内存元数据丢失。源码注释也标明多条路径尚未完成。应把 0.0.1 视为废弃研究代码，绝不能作为生产分片层；任何评估都必须与真实数据和凭据隔离。
