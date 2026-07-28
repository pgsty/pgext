## Usage

Sources:

- [Official upstream README](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/README.md)
- [Official extension control file (gp_toolkit.control)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_toolkit/gp_toolkit.control)
- [Official extension SQL (gp_toolkit--1.0.sql)](https://github.com/apache/cloudberry/blob/83022b58b2f4e62fb8fd2ebf9137b5c29e4fcb5b/gpcontrib/gp_toolkit/gp_toolkit--1.0.sql)

`gp_toolkit` — Administrative views and functions for Greenplum-family databases. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gp_toolkit;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gp_toolkit.get_column_size(ao_oid oid, OUT segment int, OUT attnum int, OUT size bigint, OUT size_uncompressed bigint, OUT compression_ratio numeric)` is an extension function and returns `SETOF`.
- `gp_toolkit.gp_param_setting(varchar)` is an extension function and returns `SETOF`.
- `gp_toolkit.gp_param_settings()` is an extension function and returns `SETOF`.
- `gp_toolkit.gp_skew_coefficient(targetoid oid, OUT skcoid oid, OUT skccoeff numeric)` is an extension function and returns `record`.
- `gp_toolkit.gp_skew_details(oid)` is an extension function and returns `setof`.
- `gp_toolkit.gp_skew_idle_fraction(targetoid oid, OUT sifoid oid, OUT siffraction numeric)` is an extension function and returns `record`.
- `gp_toolkit.session_state_memory_entries_f_on_master()` is an extension function and returns `SETOF`.
- `gp_toolkit.session_state_memory_entries_f_on_segments()` is an extension function and returns `SETOF`.
- `gp_toolkit.gp_param_setting_t` is an extension-defined type.
- `gp_toolkit.gp_skew_analysis_t` is an extension-defined type.
- `gp_toolkit.gp_skew_details_t` is an extension-defined type.
- `gp_toolkit.gp_bloat_diag` is an extension-defined view.
- `gp_toolkit.gp_bloat_expected_pages` is an extension-defined view.
- `gp_toolkit.gp_check_missing_files` is an extension-defined view.

### Requirements and Caveats

- The reviewed control file declares default version `1.6`.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
