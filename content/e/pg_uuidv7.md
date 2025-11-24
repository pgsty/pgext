---
title: "pg_uuidv7"
linkTitle: "pg_uuidv7"
description: "Create UUIDv7 values in postgres"
weight: 4540
categories: ["FUNC"]
width: full
---

[**pg_uuidv7**](https://github.com/fboulnois/pg_uuidv7) : Create UUIDv7 values in postgres


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4540** | {{< badge content="pg_uuidv7" link="https://github.com/fboulnois/pg_uuidv7" >}} | {{< ext "pg_uuidv7" >}} | `1.7.0` | {{< category "FUNC" >}} | {{< license "MPL-2.0" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_idkit" >}} {{< ext "pgx_ulid" >}} {{< ext "uuid-ossp" >}} {{< ext "sequential_uuids" >}} {{< ext "pg_hashids" >}} {{< ext "permuteseq" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="MIXED" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pg_uuidv7` | - |
| **RPM** | {{< badge content="PGDG" link="/repo/pgdg" >}} | `1.7.0` | {{< bg "18" "pg_uuidv7_18*" "green" >}} {{< bg "17" "pg_uuidv7_17*" "green" >}} {{< bg "16" "pg_uuidv7_16*" "green" >}} {{< bg "15" "pg_uuidv7_15*" "green" >}} {{< bg "14" "pg_uuidv7_14*" "green" >}} {{< bg "13" "pg_uuidv7_13*" "green" >}} | `pg_uuidv7_$v*` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.7.0` | {{< bg "18" "postgresql-18-pg-uuidv7" "green" >}} {{< bg "17" "postgresql-17-pg-uuidv7" "green" >}} {{< bg "16" "postgresql-16-pg-uuidv7" "green" >}} {{< bg "15" "postgresql-15-pg-uuidv7" "green" >}} {{< bg "14" "postgresql-14-pg-uuidv7" "green" >}} {{< bg "13" "postgresql-13-pg-uuidv7" "green" >}} | `postgresql-$v-pg-uuidv7` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_15 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_14 : AVAIL 10" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_13 : AVAIL 10" "blue" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_13 : AVAIL 9" "blue" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_13 : AVAIL 9" "blue" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_17 : AVAIL 3" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_16 : AVAIL 7" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_15 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_14 : AVAIL 9" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_13 : AVAIL 9" "blue" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_18 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_17 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_16 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_15 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_14 : AVAIL 2" "blue" >}} | {{< bg "PGDG 1.7.0" "pg_uuidv7_13 : AVAIL 2" "blue" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PGDG 1.6.0" "pg_uuidv7_18 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_uuidv7_17 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_uuidv7_16 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_uuidv7_15 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_uuidv7_14 : AVAIL 1" "blue" >}} | {{< bg "PGDG 1.6.0" "pg_uuidv7_13 : AVAIL 1" "blue" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-18-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-17-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-16-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-15-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-14-pg-uuidv7 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.7.0" "postgresql-13-pg-uuidv7 : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuidv7_18` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_uuidv7_18-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_uuidv7_18-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_18` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_18-1.6.0-2PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-x86_64/pg_uuidv7_18-1.6.0-2PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_18` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.8 KiB | [pg_uuidv7_18-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_uuidv7_18-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_18` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.3 KiB | [pg_uuidv7_18-1.6.0-2PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-8-aarch64/pg_uuidv7_18-1.6.0-2PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_18` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_18-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_uuidv7_18-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_18` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.3 KiB | [pg_uuidv7_18-1.6.0-2PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-x86_64/pg_uuidv7_18-1.6.0-2PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_18` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_uuidv7_18-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_uuidv7_18-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_18` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.1 KiB | [pg_uuidv7_18-1.6.0-2PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-9-aarch64/pg_uuidv7_18-1.6.0-2PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_18` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.5 KiB | [pg_uuidv7_18-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_uuidv7_18-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_18` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_18-1.6.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10-x86_64/pg_uuidv7_18-1.6.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_18` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_uuidv7_18-1.6.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-10.0-aarch64/pg_uuidv7_18-1.6.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.1 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.1 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.7 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.6 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-uuidv7` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-18-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuidv7_17` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_uuidv7_17-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_uuidv7_17-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_17` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [pg_uuidv7_17-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_uuidv7_17-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_17` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.8 KiB | [pg_uuidv7_17-1.5.0-3PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-x86_64/pg_uuidv7_17-1.5.0-3PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_17` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.8 KiB | [pg_uuidv7_17-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_uuidv7_17-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_17` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_17-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_uuidv7_17-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_17` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.9 KiB | [pg_uuidv7_17-1.5.0-3PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-8-aarch64/pg_uuidv7_17-1.5.0-3PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_17` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_17-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_uuidv7_17-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_17` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_17-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_uuidv7_17-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_17` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.9 KiB | [pg_uuidv7_17-1.5.0-3PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-x86_64/pg_uuidv7_17-1.5.0-3PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_17` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_uuidv7_17-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_uuidv7_17-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_17` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_17-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_uuidv7_17-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_17` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.8 KiB | [pg_uuidv7_17-1.5.0-3PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-9-aarch64/pg_uuidv7_17-1.5.0-3PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_17` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_17-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_uuidv7_17-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_17` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_17-1.6.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10-x86_64/pg_uuidv7_17-1.6.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_17` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_uuidv7_17-1.6.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-10.0-aarch64/pg_uuidv7_17-1.6.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.1 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.1 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.8 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.7 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-uuidv7` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-17-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuidv7_16` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_uuidv7_16-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [pg_uuidv7_16-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_16-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_uuidv7_16-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.0 KiB | [pg_uuidv7_16-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_16-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_uuidv7_16-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-x86_64/pg_uuidv7_16-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_16` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.8 KiB | [pg_uuidv7_16-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_16-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_uuidv7_16-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.2 KiB | [pg_uuidv7_16-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_uuidv7_16-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_16-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_16-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-8-aarch64/pg_uuidv7_16-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_16` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_16-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_16-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_16-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [pg_uuidv7_16-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_16-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pg_uuidv7_16-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_uuidv7_16-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-x86_64/pg_uuidv7_16-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_16` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_uuidv7_16-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_16-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_uuidv7_16-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_16-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_16-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_uuidv7_16-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_uuidv7_16-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-9-aarch64/pg_uuidv7_16-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_16` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_16-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_uuidv7_16-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_16` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.7 KiB | [pg_uuidv7_16-1.6.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10-x86_64/pg_uuidv7_16-1.6.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_16` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_uuidv7_16-1.6.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-10.0-aarch64/pg_uuidv7_16-1.6.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.0 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.9 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.8 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-uuidv7` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-16-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuidv7_15` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_uuidv7_15-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [pg_uuidv7_15-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_15-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_uuidv7_15-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.0 KiB | [pg_uuidv7_15-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_15-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_uuidv7_15-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 KiB | [pg_uuidv7_15-1.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.0 KiB | [pg_uuidv7_15-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.7 KiB | [pg_uuidv7_15-1.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-x86_64/pg_uuidv7_15-1.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_15` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.8 KiB | [pg_uuidv7_15-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_15-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_uuidv7_15-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.2 KiB | [pg_uuidv7_15-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_uuidv7_15-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_15-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_15-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_uuidv7_15-1.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_uuidv7_15-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-8-aarch64/pg_uuidv7_15-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_15` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_15-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_15-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_15-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [pg_uuidv7_15-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.0 KiB | [pg_uuidv7_15-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pg_uuidv7_15-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_uuidv7_15-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.1 KiB | [pg_uuidv7_15-1.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.9 KiB | [pg_uuidv7_15-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-x86_64/pg_uuidv7_15-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_15` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_uuidv7_15-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_15-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_uuidv7_15-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_15-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_15-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_uuidv7_15-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_uuidv7_15-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.9 KiB | [pg_uuidv7_15-1.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.7 KiB | [pg_uuidv7_15-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-9-aarch64/pg_uuidv7_15-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_15` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_15-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_uuidv7_15-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_15` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_15-1.6.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10-x86_64/pg_uuidv7_15-1.6.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_15` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_uuidv7_15-1.6.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-10.0-aarch64/pg_uuidv7_15-1.6.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.0 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.8 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-uuidv7` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-15-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuidv7_14` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.7 KiB | [pg_uuidv7_14-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.1 KiB | [pg_uuidv7_14-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_14-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.1 KiB | [pg_uuidv7_14-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.0 KiB | [pg_uuidv7_14-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_14-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.7 KiB | [pg_uuidv7_14-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.2 KiB | [pg_uuidv7_14-1.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.0 KiB | [pg_uuidv7_14-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.7 KiB | [pg_uuidv7_14-1.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-x86_64/pg_uuidv7_14-1.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_14` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.8 KiB | [pg_uuidv7_14-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_14-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_uuidv7_14-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_uuidv7_14-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.0 KiB | [pg_uuidv7_14-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_14-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_14-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_uuidv7_14-1.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_uuidv7_14-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-8-aarch64/pg_uuidv7_14-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_14` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_14-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_14-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_14-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.1 KiB | [pg_uuidv7_14-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_14-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pg_uuidv7_14-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_uuidv7_14-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.1 KiB | [pg_uuidv7_14-1.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.9 KiB | [pg_uuidv7_14-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-x86_64/pg_uuidv7_14-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_14` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_uuidv7_14-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.1 KiB | [pg_uuidv7_14-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_uuidv7_14-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_14-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_14-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_uuidv7_14-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_uuidv7_14-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.8 KiB | [pg_uuidv7_14-1.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.7 KiB | [pg_uuidv7_14-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-9-aarch64/pg_uuidv7_14-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_14` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_14-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_uuidv7_14-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_14` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_14-1.6.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10-x86_64/pg_uuidv7_14-1.6.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_14` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_uuidv7_14-1.6.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-10.0-aarch64/pg_uuidv7_14-1.6.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.8 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 18.0 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.7 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 18.0 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.8 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.7 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-uuidv7` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.7 KiB | [postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-14-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_uuidv7_13` | `1.7.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_13-1.7.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.7.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.6.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 21.0 KiB | [pg_uuidv7_13-1.6.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.6.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.5.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.5 KiB | [pg_uuidv7_13-1.5.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.5.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.4.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 20.0 KiB | [pg_uuidv7_13-1.4.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.4.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.4.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_13-1.4.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.4.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.3.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.8 KiB | [pg_uuidv7_13-1.3.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.3.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.2.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.6 KiB | [pg_uuidv7_13-1.2.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.2.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.1.1` | [el8.x86_64](/os/el8.x86_64) | pgdg | 19.1 KiB | [pg_uuidv7_13-1.1.1-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.1.1-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.1.0` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.9 KiB | [pg_uuidv7_13-1.1.0-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.1.0-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.0.2` | [el8.x86_64](/os/el8.x86_64) | pgdg | 18.7 KiB | [pg_uuidv7_13-1.0.2-1PGDG.rhel8.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-x86_64/pg_uuidv7_13-1.0.2-1PGDG.rhel8.x86_64.rpm) |
| `pg_uuidv7_13` | `1.7.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.8 KiB | [pg_uuidv7_13-1.7.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.7.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.6.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 21.2 KiB | [pg_uuidv7_13-1.6.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.6.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.5.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.7 KiB | [pg_uuidv7_13-1.5.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.5.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.4.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.1 KiB | [pg_uuidv7_13-1.4.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.4.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.4.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 20.0 KiB | [pg_uuidv7_13-1.4.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.4.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.3.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_13-1.3.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.3.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.2.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_13-1.2.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.2.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.1.1` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.2 KiB | [pg_uuidv7_13-1.1.1-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.1.1-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.1.0` | [el8.aarch64](/os/el8.aarch64) | pgdg | 19.1 KiB | [pg_uuidv7_13-1.1.0-1PGDG.rhel8.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-8-aarch64/pg_uuidv7_13-1.1.0-1PGDG.rhel8.aarch64.rpm) |
| `pg_uuidv7_13` | `1.7.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_13-1.7.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.7.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.6.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 21.2 KiB | [pg_uuidv7_13-1.6.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.6.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.5.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.6 KiB | [pg_uuidv7_13-1.5.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.5.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.4.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 20.0 KiB | [pg_uuidv7_13-1.4.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.4.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.4.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.9 KiB | [pg_uuidv7_13-1.4.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.4.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.3.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.8 KiB | [pg_uuidv7_13-1.3.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.3.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.2.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.6 KiB | [pg_uuidv7_13-1.2.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.2.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.1.1` | [el9.x86_64](/os/el9.x86_64) | pgdg | 19.1 KiB | [pg_uuidv7_13-1.1.1-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.1.1-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.1.0` | [el9.x86_64](/os/el9.x86_64) | pgdg | 18.9 KiB | [pg_uuidv7_13-1.1.0-1PGDG.rhel9.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-x86_64/pg_uuidv7_13-1.1.0-1PGDG.rhel9.x86_64.rpm) |
| `pg_uuidv7_13` | `1.7.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.0 KiB | [pg_uuidv7_13-1.7.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.7.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.6.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 21.1 KiB | [pg_uuidv7_13-1.6.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.6.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.5.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 20.5 KiB | [pg_uuidv7_13-1.5.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.5.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.4.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.9 KiB | [pg_uuidv7_13-1.4.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.4.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.4.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.7 KiB | [pg_uuidv7_13-1.4.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.4.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.3.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.6 KiB | [pg_uuidv7_13-1.3.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.3.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.2.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 19.3 KiB | [pg_uuidv7_13-1.2.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.2.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.1.1` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.8 KiB | [pg_uuidv7_13-1.1.1-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.1.1-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.1.0` | [el9.aarch64](/os/el9.aarch64) | pgdg | 18.7 KiB | [pg_uuidv7_13-1.1.0-1PGDG.rhel9.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-9-aarch64/pg_uuidv7_13-1.1.0-1PGDG.rhel9.aarch64.rpm) |
| `pg_uuidv7_13` | `1.7.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.5 KiB | [pg_uuidv7_13-1.7.0-1PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_uuidv7_13-1.7.0-1PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_13` | `1.6.0` | [el10.x86_64](/os/el10.x86_64) | pgdg | 21.6 KiB | [pg_uuidv7_13-1.6.0-2PGDG.rhel10.x86_64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10-x86_64/pg_uuidv7_13-1.6.0-2PGDG.rhel10.x86_64.rpm) |
| `pg_uuidv7_13` | `1.6.0` | [el10.aarch64](/os/el10.aarch64) | pgdg | 21.7 KiB | [pg_uuidv7_13-1.6.0-2PGDG.rhel10.aarch64.rpm](https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-10.0-aarch64/pg_uuidv7_13-1.6.0-2PGDG.rhel10.aarch64.rpm) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 17.9 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 17.8 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 17.8 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 17.9 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.7 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.7 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.9 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pg-uuidv7` | `1.7.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.5 KiB | [postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-uuidv7/postgresql-13-pg-uuidv7_1.7.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fboulnois/pg_uuidv7" title="Repository" icon="github" subtitle="github.com/fboulnois/pg_uuidv7" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_uuidv7-1.7.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_uuidv7;		# build spec not ready
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_uuidv7;		# install via package name, for the active PG version

pig install pg_uuidv7 -v 18;   # install for PG 18
pig install pg_uuidv7 -v 17;   # install for PG 17
pig install pg_uuidv7 -v 16;   # install for PG 16
pig install pg_uuidv7 -v 15;   # install for PG 15
pig install pg_uuidv7 -v 14;   # install for PG 14
pig install pg_uuidv7 -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_uuidv7;
```
