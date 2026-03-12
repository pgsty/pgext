

## Usage

> [pg_get_functiondef: Get function's definition](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/pg_get_functiondef)

The `pg_get_functiondef` extension provides functions to retrieve the complete definition (DDL) of PostgreSQL functions and procedures, particularly useful in Oracle compatibility contexts within IvorySQL.

### Enabling

```sql
CREATE EXTENSION pg_get_functiondef;
```

### Retrieving Function Definitions

```sql
-- Get the DDL of a function by OID
SELECT pg_get_functiondef(oid) FROM pg_proc WHERE proname = 'my_function';

-- Get function definition by name
SELECT pg_get_functiondef('my_function'::regproc);
```

This extension extends the built-in `pg_get_functiondef()` to support Oracle-compatible function and procedure syntax used by IvorySQL, including PL/iSQL procedure bodies and Oracle-style parameter declarations.
