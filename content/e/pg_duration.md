---
title: "pg_duration"
linkTitle: "pg_duration"
description: "data type for representing durations"
weight: 3830
categories: ["TYPE"]
width: full
---

data type for representing durations


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3830** | {{< badge content="pg_duration" link="https://github.com/jkosh44/pg_duration" >}} | {{< ext "pg_duration" >}} | `1.0.2` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pg_duration" >}} | `1.0.2` | {{< bg "18" "pg_duration_18*" "red" >}} {{< bg "17" "pg_duration_17*" "green" >}} {{< bg "16" "pg_duration_16*" "red" >}} {{< bg "15" "pg_duration_15*" "red" >}} {{< bg "14" "pg_duration_14*" "red" >}} | `pg_duration_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pg_duration" >}} | `1.0.2` | {{< bg "18" "postgresql-18-pg-duration" "red" >}} {{< bg "17" "postgresql-17-pg-duration" "green" >}} {{< bg "16" "postgresql-16-pg-duration" "red" >}} {{< bg "15" "postgresql-15-pg-duration" "red" >}} {{< bg "14" "postgresql-14-pg-duration" "red" >}} | `postgresql-$v-pg-duration` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |
|    `el8.aarch64`    | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |
|    `el9.x86_64`    | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 1" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |
|    `el9.aarch64`    | {{< bg "PIGSTY 1.0.2" "pg_duration_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.2" "pg_duration_17 : AVAIL 2" "green" >}} |      {{< bg "MISS" "pg_duration_16 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_15 : MISS 0" "red" >}}      |      {{< bg "MISS" "pg_duration_14 : MISS 0" "red" >}}      |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duration : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duration : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duration : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duration : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pg-duration : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pg-duration : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.2" "postgresql-17-pg-duration : AVAIL 1" "green" >}} |      {{< bg "MISS" "postgresql-16-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-duration : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-duration : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duration_18` | 1.0.2 | `el8.x86_64` | pigsty | 22.7 KiB | [pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_18-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_duration_18` | 1.0.2 | `el8.aarch64` | pigsty | 22.0 KiB | [pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_18-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_duration_18` | 1.0.2 | `el9.x86_64` | pigsty | 23.0 KiB | [pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_18-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_duration_18` | 1.0.2 | `el9.aarch64` | pigsty | 22.2 KiB | [pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_18-1.0.2-1PIGSTY.el9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_duration_17` | 1.0.2 | `el8.x86_64` | pigsty | 22.7 KiB | [pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_duration_17-1.0.2-1PIGSTY.el8.x86_64.rpm) |
| `pg_duration_17` | 1.0.2 | `el8.aarch64` | pigsty | 22.0 KiB | [pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_duration_17-1.0.2-1PIGSTY.el8.aarch64.rpm) |
| `pg_duration_17` | 1.0.2 | `el9.x86_64` | pigsty | 23.1 KiB | [pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_duration_17-1.0.2-1PIGSTY.el9.x86_64.rpm) |
| `pg_duration_17` | 1.0.2 | `el9.aarch64` | pigsty | 22.2 KiB | [pg_duration_17-1.0.2-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.2-1PIGSTY.el9.aarch64.rpm) |
| `pg_duration_17` | 1.0.1 | `el9.aarch64` | pigsty | 22.3 KiB | [pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_duration_17-1.0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-17-pg-duration` | 1.0.2 | `d12.x86_64` | pigsty | 29.2 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `d12.aarch64` | pigsty | 28.8 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u22.x86_64` | pigsty | 32.3 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u22.aarch64` | pigsty | 31.4 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u24.x86_64` | pigsty | 30.5 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-duration` | 1.0.2 | `u24.aarch64` | pigsty | 30.4 KiB | [postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-duration/postgresql-17-pg-duration_1.0.2-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/jkosh44/pg_duration" title="Repository" icon="github" subtitle="github.com/jkosh44/pg_duration" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_duration-1.0.2.tar.gz" >}}
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

