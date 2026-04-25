---
title: "address_standardizer"
linkTitle: "address_standardizer"
description: "Used to parse an address into constituent elements. Generally used to support geocoding address normalization step."
weight: 1505
categories: ["GIS"]
width: full
---

[**postgis**](https://git.osgeo.org/gitea/postgis/postgis) : Used to parse an address into constituent elements. Generally used to support geocoding address normalization step.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1505** | {{< badge content="address_standardizer" link="https://git.osgeo.org/gitea/postgis/postgis" >}} | {{< ext "address_standardizer" "postgis" >}} | `3.6.3` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "pointcloud" >}} {{< ext "pointcloud_postgis" >}} {{< ext "h3" >}} {{< ext "h3_postgis" >}} {{< ext "q3c" >}} {{< ext "ogr_fdw" >}} {{< ext "geoip" >}} |
|    **Siblings**   | {{< ext "postgis" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer_data_us" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.6.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `postgis` | - |
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
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 3.6.3" "postgresql-18-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-17-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-16-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-15-postgis-3 : AVAIL 2" "blue" >}} | {{< bg "PGDG 3.6.3" "postgresql-14-postgis-3 : AVAIL 2" "blue" >}} |


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
pig install address_standardizer;		# install by extension name, for the current active PG version

pig install address_standardizer -v 18;   # install for PG 18
pig install address_standardizer -v 17;   # install for PG 17
pig install address_standardizer -v 16;   # install for PG 16
pig install address_standardizer -v 15;   # install for PG 15
pig install address_standardizer -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION address_standardizer;
```



## Usage

> [Address Standardizer: Address parsing and standardization for PostGIS](https://github.com/postgis/postgis)

The Address Standardizer is a PostGIS extension that parses a single-line address string into a structured form using configurable lexicon, grammar, and rules tables. It is a more flexible alternative to the built-in `normalize_address` function in the TIGER geocoder.

- [Address Standardizer Reference](https://postgis.net/docs/Extras.html#Address_Standardizer)

### Setup

```sql
CREATE EXTENSION address_standardizer;
```

--------

## Standardizing Addresses

The core function takes an address string and three table references (lex, gaz, rules):

```sql
SELECT *
FROM standardize_address(
    'us_lex',        -- lexicon table
    'us_gaz',        -- gazetteer table
    'us_rules',      -- rules table
    '1600 Pennsylvania Ave NW, Washington, DC 20500'
);
```

The result contains structured fields:

| Field | Description |
|-------|-------------|
| `building` | Building name or identifier |
| `house_num` | Street number |
| `predir` | Prefix direction (N, S, E, W) |
| `qual` | Qualifier |
| `pretype` | Prefix type |
| `name` | Street name |
| `suftype` | Suffix type (St, Ave, Blvd) |
| `sufdir` | Suffix direction |
| `ruralroute` | Rural route |
| `extra` | Extra information |
| `city` | City name |
| `state` | State |
| `country` | Country |
| `postcode` | ZIP/postal code |
| `box` | PO Box |
| `unit` | Unit/apartment number |

--------

## Lexicon, Gazetteer, and Rules Tables

The standardizer is driven by three user-configurable tables:

**Lexicon (lex)** -- Maps input tokens to standardized forms and token classes:

```sql
CREATE TABLE us_lex (
    id serial PRIMARY KEY,
    seq integer,
    word text,
    stdword text,
    token integer
);
```

**Gazetteer (gaz)** -- Maps place names (cities, states) to standard forms:

```sql
CREATE TABLE us_gaz (
    id serial PRIMARY KEY,
    seq integer,
    word text,
    stdword text,
    token integer
);
```

**Rules (rules)** -- Defines grammar rules for parsing addresses:

```sql
CREATE TABLE us_rules (
    id serial PRIMARY KEY,
    rule text
);
```

For US addresses, the `address_standardizer_data_us` extension provides pre-built data for these tables.
