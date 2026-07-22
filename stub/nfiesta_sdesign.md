## Usage

Sources:

- [Official repository README](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/README.md)
- [Official extension control file](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/nfiesta_sdesign.control)
- [Official extension SQL](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/nfiesta_sdesign--1.0.0.sql)
- [Official import function](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign/-/blob/37945954242446614f08e36b958710ef965fca76/functions/fn_import_data.sql)

`nfiesta_sdesign` installs the sampling-design data model used by the NFIesta forest-inventory system. It models countries, strata, panels, clusters, plots, inventory campaigns, reference-year sets, sampling weights, and PostGIS geometries in the fixed `sdesign` schema; it is a domain schema rather than a general statistics function library.

### Installation and Inspection

Version `1.1.17` requires `plpgsql` and `postgis` and is not relocatable:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION nfiesta_sdesign;

SELECT table_name
FROM information_schema.tables
WHERE table_schema = 'sdesign'
ORDER BY table_name;
```

The main tables include `sdesign.c_country`, `sdesign.t_strata_set`, `sdesign.t_stratum`, `sdesign.t_panel`, `sdesign.t_cluster_configuration`, `sdesign.t_cluster`, `sdesign.f_p_plot`, `sdesign.t_inventory_campaign`, and `sdesign.t_reference_year_set`, with mapping tables connecting panels, clusters, campaigns, and reference years.

### Import and Validation

`sdesign.fn_import_data(text)` imports a complete dataset from a caller-supplied staging schema. That schema must contain the exact staging relations expected by the function, including countries, cluster configurations, strata sets, strata, panels, clusters, plots, inventory campaigns, measurement dates, and reference-year mappings.

Validation triggers enforce relationships such as campaign begin/end pairs, panel/reference-year mappings, plot measurement dates, cluster configuration combinations, and sampling-weight sums. `sdesign.fn_create_buffered_geometry()` derives MultiPolygon sampling-frame geometry for the supported tract layouts.

### Operational Boundaries

Treat installation, import, upgrade, and removal as application-schema migrations: they create and mutate many persistent tables, sequences, constraints, triggers, and PostGIS values. The import function builds dynamic SQL from the supplied schema string without identifier quoting, so allow only a trusted, prevalidated schema name and run it with a least-privilege owner. Import is not an idempotent row upsert and can fail late on deferred validation; use a transaction, validate staging counts and SRIDs, and test rollback on a copy. Version `1.1.17` adds a uniqueness constraint on panel/reference-year mappings, so clean duplicates before upgrading.
