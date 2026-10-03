---
title: "mobilitydb_datagen"
linkTitle: "mobilitydb_datagen"
description: "MobilityDB random data generator functions"
weight: 1651
categories: ["GIS"]
languages: ["SQL"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**mobilitydb**](https://github.com/MobilityDB/MobilityDB) : MobilityDB random data generator functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1651** | {{< badge content="mobilitydb_datagen" link="https://github.com/MobilityDB/MobilityDB" >}} | {{< ext "mobilitydb_datagen" "mobilitydb" >}} | `1.3.1` | {{< category "GIS" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "mobilitydb" >}} |
|   **See Also**    | {{< ext "pgrouting" >}} {{< ext "faker" >}} {{< ext "random" >}} {{< ext "h3" >}} {{< ext "pg_polyline" >}} {{< ext "dbt2" >}} {{< ext "tsm_system_time" >}} {{< ext "tsm_system_rows" >}} |
|    **Siblings**   | {{< ext "mobilitydb" >}} |

> [!Note] Pigsty 1.3.1 bundles MobilityDB DataGen; upgrading from 1.2 to 1.3 requires upstream backup/restore.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `mobilitydb` | `mobilitydb` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "mobilitydb_18" "green" >}} {{< bg "17" "mobilitydb_17" "green" >}} {{< bg "16" "mobilitydb_16" "green" >}} {{< bg "15" "mobilitydb_15" "green" >}} {{< bg "14" "mobilitydb_14" "green" >}} | `mobilitydb_$v` | `postgis36_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.3.1` | {{< bg "18" "postgresql-18-mobilitydb" "green" >}} {{< bg "17" "postgresql-17-mobilitydb" "green" >}} {{< bg "16" "postgresql-16-mobilitydb" "green" >}} {{< bg "15" "postgresql-15-mobilitydb" "green" >}} {{< bg "14" "postgresql-14-mobilitydb" "green" >}} | `postgresql-$v-mobilitydb` | `postgresql-$v-postgis-3` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "mobilitydb_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 2" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 2" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-18-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-17-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-16-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-15-mobilitydb : AVAIL 4" "green" >}} | {{< bg "PIGSTY 1.3.1" "postgresql-14-mobilitydb : AVAIL 4" "green" >}} |
{.matrix}


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/MobilityDB/MobilityDB" title="Repository" icon="github" subtitle="github.com/MobilityDB/MobilityDB" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="mobilitydb-1.3.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg mobilitydb;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install mobilitydb;		# install via package name, for the active PG version
pig install mobilitydb_datagen;		# install by extension name, for the current active PG version

pig install mobilitydb_datagen -v 18;   # install for PG 18
pig install mobilitydb_datagen -v 17;   # install for PG 17
pig install mobilitydb_datagen -v 16;   # install for PG 16
pig install mobilitydb_datagen -v 15;   # install for PG 15
pig install mobilitydb_datagen -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION mobilitydb_datagen CASCADE; -- requires mobilitydb
```

## Usage

Sources:

- [Version 1.3.1 data-generator manual](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/data_generator.xml)
- [Extension control file](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/datagen/mobilitydb_datagen.in.control)
- [Temporal generators](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/datagen/temporal/random_temporal.sql)
- [Temporal point generators](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/datagen/geo/random_tpoint.sql)

`mobilitydb_datagen` 1.3.1 provides PL/pgSQL functions for generating synthetic PostgreSQL, PostGIS, and MobilityDB values. It is mainly useful for regression data, demos, and benchmark fixtures that need random temporal values or trajectories.

```sql
-- After the main MobilityDB extension is loaded:
CREATE EXTENSION mobilitydb_datagen;
```

### Generating Random Temporal Values

```sql
-- A random temporal float sequence.
SELECT random_tfloat_seq(
    -100.0, 100.0,                                  -- value bounds
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',  -- time bounds
    10.0,                                           -- max value delta
    10,                                             -- max minutes between instants
    5, 10                                           -- min/max instants
);

-- Step interpolation instead of the default linear interpolation.
SELECT random_tfloat_seq(
    -100.0, 100.0,
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',
    10.0, 10, 5, 10,
    false
);

-- A random temporal geometry point sequence.
SELECT asEWKT(
    random_tgeompoint_contseq(
        2.20, 2.50, 48.80, 48.95,                  -- x/y bounds
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        0.02, 5, 20, 40,                           -- max delta, max minutes, min/max instants
        srid => 4326
    )
);
```

Other confirmed generator families include scalar helpers such as `random_bool`, `random_int`, `random_float`, `random_text`, and `random_timestamptz`; array, set, span, and range helpers; temporal helpers such as `random_tbool_inst`, `random_tint_discseq`, `random_tfloat_seq`, and `random_tfloat_seqset`; and spatial/temporal-point helpers such as `random_geom_point`, `random_geom_linestring`, `random_tgeompoint_contseq`, `random_tgeompoint_seqset`, `random_tgeogpoint_contseq`, and `random_tgeogpoint_seqset`.

### Generating Test Datasets

Create bulk test data for benchmarking trip queries:

```sql
CREATE TABLE trip_samples AS
SELECT
    vehicle_id,
    random_tgeompoint_contseq(
        2.20, 2.50, 48.80, 48.95,
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        0.02, 5, 20, 40,
        srid => 4326
    ) AS trip
FROM generate_series(1, 1000) AS vehicle_id;
```

### Caveats

- The control file requires the main `mobilitydb` extension; `mobilitydb_datagen` is not standalone.
- The generator is distributed with `mobilitydb` version `1.3.1`. Keep the dependency and generator files aligned; build flags and package availability determine which optional type families are present.
- Upstream docs intentionally omit detailed parameter lists for many generator functions and point users to the SQL source files for exact signatures.

The generator functions create synthetic values, not reproducible application datasets by default. Set a deliberate random seed when comparing runs, and review bounds, time zones, interpolation, and SRID for the workload.
