## Usage

Sources:

- [Official extension control file](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/pgtsql.control)
- [Official README](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/README.md)
- [Official extension SQL](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/sql/pgtsql--3.0.sql)
- [Official regression example](https://github.com/bigsql/pgtsql/blob/3240980dff20736a38101eb0806c8580ff169d9c/test/sql/pltsql_at_prefixed_vars.sql)

`pgtsql` 3.0 installs the `pltsql` procedural language and a small compatibility layer for Transact-SQL-style PostgreSQL functions and procedures. It is useful for porting limited procedural code that uses `@` variables, T-SQL-like control flow, `PRINT`, and local `#` temporary tables. It is not a complete SQL Server compatibility layer or wire-protocol endpoint.

### Core Workflow

Installing `pgtsql` requires a superuser because it registers a C language handler and validator. A minimal function can use T-SQL-style variable declarations and return syntax:

```sql
CREATE EXTENSION pgtsql;

CREATE FUNCTION tsql_value() RETURNS integer AS $$
DECLARE @value int = 120
BEGIN
    RETURN @value;
END
$$ LANGUAGE pltsql;

SELECT tsql_value();
SELECT sys.getdate(), sys.isnull(NULL::integer, 0);
```

Version 3.0 specifically targeted PostgreSQL 11 and added stored-procedure support. Pin this source and run its regression suite before trying it on a newer PostgreSQL release.

### Installed Surface

- Language: `pltsql`, backed by `pltsql_call_handler` and `pltsql_validator`.
- Date/time helpers: `sys.sysdatetime`, `sys.sysdatetimeoffset`, `sys.sysutcdatetime`, `sys.getdate`, and `sys.getutcdate`.
- Null helper: overloaded `sys.isnull` for a limited set of scalar types.
- Catalog-like views: `sys.tables`, `sys.views`, and `sys.objects` expose selected PostgreSQL relations with SQL Server-shaped columns.

The parser supports selected conventions documented and tested upstream, including optional semicolons, `@`-prefixed variables, T-SQL-style `IF`/`ELSE`, `PRINT`, and `#` temporary-table rewriting. Test every migrated construct; unsupported T-SQL syntax and semantics are not emulated automatically.

### Compatibility and Security Boundaries

Installation creates the `sys` schema unconditionally, so it can conflict with an existing schema. The extension grants `PUBLIC` usage on `sys`, selection from its compatibility views, and execution of its helper functions. Review those grants after installation.

The compatibility views contain placeholders for many SQL Server fields—for example epoch timestamps and constant flags—rather than equivalent server metadata. The project does not provide SQL Server authentication, TDS connectivity, full built-in coverage, SQL Server transaction behavior, or a complete system catalog. Treat `pgtsql` as a narrow PostgreSQL procedural-language porting aid and validate application behavior independently.
