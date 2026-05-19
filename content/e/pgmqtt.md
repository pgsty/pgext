---
title: "pgmqtt"
linkTitle: "pgmqtt"
description: "CDC-to-MQTT broker for PostgreSQL"
weight: 9620
categories: ["ETL"]
width: full
---

[**pgmqtt**](https://github.com/RayElg/pgmqtt) : CDC-to-MQTT broker for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9620** | {{< badge content="pgmqtt" link="https://github.com/RayElg/pgmqtt" >}} | {{< ext "pgmqtt" >}} | `0.1.0` | {{< category "ETL" >}} | {{< license "ELv2" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] manually upgraded PGRX from 0.16.1 to 0.17.0 by Vonng; requires wal_level = logical for CDC.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmqtt` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "pgmqtt_18" "green" >}} {{< bg "17" "pgmqtt_17" "green" >}} {{< bg "16" "pgmqtt_16" "green" >}} {{< bg "15" "pgmqtt_15" "green" >}} {{< bg "14" "pgmqtt_14" "green" >}} | `pgmqtt_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-pgmqtt" "green" >}} {{< bg "17" "postgresql-17-pgmqtt" "green" >}} {{< bg "16" "postgresql-16-pgmqtt" "green" >}} {{< bg "15" "postgresql-15-pgmqtt" "green" >}} {{< bg "14" "postgresql-14-pgmqtt" "green" >}} | `postgresql-$v-pgmqtt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} |      {{< bg "MISS" "postgresql-18-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgmqtt : MISS 0" "red" >}}      |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pgmqtt_18-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_18-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pgmqtt_18-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_18-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pgmqtt_18-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_18-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pgmqtt_18-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_18-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pgmqtt_18-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_18-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pgmqtt_18-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_18-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmqtt` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-18-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.4 MiB | [postgresql-18-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pgmqtt_17-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_17-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pgmqtt_17-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_17-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pgmqtt_17-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_17-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pgmqtt_17-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_17-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pgmqtt_17-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_17-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pgmqtt_17-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_17-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmqtt` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-17-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.4 MiB | [postgresql-17-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_16` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pgmqtt_16-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_16-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_16` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pgmqtt_16-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_16-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_16` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pgmqtt_16-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_16-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_16` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pgmqtt_16-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_16-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_16` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pgmqtt_16-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_16-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_16` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pgmqtt_16-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_16-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmqtt` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-16-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.4 MiB | [postgresql-16-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_15` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pgmqtt_15-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_15-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_15` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pgmqtt_15-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_15-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_15` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pgmqtt_15-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_15-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_15` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pgmqtt_15-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_15-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_15` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pgmqtt_15-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_15-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_15` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pgmqtt_15-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_15-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmqtt` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-15-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.4 MiB | [postgresql-15-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_14` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 1.7 MiB | [pgmqtt_14-0.1.0-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_14-0.1.0-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_14` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 1.4 MiB | [pgmqtt_14-0.1.0-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_14-0.1.0-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_14` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 1.7 MiB | [pgmqtt_14-0.1.0-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_14-0.1.0-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_14` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 1.5 MiB | [pgmqtt_14-0.1.0-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_14-0.1.0-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_14` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 1.6 MiB | [pgmqtt_14-0.1.0-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_14-0.1.0-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_14` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 1.4 MiB | [pgmqtt_14-0.1.0-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_14-0.1.0-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmqtt` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 1.2 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 1.4 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 1.2 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 1.5 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 1.5 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 1.3 MiB | [postgresql-14-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 1.4 MiB | [postgresql-14-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.1.0-2PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RayElg/pgmqtt" title="Repository" icon="github" subtitle="github.com/RayElg/pgmqtt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmqtt-0.1.0.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pgmqtt;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pgmqtt;		# install via package name, for the active PG version

pig install pgmqtt -v 18;   # install for PG 18
pig install pgmqtt -v 17;   # install for PG 17
pig install pgmqtt -v 16;   # install for PG 16
pig install pgmqtt -v 15;   # install for PG 15
pig install pgmqtt -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pgmqtt;
```

## Usage

Sources: [README](https://github.com/RayElg/pgmqtt/blob/main/README.md), [interfaces](https://github.com/RayElg/pgmqtt/blob/main/docs/interfaces.md), [configuration](https://github.com/RayElg/pgmqtt/blob/main/docs/configuration.md), [limitations](https://github.com/RayElg/pgmqtt/blob/main/docs/limitations.md), [Cargo.toml](https://github.com/RayElg/pgmqtt/blob/main/extension/Cargo.toml)

`pgmqtt` is a `pgrx` extension that embeds an MQTT broker into PostgreSQL and uses change data capture to turn table changes into MQTT messages. It also supports inbound topic mappings so MQTT publishes can insert rows into PostgreSQL tables.

```sql
CREATE EXTENSION pgmqtt;
```

### Outbound Mapping

Publish table changes to topics:

```sql
SELECT pgmqtt_add_outbound_mapping(
  'public',
  'my_table',
  'topics/{{ op | lower }}',
  '{{ columns | tojson }}',
  1
);
```

With that mapping, `INSERT`, `UPDATE`, and `DELETE` publish JSON payloads to topics such as `topics/insert`. The documented function signature also accepts optional `qos integer DEFAULT 0` and `template_type text DEFAULT 'jinja2'`.

### Inbound Mapping

Insert rows from MQTT publishes:

```sql
SELECT pgmqtt_add_inbound_mapping(
  'sensor/{site_id}/temperature',
  'sensor_readings',
  '{"site_id": "{site_id}", "value": "$.temperature"}'::jsonb
);
```

Publishing `{"temperature": 22.5}` to `sensor/site-1/temperature` inserts a row into `sensor_readings`.

Inbound mappings can also perform `upsert` and `delete` operations by passing `op`, `conflict_columns`, `target_schema`, `mapping_name`, and `template_type`. Topic patterns use `{variable}` captures; JSON payload fields use expressions such as `$.temperature`, `$payload`, and `$topic`.

### Inspect and Remove Mappings

```sql
SELECT * FROM pgmqtt_list_outbound_mappings();
SELECT pgmqtt_remove_outbound_mapping('public', 'my_table');

SELECT * FROM pgmqtt_list_inbound_mappings();
SELECT pgmqtt_remove_inbound_mapping('temp_readings');

SELECT * FROM pgmqtt_status();
```

`pgmqtt_status()` reports active connections, subscriptions, retained messages, pending session messages, CDC mappings, CDC slot state, inbound mappings, pending inbound writes, and dead letters.

### MQTT Client Examples

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### Configuration

The documented GUCs live under the `pgmqtt` namespace:

```sql
ALTER SYSTEM SET pgmqtt.cdc_every_n_ticks = 16;
SELECT pg_reload_conf();
```

Listener GUCs include `pgmqtt.mqtt_enabled`, `pgmqtt.mqtt_port` (`1883`), `pgmqtt.ws_enabled`, `pgmqtt.ws_port` (`9001`), `pgmqtt.mqtts_enabled`, `pgmqtt.mqtts_port` (`8883`), `pgmqtt.wss_enabled`, and `pgmqtt.wss_port` (`9002`). TLS and authentication settings include `pgmqtt.tls_cert_file`, `pgmqtt.tls_key_file`, `pgmqtt.license_key`, `pgmqtt.jwt_public_key`, `pgmqtt.jwt_required`, and `pgmqtt.jwt_required_ws`.

Performance and observability GUCs include `pgmqtt.tick_interval_ms`, `pgmqtt.max_client_buffer_bytes`, `pgmqtt.cdc_every_n_ticks`, `pgmqtt.debug_log`, `pgmqtt.metrics_snapshot_interval`, `pgmqtt.metrics_retention_days`, `pgmqtt.metrics_connections_cache_interval`, `pgmqtt.metrics_hook_function`, and `pgmqtt.metrics_notify_channel`. Listener and TLS settings are read when the MQTT background worker starts, so they require a worker restart rather than only `pg_reload_conf()`.

### Caveats

- The README requires `wal_level = logical`; without logical decoding the CDC side will not work.
- This project's CSV tracks version `0.1.0`, PostgreSQL versions 14-18, and a package-side `pgrx` `0.17.0` rebuild note. Upstream `Cargo.toml` still advertises older build defaults, so use the CSV as the package/platform authority here.
- MQTT 5.0 and MQTT 3.1.1 clients are documented as supported. QoS 0 and QoS 1 are supported; QoS 2 is not implemented and subscriptions requesting QoS 2 are downgraded to QoS 1.
- CDC captures `INSERT`, `UPDATE`, and `DELETE`; DDL changes and `TRUNCATE` are not captured. `DELETE` may require `REPLICA IDENTITY FULL`.
