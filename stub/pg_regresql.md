
## Usage

> Syntax:
>
> ```bash
> regresql init postgres://localhost/mydb
> regresql add src/sql/
> regresql update
> regresql test
> ```
>
> Sources: [README](https://github.com/boringsql/regresql), [Product page](https://boringsql.com/products/regresql/)

`RegreSQL` is documented upstream as a language-agnostic SQL regression testing tool for PostgreSQL, not as a `CREATE EXTENSION`-style in-database module. It discovers `.sql` files, runs them against PostgreSQL, snapshots expected output, and tracks query plan changes.

## Quick Start

The README's basic workflow is:

```bash
regresql init postgres://localhost/mydb
regresql discover
regresql add src/sql/
regresql update
regresql test
```

This initializes a test suite, discovers query files, creates plan definitions, captures expected output, and runs regression checks.

## What It Tracks

The upstream docs emphasize:

- expected query output snapshots
- `EXPLAIN` plan baselines
- sequential scan warnings
- migration-related query regressions
- CI-oriented output formats such as `junit`, `json`, `pgtap`, and `github-actions`

## Query Files and Plans

RegreSQL works with normal SQL files and supports multiple queries per file using `-- name:` annotations:

```sql
-- name: get-user-by-id
SELECT * FROM users WHERE id = :id;
```

Plan files provide test parameters:

```yaml
"1":
  id: 42
"2":
  id: 100
```

## Snapshots and Migrations

The tool can build and restore database snapshots and compare query behavior across migrations:

```bash
regresql snapshot build
regresql snapshot restore
regresql migrate --script db/migrations/001_add_column.sql
```

## Installation

The README documents installation via Homebrew or Go:

```bash
brew tap boringsql/boringsql
brew install regresql
```

or

```bash
go install github.com/boringsql/regresql@latest
```

PostgreSQL client tools such as `pg_dump`, `pg_restore`, and `psql` are required for snapshot commands.
