---
title: "financial"
linkTitle: "financial"
description: "Financial aggregate functions"
weight: 4840
categories: ["FUNC"]
width: full
---

[**pg_financial**](https://github.com/intgr/pg_financial) : Financial aggregate functions


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4840** | {{< badge content="financial" link="https://github.com/intgr/pg_financial" >}} | {{< ext "financial" "pg_financial" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "topn" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "count_distinct" >}} {{< ext "omnisketch" >}} {{< ext "ddsketch" >}} {{< ext "tdigest" >}} {{< ext "first_last_agg" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_financial` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "pg_financial_18*" "green" >}} {{< bg "17" "pg_financial_17*" "green" >}} {{< bg "16" "pg_financial_16*" "green" >}} {{< bg "15" "pg_financial_15*" "green" >}} {{< bg "14" "pg_financial_14*" "green" >}} {{< bg "13" "pg_financial_13*" "green" >}} | `pg_financial_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "18" "postgresql-18-pg-financial" "green" >}} {{< bg "17" "postgresql-17-pg-financial" "green" >}} {{< bg "16" "postgresql-16-pg-financial" "green" >}} {{< bg "15" "postgresql-15-pg-financial" "green" >}} {{< bg "14" "postgresql-14-pg-financial" "green" >}} {{< bg "13" "postgresql-13-pg-financial" "green" >}} | `postgresql-$v-pg-financial` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "pg_financial_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-14-pg-financial : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-13-pg-financial : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_financial_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.3 KiB | [pg_financial_18-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_18-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 KiB | [pg_financial_18-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_18-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.3 KiB | [pg_financial_18-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_18-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.4 KiB | [pg_financial_18-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_18-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.3 KiB | [pg_financial_18-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_financial_18-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_financial_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.6 KiB | [pg_financial_18-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_financial_18-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-financial` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.5 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.7 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.3 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-financial` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-18-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-18-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_financial_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.3 KiB | [pg_financial_17-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_17-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 KiB | [pg_financial_17-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_17-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.3 KiB | [pg_financial_17-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_17-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.4 KiB | [pg_financial_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.3 KiB | [pg_financial_17-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_financial_17-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_financial_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.6 KiB | [pg_financial_17-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_financial_17-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-financial` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.8 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.0 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.8 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.0 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-financial` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-17-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-17-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_financial_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.3 KiB | [pg_financial_16-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_16-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 KiB | [pg_financial_16-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_16-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.3 KiB | [pg_financial_16-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_16-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.4 KiB | [pg_financial_16-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_16-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.3 KiB | [pg_financial_16-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_financial_16-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_financial_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.6 KiB | [pg_financial_16-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_financial_16-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-financial` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.8 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.0 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.3 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-financial` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-16-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-16-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_financial_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.3 KiB | [pg_financial_15-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_15-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 KiB | [pg_financial_15-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_15-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.3 KiB | [pg_financial_15-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_15-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.4 KiB | [pg_financial_15-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_15-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.3 KiB | [pg_financial_15-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_financial_15-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_financial_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.6 KiB | [pg_financial_15-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_financial_15-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-financial` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.8 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 14.0 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-financial` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-15-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-15-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_financial_14` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.3 KiB | [pg_financial_14-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_14-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_14` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 KiB | [pg_financial_14-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_14-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_14` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.3 KiB | [pg_financial_14-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_14-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_14` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.4 KiB | [pg_financial_14-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_14-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_14` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.3 KiB | [pg_financial_14-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_financial_14-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_financial_14` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.6 KiB | [pg_financial_14-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_financial_14-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-financial` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.9 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.8 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.0 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.8 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.9 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.3 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-financial` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.3 KiB | [postgresql-14-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-14-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_financial_13` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 14.3 KiB | [pg_financial_13-1.0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_financial_13-1.0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_financial_13` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 14.5 KiB | [pg_financial_13-1.0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_financial_13-1.0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_financial_13` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 14.3 KiB | [pg_financial_13-1.0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_financial_13-1.0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_financial_13` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 14.4 KiB | [pg_financial_13-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_financial_13-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_financial_13` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 14.2 KiB | [pg_financial_13-1.0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_financial_13-1.0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_financial_13` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 14.5 KiB | [pg_financial_13-1.0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_financial_13-1.0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-financial` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.9 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 13.0 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 12.9 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 13.1 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.6 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.7 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 13.4 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-financial` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 13.4 KiB | [postgresql-13-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-financial/postgresql-13-pg-financial_1.0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/intgr/pg_financial" title="Repository" icon="github" subtitle="github.com/intgr/pg_financial" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_financial-1.0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_financial;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_financial;		# install via package name, for the active PG version
pig install financial;		# install by extension name, for the current active PG version

pig install financial -v 18;   # install for PG 18
pig install financial -v 17;   # install for PG 17
pig install financial -v 16;   # install for PG 16
pig install financial -v 15;   # install for PG 15
pig install financial -v 14;   # install for PG 14
pig install financial -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION financial;
```
