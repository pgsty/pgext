## Usage

Sources:

- [README](https://github.com/shaohuasong-fang/pgsqlauditengine/blob/4da0480d476357735f166377ee0c7179b046004c/README.md)
- [Control file](https://github.com/shaohuasong-fang/pgsqlauditengine/blob/4da0480d476357735f166377ee0c7179b046004c/pgsqlauditengine.control)
- [pgsqlauditengine--1.0.sql](https://github.com/shaohuasong-fang/pgsqlauditengine/blob/4da0480d476357735f166377ee0c7179b046004c/pgsqlauditengine--1.0.sql)

`pgsqlauditengine` checks SQL policy before execution through PostgreSQL hooks and exposes rules and audit records through an embedded REST service. The documented compatibility range is PostgreSQL 11–18.

### Enablement

Add the library to the existing preload list, enable checks, and restart. The unusual spelling of the GUC prefix below is intentional. Configure a nonempty API token before enabling remote access.

```conf
shared_preload_libraries = 'pgsqlauditengine'
PGSAUDAUDITENGINE.enabled = on
PGSAUDAUDITENGINE.check_dml = on
PGSAUDAUDITENGINE.api_listen = '127.0.0.1'
PGSAUDAUDITENGINE.api_port = 8918
```

```sql
CREATE EXTENSION pgsqlauditengine;
SELECT name, setting FROM pg_settings WHERE name LIKE 'PGSAUDAUDITENGINE.%';
```

### Rules and API

The extension DDL registers metadata; preloading activates the hooks. `PGSAUDAUDITENGINE.enabled` defaults to off. Rules cover DDL, DML, privileges, transaction statements, and programmable objects. Both `ERROR` and `WARNING` rule severities block execution; `NOTICE` allows it.

The API exposes `/api/v1/health`, `/api/v1/rules`, `/api/v1/audit-logs`, and `/api/v1/config`. Rule changes take effect at runtime. Audit records use a ring buffer rather than a durable audit archive. An empty `PGSAUDAUDITENGINE.api_token` disables authentication, and the health endpoint is unauthenticated. Keep the service restricted and manage grants/configuration as an administrator.
