## Usage

Sources:

- [Official usage documentation](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/doc/osm_fdw.md)
- [Extension SQL](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/sql/osm_fdw.sql)
- [Table helper SQL](https://github.com/vpikulik/postgres_osm_pbf_fdw/blob/d46960460ad3652ebe17dab28b50c0ca650522cd/sql/utils.sql)

`osm_fdw` is a read-only foreign data wrapper for OpenStreetMap PBF files. It exposes OSM nodes, ways, and relations as rows so they can be transformed with SQL before being copied into ordinary PostgreSQL tables or materialized views.

### Core Workflow

```sql
CREATE EXTENSION osm_fdw;
CREATE SERVER osm_files FOREIGN DATA WRAPPER osm_fdw;

SELECT create_osm_table(
    'malta_osm',
    'osm_files',
    '/srv/osm/malta-latest.osm.pbf'
);

SELECT id, type, lat, lon, tags
FROM malta_osm
WHERE type = 'NODE'
  AND tags ? 'amenity';

CREATE MATERIALIZED VIEW malta_osm_data AS
SELECT * FROM malta_osm WITH DATA;
```

`create_osm_table(table_name, server_name, filename)` creates the required foreign-table layout. A manual definition must keep the documented column order and types:

```sql
CREATE FOREIGN TABLE osm_data (
    id bigint,
    type text,
    lat double precision,
    lon double precision,
    tags jsonb,
    refs bigint[],
    members jsonb,
    version integer,
    modified timestamp,
    changeset bigint,
    user_id integer,
    username text,
    visible boolean
)
SERVER osm_files
OPTIONS (filename '/srv/osm/region.osm.pbf');
```

### Row Model

`type` is `NODE`, `WAY`, or `RELATION`. Coordinates are populated for nodes, `refs` contains node identifiers for ways, and `members` contains relation members. `tags` stores the OSM tag map as `jsonb`; the remaining columns expose version and contributor metadata.

### Caveats

The PostgreSQL server process must be able to read the file path. The FDW rereads the PBF file for every query and does not provide indexes over the source, so upstream recommends materializing or copying data and then creating normal PostgreSQL indexes. Treat the file as an external snapshot: refresh derived tables explicitly when it changes. Preserve column position as well as type in manual foreign tables, and test large extracts because scans decode the source file again. No preload or restart is required.
