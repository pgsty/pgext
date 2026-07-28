## Usage

Sources:

- [Official upstream README](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/docs/README.md)
- [Official extension control file (veil2_minimal_demo.control)](https://github.com/marcmunro/veil2/blob/99d3c931e22475e1abd35c687b9cc08f1111f7ef/veil2_minimal_demo.control)

`veil2_minimal_demo` — Provides a minimal but complete demo database for the veil2 extension. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION veil2_minimal_demo;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.9.3`.
- Install the confirmed extension dependencies first: `veil2`.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
