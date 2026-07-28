## Usage

Sources:

- [Official upstream README](https://api.pgxn.org/src/weighted_stats/weighted_stats-1.0.0/README.md)
- [Official extension control file (weighted_stats.control)](https://api.pgxn.org/src/weighted_stats/weighted_stats-1.0.0/weighted_stats.control)
- [Official extension SQL (weighted_stats.sql)](https://api.pgxn.org/src/weighted_stats/weighted_stats-1.0.0/sql/weighted_stats.sql)

`weighted_stats` — Weighted aggregate functions. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION weighted_stats;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `weighted_mean` is an aggregate exposed by the extension.
- `weighted_stddev_samp` is an aggregate exposed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
