## 用法

来源：

- [0.1.0 版本官方 README](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/README.md)
- [官方扩展 SQL](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/sql/pgclaw--0.1.0.sql)
- [官方后台工作进程实现](https://github.com/calebwin/pgclaw/blob/ea6c3abe27602724f01de199b87eac8cfcd47212/src/worker.rs)

`pgclaw` 使用 `claw` 数据类型保存 LLM 智能体定义，并由 PostgreSQL 后台工作进程异步处理插入或更新的行。它适合模型回写行字段的实验性智能体工作流，也支持可复用智能体、渠道会话和可选的 Claude Code 工作区。

### 配置与核心流程

后台工作进程需要在预加载阶段配置并重启：

```conf
shared_preload_libraries = 'pgclaw'
pgclaw.database = 'appdb'
pgclaw.api_provider = 'anthropic'
pgclaw.api_key = 'replace-with-a-secret'
```

然后创建扩展，建立具有主键和 `claw` 列的表，并注册监控：

```sql
CREATE EXTENSION pgclaw;

CREATE TABLE tickets (
    id bigserial PRIMARY KEY,
    title text,
    priority text,
    agent claw DEFAULT claw(
        'Set priority to low, medium, high, or critical from the ticket data.'
    )
);

SELECT claw_watch('tickets'::regclass);
INSERT INTO tickets(title) VALUES ('Production login returns HTTP 500');

SELECT id, priority FROM tickets;
SELECT id, error, created_at, done_at FROM claw.queue ORDER BY id DESC;
```

触发器把行 JSON 写入 `claw.queue`；工作进程调用已配置服务商、解析响应并更新源行。`claw_unwatch(regclass)` 删除触发器。`claw_model()`、`claw_prompt()` 与 `claw_agent_id()` 检查存储值。可复用定义位于 `claw.agents`；会话、消息、收件箱/发件箱、渠道绑定、心跳与定时任务分别保存在对应 `claw` 表中。

### 运维与安全边界

`claw_watch()` 要求表具有主键及至少一个 `claw` 列。处理是异步的，可能失败；应监控 `claw.queue.error` 并设计幂等重试。行内容会发送到外部 LLM 端点，因此应最小化暴露字段并取得适当的数据处理授权。API 密钥 GUC 仅允许超级用户设置，但仍驻留在服务器配置和进程内存中。

具有 `workspace` 的智能体可以创建目录并调用 Claude Code，后者能以 PostgreSQL 服务账户读写文件和运行工具。应将此模式视为远程代码执行，并与承载重要数据的数据库主机隔离。0.1.0 仍属早期设计；依赖前应测试服务商响应、权限、触发器递归防护、速率限制、事务可见性与崩溃恢复。
