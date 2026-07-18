## Usage

Sources:

- [Rust implementation at the reviewed commit](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/src/lib.rs)
- [Generated version 0.0.0 SQL](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/open_meteo_extension/usr/share/postgresql/17/extension/pg_open_meteo--0.0.0.sql)
- [Sample weather migration](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/migrations/0003_weather_table/up.sql)
- [Cargo feature matrix](https://github.com/earth-metabolome-initiative/pg-open-meteo/blob/441c7d17d74809f02d721e79307cca322b9d2d94/Cargo.toml)

`pg_open_meteo` is an unfinished pgrx scaffold. Despite its name, version 0.0.0 does not call the Open-Meteo API, fetch weather, schedule updates, or expose a weather function. Its generated extension SQL contains only three demonstration functions.

```sql
CREATE EXTENSION pg_open_meteo;
SELECT hello_pg_open_meteo();
SELECT strictly_positive(1);
SELECT x_must_be_bigger_than_y(2, 1);
```

Repository migrations define generic `position` and `sample_weather_data` tables, but those migrations are separate from the generated extension install SQL. Do not assume they are created by `CREATE EXTENSION`, and do not run them in an application schema without reviewing their unqualified object names.

The repository has no README, public API contract, release, declared license, or working weather integration. Treat it only as sample code; choose a maintained client or FDW for actual weather ingestion.
