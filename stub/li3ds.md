## Usage

Sources:

- [Official upstream README](https://github.com/li3ds/pg-li3ds/blob/60a45f4e291aa2f14ef702bcb8d6ffc5811d0576/README.rst)
- [Official extension control file (li3ds.control)](https://github.com/li3ds/pg-li3ds/blob/60a45f4e291aa2f14ef702bcb8d6ffc5811d0576/extension/li3ds.control)
- [Official extension SQL (li3ds--1.0.0.sql)](https://github.com/li3ds/pg-li3ds/blob/60a45f4e291aa2f14ef702bcb8d6ffc5811d0576/extension/li3ds--1.0.0.sql)

`li3ds` — PostgreSQL extension for managing 3D sensor data. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION li3ds;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `check_datasource_uri(uri text)` is an extension function and returns `boolean`.
- `check_pcpatch_column(schema_table_column varchar)` is an extension function and returns `boolean`.
- `check_timezone_name(timezone varchar)` is an extension function and returns `boolean`.
- `check_transfo_args(parameters jsonb, transfo_type_id int)` is an extension function and returns `boolean`.
- `check_transfotree_istree(transfo_trees integer[])` is an extension function and returns `boolean`.
- `dijkstra(config integer, source integer, target integer, stoptosensor varchar default '')` is an extension function and returns `integer[]`.
- `foreign_key_array(arr integer[], foreign_table regclass)` is an extension function and returns `boolean`.
- `isconnected(transfos integer[], doubletransfo boolean default False)` is an extension function and returns `boolean`.
- `postgres_version()` is an extension function and returns `text`.
- `transform(box4d libox4d, config integer, source integer, target integer, ttime float8 default 0.0)` is an extension function and returns `libox4d`.
- `transform(box4d libox4d, config integer, source integer, target integer, ttime text)` is an extension function and returns `libox4d`.
- `transform(box4d libox4d, func_name text, func_sign text[], params text)` is an extension function and returns `libox4d`.
- `transform(box4d libox4d, transfo integer, ttime float8 default 0.0)` is an extension function and returns `libox4d`.
- `transform(box4d libox4d, transfo integer, ttime text)` is an extension function and returns `libox4d`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- Install the confirmed extension dependencies first: `postgis`, `plpython2u`, `pointcloud`, `pointcloud_postgis`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
