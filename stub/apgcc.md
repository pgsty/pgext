## Usage

Sources:

- [Current AWS Aurora PostgreSQL supported-extension history](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraPostgreSQLReleaseNotes/AuroraPostgreSQL.Extensions.html)

`apgcc` was an AWS-internal extension, not a generally installable open-source extension. The current AWS history records version `1.0` only for Aurora PostgreSQL 9.6.3 and 9.6.6, marks it `NA` in later 9.6 releases, and explicitly says it is no longer supported.

There is no current functional API to document. On an Aurora cluster, this diagnostic query can establish whether the provider still exposes the extension to that particular engine release:

```sql
SELECT name, default_version, installed_version
FROM pg_available_extensions
WHERE name = 'apgcc';
```

An empty result is expected on current supported releases.

### Caveats

- Do not attempt to supply a third-party binary under the same name. AWS identifies `apgcc` as an internal provider extension and publishes no source or SQL interface.
- Catalog version `1.0` is historical metadata, not a current availability promise. Use the engine-specific AWS table and `pg_available_extensions` for live availability.
- Do not confuse this retired extension with `apg_plan_mgmt` or other currently supported Aurora features.
