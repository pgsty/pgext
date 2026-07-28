## Usage

Sources:

- [Official upstream README](https://github.com/mmiranda2/c-postgres-gis/blob/93ef09fbee9bb6b30282b85001d968f60dd7ec7e/README.md)
- [Official extension control file (gis_objects.control)](https://github.com/mmiranda2/c-postgres-gis/blob/93ef09fbee9bb6b30282b85001d968f60dd7ec7e/gis_objects.control)
- [Official extension SQL (gis_objects--1.0.sql)](https://github.com/mmiranda2/c-postgres-gis/blob/93ef09fbee9bb6b30282b85001d968f60dd7ec7e/gis_objects--1.0.sql)

`gis_objects` — Postgres extension for Point, Polygon, and Point-in-Polygon operations in C. Use it for the corresponding spatial data or geospatial workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION gis_objects;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `my_log()` is an extension function and returns `varchar`.
- `my_point_in_polygon(mypoint, mypoint[])` is an extension function and returns `boolean`.
- `mypoint_in(cstring)` is an extension function and returns `mypoint`.
- `mypoint_out(mypoint)` is an extension function and returns `cstring`.
- `mypoint` is an extension-defined type.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
