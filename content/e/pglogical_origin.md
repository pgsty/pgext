---
title: "pglogical_origin"
linkTitle: "pglogical_origin"
description: "Dummy extension for compatibility when upgrading from Postgres 9.4"
weight: 9501
categories: ["ETL"]
width: full
---

Dummy extension for compatibility when upgrading from Postgres 9.4


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9501** | {{< badge content="pglogical_origin" link="https://github.com/2ndQuadrant/pglogical" >}} | {{< ext "pglogical_origin" "pglogical" >}} | `2.4.5` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pglogical_ticker" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "pg_failover_slots" >}} {{< ext "pgactive" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "repmgr" >}} {{< ext "decoder_raw" >}} |
|    **Siblings**   | {{< ext "pglogical" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/pglogical" >}} | `2.4.5` | {{< bg "18" "pglogical_18*" "red" >}} {{< bg "17" "pglogical_17*" "green" >}} {{< bg "16" "pglogical_16*" "green" >}} {{< bg "15" "pglogical_15*" "green" >}} {{< bg "14" "pglogical_14*" "green" >}} | `pglogical_$v*` | - |
| **Debian** | {{< badge content="PGDG" link="/e/pglogical" >}} | `2.4.5` | {{< bg "18" "postgresql-18-pglogical" "red" >}} {{< bg "17" "postgresql-17-pglogical" "green" >}} {{< bg "16" "postgresql-16-pglogical" "green" >}} {{< bg "15" "postgresql-15-pglogical" "green" >}} {{< bg "14" "postgresql-14-pglogical" "green" >}} | `postgresql-$v-pglogical` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 4" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 3" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 2.4.6" "pglogical_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.5" "pglogical_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.4" "pglogical_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 2.4.3" "pglogical_14 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
|    `d12.aarch64`    | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
|    `u22.x86_64`    | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
|    `u22.aarch64`    | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
|    `u24.x86_64`    | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |
|    `u24.aarch64`    | {{< bg "PGDG 2.4.6" "postgresql-18-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-17-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-16-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-15-pglogical : AVAIL 1" "blue" >}} | {{< bg "PGDG 2.4.6" "postgresql-14-pglogical : AVAIL 1" "blue" >}} |


## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/2ndQuadrant/pglogical" title="Repository" icon="github" subtitle="github.com/2ndQuadrant/pglogical" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglogical_ticker-1.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pglogical_origin; # get pglogical_origin source code
pig build dep pglogical_origin; # install build dependencies
pig build pkg pglogical_origin; # build extension rpm or deb
pig build ext pglogical_origin; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pglogical_origin; # install by extension name, for the current active PG version
pig ext install pglogical; # install via package alias, for the active PG version
pig ext install pglogical_origin -v 17;   # install for PG 17
pig ext install pglogical_origin -v 16;   # install for PG 16
pig ext install pglogical_origin -v 15;   # install for PG 15
pig ext install pglogical_origin -v 14;   # install for PG 14
pig ext install pglogical_origin -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pglogical_origin CASCADE SCHEMA pglogical_origin;
```

