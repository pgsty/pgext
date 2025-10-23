---
title: "sequential_uuids"
linkTitle: "sequential_uuids"
description: "generator of sequential UUIDs"
weight: 4570
categories: ["FUNC"]
width: full
---

generator of sequential UUIDs


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4570** | {{< badge content="sequential_uuids" link="https://github.com/tvondra/sequential-uuids" >}} | {{< ext "sequential_uuids" >}} | `1.0.3` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pg_uuidv7" >}} {{< ext "pgx_ulid" >}} {{< ext "uuid-ossp" >}} {{< ext "pg_hashids" >}} {{< ext "permuteseq" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PGDG" link="/e/sequential_uuids" >}} | `1.0.3` | {{< bg "18" "sequential_uuids_18" "red" >}} {{< bg "17" "sequential_uuids_17" "green" >}} {{< bg "16" "sequential_uuids_16" "green" >}} {{< bg "15" "sequential_uuids_15" "green" >}} {{< bg "14" "sequential_uuids_14" "green" >}} | `sequential_uuids_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/sequential_uuids" >}} | `1.0.3` | {{< bg "18" "postgresql-18-sequential-uuids" "red" >}} {{< bg "17" "postgresql-17-sequential-uuids" "green" >}} {{< bg "16" "postgresql-16-sequential-uuids" "green" >}} {{< bg "15" "postgresql-15-sequential-uuids" "green" >}} {{< bg "14" "postgresql-14-sequential-uuids" "green" >}} | `postgresql-$v-sequential-uuids` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_14 : AVAIL 2" "green" >}} |
|    `el8.aarch64`    | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_14 : AVAIL 1" "green" >}} |
|    `el9.x86_64`    | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_14 : AVAIL 2" "green" >}} |
|    `el9.aarch64`    | {{< bg "PGDG 1.0.3" "sequential_uuids_18 : AVAIL 1" "blue" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_17 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_16 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_15 : AVAIL 2" "green" >}} | {{< bg "PIGSTY 1.0.3" "sequential_uuids_14 : AVAIL 2" "green" >}} |
|    `d12.x86_64`    |      {{< bg "MISS" "postgresql-18-sequential-uuids : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-sequential-uuids : AVAIL 1" "green" >}} |
|    `d12.aarch64`    |      {{< bg "MISS" "postgresql-18-sequential-uuids : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-sequential-uuids : AVAIL 1" "green" >}} |
|    `u22.x86_64`    |      {{< bg "MISS" "postgresql-18-sequential-uuids : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-sequential-uuids : AVAIL 1" "green" >}} |
|    `u22.aarch64`    |      {{< bg "MISS" "postgresql-18-sequential-uuids : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-sequential-uuids : AVAIL 1" "green" >}} |
|    `u24.x86_64`    |      {{< bg "MISS" "postgresql-18-sequential-uuids : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-sequential-uuids : AVAIL 1" "green" >}} |
|    `u24.aarch64`    |      {{< bg "MISS" "postgresql-18-sequential-uuids : MISS 0" "red" >}}      | {{< bg "PIGSTY 1.0.3" "postgresql-17-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-16-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-15-sequential-uuids : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.3" "postgresql-14-sequential-uuids : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_18` | 1.0.3 | `el8.x86_64` | pgdg | 16.2 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/sequential_uuids_18-1.0.3-2PGDG.rhel8.x86_64.rpm) |
| `sequential_uuids_18` | 1.0.3 | `el8.aarch64` | pgdg | 16.2 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/sequential_uuids_18-1.0.3-2PGDG.rhel8.aarch64.rpm) |
| `sequential_uuids_18` | 1.0.3 | `el9.x86_64` | pgdg | 16.0 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/sequential_uuids_18-1.0.3-2PGDG.rhel9.x86_64.rpm) |
| `sequential_uuids_18` | 1.0.3 | `el9.aarch64` | pgdg | 15.7 KiB | [sequential_uuids_18-1.0.3-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/sequential_uuids_18-1.0.3-2PGDG.rhel9.aarch64.rpm) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_17` | 1.0.3 | `el8.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_17-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_17-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `sequential_uuids_17` | 1.0.3 | `el8.aarch64` | pigsty | 12.8 KiB | [sequential_uuids_17-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_17-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `sequential_uuids_17` | 1.0.3 | `el9.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_17-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_17-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `sequential_uuids_17` | 1.0.2 | `el9.x86_64` | pgdg | 15.6 KiB | [sequential_uuids_17-1.0.2-5PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/sequential_uuids_17-1.0.2-5PGDG.rhel9.x86_64.rpm) |
| `sequential_uuids_17` | 1.0.3 | `el9.aarch64` | pigsty | 12.7 KiB | [sequential_uuids_17-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_17-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `sequential_uuids_17` | 1.0.2 | `el9.aarch64` | pgdg | 15.4 KiB | [sequential_uuids_17-1.0.2-5PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/sequential_uuids_17-1.0.2-5PGDG.rhel9.aarch64.rpm) |
| `postgresql-17-sequential-uuids` | 1.0.3 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-17-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-sequential-uuids` | 1.0.3 | `d12.aarch64` | pigsty | 13.0 KiB | [postgresql-17-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-sequential-uuids` | 1.0.3 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-17-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-sequential-uuids` | 1.0.3 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-17-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-sequential-uuids` | 1.0.3 | `u24.x86_64` | pigsty | 13.4 KiB | [postgresql-17-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-sequential-uuids` | 1.0.3 | `u24.aarch64` | pigsty | 13.1 KiB | [postgresql-17-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-17-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_16` | 1.0.3 | `el8.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_16-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_16-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `sequential_uuids_16` | 1.0.3 | `el8.aarch64` | pigsty | 12.8 KiB | [sequential_uuids_16-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_16-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `sequential_uuids_16` | 1.0.3 | `el9.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_16-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_16-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `sequential_uuids_16` | 1.0.2 | `el9.x86_64` | pgdg | 15.5 KiB | [sequential_uuids_16-1.0.2-4PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/sequential_uuids_16-1.0.2-4PGDG.rhel9.x86_64.rpm) |
| `sequential_uuids_16` | 1.0.3 | `el9.aarch64` | pigsty | 12.6 KiB | [sequential_uuids_16-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_16-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `sequential_uuids_16` | 1.0.2 | `el9.aarch64` | pgdg | 15.2 KiB | [sequential_uuids_16-1.0.2-4PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/sequential_uuids_16-1.0.2-4PGDG.rhel9.aarch64.rpm) |
| `postgresql-16-sequential-uuids` | 1.0.3 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-16-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-sequential-uuids` | 1.0.3 | `d12.aarch64` | pigsty | 13.0 KiB | [postgresql-16-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-sequential-uuids` | 1.0.3 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-16-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-sequential-uuids` | 1.0.3 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-16-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-sequential-uuids` | 1.0.3 | `u24.x86_64` | pigsty | 13.4 KiB | [postgresql-16-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-sequential-uuids` | 1.0.3 | `u24.aarch64` | pigsty | 13.1 KiB | [postgresql-16-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-16-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_15` | 1.0.3 | `el8.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_15-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_15-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `sequential_uuids_15` | 1.0.3 | `el8.aarch64` | pigsty | 12.8 KiB | [sequential_uuids_15-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_15-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `sequential_uuids_15` | 1.0.3 | `el9.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_15-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_15-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `sequential_uuids_15` | 1.0.2 | `el9.x86_64` | pgdg | 15.4 KiB | [sequential_uuids_15-1.0.2-2.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/sequential_uuids_15-1.0.2-2.rhel9.x86_64.rpm) |
| `sequential_uuids_15` | 1.0.3 | `el9.aarch64` | pigsty | 12.7 KiB | [sequential_uuids_15-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_15-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `sequential_uuids_15` | 1.0.2 | `el9.aarch64` | pgdg | 15.1 KiB | [sequential_uuids_15-1.0.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/sequential_uuids_15-1.0.2-2.rhel9.aarch64.rpm) |
| `postgresql-15-sequential-uuids` | 1.0.3 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-15-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-sequential-uuids` | 1.0.3 | `d12.aarch64` | pigsty | 13.0 KiB | [postgresql-15-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-sequential-uuids` | 1.0.3 | `u22.x86_64` | pigsty | 13.3 KiB | [postgresql-15-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-sequential-uuids` | 1.0.3 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-15-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-sequential-uuids` | 1.0.3 | `u24.x86_64` | pigsty | 13.4 KiB | [postgresql-15-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-sequential-uuids` | 1.0.3 | `u24.aarch64` | pigsty | 13.1 KiB | [postgresql-15-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-15-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `sequential_uuids_14` | 1.0.3 | `el8.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_14-1.0.3-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/sequential_uuids_14-1.0.3-1PIGSTY.el8.x86_64.rpm) |
| `sequential_uuids_14` | 1.0.2 | `el8.x86_64` | pgdg | 21.3 KiB | [sequential_uuids_14-1.0.2-1.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/sequential_uuids_14-1.0.2-1.rhel8.x86_64.rpm) |
| `sequential_uuids_14` | 1.0.3 | `el8.aarch64` | pigsty | 12.8 KiB | [sequential_uuids_14-1.0.3-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/sequential_uuids_14-1.0.3-1PIGSTY.el8.aarch64.rpm) |
| `sequential_uuids_14` | 1.0.3 | `el9.x86_64` | pigsty | 12.7 KiB | [sequential_uuids_14-1.0.3-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/sequential_uuids_14-1.0.3-1PIGSTY.el9.x86_64.rpm) |
| `sequential_uuids_14` | 1.0.2 | `el9.x86_64` | pgdg | 21.1 KiB | [sequential_uuids_14-1.0.2-1.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/sequential_uuids_14-1.0.2-1.rhel9.x86_64.rpm) |
| `sequential_uuids_14` | 1.0.3 | `el9.aarch64` | pigsty | 12.6 KiB | [sequential_uuids_14-1.0.3-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/sequential_uuids_14-1.0.3-1PIGSTY.el9.aarch64.rpm) |
| `sequential_uuids_14` | 1.0.2 | `el9.aarch64` | pgdg | 15.1 KiB | [sequential_uuids_14-1.0.2-2.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/sequential_uuids_14-1.0.2-2.rhel9.aarch64.rpm) |
| `postgresql-14-sequential-uuids` | 1.0.3 | `d12.x86_64` | pigsty | 13.1 KiB | [postgresql-14-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-14-sequential-uuids_1.0.3-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-sequential-uuids` | 1.0.3 | `d12.aarch64` | pigsty | 13.0 KiB | [postgresql-14-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/sequential-uuids/postgresql-14-sequential-uuids_1.0.3-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-sequential-uuids` | 1.0.3 | `u22.x86_64` | pigsty | 13.2 KiB | [postgresql-14-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-14-sequential-uuids_1.0.3-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-sequential-uuids` | 1.0.3 | `u22.aarch64` | pigsty | 13.2 KiB | [postgresql-14-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/sequential-uuids/postgresql-14-sequential-uuids_1.0.3-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-sequential-uuids` | 1.0.3 | `u24.x86_64` | pigsty | 13.3 KiB | [postgresql-14-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-14-sequential-uuids_1.0.3-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-sequential-uuids` | 1.0.3 | `u24.aarch64` | pigsty | 13.1 KiB | [postgresql-14-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/sequential-uuids/postgresql-14-sequential-uuids_1.0.3-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/sequential-uuids" title="Repository" icon="github" subtitle="github.com/tvondra/sequential-uuids" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="sequential-uuids-1.0.3.tar.gz" >}}
{{< /cards >}}


```bash
pig build get sequential_uuids; # get sequential_uuids source code
pig build dep sequential_uuids; # install build dependencies
pig build pkg sequential_uuids; # build extension rpm or deb
pig build ext sequential_uuids; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install sequential_uuids; # install by extension name, for the current active PG version
pig ext install sequential_uuids; # install via package alias, for the active PG version
pig ext install sequential_uuids -v 18;   # install for PG 18
pig ext install sequential_uuids -v 17;   # install for PG 17
pig ext install sequential_uuids -v 16;   # install for PG 16
pig ext install sequential_uuids -v 15;   # install for PG 15
pig ext install sequential_uuids -v 14;   # install for PG 14
pig ext install sequential_uuids -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION sequential_uuids;
```

