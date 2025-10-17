---
title: "pg_duration"
linkTitle: "pg_duration"
description: "data type for representing durations"
weight: 3830
categories: ["Type"]
width: full
---

data type for representing durations

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3830** | {{< badge content="pg_duration" link="https://github.com/jkosh44/pg_duration" >}} | {{< ext "pg_duration" "pg_duration" >}} | `1.0.2` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="---s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_duration" >}} | `1.0.2` | {{< badge content="18" color="red" alt="pg_duration_18*" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="red" alt="pg_duration_16*" >}} {{< badge content="15" color="red" alt="pg_duration_15*" >}} {{< badge content="14" color="red" alt="pg_duration_14*" >}} | `pg_duration_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_duration" >}} | `1.0.2` | {{< badge content="18" color="red" alt="postgresql-18-pg-duration" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="red" alt="postgresql-16-pg-duration" >}} {{< badge content="15" color="red" alt="postgresql-15-pg-duration" >}} {{< badge content="14" color="red" alt="postgresql-14-pg-duration" >}} | `postgresql-$v-pg-duration` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< pkg "pg_duration_18" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pg_duration_17" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm" >}} |    {{< pkg "pg_duration_16" >}}     |    {{< pkg "pg_duration_15" >}}     |    {{< pkg "pg_duration_14" >}}     |
|    `el8.aarch64`    | {{< pkg "pg_duration_18" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pg_duration_17" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm" >}} |    {{< pkg "pg_duration_16" >}}     |    {{< pkg "pg_duration_15" >}}     |    {{< pkg "pg_duration_14" >}}     |
|    `el9.x86_64`    | {{< pkg "pg_duration_18" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pg_duration_17" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm" >}} |    {{< pkg "pg_duration_16" >}}     |    {{< pkg "pg_duration_15" >}}     |    {{< pkg "pg_duration_14" >}}     |
|    `el9.aarch64`    | {{< pkg "pg_duration_18" "1.0.2" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pg_duration_17" "1.0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm" >}} |    {{< pkg "pg_duration_16" >}}     |    {{< pkg "pg_duration_15" >}}     |    {{< pkg "pg_duration_14" >}}     |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pg-duration" >}}     | {{< pkg "postgresql-17-pg-duration" "1.0.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb" >}} |    {{< pkg "postgresql-16-pg-duration" >}}     |    {{< pkg "postgresql-15-pg-duration" >}}     |    {{< pkg "postgresql-14-pg-duration" >}}     |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pg-duration" >}}     | {{< pkg "postgresql-17-pg-duration" "1.0.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb" >}} |    {{< pkg "postgresql-16-pg-duration" >}}     |    {{< pkg "postgresql-15-pg-duration" >}}     |    {{< pkg "postgresql-14-pg-duration" >}}     |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pg-duration" >}}     | {{< pkg "postgresql-17-pg-duration" "1.0.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb" >}} |    {{< pkg "postgresql-16-pg-duration" >}}     |    {{< pkg "postgresql-15-pg-duration" >}}     |    {{< pkg "postgresql-14-pg-duration" >}}     |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pg-duration" >}}     | {{< pkg "postgresql-17-pg-duration" "1.0.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb" >}} |    {{< pkg "postgresql-16-pg-duration" >}}     |    {{< pkg "postgresql-15-pg-duration" >}}     |    {{< pkg "postgresql-14-pg-duration" >}}     |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pg-duration" >}}     | {{< pkg "postgresql-17-pg-duration" "1.0.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb" >}} |    {{< pkg "postgresql-16-pg-duration" >}}     |    {{< pkg "postgresql-15-pg-duration" >}}     |    {{< pkg "postgresql-14-pg-duration" >}}     |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pg-duration" >}}     | {{< pkg "postgresql-17-pg-duration" "1.0.2" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb" >}} |    {{< pkg "postgresql-16-pg-duration" >}}     |    {{< pkg "postgresql-15-pg-duration" >}}     |    {{< pkg "postgresql-14-pg-duration" >}}     |


{{< tabs items="PG18,PG17" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_duration_18` | 1.0.2 | `el8.x86_64` | pigsty | 22.7 KiB | [pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_duration_18` | 1.0.2 | `el8.aarch64` | pigsty | 22.0 KiB | [pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_duration_18` | 1.0.2 | `el9.aarch64` | pigsty | 22.2 KiB | [pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_duration_18` | 1.0.2 | `el9.x86_64` | pigsty | 23.0 KiB | [pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pg_duration_17` | 1.0.2 | `el8.x86_64` | pigsty | 22.7 KiB | [pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_duration_17` | 1.0.2 | `el8.aarch64` | pigsty | 22.0 KiB | [pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_duration_17` | 1.0.2 | `el9.aarch64` | pigsty | 22.2 KiB | [pg_duration_17-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_duration_17` | 1.0.2 | `el9.x86_64` | pigsty | 23.1 KiB | [pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_duration_17` | 1.0.1 | `el9.aarch64` | pigsty | 22.3 KiB | [pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-duration` | 1.0.2 | `d12.aarch64` | pigsty | 28.8 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `d12.x86_64` | pigsty | 29.2 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u22.aarch64` | pigsty | 31.4 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u22.x86_64` | pigsty | 32.3 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u24.aarch64` | pigsty | 30.4 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u24.x86_64` | pigsty | 30.5 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jkosh44/pg_duration" title="Repository" icon="github" subtitle="github.com/jkosh44/pg_duration" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pg_duration-1.0.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pg_duration; # get pg_duration source code
pig build dep pg_duration; # install build dependencies
pig build pkg pg_duration; # build extension rpm or deb
pig build ext pg_duration; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pg_duration; # install by extension name, for the current active PG version
pig ext install pg_duration; # install via package alias, for the active PG version
pig ext install pg_duration -v 18;   # install for PG 18
pig ext install pg_duration -v 17;   # install for PG 17

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pg_duration;
```

