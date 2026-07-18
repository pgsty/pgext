## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/readme.md)
- [Version 1.1.0 SQL API](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/pg_osgr--1.1.0.sql)
- [Grid-reference conversion documentation](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/docs/function/to_gridref.md)
- [Extension control file](https://github.com/Edwards-R/pg_osgr/blob/253fa94dea4ae2e13ae3ec353c3fa9c42459dc25/pg_osgr.control)

`pg_osgr` is a pure-PL/pgSQL toolkit for converting Ordnance Survey grid references to easting, northing, datum, and accuracy, and for constructing references from those components.

```sql
CREATE EXTENSION pg_osgr;
SELECT osgr_find_datum('SU387148');
SELECT osgr_process_easting('SU387148'), osgr_process_northing('SU387148');
SELECT osgr_to_gridref(438700, 114800, 100, 27700);
```

The generic functions detect or accept EPSG 27700 for Great Britain, EPSG 29901 for Ireland and Northern Ireland, and EPSG 32630 for the Channel Islands. Resolution is limited to powers of ten and does not go below one metre.

The extension parses and formats grid references; it does not perform general coordinate-system transformations or datum shifts. Keep the EPSG code with stored coordinates, validate references at coverage boundaries, and use PostGIS/PROJ when converting between coordinate reference systems. Upstream states the rules were current on 2023-08-01, so verify later specification changes for regulated datasets.
