## 用法

来源：

- [README.md](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/README.md)
- [Cargo.toml](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/Cargo.toml)
- [src/lib.rs](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/src/lib.rs)
- [sql/pg_agent_gate--0.2.0--0.2.1.sql](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/sql/pg_agent_gate--0.2.0--0.2.1.sql)
- [pg_agent_gate.control](https://github.com/Manuelreyesbravo/pg_agent_gate/blob/2b06b19d0ed527e0e8c89aa15796c92964671e3d/pg_agent_gate.control)

`pg_agent_gate` 在 PostgreSQL 18+ 上通过提议、试运行和受控执行函数处理已注册 Agent 角色的请求。该角色执行第一条语句前必须加载动态库。

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

管理员将库合并到 `shared_preload_libraries` 并重启，或由 `agent_gate.register_agent` 为角色设置 `session_preload_libraries`，随后重新连接。安装需要超级用户。应注册普通角色，只授予必要的表权限，并设置影响行数上限与 DDL 策略。Agent 接口包括 `agent_gate.discover`、`agent_gate.propose`、`agent_gate.dry_run`、`agent_gate.commit`、`agent_gate.acts` 和 `agent_gate.whoami`；可选断言绑定使用 `pg_living_assertions`。试运行可以回滚数据库修改，但无法撤销序列推进、会话锁及外部副作用。已有特权函数、视图所有者权限和快速路径函数调用仍是安全边界，不能将其视为完整沙箱。默认 `agent_gate.attempt_durability = fast` 可能在崩溃时丢失尝试记录，外层事务回滚也会撤销记录。0.2.1 修复了写入 CTE 绕过行数限制和直接修改参数的漏洞。
