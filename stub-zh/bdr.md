## 用法

来源：

- [PGD 官方文档](https://www.enterprisedb.com/docs/pgd/6.5/)
- [官方节点创建指南](https://www.enterprisedb.com/docs/pgd/6.5/node_management/creating_nodes/)
- [官方节点管理 API](https://www.enterprisedb.com/docs/pgd/6.5/reference/nodes-management-interfaces/)

`bdr` 是 EDB Postgres Distributed (PGD) 的核心 provider extension，提供 active-active logical replication、node group、conflict management、consensus 与 distributed commit policy。

### 安装与预加载

在每个节点安装匹配的 PGD 6.5 软件包，启用 commit timestamp，预加载 `bdr`，重启全部节点后再创建扩展：

```ini
shared_preload_libraries = 'bdr'
track_commit_timestamp = on
```

```sql
CREATE EXTENSION bdr;
SELECT bdr.bdr_version();
```

只有超级用户可以创建扩展。其他 PGD 管理操作应使用预定义 `bdr_` 角色。

### 初始化并加入 Group

```sql
SELECT bdr.create_node(
  node_name := 'node_a',
  dsn       := 'host=node-a dbname=app'
);

SELECT bdr.create_node_group(
  node_group_name := 'app_group'
);
```

在另一个已准备节点创建 `bdr`，调用 `bdr.create_node` 并通过 `bdr.join_node_group` 加入。导入流量前使用 `bdr.wait_for_join_completion`。

### 分布式系统边界

Group create、join、part 与 consensus 操作是异步的，不会随外围 SQL 事务回滚。DDL replication 到达前，每个节点都必须具备兼容的 PGD/PostgreSQL binary 与 extension package。应规划 replication slot/origin、conflict policy、commit scope、sequence behavior、DDL、role propagation、backup/restore、node loss 与 network partition；`CREATE EXTENSION` 成功只证明本地安装，不证明集群健康。
