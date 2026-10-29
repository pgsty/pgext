---
title: "pg_tre"
linkTitle: "pg_tre"
description: "Native PostgreSQL index access method for approximate regular-expression matching"
weight: 2330
categories: ["FTS"]
languages: ["C"]
licenses: ["MIT AND BSD-2-Clause"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_tre**](https://codeberg.org/gregburd/pg_tre) : Native PostgreSQL index access method for approximate regular-expression matching


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2330** | {{< badge content="pg_tre" link="https://codeberg.org/gregburd/pg_tre" >}} | {{< ext "pg_tre" >}} | `4.5.3` | {{< category "FTS" >}} | {{< license "MIT AND BSD-2-Clause" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "pg_textsearch" >}} {{< ext "pg_fts" >}} {{< ext "vector" >}} {{< ext "pg_mentat" >}} |

> [!Note] Upstream 4.5.3 fixes a TRE approximate-matcher worklist overflow reachable by query users that could crash a backend and restart the cluster. SQL 4.5.3; update matching library and extension SQL. No on-disk change or new REINDEX requirement; older WAL/format and omitted-posting repair boundaries still apply.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_tre` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.3` | {{< bg "18" "pg_tre_18" "green" >}} {{< bg "17" "pg_tre_17" "green" >}} {{< bg "16" "pg_tre_16" "red" >}} {{< bg "15" "pg_tre_15" "red" >}} {{< bg "14" "pg_tre_14" "red" >}} | `pg_tre_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `4.5.3` | {{< bg "18" "postgresql-18-pg-tre" "green" >}} {{< bg "17" "postgresql-17-pg-tre" "green" >}} {{< bg "16" "postgresql-16-pg-tre" "red" >}} {{< bg "15" "postgresql-15-pg-tre" "red" >}} {{< bg "14" "postgresql-14-pg-tre" "red" >}} | `postgresql-$v-pg-tre` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "pg_tre_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "pg_tre_16 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_15 : N/A 0" "gray" >}} | {{< bg "N/A" "pg_tre_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-18-pg-tre : AVAIL 1" "green" >}} | {{< bg "PIGSTY 4.5.3" "postgresql-17-pg-tre : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-pg-tre : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-pg-tre : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tre_18` | `4.5.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 592.5 KiB | [pg_tre_18-4.5.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tre_18-4.5.3-1PGSTY.el8.x86_64.rpm) |
| `pg_tre_18` | `4.5.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 576.4 KiB | [pg_tre_18-4.5.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tre_18-4.5.3-1PGSTY.el8.aarch64.rpm) |
| `pg_tre_18` | `4.5.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 562.6 KiB | [pg_tre_18-4.5.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tre_18-4.5.3-1PGSTY.el9.x86_64.rpm) |
| `pg_tre_18` | `4.5.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 553.8 KiB | [pg_tre_18-4.5.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tre_18-4.5.3-1PGSTY.el9.aarch64.rpm) |
| `pg_tre_18` | `4.5.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 569.5 KiB | [pg_tre_18-4.5.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tre_18-4.5.3-1PGSTY.el10.x86_64.rpm) |
| `pg_tre_18` | `4.5.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 558.1 KiB | [pg_tre_18-4.5.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tre_18-4.5.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-tre` | `4.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 574.2 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 555.0 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 576.5 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 558.7 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 588.6 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 578.9 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 572.4 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 562.3 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 569.5 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-pg-tre` | `4.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 557.7 KiB | [postgresql-18-pg-tre_4.5.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tre/postgresql-18-pg-tre_4.5.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_tre_17` | `4.5.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 593.2 KiB | [pg_tre_17-4.5.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_tre_17-4.5.3-1PGSTY.el8.x86_64.rpm) |
| `pg_tre_17` | `4.5.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 577.1 KiB | [pg_tre_17-4.5.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_tre_17-4.5.3-1PGSTY.el8.aarch64.rpm) |
| `pg_tre_17` | `4.5.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 563.4 KiB | [pg_tre_17-4.5.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_tre_17-4.5.3-1PGSTY.el9.x86_64.rpm) |
| `pg_tre_17` | `4.5.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 554.6 KiB | [pg_tre_17-4.5.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_tre_17-4.5.3-1PGSTY.el9.aarch64.rpm) |
| `pg_tre_17` | `4.5.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 570.1 KiB | [pg_tre_17-4.5.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_tre_17-4.5.3-1PGSTY.el10.x86_64.rpm) |
| `pg_tre_17` | `4.5.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 559.1 KiB | [pg_tre_17-4.5.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_tre_17-4.5.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-tre` | `4.5.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 575.1 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 555.5 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 577.2 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 558.5 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 645.5 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 632.8 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 573.5 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 563.0 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 570.1 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-pg-tre` | `4.5.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 558.1 KiB | [postgresql-17-pg-tre_4.5.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-tre/postgresql-17-pg-tre_4.5.3-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://codeberg.org/gregburd/pg_tre" title="Repository" icon="link" subtitle="codeberg.org/gregburd/pg_tre" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_tre-4.5.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_tre;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_tre;		# install via package name, for the active PG version

pig install pg_tre -v 18;   # install for PG 18
pig install pg_tre -v 17;   # install for PG 17

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION pg_tre;
```

## Usage

Sources:

- [README.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/README.md)
- [pg_tre.control](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/pg_tre.control)
- [sql/pg_tre--4.5.3.sql](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/sql/pg_tre--4.5.3.sql)
- [CHANGELOG.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/CHANGELOG.md)
- [LIMITATIONS.md](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/LIMITATIONS.md)
- [META.json](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/META.json)
- [.github/workflows/ci.yml](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/.github/workflows/ci.yml)
- [Library initialization](https://api.pgxn.org/src/pg_tre/pg_tre-4.5.3/src/module.c)

`pg_tre` 4.5.3 provides the `tre` native index access method for approximate regular-expression matching over text. Upstream supports PostgreSQL 17–19; Pigsty's 4.5.2 packages target PostgreSQL 17–18. It narrows candidates with codepoint trigrams and postings, then checks matches against the heap. Use it for typo-tolerant identifiers, log messages and lexical patterns with explicit edit costs.

### Enable and Search

The index write path uses PostgreSQL's generic WAL facility. Since 4.0.0, `CREATE EXTENSION` is sufficient: `shared_preload_libraries` and a restart are not required to enable a new installation. Extension creation requires a superuser, and the control file is untrusted and non-relocatable.

```sql
CREATE EXTENSION pg_tre;
CREATE TABLE documents (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  body text NOT NULL
);
INSERT INTO documents (body) VALUES
  ('The PostgreSQL database system'),
  ('The Postgres databse system');
CREATE INDEX documents_body_tre ON documents USING tre (body);
SELECT id, body, body <@> tre_pattern('database', 1) AS distance
FROM documents
WHERE body %~~ tre_pattern('database', 1)
ORDER BY distance ASC NULLS LAST, id;
```

PostgreSQL 19 uses the core regular-expression character classifier; ICU character-class edge cases on PostgreSQL 17–18 may differ from native PostgreSQL regex semantics, as described upstream.

### Patterns and Objects

- `tre_pattern(text)` uses `pg_tre.default_max_cost`; `tre_pattern(text, int4)` supplies a maximum cost; the five-argument form also sets insertion, deletion and substitution costs. Patterns use POSIX extended regular expressions and TRE approximate budgets.
- `text %~~ tre_pattern` is the indexable approximate-match predicate; `text <@> tre_pattern` returns the edit cost or NULL when no match fits the budget.
- `tre_distance`, `tre_similarity`, `tre_amatch`, `tre_amatch_cost` and `tre_amatch_detail` expose matching and compatibility functions.
- The index operator family also supports LIKE, ILIKE, regular-expression and equality predicates with heap rechecks. The separately named trigram functions coexist with `pg_trgm`; `pg_tre` does not install the colliding bare trigram operators removed in 3.0.

Indexes accept one text column with `tre_text_ops`. Expression and partial indexes are available, but index-only scans are not. Short patterns, high edit budgets and non-literal regexes can produce broad candidate sets; inspect real plans and keep cancellation and query-time limits enabled.

### Maintenance and Limits

`fastupdate` must be true; setting it false is now rejected because the direct path could not preserve the index format. Existing off settings can still be opened, but inserts use the pending list. `pending_list_limit` is in KiB and defaults to 0, meaning unlimited. A positive value triggers a synchronous merge when the threshold is crossed; that inserting backend pays the work, and if another merge is active the limit can temporarily be exceeded.

`pg_tre.min_trigram_freq` must remain 1. Values above 1 formerly omitted rare postings and could lose results; they are now rejected. Keep regex state and compilation/matching limits bounded for untrusted patterns. Index construction uses tuplesort and temporary disk, with memory governed by maintenance_work_mem; estimate large builds before scheduling them.

```sql
SELECT * FROM tre_estimate_index_build('documents'::regclass, 2);
SELECT * FROM pg_tre_index_pending('documents_body_tre'::regclass);
SELECT * FROM pg_tre_index_check('documents_body_tre'::regclass, 2);
```

`pg_tre_index_check` compares index answers with sequential scans. Each sample performs multiple sequential scans, so bound the sample count on large tables. Version 4.5.1 fixes its search-path handling, pins trusted lookup paths and qualifies calls; use the fixed version before running this helper as a privileged user.

### Upgrade Boundaries

Install matching files, ensure sessions use the new library, and update each database:

```sql
ALTER EXTENSION pg_tre UPDATE TO '4.5.3';
```

Crossing from 3.x to 4.x changes WAL record types even though the index page format is unchanged. Shut down cleanly before replacing the library; a physical standby cannot run a different major extension version from its primary and must be rebuilt from a fresh base backup after the primary upgrade. There is no rolling upgrade across this boundary. Remove any old `wal_consistency_checking = 'pg_tre'` setting; it prevents startup in 4.x. Upstream recommends `wal_consistency_checking = 'all'` when consistency checking is needed.

Version 4.4 writes format 11 and reads formats 6–10. Upgrading from those formats does not require a routine rebuild; optional REINDEX obtains the newer dense-posting representation. Downgrading to 4.3 after a 4.4+ write requires rebuilding the affected indexes for the older format. Older upgrades may cross additional boundaries described in the changelog.

REINDEX every index built with min_trigram_freq greater than 1: installing the fix cannot restore omitted postings. Version 4.5 also disables stale SuRF acceleration on indexes from earlier affected versions, restoring correctness immediately; a rebuild is optional to regain that acceleration. Upgrading 4.5.0 to 4.5.1 changes neither format nor ordinary query behavior and requires no additional REINDEX. Old definitions using the removed tuple bloom options must still be corrected before rebuilding.

### 4.5.2 Regex Correction

Version 4.5.2 fixes newline-mode `(?n)` anchors and negated character classes. It retains the index format and needs no additional REINDEX when upgrading from 4.5.1. The older WAL-transition and omitted-posting rebuild requirements above still apply when crossing those affected versions.

```sql
SELECT E'a\nfoo' %~~ tre_pattern('(?n)^foo', 0) AS line_start,
       E'foo\n' %~~ tre_pattern('(?n)foo$', 0) AS line_end,
       E'\n' %~~ tre_pattern('(?n)[^x]', 0) AS negated_class;
```

The expected results are true, true and false. The fix affects matching and index rechecks alike; patterns without newline mode retain their previous behavior.

### 4.5.3 Security Fix

An approximate pattern with many reachable states could overflow TRE's fixed work list and crash a backend, causing the postmaster to restart the cluster's other backends. Version 4.5.3 grows that work list within the existing pattern limits. Install the matching library, replace sessions using the old library, and update extension SQL to 4.5.3. The patch changes no on-disk format and adds no REINDEX requirement; the older migration and damaged-index boundaries above remain.
