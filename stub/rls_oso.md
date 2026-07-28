## Usage

Sources:

- [Official upstream README](https://github.com/mfashby/rls_oso/blob/67edad478f4d985869d2baee5340fe97fb7cb2f4/README.md)
- [Official extension control file (rls_oso.control)](https://github.com/mfashby/rls_oso/blob/67edad478f4d985869d2baee5340fe97fb7cb2f4/rls_oso.control)
- [Official implementation source](https://github.com/mfashby/rls_oso/blob/67edad478f4d985869d2baee5340fe97fb7cb2f4/src/lib.rs)

`rls_oso` — Plugin to use Oso authorization library in postgres' row level security policies. Use it when implementing the corresponding security, audit, or access-control workflow. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION rls_oso;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `oso_configure_rls` is an extension function.
- `oso_is_allowed` is an extension function.
- `oso_reload()` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
