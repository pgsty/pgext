## Usage

Sources:

- [README.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/README.md)
- [pg_duckdb.control](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/pg_duckdb.control)
- [settings.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/settings.md)
- [functions.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/functions.md)
- [transactions.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/transactions.md)
- [gotchas_and_syntax.md](https://github.com/duckdb/pg_duckdb/blob/v1.2.1/docs/gotchas_and_syntax.md)
- [v1.2.1](https://github.com/duckdb/pg_duckdb/releases/tag/v1.2.1)

`pg_duckdb` embeds DuckDB's analytical engine in PostgreSQL and can query ordinary PostgreSQL tables and external files. Release 1.2.1 embeds DuckDB 2.0.1; upstream lists PostgreSQL 14–18. It is not relocatable and needs cluster preload.

### Enable and Compare a Query

Add the library to the existing setting, then restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_duckdb'
```

As a superuser, install it and inspect an analytical query:

```sql
CREATE EXTENSION pg_duckdb;
CREATE TABLE orders (category text, amount numeric);
INSERT INTO orders VALUES ('a', 10), ('a', 20), ('b', 5);
SET duckdb.force_execution = true;
EXPLAIN SELECT category, sum(amount) FROM orders GROUP BY category;
SELECT category, sum(amount) FROM orders GROUP BY category;
SET duckdb.force_execution = false;
```

Ordinary queries use PostgreSQL unless forced; DuckDB-only functions or remote storage operations choose DuckDB automatically. Inspect the plan and benchmark representative data. `read_parquet`, `read_csv` and `read_json` read external data, while `duckdb.create_simple_secret` configures storage credentials. Keep secrets out of shared SQL history and verify the access boundary before granting execution.

### Privileges and Configuration

`duckdb.postgres_role` selects the role allowed to execute DuckDB, manage secrets and use MotherDuck data; by default only superusers may do so. A configured role can be created during installation. Grant membership deliberately because it exposes shared DuckDB secrets and data. Local filesystem access is restricted separately through PostgreSQL server-file roles and DuckDB filesystem settings. Community extension loading is an explicit administrator choice.

### Transaction and Compatibility Limits

A transaction may not write both PostgreSQL and DuckDB tables, or mix DDL for both engines. Keep `duckdb.unsafe_allow_mixed_transactions` disabled: bypassing the rule can commit one side only and lose consistency. Execution inside PostgreSQL functions is also disabled by default; the unsafe setting for it has known crash risks.

Version 1.2.1 supports nested fixed-length arrays and improves numeric and floating conversions. It fixes parallel-worker memory reads, cancellation cleanup and secret-bearing log/statistics handling. These fixes do not remove the need to restrict credentials and review type/query compatibility. Install matching libraries, follow preload restart requirements, and update the SQL extension in each database before relying on new behavior.
