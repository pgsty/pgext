## Usage

Sources:

- [Official upstream README](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/README)
- [Official extension SQL](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/sybase_fdw--1.0.sql)
- [Official implementation](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/sybase_fdw.c)

`sybase_fdw` is an incomplete foreign-data-wrapper skeleton, not a usable Sybase connector. Version 1.0 registers a handler and validator, but every foreign-scan callback is NULL and the validator accepts options without implementing validation. It cannot plan or read a foreign table.

### What Installation Creates

In a disposable development database, installation only demonstrates the registered wrapper objects:

```sql
CREATE EXTENSION sybase_fdw;

SELECT fdwname, fdwhandler::regproc, fdwvalidator::regproc
FROM pg_foreign_data_wrapper
WHERE fdwname = 'sybase_fdw';
```

The extension creates `sybase_fdw_handler()`, `sybase_fdw_validator(text[], oid)`, and the `sybase_fdw` foreign data wrapper. It does not document connection, authentication, server, user-mapping, or foreign-table options.

### Nonfunctional Boundary

The handler returns an `FdwRoutine` whose planning, explanation, begin, iterate, rescan, and end callbacks are all NULL. Therefore, do not publish a `CREATE SERVER` or `CREATE FOREIGN TABLE` workflow as if queries could succeed. The dummy validator also provides no protection from misspelled or unsafe options.

Use the project only as historical FDW scaffolding or as a starting point for development after implementing and testing the complete PostgreSQL FDW lifecycle, a Sybase client connection layer, type conversion, transactions, credentials, errors, cancellation, and cleanup. For actual Sybase access, select a maintained connector with documented compatibility and security behavior; the 1.0 label does not indicate production readiness.
