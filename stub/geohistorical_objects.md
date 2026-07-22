## Usage

Sources:

- [Official README](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/README.md)
- [Version 1.0 control file](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/geohistorical_objects.control)
- [Version 1.0 SQL implementation](https://github.com/GeoHistoricalData/geohistorical_objects/blob/5e7ec83dc568a509a88a0a66e3add8310bba13b3/geohistorical_objects.sql)

`geohistorical_objects` provides an inheritance-based data model for geographic objects whose names, geometry, source, date, and precision vary through history. It creates objects in the fixed `geohistorical_object` schema and depends on `postgis`, `pgsfti`, `unaccent`, and `pg_trgm`.

### Core Model

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pgsfti;
CREATE EXTENSION unaccent;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION geohistorical_objects;

CREATE TABLE public.historic_places (
  place_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY
) INHERITS (geohistorical_object.geohistorical_object);

SELECT geohistorical_object.register_geohistorical_object_table(
  'public', 'historic_places'
);
```

Applications first populate `historical_source` and `numerical_origin_process`, including a `default_fuzzy_date` of type `sfti` and a `default_spatial_precision` JSON object containing a nonnegative `default` key. Rows belong in child tables, not the parent `geohistorical_object` table. The registration helper adds foreign keys and B-tree, GIN, and GiST indexes to a child.

The extension also supplies SFTI conversion functions and casts, relation tables, name-normalization helpers, and fuzzy-date distance helpers. Its design uses PostgreSQL table inheritance rather than declarative partitioning; uniqueness and foreign-key behavior do not automatically span every child.

### Installation Risk

The version 1.0 install script contains unconditional `DROP TABLE ... CASCADE`, creates broadly named casts and helper functions, and embeds test-style statements. Installing or reinstalling it can remove existing objects in the target schema. The registration function builds dynamic DDL from text identifiers without identifier quoting.

Review the complete SQL file and rehearse installation, registration, dump/restore, and removal in a disposable database. Take a tested backup before using it with existing data, pass only trusted simple identifiers to the registration function, and validate the old upstream compatibility claim on the exact PostgreSQL/PostGIS/pgsfti versions in use.
