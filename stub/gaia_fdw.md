## Usage

Sources:

- [production/fdw/README.md](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/README.md)
- [production/CMakeLists.txt](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/CMakeLists.txt)
- [production/fdw/CMakeLists.txt](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/CMakeLists.txt)
- [production/fdw/src/gaia_fdw.control](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/src/gaia_fdw.control)
- [production/fdw/src/gaia_fdw.sql](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/src/gaia_fdw.sql)
- [production/fdw/src/CMakeLists.txt](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/src/CMakeLists.txt)
- [production/fdw/tests/airport_setup.sql](https://github.com/gaia-platform/GaiaPlatform/blob/d0473e6e5710b3dbfc4ab674dda9ebaee8a5a5d9/production/fdw/tests/airport_setup.sql)

`gaia_fdw` 0.6 exposes a Gaia Platform database as PostgreSQL foreign tables. CMake derives this extension version from the major and minor parts of production 0.6.0. The pinned source is from 2022; it does not establish compatibility with current PostgreSQL majors.

### Core Workflow

After the Gaia database is running and its airport dataset is available, install the wrapper and import its tables:

```sql
CREATE EXTENSION gaia_fdw;
CREATE SERVER gaia FOREIGN DATA WRAPPER gaia_fdw;
CREATE SCHEMA airport_fdw;
IMPORT FOREIGN SCHEMA airport_fdw FROM SERVER gaia INTO airport_fdw;
```

### Prerequisites and Objects

Installation requires superuser privileges by default. The relocatable control loads `gaia_fdw-0.6`; the build links Gaia client, catalog and payload libraries. The SQL creates `gaia_fdw_handler`, `gaia_fdw_validator` and the `gaia_fdw` foreign data wrapper. No extension-specific preload is declared. The Gaia client must be able to connect to the external database before accessing its tables.

### Data Semantics

The wrapper supports reading and modifying Gaia data, but NULL semantics differ from PostgreSQL. Existing NULL strings may be read as NULL; writing NULL to a non-reference field acts as a no-op and can leave the serialization default, such as an empty string or zero. Writing NULL to a reference removes that reference. `gaia_id` cannot be updated. Check these behaviors before using imported tables for applications that depend on SQL NULL semantics.
