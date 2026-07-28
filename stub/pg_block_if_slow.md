## Usage

Sources:

- [Official upstream README](https://github.com/eshwar-333/pg_block_if_slow/blob/2313c00bb439332fe3d54321037c13d9522f53ee/README.md)
- [Official extension control file (pg_block_if_slow.control)](https://github.com/eshwar-333/pg_block_if_slow/blob/2313c00bb439332fe3d54321037c13d9522f53ee/pg_block_if_slow.control)
- [Official extension SQL (pg_block_if_slow--1.0.sql)](https://github.com/eshwar-333/pg_block_if_slow/blob/2313c00bb439332fe3d54321037c13d9522f53ee/pg_block_if_slow--1.0.sql)

`pg_block_if_slow` — When enabled, the pg_block_if_slow extension prevents the execution of any query whose **estimated cost exceeds a defined threshold** — even before it runs! Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_block_if_slow;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
