## Usage

Sources:

- [Official upstream README](https://github.com/samovers/ofarm2/blob/aa91d6097384838e8ba20efa0dbff41e540364da/deployment/postgresql/ofarm_ed25519/README.md)
- [Official extension control file (ofarm_ed25519.control)](https://github.com/samovers/ofarm2/blob/aa91d6097384838e8ba20efa0dbff41e540364da/deployment/postgresql/ofarm_ed25519/ofarm_ed25519.control)
- [Official extension SQL (ofarm_ed25519--1.0.sql)](https://github.com/samovers/ofarm2/blob/aa91d6097384838e8ba20efa0dbff41e540364da/deployment/postgresql/ofarm_ed25519/ofarm_ed25519--1.0.sql)

`ofarm_ed25519` — The build has a closed linux/amd64/linux/arm64 input set. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION ofarm_ed25519;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ed25519_verify(public_key pg_catalog.bytea, signed_bytes pg_catalog.bytea, signature pg_catalog.bytea)` is an extension function and returns `pg_catalog`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
