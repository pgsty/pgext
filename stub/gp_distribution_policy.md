## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [Official extension control file (gp_distribution_policy.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_distribution_policy/gp_distribution_policy.control)
- [Official extension SQL (gp_distribution_policy--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_distribution_policy/gp_distribution_policy--1.0.sql)

`gp_distribution_policy` — Check table distribution policy in a Greenplum-family cluster. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_distribution_policy;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gp_distribution_policy_table_check(relid regclass)` is an extension function and returns `boolean`.
- `gp_table_distribution_check(regclass)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
