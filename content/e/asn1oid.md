---
title: "asn1oid"
linkTitle: "asn1oid"
description: "asn1oid extension"
weight: 3560
categories: ["TYPE"]
width: full
---

[**asn1oid**](https://github.com/df7cb/pgsql-asn1oid) : asn1oid extension


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3560** | {{< badge content="asn1oid" link="https://github.com/df7cb/pgsql-asn1oid" >}} | {{< ext "asn1oid" >}} | `1.6` | {{< category "TYPE" >}} | {{< license "GPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pguecc" >}} {{< ext "pgcrypto" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.6` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `asn1oid` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.6` | {{< bg "18" "asn1oid_18" "green" >}} {{< bg "17" "asn1oid_17" "green" >}} {{< bg "16" "asn1oid_16" "green" >}} {{< bg "15" "asn1oid_15" "green" >}} {{< bg "14" "asn1oid_14" "green" >}} {{< bg "13" "asn1oid_13" "green" >}} | `asn1oid_$v` | - |
| **DEB** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.6` | {{< bg "18" "postgresql-18-asn1oid" "green" >}} {{< bg "17" "postgresql-17-asn1oid" "green" >}} {{< bg "16" "postgresql-16-asn1oid" "green" >}} {{< bg "15" "postgresql-15-asn1oid" "green" >}} {{< bg "14" "postgresql-14-asn1oid" "green" >}} {{< bg "13" "postgresql-13-asn1oid" "green" >}} | `postgresql-$v-asn1oid` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.6" "asn1oid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.6" "asn1oid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.6" "asn1oid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.6" "asn1oid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.6" "asn1oid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.6" "asn1oid_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.6" "asn1oid_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PGDG 1.6" "postgresql-18-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-17-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-16-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-15-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-14-asn1oid : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6" "postgresql-13-asn1oid : AVAIL 1" "blue" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `asn1oid_18` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [asn1oid_18-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_18-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_18` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.7 KiB | [asn1oid_18-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_18-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_18` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [asn1oid_18-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_18-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_18` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [asn1oid_18-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_18-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_18` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [asn1oid_18-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/asn1oid_18-1.6-1PIGSTY.el10.x86_64.rpm) |
| `asn1oid_18` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [asn1oid_18-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/asn1oid_18-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-asn1oid` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-18-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.9 KiB | [postgresql-18-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-18-asn1oid_1.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg13+1_amd64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 13.0 KiB | [postgresql-18-asn1oid_1.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg13+1_arm64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.1 KiB | [postgresql-18-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.0 KiB | [postgresql-18-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.9 KiB | [postgresql-18-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-18-asn1oid` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.1 KiB | [postgresql-18-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-18-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `asn1oid_17` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [asn1oid_17-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_17-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_17` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [asn1oid_17-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_17-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_17` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [asn1oid_17-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_17-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_17` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [asn1oid_17-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_17-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_17` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [asn1oid_17-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/asn1oid_17-1.6-1PIGSTY.el10.x86_64.rpm) |
| `asn1oid_17` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [asn1oid_17-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/asn1oid_17-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-asn1oid` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-17-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-17-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-17-asn1oid_1.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg13+1_amd64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.9 KiB | [postgresql-17-asn1oid_1.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg13+1_arm64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.4 KiB | [postgresql-17-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.3 KiB | [postgresql-17-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.9 KiB | [postgresql-17-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-17-asn1oid` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-17-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-17-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `asn1oid_16` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [asn1oid_16-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_16-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_16` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [asn1oid_16-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_16-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_16` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [asn1oid_16-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_16-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_16` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [asn1oid_16-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_16-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_16` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.3 KiB | [asn1oid_16-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/asn1oid_16-1.6-1PIGSTY.el10.x86_64.rpm) |
| `asn1oid_16` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [asn1oid_16-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/asn1oid_16-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-asn1oid` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-16-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-16-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-16-asn1oid_1.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg13+1_amd64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.9 KiB | [postgresql-16-asn1oid_1.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg13+1_arm64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.4 KiB | [postgresql-16-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.3 KiB | [postgresql-16-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.9 KiB | [postgresql-16-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-16-asn1oid` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.9 KiB | [postgresql-16-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-16-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `asn1oid_15` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [asn1oid_15-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_15-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_15` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [asn1oid_15-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_15-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_15` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [asn1oid_15-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_15-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_15` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.5 KiB | [asn1oid_15-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_15-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_15` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [asn1oid_15-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/asn1oid_15-1.6-1PIGSTY.el10.x86_64.rpm) |
| `asn1oid_15` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [asn1oid_15-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/asn1oid_15-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-asn1oid` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-15-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-15-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-15-asn1oid_1.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg13+1_amd64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.9 KiB | [postgresql-15-asn1oid_1.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg13+1_arm64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.4 KiB | [postgresql-15-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.3 KiB | [postgresql-15-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.9 KiB | [postgresql-15-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-15-asn1oid` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-15-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-15-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `asn1oid_14` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.5 KiB | [asn1oid_14-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_14-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_14` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [asn1oid_14-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_14-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_14` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [asn1oid_14-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_14-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_14` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [asn1oid_14-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_14-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_14` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [asn1oid_14-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/asn1oid_14-1.6-1PIGSTY.el10.x86_64.rpm) |
| `asn1oid_14` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [asn1oid_14-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/asn1oid_14-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-asn1oid` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg13+1_amd64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.9 KiB | [postgresql-14-asn1oid_1.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg13+1_arm64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.4 KiB | [postgresql-14-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.3 KiB | [postgresql-14-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-14-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-14-asn1oid` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 13.0 KiB | [postgresql-14-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-14-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `asn1oid_13` | `1.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.4 KiB | [asn1oid_13-1.6-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/asn1oid_13-1.6-1PIGSTY.el8.x86_64.rpm) |
| `asn1oid_13` | `1.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 13.6 KiB | [asn1oid_13-1.6-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/asn1oid_13-1.6-1PIGSTY.el8.aarch64.rpm) |
| `asn1oid_13` | `1.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 13.4 KiB | [asn1oid_13-1.6-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/asn1oid_13-1.6-1PIGSTY.el9.x86_64.rpm) |
| `asn1oid_13` | `1.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 13.4 KiB | [asn1oid_13-1.6-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/asn1oid_13-1.6-1PIGSTY.el9.aarch64.rpm) |
| `asn1oid_13` | `1.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 13.4 KiB | [asn1oid_13-1.6-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/asn1oid_13-1.6-1PIGSTY.el10.x86_64.rpm) |
| `asn1oid_13` | `1.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 13.6 KiB | [asn1oid_13-1.6-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/asn1oid_13-1.6-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-asn1oid` | `1.6` | [d12.x86_64](/os/d12.x86_64) | pgdg | 12.7 KiB | [postgresql-13-asn1oid_1.6-3.pgdg12+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg12+1_amd64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [d12.aarch64](/os/d12.aarch64) | pgdg | 12.8 KiB | [postgresql-13-asn1oid_1.6-3.pgdg12+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg12+1_arm64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [d13.x86_64](/os/d13.x86_64) | pgdg | 12.7 KiB | [postgresql-13-asn1oid_1.6-3.pgdg13+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg13+1_amd64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [d13.aarch64](/os/d13.aarch64) | pgdg | 12.9 KiB | [postgresql-13-asn1oid_1.6-3.pgdg13+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg13+1_arm64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [u22.x86_64](/os/u22.x86_64) | pgdg | 13.3 KiB | [postgresql-13-asn1oid_1.6-3.pgdg22.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg22.04+1_amd64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [u22.aarch64](/os/u22.aarch64) | pgdg | 13.2 KiB | [postgresql-13-asn1oid_1.6-3.pgdg22.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg22.04+1_arm64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [u24.x86_64](/os/u24.x86_64) | pgdg | 12.8 KiB | [postgresql-13-asn1oid_1.6-3.pgdg24.04+1_amd64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg24.04+1_amd64.deb) |
| `postgresql-13-asn1oid` | `1.6` | [u24.aarch64](/os/u24.aarch64) | pgdg | 12.9 KiB | [postgresql-13-asn1oid_1.6-3.pgdg24.04+1_arm64.deb](https://apt.postgresql.org/pub/repos/apt/pool/main/p/pgsql-asn1oid/postgresql-13-asn1oid_1.6-3.pgdg24.04+1_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/df7cb/pgsql-asn1oid" title="Repository" icon="github" subtitle="github.com/df7cb/pgsql-asn1oid" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql-asn1oid-1.6.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg asn1oid;		# build rpm
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install asn1oid;		# install via package name, for the active PG version

pig install asn1oid -v 18;   # install for PG 18
pig install asn1oid -v 17;   # install for PG 17
pig install asn1oid -v 16;   # install for PG 16
pig install asn1oid -v 15;   # install for PG 15
pig install asn1oid -v 14;   # install for PG 14
pig install asn1oid -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION asn1oid;
```
