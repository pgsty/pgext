## Usage

Sources:

- [Official upstream README](https://github.com/accarniel/festival/blob/b4ae7e2a3eff43e1be65806dc84242df8be88e3b/README.md)
- [Official extension control file (festival.control)](https://github.com/accarniel/festival/blob/b4ae7e2a3eff43e1be65806dc84242df8be88e3b/festival.control)
- [Official extension SQL (festival--1.1.1.sql)](https://github.com/accarniel/festival/blob/b4ae7e2a3eff43e1be65806dc84242df8be88e3b/festival--1.1.1.sql)

`festival` — FESTIval is a framework, implemented as a PostgreSQL extension, for conducting experimental evaluations of spatial index structures. The complete documentation of FESTIval is available here. Use it for the corresponding spatial data or geospatial workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION festival;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `FT_ADelete(absolute_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `int4`.
- `FT_ADelete(index_name text, index_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `int4`.
- `FT_AInsert(absolute_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `int4`.
- `FT_AInsert(index_name text, index_path text, p int4, geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `int4`.
- `FT_ApplyAllModificationsForFAI(absolute_path text)` is an extension function.
- `FT_ApplyAllModificationsForFAI(index_name text, index_path text)` is an extension function.
- `FT_ApplyAllModificationsFromBuffer(absolute_path text)` is an extension function.
- `FT_ApplyAllModificationsFromBuffer(index_name text, index_path text)` is an extension function.
- `FT_AQuerySpatialIndex(absolute_path text, type_query int4, obj geometry, predicate int4, processing_option int4 default 1, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `SETOF`.
- `FT_AQuerySpatialIndex(index_name text, index_path text, type_query int4, obj geometry, predicate int4, processing_option int4 default 1, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `SETOF`.
- `FT_AUpdate(absolute_path text, old_p int4, old_geom geometry, new_p int4, new_geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `int4`.
- `FT_AUpdate(index_name text, index_path text, old_p int4, old_geom geometry, new_p int4, new_geom geometry, statistic_options int4 default 1, location_statistics int4 default 1, file_statistics text default NULL)` is an extension function and returns `int4`.
- `FT_CollectOrderOfReadWrite()` is an extension function.
- `FT_CreateEmptySpatialIndex(index_id int4, absolute_path text, src_id int4, bc_id int4, sc_id int4, buf_id int4)` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.1.1`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
