## 用法

来源：

- [上游 README](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/README.md)
- [后台工作进程实现](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/src/pg_web.c)
- [HTTP 处理器](https://github.com/le0pard/pg_web/blob/6e936f451090482bdab8d6f749ccccd4c1cbbc60/src/pg_web_handler.c)

`pg_web` 版本 `0.1.0` 是一个 2013 年的演示项目，在 PostgreSQL 后台工作进程中运行极小的嵌入式 HTTP 服务器。尽管仓库描述如此，它并不提供数据库查询接口，也不应作为 Web 管理工具部署。

### 核心流程

库加载时会注册后台工作进程。将其配置进 `shared_preload_libraries`，选择在 postmaster 启动阶段生效的端口，然后重启 PostgreSQL。

```conf
shared_preload_libraries = 'pg_web'
pg_web.port = 8080
```

```sql
CREATE EXTENSION pg_web;
```

版本化 SQL 脚本不会创建任何 SQL 对象。恢复结束后，工作进程连接到硬编码的 `postgres` 数据库，并为 `/`、`/date`、`/count` 和 `/ip` 提供 GET 响应。计数器只属于当前进程；IP 端点返回对端地址。

### 安全与注意事项

实现没有记录绑定地址设置、认证、授权、TLS、请求限制或数据库端点。未知路径仍会收到 HTTP 200，并在正文中包含错误字符串。除非主机防火墙能证明并非如此，应把所配置的端口视为可从远端访问。

这个老旧原型使用历史后台工作进程 API 与内嵌网络库。绝不能将其暴露到不可信网络或用于生产。若仅在隔离实验室评估，应验证端口绑定、进程重启行为、关闭流程、畸形请求与目标 PostgreSQL 构建兼容性。
