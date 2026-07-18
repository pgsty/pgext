## Usage

Sources:

- [Upstream README](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/README.md)
- [Upstream documentation](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/doc/pg_popyramids_datamarts.md)
- [Extension control file](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/pg_popyramids_datamarts.control)
- [Extension SQL](https://github.com/torrespri/pg_popyramids_datamarts/blob/24a3360122d8c3ba9503b36cc41b50d48a248552/sql/pg_popyramids_datamarts--1.0.0.sql)

`pg_popyramids_datamarts` is the data-mart layer of the popyramids population-pyramid application, not a standalone general-purpose extension. Version `1.0.0` requires PostGIS and assumes an existing application database with `dms` and `ods` schemas, the `ods.pyrint` type, and the `ods.main` source table.

Check the application prerequisites before creating it; only a matching popyramids database should proceed to the activation and query steps:

```sql
SELECT
  to_regnamespace('dms') AS dms_schema,
  to_regnamespace('ods') AS ods_schema,
  to_regtype('ods.pyrint') AS ods_pyrint,
  to_regclass('ods.main') AS ods_main;

CREATE EXTENSION postgis;
CREATE EXTENSION pg_popyramids_datamarts;

SELECT dms.chibo_give_me_pyramids(ARRAY[1, 2]);
```

### Deployment constraints

The install script creates application-specific types, functions, casts, and materialized views, and it reads directly from `ods.main`. It also contains ownership and grant statements for the roles `postgres` and `chichinabo`. Creation will fail in a generic database when expected schemas, types, tables, or roles are absent; review and patch those environment assumptions before deployment.
