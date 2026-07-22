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
| **9620** | {{< badge content="pgmqtt" link="https://github.com/RayElg/pgmqtt" >}} | {{< ext "pgmqtt" >}} | `0.4.1` | {{< category "ETL" >}} | {{< license "Elastic-2.0" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] requires wal_level = logical for CDC.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `pgmqtt` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "pgmqtt_18" "green" >}} {{< bg "17" "pgmqtt_17" "green" >}} {{< bg "16" "pgmqtt_16" "green" >}} {{< bg "15" "pgmqtt_15" "green" >}} {{< bg "14" "pgmqtt_14" "green" >}} | `pgmqtt_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.4.1` | {{< bg "18" "postgresql-18-pgmqtt" "green" >}} {{< bg "17" "postgresql-17-pgmqtt" "green" >}} {{< bg "16" "postgresql-16-pgmqtt" "green" >}} {{< bg "15" "postgresql-15-pgmqtt" "green" >}} {{< bg "14" "postgresql-14-pgmqtt" "green" >}} | `postgresql-$v-pgmqtt` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "pgmqtt_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-18-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-17-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-16-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-15-pgmqtt : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.4.1" "postgresql-14-pgmqtt : AVAIL 1" "green" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_18` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.6 MiB | [pgmqtt_18-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_18-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_18` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.2 MiB | [pgmqtt_18-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_18-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_18` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [pgmqtt_18-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_18-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_18` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [pgmqtt_18-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_18-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_18` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [pgmqtt_18-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_18-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_18` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.3 MiB | [pgmqtt_18-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_18-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pgmqtt` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-18-pgmqtt` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-18-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-18-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_17` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.6 MiB | [pgmqtt_17-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_17-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_17` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.2 MiB | [pgmqtt_17-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_17-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_17` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [pgmqtt_17-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_17-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_17` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [pgmqtt_17-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_17-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_17` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [pgmqtt_17-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_17-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_17` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.3 MiB | [pgmqtt_17-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_17-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pgmqtt` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.9 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.9 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-17-pgmqtt` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-17-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-17-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_16` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.6 MiB | [pgmqtt_16-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_16-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_16` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.2 MiB | [pgmqtt_16-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_16-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_16` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [pgmqtt_16-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_16-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_16` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [pgmqtt_16-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_16-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_16` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [pgmqtt_16-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_16-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_16` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.3 MiB | [pgmqtt_16-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_16-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pgmqtt` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.9 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.9 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-16-pgmqtt` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-16-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-16-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_15` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.6 MiB | [pgmqtt_15-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_15-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_15` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.2 MiB | [pgmqtt_15-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_15-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_15` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [pgmqtt_15-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_15-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_15` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [pgmqtt_15-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_15-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_15` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [pgmqtt_15-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_15-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_15` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.3 MiB | [pgmqtt_15-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_15-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pgmqtt` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.9 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.9 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-15-pgmqtt` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-15-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-15-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgmqtt_14` | `0.4.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.6 MiB | [pgmqtt_14-0.4.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgmqtt_14-0.4.1-1PIGSTY.el8.x86_64.rpm) |
| `pgmqtt_14` | `0.4.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.2 MiB | [pgmqtt_14-0.4.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgmqtt_14-0.4.1-1PIGSTY.el8.aarch64.rpm) |
| `pgmqtt_14` | `0.4.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [pgmqtt_14-0.4.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgmqtt_14-0.4.1-1PIGSTY.el9.x86_64.rpm) |
| `pgmqtt_14` | `0.4.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [pgmqtt_14-0.4.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgmqtt_14-0.4.1-1PIGSTY.el9.aarch64.rpm) |
| `pgmqtt_14` | `0.4.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [pgmqtt_14-0.4.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgmqtt_14-0.4.1-1PIGSTY.el10.x86_64.rpm) |
| `pgmqtt_14` | `0.4.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.3 MiB | [pgmqtt_14-0.4.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgmqtt_14-0.4.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pgmqtt` | `0.4.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.9 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 2.9 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.2 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 2.9 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 2.9 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~resolute_amd64.deb) |
| `postgresql-14-pgmqtt` | `0.4.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 2.9 MiB | [postgresql-14-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgmqtt/postgresql-14-pgmqtt_0.4.1-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/RayElg/pgmqtt" title="Repository" icon="github" subtitle="github.com/RayElg/pgmqtt" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pgmqtt-0.4.1.tar.gz" >}}
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

Sources:

- [pgmqtt 0.4.1 README](https://github.com/RayElg/pgmqtt/blob/0.4.1/README.md)
- [pgmqtt 0.4.1 interfaces](https://github.com/RayElg/pgmqtt/blob/0.4.1/docs/interfaces.md)
- [pgmqtt 0.4.1 configuration](https://github.com/RayElg/pgmqtt/blob/0.4.1/docs/configuration.md)
- [pgmqtt 0.4.1 limitations](https://github.com/RayElg/pgmqtt/blob/0.4.1/docs/limitations.md)
- [pgmqtt 0.4.1 release notes](https://github.com/RayElg/pgmqtt/releases/tag/0.4.1)

pgmqtt embeds an MQTT broker in PostgreSQL. It can publish INSERT, UPDATE, and DELETE changes through logical decoding and can map inbound MQTT topics and JSON payloads to table writes. Use it when database and MQTT integration justify running a network broker inside the PostgreSQL server process.

### Preload and Create the Extension

Set logical WAL and preload the worker, then restart PostgreSQL:

    wal_level = logical
    shared_preload_libraries = 'pgmqtt'

Create the extension after restart:

    CREATE EXTENSION pgmqtt;

Listener addresses, ports, authentication, and TLS settings are read by the background worker. Settings documented as startup-only require a worker/server restart, not just pg_reload_conf().

### Publish Table Changes

Create an outbound mapping:

    SELECT pgmqtt_add_outbound_mapping(
      'public',
      'orders',
      'orders/{{ op | lower }}',
      '{{ columns | tojson }}',
      1
    );

The mapping publishes row changes to topics such as orders/insert. The interface also accepts a QoS and template type where supported. Version 0.4.1 drains CDC changes in batches of up to 4096 records.

Inspect or remove outbound mappings:

    SELECT * FROM pgmqtt_list_outbound_mappings();
    SELECT pgmqtt_remove_outbound_mapping('public', 'orders');

### Write Rows from MQTT

Map captured topic segments and JSON fields to a target table:

    SELECT pgmqtt_add_inbound_mapping(
      'sensor/{site_id}/temperature',
      'sensor_readings',
      '{"site_id":"{site_id}","value":"$.temperature"}'::jsonb
    );

Inbound mappings support insert and documented upsert/delete modes with options such as target_schema, conflict_columns, mapping_name, and template_type. Grant the worker role only the required table privileges and validate payload types and constraints.

    SELECT * FROM pgmqtt_list_inbound_mappings();
    SELECT pgmqtt_remove_inbound_mapping('temp_readings');

### Administration and Status

    SELECT * FROM pgmqtt_status();
    SELECT pgmqtt_disconnect_client('device-42');
    SELECT pgmqtt_disconnect_role('mqtt_devices');
    SELECT pgmqtt_reload_acls('*');

pgmqtt_status reports listener, client, subscription, retained-message, CDC, inbound-write, and dead-letter state. Administrative calls are queued for asynchronous processing by the worker.

### Configuration Index

- pgmqtt.mqtt_enabled and pgmqtt.mqtt_port: TCP MQTT listener.
- pgmqtt.ws_enabled and pgmqtt.ws_port: WebSocket listener.
- pgmqtt.tick_interval_ms and pgmqtt.cdc_every_n_ticks: worker cadence.
- pgmqtt.max_client_buffer_bytes: per-client flow-control boundary.
- pgmqtt.debug_log and pgmqtt.metrics_*: diagnostics and metrics integration.
- pgmqtt TLS, JWT, password-authentication, and ACL settings: transport and client access controls; availability differs between community and enterprise features.

### Protocol and CDC Boundaries

- MQTT 5.0 and 3.1.1 are supported. QoS 0 and 1 are implemented; requested QoS 2 is downgraded to QoS 1.
- CDC covers INSERT, UPDATE, and DELETE, not DDL or TRUNCATE. DELETE payloads may require REPLICA IDENTITY FULL.
- The CDC ring has a finite capacity of 8192 and drops the oldest records on overflow. The QoS 0 topic buffer is capped at 4096 and also drops oldest entries; QoS 1 buffering can grow without a fixed bound.
- The community edition documents TLS through a proxy, while native TLS and some JWT features are enterprise boundaries. Verify the edition before setting listener expectations.

### Version 0.4.1 and Operations

The 0.4 line consolidates HTTP/worker handling and reduces panic paths; 0.4.1 raises CDC batch processing to 4096. These changes improve throughput and structure but do not make the embedded broker lossless under every overload or crash.

Running a broker inside PostgreSQL expands the database network and resource boundary. Isolate listener interfaces, enforce authentication and topic ACLs, monitor worker lag and dropped buffers, and test failover and restart behavior before production use.
