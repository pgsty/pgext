## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/pg_chart.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/pg_chart--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/meson.build)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_chart/pg_chart.c)

`pg_chart` `1.0` provides svg charts rendered from sql query results. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_chart;
SELECT chart_bar('SELECT ''east'' AS region, 12 AS amount', 'Orders');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

chart_bar, chart_line, chart_pie, chart_scatter and chart_area return SVG text. Supply a query with the appropriate label/numeric columns. The query runs through SPI with caller privileges; restrict who can submit query text and account for backend memory and runtime when charting large result sets.

This module corresponds to the former upstream name `alohadb_chart`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
