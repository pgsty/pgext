


## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_output SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output--6.6.sql)

`pgh_output` creates export tables for PgHydro drainage-network products. It materializes drainage lines, drainage areas, watercourses, drainage points, and shorelines in a schema intended for downstream reporting or GIS exchange.

### Enable

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output;
```

### Export Tables

The extension creates the `pgh_output` schema with tables including:

| Table | Purpose |
|-------|---------|
| `geoft_bho_drainage_line` | Exported drainage-line features |
| `geoft_bho_drainage_area` | Exported drainage-area features |
| `geoft_bho_watercourse` | Exported watercourse features |
| `geoft_bho_drainage_point` | Exported drainage-point features |
| `geoft_bho_shoreline` | Exported shoreline features |

### Refresh Output

```sql
SELECT pgh_output.pghfn_updateexporttables();
```

Run the refresh function after loading or recalculating the base PgHydro network.

### Notes

Validate the PgHydro network first with consistency checks when possible. The export tables are derived products and should be regenerated after upstream drainage geometry, topology, or coding changes.
