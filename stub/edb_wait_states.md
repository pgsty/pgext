## Usage

Sources:

- [Official EDB Wait States documentation](https://www.enterprisedb.com/docs/pg_extensions/wait_states/)
- [Official installation guide](https://www.enterprisedb.com/docs/pg_extensions/wait_states/installing/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/wait_states/using/)

`edb_wait_states` samples backend wait events over time and persists query, session, system, and wait-event data for performance analysis.

### Enablement

Install EDB Wait States 1.6.0, preload its worker, restart, and create the extension:

```ini
shared_preload_libraries = 'edb_wait_states'
edb_wait_states.sampling_interval = '1s'
edb_wait_states.retention_period = 604800
```

```sql
CREATE EXTENSION edb_wait_states;
```

Grant `USAGE` on the installation schema to `pg_monitor` when EDB diagnostic tools must read the data.

### Analyze Samples

```sql
SELECT query, session_id, wait_event_type, wait_event
FROM edb_wait_states_data(
  now() - interval '15 minutes',
  now()
);

SELECT * FROM edb_wait_states_wait_events();
SELECT edb_wait_states_directory_size();
```

`edb_wait_states_queries`, `edb_wait_states_sessions`, and `edb_wait_states_sql_statements` provide higher-level views over sampled files.

### Retention and Privacy

Samples are stored under `edb_wait_states.directory`; changing that path requires a restart. Sampling increases disk usage and can retain query text, role, database, and session details. Set a finite retention period, protect the directory and SQL functions, monitor its size, and use `edb_wait_states_purge` when policy requires early removal.

