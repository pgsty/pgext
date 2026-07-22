## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/README.md)
- [0.0.1 版本 SQL API](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router--0.0.1.sql)
- [分片元数据实现](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/src/mdbr/shard.c)
- [规划器钩子实现](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/src/planner/planner.c)
- [扩展 control 文件](https://github.com/reshke/router/blob/4f0f1654a925f70b3496abaaad91780d7b45c7c2/router.control)

`router` 是一个未完成的 OLTP 分片请求路由概念验证。规范扩展名是 `router`，但其 SQL 文件中的旧注释仍称 `mdb_router`。规划器钩子与内存元数据代码不能安全用于生产数据或凭据；只能在一次性的隔离实验室中评估。

### 预期接口

SQL 接口包括 `mdbr_add_shard`、`mdbr_reset_meta`、`create_sharding_key`、`assign_key_range_2_shard`、`show_shkey` 和 `add_local_table`。下面的最小实验配置说明其预期模型：

```sql
CREATE EXTENSION router;

SELECT mdbr_add_shard(
  'shard1', '127.0.0.1', '5432', 'sharddb', 'router_user', 'lab-only-password'
);
SELECT create_sharding_key('customer_key', 'customer_id');
SELECT assign_key_range_2_shard('shard1', 'customer_key', 1, 100000);
SELECT show_shkey('customer_key');
```

上游演示采用面向预加载的启动方式和 make 驱动的集群工具；仅创建 SQL 对象并不能形成受支持的运维流程。元数据保存在共享结构中，没有持久一致性，也不会在重启后恢复。

### 关键安全边界

分片主机、用户与明文密码通过无界字符串操作复制到固定大小的共享内存缓冲区。超长输入可能破坏内存，调试日志还可能泄露包括密码在内的连接参数。绝不能提供真实凭据，也不能向不可信调用者开放元数据函数。

代码对分片、分片键、列与键范围都有很小的固定上限，同时含有未完成且从特定版本复制的规划器路径。README 自身也列出了糟糕的内存管理和缺少元数据一致性。不能推断它具备事务、故障转移、重启恢复、模式变更或路由正确性保证。`router` 只能用于检查历史设计思路，绝不能作为生产分片层。
