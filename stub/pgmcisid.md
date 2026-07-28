## Usage

Sources:

- [Official extension control file (pgmcisid.control)](https://github.com/morkato/mcisid/blob/5b67978b152806ffc6943b7d351f4f36a1a2d3ba/pgmcisid/pgmcisid.control)
- [Official extension SQL (pgmcisid--1.0.sql)](https://github.com/morkato/mcisid/blob/5b67978b152806ffc6943b7d351f4f36a1a2d3ba/pgmcisid/pgmcisid--1.0.sql)

`pgmcisid` — MCIS ID V1 generator. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgmcisid;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `mcisidv1_created_at(mcisidv1)` is an extension function and returns `TIMESTAMP`.
- `mcisidv1_gen` is an extension function.
- `mcisidv1_get_epoch()` is an extension function and returns `TIMESTAMP`.
- `mcisidv1_get_instant(mcisidv1)` is an extension function and returns `BIGINT`.
- `mcisidv1_instant_sequence(mcisidv1)` is an extension function and returns `BIGINT`.
- `mcisidv1_origin_model(mcisidv1)` is an extension function and returns `SMALLINT`.
- `mcisidv1_type_cmp(mcisidv1, mcisidv1)` is an extension function and returns `INTEGER`.
- `mcisidv1_type_eq(mcisidv1, mcisidv1)` is an extension function and returns `BOOLEAN`.
- `mcisidv1_type_gt(mcisidv1, mcisidv1)` is an extension function and returns `BOOLEAN`.
- `mcisidv1_type_in(CSTRING)` is an extension function and returns `mcisidv1`.
- `mcisidv1_type_lt(mcisidv1, mcisidv1)` is an extension function and returns `BOOLEAN`.
- `mcisidv1_type_out(mcisidv1)` is an extension function and returns `CSTRING`.
- `text_to_mcisidv1(TEXT)` is an extension function and returns `mcisidv1`.
- `mcisidv1` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
