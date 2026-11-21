## Usage

Sources:

- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.0/README.md)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.0/pg_parquet.control)
- [Official source](https://github.com/CrunchyData/pg_parquet/blob/v0.6.0/Cargo.toml)
- [Official source](https://github.com/CrunchyData/pg_parquet/releases/tag/v0.6.0)

`pg_parquet` 0.6.0 reads and writes Parquet through PostgreSQL COPY and inspects file metadata. The source targets PostgreSQL 14–19. It uses server-side local files or S3, Azure, GCS and HTTP(S) endpoints; these locations and credentials are separate from a PostgreSQL backup.

### Enable and Copy

Preserve other preload entries, add the library and restart. A superuser creates the non-relocatable extension. The server process must have access to the chosen local path or remote object. Run the example as an administrator with the relevant COPY privileges:

```ini
shared_preload_libraries = 'pg_parquet'
```

```sql
CREATE EXTENSION pg_parquet;
CREATE TABLE parquet_demo (id bigint, label text);
INSERT INTO parquet_demo VALUES (1, 'first'), (2, 'second');
COPY parquet_demo TO '/tmp/parquet_demo.parquet' WITH (format parquet);
CREATE TABLE parquet_copy (LIKE parquet_demo);
COPY parquet_copy FROM '/tmp/parquet_demo.parquet' WITH (format parquet, match_by name);
SELECT * FROM parquet.schema('/tmp/parquet_demo.parquet');
SELECT * FROM parquet.file_metadata('/tmp/parquet_demo.parquet');
```

Use COPY of a query for a filtered export. STDIN, STDOUT and PROGRAM streams require an explicit Parquet format. External output is not rolled back with the SQL transaction: failed or rolled-back exports can leave files or incomplete uploads.

### Schema and Type Handling

```sql
CREATE TABLE inferred_demo () WITH (load_from = '/tmp/parquet_demo.parquet');
SET pg_parquet.max_parallel_workers = 2;
SELECT * FROM pg_stat_progress_copy;
```

The `load_from` table option infers columns and loads rows. Composite types are created in `parquet_structs`; the creator needs CREATE there, which installation grants to PUBLIC. Generated types survive dropping the table. Dropping the extension with CASCADE can remove dependent columns. Maps require `crunchy_map`; unsupported inferred types need an explicitly defined target table.

`match_by name` matches nested composite fields by name; the default is position. Version 0.6.0 improves dictionary/decimal/INT96 and infinite date/timestamp reads, preserves JSON text and rejects embedded NUL strings. COPY TO rejects multidimensional arrays. PostGIS geometry/geography use Parquet geospatial types; curves and other unsupported geometries must be converted or exported explicitly as binary.

### Files, Performance and Credentials

`parquet.metadata`, `parquet.kv_metadata`, `parquet.column_stats` and `parquet.schema` inspect files. The `bloom_filter` and `dictionary` export options select columns; COPY FROM reads all rows and does not use bloom filters to filter them.

`pg_parquet.max_parallel_workers` defaults to zero. Eligible copies can use parallel workers; query exports, temporary tables, triggers/defaults and other cases can keep the serial path. Import workers decode while the leader inserts. Inspect `pg_stat_progress_copy` and size resources for the actual file and table.

`file_size_bytes` writes a folder of files. Nonempty folders are refused unless `overwrite true` is specified, except that a denied listing cannot establish emptiness and retains the older write behavior. Overwrite can remove files from an earlier copy. Concurrent writers are excluded within one PostgreSQL instance; other instances require external coordination. A single-file destination retains its replacement behavior. Plan cleanup for failed exports.

`pg_parquet.profile_name` selects an object-store identity only if the role is allowed by `parquet.user_profile_mapping`. Only a superuser or `parquet_profile_manager` may manage that mapping; privileges are rechecked for each copy. Keep credentials in protected server-side files/environment. `pg_parquet.allow_credentials_in_parameters` defaults off because SQL credentials can appear in logs and activity views. HTTP exports buffer one file in memory and cap it at 10 MiB; object-store parallel multipart requirements vary, including limitations for stores demanding equally sized parts.

### Upgrade

Deploy files matching the server major, follow the preload maintenance procedure and use the supplied SQL update path. Check existing folder-export jobs, inferred composite-type dependencies, credential mappings and type coercion behavior before moving to 0.6.0.
