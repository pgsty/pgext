---
title: "pg_checksums"
linkTitle: "pg_checksums"
description: "Enable/disable checksums offline and verify them online or offline"
weight: 5110
categories: ["ADMIN"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PGDG"]
page_width: full
---

[**pg_checksums**](https://github.com/credativ/pg_checksums) : Enable/disable checksums offline and verify them online or offline


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5110** | {{< badge content="pg_checksums" link="https://github.com/credativ/pg_checksums" >}} | {{< ext "pg_checksums" >}} | `1.4` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-------" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "amcheck" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_orphaned" >}} {{< ext "pg_pathcheck" >}} {{< ext "pg_isok" >}} {{< ext "pg_surgery" >}} {{< ext "pageinspect" >}} {{< ext "pg_repack" >}} {{< ext "adminpack" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `pg_checksums` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4` | {{< bg "19" "pg_checksums_19" "green" >}} {{< bg "18" "pg_checksums_18" "green" >}} {{< bg "17" "pg_checksums_17" "green" >}} {{< bg "16" "pg_checksums_16" "green" >}} {{< bg "15" "pg_checksums_15" "green" >}} | `pg_checksums_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.4` | {{< bg "19" "postgresql-19-pg-checksums" "green" >}} {{< bg "18" "postgresql-18-pg-checksums" "green" >}} {{< bg "17" "postgresql-17-pg-checksums" "green" >}} {{< bg "16" "postgresql-16-pg-checksums" "green" >}} {{< bg "15" "postgresql-15-pg-checksums" "green" >}} | `postgresql-$v-pg-checksums` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "pg_checksums_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.4" "pg_checksums_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_15 : AVAIL 4" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "pg_checksums_19 : N/A 0" "gray" >}} | {{< bg "PGDG 1.4" "pg_checksums_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_16 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_15 : AVAIL 4" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.4" "pg_checksums_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_15 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.4" "pg_checksums_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.4" "pg_checksums_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_15 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.4" "pg_checksums_19 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_18 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.4" "pg_checksums_15 : AVAIL 5" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PGDG 1.4" "postgresql-19-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-18-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-17-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-16-pg-checksums : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.4" "postgresql-15-pg-checksums : AVAIL 3" "blue" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_19` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.7 KiB | [pg_checksums_19-1.4-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/pg_checksums_19-1.4-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_checksums_19` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.8 KiB | [pg_checksums_19-1.4-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/pg_checksums_19-1.4-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_checksums_19` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 34.0 KiB | [pg_checksums_19-1.4-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/pg_checksums_19-1.4-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_checksums_19` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.3 KiB | [pg_checksums_19-1.4-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/pg_checksums_19-1.4-4PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.5 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg12+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg12+4_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.5 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg12+3_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 39.5 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg12+2_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg12+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg12+4_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg12+3_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.6 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg12+2_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 39.7 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg13+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg13+4_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 39.7 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg13+3_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 39.7 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg13+2_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.1 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg13+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg13+4_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.0 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg13+3_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 37.0 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg13+2_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.7 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.7 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 40.7 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.4 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.3 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.3 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 39.5 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 39.5 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 39.5 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.9 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.9 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.9 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 39.2 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 39.1 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 39.1 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 36.0 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.9 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb) |
| `postgresql-19-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.9 KiB | [postgresql-19-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-19-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_18` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.7 KiB | [pg_checksums_18-1.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_checksums_18-1.4-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_checksums_18` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.3 KiB | [pg_checksums_18-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_checksums_18-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_18` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.9 KiB | [pg_checksums_18-1.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_checksums_18-1.4-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_checksums_18` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 46.8 KiB | [pg_checksums_18-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_checksums_18-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_18` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 33.0 KiB | [pg_checksums_18-1.4-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_checksums_18-1.4-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_checksums_18` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_checksums_18-1.4-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_checksums_18-1.4-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_checksums_18` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 33.3 KiB | [pg_checksums_18-1.4-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_checksums_18-1.4-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_checksums_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.8 KiB | [pg_checksums_18-1.4-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_checksums_18-1.4-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_checksums_18` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.8 KiB | [pg_checksums_18-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_checksums_18-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [pg_checksums_18-1.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_checksums_18-1.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_18` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.6 KiB | [pg_checksums_18-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/pg_checksums_18-1.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg12+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg12+4_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.3 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg12+3_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 37.3 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg12+2_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.1 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg12+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg12+4_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.0 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg12+3_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.0 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg12+2_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.6 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg13+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg13+4_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.6 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg13+3_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.6 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg13+2_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.5 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg13+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg13+4_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.5 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg13+3_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.5 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg13+2_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 36.9 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 37.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 37.3 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 37.3 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.4 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 36.8 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.6 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.6 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb) |
| `postgresql-18-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.6 KiB | [postgresql-18-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-18-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_17` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 48.8 KiB | [pg_checksums_17-1.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_checksums_17-1.4-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_checksums_17` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 47.4 KiB | [pg_checksums_17-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_checksums_17-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_17` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 48.1 KiB | [pg_checksums_17-1.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_checksums_17-1.4-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_checksums_17` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 47.0 KiB | [pg_checksums_17-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_checksums_17-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_17` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 32.3 KiB | [pg_checksums_17-1.4-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_checksums_17-1.4-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_checksums_17` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.7 KiB | [pg_checksums_17-1.4-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_checksums_17-1.4-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_checksums_17` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.5 KiB | [pg_checksums_17-1.4-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_checksums_17-1.4-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_checksums_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [pg_checksums_17-1.4-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_checksums_17-1.4-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_checksums_17` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 41.1 KiB | [pg_checksums_17-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_checksums_17-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [pg_checksums_17-1.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_checksums_17-1.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_17` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.9 KiB | [pg_checksums_17-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/pg_checksums_17-1.3-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 36.9 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg12+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg12+4_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 36.9 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg12+3_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 36.9 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg12+2_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.2 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg12+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg12+4_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.2 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg12+3_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 36.2 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg12+2_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.2 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg13+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg13+4_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.2 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg13+3_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 37.2 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg13+2_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg13+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg13+4_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg13+3_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg13+2_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.1 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.0 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 38.0 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.1 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.1 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 37.1 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 37.0 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 36.7 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 36.4 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 36.4 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 36.3 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.6 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.6 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb) |
| `postgresql-17-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 35.6 KiB | [postgresql-17-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-17-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_16` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.6 KiB | [pg_checksums_16-1.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_checksums_16-1.4-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_checksums_16` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.2 KiB | [pg_checksums_16-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_checksums_16-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.5 KiB | [pg_checksums_16-1.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_checksums_16-1.1-3PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_16` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.9 KiB | [pg_checksums_16-1.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_checksums_16-1.4-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_checksums_16` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.7 KiB | [pg_checksums_16-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_checksums_16-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.1 KiB | [pg_checksums_16-1.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_checksums_16-1.1-3PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_16` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.7 KiB | [pg_checksums_16-1.4-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_checksums_16-1.4-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_checksums_16` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.0 KiB | [pg_checksums_16-1.4-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_checksums_16-1.4-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_checksums_16` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 31.9 KiB | [pg_checksums_16-1.4-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_checksums_16-1.4-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_checksums_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [pg_checksums_16-1.4-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.4-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_checksums_16` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.5 KiB | [pg_checksums_16-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 39.9 KiB | [pg_checksums_16-1.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_16` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.2 KiB | [pg_checksums_16-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_checksums_16` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [pg_checksums_16-1.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/pg_checksums_16-1.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.9 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg12+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg12+4_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.9 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg12+3_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.9 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg12+2_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.2 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg12+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg12+4_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.2 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg12+3_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 34.2 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg12+2_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 35.3 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg13+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg13+4_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 35.2 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg13+3_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 35.2 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg13+2_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.7 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg13+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg13+4_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.7 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg13+3_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.7 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg13+2_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.2 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.1 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 36.1 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.0 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.0 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 35.0 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.1 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.1 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 35.1 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.7 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.6 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.7 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.5 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.4 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.5 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.8 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.8 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb) |
| `postgresql-16-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.8 KiB | [postgresql-16-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-16-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_checksums_15` | `1.4` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.1 KiB | [pg_checksums_15-1.4-4PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.4-4PGDG.rhel8.10.x86_64.rpm) |
| `pg_checksums_15` | `1.3` | [el8.x86_64](/os/el8.x86_64) | pgdg | 44.8 KiB | [pg_checksums_15-1.3-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.3-1PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.1 KiB | [pg_checksums_15-1.1-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.1-3PGDG.rhel8.x86_64.rpm) |
| `pg_checksums_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.0 KiB | [pg_checksums_15-1.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_checksums_15-1.1-1.rhel8.x86_64.rpm) |
| `pg_checksums_15` | `1.4` | [el8.aarch64](/os/el8.aarch64) | pgdg | 45.5 KiB | [pg_checksums_15-1.4-4PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.4-4PGDG.rhel8.10.aarch64.rpm) |
| `pg_checksums_15` | `1.3` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.3 KiB | [pg_checksums_15-1.3-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.3-1PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.7 KiB | [pg_checksums_15-1.1-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.1-3PGDG.rhel8.aarch64.rpm) |
| `pg_checksums_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 44.5 KiB | [pg_checksums_15-1.1-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_checksums_15-1.1-1.rhel8.aarch64.rpm) |
| `pg_checksums_15` | `1.4` | [el9.x86_64](/os/el9.x86_64) | pgdg | 31.8 KiB | [pg_checksums_15-1.4-4PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_checksums_15-1.4-4PGDG.rhel9.9.x86_64.rpm) |
| `pg_checksums_15` | `1.4` | [el9.aarch64](/os/el9.aarch64) | pgdg | 40.3 KiB | [pg_checksums_15-1.4-4PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_checksums_15-1.4-4PGDG.rhel9.9.aarch64.rpm) |
| `pg_checksums_15` | `1.4` | [el10.x86_64](/os/el10.x86_64) | pgdg | 32.0 KiB | [pg_checksums_15-1.4-4PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_checksums_15-1.4-4PGDG.rhel10.3.x86_64.rpm) |
| `pg_checksums_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.9 KiB | [pg_checksums_15-1.4-4PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.4-4PGDG.rhel10.3.aarch64.rpm) |
| `pg_checksums_15` | `1.4` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.9 KiB | [pg_checksums_15-1.4-4PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.4-4PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.3 KiB | [pg_checksums_15-1.3-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.3-1PGDG.rhel10.2.aarch64.rpm) |
| `pg_checksums_15` | `1.3` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_checksums_15-1.3-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.3-1PGDG.rhel10.aarch64.rpm) |
| `pg_checksums_15` | `1.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 40.7 KiB | [pg_checksums_15-1.2-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/pg_checksums_15-1.2-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.6 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg12+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg12+4_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.6 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg12+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg12+3_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d12.x86_64](/os/d12.x86_64) | pgdg | 34.6 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg12+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg12+2_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.9 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg12+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg12+4_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.9 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg12+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg12+3_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d12.aarch64](/os/d12.aarch64) | pgdg | 33.9 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg12+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg12+2_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.9 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg13+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg13+4_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.9 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg13+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg13+3_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d13.x86_64](/os/d13.x86_64) | pgdg | 34.9 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg13+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg13+2_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.5 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg13+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg13+4_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.5 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg13+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg13+3_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [d13.aarch64](/os/d13.aarch64) | pgdg | 34.5 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg13+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg13+2_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.7 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg22.04+4_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.8 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg22.04+3_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u22.x86_64](/os/u22.x86_64) | pgdg | 35.7 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg22.04+2_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 34.7 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg22.04+4_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 34.7 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg22.04+3_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u22.aarch64](/os/u22.aarch64) | pgdg | 34.7 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg22.04+2_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.8 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg24.04+4_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.8 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg24.04+3_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u24.x86_64](/os/u24.x86_64) | pgdg | 34.8 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg24.04+2_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.5 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg24.04+4_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.4 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg24.04+3_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u24.aarch64](/os/u24.aarch64) | pgdg | 34.4 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg24.04+2_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.3 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg26.04+4_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.2 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg26.04+3_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u26.x86_64](/os/u26.x86_64) | pgdg | 34.2 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg26.04+2_amd64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.6 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg26.04+4_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.6 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg26.04+3_arm64.deb) |
| `postgresql-15-pg-checksums` | `1.4` | [u26.aarch64](/os/u26.aarch64) | pgdg | 33.6 KiB | [postgresql-15-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-checksums/postgresql-15-pg-checksums_1.4-1.pgdg26.04+2_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/credativ/pg_checksums" title="Repository" icon="github" subtitle="github.com/credativ/pg_checksums" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_checksums-1.4.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_checksums;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_checksums;		# install via package name, for the active PG version

pig install pg_checksums -v 19;   # install for PG 19
pig install pg_checksums -v 18;   # install for PG 18
pig install pg_checksums -v 17;   # install for PG 17
pig install pg_checksums -v 16;   # install for PG 16
pig install pg_checksums -v 15;   # install for PG 15

```


This extension does not need `CREATE EXTENSION` DDL command






## Usage

> [pg_checksums: Activate/deactivate/verify checksums in offline Postgres clusters](https://github.com/credativ/pg_checksums)

`pg_checksums_ext` is a command-line tool (based on PostgreSQL's built-in `pg_checksums`) that can verify, activate, or deactivate page-level checksums for a PostgreSQL cluster. It extends the built-in tool with online verification, `SIGUSR1` progress toggling, fine-grained progress reporting, and I/O rate limiting.

### Verify Checksums (can be done online)

```bash
pg_checksums_ext -D /path/to/data --check
```

### Enable Checksums (requires clean shutdown)

```bash
pg_checksums_ext -D /path/to/data --enable
```

### Disable Checksums (requires clean shutdown)

```bash
pg_checksums_ext -D /path/to/data --disable
```

### Additional Options

- `-D, --pgdata` -- data directory path
- `--check` / `--enable` / `--disable` -- operation mode
- `--progress` -- show progress reporting
- `--filenode` -- check only a specific filenode
- `--no-sync` -- skip fsync
- `--verbose` -- verbose output
- `--debug` -- debug output
- Send `SIGUSR1` to toggle progress reporting during operation
