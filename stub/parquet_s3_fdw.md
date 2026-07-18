## Usage

Sources:

- [Official extension control file](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/parquet_s3_fdw.control)
- [Official upstream documentation](https://github.com/pgspider/parquet_s3_fdw/blob/cf4384dba7d1f1b6252d5357311df17f5eeb1775/README.md)

`parquet_s3_fdw` — Foreign data wrapper for accessing Parquet files on local filesystems and Amazon S3 from PostgreSQL.

The reviewed catalog snapshot records version `0.3`, kind `standard`, and implementation language `C++`.
The curated compatibility set is `13,14,15,16,17`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "parquet_s3_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'parquet_s3_fdw';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
