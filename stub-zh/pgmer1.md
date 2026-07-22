## 用法

来源：

- [官方 README](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/README.md)
- [扩展实现](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/src/lib.rs)
- [Cargo 与 PostgreSQL 兼容性](https://github.com/shestero/pgmer1/blob/4339292ae4d6807de1d04b15cec2b83e71d53ca0/Cargo.toml)

`pgmer1` 0.1.0 是面向独立 MeritRank 服务的 pgrx HTTP 客户端，可从 SQL 提交带权图边并查询排名分数。端点硬编码为 `http://localhost:8000`；没有 GUC、逐调用 URL、认证或 TLS 配置。

### 核心流程

先在 PostgreSQL 服务器的回环接口上运行兼容 MeritRank 服务，然后执行：

```sql
CREATE EXTENSION pgmer1;

SELECT mr_service_url();

SELECT mr_edge('alice', 'bob', 1.0);

SELECT * FROM mr_scores('alice');

SELECT mr_node_score('alice', 'bob');
```

`mr_edge(src, dest, weight)` 发送 HTTP `PUT /edge` JSON 请求，并返回响应中的 message 字段。`mr_scores(ego)` 发送 `GET /scores/{ego}`，返回 `(node text, ego text, score double precision)` 行。`mr_node_score(ego, target)` 调用 `GET /node_score/{ego}/{target}` 并返回单个分数。`mr_service_url()` 显示固定端点。

### 事务与故障边界

调用会在 PostgreSQL 后端中同步执行。HTTP 服务缓慢或不可用时，SQL 会话也会被占用，而且实现没有配置请求超时。GET 路径会对传输初始化错误使用 unwrap，因此连接故障可能变成 Rust panic，而不是可控 SQL 错误。

`mr_edge` 会立即修改外部服务，不能参与 PostgreSQL 回滚：若超时结果不明确，重试可能重复变更，除非服务自身保证操作幂等。该扩展会把图标识符直接放入 URL 路径，没有显式百分号编码。只能接受可信标识符，限制执行权限，并将回环服务与不可信本地进程隔离。

固定版本的 Cargo 功能覆盖 PostgreSQL 11–15，并使用较旧 pgrx；目录状态为已放弃。应针对精确服务器和 MeritRank API 修订重新构建测试；需要认证、可观测性、超时控制或可靠投递时，应优先在应用侧集成。
