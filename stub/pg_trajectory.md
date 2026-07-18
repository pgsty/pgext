## Usage

Sources:

- [Project README](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/README.md)
- [Extension control file](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/pg_trajectory.control)
- [Version 0.0.1 installation SQL](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/pg_trajectory--0.0.1.sql)
- [Trajectory data model](https://github.com/ahmetkucuk/pg-trajectory/blob/1b73624f752149f5c2cf25b1e2aabb616d967b27/model/dataTypes.sql)

`pg_trajectory` 0.0.1 is an experimental pure-SQL PostgreSQL/PostGIS model for spatiotemporal trajectories. It defines `tg_pair` as a timestamp and geometry pair, and `trajectory` as start time, end time, geometry type, bounding geometry, and an array of samples. The source also contains construction, modification, distance, and similarity routines.

### Installation boundary

The control file does not declare its PostGIS dependency, but the installation SQL references `geometry`. Install PostGIS first and test the entire script on the exact PostgreSQL and PostGIS majors:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_trajectory;

SELECT pg_typeof(NULL::tg_pair), pg_typeof(NULL::trajectory);
```

The sample timestamps use `timestamp` without time zone. Establish one time-zone convention before ingesting data, and require consistent geometry type and SRID at the application boundary.

### Prototype defects and compatibility

The reviewed 0.0.1 SQL has internal inconsistencies. The `trajectory` composite omits `sampling_interval`, yet `_trajectory()` assigns that field; other routines call mixed or undefined function names. These paths may fail only when first invoked because PL/pgSQL compiles lazily. Do not treat a successful `CREATE EXTENSION` as a functional acceptance test.

The repository has no current release, upgrade chain, compatibility matrix, or modern automated test evidence. Before any use, exercise every required constructor and operator, malformed and empty arrays, mixed geometry types, timestamps, SRIDs, nulls, and large trajectories. Avoid durable storage in its custom composite layout until dump and restore plus extension-upgrade behavior are proven. For production spatiotemporal workloads, prefer a maintained PostGIS-native schema unless this prototype is audited and repaired locally.
