## Usage

Sources:

- [Current version source](https://api.pgxn.org/src/pgclone/pgclone-4.9.0/README.md)
- [Current version source](https://api.pgxn.org/src/pgclone/pgclone-4.9.0/pgclone.control)
- [Current version source](https://api.pgxn.org/src/pgclone/pgclone-4.9.0/CHANGELOG.md)
- [Current version source](https://api.pgxn.org/src/pgclone/pgclone-4.9.0/sql/pgclone--4.8.0--4.9.0.sql)

- [pgclone 4.8.0 changelog](https://api.pgxn.org/src/pgclone/pgclone-4.8.0/CHANGELOG.md)
- [4.7.0 to 4.8.0 SQL](https://api.pgxn.org/src/pgclone/pgclone-4.8.0/sql/pgclone--4.7.0--4.8.0.sql)
- [pgclone 4.8.0 control](https://api.pgxn.org/src/pgclone/pgclone-4.8.0/pgclone.control)
- [pgclone 4.8.0 README](https://api.pgxn.org/src/pgclone/pgclone-4.8.0/README.md)
- [Upstream source](https://api.pgxn.org/src/pgclone/pgclone-4.7.0/README.md)
- [Upstream source](https://api.pgxn.org/src/pgclone/pgclone-4.7.0/docs/USAGE.md)
- [Upstream source](https://api.pgxn.org/src/pgclone/pgclone-4.7.0/pgclone.control)

- [pgclone v4.5.0 README](https://github.com/valehdba/pgclone/blob/v4.5.0/README.md)
- [pgclone v4.5.0 usage guide](https://github.com/valehdba/pgclone/blob/v4.5.0/docs/USAGE.md)
- [Async cloning guide](https://github.com/valehdba/pgclone/blob/v4.5.0/docs/ASYNC.md)
- [pgclone v4.5.0 release notes](https://github.com/valehdba/pgclone/releases/tag/v4.5.0)

pgclone clones tables, schemas, functions, roles, or whole databases over a PostgreSQL connection. It also provides preflight checks, structural diffs, masking, consistent snapshots, and optional background jobs. Use it for controlled database copies, not as an unattended substitute for backup and recovery.

### Create and Run a Clone

    CREATE EXTENSION pgclone;
    SELECT pgclone.version();

    SELECT pgclone.table(
      'host=source.example dbname=app user=clone_user',
      'public',
      'customers',
      true
    );

Schema and database entry points follow the same connection-first pattern:

    SELECT pgclone.schema(
      'host=source.example dbname=app user=clone_user',
      'sales',
      true
    );

    SELECT pgclone.database(
      'host=source.example dbname=app user=clone_user',
      true
    );

The main API includes pgclone.table, pgclone.schema, pgclone.functions, pgclone.database, and pgclone.database_create. The _ex variants expose explicit choices for indexes, constraints, and triggers.

### Filter and Mask Data

JSON options can restrict columns and rows:

    SELECT pgclone.table(
      'host=source.example dbname=app user=clone_user',
      'public',
      'users',
      true,
      'users_lite',
      '{"columns":["id","name","email"],"where":"active"}'
    );

Version 4.4 adds schema- and database-level masks, table inclusion patterns, and exclude_tables. Mask expressions run in the source-side COPY query, so values that are successfully masked do not reach the target unmasked.

The 4.4.2 mask validator skips unsafe or incompatible masks: constant values that cannot cast to the column, NULL for NOT NULL columns, non-hash masks on unique or primary-key columns, and masks on foreign-key columns. A skipped mask leaves that column unmasked. Treat warnings as a failed privacy gate and inspect the result before distributing a clone.

### Conflict Handling and Remote Mask Reports

From 4.4.3 onward, synchronous table/schema/database clones honor an explicitly supplied `conflict` option. `replace` drops the existing object with cascading dependencies before rebuilding it; `skip`, `error` and `rename` select the other documented behaviors. A caller that previously supplied this ignored option can now perform a destructive operation, so review options before upgrading.

Version 4.4.4 adds bank-account and IBAN discovery and a remote report overload:

```sql
SELECT * FROM pgclone.masking_report(
    'host=source.example dbname=app user=clone_user', 'public');
```

### Oracle Sources in 4.5.0

`pgclone.table_from_oracle()` requires a library compiled with optional Oracle OCI support and a loadable Oracle client on the database host. An ordinary build exposes the function but rejects the operation. The separate `pgclone.oracle_type_to_pg()` helper works without an Oracle client.

```sql
SELECT pgclone.table_from_oracle(
    'dsn=ORACLE_SOURCE user=clone_user',
    'HR', 'EMPLOYEES', dry_run => true);
```

An omitted Oracle password is read from the server process's `PGCLONE_ORACLE_PASSWORD` environment variable. The dry run returns DDL without creating the target. Actual cloning requires an existing target schema and an absent target table; it never replaces that table. DDL and copied rows share one PostgreSQL transaction, with a read-only Oracle snapshot.

Unsupported Oracle column types are skipped with warnings, so validate the resulting column set. Defaults, foreign/check constraints, non-unique indexes, triggers, partitioning, identity properties and masking/filtering are not cloned. Replace the library and update the extension to install the new SQL functions:

```sql
ALTER EXTENSION pgclone UPDATE TO '4.5.0';
```

### Preflight, Diff, and Consistency

    SELECT pgclone.preflight(
      'host=source.example dbname=app user=clone_user',
      'public'
    )::jsonb;

    SELECT pgclone.diff(
      'host=source.example dbname=app user=clone_user',
      'public'
    )::jsonb;

preflight checks connectivity, versions, privileges, capacity, names, roles, extensions, and tablespaces. diff reports DDL differences without applying changes.

Schema and database clones use a shared exported snapshot by default so related tables are copied consistently. A long snapshot can delay source vacuum cleanup and WAL recycling. Set the consistent option to false only when accepting cross-table inconsistency is an explicit tradeoff.

### Async Jobs

Async execution requires preload and a restart:

    shared_preload_libraries = 'pgclone'

    SELECT pgclone.schema_async(
      'host=source.example dbname=app user=clone_user',
      'sales',
      true,
      '{"parallel":4}'
    );

    SELECT * FROM pgclone.jobs_view;
    SELECT pgclone.progress(1);
    SELECT pgclone.cancel(1);

pgclone also exposes progress_detail, resume, and clear_jobs for job administration. Size max_worker_processes for the requested parallelism.

### Important Boundaries

- The upstream usage guide requires superuser privileges to install and use pgclone.
- Async schema/database/parallel paths do not honor masks, tables, or exclude_tables in v4.5.0. Use the documented synchronous path when those controls are a security requirement.
- Keep passwords out of stored SQL and logs; prefer libpq service files, passfiles, or another controlled credential mechanism.
- Version 4.4.2 improves sequence-state copying and protects PostgreSQL 17 source sessions from transaction_timeout, but callers must still validate object ownership, extensions, roles, large objects, and post-clone application behavior.

### SQL Server and MySQL sources in 4.7.0

The 4.6 SQL Server interface requires optional FreeTDS support; 4.7 adds MySQL/MariaDB through an optional client library. The public entry points are pgclone.table_from_sqlserver() and pgclone.table_from_mysql(), with corresponding type-mapping helpers. Builds without the client support reject cloning while retaining the SQL interface. An omitted password uses the server's PGCLONE_SQLSERVER_PASSWORD or PGCLONE_MYSQL_PASSWORD environment variable.

Use dry_run to inspect generated DDL first. Actual imports require an existing target schema and an absent target table, and never overwrite the table. Inspect skipped/unsupported source column types and the documented mapping of keys and indexes; defaults, triggers and full source-database semantics are not reproduced. PostgreSQL DDL and copied rows are transactional, but source snapshot guarantees depend on the selected external engine and table type. Do not assume nontransactional MySQL tables share an InnoDB snapshot.

Install the matching library and optional clients, reconnect sessions, and update the SQL extension to 4.7.0. The older Oracle example above describes that interface's introduction; this is the current SQL upgrade:

```sql
ALTER EXTENSION pgclone UPDATE TO '4.7.0';
```

### 4.8.0 Partition Cloning and Restart Boundary

Version 4.8.0 has no SQL-level object changes, but after `make install` PostgreSQL must be restarted before `ALTER EXTENSION pgclone UPDATE`; a stale preloaded library can otherwise fail to find updated functions. The update adds `"conflict": "ignore"`: when the target table already exists, pgclone keeps its definition and still loads data in both sync and async paths.

Partitioned schema and async clones now create parents before partitions, copy rows once from leaf partitions, and propagate parent primary keys, indexes and triggers rather than duplicating them. Generated columns remain generated and are omitted from COPY so the target recomputes them; PostgreSQL 18 virtual generated columns require PostgreSQL 18 on the target. Async single-table cloning pre-creates serial/bigserial sequences and replays their values, and background workers honour SIGTERM.


### 4.9.0 Partial Clone Failures

Synchronous schema and database clones accept `on_error: continue`. Each failed table rolls back in its own subtransaction, emits its message and SQLSTATE, and is excluded from later foreign-key and trigger passes. The result reports `COMPLETED WITH ERRORS: N table(s) failed: ...`; it is not a fully successful clone. The default remains abort, and cancellation/shutdown errors are not swallowed. Async schema clones already continued after individual table errors. This release adds no SQL objects: install matching files, restart PostgreSQL as directed upstream, then run `ALTER EXTENSION pgclone UPDATE TO '4.9.0'`.
