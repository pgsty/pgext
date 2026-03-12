


## Usage

> [sepgsql: SELinux label-based mandatory access control for PostgreSQL](https://www.postgresql.org/docs/current/sepgsql.html)

`sepgsql` provides label-based mandatory access control (MAC) based on SELinux security policy. It adds an extra layer of security checking above PostgreSQL's standard SQL permissions.

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `sepgsql.permissive` | `off` | Enable permissive mode regardless of system SELinux settings |
| `sepgsql.debug_audit` | `off` | Force all possible logging regardless of policy |

### Functions

| Function | Returns | Description |
|----------|---------|-------------|
| `sepgsql_getcon()` | `text` | Get current client security label |
| `sepgsql_setcon(text)` | `boolean` | Switch client domain to new label (NULL to revert) |
| `sepgsql_mcstrans_in(text)` | `text` | Translate qualified MLS/MCS range to raw format |
| `sepgsql_mcstrans_out(text)` | `text` | Translate raw MLS/MCS range to qualified format |
| `sepgsql_restorecon(text)` | `boolean` | Set initial security labels for all objects in database |

### Security Labels

Security labels can be assigned to schemas, tables, columns, sequences, views, and functions:

```sql
SECURITY LABEL ON COLUMN customer.credit
    IS 'system_u:object_r:sepgsql_secret_table_t:s0';
```

### Dynamic Domain Transitions

```sql
SELECT sepgsql_getcon();
-- unconfined_u:unconfined_r:unconfined_t:s0-s0:c0.c1023

SELECT sepgsql_setcon('unconfined_u:unconfined_r:unconfined_t:s0-s0:c1.c4');
-- t
```

### Trusted Procedures

```sql
-- Create function to access sensitive data with masking
CREATE FUNCTION show_credit(int) RETURNS text
    AS 'SELECT regexp_replace(credit, ''-[0-9]+$'', ''-xxxx'', ''g'')
         FROM customer WHERE cid = $1'
    LANGUAGE sql;

-- Mark as trusted procedure
SECURITY LABEL ON FUNCTION show_credit(int)
    IS 'system_u:object_r:sepgsql_trusted_proc_exec_t:s0';
```

### Permission Classes

DML operations check: `db_table:{select|insert|update|delete}` and `db_column:{select|update|insert}`.
DDL operations check: `create`, `drop`, `setattr`, `add_name`, `remove_name`.
Schema access requires: `db_schema:search`.
