---
title: "spat"
linkTitle: "spat"
description: "Redis-like In-Memory DB Embedded in Postgres"
weight: 9400
categories: ["Sim"]
width: full
---

Redis-like In-Memory DB Embedded in Postgres

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9400** | {{< badge content="spat" link="https://github.com/Florents-Tselai/spat" >}} | {{< ext "spat" "spat" >}} | `0.1.0a4` | {{< category "SIM" >}} | {{< license "AGPL-3.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "redis_fdw" >}} {{< ext "redis" >}} {{< ext "pgmemcache" >}} {{< ext "mongo_fdw" >}} {{< ext "kafka_fdw" >}} {{< ext "documentdb" >}} {{< ext "documentdb_core" >}} {{< ext "documentdb_distributed" >}} |

> [!Note] Alpha Stage!


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/spat" >}} | `0.1.0a4` | {{< badge content="18" color="red" alt="spat_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="red" alt="spat_16*" >}} {{< badge content="15" color="red" alt="spat_15*" >}} {{< badge content="14" color="red" alt="spat_14*" >}} | `spat_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/spat" >}} | `0.1.0a4` | {{< badge content="18" color="red" alt="postgresql-18-spat" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="red" alt="postgresql-16-spat" >}} {{< badge content="15" color="red" alt="postgresql-15-spat" >}} {{< badge content="14" color="red" alt="postgresql-14-spat" >}} | `postgresql-$v-spat` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "spat_18" >}}     | {{< pkg "spat_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/spat_17-0.1.0a4-1PIGSTY.el8.x86_64.rpm" >}} |    {{< pkg "spat_16" >}}     |    {{< pkg "spat_15" >}}     |    {{< pkg "spat_14" >}}     |
|    `el8.aarch64`    |    {{< pkg "spat_18" >}}     | {{< pkg "spat_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/spat_17-0.1.0a4-1PIGSTY.el8.aarch64.rpm" >}} |    {{< pkg "spat_16" >}}     |    {{< pkg "spat_15" >}}     |    {{< pkg "spat_14" >}}     |
|    `el9.x86_64`    |    {{< pkg "spat_18" >}}     | {{< pkg "spat_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/spat_17-0.1.0a4-1PIGSTY.el9.x86_64.rpm" >}} |    {{< pkg "spat_16" >}}     |    {{< pkg "spat_15" >}}     |    {{< pkg "spat_14" >}}     |
|    `el9.aarch64`    |    {{< pkg "spat_18" >}}     | {{< pkg "spat_17" "0.1.0" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/spat_17-0.1.0a4-1PIGSTY.el9.aarch64.rpm" >}} |    {{< pkg "spat_16" >}}     |    {{< pkg "spat_15" >}}     |    {{< pkg "spat_14" >}}     |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-spat" >}}     | {{< pkg "postgresql-17-spat" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_amd64.deb" >}} |    {{< pkg "postgresql-16-spat" >}}     |    {{< pkg "postgresql-15-spat" >}}     |    {{< pkg "postgresql-14-spat" >}}     |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-spat" >}}     | {{< pkg "postgresql-17-spat" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_arm64.deb" >}} |    {{< pkg "postgresql-16-spat" >}}     |    {{< pkg "postgresql-15-spat" >}}     |    {{< pkg "postgresql-14-spat" >}}     |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-spat" >}}     | {{< pkg "postgresql-17-spat" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_amd64.deb" >}} |    {{< pkg "postgresql-16-spat" >}}     |    {{< pkg "postgresql-15-spat" >}}     |    {{< pkg "postgresql-14-spat" >}}     |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-spat" >}}     | {{< pkg "postgresql-17-spat" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_arm64.deb" >}} |    {{< pkg "postgresql-16-spat" >}}     |    {{< pkg "postgresql-15-spat" >}}     |    {{< pkg "postgresql-14-spat" >}}     |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-spat" >}}     | {{< pkg "postgresql-17-spat" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~noble_amd64.deb" >}} |    {{< pkg "postgresql-16-spat" >}}     |    {{< pkg "postgresql-15-spat" >}}     |    {{< pkg "postgresql-14-spat" >}}     |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-spat" >}}     | {{< pkg "postgresql-17-spat" "0.1.0" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~noble_arm64.deb" >}} |    {{< pkg "postgresql-16-spat" >}}     |    {{< pkg "postgresql-15-spat" >}}     |    {{< pkg "postgresql-14-spat" >}}     |


{{< tabs items="PG17" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `spat_17` | 0.1.0 | `el8.aarch64` | pigsty | 35.8 KiB | [spat_17-0.1.0a4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/spat_17-0.1.0a4-1PIGSTY.el8.aarch64.rpm) |
| `spat_17` | 0.1.0 | `el8.x86_64` | pigsty | 36.4 KiB | [spat_17-0.1.0a4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/spat_17-0.1.0a4-1PIGSTY.el8.x86_64.rpm) |
| `spat_17` | 0.1.0 | `el9.aarch64` | pigsty | 35.5 KiB | [spat_17-0.1.0a4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/spat_17-0.1.0a4-1PIGSTY.el9.aarch64.rpm) |
| `spat_17` | 0.1.0 | `el9.x86_64` | pigsty | 36.3 KiB | [spat_17-0.1.0a4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/spat_17-0.1.0a4-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-spat` | 0.1.0 | `d12.aarch64` | pigsty | 45.6 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-spat` | 0.1.0 | `d12.x86_64` | pigsty | 46.3 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u22.aarch64` | pigsty | 50.8 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u22.x86_64` | pigsty | 51.3 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u24.aarch64` | pigsty | 47.2 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-spat` | 0.1.0 | `u24.x86_64` | pigsty | 47.7 KiB | [postgresql-17-spat_0.1.0a4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/spat/postgresql-17-spat_0.1.0a4-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/Florents-Tselai/spat" title="Repository" icon="github" subtitle="github.com/Florents-Tselai/spat" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="spat-0.1.0a4.tar.gz" >}}
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

