---
title: "pg_orphaned"
linkTitle: "pg_orphaned"
description: "Deal with orphaned files"
weight: 5200
categories: ["ADMIN"]
width: full
---

[**pg_orphaned**](https://github.com/bdrouvot/pg_orphaned) : Deal with orphaned files


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5200** | {{< badge content="pg_orphaned" link="https://github.com/bdrouvot/pg_orphaned" >}} | {{< ext "pg_orphaned" >}} | `1.0` | {{< category "ADMIN" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pg_dirtyread" >}} {{< ext "pg_surgery" >}} {{< ext "amcheck" >}} {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "pg_checksums" >}} {{< ext "pg_catcheck" >}} {{< ext "pg_repack" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pg_orphaned` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "pg_orphaned_18" "green" >}} {{< bg "17" "pg_orphaned_17" "green" >}} {{< bg "16" "pg_orphaned_16" "green" >}} {{< bg "15" "pg_orphaned_15" "green" >}} {{< bg "14" "pg_orphaned_14" "green" >}} | `pg_orphaned_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "postgresql-18-pg-orphaned" "green" >}} {{< bg "17" "postgresql-17-pg-orphaned" "green" >}} {{< bg "16" "postgresql-16-pg-orphaned" "green" >}} {{< bg "15" "postgresql-15-pg-orphaned" "green" >}} {{< bg "14" "postgresql-14-pg-orphaned" "green" >}} | `postgresql-$v-pg-orphaned` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "pg_orphaned_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0" "postgresql-18-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-17-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-16-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-15-pg-orphaned : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0" "postgresql-14-pg-orphaned : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-orphaned : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pg-orphaned : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pg-orphaned : MISS 0" "red" >}}      |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_18` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_18-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_18-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_18` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.1 KiB | [pg_orphaned_18-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_18-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_18` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.4 KiB | [pg_orphaned_18-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_18-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_18` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.4 KiB | [pg_orphaned_18-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_18-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_18` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.4 KiB | [pg_orphaned_18-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_18-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_18` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.5 KiB | [pg_orphaned_18-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_18-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pg-orphaned` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.1 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.9 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.9 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.0 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 30.8 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 30.9 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pg-orphaned` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.3 KiB | [postgresql-18-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-18-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_17` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_17-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_17` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.1 KiB | [pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_17-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_17` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_17-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_17` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.4 KiB | [pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_17-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_17` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.4 KiB | [pg_orphaned_17-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_17-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_17` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.5 KiB | [pg_orphaned_17-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_17-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pg-orphaned` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.1 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.9 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.9 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.0 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 34.8 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.0 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pg-orphaned` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.3 KiB | [postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-17-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_16` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_16-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_16` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.1 KiB | [pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_16-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_16` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_16-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_16` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.4 KiB | [pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_16-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_16` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.4 KiB | [pg_orphaned_16-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_16-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_16` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.5 KiB | [pg_orphaned_16-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_16-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pg-orphaned` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.1 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.8 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.9 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.0 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 34.4 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 34.6 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pg-orphaned` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.3 KiB | [postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-16-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_15` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_15-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_15` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.2 KiB | [pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_15-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_15` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.4 KiB | [pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_15-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_15` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.5 KiB | [pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_15-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_15` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.5 KiB | [pg_orphaned_15-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_15-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_15` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.6 KiB | [pg_orphaned_15-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_15-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pg-orphaned` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.2 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.9 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.9 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.0 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 34.5 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 34.7 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.2 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pg-orphaned` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.4 KiB | [postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-15-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pg_orphaned_14` | `1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 21.3 KiB | [pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pg_orphaned_14-1.0-1PIGSTY.el8.x86_64.rpm) |
| `pg_orphaned_14` | `1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 21.2 KiB | [pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pg_orphaned_14-1.0-1PIGSTY.el8.aarch64.rpm) |
| `pg_orphaned_14` | `1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 21.4 KiB | [pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pg_orphaned_14-1.0-1PIGSTY.el9.x86_64.rpm) |
| `pg_orphaned_14` | `1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 21.5 KiB | [pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pg_orphaned_14-1.0-1PIGSTY.el9.aarch64.rpm) |
| `pg_orphaned_14` | `1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 21.5 KiB | [pg_orphaned_14-1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pg_orphaned_14-1.0-1PIGSTY.el10.x86_64.rpm) |
| `pg_orphaned_14` | `1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 21.6 KiB | [pg_orphaned_14-1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pg_orphaned_14-1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pg-orphaned` | `1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 29.1 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 28.9 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 28.9 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 29.0 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 34.4 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 34.6 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 30.1 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pg-orphaned` | `1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 30.3 KiB | [postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pg-orphaned/postgresql-14-pg-orphaned_1.0-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/bdrouvot/pg_orphaned" title="Repository" icon="github" subtitle="github.com/bdrouvot/pg_orphaned" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pg_orphaned-1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pg_orphaned;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pg_orphaned;		# install via package name, for the active PG version

pig install pg_orphaned -v 18;   # install for PG 18
pig install pg_orphaned -v 17;   # install for PG 17
pig install pg_orphaned -v 16;   # install for PG 16
pig install pg_orphaned -v 15;   # install for PG 15
pig install pg_orphaned -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pg_orphaned;
```




## Usage

> [pg_orphaned: Deal with orphaned files](https://github.com/bdrouvot/pg_orphaned)

pg_orphaned provides functions to detect and manage orphaned data files in PostgreSQL. It handles corner cases like in-progress transactions that could cause false positives by using a dirty snapshot.

### List Orphaned Files

```sql
-- List orphaned files (default: older than 1 day marked as "older")
SELECT * FROM pg_list_orphaned();

-- Custom age threshold
SELECT * FROM pg_list_orphaned('10 seconds');
SELECT * FROM pg_list_orphaned('1 minute');
```

Returns: `dbname`, `path`, `name`, `size`, `mod_time`, `relfilenode`, `reloid`, `older` (boolean).

### Move Orphaned Files to Backup

```sql
-- Move files older than the threshold to orphaned_backup directory
SELECT pg_move_orphaned('1 minute');
```

### List Moved Files

```sql
SELECT * FROM pg_list_orphaned_moved();
```

### Move Files Back (if still orphaned)

```sql
SELECT pg_move_back_orphaned();
```

### Remove Moved Files

```sql
SELECT pg_remove_moved_orphaned();
```

### Typical Workflow

```sql
-- 1. Check for orphaned files
SELECT * FROM pg_list_orphaned('1 minute');

-- 2. Move them to backup (only those older than threshold)
SELECT pg_move_orphaned('1 minute');

-- 3. Verify what was moved
SELECT * FROM pg_list_orphaned_moved();

-- 4. After confirming, remove the backup files
SELECT pg_remove_moved_orphaned();
```

Note: functions operate on orphaned files for the database you are connected to. Always double-check carefully before moving or removing files.
