## Usage

Source: [README](https://github.com/anfiau/pgelog/blob/master/README.md), [control file](https://github.com/anfiau/pgelog/blob/master/pgelog.control), [extension SQL 1.0.2](https://github.com/anfiau/pgelog/blob/master/pgelog--1.0.2.sql), [Tag v1.0.2](https://github.com/anfiau/pgelog/tree/v1.0.2)

`pgelog` writes rollback-resistant log rows through pseudo-autonomous transactions implemented with `dblink`. It caches the extra connection in `pg_variables` so repeated logging in the same session is cheaper.

### Requirements and install

- PostgreSQL 11+
- `dblink`
- `pg_variables`
- passwordless local `dblink` access, typically via `peer`

```sql
CREATE EXTENSION IF NOT EXISTS dblink;
CREATE EXTENSION IF NOT EXISTS pg_variables;
CREATE EXTENSION pgelog;
```

### Main objects and functions

```sql
SELECT pgelog_to_log('SQL', 'standalone', 'Test of logging by pgelog', '1');

SELECT pgelog_get_param('pgelog_ttl_minutes');
SELECT pgelog_set_param('pgelog_ttl_minutes', '2880');
```

- `pgelog_logs`: base log table.
- `pgelog_vw_logs`: log view with timing deltas.
- `pgelog_params`: configuration table.
- `pgelog_to_log(...)`: write a log row that survives rollback.
- `pgelog_get_param(...)`, `pgelog_set_param(...)`, `pgelog_delete_param(...)`: manage extension settings.

### Typical use

The official README shows `pgelog_to_log(...)` in PL/pgSQL exception handlers: collect diagnostics with `GET STACKED DIAGNOSTICS`, write a `FAIL` log row, then re-raise the error.

### Caveats

- Each session can open one additional `dblink` connection, so `max_connections` should account for that.
- Upstream `v1.0.2` still ships extension SQL under the same user-visible object family; Pigsty's note about runtime `dblink` plus `pg_variables` requirements is confirmed by the control file and README.
