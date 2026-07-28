## Usage

Sources:

- [Official upstream README](https://github.com/percona/percona_pg_telemetry/blob/6c8c778d7eca74189770bbad5919d5eabc8cb99e/README.md)
- [Official extension control file (percona_pg_telemetry.control)](https://github.com/percona/percona_pg_telemetry/blob/6c8c778d7eca74189770bbad5919d5eabc8cb99e/percona_pg_telemetry.control)
- [Official extension SQL (percona_pg_telemetry--1.0.sql)](https://github.com/percona/percona_pg_telemetry/blob/6c8c778d7eca74189770bbad5919d5eabc8cb99e/percona_pg_telemetry--1.0.sql)

`percona_pg_telemetry` — > [!CAUTION] > This extension has been deprecated and replaced with a backwards compatibility stub. > > No telemetry data will be gathered, it will not be maintained going forward and it should not be used in new deployments. Use it when collecting or interpreting the corresponding PostgreSQL statistics. The reviewed upstream material marks this capability deprecated.

### Core Workflow

```sql
CREATE EXTENSION percona_pg_telemetry;

ALTER SYSTEM SET percona_pg_telemetry.enabled = 0;
    SELECT pg_reload_conf();
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `percona_pg_telemetry_status(OUT latest_output_filename text, OUT pt_enabled boolean)` is an extension function and returns `record`.
- `percona_pg_telemetry_version()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.2`.
- The control file marks the extension as relocatable.
- Upstream material contains an explicit deprecation boundary.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
