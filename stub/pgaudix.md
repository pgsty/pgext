## Usage

Sources:

- [Official upstream README](https://github.com/sipesistemas/pgaudix/blob/95d386c4329e76e78ee67e2478a6660673193960/README.md)
- [Official extension control file (pgaudix.control)](https://github.com/sipesistemas/pgaudix/blob/95d386c4329e76e78ee67e2478a6660673193960/pgaudix.control)
- [Official extension SQL (pgaudix--0.1.0--0.2.0.sql)](https://github.com/sipesistemas/pgaudix/blob/95d386c4329e76e78ee67e2478a6660673193960/pgaudix--0.1.0--0.2.0.sql)

`pgaudix` — A native PostgreSQL extension for automatic table auditing. It mirrors table columns into audit tables and automatically keeps them in sync when the source table structure changes. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgaudix;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgaudix.audit_trigger()` is an extension function and returns `trigger`.
- `pgaudix.ddl_sync()` is an extension function and returns `event_trigger`.
- `pgaudix.disable(target_table regclass, drop_data boolean DEFAULT false)` is an extension function and returns `void`.
- `pgaudix.drop_cleanup()` is an extension function and returns `event_trigger`.
- `pgaudix.enable(target_table regclass)` is an extension function and returns `void`.
- `pgaudix.status()` is an extension function and returns `TABLE`.
- `pgaudix.truncate_trigger()` is an extension function and returns `trigger`.
- `pgaudix.monitored_tables` is a table installed or managed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.2.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
