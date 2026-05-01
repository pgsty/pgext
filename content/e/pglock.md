---
title: "pglock"
linkTitle: "pglock"
description: "Lightweight distributed lock service inside PostgreSQL"
weight: 4140
categories: ["UTIL"]
width: full
---

[**pglock**](https://github.com/fraruiz/pglock) : Lightweight distributed lock service inside PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4140** | {{< badge content="pglock" link="https://github.com/fraruiz/pglock" >}} | {{< ext "pglock" >}} | `1.0.0` | {{< category "UTIL" >}} | {{< license "PostgreSQL" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pglock` |
|   **Requires**    | {{< ext "pg_cron" >}} |
|   **See Also**    | {{< ext "pgmb" >}} {{< ext "pgmq" >}} {{< ext "pgq" >}} {{< ext "pg_cron" >}} |

> [!Note] Packaging patches the upstream pgmb.control mismatch and installs the extension as pglock.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pglock` | `pg_cron` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "pglock_18" "green" >}} {{< bg "17" "pglock_17" "green" >}} {{< bg "16" "pglock_16" "green" >}} {{< bg "15" "pglock_15" "green" >}} {{< bg "14" "pglock_14" "green" >}} | `pglock_$v` | `pg_cron_$v` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-pglock" "green" >}} {{< bg "17" "postgresql-17-pglock" "green" >}} {{< bg "16" "postgresql-16-pglock" "green" >}} {{< bg "15" "postgresql-15-pglock" "green" >}} {{< bg "14" "postgresql-14-pglock" "green" >}} | `postgresql-$v-pglock` | `postgresql-$v-cron` |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pglock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pglock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pglock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pglock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "pglock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "pglock_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "pglock_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-pglock : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-pglock : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglock_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pglock_18-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglock_18-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pglock_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pglock_18-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglock_18-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pglock_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pglock_18-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglock_18-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pglock_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pglock_18-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglock_18-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pglock_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pglock_18-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglock_18-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pglock_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pglock_18-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglock_18-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pglock` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.6 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.6 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pglock` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.0 KiB | [postgresql-18-pglock_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-18-pglock_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglock_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pglock_17-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglock_17-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pglock_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pglock_17-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglock_17-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pglock_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pglock_17-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglock_17-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pglock_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pglock_17-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglock_17-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pglock_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pglock_17-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglock_17-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pglock_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pglock_17-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglock_17-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pglock` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.6 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.6 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pglock` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.0 KiB | [postgresql-17-pglock_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-17-pglock_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglock_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pglock_16-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglock_16-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pglock_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pglock_16-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglock_16-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pglock_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pglock_16-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglock_16-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pglock_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pglock_16-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglock_16-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pglock_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pglock_16-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglock_16-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pglock_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pglock_16-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglock_16-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pglock` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.6 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.6 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pglock` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.0 KiB | [postgresql-16-pglock_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-16-pglock_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglock_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pglock_15-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglock_15-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pglock_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pglock_15-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglock_15-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pglock_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pglock_15-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglock_15-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pglock_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pglock_15-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglock_15-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pglock_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pglock_15-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglock_15-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pglock_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pglock_15-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglock_15-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pglock` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.6 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.6 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pglock` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.0 KiB | [postgresql-15-pglock_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-15-pglock_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pglock_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 KiB | [pglock_14-1.0.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pglock_14-1.0.0-1PIGSTY.el8.x86_64.rpm) |
| `pglock_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.3 KiB | [pglock_14-1.0.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pglock_14-1.0.0-1PIGSTY.el8.aarch64.rpm) |
| `pglock_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 KiB | [pglock_14-1.0.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pglock_14-1.0.0-1PIGSTY.el9.x86_64.rpm) |
| `pglock_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.3 KiB | [pglock_14-1.0.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pglock_14-1.0.0-1PIGSTY.el9.aarch64.rpm) |
| `pglock_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.5 KiB | [pglock_14-1.0.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pglock_14-1.0.0-1PIGSTY.el10.x86_64.rpm) |
| `pglock_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.4 KiB | [pglock_14-1.0.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pglock_14-1.0.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pglock` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 4.6 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 4.6 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 4.6 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 4.6 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 4.0 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pglock` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 4.0 KiB | [postgresql-14-pglock_1.0.0-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pglock/postgresql-14-pglock_1.0.0-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/fraruiz/pglock" title="Repository" icon="github" subtitle="github.com/fraruiz/pglock" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pglock-1.0.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pglock;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pglock;		# install via package name, for the active PG version

pig install pglock -v 18;   # install for PG 18
pig install pglock -v 17;   # install for PG 17
pig install pglock -v 16;   # install for PG 16
pig install pglock -v 15;   # install for PG 15
pig install pglock -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pglock CASCADE; -- requires pg_cron
```

## Usage

- Source: [README](https://github.com/fraruiz/pglock/blob/master/README.md)

`pglock` is a lightweight distributed lock service implemented inside PostgreSQL. It stores lock state in `pglock.locks` and uses TTL-based cleanup for stale rows.

### Create The Extension

```sql
CREATE EXTENSION pglock;
```

The upstream README lists PostgreSQL 9.1+ and `plpgsql` as requirements.

### Acquire And Release Locks

```sql
SELECT pglock.lock('worker-1', 'users');
SELECT pglock.unlock('worker-1', 'users');
```

`pglock.lock(id, resource)` returns `true` if the lock is acquired and `false` if another process already holds it. `pglock.unlock(id, resource)` is documented as idempotent.

### Isolation And Expiration

The README recommends serializable isolation for correctness under concurrency:

```sql
SELECT pglock.set_serializable();

BEGIN ISOLATION LEVEL SERIALIZABLE;
SELECT pglock.lock('my-id', 'my-resource');
SELECT pglock.unlock('my-id', 'my-resource');
COMMIT;
```

Stale locks are expired with:

```sql
SELECT pglock.ttl();
```

The documented default TTL is 5 minutes. If `pg_cron` is available, the README says `pglock.ttl()` can be scheduled every minute.

### Lock Table

The upstream schema is `pglock.locks` with columns `id`, `resource`, `locked`, `ttl`, `created_at`, and `updated_at`. The primary key is `(id, resource)`.
