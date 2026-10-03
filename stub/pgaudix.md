## Usage

Sources:

- [README.md](https://github.com/sipesistemas/pgaudix/blob/b552a8afe614e403aec7e8784806d344782969e3/README.md)
- [pgaudix.control](https://github.com/sipesistemas/pgaudix/blob/b552a8afe614e403aec7e8784806d344782969e3/pgaudix.control)
- [pgaudix--1.0.0.sql](https://github.com/sipesistemas/pgaudix/blob/b552a8afe614e403aec7e8784806d344782969e3/pgaudix--1.0.0.sql)

`pgaudix` 1.0.0 mirrors ordinary table rows into audit tables and follows supported DDL changes. This release explicitly refuses partitioned tables, partitions and inherited tables.

### Core Workflow

```sql
CREATE EXTENSION pgaudix;
CREATE TABLE audit_example (id bigint PRIMARY KEY, amount numeric);
SELECT pgaudix.enable('audit_example'::regclass);
INSERT INTO audit_example VALUES (1, 10);
UPDATE audit_example SET amount = 20 WHERE id = 1;
SELECT audit_operation, id, amount FROM audit_example_audit ORDER BY audit_id;
SELECT * FROM pgaudix.status();
```

### Privileges and Identity

The upstream requirement is PostgreSQL 16+. Installation needs a superuser; no preload is required. Functions are not executable by `PUBLIC`. Grant schema usage and specific execution rights to table owners that need `pgaudix.enable`, `pgaudix.disable` and `pgaudix.status`; grant audit-table read access separately.

`pgaudix.app_user` and `pgaudix.app_user_ip` can record application-provided identity within a transaction. They are untrusted application declarations, distinct from the authenticated database role and client address.

### Retention and Upgrade

`pgaudix.disable(table)` stops auditing while retaining history; passing `drop_data := true` removes the audit table. Dropping a source column or table can remove its corresponding history, so disable auditing first if it must be preserved. Updates record the new row; truncation records an operation without individual deleted rows. There is no automatic retention policy.

The reviewed release includes only the 1.0.0 install script, not a 0.2.0 upgrade path. Do not assume `ALTER EXTENSION` performs a lossless migration: preserve existing audit data and plan the transition explicitly. PostgreSQL extension dependencies and mirrored table layout must be restored with the extension library available. No explicit upstream license was found.
