---
title: "uint"
linkTitle: "uint"
description: "unsigned integer types"
weight: 3730
categories: ["TYPE"]
width: full
---

unsigned integer types


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3730** | {{< badge content="uint" link="https://github.com/petere/pguint" >}} | {{< ext "uint" "pguint" >}} | `1.20250815` | {{< category "TYPE" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |

> [!Note] no pg14 for el8/el9 pgdg repo


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/uint" >}} | `1.20250815` | {{< bg "18" "pguint_18*" "green" >}} {{< bg "17" "pguint_17*" "green" >}} {{< bg "16" "pguint_16*" "green" >}} {{< bg "15" "pguint_15*" "green" >}} {{< bg "14" "pguint_14*" "green" >}} | `pguint_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/uint" >}} | `1.20250815` | {{< bg "18" "postgresql-18-pguint" "red" >}} {{< bg "17" "postgresql-17-pguint" "green" >}} {{< bg "16" "postgresql-16-pguint" "green" >}} {{< bg "15" "postgresql-15-pguint" "green" >}} {{< bg "14" "postgresql-14-pguint" "green" >}} | `postgresql-$v-pguint` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.20250815" "pguint_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_15 : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 1.20231206" "pguint_14 : AVAIL 1" "green" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.20250815" "pguint_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_15 : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 1.20231206" "pguint_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.20250815" "pguint_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_15 : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 1.20231206" "pguint_14 : AVAIL 1" "green" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.20250815" "pguint_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_16 : AVAIL 4" "blue" >}} | {{< bg "PGDG 1.20250815" "pguint_15 : AVAIL 3" "blue" >}} | {{< bg "PIGSTY 1.20231206" "pguint_14 : AVAIL 1" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-pguint : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.20231206" "postgresql-17-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-16-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-15-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-14-pguint : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-pguint : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.20231206" "postgresql-17-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-16-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-15-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-14-pguint : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-pguint : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.20231206" "postgresql-17-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-16-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-15-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-14-pguint : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-pguint : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.20231206" "postgresql-17-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-16-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-15-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-14-pguint : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-pguint : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.20231206" "postgresql-17-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-16-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-15-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-14-pguint : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-pguint : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.20231206" "postgresql-17-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-16-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-15-pguint : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.20231206" "postgresql-14-pguint : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pguint_18` | 1.20250815 | `el8.x86_64` | pgdg | 72.4 KiB | [pguint_18-1.20250815-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pguint_18-1.20250815-1PGDG.rhel8.x86_64.rpm) |
| `pguint_18` | 1.20250815 | `el8.aarch64` | pgdg | 66.5 KiB | [pguint_18-1.20250815-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pguint_18-1.20250815-1PGDG.rhel8.aarch64.rpm) |
| `pguint_18` | 1.20250815 | `el9.x86_64` | pgdg | 75.8 KiB | [pguint_18-1.20250815-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pguint_18-1.20250815-1PGDG.rhel9.x86_64.rpm) |
| `pguint_18` | 1.20250815 | `el9.aarch64` | pgdg | 70.2 KiB | [pguint_18-1.20250815-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pguint_18-1.20250815-1PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pguint_17` | 1.20250815 | `el8.x86_64` | pgdg | 72.4 KiB | [pguint_17-1.20250815-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pguint_17-1.20250815-1PGDG.rhel8.x86_64.rpm) |
| `pguint_17` | 1.20231206 | `el8.x86_64` | pigsty | 68.7 KiB | [pguint_17-1.20231206-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pguint_17-1.20231206-1PIGSTY.el8.x86_64.rpm) |
| `pguint_17` | 1.20231206 | `el8.x86_64` | pgdg | 71.4 KiB | [pguint_17-1.20231206-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pguint_17-1.20231206-2PGDG.rhel8.x86_64.rpm) |
| `pguint_17` | 1.20250815 | `el8.aarch64` | pgdg | 66.5 KiB | [pguint_17-1.20250815-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pguint_17-1.20250815-1PGDG.rhel8.aarch64.rpm) |
| `pguint_17` | 1.20231206 | `el8.aarch64` | pigsty | 62.8 KiB | [pguint_17-1.20231206-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pguint_17-1.20231206-1PIGSTY.el8.aarch64.rpm) |
| `pguint_17` | 1.20231206 | `el8.aarch64` | pgdg | 65.4 KiB | [pguint_17-1.20231206-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pguint_17-1.20231206-2PGDG.rhel8.aarch64.rpm) |
| `pguint_17` | 1.20250815 | `el9.x86_64` | pgdg | 75.7 KiB | [pguint_17-1.20250815-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pguint_17-1.20250815-1PGDG.rhel9.x86_64.rpm) |
| `pguint_17` | 1.20231206 | `el9.x86_64` | pigsty | 72.8 KiB | [pguint_17-1.20231206-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pguint_17-1.20231206-1PIGSTY.el9.x86_64.rpm) |
| `pguint_17` | 1.20231206 | `el9.x86_64` | pgdg | 74.9 KiB | [pguint_17-1.20231206-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pguint_17-1.20231206-2PGDG.rhel9.x86_64.rpm) |
| `pguint_17` | 1.20250815 | `el9.aarch64` | pgdg | 70.3 KiB | [pguint_17-1.20250815-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pguint_17-1.20250815-1PGDG.rhel9.aarch64.rpm) |
| `pguint_17` | 1.20231206 | `el9.aarch64` | pigsty | 67.6 KiB | [pguint_17-1.20231206-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pguint_17-1.20231206-1PIGSTY.el9.aarch64.rpm) |
| `pguint_17` | 1.20231206 | `el9.aarch64` | pgdg | 69.6 KiB | [pguint_17-1.20231206-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pguint_17-1.20231206-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-pguint` | 1.20231206 | `d12.x86_64` | pigsty | 163.8 KiB | [postgresql-17-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-17-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pguint` | 1.20231206 | `d12.aarch64` | pigsty | 159.6 KiB | [postgresql-17-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-17-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pguint` | 1.20231206 | `u22.x86_64` | pigsty | 193.1 KiB | [postgresql-17-pguint_1.20231206-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-17-pguint_1.20231206-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pguint` | 1.20231206 | `u22.aarch64` | pigsty | 188.4 KiB | [postgresql-17-pguint_1.20231206-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-17-pguint_1.20231206-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pguint` | 1.20231206 | `u24.x86_64` | pigsty | 177.9 KiB | [postgresql-17-pguint_1.20231206-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-17-pguint_1.20231206-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pguint` | 1.20231206 | `u24.aarch64` | pigsty | 175.6 KiB | [postgresql-17-pguint_1.20231206-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-17-pguint_1.20231206-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pguint_16` | 1.20250815 | `el8.x86_64` | pgdg | 72.5 KiB | [pguint_16-1.20250815-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pguint_16-1.20250815-1PGDG.rhel8.x86_64.rpm) |
| `pguint_16` | 1.20231206 | `el8.x86_64` | pigsty | 68.7 KiB | [pguint_16-1.20231206-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pguint_16-1.20231206-1PIGSTY.el8.x86_64.rpm) |
| `pguint_16` | 1.20231206 | `el8.x86_64` | pgdg | 71.3 KiB | [pguint_16-1.20231206-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pguint_16-1.20231206-1PGDG.rhel8.x86_64.rpm) |
| `pguint_16` | 1.20220601 | `el8.x86_64` | pgdg | 71.2 KiB | [pguint_16-1.20220601-3.rhel8.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pguint_16-1.20220601-3.rhel8.1.x86_64.rpm) |
| `pguint_16` | 1.20250815 | `el8.aarch64` | pgdg | 66.5 KiB | [pguint_16-1.20250815-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pguint_16-1.20250815-1PGDG.rhel8.aarch64.rpm) |
| `pguint_16` | 1.20231206 | `el8.aarch64` | pigsty | 62.8 KiB | [pguint_16-1.20231206-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pguint_16-1.20231206-1PIGSTY.el8.aarch64.rpm) |
| `pguint_16` | 1.20231206 | `el8.aarch64` | pgdg | 65.3 KiB | [pguint_16-1.20231206-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pguint_16-1.20231206-1PGDG.rhel8.aarch64.rpm) |
| `pguint_16` | 1.20220601 | `el8.aarch64` | pgdg | 64.9 KiB | [pguint_16-1.20220601-3.rhel8.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pguint_16-1.20220601-3.rhel8.1.aarch64.rpm) |
| `pguint_16` | 1.20250815 | `el9.x86_64` | pgdg | 75.7 KiB | [pguint_16-1.20250815-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pguint_16-1.20250815-1PGDG.rhel9.x86_64.rpm) |
| `pguint_16` | 1.20231206 | `el9.x86_64` | pigsty | 72.8 KiB | [pguint_16-1.20231206-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pguint_16-1.20231206-1PIGSTY.el9.x86_64.rpm) |
| `pguint_16` | 1.20231206 | `el9.x86_64` | pgdg | 74.8 KiB | [pguint_16-1.20231206-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pguint_16-1.20231206-1PGDG.rhel9.x86_64.rpm) |
| `pguint_16` | 1.20220601 | `el9.x86_64` | pgdg | 74.5 KiB | [pguint_16-1.20220601-3.rhel9.1.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pguint_16-1.20220601-3.rhel9.1.x86_64.rpm) |
| `pguint_16` | 1.20250815 | `el9.aarch64` | pgdg | 70.4 KiB | [pguint_16-1.20250815-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pguint_16-1.20250815-1PGDG.rhel9.aarch64.rpm) |
| `pguint_16` | 1.20231206 | `el9.aarch64` | pigsty | 67.5 KiB | [pguint_16-1.20231206-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pguint_16-1.20231206-1PIGSTY.el9.aarch64.rpm) |
| `pguint_16` | 1.20231206 | `el9.aarch64` | pgdg | 69.6 KiB | [pguint_16-1.20231206-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pguint_16-1.20231206-1PGDG.rhel9.aarch64.rpm) |
| `pguint_16` | 1.20220601 | `el9.aarch64` | pgdg | 69.1 KiB | [pguint_16-1.20220601-3.rhel9.1.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pguint_16-1.20220601-3.rhel9.1.aarch64.rpm) |
| `postgresql-16-pguint` | 1.20231206 | `d12.x86_64` | pigsty | 164.3 KiB | [postgresql-16-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-16-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pguint` | 1.20231206 | `d12.aarch64` | pigsty | 159.9 KiB | [postgresql-16-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-16-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pguint` | 1.20231206 | `u22.x86_64` | pigsty | 193.1 KiB | [postgresql-16-pguint_1.20231206-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-16-pguint_1.20231206-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pguint` | 1.20231206 | `u22.aarch64` | pigsty | 188.4 KiB | [postgresql-16-pguint_1.20231206-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-16-pguint_1.20231206-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pguint` | 1.20231206 | `u24.x86_64` | pigsty | 178.0 KiB | [postgresql-16-pguint_1.20231206-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-16-pguint_1.20231206-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pguint` | 1.20231206 | `u24.aarch64` | pigsty | 175.6 KiB | [postgresql-16-pguint_1.20231206-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-16-pguint_1.20231206-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pguint_15` | 1.20250815 | `el8.x86_64` | pgdg | 72.4 KiB | [pguint_15-1.20250815-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pguint_15-1.20250815-1PGDG.rhel8.x86_64.rpm) |
| `pguint_15` | 1.20231206 | `el8.x86_64` | pigsty | 68.6 KiB | [pguint_15-1.20231206-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pguint_15-1.20231206-1PIGSTY.el8.x86_64.rpm) |
| `pguint_15` | 1.20231206 | `el8.x86_64` | pgdg | 71.2 KiB | [pguint_15-1.20231206-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pguint_15-1.20231206-1PGDG.rhel8.x86_64.rpm) |
| `pguint_15` | 1.20250815 | `el8.aarch64` | pgdg | 66.4 KiB | [pguint_15-1.20250815-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pguint_15-1.20250815-1PGDG.rhel8.aarch64.rpm) |
| `pguint_15` | 1.20231206 | `el8.aarch64` | pigsty | 62.9 KiB | [pguint_15-1.20231206-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pguint_15-1.20231206-1PIGSTY.el8.aarch64.rpm) |
| `pguint_15` | 1.20231206 | `el8.aarch64` | pgdg | 65.3 KiB | [pguint_15-1.20231206-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pguint_15-1.20231206-1PGDG.rhel8.aarch64.rpm) |
| `pguint_15` | 1.20250815 | `el9.x86_64` | pgdg | 75.6 KiB | [pguint_15-1.20250815-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pguint_15-1.20250815-1PGDG.rhel9.x86_64.rpm) |
| `pguint_15` | 1.20231206 | `el9.x86_64` | pigsty | 72.7 KiB | [pguint_15-1.20231206-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pguint_15-1.20231206-1PIGSTY.el9.x86_64.rpm) |
| `pguint_15` | 1.20231206 | `el9.x86_64` | pgdg | 74.8 KiB | [pguint_15-1.20231206-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pguint_15-1.20231206-1PGDG.rhel9.x86_64.rpm) |
| `pguint_15` | 1.20250815 | `el9.aarch64` | pgdg | 70.4 KiB | [pguint_15-1.20250815-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pguint_15-1.20250815-1PGDG.rhel9.aarch64.rpm) |
| `pguint_15` | 1.20231206 | `el9.aarch64` | pigsty | 67.4 KiB | [pguint_15-1.20231206-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pguint_15-1.20231206-1PIGSTY.el9.aarch64.rpm) |
| `pguint_15` | 1.20231206 | `el9.aarch64` | pgdg | 69.4 KiB | [pguint_15-1.20231206-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pguint_15-1.20231206-1PGDG.rhel9.aarch64.rpm) |
| `postgresql-15-pguint` | 1.20231206 | `d12.x86_64` | pigsty | 162.3 KiB | [postgresql-15-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-15-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pguint` | 1.20231206 | `d12.aarch64` | pigsty | 158.1 KiB | [postgresql-15-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-15-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pguint` | 1.20231206 | `u22.x86_64` | pigsty | 192.3 KiB | [postgresql-15-pguint_1.20231206-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-15-pguint_1.20231206-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pguint` | 1.20231206 | `u22.aarch64` | pigsty | 187.5 KiB | [postgresql-15-pguint_1.20231206-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-15-pguint_1.20231206-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pguint` | 1.20231206 | `u24.x86_64` | pigsty | 177.3 KiB | [postgresql-15-pguint_1.20231206-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-15-pguint_1.20231206-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pguint` | 1.20231206 | `u24.aarch64` | pigsty | 175.0 KiB | [postgresql-15-pguint_1.20231206-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-15-pguint_1.20231206-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pguint_14` | 1.20231206 | `el8.x86_64` | pigsty | 68.8 KiB | [pguint_14-1.20231206-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pguint_14-1.20231206-1PIGSTY.el8.x86_64.rpm) |
| `pguint_14` | 1.20231206 | `el8.aarch64` | pigsty | 62.8 KiB | [pguint_14-1.20231206-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pguint_14-1.20231206-1PIGSTY.el8.aarch64.rpm) |
| `pguint_14` | 1.20231206 | `el9.x86_64` | pigsty | 72.4 KiB | [pguint_14-1.20231206-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pguint_14-1.20231206-1PIGSTY.el9.x86_64.rpm) |
| `pguint_14` | 1.20231206 | `el9.aarch64` | pigsty | 67.5 KiB | [pguint_14-1.20231206-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pguint_14-1.20231206-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pguint` | 1.20231206 | `d12.x86_64` | pigsty | 163.1 KiB | [postgresql-14-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-14-pguint_1.20231206-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pguint` | 1.20231206 | `d12.aarch64` | pigsty | 158.7 KiB | [postgresql-14-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pguint/postgresql-14-pguint_1.20231206-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pguint` | 1.20231206 | `u22.x86_64` | pigsty | 192.3 KiB | [postgresql-14-pguint_1.20231206-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-14-pguint_1.20231206-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pguint` | 1.20231206 | `u22.aarch64` | pigsty | 187.4 KiB | [postgresql-14-pguint_1.20231206-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pguint/postgresql-14-pguint_1.20231206-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pguint` | 1.20231206 | `u24.x86_64` | pigsty | 177.3 KiB | [postgresql-14-pguint_1.20231206-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-14-pguint_1.20231206-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pguint` | 1.20231206 | `u24.aarch64` | pigsty | 175.1 KiB | [postgresql-14-pguint_1.20231206-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pguint/postgresql-14-pguint_1.20231206-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/petere/pguint" title="Repository" icon="github" subtitle="github.com/petere/pguint" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pguint-1.20231206.tar.gz" >}}
{{< /cards >}}


```bash
pig build get uint; # get uint source code
pig build dep uint; # install build dependencies
pig build pkg uint; # build extension rpm or deb
pig build ext uint; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install uint; # install by extension name, for the current active PG version
pig ext install pguint; # install via package alias, for the active PG version
pig ext install uint -v 18;   # install for PG 18
pig ext install uint -v 17;   # install for PG 17
pig ext install uint -v 16;   # install for PG 16
pig ext install uint -v 15;   # install for PG 15
pig ext install uint -v 14;   # install for PG 14
pig ext install uint -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION uint;
```

