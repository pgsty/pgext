## Usage

Sources:

- [postgresql/contrib/temporalbox/temporalbox--1.0.sql](https://github.com/yash7312/DBIS_Assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporalbox/temporalbox--1.0.sql)
- [README.md](https://github.com/yash7312/DBIS_Assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/README.md)
- [postgresql/contrib/temporalbox/temporalbox.control](https://github.com/yash7312/DBIS_Assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporalbox/temporalbox.control)

`temporalbox` is a small SQL helper extension in an academic temporal-indexing project. It maps an integer key and timestamp interval to a two-dimensional `cube` value for spatial-index experiments.

### Core Workflow

```sql
CREATE EXTENSION cube WITH SCHEMA public;
CREATE EXTENSION temporalbox;
SELECT temporalbox(7, '[2026-01-01,2026-02-01)'::tsrange);
SELECT temporalbox_point(7, TIMESTAMP '2026-01-15');
```

### Operational Boundaries

Install `cube` in `public` first: the functions explicitly call `public.cube`, and the install script’s dependency creation is not a control-file dependency. `temporalbox` maps a `tsrange`; `temporalbox_point` maps one timestamp; `temporalbox_range` maps two bounds. Time is represented as floating-point epoch seconds. An unbounded upper range becomes infinity, but range inclusivity is not retained by the box representation; use an exact range predicate when endpoint semantics matter. No preload or native library is involved. The surrounding project contains modified PostgreSQL index code; this entry describes only the SQL helper and does not certify its broader kernel changes or a PostgreSQL-major matrix.
