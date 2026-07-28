## Usage

Sources:

- [Official upstream README](https://github.com/evancarroll/pg-generate-up-down-series/blob/f7e5a0f9f1083efcd12dc4bbb906382d5acffc1e/README.md)
- [Official extension control file (generate_up_down_series.control)](https://github.com/evancarroll/pg-generate-up-down-series/blob/f7e5a0f9f1083efcd12dc4bbb906382d5acffc1e/generate_up_down_series.control)
- [Official extension SQL (generate_up_down_series--0.0.1.sql)](https://github.com/evancarroll/pg-generate-up-down-series/blob/f7e5a0f9f1083efcd12dc4bbb906382d5acffc1e/generate_up_down_series--0.0.1.sql)

`generate_up_down_series` — First you'll need to compile. Note on Debian you'll need postgresql-server-dev-all and build-essentials. Then you can install it with. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION generate_up_down_series;

SELECT *
FROM generate_up_down_series_evan(n,m);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `generate_up_down_series_evan(n int4, m int4)` is an extension function and returns `TABLE`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
