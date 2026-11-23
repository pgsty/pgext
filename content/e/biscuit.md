---
title: "biscuit"
linkTitle: "biscuit"
description: "IAM-LIKE pattern matching with bitmap indexing"
weight: 2180
categories: ["FTS"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**pg_biscuit**](https://github.com/CrystallineCore/Biscuit) : IAM-LIKE pattern matching with bitmap indexing


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2180** | {{< badge content="biscuit" link="https://github.com/CrystallineCore/Biscuit" >}} | {{< ext "biscuit" "pg_biscuit" >}} | `3.1.1` | {{< category "FTS" >}} | {{< license "MIT" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **Requires**    | {{< ext "plpgsql" >}} |
|   **See Also**    | {{< ext "pg_trgm" >}} {{< ext "pg_similarity" >}} {{< ext "fuzzystrmatch" >}} {{< ext "smlar" >}} {{< ext "pg_bigm" >}} {{< ext "pgpcre" >}} {{< ext "re2" >}} {{< ext "pgroonga" >}} |

> [!Note] REINDEX 2.x and crash-recovered unlogged indexes; 3.0 format unchanged.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} | `pg_biscuit` | `plpgsql` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.1` | {{< bg "19" "biscuit_19" "green" >}} {{< bg "18" "biscuit_18" "green" >}} {{< bg "17" "biscuit_17" "green" >}} {{< bg "16" "biscuit_16" "green" >}} {{< bg "15" "biscuit_15" "red" >}} | `biscuit_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `3.1.1` | {{< bg "19" "postgresql-19-biscuit" "green" >}} {{< bg "18" "postgresql-18-biscuit" "green" >}} {{< bg "17" "postgresql-17-biscuit" "green" >}} {{< bg "16" "postgresql-16-biscuit" "green" >}} {{< bg "15" "postgresql-15-biscuit" "red" >}} | `postgresql-$v-biscuit` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "biscuit_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_16 : AVAIL 6" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "biscuit_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_18 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_17 : AVAIL 6" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_16 : AVAIL 6" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_19 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_16 : AVAIL 3" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_19 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_16 : AVAIL 3" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_19 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_18 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_17 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_16 : AVAIL 3" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_19 : AVAIL 3" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_18 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_17 : AVAIL 9" "green" >}} | {{< bg "PIGSTY 3.1.1" "biscuit_16 : AVAIL 9" "green" >}} | {{< bg "N/A" "biscuit_15 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-19-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-18-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-17-biscuit : AVAIL 1" "green" >}} | {{< bg "PIGSTY 3.1.1" "postgresql-16-biscuit : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-biscuit : N/A 0" "gray" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_19` | `3.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 266.1 KiB | [biscuit_19-3.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_19-3.1.1-1PGSTY.el9.x86_64.rpm) |
| `biscuit_19` | `3.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 102.9 KiB | [biscuit_19-3.1.1-1PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/biscuit_19-3.1.1-1PGDG.rhel9.9.x86_64.rpm) |
| `biscuit_19` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.6 KiB | [biscuit_19-3.0.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-x86_64/biscuit_19-3.0.0-2PGDG.rhel9.9.x86_64.rpm) |
| `biscuit_19` | `3.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 260.5 KiB | [biscuit_19-3.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_19-3.1.1-1PGSTY.el9.aarch64.rpm) |
| `biscuit_19` | `3.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 99.7 KiB | [biscuit_19-3.1.1-1PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/biscuit_19-3.1.1-1PGDG.rhel9.9.aarch64.rpm) |
| `biscuit_19` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.8 KiB | [biscuit_19-3.0.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-9-aarch64/biscuit_19-3.0.0-2PGDG.rhel9.9.aarch64.rpm) |
| `biscuit_19` | `3.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 269.9 KiB | [biscuit_19-3.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_19-3.1.1-1PGSTY.el10.x86_64.rpm) |
| `biscuit_19` | `3.1.1` | [el10.x86_64](/os/el10.x86_64) | pgdg | 105.7 KiB | [biscuit_19-3.1.1-1PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/biscuit_19-3.1.1-1PGDG.rhel10.3.x86_64.rpm) |
| `biscuit_19` | `3.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 100.2 KiB | [biscuit_19-3.0.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-x86_64/biscuit_19-3.0.0-2PGDG.rhel10.3.x86_64.rpm) |
| `biscuit_19` | `3.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 263.2 KiB | [biscuit_19-3.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_19-3.1.1-1PGSTY.el10.aarch64.rpm) |
| `biscuit_19` | `3.1.1` | [el10.aarch64](/os/el10.aarch64) | pgdg | 101.6 KiB | [biscuit_19-3.1.1-1PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/biscuit_19-3.1.1-1PGDG.rhel10.3.aarch64.rpm) |
| `biscuit_19` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.2 KiB | [biscuit_19-3.0.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/19/redhat/rhel-10-aarch64/biscuit_19-3.0.0-2PGDG.rhel10.3.aarch64.rpm) |
| `postgresql-19-biscuit` | `3.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 269.9 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 261.2 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 270.3 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 262.3 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 270.3 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 264.2 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 262.2 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 257.6 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 260.7 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-biscuit` | `3.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 254.0 KiB | [postgresql-19-biscuit_3.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-19-biscuit_3.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_18` | `3.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 265.9 KiB | [biscuit_18-3.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_18-3.1.1-1PGSTY.el8.x86_64.rpm) |
| `biscuit_18` | `3.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 270.2 KiB | [biscuit_18-3.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_18-3.1.0-1PGSTY.el8.x86_64.rpm) |
| `biscuit_18` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.0 KiB | [biscuit_18-3.0.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-3.0.0-2PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.0 KiB | [biscuit_18-3.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-3.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.4 KiB | [biscuit_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.6 KiB | [biscuit_18-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/biscuit_18-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_18` | `3.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 256.6 KiB | [biscuit_18-3.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_18-3.1.1-1PGSTY.el8.aarch64.rpm) |
| `biscuit_18` | `3.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 262.0 KiB | [biscuit_18-3.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_18-3.1.0-1PGSTY.el8.aarch64.rpm) |
| `biscuit_18` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.2 KiB | [biscuit_18-3.0.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-3.0.0-2PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.2 KiB | [biscuit_18-3.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-3.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.2 KiB | [biscuit_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_18-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/biscuit_18-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_18` | `3.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 265.6 KiB | [biscuit_18-3.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_18-3.1.1-1PGSTY.el9.x86_64.rpm) |
| `biscuit_18` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 270.4 KiB | [biscuit_18-3.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_18-3.1.0-1PGSTY.el9.x86_64.rpm) |
| `biscuit_18` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.7 KiB | [biscuit_18-3.0.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/biscuit_18-3.0.0-2PGDG.rhel9.9.x86_64.rpm) |
| `biscuit_18` | `3.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 260.4 KiB | [biscuit_18-3.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_18-3.1.1-1PGSTY.el9.aarch64.rpm) |
| `biscuit_18` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 265.3 KiB | [biscuit_18-3.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_18-3.1.0-1PGSTY.el9.aarch64.rpm) |
| `biscuit_18` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.9 KiB | [biscuit_18-3.0.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/biscuit_18-3.0.0-2PGDG.rhel9.9.aarch64.rpm) |
| `biscuit_18` | `3.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 269.5 KiB | [biscuit_18-3.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_18-3.1.1-1PGSTY.el10.x86_64.rpm) |
| `biscuit_18` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 273.5 KiB | [biscuit_18-3.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_18-3.1.0-1PGSTY.el10.x86_64.rpm) |
| `biscuit_18` | `3.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 100.3 KiB | [biscuit_18-3.0.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/biscuit_18-3.0.0-2PGDG.rhel10.3.x86_64.rpm) |
| `biscuit_18` | `3.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 263.1 KiB | [biscuit_18-3.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_18-3.1.1-1PGSTY.el10.aarch64.rpm) |
| `biscuit_18` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 267.1 KiB | [biscuit_18-3.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_18-3.1.0-1PGSTY.el10.aarch64.rpm) |
| `biscuit_18` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.5 KiB | [biscuit_18-3.0.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-3.0.0-2PGDG.rhel10.3.aarch64.rpm) |
| `biscuit_18` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.4 KiB | [biscuit_18-3.0.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-3.0.0-2PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.2 KiB | [biscuit_18-3.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-3.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.4 KiB | [biscuit_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_18` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.8 KiB | [biscuit_18-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-aarch64/biscuit_18-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-18-biscuit` | `3.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 269.1 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 260.9 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 269.6 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 262.1 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 269.6 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 263.9 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 261.6 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 257.3 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 260.1 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-biscuit` | `3.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 254.0 KiB | [postgresql-18-biscuit_3.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-18-biscuit_3.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_17` | `3.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 265.4 KiB | [biscuit_17-3.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_17-3.1.1-1PGSTY.el8.x86_64.rpm) |
| `biscuit_17` | `3.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 269.7 KiB | [biscuit_17-3.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_17-3.1.0-1PGSTY.el8.x86_64.rpm) |
| `biscuit_17` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.0 KiB | [biscuit_17-3.0.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-3.0.0-2PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.0 KiB | [biscuit_17-3.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-3.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.2 KiB | [biscuit_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.6 KiB | [biscuit_17-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/biscuit_17-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_17` | `3.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 256.1 KiB | [biscuit_17-3.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_17-3.1.1-1PGSTY.el8.aarch64.rpm) |
| `biscuit_17` | `3.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 261.4 KiB | [biscuit_17-3.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_17-3.1.0-1PGSTY.el8.aarch64.rpm) |
| `biscuit_17` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.2 KiB | [biscuit_17-3.0.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-3.0.0-2PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.1 KiB | [biscuit_17-3.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-3.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_17-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/biscuit_17-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_17` | `3.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 265.1 KiB | [biscuit_17-3.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_17-3.1.1-1PGSTY.el9.x86_64.rpm) |
| `biscuit_17` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 269.9 KiB | [biscuit_17-3.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_17-3.1.0-1PGSTY.el9.x86_64.rpm) |
| `biscuit_17` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.7 KiB | [biscuit_17-3.0.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/biscuit_17-3.0.0-2PGDG.rhel9.9.x86_64.rpm) |
| `biscuit_17` | `3.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 259.9 KiB | [biscuit_17-3.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_17-3.1.1-1PGSTY.el9.aarch64.rpm) |
| `biscuit_17` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 264.6 KiB | [biscuit_17-3.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_17-3.1.0-1PGSTY.el9.aarch64.rpm) |
| `biscuit_17` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.8 KiB | [biscuit_17-3.0.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/biscuit_17-3.0.0-2PGDG.rhel9.9.aarch64.rpm) |
| `biscuit_17` | `3.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 269.0 KiB | [biscuit_17-3.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_17-3.1.1-1PGSTY.el10.x86_64.rpm) |
| `biscuit_17` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 273.0 KiB | [biscuit_17-3.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_17-3.1.0-1PGSTY.el10.x86_64.rpm) |
| `biscuit_17` | `3.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 100.3 KiB | [biscuit_17-3.0.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/biscuit_17-3.0.0-2PGDG.rhel10.3.x86_64.rpm) |
| `biscuit_17` | `3.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 262.6 KiB | [biscuit_17-3.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_17-3.1.1-1PGSTY.el10.aarch64.rpm) |
| `biscuit_17` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 266.5 KiB | [biscuit_17-3.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_17-3.1.0-1PGSTY.el10.aarch64.rpm) |
| `biscuit_17` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.4 KiB | [biscuit_17-3.0.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-3.0.0-2PGDG.rhel10.3.aarch64.rpm) |
| `biscuit_17` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.4 KiB | [biscuit_17-3.0.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-3.0.0-2PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.2 KiB | [biscuit_17-3.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-3.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.1 KiB | [biscuit_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_17` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_17-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-aarch64/biscuit_17-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-17-biscuit` | `3.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 268.6 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 260.8 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 269.5 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 261.6 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 313.3 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 308.1 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 260.9 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 256.7 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 259.4 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-biscuit` | `3.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 253.3 KiB | [postgresql-17-biscuit_3.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-17-biscuit_3.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `biscuit_16` | `3.1.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 265.3 KiB | [biscuit_16-3.1.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_16-3.1.1-1PGSTY.el8.x86_64.rpm) |
| `biscuit_16` | `3.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 269.7 KiB | [biscuit_16-3.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/biscuit_16-3.1.0-1PGSTY.el8.x86_64.rpm) |
| `biscuit_16` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.0 KiB | [biscuit_16-3.0.0-2PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-3.0.0-2PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `3.0.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 99.0 KiB | [biscuit_16-3.0.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-3.0.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `2.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 62.3 KiB | [biscuit_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-2.4.0-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `2.2.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 63.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel8.10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/biscuit_16-2.2.2-1PGDG.rhel8.10.x86_64.rpm) |
| `biscuit_16` | `3.1.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 256.1 KiB | [biscuit_16-3.1.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_16-3.1.1-1PGSTY.el8.aarch64.rpm) |
| `biscuit_16` | `3.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 261.3 KiB | [biscuit_16-3.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/biscuit_16-3.1.0-1PGSTY.el8.aarch64.rpm) |
| `biscuit_16` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.2 KiB | [biscuit_16-3.0.0-2PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-3.0.0-2PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `3.0.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 92.1 KiB | [biscuit_16-3.0.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-3.0.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-2.4.0-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el8.aarch64](/os/el8.aarch64) | pgdg | 59.1 KiB | [biscuit_16-2.2.2-1PGDG.rhel8.10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/biscuit_16-2.2.2-1PGDG.rhel8.10.aarch64.rpm) |
| `biscuit_16` | `3.1.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 265.0 KiB | [biscuit_16-3.1.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_16-3.1.1-1PGSTY.el9.x86_64.rpm) |
| `biscuit_16` | `3.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 270.0 KiB | [biscuit_16-3.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/biscuit_16-3.1.0-1PGSTY.el9.x86_64.rpm) |
| `biscuit_16` | `3.0.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 97.7 KiB | [biscuit_16-3.0.0-2PGDG.rhel9.9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/biscuit_16-3.0.0-2PGDG.rhel9.9.x86_64.rpm) |
| `biscuit_16` | `3.1.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 259.9 KiB | [biscuit_16-3.1.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_16-3.1.1-1PGSTY.el9.aarch64.rpm) |
| `biscuit_16` | `3.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 264.7 KiB | [biscuit_16-3.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/biscuit_16-3.1.0-1PGSTY.el9.aarch64.rpm) |
| `biscuit_16` | `3.0.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 94.9 KiB | [biscuit_16-3.0.0-2PGDG.rhel9.9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/biscuit_16-3.0.0-2PGDG.rhel9.9.aarch64.rpm) |
| `biscuit_16` | `3.1.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 269.0 KiB | [biscuit_16-3.1.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_16-3.1.1-1PGSTY.el10.x86_64.rpm) |
| `biscuit_16` | `3.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 273.0 KiB | [biscuit_16-3.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/biscuit_16-3.1.0-1PGSTY.el10.x86_64.rpm) |
| `biscuit_16` | `3.0.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 100.3 KiB | [biscuit_16-3.0.0-2PGDG.rhel10.3.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/biscuit_16-3.0.0-2PGDG.rhel10.3.x86_64.rpm) |
| `biscuit_16` | `3.1.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 262.6 KiB | [biscuit_16-3.1.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_16-3.1.1-1PGSTY.el10.aarch64.rpm) |
| `biscuit_16` | `3.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 266.5 KiB | [biscuit_16-3.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/biscuit_16-3.1.0-1PGSTY.el10.aarch64.rpm) |
| `biscuit_16` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.4 KiB | [biscuit_16-3.0.0-2PGDG.rhel10.3.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-3.0.0-2PGDG.rhel10.3.aarch64.rpm) |
| `biscuit_16` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.4 KiB | [biscuit_16-3.0.0-2PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-3.0.0-2PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `3.0.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 96.1 KiB | [biscuit_16-3.0.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-3.0.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.4.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 63.2 KiB | [biscuit_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.4.0-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.2.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.2.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.1.aarch64.rpm) |
| `biscuit_16` | `2.2.2` | [el10.aarch64](/os/el10.aarch64) | pgdg | 64.7 KiB | [biscuit_16-2.2.2-1PGDG.rhel10.0.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/biscuit_16-2.2.2-1PGDG.rhel10.0.aarch64.rpm) |
| `postgresql-16-biscuit` | `3.1.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 268.9 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 260.8 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 269.5 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 261.7 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 312.7 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 307.5 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 261.0 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 256.8 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 259.5 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-biscuit` | `3.1.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 253.3 KiB | [postgresql-16-biscuit_3.1.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pg-biscuit/postgresql-16-biscuit_3.1.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrystallineCore/Biscuit" title="Repository" icon="github" subtitle="github.com/CrystallineCore/Biscuit" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="Biscuit-3.1.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pg_biscuit;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pg_biscuit;		# install via package name, for the active PG version
pig install biscuit;		# install by extension name, for the current active PG version

pig install biscuit -v 19;   # install for PG 19
pig install biscuit -v 18;   # install for PG 18
pig install biscuit -v 17;   # install for PG 17
pig install biscuit -v 16;   # install for PG 16

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION biscuit CASCADE; -- requires plpgsql
```

## Usage

Sources:

- [Biscuit 3.1.0 README](https://api.pgxn.org/src/biscuit/biscuit-3.1.1/README.md)
- [Biscuit 3.1.0 CHANGELOG](https://api.pgxn.org/src/biscuit/biscuit-3.1.1/CHANGELOG.md)
- [Biscuit 3.1.0 SQL](https://api.pgxn.org/src/biscuit/biscuit-3.1.1/sql/biscuit--3.1.0.sql)
- [Biscuit 3.0.0 to 3.1.0 SQL](https://api.pgxn.org/src/biscuit/biscuit-3.1.1/sql/biscuit--3.0.0--3.1.0.sql)
- [Biscuit 3.0.0 on PGXN](https://pgxn.org/dist/biscuit/3.0.0/)
- [Biscuit 3.0.0 release](https://github.com/CrystallineCore/Biscuit/releases/tag/v3.0.0)
- [Biscuit 3.0.0 README](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/README.md)
- [Biscuit 3.0.0 changelog](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/CHANGELOG.md)
- [Biscuit 3.0.0 metadata](https://api.pgxn.org/dist/biscuit/3.0.0/META.json)
- [Biscuit 3.0.0 control file](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/biscuit.control)
- [Biscuit 3.0.0 Makefile](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/Makefile)
- [Biscuit 3.0.0 installation SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit.sql)
- [Biscuit 2.5.0 to 3.0.0 upgrade SQL](https://github.com/CrystallineCore/Biscuit/blob/v3.0.0/sql/biscuit--2.5.0--3.0.0.sql)

`biscuit` 3.1.0 is a PostgreSQL 16+ positional-bitmap index access method for exact `LIKE` and `ILIKE` filtering. It is strongest for anchored patterns, `_` wildcards, length predicates, and multi-column conjunctions. Version 3.0.0 stores index state in WAL-logged relation pages, so crash recovery, point-in-time recovery, physical replication, and hot-standby reads use PostgreSQL's ordinary recovery path. It does not require `shared_preload_libraries` or a restart.

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

Biscuit is a filtering index: it does not provide ordered, backward, index-only, or unique scans, cannot back `CLUSTER`, and does not support similarity search, fuzzy search, or general locale-aware collation. A B-tree with `text_pattern_ops` is usually a better fit for selective prefix lookups, while `pg_trgm` is designed for unanchored substring, regular-expression, and similarity searches.

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

### Version 3.1.0

Version 3.1.0 adds `~`, `!~`, `~*`, and `!~*` operator-class entries, but only regexes exactly reducible to a `LIKE` glob are candidates for acceleration. Anchors, literals, dots, and selected repetition forms are supported; alternation, bracket classes, groups, and backreferences are not. Unsupported expressions retain PostgreSQL semantics and normally use a sequential scan. Case-sensitive `~` and `!~` can be exact. `~*` is limited to pure-ASCII patterns under a suitable collation and needs an executor recheck; `!~*` is not accelerated. Narrow operator classes retain their case-sensitivity boundary.

```sql
ALTER EXTENSION biscuit UPDATE TO '3.1.0';
SELECT id, body FROM message WHERE body ~ '^timeout.*$';
```

The update from 3.0.0 adds SQL operator-family entries without changing the index format, so it does not require a format rebuild. The earlier 2.x-to-3.0.0 rebuild requirement above still applies. This release fixes HOT-chain and partial-index build correctness, anchored-match recovery, and pending-list memory safety. Unlogged Biscuit indexes still require `REINDEX` after crash recovery.

### Version 3.1.1

Version 3.1.1 adds PostgreSQL 19 support while retaining PostgreSQL 16–18. From 3.1.0, deploy the new shared library; this patch needs neither ALTER EXTENSION nor REINDEX because it changes neither SQL objects nor the disk format. Recycle sessions that already loaded the older library, or restart through the normal maintenance procedure. `SELECT biscuit_version();` reports the loaded library version. This does not remove the older 2.x migration or crash-recovery requirements above.
