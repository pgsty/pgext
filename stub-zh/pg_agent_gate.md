## 用法

来源：

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/af4ace3098e5b93d5f0f523c02f67523e424a7b7/sql/pg_agent_gate--0.2.14--0.2.15.sql)

`pg_agent_gate` 0.2.15 在 PostgreSQL 18/19 上通过提议、检查、试运行和受控执行处理已注册 Agent 角色。必须在第一条 Agent 语句之前加载动态库，可采用共享预加载，或配置角色会话预加载后重新连接。

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

### 0.2.15 安全修复

0.2.15 拒绝 gate 参数外围可能在验证前执行 Agent 所选函数的显式转换和 I/O 转换，并拒绝单独的分号，避免后端崩溃。这些修复位于动态库中，SQL 升级脚本为空操作。应替换动态库，重启或重连所有预加载它的进程，再逐库升级到 0.2.15。DDL 权限属于模式修改能力，行数限制无法约束它。
