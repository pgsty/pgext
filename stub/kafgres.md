## Usage

Sources:

- [extension/sql/kafgres--0.2.0--0.3.0.sql](https://github.com/RayElg/kafgres/blob/0f343916988ad0b4b99d3427bd0949e69b5686b3/extension/sql/kafgres--0.2.0--0.3.0.sql)
- [.github/workflows/ci.yml](https://github.com/RayElg/kafgres/blob/0f343916988ad0b4b99d3427bd0949e69b5686b3/.github/workflows/ci.yml)
- [README 0.3.2](https://github.com/RayElg/kafgres/blob/0f343916988ad0b4b99d3427bd0949e69b5686b3/README.md)
- [Configuration 0.3.2](https://github.com/RayElg/kafgres/blob/0f343916988ad0b4b99d3427bd0949e69b5686b3/docs/configuration.md)
- [0.3.2 release](https://github.com/RayElg/kafgres/releases/tag/0.3.2)
- [0.3.1 to 0.3.2 SQL](https://github.com/RayElg/kafgres/blob/0.3.2/extension/sql/kafgres--0.3.1--0.3.2.sql)
- [0.3.1 to 0.3.2 changes](https://github.com/RayElg/kafgres/compare/0.3.1...0.3.2)

`kafgres` embeds a Kafka protocol broker in PostgreSQL. Upstream release 0.3.2 builds and tests PostgreSQL 13–18. It needs superuser installation, shared preload and a restart. Its license is Elastic License 2.0.

### Enable the broker

```ini
shared_preload_libraries = 'kafgres'
kafgres.database = 'postgres'
kafgres.bind_host = '127.0.0.1'
kafgres.advertised_host = '127.0.0.1'
kafgres.port = 9092
```

```sql
CREATE EXTENSION kafgres;
SELECT kafgres_create_topic('demo', 1);
BEGIN;
SELECT kafgres_produce('demo', 'key', 'value');
COMMIT;
SELECT * FROM kafgres_partition_offsets('demo');
```

Kafka clients connect to the configured broker port. SQL production participates in the caller's transaction. Configure TLS, authentication and ACLs before exposing the listener beyond a trusted local environment.

### Storage and CDC

`kafgres.storage_engine` defaults to segment; its log uses separate files and requires the extension's replication and archive procedures. Configure `kafgres.segment_archive_command` and monitor `kafgres_archive_status()` before relying on segment retention and recovery. Ordinary PostgreSQL WAL/PITR does not cover the entire segment log. The table engine keeps its log in PostgreSQL tables; changing engines does not migrate existing records.

CDC additionally requires `wal_level = logical`; some PostgreSQL builds also require an `output_plugin_libraries` allowlist. Version 0.3.0 supports SQL CDC mappings with projection and filtering. Review the mapping and recovery procedures before deployment; the tagged CI matrix and release artifacts cover PostgreSQL 13–18.

### Durability Settings

Version 0.3.0 defaults `kafgres.fsync_before_ack` to on and `kafgres.relaxed_produce_commit` to off. Relaxing the first can lose acknowledged segment records on a power failure; relaxing the second can lose the newest idempotent-producer state after a crash and permit duplicates after retries. These settings have narrower scope than transactional SQL production and do not apply uniformly to the table engine. Preserve the strict defaults until the durability tradeoff is deliberate.

### 0.3.0 Upgrade

Install the matching 0.3.0 library, restart, then run `ALTER EXTENSION kafgres UPDATE` in the broker database. The SQL update restricts administrative-function grants and pins the producer search path. The first segment-engine start removes old-format timestamp indexes from 0.2.0; old segments remain readable by scanning, while new segments receive current indexes.

### 0.3.1 CDC Installation Boundary

CDC needs wal_level = logical. On PostgreSQL builds that enforce output-plugin allowlists, permit kafgres in output_plugin_libraries and reload configuration if the broker reports that it may not be used as an output plugin. Library preloading still needs a restart. Follow the upstream 0.3.1 upgrade notes for existing transactional segment records; copying the library alone is not a storage migration.

### 0.3.2 Fixes and Upgrade

0.3.2 fixes visibility during concurrent SQL production, compacted-topic handling, CDC draining and fencing of stale producers. Install the matching library and restart PostgreSQL, then run this in the broker database:

```sql
ALTER EXTENSION kafgres UPDATE TO '0.3.2';
```

The 0.3.1-to-0.3.2 SQL script advances the extension version without changing SQL objects. When upgrading from 0.3.0, account for old transaction markers: Sarama cannot decode markers lacking an EndTransactionMarker value, and they remain until retention removes their segments. Follow the upstream guidance for the consumer start position or topic recreation.
