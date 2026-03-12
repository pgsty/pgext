


## Usage

> [pgaudit: Open Source PostgreSQL Audit Logging](https://github.com/pgaudit/pgaudit)

pgAudit provides detailed session and/or object audit logging via the standard PostgreSQL logging facility, producing audit trails required for government, financial, or ISO certifications.

```sql
CREATE EXTENSION pgaudit;
```

### Configuration Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `pgaudit.log` | `none` | Statement classes to log: `READ`, `WRITE`, `FUNCTION`, `ROLE`, `DDL`, `MISC`, `MISC_SET`, `ALL` |
| `pgaudit.log_catalog` | `on` | Log statements when all relations are in pg_catalog |
| `pgaudit.log_client` | `off` | Show audit log messages to client |
| `pgaudit.log_level` | `log` | Log level for audit entries |
| `pgaudit.log_parameter` | `off` | Include statement parameters in log |
| `pgaudit.log_parameter_max_size` | `0` | Max parameter size in bytes (0=unlimited) |
| `pgaudit.log_relation` | `off` | Separate log entry per relation in SELECT/DML |
| `pgaudit.log_rows` | `off` | Include row count in log |
| `pgaudit.log_statement` | `on` | Include statement text in log |
| `pgaudit.log_statement_once` | `off` | Log statement text only with first entry |
| `pgaudit.role` | (none) | Master role for object audit logging |

### Session Audit Logging

Log all DML and DDL with per-relation detail:

```sql
SET pgaudit.log = 'write, ddl';
SET pgaudit.log_relation = on;
```

Log everything except miscellaneous commands:

```sql
SET pgaudit.log = 'all, -misc';
```

Example output:
```
AUDIT: SESSION,1,1,DDL,CREATE TABLE,TABLE,public.account,create table account(...)
AUDIT: SESSION,2,1,READ,SELECT,,,select * from account
```

### Object Audit Logging

Grant permissions to an audit role to control which relations are logged:

```sql
SET pgaudit.role = 'auditor';

GRANT SELECT, DELETE
   ON public.account
   TO auditor;
```

Now any `SELECT` or `DELETE` on the `account` table will be audit logged.

### Log Format

Entries are CSV with fields: `AUDIT_TYPE`, `STATEMENT_ID`, `SUBSTATEMENT_ID`, `CLASS`, `COMMAND`, `OBJECT_TYPE`, `OBJECT_NAME`, `STATEMENT`, `PARAMETER`.
