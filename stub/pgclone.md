
## Usage

> Syntax:
>
> ```sql
> CREATE EXTENSION pgclone;
> SELECT pgclone_table('host=source-server dbname=mydb user=postgres password=secret',
>                      'public', 'customers', true);
> ```
>
> Sources: [README](https://github.com/valehdba/pgclone), [Usage Guide](https://github.com/valehdba/pgclone/blob/main/docs/USAGE.md)

`pgclone` clones PostgreSQL databases, schemas, tables, functions, roles, and permissions directly from SQL. The upstream README emphasizes that it uses PostgreSQL's COPY protocol and avoids external `pg_dump` / `pg_restore` workflows.

## Core Capabilities

The README lists support for:

- cloning tables, schemas, functions, and full databases
- indexes, constraints, triggers, views, materialized views, and sequences
- selective cloning with column selection and `WHERE` filters
- conflict handling with error, skip, replace, or rename strategies
- data masking and sensitive-column discovery
- async and parallel cloning with background workers

## Basic Functions

```sql
SELECT pgclone_table(
    'host=source-server dbname=mydb user=postgres password=secret',
    'public', 'customers', true
);

SELECT pgclone_schema(
    'host=source-server dbname=mydb user=postgres password=secret',
    'sales', true
);

SELECT pgclone_database(
    'host=source-server dbname=mydb user=postgres password=secret',
    true
);
```

The README also documents `pgclone_version()` for verification after installation.

## Async Mode

For background-worker features, the extension must be preloaded:

```ini
shared_preload_libraries = 'pgclone'
```

The upstream docs describe async operations, progress tracking, and multi-worker parallel cloning in separate documentation pages.

## Requirements

Current upstream requirements are:

- PostgreSQL 14 or newer
- PostgreSQL development headers
- `libpq` development library
- a C compiler such as GCC

The repository homepage is [postgresql.az](https://postgresql.az/), and the project README links separate documents for usage, async operations, testing, and architecture.
