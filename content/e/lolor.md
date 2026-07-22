---
title: "lolor"
linkTitle: "lolor"
description: "Logical-replication-friendly replacement for PostgreSQL large objects"
weight: 9580
categories: ["ETL"]
width: full
---

[**lolor**](https://github.com/pgEdge/lolor) : Logical-replication-friendly replacement for PostgreSQL large objects


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **9580** | {{< badge content="lolor" link="https://github.com/pgEdge/lolor" >}} | {{< ext "lolor" >}} | `1.2.2` | {{< category "ETL" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dt-" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="yes" color="green" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `lolor` |
|   **See Also**    | {{< ext "spock" >}} {{< ext "snowflake" >}} |

> [!Note] works on pgedge kernel fork. Requires lolor.node


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `lolor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |


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
{{< card link="https://github.com/pgEdge/lolor" title="Repository" icon="github" subtitle="github.com/pgEdge/lolor" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="lolor-1.2.2.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg lolor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install lolor;		# install via package name, for the active PG version

pig install lolor -v 18;   # install for PG 18
pig install lolor -v 17;   # install for PG 17
pig install lolor -v 16;   # install for PG 16
pig install lolor -v 15;   # install for PG 15

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION lolor;
```




## Usage

> [lolor: Logical-replication-friendly replacement for PostgreSQL large objects](https://github.com/pgEdge/lolor)

Makes PostgreSQL large objects compatible with logical replication by storing them in non-catalog tables.

### Enabling

```sql
CREATE EXTENSION lolor;
```

Configure the node identifier in `postgresql.conf`:

```ini
lolor.node = 1  -- unique node ID (1 to 2^28)
```

Optionally adjust the search path:

```sql
SET search_path = lolor, "$user", public, pg_catalog;
```

### Large Object Operations

Once installed, the standard `lo_*` functions are redirected to use lolor's tables:

```sql
-- Create a large object
SELECT lo_create(0);

-- Import a file into a large object
SELECT lo_import('/path/to/file.bin');

-- Export a large object to a file
SELECT lo_export(oid, '/path/to/output.bin');

-- Open, read, write, seek, close
SELECT lo_open(oid, x'40000'::int);  -- INV_WRITE
SELECT lowrite(fd, 'data'::bytea);
SELECT loread(fd, 1024);
SELECT lo_close(fd);

-- Delete a large object
SELECT lo_unlink(oid);
```

### Replication Setup

Add lolor tables to your replication set:

```sql
-- For spock/pgedge replication
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject');
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject_metadata');
```

### Internal Tables

The extension manages large objects in:

- `lolor.pg_largeobject` - stores object data chunks
- `lolor.pg_largeobject_metadata` - stores object metadata

### Limitations

- Native PostgreSQL large object functionality cannot be used while lolor is active
- Migration of existing native large objects to lolor is not supported
- `ALTER LARGE OBJECT`, `GRANT ON LARGE OBJECT`, `COMMENT ON LARGE OBJECT`, and `REVOKE ON LARGE OBJECT` are not supported
- Requires PostgreSQL 16 or newer
