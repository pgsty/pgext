## Usage

Sources:

- [MobilityDB MEST README](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/mobilitydb/README.md)
- [Control](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/mobilitydb/mobilitydb_mest.control)
- [SQL 1.0](https://github.com/MobilityDB/mest/blob/26503db3d1e6a019094da2316eb2862a8215f053/contrib/mobilitydb/mobilitydb_mest--1.0.sql)

`mobilitydb_mest` 1.0 supplies multi-entry R-tree, quadtree and k-d-tree operator classes for MobilityDB span sets, temporal values and temporal points.

### Prerequisites

The reviewed source requires PostgreSQL 17 or 18, MobilityDB/MEOS 1.4, `mest`, and development headers for json-c when compiling. Its control dependencies are `mobilitydb` and `mest`.

Preload the installed PostGIS and MobilityDB libraries, retaining other existing entries, and restart PostgreSQL. Upstream's example is:

```ini
shared_preload_libraries = 'postgis-3,libMobilityDB-1.4'
```

### Core Workflow

Enable the extension with a privileged session. The following table is an example input for temporal span-set indexing:

```sql
CREATE EXTENSION mobilitydb_mest CASCADE;
CREATE TABLE recorded_spans (id bigint, t tstzspanset);
CREATE INDEX recorded_spans_mrtree ON recorded_spans
USING mgist (t tstzspanset_mrtree_equisplit_ops (num_spans = 3));
```

For an existing table containing a temporal-point column:

```sql
CREATE TABLE recorded_trips (id bigint, trip tgeompoint);
CREATE INDEX recorded_trips_mquadtree ON recorded_trips
USING mspgist (trip tgeompoint_mquadtree_equisplit_ops (num_boxes = 3));
```

### Classes and Maintenance

`mgist` and `mspgist` are supplied by `mest`. The companion adds type-specific classes such as `tstzspanset_mrtree_equisplit_ops` and `tgeompoint_mquadtree_equisplit_ops`; the SQL also defines split strategies for numeric temporal values and temporal geographic points.

`num_spans` and `num_boxes` control decomposition and affect index size. Select a class and splitting strategy against representative queries instead of assuming one configuration is optimal. The reviewed source requires MobilityDB 1.4 even when a distribution packages an older MobilityDB release; that older package is not a verified substitute.
