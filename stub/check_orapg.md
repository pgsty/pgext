## Usage

Sources:

- [Official extension control file](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/check_orapg.control)
- [Official extension SQL](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/check_orapg--3.0.sql)
- [Official README](https://github.com/yvazquezo/check_orapg/blob/d150ae5dc822c4e8f0eddf8a29e7b3d7055e562f/README.md)

`check_orapg` 3.0 helps validate an Oracle-to-PostgreSQL or Oracle-to-EDB migration. It copies selected Oracle catalog information through `oracle_fdw`, counts objects, privileges, attributes, comments, and table rows on both sides, and emits two text reports for external comparison. It reports migration coverage; it does not convert schemas or prove application-level equivalence.

### Core Workflow

The control file requires `oracle_fdw`, marks installation as superuser-only, and records testing only on PostgreSQL 11–13 and EDB Postgres 11–13. Install Oracle client libraries and configure `oracle_fdw` before creating `check_orapg`.

```sql
CREATE EXTENSION oracle_fdw;
CREATE EXTENSION check_orapg;

SELECT check_orapg.create_oracle_server(
    'ora_migration', '192.0.2.10', 1521, 'ORCL'
);
SELECT check_orapg.create_user(
    'ora_migration', 'migration_reader', 'oracle_password'
);

SELECT check_orapg.create_oracle_tables(
    'ora_migration', NULL, '''HR'',''SALES'''
);

SELECT *
FROM check_orapg.schemas_objects_postgres('''HR'',''SALES''') AS (
    schema_name text, tables int, ordinary_views int,
    materialized_views int, triggers int, sequences int,
    indexes int, functions int, procedures int, constraints int,
    table_comments int, column_comments int
);
```

After importing the Oracle catalogs, refresh them with `check_orapg.update_all_oracle_tables`, register both snapshots with `check_orapg.register_postgres_validation` and `check_orapg.register_oracle_validation`, then compare the rows returned by `check_orapg.postgres_file` and `check_orapg.oracle_file`. The `generate_postgres_file` and `generate_oracle_file` variants write those reports to the database server filesystem.

### Important Objects

- Server and mapping helpers: `create_oracle_server`, `update_server`, `show_servers`, `create_user`, `update_user_password`, and `show_users`.
- Oracle catalog refresh: `create_oracle_tables`, `update_all_oracle_tables`, `update_oracle_table`, and `update_oracle_tables_rows`.
- Comparisons: `cluster_postgres`, `cluster_oracle`, privilege reports, schema object reports, and row-count reports.
- Persistent results: 31 imported Oracle catalog tables plus `postgres_validation` and `oracle_validation` in schema `check_orapg`.

### Security and Accuracy Boundaries

`create_user` creates a user mapping for `PUBLIC`, not just the caller, and the password is supplied as function text and stored in foreign-server metadata. The SQL also constructs dynamic DDL from identifiers and credentials without robust quoting. Restrict all management functions to a tightly controlled administrative role and use a read-only Oracle account.

The file generators use server-side `COPY` and can write only where the PostgreSQL operating-system account has permission. Output paths are server paths. Counts can differ legitimately because PostgreSQL does not implement Oracle packages or synonyms and because object definitions, data values, constraints, and behavior are not compared. Review the generated reports as migration evidence, not as a pass/fail proof.
