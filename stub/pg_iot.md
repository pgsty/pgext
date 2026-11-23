## Usage

Sources:

- [README.md](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/README.md)
- [pg_iot.control](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/pg_iot.control)
- [pg_iot--0.0.1.sql](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/pg_iot--0.0.1.sql)
- [pg_iot.c](https://github.com/pierreforstmann/pg_iot/blob/862938f0f7ad3824b3bd125154483b91bd720356/pg_iot.c)

`pg_iot` marks selected tables as insert-only tables. Version 0.0.1 uses server hooks to reject UPDATE, DELETE, MERGE, TRUNCATE and ALTER TABLE for registered tables while retaining INSERT and SELECT. Upstream reports validation on PostgreSQL 14–19.

### Enable and Register a Table

Add the library to the existing preload list and restart PostgreSQL:

```ini
shared_preload_libraries = 'pg_iot'
```

As a superuser, install the extension in the database and register an existing table by schema and name:

```sql
CREATE EXTENSION pg_iot;
CREATE TABLE public.audit_log (event text);
SELECT iot.pg_iot_set('public', 'audit_log');
INSERT INTO public.audit_log VALUES ('opened');
SELECT * FROM public.audit_log;
SELECT * FROM iot.tables;
```

The non-relocatable extension creates the fixed `iot` schema. `iot.pg_iot_set(cstring, cstring)` returns boolean and is revoked from PUBLIC. `iot.tables` records the namespace, relation name and OID of registered tables. The preload setting must preserve other required libraries.

### Maintenance Boundary

The documented interface has no operation to turn an insert-only table back into an ordinary table. Plan schema changes and data retention before registration. This extension depends on its library being preloaded and on its catalog; it is not a tamper-proof archive against database or operating-system administrators. Ordinary table privileges still determine who may read or insert data.
