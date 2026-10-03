---
title: "lolor"
linkTitle: "lolor"
description: "Logical-replication-friendly replacement for PostgreSQL large objects"
weight: 9580
categories: ["ETL"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
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
|   **See Also**    | {{< ext "lo" >}} {{< ext "pglogical" >}} {{< ext "spock" >}} {{< ext "mimeo" >}} {{< ext "pgl_ddl_deploy" >}} {{< ext "logical_ddl" >}} {{< ext "pg_surgery" >}} {{< ext "pg_repack" >}} |

> [!Note] works on pgedge kernel fork. Requires lolor.node


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.2.2` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "red" >}} | `lolor` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.6` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `18.6` | {{< bg "18" "pgedge-18" "green" >}} {{< bg "17" "pgedge-17" "green" >}} {{< bg "16" "pgedge-16" "green" >}} {{< bg "15" "pgedge-15" "green" >}} {{< bg "14" "pgedge-14" "red" >}} | `pgedge-$v` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 18.6" "pgedge-18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 17.11" "pgedge-17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 16.15" "pgedge-16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 15.19" "pgedge-15 : AVAIL 1" "green" >}} | {{< bg "N/A" "pgedge-14 : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-18` | `18.6` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.6 MiB | [pgedge-18-18.6-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-18-18.6-2PGSTY.el8.x86_64.rpm) |
| `pgedge-18` | `18.6` | [el8.aarch64](/os/el8.aarch64) | pigsty | 12.2 MiB | [pgedge-18-18.6-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-18-18.6-2PGSTY.el8.aarch64.rpm) |
| `pgedge-18` | `18.6` | [el9.x86_64](/os/el9.x86_64) | pigsty | 12.0 MiB | [pgedge-18-18.6-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-18-18.6-2PGSTY.el9.x86_64.rpm) |
| `pgedge-18` | `18.6` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.7 MiB | [pgedge-18-18.6-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-18-18.6-2PGSTY.el9.aarch64.rpm) |
| `pgedge-18` | `18.6` | [el10.x86_64](/os/el10.x86_64) | pigsty | 12.1 MiB | [pgedge-18-18.6-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-18-18.6-2PGSTY.el10.x86_64.rpm) |
| `pgedge-18` | `18.6` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.9 MiB | [pgedge-18-18.6-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-18-18.6-2PGSTY.el10.aarch64.rpm) |
| `pgedge-18` | `18.6` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.3 MiB | [pgedge-18_18.6-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~bookworm_amd64.deb) |
| `pgedge-18` | `18.6` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.7 MiB | [pgedge-18_18.6-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~bookworm_arm64.deb) |
| `pgedge-18` | `18.6` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.3 MiB | [pgedge-18_18.6-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~trixie_amd64.deb) |
| `pgedge-18` | `18.6` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.8 MiB | [pgedge-18_18.6-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~trixie_arm64.deb) |
| `pgedge-18` | `18.6` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.6 MiB | [pgedge-18_18.6-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~jammy_amd64.deb) |
| `pgedge-18` | `18.6` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.4 MiB | [pgedge-18_18.6-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~jammy_arm64.deb) |
| `pgedge-18` | `18.6` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.4 MiB | [pgedge-18_18.6-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~noble_amd64.deb) |
| `pgedge-18` | `18.6` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.3 MiB | [pgedge-18_18.6-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~noble_arm64.deb) |
| `pgedge-18` | `18.6` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.5 MiB | [pgedge-18_18.6-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~resolute_amd64.deb) |
| `pgedge-18` | `18.6` | [u26.aarch64](/os/u26.aarch64) | pigsty | 11.3 MiB | [pgedge-18_18.6-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-18/pgedge-18_18.6-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-17` | `17.11` | [el8.x86_64](/os/el8.x86_64) | pigsty | 12.2 MiB | [pgedge-17-17.11-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-17-17.11-2PGSTY.el8.x86_64.rpm) |
| `pgedge-17` | `17.11` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.8 MiB | [pgedge-17-17.11-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-17-17.11-2PGSTY.el8.aarch64.rpm) |
| `pgedge-17` | `17.11` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.7 MiB | [pgedge-17-17.11-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-17-17.11-2PGSTY.el9.x86_64.rpm) |
| `pgedge-17` | `17.11` | [el9.aarch64](/os/el9.aarch64) | pigsty | 11.5 MiB | [pgedge-17-17.11-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-17-17.11-2PGSTY.el9.aarch64.rpm) |
| `pgedge-17` | `17.11` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.8 MiB | [pgedge-17-17.11-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-17-17.11-2PGSTY.el10.x86_64.rpm) |
| `pgedge-17` | `17.11` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.6 MiB | [pgedge-17-17.11-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-17-17.11-2PGSTY.el10.aarch64.rpm) |
| `pgedge-17` | `17.11` | [d12.x86_64](/os/d12.x86_64) | pigsty | 10.0 MiB | [pgedge-17_17.11-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~bookworm_amd64.deb) |
| `pgedge-17` | `17.11` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.5 MiB | [pgedge-17_17.11-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~bookworm_arm64.deb) |
| `pgedge-17` | `17.11` | [d13.x86_64](/os/d13.x86_64) | pigsty | 10.0 MiB | [pgedge-17_17.11-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~trixie_amd64.deb) |
| `pgedge-17` | `17.11` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.5 MiB | [pgedge-17_17.11-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~trixie_arm64.deb) |
| `pgedge-17` | `17.11` | [u22.x86_64](/os/u22.x86_64) | pigsty | 11.3 MiB | [pgedge-17_17.11-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~jammy_amd64.deb) |
| `pgedge-17` | `17.11` | [u22.aarch64](/os/u22.aarch64) | pigsty | 11.1 MiB | [pgedge-17_17.11-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~jammy_arm64.deb) |
| `pgedge-17` | `17.11` | [u24.x86_64](/os/u24.x86_64) | pigsty | 11.2 MiB | [pgedge-17_17.11-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~noble_amd64.deb) |
| `pgedge-17` | `17.11` | [u24.aarch64](/os/u24.aarch64) | pigsty | 11.0 MiB | [pgedge-17_17.11-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~noble_arm64.deb) |
| `pgedge-17` | `17.11` | [u26.x86_64](/os/u26.x86_64) | pigsty | 11.2 MiB | [pgedge-17_17.11-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~resolute_amd64.deb) |
| `pgedge-17` | `17.11` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.9 MiB | [pgedge-17_17.11-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-17/pgedge-17_17.11-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-16` | `16.15` | [el8.x86_64](/os/el8.x86_64) | pigsty | 11.5 MiB | [pgedge-16-16.15-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-16-16.15-2PGSTY.el8.x86_64.rpm) |
| `pgedge-16` | `16.15` | [el8.aarch64](/os/el8.aarch64) | pigsty | 11.1 MiB | [pgedge-16-16.15-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-16-16.15-2PGSTY.el8.aarch64.rpm) |
| `pgedge-16` | `16.15` | [el9.x86_64](/os/el9.x86_64) | pigsty | 11.2 MiB | [pgedge-16-16.15-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-16-16.15-2PGSTY.el9.x86_64.rpm) |
| `pgedge-16` | `16.15` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.9 MiB | [pgedge-16-16.15-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-16-16.15-2PGSTY.el9.aarch64.rpm) |
| `pgedge-16` | `16.15` | [el10.x86_64](/os/el10.x86_64) | pigsty | 11.3 MiB | [pgedge-16-16.15-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-16-16.15-2PGSTY.el10.x86_64.rpm) |
| `pgedge-16` | `16.15` | [el10.aarch64](/os/el10.aarch64) | pigsty | 11.1 MiB | [pgedge-16-16.15-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-16-16.15-2PGSTY.el10.aarch64.rpm) |
| `pgedge-16` | `16.15` | [d12.x86_64](/os/d12.x86_64) | pigsty | 9.5 MiB | [pgedge-16_16.15-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~bookworm_amd64.deb) |
| `pgedge-16` | `16.15` | [d12.aarch64](/os/d12.aarch64) | pigsty | 9.0 MiB | [pgedge-16_16.15-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~bookworm_arm64.deb) |
| `pgedge-16` | `16.15` | [d13.x86_64](/os/d13.x86_64) | pigsty | 9.5 MiB | [pgedge-16_16.15-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~trixie_amd64.deb) |
| `pgedge-16` | `16.15` | [d13.aarch64](/os/d13.aarch64) | pigsty | 9.1 MiB | [pgedge-16_16.15-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~trixie_arm64.deb) |
| `pgedge-16` | `16.15` | [u22.x86_64](/os/u22.x86_64) | pigsty | 10.8 MiB | [pgedge-16_16.15-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~jammy_amd64.deb) |
| `pgedge-16` | `16.15` | [u22.aarch64](/os/u22.aarch64) | pigsty | 10.6 MiB | [pgedge-16_16.15-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~jammy_arm64.deb) |
| `pgedge-16` | `16.15` | [u24.x86_64](/os/u24.x86_64) | pigsty | 10.7 MiB | [pgedge-16_16.15-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~noble_amd64.deb) |
| `pgedge-16` | `16.15` | [u24.aarch64](/os/u24.aarch64) | pigsty | 10.5 MiB | [pgedge-16_16.15-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~noble_arm64.deb) |
| `pgedge-16` | `16.15` | [u26.x86_64](/os/u26.x86_64) | pigsty | 10.7 MiB | [pgedge-16_16.15-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~resolute_amd64.deb) |
| `pgedge-16` | `16.15` | [u26.aarch64](/os/u26.aarch64) | pigsty | 10.4 MiB | [pgedge-16_16.15-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-16/pgedge-16_16.15-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pgedge-15` | `15.19` | [el8.x86_64](/os/el8.x86_64) | pigsty | 10.3 MiB | [pgedge-15-15.19-2PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pgedge-15-15.19-2PGSTY.el8.x86_64.rpm) |
| `pgedge-15` | `15.19` | [el8.aarch64](/os/el8.aarch64) | pigsty | 9.9 MiB | [pgedge-15-15.19-2PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pgedge-15-15.19-2PGSTY.el8.aarch64.rpm) |
| `pgedge-15` | `15.19` | [el9.x86_64](/os/el9.x86_64) | pigsty | 10.2 MiB | [pgedge-15-15.19-2PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pgedge-15-15.19-2PGSTY.el9.x86_64.rpm) |
| `pgedge-15` | `15.19` | [el9.aarch64](/os/el9.aarch64) | pigsty | 10.0 MiB | [pgedge-15-15.19-2PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pgedge-15-15.19-2PGSTY.el9.aarch64.rpm) |
| `pgedge-15` | `15.19` | [el10.x86_64](/os/el10.x86_64) | pigsty | 10.3 MiB | [pgedge-15-15.19-2PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pgedge-15-15.19-2PGSTY.el10.x86_64.rpm) |
| `pgedge-15` | `15.19` | [el10.aarch64](/os/el10.aarch64) | pigsty | 10.1 MiB | [pgedge-15-15.19-2PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pgedge-15-15.19-2PGSTY.el10.aarch64.rpm) |
| `pgedge-15` | `15.19` | [d12.x86_64](/os/d12.x86_64) | pigsty | 8.5 MiB | [pgedge-15_15.19-2PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~bookworm_amd64.deb) |
| `pgedge-15` | `15.19` | [d12.aarch64](/os/d12.aarch64) | pigsty | 8.2 MiB | [pgedge-15_15.19-2PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~bookworm_arm64.deb) |
| `pgedge-15` | `15.19` | [d13.x86_64](/os/d13.x86_64) | pigsty | 8.6 MiB | [pgedge-15_15.19-2PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~trixie_amd64.deb) |
| `pgedge-15` | `15.19` | [d13.aarch64](/os/d13.aarch64) | pigsty | 8.2 MiB | [pgedge-15_15.19-2PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~trixie_arm64.deb) |
| `pgedge-15` | `15.19` | [u22.x86_64](/os/u22.x86_64) | pigsty | 9.9 MiB | [pgedge-15_15.19-2PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~jammy_amd64.deb) |
| `pgedge-15` | `15.19` | [u22.aarch64](/os/u22.aarch64) | pigsty | 9.7 MiB | [pgedge-15_15.19-2PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~jammy_arm64.deb) |
| `pgedge-15` | `15.19` | [u24.x86_64](/os/u24.x86_64) | pigsty | 9.8 MiB | [pgedge-15_15.19-2PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~noble_amd64.deb) |
| `pgedge-15` | `15.19` | [u24.aarch64](/os/u24.aarch64) | pigsty | 9.6 MiB | [pgedge-15_15.19-2PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~noble_arm64.deb) |
| `pgedge-15` | `15.19` | [u26.x86_64](/os/u26.x86_64) | pigsty | 9.8 MiB | [pgedge-15_15.19-2PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~resolute_amd64.deb) |
| `pgedge-15` | `15.19` | [u26.aarch64](/os/u26.aarch64) | pigsty | 9.6 MiB | [pgedge-15_15.19-2PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/p/pgedge-15/pgedge-15_15.19-2PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/pgEdge/lolor" title="Repository" icon="github" subtitle="github.com/pgEdge/lolor" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="lolor-1.2.2.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg lolor;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

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

Sources:

- [v1.2.2 README](https://github.com/pgEdge/lolor/blob/v1.2.2/README.md)
- [Control file](https://github.com/pgEdge/lolor/blob/v1.2.2/lolor.control)
- [Version 1.2.2 migration](https://github.com/pgEdge/lolor/blob/v1.2.2/lolor--1.2.1--1.2.2.sql)
- [Usage reference](https://github.com/pgEdge/lolor/blob/v1.2.2/docs/using_lolor.md)
- [Release notes](https://github.com/pgEdge/lolor/blob/v1.2.2/docs/lolor_release_notes.md)
`lolor` 1.2.2 stores large-object chunks and metadata in ordinary tables so logical replication can include them. Enabling it replaces the database's native large-object routines; it is a database-wide behavior change, not an independent object store.

### Configure and Enable

Assign each writing node a distinct nonzero `lolor.node` before creating large objects. Upstream documents values from 1 through 2^28. Configure the server's parameter and create the extension in each participating database:

```conf
lolor.node = 1
```

```sql
CREATE EXTENSION lolor;
SET search_path = lolor, "$user", public, pg_catalog;
```

The control file fixes schema `lolor`, declares trusted installation and disallows relocation. The upstream README requires PostgreSQL 16 or newer; the current Pigsty pgEdge bundle has separately tested builds for 15–18. That packaging result does not establish support for arbitrary stock PostgreSQL 15 installations.

### Large Object Workflow

```sql
WITH created AS (
  SELECT lo_from_bytea(0, convert_to('example data', 'UTF8')) AS oid
)
SELECT oid, convert_from(lo_get(oid), 'UTF8') AS contents FROM created;
```

Standard calls such as `lo_create()`, `lo_get()`, `lo_put()` and `lo_unlink()` use the replacement routines. Descriptor-based access through `lo_open()`, `loread()`, `lowrite()` and `lo_close()` must stay within the same transaction. File import/export operates on server-side paths and remains subject to the relevant function and filesystem privileges.

The extension uses `lolor.pg_largeobject` and `lolor.pg_largeobject_metadata`. Existing catalog routines are retained with renamed originals so disabling or removing the extension can restore the native function names.

### Replication

With a configured Spock replication set, include both tables:

```sql
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject');
SELECT spock.repset_add_table('default', 'lolor.pg_largeobject_metadata');
```

Install compatible versions and coordinate node identifiers and object ownership on every node. Creating the extension alone does not create a logical-replication topology.

### Upgrade and Boundaries

Version 1.2.2 repairs extension and major-version upgrade behavior and adds `lolor.disable()`, `lolor.enable()` and `lolor.is_enabled()`. Update lolor **before** running pg_upgrade; the fix does not retroactively repair a major upgrade already attempted with older extension files:

```sql
ALTER EXTENSION lolor UPDATE TO '1.2.2';
SELECT lolor.is_enabled();
```

Disabling changes which native function names are active; it does not migrate stored large objects. Native large-object migration is not provided. Native large-object functionality and lolor storage cannot be used interchangeably while lolor is enabled. Upstream excludes ALTER LARGE OBJECT, GRANT ON LARGE OBJECT, COMMENT ON LARGE OBJECT and REVOKE ON LARGE OBJECT. Plan backup, restore and replication for both ordinary tables.
