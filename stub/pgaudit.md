## Usage

Sources:

- [pgAudit 19.0 README](https://github.com/pgaudit/pgaudit/blob/19.0/README.md)
- [PostgreSQL 19 release](https://github.com/pgaudit/pgaudit/releases/tag/19.0)
- [Release control file](https://github.com/pgaudit/pgaudit/blob/19.0/pgaudit.control)

`pgaudit` records session and object events through PostgreSQL logging. Use the branch matching the server: 19.0 for PostgreSQL 19, 18.1 for 18, 17.2 for 17, 16.2 for 16, 1.7.2 for 15 and 1.6.4 for 14. The branches are separate server builds; upgrading PostgreSQL does not make an older library ABI-compatible.

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

### Logging boundary

Audit output is best-effort PostgreSQL logging, outside transaction durability. A rolled-back statement can leave an audit record, and a committed transaction is not guaranteed to have a durable audit record. Protect and collect logs separately; do not treat them as a transactionally complete ledger.
