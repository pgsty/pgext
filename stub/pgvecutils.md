## Usage

Sources:

- [Official upstream README](https://github.com/theshubhendra/pgvecutils/blob/131cddb0dee83c3aa18be6726b7f46c845dca7f6/README.md)
- [Official extension control file (pgvecutils.control)](https://github.com/theshubhendra/pgvecutils/blob/131cddb0dee83c3aa18be6726b7f46c845dca7f6/pgvecutils.control)
- [Official extension SQL (pgvecutils--0.0.1.sql)](https://github.com/theshubhendra/pgvecutils/blob/131cddb0dee83c3aa18be6726b7f46c845dca7f6/pgvecutils--0.0.1.sql)

`pgvecutils` — This project is a PostgreSQL extension that provides vector manipulation utilities. Use it for the corresponding vector, model, or retrieval workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pgvecutils;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `vector_rand(int)` is an extension function and returns `vector`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- Install the confirmed extension dependencies first: `vector`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
