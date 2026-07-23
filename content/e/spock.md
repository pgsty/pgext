---
title: "spock"
linkTitle: "spock"
description: "Multi-master logical replication extension for PostgreSQL"
weight: 9570
categories: ["ETL"]
width: full
---

[**spock**](https://github.com/pgEdge/spock) : Multi-master logical replication extension for PostgreSQL


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9570** | {{< badge content="spock" link="https://github.com/pgEdge/spock" >}} | {{< ext "spock" >}} | `5.0.10` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--sLd--" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="Yes" color="orange" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `spock` |
|   **See Also**    | {{< ext "lolor" >}} {{< ext "snowflake" >}} |

> [!Note] works on pgedge kernel fork


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `5.0.10` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `spock` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.4` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.4` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 18.4" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.10" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.14" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.18" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |


{{< tabs >}}
{{< tab name="PG18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-18` | `18.4` | [el8.x86_64](/os/el8.x86_64) | pigsty | 13.2 MiB | [pgedge-18-18.4-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-18-18.4-1PIGSTY.el8.x86_64.rpm) |
| `pgedge-18` | `18.4` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.8 MiB | [pgedge-18-18.4-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-18-18.4-1PIGSTY.el8.aarch64.rpm) |
| `pgedge-18` | `18.4` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.1 MiB | [pgedge-18-18.4-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-18-18.4-1PIGSTY.el9.x86_64.rpm) |
| `pgedge-18` | `18.4` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.8 MiB | [pgedge-18-18.4-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-18-18.4-1PIGSTY.el9.aarch64.rpm) |
| `pgedge-18` | `18.4` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.2 MiB | [pgedge-18-18.4-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-18-18.4-1PIGSTY.el10.x86_64.rpm) |
| `pgedge-18` | `18.4` | [el10.aarch64](/os/el10.aarch64) | pigsty | 12.0 MiB | [pgedge-18-18.4-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-18-18.4-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-18` | `18.4` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.5 MiB | [pgedge-18_18.4-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-18` | `18.4` | [d12.aarch64](/os/d12.aarch64) | pigsty | 10.0 MiB | [pgedge-18_18.4-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-18` | `18.4` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.5 MiB | [pgedge-18_18.4-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~trixie_amd64.deb) |
| `pgedge-18` | `18.4` | [d13.aarch64](/os/d13.aarch64) | pigsty | 10.1 MiB | [pgedge-18_18.4-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~trixie_arm64.deb) |
| `pgedge-18` | `18.4` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.9 MiB | [pgedge-18_18.4-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~jammy_amd64.deb) |
| `pgedge-18` | `18.4` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.7 MiB | [pgedge-18_18.4-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~jammy_arm64.deb) |
| `pgedge-18` | `18.4` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.7 MiB | [pgedge-18_18.4-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~noble_amd64.deb) |
| `pgedge-18` | `18.4` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.6 MiB | [pgedge-18_18.4-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~noble_arm64.deb) |
| `pgedge-18` | `18.4` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.8 MiB | [pgedge-18_18.4-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~resolute_amd64.deb) |
| `pgedge-18` | `18.4` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.5 MiB | [pgedge-18_18.4-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-18/pgedge-18_18.4-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-17` | `17.10` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.8 MiB | [pgedge-17-17.10-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-17-17.10-1PIGSTY.el8.x86_64.rpm) |
| `pgedge-17` | `17.10` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.4 MiB | [pgedge-17-17.10-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-17-17.10-1PIGSTY.el8.aarch64.rpm) |
| `pgedge-17` | `17.10` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 MiB | [pgedge-17-17.10-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-17-17.10-1PIGSTY.el9.x86_64.rpm) |
| `pgedge-17` | `17.10` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 MiB | [pgedge-17-17.10-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-17-17.10-1PIGSTY.el9.aarch64.rpm) |
| `pgedge-17` | `17.10` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.8 MiB | [pgedge-17-17.10-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-17-17.10-1PIGSTY.el10.x86_64.rpm) |
| `pgedge-17` | `17.10` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.7 MiB | [pgedge-17-17.10-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-17-17.10-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-17` | `17.10` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.2 MiB | [pgedge-17_17.10-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-17` | `17.10` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [pgedge-17_17.10-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-17` | `17.10` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.2 MiB | [pgedge-17_17.10-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~trixie_amd64.deb) |
| `pgedge-17` | `17.10` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.7 MiB | [pgedge-17_17.10-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~trixie_arm64.deb) |
| `pgedge-17` | `17.10` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 MiB | [pgedge-17_17.10-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~jammy_amd64.deb) |
| `pgedge-17` | `17.10` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.4 MiB | [pgedge-17_17.10-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~jammy_arm64.deb) |
| `pgedge-17` | `17.10` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.4 MiB | [pgedge-17_17.10-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~noble_amd64.deb) |
| `pgedge-17` | `17.10` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 MiB | [pgedge-17_17.10-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~noble_arm64.deb) |
| `pgedge-17` | `17.10` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.4 MiB | [pgedge-17_17.10-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~resolute_amd64.deb) |
| `pgedge-17` | `17.10` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.2 MiB | [pgedge-17_17.10-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-17/pgedge-17_17.10-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-16` | `16.14` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.1 MiB | [pgedge-16-16.14-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-16-16.14-1PIGSTY.el8.x86_64.rpm) |
| `pgedge-16` | `16.14` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.7 MiB | [pgedge-16-16.14-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-16-16.14-1PIGSTY.el8.aarch64.rpm) |
| `pgedge-16` | `16.14` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.2 MiB | [pgedge-16-16.14-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-16-16.14-1PIGSTY.el9.x86_64.rpm) |
| `pgedge-16` | `16.14` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.0 MiB | [pgedge-16-16.14-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-16-16.14-1PIGSTY.el9.aarch64.rpm) |
| `pgedge-16` | `16.14` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.3 MiB | [pgedge-16-16.14-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-16-16.14-1PIGSTY.el10.x86_64.rpm) |
| `pgedge-16` | `16.14` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.1 MiB | [pgedge-16-16.14-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-16-16.14-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-16` | `16.14` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.7 MiB | [pgedge-16_16.14-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-16` | `16.14` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.3 MiB | [pgedge-16_16.14-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-16` | `16.14` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.7 MiB | [pgedge-16_16.14-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~trixie_amd64.deb) |
| `pgedge-16` | `16.14` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.3 MiB | [pgedge-16_16.14-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~trixie_arm64.deb) |
| `pgedge-16` | `16.14` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.0 MiB | [pgedge-16_16.14-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~jammy_amd64.deb) |
| `pgedge-16` | `16.14` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.8 MiB | [pgedge-16_16.14-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~jammy_arm64.deb) |
| `pgedge-16` | `16.14` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.9 MiB | [pgedge-16_16.14-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~noble_amd64.deb) |
| `pgedge-16` | `16.14` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.7 MiB | [pgedge-16_16.14-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~noble_arm64.deb) |
| `pgedge-16` | `16.14` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.9 MiB | [pgedge-16_16.14-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~resolute_amd64.deb) |
| `pgedge-16` | `16.14` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.7 MiB | [pgedge-16_16.14-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-16/pgedge-16_16.14-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}
{{< tab name="PG15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-15` | `15.18` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.8 MiB | [pgedge-15-15.18-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-15-15.18-1PIGSTY.el8.x86_64.rpm) |
| `pgedge-15` | `15.18` | [el8.aarch64](/os/el8.aarch64) | pigsty | 10.5 MiB | [pgedge-15-15.18-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-15-15.18-1PIGSTY.el8.aarch64.rpm) |
| `pgedge-15` | `15.18` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.3 MiB | [pgedge-15-15.18-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-15-15.18-1PIGSTY.el9.x86_64.rpm) |
| `pgedge-15` | `15.18` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.1 MiB | [pgedge-15-15.18-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-15-15.18-1PIGSTY.el9.aarch64.rpm) |
| `pgedge-15` | `15.18` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.4 MiB | [pgedge-15-15.18-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-15-15.18-1PIGSTY.el10.x86_64.rpm) |
| `pgedge-15` | `15.18` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.2 MiB | [pgedge-15-15.18-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-15-15.18-1PIGSTY.el10.aarch64.rpm) |
| `pgedge-15` | `15.18` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.7 MiB | [pgedge-15_15.18-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~bookworm_amd64.deb) |
| `pgedge-15` | `15.18` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.4 MiB | [pgedge-15_15.18-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~bookworm_arm64.deb) |
| `pgedge-15` | `15.18` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.8 MiB | [pgedge-15_15.18-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~trixie_amd64.deb) |
| `pgedge-15` | `15.18` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.4 MiB | [pgedge-15_15.18-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~trixie_arm64.deb) |
| `pgedge-15` | `15.18` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.1 MiB | [pgedge-15_15.18-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~jammy_amd64.deb) |
| `pgedge-15` | `15.18` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.9 MiB | [pgedge-15_15.18-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~jammy_arm64.deb) |
| `pgedge-15` | `15.18` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.0 MiB | [pgedge-15_15.18-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~noble_amd64.deb) |
| `pgedge-15` | `15.18` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.8 MiB | [pgedge-15_15.18-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~noble_arm64.deb) |
| `pgedge-15` | `15.18` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.0 MiB | [pgedge-15_15.18-1PIGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~resolute_amd64.deb) |
| `pgedge-15` | `15.18` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.8 MiB | [pgedge-15_15.18-1PIGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-15/pgedge-15_15.18-1PIGSTY~resolute_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/spock" title="Repository" icon="github" subtitle="github.com/pgEdge/spock" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="spock-5.0.10.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg spock;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install spock;		# install via package name, for the active PG version

pig install spock -v 18;   # install for PG 18
pig install spock -v 17;   # install for PG 17
pig install spock -v 16;   # install for PG 16
pig install spock -v 15;   # install for PG 15

```


[**Config**](https://ext.pgsty.com/usage/config/) this extension to [**`shared_preload_libraries`**](https://www.postgresql.org/docs/current/runtime-config-client.html#GUC-SHARED-PRELOAD-LIBRARIES):

```ini
shared_preload_libraries = 'spock';
```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION spock;
```




## Usage

Sources:

- [Spock v5.0.10 README](https://github.com/pgEdge/spock/blob/v5.0.10/README.md)
- [Getting started](https://github.com/pgEdge/spock/blob/v5.0.10/docs/getting_started.md)
- [Configuration reference](https://github.com/pgEdge/spock/blob/v5.0.10/docs/configuring.md)
- [Limitations](https://github.com/pgEdge/spock/blob/v5.0.10/docs/limitations.md)
- [Release notes](https://github.com/pgEdge/spock/blob/v5.0.10/docs/spock_release_notes.md)

`spock` provides active-active logical replication for PostgreSQL 15 through 18. Each participating database is a Spock node; a multi-master topology is formed by creating directed subscriptions between nodes.

### Configuration

In `postgresql.conf`:

```ini
wal_level = 'logical'
max_worker_processes = 10
max_replication_slots = 10
max_wal_senders = 10
shared_preload_libraries = 'spock'
track_commit_timestamp = on
spock.enable_ddl_replication = on
spock.include_ddl_repset = on
```

### Enabling

```sql
CREATE EXTENSION spock;
```

### Creating Nodes

On each node, create a node identity:

```sql
-- Node 1
SELECT spock.node_create(
    node_name := 'n1',
    dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);

-- Node 2
SELECT spock.node_create(
    node_name := 'n2',
    dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);
```

### Creating Subscriptions

For multi-master, each node subscribes to every other node:

```sql
-- On n1: subscribe to n2
SELECT spock.sub_create(
    subscription_name := 'sub_n1n2',
    provider_dsn := 'host=10.0.0.7 port=5432 dbname=mydb'
);

-- On n2: subscribe to n1
SELECT spock.sub_create(
    subscription_name := 'sub_n2n1',
    provider_dsn := 'host=10.0.0.5 port=5432 dbname=mydb'
);
```

### Replication Set Management

```sql
-- Add table to replication
SELECT spock.repset_add_table('default', 'my_table');

-- Remove table from replication
SELECT spock.repset_remove_table('default', 'my_table');

-- Add all tables in a schema
SELECT spock.repset_add_all_tables('default', '{public}');
```

### Key Features

- Multi-master (active-active) replication
- Automatic DDL replication
- Conflict detection and resolution using commit timestamps
- Row and column filtering
- Supports PostgreSQL 15, 16, 17, and 18
- Tables must have primary keys and matching schemas across nodes

### Operations and Caveats

- Install `spock` and add it to `shared_preload_libraries` on every participating server before creating nodes or subscriptions.
- Keep table definitions, data types, primary keys, and relevant unique indexes identical across nodes. Coordinate DDL even when DDL replication is enabled.
- Replicated tables need a primary key or another usable replica identity. Temporary and unlogged tables are not replication targets.
- Spock operates per database. Repeat extension and topology setup for each database that participates.
- Active-active conflict handling depends on commit timestamps and policy. Test simultaneous inserts and updates, especially nullable unique keys, before production use.
- Upstream documents platform/build requirements in the README; verify that the PostgreSQL build and Spock package used on every node are compatible.

### Version 5.0.10

`5.0.10` is a patch release in the 5.0 line. Its release notes include fixes for unique indexes containing `NULL`, `NULLS NOT DISTINCT` conflict handling, refreshing cached index metadata after an index is dropped, exception-path memory handling, and numerical version checks used during rolling patch upgrades. Upgrade every node to a compatible patch level and validate subscriptions after the rolling change.
