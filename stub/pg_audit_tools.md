## Usage

Sources:

- [Upstream README](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/README.md)
- [Version 0.1 control file](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/pg_audit_tools.control)
- [Version 0.1 install SQL](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/sql/pg_audit_tools--0.1.sql)
- [MIT license](https://github.com/alexandersamoylov/pg_audit_tools/blob/25ec4c7f216980825451a2352005567d03e9bbf4/LICENSE)

pg_audit_tools 0.1 is a 2017 pure-SQL prototype intended to add audit columns, a history table, a sequence, and row triggers to an existing table. The published helper is not safe or functional enough to run unchanged.

### Restrict installation

If inspecting it in a disposable database, immediately revoke the default public execution rights:

```sql
BEGIN;
CREATE EXTENSION pg_audit_tools;
REVOKE EXECUTE ON FUNCTION
  audit_tools.table_history_create(character varying, character varying)
  FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION audit_tools.table_history_tf() FROM PUBLIC;
COMMIT;
```

Repair and review the functions before calling table_history_create.

### Caveats

- Both functions are SECURITY DEFINER, are public by default, and have no fixed search path. With a superuser installer, this creates a serious privilege-escalation and object-resolution risk.
- table_history_create defines aud_user as type name but gives it a now() timestamp default. On normal PostgreSQL this expression is type-incompatible and prevents the helper from completing.
- The helper alters the source table and creates several predictably named objects without an idempotent rollback path. Partial failures can leave columns or objects behind.
- History rows are written only for the old versions of UPDATE and DELETE operations. Existing rows and INSERT events are not copied to history.
- This is mutable application history, not tamper-evident auditing. Table owners and privileged roles can change or delete it, and row triggers can be disabled or bypassed.
- Upstream provides only a one-line README and no tests, upgrade scripts, concurrency guidance, retention policy, or current compatibility statement.
