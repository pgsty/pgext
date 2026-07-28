## Usage

Sources:

- [Official upstream README](https://github.com/x4m/pg_tm_aux/blob/857b9173069068741a52e8343bc532bd094fa2b9/README.md)
- [Official extension control file (pg_tm_aux.control)](https://github.com/x4m/pg_tm_aux/blob/857b9173069068741a52e8343bc532bd094fa2b9/pg_tm_aux.control)
- [Official extension SQL (pg_tm_aux--1.0.sql)](https://github.com/x4m/pg_tm_aux/blob/857b9173069068741a52e8343bc532bd094fa2b9/pg_tm_aux--1.0.sql)

`pg_tm_aux` — **(Not needed since Postgres 17, use logical slots failover)**. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_tm_aux;

SELECT * from  pg_create_logical_replication_slot_lsn('dtt3gjq2tfmocenb6vru', 'wal2json', false, pg_lsn('1/20030948'));
SELECT * from pg_logical_slot_peek_changes('dtt3gjq2tfmocenb6vru', null, null);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pg_create_logical_replication_slot_lsn(IN slot_name name, IN plugin name, IN temporary boolean DEFAULT false, IN restart_lsn pg_lsn DEFAULT null, IN force boolean DEFAULT true, OUT slot_name name, OUT lsn pg_lsn)` is an extension function and returns `RECORD`.
- `pg_create_logical_replication_slot_lsn(IN slot_name name, IN plugin name, IN temporary boolean DEFAULT false, IN restart_lsn pg_lsn DEFAULT null, OUT slot_name name, OUT lsn pg_lsn)` is an extension function and returns `RECORD`.

### Requirements and Caveats

- The reviewed control file declares default version `1.1.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
