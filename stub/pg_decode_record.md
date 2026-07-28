## Usage

Sources:

- [Official extension control file (pg_decode_record.control)](https://github.com/myzel394/pg_decode_record/blob/d6053f846d8bda3eb487e24e39cbd5ec59e66aeb/pg_decode_record.control)
- [Official extension SQL (pg_decode_record--0.0.1.sql)](https://github.com/myzel394/pg_decode_record/blob/d6053f846d8bda3eb487e24e39cbd5ec59e66aeb/pg_decode_record--0.0.1.sql)

`pg_decode_record` — Hacky extension to extract SQL from WAL INSERT records; Originally from https://github.com/rjuju. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_decode_record;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
