---
title: "jev"
linkTitle: "jev"
description: "Natural-language row filtering, ranking and classification through TypeSafe"
weight: 1900
categories: ["RAG"]
languages: ["Python"]
licenses: ["PostgreSQL"]
repos: ["PIGSTY"]
page_width: full
---

[**jev**](https://github.com/realZachi/pg-jev) : Natural-language row filtering, ranking and classification through TypeSafe


## Overview

|    ID    | Extension |  Package   | Version |        Category        |           License            |       Language       |
|:--------:|:---------:|:----------:|:-------:|:----------------------:|:----------------------------:|:--------------------:|
| **1900** | {{< badge content="jev" link="https://github.com/realZachi/pg-jev" >}} | {{< ext "jev" >}} | `0.2.0` | {{< category "RAG" >}} | {{< license "PostgreSQL" >}} | {{< language "Python" >}} |


|  Attribute | Has Binary | Has Library | Need Load | Has DDL | Relocatable | Trusted |
|:----------:|:----------:|:-----------:|:---------:|:-------:|:-----------:|:-------:|
| {{< badge content="----d--" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="No" color="blue" >}} | {{< badge content="Yes" color="green" >}} | {{< badge content="no" color="orange" >}} | {{< badge content="no" color="orange" >}} |


| **Relationships** |   |
|:-----------------:|:----|
|    **Schemas**    | `public` |
|   **Requires**    | {{< ext "plpython3u" >}} |

> [!Note] Requires plpython3u and API credentials; sends row contents to an external model service. PG14-17.


## Packages

| Type | Repo | Version | PG Major Compatibility | Package Pattern | Dependencies |
|:----:|:----:|:-------:|:---------------------:|:----------------|:------------:|
| **EXT** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "" "red" >}} {{< bg "17" "" "green" >}} {{< bg "16" "" "green" >}} {{< bg "15" "" "green" >}} {{< bg "14" "" "green" >}} | `jev` | `plpython3u` |
| **RPM** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "jev_18" "red" >}} {{< bg "17" "jev_17" "green" >}} {{< bg "16" "jev_16" "green" >}} {{< bg "15" "jev_15" "green" >}} {{< bg "14" "jev_14" "green" >}} | `jev_$v` | `postgresql$v-plpython3` |
| **DEB** | {{< badge content="PIGSTY" link="/repo/pgsql" >}} | `0.2.0` | {{< bg "18" "postgresql-18-jev" "red" >}} {{< bg "17" "postgresql-17-jev" "green" >}} {{< bg "16" "postgresql-16-jev" "green" >}} {{< bg "15" "postgresql-15-jev" "green" >}} {{< bg "14" "postgresql-14-jev" "green" >}} | `postgresql-$v-jev` | `postgresql-plpython3-$v` |
{.packages}


