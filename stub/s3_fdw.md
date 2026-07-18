## Usage

Sources:

- [Official PGXN distribution page](https://pgxn.org/dist/s3_fdw/)

`s3_fdw` — Foreign data wrapper for reading Amazon S3 files through PostgreSQL COPY format.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "s3_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 's3_fdw';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
