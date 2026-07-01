---
title: "biscuit"
linkTitle: "biscuit"
description: "IAM-LIKE pattern matching with bitmap indexing"
weight: 2170
categories: ["FTS"]
width: full
---

[**pg_biscuit**](https://github.com/CrystallineCore/Biscuit) : IAM-LIKE pattern matching with bitmap indexing


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2170** | {{< badge content="biscuit" link="https://github.com/CrystallineCore/Biscuit" >}} | {{< ext "biscuit" "pg_biscuit" >}} | `2.4.1` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_textsearch" >}} |

> [!Note] rename from pg_biscuit to biscuit to keep up with PGDG RPM name


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_biscuit` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.1` | {{< bg "18" "biscuit_18" "green" >}} {{< bg "17" "biscuit_17" "green" >}} {{< bg "16" "biscuit_16" "green" >}} {{< bg "15" "biscuit_15" "red" >}} {{< bg "14" "biscuit_14" "red" >}} | `biscuit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.1` | {{< bg "18" "postgresql-18-biscuit" "green" >}} {{< bg "17" "postgresql-17-biscuit" "green" >}} {{< bg "16" "postgresql-16-biscuit" "green" >}} {{< bg "15" "postgresql-15-biscuit" "red" >}} {{< bg "14" "postgresql-14-biscuit" "red" >}} | `postgresql-$v-biscuit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_16 : AVAIL 3" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_16 : AVAIL 3" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.1" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.0" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_18` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [biscuit_18-2.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_18-2.4.1-1PIGSTY.el8.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.4 KiB | [biscuit_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.7 KiB | [biscuit_18-2.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_18-2.4.1-1PIGSTY.el8.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.2 KiB | [biscuit_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_18-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.3 KiB | [biscuit_18-2.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_18-2.4.1-1PIGSTY.el9.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.8 KiB | [biscuit_18-2.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.2 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.3 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.2.2-1PGDG.rhel9.7.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.4 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.2.2-1PGDG.rhel9.6.x86_64.rpm) |
| `biscuit_18` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 63.3 KiB | [biscuit_18-2.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_18-2.4.1-1PIGSTY.el9.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.9 KiB | [biscuit_18-2.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.2.2-1PGDG.rhel9.7.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.2.2-1PGDG.rhel9.6.aarch64.rpm) |
| `biscuit_18` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.7 KiB | [biscuit_18-2.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_18-2.4.1-1PIGSTY.el10.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.7 KiB | [biscuit_18-2.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.9 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.9 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.2.2-1PGDG.rhel10.1.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.2.2-1PGDG.rhel10.0.x86_64.rpm) |
| `biscuit_18` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.9 KiB | [biscuit_18-2.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_18-2.4.1-1PIGSTY.el10.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.4 KiB | [biscuit_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-biscuit` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.8 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 138.7 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.9 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 139.7 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 146.1 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 143.3 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 141.9 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 140.2 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 141.7 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.9 KiB | [postgresql-18-biscuit_2.4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_17` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.3 KiB | [biscuit_17-2.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_17-2.4.1-1PIGSTY.el8.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.2 KiB | [biscuit_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.6 KiB | [biscuit_17-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.5 KiB | [biscuit_17-2.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_17-2.4.1-1PIGSTY.el8.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_17-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.2 KiB | [biscuit_17-2.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_17-2.4.1-1PIGSTY.el9.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.5 KiB | [biscuit_17-2.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.2.2-1PGDG.rhel9.7.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.1 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.2.2-1PGDG.rhel9.6.x86_64.rpm) |
| `biscuit_17` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.8 KiB | [biscuit_17-2.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_17-2.4.1-1PIGSTY.el9.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.4 KiB | [biscuit_17-2.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.5 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.2.2-1PGDG.rhel9.7.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.2.2-1PGDG.rhel9.6.aarch64.rpm) |
| `biscuit_17` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.5 KiB | [biscuit_17-2.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_17-2.4.1-1PIGSTY.el10.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.5 KiB | [biscuit_17-2.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.2.2-1PGDG.rhel10.1.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.5 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.2.2-1PGDG.rhel10.0.x86_64.rpm) |
| `biscuit_17` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.5 KiB | [biscuit_17-2.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_17-2.4.1-1PIGSTY.el10.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [biscuit_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-biscuit` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.7 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 138.2 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.7 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 139.2 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 169.7 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.2 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 141.9 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 139.5 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 141.2 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.3 KiB | [postgresql-17-biscuit_2.4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_16` | `2.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 65.4 KiB | [biscuit_16-2.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_16-2.4.1-1PIGSTY.el8.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.3 KiB | [biscuit_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `2.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.5 KiB | [biscuit_16-2.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_16-2.4.1-1PIGSTY.el8.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_16-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `2.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 64.3 KiB | [biscuit_16-2.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_16-2.4.1-1PIGSTY.el9.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.6 KiB | [biscuit_16-2.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.2.2-1PGDG.rhel9.7.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.1 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.2.2-1PGDG.rhel9.6.x86_64.rpm) |
| `biscuit_16` | `2.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.8 KiB | [biscuit_16-2.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_16-2.4.1-1PIGSTY.el9.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.4 KiB | [biscuit_16-2.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.2.2-1PGDG.rhel9.7.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.2.2-1PGDG.rhel9.6.aarch64.rpm) |
| `biscuit_16` | `2.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.6 KiB | [biscuit_16-2.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_16-2.4.1-1PIGSTY.el10.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.5 KiB | [biscuit_16-2.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.2.2-1PGDG.rhel10.1.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.6 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.2.2-1PGDG.rhel10.0.x86_64.rpm) |
| `biscuit_16` | `2.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.6 KiB | [biscuit_16-2.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_16-2.4.1-1PIGSTY.el10.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [biscuit_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-biscuit` | `2.4.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.7 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 138.3 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.8 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 139.1 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 169.6 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 166.2 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 141.8 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 139.5 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 141.3 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.1 KiB | [postgresql-16-biscuit_2.4.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrystallineCore/Biscuit" title="Repository" icon="github" subtitle="github.com/CrystallineCore/Biscuit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="Biscuit-2.4.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_biscuit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_biscuit;		# install via package name, for the active PG version
pig install biscuit;		# install by extension name, for the current active PG version

pig install biscuit -v 18;   # install for PG 18
pig install biscuit -v 17;   # install for PG 17
pig install biscuit -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION biscuit CASCADE; -- requires plpgsql
```




## Usage

Sources:

- [PGXN biscuit 2.4.1](https://pgxn.org/dist/biscuit/2.4.1/)
- [Biscuit README](https://github.com/CrystallineCore/Biscuit)
- [Biscuit CHANGELOG](https://github.com/CrystallineCore/Biscuit/blob/main/CHANGELOG.md)
- [Biscuit documentation](https://biscuit.readthedocs.io/)
- [Local package metadata](../db/extension.csv)

`biscuit` is a PostgreSQL index access method for accelerating `LIKE`, `NOT LIKE`, `ILIKE`, and `NOT ILIKE` pattern matching on text. It uses bitmap-style position indexes to avoid the heap recheck overhead common in trigram searches and supports multi-column indexes for wildcard-heavy workloads.

PGXN package 2.4.1 ships the SQL/control version `2.4.0`; the extension's visible `default_version` is therefore still `2.4.0`. The local Pigsty extension name is `biscuit`, while older package metadata may mention `pg_biscuit`.

> [!WARNING]
> Upstream marks Biscuit as actively developed and recommends thorough staging validation before production use. Test representative datasets, query patterns, upgrades, backup/restore, and performance behavior before relying on it for critical workloads.

### Quick Start

```sql
CREATE EXTENSION IF NOT EXISTS biscuit;

CREATE TABLE users (
  id bigserial PRIMARY KEY,
  name text,
  email text,
  bio text
);

CREATE INDEX users_name_biscuit
ON users USING biscuit (name);

SELECT *
FROM users
WHERE name LIKE '%john%';
```

`biscuit` supports ordinary wildcard patterns with `%` and `_`:

```sql
SELECT * FROM users WHERE name LIKE 'john%';
SELECT * FROM users WHERE name LIKE '%smith';
SELECT * FROM users WHERE name LIKE '%oh_';
SELECT * FROM users WHERE name ILIKE '%john%';
SELECT * FROM users WHERE name NOT LIKE '%test%';
```

### Multi-Column Indexes

```sql
CREATE INDEX users_search_biscuit
ON users USING biscuit (name, email, bio);

SELECT *
FROM users
WHERE name ILIKE '%john%'
  AND email LIKE '%example.com'
  AND bio NOT LIKE '%inactive%';
```

Biscuit can combine bitmap matches from multiple indexed columns and may reorder predicates by estimated selectivity.

### Expression Indexes

Version 2.4.0 adds expression-index support:

```sql
CREATE INDEX users_lower_name_biscuit
ON users USING biscuit (lower(name));

SELECT *
FROM users
WHERE lower(name) LIKE '%john%';
```

For `char(n)` / `bpchar` columns, upstream recommends expression indexes that cast to `text`, because native `bpchar` operator classes are not yet available:

```sql
CREATE INDEX legacy_code_biscuit
ON legacy_table USING biscuit ((code::text));
```

### Introspection

```sql
SELECT *
FROM biscuit_operators;

SELECT *
FROM biscuit_version_history;
```

The `biscuit_operators` view lists the operators registered for the Biscuit access method. In 2.4.0, the view was fixed to remain correct if additional operator classes or families are added.

### Operational Notes

Biscuit's design is optimized for:

- prefix, suffix, substring, and mixed wildcard `LIKE` / `ILIKE` patterns
- multi-column predicates where bitmap intersections can reduce candidate sets
- exact pattern matching without trigram false-positive rechecks
- workloads where text-pattern search dominates query latency

It is not a general full-text search engine and does not replace ranking, stemming, tokenization, or phrase search. Use PostgreSQL full text search, trigram indexes, or dedicated search extensions when those semantics are required.

### Caveats

- Upstream requires PostgreSQL 16 or newer and standard build tools. Pigsty local metadata currently packages Biscuit for PostgreSQL 16-18.
- PGXN package version 2.4.1 carries SQL/control `default_version = '2.4.0'`; this is expected for the current source package.
- Biscuit only targets `LIKE` / `ILIKE`-style wildcard matching. Regular expressions are not the supported search surface.
- Non-text columns should be indexed through explicit text expressions when needed.
- Benchmark against `pg_trgm` and your actual data distribution before replacing existing production indexes.
