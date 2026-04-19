---
title: "postgis_tiger_geocoder"
linkTitle: "postgis_tiger_geocoder"
description: "PostGIS tiger geocoder and reverse geocoder"
weight: 1504
categories: ["GIS"]
width: full
---

[**postgis**](https://git.osgeo.org/gitea/postgis/postgis) : PostGIS tiger geocoder and reverse geocoder


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1504** | {{< badge content="postgis_tiger_geocoder" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "postgis_tiger_geocoder" "postgis" >}} | `3.6.3` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `tiger` |
|   **Requires**    | {{< ext "postgis" >}} {{< ext "fuzzystrmatch" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "address_standardizer" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postgis` | `postgis`, `fuzzystrmatch` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "postgis36_18" "green" >}} {{< bg "17" "postgis36_17" "green" >}} {{< bg "16" "postgis36_16" "green" >}} {{< bg "15" "postgis36_15" "green" >}} {{< bg "14" "postgis36_14" "green" >}} | `postgis36_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "postgresql-18-postgis-3" "green" >}} {{< bg "17" "postgresql-17-postgis-3" "green" >}} {{< bg "16" "postgresql-16-postgis-3" "green" >}} {{< bg "15" "postgresql-15-postgis-3" "green" >}} {{< bg "14" "postgresql-14-postgis-3" "green" >}} | `postgresql-$v-postgis-3` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 3" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 3" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 6" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgis36_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_15 : AVAIL 5" "blue" >}} | {{< bg "PGDG 3.6.3" "postgis36_14 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://git.osgeo.org/gitea/postgis/postgis" title="Repository" icon="link" subtitle="git.osgeo.org/gitea/postgis/postgis" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install postgis;		# install via package name, for the active PG version
pig install postgis_tiger_geocoder;		# install by extension name, for the current active PG version

pig install postgis_tiger_geocoder -v 18;   # install for PG 18
pig install postgis_tiger_geocoder -v 17;   # install for PG 17
pig install postgis_tiger_geocoder -v 16;   # install for PG 16
pig install postgis_tiger_geocoder -v 15;   # install for PG 15
pig install postgis_tiger_geocoder -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION postgis_tiger_geocoder CASCADE; -- requires postgis, fuzzystrmatch
```



## Usage

> [PostGIS TIGER Geocoder: US Census TIGER/Line geocoding for PostGIS](https://github.com/postgis/postgis)

The PostGIS TIGER Geocoder provides geocoding and reverse geocoding capabilities for US addresses using US Census TIGER/Line data. It can parse an address string into a normalized form, find the geographic coordinates, and reverse-geocode coordinates back to an address.

- [TIGER Geocoder Reference](https://postgis.net/docs/Extras.html)

### Setup

```sql
CREATE EXTENSION postgis_tiger_geocoder CASCADE;
```

This creates the `tiger` schema with the geocoder tables and functions.

--------

## Loading TIGER Data

Before geocoding, TIGER/Line data must be loaded for the states you need. The extension provides helper functions to generate the loading scripts:

```sql
-- Generate a script to download and load data for a state
-- (e.g., Massachusetts = 'MA')
SELECT loader_generate_script(ARRAY['MA'], 'sh');
```

This generates a shell script that uses `shp2pgsql` to load TIGER shapefiles. Run the generated script to populate the `tiger_data` schema with address ranges, edges, faces, and other data.

After loading:

```sql
-- Install missing indexes for performance
SELECT install_missing_indexes();

-- Update statistics
ANALYZE tiger.addr;
ANALYZE tiger.edges;
ANALYZE tiger.faces;
```

--------

## Geocoding

Convert a US address string to geographic coordinates:

```sql
-- Basic geocoding
SELECT g.rating, ST_X(g.geomout) AS lon, ST_Y(g.geomout) AS lat,
       pprint_addy(g.addy) AS address
FROM geocode('1600 Pennsylvania Ave NW, Washington, DC 20500') AS g;
```

The `rating` indicates match quality (lower is better, 0 = exact match).

```sql
-- Geocode with a limit on results
SELECT g.rating, ST_AsText(g.geomout), pprint_addy(g.addy)
FROM geocode('100 Main St, Boston, MA', 3) AS g;

-- Batch geocode from a table
SELECT a.id, g.rating, g.geomout, pprint_addy(g.addy)
FROM addresses a, LATERAL geocode(a.address_string, 1) AS g;
```

--------

## Reverse Geocoding

Convert coordinates back to a street address:

```sql
SELECT pprint_addy(r.addy[1]) AS address
FROM reverse_geocode(ST_SetSRID(ST_MakePoint(-77.0365, 38.8977), 4326)) AS r;
```

--------

## Address Normalization

Parse and normalize address strings without geocoding:

```sql
SELECT *
FROM normalize_address('1600 Pennsylvania Avenue NW, Washington, DC 20500');
```

Returns components: `address` (number), `predirAbbrev`, `streetName`, `streetTypeAbbrev`, `postdirAbbrev`, `internal`, `location` (city), `stateAbbrev`, `zip`, `parsed`.

```sql
-- Pretty-print a normalized address
SELECT pprint_addy(normalize_address('100 main street boston ma 02101'));
```

--------

## Configuration

The `tiger.geocode_settings` table controls geocoder behavior:

```sql
-- View current settings
SELECT * FROM tiger.geocode_settings;

-- Adjust settings (e.g., increase debug level)
UPDATE tiger.geocode_settings SET val = 'true' WHERE name = 'debug_geocode_address';
```
