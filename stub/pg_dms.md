## Usage

Sources:

- [Pinned upstream README](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/README.md)
- [Version 0.0.1 installation SQL](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/pg_dms--0.0.1.sql)
- [Pinned C implementation](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/pg_dms.c)
- [Pinned regression example](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/sql/pg_dms_test.sql)
- [Pinned extension control file](https://github.com/abris-platform/pg_dms/blob/ec456d7a0f0ded981d7e4c516ccdeb20fd8a1197/pg_dms.control)

pg_dms version 0.0.1 is an application-specific document management data model. It implements pg_dms_id, pg_dms_family, and pg_dms_ref types, cross-type B-tree comparisons, UUID casts, embedded action histories, record JSON and hash helpers, and registry import/export tables and triggers. It is not a general document store.

### Identifier and action example

```sql
CREATE EXTENSION pg_dms CASCADE;

CREATE TABLE public.directory (
  key pg_dms_id NOT NULL DEFAULT uuid_generate_v4(),
  num integer NOT NULL
);

INSERT INTO public.directory (num) VALUES (1);

UPDATE public.directory
SET key = pg_dms_setaction(
  key,
  100,
  'public.directory'::regclass,
  uuid_generate_v4()
);

SELECT key, pg_dms_getstatus(key), pg_dms_getlevel(key)
FROM public.directory;
```

A pg_dms_id text value consists of two UUIDs separated by an underscore. A new ID records the creating role and transaction-start timestamp. Each setaction call appends another action structure to the variable-length value, so indexed keys and row size grow with history.

### SQL safety, schema, and compatibility limits

The control file declares the extension relocatable, but its SQL hard-codes tables and functions in public and operators in pg_catalog. It also creates generic action_list and registry objects. Install only in an isolated application database after checking for name and operator conflicts; do not move the extension schema.

Registry helpers and triggers construct dynamic INSERT statements by concatenating JSON-supplied schema names, table names, column names, and values without identifier or literal quoting. Untrusted input can change the generated SQL. Revoke all functions and registry tables from PUBLIC, and do not use the import path until it has been replaced with parameterized, schema-allowlisted code.

Several C functions are declared IMMUTABLE even though they read the current role, transaction time, or current clock. The planner can cache or reorder those calls incorrectly. The record helpers depend on PostgreSQL internal tuple APIs, the SQL creates obsolete PostgreSQL-10-era table options, and no modern-major compatibility matrix is supplied.

The repository has not changed since 2018 and contains no license file. Treat it as unmaintained historical application code: build against the exact target server, inspect every created object and grant, validate binary ordering and hashing, and use only with independent backups and application-specific regression tests.
