## Usage

Sources:

- [Official Advanced Storage Pack documentation](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/configuring/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/using/)

`autocluster` is an EDB table access method that places rows with the same clustering key near one another as new data is inserted.

### Enablement

Install Advanced Storage Pack 1.7.0, preload `autocluster`, restart, and create the extension:

```ini
shared_preload_libraries = 'autocluster'
```

```sql
CREATE EXTENSION autocluster;
```

### Create a Clustered Table

```sql
CREATE TABLE iot (
  thermostat_id        bigint NOT NULL,
  recordtime           time NOT NULL,
  measured_temperature float4
) USING autocluster;

CREATE INDEX ON iot (thermostat_id);

SELECT autocluster.autocluster(
  rel := 'iot'::regclass,
  cols := '{1}',
  max_objects := 10000
);
```

`cols` identifies clustering-key attribute numbers. Insertions with the same key are directed toward nearby blocks, reducing page reads for key-local access patterns.

### Operational Boundaries

Autocluster is a physical table layout, not a query hint. Choose keys from stable access patterns, validate write concurrency and space behavior, and test backup, restore, logical replication, and major upgrades with the exact EDB package. Preloading is cluster-wide, while `CREATE EXTENSION` and TAM tables are database-local.

