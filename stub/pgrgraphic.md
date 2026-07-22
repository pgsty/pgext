## Usage

Sources:

- [Official README](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/README.md)
- [Extension SQL](https://github.com/asotolongo/pgrgraphic/blob/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/pgrgraphic--1.0.0.sql)
- [Official examples](https://github.com/asotolongo/pgrgraphic/tree/ff2a6ddcf42d4d73be4ec3c694b3e330ffcb3c0c/ejemplo)

pgrgraphic is an old PL/R extension that renders PostgreSQL query results to PNG or JPEG chart files on the database server. It covers bars, lines, points, histograms, contours, heat maps, box plots, and pie charts. It is tied to a PostgreSQL 9.x, R, PL/R, and X-server-era environment and is unsafe for untrusted use.

### Core Workflow

After the server administrator installs R, PL/R, lattice, gplots, and a suitable graphics environment, a trusted user can pass a SQL text query to a chart wrapper:

```sql
CREATE EXTENSION plr;
CREATE EXTENSION pgrgraphic;

SELECT pgrgraphic.grafica_barras_simples(
  'orders_by_status',
  'Orders by status',
  'Status',
  'Count',
  'SELECT count(*) FROM orders GROUP BY status ORDER BY status',
  ARRAY['new', 'done'],
  'vertical',
  'png'
);
```

The function writes a file below the PostgreSQL operating-system account's home or working directory and returns text; upstream uses 10 as its success marker.

### Important Objects

- `grafica_barras_simples` and `grafica_barras_multiples` render bar charts.
- `grafica_lineas`, `grafica_lineas_puntos`, and `grafica_puntos` render line and point charts.
- `grafica_histograma` and `grafica_histograma_3d` render histograms.
- `grafica_caja`, `grafica_contorno`, `grafica_dispersion`, `grafica_mapa_de_calor`, and `grafica_pie` cover other chart forms.
- The parallel `r_*` functions are the PL/R implementations behind the SQL wrappers.

### Security and Operational Notes

PL/R is untrusted, and pgrgraphic accepts executable SQL text and writes arbitrary chart names through the database-server account. Do not grant these functions to untrusted roles. The installation SQL hard-codes the `pgrgraphic` schema and ownership by `postgres`, which can fail or create surprising ownership. Files are not lifecycle-managed and can fill the server disk. Prefer exporting query data to a separately sandboxed visualization service.
