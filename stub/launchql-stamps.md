## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/stamps/README.md)
- [Official extension control file (launchql-stamps.control)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/stamps/launchql-stamps.control)
- [Official extension SQL (launchql-stamps--0.4.5.sql)](https://github.com/constructive-io/utils/blob/949bad998f1ac6eaf873463ce2f788765068bc1a/packages/stamps/sql/launchql-stamps--0.4.5.sql)

`launchql-stamps` — PostgreSQL extension providing trigger functions for automatically adding timestamps and user tracking to database tables. This extension simplifies the implementation of audit trails by automatically recording creation and update timestamps, as well as the users who performed these actions. Use it when SQL needs these specialized functions or aggregates. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "launchql-stamps";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `stamps.peoplestamps()` is an extension function and returns `trigger`.
- `stamps.timestamps()` is an extension function and returns `trigger`.
- `stamps` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.4.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `launchql-jwt-claims`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
