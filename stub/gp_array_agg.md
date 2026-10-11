## Usage

Sources:

- [gpdb-doc/markdown/ref_guide/modules/gp_array_agg.html.md](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpdb-doc/markdown/ref_guide/modules/gp_array_agg.html.md)
- [gpcontrib/gp_array_agg/gp_array_agg--1.0.0.sql](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_array_agg/gp_array_agg--1.0.0.sql)
- [gpcontrib/gp_array_agg/Makefile](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_array_agg/Makefile)
- [LICENSE](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/LICENSE)
- [gpcontrib/gp_array_agg/gp_array_agg.control](https://github.com/open-gpdb/gpdb/blob/c6e63e9a53d01ed6ea3a93e5e280f38e8b9d808b/gpcontrib/gp_array_agg/gp_array_agg.control)

`gp_array_agg` provides an array aggregate that can combine partial results across Greenplum workers. The reviewed 1.0.0 module belongs to the Greenplum 6 distribution maintained by OpenGPDB.

### Core Workflow

```sql
CREATE EXTENSION gp_array_agg;
SELECT gp_array_agg(v ORDER BY v)
FROM (VALUES (3), (1), (NULL::integer), (2)) AS input(v);
```

### Operational Boundaries

`gp_array_agg(anyelement)` collects input values, including nulls; an empty input produces null. Element order is unspecified unless the aggregate has an explicit `ORDER BY`. The supporting `array_agg_combine`, `array_agg_serialize`, and `array_agg_deserialize` functions serve the aggregate implementation, not a separate application workflow.

Register the supplied module as a superuser in each database that needs it. It uses an on-demand C library, requires no preload configuration, and can be installed in a chosen schema. This entry describes the Greenplum kernel module; it does not establish support for an upstream PostgreSQL major or availability in Pigsty packages.
