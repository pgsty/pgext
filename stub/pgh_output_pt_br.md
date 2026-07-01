


## Usage

Sources: [PGHydro README](https://github.com/pghydro/pghydro/blob/master/README.md), [pgh_output_pt_br SQL](https://github.com/pghydro/pghydro/blob/master/pgh_output_pt_br--6.6.sql)

`pgh_output_pt_br` provides a Brazilian-Portuguese PgHydro output schema. It creates localized export tables for drainage-network products and a refresh function that populates them from the base PgHydro model.

### Enable

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pghydro;
CREATE EXTENSION pgh_output_pt_br;
```

### Output Tables

The extension creates the `pgh_output_pt_br` schema with tables including:

| Table | Purpose |
|-------|---------|
| `geoft_bho_trecho_drenagem` | Drainage-line export table |
| `geoft_bho_area_drenagem` | Drainage-area export table |
| `geoft_bho_curso_dagua` | Watercourse export table |
| `geoft_bho_ponto_drenagem` | Drainage-point export table |
| `geoft_bho_linha_costa` | Shoreline export table |

### Refresh Output

```sql
SELECT pgh_output_pt_br.pghfn_updateexporttables();
```

### Notes

Use this extension when exported PgHydro products need Brazilian-Portuguese object names. Regenerate the output after source drainage geometry, topology, or coding changes.
