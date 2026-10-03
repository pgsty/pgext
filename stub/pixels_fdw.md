## Usage

Sources:

- [Official README](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/README.md)
- [Extension control file](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/pixels_fdw.control)
- [Installation SQL](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/pixels_fdw--1.0.sql)
- [Build configuration](https://github.com/pixelsdb/pixels-postgres/blob/2407a7acda65d71c8b646d51b51564fdb0ef64ea/pixels_fdw/Makefile)

`pixels_fdw` reads Pixels columnar files through PostgreSQL foreign tables. This research integration uses the Pixels C++ reader and server-local file paths. Upstream explicitly requires preload and a server restart.

### Enablement

Append the library to the existing preload list and restart after installing matching library dependencies.

```conf
shared_preload_libraries = 'pixels_fdw'
```

### Core Workflow

The foreign table columns must match the input files. Paths and filter expressions use the upstream reader syntax.

```sql
CREATE EXTENSION pixels_fdw;
CREATE SERVER pixels_server FOREIGN DATA WRAPPER pixels_fdw;
CREATE FOREIGN TABLE example (id int, name varchar, birthday date, score decimal(15,2))
  SERVER pixels_server
  OPTIONS (filename '|/srv/pixels/data|', filters 'id > 1 & score < 90');
SELECT * FROM example;
```

### Dependencies and Limits

The source build links Pixels C++ components, Protocol Buffers and the C++ runtime. `filename` selects files and `filters` provides reader-side filtering. The control declares version `1.0` and relocatability, but no PostgreSQL-major matrix is documented. Restrict who can define file-backed foreign tables; the database service account must be able to read the files. No write workflow is documented.
