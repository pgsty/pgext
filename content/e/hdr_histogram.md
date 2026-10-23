---
title: "hdr_histogram"
linkTitle: "hdr_histogram"
description: "HDR histogram type, percentile queries and mergeable aggregation."
weight: 4715
categories: ["FUNC"]
languages: ["C"]
licenses: ["MIT"]
repos: ["PIGSTY"]
page_width: full
---

[**hdr_histogram**](https://github.com/vovcke/pg_hdr_histogram) : HDR histogram type, percentile queries and mergeable aggregation.


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **4715** | {{< badge content="hdr_histogram" link="https://github.com/vovcke/pg_hdr_histogram" >}} | {{< ext "hdr_histogram" >}} | `1.0` | {{< category "FUNC" >}} | {{< license "MIT" >}} | {{< language "C" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="--s-dtr" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="yes" color="green" >}} |

> [!Note] Official v1.0.0 release asset installs control/SQL 1.0; the Git tag declares 1.0.0. Extension license is MIT; packages also include the CC0-1.0 HdrHistogram_c library. PG14-18; requires zlib. Default lossless merge depends on accumulator range/order; left mode can discard samples.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `hdr_histogram` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "hdr_histogram_18" "green" >}} {{< bg "17" "hdr_histogram_17" "green" >}} {{< bg "16" "hdr_histogram_16" "green" >}} {{< bg "15" "hdr_histogram_15" "green" >}} {{< bg "14" "hdr_histogram_14" "green" >}} | `hdr_histogram_$v` | `zlib` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `1.0.0` | {{< bg "18" "postgresql-18-hdr-histogram" "green" >}} {{< bg "17" "postgresql-17-hdr-histogram" "green" >}} {{< bg "16" "postgresql-16-hdr-histogram" "green" >}} {{< bg "15" "postgresql-15-hdr-histogram" "green" >}} {{< bg "14" "postgresql-14-hdr-histogram" "green" >}} | `postgresql-$v-hdr-histogram` | `zlib1g` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "hdr_histogram_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-18-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-17-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-16-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-15-hdr-histogram : AVAIL 1" "green" >}} | {{< bg "PIGSTY 1.0.0" "postgresql-14-hdr-histogram : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG18" value="pg18" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdr_histogram_18` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 79.4 KiB | [hdr_histogram_18-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hdr_histogram_18-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `hdr_histogram_18` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.4 KiB | [hdr_histogram_18-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hdr_histogram_18-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `hdr_histogram_18` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 76.0 KiB | [hdr_histogram_18-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hdr_histogram_18-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `hdr_histogram_18` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 72.1 KiB | [hdr_histogram_18-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hdr_histogram_18-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `hdr_histogram_18` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.8 KiB | [hdr_histogram_18-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hdr_histogram_18-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `hdr_histogram_18` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 72.7 KiB | [hdr_histogram_18-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hdr_histogram_18-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.5 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.8 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.7 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.9 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 69.4 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 65.7 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.3 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.8 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.7 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-18-hdr-histogram` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.1 KiB | [postgresql-18-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-18-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdr_histogram_17` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 79.4 KiB | [hdr_histogram_17-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hdr_histogram_17-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `hdr_histogram_17` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.4 KiB | [hdr_histogram_17-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hdr_histogram_17-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `hdr_histogram_17` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 76.1 KiB | [hdr_histogram_17-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hdr_histogram_17-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `hdr_histogram_17` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 72.1 KiB | [hdr_histogram_17-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hdr_histogram_17-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `hdr_histogram_17` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.7 KiB | [hdr_histogram_17-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hdr_histogram_17-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `hdr_histogram_17` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 72.7 KiB | [hdr_histogram_17-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hdr_histogram_17-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.5 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.8 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.7 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.9 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 72.6 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 68.5 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.2 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 64.8 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.7 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-17-hdr-histogram` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.2 KiB | [postgresql-17-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-17-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdr_histogram_16` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 79.3 KiB | [hdr_histogram_16-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hdr_histogram_16-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `hdr_histogram_16` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.4 KiB | [hdr_histogram_16-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hdr_histogram_16-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `hdr_histogram_16` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 76.1 KiB | [hdr_histogram_16-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hdr_histogram_16-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `hdr_histogram_16` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 72.1 KiB | [hdr_histogram_16-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hdr_histogram_16-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `hdr_histogram_16` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 76.7 KiB | [hdr_histogram_16-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hdr_histogram_16-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `hdr_histogram_16` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 72.6 KiB | [hdr_histogram_16-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hdr_histogram_16-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 70.5 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 66.8 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 71.7 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 67.9 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 72.6 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 68.6 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 68.2 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 65.6 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 68.7 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-16-hdr-histogram` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 65.2 KiB | [postgresql-16-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-16-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdr_histogram_15` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 79.7 KiB | [hdr_histogram_15-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hdr_histogram_15-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `hdr_histogram_15` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.9 KiB | [hdr_histogram_15-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hdr_histogram_15-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `hdr_histogram_15` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.9 KiB | [hdr_histogram_15-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hdr_histogram_15-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `hdr_histogram_15` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 79.1 KiB | [hdr_histogram_15-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hdr_histogram_15-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `hdr_histogram_15` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 83.6 KiB | [hdr_histogram_15-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hdr_histogram_15-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `hdr_histogram_15` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 79.9 KiB | [hdr_histogram_15-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hdr_histogram_15-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 71.1 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 71.1 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 72.3 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.5 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 78.6 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 75.1 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.3 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.4 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 75.0 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-15-hdr-histogram` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.7 KiB | [postgresql-15-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-15-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `hdr_histogram_14` | `1.0.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 79.7 KiB | [hdr_histogram_14-1.0.0-1PGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/hdr_histogram_14-1.0.0-1PGSTY.el8.x86_64.rpm) |
| `hdr_histogram_14` | `1.0.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 76.8 KiB | [hdr_histogram_14-1.0.0-1PGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/hdr_histogram_14-1.0.0-1PGSTY.el8.aarch64.rpm) |
| `hdr_histogram_14` | `1.0.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 82.9 KiB | [hdr_histogram_14-1.0.0-1PGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/hdr_histogram_14-1.0.0-1PGSTY.el9.x86_64.rpm) |
| `hdr_histogram_14` | `1.0.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 79.2 KiB | [hdr_histogram_14-1.0.0-1PGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/hdr_histogram_14-1.0.0-1PGSTY.el9.aarch64.rpm) |
| `hdr_histogram_14` | `1.0.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 83.6 KiB | [hdr_histogram_14-1.0.0-1PGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/hdr_histogram_14-1.0.0-1PGSTY.el10.x86_64.rpm) |
| `hdr_histogram_14` | `1.0.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 79.9 KiB | [hdr_histogram_14-1.0.0-1PGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/hdr_histogram_14-1.0.0-1PGSTY.el10.aarch64.rpm) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 71.0 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~bookworm_amd64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 67.5 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~bookworm_arm64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 72.2 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~trixie_amd64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 68.5 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~trixie_arm64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 78.5 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~jammy_amd64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 75.1 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~jammy_arm64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 74.3 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~noble_amd64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 71.3 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~noble_arm64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 75.0 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~resolute_amd64.deb) |
| `postgresql-14-hdr-histogram` | `1.0.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 71.6 KiB | [postgresql-14-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/h/hdr-histogram/postgresql-14-hdr-histogram_1.0.0-1PGSTY~resolute_arm64.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/vovcke/pg_hdr_histogram" title="Repository" icon="github" subtitle="github.com/vovcke/pg_hdr_histogram" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="hdr_histogram-1.0.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg hdr_histogram;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://pigsty.io/docs/pgsql/ext/install/) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install hdr_histogram;		# install via package name, for the active PG version

pig install hdr_histogram -v 18;   # install for PG 18
pig install hdr_histogram -v 17;   # install for PG 17
pig install hdr_histogram -v 16;   # install for PG 16
pig install hdr_histogram -v 15;   # install for PG 15
pig install hdr_histogram -v 14;   # install for PG 14

```


[**Create**](https://pigsty.io/docs/pgsql/ext/create/) this extension with:

```sql
CREATE EXTENSION hdr_histogram;
```

## Usage

Sources:

- [README.md](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/README.md)
- [hdr_histogram--1.0.0.sql](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram--1.0.0.sql)
- [hdr_histogram.c](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram.c)
- [META.json](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/META.json)
- [.github/workflows/ci.yml](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/.github/workflows/ci.yml)
- [hdr_histogram.control](https://github.com/vovcke/pg_hdr_histogram/blob/87686cfe3a06d48addd052e50f8e41a5bdc4347b/hdr_histogram.control)
- [Official 1.0.0 release archive](https://github.com/vovcke/pg_hdr_histogram/releases/download/v1.0.0/hdr_histogram-1.0.0.tar.gz)

`hdr_histogram` stores compressed HDR histograms and computes percentiles with configurable precision. `hdr_add` aggregates numeric samples; `hdr_merge` combines stored histograms. PostgreSQL 14–18 appear in upstream CI. The trusted, relocatable extension requires zlib and can be installed by a role with CREATE privilege on the database.

The official 1.0.0 release archive used for Pigsty packages installs control/SQL version 1.0; the Git tag's control file declares 1.0.0. Check the installed extension version rather than inferring it from the package version.

### Core Workflow

```sql
CREATE EXTENSION hdr_histogram;
SELECT hdr_percentile(hdr_add(1.0), 50);
SET hdr_histogram.merge_mode = 'lossless';
SELECT hdr_percentile(hdr_add(latency::double precision), 99)
FROM (VALUES (100), (200), (300), (500)) AS samples(latency);
CREATE TABLE latency_histograms (recorded_at timestamptz, hist hdr_histogram);
SELECT hdr_percentile(hdr_merge(hist), 99) FROM latency_histograms;
```

### Operational Boundaries

Samples are converted to int64, so choose one unit consistently. The default highest value is 1e9 with three significant figures; hdr_histogram.lowest, hdr_histogram.highest and hdr_histogram.sig_figs configure new histograms. Nonfinite and out-of-range values error. `hdr_percentile` accepts a percentile from 0 to 100; `hdr_min`, `hdr_max`, `hdr_mean`, `hdr_stddev`, `hdr_total_count`, `hdr_sig_figs`, `hdr_lowest` and `hdr_highest` inspect a histogram. The default merge_mode lossless rejects incompatible merges; strict requires identical parameters, left can discard values outside the accumulator range, and union expands range/precision. The first histogram function call loads the library; the example calls it before applying session GUCs. No shared preload is required for core use. Percentiles are bounded-error estimates, not exact order statistics.
