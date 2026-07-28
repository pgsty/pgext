## Usage

Sources:

- [Official upstream README](https://github.com/louisja1/bacon/blob/0610f49a71fc1ec5bafee5c1549725c79bc8e7d0/README.md)
- [Official extension control file (find_bucket_ts.control)](https://github.com/louisja1/bacon/blob/0610f49a71fc1ec5bafee5c1549725c79bc8e7d0/src/find_bucket_ext/find_bucket_ts/find_bucket_ts.control)
- [Official extension SQL (find_bucket_ts--1.0.sql)](https://github.com/louisja1/bacon/blob/0610f49a71fc1ec5bafee5c1549725c79bc8e7d0/src/find_bucket_ext/find_bucket_ts/find_bucket_ts--1.0.sql)

`find_bucket_ts` — Artifacts of BaCon. For VLDB26 reviewers only. Use it for the corresponding scheduling, temporal, or time-series workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION find_bucket_ts;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `find_bucket_ts(val timestamp without time zone, buckets timestamp without time zone[][], need_null boolean)` is an extension function and returns `int`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
