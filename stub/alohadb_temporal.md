## Usage

Sources:

- [contrib/alohadb_temporal/alohadb_temporal--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_temporal/alohadb_temporal--1.0.sql)
- [contrib/alohadb_temporal/alohadb_temporal.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_temporal/alohadb_temporal.c)
- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_temporal/alohadb_temporal.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_temporal/alohadb_temporal.control)

`alohadb_temporal` is AlohaDB’s system-versioning module. It attaches a history table and row trigger to a base table so updates and deletes retain previous row versions.

### Core Workflow

```sql
CREATE EXTENSION alohadb_temporal;
CREATE TABLE temporal_demo (id integer PRIMARY KEY, label text);
SELECT alohadb_enable_system_versioning('temporal_demo'::regclass);
INSERT INTO temporal_demo (id, label) VALUES (1, 'first');
UPDATE temporal_demo SET label = 'second' WHERE id = 1;
SELECT * FROM alohadb_temporal_status();
```

### Operational Boundaries

This is a module of the PostgreSQL-18-based AlohaDB fork; standard PostgreSQL compatibility has not been verified. Administrator installation is required for its native functions. `alohadb_enable_system_versioning` changes the base table and adds temporal columns, history storage and a trigger, so use a table the caller can alter. `alohadb_temporal_status` lists registrations. `alohadb_as_of` and `alohadb_versions_between` return records whose column types the caller must declare. History grows with writes and needs retention planning. Review the implementation before `alohadb_disable_system_versioning`: it removes versioning infrastructure and temporal columns. This source prototype does not establish an application-wide audit or tamper-resistance guarantee.
