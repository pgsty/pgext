## Usage

Sources:

- [README.md](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/README.md)
- [docs/API.md](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/docs/API.md)
- [Cargo.toml](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/Cargo.toml)
- [sql/geo_vec--0.1.0.sql](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/sql/geo_vec--0.1.0.sql)
- [geo_vec.control](https://github.com/decision-labs/pg_geo_vec/blob/199a2412ee75e69897be68e48d375c828b849345/geo_vec.control)

`geo_vec` combines vector nearest-neighbor search and geographic bounding-box filtering in one index. This source preview forks the DiskANN implementation from pgvectorscale and declares a separate extension identity.

### Core Workflow

```sql
CREATE EXTENSION vector;
CREATE EXTENSION postgis;
CREATE EXTENSION geo_vec;
CREATE TABLE places (id bigint PRIMARY KEY, embedding vector(3), geom geometry(Point,4326));
CREATE INDEX ON places USING geo_vec (embedding vector_cosine_ops, geom);
SELECT id FROM places
WHERE geom && ST_MakeEnvelope(-122.5,37.7,-122.4,37.8,4326)
ORDER BY embedding <=> '[0.1,0.2,0.3]'::vector LIMIT 20;
```

### Operational Boundaries

Requires `vector` and `postgis`, and installation requires a superuser. The README supports PostgreSQL 17 and later; the reviewed manifest includes 17 and 18 features. Older feature flags alone are not a compatibility guarantee. No preload or restart is documented.

The access method supports cosine, L2 and inner-product vector operator classes, geometry bounding boxes and optional smallint-array labels. `geo_vec.query_search_list_size` and `geo_vec.query_rescore` trade query work for recall. `geo_vec.spatial_brute_force_threshold` selects exhaustive scanning for small spatial candidate sets; larger regions use approximate graph search. Measure recall on the target workload and use matching geometry coordinate systems.

Index builds consume memory and storage. Native builds require suitable AVX2/FMA or NEON support. The catalog records source availability, not a tested Pigsty package matrix.
