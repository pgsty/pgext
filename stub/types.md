## Usage

Sources:

- [Official upstream README](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/readme.md)
- [Official extension control file (types.control)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/types/types.control)
- [Official extension SQL (types--0.0.1.sql)](https://github.com/pyramation/pg-utils/blob/a35cf5f431e09cd222085e2f24aeb308dde4d0e3/packages/types/sql/types--0.0.1.sql)

`types` — Reusable domains for attachments, email addresses, hostnames, images, uploads, and URLs. Use it when application data needs this type, domain, or its operators. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION types;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `attachment` is an extension-defined domain.
- `email` is an extension-defined domain.
- `hostname` is an extension-defined domain.
- `image` is an extension-defined domain.
- `upload` is an extension-defined domain.
- `url` is an extension-defined domain.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `plpgsql`, `uuid-ossp`, `citext`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
