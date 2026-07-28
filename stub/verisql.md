## Usage

Sources:

- [Official upstream README](https://github.com/sneha21032004/verisql/blob/c13e6c804012b02af45e49aea33c9d8ab6526180/postgres-extension/README.md)
- [Official extension control file (verisql.control)](https://github.com/sneha21032004/verisql/blob/c13e6c804012b02af45e49aea33c9d8ab6526180/postgres-extension/verisql.control)
- [Official extension SQL (verisql--0.1.0.sql)](https://github.com/sneha21032004/verisql/blob/c13e6c804012b02af45e49aea33c9d8ab6526180/postgres-extension/verisql--0.1.0.sql)

`verisql` — Deterministic verification oracle for AI-generated SQL, **inside the database** — no extra service, no client library required. Pure PL/pgSQL, runs on any PostgreSQL 13+, distributed as a CREATE EXTENSION install. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION verisql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `verisql.check(p_sql text)` is an extension function and returns `TABLE`.
- `verisql.diff(p_sql_a text, p_sql_b text)` is an extension function and returns `boolean`.
- `verisql.explain_sanity(p_sql text)` is an extension function and returns `TABLE`.
- `verisql.fingerprint(p_sql text)` is an extension function and returns `text`.
- `verisql.history_check(p_sql text, p_tag text)` is an extension function and returns `TABLE`.
- `verisql.history_record(p_sql text, p_tag text)` is an extension function and returns `void`.
- `verisql.query_history` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- The control file marks the extension as trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
