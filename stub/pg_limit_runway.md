## Usage

Sources:

- [README](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/README.md)
- [Control](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/pg_limit_runway.control)
- [Install and grants](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/INSTALL.md)
- [Install SQL](https://github.com/apjapj/pg_limit_runway/blob/001157bde784bd3f9beb002b7a392e740c480f69/pg_limit_runway--0.7.0.sql)

`pg_limit_runway` tracks database hard limits and estimates exhaustion from sampled rates. Release tag v0.7.1 ships control and extension SQL version 0.7.0. This pure SQL extension is tested on PostgreSQL 14–19rc1, requires superuser installation and needs no preload or restart.

### Inspect Headroom

```sql
CREATE EXTENSION pg_limit_runway;
SELECT * FROM limit_runway_findings();
SELECT * FROM limit_runway_unwatched();
```

The catalog readers cover sequence/identity ranges, narrow foreign-key columns, transaction and MultiXact age, partition runway, replication-slot retention, connection headroom and role-password expiry. Unwatched objects and missing rates are explicit; a lack of findings is not proof that all resource limits are safe. Rate-based dates require repeated scheduled samples and can disagree across time windows.

### Access and Scheduling

Grant the documented object/function permissions, statistics access and sequence permissions deliberately. Sampling, pruning and alerts need their additional write grants and an external scheduler. Action helpers return commands for review, not automatic remediation. The generated script-install route is an alternative for environments without extension-file installation; do not combine it with CREATE EXTENSION in the same database. Follow the SQL 0.7.0 upgrade path rather than inventing a 0.7.1 SQL version.
