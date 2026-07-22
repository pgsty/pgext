## Usage

Sources:

- [Official README](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/README.md)
- [Extension SQL](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/hba_fdw--1.0.sql)
- [FDW implementation](https://github.com/fabriziomello/hba_fdw/blob/fcc1cd8eee8f997538b917f095caef5373a83e7d/hba_fdw.c)

hba_fdw is a work-in-progress foreign-data-wrapper skeleton that was intended to manage pg_hba.conf through SQL. The checked implementation neither reads nor changes authentication rules: every scan immediately returns no rows, the validator accepts every option, and no insert, update, or delete callbacks exist.

### Core Workflow

The only completed operation is registering the empty wrapper:

```sql
CREATE EXTENSION hba_fdw;

SELECT fdwname
FROM pg_foreign_data_wrapper
WHERE fdwname = 'hba_fdw';
```

There is no authoritative server-option or foreign-table schema to configure, and creating a foreign table does not make pg_hba.conf queryable.

### Important Objects

- `hba_fdw_handler` returns the minimal scan callback table.
- `hba_fdw_validator` currently performs no validation.
- `hba_fdw` is registered as a foreign data wrapper by the installation SQL.

### Operational Notes

Do not build authentication administration around hba_fdw. It has no parser, reload mechanism, concurrency protection, privilege model, or safe writer, and the README only promises future functionality. Manage pg_hba.conf with supported configuration-management tooling and validate changes with PostgreSQL's own facilities.
