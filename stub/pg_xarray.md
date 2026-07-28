## Usage

Sources:

- [Official upstream README](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/README.md)
- [Official extension control file (pg_xarray.control)](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_xarray/pg_xarray.control)
- [Official implementation source](https://github.com/matroidbe/pg_extensions-releases/blob/bbc2398a3e45c722beef6dd26f698bc2a017e241/extensions/pg_xarray/src/lib.rs)

`pg_xarray` — Catalog and query layer for chunked scientific arrays (NetCDF, Zarr, HDF5, GRIB, COG, SELAFIN, MED, CGNS, FITS). Use it for the corresponding analytical or storage workflow. Upstream explicitly says it is not production-ready.

### Core Workflow

```sql
CREATE EXTENSION pg_xarray;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `chunk_count` is an extension function.
- `list_datasets()` is an extension function.
- `register_chunk` is an extension function.
- `register_dataset` is an extension function.
- `register_file` is an extension function.
- `register_mesh` is an extension function.
- `register_mesh_cell` is an extension function.
- `register_mesh_node` is an extension function.
- `register_mesh_version` is an extension function.
- `register_variable` is an extension function.
- `xarray_to_glb` is an extension function.
- `xarray_to_png` is an extension function.

### Requirements and Caveats

- The catalog records version `0.2.0`.
- Install the confirmed extension dependencies first: `postgis`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream explicitly says the project is not production-ready.
- Upstream describes the project as a proof of concept.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
