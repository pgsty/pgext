---
title: "ddsketch"
linkTitle: "ddsketch"
description: "Provides ddsketch aggregate function"
weight: 4650
categories: ["FUNC"]
languages: ["C"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**ddsketch**](https://github.com/tvondra/ddsketch) : Provides ddsketch aggregate function


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4650** | {{< badge content="ddsketch" link="https://github.com/tvondra/ddsketch" >}} | {{< ext "ddsketch" >}} | `1.0.1` | {{< category "FUNC" >}} | {{< license "PostgreSQL" >}} | {{< language "C" >}} |
{.wide .expand}


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |
{.wide .expand}


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "tdigest" >}} {{< ext "count_distinct" >}} {{< ext "topn" >}} {{< ext "omnisketch" >}} {{< ext "datasketches" >}} {{< ext "hll" >}} {{< ext "quantile" >}} {{< ext "lower_quantile" >}} {{< ext "weighted_statistics" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "19" "" "green" >}} {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} | `ddsketch` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "19" "ddsketch_19" "green" >}} {{< bg "18" "ddsketch_18" "green" >}} {{< bg "17" "ddsketch_17" "green" >}} {{< bg "16" "ddsketch_16" "green" >}} {{< bg "15" "ddsketch_15" "green" >}} | `ddsketch_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.1` | {{< bg "19" "postgresql-19-ddsketch" "green" >}} {{< bg "18" "postgresql-18-ddsketch" "green" >}} {{< bg "17" "postgresql-17-ddsketch" "green" >}} {{< bg "16" "postgresql-16-ddsketch" "green" >}} {{< bg "15" "postgresql-15-ddsketch" "green" >}} | `postgresql-$v-ddsketch` | - |
{.packages .wide .expand}


| **Linux** / **PG** |                  **PG19**                   |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "ddsketch_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "ddsketch_19 : N/A 0" "gray" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_19 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "ddsketch_15 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-19-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-18-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-17-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-16-ddsketch : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.1" "postgresql-15-ddsketch : AVAIL 1" "green" >}} |
{.matrix .wide .expand}


