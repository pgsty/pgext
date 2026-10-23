## 用法

来源：

- [官方资料](https://github.com/MasahikoSawada/pg_raft/blob/320d2cb89a5a971ed2b3f9430b16a5ed20344da1/README.md)
- [官方资料](https://github.com/MasahikoSawada/pg_raft/blob/320d2cb89a5a971ed2b3f9430b16a5ed20344da1/pg_raft.control)
- [官方资料](https://github.com/MasahikoSawada/pg_raft/blob/320d2cb89a5a971ed2b3f9430b16a5ed20344da1/pg_raft--1.0.sql)

`pg_raft` 是尚未完成的 Raft 实验项目，SQL 版本为 1.0，本次核对源码修订 320d2cb8。它提供节点管理与选主机制；上游待办清单中的日志复制、加入或离开组以及启动器仍未完成。

### 实验性配置

上游流程要求预加载库、指定工作数据库并重启服务器：

```ini
shared_preload_libraries = 'pg_raft'
pg_raft.database = 'postgres'
```

随后超级用户可在该数据库中执行原型的节点创建示例：

```sql
CREATE EXTENSION pg_raft;
SELECT pgraft.create_node('node-1', 'port=5432 dbname=postgres');
```

应在现有预加载列表中加入该库并保留其他库。这只是实验流程，不是搭建完整生产复制集群的操作手册。

### 对象与边界

`pgraft` 模式包含节点、组、状态和日志对象，以及内部投票和追加日志 RPC 函数。这些接口属于原型实现细节；存在 RPC 函数不代表已经具备完整共识或持久化复制日志能力。

所核对的仓库没有明确支持的 PostgreSQL 大版本范围，也没有明确许可证。采用前应向上游确认这两点。请使用隔离的测试集群，保护连接信息，不要以此原型承诺可用性或数据持久性。
