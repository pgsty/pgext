## Usage

Sources:

- [pg_policy 0.1.0 on PGXN](https://pgxn.org/dist/pg_policy/0.1.0/)
- [pg_policy 0.1.0 README](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/README.md)
- [Agent Policy Language reference](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/doc/language.md)
- [pg_policy 0.1.0 security policy](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/SECURITY.md)
- [pg_policy 0.1.0 control file](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/pg_policy.control)
- [pg_policy 0.1.0 extension SQL](https://api.pgxn.org/src/pg_policy/pg_policy-0.1.0/sql/pg_policy--0.1.0.sql)
- [Pigsty pg_policy package page](https://pgext.cloud/ext/pg_policy)

`pg_policy` 0.1.0 is an experimental SQL and PL/pgSQL policy evaluator for agent and tool actions. It stores Agent Policy Language rules, evaluates context and session history, records every decision, and returns obligations for a gateway to enforce. It complements PostgreSQL roles and row-level security; it does not intercept SQL or tool calls by itself.

### Pigsty Schema Compatibility

Upstream 0.1.0 declares the reserved schema name `pg_policy` and defines an unquoted function named `check`. Pigsty packages patch the installed schema to `policy`, quote the reserved function name as `policy."check"()`, and fix function search paths. The upstream examples therefore cannot be copied verbatim into a Pigsty installation.

```sql
CREATE EXTENSION pg_policy;

SELECT policy.set_setting('enforcement_mode', 'log_only');
```

The extension is not relocatable, requires PostgreSQL 14 or later, and does not require `shared_preload_libraries` or a PostgreSQL restart. Current Pigsty packages cover PostgreSQL 14–18.

### Define and Evaluate a Guardrail

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

`policy.evaluate(...)` returns JSON containing `decision`, `allowed`, `matched_policies`, `obligations`, `reasons`, and `mode`. The convenience wrapper `policy."check"()` returns only a boolean. `policy.enforce()` requests exception-on-deny behavior when the mode is `enforce`.

### APL Surface

An APL document begins with one effect: `permit`, `forbid`, or `guide`. It can match principal, action, and resource types and identifiers. In 0.1.0, context conditions support only `==`, `in [...]`, and `and`. A temporal clause can count matching session events inside an interval when evaluation receives a session identifier.

`forbid` overrides matching `permit` rules. `guide` allows the action and can return `advice`, `prefer_tool`, or `max_rows` obligations. The caller—not the extension—must interpret and apply those obligations.

### Sessions, Temporal Limits, and Audit

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

`policy.open_session()` creates or updates a session. Evaluations with a session identifier append an event and can satisfy temporal predicates. Every evaluation writes `policy.decision_log`; other important relations are `policy.policies`, `policy.sessions`, `policy.events`, and `policy.settings`.

### Enforcement and Security Boundaries

- The default `enforcement_mode` is `log_only` and the default decision is `permit`. A matched deny becomes an allow with a `shadow_deny` obligation.
- In `guide` mode, a matched deny becomes an allow with `would_deny`. Only `enforce` preserves a deny and allows `policy.enforce()` to raise an error.
- A gateway must call the evaluator before the protected action and hard-fail on deny. Calling `policy.evaluate(...)` after executing a tool is only auditing.
- Keep PostgreSQL `GRANT` and `REVOKE`, row-level security, network controls, and least-privilege credentials as the authoritative data-plane controls. Superusers and roles with `BYPASSRLS` can bypass row-level controls.
- The 0.1 line is explicitly an experimental MVP, not a hardened production security boundary. Shadow-test policies, restrict who can change `policy.settings` or `policy.policies`, and monitor `policy.decision_log` before switching to `enforce`.
