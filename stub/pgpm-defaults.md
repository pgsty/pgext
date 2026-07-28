## Usage

Sources:

- [Official upstream README](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/defaults/README.md)
- [Official extension control file (pgpm-defaults.control)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/defaults/pgpm-defaults.control)
- [Official extension SQL (pgpm-defaults--0.15.5.sql)](https://github.com/constructive-io/pgpm-modules/blob/45d62108f508af564b72bd6512e86a5ec7cdc923/packages/defaults/sql/pgpm-defaults--0.15.5.sql)

`pgpm-defaults` — @pgpm/defaults establishes a secure baseline configuration for PostgreSQL databases by revoking default public access. This package implements the principle of least privilege by removing PostgreSQL's default permissive settings and requiring explicit permission grants. Use it when implementing the corresponding security, audit, or access-control workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION "pgpm-defaults";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.15.5`.
- Install the confirmed extension dependencies first: `plpgsql`, `pgpm-verify`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
