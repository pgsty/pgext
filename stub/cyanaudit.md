## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/cyanaudit/cyanaudit-2.2.0/README.md)
- [Official extension SQL (cyanaudit--2.0.sql)](https://api.pgxn.org/src/cyanaudit/cyanaudit-2.2.0/sql/cyanaudit--2.0.sql)

`cyanaudit` — Cyan Audit is a PostgreSQL utility providing comprehensive and easily-searchable logs of DML (INSERT/UPDATE/DELETE) activity in your database. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

This component has no confirmed standalone `CREATE EXTENSION` workflow in the reviewed source. Build, load, or enable it only through the exact upstream mechanism, then verify the resulting server behavior in an isolated database.

### Important Objects

- `cyanaudit.fn_activate_partition(in_partition_name varchar)` is an extension function and returns `void`.
- `cyanaudit.fn_add_trigger_to_extension(in_table_schema varchar, in_table_name varchar)` is an extension function and returns `void`.
- `cyanaudit.fn_after_audit_field_change()` is an extension function and returns `trigger`.
- `cyanaudit.fn_archive_partition(in_partition_name varchar)` is an extension function and returns `void`.
- `cyanaudit.fn_before_audit_field_change()` is an extension function and returns `trigger`.
- `cyanaudit.fn_create_event_trigger()` is an extension function and returns `void`.
- `cyanaudit.fn_create_new_partition(in_new_table_name varchar default 'tb_audit_event_' || to_char(now(), 'YYYYMMDD_HH24MI'))` is an extension function and returns `varchar`.
- `cyanaudit.fn_create_partition_indexes(in_table_name varchar)` is an extension function and returns `void`.
- `cyanaudit.fn_get_active_partition_name()` is an extension function and returns `varchar`.
- `cyanaudit.fn_get_current_uid()` is an extension function and returns `integer`.
- `cyanaudit.fn_get_email_by_uid(in_uid integer)` is an extension function and returns `varchar`.
- `cyanaudit.fn_get_last_txid()` is an extension function and returns `bigint`.
- `cyanaudit.fn_get_or_create_audit_field(in_table_schema varchar, in_table_name varchar, in_column_name varchar)` is an extension function and returns `integer`.
- `cyanaudit.fn_get_or_create_audit_transaction_type(in_label varchar)` is an extension function and returns `integer`.

### Requirements and Caveats

- The catalog records version `2.2.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
