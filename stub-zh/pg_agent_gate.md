## 用法

来源：

- [README.md](https://api.pgxn.org/src/pg_agent_gate/pg_agent_gate-0.2.14/README.md)
- [pg_agent_gate.control](https://api.pgxn.org/src/pg_agent_gate/pg_agent_gate-0.2.14/pg_agent_gate.control)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_agent_gate/pg_agent_gate-0.2.14/CHANGELOG.md)

`pg_agent_gate` 0.2.14 在 PostgreSQL 18/19 上通过提议、检查、试运行和受控执行处理已注册 Agent 角色。必须在第一条 Agent 语句之前加载动态库，可采用共享预加载，或配置角色会话预加载后重新连接。

### 核心用法

```sql
CREATE EXTENSION pg_agent_gate;
```

```sql
SELECT agent_gate.whoami();
SELECT agent_gate.discover('customers');
SELECT agent_gate.propose('SELECT 1', 'Check connectivity');
```

### 运行边界

由超级用户安装，为普通角色登记有限的表权限、行数限制与 DDL 允许列表。API 包括 `agent_gate.whoami`、`agent_gate.discover`、`agent_gate.propose`、`agent_gate.dry_run`、`agent_gate.commit` 和 `agent_gate.acts`。这些操作要求 READ COMMITTED，试运行与提交前会再次核对提议时的会话设置，调用已知状态修改内置函数的读取请求会被拒绝。重新登记同名角色不会继承旧提议。`agent_gate.max_proposal_bytes` 与 `agent_gate.max_intent_bytes` 限制输入大小。已有重复提交的提议可能阻止升级，需要检查处理。可选绑定检查使用 `pg_living_assertions`。回滚无法撤销所有序列或外部副作用，已有特权代码仍须可信。服务器日志可能包含 SQL 字面量；尝试记录受外层事务影响，默认快速持久化模式可能在崩溃时丢失近期记录。
