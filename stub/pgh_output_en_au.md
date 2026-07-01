


## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_output_en_au SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output_en-au--6.6.sql)

`pgh_output_en_au` provides an Australian-English PgHydro output schema. It creates localized export tables for drainage lines, drainage areas, and drainage points, plus a refresh function that populates them from the base PgHydro network.

### Enable

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output_en_au;
```

### Output Tables

The extension creates the `pgh_output_en_au` schema with tables including:

| Table | Purpose |
|-------|---------|
| `geoft_bho_drainage_line` | Drainage-line export table |
| `geoft_bho_drainage_area` | Drainage-area export table |
| `geoft_bho_drainage_point` | Drainage-point export table |

### Refresh Output

```sql
SELECT pgh_output_en_au.pghfn_updateexporttables();
```

Run the refresh after recalculating or editing the source PgHydro objects.

### Notes

Use this extension when the exported schema should expose Australian-English naming conventions while still being derived from the PgHydro drainage network.
