---
title: "pg_cooldown"
linkTitle: "pg_cooldown"
description: "remove buffered pages for specific relations"
weight: 5070
categories: ["ADMIN"]
width: full
---

[**pg_cooldown**](https://github.com/rbergm/pg_cooldown) : remove buffered pages for specific relations


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5070** | {{< badge content="pg_cooldown" link="https://github.com/rbergm/pg_cooldown" >}} | {{< ext "pg_cooldown" >}} | `0.1` | {{< category "ADMIN" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgfincore" >}} {{< ext "pgcozy" >}} {{< ext "pg_prewarm" >}} {{< ext "pg_buffercache" >}} {{< ext "pg_repack" >}} {{< ext "pg_squeeze" >}} {{< ext "system_stats" >}} {{< ext "pgmeminfo" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "red" >}} | `pg_cooldown` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1` | {{< bg "18" "pg_cooldown_18" "green" >}} {{< bg "17" "pg_cooldown_17" "green" >}} {{< bg "16" "pg_cooldown_16" "green" >}} {{< bg "15" "pg_cooldown_15" "green" >}} {{< bg "14" "pg_cooldown_14" "green" >}} {{< bg "13" "pg_cooldown_13" "red" >}} | `pg_cooldown_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1` | {{< bg "18" "postgresql-18-pg-cooldown" "green" >}} {{< bg "17" "postgresql-17-pg-cooldown" "green" >}} {{< bg "16" "postgresql-16-pg-cooldown" "green" >}} {{< bg "15" "postgresql-15-pg-cooldown" "green" >}} {{< bg "14" "postgresql-14-pg-cooldown" "green" >}} {{< bg "13" "postgresql-13-pg-cooldown" "red" >}} | `postgresql-$v-pg-cooldown` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_cooldown_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pg_cooldown_14 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_cooldown_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pg-cooldown : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pg-cooldown : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-cooldown : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-13-pg-cooldown : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pg-cooldown : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pg-cooldown : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pg-cooldown : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pg-cooldown : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pg-cooldown : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cooldown_18` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_cooldown_18-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_18-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_18` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pg_cooldown_18-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_18-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_18-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_18-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.1 KiB | [pg_cooldown_18-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_18-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_18-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cooldown_18-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cooldown_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.2 KiB | [pg_cooldown_18-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cooldown_18-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-cooldown` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.9 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.9 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 12.4 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 12.4 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-cooldown` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-18-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-18-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cooldown_17` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_cooldown_17-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_17-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_17` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pg_cooldown_17-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_17-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_17` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_17-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_17-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_17` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.1 KiB | [pg_cooldown_17-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_17-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_17` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_17-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cooldown_17-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cooldown_17` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.2 KiB | [pg_cooldown_17-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cooldown_17-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-cooldown` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.8 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.9 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-cooldown` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-17-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-17-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cooldown_16` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_cooldown_16-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_16-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_16` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pg_cooldown_16-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_16-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_16` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_16-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_16-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_16` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.1 KiB | [pg_cooldown_16-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_16-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_16` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_16-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cooldown_16-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cooldown_16` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.3 KiB | [pg_cooldown_16-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cooldown_16-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-cooldown` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.8 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.9 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-cooldown` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-16-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-16-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cooldown_15` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_cooldown_15-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_15-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_15` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pg_cooldown_15-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_15-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_15` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 KiB | [pg_cooldown_15-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_15-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_15` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.1 KiB | [pg_cooldown_15-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_15-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_15` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_15-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cooldown_15-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cooldown_15` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.3 KiB | [pg_cooldown_15-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cooldown_15-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-cooldown` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.8 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 12.0 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.2 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-cooldown` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-15-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-15-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cooldown_14` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.2 KiB | [pg_cooldown_14-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_14-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_14` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pg_cooldown_14-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_14-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_14` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.1 KiB | [pg_cooldown_14-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_14-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_14` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.1 KiB | [pg_cooldown_14-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_14-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pg_cooldown_14` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 16.0 KiB | [pg_cooldown_14-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_cooldown_14-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pg_cooldown_14` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 16.2 KiB | [pg_cooldown_14-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_cooldown_14-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-cooldown` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 11.8 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 11.9 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 11.8 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 11.9 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.2 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.4 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-cooldown` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.2 KiB | [postgresql-14-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-14-pg-cooldown_0.1-2PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_cooldown_13` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 16.3 KiB | [pg_cooldown_13-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_cooldown_13-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pg_cooldown_13` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 16.4 KiB | [pg_cooldown_13-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_cooldown_13-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pg_cooldown_13` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 16.4 KiB | [pg_cooldown_13-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_cooldown_13-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pg_cooldown_13` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 16.3 KiB | [pg_cooldown_13-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_cooldown_13-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-13-pg-cooldown` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 12.7 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-cooldown` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 12.8 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-cooldown` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 13.1 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-cooldown` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 13.1 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-cooldown` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 12.5 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-cooldown` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 12.3 KiB | [postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-cooldown/postgresql-13-pg-cooldown_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/rbergm/pg_cooldown" title="Repository" icon="github" subtitle="github.com/rbergm/pg_cooldown" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_cooldown-0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_cooldown;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_cooldown;		# install via package name, for the active PG version

pig install pg_cooldown -v 18;   # install for PG 18
pig install pg_cooldown -v 17;   # install for PG 17
pig install pg_cooldown -v 16;   # install for PG 16
pig install pg_cooldown -v 15;   # install for PG 15
pig install pg_cooldown -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_cooldown;
```
