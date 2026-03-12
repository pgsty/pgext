---
title: "arraymath"
linkTitle: "arraymath"
description: "Array math and operators that work element by element on the contents of arrays"
weight: 4770
categories: ["FUNC"]
width: full
---

[**pg_arraymath**](https://github.com/pramsey/pgsql-arraymath) : Array math and operators that work element by element on the contents of arrays


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4770** | {{< badge content="arraymath" link="https://github.com/pramsey/pgsql-arraymath" >}} | {{< ext "arraymath" "pg_arraymath" >}} | `1.1` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "aggs_for_arrays" >}} {{< ext "aggs_for_vecs" >}} {{< ext "intarray" >}} {{< ext "first_last_agg" >}} {{< ext "floatvec" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_arraymath` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "pg_arraymath_18" "green" >}} {{< bg "17" "pg_arraymath_17" "green" >}} {{< bg "16" "pg_arraymath_16" "green" >}} {{< bg "15" "pg_arraymath_15" "green" >}} {{< bg "14" "pg_arraymath_14" "green" >}} | `pg_arraymath_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.1` | {{< bg "18" "postgresql-18-pg-arraymath" "green" >}} {{< bg "17" "postgresql-17-pg-arraymath" "green" >}} {{< bg "16" "postgresql-16-pg-arraymath" "green" >}} {{< bg "15" "postgresql-15-pg-arraymath" "green" >}} {{< bg "14" "postgresql-14-pg-arraymath" "green" >}} | `postgresql-$v-pg-arraymath` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "pg_arraymath_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.1" "postgresql-18-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-17-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-16-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-15-pg-arraymath : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.1" "postgresql-14-pg-arraymath : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_arraymath_18` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_arraymath_18-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_18-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_18` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.2 KiB | [pg_arraymath_18-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_18-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_18` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pg_arraymath_18-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_18-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_18` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_arraymath_18-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_18-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_18` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.6 KiB | [pg_arraymath_18-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_arraymath_18-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_arraymath_18` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.4 KiB | [pg_arraymath_18-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_arraymath_18-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-arraymath` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.1 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.8 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.2 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.0 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 26.6 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 26.3 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.3 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-arraymath` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.1 KiB | [postgresql-18-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-18-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_arraymath_17` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_arraymath_17-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_17-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_17` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.3 KiB | [pg_arraymath_17-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_17-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_17` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_arraymath_17-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_17-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_17` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_arraymath_17-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_17-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_17` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [pg_arraymath_17-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_arraymath_17-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_arraymath_17` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pg_arraymath_17-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_arraymath_17-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-arraymath` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.0 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.8 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.1 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.0 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.9 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.6 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.3 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-arraymath` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.1 KiB | [postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-17-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_arraymath_16` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.7 KiB | [pg_arraymath_16-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_16-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_16` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.3 KiB | [pg_arraymath_16-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_16-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_16` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.5 KiB | [pg_arraymath_16-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_16-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_16` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_arraymath_16-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_16-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_16` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.7 KiB | [pg_arraymath_16-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_arraymath_16-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_arraymath_16` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pg_arraymath_16-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_arraymath_16-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-arraymath` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 25.1 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.9 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 25.2 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 25.0 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.9 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.6 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.4 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-arraymath` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 26.2 KiB | [postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-16-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_arraymath_15` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.6 KiB | [pg_arraymath_15-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_15-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_15` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.1 KiB | [pg_arraymath_15-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_15-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_15` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pg_arraymath_15-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_15-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_15` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_arraymath_15-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_15-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_15` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.6 KiB | [pg_arraymath_15-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_arraymath_15-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_arraymath_15` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pg_arraymath_15-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_arraymath_15-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-arraymath` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.7 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.5 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.8 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.8 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.4 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.3 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 26.0 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-arraymath` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.7 KiB | [postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-15-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_arraymath_14` | `1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 19.5 KiB | [pg_arraymath_14-1.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_arraymath_14-1.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_arraymath_14` | `1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 19.1 KiB | [pg_arraymath_14-1.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_arraymath_14-1.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_arraymath_14` | `1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 19.4 KiB | [pg_arraymath_14-1.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_arraymath_14-1.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_arraymath_14` | `1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 19.0 KiB | [pg_arraymath_14-1.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_arraymath_14-1.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_arraymath_14` | `1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 19.6 KiB | [pg_arraymath_14-1.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_arraymath_14-1.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_arraymath_14` | `1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 19.5 KiB | [pg_arraymath_14-1.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_arraymath_14-1.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-arraymath` | `1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 24.6 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 24.5 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 24.8 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 24.8 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 27.4 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 27.3 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 25.9 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-arraymath` | `1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 25.7 KiB | [postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-arraymath/postgresql-14-pg-arraymath_1.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pramsey/pgsql-arraymath" title="Repository" icon="github" subtitle="github.com/pramsey/pgsql-arraymath" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgsql-arraymath-1.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_arraymath;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_arraymath;		# install via package name, for the active PG version
pig install arraymath;		# install by extension name, for the current active PG version

pig install arraymath -v 18;   # install for PG 18
pig install arraymath -v 17;   # install for PG 17
pig install arraymath -v 16;   # install for PG 16
pig install arraymath -v 15;   # install for PG 15
pig install arraymath -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION arraymath;
```



## Usage

> [arraymath: element-by-element array operations for PostgreSQL](https://github.com/pramsey/pgsql-arraymath)

Provides element-by-element operators and utility functions for integer, float, and numeric arrays.

```sql
CREATE EXTENSION arraymath;
```

### Operators

All operators are prefixed with `@` to indicate element-by-element behavior. Works with array-vs-array (same length or cycling) and array-vs-scalar.

| Operator | Description | Returns |
|---|---|---|
| `@=` | Element-by-element equality | `boolean[]` |
| `@<` | Element-by-element less than | `boolean[]` |
| `@>` | Element-by-element greater than | `boolean[]` |
| `@<=` | Element-by-element less than or equals | `boolean[]` |
| `@>=` | Element-by-element greater than or equals | `boolean[]` |
| `@+` | Element-by-element addition | same type |
| `@-` | Element-by-element subtraction | same type |
| `@*` | Element-by-element multiplication | same type |
| `@/` | Element-by-element division | same type |

### Functions

| Function | Description |
|---|---|
| `array_sum(anyarray)` | Sum of all elements |
| `array_avg(anyarray)` | Average of all elements |
| `array_min(anyarray)` | Minimum element |
| `array_max(anyarray)` | Maximum element |
| `array_median(anyarray)` | Median element |
| `array_sort(anyarray)` | Sort ascending |
| `array_rsort(anyarray)` | Sort descending |

### Examples

```sql
-- Array vs scalar
SELECT ARRAY[1,2,3,4] @< 4;             -- {t,t,t,f}
SELECT ARRAY[3.4,5.6,7.6] @* 8.1;       -- {27.54,45.36,61.56}

-- Array vs array (cycling shorter array)
SELECT ARRAY[1,2,3,4,5,6] @* ARRAY[1,2]; -- {1,4,3,8,5,12}
SELECT ARRAY[1,2,3] @= ARRAY[3,2,1];     -- {f,t,f}

-- Utility functions
SELECT array_sort(ARRAY[9,1,8,2,7]);     -- {1,2,7,8,9}
SELECT array_sum(ARRAY[1,2,3,4,5]);      -- 15
SELECT array_median(ARRAY[1,2,3,4,5]);   -- 3
```
