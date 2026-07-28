## Usage

Sources:

- [Official upstream README](https://github.com/sasasu/pgts/blob/95230b6f7912c8fa94b6994c6b71d83d5c3644b6/README.md)
- [Official extension control file (pgts.control)](https://github.com/sasasu/pgts/blob/95230b6f7912c8fa94b6994c6b71d83d5c3644b6/src/pgts.control)
- [Official extension SQL (pgts--0.1.0.sql)](https://github.com/sasasu/pgts/blob/95230b6f7912c8fa94b6994c6b71d83d5c3644b6/src/pgts--0.1.0.sql)

`pgts` — time encode for PostgreSQL. Use it for the corresponding scheduling, temporal, or time-series workflow. Upstream describes it as a proof of concept.

### Core Workflow

```sql
CREATE EXTENSION pgts;

select
  hostname,
  unnest(    ts.timestamp_decode(ctime)   )                 as ctime
from
  x
group by
  hostname,
  ctime
order by
  ctime;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ts.f8_decode(v bytea)` is an extension function and returns `table`.
- `ts.f8_encode(v double precision[])` is an extension function and returns `bytea`.
- `ts.timestamp_decode(v bytea)` is an extension function and returns `timestamp[]`.
- `ts.timestamp_encode(v timestamp[])` is an extension function and returns `bytea`.
- `ts.u8_decode(v bytea)` is an extension function and returns `bigint[]`.
- `ts.u8_encode(v bigint[])` is an extension function and returns `bytea`.
- `ts` is a schema created by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
