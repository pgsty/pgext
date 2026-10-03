## Usage

Sources:

- [extensions/pg_solid/pg_solid.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/pg_solid.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_solid/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/Cargo.toml)
- [extensions/pg_solid/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_solid/pgbrew.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/pgbrew.toml)
- [extensions/pg_solid/build.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_solid/build.rs)

`pg_solid` 0.3.0 provides solid geometry in `pgsolid`, including constructors, Boolean operations, measurements, transforms, spatial indexing and CAD/BIM import/export.

### Core Workflow

```sql
CREATE EXTENSION pg_solid;
SELECT pgsolid.solid_volume(pgsolid.solid_box(100, 200, 300));
SELECT pgsolid.solid_is_valid(pgsolid.solid_sphere(10));
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is declared. OpenCASCADE 7.6+ native libraries are required; file/model imports must be restricted and resource use bounded. Constructors and measurement functions operate in the coordinates and units supplied by the caller. Validity checks and tolerance handling matter for imported or computed shapes. This is a separate solid-geometry implementation, not an assertion of PostGIS dependency or interchangeable geometry encoding. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
