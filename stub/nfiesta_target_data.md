## Usage

Sources:

- [Official upstream README](https://gitlab.com/nfiesta/nfiesta_target_data/-/blob/main/README.md)
- [Official extension control file](https://gitlab.com/nfiesta/nfiesta_target_data/-/blob/main/nfiesta_target_data.control)
- [Official project page](https://gitlab.com/nfiesta/nfiesta_target_data)

`nfiesta_target_data` — PostgreSQL extension for manipulation of target data - aggregation of local density contributions at the plot level. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION nfiesta_target_data;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `2.33.2`.
- Install the confirmed extension dependencies first: `plpgsql`, `plpython3u`, `nfiesta_sdesign`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
