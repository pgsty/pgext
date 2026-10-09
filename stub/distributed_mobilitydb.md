## Usage

Sources:

- [README.md](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/README.md)
- [sql/distributed_mobilitydb--1.0.sql](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/sql/distributed_mobilitydb--1.0.sql)
- [sql/partitioning/tiling.sql](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/sql/partitioning/tiling.sql)
- [CMakeLists.txt](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/CMakeLists.txt)
- [src/general/shared_library_init.c](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/src/general/shared_library_init.c)
- [src/distributed_mobilitydb.control](https://github.com/mbakli/DistributedMobilityDB/blob/f13cd2f1b4e9aeb8c83a132914560a53aa5c99f9/src/distributed_mobilitydb.control)

`distributed_mobilitydb` distributes PostGIS and MobilityDB data through Citus. It supplies spatial/temporal tiling, distribution metadata and planner hooks; upstream labels it under development.

### Core Workflow

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION mobilitydb;
CREATE EXTENSION citus;
CREATE EXTENSION distributed_mobilitydb;
```

### Operational Boundaries

Prepare a working Citus cluster and compatible PostGIS 3+, MobilityDB 1.1+ and Citus 10+ dependencies. Create the extension as a superuser. `create_spatiotemporal_distributed_table` creates a tiled output table from an input table; its table names, tile count, method, geometry column and segmentation options determine the physical layout. The reference-table option replicates a table without tiling and rejects a non-default tile count.

The README requires PostgreSQL 13+ while build metadata names 13–14; the catalog records that narrower declared build range, not an assertion about later releases. Some extension metadata and types are placed in pg_catalog. Plan data redistribution, disk usage and privileges before converting real tables, and preserve distributed data and metadata in coordinated backups. Citus has its own cluster preload/setup requirements; the source does not declare a separate preload requirement for this library.
