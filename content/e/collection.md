---
title: "collection"
linkTitle: "collection"
description: "Associative-array data types for high-performance PL/pgSQL workflows"
weight: 3700
categories: ["TYPE"]
languages: ["C"]
licenses: ["Apache-2.0"]
repos: ["PIGSTY"]
page_width: full
---

[**pgcollection**](https://github.com/aws/pgcollection) : Associative-array data types for high-performance PL/pgSQL workflows


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3700** | {{< badge content="collection" link="https://github.com/aws/pgcollection" >}} | {{< ext "collection" "pgcollection" >}} | `2.1.3` | {{< category "TYPE" >}} | {{< license "Apache-2.0" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pgbson" >}} {{< ext "intarray" >}} {{< ext "hstore" >}} {{< ext "arraymath" >}} {{< ext "pgjq" >}} {{< ext "jsquery" >}} {{< ext "pg_jsonschema" >}} {{< ext "jsonschema" >}} {{< ext "floatvec" >}} {{< ext "pg_projection" >}} {{< ext "istore" >}} |

> [!Note] PGXN name: pgcollection.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.3` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgcollection` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.3` | {{< bg "18" "pgcollection_18" "green" >}} {{< bg "17" "pgcollection_17" "green" >}} {{< bg "16" "pgcollection_16" "green" >}} {{< bg "15" "pgcollection_15" "green" >}} {{< bg "14" "pgcollection_14" "green" >}} | `pgcollection_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `2.1.3` | {{< bg "18" "postgresql-18-collection" "green" >}} {{< bg "17" "postgresql-17-collection" "green" >}} {{< bg "16" "postgresql-16-collection" "green" >}} {{< bg "15" "postgresql-15-collection" "green" >}} {{< bg "14" "postgresql-14-collection" "green" >}} | `postgresql-$v-collection` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "pgcollection_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-18-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-17-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-16-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-15-collection : AVAIL 1" "green" >}} | {{< bg "PIGSTY 2.1.3" "postgresql-14-collection : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_18` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 157.2 KiB | [pgcollection_18-2.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_18-2.1.3-1PGSTY.el8.x86_64.rpm) |
| `pgcollection_18` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.3 KiB | [pgcollection_18-2.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_18-2.1.3-1PGSTY.el8.aarch64.rpm) |
| `pgcollection_18` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 159.5 KiB | [pgcollection_18-2.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_18-2.1.3-1PGSTY.el9.x86_64.rpm) |
| `pgcollection_18` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 158.9 KiB | [pgcollection_18-2.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_18-2.1.3-1PGSTY.el9.aarch64.rpm) |
| `pgcollection_18` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 161.4 KiB | [pgcollection_18-2.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_18-2.1.3-1PGSTY.el10.x86_64.rpm) |
| `pgcollection_18` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 159.1 KiB | [pgcollection_18-2.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_18-2.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-collection` | `2.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.9 KiB | [postgresql-18-collection_2.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-collection` | `2.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.4 KiB | [postgresql-18-collection_2.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-collection` | `2.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 144.6 KiB | [postgresql-18-collection_2.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-collection` | `2.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.9 KiB | [postgresql-18-collection_2.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-collection` | `2.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 159.3 KiB | [postgresql-18-collection_2.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-collection` | `2.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 158.1 KiB | [postgresql-18-collection_2.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-collection` | `2.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 152.3 KiB | [postgresql-18-collection_2.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-18-collection` | `2.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 150.1 KiB | [postgresql-18-collection_2.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-18-collection` | `2.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 149.8 KiB | [postgresql-18-collection_2.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-collection` | `2.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 147.5 KiB | [postgresql-18-collection_2.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-18-collection_2.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_17` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 157.0 KiB | [pgcollection_17-2.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_17-2.1.3-1PGSTY.el8.x86_64.rpm) |
| `pgcollection_17` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 155.1 KiB | [pgcollection_17-2.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_17-2.1.3-1PGSTY.el8.aarch64.rpm) |
| `pgcollection_17` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 159.2 KiB | [pgcollection_17-2.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_17-2.1.3-1PGSTY.el9.x86_64.rpm) |
| `pgcollection_17` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 158.6 KiB | [pgcollection_17-2.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_17-2.1.3-1PGSTY.el9.aarch64.rpm) |
| `pgcollection_17` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 161.3 KiB | [pgcollection_17-2.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_17-2.1.3-1PGSTY.el10.x86_64.rpm) |
| `pgcollection_17` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 158.8 KiB | [pgcollection_17-2.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_17-2.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-collection` | `2.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.9 KiB | [postgresql-17-collection_2.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-collection` | `2.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.3 KiB | [postgresql-17-collection_2.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-collection` | `2.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 144.2 KiB | [postgresql-17-collection_2.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-collection` | `2.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.5 KiB | [postgresql-17-collection_2.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-collection` | `2.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 175.6 KiB | [postgresql-17-collection_2.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-collection` | `2.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 174.2 KiB | [postgresql-17-collection_2.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-collection` | `2.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 152.3 KiB | [postgresql-17-collection_2.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-17-collection` | `2.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 149.7 KiB | [postgresql-17-collection_2.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-17-collection` | `2.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 149.6 KiB | [postgresql-17-collection_2.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-collection` | `2.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 147.5 KiB | [postgresql-17-collection_2.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-17-collection_2.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_16` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 156.0 KiB | [pgcollection_16-2.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_16-2.1.3-1PGSTY.el8.x86_64.rpm) |
| `pgcollection_16` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 153.8 KiB | [pgcollection_16-2.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_16-2.1.3-1PGSTY.el8.aarch64.rpm) |
| `pgcollection_16` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 157.9 KiB | [pgcollection_16-2.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_16-2.1.3-1PGSTY.el9.x86_64.rpm) |
| `pgcollection_16` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 157.4 KiB | [pgcollection_16-2.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_16-2.1.3-1PGSTY.el9.aarch64.rpm) |
| `pgcollection_16` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 160.3 KiB | [pgcollection_16-2.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_16-2.1.3-1PGSTY.el10.x86_64.rpm) |
| `pgcollection_16` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 157.5 KiB | [pgcollection_16-2.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_16-2.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-collection` | `2.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 142.5 KiB | [postgresql-16-collection_2.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-collection` | `2.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 139.0 KiB | [postgresql-16-collection_2.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-collection` | `2.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 142.8 KiB | [postgresql-16-collection_2.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-collection` | `2.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 139.3 KiB | [postgresql-16-collection_2.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-collection` | `2.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 173.6 KiB | [postgresql-16-collection_2.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-collection` | `2.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 172.2 KiB | [postgresql-16-collection_2.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-collection` | `2.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 150.8 KiB | [postgresql-16-collection_2.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-16-collection` | `2.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 153.0 KiB | [postgresql-16-collection_2.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-16-collection` | `2.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 148.5 KiB | [postgresql-16-collection_2.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-collection` | `2.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 146.0 KiB | [postgresql-16-collection_2.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-16-collection_2.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_15` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 156.6 KiB | [pgcollection_15-2.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_15-2.1.3-1PGSTY.el8.x86_64.rpm) |
| `pgcollection_15` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 154.4 KiB | [pgcollection_15-2.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_15-2.1.3-1PGSTY.el8.aarch64.rpm) |
| `pgcollection_15` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 158.6 KiB | [pgcollection_15-2.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_15-2.1.3-1PGSTY.el9.x86_64.rpm) |
| `pgcollection_15` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 157.8 KiB | [pgcollection_15-2.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_15-2.1.3-1PGSTY.el9.aarch64.rpm) |
| `pgcollection_15` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 160.5 KiB | [pgcollection_15-2.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_15-2.1.3-1PGSTY.el10.x86_64.rpm) |
| `pgcollection_15` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 157.9 KiB | [pgcollection_15-2.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_15-2.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-collection` | `2.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.4 KiB | [postgresql-15-collection_2.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-collection` | `2.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 162.1 KiB | [postgresql-15-collection_2.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-collection` | `2.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.8 KiB | [postgresql-15-collection_2.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-collection` | `2.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.1 KiB | [postgresql-15-collection_2.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-collection` | `2.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 174.1 KiB | [postgresql-15-collection_2.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-collection` | `2.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 173.0 KiB | [postgresql-15-collection_2.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-collection` | `2.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 151.0 KiB | [postgresql-15-collection_2.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-15-collection` | `2.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 148.9 KiB | [postgresql-15-collection_2.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-15-collection` | `2.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 148.7 KiB | [postgresql-15-collection_2.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-collection` | `2.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 146.3 KiB | [postgresql-15-collection_2.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-15-collection_2.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgcollection_14` | `2.1.3` | [el8.x86_64](/os/el8.x86_64) | pigsty | 156.4 KiB | [pgcollection_14-2.1.3-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgcollection_14-2.1.3-1PGSTY.el8.x86_64.rpm) |
| `pgcollection_14` | `2.1.3` | [el8.aarch64](/os/el8.aarch64) | pigsty | 154.2 KiB | [pgcollection_14-2.1.3-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgcollection_14-2.1.3-1PGSTY.el8.aarch64.rpm) |
| `pgcollection_14` | `2.1.3` | [el9.x86_64](/os/el9.x86_64) | pigsty | 158.4 KiB | [pgcollection_14-2.1.3-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgcollection_14-2.1.3-1PGSTY.el9.x86_64.rpm) |
| `pgcollection_14` | `2.1.3` | [el9.aarch64](/os/el9.aarch64) | pigsty | 157.7 KiB | [pgcollection_14-2.1.3-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgcollection_14-2.1.3-1PGSTY.el9.aarch64.rpm) |
| `pgcollection_14` | `2.1.3` | [el10.x86_64](/os/el10.x86_64) | pigsty | 160.4 KiB | [pgcollection_14-2.1.3-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgcollection_14-2.1.3-1PGSTY.el10.x86_64.rpm) |
| `pgcollection_14` | `2.1.3` | [el10.aarch64](/os/el10.aarch64) | pigsty | 157.8 KiB | [pgcollection_14-2.1.3-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgcollection_14-2.1.3-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-collection` | `2.1.3` | [d12.x86_64](/os/d12.x86_64) | pigsty | 143.4 KiB | [postgresql-14-collection_2.1.3-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-collection` | `2.1.3` | [d12.aarch64](/os/d12.aarch64) | pigsty | 140.0 KiB | [postgresql-14-collection_2.1.3-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-collection` | `2.1.3` | [d13.x86_64](/os/d13.x86_64) | pigsty | 143.9 KiB | [postgresql-14-collection_2.1.3-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-collection` | `2.1.3` | [d13.aarch64](/os/d13.aarch64) | pigsty | 140.2 KiB | [postgresql-14-collection_2.1.3-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-collection` | `2.1.3` | [u22.x86_64](/os/u22.x86_64) | pigsty | 173.7 KiB | [postgresql-14-collection_2.1.3-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-collection` | `2.1.3` | [u22.aarch64](/os/u22.aarch64) | pigsty | 172.6 KiB | [postgresql-14-collection_2.1.3-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-collection` | `2.1.3` | [u24.x86_64](/os/u24.x86_64) | pigsty | 151.0 KiB | [postgresql-14-collection_2.1.3-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~noble_amd64.deb) |
| `postgresql-14-collection` | `2.1.3` | [u24.aarch64](/os/u24.aarch64) | pigsty | 148.8 KiB | [postgresql-14-collection_2.1.3-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~noble_arm64.deb) |
| `postgresql-14-collection` | `2.1.3` | [u26.x86_64](/os/u26.x86_64) | pigsty | 148.7 KiB | [postgresql-14-collection_2.1.3-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-collection` | `2.1.3` | [u26.aarch64](/os/u26.aarch64) | pigsty | 146.3 KiB | [postgresql-14-collection_2.1.3-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/c/collection/postgresql-14-collection_2.1.3-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/aws/pgcollection" title="Repository" icon="github" subtitle="github.com/aws/pgcollection" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgcollection-2.1.3.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg pgcollection;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install pgcollection;		# install via package name, for the active PG version
pig install collection;		# install by extension name, for the current active PG version

pig install collection -v 18;   # install for PG 18
pig install collection -v 17;   # install for PG 17
pig install collection -v 16;   # install for PG 16
pig install collection -v 15;   # install for PG 15
pig install collection -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION collection;
```




## Usage

> [collection: key-value collection data types for PL/pgSQL](https://github.com/aws/pgcollection)

The `collection` extension provides two memory-optimized collection data types for use within PL/pgSQL functions.

```sql
CREATE EXTENSION collection;
```

### Data Types

- **`collection`**: Key-value pairs with text keys (max 32,767 chars), stored in creation order
- **`icollection`**: Key-value pairs with 64-bit integer keys, enabling sparse arrays

Both types support type modifiers to specify element types:

```sql
DECLARE
    c1  collection('date');
    ic1 icollection('int4');
```

### Subscript Access

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['Japan'] := 'Tokyo';
    RAISE NOTICE '%', t_capital['USA'];  -- Washington, D.C.
END $$;
```

### Core Functions

| Function | Description |
|----------|-------------|
| `add(coll, key, value)` | Add element |
| `count(coll)` | Element count |
| `delete(coll, key)` | Remove element |
| `exist(coll, key)` | Check key existence |
| `find(coll, key)` | Retrieve value |
| `first(coll)` | Move iterator to start |
| `last(coll)` | Move iterator to end |
| `next(coll)` | Advance iterator |
| `prev(coll)` | Reverse iterator |
| `key(coll)` | Current key |
| `value(coll)` | Current value |
| `copy(coll)` | Create copy |
| `sort(coll)` | Sort by keys |
| `keys_to_table(coll)` | All keys as set |
| `values_to_table(coll)` | All values as set |
| `to_table(coll)` | Keys and values as table |
{.wide .expand}

### Iterator Example

```sql
DO $$
DECLARE t_capital collection;
BEGIN
    t_capital['USA'] := 'Washington, D.C.';
    t_capital['United Kingdom'] := 'London';
    t_capital['Japan'] := 'Tokyo';

    t_capital := first(t_capital);
    WHILE NOT isnull(t_capital) LOOP
        RAISE NOTICE 'Capital of % is %', key(t_capital), value(t_capital);
        t_capital := next(t_capital);
    END LOOP;
END $$;
```

### Sparse Arrays (icollection)

`icollection` supports non-contiguous integer indices and distinguishes between NULL values and uninitialized keys:

```sql
DECLARE sparse icollection;
BEGIN
    sparse[1] := 'first';
    sparse[1000000] := 'millionth';  -- no memory wasted on gaps
END;
```
