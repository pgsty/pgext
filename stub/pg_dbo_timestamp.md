## Usage

Sources:

- [Official README at the reviewed commit](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/README.md)
- [Version 1.0.1 install SQL](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/scripts/pg_dbo_timestamp--1.0.1.sql)
- [DDL event function source](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/DATABASE/SCHEMA/public/FUNCTION/dbots_on_ddl_event.sql)
- [Reporting view source](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/DATABASE/SCHEMA/public/VIEW/dbots_object_timestamps.sql)
- [Extension control file](https://github.com/pgcodekeeper/pg_dbo_timestamp/blob/8392fee066d972cb97363a1e822a54251cd55840/pg_dbo_timestamp.control)

`pg_dbo_timestamp` records the latest observed DDL modification time, current user, session user, and client address for database objects. It is intended to support schema-comparison tools such as pgCodeKeeper. It is not an immutable or tamper-evident security audit log.

### Core Workflow

The DDL event trigger is deliberately installed disabled and must be enabled manually after extension creation:

```sql
CREATE SCHEMA schema_audit;
CREATE EXTENSION pg_dbo_timestamp SCHEMA schema_audit;
ALTER EVENT TRIGGER dbots_tg_on_ddl_event ENABLE;

GRANT USAGE ON SCHEMA schema_audit TO app_ddl;
GRANT SELECT, INSERT, UPDATE, DELETE
  ON schema_audit.dbots_event_data TO app_ddl;
GRANT SELECT
  ON schema_audit.dbots_object_timestamps TO schema_reader;

SELECT type, schema, name, last_modified, cur_user, ses_user, ip_address
FROM schema_audit.dbots_object_timestamps
ORDER BY last_modified DESC;
```

Every role whose DDL should be recorded needs access to the extension schema and read/write access to `dbots_event_data`; otherwise the event function catches the error, emits a warning, and lets DDL continue with stale metadata.

### Main Objects

`dbots_tg_on_ddl_event` runs `dbots_on_ddl_event()` at `ddl_command_end` for creates and alterations. `dbots_tg_on_drop_event` runs `dbots_on_drop_event()` for drops and is installed enabled. `dbots_event_data` stores catalog class/object identifiers and the latest metadata. `dbots_object_timestamps` resolves human-readable identities and joins current object ACLs.

Column-level drops update the parent object's timestamp. GRANT and REVOKE changes are not recorded as modification events because PostgreSQL does not provide the required event-trigger data; the view instead reads current ACLs at query time.

### Trust and Maintenance Boundaries

Roles granted DML on `dbots_event_data` can alter or delete the recorded history, and the functions intentionally downgrade internal failures to warnings. The table keeps only the latest observation per object, not an append-only event stream. Treat the data as a cache for tooling, monitor warnings, and reconcile it when privileges, extension upgrades, restores, or unsupported DDL can make it stale.

Event triggers are database-wide and installation requires elevated privileges. Test the extension with every DDL deployment path, extension-management operation, and PostgreSQL major version used by the database.
