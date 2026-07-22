## Usage

Sources:

- [Official README](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/readme.md)
- [Official usage example](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/docs/example.md)
- [Official extension SQL](https://github.com/dariubs/town/blob/3944f2f34bd509588907ec845398e18796939379/sql/town--0.1.0.sql)

`town` is a small PL/pgSQL helper that creates a conventional time-series table with timestamp, tag-array, and JSONB payload columns. It is useful for prototypes that want this exact layout; it is not a time-series engine and does not provide partitioning, retention, compression, rollups, or background maintenance.

### Core Workflow

Install `town`, then call `create_town_table()` with the table name:

```sql
CREATE EXTENSION town;
SELECT create_town_table('metrics');

INSERT INTO metrics(ts, tags, data)
VALUES (now(), '{town,db,timeseries}', '{"val":21324}');

SELECT ts, tags, data->>'val' AS val
FROM metrics
WHERE 'town' = ANY(tags)
  AND ts BETWEEN timestamptz '2026-07-01' AND now();
```

The generated table has `id BIGSERIAL PRIMARY KEY`, `ts TIMESTAMP WITH TIME ZONE`, `tags TEXT[]`, and `data JSONB DEFAULT '{}'::jsonb`. The helper also installs `btree_gist` if needed and creates a GIN index on `tags` plus a GiST index on `ts`.

### Operational Boundaries

Calling `create_town_table()` requires enough privilege to create an extension, table, sequence, and indexes in the current database and schema. The table identifier is quoted with `format('%I', ...)`, but the function accepts only names up to 30 characters and does not accept a schema-qualified name.

Although table creation uses `IF NOT EXISTS`, index creation does not, so repeating the call for an existing table can fail on duplicate index names. Add application-specific constraints, non-null rules, indexes, partitioning, retention, and access control yourself; benchmark the GIN/GiST write cost and JSONB shape against the intended workload.
