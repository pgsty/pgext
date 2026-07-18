## 用法

来源：

- [已复核 commit 的连接器 README](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/README.md)
- [SQL 函数实现](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/src/lib.rs)
- [RPC 客户端实现](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/src/rpc.rs)
- [Cargo 元数据](https://github.com/Intersubjective/meritrank-rust/blob/d1a2ada875e6c1d546fd07fdd2d61e9ba2164081/psql-connector/Cargo.toml)

`pgmer2` 将 SQL 会话连接到单独运行的 MeritRank 图排名服务。在后端启动前，应在 PostgreSQL 服务器环境中设置 `MERITRANK_SERVICE_URL`；默认值为 `tcp://127.0.0.1:8080`。

```sql
CREATE EXTENSION pgmer2;
SELECT mr_service_url();
SELECT * FROM mr_put_edge('alice', 'bob', 1.0, '', -1);
SELECT * FROM mr_node_score('alice', 'bob', '');
```

API 可管理图上下文与边、计算分数、列出节点和邻居，并通过 `mr_bulk_load_edges` 批量加载。服务属于外部状态，必须独立于 PostgreSQL 进行部署、保护、监控与备份。

当前传输是未加密 TCP，使用自定义二进制协议，连接器也未展示认证。多个执行网络调用的读取函数被声明为 `IMMUTABLE`，即使远端状态变化，PostgreSQL 仍可能复用或折叠结果。应限制执行权、隔离服务网络，且不能依赖普通 SQL 易变性或事务语义来约束远端变更。
