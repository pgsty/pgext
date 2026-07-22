## Usage

Sources:

- [Official README](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/README.md)
- [Extension SQL for version 1.0](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/postgis_domains--1.0.sql)
- [Extension control file](https://github.com/cahutton/postgis_domains/blob/cb9290194eebb0b63c85c7e59c406e5c0049f48e/postgis_domains.control)

`postgis_domains` defines reusable PostGIS domains for common WGS 84 geometry and geography shapes. Use them when a schema should reject values with the wrong geometry family, dimensions, SRID, validity, or simplicity at the column boundary instead of repeating those checks in every table.

### Core Workflow

PostGIS is a declared dependency, so installing this extension also requires `postgis` to be available. Choose the domain whose type and dimensional suffix match the intended data:

```sql
CREATE EXTENSION postgis_domains;

CREATE TABLE places (
  place_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  location point_geometry NOT NULL
);

INSERT INTO places (location)
VALUES (ST_SetSRID(ST_MakePoint(-73.9857, 40.7484), 4326));
```

An input is accepted only if it satisfies the domain's checks. This makes the constraint apply consistently to inserts, updates, casts, and every table that uses the domain.

### Domain Families

The extension defines geometry and geography counterparts for point, multipoint, linestring, multilinestring, polygon, multipolygon, and general geometry values. The base two-dimensional form uses names such as `point_geometry`, `linestring_geometry`, and `polygon_geography`; `m`, `z`, and `zm` variants encode the required coordinate dimensions, for example `pointz_geometry` and `polygonzm_geography`.

The SQL definitions combine checks appropriate to each domain:

- the PostGIS object family must match the domain;
- the coordinate dimension must match the plain, M, Z, or ZM variant;
- the SRID must be 4326;
- the value must satisfy the extension's validity and simplicity checks.

Inspect the exact installed definition before relying on a particular domain in generated DDL:

```sql
SELECT domain_name, data_type
FROM information_schema.domains
WHERE domain_schema = 'public'
  AND domain_name LIKE '%\_geometry' ESCAPE '\\'
ORDER BY domain_name;
```

### Operational Notes

Version 1.0 was developed with PostgreSQL 10 and PostGIS 2.5; the README only states that it should work with recent versions. That statement is not a tested compatibility matrix, so exercise the chosen domains against the exact PostgreSQL/PostGIS versions in production.

The control file declares `postgis` as a dependency, allows relocation, and does not require superuser installation or shared-library preload. Because domains participate in table column types, removing or replacing the extension requires addressing dependent columns first. All bundled domains enforce WGS 84 SRID 4326; transform source data explicitly rather than relabeling coordinates with a new SRID.