{{< tabs group="pgmajor" >}}
{{< tab label="PG19" value="pg19" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_19` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 71.1 KiB | [ddsketch_19-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_19-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `ddsketch_19` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.3 KiB | [ddsketch_19-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_19-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `ddsketch_19` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.0 KiB | [ddsketch_19-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_19-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `ddsketch_19` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 70.4 KiB | [ddsketch_19-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_19-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-19-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.8 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.5 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.7 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.8 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.9 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.7 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.6 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.2 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 62.9 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-19-ddsketch` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 63.4 KiB | [postgresql-19-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-19-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_18` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.5 KiB | [ddsketch_18-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_18-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `ddsketch_18` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.0 KiB | [ddsketch_18-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_18-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `ddsketch_18` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 71.1 KiB | [ddsketch_18-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_18-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `ddsketch_18` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.3 KiB | [ddsketch_18-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_18-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `ddsketch_18` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.0 KiB | [ddsketch_18-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_18-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `ddsketch_18` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 70.5 KiB | [ddsketch_18-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_18-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.8 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.5 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.7 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.8 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 64.9 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.7 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.6 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.1 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 63.0 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-ddsketch` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 63.4 KiB | [postgresql-18-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-18-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_17` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.8 KiB | [ddsketch_17-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_17-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `ddsketch_17` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.1 KiB | [ddsketch_17-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_17-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `ddsketch_17` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 71.1 KiB | [ddsketch_17-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_17-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `ddsketch_17` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.3 KiB | [ddsketch_17-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_17-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `ddsketch_17` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.3 KiB | [ddsketch_17-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_17-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `ddsketch_17` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 70.4 KiB | [ddsketch_17-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_17-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.8 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.5 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.6 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.8 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.8 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.4 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.7 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.2 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 63.2 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-ddsketch` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 63.6 KiB | [postgresql-17-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-17-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_16` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.8 KiB | [ddsketch_16-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_16-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `ddsketch_16` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.1 KiB | [ddsketch_16-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_16-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `ddsketch_16` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 71.1 KiB | [ddsketch_16-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_16-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `ddsketch_16` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.2 KiB | [ddsketch_16-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_16-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `ddsketch_16` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.3 KiB | [ddsketch_16-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_16-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `ddsketch_16` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 70.4 KiB | [ddsketch_16-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_16-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.8 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.5 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.6 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.8 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.8 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.3 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.7 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.2 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 63.2 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-ddsketch` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 63.6 KiB | [postgresql-16-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-16-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `ddsketch_15` | `1.0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 68.7 KiB | [ddsketch_15-1.0.1-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/ddsketch_15-1.0.1-1PGSTY.el8.x86_64.rpm) |
| `ddsketch_15` | `1.0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 68.0 KiB | [ddsketch_15-1.0.1-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/ddsketch_15-1.0.1-1PGSTY.el8.aarch64.rpm) |
| `ddsketch_15` | `1.0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 71.1 KiB | [ddsketch_15-1.0.1-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/ddsketch_15-1.0.1-1PGSTY.el9.x86_64.rpm) |
| `ddsketch_15` | `1.0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 70.2 KiB | [ddsketch_15-1.0.1-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/ddsketch_15-1.0.1-1PGSTY.el9.aarch64.rpm) |
| `ddsketch_15` | `1.0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 70.2 KiB | [ddsketch_15-1.0.1-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/ddsketch_15-1.0.1-1PGSTY.el10.x86_64.rpm) |
| `ddsketch_15` | `1.0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 70.4 KiB | [ddsketch_15-1.0.1-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/ddsketch_15-1.0.1-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-ddsketch` | `1.0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 60.7 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 60.5 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 60.5 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 60.8 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.8 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 69.5 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 63.6 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~noble_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.1 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~noble_arm64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u26.x86_64](/os/u26.x86_64) | pigsty | 63.5 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-ddsketch` | `1.0.1` | [u26.aarch64](/os/u26.aarch64) | pigsty | 63.6 KiB | [postgresql-15-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/d/ddsketch/postgresql-15-ddsketch_1.0.1-1PGSTY~resolute_arm64.deb) |
{.downloads .wide .expand}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/tvondra/ddsketch" title="Repository" icon="github" subtitle="github.com/tvondra/ddsketch" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="ddsketch-1.0.1.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg ddsketch;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install ddsketch;		# install via package name, for the active PG version

pig install ddsketch -v 19;   # install for PG 19
pig install ddsketch -v 18;   # install for PG 18
pig install ddsketch -v 17;   # install for PG 17
pig install ddsketch -v 16;   # install for PG 16
pig install ddsketch -v 15;   # install for PG 15

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION ddsketch;
```




## Usage

> [ddsketch: DDSketch quantile estimation for PostgreSQL](https://github.com/tvondra/ddsketch)

Implements DDSketch, a fully-mergeable quantile sketch with relative-error guarantees. Much faster than `percentile_cont` and supports parallelism.

```sql
CREATE EXTENSION ddsketch;
```

### Direct Aggregation Functions

| Function | Description |
|---|---|
| `ddsketch_percentile(value, alpha, nbuckets, quantile)` | Estimate a single percentile |
| `ddsketch_percentile(value, alpha, nbuckets, quantiles[])` | Estimate multiple percentiles |
| `ddsketch_percentile_of(value, alpha, nbuckets, value)` | Estimate percentile rank of a value |
| `ddsketch_percentile_of(value, alpha, nbuckets, values[])` | Estimate percentile ranks of multiple values |

### Pre-aggregation Functions

| Function | Description |
|---|---|
| `ddsketch(value, alpha, nbuckets)` | Build a ddsketch from values |
| `ddsketch_percentile(sketch, quantile)` | Estimate percentile from a pre-built sketch |
| `ddsketch_percentile(sketch, quantiles[])` | Estimate multiple percentiles from a pre-built sketch |

### Utility Functions

| Function | Description |
|---|---|
| `ddsketch_count(sketch)` | Return the number of items in the sketch |
| `ddsketch_sum(sketch, low, high)` | Trimmed sum within a value range |
| `ddsketch_avg(sketch, low, high)` | Trimmed average within a value range |

### Parameters

- `alpha` -- controls accuracy and sketch size (lower = more accurate, larger)
- `nbuckets` -- maximum number of buckets (each 8 bytes)

### Examples

```sql
-- Instead of: SELECT percentile_cont(0.95) WITHIN GROUP (ORDER BY a) FROM t;
SELECT ddsketch_percentile(a, 0.05, 1024, 0.95) FROM t;

-- Multiple percentiles at once
SELECT ddsketch_percentile(a, 0.05, 1024, ARRAY[0.5, 0.95, 0.99]) FROM t;

-- Pre-aggregate for fast repeated queries
CREATE TABLE p AS SELECT a, b, ddsketch(c, 0.05, 1024) AS d FROM t GROUP BY a, b;

-- Query pre-aggregated data (~1.5ms vs ~7s for exact)
SELECT a, ddsketch_percentile(d, 0.95) FROM p GROUP BY a ORDER BY a;
```
