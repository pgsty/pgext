## Usage

Sources:

- [Upstream workflow documentation at the reviewed commit](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/doc/overall.md)
- [Declared and operational dependency script](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/creating_dependency_extensions.sql)
- [Geocoding API source](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/src/050_geocoding_api.sql)
- [Extension control file](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/historical_geocoding.control)

`historical_geocoding` is a research-oriented SQL schema for matching an address and fuzzy date against incomplete historical address data. It ranks candidate places using semantic, temporal, number, scale, spatial-distance, and source-precision terms; it does not install a ready-to-query gazetteer.

### Installation and Workflow

The control file declares `geohistorical_objects`, `postal`, and `unaccent`. Upstream's dependency script also installs `postgis`, `pgsfti`, and `pg_trgm`, which the generated SQL uses but does not declare in `requires`. Prepare and validate all six before creating the extension.

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgsfti;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION geohistorical_objects;
CREATE EXTENSION postal;
CREATE EXTENSION unaccent;
CREATE EXTENSION historical_geocoding;

SELECT *
FROM historical_geocoding.geocode_name_optimised(
  '10 rue du temple, Paris',
  sfti_makesfti('1872-11-15'::date)
)
LIMIT 10;
```

The main API family includes `geocode_name_base`, `geocode_name_optimised_inter`, and `geocode_name_optimised`. A real workflow must normalize address/date input, import and validate road and building-number sources, populate precise and rough localization tables, retain source lineage, and evaluate ranked candidates against manually verified examples.

### Security and Installation Hazards

`ordering_priority_function` is interpolated directly into dynamic SQL after a deny-list function removes selected characters and keywords. That is not a safe SQL-injection boundary. Never pass user-controlled text; expose fixed, reviewed scoring expressions through a wrapper and revoke direct API execution from untrusted roles.

The PGXS build concatenates source files that contain unguarded sample `SELECT` statements and test `CREATE TABLE` statements into the extension install SQL. Inspect the generated `historical_geocoding--1.0.sql` and perform the first install only in a disposable database; do not assume creation is side-effect-free.

The reviewed project focuses on French and Paris historical data and includes partially manual research workflows. Results are probabilistic: evaluate bias, missing data, aliases, temporal uncertainty, SRID 2154 assumptions, libpostal behavior, reproducibility, and human correction with a gold-standard sample. Preserve provenance and confidence rather than overwriting source facts.
