## 用法

来源：

- [README](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/README.md)
- [Control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/pg_agent_gate.control)
- [CHANGELOG](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/CHANGELOG.md)
- [SQL upgrade](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/v0.4.3/sql/pg_agent_gate--0.4.1--0.4.3.sql)

`pg_agent_gate` 0.4.3 在 PostgreSQL 18/19 上通过提议、检查、试运行和受控执行处理已注册 Agent 角色，并可要求提交前经过人工审批。必须在第一条 Agent 语句之前加载动态库，可采用共享预加载，或配置角色会话预加载后重新连接。

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

### 写入预算与角色成员

0.4.0 引入默认累计写入预算：每小时允许修改该 Agent 单条语句行数上限所规定的行数。`agent_gate.set_budget` 调整额度和窗口；`agent_gate.reset_budget` 重新开始计算额度，不改写记录。已提交写入按实际行影响计数，包括级联与触发器；并发提交会串行核对预算。

Agent 角色不能有成员。登记函数会拒绝这样的角色；升级时会对已有成员关系发出警告，管理员必须撤销这些关系。每个 Agent 应使用自己的登录角色连接，连接代理不能通过角色切换进入 Agent 身份。

### 0.4.3 的函数与会话检查

0.4.3 动态库在隐藏函数调用实际执行时进行检查，包括类型支持函数、聚合、默认值和约束。它拒绝传给已知执行函数的 SQL 文本、Agent 可修改的函数、受保护上下文变化以及未计入控制范围的状态变化。DDL 引起的行变化也计入行数限制。这些检查会有意拒绝普通 PostgreSQL 角色原本能够执行的部分语句。

受保护上下文包括角色身份、行级安全设置、搜索路径和应用占位参数。上下文变化会导致拒绝，而不是自动恢复原值。Agent 所有对象附带的原生代码受到限制。可选的 `pg_living_assertions` 0.5.5 可能泄漏上下文并导致拒绝，0.5.6 已修复该问题。不要将 Agent 可控制的模式放入其他用户的搜索路径，因为 DDL 可以创建之后以其他用户权限运行的代码。

### 动态库与 SQL 升级

替换动态库后，重启或重新连接所有预加载它的进程，再逐库升级：

```sql
ALTER EXTENSION pg_agent_gate UPDATE TO '0.4.3';
```

0.4.1 至 0.4.3 迁移不改变模式对象。更早版本的升级路径仍会添加预算和审批对象；升级后应检查已有 Agent 的预算。0.4.2 未单独发布。超级用户、操作系统和可信 DBA 代码不在这个保护边界内；获准执行的语句仍可能违背人的意图。
