## Usage

Sources:

- [pgAudit 18.0 README](https://github.com/pgaudit/pgaudit/blob/18.0/README.md)
- [PostgreSQL 19 preview](https://github.com/pgaudit/pgaudit/releases/tag/19rc1)
- [Preview control file](https://github.com/pgaudit/pgaudit/blob/2a7322004a90ebb8cf51d929e3dcf79e791b557e/pgaudit.control)

`pgaudit` records session and object audit events through PostgreSQL logging. Choose the branch matching the server: 18.x for PostgreSQL 18, 17.x for 17, 16.x for 16, 1.7.x for 15 and 1.6.x for 14. Upstream 19rc1 is a preview for PostgreSQL 19 with SQL version 19.0; the catalog's packaged version remains stable 18.0.

### Enable auditing

Add `pgaudit` to `shared_preload_libraries` and restart the server. As a superuser, create the extension before setting `pgaudit.log`:

```sql
CREATE EXTENSION pgaudit;
```

Only a superuser can change audit settings. Size and protect the log volume: SQL text and parameters can contain sensitive information, and broad logging can consume substantial space. The extension provides audit records; it does not itself establish compliance.

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
```text
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
