## 用法

来源：

- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_shard/README.md)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_shard/kingres_shard.control)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/contrib/kingres_shard/kingres_shard--0.5.sql)
- [官方来源](https://github.com/KingresDB/Kingres/blob/3c98d1cd4a8d81e2fee4d859da0951b65ac44e72/kingres/README.md)

`kingres_shard` 0.5 为 KingresDB 提供哈希桶分布、参考表、语句路由与两阶段提交。它要求 KingresDB 内核标记，不能针对原版 PostgreSQL 构建。源码支持 PostgreSQL 18 与 19 分支，后者基于候选发布版。

### 配置集群

每个节点都须预加载并重启。所有节点使用同一 PostgreSQL 主版本及兼容扩展文件，具备 libpq 连通性和匹配的角色、数据库。准备事务上限必须覆盖并发分布式事务。管理 API 要求超级用户权限。

```ini
shared_preload_libraries = 'kingres_shard'
max_prepared_transactions = 100
```

```sql
CREATE EXTENSION kingres_shard;
SELECT kingres.init_cluster('host1', 5432);
SELECT kingres.add_node('host2', 5432);
CREATE TABLE public.users (id int PRIMARY KEY, name text);
SELECT kingres.create_distributed_table('public.users', 'id');
SELECT * FROM kingres.check_cluster();
```

### 路由与维护

`kingres.create_reference_table` 将小表复制到各节点，`kingres.move_bucket` 与 `kingres.rebalance_plan` 支持位置调整，`kingres.recover_prepared_transactions` 处理已记录的提交决策。DDL 经由映射组主节点协调。共享 `kingres` schema 可与家族中的其他扩展共存。依赖新辅助函数前，应在所有节点升级库、重启并应用提供的 SQL 升级链。

### 正确性边界

两阶段提交不提供全局快照：分布式读取可能看到不同快照，同一语句的不同扫描也可能如此。不支持的多行写入会被拒绝。不支持跨 PostgreSQL 主版本滚动升级。所审查修订仍披露高严重性故障转移缺陷，控制平面的自动故障转移尚未完成。不要依赖其故障转移持久性，也不要将此预览视为生产就绪的分布式数据库。
