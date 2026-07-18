## Usage

Sources:

- [Pinned upstream README](https://github.com/fabriziomello/session_variables/blob/1022a217962bea72e9313478d1e0c27c1b07a3a4/README.md)
- [Pinned installation SQL](https://github.com/fabriziomello/session_variables/blob/1022a217962bea72e9313478d1e0c27c1b07a3a4/sql/session_variables.sql)
- [Pinned distribution metadata](https://github.com/fabriziomello/session_variables/blob/1022a217962bea72e9313478d1e0c27c1b07a3a4/META.json)
- [Formal PGXN documentation](https://pgxn.org/dist/session_variables/0.0.4/doc/session_variables.html)

session_variables version 0.0.4 is a two-function SQL wrapper around PostgreSQL custom settings. set_value writes a text value under the session_variables namespace, and get_value reads it back.

### Example

```sql
CREATE EXTENSION session_variables;

SELECT set_value('request_id', 'req-123');
SELECT get_value('request_id');

BEGIN;
SELECT set_value('request_id', 'req-456');
ROLLBACK;

SELECT get_value('request_id');
```

The final read still returns req-456: set_value calls set_config with transaction-local mode disabled, so the setting lasts for the entire PostgreSQL session and is not rolled back with the transaction. Values are text. Reading a name that has never been set raises an error because get_value does not request missing-ok behavior.

### Pooling and security limits

The state belongs to a backend connection, not an application user or logical request. With session or transaction pooling, a later client can inherit a value left by an earlier client unless every checkout resets it. The extension has no list, unset, type, expiry, ownership, privilege, or isolation model. Its functions are executable by PUBLIC unless privileges are changed.

Do not store authentication decisions, tenant identity, secrets, or authorization context in session_variables. Prefer transaction-local set_config with disciplined reset logic or native application context designed for the pooler mode. If legacy code depends on it, revoke public execution, expose validated wrappers, initialize every variable on every checkout, and clear the backend before reuse.

The repository last changed in 2012, describes very old PostgreSQL installation paths, and is classified as abandoned. The SQL is simple, but compatibility and operational semantics still require tests against the target pooler and PostgreSQL version.
