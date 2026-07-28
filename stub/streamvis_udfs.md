## Usage

Sources:

- [Official upstream README](https://github.com/hrbigelow/streamvis/blob/51018a5199d3a26196b96836c69ca8142c7338ba/db/README.md)
- [Official extension control file (streamvis_udfs.control)](https://github.com/hrbigelow/streamvis/blob/51018a5199d3a26196b96836c69ca8142c7338ba/db/udf/streamvis_udfs.control)
- [Official extension SQL (streamvis_udfs--1.0.sql)](https://github.com/hrbigelow/streamvis/blob/51018a5199d3a26196b96836c69ca8142c7338ba/db/udf/streamvis_udfs--1.0.sql)

`streamvis_udfs` — A logging client and server for self-hosted data logging and visualization. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION streamvis_udfs;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `decode_float_enc(e enc_typ)` is an extension function and returns `REAL[]`.
- `decode_int_enc(e enc_typ)` is an extension function and returns `INT[]`.
- `decode_text_enc(e enc_typ)` is an extension function and returns `TEXT[]`.
- `encode_bool_enc(p_vals BOOLEAN[])` is an extension function and returns `enc_typ`.
- `encode_float_enc(p_vals FLOAT[])` is an extension function and returns `enc_typ`.
- `encode_int_enc(p_vals INT[])` is an extension function and returns `enc_typ`.
- `encode_text_enc(p_vals TEXT[])` is an extension function and returns `enc_typ`.
- `window_avg_finalfunc(internal)` is an extension function and returns `enc_typ`.
- `window_avg_sfunc(internal, int, int, enc_typ, enc_typ[])` is an extension function and returns `internal`.
- `window_avg` is an aggregate exposed by the extension.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
