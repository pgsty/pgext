## 用法

来源：

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.3.0/sql/pg_agent_gate--0.2.15--0.3.0.sql)

`pg_agent_gate` 0.3.0 在 PostgreSQL 18/19 上通过提议、检查、试运行和受控执行处理已注册 Agent 角色，并可要求提交前经过人工审批。必须在第一条 Agent 语句之前加载动态库，可采用共享预加载，或配置角色会话预加载后重新连接。

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

### 人工审批

调用 `agent_gate.register_agent` 时设置 `p_requires_approval`，即可要求该 Agent 的提议经过审批。`agent_gate_approver` 的成员通过 `agent_gate.approve` 记录同意后，仍由 Agent 提交自己的提议。审批本身不执行 SQL，也不绕过行数限制、断言或提交时检查。Agent 不能自行审批，`propose_and_commit` 也不能绕过审批要求。

固定模式 `agent_gate_internal` 保存提议、执行和审批记录，Agent 不直接访问这些表。执行记录的 `approved_by` 字段保存审批者身份。

### 动态库与 SQL 升级

动态库保留了 0.2.15 对可执行函数的参数转换和单分号崩溃的防护。应替换动态库，重启或重连所有预加载它的进程，再通过上游 SQL 升级脚本逐库升级到 0.3.0。0.3.0 迁移新增审批对象，不是空操作。DDL 权限属于模式修改能力，行数限制无法约束它。
