---
title: "spat"
linkTitle: "spat"
description: "Redis-like In-Memory DB Embedded in Postgres"
weight: 9400
categories: ["SIM"]
width: full
---

Redis-like In-Memory DB Embedded in Postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9400** | {{< badge content="spat" link="https://github.com/Florents-Tselai/spat" >}} | {{< ext "spat" >}} | `0.1.0a4` | {{< category "SIM" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "redis_fdw" >}} {{< ext "redis" >}} {{< ext "pgmemcache" >}} {{< ext "mongo_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] Alpha Stage!


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/spat" >}} | `0.1.0a4` | {{< bg "18" "spat_18*" "red" >}} {{< bg "17" "spat_17*" "green" >}} {{< bg "16" "spat_16*" "red" >}} {{< bg "15" "spat_15*" "red" >}} {{< bg "14" "spat_14*" "red" >}} {{< bg "13" "spat_13*" "red" >}} | `spat_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/spat" >}} | `0.1.0a4` | {{< bg "18" "postgresql-18-spat" "red" >}} {{< bg "17" "postgresql-17-spat" "green" >}} {{< bg "16" "postgresql-16-spat" "red" >}} {{< bg "15" "postgresql-15-spat" "red" >}} {{< bg "14" "postgresql-14-spat" "red" >}} {{< bg "13" "postgresql-13-spat" "red" >}} | `postgresql-$v-spat` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} |      {{< bg "MISS" "spat_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "spat_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "spat_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_13 : MISS 0" "red" >}}      |
| {{< os "el8.aarch64" >}} |      {{< bg "MISS" "spat_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "spat_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "spat_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_13 : MISS 0" "red" >}}      |
| {{< os "el9.x86_64" >}} |      {{< bg "MISS" "spat_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "spat_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "spat_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_13 : MISS 0" "red" >}}      |
| {{< os "el9.aarch64" >}} |      {{< bg "MISS" "spat_18 : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "spat_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "spat_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_13 : MISS 0" "red" >}}      |
| {{< os "el10.x86_64" >}} |      {{< bg "MISS" "spat_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_13 : MISS 0" "red" >}}      |
| {{< os "el10.aarch64" >}} |      {{< bg "MISS" "spat_18 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_17 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_14 : MISS 0" "red" >}}      |      {{< bg "MISS" "spat_13 : MISS 0" "red" >}}      |
| {{< os "d12.x86_64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-spat : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "d12.aarch64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-spat : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "d13.x86_64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "d13.aarch64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "u22.x86_64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-spat : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "u22.aarch64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-spat : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "u24.x86_64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-spat : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |
| {{< os "u24.aarch64" >}} |      {{< bg "MISS" "postgresql-18-spat : MISS 0" "red" >}}      | {{< bg "PIGSTY 0.1.0" "postgresql-17-spat : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-spat : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-13-spat : MISS 0" "red" >}}      |


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `spat_17` | 0.1.0 | `el8.x86_64` | pigsty | 36.4 KiB | [spat_17-0.1.0a4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/spat_17-0.1.0a4-1PIGSTY.el8.x86_64.rpm) |
| `spat_17` | 0.1.0 | `el8.aarch64` | pigsty | 35.8 KiB | [spat_17-0.1.0a4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/spat_17-0.1.0a4-1PIGSTY.el8.aarch64.rpm) |
| `spat_17` | 0.1.0 | `el9.x86_64` | pigsty | 36.3 KiB | [spat_17-0.1.0a4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/spat_17-0.1.0a4-1PIGSTY.el9.x86_64.rpm) |
| `spat_17` | 0.1.0 | `el9.aarch64` | pigsty | 35.5 KiB | [spat_17-0.1.0a4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/spat_17-0.1.0a4-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-spat` | 0.1.0 | `d12.x86_64` | pigsty | 46.3 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-spat` | 0.1.0 | `d12.aarch64` | pigsty | 45.6 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u22.x86_64` | pigsty | 51.3 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u22.aarch64` | pigsty | 50.8 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u24.x86_64` | pigsty | 47.7 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u24.aarch64` | pigsty | 47.2 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Florents-Tselai/spat" title="Repository" icon="github" subtitle="github.com/Florents-Tselai/spat" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="spat-0.1.0a4.tar.gz" >}}
{{< /cards >}}


```bash
pig build get spat; # get spat source code
pig build dep spat; # install build dependencies
pig build pkg spat; # build extension rpm or deb
pig build ext spat; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install spat; # install by extension name, for the current active PG version
pig ext install spat; # install via package alias, for the active PG version
pig ext install spat -v 17;   # install for PG 17

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION spat;
```

