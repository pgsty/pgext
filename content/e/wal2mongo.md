---
title: "wal2mongo"
linkTitle: "wal2mongo"
description: "PostgreSQL logical decoding output plugin for MongoDB"
weight: 9640
categories: ["ETL"]
width: full
---

[**wal2mongo**](https://github.com/HighgoSoftware/wal2mongo) : PostgreSQL logical decoding output plugin for MongoDB


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9640** | {{< badge content="wal2mongo" link="https://github.com/HighgoSoftware/wal2mongo" >}} | {{< ext "wal2mongo" >}} | `1.0.7` | {{< category "ETL" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s----" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="orange" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "mongo_fdw" >}} {{< ext "wal2json" >}} {{< ext "decoderbufs" >}} {{< ext "decoder_raw" >}} {{< ext "documentdb" >}} {{< ext "pglogical" >}} {{< ext "test_decoding" >}} {{< ext "pgoutput" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.7` | {{< bg "18" "" "red" >}} {{< bg "17" "" "red" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `wal2mongo` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.7` | {{< bg "18" "wal2mongo_18" "red" >}} {{< bg "17" "wal2mongo_17" "red" >}} {{< bg "16" "wal2mongo_16" "green" >}} {{< bg "15" "wal2mongo_15" "green" >}} {{< bg "14" "wal2mongo_14" "green" >}} | `wal2mongo_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.7` | {{< bg "18" "postgresql-18-wal2mongo" "red" >}} {{< bg "17" "postgresql-17-wal2mongo" "red" >}} {{< bg "16" "postgresql-16-wal2mongo" "green" >}} {{< bg "15" "postgresql-15-wal2mongo" "green" >}} {{< bg "14" "postgresql-14-wal2mongo" "green" >}} | `postgresql-$v-wal2mongo` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "wal2mongo_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "wal2mongo_17 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.7" "wal2mongo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_14 : AVAIL 1" "blue" >}} |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "wal2mongo_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "wal2mongo_17 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.7" "wal2mongo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "wal2mongo_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "wal2mongo_17 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.7" "wal2mongo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_14 : AVAIL 1" "blue" >}} |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "wal2mongo_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "wal2mongo_17 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.7" "wal2mongo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "wal2mongo_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "wal2mongo_17 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.7" "wal2mongo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_14 : AVAIL 1" "blue" >}} |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "wal2mongo_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "wal2mongo_17 : MISS 0" "red" >}}      | {{< bg "PGDG 1.0.7" "wal2mongo_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.0.7" "wal2mongo_14 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.7" "postgresql-16-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-15-wal2mongo : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.7" "postgresql-14-wal2mongo : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-wal2mongo : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-wal2mongo : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-wal2mongo : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2mongo_16` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.2 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/wal2mongo_16-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_16` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.0 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/wal2mongo_16-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_16` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/wal2mongo_16-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `wal2mongo_16` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [wal2mongo_16-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/wal2mongo_16-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `wal2mongo_16` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.6 KiB | [wal2mongo_16-1.0.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/wal2mongo_16-1.0.7-3PGDG.rhel10.x86_64.rpm) |
| `wal2mongo_16` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.3 KiB | [wal2mongo_16-1.0.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-aarch64/wal2mongo_16-1.0.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-wal2mongo` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.4 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.0 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.4 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 34.0 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.6 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.1 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.3 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-wal2mongo` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.2 KiB | [postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-16-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2mongo_15` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.2 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/wal2mongo_15-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_15` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/wal2mongo_15-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_15` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.0 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/wal2mongo_15-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `wal2mongo_15` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [wal2mongo_15-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/wal2mongo_15-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `wal2mongo_15` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.5 KiB | [wal2mongo_15-1.0.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/wal2mongo_15-1.0.7-3PGDG.rhel10.x86_64.rpm) |
| `wal2mongo_15` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.2 KiB | [wal2mongo_15-1.0.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-aarch64/wal2mongo_15-1.0.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-wal2mongo` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.1 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.8 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.1 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.7 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.3 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.7 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.1 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-wal2mongo` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-15-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `wal2mongo_14` | `1.0.7` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.2 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/wal2mongo_14-1.0.7-1PGDG.rhel8.x86_64.rpm) |
| `wal2mongo_14` | `1.0.7` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/wal2mongo_14-1.0.7-1PGDG.rhel8.aarch64.rpm) |
| `wal2mongo_14` | `1.0.7` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.0 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/wal2mongo_14-1.0.7-1PGDG.rhel9.x86_64.rpm) |
| `wal2mongo_14` | `1.0.7` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [wal2mongo_14-1.0.7-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/wal2mongo_14-1.0.7-1PGDG.rhel9.aarch64.rpm) |
| `wal2mongo_14` | `1.0.7` | [el10.x86_64](/os/el10.x86_64) | pgdg | 20.5 KiB | [wal2mongo_14-1.0.7-3PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/wal2mongo_14-1.0.7-3PGDG.rhel10.x86_64.rpm) |
| `wal2mongo_14` | `1.0.7` | [el10.aarch64](/os/el10.aarch64) | pgdg | 20.2 KiB | [wal2mongo_14-1.0.7-3PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-aarch64/wal2mongo_14-1.0.7-3PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-wal2mongo` | `1.0.7` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.1 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [d12.aarch64](/os/d12.aarch64) | pigsty | 33.7 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.1 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [d13.aarch64](/os/d13.aarch64) | pigsty | 33.7 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [u22.x86_64](/os/u22.x86_64) | pigsty | 39.2 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [u22.aarch64](/os/u22.aarch64) | pigsty | 38.6 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.1 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-wal2mongo` | `1.0.7` | [u24.aarch64](/os/u24.aarch64) | pigsty | 34.9 KiB | [postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/w/wal2mongo/postgresql-14-wal2mongo_1.0.7-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/HighgoSoftware/wal2mongo" title="Repository" icon="github" subtitle="github.com/HighgoSoftware/wal2mongo" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="wal2mongo-1.0.7.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg wal2mongo;		# build deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install wal2mongo;		# install via package name, for the active PG version

pig install wal2mongo -v 16;   # install for PG 16
pig install wal2mongo -v 15;   # install for PG 15
pig install wal2mongo -v 14;   # install for PG 14

```


This extension does not need `CREATE EXTENSION` DDL command





## Usage

> [wal2mongo: PostgreSQL logical decoding output plugin for MongoDB](https://github.com/HighgoSoftware/wal2mongo)

A logical decoding output plugin that formats PostgreSQL WAL changes as MongoDB-compatible commands, enabling data replication from PostgreSQL to MongoDB.

### Configuration

In `postgresql.conf`:

```ini
wal_level = logical
max_replication_slots = 10
```

### Using with SQL Functions

```sql
-- Create a logical replication slot
SELECT * FROM pg_create_logical_replication_slot('w2m_slot', 'wal2mongo');

-- Perform DML operations
CREATE TABLE books (id SERIAL PRIMARY KEY, title VARCHAR(100), author VARCHAR(100));
INSERT INTO books (id, title, author) VALUES (123, 'My Book', 'Author');

-- Peek at changes (MongoDB format)
SELECT * FROM pg_logical_slot_peek_changes('w2m_slot', NULL, NULL);
-- Output: db.books.insertOne( { id:123, title:"My Book", author:"Author" } )

-- Consume changes
SELECT * FROM pg_logical_slot_get_changes('w2m_slot', NULL, NULL);

-- Drop the slot
SELECT pg_drop_replication_slot('w2m_slot');
```

### Using with pg_recvlogical

```bash
pg_recvlogical -d postgres --slot w2m_slot --create-slot -P wal2mongo
pg_recvlogical -d postgres --slot w2m_slot --start -f -
```

### Replicating to MongoDB

The output can be applied directly in the MongoDB shell:

```javascript
// Copy the output from pg_logical_slot_get_changes
db.books.insertOne( { id:123, title:"My Book", author:"Author" } )
```

Or save to a `.js` file and import:

```bash
mongo < changes.js
```

### Output Format

- **INSERT**: `db.<table>.insertOne( { <columns> } )`
- **UPDATE**: `db.<table>.updateOne( { <key> }, { $set: { <changes> } } )`
- **DELETE**: `db.<table>.deleteOne( { <key> } )`

Tables need a primary key or replica identity for UPDATE/DELETE operations to be captured.
