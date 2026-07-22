## 用法

来源：

- [Pinned official README](https://github.com/nuko-yokohama/pg_sulog/blob/e07a87f76261ccfd884da42bff84419b395cba81/README.md)
- [Pinned hook implementation](https://github.com/nuko-yokohama/pg_sulog/blob/e07a87f76261ccfd884da42bff84419b395cba81/pg_sulog.c)

`pg_sulog` 是只能预加载的钩子模块，可把超级用户语句写入 PostgreSQL 服务端日志，并尝试阻止它们。它不创建 SQL 对象，因此没有面向应用的 `CREATE EXTENSION` 流程；共享库在 postmaster 启动时预加载后，相关行为便会生效。

### 配置

安装匹配的共享库，将其加入 `postgresql.conf`，选择模式，然后重启 PostgreSQL：

```ini
shared_preload_libraries = 'pg_sulog'
pg_sulog.mode = 'LOGGING'
```

文档列出的模式是：

- `LOGGING`：记录超级用户发出的每条语句并允许执行；这是默认值。
- `MAINTENANCE`：允许并记录 `VACUUM`、`REINDEX`、`ANALYZE`、`CLUSTER`，把其他超级用户语句视为被阻止。
- `BLOCK`：把全部超级用户语句视为被阻止。

消息以 PostgreSQL 警告形式发出，包含时间戳、`logging` 或 `blocked`、登录角色和原始查询文本。服务端日志的保留与访问控制必须考虑语句中的敏感字面量。

### 运维边界

实现挂接执行器和工具命令处理流程，并在内存中构建超级用户角色名称列表。列表最多 64 项。它从连接中识别登录角色，不提供 SQL 审计表、结构化事件模式、按角色策略、脱敏或防篡改持久化。

这个已弃用模块在上游只验证过 CentOS 7 与 PostgreSQL 9.3–9.5，并提及 9.6 开发版。此后 PostgreSQL 的钩子签名已经变化。更重要的是，固定版本的阻止代码包含不一致的模式比较和被注释掉的执行路径；若未在完全一致的服务端构建上完成源码审查和破坏性隔离测试，不应把 `MAINTENANCE` 或 `BLOCK` 当作安全控制。

只有在验证兼容性和日志量后才使用 `LOGGING`。预加载失败可能导致 PostgreSQL 无法启动，因此应保留离线修改配置的手段。
