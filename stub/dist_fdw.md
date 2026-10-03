## Usage

Sources:

- [Installation SQL](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/bulkload/dist_fdw--1.0.sql)
- [Extension control file](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/bulkload/dist_fdw.control)
- [Implementation (dist_fdw.cpp)](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/bulkload/dist_fdw.cpp)

`dist_fdw` is the openGauss bulk-load adapter for external files and distributed storage. Its source and SQL depend on openGauss kernel facilities; it is not a portable PostgreSQL extension.

### Enablement and Objects

Use a compatible openGauss installation that supplies this module. The versioned SQL registers `pg_catalog.dist_fdw_handler()`, `pg_catalog.dist_fdw_validator(text[], oid)` and the `dist_fdw` foreign data wrapper. Inspect the installed wrapper before configuring an import.

```sql
SELECT extname, extversion FROM pg_extension WHERE extname = 'dist_fdw';
SELECT fdwname FROM pg_foreign_data_wrapper WHERE fdwname = 'dist_fdw';
```

### Import Workflow and Limits

Create an external server and foreign table using the storage protocol, location, format and error-handling options supported by the exact openGauss deployment; then load the intended local target from that foreign table. Source supports separate local, remote and object-storage paths, with deployment-dependent behavior. A generic file or object-store recipe is unsafe across kernel variants. File paths are interpreted on database hosts, and import/export privileges and error-log retention require administrator review. The extension/control version is `1.0`; this does not identify the surrounding kernel release.
