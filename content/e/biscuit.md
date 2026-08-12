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
| **2170** | {{< badge content="biscuit" link="https://github.com/CrystallineCore/Biscuit" >}} | {{< ext "biscuit" "pg_biscuit" >}} | `3.0.0` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "pg_similarity" >}} {{< ext "fuzzystrmatch" >}} {{< ext "smlar" >}} {{< ext "pg_bigm" >}} {{< ext "pgpcre" >}} {{< ext "re2" >}} {{< ext "pgroonga" >}} |

> [!Note] Latest stable PGXN distribution and packaged extension version are 3.0.0; upgrading from 2.x requires REINDEX; package name is biscuit.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `pg_biscuit` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.0` | {{< bg "18" "biscuit_18" "green" >}} {{< bg "17" "biscuit_17" "green" >}} {{< bg "16" "biscuit_16" "green" >}} {{< bg "15" "biscuit_15" "red" >}} {{< bg "14" "biscuit_14" "red" >}} | `biscuit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.0.0` | {{< bg "18" "postgresql-18-biscuit" "green" >}} {{< bg "17" "postgresql-17-biscuit" "green" >}} {{< bg "16" "postgresql-16-biscuit" "green" >}} {{< bg "15" "postgresql-15-biscuit" "red" >}} {{< bg "14" "postgresql-14-biscuit" "red" >}} | `postgresql-$v-biscuit` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 3" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "biscuit_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 3" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "biscuit_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "biscuit_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "biscuit_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "biscuit_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_18 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_17 : AVAIL 5" "green" >}} | {{< bg "PIGSTY 2.4.3" "biscuit_16 : AVAIL 5" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} | {{< bg "N/A" "biscuit_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.4.3" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-biscuit : N/A 0" "gray" >}} |


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
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="Biscuit-3.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_biscuit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

- [Biscuit 3.0.0 on PGXN](https://pgxn.org/dist/biscuit/3.0.0/)
- [Biscuit 3.0.0 release](https://github.com/CrystallineCore/Biscuit/releases/tag/v3.0.0)
- [Biscuit 3.0.0 README](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/README.md)
- [Biscuit 3.0.0 changelog](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/CHANGELOG.md)
- [Biscuit 3.0.0 metadata](https://api.pgxn.org/dist/biscuit/3.0.0/META.json)
- [Biscuit 3.0.0 control file](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/biscuit.control)
- [Biscuit 3.0.0 Makefile](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/Makefile)
- [Biscuit 3.0.0 installation SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit.sql)
- [Biscuit 2.5.0 to 3.0.0 upgrade SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit--2.5.0--3.0.0.sql)

`biscuit` 3.0.0 is a PostgreSQL 16+ positional-bitmap index access method for exact `LIKE` and `ILIKE` filtering. It is strongest for anchored patterns, `_` wildcards, length predicates, and multi-column conjunctions. Version 3.0.0 stores index state in WAL-logged relation pages, so crash recovery, point-in-time recovery, physical replication, and hot-standby reads use PostgreSQL's ordinary recovery path. It does not require `shared_preload_libraries` or a restart.

The project remains under active development and recommends representative staging tests. Its per-connection memory, write amplification, and cache-reload behavior make it best suited to read-mostly analytical workloads rather than continuously updated OLTP tables or very large connection pools.

### Build and Query an Index

Load the data first, then create the index. The default `biscuit_ops` supports both case-sensitive and case-insensitive predicates. Use `biscuit_like_ops` or `biscuit_ilike_ops` when only one mode is required, avoiding the unused structure set.

```sql
CREATE EXTENSION biscuit;

CREATE INDEX message_body_biscuit_idx
ON message USING biscuit (body biscuit_like_ops);

ANALYZE message;

EXPLAIN (ANALYZE, BUFFERS)
SELECT id, body
FROM message
WHERE body LIKE 'timeout%';
```

Expression and multi-column indexes are supported. The query must use expressions and operators compatible with the chosen operator class. Check representative plans after loading statistics, especially for unanchored patterns.

### Operator Classes and Query Boundaries

- `biscuit_ops` is the default text operator class and indexes `LIKE`, `NOT LIKE`, `ILIKE`, and `NOT ILIKE`.
- `biscuit_like_ops` indexes only `LIKE` and `NOT LIKE`.
- `biscuit_ilike_ops` indexes only `ILIKE` and `NOT ILIKE`.

Biscuit returns exact matches without a heap recheck, but it is a filtering index: it does not provide ordered, backward, index-only, or unique scans, cannot back `CLUSTER`, and does not support regular expressions, similarity search, fuzzy search, or locale-aware collation. A B-tree with `text_pattern_ops` is usually a better fit for selective prefix lookups, while `pg_trgm` is designed for unanchored substring, regular-expression, and similarity searches.

### Diagnostics and Configuration

Important inspection objects include `biscuit_indexes`, `biscuit_status`, `biscuit_index_stats(oid)`, `biscuit_index_memory_size()`, `biscuit_pending_list_stats(oid)`, and `biscuit_pending_list_usage`. The memory function reports the current backend's session-local copy. `total_pending_bytes` is refreshed during `VACUUM`, so pending-list figures can lag live writes by up to one vacuum cycle.

- `biscuit.delta_compaction_slots` defaults to 20000 and controls how many pending rows are tolerated before compaction. It is a privileged setting because raising it can increase reload work for other sessions.
- `biscuit.diag_scan_trace` defaults to off and emits verbose per-scan candidate accounting. Enable it only for a focused reproducer.

Every backend lazily loads its own copy of an index and keeps it for the connection lifetime. A committed write invalidates other cached copies; their next access reloads the index rather than refreshing it incrementally. Size pools for this memory behavior and avoid interleaving frequent writes with latency-sensitive reads.

Live-index `INSERT` and `UPDATE` generate substantial WAL; monitor `pg_wal`, replication lag, and replication-slot retention, and consider a bounded `max_slot_wal_keep_size`. Bulk loading before index creation is substantially cheaper. `VACUUM` drains pending work but does not shrink the index; use `REINDEX` to reclaim index space.

### Upgrade to 3.0.0

Version 3.0.0 is an incompatible on-disk format change. Updating the extension catalog does not convert existing index pages: every Biscuit index created under 2.x must be rebuilt. Plan enough maintenance time and WAL capacity for the rebuild.

```sql
ALTER EXTENSION biscuit UPDATE TO '3.0.0';

SELECT schema_name, index_name
FROM biscuit_indexes;

REINDEX INDEX CONCURRENTLY public.message_body_biscuit_idx;
```

The unpatched upstream 3.0.0 archive ships and installs only the `2.5.0--3.0.0` step, while earlier stable packages exposed catalog versions `2.4.0` or `2.4.1`. Pigsty's 3.0.0 RPM and DEB packages restore that missing catalog path before applying the upstream step. For another source build or package, inspect `pg_extension_update_paths('biscuit')` before `ALTER EXTENSION`; regardless of the available SQL path, the mandatory `REINDEX` or `REINDEX CONCURRENTLY` remains a separate manual operation.
