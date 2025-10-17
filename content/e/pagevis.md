---
title: "pagevis"
linkTitle: "pagevis"
description: "Visualise database pages in ascii code"
weight: 6800
categories: ["Stat"]
width: full
---

Visualise database pages in ascii code

## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **6800** | {{< badge content="pagevis" link="https://github.com/hollobon/pagevis" >}} | {{< ext "pagevis" "pagevis" >}} | `0.1` | {{< category "STAT" >}} | {{< license "MIT" >}} | {{< language "SQL" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="-----d-r" color="blue" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="No" color="green" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="yes" color="green" >}} | {{< badge content="no" color="red" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|   **See Also**    | {{< ext "pageinspect" >}} {{< ext "pg_visibility" >}} {{< ext "amcheck" >}} {{< ext "pg_surgery" >}} {{< ext "pgstattuple" >}} {{< ext "pg_dirtyread" >}} {{< ext "toastinfo" >}} {{< ext "pg_profile" >}} |


## Packages

| Type | Repo | Version | PG Major Availability | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EL** | {{< badge content="PIGSTY" link="/e/pagevis" >}} | `0.1` | {{< badge content="18" color="red" alt="pagevis_18" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `pagevis_$v` | - |
| **Debian** | {{< badge content="PIGSTY" link="/e/pagevis" >}} | `0.1` | {{< badge content="18" color="red" alt="postgresql-18-pagevis" >}} {{< badge content="17" color="green" >}} {{< badge content="16" color="green" >}} {{< badge content="15" color="green" >}} {{< badge content="14" color="green" >}} | `postgresql-$v-pagevis` | - |


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
|    `el8.x86_64`    |    {{< pkg "pagevis_18" >}}     | {{< pkg "pagevis_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_17-0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pagevis_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_16-0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pagevis_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_15-0.1-1PIGSTY.el8.x86_64.rpm" >}} | {{< pkg "pagevis_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_14-0.1-1PIGSTY.el8.x86_64.rpm" >}} |
|    `el8.aarch64`    |    {{< pkg "pagevis_18" >}}     | {{< pkg "pagevis_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_17-0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pagevis_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_16-0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pagevis_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_15-0.1-1PIGSTY.el8.aarch64.rpm" >}} | {{< pkg "pagevis_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_14-0.1-1PIGSTY.el8.aarch64.rpm" >}} |
|    `el9.x86_64`    |    {{< pkg "pagevis_18" >}}     | {{< pkg "pagevis_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_17-0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pagevis_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_16-0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pagevis_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_15-0.1-1PIGSTY.el9.x86_64.rpm" >}} | {{< pkg "pagevis_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_14-0.1-1PIGSTY.el9.x86_64.rpm" >}} |
|    `el9.aarch64`    |    {{< pkg "pagevis_18" >}}     | {{< pkg "pagevis_17" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_17-0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pagevis_16" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_16-0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pagevis_15" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_15-0.1-1PIGSTY.el9.aarch64.rpm" >}} | {{< pkg "pagevis_14" "0.1" "pigsty" "https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_14-0.1-1PIGSTY.el9.aarch64.rpm" >}} |
|    `d12.x86_64`    |    {{< pkg "postgresql-18-pagevis" >}}     | {{< pkg "postgresql-17-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-16-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-15-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~bookworm_amd64.deb" >}} | {{< pkg "postgresql-14-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~bookworm_amd64.deb" >}} |
|    `d12.aarch64`    |    {{< pkg "postgresql-18-pagevis" >}}     | {{< pkg "postgresql-17-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-16-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-15-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~bookworm_arm64.deb" >}} | {{< pkg "postgresql-14-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~bookworm_arm64.deb" >}} |
|    `u22.x86_64`    |    {{< pkg "postgresql-18-pagevis" >}}     | {{< pkg "postgresql-17-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-16-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-15-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~jammy_amd64.deb" >}} | {{< pkg "postgresql-14-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~jammy_amd64.deb" >}} |
|    `u22.aarch64`    |    {{< pkg "postgresql-18-pagevis" >}}     | {{< pkg "postgresql-17-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-16-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-15-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~jammy_arm64.deb" >}} | {{< pkg "postgresql-14-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~jammy_arm64.deb" >}} |
|    `u24.x86_64`    |    {{< pkg "postgresql-18-pagevis" >}}     | {{< pkg "postgresql-17-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-16-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-15-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~noble_amd64.deb" >}} | {{< pkg "postgresql-14-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~noble_amd64.deb" >}} |
|    `u24.aarch64`    |    {{< pkg "postgresql-18-pagevis" >}}     | {{< pkg "postgresql-17-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-16-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-15-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~noble_arm64.deb" >}} | {{< pkg "postgresql-14-pagevis" "0.1" "pigsty" "https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~noble_arm64.deb" >}} |


{{< tabs items="PG17,PG16,PG15,PG14,PG13" >}}


{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pagevis_17` | 0.1 | `el8.x86_64` | pigsty | 8.1 KiB | [pagevis_17-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_17-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_17` | 0.1 | `el8.aarch64` | pigsty | 8.1 KiB | [pagevis_17-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_17-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_17` | 0.1 | `el9.aarch64` | pigsty | 8.1 KiB | [pagevis_17-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_17-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_17` | 0.1 | `el9.x86_64` | pigsty | 8.2 KiB | [pagevis_17-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_17-0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-17-pagevis` | 0.1 | `d12.x86_64` | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-17-pagevis` | 0.1 | `d12.aarch64` | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-17-pagevis` | 0.1 | `u22.x86_64` | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-17-pagevis` | 0.1 | `u22.aarch64` | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-17-pagevis` | 0.1 | `u24.x86_64` | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-17-pagevis` | 0.1 | `u24.aarch64` | pigsty | 5.2 KiB | [postgresql-17-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-17-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pagevis_16` | 0.1 | `el8.x86_64` | pigsty | 8.1 KiB | [pagevis_16-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_16-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_16` | 0.1 | `el8.aarch64` | pigsty | 8.1 KiB | [pagevis_16-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_16-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_16` | 0.1 | `el9.x86_64` | pigsty | 8.2 KiB | [pagevis_16-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_16-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_16` | 0.1 | `el9.aarch64` | pigsty | 8.1 KiB | [pagevis_16-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_16-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-16-pagevis` | 0.1 | `d12.x86_64` | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-16-pagevis` | 0.1 | `d12.aarch64` | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-16-pagevis` | 0.1 | `u22.aarch64` | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-16-pagevis` | 0.1 | `u22.x86_64` | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-16-pagevis` | 0.1 | `u24.x86_64` | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-16-pagevis` | 0.1 | `u24.aarch64` | pigsty | 5.2 KiB | [postgresql-16-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-16-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pagevis_15` | 0.1 | `el8.x86_64` | pigsty | 8.1 KiB | [pagevis_15-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_15-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_15` | 0.1 | `el8.aarch64` | pigsty | 8.1 KiB | [pagevis_15-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_15-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_15` | 0.1 | `el9.x86_64` | pigsty | 8.2 KiB | [pagevis_15-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_15-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_15` | 0.1 | `el9.aarch64` | pigsty | 8.1 KiB | [pagevis_15-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_15-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-15-pagevis` | 0.1 | `d12.aarch64` | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-15-pagevis` | 0.1 | `d12.x86_64` | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-15-pagevis` | 0.1 | `u22.aarch64` | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-15-pagevis` | 0.1 | `u22.x86_64` | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-15-pagevis` | 0.1 | `u24.x86_64` | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-15-pagevis` | 0.1 | `u24.aarch64` | pigsty | 5.2 KiB | [postgresql-15-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-15-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pagevis_14` | 0.1 | `el8.x86_64` | pigsty | 8.1 KiB | [pagevis_14-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_14-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_14` | 0.1 | `el8.aarch64` | pigsty | 8.1 KiB | [pagevis_14-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_14-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_14` | 0.1 | `el9.x86_64` | pigsty | 8.2 KiB | [pagevis_14-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_14-0.1-1PIGSTY.el9.x86_64.rpm) |
| `pagevis_14` | 0.1 | `el9.aarch64` | pigsty | 8.1 KiB | [pagevis_14-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_14-0.1-1PIGSTY.el9.aarch64.rpm) |
| `postgresql-14-pagevis` | 0.1 | `d12.x86_64` | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-14-pagevis` | 0.1 | `d12.aarch64` | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-14-pagevis` | 0.1 | `u22.x86_64` | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-14-pagevis` | 0.1 | `u22.aarch64` | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-14-pagevis` | 0.1 | `u24.x86_64` | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~noble_amd64.deb) |
| `postgresql-14-pagevis` | 0.1 | `u24.aarch64` | pigsty | 5.2 KiB | [postgresql-14-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-14-pagevis_0.1-1PIGSTY~noble_arm64.deb) |

{{< /tab >}}

{{< tab >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:-------------|
| `pagevis_13` | 0.1 | `el8.aarch64` | pigsty | 8.1 KiB | [pagevis_13-0.1-1PIGSTY.el8.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/pagevis_13-0.1-1PIGSTY.el8.aarch64.rpm) |
| `pagevis_13` | 0.1 | `el8.x86_64` | pigsty | 8.1 KiB | [pagevis_13-0.1-1PIGSTY.el8.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/pagevis_13-0.1-1PIGSTY.el8.x86_64.rpm) |
| `pagevis_13` | 0.1 | `el9.aarch64` | pigsty | 8.1 KiB | [pagevis_13-0.1-1PIGSTY.el9.aarch64.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/pagevis_13-0.1-1PIGSTY.el9.aarch64.rpm) |
| `pagevis_13` | 0.1 | `el9.x86_64` | pigsty | 8.2 KiB | [pagevis_13-0.1-1PIGSTY.el9.x86_64.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/pagevis_13-0.1-1PIGSTY.el9.x86_64.rpm) |
| `postgresql-13-pagevis` | 0.1 | `d12.aarch64` | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~bookworm_arm64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~bookworm_arm64.deb) |
| `postgresql-13-pagevis` | 0.1 | `d12.x86_64` | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~bookworm_amd64.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~bookworm_amd64.deb) |
| `postgresql-13-pagevis` | 0.1 | `u22.aarch64` | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~jammy_arm64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~jammy_arm64.deb) |
| `postgresql-13-pagevis` | 0.1 | `u22.x86_64` | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~jammy_amd64.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~jammy_amd64.deb) |
| `postgresql-13-pagevis` | 0.1 | `u24.aarch64` | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~noble_arm64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~noble_arm64.deb) |
| `postgresql-13-pagevis` | 0.1 | `u24.x86_64` | pigsty | 5.2 KiB | [postgresql-13-pagevis_0.1-1PIGSTY~noble_amd64.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/p/pagevis/postgresql-13-pagevis_0.1-1PIGSTY~noble_amd64.deb) |

{{< /tab >}}

{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/hollobon/pagevis" title="Repository" icon="github" subtitle="github.com/hollobon/pagevis" >}}
{{< card link="/list" icon="clipboard-list"  title="Source Tarball" subtitle="pagevis-0.1.tar.gz" >}}
{{< /cards >}}


```bash
pig build get pagevis; # get pagevis source code
pig build dep pagevis; # install build dependencies
pig build pkg pagevis; # build extension rpm or deb
pig build ext pagevis; # build extension rpms
```


## Install

To add the required PGDG / PIGSTY upstream repository, use:

```bash
pig repo add pgsql -u   # add PGDG + Pigsty repo and update cache (leave existing repos)
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with:

```bash
pig ext install pagevis; # install by extension name, for the current active PG version
pig ext install pagevis; # install via package alias, for the active PG version
pig ext install pagevis -v 18;   # install for PG 18
pig ext install pagevis -v 17;   # install for PG 17
pig ext install pagevis -v 16;   # install for PG 16
pig ext install pagevis -v 15;   # install for PG 15
pig ext install pagevis -v 14;   # install for PG 14
pig ext install pagevis -v 13;   # install for PG 13

```

[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```bash
CREATE EXTENSION pagevis;
```

