## Usage

Sources:

- [Official extension control file](https://gitlab.com/nfiesta/nfiesta_gisdata_old/-/blob/main/extension/nfiesta_gisdata.control)
- [Official project page](https://gitlab.com/nfiesta/nfiesta_gisdata_old)

`nfiesta_gisdata` — nfi gisdata database. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION nfiesta_gisdata;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `3.0.7`.
- Install the confirmed extension dependencies first: `plpgsql`, `postgis`, `postgis_raster`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
