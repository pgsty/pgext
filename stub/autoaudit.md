## Usage

Sources:

- [README.md](https://github.com/Khraben/AutoAudit-Extension/blob/0ceda962dcc56f2c35fcd563a5814c2e9bfc6f13/README.md)
- [autoaudit--1.0.sql](https://github.com/Khraben/AutoAudit-Extension/blob/0ceda962dcc56f2c35fcd563a5814c2e9bfc6f13/autoaudit--1.0.sql)
- [autoaudit.control](https://github.com/Khraben/AutoAudit-Extension/blob/0ceda962dcc56f2c35fcd563a5814c2e9bfc6f13/autoaudit.control)

`autoaudit` 1.0 is a course-project SQL extension that records row inserts, updates and deletes in an audit table. Installation enrolls existing user tables and installs an event trigger for subsequently created tables, so review the database-wide effect before enabling it.

### Core Workflow

```sql
CREATE EXTENSION autoaudit;
SELECT operationtype, tablename, username, databefore, dataafter
FROM autoaudit.audit_log ORDER BY id DESC LIMIT 20;
```

### Operational Boundaries

`autoaudit.audit_log` stores operation, table name, timestamp, user, client address and before/after JSONB values. `autoaudit.audit_function` is the row trigger; `autoaudit.ddl_handler` creates triggers after table creation. Auditing is transactional: rolled-back data changes and their audit rows roll back together.

Installation needs superuser rights for the event trigger. The functions use invoker privileges, and the installer revokes public schema access; other application roles need deliberate audit-schema and audit-table privileges or their writes can fail. Check identifier handling, existing trigger names and table ownership in an isolated database first.

No preload or shared library is used by the SQL, despite an unused library name in the control file. The repository does not declare a license or PostgreSQL-major matrix. This educational implementation is not an independently verified, tamper-resistant compliance audit system.
