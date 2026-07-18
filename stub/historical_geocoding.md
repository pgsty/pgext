## Usage

Sources:

- [Upstream workflow documentation at the reviewed commit](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/doc/overall.md)
- [Upstream README](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/README.md)
- [Extension control file](https://github.com/GeoHistoricalData/historical_geocoding/blob/596778df15067e6494c63c5e25d42b2843a90165/historical_geocoding.control)

`historical_geocoding` is a research-oriented SQL schema for matching a textual address and date against incomplete historical address data and ranking probable locations with confidence. It requires `geohistorical_objects`, `postal`, and `unaccent` according to its control file.

```sql
CREATE EXTENSION historical_geocoding CASCADE;
SELECT e.extname, e.extversion
FROM pg_extension AS e
WHERE e.extname IN (
  'historical_geocoding', 'geohistorical_objects', 'postal', 'unaccent'
);
```

Installation supplies a template of tables and functions; it does not provide a ready historical gazetteer. The documented workflow requires normalizing address/date input, importing and validating map-derived road and building-number data, matching by text, time and space, and ranking candidates. Results are probabilistic and should retain provenance and confidence rather than overwrite source facts.

The reviewed project is centered on French and Paris historical data and describes partially manual research workflows. Audit every function and required extension, especially dynamic SQL, inheritance, fuzzy time models, PostGIS operations, and libpostal behavior. Evaluate bias, missing data, aliases, temporal uncertainty, reproducibility, and human correction with a gold-standard sample before any operational use.
