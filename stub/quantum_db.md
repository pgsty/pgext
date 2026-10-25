## Usage

Sources:

- [Official source](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/README.md)
- [Official source](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/pg_extension/README.md)
- [Official source](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/pg_extension/quantum_db.control)
- [Official source](https://github.com/OpenDataBox/Qute/blob/0108b98c04643d121f4e329c946fa50b68911dd1/systems/quantum_db/pg_extension/sql/quantum_db--0.1.1.sql)

`quantum_db` 0.1.1 is a PostgreSQL 16 research bridge to a local tensor-network/quantum sidecar. PostgreSQL remains authoritative and rechecks candidate predicates. The project explicitly has no license; public source availability does not supply reuse permission.

### Start the Local Sidecar

The sidecar requires Python 3.11 or later. Real quantum hardware is optional and separately requires provider access; ordinary PostgreSQL integration does not establish that hardware has run.

```bash
qdb-sidecar --socket /tmp/qdb.sock --state outputs/qdb_jobs.sqlite3
```

Configure PostgreSQL and restart to load planner hooks in all backends:

```ini
shared_preload_libraries = 'quantum_db'
quantum_db.socket_path = '/tmp/qdb.sock'
```

Install as a superuser in each target database; the extension requires the standard procedural language dependency:

```sql
CREATE EXTENSION quantum_db;
SELECT qdb_explain('SELECT 1');
```

### Dataset and Job Interfaces

`qdb_register_dataset` registers a relation, its key column and indexed columns with the sidecar. Registered relations reject inserts, updates, deletes and truncation until `qdb_unregister_dataset` removes registration. Do not register a live write-heavy application table without planning that restriction.

`qdb_submit`, `qdb_status`, `qdb_result` and `qdb_cancel` manage asynchronous jobs. The `qdb` schema stores datasets, artifacts and job audit records. SQL text and selected dataset information cross a local Unix-socket boundary; protect the socket, sidecar state and any cloud credentials. Some SQL wrappers run with definer privileges, so review execution grants before exposing them to untrusted roles.

### Failure and Upgrade Boundaries

The 0.1.1 update adds artifact-checksum checks and SQL job-audit mirroring. An old table layout retains native scanning until upgraded. Install matching files, restart old library users, then update each database:

```sql
ALTER EXTENSION quantum_db UPDATE TO '0.1.1';
```

The sidecar is a separate operational dependency, and its SQLite state is outside a PostgreSQL-only backup. Verify readiness, dataset snapshot alignment and fallback behavior. Research telemetry or candidate generation is not evidence of complete-query quantum acceleration.
