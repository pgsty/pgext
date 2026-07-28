## Usage

Sources:

- [Official upstream README](https://github.com/thebf/pgx-s3sign/blob/90a0d59f9655470df8c3a6700b94ed711207bf4e/README.md)
- [Official extension control file (pgx-s3sign.control)](https://github.com/thebf/pgx-s3sign/blob/90a0d59f9655470df8c3a6700b94ed711207bf4e/pgx-s3sign.control)
- [Official implementation source](https://github.com/thebf/pgx-s3sign/blob/90a0d59f9655470df8c3a6700b94ed711207bf4e/src/lib.rs)

`pgx-s3sign` — Sign S3 requests, fast: Created by pgx. Use it when implementing the corresponding security, audit, or access-control workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "pgx-s3sign";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgx_s3sign_pre_get` is an extension function.
- `pgx_s3sign_pre_put` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
