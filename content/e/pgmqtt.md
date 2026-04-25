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
| **9620** | {{< badge content="pgmqtt" link="https://github.com/RayElg/pgmqtt" >}} | {{< ext "pgmqtt" >}} | `0.1.0` | {{< category "ETL" >}} | {{< license "Elastic License 2.0" >}} | {{< language "Rust" >}} |


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
| {{< os "u26.aarch64" >}} |      {{< bg "MISS" "postgresql-18-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-17-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-16-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-15-pgmqtt : MISS 0" "red" >}}      |      {{< bg "MISS" "postgresql-14-pgmqtt : MISS 0" "red" >}}      |


{{< tabs items="PG18,PG17,PG16,PG15,PG14" >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

{{< /tab >}}
{{< tab >}}

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

Sources: [official README](https://github.com/RayElg/pgmqtt/blob/main/README.md), [official repo](https://github.com/RayElg/pgmqtt)

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
  '{{ columns | tojson }}'
);
```

With that mapping, `INSERT`, `UPDATE`, and `DELETE` publish JSON payloads to topics such as `topics/insert`.

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

### MQTT Client Examples

```bash
mosquitto_sub -h localhost -t 'topics/#'
mosquitto_pub -h localhost -t 'sensor/site-1/temperature' -m '{"temperature": 22.5}'
```

### Caveats

- The README requires `wal_level = logical`; without logical decoding the CDC side will not work.
- Upstream documentation is currently README-level only, so the documented SQL surface is limited to the inbound and outbound mapping workflow.
