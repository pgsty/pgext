## Usage

Sources:

- [Official Advanced Storage Pack documentation](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/)
- [Official configuration guide](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/configuring/)
- [Official usage guide](https://www.enterprisedb.com/docs/pg_extensions/advanced_storage_pack/using/)

`bluefin` is an append-only EDB table access method that compacts tuple headers and delta-compresses adjacent data, targeting time-series and monitoring workloads.

### Enablement

Bluefin requires PostgreSQL-family version 15 or later. Install Advanced Storage Pack 1.7.0, preload it, restart, and create the extension:

```ini
shared_preload_libraries = 'bluefin'
```

```sql
CREATE EXTENSION bluefin;
```

### Use Partitioned Append-Only Storage

```sql
CREATE TABLE truck_logs (
  ts        timestamptz,
  truck_id  integer,
  latitude  float8,
  longitude float8
) PARTITION BY RANGE (ts);

CREATE TABLE truck_logs_2026_08
PARTITION OF truck_logs
FOR VALUES FROM ('2026-08-01') TO ('2026-09-01')
USING bluefin;
```

Bluefin does not permit `UPDATE` or `DELETE`. Retention should use partition detach/drop rather than row deletion.

### Estimate and Operate

```sql
SELECT *
FROM bluefin.estimate_compression('existing_heap', 10.0);
```

Compression estimates sample heap pages and do not move data. Before migration, validate insert ordering, partition rotation, index support, crash recovery, backup/restore, and replica behavior. The append-only contract is a schema-design boundary; applications that require corrections must write compensating rows or use a different TAM.

