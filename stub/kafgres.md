## Usage

Sources:

- [README 0.1.0](https://github.com/RayElg/kafgres/blob/0.1.0/README.md)
- [Configuration 0.1.0](https://github.com/RayElg/kafgres/blob/0.1.0/docs/configuration.md)

`kafgres` embeds a Kafka protocol broker in PostgreSQL. Release 0.1.0 is packaged for PostgreSQL 16 using the upstream pgrx 0.16.1 source. It needs superuser installation, shared preload and a restart. Its license is Elastic License 2.0.

### Enable the broker

```conf
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

CDC additionally requires `wal_level = logical` and an appropriate `output_plugin_libraries` allowlist. Review the mapping and recovery procedures before deployment. The package matrix excludes other majors: upstream 0.1.0 decoding code fails to compile on PostgreSQL 17.
