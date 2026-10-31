---
title: "postgis_tiger_geocoder"
linkTitle: "postgis_tiger_geocoder"
description: "PostGIS tiger geocoder and reverse geocoder"
weight: 1504
categories: ["GIS"]
languages: ["PLpgSQL"]
licenses: ["GPL-2.0"]
repos: ["PGDG"]
page_width: full
---

[**postgis_tiger_geocoder**](https://gitea.osgeo.org/postgis/postgis_tiger_geocoder) : PostGIS tiger geocoder and reverse geocoder


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1504** | {{< badge content="postgis_tiger_geocoder" link="https://gitea.osgeo.org/postgis/postgis_tiger_geocoder" >}} | {{< ext "postgis_tiger_geocoder" >}} | `2025.2` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "PLpgSQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `tiger` |
|   **Requires**    | {{< ext "postgis" >}} {{< ext "fuzzystrmatch" >}} |
|   **See Also**    | {{< ext "nominatim_fdw" >}} {{< ext "geoip" >}} {{< ext "tzf" >}} {{< ext "pg_geohash" >}} {{< ext "ogr_fdw" >}} |

> [!Note] Standalone since PostGIS 3.7; control/package 2025.2; PG16+; no local recipe.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2025.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `postgis_tiger_geocoder` | `postgis`, `fuzzystrmatch` |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2025.2` | {{< bg "18" "postgis_tiger_geocoder_18" "green" >}} {{< bg "17" "postgis_tiger_geocoder_17" "green" >}} {{< bg "16" "postgis_tiger_geocoder_16" "green" >}} {{< bg "15" "postgis_tiger_geocoder_15" "red" >}} {{< bg "14" "postgis_tiger_geocoder_14" "red" >}} | `postgis_tiger_geocoder_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `2025.2` | {{< bg "18" "postgresql-18-postgis-tiger-geocoder" "green" >}} {{< bg "17" "postgresql-17-postgis-tiger-geocoder" "green" >}} {{< bg "16" "postgresql-16-postgis-tiger-geocoder" "green" >}} {{< bg "15" "postgresql-15-postgis-tiger-geocoder" "red" >}} {{< bg "14" "postgresql-14-postgis-tiger-geocoder" "red" >}} | `postgresql-$v-postgis-tiger-geocoder` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "MISS" "postgis_tiger_geocoder_18 : MISS 0" "red" >}} | {{< bg "MISS" "postgis_tiger_geocoder_17 : MISS 0" "red" >}} | {{< bg "MISS" "postgis_tiger_geocoder_16 : MISS 0" "red" >}} | {{< bg "N/A" "postgis_tiger_geocoder_15 : N/A 0" "gray" >}} | {{< bg "N/A" "postgis_tiger_geocoder_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "MISS" "postgis_tiger_geocoder_18 : MISS 0" "red" >}} | {{< bg "MISS" "postgis_tiger_geocoder_17 : MISS 0" "red" >}} | {{< bg "MISS" "postgis_tiger_geocoder_16 : MISS 0" "red" >}} | {{< bg "N/A" "postgis_tiger_geocoder_15 : N/A 0" "gray" >}} | {{< bg "N/A" "postgis_tiger_geocoder_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_16 : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgis_tiger_geocoder_15 : N/A 0" "gray" >}} | {{< bg "N/A" "postgis_tiger_geocoder_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_16 : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgis_tiger_geocoder_15 : N/A 0" "gray" >}} | {{< bg "N/A" "postgis_tiger_geocoder_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_16 : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgis_tiger_geocoder_15 : N/A 0" "gray" >}} | {{< bg "N/A" "postgis_tiger_geocoder_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgis_tiger_geocoder_16 : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgis_tiger_geocoder_15 : N/A 0" "gray" >}} | {{< bg "N/A" "postgis_tiger_geocoder_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 2025.2" "postgresql-18-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-17-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "PGDG 2025.2" "postgresql-16-postgis-tiger-geocoder : AVAIL 1" "blue" >}} | {{< bg "N/A" "postgresql-15-postgis-tiger-geocoder : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-postgis-tiger-geocoder : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis_tiger_geocoder_18` | `2025.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.4 KiB | [postgis_tiger_geocoder_18-2025.2-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/postgis_tiger_geocoder_18-2025.2-1PGDG.rhel9.8.noarch.rpm) |
| `postgis_tiger_geocoder_18` | `2025.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 104.3 KiB | [postgis_tiger_geocoder_18-2025.2-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/postgis_tiger_geocoder_18-2025.2-1PGDG.rhel9.8.noarch.rpm) |
| `postgis_tiger_geocoder_18` | `2025.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.5 KiB | [postgis_tiger_geocoder_18-2025.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/postgis_tiger_geocoder_18-2025.2-1PGDG.rhel10.2.noarch.rpm) |
| `postgis_tiger_geocoder_18` | `2025.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.5 KiB | [postgis_tiger_geocoder_18-2025.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/postgis_tiger_geocoder_18-2025.2-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.9 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.9 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.1 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.1 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 95.0 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 95.0 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 95.1 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 95.1 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 95.0 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb) |
| `postgresql-18-postgis-tiger-geocoder` | `2025.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 95.0 KiB | [postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-18-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis_tiger_geocoder_17` | `2025.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.4 KiB | [postgis_tiger_geocoder_17-2025.2-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/postgis_tiger_geocoder_17-2025.2-1PGDG.rhel9.8.noarch.rpm) |
| `postgis_tiger_geocoder_17` | `2025.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 104.3 KiB | [postgis_tiger_geocoder_17-2025.2-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/postgis_tiger_geocoder_17-2025.2-1PGDG.rhel9.8.noarch.rpm) |
| `postgis_tiger_geocoder_17` | `2025.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.6 KiB | [postgis_tiger_geocoder_17-2025.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/postgis_tiger_geocoder_17-2025.2-1PGDG.rhel10.2.noarch.rpm) |
| `postgis_tiger_geocoder_17` | `2025.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.5 KiB | [postgis_tiger_geocoder_17-2025.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/postgis_tiger_geocoder_17-2025.2-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 95.1 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 95.1 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb) |
| `postgresql-17-postgis-tiger-geocoder` | `2025.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.9 KiB | [postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-17-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `postgis_tiger_geocoder_16` | `2025.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 104.4 KiB | [postgis_tiger_geocoder_16-2025.2-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/postgis_tiger_geocoder_16-2025.2-1PGDG.rhel9.8.noarch.rpm) |
| `postgis_tiger_geocoder_16` | `2025.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 104.3 KiB | [postgis_tiger_geocoder_16-2025.2-1PGDG.rhel9.8.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/postgis_tiger_geocoder_16-2025.2-1PGDG.rhel9.8.noarch.rpm) |
| `postgis_tiger_geocoder_16` | `2025.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 104.6 KiB | [postgis_tiger_geocoder_16-2025.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/postgis_tiger_geocoder_16-2025.2-1PGDG.rhel10.2.noarch.rpm) |
| `postgis_tiger_geocoder_16` | `2025.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 104.5 KiB | [postgis_tiger_geocoder_16-2025.2-1PGDG.rhel10.2.noarch.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/postgis_tiger_geocoder_16-2025.2-1PGDG.rhel10.2.noarch.rpm) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [d12.x86_64](/os/d12.x86_64) | pgdg | 95.0 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [d12.aarch64](/os/d12.aarch64) | pgdg | 95.0 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg12+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [d13.x86_64](/os/d13.x86_64) | pgdg | 95.0 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [d13.aarch64](/os/d13.aarch64) | pgdg | 95.0 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg13+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [u22.x86_64](/os/u22.x86_64) | pgdg | 94.9 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [u22.aarch64](/os/u22.aarch64) | pgdg | 94.9 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg22.04+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [u24.x86_64](/os/u24.x86_64) | pgdg | 95.2 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [u24.aarch64](/os/u24.aarch64) | pgdg | 95.2 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg24.04+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [u26.x86_64](/os/u26.x86_64) | pgdg | 94.9 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb) |
| `postgresql-16-postgis-tiger-geocoder` | `2025.2` | [u26.aarch64](/os/u26.aarch64) | pgdg | 94.9 KiB | [postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/postgis-tiger-geocoder/postgresql-16-postgis-tiger-geocoder_2025.2-1.pgdg26.04+1_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://gitea.osgeo.org/postgis/postgis_tiger_geocoder" title="Repository" icon="link" subtitle="gitea.osgeo.org/postgis/postgis_tiger_geocoder" />}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install postgis_tiger_geocoder;		# install via package name, for the active PG version

pig install postgis_tiger_geocoder -v 18;   # install for PG 18
pig install postgis_tiger_geocoder -v 17;   # install for PG 17
pig install postgis_tiger_geocoder -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

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
