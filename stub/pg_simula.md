## Usage

Sources:

- [Official README](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/README.md)
- [Official extension SQL](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/pg_simula--1.0.sql)
- [Official C implementation](https://github.com/MasahikoSawada/pg_simula/blob/532efbc57ebd57de8308af8278e77581bb6be885/pg_simula.c)

`pg_simula` is a failure-injection extension for PostgreSQL testing. It can delay selected command tags or raise `ERROR`, `FATAL`, or `PANIC`, and it can reject new connections. Use it only on disposable test systems: some actions terminate sessions or crash the server.

### Setup and Core Workflow

The library must be preloaded because it installs planner, utility-command, and connection hooks. Add it to `shared_preload_libraries`, restart PostgreSQL, and then create the extension:

```sql
CREATE EXTENSION pg_simula;

SELECT add_simula_event('INSERT', 'WAIT', 10);
SELECT add_simula_event('TRUNCATE TABLE', 'ERROR', 0);

SET pg_simula.enabled = on;
INSERT INTO test_table VALUES (1);
TRUNCATE test_table;

SELECT clear_all_events();
```

Rules are stored in `simula_events(operation, action, sec)`. The `operation` is a PostgreSQL command tag. `WAIT` sleeps for `sec`; `ERROR`, `FATAL`, and `PANIC` inject the corresponding server error level. `clear_all_events()` removes every rule. Prefer the management functions over direct table edits, which can recursively trigger simulations.

### Connection Refusal and Safety

`pg_simula.enabled` controls command simulation. `pg_simula.connection_refuse` rejects new connections with an authorization error; configure and test a recovery path through an existing session or configuration reload before enabling it.

The upstream project reports testing only on Linux with PostgreSQL 10. Hooks and internal APIs are version-sensitive, so do not infer support for later releases. A `PANIC` rule can force PostgreSQL recovery, and a `FATAL` rule disconnects the session. Keep the extension, its configuration, and its event table out of production and remove all rules after each test.
