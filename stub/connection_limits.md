## Usage

Sources:

- [Upstream README](https://github.com/tvondra/connection_limits/blob/ab816a0a386978e124c24b2695088aef661f0e84/README.md)
- [Version 1.0.0 install SQL](https://github.com/tvondra/connection_limits/blob/ab816a0a386978e124c24b2695088aef661f0e84/sql/connection_limits--1.0.0.sql)
- [Connection hook implementation](https://github.com/tvondra/connection_limits/blob/ab816a0a386978e124c24b2695088aef661f0e84/src/connection_limits.c)

connection_limits enforces concurrent-connection quotas by database name, role name, client address, or a combination of those fields. It is useful when PostgreSQL's built-in role and database connection limits are not expressive enough.

### Server configuration

The library must reserve shared memory during server startup. Add it to the preload list, optionally define defaults, and restart PostgreSQL:

```ini
shared_preload_libraries = 'connection_limits'
connection_limits.per_user = 5
```

Specific rules are read from PGDATA/pg_limits.conf. The mask column is optional, and all acts as a wildcard:

```text
database  username  IP              [mask]         limit
all       foouser   all                             10
foodb     all       all                             20
all       all       192.168.1.0/24                  10
```

After installing the SQL extension, inspect counters or reload the rule file with:

```sql
CREATE EXTENSION connection_limits;

SELECT * FROM connection_limits;
SELECT connection_limits_reload();
```

The preload setting and shared-memory allocation still require a restart. The SQL reload function only rereads rules after the library is active.

### Caveats

- Superuser connections are counted but not rejected. They can therefore consume a quota and prevent ordinary roles from connecting.
- Rules are keyed by the database and role names supplied during login, not by object identifiers. Update the rule file after a rename.
- File rules override the per-database, per-user, and per-IP defaults. Review overlapping rules carefully.
- The install SQL leaves `connection_limits_reload()` executable by `PUBLIC`, and the C function performs no superuser check before replacing shared rules. Revoke that privilege from untrusted roles; otherwise any database user can trigger rule reloads.
- Prefer PostgreSQL's built-in CONNECTION LIMIT for simple per-role or per-database quotas. This extension uses server hooks and old internal APIs, and upstream provides no current major-version compatibility matrix.
