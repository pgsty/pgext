## Usage

Sources:

- [Official upstream README](https://github.com/oncox/pggeojson/blob/e1a1c64218ab215b08325f2f8b38cb333342a786/README.md)
- [Official extension control file (pggeojson.control)](https://github.com/oncox/pggeojson/blob/e1a1c64218ab215b08325f2f8b38cb333342a786/pggeojson.control)
- [Official extension SQL (pggeojson--1.0.sql)](https://github.com/oncox/pggeojson/blob/e1a1c64218ab215b08325f2f8b38cb333342a786/pggeojson--1.0.sql)

`pggeojson` — pgGeoJSON is a PostgreSQL module providing additional functionality for generating GeoJSON output. Use it for the corresponding spatial data or geospatial workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pggeojson;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `PGG_AsGeoJSON` is an extension function.
- `PGG_AsGeoJSON` is an extension function.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
