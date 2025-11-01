---
title: "orioledb"
linkTitle: "orioledb"
description: "OrioleDB, the next generation transactional engine"
weight: 2920
categories: ["FEAT"]
width: full
---

OrioleDB, the next generation transactional engine


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2920** | {{< badge content="orioledb" link="https://github.com/orioledb/orioledb" >}} | {{< ext "orioledb" >}} | `1.5` | {{< category "FEAT" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="red" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "columnar" >}} {{< ext "pg_mooncake" >}} {{< ext "citus_columnar" >}} {{< ext "pg_analytics" >}} {{< ext "pg_duckdb" >}} {{< ext "timescaledb" >}} {{< ext "citus" >}} {{< ext "pg_strom" >}} |

> [!Note] only works on orioledb patched postgres kernel


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/orioledb" >}} | `1.5` | {{< bg "18" "orioledb_18*" "red" >}} {{< bg "17" "orioledb_17*" "green" >}} {{< bg "16" "orioledb_16*" "red" >}} {{< bg "15" "orioledb_15*" "red" >}} {{< bg "14" "orioledb_14*" "red" >}} {{< bg "13" "orioledb_13*" "red" >}} | `orioledb_$v*` | `oriolepg_$v` |
| **Debian** | {{< badge content="PIGSTY" link="/e/orioledb" >}} | `1.5` | {{< bg "18" "oriolepg-18-orioledb" "red" >}} {{< bg "17" "oriolepg-17-orioledb" "green" >}} {{< bg "16" "oriolepg-16-orioledb" "red" >}} {{< bg "15" "oriolepg-15-orioledb" "red" >}} {{< bg "14" "oriolepg-14-orioledb" "red" >}} {{< bg "13" "oriolepg-13-orioledb" "red" >}} | `oriolepg-$v-orioledb` | `oriolepg-$v` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |      {{< bg "MISS" "orioledb_18 : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : THROW 1" "purple" >}}      |      {{< bg "MISS" "orioledb_16 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : THROW 0" "red" >}}      |
|    `el8.aarch64`    |      {{< bg "MISS" "orioledb_18 : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : THROW 1" "purple" >}}      |      {{< bg "MISS" "orioledb_16 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : THROW 0" "red" >}}      |
|    `el9.x86_64`    |      {{< bg "MISS" "orioledb_18 : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : THROW 1" "purple" >}}      |      {{< bg "MISS" "orioledb_16 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : THROW 0" "red" >}}      |
|    `el9.aarch64`    |      {{< bg "MISS" "orioledb_18 : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : THROW 1" "purple" >}}      |      {{< bg "MISS" "orioledb_16 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : THROW 0" "red" >}}      |
|    `el10.x86_64`    |      {{< bg "MISS" "orioledb_18 : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : THROW 1" "purple" >}}      |      {{< bg "MISS" "orioledb_16 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : THROW 0" "red" >}}      |
|    `el10.aarch64`    |      {{< bg "MISS" "orioledb_18 : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "orioledb_17 : THROW 1" "purple" >}}      |      {{< bg "MISS" "orioledb_16 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_15 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_14 : THROW 0" "red" >}}      |      {{< bg "MISS" "orioledb_13 : THROW 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : THROW 1" "purple" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : THROW 1" "purple" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `d13.x86_64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-17-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `d13.aarch64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-17-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : THROW 1" "purple" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : THROW 1" "purple" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : THROW 1" "purple" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "oriolepg-18-orioledb : THROW 0" "red" >}}      |      {{< bg "PIGSTY 1.5" "oriolepg-17-orioledb : THROW 1" "purple" >}}      |      {{< bg "MISS" "oriolepg-16-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-15-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-14-orioledb : THROW 0" "red" >}}      |      {{< bg "MISS" "oriolepg-13-orioledb : THROW 0" "red" >}}      |


{{< tabs items="PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `orioledb_17` | 1.5 | `el8.x86_64` | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/orioledb_17-1.5-0.beta12PIGSTY.el8.x86_64.rpm) |
| `orioledb_17` | 1.5 | `el8.aarch64` | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/orioledb_17-1.5-0.beta12PIGSTY.el8.aarch64.rpm) |
| `orioledb_17` | 1.5 | `el9.x86_64` | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/orioledb_17-1.5-0.beta12PIGSTY.el9.x86_64.rpm) |
| `orioledb_17` | 1.5 | `el9.aarch64` | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/orioledb_17-1.5-0.beta12PIGSTY.el9.aarch64.rpm) |
| `orioledb_17` | 1.5 | `el10.x86_64` | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/orioledb_17-1.5-0.beta12PIGSTY.el10.x86_64.rpm) |
| `orioledb_17` | 1.5 | `el10.aarch64` | pigsty | 1.3 MiB | [orioledb_17-1.5-0.beta12PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/orioledb_17-1.5-0.beta12PIGSTY.el10.aarch64.rpm) |
| `oriolepg-17-orioledb` | 1.5 | `d12.x86_64` | pigsty | 1.5 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_amd64.deb) |
| `oriolepg-17-orioledb` | 1.5 | `d12.aarch64` | pigsty | 1.4 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~bookworm_arm64.deb) |
| `oriolepg-17-orioledb` | 1.5 | `u22.x86_64` | pigsty | 1.6 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_amd64.deb) |
| `oriolepg-17-orioledb` | 1.5 | `u22.aarch64` | pigsty | 1.5 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~jammy_arm64.deb) |
| `oriolepg-17-orioledb` | 1.5 | `u24.x86_64` | pigsty | 1.4 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_amd64.deb) |
| `oriolepg-17-orioledb` | 1.5 | `u24.aarch64` | pigsty | 1.3 MiB | [oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/o/oriolepg-17-orioledb/oriolepg-17-orioledb_1.5-0.beta12PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/orioledb/orioledb" title="Repository" icon="github" subtitle="github.com/orioledb/orioledb" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="orioledb-beta12.tar.gz" >}}
{{< /cards >}}


```bash
pig build get orioledb; # get orioledb source code
pig build dep orioledb; # install build dependencies
pig build pkg orioledb; # build extension rpm or deb
pig build ext orioledb; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install orioledb; # install by extension name, for the current active PG version
pig ext install orioledb; # install via package alias, for the active PG version
pig ext install orioledb -v 17;   # install for PG 17

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION orioledb;
```

