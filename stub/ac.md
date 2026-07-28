## Usage

Sources:

- [Official upstream README](https://github.com/darh/pgxs-acl/blob/facb10146c6a23cf2f3c9580c7d408677c5a8641/README.md)
- [Official extension control file (ac.control)](https://github.com/darh/pgxs-acl/blob/facb10146c6a23cf2f3c9580c7d408677c5a8641/src/ac.control)
- [Official extension SQL (ac--0.0.1.sql)](https://github.com/darh/pgxs-acl/blob/facb10146c6a23cf2f3c9580c7d408677c5a8641/src/ac--0.0.1.sql)

`ac` — Access Control Postgres extension (with dev environment). Use it when implementing the corresponding security, audit, or access-control workflow. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION ac;

SELECT ac_policy(
    'user:1',
    ARRAY['read', 'write'],
    ARRAY['delete']
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ac_check(ac_operation, ac_subject, ac_policy[])` is an extension function and returns `BOOLEAN`.
- `ac_check(op ac_operation , bindings ac_subject[] , list ac_policy[])` is an extension function and returns `BOOLEAN`.
- `ac_list_cleanup(dirty ac_policy[])` is an extension function and returns `ac_policy[]`.
- `ac_policy(ac_subject, ac_operation)` is an extension function and returns `ac_policy`.
- `ac_policy(ac_subject, ac_operation, ac_operation)` is an extension function and returns `ac_policy`.
- `ac_policy(ac_subject, ac_operation[])` is an extension function and returns `ac_policy`.
- `ac_policy(ac_subject, ac_operation[], ac_operation)` is an extension function and returns `ac_policy`.
- `ac_policy(ac_subject, ac_operation[], ac_operation[])` is an extension function and returns `ac_policy`.
- `ac_policy` is an extension-defined type.
- `ac_operation` is an extension-defined domain.
- `ac_subject` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Upstream explicitly says the project is not production-ready.
- Upstream describes the project as a work in progress.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
