---
title: "pagevis"
linkTitle: "pagevis"
description: "Visualise database pages in ascii code"
weight: 6800
categories: ["STAT"]
width: full
---

[**pagevis**](https://github.com/hollobon/pagevis) : Visualise database pages in ascii code


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6800** | {{< badge content="pagevis" link="https://github.com/hollobon/pagevis" >}} | {{< ext "pagevis" >}} | `0.1` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d-r" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} {{< ext "pgstattuple" >}} {{< ext "pg_dirtyread" >}} {{< ext "toastinfo" >}} {{< ext "pg_profile" >}} |


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1` | {{< bg "18" "" "green" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} {{< bg "13" "" "green" >}} | `pagevis` | - |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1` | {{< bg "18" "pagevis_18" "green" >}} {{< bg "17" "pagevis_17" "green" >}} {{< bg "16" "pagevis_16" "green" >}} {{< bg "15" "pagevis_15" "green" >}} {{< bg "14" "pagevis_14" "green" >}} {{< bg "13" "pagevis_13" "green" >}} | `pagevis_$v` | - |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.1` | {{< bg "18" "postgresql-18-pagevis" "green" >}} {{< bg "17" "postgresql-17-pagevis" "green" >}} {{< bg "16" "postgresql-16-pagevis" "green" >}} {{< bg "15" "postgresql-15-pagevis" "green" >}} {{< bg "14" "postgresql-14-pagevis" "green" >}} {{< bg "13" "postgresql-13-pagevis" "green" >}} | `postgresql-$v-pagevis` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |                  **PG13**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "PIGSTY 0.1" "pagevis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_13 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "PIGSTY 0.1" "pagevis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_13 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "PIGSTY 0.1" "pagevis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_13 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "PIGSTY 0.1" "pagevis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_13 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "PIGSTY 0.1" "pagevis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_13 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "PIGSTY 0.1" "pagevis_18 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_14 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "pagevis_13 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "PIGSTY 0.1" "postgresql-18-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-17-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-16-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-15-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-14-pagevis : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.1" "postgresql-13-pagevis : AVAIL 1" "green" >}} |


{{< tabs items="PG18,PG17,PG16,PG15,PG14,PG13" >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pagevis_18` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 8.1 KiB | [pagevis_18-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_18-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_18` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.1 KiB | [pagevis_18-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_18-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_18` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 8.2 KiB | [pagevis_18-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_18-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_18` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 8.1 KiB | [pagevis_18-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_18-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_18` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 8.2 KiB | [pagevis_18-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pagevis_18-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pagevis_18` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 8.2 KiB | [pagevis_18-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pagevis_18-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-18-pagevis` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-18-pagevis` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-18-pagevis` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-18-pagevis` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-18-pagevis` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-18-pagevis` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-18-pagevis` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-18-pagevis` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-18-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-18-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pagevis_17` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 8.1 KiB | [pagevis_17-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_17-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_17` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.1 KiB | [pagevis_17-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_17-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_17` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 8.2 KiB | [pagevis_17-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_17-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_17` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 8.1 KiB | [pagevis_17-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_17-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_17` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 8.2 KiB | [pagevis_17-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pagevis_17-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pagevis_17` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 8.2 KiB | [pagevis_17-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pagevis_17-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-17-pagevis` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pagevis` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pagevis` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-17-pagevis` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-17-pagevis` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pagevis` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pagevis` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pagevis` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pagevis_16` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 8.1 KiB | [pagevis_16-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_16-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_16` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.1 KiB | [pagevis_16-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_16-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_16` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 8.2 KiB | [pagevis_16-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_16-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_16` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 8.1 KiB | [pagevis_16-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_16-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_16` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 8.2 KiB | [pagevis_16-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pagevis_16-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pagevis_16` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 8.2 KiB | [pagevis_16-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pagevis_16-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-16-pagevis` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pagevis` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pagevis` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-16-pagevis` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-16-pagevis` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pagevis` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pagevis` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pagevis` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pagevis_15` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 8.1 KiB | [pagevis_15-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_15-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_15` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.1 KiB | [pagevis_15-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_15-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_15` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 8.2 KiB | [pagevis_15-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_15-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_15` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 8.1 KiB | [pagevis_15-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_15-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_15` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 8.2 KiB | [pagevis_15-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pagevis_15-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pagevis_15` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 8.2 KiB | [pagevis_15-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pagevis_15-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-15-pagevis` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pagevis` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pagevis` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-15-pagevis` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-15-pagevis` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pagevis` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pagevis` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pagevis` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pagevis_14` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 8.1 KiB | [pagevis_14-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_14-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_14` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.1 KiB | [pagevis_14-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_14-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_14` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 8.2 KiB | [pagevis_14-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_14-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_14` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 8.1 KiB | [pagevis_14-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_14-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_14` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 8.2 KiB | [pagevis_14-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pagevis_14-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pagevis_14` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 8.2 KiB | [pagevis_14-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pagevis_14-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-14-pagevis` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pagevis` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pagevis` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-14-pagevis` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-14-pagevis` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pagevis` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pagevis` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pagevis` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}
{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `pagevis_13` | `0.1` | [el8.x86_64](/os/el8.x86_64) | pigsty | 8.1 KiB | [pagevis_13-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_13-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_13` | `0.1` | [el8.aarch64](/os/el8.aarch64) | pigsty | 8.1 KiB | [pagevis_13-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_13-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_13` | `0.1` | [el9.x86_64](/os/el9.x86_64) | pigsty | 8.2 KiB | [pagevis_13-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_13-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_13` | `0.1` | [el9.aarch64](/os/el9.aarch64) | pigsty | 8.1 KiB | [pagevis_13-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_13-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_13` | `0.1` | [el10.x86_64](/os/el10.x86_64) | pigsty | 8.2 KiB | [pagevis_13-0.1-1PIGSTY.el10.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/pagevis_13-0.1-1PIGSTY.el10.x86_64.rpm) |
| `pagevis_13` | `0.1` | [el10.aarch64](/os/el10.aarch64) | pigsty | 8.2 KiB | [pagevis_13-0.1-1PIGSTY.el10.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/pagevis_13-0.1-1PIGSTY.el10.aarch64.rpm) |
| `postgresql-13-pagevis` | `0.1` | [d12.x86_64](/os/d12.x86_64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pagevis` | `0.1` | [d12.aarch64](/os/d12.aarch64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pagevis` | `0.1` | [d13.x86_64](/os/d13.x86_64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~trixie_amd64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~trixie_amd64.deb) |
| `postgresql-13-pagevis` | `0.1` | [d13.aarch64](/os/d13.aarch64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~trixie_arm64.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~trixie_arm64.deb) |
| `postgresql-13-pagevis` | `0.1` | [u22.x86_64](/os/u22.x86_64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pagevis` | `0.1` | [u22.aarch64](/os/u22.aarch64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pagevis` | `0.1` | [u24.x86_64](/os/u24.x86_64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-13-pagevis` | `0.1` | [u24.aarch64](/os/u24.aarch64) | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hollobon/pagevis" title="Repository" icon="github" subtitle="github.com/hollobon/pagevis" >}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="pagevis-0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build pkg pagevis;		# build rpm / deb with pig
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgdg pigsty -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](/pig):

```bash
pig install pagevis;		# install via package name, for the active PG version

pig install pagevis -v 18;   # install for PG 18
pig install pagevis -v 17;   # install for PG 17
pig install pagevis -v 16;   # install for PG 16
pig install pagevis -v 15;   # install for PG 15
pig install pagevis -v 14;   # install for PG 14
pig install pagevis -v 13;   # install for PG 13

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION pagevis;
```
