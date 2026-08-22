---
title: "qdgc"
linkTitle: "qdgc"
description: "Encode, decode, navigate, and fill Extended Quarter Degree Grid Cell codes in pure SQL."
weight: 1700
categories: ["GIS"]
languages: ["SQL"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**qdgc**](https://pgxn.org/dist/qdgc/0.1.0/) : Encode, decode, navigate, and fill Extended Quarter Degree Grid Cell codes in pure SQL.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1700** | {{< badge content="qdgc" link="https://pgxn.org/dist/qdgc/0.1.0/" >}} | {{< ext "qdgc" >}} | `0.1.0` | {{< category "GIS" >}} | {{< license "Apache-2.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "qdgc_postgis" >}} |
|   **See Also**    | {{< ext "h3" >}} {{< ext "pgrouting" >}} {{< ext "pg_geohash" >}} {{< ext "q3c" >}} {{< ext "postgis_topology" >}} {{< ext "pg_polyline" >}} {{< ext "pg_eviltransform" >}} {{< ext "mobilitydb" >}} {{< ext "earthdistance" >}} {{< ext "pointcloud" >}} |
|    **Siblings**   | {{< ext "qdgc_postgis" >}} |

> [!Note] PGXN distribution qdgc also ships qdgc_postgis; the GitHub v0.1.0 tag belongs to qdgc-py and is not this PGXN release.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `qdgc` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "qdgc_18" "green" >}} {{< bg "17" "qdgc_17" "green" >}} {{< bg "16" "qdgc_16" "green" >}} {{< bg "15" "qdgc_15" "green" >}} {{< bg "14" "qdgc_14" "green" >}} | `qdgc_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-qdgc" "green" >}} {{< bg "17" "postgresql-17-qdgc" "green" >}} {{< bg "16" "postgresql-16-qdgc" "green" >}} {{< bg "15" "postgresql-15-qdgc" "green" >}} {{< bg "14" "postgresql-14-qdgc" "green" >}} | `postgresql-$v-qdgc` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "qdgc_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-qdgc : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-qdgc : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `qdgc_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.0 KiB | [qdgc_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/qdgc_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.9 KiB | [qdgc_18-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/qdgc_18-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.5 KiB | [qdgc_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/qdgc_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [qdgc_18-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/qdgc_18-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.7 KiB | [qdgc_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/qdgc_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `qdgc_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.6 KiB | [qdgc_18-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/qdgc_18-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-qdgc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.1 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.1 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.5 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.5 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-18-qdgc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 16.5 KiB | [postgresql-18-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-18-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `qdgc_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.0 KiB | [qdgc_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/qdgc_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.9 KiB | [qdgc_17-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/qdgc_17-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.5 KiB | [qdgc_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/qdgc_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [qdgc_17-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/qdgc_17-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.7 KiB | [qdgc_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/qdgc_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `qdgc_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.6 KiB | [qdgc_17-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/qdgc_17-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-qdgc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.1 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.1 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.5 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.5 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-17-qdgc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 16.5 KiB | [postgresql-17-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-17-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `qdgc_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.0 KiB | [qdgc_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/qdgc_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.9 KiB | [qdgc_16-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/qdgc_16-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.5 KiB | [qdgc_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/qdgc_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [qdgc_16-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/qdgc_16-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.7 KiB | [qdgc_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/qdgc_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `qdgc_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.6 KiB | [qdgc_16-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/qdgc_16-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-qdgc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.1 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.1 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.5 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.5 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-16-qdgc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 16.5 KiB | [postgresql-16-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-16-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `qdgc_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.0 KiB | [qdgc_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/qdgc_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.9 KiB | [qdgc_15-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/qdgc_15-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.5 KiB | [qdgc_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/qdgc_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [qdgc_15-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/qdgc_15-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.7 KiB | [qdgc_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/qdgc_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `qdgc_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.6 KiB | [qdgc_15-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/qdgc_15-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-qdgc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.1 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.1 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.5 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.5 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-15-qdgc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 16.5 KiB | [postgresql-15-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-15-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `qdgc_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 25.0 KiB | [qdgc_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/qdgc_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 24.9 KiB | [qdgc_14-0.1.0-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/qdgc_14-0.1.0-1PIGSTY.el8.noarch.rpm) |
| `qdgc_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 24.5 KiB | [qdgc_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/qdgc_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 24.5 KiB | [qdgc_14-0.1.0-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/qdgc_14-0.1.0-1PIGSTY.el9.noarch.rpm) |
| `qdgc_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 24.7 KiB | [qdgc_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/qdgc_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `qdgc_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 24.6 KiB | [qdgc_14-0.1.0-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/qdgc_14-0.1.0-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-qdgc` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 16.1 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 16.1 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 16.1 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 16.1 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~trixie_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 16.5 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 16.5 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~jammy_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 16.5 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 16.5 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~noble_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 16.5 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
| `postgresql-14-qdgc` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 16.5 KiB | [postgresql-14-qdgc_0.1.0-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/q/qdgc/postgresql-14-qdgc_0.1.0-1PIGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://pgxn.org/dist/qdgc/0.1.0/" title="Repository" icon="link" subtitle="pgxn.org/dist/qdgc/0.1.0/" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="qdgc-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg qdgc;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install qdgc;		# install via package name, for the active PG version

pig install qdgc -v 18;   # install for PG 18
pig install qdgc -v 17;   # install for PG 17
pig install qdgc -v 16;   # install for PG 16
pig install qdgc -v 15;   # install for PG 15
pig install qdgc -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION qdgc;
```

## Usage

Sources:

- [PGXN qdgc 0.1.0 release](https://pgxn.org/dist/qdgc/0.1.0/)
- [Official 0.1.0 README](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/README.md)
- [Official qdgc control file](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc.control)
- [Official qdgc 0.1.0 extension SQL](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc--0.1.0.sql)

`qdgc` 0.1.0 is the trusted, relocatable, pure-SQL core of the QDGC extension family. It encodes longitude and latitude as Extended Quarter Degree Grid Cell codes, decodes their bounds, navigates the prefix hierarchy, reports level metrics, and fills longitude/latitude bounding boxes. It has no PostGIS or compiled-library dependency; geometry, geography, and polygon-fill operations belong to the companion `qdgc_postgis` extension.

### Core Workflow

```sql
CREATE EXTENSION qdgc;

-- qdgc_encode uses (longitude, latitude, level).
SELECT qdgc_encode(31.4, 2.7, 5);
-- E031N02ADBAC

-- The h3-style alias reverses the coordinate arguments.
SELECT qdgc_latlng_to_cell(2.7, 31.4, 5);

SELECT *
FROM qdgc_cell_to_bounds('E031N02ADBAC');

SELECT qdgc_cell_to_parent('E031N02ADBAC', 3);
SELECT * FROM qdgc_cell_to_children('E031N02AD', 5);
```

QDGC hierarchy is encoded directly in the text: a child code begins with its parent code. That makes prefix filtering useful for rollups and descendant lookups:

```sql
CREATE TABLE observations (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    qdgc_code text NOT NULL
);

CREATE INDEX observations_qdgc_idx ON observations (qdgc_code);

SELECT qdgc_cell_to_parent(qdgc_code, 3) AS level_3_cell,
       count(*)
FROM observations
GROUP BY 1;

SELECT *
FROM observations
WHERE qdgc_code LIKE 'E031N02AB%';
```

### Bounding Boxes and Level Metrics

The core extension can enumerate rectangular coverage without PostGIS. Pass `min_lon > max_lon` for a box that crosses the antimeridian.

```sql
SELECT qdgc_bbox_cell_count(30.0, 1.0, 32.0, 3.0, 7);

SELECT *
FROM qdgc_bbox_to_cells(30.0, 1.0, 32.0, 3.0, 7);

SELECT qdgc_level_degrees(7);
SELECT qdgc_get_num_cells(7);
SELECT qdgc_average_cell_area(7, 2.0, 'km^2');
SELECT qdgc_version();
```

`qdgc_average_cell_area` is a spherical estimate. Use `qdgc_cell_area_km2` from `qdgc_postgis` when a cell-specific WGS84 spheroid measurement is required.

### Important Objects

- `qdgc_encode(lon, lat, level)` and `qdgc_latlng_to_cell(lat, lng, level)` create codes; the argument order is intentionally different.
- `qdgc_is_valid_cell`, `qdgc_get_level`, `qdgc_cell_to_bounds`, `qdgc_cell_to_lonlat`, and `qdgc_cell_to_latlng` inspect or decode a code.
- `qdgc_cell_to_parent` and `qdgc_cell_to_children` navigate the four-way prefix hierarchy.
- `qdgc_bbox_to_cells` enumerates cells meeting a bounding box, while `qdgc_bbox_cell_count` calculates the count without materializing the set.
- `qdgc_level_degrees`, `qdgc_get_num_cells`, and `qdgc_average_cell_area` report grid-level metrics.

### Operational Notes

- Upstream requires PostgreSQL 13 or newer and tests PostgreSQL 13 through 17. PostgreSQL 18 is not part of the published 0.1.0 test matrix.
- The control file sets `trusted = true` and `relocatable = true`. No `shared_preload_libraries`, `LOAD`, server restart, or native library is required.
- Relocatable functions call one another by unqualified name. Install `qdgc` into a schema on the active `search_path`; the default `public` schema satisfies this boundary.
- Coordinates are longitude/latitude degrees. `qdgc_encode` takes longitude first, while `qdgc_latlng_to_cell` takes latitude first.
- Result cardinality grows by four for every additional child level. Count a bounding-box fill before materializing it, and avoid requesting deep descendants without a deliberate result-size bound.
