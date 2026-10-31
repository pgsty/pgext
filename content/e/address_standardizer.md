---
title: "address_standardizer"
linkTitle: "address_standardizer"
description: "Used to parse an address into constituent elements. Generally used to support geocoding address normalization step."
weight: 1505
categories: ["GIS"]
languages: ["C"]
licenses: ["GPL-2.0"]
repos: ["PGDG"]
page_width: full
---

[**address_standardizer**](https://github.com/postgis/address_standardizer) : Used to parse an address into constituent elements. Generally used to support geocoding address normalization step.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1505** | {{< badge content="address_standardizer" link="https://github.com/postgis/address_standardizer" >}} | {{< ext "address_standardizer" >}} | `3.7` | {{< category "GIS" >}} | {{< license "GPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "nominatim_fdw" >}} {{< ext "geoip" >}} {{< ext "country" >}} {{< ext "tzf" >}} {{< ext "pg_geohash" >}} {{< ext "ogr_fdw" >}} {{< ext "addressing_dictionary" >}} |
|    **Siblings**   | {{< ext "address_standardizer_data_us" >}} |

> [!Note] Standalone since PostGIS 3.7; control 3.7; PGDG package 3.7.0; no local recipe.


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


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `address_standardizer_18` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 216.0 KiB | [address_standardizer_18-3.7.0-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/address_standardizer_18-3.7.0-4PGDG.rhel9.8.x86_64.rpm) |
| `address_standardizer_18` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 214.6 KiB | [address_standardizer_18-3.7.0-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/address_standardizer_18-3.7.0-4PGDG.rhel9.8.aarch64.rpm) |
| `address_standardizer_18` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.9 KiB | [address_standardizer_18-3.7.0-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/address_standardizer_18-3.7.0-4PGDG.rhel10.2.x86_64.rpm) |
| `address_standardizer_18` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.1 KiB | [address_standardizer_18-3.7.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/address_standardizer_18-3.7.0-4PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-18-address-standardizer` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 528.3 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 524.6 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 529.0 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 526.1 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 526.6 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 523.1 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 523.9 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 520.3 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 522.7 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-18-address-standardizer` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 519.4 KiB | [postgresql-18-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-18-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `address_standardizer_17` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 215.9 KiB | [address_standardizer_17-3.7.0-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/address_standardizer_17-3.7.0-4PGDG.rhel9.8.x86_64.rpm) |
| `address_standardizer_17` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 214.5 KiB | [address_standardizer_17-3.7.0-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/address_standardizer_17-3.7.0-4PGDG.rhel9.8.aarch64.rpm) |
| `address_standardizer_17` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.8 KiB | [address_standardizer_17-3.7.0-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/address_standardizer_17-3.7.0-4PGDG.rhel10.2.x86_64.rpm) |
| `address_standardizer_17` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.1 KiB | [address_standardizer_17-3.7.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/address_standardizer_17-3.7.0-4PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-17-address-standardizer` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 528.0 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 524.7 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 529.1 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 526.1 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 546.3 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 542.5 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 523.3 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 520.0 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 522.2 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-17-address-standardizer` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 518.8 KiB | [postgresql-17-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-17-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `address_standardizer_16` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 215.9 KiB | [address_standardizer_16-3.7.0-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/address_standardizer_16-3.7.0-4PGDG.rhel9.8.x86_64.rpm) |
| `address_standardizer_16` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 214.5 KiB | [address_standardizer_16-3.7.0-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/address_standardizer_16-3.7.0-4PGDG.rhel9.8.aarch64.rpm) |
| `address_standardizer_16` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 216.8 KiB | [address_standardizer_16-3.7.0-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/address_standardizer_16-3.7.0-4PGDG.rhel10.2.x86_64.rpm) |
| `address_standardizer_16` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 215.1 KiB | [address_standardizer_16-3.7.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/address_standardizer_16-3.7.0-4PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-16-address-standardizer` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 528.1 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 524.6 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 529.2 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 526.4 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 545.8 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 542.6 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 523.2 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 520.2 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 522.4 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-16-address-standardizer` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 518.9 KiB | [postgresql-16-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-16-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `address_standardizer_15` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 224.5 KiB | [address_standardizer_15-3.7.0-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/address_standardizer_15-3.7.0-4PGDG.rhel9.8.x86_64.rpm) |
| `address_standardizer_15` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 222.1 KiB | [address_standardizer_15-3.7.0-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/address_standardizer_15-3.7.0-4PGDG.rhel9.8.aarch64.rpm) |
| `address_standardizer_15` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 225.0 KiB | [address_standardizer_15-3.7.0-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/address_standardizer_15-3.7.0-4PGDG.rhel10.2.x86_64.rpm) |
| `address_standardizer_15` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 222.6 KiB | [address_standardizer_15-3.7.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/address_standardizer_15-3.7.0-4PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-15-address-standardizer` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 528.8 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 525.8 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 529.5 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 527.1 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 551.9 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 548.3 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 528.9 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 525.9 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 527.1 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-15-address-standardizer` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 524.8 KiB | [postgresql-15-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-15-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `address_standardizer_14` | `3.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 224.4 KiB | [address_standardizer_14-3.7.0-4PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/address_standardizer_14-3.7.0-4PGDG.rhel9.8.x86_64.rpm) |
| `address_standardizer_14` | `3.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 222.1 KiB | [address_standardizer_14-3.7.0-4PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/address_standardizer_14-3.7.0-4PGDG.rhel9.8.aarch64.rpm) |
| `address_standardizer_14` | `3.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 224.9 KiB | [address_standardizer_14-3.7.0-4PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/address_standardizer_14-3.7.0-4PGDG.rhel10.2.x86_64.rpm) |
| `address_standardizer_14` | `3.7.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 222.6 KiB | [address_standardizer_14-3.7.0-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/address_standardizer_14-3.7.0-4PGDG.rhel10.2.aarch64.rpm) |
| `postgresql-14-address-standardizer` | `3.7.0` | [d12.x86_64](/os/d12.x86_64) | pgdg | 528.9 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg12+1_amd64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [d12.aarch64](/os/d12.aarch64) | pgdg | 525.7 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg12+1_arm64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [d13.x86_64](/os/d13.x86_64) | pgdg | 529.4 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg13+1_amd64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [d13.aarch64](/os/d13.aarch64) | pgdg | 527.2 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg13+1_arm64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [u22.x86_64](/os/u22.x86_64) | pgdg | 550.7 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg22.04+1_amd64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [u22.aarch64](/os/u22.aarch64) | pgdg | 547.3 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg22.04+1_arm64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [u24.x86_64](/os/u24.x86_64) | pgdg | 528.8 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg24.04+1_amd64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [u24.aarch64](/os/u24.aarch64) | pgdg | 525.9 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg24.04+1_arm64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [u26.x86_64](/os/u26.x86_64) | pgdg | 527.3 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg26.04+1_amd64.deb) |
| `postgresql-14-address-standardizer` | `3.7.0` | [u26.aarch64](/os/u26.aarch64) | pgdg | 524.8 KiB | [postgresql-14-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/a/address-standardizer/postgresql-14-address-standardizer_3.7.0-1.pgdg26.04+1_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

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

pig install address_standardizer -v 18;   # install for PG 18
pig install address_standardizer -v 17;   # install for PG 17
pig install address_standardizer -v 16;   # install for PG 16
pig install address_standardizer -v 15;   # install for PG 15
pig install address_standardizer -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

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
