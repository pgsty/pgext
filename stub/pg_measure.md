## Usage

Sources:

- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_measure/pg_measure.control)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_measure/pg_measure--1.0.sql)
- [Official source](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_measure/meson.build)

`pg_measure` `1.0` provides measurements with units, uncertainty and provenance. This preview module is scoped to the PostgreSQL 18 fork in AlohaDB 21.1.0; stock PostgreSQL compatibility is not established.

### Workflow

```sql
CREATE EXTENSION pg_measure;
SELECT m_convert(measurement(1, 0.01, 'm'), 'cm');
```

Installation requires a superuser; function use also requires the appropriate object privileges.

### Operational boundaries

The measurement type provides unit conversion, uncertainty propagation, arithmetic and aggregates. m_wmean computes weighted means; prov_track and prov_lineage record and inspect row provenance. Uncertainty follows first-order propagation and is not a complete statistical dependence model. Review the provenance table policies and grants when sharing results; retain history required for traceability.

This module corresponds to the former upstream name `alohadb_measure`. Follow AlohaDB’s cluster/object conversion procedure and do not substitute an unrelated implementation with a similar name. Rehearse dependencies, privileges and backup restoration before upgrades.
