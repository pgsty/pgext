---
title: "stannum"
linkTitle: "stannum"
description: "Experimental TINQL text search with BM25 and segmented indexes"
weight: 2320
categories: ["FTS"]
languages: ["Rust"]
licenses: ["AGPL-3.0-or-later"]
repos: ["PIGSTY"]
page_width: full
---

[**stannum**](https://github.com/TeamSpringbird/stannum) : Experimental TINQL text search with BM25 and segmented indexes


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **2320** | {{< badge content="stannum" link="https://github.com/TeamSpringbird/stannum" >}} | {{< ext "stannum" >}} | `0.1.0` | {{< category "FTS" >}} | {{< license "AGPL-3.0-or-later" >}} | {{< language "Rust" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `stannum` |
|   **See Also**    | {{< ext "tin" >}} {{< ext "pg_search" >}} |

> [!Note] Unreleased 0.1.0-dev snapshot; no version release/tag. Conflicts with tin in the same database. Optional preload enables maintenance workers and standby index reads; preload must match on primary and standbys.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "red" >}} {{< bg "15" "" "red" >}} {{< bg "14" "" "red" >}} | `stannum` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "stannum_18" "green" >}} {{< bg "17" "stannum_17" "green" >}} {{< bg "16" "stannum_16" "red" >}} {{< bg "15" "stannum_15" "red" >}} {{< bg "14" "stannum_14" "red" >}} | `stannum_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1.0` | {{< bg "18" "postgresql-18-stannum" "green" >}} {{< bg "17" "postgresql-17-stannum" "green" >}} {{< bg "16" "postgresql-16-stannum" "red" >}} {{< bg "15" "postgresql-15-stannum" "red" >}} {{< bg "14" "postgresql-14-stannum" "red" >}} | `postgresql-$v-stannum` | - |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "stannum_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "stannum_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "stannum_16 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_15 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_14 : N/A 0" "gray" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "stannum_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "stannum_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "stannum_16 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_15 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_14 : N/A 0" "gray" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "stannum_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "stannum_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "stannum_16 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_15 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_14 : N/A 0" "gray" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "stannum_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "stannum_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "stannum_16 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_15 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_14 : N/A 0" "gray" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "stannum_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "stannum_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "stannum_16 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_15 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_14 : N/A 0" "gray" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "stannum_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "stannum_17 : AVAIL 1" "green" >}} | {{< bg "N/A" "stannum_16 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_15 : N/A 0" "gray" >}} | {{< bg "N/A" "stannum_14 : N/A 0" "gray" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-18-stannum : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1.0" "postgresql-17-stannum : AVAIL 1" "green" >}} | {{< bg "N/A" "postgresql-16-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-15-stannum : N/A 0" "gray" >}} | {{< bg "N/A" "postgresql-14-stannum : N/A 0" "gray" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `stannum_18` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [stannum_18-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/stannum_18-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `stannum_18` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.1 MiB | [stannum_18-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/stannum_18-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `stannum_18` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [stannum_18-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/stannum_18-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `stannum_18` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [stannum_18-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/stannum_18-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `stannum_18` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [stannum_18-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/stannum_18-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `stannum_18` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.3 MiB | [stannum_18-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/stannum_18-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-stannum` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 2.9 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.0 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.2 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-stannum` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.0 MiB | [postgresql-18-stannum_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/stannum/postgresql-18-stannum_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `stannum_17` | `0.1.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 3.5 MiB | [stannum_17-0.1.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/stannum_17-0.1.0-1PGSTY.el8.x86_64.rpm) |
| `stannum_17` | `0.1.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 3.1 MiB | [stannum_17-0.1.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/stannum_17-0.1.0-1PGSTY.el8.aarch64.rpm) |
| `stannum_17` | `0.1.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 3.5 MiB | [stannum_17-0.1.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/stannum_17-0.1.0-1PGSTY.el9.x86_64.rpm) |
| `stannum_17` | `0.1.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 3.3 MiB | [stannum_17-0.1.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/stannum_17-0.1.0-1PGSTY.el9.aarch64.rpm) |
| `stannum_17` | `0.1.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 3.5 MiB | [stannum_17-0.1.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/stannum_17-0.1.0-1PGSTY.el10.x86_64.rpm) |
| `stannum_17` | `0.1.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 3.2 MiB | [stannum_17-0.1.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/stannum_17-0.1.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-stannum` | `0.1.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 3.0 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 2.5 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 3.0 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 2.5 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 3.3 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 3.0 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 3.3 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 3.0 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 3.2 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-stannum` | `0.1.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 3.0 MiB | [postgresql-17-stannum_0.1.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/s/stannum/postgresql-17-stannum_0.1.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/TeamSpringbird/stannum" title="Repository" icon="github" subtitle="github.com/TeamSpringbird/stannum" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="stannum-0.1.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg stannum;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install stannum;		# install via package name, for the active PG version

pig install stannum -v 18;   # install for PG 18
pig install stannum -v 17;   # install for PG 17

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION stannum;
```

## Usage

Sources:

- [README](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/README.md)
- [Control file](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/postgres/stannum.control)
- [Cargo.toml](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/Cargo.toml)
- [postgres/sql/stannum--0.1.0.sql](https://github.com/TeamSpringbird/stannum/blob/e163585cb9d6b6f78b8de067a9c4aa33ea238063/postgres/sql/stannum--0.1.0.sql)

`stannum` is an experimental text-search extension derived from Lead, with persistent indexes, Boolean and phrase queries, and BM25 scoring. The reviewed 0.1.0 source has its own extension identity and storage format.

### Core Workflow

```sql
CREATE EXTENSION stannum;
CREATE TABLE documents (id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY, body text);
INSERT INTO documents (body) VALUES
  ('PostgreSQL supports full text search'), ('Search with exact phrase matching');
CREATE INDEX documents_search ON documents USING stannum (body);
ANALYZE documents;
SELECT id, stannum.full_score(ctid) AS score
FROM documents WHERE body ==> 'search' ORDER BY score DESC LIMIT 10;
SELECT * FROM stannum.segment_info('documents_search');
SELECT * FROM stannum.verify_index('documents_search', heap_check => true);
```

### Operation and Compatibility

The `==>` operator accepts TINQL. `stannum.highlight` marks matching text; `stannum.segment_info` inspects index segments and `stannum.verify_index` checks their consistency, optionally against the heap. Index maintenance follows PostgreSQL writes and vacuum.

PostgreSQL 17 and 18 are build targets, with most current lifecycle evidence on 18. A superuser installs the extension in its fixed `stannum` schema. Primary-server use loads on demand. For standby index reads, add the library to `shared_preload_libraries` on both primary and standbys and restart those servers.

### Migration Boundary

Do not install `tin` and `stannum` in the same database: their global `==>` operator conflicts. There is no supported in-place migration from the earlier extension names or storage layouts. Treat this as experimental software and validate recovery and maintenance on the intended workload.
