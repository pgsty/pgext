## Usage

Sources:

- [extension/sql/kafgres--0.2.0--0.3.0.sql](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/extension/sql/kafgres--0.2.0--0.3.0.sql)
- [.github/workflows/ci.yml](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/.github/workflows/ci.yml)
- [README 0.3.0](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/README.md)
- [Configuration 0.3.0](https://github.com/RayElg/kafgres/blob/8def1919c65f71217665cdc9b72df4c693811cc2/docs/configuration.md)

- [0.3.0 release](https://github.com/RayElg/kafgres/releases/tag/0.3.0)

`kafgres` embeds a Kafka protocol broker in PostgreSQL. Upstream release 0.3.0 builds and tests PostgreSQL 13–18. It needs superuser installation, shared preload and a restart. Its license is Elastic License 2.0.

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
