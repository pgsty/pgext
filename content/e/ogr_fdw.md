---
title: "ogr_fdw"
linkTitle: "ogr_fdw"
description: "foreign-data wrapper for GIS data access"
weight: 1550
categories: ["GIS"]
width: full
---

[**ogr_fdw**](https://github.com/pramsey/pgsql-ogr-fdw) : foreign-data wrapper for GIS data access


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1550** | {{< badge content="ogr_fdw" link="https://github.com/pramsey/pgsql-ogr-fdw" >}} | {{< ext "ogr_fdw" >}} | `1.1.9` | {{< category "GIS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "postgis" >}} {{< ext "file_fdw" >}} {{< ext "postgres_fdw" >}} {{< ext "postgis_topology" >}} {{< ext "postgis_raster" >}} {{< ext "postgis_sfcgal" >}} {{< ext "postgis_tiger_geocoder" >}} {{< ext "address_standardizer" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.9` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `ogr_fdw` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.9` | {{< bg "18" "ogr_fdw_18" "green" >}} {{< bg "17" "ogr_fdw_17" "green" >}} {{< bg "16" "ogr_fdw_16" "green" >}} {{< bg "15" "ogr_fdw_15" "green" >}} {{< bg "14" "ogr_fdw_14" "green" >}} | `ogr_fdw_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.1.9` | {{< bg "18" "postgresql-18-ogr-fdw" "green" >}} {{< bg "17" "postgresql-17-ogr-fdw" "green" >}} {{< bg "16" "postgresql-16-ogr-fdw" "green" >}} {{< bg "15" "postgresql-15-ogr-fdw" "green" >}} {{< bg "14" "postgresql-14-ogr-fdw" "green" >}} | `postgresql-$v-ogr-fdw` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_14 : AVAIL 11" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_18 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_17 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_15 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_14 : AVAIL 7" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_16 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_15 : AVAIL 12" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_14 : AVAIL 13" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_17 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_16 : AVAIL 8" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_14 : AVAIL 9" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_18 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_17 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_16 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.1.9" "ogr_fdw_14 : AVAIL 6" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_15 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.1.7" "ogr_fdw_14 : AVAIL 4" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.1.9" "postgresql-18-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-17-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-16-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-15-ogr-fdw : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.1.9" "postgresql-14-ogr-fdw : AVAIL 3" "blue" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_18` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.3 KiB | [ogr_fdw_18-1.1.9-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ogr_fdw_18-1.1.9-1PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/ogr_fdw_18-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.3 KiB | [ogr_fdw_18-1.1.9-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ogr_fdw_18-1.1.9-1PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/ogr_fdw_18-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_18-1.1.9-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ogr_fdw_18-1.1.9-1PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel9.7.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel9.6.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/ogr_fdw_18-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_18-1.1.9-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ogr_fdw_18-1.1.9-1PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel9.7.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.6 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel9.6.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/ogr_fdw_18-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_18-1.1.9-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ogr_fdw_18-1.1.9-1PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.1 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.2 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel10.1.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ogr_fdw_18-1.1.7-6PGDG.rhel10.0.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/ogr_fdw_18-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel10.2.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel10.1.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_18-1.1.7-6PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ogr_fdw_18-1.1.7-6PGDG.rhel10.0.aarch64.rpm) |
| `ogr_fdw_18` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_18-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/ogr_fdw_18-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.3 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.3 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.5 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.5 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.5 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.5 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.7 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.8 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.7 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.3 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.3 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.3 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 92.4 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 92.4 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 92.4 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 90.4 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 89.9 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 89.8 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.7 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.6 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.7 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.8 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.9 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.9 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.8 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.4 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.4 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.9` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.5 KiB | [postgresql-18-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.2 KiB | [postgresql-18-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb) |
| `postgresql-18-ogr-fdw` | `1.1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.3 KiB | [postgresql-18-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-18-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_17` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.2 KiB | [ogr_fdw_17-1.1.9-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.9-1PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/ogr_fdw_17-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.3 KiB | [ogr_fdw_17-1.1.9-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.9-1PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.3 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.8 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/ogr_fdw_17-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_17-1.1.9-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.9-1PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel9.7.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel9.6.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_17-1.1.5-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/ogr_fdw_17-1.1.5-3PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.9-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.9-1PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel9.7.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.6 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel9.6.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.7 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [ogr_fdw_17-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_17-1.1.5-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/ogr_fdw_17-1.1.5-3PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.3 KiB | [ogr_fdw_17-1.1.9-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.9-1PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.2 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel10.1.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.7-6PGDG.rhel10.0.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_17-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/ogr_fdw_17-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel10.2.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel10.1.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_17-1.1.7-6PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ogr_fdw_17-1.1.7-6PGDG.rhel10.0.aarch64.rpm) |
| `ogr_fdw_17` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_17-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/ogr_fdw_17-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.1 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.1 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.2 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.4 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.2 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.4 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.5 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.4 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.4 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.1 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.0 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.0 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.5 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.5 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.4 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.1 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.1 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.1 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.5 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.6 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.4 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.0 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.0 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.3 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.7 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.5 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.4 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.9` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.1 KiB | [postgresql-17-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.3 KiB | [postgresql-17-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb) |
| `postgresql-17-ogr-fdw` | `1.1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 86.9 KiB | [postgresql-17-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-17-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_16` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.2 KiB | [ogr_fdw_16-1.1.9-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.9-1PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.0 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.5 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.2 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/ogr_fdw_16-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.4 KiB | [ogr_fdw_16-1.1.9-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.9-1PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.8 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.5 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/ogr_fdw_16-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_16-1.1.9-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.9-1PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel9.7.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel9.6.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.3 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_16-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 48.1 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/ogr_fdw_16-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.7 KiB | [ogr_fdw_16-1.1.9-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.9-1PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.5 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel9.7.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.6 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel9.6.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.7 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.8 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.4 KiB | [ogr_fdw_16-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 46.9 KiB | [ogr_fdw_16-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/ogr_fdw_16-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.3 KiB | [ogr_fdw_16-1.1.9-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.9-1PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.1 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.1 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel10.1.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.7-6PGDG.rhel10.0.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.4 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_16-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/ogr_fdw_16-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel10.2.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel10.1.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_16-1.1.7-6PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ogr_fdw_16-1.1.7-6PGDG.rhel10.0.aarch64.rpm) |
| `ogr_fdw_16` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 49.7 KiB | [ogr_fdw_16-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/ogr_fdw_16-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.1 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.1 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.2 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.4 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.6 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.1 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.6 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.5 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 90.6 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.1 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 88.9 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.2 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.2 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.1 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 106.0 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.9 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.7 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 103.7 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.6 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.7 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 89.6 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.5 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.7 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 87.5 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.7 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.5 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 88.6 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.9` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.0 KiB | [postgresql-16-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 86.9 KiB | [postgresql-16-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb) |
| `postgresql-16-ogr-fdw` | `1.1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 87.1 KiB | [postgresql-16-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-16-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_15` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.8 KiB | [ogr_fdw_15-1.1.9-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.9-1PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.6 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.5 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.4-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_15-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.3-1.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_15-1.1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/ogr_fdw_15-1.1.2-2.rhel8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.9 KiB | [ogr_fdw_15-1.1.9-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.9-1PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.7 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.9 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/ogr_fdw_15-1.1.4-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.2 KiB | [ogr_fdw_15-1.1.9-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.9-1PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.8 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.9 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel9.7.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel9.6.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.9 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_15-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.4-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.8 KiB | [ogr_fdw_15-1.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.3-1.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.9 KiB | [ogr_fdw_15-1.1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/ogr_fdw_15-1.1.2-2.rhel9.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_15-1.1.9-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.9-1PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel9.7.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel9.6.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_15-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_15-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_15-1.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/ogr_fdw_15-1.1.4-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.9 KiB | [ogr_fdw_15-1.1.9-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.9-1PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel10.1.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.2 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.7-6PGDG.rhel10.0.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_15-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/ogr_fdw_15-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.8 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel10.2.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.8 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel10.1.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.8 KiB | [ogr_fdw_15-1.1.7-6PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ogr_fdw_15-1.1.7-6PGDG.rhel10.0.aarch64.rpm) |
| `ogr_fdw_15` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.0 KiB | [ogr_fdw_15-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/ogr_fdw_15-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.9 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.5 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.0 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.1 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.0 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.9 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.4 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.3 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.4 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.6 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.6 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.6 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.1 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.0 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.0 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.6 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.5 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.6 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.8 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.8 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.7 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.6 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.7 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.4 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [u26.x86_64](/os/u26.x86_64) | pgdg | 89.5 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 89.4 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 89.4 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.9` | [u26.aarch64](/os/u26.aarch64) | pgdg | 88.2 KiB | [postgresql-15-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 88.2 KiB | [postgresql-15-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb) |
| `postgresql-15-ogr-fdw` | `1.1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 88.0 KiB | [postgresql-15-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-15-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ogr_fdw_14` | `1.1.9` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.8 KiB | [ogr_fdw_14-1.1.9-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.9-1PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.6 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel8.10.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.7-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.6-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el8.x86_64](/os/el8.x86_64) | pgdg | 51.4 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.5-4PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.4-2PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.4 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.4-1PGDG.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.3-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.3-1.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.2-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.2-2.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 49.5 KiB | [ogr_fdw_14-1.1.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.2-1.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.7 KiB | [ogr_fdw_14-1.1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/ogr_fdw_14-1.1.1-1.rhel8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.9` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.9 KiB | [ogr_fdw_14-1.1.9-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.9-1PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.7 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel8.10.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.7-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.9 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.6-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el8.aarch64](/os/el8.aarch64) | pgdg | 49.4 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.5-4PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.4-2PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/ogr_fdw_14-1.1.4-1PGDG.rhel8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.9` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_14-1.1.9-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.9-1PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.0 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel9.8.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.0 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel9.7.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel9.6.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.1 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.7-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el9.x86_64](/os/el9.x86_64) | pgdg | 51.3 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.6-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.5 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.5-4PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.7 KiB | [ogr_fdw_14-1.1.5-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.5-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.2 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.4-2PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.1 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.4-1PGDG.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.3` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_14-1.1.3-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.3-1.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 49.6 KiB | [ogr_fdw_14-1.1.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.2-2.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 50.0 KiB | [ogr_fdw_14-1.1.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/ogr_fdw_14-1.1.2-1.rhel9.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.9` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.8 KiB | [ogr_fdw_14-1.1.9-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.9-1PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel9.8.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.0 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel9.7.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel9.6.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.2 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.7-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.1 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.6-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.5` | [el9.aarch64](/os/el9.aarch64) | pgdg | 49.5 KiB | [ogr_fdw_14-1.1.5-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.5-4PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.3 KiB | [ogr_fdw_14-1.1.4-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.4-2PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 48.0 KiB | [ogr_fdw_14-1.1.4-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/ogr_fdw_14-1.1.4-1PGDG.rhel9.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.9` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.0 KiB | [ogr_fdw_14-1.1.9-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.9-1PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel10.2.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 51.8 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel10.1.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.2 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.7-6PGDG.rhel10.0.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.7-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.6` | [el10.x86_64](/os/el10.x86_64) | pgdg | 52.1 KiB | [ogr_fdw_14-1.1.6-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/ogr_fdw_14-1.1.6-1PGDG.rhel10.x86_64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.8 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel10.2.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.8 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel10.1.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 50.8 KiB | [ogr_fdw_14-1.1.7-6PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ogr_fdw_14-1.1.7-6PGDG.rhel10.0.aarch64.rpm) |
| `ogr_fdw_14` | `1.1.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 51.0 KiB | [ogr_fdw_14-1.1.7-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/ogr_fdw_14-1.1.7-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.9 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg12+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [d12.x86_64](/os/d12.x86_64) | pgdg | 91.2 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg12+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d12.x86_64](/os/d12.x86_64) | pgdg | 90.8 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg12+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.9 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg12+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [d12.aarch64](/os/d12.aarch64) | pgdg | 89.1 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg12+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d12.aarch64](/os/d12.aarch64) | pgdg | 88.8 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg12+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.3 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg13+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.3 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg13+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d13.x86_64](/os/d13.x86_64) | pgdg | 91.3 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg13+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.3 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg13+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.6 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg13+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [d13.aarch64](/os/d13.aarch64) | pgdg | 89.5 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg13+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.0 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.0 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u22.x86_64](/os/u22.x86_64) | pgdg | 107.0 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [u22.aarch64](/os/u22.aarch64) | pgdg | 105.0 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.5 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u22.aarch64](/os/u22.aarch64) | pgdg | 104.6 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.5 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.7 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u24.x86_64](/os/u24.x86_64) | pgdg | 90.6 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.6 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.6 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg24.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u24.aarch64](/os/u24.aarch64) | pgdg | 88.6 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg24.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [u26.x86_64](/os/u26.x86_64) | pgdg | 89.6 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [u26.x86_64](/os/u26.x86_64) | pgdg | 89.7 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg26.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u26.x86_64](/os/u26.x86_64) | pgdg | 89.3 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg26.04+1_amd64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.9` | [u26.aarch64](/os/u26.aarch64) | pgdg | 88.0 KiB | [postgresql-14-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.9-1.pgdg26.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.8` | [u26.aarch64](/os/u26.aarch64) | pgdg | 88.0 KiB | [postgresql-14-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.8-2.pgdg26.04+1_arm64.deb) |
| `postgresql-14-ogr-fdw` | `1.1.7` | [u26.aarch64](/os/u26.aarch64) | pgdg | 88.3 KiB | [postgresql-14-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-ogr-fdw/postgresql-14-ogr-fdw_1.1.7-3.pgdg26.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-ogr-fdw" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-ogr-fdw" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install ogr_fdw;		# install via package name, for the active PG version

pig install ogr_fdw -v 18;   # install for PG 18
pig install ogr_fdw -v 17;   # install for PG 17
pig install ogr_fdw -v 16;   # install for PG 16
pig install ogr_fdw -v 15;   # install for PG 15
pig install ogr_fdw -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION ogr_fdw;
```

## Usage

Sources:

- [ogr_fdw v1.1.9 README](https://github.com/pramsey/pgsql-ogr-fdw/blob/v1.1.9/README.md)
- [Extension control file](https://github.com/pramsey/pgsql-ogr-fdw/blob/v1.1.9/ogr_fdw.control)
- [v1.1.8 to v1.1.9 comparison](https://github.com/pramsey/pgsql-ogr-fdw/compare/v1.1.8...v1.1.9)
- [GDAL vector drivers](https://gdal.org/en/stable/drivers/vector/index.html)

`ogr_fdw` exposes vector data supported by GDAL/OGR as PostgreSQL foreign tables. It can read files and remote data sources through OGR drivers and can write when the selected driver and data source support updates. Install PostGIS before `ogr_fdw` for native geometry columns; otherwise geometry is exposed as WKB `bytea`.

### Discover and Import a Layer

Use the installed helper to inspect a source and generate matching SQL:

```console
ogr_fdw_info -s /srv/gis/cities.gpkg
ogr_fdw_info -s /srv/gis/cities.gpkg -l cities
```

A minimal equivalent definition is:

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ogr_fdw;

CREATE SERVER city_source
  FOREIGN DATA WRAPPER ogr_fdw
  OPTIONS (
    datasource '/srv/gis/cities.gpkg',
    format 'GPKG'
  );

CREATE FOREIGN TABLE city (
  fid bigint,
  geom geometry,
  name text
) SERVER city_source
  OPTIONS (layer 'cities');

SELECT fid, name FROM city WHERE geom && ST_MakeEnvelope(-10, 35, 30, 60, 4326);
```

The PostgreSQL server account needs filesystem permissions for file-backed data sources and network/credential access for remote drivers.

### Import and Mapping

```sql
CREATE SCHEMA gis_import;

IMPORT FOREIGN SCHEMA ogr_all
  LIMIT TO (cities)
  FROM SERVER city_source
  INTO gis_import;
```

`ogr_all` means all OGR layers. Import normally launders table and column names; use `launder_table_names` and `launder_column_names` options when exact remote names are required. A foreign column can map to a different source name with `OPTIONS (column_name 'RemoteName')`.

### Important Options and Objects

- Server options: required `datasource`, optional `format`, `updateable`, `config_options`, `open_options`, and `character_encoding`.
- Table options: `layer` identifies the OGR layer and `updateable` can disable writes.
- `fid` identifies a feature and is required for writable foreign tables.
- `ogr_fdw_info` lists drivers and layers and emits server/table definitions.
- `ogr_fdw_version()` reports the extension and GDAL version.
- `ogr_fdw_drivers()` lists the compiled OGR drivers.

### Performance and Write Boundaries

Simple comparisons and bounding-box `&&` predicates can be pushed down, but more complex filters may be evaluated locally. The FDW retrieves all selected source columns and opens two OGR connections per query rather than pooling them. Use `EXPLAIN`, project only needed columns, and benchmark the actual driver and data source.

Writes depend on driver capability and require source-level write permissions plus `fid`. Set `updateable = false` when a source must remain read-only. Version 1.1.9 simplifies the version string relative to 1.1.8 and has no documented SQL workflow change; the control file remains at SQL extension version 1.1.
