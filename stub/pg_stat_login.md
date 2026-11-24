## Usage

Sources:

- [README](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/README.md)
- [Control](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/pg_stat_login.control)
- [SQL 1.0](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/pg_stat_login--1.0.sql)
- [Authentication hook](https://github.com/asotolongo/pg_stat_login/blob/0c97157ab7e621410c5f7cc4c3f767b72acb86ac/pg_stat_login.c)

`pg_stat_login` 1.0 records successful and failed authentication by user/database pair using PostgreSQL's Custom Cumulative Statistics API. It requires PostgreSQL 18 or later. It is a distinct extension from the older `pg_login_stat` for PostgreSQL 15–17.

### Enable and Inspect

Append the library to the existing preload list, restart PostgreSQL and install the SQL objects as a superuser:

```ini
shared_preload_libraries = 'pg_stat_login'
```

```sql
CREATE EXTENSION pg_stat_login;
SELECT username, datname, login_ok, login_fail,
       last_login_ok, last_login_fail
FROM pg_stat_logins;
SELECT * FROM pg_stat_login_info();
```

Members of `pg_read_all_stats` can read the view and inspection functions. Only superusers can reset the statistics. The extension is not relocatable.

### Capacity and Maintenance

The shared table has a fixed capacity of 10,000 distinct pairs. `pg_stat_login_info()` reports `max_entries`, `num_entries`, `n_dropped` and `stats_reset`. At capacity, existing pairs keep accumulating but events for new pairs are dropped. Failed attempts also occupy slots.

```sql
SELECT pg_stat_login_reset();
ALTER SYSTEM SET pg_stat_login.enable = off;
SELECT pg_reload_conf();
```

Reset clears all accumulated counts and reclaims slots. The reloadable `pg_stat_login.enable` defaults on. Counts survive clean restarts but are discarded after crash recovery; this is cumulative monitoring, not a durable audit trail.

### Authentication Boundaries

`login_ok` means authentication passed, not that a usable session opened. Later failures, including an unknown database or reserved connection slots, can still count as successful authentication. Rejections before the authentication hook, including connection exhaustion and `pg_hba.conf` rejection, are not counted. A client disconnecting before supplying a password is also excluded.

Names are recorded as received, including nonexistent roles or databases. Invalid database-encoding bytes are escaped; valid names that cannot be represented in the client encoding can make inspection fail. Check server logs for events outside the hook and correlate names with PostgreSQL catalogs when interpreting counts. The extension uses custom statistics kind ID 30.
