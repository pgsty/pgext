## 用法

来源：

- [官方扩展控制文件](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/pg_graphql_server.control)
- [官方后台 worker 实现](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/src/lib.rs)
- [官方 HTTP server 实现](https://github.com/marcelmeulemans/pg_graphql_server/blob/73a00a3b7dbda0036eefee4903f0b3d4710b0335/src/server.rs)

`pg_graphql_server` 0.0.1 是实验性的 pgrx 后台 worker，在 PostgreSQL postmaster 内启动 HTTP endpoint，并把 GraphQL 查询字符串交给独立的 `pg_graphql` 扩展。固定源码没有认证或 TLS 层，未经大量加固不应作为生产 API 暴露。

### 核心流程

安装库，把 `pg_graphql_server` 加入 `shared_preload_libraries`，然后重启 PostgreSQL。worker 会连接 `postgres` 数据库并在其中创建管理对象。安装两个扩展后添加 listener 配置：

```sql
SHOW shared_preload_libraries;

CREATE EXTENSION pg_graphql;
CREATE EXTENSION pg_graphql_server;

INSERT INTO http_server.servers (
  listen_port,
  postgres_user,
  postgres_pass,
  database_name
)
VALUES (8080, 'graphql_app', 'change-this-password', 'appdb');
```

事务提交后，表触发器会请求配置 reload。worker 启动 listener，并使用已存凭据回连本机 PostgreSQL 端口。它暴露 `GET /health`，以及 `GET /graphql` 和 `POST /graphql`；GraphQL 请求在配置的目标数据库中通过 `graphql.resolve($1)` 求值。

### 重要对象

- `http_server.servers` 按 `listen_port` 保存 listener，以及 PostgreSQL 用户名、密码和目标数据库。
- `http_server.on_servers_changed()` 及其 statement trigger 在配置变化后调用 `pg_reload_conf()`。
- `/health` 报告连接池大小和空闲连接数。
- `/graphql` 接收 GraphQL 请求，并返回 `graphql.resolve($1)` 产生的 JSON 值。

### 预加载与兼容性

worker 由 `_PG_init` 注册，在 recovery 完成后启动，并且需要 SPI 访问，因此必须预加载并重启服务器。Cargo manifest 为 PostgreSQL 14、15 和 16 提供构建 feature。每个配置的目标数据库都必须能够安装且已经安装 `pg_graphql`，否则 listener 会失败并重试。

### 安全与运维说明

固定实现把每个 HTTP server 绑定到全部 IPv4 接口，没有认证、授权、TLS、origin 过滤、请求大小限制或查询 allowlist，并把数据库密码以普通表值保存。它还只从名为 `postgres` 的数据库启动，并通过 `::1` 回连。评估该代码时，应限制网络可达性、使用最小权限数据库角色、保护管理表，并在前端放置加固代理。应把 0.0.1 视为原型，任何非实验环境部署前都要审计精确源码。
