


## Usage

> [pgpool_regclass: replacement for regclass](https://pgpool.net/)

The `pgpool_regclass` extension provides a replacement `regclass` function used internally by Pgpool-II to handle relation name resolution across multiple backends.

### Function

```sql
-- Resolve a relation name to its OID, similar to PostgreSQL's regclass cast
SELECT pgpool_regclass('my_table');
SELECT pgpool_regclass('my_schema.my_table');
```

### Purpose

In standard PostgreSQL, casting a string to `regclass` (e.g., `'my_table'::regclass`) resolves the relation name to an OID. However, Pgpool-II needs to determine whether a SQL statement references a temporary table or a regular table to route queries correctly.

The `pgpool_regclass` function provides this resolution capability as a regular function call rather than a type cast, allowing Pgpool-II to:

- Determine if a referenced table exists on the backend
- Distinguish temporary tables from permanent tables for proper query routing
- Handle schema-qualified table names correctly across pooled connections

### Notes

- This extension is primarily used internally by Pgpool-II and is not typically called directly by applications
- It should be installed on all PostgreSQL backend nodes managed by Pgpool-II
- For Pgpool-II 3.0+, this function helps with the `check_temp_table` feature
