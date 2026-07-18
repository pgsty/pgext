## Usage

Sources:

- [Extension README](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/README.md)
- [Extension control file](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/pg_address.control)
- [Version 0.0.3 SQL definition](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/sql/pg_address.sql)
- [Smoke test](https://github.com/lotint/dbtools/blob/3966ce98177cfa3b2e99a1871f08d6bb52a1a4ce/pg_extensions/pg_address/tests/smoke.sql)

`pg_address` 0.0.3 defines a composite type named `pg_address` for storing an address as country, region, city, postal code, street, number, suburb, latitude, and longitude fields.

### Store and read a composite address

```sql
CREATE EXTENSION pg_address;

CREATE TABLE customer_location (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    address pg_address
);

INSERT INTO customer_location (address)
VALUES (ROW('Germany', 'Berlin', 'Berlin', '10117',
            'Hauptstrasse', '57b', 'Mitte', 52.5200, 13.4050)::pg_address);

SELECT (address).city, (address).zip_code, (address).lat, (address).lon
FROM customer_location;
```

The field order is `country`, `region`, `city`, `zip_code`, `street`, `num`, `suburb`, `lat`, and `lon`. Use an explicit cast or named-field assignment so positional values do not silently drift.

### Data-model boundary

This is a storage type, not an address parser, geocoder, or validator. Text fields accept arbitrary values, coordinates have no range constraint, and no normalization or index operator class is supplied. Add application-specific constraints for latitude, longitude, country codes, and required fields.

Composite-type changes can affect dependent tables, functions, clients, and positional row literals. The repository's installation notes target PostgreSQL 9.6-era development packages and publish no current compatibility matrix; validate extension upgrades and client type decoding on the exact PostgreSQL release before adoption.
