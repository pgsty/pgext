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
| **2170** | {{< badge content="biscuit" link="https://github.com/CrystallineCore/Biscuit" >}} | {{< ext "biscuit" "pg_biscuit" >}} | `2.4.3` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "hll" >}} {{< ext "rum" >}} {{< ext "pg_textsearch" >}} |

> [!Note] Latest stable PGXN distribution and package release is 2.4.3; 2.5.0 is testing; packaged control and SQL default version remain 2.4.1; package name is biscuit.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_biscuit` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.3` | {{< bg "18" "biscuit_18" "green" >}} {{< bg "17" "biscuit_17" "green" >}} {{< bg "16" "biscuit_16" "green" >}} {{< bg "15" "biscuit_15" "red" >}} {{< bg "14" "biscuit_14" "red" >}} | `biscuit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.4.3` | {{< bg "18" "postgresql-18-biscuit" "green" >}} {{< bg "17" "postgresql-17-biscuit" "green" >}} {{< bg "16" "postgresql-16-biscuit" "green" >}} {{< bg "15" "postgresql-15-biscuit" "red" >}} {{< bg "14" "postgresql-14-biscuit" "red" >}} | `postgresql-$v-biscuit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 3" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 3" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} |      {{< bg "MISS" "biscuit_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "biscuit_14 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-15-biscuit : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-biscuit : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_18` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.8 KiB | [biscuit_18-2.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_18-2.4.3-1PIGSTY.el8.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.4 KiB | [biscuit_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 62.1 KiB | [biscuit_18-2.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_18-2.4.3-1PIGSTY.el8.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.2 KiB | [biscuit_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_18-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.6 KiB | [biscuit_18-2.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_18-2.4.3-1PIGSTY.el9.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.8 KiB | [biscuit_18-2.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.2 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.3 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.2.2-1PGDG.rhel9.7.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.4 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-2.2.2-1PGDG.rhel9.6.x86_64.rpm) |
| `biscuit_18` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.6 KiB | [biscuit_18-2.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_18-2.4.3-1PIGSTY.el9.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.9 KiB | [biscuit_18-2.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.2.2-1PGDG.rhel9.7.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-2.2.2-1PGDG.rhel9.6.aarch64.rpm) |
| `biscuit_18` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.2 KiB | [biscuit_18-2.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_18-2.4.3-1PIGSTY.el10.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.7 KiB | [biscuit_18-2.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.9 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 67.9 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.2.2-1PGDG.rhel10.1.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-2.2.2-1PGDG.rhel10.0.x86_64.rpm) |
| `biscuit_18` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 64.2 KiB | [biscuit_18-2.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_18-2.4.3-1PIGSTY.el10.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.4 KiB | [biscuit_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-biscuit` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.0 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 138.0 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.2 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 138.8 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 145.2 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 142.4 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 140.9 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 139.2 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 140.7 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-biscuit` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 138.0 KiB | [postgresql-18-biscuit_2.4.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-18-biscuit_2.4.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_17` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.7 KiB | [biscuit_17-2.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_17-2.4.3-1PIGSTY.el8.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.2 KiB | [biscuit_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.6 KiB | [biscuit_17-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.9 KiB | [biscuit_17-2.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_17-2.4.3-1PIGSTY.el8.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_17-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.6 KiB | [biscuit_17-2.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_17-2.4.3-1PIGSTY.el9.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.5 KiB | [biscuit_17-2.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.2.2-1PGDG.rhel9.7.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.1 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-2.2.2-1PGDG.rhel9.6.x86_64.rpm) |
| `biscuit_17` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.2 KiB | [biscuit_17-2.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_17-2.4.3-1PIGSTY.el9.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.4 KiB | [biscuit_17-2.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.5 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.2.2-1PGDG.rhel9.7.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-2.2.2-1PGDG.rhel9.6.aarch64.rpm) |
| `biscuit_17` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.0 KiB | [biscuit_17-2.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_17-2.4.3-1PIGSTY.el10.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.5 KiB | [biscuit_17-2.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.2.2-1PGDG.rhel10.1.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.5 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-2.2.2-1PGDG.rhel10.0.x86_64.rpm) |
| `biscuit_17` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 63.9 KiB | [biscuit_17-2.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_17-2.4.3-1PIGSTY.el10.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [biscuit_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-biscuit` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 142.9 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 137.5 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.0 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 138.5 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.6 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.2 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 140.8 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 138.5 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 140.2 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-biscuit` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 137.2 KiB | [postgresql-17-biscuit_2.4.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-17-biscuit_2.4.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_16` | `2.4.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 64.7 KiB | [biscuit_16-2.4.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_16-2.4.3-1PIGSTY.el8.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.3 KiB | [biscuit_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `2.4.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 61.9 KiB | [biscuit_16-2.4.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_16-2.4.3-1PIGSTY.el8.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_16-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `2.4.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 63.6 KiB | [biscuit_16-2.4.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_16-2.4.3-1PIGSTY.el9.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 62.6 KiB | [biscuit_16-2.4.0-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.4.0-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.2.2-1PGDG.rhel9.8.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.7.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.2.2-1PGDG.rhel9.7.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.x86_64](/os/el9.x86_64) | pgdg | 65.1 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.6.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-2.2.2-1PGDG.rhel9.6.x86_64.rpm) |
| `biscuit_16` | `2.4.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 62.2 KiB | [biscuit_16-2.4.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_16-2.4.3-1PIGSTY.el9.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 61.4 KiB | [biscuit_16-2.4.0-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.4.0-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.2.2-1PGDG.rhel9.8.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.6 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.7.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.2.2-1PGDG.rhel9.7.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el9.aarch64](/os/el9.aarch64) | pgdg | 62.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel9.6.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-2.2.2-1PGDG.rhel9.6.aarch64.rpm) |
| `biscuit_16` | `2.4.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 65.0 KiB | [biscuit_16-2.4.3-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_16-2.4.3-1PIGSTY.el10.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 64.5 KiB | [biscuit_16-2.4.0-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.4.0-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.2.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.2.2-1PGDG.rhel10.2.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.0 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.2.2-1PGDG.rhel10.1.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.x86_64](/os/el10.x86_64) | pgdg | 68.6 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.0.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-2.2.2-1PGDG.rhel10.0.x86_64.rpm) |
| `biscuit_16` | `2.4.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 63.9 KiB | [biscuit_16-2.4.3-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_16-2.4.3-1PIGSTY.el10.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [biscuit_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-biscuit` | `2.4.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 142.9 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 137.6 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.0 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 138.5 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 168.6 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 165.1 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 140.8 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 138.6 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 140.2 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-biscuit` | `2.4.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 137.2 KiB | [postgresql-16-biscuit_2.4.3-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-16-biscuit_2.4.3-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrystallineCore/Biscuit" title="Repository" icon="github" subtitle="github.com/CrystallineCore/Biscuit" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="Biscuit-2.4.3.tar.gz" >}}
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

- [PGXN biscuit 2.4.3 distribution](https://pgxn.org/dist/biscuit/2.4.3/)
- [PGXN 2.4.3 metadata](https://api.pgxn.org/dist/biscuit/2.4.3/META.json)
- [PGXN 2.4.3 control file](https://api.pgxn.org/src/biscuit/biscuit-2.4.3/biscuit.control)
- [PGXN 2.4.3 changelog](https://api.pgxn.org/src/biscuit/biscuit-2.4.3/CHANGELOG.md)
- [Official documentation](https://biscuit.readthedocs.io/)

`biscuit` is an experimental PostgreSQL index access method optimized for pattern filters on text. It targets selective `LIKE`, `ILIKE`, `NOT LIKE`, and `NOT ILIKE` predicates, including multi-column and expression indexes, while trading additional memory and write work for faster filtering.

### Core Workflow

```sql
CREATE EXTENSION biscuit;

CREATE INDEX message_body_biscuit_idx
ON message USING biscuit (body);

SELECT id, body
FROM message
WHERE body ILIKE '%timeout%';
```

Expression indexes work when the query uses the same expression:

```sql
CREATE INDEX customer_email_biscuit_idx
ON customer USING biscuit (lower(email));

SELECT *
FROM customer
WHERE lower(email) LIKE '%@example.com';
```

For predicates spanning several indexed text columns, use a multi-column index and confirm the chosen plan with `EXPLAIN (ANALYZE, BUFFERS)` on representative data.

### Important Objects

- `biscuit` is the index access method used in `CREATE INDEX ... USING biscuit`.
- `biscuit_operators` reports the supported operators.
- `biscuit_version` and `biscuit_build_info` expose build information; `biscuit_build_info_json` returns it as JSON.
- `biscuit_status` reports the installed build and bitmap configuration.
- `biscuit_index_stats` and `biscuit_index_memory_size` inspect an index and its memory footprint.
- `biscuit_memory_usage` is a view of extension memory use.
- `biscuit_has_roaring` and `biscuit_roaring_version` report optional Roaring bitmap support.

### Limits and Operations

`biscuit` is for filtering, not ordered index scans. It does not provide regular-expression or similarity search. Indexes can be larger and more expensive to maintain than a B-tree; benchmark read selectivity, ingest cost, memory use, and vacuum behavior before production use. Keep a conventional index where ordering, equality, uniqueness, or another access method is still required.

The upstream project labels Biscuit as actively developed. PGXN publishes `2.4.3` as a stable distribution, but that archive's changelog stops at `2.4.2`, and its metadata and control file expose SQL extension version `2.4.1`. Treat `2.4.3` as a distribution/package refresh: no additional SQL API delta is claimed. The material `2.4.2` change fixes a use-after-free in the index cache plus compiler warnings.
