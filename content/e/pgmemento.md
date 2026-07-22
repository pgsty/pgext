---
title: "pgmemento"
linkTitle: "pgmemento"
description: "Transaction-based audit trail with schema versioning"
weight: 7190
categories: ["SEC"]
width: full
---

[**pgmemento**](https://github.com/pgMemento/pgMemento) : Transaction-based audit trail with schema versioning


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **7190** | {{< badge content="pgmemento" link="https://github.com/pgMemento/pgMemento" >}} | {{< ext "pgmemento" >}} | `0.7.4` | {{< category "SEC" >}} | {{< license "LGPL-3.0" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `pgmemento` |
|   **See Also**    | {{< ext "table_log" >}} {{< ext "table_version" >}} {{< ext "ddl_historization" >}} |

> [!Note] Packages upgrade scripts from 0.7, 0.7.1, 0.7.2, and 0.7.3 to 0.7.4.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.4` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmemento` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.4` | {{< bg "18" "pgmemento_18" "green" >}} {{< bg "17" "pgmemento_17" "green" >}} {{< bg "16" "pgmemento_16" "green" >}} {{< bg "15" "pgmemento_15" "green" >}} {{< bg "14" "pgmemento_14" "green" >}} | `pgmemento_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.7.4` | {{< bg "18" "postgresql-18-pgmemento" "green" >}} {{< bg "17" "postgresql-17-pgmemento" "green" >}} {{< bg "16" "postgresql-16-pgmemento" "green" >}} {{< bg "15" "postgresql-15-pgmemento" "green" >}} {{< bg "14" "postgresql-14-pgmemento" "green" >}} | `postgresql-$v-pgmemento` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "pgmemento_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-18-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-17-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-16-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-15-pgmemento : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.7.4" "postgresql-14-pgmemento : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemento_18` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.4 KiB | [pgmemento_18-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmemento_18-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_18` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 46.3 KiB | [pgmemento_18-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmemento_18-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_18` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pgmemento_18-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmemento_18-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_18` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.7 KiB | [pgmemento_18-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmemento_18-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_18` | `0.7.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.9 KiB | [pgmemento_18-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmemento_18-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `pgmemento_18` | `0.7.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pgmemento_18-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmemento_18-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `postgresql-18-pgmemento` | `0.7.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.2 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.2 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.2 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 34.2 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.0 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.0 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.0 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |
| `postgresql-18-pgmemento` | `0.7.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.0 KiB | [postgresql-18-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-18-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemento_17` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.4 KiB | [pgmemento_17-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmemento_17-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_17` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 46.3 KiB | [pgmemento_17-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmemento_17-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_17` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pgmemento_17-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmemento_17-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_17` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.7 KiB | [pgmemento_17-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmemento_17-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_17` | `0.7.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.9 KiB | [pgmemento_17-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmemento_17-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `pgmemento_17` | `0.7.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pgmemento_17-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmemento_17-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `postgresql-17-pgmemento` | `0.7.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.2 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.2 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.2 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 34.2 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.0 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.0 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.1 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |
| `postgresql-17-pgmemento` | `0.7.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.1 KiB | [postgresql-17-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-17-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemento_16` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.4 KiB | [pgmemento_16-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmemento_16-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_16` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 46.3 KiB | [pgmemento_16-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmemento_16-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_16` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pgmemento_16-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmemento_16-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_16` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.7 KiB | [pgmemento_16-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmemento_16-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_16` | `0.7.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.9 KiB | [pgmemento_16-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmemento_16-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `pgmemento_16` | `0.7.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pgmemento_16-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmemento_16-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `postgresql-16-pgmemento` | `0.7.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.2 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.2 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.2 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 34.2 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.0 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.0 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.0 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |
| `postgresql-16-pgmemento` | `0.7.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.0 KiB | [postgresql-16-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-16-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemento_15` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.4 KiB | [pgmemento_15-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmemento_15-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_15` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 46.3 KiB | [pgmemento_15-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmemento_15-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_15` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pgmemento_15-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmemento_15-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_15` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.7 KiB | [pgmemento_15-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmemento_15-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_15` | `0.7.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.9 KiB | [pgmemento_15-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmemento_15-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `pgmemento_15` | `0.7.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pgmemento_15-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmemento_15-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `postgresql-15-pgmemento` | `0.7.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.2 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.2 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.2 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 34.2 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.0 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.0 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.0 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |
| `postgresql-15-pgmemento` | `0.7.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.0 KiB | [postgresql-15-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-15-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmemento_14` | `0.7.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 46.4 KiB | [pgmemento_14-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmemento_14-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_14` | `0.7.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 46.3 KiB | [pgmemento_14-0.7.4-1PIGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmemento_14-0.7.4-1PIGSTY.el8.noarch.rpm) |
| `pgmemento_14` | `0.7.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 43.7 KiB | [pgmemento_14-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmemento_14-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_14` | `0.7.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 43.7 KiB | [pgmemento_14-0.7.4-1PIGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmemento_14-0.7.4-1PIGSTY.el9.noarch.rpm) |
| `pgmemento_14` | `0.7.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 43.9 KiB | [pgmemento_14-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmemento_14-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `pgmemento_14` | `0.7.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 43.8 KiB | [pgmemento_14-0.7.4-1PIGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmemento_14-0.7.4-1PIGSTY.el10.noarch.rpm) |
| `postgresql-14-pgmemento` | `0.7.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 34.2 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 34.2 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~bookworm_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 34.2 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 34.2 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~trixie_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 35.0 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 35.0 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~jammy_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 35.0 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 35.0 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~noble_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 35.0 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |
| `postgresql-14-pgmemento` | `0.7.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 35.0 KiB | [postgresql-14-pgmemento_0.7.4-1PIGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmemento/postgresql-14-pgmemento_0.7.4-1PIGSTY~resolute_all.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgMemento/pgMemento" title="Repository" icon="github" subtitle="github.com/pgMemento/pgMemento" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmemento-0.7.4.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmemento;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmemento;		# install via package name, for the active PG version

pig install pgmemento -v 18;   # install for PG 18
pig install pgmemento -v 17;   # install for PG 17
pig install pgmemento -v 16;   # install for PG 16
pig install pgmemento -v 15;   # install for PG 15
pig install pgmemento -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmemento;
```

## Usage

Sources:

- [pgMemento v0.7.4 README](https://github.com/pgMemento/pgMemento/blob/v0.7.4/README.md)
- [pgMemento v0.7.4 control file](https://github.com/pgMemento/pgMemento/blob/v0.7.4/extension/pgmemento.control)
- [pgMemento wiki](https://github.com/pgMemento/pgMemento/wiki)
- [Changes from v0.7.3 to v0.7.4](https://github.com/pgMemento/pgMemento/compare/v0.7.3...v0.7.4)

pgmemento is a trigger-based audit trail for PostgreSQL. It records DML changes as JSONB deltas, groups them by transaction and table event, tracks audited schema changes, and provides restore and revert helpers. Use it when row history and transaction context must be queried inside PostgreSQL.

### Create and Initialize

    CREATE EXTENSION pgmemento;
    SELECT pgmemento.init('public');

init instruments eligible tables in the selected schema and adds the pgmemento_audit_id identity column used to track row history. Run it first on a staging copy: auditing changes table definitions and installs event and row triggers.

Use start and stop to control auditing for a schema, and use the documented drop function when intentionally removing pgMemento's instrumentation. Do not manually delete extension triggers or audit identifiers.

### Inspect the Audit Trail

The central data model includes:

- transaction_log: transaction metadata and optional application context.
- table_event_log: table-level events within a transaction.
- row_log: JSONB row deltas linked to a table event.
- audited_schema and audited_table metadata: tracked schemas, tables, columns, and versions.

A typical investigation joins a transaction to its table events and row deltas:

    SELECT t.id,
           t.txid_time,
           e.table_operation,
           r.audit_id,
           r.old_data,
           r.new_data
    FROM pgmemento.transaction_log AS t
    JOIN pgmemento.table_event_log AS e
      ON e.transaction_id = t.id
    JOIN pgmemento.row_log AS r
      ON r.event_key = e.event_key
    WHERE t.id = 12345;

Inspect the installed views and column names before embedding this query because audit schema details can differ across pgmemento versions.

### Restore and Revert

pgmemento provides restore functions that reconstruct table state from the audit trail and revert_transaction or related helpers that apply compensating changes. Treat these as recovery operations:

1. take and verify a backup;
2. identify the exact transaction and dependent changes;
3. preview reconstructed data where possible;
4. run the operation in a controlled transaction;
5. validate constraints, sequences, and application invariants.

### Version 0.7.4

Version 0.7.4 changes row serialization to avoid PostgreSQL's jsonb_build_object argument limit for very wide payloads and adds PostgreSQL 15 support. Upgrade using ALTER EXTENSION pgmemento UPDATE only after testing the version-specific upgrade script.

### Operational Boundaries

- Audit triggers add latency and write volume to every tracked change. Monitor row_log growth and index maintenance.
- The audit trail resides in the same database and is not a substitute for backups, WAL archives, or tamper-resistant external audit storage.
- Schema initialization and DDL tracking alter application tables. Coordinate migrations with pgmemento rather than bypassing its event triggers.
- Limit direct writes to the pgmemento schema and protect any transaction metadata that can contain user or application information.
