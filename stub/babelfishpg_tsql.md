

## Usage

> [babelfishpg_tsql: SQL Server Transact SQL compatibility](https://babelfishpg.org/)

The `babelfishpg_tsql` extension provides Microsoft SQL Server Transact-SQL (T-SQL) compatibility for PostgreSQL as part of the Babelfish project. Applications written for SQL Server can connect and run queries against PostgreSQL with minimal changes.

### Enabling

```sql
CREATE EXTENSION babelfishpg_tsql;
```

### Key Features

- **T-SQL Language Support**: Understands T-SQL syntax including stored procedures, functions, triggers, and batches
- **SQL Server Wire Protocol**: Applications can connect using the TDS (Tabular Data Stream) protocol on port 1433
- **System Procedures**: Common `sp_` system stored procedures are available
- **System Views**: SQL Server catalog views (e.g., `sys.tables`, `sys.columns`, `sys.objects`)
- **Multi-Database Semantics**: Supports SQL Server-style database/schema separation

### Supported T-SQL Features

- `BEGIN...END` blocks, `IF...ELSE`, `WHILE` loops
- `TRY...CATCH` error handling
- Temporary tables (`#temp`, `##global_temp`)
- Table variables (`DECLARE @t TABLE (...)`)
- `IDENTITY` columns and `@@IDENTITY` / `SCOPE_IDENTITY()`
- `TOP` clause, `OUTPUT` clause
- `MERGE` statements
- Common Table Expressions (CTEs)
- Cross-database queries within the same instance
- `EXEC` / `EXECUTE` for dynamic SQL
- SQL Server-style string concatenation and NULL handling
- `PRINT` and `RAISERROR` statements

### Connecting via TDS Protocol

Applications can connect using SQL Server drivers (JDBC, ODBC, ADO.NET) to the TDS listener port (default 1433):

```
Server: hostname
Port: 1433
Database: mydb
```

### Notes

- Requires `babelfishpg_common` extension
- Part of the Babelfish for PostgreSQL project (Apache 2.0 / PostgreSQL license)
- Not all T-SQL features are supported; check the Babelfish compatibility reference
