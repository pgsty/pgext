## Usage

Sources:

- [MobilityDB v1.3.1 README](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/README.md)
- [Extension control file](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/sql/mobilitydb.in.control)
- [Version 1.3 migration manual](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/introduction.xml)
- [Temporal spatial API](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/temporal_spatial_p1.xml)
- [Version 1.3.1 release and upgrade](https://github.com/MobilityDB/MobilityDB/releases/tag/v1.3.1)
- [1.3.0 to 1.3.1 SQL migration](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/sql/mobilitydb--1.3.0--1.3.1.sql)

`mobilitydb` 1.3.1 extends PostgreSQL and PostGIS with temporal values and moving-object trajectories. It supports storing changing attributes, reconstructing positions at a timestamp, and indexing space-time bounds. This patch fixes a backend-crashing binary-input vulnerability; installations on 1.3.0 should upgrade.

### Enable the Extension

The release requires PostgreSQL 14 or later and PostGIS 3 or later; it also adds PostgreSQL 19 build support. Package availability is tracked separately. Upstream requires loading the matching PostGIS library and recommends this lock allocation:

```conf
shared_preload_libraries = 'postgis-3'
max_locks_per_transaction = 128
```

Append the PostGIS library to the existing preload list, restart PostgreSQL, and enable both extensions in the target database with an authorized administrative role:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
```

### Store and Query a Trajectory

The example uses projected coordinates and complete UTC timestamps. Choose the coordinate reference system appropriate for the application; geographic coordinates require different distance semantics.

```sql
CREATE TABLE trips (
    trip_id bigint PRIMARY KEY,
    trip tgeompoint NOT NULL
);

INSERT INTO trips VALUES (
    1,
    tgeompoint 'SRID=3857;[Point(0 0)@2026-01-01 08:00:00+00,
                         Point(1000 0)@2026-01-01 09:00:00+00]'
);

SELECT valueAtTimestamp(trip, '2026-01-01 08:30:00+00'),
       ST_AsText(trajectory(trip)),
       length(trip),
       speed(trip)
FROM trips;

CREATE INDEX trips_space_time_idx ON trips USING gist (trip);

SELECT trip_id
FROM trips
WHERE trip && stbox(
    ST_MakeEnvelope(-100, -100, 1100, 100, 3857),
    tstzspan '[2026-01-01 08:00:00+00, 2026-01-01 09:00:00+00]'
);
```

The bounding-box operator supplies an indexable filter. Apply the appropriate exact temporal or spatial predicate afterward when bounding overlap is insufficient.

### Type and Function Index

- `tbool`, `tint`, `tfloat`, and `ttext`: time-varying scalar values.
- `tgeompoint` and `tgeogpoint`: moving geometry or geography points; `tnpoint` represents a network point when that optional family is built.
- `tgeometry` and `tgeography`: arbitrary changing spatial values with discrete or step interpolation.
- `tcbuffer`, `tpose`, and `trgeometry`: optional experimental spatial families in the 1.3 line; do not assume every build contains them.
- Instant, sequence, and sequence-set representations describe one timestamp, one sequence, or multiple non-overlapping sequences. Linear interpolation is type-dependent.
- `valueAtTimestamp`, `startTimestamp`, `endTimestamp`, and `duration`: inspect temporal extent and values.
- `atTime` and `atGeometry`: restrict values to a time domain or geometry.
- `trajectory`, `length`, and `speed`: inspect the spatial path and motion.
- `twAvg` and `tUnion`: time-weighted summaries and temporal aggregation.
- GiST and SP-GiST operator classes accelerate supported temporal and space-time bounding queries.

### Upgrade and Safety Boundaries

Install the new library and SQL files, then update each database:

```sql
ALTER EXTENSION mobilitydb UPDATE TO '1.3.1';
SELECT extversion FROM pg_extension WHERE extname = 'mobilitydb';
```

- Version 1.3.1 fixes CVE-2026-102639: malformed WKB temporal, set, or span input could read beyond the input buffer and crash a backend. The SQL migration alone does not replace the vulnerable library; reconnect or restart processes that loaded the older binary.
- The migration removes five same-base-type `<->` operators and their `set_distance` functions because they conflict with operators supplied by `btree_gist`. Review dependent objects before updating and use the appropriate `btree_gist` operators where needed.
- Upgrading from the 1.2 line to 1.3 changes the temporal binary format and requires the upstream backup-and-restore procedure. An in-place 1.3.0-to-1.3.1 SQL update does not replace that major-line migration.
- Coordinate systems, interpolation, gaps, inclusive bounds, and units affect results. Validate them against the data model rather than treating every trajectory as a continuous geographical line.
