## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tsfunc/pg_tsfunc.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tsfunc/pg_tsfunc--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_tsfunc/meson.build)

`pg_tsfunc` `1.0` provides time-series resampling, downsampling and value codecs. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_tsfunc;
SELECT ts_decode_floats(ts_encode_floats(ARRAY[1.0,2.0,3.0]::float8[]));
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

ts_lttb and ts_minmax downsample; ts_resample aligns observations; ts_asof and ts_asof_join perform time-based lookups. Timestamp/float codecs store a series in bytea rather than providing a table access method. Verify ordering, duplicate handling and max-gap choices for scientific calculations. Array and aggregate operations still consume backend memory; compression ratios depend on data.

This module corresponds to the former upstream name `alohadb_tsfunc`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
