## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/README.md)
- [Extension control file](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/pg_tier.control)
- [SQL API and tiering workflow](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/sql/pg_tier.sql)
- [Rust implementation](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/src/lib.rs)
- [Cargo metadata](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/Cargo.toml)

`pg_tier` moves a local table to S3-compatible storage through its required `parquet_s3_fdw` dependency, then exposes the data under the original table name as a foreign table.

### Core Workflow

```sql
CREATE EXTENSION pg_tier CASCADE;
SELECT tier.set_tier_config(
  :'bucket_name', :'access_key', :'secret_key', :'region'
);
SELECT tier.enable_tiering('archive_source');
SELECT tier.execute_s3_tiering('archive_source');
```

The example uses psql variables so no real credential is hard-coded in the document. Variable expansion still sends values to PostgreSQL, where statement logging can capture them; prefer client protocol parameters and appropriate logging controls.

### Safety and Caveats

- `tier.set_tier_config` stores the bucket, access key, secret key, and region as text in `tier.server_credential`, and creates an FDW user mapping `FOR public`. Restrict extension tables and management functions, protect backups, and use a dedicated least-privilege bucket identity.
- The setup function creates a fixed server named `pg_tier_s3_srv`. Repeated configuration or an existing object with that name can fail; test lifecycle and recovery procedures before operating it.
- Enabling tiering renames the local table, creates a foreign table under the original name, and records both in extension catalogs. Execution writes the remote data and normally truncates the renamed local table.
- Verify the dependency's credential, endpoint, and TLS behavior. The reviewed build targets PostgreSQL 14 through 16 and has no PostgreSQL 17 or later feature.

Upstream explicitly labels version 0.0.5 beta and unsuitable for production because bugs and concurrency issues may exist. Tiering changes the table's storage and object identity and can expose partial-failure boundaries between PostgreSQL and S3. Test row counts, constraints, indexes, permissions, concurrent writes, rollback, backup, restore, and recovery. Keep an independently verified copy until remote data and replacement foreign table are proven complete.
