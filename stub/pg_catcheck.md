


## Usage

> [pg_catcheck: Diagnosing system catalog corruption](https://github.com/EnterpriseDB/pg_catcheck)

`pg_catcheck` is a command-line tool that checks PostgreSQL system catalogs for corruption by verifying cross-references between catalog tables. It accepts the same connection parameters as other PostgreSQL utilities (`-h`, `-p`, `-U`, `-d`).

### Basic Usage

```bash
pg_catcheck -h localhost -p 5432 -d mydb
```

Normal output when no issues are found:

```
progress: done (0 inconsistencies, 0 warnings, 0 errors)
```

### Example Output with Corruption

```
notice: pg_class row has invalid relnamespace "24580": no matching entry in pg_namespace
row identity: oid="24581" relname="foo" relkind="r"
notice: pg_type row has invalid typnamespace "24580": no matching entry in pg_namespace
row identity: oid="24583"
progress: done (4 inconsistencies, 0 warnings, 0 errors)
```

### Result Categories

- **Inconsistencies**: logical problems in catalog cross-references (e.g., dangling OID references)
- **Warnings**: more serious issues
- **Errors**: inability to read catalogs

### Options

```bash
pg_catcheck --help                      # Full list of options
pg_catcheck --select-from-relations     # Also check for missing/inaccessible relation files
```

### Connection

Supports the same options as `psql`: `-h` host, `-p` port, `-U` user, `-d` database, or connection strings/URLs.
