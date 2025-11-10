---
title: "pg_xenophile"
linkTitle: "pg_xenophile"
description: "More than the bare necessities for PostgreSQL i18n and l10n."
weight: 3610
categories: ["TYPE"]
width: full
---

[**pg_xenophile**](https://github.com/bigsmoke/pg_xenophile)


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3610** | {{< badge content="pg_xenophile" link="https://github.com/bigsmoke/pg_xenophile" >}} | {{< ext "pg_xenophile" >}} | `0.8.3` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Need By**    | {{< ext "l10n_table_dependent_extension" >}} |
|   **See Also**    | {{< ext "country" >}} {{< ext "currency" >}} {{< ext "icu_ext" >}} {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} |
|    **Siblings**   | {{< ext "l10n_table_dependent_extension" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_xenophile" >}} | `0.8.3` | {{< bg "18" "pg_xenophile_18" "green" >}} {{< bg "17" "pg_xenophile_17" "green" >}} {{< bg "16" "pg_xenophile_16" "green" >}} {{< bg "15" "pg_xenophile_15" "green" >}} {{< bg "14" "pg_xenophile_14" "green" >}} {{< bg "13" "pg_xenophile_13" "green" >}} | `pg_xenophile_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_xenophile" >}} | `0.8.3` | {{< bg "18" "postgresql-18-pg-xenophile" "green" >}} {{< bg "17" "postgresql-17-pg-xenophile" "green" >}} {{< bg "16" "postgresql-16-pg-xenophile" "green" >}} {{< bg "15" "postgresql-15-pg-xenophile" "green" >}} {{< bg "14" "postgresql-14-pg-xenophile" "green" >}} {{< bg "13" "postgresql-13-pg-xenophile" "green" >}} | `postgresql-$v-pg-xenophile` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "pg_xenophile_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-18-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-17-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-16-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-15-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-14-pg-xenophile : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.8.3" "postgresql-13-pg-xenophile : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xenophile_18` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.1 KiB | [pg_xenophile_18-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xenophile_18-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_xenophile_18` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_xenophile_18-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xenophile_18-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_xenophile_18` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.5 KiB | [pg_xenophile_18-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xenophile_18-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_xenophile_18` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.4 KiB | [pg_xenophile_18-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xenophile_18-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_xenophile_18` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.6 KiB | [pg_xenophile_18-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xenophile_18-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_xenophile_18` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [pg_xenophile_18-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xenophile_18-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 46.0 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.0 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 45.9 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.0 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.6 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.7 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 48.2 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-xenophile` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.2 KiB | [postgresql-18-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-18-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xenophile_17` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.1 KiB | [pg_xenophile_17-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xenophile_17-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_xenophile_17` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_xenophile_17-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xenophile_17-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_xenophile_17` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.5 KiB | [pg_xenophile_17-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xenophile_17-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_xenophile_17` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.4 KiB | [pg_xenophile_17-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xenophile_17-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_xenophile_17` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.6 KiB | [pg_xenophile_17-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xenophile_17-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_xenophile_17` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [pg_xenophile_17-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xenophile_17-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 46.0 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.0 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.0 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.7 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.7 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 48.2 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-xenophile` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.2 KiB | [postgresql-17-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-17-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xenophile_16` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.1 KiB | [pg_xenophile_16-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xenophile_16-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_xenophile_16` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_xenophile_16-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xenophile_16-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_xenophile_16` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.5 KiB | [pg_xenophile_16-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xenophile_16-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_xenophile_16` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.4 KiB | [pg_xenophile_16-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xenophile_16-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_xenophile_16` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.6 KiB | [pg_xenophile_16-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xenophile_16-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_xenophile_16` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [pg_xenophile_16-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xenophile_16-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 46.0 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.0 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.0 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.6 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.7 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 48.2 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-xenophile` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.2 KiB | [postgresql-16-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-16-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xenophile_15` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.1 KiB | [pg_xenophile_15-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xenophile_15-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_xenophile_15` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_xenophile_15-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xenophile_15-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_xenophile_15` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.5 KiB | [pg_xenophile_15-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xenophile_15-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_xenophile_15` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.5 KiB | [pg_xenophile_15-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xenophile_15-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_xenophile_15` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.6 KiB | [pg_xenophile_15-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xenophile_15-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_xenophile_15` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [pg_xenophile_15-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xenophile_15-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 46.0 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.0 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.0 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.7 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.7 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 48.2 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-xenophile` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.2 KiB | [postgresql-15-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-15-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xenophile_14` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.1 KiB | [pg_xenophile_14-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xenophile_14-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_xenophile_14` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_xenophile_14-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xenophile_14-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_xenophile_14` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.5 KiB | [pg_xenophile_14-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xenophile_14-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_xenophile_14` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.5 KiB | [pg_xenophile_14-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xenophile_14-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_xenophile_14` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.6 KiB | [pg_xenophile_14-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xenophile_14-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_xenophile_14` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [pg_xenophile_14-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xenophile_14-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 46.0 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 46.0 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.0 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.7 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.7 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 48.2 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-xenophile` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.2 KiB | [postgresql-14-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-14-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_xenophile_13` | `0.8.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 49.1 KiB | [pg_xenophile_13-0.8.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_xenophile_13-0.8.3-1PIGSTY.el8.x86_64.rpm) |
| `pg_xenophile_13` | `0.8.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 49.1 KiB | [pg_xenophile_13-0.8.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_xenophile_13-0.8.3-1PIGSTY.el8.aarch64.rpm) |
| `pg_xenophile_13` | `0.8.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.5 KiB | [pg_xenophile_13-0.8.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_xenophile_13-0.8.3-1PIGSTY.el9.x86_64.rpm) |
| `pg_xenophile_13` | `0.8.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.5 KiB | [pg_xenophile_13-0.8.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_xenophile_13-0.8.3-1PIGSTY.el9.aarch64.rpm) |
| `pg_xenophile_13` | `0.8.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.6 KiB | [pg_xenophile_13-0.8.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_xenophile_13-0.8.3-1PIGSTY.el10.x86_64.rpm) |
| `pg_xenophile_13` | `0.8.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [pg_xenophile_13-0.8.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_xenophile_13-0.8.3-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 45.9 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 45.9 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 46.0 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 46.0 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 48.6 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 48.7 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 48.2 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-xenophile` | `0.8.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 48.2 KiB | [postgresql-13-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-xenophile/postgresql-13-pg-xenophile_0.8.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bigsmoke/pg_xenophile" title="Repository" icon="github" subtitle="github.com/bigsmoke/pg_xenophile" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_xenophile-0.8.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_xenophile; # get pg_xenophile source code
pig build dep pg_xenophile; # install build dependencies
pig build pkg pg_xenophile; # build extension rpm or deb
pig build ext pg_xenophile; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_xenophile; # install by extension name, for the current active PG version
pig ext install pg_xenophile; # install via package alias, for the active PG version
pig ext install pg_xenophile -v 18;   # install for PG 18
pig ext install pg_xenophile -v 17;   # install for PG 17
pig ext install pg_xenophile -v 16;   # install for PG 16
pig ext install pg_xenophile -v 15;   # install for PG 15
pig ext install pg_xenophile -v 14;   # install for PG 14
pig ext install pg_xenophile -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_xenophile CASCADE SCHEMA xeno;
```

