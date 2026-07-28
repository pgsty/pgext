## Usage

Sources:

- [Official upstream README](https://github.com/gabbasb/pg_gnuplot/blob/1716363eb3138f63b353eb749bee82782307e181/README.md)
- [Official extension control file (pg_gnuplot.control)](https://github.com/gabbasb/pg_gnuplot/blob/1716363eb3138f63b353eb749bee82782307e181/pg_gnuplot.control)
- [Official extension SQL (pg_gnuplot--1.0.sql)](https://github.com/gabbasb/pg_gnuplot/blob/1716363eb3138f63b353eb749bee82782307e181/pg_gnuplot--1.0.sql)

`pg_gnuplot` — PostgreSQL extension to plot graphs using GNUPlot. Use it for the corresponding SQL or database utility workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_gnuplot;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gnuplot_version()` is an extension function and returns `cstring`.
- `pg_gnuplot_version()` is an extension function and returns `pg_catalog`.
- `pg_plot(db_query pg_catalog.text, plot_cmd pg_catalog.text)` is an extension function and returns `pg_catalog`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
