---
title: "macavity"
linkTitle: "macavity"
description: "Deterministic session-local fault injection for PostgreSQL test clusters"
weight: 5215
categories: ["ADMIN"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**macavity**](https://github.com/CrystallineCore/Macavity) : Deterministic session-local fault injection for PostgreSQL test clusters


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **5215** | {{< badge content="macavity" link="https://github.com/CrystallineCore/Macavity" >}} | {{< ext "macavity" >}} | `0.2.0` | {{< category "ADMIN" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |

> [!Note] Testing release; destructive crash action affects the entire instance. Disposable test clusters only.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `macavity` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "macavity_18" "green" >}} {{< bg "17" "macavity_17" "green" >}} {{< bg "16" "macavity_16" "green" >}} {{< bg "15" "macavity_15" "red" >}} {{< bg "14" "macavity_14" "red" >}} | `macavity_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-macavity" "green" >}} {{< bg "17" "postgresql-17-macavity" "green" >}} {{< bg "16" "postgresql-16-macavity" "green" >}} {{< bg "15" "postgresql-15-macavity" "red" >}} {{< bg "14" "postgresql-14-macavity" "red" >}} | `postgresql-$v-macavity` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "macavity_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "macavity_15 : N/A 0" "gray" >}} | {{< bg "N/A" "macavity_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "macavity_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "macavity_15 : N/A 0" "gray" >}} | {{< bg "N/A" "macavity_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "macavity_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "macavity_15 : N/A 0" "gray" >}} | {{< bg "N/A" "macavity_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "macavity_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "macavity_15 : N/A 0" "gray" >}} | {{< bg "N/A" "macavity_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "macavity_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "macavity_15 : N/A 0" "gray" >}} | {{< bg "N/A" "macavity_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "macavity_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "macavity_16 : AVAIL 1" "green" >}} | {{< bg "N/A" "macavity_15 : N/A 0" "gray" >}} | {{< bg "N/A" "macavity_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-18-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-macavity : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-macavity : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-15-macavity : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-macavity : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `macavity_18` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 47.8 KiB | [macavity_18-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/macavity_18-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `macavity_18` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 47.2 KiB | [macavity_18-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/macavity_18-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `macavity_18` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.8 KiB | [macavity_18-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/macavity_18-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `macavity_18` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.3 KiB | [macavity_18-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/macavity_18-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `macavity_18` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.8 KiB | [macavity_18-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/macavity_18-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `macavity_18` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [macavity_18-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/macavity_18-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-macavity` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.3 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.4 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 40.1 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 39.6 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.3 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.4 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-macavity` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.3 KiB | [postgresql-18-macavity_0.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/macavity/postgresql-18-macavity_0.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `macavity_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 47.8 KiB | [macavity_17-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/macavity_17-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `macavity_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 47.2 KiB | [macavity_17-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/macavity_17-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `macavity_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.7 KiB | [macavity_17-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/macavity_17-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `macavity_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.3 KiB | [macavity_17-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/macavity_17-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `macavity_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.7 KiB | [macavity_17-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/macavity_17-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `macavity_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [macavity_17-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/macavity_17-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-macavity` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.3 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.4 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.9 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.4 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.2 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.4 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-macavity` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.1 KiB | [postgresql-17-macavity_0.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/macavity/postgresql-17-macavity_0.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `macavity_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 47.8 KiB | [macavity_16-0.2.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/macavity_16-0.2.0-1PGSTY.el8.x86_64.rpm) |
| `macavity_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 47.2 KiB | [macavity_16-0.2.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/macavity_16-0.2.0-1PGSTY.el8.aarch64.rpm) |
| `macavity_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 47.7 KiB | [macavity_16-0.2.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/macavity_16-0.2.0-1PGSTY.el9.x86_64.rpm) |
| `macavity_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 47.3 KiB | [macavity_16-0.2.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/macavity_16-0.2.0-1PGSTY.el9.aarch64.rpm) |
| `macavity_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 47.7 KiB | [macavity_16-0.2.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/macavity_16-0.2.0-1PGSTY.el10.x86_64.rpm) |
| `macavity_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 47.6 KiB | [macavity_16-0.2.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/macavity_16-0.2.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-macavity` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 40.6 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 40.3 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 40.6 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 40.4 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 44.7 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 44.3 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 39.5 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 39.2 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 39.4 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-macavity` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 39.1 KiB | [postgresql-16-macavity_0.2.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/m/macavity/postgresql-16-macavity_0.2.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/CrystallineCore/Macavity" title="Repository" icon="github" subtitle="github.com/CrystallineCore/Macavity" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="macavity-0.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg macavity;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install macavity;		# install via package name, for the active PG version

pig install macavity -v 18;   # install for PG 18
pig install macavity -v 17;   # install for PG 17
pig install macavity -v 16;   # install for PG 16

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION macavity;
```

## Usage

Sources:

- [README.md](https://github.com/CrystallineCore/Macavity/blob/v0.2.0/README.md)
- [CHANGELOG.md](https://github.com/CrystallineCore/Macavity/blob/v0.2.0/CHANGELOG.md)
- [macavity.control](https://github.com/CrystallineCore/Macavity/blob/v0.2.0/macavity.control)
- [sql/macavity--0.2.0.sql](https://github.com/CrystallineCore/Macavity/blob/v0.2.0/sql/macavity--0.2.0.sql)
- [sql/macavity--0.1.0--0.2.0.sql](https://github.com/CrystallineCore/Macavity/blob/v0.2.0/sql/macavity--0.1.0--0.2.0.sql)

`macavity` 0.2.0 provides deterministic fault injection for disposable PostgreSQL test clusters. Upstream tests PostgreSQL 16–18 on Linux; PostgreSQL 19+ and other platforms are not validated, and Windows crash injection is unsupported. The crash action can disconnect every session and force crash recovery.

### Event workflow

```sql
CREATE EXTENSION macavity;
SELECT * FROM macavity_points();
SELECT macavity_arm('executor_start', 'error', 1);
SELECT 1; -- expected injected error
SELECT * FROM macavity_status();
SELECT macavity_disarm();
SELECT macavity_reset();
```

### Registry and upgrade

Each session keeps multiple events with stable IDs, separate counters and states `armed`, `completed` or `disarmed`. `macavity_arm` returns an integer event ID; its single-ID overload re-arms a completed or disarmed event. `macavity_status` returns zero or more event rows. `macavity_disarm` retains event history; `macavity_reset` clears it and restarts IDs.

Points are `executor_start`, `executor_end`, `before_commit` and `before_abort`; actions are `error`, `delay` and `crash`. Delay lasts one second. Abort-time error injection is rejected. At a shared hit, delay precedes crash, then error, with IDs breaking ties.

Creation requires a superuser; no shared preload is needed. Only point enumeration is granted to `PUBLIC`. Upgrading with `ALTER EXTENSION macavity UPDATE` recreates functions with changed signatures: restore explicit grants, update callers and reconnect old sessions. Current Pigsty package fields still refer to 0.1.0.
