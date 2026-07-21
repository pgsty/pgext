## Usage

Sources:

- [pgclone v4.4.2 README](https://github.com/valehdba/pgclone/blob/v4.4.2/README.md)
- [pgclone v4.4.2 usage guide](https://github.com/valehdba/pgclone/blob/v4.4.2/docs/USAGE.md)
- [Async cloning guide](https://github.com/valehdba/pgclone/blob/v4.4.2/docs/ASYNC.md)
- [pgclone v4.4.2 release notes](https://github.com/valehdba/pgclone/releases/tag/v4.4.2)

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
- Async schema/database/parallel paths do not honor masks, tables, or exclude_tables in v4.4.2. Use the documented synchronous path when those controls are a security requirement.
- Keep passwords out of stored SQL and logs; prefer libpq service files, passfiles, or another controlled credential mechanism.
- Version 4.4.2 improves sequence-state copying and protects PostgreSQL 17 source sessions from transaction_timeout, but callers must still validate object ownership, extensions, roles, large objects, and post-clone application behavior.
