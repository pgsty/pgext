## Usage

Source: [README](https://github.com/Snehil-Shah/pg_dispatch/blob/master/README.md), [database.dev page](https://database.dev/Snehil_Shah/pg_dispatch)

`pg_dispatch` is a TLE-compatible async SQL dispatcher built on `pg_cron`. It is intended for deferring side effects out of the caller's main transaction in sandboxed environments such as RDS or Supabase.

### Requirements and install

- PostgreSQL 13+
- `pg_cron` 1.5.0+
- `pgcrypto`

```sql
SELECT dbdev.install(Snehil_Shah@pg_dispatch);
CREATE EXTENSION "Snehil_Shah@pg_dispatch";
```

The extension installs into the `pgdispatch` schema.

### Main functions

```sql
SELECT pgdispatch.fire('SELECT pg_sleep(40);');
SELECT pgdispatch.snooze('SELECT pg_sleep(20);', '20 seconds');
```

- `pgdispatch.fire(command text)`: enqueue SQL for immediate async execution.
- `pgdispatch.snooze(command text, delay interval)`: enqueue SQL for delayed async execution.

### Typical use

The official README positions `pg_dispatch` for PL/pgSQL or trigger-based workflows where a foreground RPC should commit quickly while notifications, analytics updates, or other expensive SQL run later in a separate transaction.

### Caveats

- The runtime dependency on `pgcrypto` is in addition to `pg_cron`.
- The `delay` argument truncates to seconds precision.
- The project documents TLE/database.dev installation first; manual PGXN installation is also linked from the README.
