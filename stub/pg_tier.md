## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/README.md)
- [Extension control file](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/pg_tier.control)
- [Rust implementation](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/src/lib.rs)
- [Cargo metadata](https://github.com/Pandry/pg_tier/blob/1ffdbfd20823a99c7e0cff004e05060309848242/Cargo.toml)

`pg_tier` moves a local table to S3-compatible storage through its required `parquet_s3_fdw` dependency, then exposes the data under the original table name as a foreign table.

```sql
CREATE EXTENSION pg_tier CASCADE;
SELECT tier.set_tier_config(
  'archive-bucket', 'ACCESS_KEY', 'SECRET_KEY', 'REGION'
);
SELECT tier.enable_tiering('archive_source');
SELECT tier.execute_s3_tiering('archive_source');
```

The credential strings above are placeholders. Do not place real cloud keys in query text, logs, shell history, or broadly readable catalogs. Use a dedicated least-privilege bucket identity and verify the dependency's credential and TLS behavior.

Upstream explicitly labels version 0.0.5 beta and unsuitable for production because bugs and concurrency issues may exist. Tiering changes the table's storage and object identity and can expose partial-failure boundaries between PostgreSQL and S3. Test row counts, constraints, indexes, permissions, concurrent writes, rollback, backup, restore, and recovery. Keep an independently verified copy until remote data and replacement foreign table are proven complete.
