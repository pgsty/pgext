
## Usage

Sources: [README](https://github.com/rodo/pg_ddl_historization/blob/master/README.md), [releases](https://github.com/rodo/pg_ddl_historization/releases)

`ddl_historization` is a PostgreSQL extension that records database DDL changes in a historization table. The upstream README documents installation via `make install`, `pgxn install ddl_historization`, and an AWS RDS path via `pg_tle`.

### Enable logging

```sql
CREATE EXTENSION ddl_historization;
```

The README describes the extension as using PostgreSQL event triggers to historize DDL changes made in the database.

### What upstream currently documents

- Cluster-local install: `make install`
- PGXN install: `pgxn install ddl_historization`
- AWS RDS / `pg_tle`: build `pgtle.ddl_historization-0.3.sql` with `make pgtle`
- Test suite: `make test` with pgTAP

### Release notes worth knowing

- Release `0.2` is the version requested by this refresh task.
- Release `0.0.4` says it added functions to start and stop logging.
- Release `0.0.6` says it added a `ddl_history_column` table.
- Release `0.0.7` says it fixed a foreign-key related logging bug.

### Caveat

The current upstream README is minimal and does not document the exact SQL signatures for the start/stop logging functions or the schema of the historization tables added in later releases. Keep this stub conservative unless the upstream README or release notes become more explicit.
