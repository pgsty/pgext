---
title: "address_standardizer_data_us"
linkTitle: "address_standardizer_data_us"
description: "Address Standardizer US dataset example"
weight: 1506
categories: ["GIS"]
languages: ["SQL"]
licenses: ["GPL-2.0"]
repos: ["PGDG"]
page_width: full
---

[**address_standardizer**](https://github.com/postgis/address_standardizer) : Address Standardizer US dataset example


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1506** | {{< badge content="address_standardizer_data_us" link="https://github.com/postgis/address_standardizer" >}} | {{< ext "address_standardizer_data_us" "address_standardizer" >}} | `3.7` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "nominatim_fdw" >}} {{< ext "country" >}} {{< ext "geoip" >}} {{< ext "tzf" >}} {{< ext "pg_geohash" >}} {{< ext "ogr_fdw" >}} |
|    **Siblings**   | {{< ext "address_standardizer" >}} |

> [!Note] Standalone with address_standardizer since PostGIS 3.7; control 3.7/package 3.7.0; no local recipe.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.7` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `address_standardizer` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.7.0` | {{< bg "18" "address_standardizer_18" "green" >}} {{< bg "17" "address_standardizer_17" "green" >}} {{< bg "16" "address_standardizer_16" "green" >}} {{< bg "15" "address_standardizer_15" "green" >}} {{< bg "14" "address_standardizer_14" "green" >}} | `address_standardizer_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `3.7.0` | {{< bg "18" "postgresql-18-address-standardizer" "green" >}} {{< bg "17" "postgresql-17-address-standardizer" "green" >}} {{< bg "16" "postgresql-16-address-standardizer" "green" >}} {{< bg "15" "postgresql-15-address-standardizer" "green" >}} {{< bg "14" "postgresql-14-address-standardizer" "green" >}} | `postgresql-$v-address-standardizer` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "MISS" "address_standardizer_18 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_17 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_16 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_15 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_14 : MISS 0" "red" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "MISS" "address_standardizer_18 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_17 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_16 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_15 : MISS 0" "red" >}} | {{< bg "MISS" "address_standardizer_14 : MISS 0" "red" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "address_standardizer_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 3.7.0" "postgresql-18-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-17-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-16-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-15-address-standardizer : AVAIL 1" "blue" >}} | {{< bg "PGDG 3.7.0" "postgresql-14-address-standardizer : AVAIL 1" "blue" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/postgis/address_standardizer" title="Repository" icon="github" subtitle="github.com/postgis/address_standardizer" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install address_standardizer;		# install via package name, for the active PG version
pig install address_standardizer_data_us;		# install by extension name, for the current active PG version

pig install address_standardizer_data_us -v 18;   # install for PG 18
pig install address_standardizer_data_us -v 17;   # install for PG 17
pig install address_standardizer_data_us -v 16;   # install for PG 16
pig install address_standardizer_data_us -v 15;   # install for PG 15
pig install address_standardizer_data_us -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION address_standardizer_data_us;
```




## Usage

> [Address Standardizer Data US: US address data for the address_standardizer extension](https://github.com/postgis/postgis)

This extension provides pre-built US lexicon, gazetteer, and rules data for use with the `address_standardizer` extension. It includes tables with common US street types, directional abbreviations, state names, and grammar rules needed to parse US addresses.

- [Address Standardizer Reference](https://postgis.net/docs/Extras.html#Address_Standardizer)

### Setup

```sql
CREATE EXTENSION address_standardizer_data_us;
```

This creates the `us_lex`, `us_gaz`, and `us_rules` tables in the public schema, pre-populated with US address data.

--------

## Using with address_standardizer

Once installed, you can immediately standardize US addresses:

```sql
SELECT *
FROM standardize_address(
    'us_lex', 'us_gaz', 'us_rules',
    '123 Main Street NW, Apt 4B, Springfield, IL 62704'
);
```

The provided data covers common US address patterns including:

- Street types (Street, Avenue, Boulevard, Drive, Lane, Court, etc.)
- Directional prefixes and suffixes (North, South, N, S, NW, SE, etc.)
- State names and abbreviations
- Unit designators (Apt, Suite, Unit, etc.)
- Highway designators (US, State, County, Interstate, etc.)