| **Linux** / **PG** |                  **PG18**                   |                  **PG17**                   |                  **PG16**                   |                  **PG15**                   |                  **PG14**                   |
|:------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|:-------------------------------------------:|
| {{< os "el8.x86_64" >}} | {{< bg "N/A" "jev_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "jev_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_14 : AVAIL 1" "green" >}} |
| {{< os "el8.aarch64" >}} | {{< bg "N/A" "jev_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "jev_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_14 : AVAIL 1" "green" >}} |
| {{< os "el9.x86_64" >}} | {{< bg "N/A" "jev_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "jev_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_14 : AVAIL 1" "green" >}} |
| {{< os "el9.aarch64" >}} | {{< bg "N/A" "jev_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "jev_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_14 : AVAIL 1" "green" >}} |
| {{< os "el10.x86_64" >}} | {{< bg "N/A" "jev_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "jev_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_14 : AVAIL 1" "green" >}} |
| {{< os "el10.aarch64" >}} | {{< bg "N/A" "jev_18 : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "jev_17 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_16 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_15 : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "jev_14 : AVAIL 1" "green" >}} |
| {{< os "d12.x86_64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "d12.aarch64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "d13.x86_64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "d13.aarch64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "u22.x86_64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "u22.aarch64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "u24.x86_64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "u24.aarch64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "u26.x86_64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
| {{< os "u26.aarch64" >}} | {{< bg "N/A" "postgresql-18-jev : N/A 0" "gray" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-17-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-16-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-15-jev : AVAIL 1" "green" >}} | {{< bg "PIGSTY 0.2.0" "postgresql-14-jev : AVAIL 1" "green" >}} |
{.matrix}


{{< tabs group="pgmajor" >}}
{{< tab label="PG17" value="pg17" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jev_17` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.3 KiB | [jev_17-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jev_17-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_17` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.3 KiB | [jev_17-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jev_17-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_17` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.0 KiB | [jev_17-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jev_17-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_17` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [jev_17-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jev_17-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_17` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.2 KiB | [jev_17-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jev_17-0.2.0-1PGSTY.el10.noarch.rpm) |
| `jev_17` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [jev_17-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jev_17-0.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-17-jev` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.6 KiB | [postgresql-17-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.6 KiB | [postgresql-17-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.6 KiB | [postgresql-17-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.6 KiB | [postgresql-17-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.9 KiB | [postgresql-17-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.9 KiB | [postgresql-17-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-17-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.8 KiB | [postgresql-17-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 17.8 KiB | [postgresql-17-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-17-jev` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.8 KiB | [postgresql-17-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-17-jev_0.2.0-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG16" value="pg16" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jev_16` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.3 KiB | [jev_16-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jev_16-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_16` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.3 KiB | [jev_16-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jev_16-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_16` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.0 KiB | [jev_16-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jev_16-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_16` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [jev_16-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jev_16-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_16` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.2 KiB | [jev_16-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jev_16-0.2.0-1PGSTY.el10.noarch.rpm) |
| `jev_16` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [jev_16-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jev_16-0.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-16-jev` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.6 KiB | [postgresql-16-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.6 KiB | [postgresql-16-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.6 KiB | [postgresql-16-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.6 KiB | [postgresql-16-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.9 KiB | [postgresql-16-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.9 KiB | [postgresql-16-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-16-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.8 KiB | [postgresql-16-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 17.8 KiB | [postgresql-16-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-16-jev` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.8 KiB | [postgresql-16-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-16-jev_0.2.0-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG15" value="pg15" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jev_15` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.3 KiB | [jev_15-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jev_15-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_15` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.3 KiB | [jev_15-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jev_15-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_15` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.0 KiB | [jev_15-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jev_15-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_15` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [jev_15-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jev_15-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_15` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.2 KiB | [jev_15-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jev_15-0.2.0-1PGSTY.el10.noarch.rpm) |
| `jev_15` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [jev_15-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jev_15-0.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-15-jev` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.6 KiB | [postgresql-15-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.6 KiB | [postgresql-15-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.6 KiB | [postgresql-15-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.6 KiB | [postgresql-15-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.9 KiB | [postgresql-15-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.9 KiB | [postgresql-15-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-15-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.8 KiB | [postgresql-15-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 17.8 KiB | [postgresql-15-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-15-jev` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.8 KiB | [postgresql-15-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-15-jev_0.2.0-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}
{{< tab label="PG14" value="pg14" >}}

| **Package** | **Version** | **OS** | **ORG** | **SIZE** | **File URL** |
|:------------|:-----------:|:------:|:-------:|:--------:|:--------------|
| `jev_14` | `0.2.0` | [el8.x86_64](/os/el8.x86_64) | pigsty | 23.3 KiB | [jev_14-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.x86_64/jev_14-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_14` | `0.2.0` | [el8.aarch64](/os/el8.aarch64) | pigsty | 23.3 KiB | [jev_14-0.2.0-1PGSTY.el8.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el8.aarch64/jev_14-0.2.0-1PGSTY.el8.noarch.rpm) |
| `jev_14` | `0.2.0` | [el9.x86_64](/os/el9.x86_64) | pigsty | 23.0 KiB | [jev_14-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.x86_64/jev_14-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_14` | `0.2.0` | [el9.aarch64](/os/el9.aarch64) | pigsty | 23.0 KiB | [jev_14-0.2.0-1PGSTY.el9.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el9.aarch64/jev_14-0.2.0-1PGSTY.el9.noarch.rpm) |
| `jev_14` | `0.2.0` | [el10.x86_64](/os/el10.x86_64) | pigsty | 23.2 KiB | [jev_14-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.x86_64/jev_14-0.2.0-1PGSTY.el10.noarch.rpm) |
| `jev_14` | `0.2.0` | [el10.aarch64](/os/el10.aarch64) | pigsty | 23.1 KiB | [jev_14-0.2.0-1PGSTY.el10.noarch.rpm](https://repo.pigsty.io/yum/pgsql/el10.aarch64/jev_14-0.2.0-1PGSTY.el10.noarch.rpm) |
| `postgresql-14-jev` | `0.2.0` | [d12.x86_64](/os/d12.x86_64) | pigsty | 19.6 KiB | [postgresql-14-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [d12.aarch64](/os/d12.aarch64) | pigsty | 19.6 KiB | [postgresql-14-jev_0.2.0-1PGSTY~bookworm_all.deb](https://repo.pigsty.io/apt/pgsql/bookworm/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~bookworm_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [d13.x86_64](/os/d13.x86_64) | pigsty | 19.6 KiB | [postgresql-14-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [d13.aarch64](/os/d13.aarch64) | pigsty | 19.6 KiB | [postgresql-14-jev_0.2.0-1PGSTY~trixie_all.deb](https://repo.pigsty.io/apt/pgsql/trixie/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~trixie_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [u22.x86_64](/os/u22.x86_64) | pigsty | 17.9 KiB | [postgresql-14-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [u22.aarch64](/os/u22.aarch64) | pigsty | 17.9 KiB | [postgresql-14-jev_0.2.0-1PGSTY~jammy_all.deb](https://repo.pigsty.io/apt/pgsql/jammy/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~jammy_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [u24.x86_64](/os/u24.x86_64) | pigsty | 17.8 KiB | [postgresql-14-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [u24.aarch64](/os/u24.aarch64) | pigsty | 17.8 KiB | [postgresql-14-jev_0.2.0-1PGSTY~noble_all.deb](https://repo.pigsty.io/apt/pgsql/noble/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~noble_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [u26.x86_64](/os/u26.x86_64) | pigsty | 17.8 KiB | [postgresql-14-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~resolute_all.deb) |
| `postgresql-14-jev` | `0.2.0` | [u26.aarch64](/os/u26.aarch64) | pigsty | 17.8 KiB | [postgresql-14-jev_0.2.0-1PGSTY~resolute_all.deb](https://repo.pigsty.io/apt/pgsql/resolute/pool/main/j/jev/postgresql-14-jev_0.2.0-1PGSTY~resolute_all.deb) |
{.downloads}

{{< /tab >}}{{< /tabs >}}

## Source

{{< cards cols=3 >}}
{{< card link="https://github.com/realZachi/pg-jev" title="Repository" icon="github" subtitle="github.com/realZachi/pg-jev" />}}
{{< card link="/list" title="Source Tarball" icon="clipboard-list" subtitle="jev-0.2.0.tar.gz" />}}
{{< /cards >}}


```bash
pig build pkg jev;		# build rpm/deb
```


## Install

Make sure [**PGDG**](/repo/pgdg) and [**PIGSTY**](/repo/pgsql) repo available:

```bash
pig repo add pgsql -u   # add both repo and update cache
```

[**Install**](https://ext.pgsty.com/usage/install) this extension with [**pig**](https://pig.pgsty.com):

```bash
pig install jev;		# install via package name, for the active PG version

pig install jev -v 17;   # install for PG 17
pig install jev -v 16;   # install for PG 16
pig install jev -v 15;   # install for PG 15
pig install jev -v 14;   # install for PG 14

```


[**Create**](https://ext.pgsty.com/usage/create) this extension with:

```sql
CREATE EXTENSION jev CASCADE; -- requires plpython3u
```

## Usage

Sources:

- [PGXN 0.2.0](https://pgxn.org/dist/jev/0.2.0/)

`jev` filters, ranks and classifies rows using natural-language conditions evaluated by a TypeSafe API. It requires `plpython3u`, superuser installation and an API key. The upstream documented PostgreSQL range is 14–17; no preload is needed.

### Query rows

```sql
CREATE EXTENSION jev CASCADE;
SET jev.api_key = 'your-key';
CREATE TABLE jev_demo (id integer, body text);
INSERT INTO jev_demo VALUES (1, 'The customer requests a refund');
SELECT id, jev_prob(jev_demo, 'the customer requests a refund')
FROM jev_demo;
SELECT jev_stats();
```

`jev()` returns a boolean predicate; `jev_prob()` returns a probability. `jev_choice()` classifies among options, and `jev_score()` evaluates ordered levels. Cache and connection pools are session-local; `jev_cache_clear()` clears session state.

### Service and data boundaries

Row contents are transmitted to the configured API. Set `jev.api_url` for the intended service and use `jev.max_rows_per_statement` and `jev.max_chars_per_statement` to cap work. API latency and charges depend on the service and data volume. The API key must be handled as a credential.

PL/Python runs with server operating-system privileges. Hosts that withhold superuser access or PL/Python cannot run this extension. Local package tests use the upstream mock API; they do not validate the remote model's judgment quality.
