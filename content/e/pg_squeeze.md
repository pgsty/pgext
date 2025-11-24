---
title: "pg_squeeze"
linkTitle: "pg_squeeze"
description: "A tool to remove unused space from a relation."
weight: 5040
categories: ["ADMIN"]
width: full
---

[**pg_squeeze**](https://github.com/cybertec-postgresql/pg_squeeze) : A tool to remove unused space from a relation.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5040** | {{< badge content="pg_squeeze" link="https://github.com/cybertec-postgresql/pg_squeeze" >}} | {{< ext "pg_squeeze" >}} | `1.9.1` | {{< category "ADMIN" >}} | {{< license "BSD 2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_repack" >}} {{< ext "pgfincore" >}} {{< ext "pg_prewarm" >}} {{< ext "pgstattuple" >}} {{< ext "pg_cooldown" >}} {{< ext "pgcozy" >}} {{< ext "amcheck" >}} {{< ext "pageinspect" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_squeeze` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.1` | {{< bg "18" "pg_squeeze_18*" "green" >}} {{< bg "17" "pg_squeeze_17*" "green" >}} {{< bg "16" "pg_squeeze_16*" "green" >}} {{< bg "15" "pg_squeeze_15*" "green" >}} {{< bg "14" "pg_squeeze_14*" "green" >}} {{< bg "13" "pg_squeeze_13*" "green" >}} | `pg_squeeze_$v*` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.9.1` | {{< bg "18" "postgresql-18-squeeze" "green" >}} {{< bg "17" "postgresql-17-squeeze" "green" >}} {{< bg "16" "postgresql-16-squeeze" "green" >}} {{< bg "15" "postgresql-15-squeeze" "green" >}} {{< bg "14" "postgresql-14-squeeze" "green" >}} {{< bg "13" "postgresql-13-squeeze" "green" >}} | `postgresql-$v-squeeze` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_13 : AVAIL 8" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_13 : AVAIL 6" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_14 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_13 : AVAIL 7" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_17 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_16 : AVAIL 5" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_15 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_14 : AVAIL 6" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_13 : AVAIL 6" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.9.1" "pg_squeeze_13 : AVAIL 2" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.9.1" "postgresql-18-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-17-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-16-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-15-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-14-squeeze : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.9.1" "postgresql-13-squeeze : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_18` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_squeeze_18-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_squeeze_18-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_squeeze_18-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_squeeze_18-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_squeeze_18-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_18` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.6 KiB | [pg_squeeze_18-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/pg_squeeze_18-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-squeeze` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 115.7 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg12+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.0 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg12+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 115.6 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg13+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.1 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg13+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 118.2 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 113.3 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 115.4 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-squeeze` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 110.8 KiB | [postgresql-18-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-18-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_17` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.9 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.7.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_squeeze_17-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.2 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.7.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_squeeze_17-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.2 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.0 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.9 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_squeeze_17-1.7.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.4 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_17-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.2 KiB | [pg_squeeze_17-1.7.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_squeeze_17-1.7.0-2PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_squeeze_17-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_17` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_17-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pg_squeeze_17-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_17` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.1 KiB | [pg_squeeze_17-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pg_squeeze_17-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-squeeze` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 115.7 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg12+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 110.9 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg12+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 115.6 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg13+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.2 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg13+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 138.8 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 133.7 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 115.3 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-squeeze` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 110.7 KiB | [postgresql-17-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-17-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_16` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.5 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_squeeze_16-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.1 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.8 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_squeeze_16-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.5 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.1 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 55.9 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.1 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.0 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_squeeze_16-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.4 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.2 KiB | [pg_squeeze_16-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.4 KiB | [pg_squeeze_16-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.4 KiB | [pg_squeeze_16-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_squeeze_16-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_squeeze_16-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_16` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_16-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pg_squeeze_16-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_16` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.1 KiB | [pg_squeeze_16-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pg_squeeze_16-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-squeeze` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 115.3 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg12+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 110.6 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg12+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 115.4 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg13+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.1 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg13+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 136.9 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 131.7 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 115.0 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-squeeze` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 110.4 KiB | [postgresql-16-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-16-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_15` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.9 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.8 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.1 KiB | [pg_squeeze_15-1.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_squeeze_15-1.5.0-1.rhel8.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.5 KiB | [pg_squeeze_15-1.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_squeeze_15-1.5.0-1.rhel8.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.9 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.4 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.3 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.1 KiB | [pg_squeeze_15-1.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_squeeze_15-1.5.0-1.rhel9.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.9 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_15-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [pg_squeeze_15-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [pg_squeeze_15-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.5 KiB | [pg_squeeze_15-1.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_squeeze_15-1.5.0-1.rhel9.aarch64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.7 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.2 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_squeeze_15-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_15` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.9 KiB | [pg_squeeze_15-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pg_squeeze_15-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_15` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.5 KiB | [pg_squeeze_15-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pg_squeeze_15-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-squeeze` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 115.5 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg12+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 110.6 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg12+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 115.5 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg13+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.0 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg13+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 137.2 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 132.7 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 115.2 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-squeeze` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 110.6 KiB | [postgresql-15-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-15-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_14` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.9 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.8 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.1 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 53.0 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 46.0 KiB | [pg_squeeze_14-1.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.5.0-1.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 112.2 KiB | [pg_squeeze_14-1.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_squeeze_14-1.4.1-2.rhel8.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 55.0 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.0 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.3 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.2 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.5 KiB | [pg_squeeze_14-1.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_squeeze_14-1.5.0-1.rhel8.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.7 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.6 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.7 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.8 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 46.2 KiB | [pg_squeeze_14-1.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.5.0-1.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 112.8 KiB | [pg_squeeze_14-1.4.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_squeeze_14-1.4.1-2.rhel9.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.9 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.7 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_14-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.8 KiB | [pg_squeeze_14-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.9 KiB | [pg_squeeze_14-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.4 KiB | [pg_squeeze_14-1.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_squeeze_14-1.5.0-1.rhel9.aarch64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.8 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_squeeze_14-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_14` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 56.0 KiB | [pg_squeeze_14-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pg_squeeze_14-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_14` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_14-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pg_squeeze_14-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-squeeze` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 115.7 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg12+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 111.1 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg12+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 115.8 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg13+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 111.3 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg13+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 137.6 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 132.5 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 115.6 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-squeeze` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 110.8 KiB | [postgresql-14-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-14-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_squeeze_13` | `1.9.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 57.3 KiB | [pg_squeeze_13-1.9.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.9.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.8.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.7 KiB | [pg_squeeze_13-1.8.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.8.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 56.2 KiB | [pg_squeeze_13-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.6.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.3 KiB | [pg_squeeze_13-1.6.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.6.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.6.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 52.3 KiB | [pg_squeeze_13-1.6.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.6.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 45.5 KiB | [pg_squeeze_13-1.5.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.5.0-1.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 110.3 KiB | [pg_squeeze_13-1.4.1-2.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.4.1-2.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.3.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 107.8 KiB | [pg_squeeze_13-1.3.1-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_squeeze_13-1.3.1-1.rhel8.x86_64.rpm) |
| `pg_squeeze_13` | `1.9.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.8 KiB | [pg_squeeze_13-1.9.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_squeeze_13-1.9.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_13` | `1.8.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 54.3 KiB | [pg_squeeze_13-1.8.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_squeeze_13-1.8.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_13` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 53.9 KiB | [pg_squeeze_13-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_squeeze_13-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_13` | `1.6.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.1 KiB | [pg_squeeze_13-1.6.2-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_squeeze_13-1.6.2-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_13` | `1.6.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 50.0 KiB | [pg_squeeze_13-1.6.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_squeeze_13-1.6.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_squeeze_13` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 43.3 KiB | [pg_squeeze_13-1.5.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_squeeze_13-1.5.0-1.rhel8.aarch64.rpm) |
| `pg_squeeze_13` | `1.9.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.9 KiB | [pg_squeeze_13-1.9.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.9.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.8.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.5 KiB | [pg_squeeze_13-1.8.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.8.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 56.4 KiB | [pg_squeeze_13-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.6.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.3 KiB | [pg_squeeze_13-1.6.2-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.6.2-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.6.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 52.5 KiB | [pg_squeeze_13-1.6.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.6.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 45.9 KiB | [pg_squeeze_13-1.5.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.5.0-1.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 111.1 KiB | [pg_squeeze_13-1.4.1-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_squeeze_13-1.4.1-2.rhel9.x86_64.rpm) |
| `pg_squeeze_13` | `1.9.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.9 KiB | [pg_squeeze_13-1.9.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_squeeze_13-1.9.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_13` | `1.8.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.6 KiB | [pg_squeeze_13-1.8.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_squeeze_13-1.8.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_13` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 54.5 KiB | [pg_squeeze_13-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_squeeze_13-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_13` | `1.6.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [pg_squeeze_13-1.6.2-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_squeeze_13-1.6.2-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_13` | `1.6.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 50.7 KiB | [pg_squeeze_13-1.6.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_squeeze_13-1.6.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_squeeze_13` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 44.3 KiB | [pg_squeeze_13-1.5.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_squeeze_13-1.5.0-1.rhel9.aarch64.rpm) |
| `pg_squeeze_13` | `1.9.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.5 KiB | [pg_squeeze_13-1.9.1-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_squeeze_13-1.9.1-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_13` | `1.8.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 57.1 KiB | [pg_squeeze_13-1.8.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_squeeze_13-1.8.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_squeeze_13` | `1.9.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.7 KiB | [pg_squeeze_13-1.9.1-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pg_squeeze_13-1.9.1-1PGDG.rhel10.aarch64.rpm) |
| `pg_squeeze_13` | `1.8.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 55.4 KiB | [pg_squeeze_13-1.8.0-1PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pg_squeeze_13-1.8.0-1PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-squeeze` | `1.9.1` | [d12.x86_64](/os/d12.x86_64) | pgdg | 114.9 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg12+1_amd64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [d12.aarch64](/os/d12.aarch64) | pgdg | 110.0 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg12+1_arm64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [d13.x86_64](/os/d13.x86_64) | pgdg | 114.9 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg13+1_amd64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [d13.aarch64](/os/d13.aarch64) | pgdg | 110.8 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg13+1_arm64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [u22.x86_64](/os/u22.x86_64) | pgdg | 136.0 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [u22.aarch64](/os/u22.aarch64) | pgdg | 131.0 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [u24.x86_64](/os/u24.x86_64) | pgdg | 114.8 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-squeeze` | `1.9.1` | [u24.aarch64](/os/u24.aarch64) | pgdg | 109.9 KiB | [postgresql-13-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pg-squeeze/postgresql-13-squeeze_1.9.1-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/cybertec-postgresql/pg_squeeze" title="Repository" icon="github" subtitle="github.com/cybertec-postgresql/pg_squeeze" >}}
{{< /cards >}}


## Install

Make sure [**PGDG**](/repo/pgdg) repo available:

```bash
pig repo add pgdg -u    # add pgdg repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_squeeze;		# install via package name, for the active PG version

pig install pg_squeeze -v 18;   # install for PG 18
pig install pg_squeeze -v 17;   # install for PG 17
pig install pg_squeeze -v 16;   # install for PG 16
pig install pg_squeeze -v 15;   # install for PG 15
pig install pg_squeeze -v 14;   # install for PG 14
pig install pg_squeeze -v 13;   # install for PG 13

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```sql
shared_preload_libraries = 'pg_squeeze';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_squeeze;
```
