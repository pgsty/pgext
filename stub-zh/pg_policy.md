## 用法

来源：

- [PGXN 上的 pg_policy 0.1.0](https://pgxn.org/dist/pg_policy/0.1.0/)
- [pg_policy 0.1.0 README](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/README.md)
- [Agent Policy Language 参考](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/doc/language.md)
- [pg_policy 0.1.0 安全策略](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/SECURITY.md)
- [pg_policy 0.1.0 控制文件](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/pg_policy.control)
- [pg_policy 0.1.0 扩展 SQL](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/sql/pg_policy--0.1.0.sql)
- [Pigsty pg_policy 软件包页面](https://pgext.cloud/ext/pg_policy)

`pg_policy` 0.1.0 是一个实验性的 SQL 与 PL/pgSQL 策略求值器，用于代理和工具动作。它存储 Agent Policy Language 规则，依据上下文和会话历史求值，记录每次决策，并返回供网关执行的义务。它用于补充 PostgreSQL 角色与行级安全，而不会自行拦截 SQL 或工具调用。

### Pigsty 模式兼容性

上游 0.1.0 声明了保留模式名 `pg_policy`，并定义了名为 `check` 的未加引号函数。Pigsty 软件包把安装模式修补为 `policy`，将保留函数名加引号为 `policy."check"()`，并固定函数搜索路径。因此，上游示例不能原样复制到 Pigsty 安装中。

```sql
CREATE EXTENSION pg_policy;

SELECT policy.set_setting('enforcement_mode', 'log_only');
```

该扩展不可重定位，要求 PostgreSQL 14 或以上版本，不需要 `shared_preload_libraries`，也无需重启 PostgreSQL。当前 Pigsty 软件包覆盖 PostgreSQL 14–18。

### 定义并求值一条护栏

```sql
SELECT policy.upsert_policy('block_ddl', $apl$
forbid
  principal agent "research_bot"
  action tool "execute_sql"
  when { context.statement_type in ["DROP", "TRUNCATE", "ALTER", "CREATE"] }
  reason "Research agents may not run DDL"
$apl$);

SELECT policy.set_setting('enforcement_mode', 'enforce');

SELECT policy.evaluate(
  'agent', 'research_bot',
  'tool', 'execute_sql',
  '*', '*',
  '{"statement_type":"DROP"}'::jsonb,
  NULL
);

SELECT policy."check"(
  'research_bot',
  'execute_sql',
  '{"statement_type":"DROP"}'::jsonb
);
```

`policy.evaluate(...)` 返回包含 `decision`、`allowed`、`matched_policies`、`obligations`、`reasons` 与 `mode` 的 JSON。便捷封装 `policy."check"()` 只返回布尔值。`policy.enforce()` 会在模式为 `enforce` 时请求遇到拒绝即抛出异常。

### APL 能力边界

APL 文档以 `permit`、`forbid` 或 `guide` 三种效果之一开头，可以匹配主体、动作和资源的类型与标识符。在 0.1.0 中，上下文条件只支持 `==`、`in [...]` 与 `and`。当求值时传入会话标识符，时间子句可以统计给定时间间隔内匹配的会话事件。

匹配的 `forbid` 会覆盖 `permit`。`guide` 允许动作，并可返回 `advice`、`prefer_tool` 或 `max_rows` 义务。这些义务必须由调用方解释和执行，而不是由扩展自动处理。

### 会话、时间限制与审计

```sql
SELECT policy.open_session(
  'sess-1',
  'agent',
  'research_bot'
);

SELECT policy.upsert_policy('export_budget', $apl$
forbid
  principal agent "research_bot"
  action tool "export_csv"
  when temporal {
    count(action == "export_csv") within interval '1 hour' >= 3
  }
  reason "Export budget exceeded"
$apl$);

SELECT policy.evaluate(
  'agent', 'research_bot',
  'tool', 'export_csv',
  '*', '*',
  '{}'::jsonb,
  'sess-1'
);
```

`policy.open_session()` 创建或更新会话。带会话标识符的求值会追加事件，并可满足时间谓词。每次求值都会写入 `policy.decision_log`；其他重要关系包括 `policy.policies`、`policy.sessions`、`policy.events` 与 `policy.settings`。

### 执行与安全边界

- 默认 `enforcement_mode` 是 `log_only`，默认决策是 `permit`。匹配的拒绝会变成允许，并附加 `shadow_deny` 义务。
- 在 `guide` 模式下，匹配的拒绝会变成允许，并附加 `would_deny`。只有 `enforce` 会保留拒绝，并允许 `policy.enforce()` 抛出错误。
- 网关必须在受保护动作之前调用求值器，并在拒绝时硬失败。工具执行后才调用 `policy.evaluate(...)` 只能提供审计。
- 应继续把 PostgreSQL `GRANT` 与 `REVOKE`、行级安全、网络控制和最小权限凭证作为权威的数据面控制。超级用户以及带有 `BYPASSRLS` 的角色可以绕过行级控制。
- 0.1 系列明确是实验性 MVP，而不是加固过的生产安全边界。切换到 `enforce` 前，应影子测试策略、限制能修改 `policy.settings` 或 `policy.policies` 的角色，并监控 `policy.decision_log`。
