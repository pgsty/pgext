## Usage

Sources:

- [Official pg_lake README](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/README.md)
- [Version 3.4 control file](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/pg_lake/pg_lake.control)
- [Official build and startup guide](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/building-from-source.md)
- [Official project documentation index](https://github.com/Snowflake-Labs/pg_lake/blob/44134cc33fb152716e10752d0a345c6e1acb8725/docs/README.md)
- [DuckDB secrets manager](https://duckdb.org/docs/stable/configuration/secrets_manager.html)

`pg_lake` is the top-level extension for Snowflake's PostgreSQL lakehouse stack. It installs the table, Iceberg, copy, query-engine, extension-base, and map components needed to query object-store files and create transactional Iceberg tables. The PostgreSQL extensions orchestrate planning and transactions while a separate local `pgduck_server` process executes vectorized work with DuckDB.

### Start the Packaged Stack

Version `3.4` supports PostgreSQL 16 through 18. The PIGSTY RPM and DEB packages install the extension files and a versioned `pgduck_server` binary, but they do not currently install or auto-start a `systemd` service. Running `CREATE EXTENSION` does not start `pgduck_server` either.

Add `pg_extension_base` to `shared_preload_libraries` and restart PostgreSQL:

```conf
shared_preload_libraries = 'pg_extension_base'
```

`pgduck_server` listens on `/tmp/.s.PGSQL.5332` with mode `0770` by default. Run it as the PostgreSQL operating-system user so PostgreSQL can access the socket. Do not start it as an unrelated login user with the bare command.

```shell
# Debian/Ubuntu with PostgreSQL 18; use 16 or 17 as appropriate.
PG_LAKE_SERVER=/usr/lib/postgresql/18/bin/pgduck_server
# RHEL-compatible systems use /usr/pgsql-18/bin/pgduck_server.

sudo install -d -o postgres -g postgres -m 0700 \
  /var/lib/pg_lake /var/lib/pg_lake/extensions
sudo install -d -o postgres -g postgres -m 0750 /var/cache/pg_lake
sudo -u postgres -H "$PG_LAKE_SERVER" \
  --duckdb_database_file_path /var/lib/pg_lake/pgduck_server.db \
  --extensions_dir /var/lib/pg_lake/extensions \
  --cache_dir /var/cache/pg_lake
```

This command runs in the foreground and must remain running. Use a service supervisor for production. If you use a dedicated service account instead, make it a member of the `postgres` group and start the server with `--unix_socket_group postgres --unix_socket_permissions 0770`.

In another terminal, verify the query engine before creating the extensions:

```shell
sudo -u postgres psql -X \
  "host=/tmp port=5332 dbname=postgres connect_timeout=2" \
  -c 'SELECT version();'
```

Then create the complete dependency tree in the target database:

```sql
CREATE EXTENSION pg_lake CASCADE;
SELECT lake.version();
```

### Configure Object-Store Access

Object-store credentials are resolved by `pgduck_server`, not by the PostgreSQL backend. AWS and GCP can use their normal credential chains. For a local S3-compatible endpoint such as MinIO, first create the bucket, connect directly to `pgduck_server`, and create a persistent DuckDB secret:

```shell
sudo -u postgres psql -X -h /tmp -p 5332 -d postgres
```

```sql
CREATE PERSISTENT SECRET pglake_object_store (
    TYPE S3,
    KEY_ID 'access-key',
    SECRET 'secret-key',
    REGION 'us-east-1',
    ENDPOINT 'minio.example.com:9000',
    SCOPE 's3://analytics-bucket',
    URL_STYLE 'path',
    USE_SSL false
);
```

Connect to PostgreSQL, then choose the managed Iceberg location in the same session that creates the table:

```sql
SET pg_lake_iceberg.default_location_prefix =
    's3://analytics-bucket/warehouse';
```

### Core Workflows

Create and modify a transactional Iceberg table:

```sql
CREATE TABLE measurements (
    station_name text NOT NULL,
    measured_at timestamptz NOT NULL,
    value double precision
) USING iceberg;

INSERT INTO measurements VALUES
    ('Istanbul', now(), 18.5),
    ('Haarlem', now(), 9.3);
```

Import or export Parquet, CSV, or newline-delimited JSON through `COPY`:

```sql
COPY (SELECT * FROM measurements)
TO 's3://analytics-bucket/export/measurements.parquet';

COPY measurements
FROM 's3://analytics-bucket/import/measurements.parquet';
```

Query files without loading them into PostgreSQL:

```sql
CREATE FOREIGN TABLE external_events ()
SERVER pg_lake
OPTIONS (path 's3://analytics-bucket/events/*.parquet');

SELECT count(*) FROM external_events;
```

### Component Index

- `pg_lake`: meta-extension and `lake.version()`.
- `pg_lake_table`: data-lake FDW, Iceberg table syntax, file utilities, and table catalogs.
- `pg_lake_iceberg`: Iceberg metadata, snapshots, manifests, and catalog integration.
- `pg_lake_copy`: `COPY` interception for object-store files and lake formats.
- `pg_lake_engine`: shared query rewrite, type conversion, cleanup, and `pgduck_server` client layer.
- `pg_extension_base`: preload and lifecycle-worker infrastructure.
- `pg_map`: generated PostgreSQL map types used for nested lake data.

### Operational Caveats

- `pgduck_server` is required on every PostgreSQL host that can execute lake queries. Keep it supervised and verify its local socket before serving traffic.
- The default socket mode is `0770`; its owner and group come from the account that starts `pgduck_server`. A mismatched service user causes `ERROR: could not start query engine`.
- S3 and compatible credentials are resolved by the DuckDB secrets/credential chain. Grant only the bucket permissions required by the workload.
- The first start may download the DuckDB spatial extension. Ensure the service account has the required network access and writable state/cache directories.
- `CREATE PERSISTENT SECRET` survives server restarts, but DuckDB stores it unencrypted under `~/.duckdb/stored_secrets`. Keep the service account and its home directory stable, restrict permissions, and protect those files as credentials.
- The default memory limit is 80 percent of system memory. Set `--memory_limit` explicitly when PostgreSQL and `pgduck_server` share a production host.
- Iceberg writes create Parquet files per statement. Batch inserts and run regular `VACUUM` to avoid many small files.
- The PostgreSQL extensions, `pgduck_server`, object-store data, and Iceberg catalog form one deployment unit. Back up and upgrade them as separate evidence layers; creating the extension alone does not prove the external services are usable.
