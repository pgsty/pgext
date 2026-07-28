## Usage

Sources:

- [Official upstream README](https://github.com/rajeshkumarblr/pg_bloom/blob/00dc31b46450d2032cc22069e4c8b83f20285f73/README.md)
- [Official extension control file (pg_bloom.control)](https://github.com/rajeshkumarblr/pg_bloom/blob/00dc31b46450d2032cc22069e4c8b83f20285f73/pg_bloom.control)
- [Official extension SQL (pg_bloom--0.0.1.sql)](https://github.com/rajeshkumarblr/pg_bloom/blob/00dc31b46450d2032cc22069e4c8b83f20285f73/pg_bloom--0.0.1.sql)

`pg_bloom` — A PostgreSQL extension implementing a Bloom filter data structure. Use it when application data needs this type, domain, or its operators. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_bloom;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `bloom_contains(bloom, text)` is an extension function and returns `boolean`.
- `bloom_in(cstring)` is an extension function and returns `bloom`.
- `bloom_out(bloom)` is an extension function and returns `cstring`.
- `bloom` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
