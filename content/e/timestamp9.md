---
title: "timestamp9"
linkTitle: "timestamp9"
description: "timestamp nanosecond resolution"
weight: 3890
categories: ["TYPE"]
width: full
---

timestamp nanosecond resolution


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **3890** | {{< badge content="timestamp9" link="https://github.com/optiver/timestamp9" >}} | {{< ext "timestamp9" >}} | `1.4.0` | {{< category "TYPE" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="red" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "prefix" >}} {{< ext "semver" >}} {{< ext "unit" >}} {{< ext "pgpdf" >}} {{< ext "pglite_fusion" >}} {{< ext "md5hash" >}} {{< ext "asn1oid" >}} {{< ext "roaringbitmap" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/timestamp9" >}} | `1.4.0` | {{< bg "18" "timestamp9_18*" "green" >}} {{< bg "17" "timestamp9_17*" "green" >}} {{< bg "16" "timestamp9_16*" "green" >}} {{< bg "15" "timestamp9_15*" "green" >}} {{< bg "14" "timestamp9_14*" "green" >}} | `timestamp9_$v*` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/timestamp9" >}} | `1.4.0` | {{< bg "18" "postgresql-18-timestamp9" "green" >}} {{< bg "17" "postgresql-17-timestamp9" "green" >}} {{< bg "16" "postgresql-16-timestamp9" "green" >}} {{< bg "15" "postgresql-15-timestamp9" "green" >}} {{< bg "14" "postgresql-14-timestamp9" "green" >}} | `postgresql-$v-timestamp9` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.4.0" "timestamp9_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_14 : AVAIL 2" "blue" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.4.0" "timestamp9_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_14 : AVAIL 2" "blue" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.4.0" "timestamp9_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_14 : AVAIL 2" "blue" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.4.0" "timestamp9_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.4.0" "timestamp9_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.3.0" "timestamp9_14 : AVAIL 2" "blue" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-timestamp9 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-timestamp9 : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-timestamp9 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-timestamp9 : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-timestamp9 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-timestamp9 : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-timestamp9 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-timestamp9 : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-timestamp9 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-timestamp9 : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-timestamp9 : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.4.0" "postgresql-17-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-16-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-15-timestamp9 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.4.0" "postgresql-14-timestamp9 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timestamp9_18` | 1.4.0 | `el8.x86_64` | pgdg | 17.6 KiB | [timestamp9_18-1.4.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/timestamp9_18-1.4.0-3PGDG.rhel8.x86_64.rpm) |
| `timestamp9_18` | 1.4.0 | `el8.aarch64` | pgdg | 17.4 KiB | [timestamp9_18-1.4.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/timestamp9_18-1.4.0-3PGDG.rhel8.aarch64.rpm) |
| `timestamp9_18` | 1.4.0 | `el9.x86_64` | pgdg | 17.6 KiB | [timestamp9_18-1.4.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/timestamp9_18-1.4.0-3PGDG.rhel9.x86_64.rpm) |
| `timestamp9_18` | 1.4.0 | `el9.aarch64` | pgdg | 17.3 KiB | [timestamp9_18-1.4.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/timestamp9_18-1.4.0-3PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timestamp9_17` | 1.4.0 | `el8.x86_64` | pgdg | 17.5 KiB | [timestamp9_17-1.4.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/timestamp9_17-1.4.0-2PGDG.rhel8.x86_64.rpm) |
| `timestamp9_17` | 1.4.0 | `el8.aarch64` | pgdg | 17.3 KiB | [timestamp9_17-1.4.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/timestamp9_17-1.4.0-2PGDG.rhel8.aarch64.rpm) |
| `timestamp9_17` | 1.4.0 | `el9.x86_64` | pgdg | 17.7 KiB | [timestamp9_17-1.4.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/timestamp9_17-1.4.0-2PGDG.rhel9.x86_64.rpm) |
| `timestamp9_17` | 1.4.0 | `el9.aarch64` | pgdg | 17.6 KiB | [timestamp9_17-1.4.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/timestamp9_17-1.4.0-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-timestamp9` | 1.4.0 | `d12.x86_64` | pigsty | 8.4 KiB | [postgresql-17-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-17-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-timestamp9` | 1.4.0 | `d12.aarch64` | pigsty | 8.5 KiB | [postgresql-17-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-17-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-timestamp9` | 1.4.0 | `u22.x86_64` | pigsty | 9.1 KiB | [postgresql-17-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-17-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-timestamp9` | 1.4.0 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-17-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-17-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-timestamp9` | 1.4.0 | `u24.x86_64` | pigsty | 9.1 KiB | [postgresql-17-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-17-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-timestamp9` | 1.4.0 | `u24.aarch64` | pigsty | 9.2 KiB | [postgresql-17-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-17-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timestamp9_16` | 1.4.0 | `el8.x86_64` | pgdg | 17.5 KiB | [timestamp9_16-1.4.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/timestamp9_16-1.4.0-2PGDG.rhel8.x86_64.rpm) |
| `timestamp9_16` | 1.4.0 | `el8.aarch64` | pgdg | 17.3 KiB | [timestamp9_16-1.4.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/timestamp9_16-1.4.0-2PGDG.rhel8.aarch64.rpm) |
| `timestamp9_16` | 1.4.0 | `el9.x86_64` | pgdg | 17.6 KiB | [timestamp9_16-1.4.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/timestamp9_16-1.4.0-2PGDG.rhel9.x86_64.rpm) |
| `timestamp9_16` | 1.4.0 | `el9.aarch64` | pgdg | 17.2 KiB | [timestamp9_16-1.4.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/timestamp9_16-1.4.0-2PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-timestamp9` | 1.4.0 | `d12.x86_64` | pigsty | 8.4 KiB | [postgresql-16-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-16-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-timestamp9` | 1.4.0 | `d12.aarch64` | pigsty | 8.6 KiB | [postgresql-16-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-16-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-timestamp9` | 1.4.0 | `u22.x86_64` | pigsty | 9.1 KiB | [postgresql-16-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-16-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-timestamp9` | 1.4.0 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-16-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-16-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-timestamp9` | 1.4.0 | `u24.x86_64` | pigsty | 9.1 KiB | [postgresql-16-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-16-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-timestamp9` | 1.4.0 | `u24.aarch64` | pigsty | 9.2 KiB | [postgresql-16-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-16-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timestamp9_15` | 1.3.0 | `el8.x86_64` | pgdg | 17.3 KiB | [timestamp9_15-1.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/timestamp9_15-1.3.0-1.rhel8.x86_64.rpm) |
| `timestamp9_15` | 1.1.0 | `el8.x86_64` | pgdg | 16.4 KiB | [timestamp9_15-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/timestamp9_15-1.1.0-1.rhel8.x86_64.rpm) |
| `timestamp9_15` | 1.3.0 | `el8.aarch64` | pgdg | 17.1 KiB | [timestamp9_15-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/timestamp9_15-1.3.0-1.rhel8.aarch64.rpm) |
| `timestamp9_15` | 1.1.0 | `el8.aarch64` | pgdg | 16.3 KiB | [timestamp9_15-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/timestamp9_15-1.1.0-1.rhel8.aarch64.rpm) |
| `timestamp9_15` | 1.3.0 | `el9.x86_64` | pgdg | 17.4 KiB | [timestamp9_15-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/timestamp9_15-1.3.0-1.rhel9.x86_64.rpm) |
| `timestamp9_15` | 1.1.0 | `el9.x86_64` | pgdg | 16.8 KiB | [timestamp9_15-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/timestamp9_15-1.1.0-1.rhel9.x86_64.rpm) |
| `timestamp9_15` | 1.3.0 | `el9.aarch64` | pgdg | 17.0 KiB | [timestamp9_15-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/timestamp9_15-1.3.0-1.rhel9.aarch64.rpm) |
| `timestamp9_15` | 1.1.0 | `el9.aarch64` | pgdg | 16.6 KiB | [timestamp9_15-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/timestamp9_15-1.1.0-1.rhel9.aarch64.rpm) |
| `postgresql-15-timestamp9` | 1.4.0 | `d12.x86_64` | pigsty | 8.5 KiB | [postgresql-15-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-15-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-timestamp9` | 1.4.0 | `d12.aarch64` | pigsty | 8.6 KiB | [postgresql-15-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-15-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-timestamp9` | 1.4.0 | `u22.x86_64` | pigsty | 9.2 KiB | [postgresql-15-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-15-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-timestamp9` | 1.4.0 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-15-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-15-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-timestamp9` | 1.4.0 | `u24.x86_64` | pigsty | 9.2 KiB | [postgresql-15-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-15-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-timestamp9` | 1.4.0 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-15-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-15-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `timestamp9_14` | 1.3.0 | `el8.x86_64` | pgdg | 17.3 KiB | [timestamp9_14-1.3.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/timestamp9_14-1.3.0-1.rhel8.x86_64.rpm) |
| `timestamp9_14` | 1.1.0 | `el8.x86_64` | pgdg | 16.4 KiB | [timestamp9_14-1.1.0-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/timestamp9_14-1.1.0-1.rhel8.x86_64.rpm) |
| `timestamp9_14` | 1.3.0 | `el8.aarch64` | pgdg | 17.1 KiB | [timestamp9_14-1.3.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/timestamp9_14-1.3.0-1.rhel8.aarch64.rpm) |
| `timestamp9_14` | 1.1.0 | `el8.aarch64` | pgdg | 16.3 KiB | [timestamp9_14-1.1.0-1.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/timestamp9_14-1.1.0-1.rhel8.aarch64.rpm) |
| `timestamp9_14` | 1.3.0 | `el9.x86_64` | pgdg | 17.4 KiB | [timestamp9_14-1.3.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/timestamp9_14-1.3.0-1.rhel9.x86_64.rpm) |
| `timestamp9_14` | 1.1.0 | `el9.x86_64` | pgdg | 16.7 KiB | [timestamp9_14-1.1.0-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/timestamp9_14-1.1.0-1.rhel9.x86_64.rpm) |
| `timestamp9_14` | 1.3.0 | `el9.aarch64` | pgdg | 16.9 KiB | [timestamp9_14-1.3.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/timestamp9_14-1.3.0-1.rhel9.aarch64.rpm) |
| `timestamp9_14` | 1.1.0 | `el9.aarch64` | pgdg | 16.5 KiB | [timestamp9_14-1.1.0-1.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/timestamp9_14-1.1.0-1.rhel9.aarch64.rpm) |
| `postgresql-14-timestamp9` | 1.4.0 | `d12.x86_64` | pigsty | 8.5 KiB | [postgresql-14-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-14-timestamp9_1.4.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-timestamp9` | 1.4.0 | `d12.aarch64` | pigsty | 8.6 KiB | [postgresql-14-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/t/timestamp9/postgresql-14-timestamp9_1.4.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-timestamp9` | 1.4.0 | `u22.x86_64` | pigsty | 9.2 KiB | [postgresql-14-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-14-timestamp9_1.4.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-timestamp9` | 1.4.0 | `u22.aarch64` | pigsty | 9.1 KiB | [postgresql-14-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/t/timestamp9/postgresql-14-timestamp9_1.4.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-timestamp9` | 1.4.0 | `u24.x86_64` | pigsty | 9.2 KiB | [postgresql-14-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-14-timestamp9_1.4.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-timestamp9` | 1.4.0 | `u24.aarch64` | pigsty | 9.3 KiB | [postgresql-14-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/t/timestamp9/postgresql-14-timestamp9_1.4.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/optiver/timestamp9" title="Repository" icon="github" subtitle="github.com/optiver/timestamp9" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="timestamp9-timestamp9-1.4.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build get timestamp9; # get timestamp9 source code
pig build dep timestamp9; # install build dependencies
pig build pkg timestamp9; # build extension rpm or deb
pig build ext timestamp9; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install timestamp9; # install by extension name, for the current active PG version
pig ext install timestamp9; # install via package alias, for the active PG version
pig ext install timestamp9 -v 18;   # install for PG 18
pig ext install timestamp9 -v 17;   # install for PG 17
pig ext install timestamp9 -v 16;   # install for PG 16
pig ext install timestamp9 -v 15;   # install for PG 15
pig ext install timestamp9 -v 14;   # install for PG 14
pig ext install timestamp9 -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION timestamp9;
```

