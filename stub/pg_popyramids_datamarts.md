## Usage

Sources:

- [Official repository](https://github.com/torrespri/pg_popyramids_datamarts)
- [Extension control file](https://github.com/torrespri/pg_popyramids_datamarts/blob/master/pg_popyramids_datamarts.control)
- [Version 1.0.0 extension SQL](https://github.com/torrespri/pg_popyramids_datamarts/blob/master/sql/pg_popyramids_datamarts--1.0.0.sql)
- [Upstream project documentation](https://github.com/torrespri/pg_popyramids_datamarts/blob/master/doc/pg_popyramids_datamarts.md)

pg_popyramids_datamarts installs the data-mart layer of the popyramids application: demographic pyramid types and transformations, PostGIS geometry helpers, GeoJSON-producing functions, and populated materialized views. It is application schema, not a standalone general-purpose extension.

### Required Application Environment

The control file requires PostGIS and marks the extension non-relocatable, but the version 1.0.0 SQL has additional prerequisites that are not encoded there. Before installation, the database must already have the destination schema, the operational-data-store schema with its pyramid type and source table, and the hard-coded roles referenced by grants.

```sql
CREATE EXTENSION postgis;
CREATE SCHEMA dms;

-- The popyramids ODS deployment must already provide ods.pyrint and ods.main.
-- Provision the roles referenced by the upstream grants before continuing.
CREATE EXTENSION pg_popyramids_datamarts;
```

Do not run the final statement in a generic database: extension creation builds materialized views with data from the existing operational store and will fail if any application-specific type, relation, column, role, or schema is absent.

### Main Object Families

- Demographic types represent age buckets, male and female counts or percentages, variables, and classified pyramid shapes.
- Array and pyramid helpers aggregate buckets, compute totals and percentages, convert operational-store values, and derive PostGIS geometry.
- Query helpers select population-pyramid identifiers and return GeoJSON-like feature collections for map clients.
- Materialized views expose raw, five-year, ten-year, and broad-group variants, percentage variants, statistics, and user-interface catalogs.

### Refresh and Ownership Boundaries

The materialized views are created with data and depend on the existing source table and on one another. The extension does not supply an automated refresh schedule; refresh them in dependency order after source changes and assess locking, runtime, and storage on production-sized data.

Version 1.0.0 embeds application-specific grants and object names, including PostgreSQL roles. Review the complete SQL against the target popyramids deployment before installation instead of treating the control-file dependency list as sufficient. Test installation and removal on a restored copy, because the extension owns a large graph of schema-qualified types, casts, functions, and populated views.

No preload or restart is declared. Database-level installation still requires authority to create all application objects and to grant privileges to the referenced roles.
