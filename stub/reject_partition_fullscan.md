## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [Official extension control file (reject_partition_fullscan.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/reject_partition_fullscan/reject_partition_fullscan.control)
- [Official extension SQL (reject_partition_fullscan--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/reject_partition_fullscan/reject_partition_fullscan--1.0.sql)

`reject_partition_fullscan` — Reject queries that scan every partition without pruning. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION reject_partition_fullscan;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
